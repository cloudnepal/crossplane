/*
Copyright 2021 The Crossplane Authors.

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

package v1

import (
	"context"
	"time"

	v1 "github.com/crossplane/crossplane/apis/apiextensions/v1"
	scheme "github.com/crossplane/crossplane/internal/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CompositionRevisionsGetter has a method to return a CompositionRevisionInterface.
// A group's client should implement this interface.
type CompositionRevisionsGetter interface {
	CompositionRevisions() CompositionRevisionInterface
}

// CompositionRevisionInterface has methods to work with CompositionRevision resources.
type CompositionRevisionInterface interface {
	Create(ctx context.Context, compositionRevision *v1.CompositionRevision, opts metav1.CreateOptions) (*v1.CompositionRevision, error)
	Update(ctx context.Context, compositionRevision *v1.CompositionRevision, opts metav1.UpdateOptions) (*v1.CompositionRevision, error)
	UpdateStatus(ctx context.Context, compositionRevision *v1.CompositionRevision, opts metav1.UpdateOptions) (*v1.CompositionRevision, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CompositionRevision, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CompositionRevisionList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CompositionRevision, err error)
	CompositionRevisionExpansion
}

// compositionRevisions implements CompositionRevisionInterface
type compositionRevisions struct {
	client rest.Interface
}

// newCompositionRevisions returns a CompositionRevisions
func newCompositionRevisions(c *ApiextensionsV1Client) *compositionRevisions {
	return &compositionRevisions{
		client: c.RESTClient(),
	}
}

// Get takes name of the compositionRevision, and returns the corresponding compositionRevision object, and an error if there is any.
func (c *compositionRevisions) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CompositionRevision, err error) {
	result = &v1.CompositionRevision{}
	err = c.client.Get().
		Resource("compositionrevisions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CompositionRevisions that match those selectors.
func (c *compositionRevisions) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CompositionRevisionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CompositionRevisionList{}
	err = c.client.Get().
		Resource("compositionrevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested compositionRevisions.
func (c *compositionRevisions) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("compositionrevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a compositionRevision and creates it.  Returns the server's representation of the compositionRevision, and an error, if there is any.
func (c *compositionRevisions) Create(ctx context.Context, compositionRevision *v1.CompositionRevision, opts metav1.CreateOptions) (result *v1.CompositionRevision, err error) {
	result = &v1.CompositionRevision{}
	err = c.client.Post().
		Resource("compositionrevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(compositionRevision).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a compositionRevision and updates it. Returns the server's representation of the compositionRevision, and an error, if there is any.
func (c *compositionRevisions) Update(ctx context.Context, compositionRevision *v1.CompositionRevision, opts metav1.UpdateOptions) (result *v1.CompositionRevision, err error) {
	result = &v1.CompositionRevision{}
	err = c.client.Put().
		Resource("compositionrevisions").
		Name(compositionRevision.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(compositionRevision).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *compositionRevisions) UpdateStatus(ctx context.Context, compositionRevision *v1.CompositionRevision, opts metav1.UpdateOptions) (result *v1.CompositionRevision, err error) {
	result = &v1.CompositionRevision{}
	err = c.client.Put().
		Resource("compositionrevisions").
		Name(compositionRevision.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(compositionRevision).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the compositionRevision and deletes it. Returns an error if one occurs.
func (c *compositionRevisions) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("compositionrevisions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *compositionRevisions) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("compositionrevisions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched compositionRevision.
func (c *compositionRevisions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CompositionRevision, err error) {
	result = &v1.CompositionRevision{}
	err = c.client.Patch(pt).
		Resource("compositionrevisions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
