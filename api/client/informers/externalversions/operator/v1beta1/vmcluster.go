/*


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
// Code generated by informer-gen-v0.32. DO NOT EDIT.

package v1beta1

import (
	context "context"
	time "time"

	internalinterfaces "github.com/VictoriaMetrics/operator/api/client/informers/externalversions/internalinterfaces"
	operatorv1beta1 "github.com/VictoriaMetrics/operator/api/client/listers/operator/v1beta1"
	versioned "github.com/VictoriaMetrics/operator/api/client/versioned"
	apioperatorv1beta1 "github.com/VictoriaMetrics/operator/api/operator/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VMClusterInformer provides access to a shared informer and lister for
// VMClusters.
type VMClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() operatorv1beta1.VMClusterLister
}

type vMClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVMClusterInformer constructs a new informer for VMCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVMClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVMClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVMClusterInformer constructs a new informer for VMCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVMClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1beta1().VMClusters(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1beta1().VMClusters(namespace).Watch(context.TODO(), options)
			},
		},
		&apioperatorv1beta1.VMCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *vMClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVMClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *vMClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apioperatorv1beta1.VMCluster{}, f.defaultInformer)
}

func (f *vMClusterInformer) Lister() operatorv1beta1.VMClusterLister {
	return operatorv1beta1.NewVMClusterLister(f.Informer().GetIndexer())
}
