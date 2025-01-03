/*
Copyright The Godel Scheduler Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kubewharf/godel-scheduler-api/pkg/apis/scheduling/v1alpha1"
	scheme "github.com/kubewharf/godel-scheduler-api/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ReservationsGetter has a method to return a ReservationInterface.
// A group's client should implement this interface.
type ReservationsGetter interface {
	Reservations(namespace string) ReservationInterface
}

// ReservationInterface has methods to work with Reservation resources.
type ReservationInterface interface {
	Create(ctx context.Context, reservation *v1alpha1.Reservation, opts v1.CreateOptions) (*v1alpha1.Reservation, error)
	Update(ctx context.Context, reservation *v1alpha1.Reservation, opts v1.UpdateOptions) (*v1alpha1.Reservation, error)
	UpdateStatus(ctx context.Context, reservation *v1alpha1.Reservation, opts v1.UpdateOptions) (*v1alpha1.Reservation, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Reservation, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ReservationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Reservation, err error)
	ReservationExpansion
}

// reservations implements ReservationInterface
type reservations struct {
	client rest.Interface
	ns     string
}

// newReservations returns a Reservations
func newReservations(c *SchedulingV1alpha1Client, namespace string) *reservations {
	return &reservations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the reservation, and returns the corresponding reservation object, and an error if there is any.
func (c *reservations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Reservation, err error) {
	result = &v1alpha1.Reservation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("reservations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Reservations that match those selectors.
func (c *reservations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ReservationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ReservationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("reservations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested reservations.
func (c *reservations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("reservations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a reservation and creates it.  Returns the server's representation of the reservation, and an error, if there is any.
func (c *reservations) Create(ctx context.Context, reservation *v1alpha1.Reservation, opts v1.CreateOptions) (result *v1alpha1.Reservation, err error) {
	result = &v1alpha1.Reservation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("reservations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(reservation).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a reservation and updates it. Returns the server's representation of the reservation, and an error, if there is any.
func (c *reservations) Update(ctx context.Context, reservation *v1alpha1.Reservation, opts v1.UpdateOptions) (result *v1alpha1.Reservation, err error) {
	result = &v1alpha1.Reservation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("reservations").
		Name(reservation.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(reservation).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *reservations) UpdateStatus(ctx context.Context, reservation *v1alpha1.Reservation, opts v1.UpdateOptions) (result *v1alpha1.Reservation, err error) {
	result = &v1alpha1.Reservation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("reservations").
		Name(reservation.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(reservation).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the reservation and deletes it. Returns an error if one occurs.
func (c *reservations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("reservations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *reservations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("reservations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched reservation.
func (c *reservations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Reservation, err error) {
	result = &v1alpha1.Reservation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("reservations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
