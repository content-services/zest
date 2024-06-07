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


// ContentConfigsAPIService ContentConfigsAPI service
type ContentConfigsAPIService service

type ContentConfigsAPIContentOstreeConfigsListRequest struct {
	ctx context.Context
	ApiService *ContentConfigsAPIService
	pulpDomain string
	limit *int32
	offset *int32
	ordering *[]string
	orphanedFor *float32
	pulpHrefIn *[]string
	pulpIdIn *[]string
	q *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ContentConfigsAPIContentOstreeConfigsListRequest) Limit(limit int32) ContentConfigsAPIContentOstreeConfigsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentConfigsAPIContentOstreeConfigsListRequest) Offset(offset int32) ContentConfigsAPIContentOstreeConfigsListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r ContentConfigsAPIContentOstreeConfigsListRequest) Ordering(ordering []string) ContentConfigsAPIContentOstreeConfigsListRequest {
	r.ordering = &ordering
	return r
}

// Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME.
func (r ContentConfigsAPIContentOstreeConfigsListRequest) OrphanedFor(orphanedFor float32) ContentConfigsAPIContentOstreeConfigsListRequest {
	r.orphanedFor = &orphanedFor
	return r
}

// Multiple values may be separated by commas.
func (r ContentConfigsAPIContentOstreeConfigsListRequest) PulpHrefIn(pulpHrefIn []string) ContentConfigsAPIContentOstreeConfigsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentConfigsAPIContentOstreeConfigsListRequest) PulpIdIn(pulpIdIn []string) ContentConfigsAPIContentOstreeConfigsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r ContentConfigsAPIContentOstreeConfigsListRequest) Q(q string) ContentConfigsAPIContentOstreeConfigsListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF
func (r ContentConfigsAPIContentOstreeConfigsListRequest) RepositoryVersion(repositoryVersion string) ContentConfigsAPIContentOstreeConfigsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentConfigsAPIContentOstreeConfigsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentConfigsAPIContentOstreeConfigsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentConfigsAPIContentOstreeConfigsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentConfigsAPIContentOstreeConfigsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentConfigsAPIContentOstreeConfigsListRequest) Fields(fields []string) ContentConfigsAPIContentOstreeConfigsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentConfigsAPIContentOstreeConfigsListRequest) ExcludeFields(excludeFields []string) ContentConfigsAPIContentOstreeConfigsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentConfigsAPIContentOstreeConfigsListRequest) Execute() (*PaginatedostreeOstreeConfigResponseList, *http.Response, error) {
	return r.ApiService.ContentOstreeConfigsListExecute(r)
}

/*
ContentOstreeConfigsList List ostree configs

A ViewSet class for OSTree repository configurations.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentConfigsAPIContentOstreeConfigsListRequest
*/
func (a *ContentConfigsAPIService) ContentOstreeConfigsList(ctx context.Context, pulpDomain string) ContentConfigsAPIContentOstreeConfigsListRequest {
	return ContentConfigsAPIContentOstreeConfigsListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedostreeOstreeConfigResponseList
func (a *ContentConfigsAPIService) ContentOstreeConfigsListExecute(r ContentConfigsAPIContentOstreeConfigsListRequest) (*PaginatedostreeOstreeConfigResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedostreeOstreeConfigResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentConfigsAPIService.ContentOstreeConfigsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/ostree/configs/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.orphanedFor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orphaned_for", r.orphanedFor, "")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
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

type ContentConfigsAPIContentOstreeConfigsReadRequest struct {
	ctx context.Context
	ApiService *ContentConfigsAPIService
	ostreeOstreeConfigHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentConfigsAPIContentOstreeConfigsReadRequest) Fields(fields []string) ContentConfigsAPIContentOstreeConfigsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentConfigsAPIContentOstreeConfigsReadRequest) ExcludeFields(excludeFields []string) ContentConfigsAPIContentOstreeConfigsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentConfigsAPIContentOstreeConfigsReadRequest) Execute() (*OstreeOstreeConfigResponse, *http.Response, error) {
	return r.ApiService.ContentOstreeConfigsReadExecute(r)
}

/*
ContentOstreeConfigsRead Inspect an ostree config

A ViewSet class for OSTree repository configurations.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ostreeOstreeConfigHref
 @return ContentConfigsAPIContentOstreeConfigsReadRequest
*/
func (a *ContentConfigsAPIService) ContentOstreeConfigsRead(ctx context.Context, ostreeOstreeConfigHref string) ContentConfigsAPIContentOstreeConfigsReadRequest {
	return ContentConfigsAPIContentOstreeConfigsReadRequest{
		ApiService: a,
		ctx: ctx,
		ostreeOstreeConfigHref: ostreeOstreeConfigHref,
	}
}

// Execute executes the request
//  @return OstreeOstreeConfigResponse
func (a *ContentConfigsAPIService) ContentOstreeConfigsReadExecute(r ContentConfigsAPIContentOstreeConfigsReadRequest) (*OstreeOstreeConfigResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OstreeOstreeConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentConfigsAPIService.ContentOstreeConfigsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ostree_ostree_config_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ostree_ostree_config_href"+"}", url.PathEscape(parameterValueToString(r.ostreeOstreeConfigHref, "ostreeOstreeConfigHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
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