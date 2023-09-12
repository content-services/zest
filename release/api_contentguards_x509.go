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


// ContentguardsX509APIService ContentguardsX509API service
type ContentguardsX509APIService service

type ContentguardsX509APIContentguardsCertguardX509CreateRequest struct {
	ctx context.Context
	ApiService *ContentguardsX509APIService
	pulpDomain string
	certguardX509CertGuard *CertguardX509CertGuard
}

func (r ContentguardsX509APIContentguardsCertguardX509CreateRequest) CertguardX509CertGuard(certguardX509CertGuard CertguardX509CertGuard) ContentguardsX509APIContentguardsCertguardX509CreateRequest {
	r.certguardX509CertGuard = &certguardX509CertGuard
	return r
}

func (r ContentguardsX509APIContentguardsCertguardX509CreateRequest) Execute() (*CertguardX509CertGuardResponse, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardX509CreateExecute(r)
}

/*
ContentguardsCertguardX509Create Create a x509 cert guard

X509CertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentguardsX509APIContentguardsCertguardX509CreateRequest
*/
func (a *ContentguardsX509APIService) ContentguardsCertguardX509Create(ctx context.Context, pulpDomain string) ContentguardsX509APIContentguardsCertguardX509CreateRequest {
	return ContentguardsX509APIContentguardsCertguardX509CreateRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return CertguardX509CertGuardResponse
func (a *ContentguardsX509APIService) ContentguardsCertguardX509CreateExecute(r ContentguardsX509APIContentguardsCertguardX509CreateRequest) (*CertguardX509CertGuardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertguardX509CertGuardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsX509APIService.ContentguardsCertguardX509Create")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/{pulp_domain}/api/v3/contentguards/certguard/x509/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
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

type ContentguardsX509APIContentguardsCertguardX509DeleteRequest struct {
	ctx context.Context
	ApiService *ContentguardsX509APIService
	certguardX509CertGuardHref string
}

func (r ContentguardsX509APIContentguardsCertguardX509DeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ContentguardsCertguardX509DeleteExecute(r)
}

/*
ContentguardsCertguardX509Delete Delete a x509 cert guard

X509CertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certguardX509CertGuardHref
 @return ContentguardsX509APIContentguardsCertguardX509DeleteRequest
*/
func (a *ContentguardsX509APIService) ContentguardsCertguardX509Delete(ctx context.Context, certguardX509CertGuardHref string) ContentguardsX509APIContentguardsCertguardX509DeleteRequest {
	return ContentguardsX509APIContentguardsCertguardX509DeleteRequest{
		ApiService: a,
		ctx: ctx,
		certguardX509CertGuardHref: certguardX509CertGuardHref,
	}
}

// Execute executes the request
func (a *ContentguardsX509APIService) ContentguardsCertguardX509DeleteExecute(r ContentguardsX509APIContentguardsCertguardX509DeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsX509APIService.ContentguardsCertguardX509Delete")
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

type ContentguardsX509APIContentguardsCertguardX509ListRequest struct {
	ctx context.Context
	ApiService *ContentguardsX509APIService
	pulpDomain string
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIexact *string
	nameIn *[]string
	nameIstartswith *string
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
func (r ContentguardsX509APIContentguardsCertguardX509ListRequest) Limit(limit int32) ContentguardsX509APIContentguardsCertguardX509ListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r ContentguardsX509APIContentguardsCertguardX509ListRequest) Name(name string) ContentguardsX509APIContentguardsCertguardX509ListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r ContentguardsX509APIContentguardsCertguardX509ListRequest) NameContains(nameContains string) ContentguardsX509APIContentguardsCertguardX509ListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r ContentguardsX509APIContentguardsCertguardX509ListRequest) NameIcontains(nameIcontains string) ContentguardsX509APIContentguardsCertguardX509ListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name matches value
func (r ContentguardsX509APIContentguardsCertguardX509ListRequest) NameIexact(nameIexact string) ContentguardsX509APIContentguardsCertguardX509ListRequest {
	r.nameIexact = &nameIexact
	return r
}

// Filter results where name is in a comma-separated list of values
func (r ContentguardsX509APIContentguardsCertguardX509ListRequest) NameIn(nameIn []string) ContentguardsX509APIContentguardsCertguardX509ListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r ContentguardsX509APIContentguardsCertguardX509ListRequest) NameIstartswith(nameIstartswith string) ContentguardsX509APIContentguardsCertguardX509ListRequest {
	r.nameIstartswith = &nameIstartswith
	return r
}

// Filter results where name starts with value
func (r ContentguardsX509APIContentguardsCertguardX509ListRequest) NameStartswith(nameStartswith string) ContentguardsX509APIContentguardsCertguardX509ListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r ContentguardsX509APIContentguardsCertguardX509ListRequest) Offset(offset int32) ContentguardsX509APIContentguardsCertguardX509ListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentguardsX509APIContentguardsCertguardX509ListRequest) Ordering(ordering []string) ContentguardsX509APIContentguardsCertguardX509ListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ContentguardsX509APIContentguardsCertguardX509ListRequest) PulpHrefIn(pulpHrefIn []string) ContentguardsX509APIContentguardsCertguardX509ListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentguardsX509APIContentguardsCertguardX509ListRequest) PulpIdIn(pulpIdIn []string) ContentguardsX509APIContentguardsCertguardX509ListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r ContentguardsX509APIContentguardsCertguardX509ListRequest) Q(q string) ContentguardsX509APIContentguardsCertguardX509ListRequest {
	r.q = &q
	return r
}

// A list of fields to include in the response.
func (r ContentguardsX509APIContentguardsCertguardX509ListRequest) Fields(fields []string) ContentguardsX509APIContentguardsCertguardX509ListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentguardsX509APIContentguardsCertguardX509ListRequest) ExcludeFields(excludeFields []string) ContentguardsX509APIContentguardsCertguardX509ListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentguardsX509APIContentguardsCertguardX509ListRequest) Execute() (*PaginatedcertguardX509CertGuardResponseList, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardX509ListExecute(r)
}

/*
ContentguardsCertguardX509List List x509 cert guards

X509CertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentguardsX509APIContentguardsCertguardX509ListRequest
*/
func (a *ContentguardsX509APIService) ContentguardsCertguardX509List(ctx context.Context, pulpDomain string) ContentguardsX509APIContentguardsCertguardX509ListRequest {
	return ContentguardsX509APIContentguardsCertguardX509ListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedcertguardX509CertGuardResponseList
func (a *ContentguardsX509APIService) ContentguardsCertguardX509ListExecute(r ContentguardsX509APIContentguardsCertguardX509ListRequest) (*PaginatedcertguardX509CertGuardResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedcertguardX509CertGuardResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsX509APIService.ContentguardsCertguardX509List")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/{pulp_domain}/api/v3/contentguards/certguard/x509/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

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
	if r.nameIexact != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__iexact", r.nameIexact, "")
	}
	if r.nameIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__in", r.nameIn, "csv")
	}
	if r.nameIstartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__istartswith", r.nameIstartswith, "")
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
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
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

type ContentguardsX509APIContentguardsCertguardX509PartialUpdateRequest struct {
	ctx context.Context
	ApiService *ContentguardsX509APIService
	certguardX509CertGuardHref string
	patchedcertguardX509CertGuard *PatchedcertguardX509CertGuard
}

func (r ContentguardsX509APIContentguardsCertguardX509PartialUpdateRequest) PatchedcertguardX509CertGuard(patchedcertguardX509CertGuard PatchedcertguardX509CertGuard) ContentguardsX509APIContentguardsCertguardX509PartialUpdateRequest {
	r.patchedcertguardX509CertGuard = &patchedcertguardX509CertGuard
	return r
}

func (r ContentguardsX509APIContentguardsCertguardX509PartialUpdateRequest) Execute() (*CertguardX509CertGuardResponse, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardX509PartialUpdateExecute(r)
}

/*
ContentguardsCertguardX509PartialUpdate Update a x509 cert guard

X509CertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certguardX509CertGuardHref
 @return ContentguardsX509APIContentguardsCertguardX509PartialUpdateRequest
*/
func (a *ContentguardsX509APIService) ContentguardsCertguardX509PartialUpdate(ctx context.Context, certguardX509CertGuardHref string) ContentguardsX509APIContentguardsCertguardX509PartialUpdateRequest {
	return ContentguardsX509APIContentguardsCertguardX509PartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		certguardX509CertGuardHref: certguardX509CertGuardHref,
	}
}

// Execute executes the request
//  @return CertguardX509CertGuardResponse
func (a *ContentguardsX509APIService) ContentguardsCertguardX509PartialUpdateExecute(r ContentguardsX509APIContentguardsCertguardX509PartialUpdateRequest) (*CertguardX509CertGuardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertguardX509CertGuardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsX509APIService.ContentguardsCertguardX509PartialUpdate")
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

type ContentguardsX509APIContentguardsCertguardX509ReadRequest struct {
	ctx context.Context
	ApiService *ContentguardsX509APIService
	certguardX509CertGuardHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentguardsX509APIContentguardsCertguardX509ReadRequest) Fields(fields []string) ContentguardsX509APIContentguardsCertguardX509ReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentguardsX509APIContentguardsCertguardX509ReadRequest) ExcludeFields(excludeFields []string) ContentguardsX509APIContentguardsCertguardX509ReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentguardsX509APIContentguardsCertguardX509ReadRequest) Execute() (*CertguardX509CertGuardResponse, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardX509ReadExecute(r)
}

/*
ContentguardsCertguardX509Read Inspect a x509 cert guard

X509CertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certguardX509CertGuardHref
 @return ContentguardsX509APIContentguardsCertguardX509ReadRequest
*/
func (a *ContentguardsX509APIService) ContentguardsCertguardX509Read(ctx context.Context, certguardX509CertGuardHref string) ContentguardsX509APIContentguardsCertguardX509ReadRequest {
	return ContentguardsX509APIContentguardsCertguardX509ReadRequest{
		ApiService: a,
		ctx: ctx,
		certguardX509CertGuardHref: certguardX509CertGuardHref,
	}
}

// Execute executes the request
//  @return CertguardX509CertGuardResponse
func (a *ContentguardsX509APIService) ContentguardsCertguardX509ReadExecute(r ContentguardsX509APIContentguardsCertguardX509ReadRequest) (*CertguardX509CertGuardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertguardX509CertGuardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsX509APIService.ContentguardsCertguardX509Read")
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

type ContentguardsX509APIContentguardsCertguardX509UpdateRequest struct {
	ctx context.Context
	ApiService *ContentguardsX509APIService
	certguardX509CertGuardHref string
	certguardX509CertGuard *CertguardX509CertGuard
}

func (r ContentguardsX509APIContentguardsCertguardX509UpdateRequest) CertguardX509CertGuard(certguardX509CertGuard CertguardX509CertGuard) ContentguardsX509APIContentguardsCertguardX509UpdateRequest {
	r.certguardX509CertGuard = &certguardX509CertGuard
	return r
}

func (r ContentguardsX509APIContentguardsCertguardX509UpdateRequest) Execute() (*CertguardX509CertGuardResponse, *http.Response, error) {
	return r.ApiService.ContentguardsCertguardX509UpdateExecute(r)
}

/*
ContentguardsCertguardX509Update Update a x509 cert guard

X509CertGuard API Viewsets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param certguardX509CertGuardHref
 @return ContentguardsX509APIContentguardsCertguardX509UpdateRequest
*/
func (a *ContentguardsX509APIService) ContentguardsCertguardX509Update(ctx context.Context, certguardX509CertGuardHref string) ContentguardsX509APIContentguardsCertguardX509UpdateRequest {
	return ContentguardsX509APIContentguardsCertguardX509UpdateRequest{
		ApiService: a,
		ctx: ctx,
		certguardX509CertGuardHref: certguardX509CertGuardHref,
	}
}

// Execute executes the request
//  @return CertguardX509CertGuardResponse
func (a *ContentguardsX509APIService) ContentguardsCertguardX509UpdateExecute(r ContentguardsX509APIContentguardsCertguardX509UpdateRequest) (*CertguardX509CertGuardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertguardX509CertGuardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentguardsX509APIService.ContentguardsCertguardX509Update")
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
