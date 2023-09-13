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


// ContentModulemdObsoletesAPIService ContentModulemdObsoletesAPI service
type ContentModulemdObsoletesAPIService service

type ContentModulemdObsoletesAPIContentRpmModulemdObsoletesCreateRequest struct {
	ctx context.Context
	ApiService *ContentModulemdObsoletesAPIService
	pulpDomain string
	rpmModulemdObsolete *RpmModulemdObsolete
}

func (r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesCreateRequest) RpmModulemdObsolete(rpmModulemdObsolete RpmModulemdObsolete) ContentModulemdObsoletesAPIContentRpmModulemdObsoletesCreateRequest {
	r.rpmModulemdObsolete = &rpmModulemdObsolete
	return r
}

func (r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentRpmModulemdObsoletesCreateExecute(r)
}

/*
ContentRpmModulemdObsoletesCreate Create a modulemd obsolete

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentModulemdObsoletesAPIContentRpmModulemdObsoletesCreateRequest
*/
func (a *ContentModulemdObsoletesAPIService) ContentRpmModulemdObsoletesCreate(ctx context.Context, pulpDomain string) ContentModulemdObsoletesAPIContentRpmModulemdObsoletesCreateRequest {
	return ContentModulemdObsoletesAPIContentRpmModulemdObsoletesCreateRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentModulemdObsoletesAPIService) ContentRpmModulemdObsoletesCreateExecute(r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentModulemdObsoletesAPIService.ContentRpmModulemdObsoletesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/{pulp_domain}/api/v3/content/rpm/modulemd_obsoletes/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rpmModulemdObsolete == nil {
		return localVarReturnValue, nil, reportError("rpmModulemdObsolete is required and must be specified")
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
	localVarPostBody = r.rpmModulemdObsolete
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

type ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest struct {
	ctx context.Context
	ApiService *ContentModulemdObsoletesAPIService
	pulpDomain string
	limit *int32
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest) Limit(limit int32) ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest) Offset(offset int32) ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest) Ordering(ordering []string) ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest) PulpHrefIn(pulpHrefIn []string) ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest) PulpIdIn(pulpIdIn []string) ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Repository Version referenced by HREF
func (r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest) RepositoryVersion(repositoryVersion string) ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest) Fields(fields []string) ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest) ExcludeFields(excludeFields []string) ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest) Execute() (*PaginatedrpmModulemdObsoleteResponseList, *http.Response, error) {
	return r.ApiService.ContentRpmModulemdObsoletesListExecute(r)
}

/*
ContentRpmModulemdObsoletesList List modulemd obsoletes

ViewSet for Modulemd.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest
*/
func (a *ContentModulemdObsoletesAPIService) ContentRpmModulemdObsoletesList(ctx context.Context, pulpDomain string) ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest {
	return ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedrpmModulemdObsoleteResponseList
func (a *ContentModulemdObsoletesAPIService) ContentRpmModulemdObsoletesListExecute(r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesListRequest) (*PaginatedrpmModulemdObsoleteResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedrpmModulemdObsoleteResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentModulemdObsoletesAPIService.ContentRpmModulemdObsoletesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/{pulp_domain}/api/v3/content/rpm/modulemd_obsoletes/"
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
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
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

type ContentModulemdObsoletesAPIContentRpmModulemdObsoletesReadRequest struct {
	ctx context.Context
	ApiService *ContentModulemdObsoletesAPIService
	rpmModulemdObsoleteHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesReadRequest) Fields(fields []string) ContentModulemdObsoletesAPIContentRpmModulemdObsoletesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesReadRequest) ExcludeFields(excludeFields []string) ContentModulemdObsoletesAPIContentRpmModulemdObsoletesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesReadRequest) Execute() (*RpmModulemdObsoleteResponse, *http.Response, error) {
	return r.ApiService.ContentRpmModulemdObsoletesReadExecute(r)
}

/*
ContentRpmModulemdObsoletesRead Inspect a modulemd obsolete

ViewSet for Modulemd.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmModulemdObsoleteHref
 @return ContentModulemdObsoletesAPIContentRpmModulemdObsoletesReadRequest
*/
func (a *ContentModulemdObsoletesAPIService) ContentRpmModulemdObsoletesRead(ctx context.Context, rpmModulemdObsoleteHref string) ContentModulemdObsoletesAPIContentRpmModulemdObsoletesReadRequest {
	return ContentModulemdObsoletesAPIContentRpmModulemdObsoletesReadRequest{
		ApiService: a,
		ctx: ctx,
		rpmModulemdObsoleteHref: rpmModulemdObsoleteHref,
	}
}

// Execute executes the request
//  @return RpmModulemdObsoleteResponse
func (a *ContentModulemdObsoletesAPIService) ContentRpmModulemdObsoletesReadExecute(r ContentModulemdObsoletesAPIContentRpmModulemdObsoletesReadRequest) (*RpmModulemdObsoleteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RpmModulemdObsoleteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentModulemdObsoletesAPIService.ContentRpmModulemdObsoletesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{rpm_modulemd_obsolete_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_modulemd_obsolete_href"+"}", url.PathEscape(parameterValueToString(r.rpmModulemdObsoleteHref, "rpmModulemdObsoleteHref")), -1)
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
