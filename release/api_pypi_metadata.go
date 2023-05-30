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


// PypiMetadataAPIService PypiMetadataAPI service
type PypiMetadataAPIService service

type PypiMetadataAPIPypiPypiReadRequest struct {
	ctx context.Context
	ApiService *PypiMetadataAPIService
	meta string
	path string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PypiMetadataAPIPypiPypiReadRequest) Fields(fields []string) PypiMetadataAPIPypiPypiReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PypiMetadataAPIPypiPypiReadRequest) ExcludeFields(excludeFields []string) PypiMetadataAPIPypiPypiReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PypiMetadataAPIPypiPypiReadRequest) Execute() (*PackageMetadataResponse, *http.Response, error) {
	return r.ApiService.PypiPypiReadExecute(r)
}

/*
PypiPypiRead Get package metadata

Retrieves the package's core-metadata specified by
https://packaging.python.org/specifications/core-metadata/.
`meta` must be a path in form of `{package}/json/` or `{package}/{version}/json/`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param meta
 @param path
 @return PypiMetadataAPIPypiPypiReadRequest
*/
func (a *PypiMetadataAPIService) PypiPypiRead(ctx context.Context, meta string, path string) PypiMetadataAPIPypiPypiReadRequest {
	return PypiMetadataAPIPypiPypiReadRequest{
		ApiService: a,
		ctx: ctx,
		meta: meta,
		path: path,
	}
}

// Execute executes the request
//  @return PackageMetadataResponse
func (a *PypiMetadataAPIService) PypiPypiReadExecute(r PypiMetadataAPIPypiPypiReadRequest) (*PackageMetadataResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PackageMetadataResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PypiMetadataAPIService.PypiPypiRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pypi/{path}/pypi/{meta}/"
	localVarPath = strings.Replace(localVarPath, "{"+"meta"+"}", url.PathEscape(parameterValueToString(r.meta, "meta")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterValueToString(r.path, "path")), -1)
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
