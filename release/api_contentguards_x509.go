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


// ContentguardsX509ApiService ContentguardsX509Api service
type ContentguardsX509ApiService service

type ContentguardsX509ApiContentguardsCertguardX509CreateRequest struct {
	ctx context.Context
	ApiService *ContentguardsX509ApiService
	certguardX509CertGuard *CertguardX509CertGuard
}

func (r ContentguardsX509ApiContentguardsCertguardX509CreateRequest) CertguardX509CertGuard(certguardX509CertGuard CertguardX509CertGuard) ContentguardsX509ApiContentguardsCertguardX509CreateRequest {
	r.certguardX509CertGuard = &certguardX509CertGuard
	return r
}

func (r ContentguardsX509ApiContentguardsCertguardX509CreateRequest) Execute() (*CertguardX509CertGuardResponse, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardX509CreateExecute(r)
}

/*
ContentguardsCertguardX509Create Create a x509 cert guard

X509CertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentguardsX509ApiContentguardsCertguardX509CreateRequest
*/
func (a *ContentguardsX509ApiService) ContentguardsCertguardX509Create(ctx context.Context) ContentguardsX509ApiContentguardsCertguardX509CreateRequest {
	return ContentguardsX509ApiContentguardsCertguardX509CreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CertguardX509CertGuardResponse
func (a *ContentguardsX509ApiService) ContentguardsCertguardX509CreateExecute(r ContentguardsX509ApiContentguardsCertguardX509CreateRequest) (*CertguardX509CertGuardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertguardX509CertGuardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsX509ApiService.ContentguardsCertguardX509Create")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/contentguards/certguard/x509/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.certguardX509CertGuard == nil {
		return localVarReturnValue, nil, reportError("certguardX509CertGuard is required and must be specified")
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
	localVarPostBody = r.certguardX509CertGuard
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

type ContentguardsX509ApiContentguardsCertguardX509DeleteRequest struct {
	ctx context.Context
	ApiService *ContentguardsX509ApiService
	certguardX509CertGuardHref string
}

func (r ContentguardsX509ApiContentguardsCertguardX509DeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ContentguardsCertguardX509DeleteExecute(r)
}

/*
ContentguardsCertguardX509Delete Delete a x509 cert guard

X509CertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certguardX509CertGuardHref
 @return ContentguardsX509ApiContentguardsCertguardX509DeleteRequest
*/
func (a *ContentguardsX509ApiService) ContentguardsCertguardX509Delete(ctx context.Context, certguardX509CertGuardHref string) ContentguardsX509ApiContentguardsCertguardX509DeleteRequest {
	return ContentguardsX509ApiContentguardsCertguardX509DeleteRequest{
		ApiService: a,
		ctx: ctx,
		certguardX509CertGuardHref: certguardX509CertGuardHref,
	}
}

// Execute executes the request
func (a *ContentguardsX509ApiService) ContentguardsCertguardX509DeleteExecute(r ContentguardsX509ApiContentguardsCertguardX509DeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsX509ApiService.ContentguardsCertguardX509Delete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{certguard_x509_cert_guard_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"certguard_x509_cert_guard_href"+"}", url.PathEscape(parameterValueToString(r.certguardX509CertGuardHref, "certguardX509CertGuardHref")), -1)
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

type ContentguardsX509ApiContentguardsCertguardX509ListRequest struct {
	ctx context.Context
	ApiService *ContentguardsX509ApiService
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
func (r ContentguardsX509ApiContentguardsCertguardX509ListRequest) Limit(limit int32) ContentguardsX509ApiContentguardsCertguardX509ListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r ContentguardsX509ApiContentguardsCertguardX509ListRequest) Name(name string) ContentguardsX509ApiContentguardsCertguardX509ListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r ContentguardsX509ApiContentguardsCertguardX509ListRequest) NameContains(nameContains string) ContentguardsX509ApiContentguardsCertguardX509ListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r ContentguardsX509ApiContentguardsCertguardX509ListRequest) NameIcontains(nameIcontains string) ContentguardsX509ApiContentguardsCertguardX509ListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r ContentguardsX509ApiContentguardsCertguardX509ListRequest) NameIn(nameIn []string) ContentguardsX509ApiContentguardsCertguardX509ListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r ContentguardsX509ApiContentguardsCertguardX509ListRequest) NameStartswith(nameStartswith string) ContentguardsX509ApiContentguardsCertguardX509ListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r ContentguardsX509ApiContentguardsCertguardX509ListRequest) Offset(offset int32) ContentguardsX509ApiContentguardsCertguardX509ListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ContentguardsX509ApiContentguardsCertguardX509ListRequest) Ordering(ordering []string) ContentguardsX509ApiContentguardsCertguardX509ListRequest {
	r.ordering = &ordering
	return r
}

// A list of fields to include in the response.
func (r ContentguardsX509ApiContentguardsCertguardX509ListRequest) Fields(fields []string) ContentguardsX509ApiContentguardsCertguardX509ListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentguardsX509ApiContentguardsCertguardX509ListRequest) ExcludeFields(excludeFields []string) ContentguardsX509ApiContentguardsCertguardX509ListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentguardsX509ApiContentguardsCertguardX509ListRequest) Execute() (*PaginatedcertguardX509CertGuardResponseList, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardX509ListExecute(r)
}

/*
ContentguardsCertguardX509List List x509 cert guards

X509CertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentguardsX509ApiContentguardsCertguardX509ListRequest
*/
func (a *ContentguardsX509ApiService) ContentguardsCertguardX509List(ctx context.Context) ContentguardsX509ApiContentguardsCertguardX509ListRequest {
	return ContentguardsX509ApiContentguardsCertguardX509ListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedcertguardX509CertGuardResponseList
func (a *ContentguardsX509ApiService) ContentguardsCertguardX509ListExecute(r ContentguardsX509ApiContentguardsCertguardX509ListRequest) (*PaginatedcertguardX509CertGuardResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedcertguardX509CertGuardResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsX509ApiService.ContentguardsCertguardX509List")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/contentguards/certguard/x509/"
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

type ContentguardsX509ApiContentguardsCertguardX509PartialUpdateRequest struct {
	ctx context.Context
	ApiService *ContentguardsX509ApiService
	certguardX509CertGuardHref string
	patchedcertguardX509CertGuard *PatchedcertguardX509CertGuard
}

func (r ContentguardsX509ApiContentguardsCertguardX509PartialUpdateRequest) PatchedcertguardX509CertGuard(patchedcertguardX509CertGuard PatchedcertguardX509CertGuard) ContentguardsX509ApiContentguardsCertguardX509PartialUpdateRequest {
	r.patchedcertguardX509CertGuard = &patchedcertguardX509CertGuard
	return r
}

func (r ContentguardsX509ApiContentguardsCertguardX509PartialUpdateRequest) Execute() (*CertguardX509CertGuardResponse, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardX509PartialUpdateExecute(r)
}

/*
ContentguardsCertguardX509PartialUpdate Update a x509 cert guard

X509CertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certguardX509CertGuardHref
 @return ContentguardsX509ApiContentguardsCertguardX509PartialUpdateRequest
*/
func (a *ContentguardsX509ApiService) ContentguardsCertguardX509PartialUpdate(ctx context.Context, certguardX509CertGuardHref string) ContentguardsX509ApiContentguardsCertguardX509PartialUpdateRequest {
	return ContentguardsX509ApiContentguardsCertguardX509PartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		certguardX509CertGuardHref: certguardX509CertGuardHref,
	}
}

// Execute executes the request
//  @return CertguardX509CertGuardResponse
func (a *ContentguardsX509ApiService) ContentguardsCertguardX509PartialUpdateExecute(r ContentguardsX509ApiContentguardsCertguardX509PartialUpdateRequest) (*CertguardX509CertGuardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertguardX509CertGuardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsX509ApiService.ContentguardsCertguardX509PartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{certguard_x509_cert_guard_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"certguard_x509_cert_guard_href"+"}", url.PathEscape(parameterValueToString(r.certguardX509CertGuardHref, "certguardX509CertGuardHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedcertguardX509CertGuard == nil {
		return localVarReturnValue, nil, reportError("patchedcertguardX509CertGuard is required and must be specified")
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
	localVarPostBody = r.patchedcertguardX509CertGuard
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

type ContentguardsX509ApiContentguardsCertguardX509ReadRequest struct {
	ctx context.Context
	ApiService *ContentguardsX509ApiService
	certguardX509CertGuardHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentguardsX509ApiContentguardsCertguardX509ReadRequest) Fields(fields []string) ContentguardsX509ApiContentguardsCertguardX509ReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentguardsX509ApiContentguardsCertguardX509ReadRequest) ExcludeFields(excludeFields []string) ContentguardsX509ApiContentguardsCertguardX509ReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentguardsX509ApiContentguardsCertguardX509ReadRequest) Execute() (*CertguardX509CertGuardResponse, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardX509ReadExecute(r)
}

/*
ContentguardsCertguardX509Read Inspect a x509 cert guard

X509CertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certguardX509CertGuardHref
 @return ContentguardsX509ApiContentguardsCertguardX509ReadRequest
*/
func (a *ContentguardsX509ApiService) ContentguardsCertguardX509Read(ctx context.Context, certguardX509CertGuardHref string) ContentguardsX509ApiContentguardsCertguardX509ReadRequest {
	return ContentguardsX509ApiContentguardsCertguardX509ReadRequest{
		ApiService: a,
		ctx: ctx,
		certguardX509CertGuardHref: certguardX509CertGuardHref,
	}
}

// Execute executes the request
//  @return CertguardX509CertGuardResponse
func (a *ContentguardsX509ApiService) ContentguardsCertguardX509ReadExecute(r ContentguardsX509ApiContentguardsCertguardX509ReadRequest) (*CertguardX509CertGuardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertguardX509CertGuardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsX509ApiService.ContentguardsCertguardX509Read")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{certguard_x509_cert_guard_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"certguard_x509_cert_guard_href"+"}", url.PathEscape(parameterValueToString(r.certguardX509CertGuardHref, "certguardX509CertGuardHref")), -1)
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

type ContentguardsX509ApiContentguardsCertguardX509UpdateRequest struct {
	ctx context.Context
	ApiService *ContentguardsX509ApiService
	certguardX509CertGuardHref string
	certguardX509CertGuard *CertguardX509CertGuard
}

func (r ContentguardsX509ApiContentguardsCertguardX509UpdateRequest) CertguardX509CertGuard(certguardX509CertGuard CertguardX509CertGuard) ContentguardsX509ApiContentguardsCertguardX509UpdateRequest {
	r.certguardX509CertGuard = &certguardX509CertGuard
	return r
}

func (r ContentguardsX509ApiContentguardsCertguardX509UpdateRequest) Execute() (*CertguardX509CertGuardResponse, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardX509UpdateExecute(r)
}

/*
ContentguardsCertguardX509Update Update a x509 cert guard

X509CertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certguardX509CertGuardHref
 @return ContentguardsX509ApiContentguardsCertguardX509UpdateRequest
*/
func (a *ContentguardsX509ApiService) ContentguardsCertguardX509Update(ctx context.Context, certguardX509CertGuardHref string) ContentguardsX509ApiContentguardsCertguardX509UpdateRequest {
	return ContentguardsX509ApiContentguardsCertguardX509UpdateRequest{
		ApiService: a,
		ctx: ctx,
		certguardX509CertGuardHref: certguardX509CertGuardHref,
	}
}

// Execute executes the request
//  @return CertguardX509CertGuardResponse
func (a *ContentguardsX509ApiService) ContentguardsCertguardX509UpdateExecute(r ContentguardsX509ApiContentguardsCertguardX509UpdateRequest) (*CertguardX509CertGuardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertguardX509CertGuardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsX509ApiService.ContentguardsCertguardX509Update")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{certguard_x509_cert_guard_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"certguard_x509_cert_guard_href"+"}", url.PathEscape(parameterValueToString(r.certguardX509CertGuardHref, "certguardX509CertGuardHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.certguardX509CertGuard == nil {
		return localVarReturnValue, nil, reportError("certguardX509CertGuard is required and must be specified")
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
	localVarPostBody = r.certguardX509CertGuard
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
