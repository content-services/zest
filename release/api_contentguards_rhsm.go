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


// ContentguardsRhsmApiService ContentguardsRhsmApi service
type ContentguardsRhsmApiService service

type ContentguardsRhsmApiContentguardsCertguardRhsmCreateRequest struct {
	ctx context.Context
	ApiService *ContentguardsRhsmApiService
	certguardRHSMCertGuard *CertguardRHSMCertGuard
}

func (r ContentguardsRhsmApiContentguardsCertguardRhsmCreateRequest) CertguardRHSMCertGuard(certguardRHSMCertGuard CertguardRHSMCertGuard) ContentguardsRhsmApiContentguardsCertguardRhsmCreateRequest {
	r.certguardRHSMCertGuard = &certguardRHSMCertGuard
	return r
}

func (r ContentguardsRhsmApiContentguardsCertguardRhsmCreateRequest) Execute() (*CertguardRHSMCertGuardResponse, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardRhsmCreateExecute(r)
}

/*
ContentguardsCertguardRhsmCreate Create a rhsm cert guard

RHSMCertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentguardsRhsmApiContentguardsCertguardRhsmCreateRequest
*/
func (a *ContentguardsRhsmApiService) ContentguardsCertguardRhsmCreate(ctx context.Context) ContentguardsRhsmApiContentguardsCertguardRhsmCreateRequest {
	return ContentguardsRhsmApiContentguardsCertguardRhsmCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CertguardRHSMCertGuardResponse
func (a *ContentguardsRhsmApiService) ContentguardsCertguardRhsmCreateExecute(r ContentguardsRhsmApiContentguardsCertguardRhsmCreateRequest) (*CertguardRHSMCertGuardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertguardRHSMCertGuardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsRhsmApiService.ContentguardsCertguardRhsmCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/contentguards/certguard/rhsm/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.certguardRHSMCertGuard == nil {
		return localVarReturnValue, nil, reportError("certguardRHSMCertGuard is required and must be specified")
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
	localVarPostBody = r.certguardRHSMCertGuard
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

type ContentguardsRhsmApiContentguardsCertguardRhsmDeleteRequest struct {
	ctx context.Context
	ApiService *ContentguardsRhsmApiService
	certguardRHSMCertGuardHref string
}

func (r ContentguardsRhsmApiContentguardsCertguardRhsmDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ContentguardsCertguardRhsmDeleteExecute(r)
}

/*
ContentguardsCertguardRhsmDelete Delete a rhsm cert guard

RHSMCertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certguardRHSMCertGuardHref
 @return ContentguardsRhsmApiContentguardsCertguardRhsmDeleteRequest
*/
func (a *ContentguardsRhsmApiService) ContentguardsCertguardRhsmDelete(ctx context.Context, certguardRHSMCertGuardHref string) ContentguardsRhsmApiContentguardsCertguardRhsmDeleteRequest {
	return ContentguardsRhsmApiContentguardsCertguardRhsmDeleteRequest{
		ApiService: a,
		ctx: ctx,
		certguardRHSMCertGuardHref: certguardRHSMCertGuardHref,
	}
}

// Execute executes the request
func (a *ContentguardsRhsmApiService) ContentguardsCertguardRhsmDeleteExecute(r ContentguardsRhsmApiContentguardsCertguardRhsmDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsRhsmApiService.ContentguardsCertguardRhsmDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{certguard_r_h_s_m_cert_guard_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"certguard_r_h_s_m_cert_guard_href"+"}", url.PathEscape(parameterValueToString(r.certguardRHSMCertGuardHref, "certguardRHSMCertGuardHref")), -1)
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

type ContentguardsRhsmApiContentguardsCertguardRhsmListRequest struct {
	ctx context.Context
	ApiService *ContentguardsRhsmApiService
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
func (r ContentguardsRhsmApiContentguardsCertguardRhsmListRequest) Limit(limit int32) ContentguardsRhsmApiContentguardsCertguardRhsmListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r ContentguardsRhsmApiContentguardsCertguardRhsmListRequest) Name(name string) ContentguardsRhsmApiContentguardsCertguardRhsmListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r ContentguardsRhsmApiContentguardsCertguardRhsmListRequest) NameContains(nameContains string) ContentguardsRhsmApiContentguardsCertguardRhsmListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r ContentguardsRhsmApiContentguardsCertguardRhsmListRequest) NameIcontains(nameIcontains string) ContentguardsRhsmApiContentguardsCertguardRhsmListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r ContentguardsRhsmApiContentguardsCertguardRhsmListRequest) NameIn(nameIn []string) ContentguardsRhsmApiContentguardsCertguardRhsmListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r ContentguardsRhsmApiContentguardsCertguardRhsmListRequest) NameStartswith(nameStartswith string) ContentguardsRhsmApiContentguardsCertguardRhsmListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r ContentguardsRhsmApiContentguardsCertguardRhsmListRequest) Offset(offset int32) ContentguardsRhsmApiContentguardsCertguardRhsmListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ContentguardsRhsmApiContentguardsCertguardRhsmListRequest) Ordering(ordering []string) ContentguardsRhsmApiContentguardsCertguardRhsmListRequest {
	r.ordering = &ordering
	return r
}

// A list of fields to include in the response.
func (r ContentguardsRhsmApiContentguardsCertguardRhsmListRequest) Fields(fields []string) ContentguardsRhsmApiContentguardsCertguardRhsmListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentguardsRhsmApiContentguardsCertguardRhsmListRequest) ExcludeFields(excludeFields []string) ContentguardsRhsmApiContentguardsCertguardRhsmListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentguardsRhsmApiContentguardsCertguardRhsmListRequest) Execute() (*PaginatedcertguardRHSMCertGuardResponseList, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardRhsmListExecute(r)
}

/*
ContentguardsCertguardRhsmList List rhsm cert guards

RHSMCertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentguardsRhsmApiContentguardsCertguardRhsmListRequest
*/
func (a *ContentguardsRhsmApiService) ContentguardsCertguardRhsmList(ctx context.Context) ContentguardsRhsmApiContentguardsCertguardRhsmListRequest {
	return ContentguardsRhsmApiContentguardsCertguardRhsmListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedcertguardRHSMCertGuardResponseList
func (a *ContentguardsRhsmApiService) ContentguardsCertguardRhsmListExecute(r ContentguardsRhsmApiContentguardsCertguardRhsmListRequest) (*PaginatedcertguardRHSMCertGuardResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedcertguardRHSMCertGuardResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsRhsmApiService.ContentguardsCertguardRhsmList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/contentguards/certguard/rhsm/"
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

type ContentguardsRhsmApiContentguardsCertguardRhsmPartialUpdateRequest struct {
	ctx context.Context
	ApiService *ContentguardsRhsmApiService
	certguardRHSMCertGuardHref string
	patchedcertguardRHSMCertGuard *PatchedcertguardRHSMCertGuard
}

func (r ContentguardsRhsmApiContentguardsCertguardRhsmPartialUpdateRequest) PatchedcertguardRHSMCertGuard(patchedcertguardRHSMCertGuard PatchedcertguardRHSMCertGuard) ContentguardsRhsmApiContentguardsCertguardRhsmPartialUpdateRequest {
	r.patchedcertguardRHSMCertGuard = &patchedcertguardRHSMCertGuard
	return r
}

func (r ContentguardsRhsmApiContentguardsCertguardRhsmPartialUpdateRequest) Execute() (*CertguardRHSMCertGuardResponse, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardRhsmPartialUpdateExecute(r)
}

/*
ContentguardsCertguardRhsmPartialUpdate Update a rhsm cert guard

RHSMCertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certguardRHSMCertGuardHref
 @return ContentguardsRhsmApiContentguardsCertguardRhsmPartialUpdateRequest
*/
func (a *ContentguardsRhsmApiService) ContentguardsCertguardRhsmPartialUpdate(ctx context.Context, certguardRHSMCertGuardHref string) ContentguardsRhsmApiContentguardsCertguardRhsmPartialUpdateRequest {
	return ContentguardsRhsmApiContentguardsCertguardRhsmPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		certguardRHSMCertGuardHref: certguardRHSMCertGuardHref,
	}
}

// Execute executes the request
//  @return CertguardRHSMCertGuardResponse
func (a *ContentguardsRhsmApiService) ContentguardsCertguardRhsmPartialUpdateExecute(r ContentguardsRhsmApiContentguardsCertguardRhsmPartialUpdateRequest) (*CertguardRHSMCertGuardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertguardRHSMCertGuardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsRhsmApiService.ContentguardsCertguardRhsmPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{certguard_r_h_s_m_cert_guard_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"certguard_r_h_s_m_cert_guard_href"+"}", url.PathEscape(parameterValueToString(r.certguardRHSMCertGuardHref, "certguardRHSMCertGuardHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedcertguardRHSMCertGuard == nil {
		return localVarReturnValue, nil, reportError("patchedcertguardRHSMCertGuard is required and must be specified")
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
	localVarPostBody = r.patchedcertguardRHSMCertGuard
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

type ContentguardsRhsmApiContentguardsCertguardRhsmReadRequest struct {
	ctx context.Context
	ApiService *ContentguardsRhsmApiService
	certguardRHSMCertGuardHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentguardsRhsmApiContentguardsCertguardRhsmReadRequest) Fields(fields []string) ContentguardsRhsmApiContentguardsCertguardRhsmReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentguardsRhsmApiContentguardsCertguardRhsmReadRequest) ExcludeFields(excludeFields []string) ContentguardsRhsmApiContentguardsCertguardRhsmReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentguardsRhsmApiContentguardsCertguardRhsmReadRequest) Execute() (*CertguardRHSMCertGuardResponse, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardRhsmReadExecute(r)
}

/*
ContentguardsCertguardRhsmRead Inspect a rhsm cert guard

RHSMCertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certguardRHSMCertGuardHref
 @return ContentguardsRhsmApiContentguardsCertguardRhsmReadRequest
*/
func (a *ContentguardsRhsmApiService) ContentguardsCertguardRhsmRead(ctx context.Context, certguardRHSMCertGuardHref string) ContentguardsRhsmApiContentguardsCertguardRhsmReadRequest {
	return ContentguardsRhsmApiContentguardsCertguardRhsmReadRequest{
		ApiService: a,
		ctx: ctx,
		certguardRHSMCertGuardHref: certguardRHSMCertGuardHref,
	}
}

// Execute executes the request
//  @return CertguardRHSMCertGuardResponse
func (a *ContentguardsRhsmApiService) ContentguardsCertguardRhsmReadExecute(r ContentguardsRhsmApiContentguardsCertguardRhsmReadRequest) (*CertguardRHSMCertGuardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertguardRHSMCertGuardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsRhsmApiService.ContentguardsCertguardRhsmRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{certguard_r_h_s_m_cert_guard_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"certguard_r_h_s_m_cert_guard_href"+"}", url.PathEscape(parameterValueToString(r.certguardRHSMCertGuardHref, "certguardRHSMCertGuardHref")), -1)
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

type ContentguardsRhsmApiContentguardsCertguardRhsmUpdateRequest struct {
	ctx context.Context
	ApiService *ContentguardsRhsmApiService
	certguardRHSMCertGuardHref string
	certguardRHSMCertGuard *CertguardRHSMCertGuard
}

func (r ContentguardsRhsmApiContentguardsCertguardRhsmUpdateRequest) CertguardRHSMCertGuard(certguardRHSMCertGuard CertguardRHSMCertGuard) ContentguardsRhsmApiContentguardsCertguardRhsmUpdateRequest {
	r.certguardRHSMCertGuard = &certguardRHSMCertGuard
	return r
}

func (r ContentguardsRhsmApiContentguardsCertguardRhsmUpdateRequest) Execute() (*CertguardRHSMCertGuardResponse, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardRhsmUpdateExecute(r)
}

/*
ContentguardsCertguardRhsmUpdate Update a rhsm cert guard

RHSMCertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certguardRHSMCertGuardHref
 @return ContentguardsRhsmApiContentguardsCertguardRhsmUpdateRequest
*/
func (a *ContentguardsRhsmApiService) ContentguardsCertguardRhsmUpdate(ctx context.Context, certguardRHSMCertGuardHref string) ContentguardsRhsmApiContentguardsCertguardRhsmUpdateRequest {
	return ContentguardsRhsmApiContentguardsCertguardRhsmUpdateRequest{
		ApiService: a,
		ctx: ctx,
		certguardRHSMCertGuardHref: certguardRHSMCertGuardHref,
	}
}

// Execute executes the request
//  @return CertguardRHSMCertGuardResponse
func (a *ContentguardsRhsmApiService) ContentguardsCertguardRhsmUpdateExecute(r ContentguardsRhsmApiContentguardsCertguardRhsmUpdateRequest) (*CertguardRHSMCertGuardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertguardRHSMCertGuardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsRhsmApiService.ContentguardsCertguardRhsmUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{certguard_r_h_s_m_cert_guard_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"certguard_r_h_s_m_cert_guard_href"+"}", url.PathEscape(parameterValueToString(r.certguardRHSMCertGuardHref, "certguardRHSMCertGuardHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.certguardRHSMCertGuard == nil {
		return localVarReturnValue, nil, reportError("certguardRHSMCertGuard is required and must be specified")
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
	localVarPostBody = r.certguardRHSMCertGuard
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