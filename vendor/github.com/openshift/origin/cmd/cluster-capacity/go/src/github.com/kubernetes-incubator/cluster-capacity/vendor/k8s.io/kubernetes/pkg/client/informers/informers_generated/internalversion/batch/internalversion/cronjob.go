/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package internalversion

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	batch "k8s.io/kubernetes/pkg/apis/batch"
	internalclientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/internalinterfaces"
	internalversion "k8s.io/kubernetes/pkg/client/listers/batch/internalversion"
	time "time"
)

// CronJobInformer provides access to a shared informer and lister for
// CronJobs.
type CronJobInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.CronJobLister
}

type cronJobInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCronJobInformer constructs a new informer for CronJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCronJobInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCronJobInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCronJobInformer constructs a new informer for CronJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCronJobInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Batch().CronJobs(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Batch().CronJobs(namespace).Watch(options)
			},
		},
		&batch.CronJob{},
		resyncPeriod,
		indexers,
	)
}

func (f *cronJobInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCronJobInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cronJobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&batch.CronJob{}, f.defaultInformer)
}

func (f *cronJobInformer) Lister() internalversion.CronJobLister {
	return internalversion.NewCronJobLister(f.Informer().GetIndexer())
}