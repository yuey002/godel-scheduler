/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package volumebinding

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"

	framework "github.com/kubewharf/godel-scheduler/pkg/framework/api"
	"github.com/kubewharf/godel-scheduler/pkg/scheduler/framework/handle"
	podutil "github.com/kubewharf/godel-scheduler/pkg/util/pod"
	"github.com/kubewharf/godel-scheduler/pkg/volume/scheduling"
)

// VolumeBinding is a plugin that binds pod volumes in scheduling.
type VolumeBinding struct {
	binder scheduling.BaseVolumeBinder
}

var _ framework.FilterPlugin = &VolumeBinding{}

// Name is the name of the plugin used in Registry and configurations.
const Name = "VolumeBinding"

// Name returns name of the plugin. It is used in logs, etc.
func (pl *VolumeBinding) Name() string {
	return Name
}

func podHasPVCs(pod *v1.Pod) bool {
	for _, vol := range pod.Spec.Volumes {
		if vol.PersistentVolumeClaim != nil {
			return true
		}
	}
	return false
}

// Filter invoked at the filter extension point.
// It evaluates if a pod can fit due to the volumes it requests,
// for both bound and unbound PVCs.
//
// For PVCs that are bound, then it checks that the corresponding PV's node affinity is
// satisfied by the given node.
//
// For PVCs that are unbound, it tries to find available PVs that can satisfy the PVC requirements
// and that the PV node affinity is satisfied by the given node.
//
// The predicate returns true if all bound PVCs have compatible PVs with the node, and if all unbound
// PVCs can be matched with an available and node-compatible PV.
func (pl *VolumeBinding) Filter(ctx context.Context, cs *framework.CycleState, pod *v1.Pod, nodeInfo framework.NodeInfo) *framework.Status {
	if nodeInfo == nil {
		return framework.NewStatus(framework.Error, "node not found")
	}
	// If pod does not request any PVC, we don't need to do anything.
	if !podHasPVCs(pod) {
		return nil
	}

	podLauncher, _ := podutil.GetPodLauncher(pod)
	reasons, err := pl.binder.FindPodVolumes(pod, nodeInfo.GetNodeName(), nodeInfo.GetNodeLabels(podLauncher))
	if err != nil {
		return framework.NewStatus(framework.Error, err.Error())
	}

	if len(reasons) > 0 {
		status := framework.NewStatus(framework.UnschedulableAndUnresolvable)
		for _, reason := range reasons {
			status.AppendReason(string(reason))
		}
		return status
	}
	return nil
}

// New initializes a new plugin with volume binder and returns it.
func New(_ runtime.Object, fh handle.PodFrameworkHandle) (framework.Plugin, error) {
	return &VolumeBinding{
		binder: scheduling.NewBaseVolumeBinder(
			fh.SharedInformerFactory().Storage().V1().CSINodes(),
			fh.SharedInformerFactory().Core().V1().PersistentVolumeClaims(),
			fh.SharedInformerFactory().Core().V1().PersistentVolumes(),
			fh.SharedInformerFactory().Storage().V1().StorageClasses(),
		),
	}, nil
}
