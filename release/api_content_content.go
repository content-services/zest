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


// ContentContentAPIService ContentContentAPI service
type ContentContentAPIService service

type ContentContentAPIContentOstreeContentCreateRequest struct {
	ctx context.Context
	ApiService *ContentContentAPIService
	pulpDomain string
	ostreeOstreeContent *OstreeOstreeContent
}

func (r ContentContentAPIContentOstreeContentCreateRequest) OstreeOstreeContent(ostreeOstreeContent OstreeOstreeContent) ContentContentAPIContentOstreeContentCreateRequest {
	r.ostreeOstreeContent = &ostreeOstreeContent
	return r
}

func (r ContentContentAPIContentOstreeContentCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentOstreeContentCreateExecute(r)
}

/*
ContentOstreeContentCreate Create an ostree content

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentContentAPIContentOstreeContentCreateRequest
*/
func (a *ContentContentAPIService) ContentOstreeContentCreate(ctx context.Context, pulpDomain string) ContentContentAPIContentOstreeContentCreateRequest {
	return ContentContentAPIContentOstreeContentCreateRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentContentAPIService) ContentOstreeContentCreateExecute(r ContentContentAPIContentOstreeContentCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentContentAPIService.ContentOstreeContentCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/ostree/content/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ostreeOstreeContent == nil {
		return localVarReturnValue, nil, reportError("ostreeOstreeContent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

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
	// body params
	localVarPostBody = r.ostreeOstreeContent
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

type ContentContentAPIContentOstreeContentListRequest struct {
	ctx context.Context
	ApiService *ContentContentAPIService
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
func (r ContentContentAPIContentOstreeContentListRequest) Limit(limit int32) ContentContentAPIContentOstreeContentListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentContentAPIContentOstreeContentListRequest) Offset(offset int32) ContentContentAPIContentOstreeContentListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r ContentContentAPIContentOstreeContentListRequest) Ordering(ordering []string) ContentContentAPIContentOstreeContentListRequest {
	r.ordering = &ordering
	return r
}

// Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME.
func (r ContentContentAPIContentOstreeContentListRequest) OrphanedFor(orphanedFor float32) ContentContentAPIContentOstreeContentListRequest {
	r.orphanedFor = &orphanedFor
	return r
}

// Multiple values may be separated by commas.
func (r ContentContentAPIContentOstreeContentListRequest) PulpHrefIn(pulpHrefIn []string) ContentContentAPIContentOstreeContentListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentContentAPIContentOstreeContentListRequest) PulpIdIn(pulpIdIn []string) ContentContentAPIContentOstreeContentListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r ContentContentAPIContentOstreeContentListRequest) Q(q string) ContentContentAPIContentOstreeContentListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF
func (r ContentContentAPIContentOstreeContentListRequest) RepositoryVersion(repositoryVersion string) ContentContentAPIContentOstreeContentListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentContentAPIContentOstreeContentListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentContentAPIContentOstreeContentListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentContentAPIContentOstreeContentListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentContentAPIContentOstreeContentListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentContentAPIContentOstreeContentListRequest) Fields(fields []string) ContentContentAPIContentOstreeContentListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentContentAPIContentOstreeContentListRequest) ExcludeFields(excludeFields []string) ContentContentAPIContentOstreeContentListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentContentAPIContentOstreeContentListRequest) Execute() (*PaginatedostreeOstreeContentResponseList, *http.Response, error) {
	return r.ApiService.ContentOstreeContentListExecute(r)
}

/*
ContentOstreeContentList List ostree contents

A ViewSet class for uncategorized content units (e.g., static deltas).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentContentAPIContentOstreeContentListRequest
*/
func (a *ContentContentAPIService) ContentOstreeContentList(ctx context.Context, pulpDomain string) ContentContentAPIContentOstreeContentListRequest {
	return ContentContentAPIContentOstreeContentListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedostreeOstreeContentResponseList
func (a *ContentContentAPIService) ContentOstreeContentListExecute(r ContentContentAPIContentOstreeContentListRequest) (*PaginatedostreeOstreeContentResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedostreeOstreeContentResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentContentAPIService.ContentOstreeContentList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/ostree/content/"
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
                               parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i).Interface(), "multi")
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
                               parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i).Interface(), "multi")
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

type ContentContentAPIContentOstreeContentReadRequest struct {
	ctx context.Context
	ApiService *ContentContentAPIService
	ostreeOstreeContentHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentContentAPIContentOstreeContentReadRequest) Fields(fields []string) ContentContentAPIContentOstreeContentReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentContentAPIContentOstreeContentReadRequest) ExcludeFields(excludeFields []string) ContentContentAPIContentOstreeContentReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentContentAPIContentOstreeContentReadRequest) Execute() (*OstreeOstreeContentResponse, *http.Response, error) {
	return r.ApiService.ContentOstreeContentReadExecute(r)
}

/*
ContentOstreeContentRead Inspect an ostree content

A ViewSet class for uncategorized content units (e.g., static deltas).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ostreeOstreeContentHref
 @return ContentContentAPIContentOstreeContentReadRequest
*/
func (a *ContentContentAPIService) ContentOstreeContentRead(ctx context.Context, ostreeOstreeContentHref string) ContentContentAPIContentOstreeContentReadRequest {
	return ContentContentAPIContentOstreeContentReadRequest{
		ApiService: a,
		ctx: ctx,
		ostreeOstreeContentHref: ostreeOstreeContentHref,
	}
}

// Execute executes the request
//  @return OstreeOstreeContentResponse
func (a *ContentContentAPIService) ContentOstreeContentReadExecute(r ContentContentAPIContentOstreeContentReadRequest) (*OstreeOstreeContentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OstreeOstreeContentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentContentAPIService.ContentOstreeContentRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ostree_ostree_content_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ostree_ostree_content_href"+"}", url.PathEscape(parameterValueToString(r.ostreeOstreeContentHref, "ostreeOstreeContentHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
                               parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i).Interface(), "multi")
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
                               parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i).Interface(), "multi")
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
