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


// PulpAnsibleApiV3CollectionsVersionsApiService PulpAnsibleApiV3CollectionsVersionsApi service
type PulpAnsibleApiV3CollectionsVersionsApiService service

type PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsDeleteRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3CollectionsVersionsApiService
	name string
	namespace string
	path string
	version string
}

func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3CollectionsVersionsDeleteExecute(r)
}

/*
PulpAnsibleGalaxyApiV3CollectionsVersionsDelete Method for PulpAnsibleGalaxyApiV3CollectionsVersionsDelete

Legacy v3 endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @param namespace
 @param path
 @param version
 @return PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsDeleteRequest

Deprecated
*/
func (a *PulpAnsibleApiV3CollectionsVersionsApiService) PulpAnsibleGalaxyApiV3CollectionsVersionsDelete(ctx context.Context, name string, namespace string, path string, version string) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsDeleteRequest {
	return PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
		namespace: namespace,
		path: path,
		version: version,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
// Deprecated
func (a *PulpAnsibleApiV3CollectionsVersionsApiService) PulpAnsibleGalaxyApiV3CollectionsVersionsDeleteExecute(r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3CollectionsVersionsApiService.PulpAnsibleGalaxyApiV3CollectionsVersionsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/collections/{namespace}/{name}/versions/{version}/"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(parameterValueToString(r.namespace, "namespace")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterValueToString(r.path, "path")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", url.PathEscape(parameterValueToString(r.version, "version")), -1)
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

type PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3CollectionsVersionsApiService
	name string
	namespace string
	path string
	isHighest *bool
	limit *int32
	name2 *string
	namespace2 *string
	offset *int32
	ordering *[]string
	q *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	tags *string
	version *string
	fields *[]string
	excludeFields *[]string
}

func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) IsHighest(isHighest bool) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.isHighest = &isHighest
	return r
}

// Number of results to return per page.
func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Limit(limit int32) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.limit = &limit
	return r
}

func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Name2(name2 string) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.name2 = &name2
	return r
}

func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Namespace2(namespace2 string) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.namespace2 = &namespace2
	return r
}

// The initial index from which to return the results.
func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Offset(offset int32) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Ordering(ordering []string) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.ordering = &ordering
	return r
}

func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Q(q string) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF
func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) RepositoryVersion(repositoryVersion string) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter by comma separate list of tags that must all be matched
func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Tags(tags string) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.tags = &tags
	return r
}

// Filter results where version matches value
func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Version(version string) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.version = &version
	return r
}

// A list of fields to include in the response.
func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Fields(fields []string) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) ExcludeFields(excludeFields []string) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) Execute() (*PaginatedCollectionVersionListResponseList, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3CollectionsVersionsListExecute(r)
}

/*
PulpAnsibleGalaxyApiV3CollectionsVersionsList Method for PulpAnsibleGalaxyApiV3CollectionsVersionsList

Legacy v3 endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @param namespace
 @param path
 @return PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest

Deprecated
*/
func (a *PulpAnsibleApiV3CollectionsVersionsApiService) PulpAnsibleGalaxyApiV3CollectionsVersionsList(ctx context.Context, name string, namespace string, path string) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest {
	return PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
		namespace: namespace,
		path: path,
	}
}

// Execute executes the request
//  @return PaginatedCollectionVersionListResponseList
// Deprecated
func (a *PulpAnsibleApiV3CollectionsVersionsApiService) PulpAnsibleGalaxyApiV3CollectionsVersionsListExecute(r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest) (*PaginatedCollectionVersionListResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedCollectionVersionListResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3CollectionsVersionsApiService.PulpAnsibleGalaxyApiV3CollectionsVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/collections/{namespace}/{name}/versions/"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(parameterValueToString(r.namespace, "namespace")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterValueToString(r.path, "path")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isHighest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_highest", r.isHighest, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name2, "")
	}
	if r.namespace2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace2, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tags", r.tags, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
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

type PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3CollectionsVersionsApiService
	name string
	namespace string
	path string
	version string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest) Fields(fields []string) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest) ExcludeFields(excludeFields []string) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest) Execute() (*CollectionVersionResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3CollectionsVersionsReadExecute(r)
}

/*
PulpAnsibleGalaxyApiV3CollectionsVersionsRead Method for PulpAnsibleGalaxyApiV3CollectionsVersionsRead

Legacy v3 endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @param namespace
 @param path
 @param version
 @return PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest

Deprecated
*/
func (a *PulpAnsibleApiV3CollectionsVersionsApiService) PulpAnsibleGalaxyApiV3CollectionsVersionsRead(ctx context.Context, name string, namespace string, path string, version string) PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest {
	return PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
		namespace: namespace,
		path: path,
		version: version,
	}
}

// Execute executes the request
//  @return CollectionVersionResponse
// Deprecated
func (a *PulpAnsibleApiV3CollectionsVersionsApiService) PulpAnsibleGalaxyApiV3CollectionsVersionsReadExecute(r PulpAnsibleApiV3CollectionsVersionsApiPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest) (*CollectionVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3CollectionsVersionsApiService.PulpAnsibleGalaxyApiV3CollectionsVersionsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/collections/{namespace}/{name}/versions/{version}/"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(parameterValueToString(r.namespace, "namespace")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterValueToString(r.path, "path")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", url.PathEscape(parameterValueToString(r.version, "version")), -1)
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
