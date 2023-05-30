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


// ImportersPulpImportsAPIService ImportersPulpImportsAPI service
type ImportersPulpImportsAPIService service

type ImportersPulpImportsAPIImportersCorePulpImportsCreateRequest struct {
	ctx context.Context
	ApiService *ImportersPulpImportsAPIService
	pulpImporterHref string
	pulpImport *PulpImport
}

func (r ImportersPulpImportsAPIImportersCorePulpImportsCreateRequest) PulpImport(pulpImport PulpImport) ImportersPulpImportsAPIImportersCorePulpImportsCreateRequest {
	r.pulpImport = &pulpImport
	return r
}

func (r ImportersPulpImportsAPIImportersCorePulpImportsCreateRequest) Execute() (*TaskGroupOperationResponse, *http.Response, error) {
	return r.ApiService.ImportersCorePulpImportsCreateExecute(r)
}

/*
ImportersCorePulpImportsCreate Create a pulp import

Trigger an asynchronous task to import a Pulp export.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpImporterHref
 @return ImportersPulpImportsAPIImportersCorePulpImportsCreateRequest
*/
func (a *ImportersPulpImportsAPIService) ImportersCorePulpImportsCreate(ctx context.Context, pulpImporterHref string) ImportersPulpImportsAPIImportersCorePulpImportsCreateRequest {
	return ImportersPulpImportsAPIImportersCorePulpImportsCreateRequest{
		ApiService: a,
		ctx: ctx,
		pulpImporterHref: pulpImporterHref,
	}
}

// Execute executes the request
//  @return TaskGroupOperationResponse
func (a *ImportersPulpImportsAPIService) ImportersCorePulpImportsCreateExecute(r ImportersPulpImportsAPIImportersCorePulpImportsCreateRequest) (*TaskGroupOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskGroupOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpImportsAPIService.ImportersCorePulpImportsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{pulp_importer_href}imports/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_importer_href"+"}", url.PathEscape(parameterValueToString(r.pulpImporterHref, "pulpImporterHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pulpImport == nil {
		return localVarReturnValue, nil, reportError("pulpImport is required and must be specified")
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
	localVarPostBody = r.pulpImport
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

type ImportersPulpImportsAPIImportersCorePulpImportsDeleteRequest struct {
	ctx context.Context
	ApiService *ImportersPulpImportsAPIService
	pulpPulpImportHref string
}

func (r ImportersPulpImportsAPIImportersCorePulpImportsDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ImportersCorePulpImportsDeleteExecute(r)
}

/*
ImportersCorePulpImportsDelete Delete a pulp import

ViewSet for PulpImports.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpPulpImportHref
 @return ImportersPulpImportsAPIImportersCorePulpImportsDeleteRequest
*/
func (a *ImportersPulpImportsAPIService) ImportersCorePulpImportsDelete(ctx context.Context, pulpPulpImportHref string) ImportersPulpImportsAPIImportersCorePulpImportsDeleteRequest {
	return ImportersPulpImportsAPIImportersCorePulpImportsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		pulpPulpImportHref: pulpPulpImportHref,
	}
}

// Execute executes the request
func (a *ImportersPulpImportsAPIService) ImportersCorePulpImportsDeleteExecute(r ImportersPulpImportsAPIImportersCorePulpImportsDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpImportsAPIService.ImportersCorePulpImportsDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{pulp_pulp_import_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_pulp_import_href"+"}", url.PathEscape(parameterValueToString(r.pulpPulpImportHref, "pulpPulpImportHref")), -1)
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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ImportersPulpImportsAPIImportersCorePulpImportsListRequest struct {
	ctx context.Context
	ApiService *ImportersPulpImportsAPIService
	pulpImporterHref string
	limit *int32
	offset *int32
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ImportersPulpImportsAPIImportersCorePulpImportsListRequest) Limit(limit int32) ImportersPulpImportsAPIImportersCorePulpImportsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ImportersPulpImportsAPIImportersCorePulpImportsListRequest) Offset(offset int32) ImportersPulpImportsAPIImportersCorePulpImportsListRequest {
	r.offset = &offset
	return r
}

// A list of fields to include in the response.
func (r ImportersPulpImportsAPIImportersCorePulpImportsListRequest) Fields(fields []string) ImportersPulpImportsAPIImportersCorePulpImportsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ImportersPulpImportsAPIImportersCorePulpImportsListRequest) ExcludeFields(excludeFields []string) ImportersPulpImportsAPIImportersCorePulpImportsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ImportersPulpImportsAPIImportersCorePulpImportsListRequest) Execute() (*PaginatedImportResponseList, *http.Response, error) {
	return r.ApiService.ImportersCorePulpImportsListExecute(r)
}

/*
ImportersCorePulpImportsList List pulp imports

ViewSet for PulpImports.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpImporterHref
 @return ImportersPulpImportsAPIImportersCorePulpImportsListRequest
*/
func (a *ImportersPulpImportsAPIService) ImportersCorePulpImportsList(ctx context.Context, pulpImporterHref string) ImportersPulpImportsAPIImportersCorePulpImportsListRequest {
	return ImportersPulpImportsAPIImportersCorePulpImportsListRequest{
		ApiService: a,
		ctx: ctx,
		pulpImporterHref: pulpImporterHref,
	}
}

// Execute executes the request
//  @return PaginatedImportResponseList
func (a *ImportersPulpImportsAPIService) ImportersCorePulpImportsListExecute(r ImportersPulpImportsAPIImportersCorePulpImportsListRequest) (*PaginatedImportResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedImportResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpImportsAPIService.ImportersCorePulpImportsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{pulp_importer_href}imports/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_importer_href"+"}", url.PathEscape(parameterValueToString(r.pulpImporterHref, "pulpImporterHref")), -1)
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

type ImportersPulpImportsAPIImportersCorePulpImportsReadRequest struct {
	ctx context.Context
	ApiService *ImportersPulpImportsAPIService
	pulpPulpImportHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ImportersPulpImportsAPIImportersCorePulpImportsReadRequest) Fields(fields []string) ImportersPulpImportsAPIImportersCorePulpImportsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ImportersPulpImportsAPIImportersCorePulpImportsReadRequest) ExcludeFields(excludeFields []string) ImportersPulpImportsAPIImportersCorePulpImportsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ImportersPulpImportsAPIImportersCorePulpImportsReadRequest) Execute() (*ImportResponse, *http.Response, error) {
	return r.ApiService.ImportersCorePulpImportsReadExecute(r)
}

/*
ImportersCorePulpImportsRead Inspect a pulp import

ViewSet for PulpImports.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpPulpImportHref
 @return ImportersPulpImportsAPIImportersCorePulpImportsReadRequest
*/
func (a *ImportersPulpImportsAPIService) ImportersCorePulpImportsRead(ctx context.Context, pulpPulpImportHref string) ImportersPulpImportsAPIImportersCorePulpImportsReadRequest {
	return ImportersPulpImportsAPIImportersCorePulpImportsReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpPulpImportHref: pulpPulpImportHref,
	}
}

// Execute executes the request
//  @return ImportResponse
func (a *ImportersPulpImportsAPIService) ImportersCorePulpImportsReadExecute(r ImportersPulpImportsAPIImportersCorePulpImportsReadRequest) (*ImportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ImportResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpImportsAPIService.ImportersCorePulpImportsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{pulp_pulp_import_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_pulp_import_href"+"}", url.PathEscape(parameterValueToString(r.pulpPulpImportHref, "pulpPulpImportHref")), -1)
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
