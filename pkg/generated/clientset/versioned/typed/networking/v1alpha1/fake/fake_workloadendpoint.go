// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/baidubce/baiducloud-cce-cni-driver/pkg/apis/networking/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWorkloadEndpoints implements WorkloadEndpointInterface
type FakeWorkloadEndpoints struct {
	Fake *FakeCceV1alpha1
	ns   string
}

var workloadendpointsResource = schema.GroupVersionResource{Group: "cce.io", Version: "v1alpha1", Resource: "workloadendpoints"}

var workloadendpointsKind = schema.GroupVersionKind{Group: "cce.io", Version: "v1alpha1", Kind: "WorkloadEndpoint"}

// Get takes name of the workloadEndpoint, and returns the corresponding workloadEndpoint object, and an error if there is any.
func (c *FakeWorkloadEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.WorkloadEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(workloadendpointsResource, c.ns, name), &v1alpha1.WorkloadEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkloadEndpoint), err
}

// List takes label and field selectors, and returns the list of WorkloadEndpoints that match those selectors.
func (c *FakeWorkloadEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WorkloadEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(workloadendpointsResource, workloadendpointsKind, c.ns, opts), &v1alpha1.WorkloadEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WorkloadEndpointList{ListMeta: obj.(*v1alpha1.WorkloadEndpointList).ListMeta}
	for _, item := range obj.(*v1alpha1.WorkloadEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workloadEndpoints.
func (c *FakeWorkloadEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(workloadendpointsResource, c.ns, opts))

}

// Create takes the representation of a workloadEndpoint and creates it.  Returns the server's representation of the workloadEndpoint, and an error, if there is any.
func (c *FakeWorkloadEndpoints) Create(ctx context.Context, workloadEndpoint *v1alpha1.WorkloadEndpoint, opts v1.CreateOptions) (result *v1alpha1.WorkloadEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(workloadendpointsResource, c.ns, workloadEndpoint), &v1alpha1.WorkloadEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkloadEndpoint), err
}

// Update takes the representation of a workloadEndpoint and updates it. Returns the server's representation of the workloadEndpoint, and an error, if there is any.
func (c *FakeWorkloadEndpoints) Update(ctx context.Context, workloadEndpoint *v1alpha1.WorkloadEndpoint, opts v1.UpdateOptions) (result *v1alpha1.WorkloadEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(workloadendpointsResource, c.ns, workloadEndpoint), &v1alpha1.WorkloadEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkloadEndpoint), err
}

// Delete takes name of the workloadEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeWorkloadEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(workloadendpointsResource, c.ns, name), &v1alpha1.WorkloadEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkloadEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(workloadendpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.WorkloadEndpointList{})
	return err
}

// Patch applies the patch and returns the patched workloadEndpoint.
func (c *FakeWorkloadEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorkloadEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(workloadendpointsResource, c.ns, name, pt, data, subresources...), &v1alpha1.WorkloadEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkloadEndpoint), err
}
