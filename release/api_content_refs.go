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


// ContentRefsAPIService ContentRefsAPI service
type ContentRefsAPIService service

type ContentRefsAPIContentOstreeRefsListRequest struct {
	ctx context.Context
	ApiService *ContentRefsAPIService
	checksum *string
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
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

func (r ContentRefsAPIContentOstreeRefsListRequest) Checksum(checksum string) ContentRefsAPIContentOstreeRefsListRequest {
	r.checksum = &checksum
	return r
}

// Number of results to return per page.
func (r ContentRefsAPIContentOstreeRefsListRequest) Limit(limit int32) ContentRefsAPIContentOstreeRefsListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r ContentRefsAPIContentOstreeRefsListRequest) Name(name string) ContentRefsAPIContentOstreeRefsListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r ContentRefsAPIContentOstreeRefsListRequest) NameContains(nameContains string) ContentRefsAPIContentOstreeRefsListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r ContentRefsAPIContentOstreeRefsListRequest) NameIcontains(nameIcontains string) ContentRefsAPIContentOstreeRefsListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r ContentRefsAPIContentOstreeRefsListRequest) NameIn(nameIn []string) ContentRefsAPIContentOstreeRefsListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r ContentRefsAPIContentOstreeRefsListRequest) NameStartswith(nameStartswith string) ContentRefsAPIContentOstreeRefsListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r ContentRefsAPIContentOstreeRefsListRequest) Offset(offset int32) ContentRefsAPIContentOstreeRefsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;relative_path&#x60; - Relative path * &#x60;-relative_path&#x60; - Relative path (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentRefsAPIContentOstreeRefsListRequest) Ordering(ordering []string) ContentRefsAPIContentOstreeRefsListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ContentRefsAPIContentOstreeRefsListRequest) PulpHrefIn(pulpHrefIn []string) ContentRefsAPIContentOstreeRefsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentRefsAPIContentOstreeRefsListRequest) PulpIdIn(pulpIdIn []string) ContentRefsAPIContentOstreeRefsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Repository Version referenced by HREF
func (r ContentRefsAPIContentOstreeRefsListRequest) RepositoryVersion(repositoryVersion string) ContentRefsAPIContentOstreeRefsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentRefsAPIContentOstreeRefsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentRefsAPIContentOstreeRefsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentRefsAPIContentOstreeRefsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentRefsAPIContentOstreeRefsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentRefsAPIContentOstreeRefsListRequest) Fields(fields []string) ContentRefsAPIContentOstreeRefsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentRefsAPIContentOstreeRefsListRequest) ExcludeFields(excludeFields []string) ContentRefsAPIContentOstreeRefsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentRefsAPIContentOstreeRefsListRequest) Execute() (*PaginatedostreeOstreeRefResponseList, *http.Response, error) {
	return r.ApiService.ContentOstreeRefsListExecute(r)
}

/*
ContentOstreeRefsList List ostree refs

A ViewSet class for OSTree head commits.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentRefsAPIContentOstreeRefsListRequest
*/
func (a *ContentRefsAPIService) ContentOstreeRefsList(ctx context.Context) ContentRefsAPIContentOstreeRefsListRequest {
	return ContentRefsAPIContentOstreeRefsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedostreeOstreeRefResponseList
func (a *ContentRefsAPIService) ContentOstreeRefsListExecute(r ContentRefsAPIContentOstreeRefsListRequest) (*PaginatedostreeOstreeRefResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedostreeOstreeRefResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentRefsAPIService.ContentOstreeRefsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/ostree/refs/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.checksum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "checksum", r.checksum, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.nameContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__contains", r.nameContains, "")
	}
	if r.nameIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__icontains", r.nameIcontains, "")
	}
	if r.nameIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__in", r.nameIn, "csv")
	}
	if r.nameStartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__startswith", r.nameStartswith, "")
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

type ContentRefsAPIContentOstreeRefsReadRequest struct {
	ctx context.Context
	ApiService *ContentRefsAPIService
	ostreeOstreeRefHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentRefsAPIContentOstreeRefsReadRequest) Fields(fields []string) ContentRefsAPIContentOstreeRefsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentRefsAPIContentOstreeRefsReadRequest) ExcludeFields(excludeFields []string) ContentRefsAPIContentOstreeRefsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentRefsAPIContentOstreeRefsReadRequest) Execute() (*OstreeOstreeRefResponse, *http.Response, error) {
	return r.ApiService.ContentOstreeRefsReadExecute(r)
}

/*
ContentOstreeRefsRead Inspect an ostree ref

A ViewSet class for OSTree head commits.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ostreeOstreeRefHref
 @return ContentRefsAPIContentOstreeRefsReadRequest
*/
func (a *ContentRefsAPIService) ContentOstreeRefsRead(ctx context.Context, ostreeOstreeRefHref string) ContentRefsAPIContentOstreeRefsReadRequest {
	return ContentRefsAPIContentOstreeRefsReadRequest{
		ApiService: a,
		ctx: ctx,
		ostreeOstreeRefHref: ostreeOstreeRefHref,
	}
}

// Execute executes the request
//  @return OstreeOstreeRefResponse
func (a *ContentRefsAPIService) ContentOstreeRefsReadExecute(r ContentRefsAPIContentOstreeRefsReadRequest) (*OstreeOstreeRefResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OstreeOstreeRefResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentRefsAPIService.ContentOstreeRefsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ostree_ostree_ref_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ostree_ostree_ref_href"+"}", url.PathEscape(parameterValueToString(r.ostreeOstreeRefHref, "ostreeOstreeRefHref")), -1)
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
