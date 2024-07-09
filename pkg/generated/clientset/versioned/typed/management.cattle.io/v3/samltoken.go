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

package v3

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SamlTokensGetter has a method to return a SamlTokenInterface.
// A group's client should implement this interface.
type SamlTokensGetter interface {
	SamlTokens() SamlTokenInterface
}

// SamlTokenInterface has methods to work with SamlToken resources.
type SamlTokenInterface interface {
	Create(ctx context.Context, samlToken *v3.SamlToken, opts v1.CreateOptions) (*v3.SamlToken, error)
	Update(ctx context.Context, samlToken *v3.SamlToken, opts v1.UpdateOptions) (*v3.SamlToken, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.SamlToken, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.SamlTokenList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.SamlToken, err error)
	SamlTokenExpansion
}

// samlTokens implements SamlTokenInterface
type samlTokens struct {
	client rest.Interface
}

// newSamlTokens returns a SamlTokens
func newSamlTokens(c *ManagementV3Client) *samlTokens {
	return &samlTokens{
		client: c.RESTClient(),
	}
}

// Get takes name of the samlToken, and returns the corresponding samlToken object, and an error if there is any.
func (c *samlTokens) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.SamlToken, err error) {
	result = &v3.SamlToken{}
	err = c.client.Get().
		Resource("samltokens").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SamlTokens that match those selectors.
func (c *samlTokens) List(ctx context.Context, opts v1.ListOptions) (result *v3.SamlTokenList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.SamlTokenList{}
	err = c.client.Get().
		Resource("samltokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested samlTokens.
func (c *samlTokens) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("samltokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a samlToken and creates it.  Returns the server's representation of the samlToken, and an error, if there is any.
func (c *samlTokens) Create(ctx context.Context, samlToken *v3.SamlToken, opts v1.CreateOptions) (result *v3.SamlToken, err error) {
	result = &v3.SamlToken{}
	err = c.client.Post().
		Resource("samltokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(samlToken).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a samlToken and updates it. Returns the server's representation of the samlToken, and an error, if there is any.
func (c *samlTokens) Update(ctx context.Context, samlToken *v3.SamlToken, opts v1.UpdateOptions) (result *v3.SamlToken, err error) {
	result = &v3.SamlToken{}
	err = c.client.Put().
		Resource("samltokens").
		Name(samlToken.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(samlToken).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the samlToken and deletes it. Returns an error if one occurs.
func (c *samlTokens) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("samltokens").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *samlTokens) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("samltokens").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched samlToken.
func (c *samlTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.SamlToken, err error) {
	result = &v3.SamlToken{}
	err = c.client.Patch(pt).
		Resource("samltokens").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
