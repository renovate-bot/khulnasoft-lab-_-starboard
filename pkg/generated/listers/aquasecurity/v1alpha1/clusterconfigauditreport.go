// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/khulnasoft-lab/starboard/pkg/apis/khulnasoft-lab/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterConfigAuditReportLister helps list ClusterConfigAuditReports.
// All objects returned here must be treated as read-only.
type ClusterConfigAuditReportLister interface {
	// List lists all ClusterConfigAuditReports in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterConfigAuditReport, err error)
	// Get retrieves the ClusterConfigAuditReport from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterConfigAuditReport, error)
	ClusterConfigAuditReportListerExpansion
}

// clusterConfigAuditReportLister implements the ClusterConfigAuditReportLister interface.
type clusterConfigAuditReportLister struct {
	indexer cache.Indexer
}

// NewClusterConfigAuditReportLister returns a new ClusterConfigAuditReportLister.
func NewClusterConfigAuditReportLister(indexer cache.Indexer) ClusterConfigAuditReportLister {
	return &clusterConfigAuditReportLister{indexer: indexer}
}

// List lists all ClusterConfigAuditReports in the indexer.
func (s *clusterConfigAuditReportLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterConfigAuditReport, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterConfigAuditReport))
	})
	return ret, err
}

// Get retrieves the ClusterConfigAuditReport from the index for a given name.
func (s *clusterConfigAuditReportLister) Get(name string) (*v1alpha1.ClusterConfigAuditReport, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterconfigauditreport"), name)
	}
	return obj.(*v1alpha1.ClusterConfigAuditReport), nil
}
