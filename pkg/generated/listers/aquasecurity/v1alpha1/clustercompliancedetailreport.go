// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/khulnasoft-lab/starboard/pkg/apis/khulnasoft-lab/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterComplianceDetailReportLister helps list ClusterComplianceDetailReports.
// All objects returned here must be treated as read-only.
type ClusterComplianceDetailReportLister interface {
	// List lists all ClusterComplianceDetailReports in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterComplianceDetailReport, err error)
	// ClusterComplianceDetailReports returns an object that can list and get ClusterComplianceDetailReports.
	ClusterComplianceDetailReports(namespace string) ClusterComplianceDetailReportNamespaceLister
	ClusterComplianceDetailReportListerExpansion
}

// clusterComplianceDetailReportLister implements the ClusterComplianceDetailReportLister interface.
type clusterComplianceDetailReportLister struct {
	indexer cache.Indexer
}

// NewClusterComplianceDetailReportLister returns a new ClusterComplianceDetailReportLister.
func NewClusterComplianceDetailReportLister(indexer cache.Indexer) ClusterComplianceDetailReportLister {
	return &clusterComplianceDetailReportLister{indexer: indexer}
}

// List lists all ClusterComplianceDetailReports in the indexer.
func (s *clusterComplianceDetailReportLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterComplianceDetailReport, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterComplianceDetailReport))
	})
	return ret, err
}

// ClusterComplianceDetailReports returns an object that can list and get ClusterComplianceDetailReports.
func (s *clusterComplianceDetailReportLister) ClusterComplianceDetailReports(namespace string) ClusterComplianceDetailReportNamespaceLister {
	return clusterComplianceDetailReportNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterComplianceDetailReportNamespaceLister helps list and get ClusterComplianceDetailReports.
// All objects returned here must be treated as read-only.
type ClusterComplianceDetailReportNamespaceLister interface {
	// List lists all ClusterComplianceDetailReports in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterComplianceDetailReport, err error)
	// Get retrieves the ClusterComplianceDetailReport from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterComplianceDetailReport, error)
	ClusterComplianceDetailReportNamespaceListerExpansion
}

// clusterComplianceDetailReportNamespaceLister implements the ClusterComplianceDetailReportNamespaceLister
// interface.
type clusterComplianceDetailReportNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterComplianceDetailReports in the indexer for a given namespace.
func (s clusterComplianceDetailReportNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterComplianceDetailReport, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterComplianceDetailReport))
	})
	return ret, err
}

// Get retrieves the ClusterComplianceDetailReport from the indexer for a given namespace and name.
func (s clusterComplianceDetailReportNamespaceLister) Get(name string) (*v1alpha1.ClusterComplianceDetailReport, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clustercompliancedetailreport"), name)
	}
	return obj.(*v1alpha1.ClusterComplianceDetailReport), nil
}
