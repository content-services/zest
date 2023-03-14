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


// ExportersFilesystemApiService ExportersFilesystemApi service
type ExportersFilesystemApiService service

type ExportersFilesystemApiExportersCoreFilesystemCreateRequest struct {
	ctx context.Context
	ApiService *ExportersFilesystemApiService
	filesystemExporter *FilesystemExporter
}

func (r ExportersFilesystemApiExportersCoreFilesystemCreateRequest) FilesystemExporter(filesystemExporter FilesystemExporter) ExportersFilesystemApiExportersCoreFilesystemCreateRequest {
	r.filesystemExporter = &filesystemExporter
	return r
}

func (r ExportersFilesystemApiExportersCoreFilesystemCreateRequest) Execute() (*FilesystemExporterResponse, *http.Response, error) {
	return r.ApiService.ExportersCoreFilesystemCreateExecute(r)
}

/*
ExportersCoreFilesystemCreate Create a filesystem exporter

Endpoint for managing FilesystemExporters. FilesystemExporters are provided as a tech preview.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ExportersFilesystemApiExportersCoreFilesystemCreateRequest
*/
func (a *ExportersFilesystemApiService) ExportersCoreFilesystemCreate(ctx context.Context) ExportersFilesystemApiExportersCoreFilesystemCreateRequest {
	return ExportersFilesystemApiExportersCoreFilesystemCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FilesystemExporterResponse
func (a *ExportersFilesystemApiService) ExportersCoreFilesystemCreateExecute(r ExportersFilesystemApiExportersCoreFilesystemCreateRequest) (*FilesystemExporterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FilesystemExporterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersFilesystemApiService.ExportersCoreFilesystemCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/exporters/core/filesystem/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.filesystemExporter == nil {
		return localVarReturnValue, nil, reportError("filesystemExporter is required and must be specified")
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
	localVarPostBody = r.filesystemExporter
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

type ExportersFilesystemApiExportersCoreFilesystemDeleteRequest struct {
	ctx context.Context
	ApiService *ExportersFilesystemApiService
	filesystemExporterHref string
}

func (r ExportersFilesystemApiExportersCoreFilesystemDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ExportersCoreFilesystemDeleteExecute(r)
}

/*
ExportersCoreFilesystemDelete Delete a filesystem exporter

Trigger an asynchronous delete task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param filesystemExporterHref
 @return ExportersFilesystemApiExportersCoreFilesystemDeleteRequest
*/
func (a *ExportersFilesystemApiService) ExportersCoreFilesystemDelete(ctx context.Context, filesystemExporterHref string) ExportersFilesystemApiExportersCoreFilesystemDeleteRequest {
	return ExportersFilesystemApiExportersCoreFilesystemDeleteRequest{
		ApiService: a,
		ctx: ctx,
		filesystemExporterHref: filesystemExporterHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ExportersFilesystemApiService) ExportersCoreFilesystemDeleteExecute(r ExportersFilesystemApiExportersCoreFilesystemDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersFilesystemApiService.ExportersCoreFilesystemDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{filesystem_exporter_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"filesystem_exporter_href"+"}", url.PathEscape(parameterValueToString(r.filesystemExporterHref, "filesystemExporterHref")), -1)
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

type ExportersFilesystemApiExportersCoreFilesystemListRequest struct {
	ctx context.Context
	ApiService *ExportersFilesystemApiService
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
	offset *int32
	ordering *[]string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ExportersFilesystemApiExportersCoreFilesystemListRequest) Limit(limit int32) ExportersFilesystemApiExportersCoreFilesystemListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r ExportersFilesystemApiExportersCoreFilesystemListRequest) Name(name string) ExportersFilesystemApiExportersCoreFilesystemListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r ExportersFilesystemApiExportersCoreFilesystemListRequest) NameContains(nameContains string) ExportersFilesystemApiExportersCoreFilesystemListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r ExportersFilesystemApiExportersCoreFilesystemListRequest) NameIcontains(nameIcontains string) ExportersFilesystemApiExportersCoreFilesystemListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r ExportersFilesystemApiExportersCoreFilesystemListRequest) NameIn(nameIn []string) ExportersFilesystemApiExportersCoreFilesystemListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r ExportersFilesystemApiExportersCoreFilesystemListRequest) NameStartswith(nameStartswith string) ExportersFilesystemApiExportersCoreFilesystemListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r ExportersFilesystemApiExportersCoreFilesystemListRequest) Offset(offset int32) ExportersFilesystemApiExportersCoreFilesystemListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ExportersFilesystemApiExportersCoreFilesystemListRequest) Ordering(ordering []string) ExportersFilesystemApiExportersCoreFilesystemListRequest {
	r.ordering = &ordering
	return r
}

// A list of fields to include in the response.
func (r ExportersFilesystemApiExportersCoreFilesystemListRequest) Fields(fields []string) ExportersFilesystemApiExportersCoreFilesystemListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ExportersFilesystemApiExportersCoreFilesystemListRequest) ExcludeFields(excludeFields []string) ExportersFilesystemApiExportersCoreFilesystemListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ExportersFilesystemApiExportersCoreFilesystemListRequest) Execute() (*PaginatedFilesystemExporterResponseList, *http.Response, error) {
	return r.ApiService.ExportersCoreFilesystemListExecute(r)
}

/*
ExportersCoreFilesystemList List filesystem exporters

Endpoint for managing FilesystemExporters. FilesystemExporters are provided as a tech preview.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ExportersFilesystemApiExportersCoreFilesystemListRequest
*/
func (a *ExportersFilesystemApiService) ExportersCoreFilesystemList(ctx context.Context) ExportersFilesystemApiExportersCoreFilesystemListRequest {
	return ExportersFilesystemApiExportersCoreFilesystemListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedFilesystemExporterResponseList
func (a *ExportersFilesystemApiService) ExportersCoreFilesystemListExecute(r ExportersFilesystemApiExportersCoreFilesystemListRequest) (*PaginatedFilesystemExporterResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedFilesystemExporterResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersFilesystemApiService.ExportersCoreFilesystemList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/exporters/core/filesystem/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ExportersFilesystemApiExportersCoreFilesystemPartialUpdateRequest struct {
	ctx context.Context
	ApiService *ExportersFilesystemApiService
	filesystemExporterHref string
	patchedFilesystemExporter *PatchedFilesystemExporter
}

func (r ExportersFilesystemApiExportersCoreFilesystemPartialUpdateRequest) PatchedFilesystemExporter(patchedFilesystemExporter PatchedFilesystemExporter) ExportersFilesystemApiExportersCoreFilesystemPartialUpdateRequest {
	r.patchedFilesystemExporter = &patchedFilesystemExporter
	return r
}

func (r ExportersFilesystemApiExportersCoreFilesystemPartialUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ExportersCoreFilesystemPartialUpdateExecute(r)
}

/*
ExportersCoreFilesystemPartialUpdate Update a filesystem exporter

Trigger an asynchronous partial update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param filesystemExporterHref
 @return ExportersFilesystemApiExportersCoreFilesystemPartialUpdateRequest
*/
func (a *ExportersFilesystemApiService) ExportersCoreFilesystemPartialUpdate(ctx context.Context, filesystemExporterHref string) ExportersFilesystemApiExportersCoreFilesystemPartialUpdateRequest {
	return ExportersFilesystemApiExportersCoreFilesystemPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		filesystemExporterHref: filesystemExporterHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ExportersFilesystemApiService) ExportersCoreFilesystemPartialUpdateExecute(r ExportersFilesystemApiExportersCoreFilesystemPartialUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersFilesystemApiService.ExportersCoreFilesystemPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{filesystem_exporter_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"filesystem_exporter_href"+"}", url.PathEscape(parameterValueToString(r.filesystemExporterHref, "filesystemExporterHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedFilesystemExporter == nil {
		return localVarReturnValue, nil, reportError("patchedFilesystemExporter is required and must be specified")
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
	localVarPostBody = r.patchedFilesystemExporter
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

type ExportersFilesystemApiExportersCoreFilesystemReadRequest struct {
	ctx context.Context
	ApiService *ExportersFilesystemApiService
	filesystemExporterHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ExportersFilesystemApiExportersCoreFilesystemReadRequest) Fields(fields []string) ExportersFilesystemApiExportersCoreFilesystemReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ExportersFilesystemApiExportersCoreFilesystemReadRequest) ExcludeFields(excludeFields []string) ExportersFilesystemApiExportersCoreFilesystemReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ExportersFilesystemApiExportersCoreFilesystemReadRequest) Execute() (*FilesystemExporterResponse, *http.Response, error) {
	return r.ApiService.ExportersCoreFilesystemReadExecute(r)
}

/*
ExportersCoreFilesystemRead Inspect a filesystem exporter

Endpoint for managing FilesystemExporters. FilesystemExporters are provided as a tech preview.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param filesystemExporterHref
 @return ExportersFilesystemApiExportersCoreFilesystemReadRequest
*/
func (a *ExportersFilesystemApiService) ExportersCoreFilesystemRead(ctx context.Context, filesystemExporterHref string) ExportersFilesystemApiExportersCoreFilesystemReadRequest {
	return ExportersFilesystemApiExportersCoreFilesystemReadRequest{
		ApiService: a,
		ctx: ctx,
		filesystemExporterHref: filesystemExporterHref,
	}
}

// Execute executes the request
//  @return FilesystemExporterResponse
func (a *ExportersFilesystemApiService) ExportersCoreFilesystemReadExecute(r ExportersFilesystemApiExportersCoreFilesystemReadRequest) (*FilesystemExporterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FilesystemExporterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersFilesystemApiService.ExportersCoreFilesystemRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{filesystem_exporter_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"filesystem_exporter_href"+"}", url.PathEscape(parameterValueToString(r.filesystemExporterHref, "filesystemExporterHref")), -1)
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

type ExportersFilesystemApiExportersCoreFilesystemUpdateRequest struct {
	ctx context.Context
	ApiService *ExportersFilesystemApiService
	filesystemExporterHref string
	filesystemExporter *FilesystemExporter
}

func (r ExportersFilesystemApiExportersCoreFilesystemUpdateRequest) FilesystemExporter(filesystemExporter FilesystemExporter) ExportersFilesystemApiExportersCoreFilesystemUpdateRequest {
	r.filesystemExporter = &filesystemExporter
	return r
}

func (r ExportersFilesystemApiExportersCoreFilesystemUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ExportersCoreFilesystemUpdateExecute(r)
}

/*
ExportersCoreFilesystemUpdate Update a filesystem exporter

Trigger an asynchronous update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param filesystemExporterHref
 @return ExportersFilesystemApiExportersCoreFilesystemUpdateRequest
*/
func (a *ExportersFilesystemApiService) ExportersCoreFilesystemUpdate(ctx context.Context, filesystemExporterHref string) ExportersFilesystemApiExportersCoreFilesystemUpdateRequest {
	return ExportersFilesystemApiExportersCoreFilesystemUpdateRequest{
		ApiService: a,
		ctx: ctx,
		filesystemExporterHref: filesystemExporterHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ExportersFilesystemApiService) ExportersCoreFilesystemUpdateExecute(r ExportersFilesystemApiExportersCoreFilesystemUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersFilesystemApiService.ExportersCoreFilesystemUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{filesystem_exporter_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"filesystem_exporter_href"+"}", url.PathEscape(parameterValueToString(r.filesystemExporterHref, "filesystemExporterHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.filesystemExporter == nil {
		return localVarReturnValue, nil, reportError("filesystemExporter is required and must be specified")
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
	localVarPostBody = r.filesystemExporter
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