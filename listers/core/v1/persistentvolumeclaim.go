//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1

import (
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

// PersistentVolumeClaimClusterLister can list PersistentVolumeClaims across all workspaces, or scope down to a PersistentVolumeClaimLister for one workspace.
type PersistentVolumeClaimClusterLister interface {
	List(selector labels.Selector) (ret []*corev1.PersistentVolumeClaim, err error)
	Cluster(cluster logicalcluster.Name) corev1listers.PersistentVolumeClaimLister
}

type persistentVolumeClaimClusterLister struct {
	indexer cache.Indexer
}

// NewPersistentVolumeClaimClusterLister returns a new PersistentVolumeClaimClusterLister.
func NewPersistentVolumeClaimClusterLister(indexer cache.Indexer) *persistentVolumeClaimClusterLister {
	return &persistentVolumeClaimClusterLister{indexer: indexer}
}

// List lists all PersistentVolumeClaims in the indexer across all workspaces.
func (s *persistentVolumeClaimClusterLister) List(selector labels.Selector) (ret []*corev1.PersistentVolumeClaim, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*corev1.PersistentVolumeClaim))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get PersistentVolumeClaims.
func (s *persistentVolumeClaimClusterLister) Cluster(cluster logicalcluster.Name) corev1listers.PersistentVolumeClaimLister {
	return &persistentVolumeClaimLister{indexer: s.indexer, cluster: cluster}
}

// persistentVolumeClaimLister implements the corev1listers.PersistentVolumeClaimLister interface.
type persistentVolumeClaimLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all PersistentVolumeClaims in the indexer for a workspace.
func (s *persistentVolumeClaimLister) List(selector labels.Selector) (ret []*corev1.PersistentVolumeClaim, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.cluster, selector, func(i interface{}) {
		ret = append(ret, i.(*corev1.PersistentVolumeClaim))
	})
	return ret, err
}

// PersistentVolumeClaims returns an object that can list and get PersistentVolumeClaims in one namespace.
func (s *persistentVolumeClaimLister) PersistentVolumeClaims(namespace string) corev1listers.PersistentVolumeClaimNamespaceLister {
	return &persistentVolumeClaimNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// persistentVolumeClaimNamespaceLister implements the corev1listers.PersistentVolumeClaimNamespaceLister interface.
type persistentVolumeClaimNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all PersistentVolumeClaims in the indexer for a given workspace and namespace.
func (s *persistentVolumeClaimNamespaceLister) List(selector labels.Selector) (ret []*corev1.PersistentVolumeClaim, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.cluster, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*corev1.PersistentVolumeClaim))
	})
	return ret, err
}

// Get retrieves the PersistentVolumeClaim from the indexer for a given workspace, namespace and name.
func (s *persistentVolumeClaimNamespaceLister) Get(name string) (*corev1.PersistentVolumeClaim, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(corev1.Resource("PersistentVolumeClaim"), name)
	}
	return obj.(*corev1.PersistentVolumeClaim), nil
}
