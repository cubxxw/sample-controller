/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	nsdddcontrollerv1beta1 "k8s.io/sample-controller/pkg/apis/nsdddcontroller/v1beta1"
	versioned "k8s.io/sample-controller/pkg/generated_blog/clientset/versioned"
	internalinterfaces "k8s.io/sample-controller/pkg/generated_blog/informers/externalversions/internalinterfaces"
	v1beta1 "k8s.io/sample-controller/pkg/generated_blog/listers/nsdddcontroller/v1beta1"
)

// BlogInformer provides access to a shared informer and lister for
// Blogs.
type BlogInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.BlogLister
}

type blogInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBlogInformer constructs a new informer for Blog type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBlogInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBlogInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBlogInformer constructs a new informer for Blog type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBlogInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ControllerV1beta1().Blogs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ControllerV1beta1().Blogs(namespace).Watch(context.TODO(), options)
			},
		},
		&nsdddcontrollerv1beta1.Blog{},
		resyncPeriod,
		indexers,
	)
}

func (f *blogInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBlogInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *blogInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&nsdddcontrollerv1beta1.Blog{}, f.defaultInformer)
}

func (f *blogInformer) Lister() v1beta1.BlogLister {
	return v1beta1.NewBlogLister(f.Informer().GetIndexer())
}
