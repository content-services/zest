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


// ExportersPulpAPIService ExportersPulpAPI service
type ExportersPulpAPIService service

type ExportersPulpAPIExportersCorePulpCreateRequest struct {
	ctx context.Context
	ApiService *ExportersPulpAPIService
	pulpDomain string
	pulpExporter *PulpExporter
}

func (r ExportersPulpAPIExportersCorePulpCreateRequest) PulpExporter(pulpExporter PulpExporter) ExportersPulpAPIExportersCorePulpCreateRequest {
	r.pulpExporter = &pulpExporter
	return r
}

func (r ExportersPulpAPIExportersCorePulpCreateRequest) Execute() (*PulpExporterResponse, *http.Response, error) {
	return r.ApiService.ExportersCorePulpCreateExecute(r)
}

/*
ExportersCorePulpCreate Create a pulp exporter

ViewSet for viewing PulpExporters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ExportersPulpAPIExportersCorePulpCreateRequest
*/
func (a *ExportersPulpAPIService) ExportersCorePulpCreate(ctx context.Context, pulpDomain string) ExportersPulpAPIExportersCorePulpCreateRequest {
	return ExportersPulpAPIExportersCorePulpCreateRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PulpExporterResponse
func (a *ExportersPulpAPIService) ExportersCorePulpCreateExecute(r ExportersPulpAPIExportersCorePulpCreateRequest) (*PulpExporterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PulpExporterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersPulpAPIService.ExportersCorePulpCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/exporters/core/pulp/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pulpExporter == nil {
		return localVarReturnValue, nil, reportError("pulpExporter is required and must be specified")
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
	localVarPostBody = r.pulpExporter
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

type ExportersPulpAPIExportersCorePulpDeleteRequest struct {
	ctx context.Context
	ApiService *ExportersPulpAPIService
	pulpExporterHref string
}

func (r ExportersPulpAPIExportersCorePulpDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ExportersCorePulpDeleteExecute(r)
}

/*
ExportersCorePulpDelete Delete a pulp exporter

Trigger an asynchronous delete task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpExporterHref
 @return ExportersPulpAPIExportersCorePulpDeleteRequest
*/
func (a *ExportersPulpAPIService) ExportersCorePulpDelete(ctx context.Context, pulpExporterHref string) ExportersPulpAPIExportersCorePulpDeleteRequest {
	return ExportersPulpAPIExportersCorePulpDeleteRequest{
		ApiService: a,
		ctx: ctx,
		pulpExporterHref: pulpExporterHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ExportersPulpAPIService) ExportersCorePulpDeleteExecute(r ExportersPulpAPIExportersCorePulpDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersPulpAPIService.ExportersCorePulpDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{pulp_exporter_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_exporter_href"+"}", url.PathEscape(parameterValueToString(r.pulpExporterHref, "pulpExporterHref")), -1)
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

type ExportersPulpAPIExportersCorePulpListRequest struct {
	ctx context.Context
	ApiService *ExportersPulpAPIService
	pulpDomain string
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIexact *string
	nameIn *[]string
	nameIregex *string
	nameIstartswith *string
	nameRegex *string
	nameStartswith *string
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	q *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ExportersPulpAPIExportersCorePulpListRequest) Limit(limit int32) ExportersPulpAPIExportersCorePulpListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r ExportersPulpAPIExportersCorePulpListRequest) Name(name string) ExportersPulpAPIExportersCorePulpListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r ExportersPulpAPIExportersCorePulpListRequest) NameContains(nameContains string) ExportersPulpAPIExportersCorePulpListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r ExportersPulpAPIExportersCorePulpListRequest) NameIcontains(nameIcontains string) ExportersPulpAPIExportersCorePulpListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name matches value
func (r ExportersPulpAPIExportersCorePulpListRequest) NameIexact(nameIexact string) ExportersPulpAPIExportersCorePulpListRequest {
	r.nameIexact = &nameIexact
	return r
}

// Filter results where name is in a comma-separated list of values
func (r ExportersPulpAPIExportersCorePulpListRequest) NameIn(nameIn []string) ExportersPulpAPIExportersCorePulpListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name matches regex value
func (r ExportersPulpAPIExportersCorePulpListRequest) NameIregex(nameIregex string) ExportersPulpAPIExportersCorePulpListRequest {
	r.nameIregex = &nameIregex
	return r
}

// Filter results where name starts with value
func (r ExportersPulpAPIExportersCorePulpListRequest) NameIstartswith(nameIstartswith string) ExportersPulpAPIExportersCorePulpListRequest {
	r.nameIstartswith = &nameIstartswith
	return r
}

// Filter results where name matches regex value
func (r ExportersPulpAPIExportersCorePulpListRequest) NameRegex(nameRegex string) ExportersPulpAPIExportersCorePulpListRequest {
	r.nameRegex = &nameRegex
	return r
}

// Filter results where name starts with value
func (r ExportersPulpAPIExportersCorePulpListRequest) NameStartswith(nameStartswith string) ExportersPulpAPIExportersCorePulpListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r ExportersPulpAPIExportersCorePulpListRequest) Offset(offset int32) ExportersPulpAPIExportersCorePulpListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;path&#x60; - Path* &#x60;-path&#x60; - Path (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r ExportersPulpAPIExportersCorePulpListRequest) Ordering(ordering []string) ExportersPulpAPIExportersCorePulpListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ExportersPulpAPIExportersCorePulpListRequest) PulpHrefIn(pulpHrefIn []string) ExportersPulpAPIExportersCorePulpListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ExportersPulpAPIExportersCorePulpListRequest) PulpIdIn(pulpIdIn []string) ExportersPulpAPIExportersCorePulpListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r ExportersPulpAPIExportersCorePulpListRequest) Q(q string) ExportersPulpAPIExportersCorePulpListRequest {
	r.q = &q
	return r
}

// A list of fields to include in the response.
func (r ExportersPulpAPIExportersCorePulpListRequest) Fields(fields []string) ExportersPulpAPIExportersCorePulpListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ExportersPulpAPIExportersCorePulpListRequest) ExcludeFields(excludeFields []string) ExportersPulpAPIExportersCorePulpListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ExportersPulpAPIExportersCorePulpListRequest) Execute() (*PaginatedPulpExporterResponseList, *http.Response, error) {
	return r.ApiService.ExportersCorePulpListExecute(r)
}

/*
ExportersCorePulpList List pulp exporters

ViewSet for viewing PulpExporters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ExportersPulpAPIExportersCorePulpListRequest
*/
func (a *ExportersPulpAPIService) ExportersCorePulpList(ctx context.Context, pulpDomain string) ExportersPulpAPIExportersCorePulpListRequest {
	return ExportersPulpAPIExportersCorePulpListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedPulpExporterResponseList
func (a *ExportersPulpAPIService) ExportersCorePulpListExecute(r ExportersPulpAPIExportersCorePulpListRequest) (*PaginatedPulpExporterResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedPulpExporterResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersPulpAPIService.ExportersCorePulpList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/exporters/core/pulp/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.nameContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__contains", r.nameContains, "form", "")
	}
	if r.nameIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__icontains", r.nameIcontains, "form", "")
	}
	if r.nameIexact != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__iexact", r.nameIexact, "form", "")
	}
	if r.nameIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__in", r.nameIn, "form", "csv")
	}
	if r.nameIregex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__iregex", r.nameIregex, "form", "")
	}
	if r.nameIstartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__istartswith", r.nameIstartswith, "form", "")
	}
	if r.nameRegex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__regex", r.nameRegex, "form", "")
	}
	if r.nameStartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__startswith", r.nameStartswith, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "form", "csv")
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

type ExportersPulpAPIExportersCorePulpPartialUpdateRequest struct {
	ctx context.Context
	ApiService *ExportersPulpAPIService
	pulpExporterHref string
	patchedPulpExporter *PatchedPulpExporter
}

func (r ExportersPulpAPIExportersCorePulpPartialUpdateRequest) PatchedPulpExporter(patchedPulpExporter PatchedPulpExporter) ExportersPulpAPIExportersCorePulpPartialUpdateRequest {
	r.patchedPulpExporter = &patchedPulpExporter
	return r
}

func (r ExportersPulpAPIExportersCorePulpPartialUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ExportersCorePulpPartialUpdateExecute(r)
}

/*
ExportersCorePulpPartialUpdate Update a pulp exporter

Trigger an asynchronous partial update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpExporterHref
 @return ExportersPulpAPIExportersCorePulpPartialUpdateRequest
*/
func (a *ExportersPulpAPIService) ExportersCorePulpPartialUpdate(ctx context.Context, pulpExporterHref string) ExportersPulpAPIExportersCorePulpPartialUpdateRequest {
	return ExportersPulpAPIExportersCorePulpPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		pulpExporterHref: pulpExporterHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ExportersPulpAPIService) ExportersCorePulpPartialUpdateExecute(r ExportersPulpAPIExportersCorePulpPartialUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersPulpAPIService.ExportersCorePulpPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{pulp_exporter_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_exporter_href"+"}", url.PathEscape(parameterValueToString(r.pulpExporterHref, "pulpExporterHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedPulpExporter == nil {
		return localVarReturnValue, nil, reportError("patchedPulpExporter is required and must be specified")
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
	localVarPostBody = r.patchedPulpExporter
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

type ExportersPulpAPIExportersCorePulpReadRequest struct {
	ctx context.Context
	ApiService *ExportersPulpAPIService
	pulpExporterHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ExportersPulpAPIExportersCorePulpReadRequest) Fields(fields []string) ExportersPulpAPIExportersCorePulpReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ExportersPulpAPIExportersCorePulpReadRequest) ExcludeFields(excludeFields []string) ExportersPulpAPIExportersCorePulpReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ExportersPulpAPIExportersCorePulpReadRequest) Execute() (*PulpExporterResponse, *http.Response, error) {
	return r.ApiService.ExportersCorePulpReadExecute(r)
}

/*
ExportersCorePulpRead Inspect a pulp exporter

ViewSet for viewing PulpExporters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpExporterHref
 @return ExportersPulpAPIExportersCorePulpReadRequest
*/
func (a *ExportersPulpAPIService) ExportersCorePulpRead(ctx context.Context, pulpExporterHref string) ExportersPulpAPIExportersCorePulpReadRequest {
	return ExportersPulpAPIExportersCorePulpReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpExporterHref: pulpExporterHref,
	}
}

// Execute executes the request
//  @return PulpExporterResponse
func (a *ExportersPulpAPIService) ExportersCorePulpReadExecute(r ExportersPulpAPIExportersCorePulpReadRequest) (*PulpExporterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PulpExporterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersPulpAPIService.ExportersCorePulpRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{pulp_exporter_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_exporter_href"+"}", url.PathEscape(parameterValueToString(r.pulpExporterHref, "pulpExporterHref")), -1)
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

type ExportersPulpAPIExportersCorePulpUpdateRequest struct {
	ctx context.Context
	ApiService *ExportersPulpAPIService
	pulpExporterHref string
	pulpExporter *PulpExporter
}

func (r ExportersPulpAPIExportersCorePulpUpdateRequest) PulpExporter(pulpExporter PulpExporter) ExportersPulpAPIExportersCorePulpUpdateRequest {
	r.pulpExporter = &pulpExporter
	return r
}

func (r ExportersPulpAPIExportersCorePulpUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ExportersCorePulpUpdateExecute(r)
}

/*
ExportersCorePulpUpdate Update a pulp exporter

Trigger an asynchronous update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpExporterHref
 @return ExportersPulpAPIExportersCorePulpUpdateRequest
*/
func (a *ExportersPulpAPIService) ExportersCorePulpUpdate(ctx context.Context, pulpExporterHref string) ExportersPulpAPIExportersCorePulpUpdateRequest {
	return ExportersPulpAPIExportersCorePulpUpdateRequest{
		ApiService: a,
		ctx: ctx,
		pulpExporterHref: pulpExporterHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ExportersPulpAPIService) ExportersCorePulpUpdateExecute(r ExportersPulpAPIExportersCorePulpUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExportersPulpAPIService.ExportersCorePulpUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{pulp_exporter_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_exporter_href"+"}", url.PathEscape(parameterValueToString(r.pulpExporterHref, "pulpExporterHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pulpExporter == nil {
		return localVarReturnValue, nil, reportError("pulpExporter is required and must be specified")
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
	localVarPostBody = r.pulpExporter
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
