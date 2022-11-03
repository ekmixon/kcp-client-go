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

package v1beta1

import (
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	flowcontrolv1beta1 "k8s.io/api/flowcontrol/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	flowcontrolv1beta1listers "k8s.io/client-go/listers/flowcontrol/v1beta1"
	"k8s.io/client-go/tools/cache"
)

// FlowSchemaClusterLister can list FlowSchemas across all workspaces, or scope down to a FlowSchemaLister for one workspace.
type FlowSchemaClusterLister interface {
	List(selector labels.Selector) (ret []*flowcontrolv1beta1.FlowSchema, err error)
	Cluster(cluster logicalcluster.Name) flowcontrolv1beta1listers.FlowSchemaLister
}

type flowSchemaClusterLister struct {
	indexer cache.Indexer
}

// NewFlowSchemaClusterLister returns a new FlowSchemaClusterLister.
func NewFlowSchemaClusterLister(indexer cache.Indexer) *flowSchemaClusterLister {
	return &flowSchemaClusterLister{indexer: indexer}
}

// List lists all FlowSchemas in the indexer across all workspaces.
func (s *flowSchemaClusterLister) List(selector labels.Selector) (ret []*flowcontrolv1beta1.FlowSchema, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*flowcontrolv1beta1.FlowSchema))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get FlowSchemas.
func (s *flowSchemaClusterLister) Cluster(cluster logicalcluster.Name) flowcontrolv1beta1listers.FlowSchemaLister {
	return &flowSchemaLister{indexer: s.indexer, cluster: cluster}
}

// flowSchemaLister implements the flowcontrolv1beta1listers.FlowSchemaLister interface.
type flowSchemaLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all FlowSchemas in the indexer for a workspace.
func (s *flowSchemaLister) List(selector labels.Selector) (ret []*flowcontrolv1beta1.FlowSchema, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.cluster, selector, func(i interface{}) {
		ret = append(ret, i.(*flowcontrolv1beta1.FlowSchema))
	})
	return ret, err
}

// Get retrieves the FlowSchema from the indexer for a given workspace and name.
func (s *flowSchemaLister) Get(name string) (*flowcontrolv1beta1.FlowSchema, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(flowcontrolv1beta1.Resource("FlowSchema"), name)
	}
	return obj.(*flowcontrolv1beta1.FlowSchema), nil
}
