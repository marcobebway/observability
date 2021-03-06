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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/knative/observability/pkg/apis/sink/v1alpha1"
	scheme "github.com/knative/observability/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MetricSinksGetter has a method to return a MetricSinkInterface.
// A group's client should implement this interface.
type MetricSinksGetter interface {
	MetricSinks(namespace string) MetricSinkInterface
}

// MetricSinkInterface has methods to work with MetricSink resources.
type MetricSinkInterface interface {
	Create(*v1alpha1.MetricSink) (*v1alpha1.MetricSink, error)
	Update(*v1alpha1.MetricSink) (*v1alpha1.MetricSink, error)
	UpdateStatus(*v1alpha1.MetricSink) (*v1alpha1.MetricSink, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MetricSink, error)
	List(opts v1.ListOptions) (*v1alpha1.MetricSinkList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MetricSink, err error)
	MetricSinkExpansion
}

// metricSinks implements MetricSinkInterface
type metricSinks struct {
	client rest.Interface
	ns     string
}

// newMetricSinks returns a MetricSinks
func newMetricSinks(c *ObservabilityV1alpha1Client, namespace string) *metricSinks {
	return &metricSinks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the metricSink, and returns the corresponding metricSink object, and an error if there is any.
func (c *metricSinks) Get(name string, options v1.GetOptions) (result *v1alpha1.MetricSink, err error) {
	result = &v1alpha1.MetricSink{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("metricsinks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MetricSinks that match those selectors.
func (c *metricSinks) List(opts v1.ListOptions) (result *v1alpha1.MetricSinkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MetricSinkList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("metricsinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested metricSinks.
func (c *metricSinks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("metricsinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a metricSink and creates it.  Returns the server's representation of the metricSink, and an error, if there is any.
func (c *metricSinks) Create(metricSink *v1alpha1.MetricSink) (result *v1alpha1.MetricSink, err error) {
	result = &v1alpha1.MetricSink{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("metricsinks").
		Body(metricSink).
		Do().
		Into(result)
	return
}

// Update takes the representation of a metricSink and updates it. Returns the server's representation of the metricSink, and an error, if there is any.
func (c *metricSinks) Update(metricSink *v1alpha1.MetricSink) (result *v1alpha1.MetricSink, err error) {
	result = &v1alpha1.MetricSink{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("metricsinks").
		Name(metricSink.Name).
		Body(metricSink).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *metricSinks) UpdateStatus(metricSink *v1alpha1.MetricSink) (result *v1alpha1.MetricSink, err error) {
	result = &v1alpha1.MetricSink{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("metricsinks").
		Name(metricSink.Name).
		SubResource("status").
		Body(metricSink).
		Do().
		Into(result)
	return
}

// Delete takes name of the metricSink and deletes it. Returns an error if one occurs.
func (c *metricSinks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("metricsinks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *metricSinks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("metricsinks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched metricSink.
func (c *metricSinks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MetricSink, err error) {
	result = &v1alpha1.MetricSink{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("metricsinks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
