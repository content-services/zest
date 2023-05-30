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


// GalaxyDetailAPIService GalaxyDetailAPI service
type GalaxyDetailAPIService service

type GalaxyDetailAPIGalaxyCollectionDetailGetRequest struct {
	ctx context.Context
	ApiService *GalaxyDetailAPIService
	ansibleCollectionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r GalaxyDetailAPIGalaxyCollectionDetailGetRequest) Fields(fields []string) GalaxyDetailAPIGalaxyCollectionDetailGetRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r GalaxyDetailAPIGalaxyCollectionDetailGetRequest) ExcludeFields(excludeFields []string) GalaxyDetailAPIGalaxyCollectionDetailGetRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r GalaxyDetailAPIGalaxyCollectionDetailGetRequest) Execute() (*GalaxyCollectionResponse, *http.Response, error) {
	return r.ApiService.GalaxyCollectionDetailGetExecute(r)
}

/*
GalaxyCollectionDetailGet Method for GalaxyCollectionDetailGet

Get the detail view of a Collection.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleCollectionHref
 @return GalaxyDetailAPIGalaxyCollectionDetailGetRequest
*/
func (a *GalaxyDetailAPIService) GalaxyCollectionDetailGet(ctx context.Context, ansibleCollectionHref string) GalaxyDetailAPIGalaxyCollectionDetailGetRequest {
	return GalaxyDetailAPIGalaxyCollectionDetailGetRequest{
		ApiService: a,
		ctx: ctx,
		ansibleCollectionHref: ansibleCollectionHref,
	}
}

// Execute executes the request
//  @return GalaxyCollectionResponse
func (a *GalaxyDetailAPIService) GalaxyCollectionDetailGetExecute(r GalaxyDetailAPIGalaxyCollectionDetailGetRequest) (*GalaxyCollectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GalaxyCollectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GalaxyDetailAPIService.GalaxyCollectionDetailGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ansible_collection_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_collection_href"+"}", url.PathEscape(parameterValueToString(r.ansibleCollectionHref, "ansibleCollectionHref")), -1)
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
