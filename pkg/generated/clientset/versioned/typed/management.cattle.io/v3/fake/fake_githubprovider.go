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

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGithubProviders implements GithubProviderInterface
type FakeGithubProviders struct {
	Fake *FakeManagementV3
}

var githubprovidersResource = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "githubproviders"}

var githubprovidersKind = schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GithubProvider"}

// Get takes name of the githubProvider, and returns the corresponding githubProvider object, and an error if there is any.
func (c *FakeGithubProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.GithubProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(githubprovidersResource, name), &v3.GithubProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GithubProvider), err
}

// List takes label and field selectors, and returns the list of GithubProviders that match those selectors.
func (c *FakeGithubProviders) List(ctx context.Context, opts v1.ListOptions) (result *v3.GithubProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(githubprovidersResource, githubprovidersKind, opts), &v3.GithubProviderList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.GithubProviderList{ListMeta: obj.(*v3.GithubProviderList).ListMeta}
	for _, item := range obj.(*v3.GithubProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested githubProviders.
func (c *FakeGithubProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(githubprovidersResource, opts))
}

// Create takes the representation of a githubProvider and creates it.  Returns the server's representation of the githubProvider, and an error, if there is any.
func (c *FakeGithubProviders) Create(ctx context.Context, githubProvider *v3.GithubProvider, opts v1.CreateOptions) (result *v3.GithubProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(githubprovidersResource, githubProvider), &v3.GithubProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GithubProvider), err
}

// Update takes the representation of a githubProvider and updates it. Returns the server's representation of the githubProvider, and an error, if there is any.
func (c *FakeGithubProviders) Update(ctx context.Context, githubProvider *v3.GithubProvider, opts v1.UpdateOptions) (result *v3.GithubProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(githubprovidersResource, githubProvider), &v3.GithubProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GithubProvider), err
}

// Delete takes name of the githubProvider and deletes it. Returns an error if one occurs.
func (c *FakeGithubProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(githubprovidersResource, name, opts), &v3.GithubProvider{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGithubProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(githubprovidersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.GithubProviderList{})
	return err
}

// Patch applies the patch and returns the patched githubProvider.
func (c *FakeGithubProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.GithubProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(githubprovidersResource, name, pt, data, subresources...), &v3.GithubProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GithubProvider), err
}
