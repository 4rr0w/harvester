/*
Copyright 2024 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1beta1 "github.com/kube-logging/logging-operator/pkg/sdk/logging/api/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LoggingsGetter has a method to return a LoggingInterface.
// A group's client should implement this interface.
type LoggingsGetter interface {
	Loggings() LoggingInterface
}

// LoggingInterface has methods to work with Logging resources.
type LoggingInterface interface {
	Create(ctx context.Context, logging *v1beta1.Logging, opts v1.CreateOptions) (*v1beta1.Logging, error)
	Update(ctx context.Context, logging *v1beta1.Logging, opts v1.UpdateOptions) (*v1beta1.Logging, error)
	UpdateStatus(ctx context.Context, logging *v1beta1.Logging, opts v1.UpdateOptions) (*v1beta1.Logging, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Logging, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.LoggingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Logging, err error)
	LoggingExpansion
}

// loggings implements LoggingInterface
type loggings struct {
	client rest.Interface
}

// newLoggings returns a Loggings
func newLoggings(c *LoggingV1beta1Client) *loggings {
	return &loggings{
		client: c.RESTClient(),
	}
}

// Get takes name of the logging, and returns the corresponding logging object, and an error if there is any.
func (c *loggings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Logging, err error) {
	result = &v1beta1.Logging{}
	err = c.client.Get().
		Resource("loggings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Loggings that match those selectors.
func (c *loggings) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.LoggingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.LoggingList{}
	err = c.client.Get().
		Resource("loggings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested loggings.
func (c *loggings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("loggings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a logging and creates it.  Returns the server's representation of the logging, and an error, if there is any.
func (c *loggings) Create(ctx context.Context, logging *v1beta1.Logging, opts v1.CreateOptions) (result *v1beta1.Logging, err error) {
	result = &v1beta1.Logging{}
	err = c.client.Post().
		Resource("loggings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(logging).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a logging and updates it. Returns the server's representation of the logging, and an error, if there is any.
func (c *loggings) Update(ctx context.Context, logging *v1beta1.Logging, opts v1.UpdateOptions) (result *v1beta1.Logging, err error) {
	result = &v1beta1.Logging{}
	err = c.client.Put().
		Resource("loggings").
		Name(logging.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(logging).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *loggings) UpdateStatus(ctx context.Context, logging *v1beta1.Logging, opts v1.UpdateOptions) (result *v1beta1.Logging, err error) {
	result = &v1beta1.Logging{}
	err = c.client.Put().
		Resource("loggings").
		Name(logging.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(logging).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the logging and deletes it. Returns an error if one occurs.
func (c *loggings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("loggings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *loggings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("loggings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched logging.
func (c *loggings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Logging, err error) {
	result = &v1beta1.Logging{}
	err = c.client.Patch(pt).
		Resource("loggings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
