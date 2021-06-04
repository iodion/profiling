package metric

import (
	"gitlab.bouncex.net/identity/id-resolution/pkg/metric"
)

const (
	// idr stands for id resolution
	namespace = "idr"
)

var (
	// LookupBySoftID metric tracks lookup api endpoint
	LookupBySoftID = metric.NewHistogramVecWithLabel(namespace, "lookup_by_softid", []string{"verb", "status"}, map[string]string{"path": "/lookup", "unit": "s"})

	// BulkLookupByBXDID collects metrics on bulk lookup by bxdid
	BulkLookupByBXDID = metric.NewHistogramVecWithLabel(namespace, "bulk_Lookup_by_bxdid", []string{"verb", "status"}, map[string]string{"path": "/client/bulk/lookup/bxdid", "unit": "s"})

	// BulkLookupRelatedBXDIDsByBXDIDThroughSoftID collects metrics on bulk lookup of related bxdids by bxdid
	BulkLookupRelatedBXDIDsByBXDIDThroughSoftID = metric.NewHistogramVecWithLabel(namespace, "bulk_Lookup_related_bxdid_by_bxdid", []string{"verb", "status"}, map[string]string{"path": "/client/bulk/lookup/bxdid/related", "unit": "s"})

	// BulkLookupBySoftID collects metrics on bulk lookup by soft id
	BulkLookupBySoftID = metric.NewHistogramVecWithLabel(namespace, "bulk_lookup_by_softid", []string{"verb", "status"}, map[string]string{"path": "/client/bulk/lookup/softid", "unit": "s"})

	// FindBXDIDBySoftID collects metrics user on repository
	FindBXDIDBySoftID = metric.NewHistogramWithLabel(namespace, "find_bxdid_by_softid", map[string]string{"method": "FindBXDIDBySoftID", "unit": "s"})

	// FindEmailHashByBXDID collects metrics on user repository
	FindEmailHashByBXDID = metric.NewHistogramWithLabel(namespace, "find_email_hash_by_bxdid", map[string]string{"method": "FindEmailHashByBXDID", "unit": "s"})

	// FetchBXDIDBySoftID collects metrics on user query
	FetchBXDIDBySoftID = metric.NewHistogramWithLabel(namespace, "fetch_bxdid_by_softid", map[string]string{"method": "FetchBXDIDBySoftID", "unit": "s"})

	// FetchEmailHashByBXDID collects metrics on user query
	FetchEmailHashByBXDID = metric.NewHistogramWithLabel(namespace, "fetch_email_hash_by_bxdid", map[string]string{"method": "FetchEmailHashByBXDID", "unit": "s"})

	GraphRequestProcessingLatency = metric.NewHistogramVecWithLabel(namespace, "graph_req_processing_latency", []string{"verb", "status", "bxwid", "isValid"}, map[string]string{"path": "/graph", "unit": "s", "help": "amount of time to complete background processing of a request"})

	Request = metric.NewHistogramVecWithLabel(namespace, "request", []string{"path", "verb", "status"}, map[string]string{"unit": "s"})
)
