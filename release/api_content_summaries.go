/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zest

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// ContentSummariesAPIService ContentSummariesAPI service
type ContentSummariesAPIService service

type ContentSummariesAPIContentOstreeSummariesListRequest struct {
	ctx context.Context
	ApiService *ContentSummariesAPIService
	pulpDomain string
	limit *int32
	offset *int32
	ordering *[]string
	orphanedFor *float32
	prnIn *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	pulpLabelSelect *string
	q *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ContentSummariesAPIContentOstreeSummariesListRequest) Limit(limit int32) ContentSummariesAPIContentOstreeSummariesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentSummariesAPIContentOstreeSummariesListRequest) Offset(offset int32) ContentSummariesAPIContentOstreeSummariesListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r ContentSummariesAPIContentOstreeSummariesListRequest) Ordering(ordering []string) ContentSummariesAPIContentOstreeSummariesListRequest {
	r.ordering = &ordering
	return r
}

// Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME.
func (r ContentSummariesAPIContentOstreeSummariesListRequest) OrphanedFor(orphanedFor float32) ContentSummariesAPIContentOstreeSummariesListRequest {
	r.orphanedFor = &orphanedFor
	return r
}

// Multiple values may be separated by commas.
func (r ContentSummariesAPIContentOstreeSummariesListRequest) PrnIn(prnIn []string) ContentSummariesAPIContentOstreeSummariesListRequest {
	r.prnIn = &prnIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentSummariesAPIContentOstreeSummariesListRequest) PulpHrefIn(pulpHrefIn []string) ContentSummariesAPIContentOstreeSummariesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentSummariesAPIContentOstreeSummariesListRequest) PulpIdIn(pulpIdIn []string) ContentSummariesAPIContentOstreeSummariesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter labels by search string
func (r ContentSummariesAPIContentOstreeSummariesListRequest) PulpLabelSelect(pulpLabelSelect string) ContentSummariesAPIContentOstreeSummariesListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Filter results by using NOT, AND and OR operations on other filters
func (r ContentSummariesAPIContentOstreeSummariesListRequest) Q(q string) ContentSummariesAPIContentOstreeSummariesListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentSummariesAPIContentOstreeSummariesListRequest) RepositoryVersion(repositoryVersion string) ContentSummariesAPIContentOstreeSummariesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentSummariesAPIContentOstreeSummariesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentSummariesAPIContentOstreeSummariesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentSummariesAPIContentOstreeSummariesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentSummariesAPIContentOstreeSummariesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentSummariesAPIContentOstreeSummariesListRequest) Fields(fields []string) ContentSummariesAPIContentOstreeSummariesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentSummariesAPIContentOstreeSummariesListRequest) ExcludeFields(excludeFields []string) ContentSummariesAPIContentOstreeSummariesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentSummariesAPIContentOstreeSummariesListRequest) Execute() (*PaginatedostreeOstreeSummaryResponseList, *http.Response, error) {
	return r.ApiService.ContentOstreeSummariesListExecute(r)
}

/*
ContentOstreeSummariesList List ostree summarys

A ViewSet class for OSTree repository summary files.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentSummariesAPIContentOstreeSummariesListRequest
*/
func (a *ContentSummariesAPIService) ContentOstreeSummariesList(ctx context.Context, pulpDomain string) ContentSummariesAPIContentOstreeSummariesListRequest {
	return ContentSummariesAPIContentOstreeSummariesListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedostreeOstreeSummaryResponseList
func (a *ContentSummariesAPIService) ContentOstreeSummariesListExecute(r ContentSummariesAPIContentOstreeSummariesListRequest) (*PaginatedostreeOstreeSummaryResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedostreeOstreeSummaryResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentSummariesAPIService.ContentOstreeSummariesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/ostree/summaries/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "form", "csv")
	}
	if r.orphanedFor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orphaned_for", r.orphanedFor, "form", "")
	}
	if r.prnIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prn__in", r.prnIn, "form", "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "form", "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "form", "csv")
	}
	if r.pulpLabelSelect != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_label_select", r.pulpLabelSelect, "form", "")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "form", "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "form", "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "form", "")
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
                               parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "form", "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
                               parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "form", "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ContentSummariesAPIContentOstreeSummariesReadRequest struct {
	ctx context.Context
	ApiService *ContentSummariesAPIService
	ostreeOstreeSummaryHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentSummariesAPIContentOstreeSummariesReadRequest) Fields(fields []string) ContentSummariesAPIContentOstreeSummariesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentSummariesAPIContentOstreeSummariesReadRequest) ExcludeFields(excludeFields []string) ContentSummariesAPIContentOstreeSummariesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentSummariesAPIContentOstreeSummariesReadRequest) Execute() (*OstreeOstreeSummaryResponse, *http.Response, error) {
	return r.ApiService.ContentOstreeSummariesReadExecute(r)
}

/*
ContentOstreeSummariesRead Inspect an ostree summary

A ViewSet class for OSTree repository summary files.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ostreeOstreeSummaryHref
 @return ContentSummariesAPIContentOstreeSummariesReadRequest
*/
func (a *ContentSummariesAPIService) ContentOstreeSummariesRead(ctx context.Context, ostreeOstreeSummaryHref string) ContentSummariesAPIContentOstreeSummariesReadRequest {
	return ContentSummariesAPIContentOstreeSummariesReadRequest{
		ApiService: a,
		ctx: ctx,
		ostreeOstreeSummaryHref: ostreeOstreeSummaryHref,
	}
}

// Execute executes the request
//  @return OstreeOstreeSummaryResponse
func (a *ContentSummariesAPIService) ContentOstreeSummariesReadExecute(r ContentSummariesAPIContentOstreeSummariesReadRequest) (*OstreeOstreeSummaryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OstreeOstreeSummaryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentSummariesAPIService.ContentOstreeSummariesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ostree_ostree_summary_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ostree_ostree_summary_href"+"}", url.PathEscape(parameterValueToString(r.ostreeOstreeSummaryHref, "ostreeOstreeSummaryHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
                               parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "form", "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
                               parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "form", "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
