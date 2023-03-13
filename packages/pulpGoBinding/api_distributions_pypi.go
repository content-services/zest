/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpGoBinding

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// DistributionsPypiApiService DistributionsPypiApi service
type DistributionsPypiApiService service

type DistributionsPypiApiDistributionsPythonPypiCreateRequest struct {
	ctx context.Context
	ApiService *DistributionsPypiApiService
	pythonPythonDistribution *PythonPythonDistribution
}

func (r DistributionsPypiApiDistributionsPythonPypiCreateRequest) PythonPythonDistribution(pythonPythonDistribution PythonPythonDistribution) DistributionsPypiApiDistributionsPythonPypiCreateRequest {
	r.pythonPythonDistribution = &pythonPythonDistribution
	return r
}

func (r DistributionsPypiApiDistributionsPythonPypiCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsPythonPypiCreateExecute(r)
}

/*
DistributionsPythonPypiCreate Create a python distribution

Trigger an asynchronous create task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DistributionsPypiApiDistributionsPythonPypiCreateRequest
*/
func (a *DistributionsPypiApiService) DistributionsPythonPypiCreate(ctx context.Context) DistributionsPypiApiDistributionsPythonPypiCreateRequest {
	return DistributionsPypiApiDistributionsPythonPypiCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsPypiApiService) DistributionsPythonPypiCreateExecute(r DistributionsPypiApiDistributionsPythonPypiCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsPypiApiService.DistributionsPythonPypiCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/distributions/python/pypi/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pythonPythonDistribution == nil {
		return localVarReturnValue, nil, reportError("pythonPythonDistribution is required and must be specified")
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
	localVarPostBody = r.pythonPythonDistribution
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

type DistributionsPypiApiDistributionsPythonPypiDeleteRequest struct {
	ctx context.Context
	ApiService *DistributionsPypiApiService
	pythonPythonDistributionHref string
}

func (r DistributionsPypiApiDistributionsPythonPypiDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsPythonPypiDeleteExecute(r)
}

/*
DistributionsPythonPypiDelete Delete a python distribution

Trigger an asynchronous delete task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pythonPythonDistributionHref
 @return DistributionsPypiApiDistributionsPythonPypiDeleteRequest
*/
func (a *DistributionsPypiApiService) DistributionsPythonPypiDelete(ctx context.Context, pythonPythonDistributionHref string) DistributionsPypiApiDistributionsPythonPypiDeleteRequest {
	return DistributionsPypiApiDistributionsPythonPypiDeleteRequest{
		ApiService: a,
		ctx: ctx,
		pythonPythonDistributionHref: pythonPythonDistributionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsPypiApiService) DistributionsPythonPypiDeleteExecute(r DistributionsPypiApiDistributionsPythonPypiDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsPypiApiService.DistributionsPythonPypiDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{python_python_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"python_python_distribution_href"+"}", url.PathEscape(parameterValueToString(r.pythonPythonDistributionHref, "pythonPythonDistributionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type DistributionsPypiApiDistributionsPythonPypiListRequest struct {
	ctx context.Context
	ApiService *DistributionsPypiApiService
	basePath *string
	basePathContains *string
	basePathIcontains *string
	basePathIn *[]string
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
	offset *int32
	ordering *[]string
	pulpLabelSelect *string
	withContent *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where base_path matches value
func (r DistributionsPypiApiDistributionsPythonPypiListRequest) BasePath(basePath string) DistributionsPypiApiDistributionsPythonPypiListRequest {
	r.basePath = &basePath
	return r
}

// Filter results where base_path contains value
func (r DistributionsPypiApiDistributionsPythonPypiListRequest) BasePathContains(basePathContains string) DistributionsPypiApiDistributionsPythonPypiListRequest {
	r.basePathContains = &basePathContains
	return r
}

// Filter results where base_path contains value
func (r DistributionsPypiApiDistributionsPythonPypiListRequest) BasePathIcontains(basePathIcontains string) DistributionsPypiApiDistributionsPythonPypiListRequest {
	r.basePathIcontains = &basePathIcontains
	return r
}

// Filter results where base_path is in a comma-separated list of values
func (r DistributionsPypiApiDistributionsPythonPypiListRequest) BasePathIn(basePathIn []string) DistributionsPypiApiDistributionsPythonPypiListRequest {
	r.basePathIn = &basePathIn
	return r
}

// Number of results to return per page.
func (r DistributionsPypiApiDistributionsPythonPypiListRequest) Limit(limit int32) DistributionsPypiApiDistributionsPythonPypiListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r DistributionsPypiApiDistributionsPythonPypiListRequest) Name(name string) DistributionsPypiApiDistributionsPythonPypiListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r DistributionsPypiApiDistributionsPythonPypiListRequest) NameContains(nameContains string) DistributionsPypiApiDistributionsPythonPypiListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r DistributionsPypiApiDistributionsPythonPypiListRequest) NameIcontains(nameIcontains string) DistributionsPypiApiDistributionsPythonPypiListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r DistributionsPypiApiDistributionsPythonPypiListRequest) NameIn(nameIn []string) DistributionsPypiApiDistributionsPythonPypiListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r DistributionsPypiApiDistributionsPythonPypiListRequest) NameStartswith(nameStartswith string) DistributionsPypiApiDistributionsPythonPypiListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r DistributionsPypiApiDistributionsPythonPypiListRequest) Offset(offset int32) DistributionsPypiApiDistributionsPythonPypiListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r DistributionsPypiApiDistributionsPythonPypiListRequest) Ordering(ordering []string) DistributionsPypiApiDistributionsPythonPypiListRequest {
	r.ordering = &ordering
	return r
}

// Filter labels by search string
func (r DistributionsPypiApiDistributionsPythonPypiListRequest) PulpLabelSelect(pulpLabelSelect string) DistributionsPypiApiDistributionsPythonPypiListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Filter distributions based on the content served by them
func (r DistributionsPypiApiDistributionsPythonPypiListRequest) WithContent(withContent string) DistributionsPypiApiDistributionsPythonPypiListRequest {
	r.withContent = &withContent
	return r
}

// A list of fields to include in the response.
func (r DistributionsPypiApiDistributionsPythonPypiListRequest) Fields(fields []string) DistributionsPypiApiDistributionsPythonPypiListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r DistributionsPypiApiDistributionsPythonPypiListRequest) ExcludeFields(excludeFields []string) DistributionsPypiApiDistributionsPythonPypiListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r DistributionsPypiApiDistributionsPythonPypiListRequest) Execute() (*PaginatedpythonPythonDistributionResponseList, *http.Response, error) {
	return r.ApiService.DistributionsPythonPypiListExecute(r)
}

/*
DistributionsPythonPypiList List python distributions


Pulp Python Distributions are used to distribute Python content from
Python Repositories or
Python Publications.  Pulp Python
Distributions should not be confused with "Python Distribution" as defined by the Python
community. In Pulp usage, Python content is referred to as Python Package Content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DistributionsPypiApiDistributionsPythonPypiListRequest
*/
func (a *DistributionsPypiApiService) DistributionsPythonPypiList(ctx context.Context) DistributionsPypiApiDistributionsPythonPypiListRequest {
	return DistributionsPypiApiDistributionsPythonPypiListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedpythonPythonDistributionResponseList
func (a *DistributionsPypiApiService) DistributionsPythonPypiListExecute(r DistributionsPypiApiDistributionsPythonPypiListRequest) (*PaginatedpythonPythonDistributionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedpythonPythonDistributionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsPypiApiService.DistributionsPythonPypiList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/distributions/python/pypi/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.basePath != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path", r.basePath, "")
	}
	if r.basePathContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path__contains", r.basePathContains, "")
	}
	if r.basePathIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path__icontains", r.basePathIcontains, "")
	}
	if r.basePathIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path__in", r.basePathIn, "csv")
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
	if r.pulpLabelSelect != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_label_select", r.pulpLabelSelect, "")
	}
	if r.withContent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "with_content", r.withContent, "")
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

type DistributionsPypiApiDistributionsPythonPypiPartialUpdateRequest struct {
	ctx context.Context
	ApiService *DistributionsPypiApiService
	pythonPythonDistributionHref string
	patchedpythonPythonDistribution *PatchedpythonPythonDistribution
}

func (r DistributionsPypiApiDistributionsPythonPypiPartialUpdateRequest) PatchedpythonPythonDistribution(patchedpythonPythonDistribution PatchedpythonPythonDistribution) DistributionsPypiApiDistributionsPythonPypiPartialUpdateRequest {
	r.patchedpythonPythonDistribution = &patchedpythonPythonDistribution
	return r
}

func (r DistributionsPypiApiDistributionsPythonPypiPartialUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsPythonPypiPartialUpdateExecute(r)
}

/*
DistributionsPythonPypiPartialUpdate Update a python distribution

Trigger an asynchronous partial update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pythonPythonDistributionHref
 @return DistributionsPypiApiDistributionsPythonPypiPartialUpdateRequest
*/
func (a *DistributionsPypiApiService) DistributionsPythonPypiPartialUpdate(ctx context.Context, pythonPythonDistributionHref string) DistributionsPypiApiDistributionsPythonPypiPartialUpdateRequest {
	return DistributionsPypiApiDistributionsPythonPypiPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		pythonPythonDistributionHref: pythonPythonDistributionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsPypiApiService) DistributionsPythonPypiPartialUpdateExecute(r DistributionsPypiApiDistributionsPythonPypiPartialUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsPypiApiService.DistributionsPythonPypiPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{python_python_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"python_python_distribution_href"+"}", url.PathEscape(parameterValueToString(r.pythonPythonDistributionHref, "pythonPythonDistributionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedpythonPythonDistribution == nil {
		return localVarReturnValue, nil, reportError("patchedpythonPythonDistribution is required and must be specified")
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
	localVarPostBody = r.patchedpythonPythonDistribution
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

type DistributionsPypiApiDistributionsPythonPypiReadRequest struct {
	ctx context.Context
	ApiService *DistributionsPypiApiService
	pythonPythonDistributionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r DistributionsPypiApiDistributionsPythonPypiReadRequest) Fields(fields []string) DistributionsPypiApiDistributionsPythonPypiReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r DistributionsPypiApiDistributionsPythonPypiReadRequest) ExcludeFields(excludeFields []string) DistributionsPypiApiDistributionsPythonPypiReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r DistributionsPypiApiDistributionsPythonPypiReadRequest) Execute() (*PythonPythonDistributionResponse, *http.Response, error) {
	return r.ApiService.DistributionsPythonPypiReadExecute(r)
}

/*
DistributionsPythonPypiRead Inspect a python distribution


Pulp Python Distributions are used to distribute Python content from
Python Repositories or
Python Publications.  Pulp Python
Distributions should not be confused with "Python Distribution" as defined by the Python
community. In Pulp usage, Python content is referred to as Python Package Content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pythonPythonDistributionHref
 @return DistributionsPypiApiDistributionsPythonPypiReadRequest
*/
func (a *DistributionsPypiApiService) DistributionsPythonPypiRead(ctx context.Context, pythonPythonDistributionHref string) DistributionsPypiApiDistributionsPythonPypiReadRequest {
	return DistributionsPypiApiDistributionsPythonPypiReadRequest{
		ApiService: a,
		ctx: ctx,
		pythonPythonDistributionHref: pythonPythonDistributionHref,
	}
}

// Execute executes the request
//  @return PythonPythonDistributionResponse
func (a *DistributionsPypiApiService) DistributionsPythonPypiReadExecute(r DistributionsPypiApiDistributionsPythonPypiReadRequest) (*PythonPythonDistributionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PythonPythonDistributionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsPypiApiService.DistributionsPythonPypiRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{python_python_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"python_python_distribution_href"+"}", url.PathEscape(parameterValueToString(r.pythonPythonDistributionHref, "pythonPythonDistributionHref")), -1)
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

type DistributionsPypiApiDistributionsPythonPypiUpdateRequest struct {
	ctx context.Context
	ApiService *DistributionsPypiApiService
	pythonPythonDistributionHref string
	pythonPythonDistribution *PythonPythonDistribution
}

func (r DistributionsPypiApiDistributionsPythonPypiUpdateRequest) PythonPythonDistribution(pythonPythonDistribution PythonPythonDistribution) DistributionsPypiApiDistributionsPythonPypiUpdateRequest {
	r.pythonPythonDistribution = &pythonPythonDistribution
	return r
}

func (r DistributionsPypiApiDistributionsPythonPypiUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsPythonPypiUpdateExecute(r)
}

/*
DistributionsPythonPypiUpdate Update a python distribution

Trigger an asynchronous update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pythonPythonDistributionHref
 @return DistributionsPypiApiDistributionsPythonPypiUpdateRequest
*/
func (a *DistributionsPypiApiService) DistributionsPythonPypiUpdate(ctx context.Context, pythonPythonDistributionHref string) DistributionsPypiApiDistributionsPythonPypiUpdateRequest {
	return DistributionsPypiApiDistributionsPythonPypiUpdateRequest{
		ApiService: a,
		ctx: ctx,
		pythonPythonDistributionHref: pythonPythonDistributionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsPypiApiService) DistributionsPythonPypiUpdateExecute(r DistributionsPypiApiDistributionsPythonPypiUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsPypiApiService.DistributionsPythonPypiUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{python_python_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"python_python_distribution_href"+"}", url.PathEscape(parameterValueToString(r.pythonPythonDistributionHref, "pythonPythonDistributionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pythonPythonDistribution == nil {
		return localVarReturnValue, nil, reportError("pythonPythonDistribution is required and must be specified")
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
	localVarPostBody = r.pythonPythonDistribution
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