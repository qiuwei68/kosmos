// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/kosmos.io/kosmos/pkg/apis/kosmos/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeShadowDaemonSets implements ShadowDaemonSetInterface
type FakeShadowDaemonSets struct {
	Fake *FakeKosmosV1alpha1
	ns   string
}

var shadowdaemonsetsResource = schema.GroupVersionResource{Group: "kosmos.io", Version: "v1alpha1", Resource: "shadowdaemonsets"}

var shadowdaemonsetsKind = schema.GroupVersionKind{Group: "kosmos.io", Version: "v1alpha1", Kind: "ShadowDaemonSet"}

// Get takes name of the shadowDaemonSet, and returns the corresponding shadowDaemonSet object, and an error if there is any.
func (c *FakeShadowDaemonSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ShadowDaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(shadowdaemonsetsResource, c.ns, name), &v1alpha1.ShadowDaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShadowDaemonSet), err
}

// List takes label and field selectors, and returns the list of ShadowDaemonSets that match those selectors.
func (c *FakeShadowDaemonSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ShadowDaemonSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(shadowdaemonsetsResource, shadowdaemonsetsKind, c.ns, opts), &v1alpha1.ShadowDaemonSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ShadowDaemonSetList{ListMeta: obj.(*v1alpha1.ShadowDaemonSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.ShadowDaemonSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested shadowDaemonSets.
func (c *FakeShadowDaemonSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(shadowdaemonsetsResource, c.ns, opts))

}

// Create takes the representation of a shadowDaemonSet and creates it.  Returns the server's representation of the shadowDaemonSet, and an error, if there is any.
func (c *FakeShadowDaemonSets) Create(ctx context.Context, shadowDaemonSet *v1alpha1.ShadowDaemonSet, opts v1.CreateOptions) (result *v1alpha1.ShadowDaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(shadowdaemonsetsResource, c.ns, shadowDaemonSet), &v1alpha1.ShadowDaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShadowDaemonSet), err
}

// Update takes the representation of a shadowDaemonSet and updates it. Returns the server's representation of the shadowDaemonSet, and an error, if there is any.
func (c *FakeShadowDaemonSets) Update(ctx context.Context, shadowDaemonSet *v1alpha1.ShadowDaemonSet, opts v1.UpdateOptions) (result *v1alpha1.ShadowDaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(shadowdaemonsetsResource, c.ns, shadowDaemonSet), &v1alpha1.ShadowDaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShadowDaemonSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeShadowDaemonSets) UpdateStatus(ctx context.Context, shadowDaemonSet *v1alpha1.ShadowDaemonSet, opts v1.UpdateOptions) (*v1alpha1.ShadowDaemonSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(shadowdaemonsetsResource, "status", c.ns, shadowDaemonSet), &v1alpha1.ShadowDaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShadowDaemonSet), err
}

// Delete takes name of the shadowDaemonSet and deletes it. Returns an error if one occurs.
func (c *FakeShadowDaemonSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(shadowdaemonsetsResource, c.ns, name, opts), &v1alpha1.ShadowDaemonSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeShadowDaemonSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(shadowdaemonsetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ShadowDaemonSetList{})
	return err
}

// Patch applies the patch and returns the patched shadowDaemonSet.
func (c *FakeShadowDaemonSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ShadowDaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(shadowdaemonsetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ShadowDaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShadowDaemonSet), err
}
