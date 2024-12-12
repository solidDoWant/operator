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
// Code generated by lister-gen-v0.31. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/VictoriaMetrics/operator/api/operator/v1beta1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// VMClusterLister helps list VMClusters.
// All objects returned here must be treated as read-only.
type VMClusterLister interface {
	// List lists all VMClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.VMCluster, err error)
	// VMClusters returns an object that can list and get VMClusters.
	VMClusters(namespace string) VMClusterNamespaceLister
	VMClusterListerExpansion
}

// vMClusterLister implements the VMClusterLister interface.
type vMClusterLister struct {
	listers.ResourceIndexer[*v1beta1.VMCluster]
}

// NewVMClusterLister returns a new VMClusterLister.
func NewVMClusterLister(indexer cache.Indexer) VMClusterLister {
	return &vMClusterLister{listers.New[*v1beta1.VMCluster](indexer, v1beta1.Resource("vmcluster"))}
}

// VMClusters returns an object that can list and get VMClusters.
func (s *vMClusterLister) VMClusters(namespace string) VMClusterNamespaceLister {
	return vMClusterNamespaceLister{listers.NewNamespaced[*v1beta1.VMCluster](s.ResourceIndexer, namespace)}
}

// VMClusterNamespaceLister helps list and get VMClusters.
// All objects returned here must be treated as read-only.
type VMClusterNamespaceLister interface {
	// List lists all VMClusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.VMCluster, err error)
	// Get retrieves the VMCluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.VMCluster, error)
	VMClusterNamespaceListerExpansion
}

// vMClusterNamespaceLister implements the VMClusterNamespaceLister
// interface.
type vMClusterNamespaceLister struct {
	listers.ResourceIndexer[*v1beta1.VMCluster]
}
