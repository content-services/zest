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


// ContentModulemdDefaultsAPIService ContentModulemdDefaultsAPI service
type ContentModulemdDefaultsAPIService service

type ContentModulemdDefaultsAPIContentRpmModulemdDefaultsCreateRequest struct {
	ctx context.Context
	ApiService *ContentModulemdDefaultsAPIService
	pulpDomain string
	rpmModulemdDefaults *RpmModulemdDefaults
}

func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsCreateRequest) RpmModulemdDefaults(rpmModulemdDefaults RpmModulemdDefaults) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsCreateRequest {
	r.rpmModulemdDefaults = &rpmModulemdDefaults
	return r
}

func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentRpmModulemdDefaultsCreateExecute(r)
}

/*
ContentRpmModulemdDefaultsCreate Create a modulemd defaults

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentModulemdDefaultsAPIContentRpmModulemdDefaultsCreateRequest
*/
func (a *ContentModulemdDefaultsAPIService) ContentRpmModulemdDefaultsCreate(ctx context.Context, pulpDomain string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsCreateRequest {
	return ContentModulemdDefaultsAPIContentRpmModulemdDefaultsCreateRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentModulemdDefaultsAPIService) ContentRpmModulemdDefaultsCreateExecute(r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentModulemdDefaultsAPIService.ContentRpmModulemdDefaultsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/rpm/modulemd_defaults/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rpmModulemdDefaults == nil {
		return localVarReturnValue, nil, reportError("rpmModulemdDefaults is required and must be specified")
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
	localVarPostBody = r.rpmModulemdDefaults
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

type ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest struct {
	ctx context.Context
	ApiService *ContentModulemdDefaultsAPIService
	pulpDomain string
	limit *int32
	module *string
	moduleIn *[]string
	offset *int32
	ordering *[]string
	orphanedFor *float32
	pulpHrefIn *[]string
	pulpIdIn *[]string
	q *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	sha256 *string
	stream *string
	streamIn *[]string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) Limit(limit int32) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.limit = &limit
	return r
}

// Filter results where module matches value
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) Module(module string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.module = &module
	return r
}

// Filter results where module is in a comma-separated list of values
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) ModuleIn(moduleIn []string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.moduleIn = &moduleIn
	return r
}

// The initial index from which to return the results.
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) Offset(offset int32) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;module&#x60; - Module* &#x60;-module&#x60; - Module (descending)* &#x60;stream&#x60; - Stream* &#x60;-stream&#x60; - Stream (descending)* &#x60;profiles&#x60; - Profiles* &#x60;-profiles&#x60; - Profiles (descending)* &#x60;digest&#x60; - Digest* &#x60;-digest&#x60; - Digest (descending)* &#x60;snippet&#x60; - Snippet* &#x60;-snippet&#x60; - Snippet (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) Ordering(ordering []string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.ordering = &ordering
	return r
}

// Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME.
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) OrphanedFor(orphanedFor float32) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.orphanedFor = &orphanedFor
	return r
}

// Multiple values may be separated by commas.
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) PulpHrefIn(pulpHrefIn []string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) PulpIdIn(pulpIdIn []string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter results by using NOT, AND and OR operations on other filters
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) Q(q string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) RepositoryVersion(repositoryVersion string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) Sha256(sha256 string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.sha256 = &sha256
	return r
}

// Filter results where stream matches value
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) Stream(stream string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.stream = &stream
	return r
}

// Filter results where stream is in a comma-separated list of values
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) StreamIn(streamIn []string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.streamIn = &streamIn
	return r
}

// A list of fields to include in the response.
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) Fields(fields []string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) ExcludeFields(excludeFields []string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) Execute() (*PaginatedrpmModulemdDefaultsResponseList, *http.Response, error) {
	return r.ApiService.ContentRpmModulemdDefaultsListExecute(r)
}

/*
ContentRpmModulemdDefaultsList List modulemd defaultss

ViewSet for Modulemd.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest
*/
func (a *ContentModulemdDefaultsAPIService) ContentRpmModulemdDefaultsList(ctx context.Context, pulpDomain string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest {
	return ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedrpmModulemdDefaultsResponseList
func (a *ContentModulemdDefaultsAPIService) ContentRpmModulemdDefaultsListExecute(r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsListRequest) (*PaginatedrpmModulemdDefaultsResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedrpmModulemdDefaultsResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentModulemdDefaultsAPIService.ContentRpmModulemdDefaultsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/rpm/modulemd_defaults/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.module != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "module", r.module, "form", "")
	}
	if r.moduleIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "module__in", r.moduleIn, "form", "csv")
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
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "form", "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "form", "csv")
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
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha256", r.sha256, "form", "")
	}
	if r.stream != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stream", r.stream, "form", "")
	}
	if r.streamIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stream__in", r.streamIn, "form", "csv")
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

type ContentModulemdDefaultsAPIContentRpmModulemdDefaultsReadRequest struct {
	ctx context.Context
	ApiService *ContentModulemdDefaultsAPIService
	rpmModulemdDefaultsHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsReadRequest) Fields(fields []string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsReadRequest) ExcludeFields(excludeFields []string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsReadRequest) Execute() (*RpmModulemdDefaultsResponse, *http.Response, error) {
	return r.ApiService.ContentRpmModulemdDefaultsReadExecute(r)
}

/*
ContentRpmModulemdDefaultsRead Inspect a modulemd defaults

ViewSet for Modulemd.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmModulemdDefaultsHref
 @return ContentModulemdDefaultsAPIContentRpmModulemdDefaultsReadRequest
*/
func (a *ContentModulemdDefaultsAPIService) ContentRpmModulemdDefaultsRead(ctx context.Context, rpmModulemdDefaultsHref string) ContentModulemdDefaultsAPIContentRpmModulemdDefaultsReadRequest {
	return ContentModulemdDefaultsAPIContentRpmModulemdDefaultsReadRequest{
		ApiService: a,
		ctx: ctx,
		rpmModulemdDefaultsHref: rpmModulemdDefaultsHref,
	}
}

// Execute executes the request
//  @return RpmModulemdDefaultsResponse
func (a *ContentModulemdDefaultsAPIService) ContentRpmModulemdDefaultsReadExecute(r ContentModulemdDefaultsAPIContentRpmModulemdDefaultsReadRequest) (*RpmModulemdDefaultsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RpmModulemdDefaultsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentModulemdDefaultsAPIService.ContentRpmModulemdDefaultsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{rpm_modulemd_defaults_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_modulemd_defaults_href"+"}", url.PathEscape(parameterValueToString(r.rpmModulemdDefaultsHref, "rpmModulemdDefaultsHref")), -1)
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
