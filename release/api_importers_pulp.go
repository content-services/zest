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


// ImportersPulpAPIService ImportersPulpAPI service
type ImportersPulpAPIService service

type ImportersPulpAPIImportersCorePulpCreateRequest struct {
	ctx context.Context
	ApiService *ImportersPulpAPIService
	pulpDomain string
	pulpImporter *PulpImporter
}

func (r ImportersPulpAPIImportersCorePulpCreateRequest) PulpImporter(pulpImporter PulpImporter) ImportersPulpAPIImportersCorePulpCreateRequest {
	r.pulpImporter = &pulpImporter
	return r
}

func (r ImportersPulpAPIImportersCorePulpCreateRequest) Execute() (*PulpImporterResponse, *http.Response, error) {
	return r.ApiService.ImportersCorePulpCreateExecute(r)
}

/*
ImportersCorePulpCreate Create a pulp importer

ViewSet for PulpImporters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ImportersPulpAPIImportersCorePulpCreateRequest
*/
func (a *ImportersPulpAPIService) ImportersCorePulpCreate(ctx context.Context, pulpDomain string) ImportersPulpAPIImportersCorePulpCreateRequest {
	return ImportersPulpAPIImportersCorePulpCreateRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PulpImporterResponse
func (a *ImportersPulpAPIService) ImportersCorePulpCreateExecute(r ImportersPulpAPIImportersCorePulpCreateRequest) (*PulpImporterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PulpImporterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpAPIService.ImportersCorePulpCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/importers/core/pulp/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pulpImporter == nil {
		return localVarReturnValue, nil, reportError("pulpImporter is required and must be specified")
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
	localVarPostBody = r.pulpImporter
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

type ImportersPulpAPIImportersCorePulpDeleteRequest struct {
	ctx context.Context
	ApiService *ImportersPulpAPIService
	pulpImporterHref string
}

func (r ImportersPulpAPIImportersCorePulpDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ImportersCorePulpDeleteExecute(r)
}

/*
ImportersCorePulpDelete Delete a pulp importer

ViewSet for PulpImporters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpImporterHref
 @return ImportersPulpAPIImportersCorePulpDeleteRequest
*/
func (a *ImportersPulpAPIService) ImportersCorePulpDelete(ctx context.Context, pulpImporterHref string) ImportersPulpAPIImportersCorePulpDeleteRequest {
	return ImportersPulpAPIImportersCorePulpDeleteRequest{
		ApiService: a,
		ctx: ctx,
		pulpImporterHref: pulpImporterHref,
	}
}

// Execute executes the request
func (a *ImportersPulpAPIService) ImportersCorePulpDeleteExecute(r ImportersPulpAPIImportersCorePulpDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpAPIService.ImportersCorePulpDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{pulp_importer_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_importer_href"+"}", url.PathEscape(parameterValueToString(r.pulpImporterHref, "pulpImporterHref")), -1)
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

type ImportersPulpAPIImportersCorePulpListRequest struct {
	ctx context.Context
	ApiService *ImportersPulpAPIService
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
func (r ImportersPulpAPIImportersCorePulpListRequest) Limit(limit int32) ImportersPulpAPIImportersCorePulpListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r ImportersPulpAPIImportersCorePulpListRequest) Name(name string) ImportersPulpAPIImportersCorePulpListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r ImportersPulpAPIImportersCorePulpListRequest) NameContains(nameContains string) ImportersPulpAPIImportersCorePulpListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r ImportersPulpAPIImportersCorePulpListRequest) NameIcontains(nameIcontains string) ImportersPulpAPIImportersCorePulpListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name matches value
func (r ImportersPulpAPIImportersCorePulpListRequest) NameIexact(nameIexact string) ImportersPulpAPIImportersCorePulpListRequest {
	r.nameIexact = &nameIexact
	return r
}

// Filter results where name is in a comma-separated list of values
func (r ImportersPulpAPIImportersCorePulpListRequest) NameIn(nameIn []string) ImportersPulpAPIImportersCorePulpListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name matches regex value
func (r ImportersPulpAPIImportersCorePulpListRequest) NameIregex(nameIregex string) ImportersPulpAPIImportersCorePulpListRequest {
	r.nameIregex = &nameIregex
	return r
}

// Filter results where name starts with value
func (r ImportersPulpAPIImportersCorePulpListRequest) NameIstartswith(nameIstartswith string) ImportersPulpAPIImportersCorePulpListRequest {
	r.nameIstartswith = &nameIstartswith
	return r
}

// Filter results where name matches regex value
func (r ImportersPulpAPIImportersCorePulpListRequest) NameRegex(nameRegex string) ImportersPulpAPIImportersCorePulpListRequest {
	r.nameRegex = &nameRegex
	return r
}

// Filter results where name starts with value
func (r ImportersPulpAPIImportersCorePulpListRequest) NameStartswith(nameStartswith string) ImportersPulpAPIImportersCorePulpListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r ImportersPulpAPIImportersCorePulpListRequest) Offset(offset int32) ImportersPulpAPIImportersCorePulpListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r ImportersPulpAPIImportersCorePulpListRequest) Ordering(ordering []string) ImportersPulpAPIImportersCorePulpListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ImportersPulpAPIImportersCorePulpListRequest) PulpHrefIn(pulpHrefIn []string) ImportersPulpAPIImportersCorePulpListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ImportersPulpAPIImportersCorePulpListRequest) PulpIdIn(pulpIdIn []string) ImportersPulpAPIImportersCorePulpListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter results by using NOT, AND and OR operations on other filters
func (r ImportersPulpAPIImportersCorePulpListRequest) Q(q string) ImportersPulpAPIImportersCorePulpListRequest {
	r.q = &q
	return r
}

// A list of fields to include in the response.
func (r ImportersPulpAPIImportersCorePulpListRequest) Fields(fields []string) ImportersPulpAPIImportersCorePulpListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ImportersPulpAPIImportersCorePulpListRequest) ExcludeFields(excludeFields []string) ImportersPulpAPIImportersCorePulpListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ImportersPulpAPIImportersCorePulpListRequest) Execute() (*PaginatedPulpImporterResponseList, *http.Response, error) {
	return r.ApiService.ImportersCorePulpListExecute(r)
}

/*
ImportersCorePulpList List pulp importers

ViewSet for PulpImporters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ImportersPulpAPIImportersCorePulpListRequest
*/
func (a *ImportersPulpAPIService) ImportersCorePulpList(ctx context.Context, pulpDomain string) ImportersPulpAPIImportersCorePulpListRequest {
	return ImportersPulpAPIImportersCorePulpListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedPulpImporterResponseList
func (a *ImportersPulpAPIService) ImportersCorePulpListExecute(r ImportersPulpAPIImportersCorePulpListRequest) (*PaginatedPulpImporterResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedPulpImporterResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpAPIService.ImportersCorePulpList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/importers/core/pulp/"
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

type ImportersPulpAPIImportersCorePulpPartialUpdateRequest struct {
	ctx context.Context
	ApiService *ImportersPulpAPIService
	pulpImporterHref string
	patchedPulpImporter *PatchedPulpImporter
}

func (r ImportersPulpAPIImportersCorePulpPartialUpdateRequest) PatchedPulpImporter(patchedPulpImporter PatchedPulpImporter) ImportersPulpAPIImportersCorePulpPartialUpdateRequest {
	r.patchedPulpImporter = &patchedPulpImporter
	return r
}

func (r ImportersPulpAPIImportersCorePulpPartialUpdateRequest) Execute() (*PulpImporterResponse, *http.Response, error) {
	return r.ApiService.ImportersCorePulpPartialUpdateExecute(r)
}

/*
ImportersCorePulpPartialUpdate Update a pulp importer

ViewSet for PulpImporters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpImporterHref
 @return ImportersPulpAPIImportersCorePulpPartialUpdateRequest
*/
func (a *ImportersPulpAPIService) ImportersCorePulpPartialUpdate(ctx context.Context, pulpImporterHref string) ImportersPulpAPIImportersCorePulpPartialUpdateRequest {
	return ImportersPulpAPIImportersCorePulpPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		pulpImporterHref: pulpImporterHref,
	}
}

// Execute executes the request
//  @return PulpImporterResponse
func (a *ImportersPulpAPIService) ImportersCorePulpPartialUpdateExecute(r ImportersPulpAPIImportersCorePulpPartialUpdateRequest) (*PulpImporterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PulpImporterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpAPIService.ImportersCorePulpPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{pulp_importer_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_importer_href"+"}", url.PathEscape(parameterValueToString(r.pulpImporterHref, "pulpImporterHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedPulpImporter == nil {
		return localVarReturnValue, nil, reportError("patchedPulpImporter is required and must be specified")
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
	localVarPostBody = r.patchedPulpImporter
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

type ImportersPulpAPIImportersCorePulpReadRequest struct {
	ctx context.Context
	ApiService *ImportersPulpAPIService
	pulpImporterHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ImportersPulpAPIImportersCorePulpReadRequest) Fields(fields []string) ImportersPulpAPIImportersCorePulpReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ImportersPulpAPIImportersCorePulpReadRequest) ExcludeFields(excludeFields []string) ImportersPulpAPIImportersCorePulpReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ImportersPulpAPIImportersCorePulpReadRequest) Execute() (*PulpImporterResponse, *http.Response, error) {
	return r.ApiService.ImportersCorePulpReadExecute(r)
}

/*
ImportersCorePulpRead Inspect a pulp importer

ViewSet for PulpImporters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpImporterHref
 @return ImportersPulpAPIImportersCorePulpReadRequest
*/
func (a *ImportersPulpAPIService) ImportersCorePulpRead(ctx context.Context, pulpImporterHref string) ImportersPulpAPIImportersCorePulpReadRequest {
	return ImportersPulpAPIImportersCorePulpReadRequest{
		ApiService: a,
		ctx: ctx,
		pulpImporterHref: pulpImporterHref,
	}
}

// Execute executes the request
//  @return PulpImporterResponse
func (a *ImportersPulpAPIService) ImportersCorePulpReadExecute(r ImportersPulpAPIImportersCorePulpReadRequest) (*PulpImporterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PulpImporterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpAPIService.ImportersCorePulpRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{pulp_importer_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_importer_href"+"}", url.PathEscape(parameterValueToString(r.pulpImporterHref, "pulpImporterHref")), -1)
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

type ImportersPulpAPIImportersCorePulpUpdateRequest struct {
	ctx context.Context
	ApiService *ImportersPulpAPIService
	pulpImporterHref string
	pulpImporter *PulpImporter
}

func (r ImportersPulpAPIImportersCorePulpUpdateRequest) PulpImporter(pulpImporter PulpImporter) ImportersPulpAPIImportersCorePulpUpdateRequest {
	r.pulpImporter = &pulpImporter
	return r
}

func (r ImportersPulpAPIImportersCorePulpUpdateRequest) Execute() (*PulpImporterResponse, *http.Response, error) {
	return r.ApiService.ImportersCorePulpUpdateExecute(r)
}

/*
ImportersCorePulpUpdate Update a pulp importer

ViewSet for PulpImporters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpImporterHref
 @return ImportersPulpAPIImportersCorePulpUpdateRequest
*/
func (a *ImportersPulpAPIService) ImportersCorePulpUpdate(ctx context.Context, pulpImporterHref string) ImportersPulpAPIImportersCorePulpUpdateRequest {
	return ImportersPulpAPIImportersCorePulpUpdateRequest{
		ApiService: a,
		ctx: ctx,
		pulpImporterHref: pulpImporterHref,
	}
}

// Execute executes the request
//  @return PulpImporterResponse
func (a *ImportersPulpAPIService) ImportersCorePulpUpdateExecute(r ImportersPulpAPIImportersCorePulpUpdateRequest) (*PulpImporterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PulpImporterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpAPIService.ImportersCorePulpUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{pulp_importer_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_importer_href"+"}", url.PathEscape(parameterValueToString(r.pulpImporterHref, "pulpImporterHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pulpImporter == nil {
		return localVarReturnValue, nil, reportError("pulpImporter is required and must be specified")
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
	localVarPostBody = r.pulpImporter
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
