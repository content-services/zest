/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zest/release/v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// ContentContentApiService ContentContentApi service
type ContentContentApiService service

type ContentContentApiContentOstreeContentListRequest struct {
	ctx context.Context
	ApiService *ContentContentApiService
	limit *int32
	offset *int32
	ordering *[]string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ContentContentApiContentOstreeContentListRequest) Limit(limit int32) ContentContentApiContentOstreeContentListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentContentApiContentOstreeContentListRequest) Offset(offset int32) ContentContentApiContentOstreeContentListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ContentContentApiContentOstreeContentListRequest) Ordering(ordering []string) ContentContentApiContentOstreeContentListRequest {
	r.ordering = &ordering
	return r
}

// Repository Version referenced by HREF
func (r ContentContentApiContentOstreeContentListRequest) RepositoryVersion(repositoryVersion string) ContentContentApiContentOstreeContentListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentContentApiContentOstreeContentListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentContentApiContentOstreeContentListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentContentApiContentOstreeContentListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentContentApiContentOstreeContentListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentContentApiContentOstreeContentListRequest) Fields(fields []string) ContentContentApiContentOstreeContentListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentContentApiContentOstreeContentListRequest) ExcludeFields(excludeFields []string) ContentContentApiContentOstreeContentListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentContentApiContentOstreeContentListRequest) Execute() (*PaginatedostreeOstreeContentResponseList, *http.Response, error) {
	return r.ApiService.ContentOstreeContentListExecute(r)
}

/*
ContentOstreeContentList List ostree contents

A ViewSet class for uncategorized content units (e.g., static deltas).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentContentApiContentOstreeContentListRequest
*/
func (a *ContentContentApiService) ContentOstreeContentList(ctx context.Context) ContentContentApiContentOstreeContentListRequest {
	return ContentContentApiContentOstreeContentListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedostreeOstreeContentResponseList
func (a *ContentContentApiService) ContentOstreeContentListExecute(r ContentContentApiContentOstreeContentListRequest) (*PaginatedostreeOstreeContentResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedostreeOstreeContentResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentContentApiService.ContentOstreeContentList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/ostree/content/"
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

type ContentContentApiContentOstreeContentReadRequest struct {
	ctx context.Context
	ApiService *ContentContentApiService
	ostreeOstreeContentHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentContentApiContentOstreeContentReadRequest) Fields(fields []string) ContentContentApiContentOstreeContentReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentContentApiContentOstreeContentReadRequest) ExcludeFields(excludeFields []string) ContentContentApiContentOstreeContentReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentContentApiContentOstreeContentReadRequest) Execute() (*OstreeOstreeContentResponse, *http.Response, error) {
	return r.ApiService.ContentOstreeContentReadExecute(r)
}

/*
ContentOstreeContentRead Inspect an ostree content

A ViewSet class for uncategorized content units (e.g., static deltas).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ostreeOstreeContentHref
 @return ContentContentApiContentOstreeContentReadRequest
*/
func (a *ContentContentApiService) ContentOstreeContentRead(ctx context.Context, ostreeOstreeContentHref string) ContentContentApiContentOstreeContentReadRequest {
	return ContentContentApiContentOstreeContentReadRequest{
		ApiService: a,
		ctx: ctx,
		ostreeOstreeContentHref: ostreeOstreeContentHref,
	}
}

// Execute executes the request
//  @return OstreeOstreeContentResponse
func (a *ContentContentApiService) ContentOstreeContentReadExecute(r ContentContentApiContentOstreeContentReadRequest) (*OstreeOstreeContentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OstreeOstreeContentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentContentApiService.ContentOstreeContentRead")
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
