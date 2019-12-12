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

package v1beta1

import (
	"time"

	v1beta1 "github.com/gravitational/gravity/lib/apis/lens/v1beta1"
	scheme "github.com/gravitational/gravity/lib/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ImageSetsGetter has a method to return a ImageSetInterface.
// A group's client should implement this interface.
type ImageSetsGetter interface {
	ImageSets(namespace string) ImageSetInterface
}

// ImageSetInterface has methods to work with ImageSet resources.
type ImageSetInterface interface {
	Create(*v1beta1.ImageSet) (*v1beta1.ImageSet, error)
	Update(*v1beta1.ImageSet) (*v1beta1.ImageSet, error)
	UpdateStatus(*v1beta1.ImageSet) (*v1beta1.ImageSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.ImageSet, error)
	List(opts v1.ListOptions) (*v1beta1.ImageSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ImageSet, err error)
	ImageSetExpansion
}

// imageSets implements ImageSetInterface
type imageSets struct {
	client rest.Interface
	ns     string
}

// newImageSets returns a ImageSets
func newImageSets(c *LensV1beta1Client, namespace string) *imageSets {
	return &imageSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the imageSet, and returns the corresponding imageSet object, and an error if there is any.
func (c *imageSets) Get(name string, options v1.GetOptions) (result *v1beta1.ImageSet, err error) {
	result = &v1beta1.ImageSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("imagesets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ImageSets that match those selectors.
func (c *imageSets) List(opts v1.ListOptions) (result *v1beta1.ImageSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ImageSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("imagesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested imageSets.
func (c *imageSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("imagesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a imageSet and creates it.  Returns the server's representation of the imageSet, and an error, if there is any.
func (c *imageSets) Create(imageSet *v1beta1.ImageSet) (result *v1beta1.ImageSet, err error) {
	result = &v1beta1.ImageSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("imagesets").
		Body(imageSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a imageSet and updates it. Returns the server's representation of the imageSet, and an error, if there is any.
func (c *imageSets) Update(imageSet *v1beta1.ImageSet) (result *v1beta1.ImageSet, err error) {
	result = &v1beta1.ImageSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("imagesets").
		Name(imageSet.Name).
		Body(imageSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *imageSets) UpdateStatus(imageSet *v1beta1.ImageSet) (result *v1beta1.ImageSet, err error) {
	result = &v1beta1.ImageSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("imagesets").
		Name(imageSet.Name).
		SubResource("status").
		Body(imageSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the imageSet and deletes it. Returns an error if one occurs.
func (c *imageSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("imagesets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *imageSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("imagesets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched imageSet.
func (c *imageSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ImageSet, err error) {
	result = &v1beta1.ImageSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("imagesets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
