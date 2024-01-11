/*
Copyright 2024 The Kubernetes Authors

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

package fake

import (
	"context"

	v1alpha1 "github.com/k8snetworkplumbingwg/whereabouts/pkg/api/whereabouts.cni.cncf.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOverlappingRangeIPReservations implements OverlappingRangeIPReservationInterface
type FakeOverlappingRangeIPReservations struct {
	Fake *FakeWhereaboutsV1alpha1
	ns   string
}

var overlappingrangeipreservationsResource = schema.GroupVersionResource{Group: "whereabouts.cni.cncf.io", Version: "v1alpha1", Resource: "overlappingrangeipreservations"}

var overlappingrangeipreservationsKind = schema.GroupVersionKind{Group: "whereabouts.cni.cncf.io", Version: "v1alpha1", Kind: "OverlappingRangeIPReservation"}

// Get takes name of the overlappingRangeIPReservation, and returns the corresponding overlappingRangeIPReservation object, and an error if there is any.
func (c *FakeOverlappingRangeIPReservations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OverlappingRangeIPReservation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(overlappingrangeipreservationsResource, c.ns, name), &v1alpha1.OverlappingRangeIPReservation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OverlappingRangeIPReservation), err
}

// List takes label and field selectors, and returns the list of OverlappingRangeIPReservations that match those selectors.
func (c *FakeOverlappingRangeIPReservations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OverlappingRangeIPReservationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(overlappingrangeipreservationsResource, overlappingrangeipreservationsKind, c.ns, opts), &v1alpha1.OverlappingRangeIPReservationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OverlappingRangeIPReservationList{ListMeta: obj.(*v1alpha1.OverlappingRangeIPReservationList).ListMeta}
	for _, item := range obj.(*v1alpha1.OverlappingRangeIPReservationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested overlappingRangeIPReservations.
func (c *FakeOverlappingRangeIPReservations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(overlappingrangeipreservationsResource, c.ns, opts))

}

// Create takes the representation of a overlappingRangeIPReservation and creates it.  Returns the server's representation of the overlappingRangeIPReservation, and an error, if there is any.
func (c *FakeOverlappingRangeIPReservations) Create(ctx context.Context, overlappingRangeIPReservation *v1alpha1.OverlappingRangeIPReservation, opts v1.CreateOptions) (result *v1alpha1.OverlappingRangeIPReservation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(overlappingrangeipreservationsResource, c.ns, overlappingRangeIPReservation), &v1alpha1.OverlappingRangeIPReservation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OverlappingRangeIPReservation), err
}

// Update takes the representation of a overlappingRangeIPReservation and updates it. Returns the server's representation of the overlappingRangeIPReservation, and an error, if there is any.
func (c *FakeOverlappingRangeIPReservations) Update(ctx context.Context, overlappingRangeIPReservation *v1alpha1.OverlappingRangeIPReservation, opts v1.UpdateOptions) (result *v1alpha1.OverlappingRangeIPReservation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(overlappingrangeipreservationsResource, c.ns, overlappingRangeIPReservation), &v1alpha1.OverlappingRangeIPReservation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OverlappingRangeIPReservation), err
}

// Delete takes name of the overlappingRangeIPReservation and deletes it. Returns an error if one occurs.
func (c *FakeOverlappingRangeIPReservations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(overlappingrangeipreservationsResource, c.ns, name, opts), &v1alpha1.OverlappingRangeIPReservation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOverlappingRangeIPReservations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(overlappingrangeipreservationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OverlappingRangeIPReservationList{})
	return err
}

// Patch applies the patch and returns the patched overlappingRangeIPReservation.
func (c *FakeOverlappingRangeIPReservations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OverlappingRangeIPReservation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(overlappingrangeipreservationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.OverlappingRangeIPReservation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OverlappingRangeIPReservation), err
}
