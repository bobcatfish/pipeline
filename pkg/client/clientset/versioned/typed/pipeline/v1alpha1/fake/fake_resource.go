/*
Copyright 2018 The Knative Authors

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
package fake

import (
	v1alpha1 "github.com/knative/build-pipeline/pkg/apis/pipeline/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeResources implements ResourceInterface
type FakeResources struct {
	Fake *FakePipelineV1alpha1
	ns   string
}

var resourcesResource = schema.GroupVersionResource{Group: "pipeline.knative.dev", Version: "v1alpha1", Resource: "resources"}

var resourcesKind = schema.GroupVersionKind{Group: "pipeline.knative.dev", Version: "v1alpha1", Kind: "Resource"}

// Get takes name of the resource, and returns the corresponding resource object, and an error if there is any.
func (c *FakeResources) Get(name string, options v1.GetOptions) (result *v1alpha1.Resource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourcesResource, c.ns, name), &v1alpha1.Resource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Resource), err
}

// List takes label and field selectors, and returns the list of Resources that match those selectors.
func (c *FakeResources) List(opts v1.ListOptions) (result *v1alpha1.ResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourcesResource, resourcesKind, c.ns, opts), &v1alpha1.ResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResourceList{ListMeta: obj.(*v1alpha1.ResourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resources.
func (c *FakeResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourcesResource, c.ns, opts))

}

// Create takes the representation of a resource and creates it.  Returns the server's representation of the resource, and an error, if there is any.
func (c *FakeResources) Create(resource *v1alpha1.Resource) (result *v1alpha1.Resource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourcesResource, c.ns, resource), &v1alpha1.Resource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Resource), err
}

// Update takes the representation of a resource and updates it. Returns the server's representation of the resource, and an error, if there is any.
func (c *FakeResources) Update(resource *v1alpha1.Resource) (result *v1alpha1.Resource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourcesResource, c.ns, resource), &v1alpha1.Resource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Resource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResources) UpdateStatus(resource *v1alpha1.Resource) (*v1alpha1.Resource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resourcesResource, "status", c.ns, resource), &v1alpha1.Resource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Resource), err
}

// Delete takes name of the resource and deletes it. Returns an error if one occurs.
func (c *FakeResources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(resourcesResource, c.ns, name), &v1alpha1.Resource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResourceList{})
	return err
}

// Patch applies the patch and returns the patched resource.
func (c *FakeResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Resource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourcesResource, c.ns, name, data, subresources...), &v1alpha1.Resource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Resource), err
}
