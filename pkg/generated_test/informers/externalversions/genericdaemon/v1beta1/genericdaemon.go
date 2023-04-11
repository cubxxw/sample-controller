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
	genericdaemonv1beta1 "k8s.io/sample-controller/pkg/apis/genericdaemon/v1beta1"
	versioned "k8s.io/sample-controller/pkg/generated_test/clientset/versioned"
	internalinterfaces "k8s.io/sample-controller/pkg/generated_test/informers/externalversions/internalinterfaces"
	v1beta1 "k8s.io/sample-controller/pkg/generated_test/listers/genericdaemon/v1beta1"
)

// GenericdaemonInformer provides access to a shared informer and lister for
// Genericdaemons.
type GenericdaemonInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.GenericdaemonLister
}

type genericdaemonInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGenericdaemonInformer constructs a new informer for Genericdaemon type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGenericdaemonInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGenericdaemonInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGenericdaemonInformer constructs a new informer for Genericdaemon type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGenericdaemonInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MydomainV1beta1().Genericdaemons(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MydomainV1beta1().Genericdaemons(namespace).Watch(context.TODO(), options)
			},
		},
		&genericdaemonv1beta1.Genericdaemon{},
		resyncPeriod,
		indexers,
	)
}

func (f *genericdaemonInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGenericdaemonInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *genericdaemonInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&genericdaemonv1beta1.Genericdaemon{}, f.defaultInformer)
}

func (f *genericdaemonInformer) Lister() v1beta1.GenericdaemonLister {
	return v1beta1.NewGenericdaemonLister(f.Informer().GetIndexer())
}
