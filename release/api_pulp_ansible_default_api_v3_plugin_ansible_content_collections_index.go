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


// PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApi service
type PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService service

type PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDeleteRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService
	distroBasePath string
	name string
	namespace string
}

func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDeleteExecute(r)
}

/*
PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDelete Method for PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDelete

Trigger an asynchronous delete task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distroBasePath
 @param name
 @param namespace
 @return PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDeleteRequest
*/
func (a *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService) PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDelete(ctx context.Context, distroBasePath string, name string, namespace string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDeleteRequest {
	return PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDeleteRequest{
		ApiService: a,
		ctx: ctx,
		distroBasePath: distroBasePath,
		name: name,
		namespace: namespace,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService) PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDeleteExecute(r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/"
	localVarPath = strings.Replace(localVarPath, "{"+"distro_base_path"+"}", url.PathEscape(parameterValueToString(r.distroBasePath, "distroBasePath")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(parameterValueToString(r.namespace, "namespace")), -1)
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

type PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService
	distroBasePath string
	deprecated *bool
	limit *int32
	name *string
	namespace *string
	offset *int32
	ordering *[]string
	fields *[]string
	excludeFields *[]string
}

func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest) Deprecated(deprecated bool) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.deprecated = &deprecated
	return r
}

// Number of results to return per page.
func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest) Limit(limit int32) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.limit = &limit
	return r
}

func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest) Name(name string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.name = &name
	return r
}

func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest) Namespace(namespace string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.namespace = &namespace
	return r
}

// The initial index from which to return the results.
func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest) Offset(offset int32) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest) Ordering(ordering []string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.ordering = &ordering
	return r
}

// A list of fields to include in the response.
func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest) Fields(fields []string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest) ExcludeFields(excludeFields []string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest) Execute() (*PaginatedCollectionResponseList, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListExecute(r)
}

/*
PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexList Method for PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexList

ViewSet for Collections.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distroBasePath
 @return PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest
*/
func (a *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService) PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexList(ctx context.Context, distroBasePath string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest {
	return PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest{
		ApiService: a,
		ctx: ctx,
		distroBasePath: distroBasePath,
	}
}

// Execute executes the request
//  @return PaginatedCollectionResponseList
func (a *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService) PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListExecute(r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexListRequest) (*PaginatedCollectionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedCollectionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/"
	localVarPath = strings.Replace(localVarPath, "{"+"distro_base_path"+"}", url.PathEscape(parameterValueToString(r.distroBasePath, "distroBasePath")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.deprecated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deprecated", r.deprecated, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.namespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace, "")
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

type PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexReadRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService
	distroBasePath string
	name string
	namespace string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexReadRequest) Fields(fields []string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexReadRequest) ExcludeFields(excludeFields []string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexReadRequest) Execute() (*CollectionResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexReadExecute(r)
}

/*
PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexRead Method for PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexRead

ViewSet for Collections.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distroBasePath
 @param name
 @param namespace
 @return PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexReadRequest
*/
func (a *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService) PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexRead(ctx context.Context, distroBasePath string, name string, namespace string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexReadRequest {
	return PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexReadRequest{
		ApiService: a,
		ctx: ctx,
		distroBasePath: distroBasePath,
		name: name,
		namespace: namespace,
	}
}

// Execute executes the request
//  @return CollectionResponse
func (a *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService) PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexReadExecute(r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexReadRequest) (*CollectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/"
	localVarPath = strings.Replace(localVarPath, "{"+"distro_base_path"+"}", url.PathEscape(parameterValueToString(r.distroBasePath, "distroBasePath")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(parameterValueToString(r.namespace, "namespace")), -1)
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

type PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdateRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService
	distroBasePath string
	name string
	namespace string
	body *map[string]interface{}
}

func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdateRequest) Body(body map[string]interface{}) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdateRequest {
	r.body = &body
	return r
}

func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdateExecute(r)
}

/*
PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdate Method for PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdate

Trigger an asynchronous update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distroBasePath
 @param name
 @param namespace
 @return PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdateRequest
*/
func (a *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService) PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdate(ctx context.Context, distroBasePath string, name string, namespace string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdateRequest {
	return PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdateRequest{
		ApiService: a,
		ctx: ctx,
		distroBasePath: distroBasePath,
		name: name,
		namespace: namespace,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService) PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdateExecute(r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/"
	localVarPath = strings.Replace(localVarPath, "{"+"distro_base_path"+"}", url.PathEscape(parameterValueToString(r.distroBasePath, "distroBasePath")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(parameterValueToString(r.namespace, "namespace")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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
