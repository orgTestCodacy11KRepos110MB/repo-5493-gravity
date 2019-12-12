/*
Copyright 2019 Gravitational, Inc.

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
	v1beta1 "github.com/gravitational/gravity/lib/apis/lens/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeImageSets implements ImageSetInterface
type FakeImageSets struct {
	Fake *FakeLensV1beta1
	ns   string
}

var imagesetsResource = schema.GroupVersionResource{Group: "lens.gravitational.io", Version: "v1beta1", Resource: "imagesets"}

var imagesetsKind = schema.GroupVersionKind{Group: "lens.gravitational.io", Version: "v1beta1", Kind: "ImageSet"}

// Get takes name of the imageSet, and returns the corresponding imageSet object, and an error if there is any.
func (c *FakeImageSets) Get(name string, options v1.GetOptions) (result *v1beta1.ImageSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(imagesetsResource, c.ns, name), &v1beta1.ImageSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ImageSet), err
}

// List takes label and field selectors, and returns the list of ImageSets that match those selectors.
func (c *FakeImageSets) List(opts v1.ListOptions) (result *v1beta1.ImageSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(imagesetsResource, imagesetsKind, c.ns, opts), &v1beta1.ImageSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ImageSetList{ListMeta: obj.(*v1beta1.ImageSetList).ListMeta}
	for _, item := range obj.(*v1beta1.ImageSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested imageSets.
func (c *FakeImageSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(imagesetsResource, c.ns, opts))

}

// Create takes the representation of a imageSet and creates it.  Returns the server's representation of the imageSet, and an error, if there is any.
func (c *FakeImageSets) Create(imageSet *v1beta1.ImageSet) (result *v1beta1.ImageSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(imagesetsResource, c.ns, imageSet), &v1beta1.ImageSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ImageSet), err
}

// Update takes the representation of a imageSet and updates it. Returns the server's representation of the imageSet, and an error, if there is any.
func (c *FakeImageSets) Update(imageSet *v1beta1.ImageSet) (result *v1beta1.ImageSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(imagesetsResource, c.ns, imageSet), &v1beta1.ImageSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ImageSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeImageSets) UpdateStatus(imageSet *v1beta1.ImageSet) (*v1beta1.ImageSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(imagesetsResource, "status", c.ns, imageSet), &v1beta1.ImageSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ImageSet), err
}

// Delete takes name of the imageSet and deletes it. Returns an error if one occurs.
func (c *FakeImageSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(imagesetsResource, c.ns, name), &v1beta1.ImageSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeImageSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(imagesetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.ImageSetList{})
	return err
}

// Patch applies the patch and returns the patched imageSet.
func (c *FakeImageSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ImageSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(imagesetsResource, c.ns, name, pt, data, subresources...), &v1beta1.ImageSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ImageSet), err
}
