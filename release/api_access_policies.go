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


// AccessPoliciesAPIService AccessPoliciesAPI service
type AccessPoliciesAPIService service

type AccessPoliciesAPIAccessPoliciesListRequest struct {
	ctx context.Context
	ApiService *AccessPoliciesAPIService
	pulpDomain string
	customized *bool
	limit *int32
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	q *string
	viewsetName *string
	viewsetNameContains *string
	viewsetNameIcontains *string
	viewsetNameIexact *string
	viewsetNameIn *[]string
	viewsetNameIregex *string
	viewsetNameIstartswith *string
	viewsetNameRegex *string
	viewsetNameStartswith *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where customized matches value
func (r AccessPoliciesAPIAccessPoliciesListRequest) Customized(customized bool) AccessPoliciesAPIAccessPoliciesListRequest {
	r.customized = &customized
	return r
}

// Number of results to return per page.
func (r AccessPoliciesAPIAccessPoliciesListRequest) Limit(limit int32) AccessPoliciesAPIAccessPoliciesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r AccessPoliciesAPIAccessPoliciesListRequest) Offset(offset int32) AccessPoliciesAPIAccessPoliciesListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;creation_hooks&#x60; - Creation hooks* &#x60;-creation_hooks&#x60; - Creation hooks (descending)* &#x60;statements&#x60; - Statements* &#x60;-statements&#x60; - Statements (descending)* &#x60;viewset_name&#x60; - Viewset name* &#x60;-viewset_name&#x60; - Viewset name (descending)* &#x60;customized&#x60; - Customized* &#x60;-customized&#x60; - Customized (descending)* &#x60;queryset_scoping&#x60; - Queryset scoping* &#x60;-queryset_scoping&#x60; - Queryset scoping (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r AccessPoliciesAPIAccessPoliciesListRequest) Ordering(ordering []string) AccessPoliciesAPIAccessPoliciesListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r AccessPoliciesAPIAccessPoliciesListRequest) PulpHrefIn(pulpHrefIn []string) AccessPoliciesAPIAccessPoliciesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r AccessPoliciesAPIAccessPoliciesListRequest) PulpIdIn(pulpIdIn []string) AccessPoliciesAPIAccessPoliciesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r AccessPoliciesAPIAccessPoliciesListRequest) Q(q string) AccessPoliciesAPIAccessPoliciesListRequest {
	r.q = &q
	return r
}

// Filter results where viewset_name matches value
func (r AccessPoliciesAPIAccessPoliciesListRequest) ViewsetName(viewsetName string) AccessPoliciesAPIAccessPoliciesListRequest {
	r.viewsetName = &viewsetName
	return r
}

// Filter results where viewset_name contains value
func (r AccessPoliciesAPIAccessPoliciesListRequest) ViewsetNameContains(viewsetNameContains string) AccessPoliciesAPIAccessPoliciesListRequest {
	r.viewsetNameContains = &viewsetNameContains
	return r
}

// Filter results where viewset_name contains value
func (r AccessPoliciesAPIAccessPoliciesListRequest) ViewsetNameIcontains(viewsetNameIcontains string) AccessPoliciesAPIAccessPoliciesListRequest {
	r.viewsetNameIcontains = &viewsetNameIcontains
	return r
}

// Filter results where viewset_name matches value
func (r AccessPoliciesAPIAccessPoliciesListRequest) ViewsetNameIexact(viewsetNameIexact string) AccessPoliciesAPIAccessPoliciesListRequest {
	r.viewsetNameIexact = &viewsetNameIexact
	return r
}

// Filter results where viewset_name is in a comma-separated list of values
func (r AccessPoliciesAPIAccessPoliciesListRequest) ViewsetNameIn(viewsetNameIn []string) AccessPoliciesAPIAccessPoliciesListRequest {
	r.viewsetNameIn = &viewsetNameIn
	return r
}

// Filter results where viewset_name matches regex value
func (r AccessPoliciesAPIAccessPoliciesListRequest) ViewsetNameIregex(viewsetNameIregex string) AccessPoliciesAPIAccessPoliciesListRequest {
	r.viewsetNameIregex = &viewsetNameIregex
	return r
}

// Filter results where viewset_name starts with value
func (r AccessPoliciesAPIAccessPoliciesListRequest) ViewsetNameIstartswith(viewsetNameIstartswith string) AccessPoliciesAPIAccessPoliciesListRequest {
	r.viewsetNameIstartswith = &viewsetNameIstartswith
	return r
}

// Filter results where viewset_name matches regex value
func (r AccessPoliciesAPIAccessPoliciesListRequest) ViewsetNameRegex(viewsetNameRegex string) AccessPoliciesAPIAccessPoliciesListRequest {
	r.viewsetNameRegex = &viewsetNameRegex
	return r
}

// Filter results where viewset_name starts with value
func (r AccessPoliciesAPIAccessPoliciesListRequest) ViewsetNameStartswith(viewsetNameStartswith string) AccessPoliciesAPIAccessPoliciesListRequest {
	r.viewsetNameStartswith = &viewsetNameStartswith
	return r
}

// A list of fields to include in the response.
func (r AccessPoliciesAPIAccessPoliciesListRequest) Fields(fields []string) AccessPoliciesAPIAccessPoliciesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r AccessPoliciesAPIAccessPoliciesListRequest) ExcludeFields(excludeFields []string) AccessPoliciesAPIAccessPoliciesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r AccessPoliciesAPIAccessPoliciesListRequest) Execute() (*PaginatedAccessPolicyResponseList, *http.Response, error) {
	return r.ApiService.AccessPoliciesListExecute(r)
}

/*
AccessPoliciesList List access policys

ViewSet for AccessPolicy.NOTE: This API endpoint is in "tech preview" and subject to change

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return AccessPoliciesAPIAccessPoliciesListRequest
*/
func (a *AccessPoliciesAPIService) AccessPoliciesList(ctx context.Context, pulpDomain string) AccessPoliciesAPIAccessPoliciesListRequest {
	return AccessPoliciesAPIAccessPoliciesListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedAccessPolicyResponseList
func (a *AccessPoliciesAPIService) AccessPoliciesListExecute(r AccessPoliciesAPIAccessPoliciesListRequest) (*PaginatedAccessPolicyResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedAccessPolicyResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessPoliciesAPIService.AccessPoliciesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/access_policies/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.customized != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customized", r.customized, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
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
	if r.viewsetName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewset_name", r.viewsetName, "form", "")
	}
	if r.viewsetNameContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewset_name__contains", r.viewsetNameContains, "form", "")
	}
	if r.viewsetNameIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewset_name__icontains", r.viewsetNameIcontains, "form", "")
	}
	if r.viewsetNameIexact != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewset_name__iexact", r.viewsetNameIexact, "form", "")
	}
	if r.viewsetNameIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewset_name__in", r.viewsetNameIn, "form", "csv")
	}
	if r.viewsetNameIregex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewset_name__iregex", r.viewsetNameIregex, "form", "")
	}
	if r.viewsetNameIstartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewset_name__istartswith", r.viewsetNameIstartswith, "form", "")
	}
	if r.viewsetNameRegex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewset_name__regex", r.viewsetNameRegex, "form", "")
	}
	if r.viewsetNameStartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewset_name__startswith", r.viewsetNameStartswith, "form", "")
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

type AccessPoliciesAPIAccessPoliciesPartialUpdateRequest struct {
	ctx context.Context
	ApiService *AccessPoliciesAPIService
	accessPolicyHref string
	patchedAccessPolicy *PatchedAccessPolicy
}

func (r AccessPoliciesAPIAccessPoliciesPartialUpdateRequest) PatchedAccessPolicy(patchedAccessPolicy PatchedAccessPolicy) AccessPoliciesAPIAccessPoliciesPartialUpdateRequest {
	r.patchedAccessPolicy = &patchedAccessPolicy
	return r
}

func (r AccessPoliciesAPIAccessPoliciesPartialUpdateRequest) Execute() (*AccessPolicyResponse, *http.Response, error) {
	return r.ApiService.AccessPoliciesPartialUpdateExecute(r)
}

/*
AccessPoliciesPartialUpdate Update an access policy

ViewSet for AccessPolicy.NOTE: This API endpoint is in "tech preview" and subject to change

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accessPolicyHref
 @return AccessPoliciesAPIAccessPoliciesPartialUpdateRequest
*/
func (a *AccessPoliciesAPIService) AccessPoliciesPartialUpdate(ctx context.Context, accessPolicyHref string) AccessPoliciesAPIAccessPoliciesPartialUpdateRequest {
	return AccessPoliciesAPIAccessPoliciesPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		accessPolicyHref: accessPolicyHref,
	}
}

// Execute executes the request
//  @return AccessPolicyResponse
func (a *AccessPoliciesAPIService) AccessPoliciesPartialUpdateExecute(r AccessPoliciesAPIAccessPoliciesPartialUpdateRequest) (*AccessPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccessPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessPoliciesAPIService.AccessPoliciesPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{access_policy_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"access_policy_href"+"}", url.PathEscape(parameterValueToString(r.accessPolicyHref, "accessPolicyHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedAccessPolicy == nil {
		return localVarReturnValue, nil, reportError("patchedAccessPolicy is required and must be specified")
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
	localVarPostBody = r.patchedAccessPolicy
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

type AccessPoliciesAPIAccessPoliciesReadRequest struct {
	ctx context.Context
	ApiService *AccessPoliciesAPIService
	accessPolicyHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r AccessPoliciesAPIAccessPoliciesReadRequest) Fields(fields []string) AccessPoliciesAPIAccessPoliciesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r AccessPoliciesAPIAccessPoliciesReadRequest) ExcludeFields(excludeFields []string) AccessPoliciesAPIAccessPoliciesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r AccessPoliciesAPIAccessPoliciesReadRequest) Execute() (*AccessPolicyResponse, *http.Response, error) {
	return r.ApiService.AccessPoliciesReadExecute(r)
}

/*
AccessPoliciesRead Inspect an access policy

ViewSet for AccessPolicy.NOTE: This API endpoint is in "tech preview" and subject to change

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accessPolicyHref
 @return AccessPoliciesAPIAccessPoliciesReadRequest
*/
func (a *AccessPoliciesAPIService) AccessPoliciesRead(ctx context.Context, accessPolicyHref string) AccessPoliciesAPIAccessPoliciesReadRequest {
	return AccessPoliciesAPIAccessPoliciesReadRequest{
		ApiService: a,
		ctx: ctx,
		accessPolicyHref: accessPolicyHref,
	}
}

// Execute executes the request
//  @return AccessPolicyResponse
func (a *AccessPoliciesAPIService) AccessPoliciesReadExecute(r AccessPoliciesAPIAccessPoliciesReadRequest) (*AccessPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccessPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessPoliciesAPIService.AccessPoliciesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{access_policy_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"access_policy_href"+"}", url.PathEscape(parameterValueToString(r.accessPolicyHref, "accessPolicyHref")), -1)
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

type AccessPoliciesAPIAccessPoliciesResetRequest struct {
	ctx context.Context
	ApiService *AccessPoliciesAPIService
	accessPolicyHref string
}

func (r AccessPoliciesAPIAccessPoliciesResetRequest) Execute() (*AccessPolicyResponse, *http.Response, error) {
	return r.ApiService.AccessPoliciesResetExecute(r)
}

/*
AccessPoliciesReset Method for AccessPoliciesReset

Reset the access policy to its uncustomized default value.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accessPolicyHref
 @return AccessPoliciesAPIAccessPoliciesResetRequest
*/
func (a *AccessPoliciesAPIService) AccessPoliciesReset(ctx context.Context, accessPolicyHref string) AccessPoliciesAPIAccessPoliciesResetRequest {
	return AccessPoliciesAPIAccessPoliciesResetRequest{
		ApiService: a,
		ctx: ctx,
		accessPolicyHref: accessPolicyHref,
	}
}

// Execute executes the request
//  @return AccessPolicyResponse
func (a *AccessPoliciesAPIService) AccessPoliciesResetExecute(r AccessPoliciesAPIAccessPoliciesResetRequest) (*AccessPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccessPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessPoliciesAPIService.AccessPoliciesReset")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{access_policy_href}reset/"
	localVarPath = strings.Replace(localVarPath, "{"+"access_policy_href"+"}", url.PathEscape(parameterValueToString(r.accessPolicyHref, "accessPolicyHref")), -1)
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

type AccessPoliciesAPIAccessPoliciesUpdateRequest struct {
	ctx context.Context
	ApiService *AccessPoliciesAPIService
	accessPolicyHref string
	accessPolicy *AccessPolicy
}

func (r AccessPoliciesAPIAccessPoliciesUpdateRequest) AccessPolicy(accessPolicy AccessPolicy) AccessPoliciesAPIAccessPoliciesUpdateRequest {
	r.accessPolicy = &accessPolicy
	return r
}

func (r AccessPoliciesAPIAccessPoliciesUpdateRequest) Execute() (*AccessPolicyResponse, *http.Response, error) {
	return r.ApiService.AccessPoliciesUpdateExecute(r)
}

/*
AccessPoliciesUpdate Update an access policy

ViewSet for AccessPolicy.NOTE: This API endpoint is in "tech preview" and subject to change

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accessPolicyHref
 @return AccessPoliciesAPIAccessPoliciesUpdateRequest
*/
func (a *AccessPoliciesAPIService) AccessPoliciesUpdate(ctx context.Context, accessPolicyHref string) AccessPoliciesAPIAccessPoliciesUpdateRequest {
	return AccessPoliciesAPIAccessPoliciesUpdateRequest{
		ApiService: a,
		ctx: ctx,
		accessPolicyHref: accessPolicyHref,
	}
}

// Execute executes the request
//  @return AccessPolicyResponse
func (a *AccessPoliciesAPIService) AccessPoliciesUpdateExecute(r AccessPoliciesAPIAccessPoliciesUpdateRequest) (*AccessPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccessPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessPoliciesAPIService.AccessPoliciesUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{access_policy_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"access_policy_href"+"}", url.PathEscape(parameterValueToString(r.accessPolicyHref, "accessPolicyHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accessPolicy == nil {
		return localVarReturnValue, nil, reportError("accessPolicy is required and must be specified")
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
	localVarPostBody = r.accessPolicy
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
