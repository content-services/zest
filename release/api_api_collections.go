/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// ApiCollectionsApiService ApiCollectionsApi service
type ApiCollectionsApiService service

type ApiCollectionsApiApiV2CollectionsGetRequest struct {
	ctx context.Context
	ApiService *ApiCollectionsApiService
	ansibleCollectionHref string
	page *int32
	fields *[]string
	excludeFields *[]string
}

// A page number within the paginated result set.
func (r ApiCollectionsApiApiV2CollectionsGetRequest) Page(page int32) ApiCollectionsApiApiV2CollectionsGetRequest {
	r.page = &page
	return r
}

// A list of fields to include in the response.
func (r ApiCollectionsApiApiV2CollectionsGetRequest) Fields(fields []string) ApiCollectionsApiApiV2CollectionsGetRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ApiCollectionsApiApiV2CollectionsGetRequest) ExcludeFields(excludeFields []string) ApiCollectionsApiApiV2CollectionsGetRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiCollectionsApiApiV2CollectionsGetRequest) Execute() (*PaginatedGalaxyCollectionResponseList, *http.Response, error) {
	return r.ApiService.ApiV2CollectionsGetExecute(r)
}

/*
ApiV2CollectionsGet Method for ApiV2CollectionsGet

View for Collection models.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleCollectionHref
 @return ApiCollectionsApiApiV2CollectionsGetRequest
*/
func (a *ApiCollectionsApiService) ApiV2CollectionsGet(ctx context.Context, ansibleCollectionHref string) ApiCollectionsApiApiV2CollectionsGetRequest {
	return ApiCollectionsApiApiV2CollectionsGetRequest{
		ApiService: a,
		ctx: ctx,
		ansibleCollectionHref: ansibleCollectionHref,
	}
}

// Execute executes the request
//  @return PaginatedGalaxyCollectionResponseList
func (a *ApiCollectionsApiService) ApiV2CollectionsGetExecute(r ApiCollectionsApiApiV2CollectionsGetRequest) (*PaginatedGalaxyCollectionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedGalaxyCollectionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiCollectionsApiService.ApiV2CollectionsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ansible_collection_href}api/v2/collections/"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_collection_href"+"}", url.PathEscape(parameterValueToString(r.ansibleCollectionHref, "ansibleCollectionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
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

type ApiCollectionsApiApiV2CollectionsPostRequest struct {
	ctx context.Context
	ApiService *ApiCollectionsApiService
	ansibleCollectionHref string
	galaxyCollection *GalaxyCollection
}

func (r ApiCollectionsApiApiV2CollectionsPostRequest) GalaxyCollection(galaxyCollection GalaxyCollection) ApiCollectionsApiApiV2CollectionsPostRequest {
	r.galaxyCollection = &galaxyCollection
	return r
}

func (r ApiCollectionsApiApiV2CollectionsPostRequest) Execute() (*GalaxyCollectionResponse, *http.Response, error) {
	return r.ApiService.ApiV2CollectionsPostExecute(r)
}

/*
ApiV2CollectionsPost Method for ApiV2CollectionsPost

Queues a task that creates a new Collection from an uploaded artifact.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleCollectionHref
 @return ApiCollectionsApiApiV2CollectionsPostRequest

Deprecated
*/
func (a *ApiCollectionsApiService) ApiV2CollectionsPost(ctx context.Context, ansibleCollectionHref string) ApiCollectionsApiApiV2CollectionsPostRequest {
	return ApiCollectionsApiApiV2CollectionsPostRequest{
		ApiService: a,
		ctx: ctx,
		ansibleCollectionHref: ansibleCollectionHref,
	}
}

// Execute executes the request
//  @return GalaxyCollectionResponse
// Deprecated
func (a *ApiCollectionsApiService) ApiV2CollectionsPostExecute(r ApiCollectionsApiApiV2CollectionsPostRequest) (*GalaxyCollectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GalaxyCollectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiCollectionsApiService.ApiV2CollectionsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ansible_collection_href}api/v2/collections/"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_collection_href"+"}", url.PathEscape(parameterValueToString(r.ansibleCollectionHref, "ansibleCollectionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.galaxyCollection == nil {
		return localVarReturnValue, nil, reportError("galaxyCollection is required and must be specified")
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
	localVarPostBody = r.galaxyCollection
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
