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


// SigningServicesAPIService SigningServicesAPI service
type SigningServicesAPIService service

type SigningServicesAPISigningServicesListRequest struct {
	ctx context.Context
	ApiService *SigningServicesAPIService
	pulpDomain string
	limit *int32
	name *string
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	q *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r SigningServicesAPISigningServicesListRequest) Limit(limit int32) SigningServicesAPISigningServicesListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r SigningServicesAPISigningServicesListRequest) Name(name string) SigningServicesAPISigningServicesListRequest {
	r.name = &name
	return r
}

// The initial index from which to return the results.
func (r SigningServicesAPISigningServicesListRequest) Offset(offset int32) SigningServicesAPISigningServicesListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;public_key&#x60; - Public key* &#x60;-public_key&#x60; - Public key (descending)* &#x60;pubkey_fingerprint&#x60; - Pubkey fingerprint* &#x60;-pubkey_fingerprint&#x60; - Pubkey fingerprint (descending)* &#x60;script&#x60; - Script* &#x60;-script&#x60; - Script (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r SigningServicesAPISigningServicesListRequest) Ordering(ordering []string) SigningServicesAPISigningServicesListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r SigningServicesAPISigningServicesListRequest) PulpHrefIn(pulpHrefIn []string) SigningServicesAPISigningServicesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r SigningServicesAPISigningServicesListRequest) PulpIdIn(pulpIdIn []string) SigningServicesAPISigningServicesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r SigningServicesAPISigningServicesListRequest) Q(q string) SigningServicesAPISigningServicesListRequest {
	r.q = &q
	return r
}

// A list of fields to include in the response.
func (r SigningServicesAPISigningServicesListRequest) Fields(fields []string) SigningServicesAPISigningServicesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r SigningServicesAPISigningServicesListRequest) ExcludeFields(excludeFields []string) SigningServicesAPISigningServicesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r SigningServicesAPISigningServicesListRequest) Execute() (*PaginatedSigningServiceResponseList, *http.Response, error) {
	return r.ApiService.SigningServicesListExecute(r)
}

/*
SigningServicesList List signing services

A ViewSet that supports browsing of existing signing services.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return SigningServicesAPISigningServicesListRequest
*/
func (a *SigningServicesAPIService) SigningServicesList(ctx context.Context, pulpDomain string) SigningServicesAPISigningServicesListRequest {
	return SigningServicesAPISigningServicesListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedSigningServiceResponseList
func (a *SigningServicesAPIService) SigningServicesListExecute(r SigningServicesAPISigningServicesListRequest) (*PaginatedSigningServiceResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedSigningServiceResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SigningServicesAPIService.SigningServicesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/signing-services/"
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

type SigningServicesAPISigningServicesReadRequest struct {
	ctx context.Context
	ApiService *SigningServicesAPIService
	signingServiceHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r SigningServicesAPISigningServicesReadRequest) Fields(fields []string) SigningServicesAPISigningServicesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r SigningServicesAPISigningServicesReadRequest) ExcludeFields(excludeFields []string) SigningServicesAPISigningServicesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r SigningServicesAPISigningServicesReadRequest) Execute() (*SigningServiceResponse, *http.Response, error) {
	return r.ApiService.SigningServicesReadExecute(r)
}

/*
SigningServicesRead Inspect a signing service

A ViewSet that supports browsing of existing signing services.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param signingServiceHref
 @return SigningServicesAPISigningServicesReadRequest
*/
func (a *SigningServicesAPIService) SigningServicesRead(ctx context.Context, signingServiceHref string) SigningServicesAPISigningServicesReadRequest {
	return SigningServicesAPISigningServicesReadRequest{
		ApiService: a,
		ctx: ctx,
		signingServiceHref: signingServiceHref,
	}
}

// Execute executes the request
//  @return SigningServiceResponse
func (a *SigningServicesAPIService) SigningServicesReadExecute(r SigningServicesAPISigningServicesReadRequest) (*SigningServiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SigningServiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SigningServicesAPIService.SigningServicesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{signing_service_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"signing_service_href"+"}", url.PathEscape(parameterValueToString(r.signingServiceHref, "signingServiceHref")), -1)
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
