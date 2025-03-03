/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/bookingcom/shipper/pkg/apis/shipper/v1alpha1"
	scheme "github.com/bookingcom/shipper/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TrafficTargetsGetter has a method to return a TrafficTargetInterface.
// A group's client should implement this interface.
type TrafficTargetsGetter interface {
	TrafficTargets(namespace string) TrafficTargetInterface
}

// TrafficTargetInterface has methods to work with TrafficTarget resources.
type TrafficTargetInterface interface {
	Create(*v1alpha1.TrafficTarget) (*v1alpha1.TrafficTarget, error)
	Update(*v1alpha1.TrafficTarget) (*v1alpha1.TrafficTarget, error)
	UpdateStatus(*v1alpha1.TrafficTarget) (*v1alpha1.TrafficTarget, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.TrafficTarget, error)
	List(opts v1.ListOptions) (*v1alpha1.TrafficTargetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TrafficTarget, err error)
	TrafficTargetExpansion
}

// trafficTargets implements TrafficTargetInterface
type trafficTargets struct {
	client rest.Interface
	ns     string
}

// newTrafficTargets returns a TrafficTargets
func newTrafficTargets(c *ShipperV1alpha1Client, namespace string) *trafficTargets {
	return &trafficTargets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the trafficTarget, and returns the corresponding trafficTarget object, and an error if there is any.
func (c *trafficTargets) Get(name string, options v1.GetOptions) (result *v1alpha1.TrafficTarget, err error) {
	result = &v1alpha1.TrafficTarget{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("traffictargets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TrafficTargets that match those selectors.
func (c *trafficTargets) List(opts v1.ListOptions) (result *v1alpha1.TrafficTargetList, err error) {
	result = &v1alpha1.TrafficTargetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("traffictargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested trafficTargets.
func (c *trafficTargets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("traffictargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a trafficTarget and creates it.  Returns the server's representation of the trafficTarget, and an error, if there is any.
func (c *trafficTargets) Create(trafficTarget *v1alpha1.TrafficTarget) (result *v1alpha1.TrafficTarget, err error) {
	result = &v1alpha1.TrafficTarget{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("traffictargets").
		Body(trafficTarget).
		Do().
		Into(result)
	return
}

// Update takes the representation of a trafficTarget and updates it. Returns the server's representation of the trafficTarget, and an error, if there is any.
func (c *trafficTargets) Update(trafficTarget *v1alpha1.TrafficTarget) (result *v1alpha1.TrafficTarget, err error) {
	result = &v1alpha1.TrafficTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("traffictargets").
		Name(trafficTarget.Name).
		Body(trafficTarget).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *trafficTargets) UpdateStatus(trafficTarget *v1alpha1.TrafficTarget) (result *v1alpha1.TrafficTarget, err error) {
	result = &v1alpha1.TrafficTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("traffictargets").
		Name(trafficTarget.Name).
		SubResource("status").
		Body(trafficTarget).
		Do().
		Into(result)
	return
}

// Delete takes name of the trafficTarget and deletes it. Returns an error if one occurs.
func (c *trafficTargets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("traffictargets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *trafficTargets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("traffictargets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched trafficTarget.
func (c *trafficTargets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TrafficTarget, err error) {
	result = &v1alpha1.TrafficTarget{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("traffictargets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
