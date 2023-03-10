/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpGoBinding

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// SigningServicesApiService SigningServicesApi service
type SigningServicesApiService service

type SigningServicesApiSigningServicesListRequest struct {
	ctx context.Context
	ApiService *SigningServicesApiService
	limit *int32
	name *string
	offset *int32
	ordering *[]string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r SigningServicesApiSigningServicesListRequest) Limit(limit int32) SigningServicesApiSigningServicesListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r SigningServicesApiSigningServicesListRequest) Name(name string) SigningServicesApiSigningServicesListRequest {
	r.name = &name
	return r
}

// The initial index from which to return the results.
func (r SigningServicesApiSigningServicesListRequest) Offset(offset int32) SigningServicesApiSigningServicesListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r SigningServicesApiSigningServicesListRequest) Ordering(ordering []string) SigningServicesApiSigningServicesListRequest {
	r.ordering = &ordering
	return r
}

// A list of fields to include in the response.
func (r SigningServicesApiSigningServicesListRequest) Fields(fields []string) SigningServicesApiSigningServicesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r SigningServicesApiSigningServicesListRequest) ExcludeFields(excludeFields []string) SigningServicesApiSigningServicesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r SigningServicesApiSigningServicesListRequest) Execute() (*PaginatedSigningServiceResponseList, *http.Response, error) {
	return r.ApiService.SigningServicesListExecute(r)
}

/*
SigningServicesList List signing services

A ViewSet that supports browsing of existing signing services.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SigningServicesApiSigningServicesListRequest
*/
func (a *SigningServicesApiService) SigningServicesList(ctx context.Context) SigningServicesApiSigningServicesListRequest {
	return SigningServicesApiSigningServicesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedSigningServiceResponseList
func (a *SigningServicesApiService) SigningServicesListExecute(r SigningServicesApiSigningServicesListRequest) (*PaginatedSigningServiceResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedSigningServiceResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SigningServicesApiService.SigningServicesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/signing-services/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
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

type SigningServicesApiSigningServicesReadRequest struct {
	ctx context.Context
	ApiService *SigningServicesApiService
	signingServiceHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r SigningServicesApiSigningServicesReadRequest) Fields(fields []string) SigningServicesApiSigningServicesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r SigningServicesApiSigningServicesReadRequest) ExcludeFields(excludeFields []string) SigningServicesApiSigningServicesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r SigningServicesApiSigningServicesReadRequest) Execute() (*SigningServiceResponse, *http.Response, error) {
	return r.ApiService.SigningServicesReadExecute(r)
}

/*
SigningServicesRead Inspect a signing service

A ViewSet that supports browsing of existing signing services.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param signingServiceHref
 @return SigningServicesApiSigningServicesReadRequest
*/
func (a *SigningServicesApiService) SigningServicesRead(ctx context.Context, signingServiceHref string) SigningServicesApiSigningServicesReadRequest {
	return SigningServicesApiSigningServicesReadRequest{
		ApiService: a,
		ctx: ctx,
		signingServiceHref: signingServiceHref,
	}
}

// Execute executes the request
//  @return SigningServiceResponse
func (a *SigningServicesApiService) SigningServicesReadExecute(r SigningServicesApiSigningServicesReadRequest) (*SigningServiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SigningServiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SigningServicesApiService.SigningServicesRead")
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
