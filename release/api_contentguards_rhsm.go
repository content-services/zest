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


// ContentguardsRhsmAPIService ContentguardsRhsmAPI service
type ContentguardsRhsmAPIService service

type ContentguardsRhsmAPIContentguardsCertguardRhsmCreateRequest struct {
	ctx context.Context
	ApiService *ContentguardsRhsmAPIService
	pulpDomain string
	certguardRHSMCertGuard *CertguardRHSMCertGuard
}

func (r ContentguardsRhsmAPIContentguardsCertguardRhsmCreateRequest) CertguardRHSMCertGuard(certguardRHSMCertGuard CertguardRHSMCertGuard) ContentguardsRhsmAPIContentguardsCertguardRhsmCreateRequest {
	r.certguardRHSMCertGuard = &certguardRHSMCertGuard
	return r
}

func (r ContentguardsRhsmAPIContentguardsCertguardRhsmCreateRequest) Execute() (*CertguardRHSMCertGuardResponse, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardRhsmCreateExecute(r)
}

/*
ContentguardsCertguardRhsmCreate Create a rhsm cert guard

RHSMCertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentguardsRhsmAPIContentguardsCertguardRhsmCreateRequest
*/
func (a *ContentguardsRhsmAPIService) ContentguardsCertguardRhsmCreate(ctx context.Context, pulpDomain string) ContentguardsRhsmAPIContentguardsCertguardRhsmCreateRequest {
	return ContentguardsRhsmAPIContentguardsCertguardRhsmCreateRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return CertguardRHSMCertGuardResponse
func (a *ContentguardsRhsmAPIService) ContentguardsCertguardRhsmCreateExecute(r ContentguardsRhsmAPIContentguardsCertguardRhsmCreateRequest) (*CertguardRHSMCertGuardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertguardRHSMCertGuardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsRhsmAPIService.ContentguardsCertguardRhsmCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/contentguards/certguard/rhsm/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
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

type ContentguardsRhsmAPIContentguardsCertguardRhsmDeleteRequest struct {
	ctx context.Context
	ApiService *ContentguardsRhsmAPIService
	certguardRHSMCertGuardHref string
}

func (r ContentguardsRhsmAPIContentguardsCertguardRhsmDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ContentguardsCertguardRhsmDeleteExecute(r)
}

/*
ContentguardsCertguardRhsmDelete Delete a rhsm cert guard

RHSMCertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certguardRHSMCertGuardHref
 @return ContentguardsRhsmAPIContentguardsCertguardRhsmDeleteRequest
*/
func (a *ContentguardsRhsmAPIService) ContentguardsCertguardRhsmDelete(ctx context.Context, certguardRHSMCertGuardHref string) ContentguardsRhsmAPIContentguardsCertguardRhsmDeleteRequest {
	return ContentguardsRhsmAPIContentguardsCertguardRhsmDeleteRequest{
		ApiService: a,
		ctx: ctx,
		certguardRHSMCertGuardHref: certguardRHSMCertGuardHref,
	}
}

// Execute executes the request
func (a *ContentguardsRhsmAPIService) ContentguardsCertguardRhsmDeleteExecute(r ContentguardsRhsmAPIContentguardsCertguardRhsmDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsRhsmAPIService.ContentguardsCertguardRhsmDelete")
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

type ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest struct {
	ctx context.Context
	ApiService *ContentguardsRhsmAPIService
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
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) Limit(limit int32) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) Name(name string) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) NameContains(nameContains string) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) NameIcontains(nameIcontains string) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name matches value
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) NameIexact(nameIexact string) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.nameIexact = &nameIexact
	return r
}

// Filter results where name is in a comma-separated list of values
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) NameIn(nameIn []string) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name matches regex value
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) NameIregex(nameIregex string) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.nameIregex = &nameIregex
	return r
}

// Filter results where name starts with value
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) NameIstartswith(nameIstartswith string) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.nameIstartswith = &nameIstartswith
	return r
}

// Filter results where name matches regex value
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) NameRegex(nameRegex string) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.nameRegex = &nameRegex
	return r
}

// Filter results where name starts with value
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) NameStartswith(nameStartswith string) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) Offset(offset int32) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;description&#x60; - Description* &#x60;-description&#x60; - Description (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) Ordering(ordering []string) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) PulpHrefIn(pulpHrefIn []string) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) PulpIdIn(pulpIdIn []string) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter results by using NOT, AND and OR operations on other filters
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) Q(q string) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.q = &q
	return r
}

// A list of fields to include in the response.
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) Fields(fields []string) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) ExcludeFields(excludeFields []string) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) Execute() (*PaginatedcertguardRHSMCertGuardResponseList, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardRhsmListExecute(r)
}

/*
ContentguardsCertguardRhsmList List rhsm cert guards

RHSMCertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest
*/
func (a *ContentguardsRhsmAPIService) ContentguardsCertguardRhsmList(ctx context.Context, pulpDomain string) ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest {
	return ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedcertguardRHSMCertGuardResponseList
func (a *ContentguardsRhsmAPIService) ContentguardsCertguardRhsmListExecute(r ContentguardsRhsmAPIContentguardsCertguardRhsmListRequest) (*PaginatedcertguardRHSMCertGuardResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedcertguardRHSMCertGuardResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsRhsmAPIService.ContentguardsCertguardRhsmList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/contentguards/certguard/rhsm/"
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

type ContentguardsRhsmAPIContentguardsCertguardRhsmPartialUpdateRequest struct {
	ctx context.Context
	ApiService *ContentguardsRhsmAPIService
	certguardRHSMCertGuardHref string
	patchedcertguardRHSMCertGuard *PatchedcertguardRHSMCertGuard
}

func (r ContentguardsRhsmAPIContentguardsCertguardRhsmPartialUpdateRequest) PatchedcertguardRHSMCertGuard(patchedcertguardRHSMCertGuard PatchedcertguardRHSMCertGuard) ContentguardsRhsmAPIContentguardsCertguardRhsmPartialUpdateRequest {
	r.patchedcertguardRHSMCertGuard = &patchedcertguardRHSMCertGuard
	return r
}

func (r ContentguardsRhsmAPIContentguardsCertguardRhsmPartialUpdateRequest) Execute() (*CertguardRHSMCertGuardResponse, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardRhsmPartialUpdateExecute(r)
}

/*
ContentguardsCertguardRhsmPartialUpdate Update a rhsm cert guard

RHSMCertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certguardRHSMCertGuardHref
 @return ContentguardsRhsmAPIContentguardsCertguardRhsmPartialUpdateRequest
*/
func (a *ContentguardsRhsmAPIService) ContentguardsCertguardRhsmPartialUpdate(ctx context.Context, certguardRHSMCertGuardHref string) ContentguardsRhsmAPIContentguardsCertguardRhsmPartialUpdateRequest {
	return ContentguardsRhsmAPIContentguardsCertguardRhsmPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		certguardRHSMCertGuardHref: certguardRHSMCertGuardHref,
	}
}

// Execute executes the request
//  @return CertguardRHSMCertGuardResponse
func (a *ContentguardsRhsmAPIService) ContentguardsCertguardRhsmPartialUpdateExecute(r ContentguardsRhsmAPIContentguardsCertguardRhsmPartialUpdateRequest) (*CertguardRHSMCertGuardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertguardRHSMCertGuardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsRhsmAPIService.ContentguardsCertguardRhsmPartialUpdate")
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

type ContentguardsRhsmAPIContentguardsCertguardRhsmReadRequest struct {
	ctx context.Context
	ApiService *ContentguardsRhsmAPIService
	certguardRHSMCertGuardHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmReadRequest) Fields(fields []string) ContentguardsRhsmAPIContentguardsCertguardRhsmReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentguardsRhsmAPIContentguardsCertguardRhsmReadRequest) ExcludeFields(excludeFields []string) ContentguardsRhsmAPIContentguardsCertguardRhsmReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentguardsRhsmAPIContentguardsCertguardRhsmReadRequest) Execute() (*CertguardRHSMCertGuardResponse, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardRhsmReadExecute(r)
}

/*
ContentguardsCertguardRhsmRead Inspect a rhsm cert guard

RHSMCertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certguardRHSMCertGuardHref
 @return ContentguardsRhsmAPIContentguardsCertguardRhsmReadRequest
*/
func (a *ContentguardsRhsmAPIService) ContentguardsCertguardRhsmRead(ctx context.Context, certguardRHSMCertGuardHref string) ContentguardsRhsmAPIContentguardsCertguardRhsmReadRequest {
	return ContentguardsRhsmAPIContentguardsCertguardRhsmReadRequest{
		ApiService: a,
		ctx: ctx,
		certguardRHSMCertGuardHref: certguardRHSMCertGuardHref,
	}
}

// Execute executes the request
//  @return CertguardRHSMCertGuardResponse
func (a *ContentguardsRhsmAPIService) ContentguardsCertguardRhsmReadExecute(r ContentguardsRhsmAPIContentguardsCertguardRhsmReadRequest) (*CertguardRHSMCertGuardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertguardRHSMCertGuardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsRhsmAPIService.ContentguardsCertguardRhsmRead")
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

type ContentguardsRhsmAPIContentguardsCertguardRhsmUpdateRequest struct {
	ctx context.Context
	ApiService *ContentguardsRhsmAPIService
	certguardRHSMCertGuardHref string
	certguardRHSMCertGuard *CertguardRHSMCertGuard
}

func (r ContentguardsRhsmAPIContentguardsCertguardRhsmUpdateRequest) CertguardRHSMCertGuard(certguardRHSMCertGuard CertguardRHSMCertGuard) ContentguardsRhsmAPIContentguardsCertguardRhsmUpdateRequest {
	r.certguardRHSMCertGuard = &certguardRHSMCertGuard
	return r
}

func (r ContentguardsRhsmAPIContentguardsCertguardRhsmUpdateRequest) Execute() (*CertguardRHSMCertGuardResponse, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardRhsmUpdateExecute(r)
}

/*
ContentguardsCertguardRhsmUpdate Update a rhsm cert guard

RHSMCertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certguardRHSMCertGuardHref
 @return ContentguardsRhsmAPIContentguardsCertguardRhsmUpdateRequest
*/
func (a *ContentguardsRhsmAPIService) ContentguardsCertguardRhsmUpdate(ctx context.Context, certguardRHSMCertGuardHref string) ContentguardsRhsmAPIContentguardsCertguardRhsmUpdateRequest {
	return ContentguardsRhsmAPIContentguardsCertguardRhsmUpdateRequest{
		ApiService: a,
		ctx: ctx,
		certguardRHSMCertGuardHref: certguardRHSMCertGuardHref,
	}
}

// Execute executes the request
//  @return CertguardRHSMCertGuardResponse
func (a *ContentguardsRhsmAPIService) ContentguardsCertguardRhsmUpdateExecute(r ContentguardsRhsmAPIContentguardsCertguardRhsmUpdateRequest) (*CertguardRHSMCertGuardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertguardRHSMCertGuardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsRhsmAPIService.ContentguardsCertguardRhsmUpdate")
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
