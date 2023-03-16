/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zest/release/v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// DistributionsMavenApiService DistributionsMavenApi service
type DistributionsMavenApiService service

type DistributionsMavenApiDistributionsMavenMavenCreateRequest struct {
	ctx context.Context
	ApiService *DistributionsMavenApiService
	mavenMavenDistribution *MavenMavenDistribution
}

func (r DistributionsMavenApiDistributionsMavenMavenCreateRequest) MavenMavenDistribution(mavenMavenDistribution MavenMavenDistribution) DistributionsMavenApiDistributionsMavenMavenCreateRequest {
	r.mavenMavenDistribution = &mavenMavenDistribution
	return r
}

func (r DistributionsMavenApiDistributionsMavenMavenCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsMavenMavenCreateExecute(r)
}

/*
DistributionsMavenMavenCreate Create a maven distribution

Trigger an asynchronous create task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DistributionsMavenApiDistributionsMavenMavenCreateRequest
*/
func (a *DistributionsMavenApiService) DistributionsMavenMavenCreate(ctx context.Context) DistributionsMavenApiDistributionsMavenMavenCreateRequest {
	return DistributionsMavenApiDistributionsMavenMavenCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsMavenApiService) DistributionsMavenMavenCreateExecute(r DistributionsMavenApiDistributionsMavenMavenCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsMavenApiService.DistributionsMavenMavenCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/distributions/maven/maven/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mavenMavenDistribution == nil {
		return localVarReturnValue, nil, reportError("mavenMavenDistribution is required and must be specified")
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
	localVarPostBody = r.mavenMavenDistribution
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

type DistributionsMavenApiDistributionsMavenMavenDeleteRequest struct {
	ctx context.Context
	ApiService *DistributionsMavenApiService
	mavenMavenDistributionHref string
}

func (r DistributionsMavenApiDistributionsMavenMavenDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsMavenMavenDeleteExecute(r)
}

/*
DistributionsMavenMavenDelete Delete a maven distribution

Trigger an asynchronous delete task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenDistributionHref
 @return DistributionsMavenApiDistributionsMavenMavenDeleteRequest
*/
func (a *DistributionsMavenApiService) DistributionsMavenMavenDelete(ctx context.Context, mavenMavenDistributionHref string) DistributionsMavenApiDistributionsMavenMavenDeleteRequest {
	return DistributionsMavenApiDistributionsMavenMavenDeleteRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenDistributionHref: mavenMavenDistributionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsMavenApiService) DistributionsMavenMavenDeleteExecute(r DistributionsMavenApiDistributionsMavenMavenDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsMavenApiService.DistributionsMavenMavenDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{maven_maven_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_distribution_href"+"}", url.PathEscape(parameterValueToString(r.mavenMavenDistributionHref, "mavenMavenDistributionHref")), -1)
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

type DistributionsMavenApiDistributionsMavenMavenListRequest struct {
	ctx context.Context
	ApiService *DistributionsMavenApiService
	basePath *string
	basePathContains *string
	basePathIcontains *string
	basePathIn *[]string
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
	offset *int32
	ordering *[]string
	pulpLabelSelect *string
	repository *string
	repositoryIn *[]string
	withContent *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where base_path matches value
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) BasePath(basePath string) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.basePath = &basePath
	return r
}

// Filter results where base_path contains value
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) BasePathContains(basePathContains string) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.basePathContains = &basePathContains
	return r
}

// Filter results where base_path contains value
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) BasePathIcontains(basePathIcontains string) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.basePathIcontains = &basePathIcontains
	return r
}

// Filter results where base_path is in a comma-separated list of values
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) BasePathIn(basePathIn []string) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.basePathIn = &basePathIn
	return r
}

// Number of results to return per page.
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) Limit(limit int32) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) Name(name string) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) NameContains(nameContains string) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) NameIcontains(nameIcontains string) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) NameIn(nameIn []string) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) NameStartswith(nameStartswith string) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) Offset(offset int32) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) Ordering(ordering []string) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.ordering = &ordering
	return r
}

// Filter labels by search string
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) PulpLabelSelect(pulpLabelSelect string) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Filter results where repository matches value
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) Repository(repository string) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.repository = &repository
	return r
}

// Filter results where repository is in a comma-separated list of values
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) RepositoryIn(repositoryIn []string) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.repositoryIn = &repositoryIn
	return r
}

// Filter distributions based on the content served by them
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) WithContent(withContent string) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.withContent = &withContent
	return r
}

// A list of fields to include in the response.
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) Fields(fields []string) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r DistributionsMavenApiDistributionsMavenMavenListRequest) ExcludeFields(excludeFields []string) DistributionsMavenApiDistributionsMavenMavenListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r DistributionsMavenApiDistributionsMavenMavenListRequest) Execute() (*PaginatedmavenMavenDistributionResponseList, *http.Response, error) {
	return r.ApiService.DistributionsMavenMavenListExecute(r)
}

/*
DistributionsMavenMavenList List maven distributions

ViewSet for Maven Distributions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DistributionsMavenApiDistributionsMavenMavenListRequest
*/
func (a *DistributionsMavenApiService) DistributionsMavenMavenList(ctx context.Context) DistributionsMavenApiDistributionsMavenMavenListRequest {
	return DistributionsMavenApiDistributionsMavenMavenListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedmavenMavenDistributionResponseList
func (a *DistributionsMavenApiService) DistributionsMavenMavenListExecute(r DistributionsMavenApiDistributionsMavenMavenListRequest) (*PaginatedmavenMavenDistributionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedmavenMavenDistributionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsMavenApiService.DistributionsMavenMavenList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/distributions/maven/maven/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.basePath != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path", r.basePath, "")
	}
	if r.basePathContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path__contains", r.basePathContains, "")
	}
	if r.basePathIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path__icontains", r.basePathIcontains, "")
	}
	if r.basePathIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path__in", r.basePathIn, "csv")
	}
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
	if r.nameIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__in", r.nameIn, "csv")
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
	if r.pulpLabelSelect != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_label_select", r.pulpLabelSelect, "")
	}
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository", r.repository, "")
	}
	if r.repositoryIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository__in", r.repositoryIn, "csv")
	}
	if r.withContent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "with_content", r.withContent, "")
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

type DistributionsMavenApiDistributionsMavenMavenPartialUpdateRequest struct {
	ctx context.Context
	ApiService *DistributionsMavenApiService
	mavenMavenDistributionHref string
	patchedmavenMavenDistribution *PatchedmavenMavenDistribution
}

func (r DistributionsMavenApiDistributionsMavenMavenPartialUpdateRequest) PatchedmavenMavenDistribution(patchedmavenMavenDistribution PatchedmavenMavenDistribution) DistributionsMavenApiDistributionsMavenMavenPartialUpdateRequest {
	r.patchedmavenMavenDistribution = &patchedmavenMavenDistribution
	return r
}

func (r DistributionsMavenApiDistributionsMavenMavenPartialUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsMavenMavenPartialUpdateExecute(r)
}

/*
DistributionsMavenMavenPartialUpdate Update a maven distribution

Trigger an asynchronous partial update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenDistributionHref
 @return DistributionsMavenApiDistributionsMavenMavenPartialUpdateRequest
*/
func (a *DistributionsMavenApiService) DistributionsMavenMavenPartialUpdate(ctx context.Context, mavenMavenDistributionHref string) DistributionsMavenApiDistributionsMavenMavenPartialUpdateRequest {
	return DistributionsMavenApiDistributionsMavenMavenPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenDistributionHref: mavenMavenDistributionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsMavenApiService) DistributionsMavenMavenPartialUpdateExecute(r DistributionsMavenApiDistributionsMavenMavenPartialUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsMavenApiService.DistributionsMavenMavenPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{maven_maven_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_distribution_href"+"}", url.PathEscape(parameterValueToString(r.mavenMavenDistributionHref, "mavenMavenDistributionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedmavenMavenDistribution == nil {
		return localVarReturnValue, nil, reportError("patchedmavenMavenDistribution is required and must be specified")
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
	localVarPostBody = r.patchedmavenMavenDistribution
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

type DistributionsMavenApiDistributionsMavenMavenReadRequest struct {
	ctx context.Context
	ApiService *DistributionsMavenApiService
	mavenMavenDistributionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r DistributionsMavenApiDistributionsMavenMavenReadRequest) Fields(fields []string) DistributionsMavenApiDistributionsMavenMavenReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r DistributionsMavenApiDistributionsMavenMavenReadRequest) ExcludeFields(excludeFields []string) DistributionsMavenApiDistributionsMavenMavenReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r DistributionsMavenApiDistributionsMavenMavenReadRequest) Execute() (*MavenMavenDistributionResponse, *http.Response, error) {
	return r.ApiService.DistributionsMavenMavenReadExecute(r)
}

/*
DistributionsMavenMavenRead Inspect a maven distribution

ViewSet for Maven Distributions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenDistributionHref
 @return DistributionsMavenApiDistributionsMavenMavenReadRequest
*/
func (a *DistributionsMavenApiService) DistributionsMavenMavenRead(ctx context.Context, mavenMavenDistributionHref string) DistributionsMavenApiDistributionsMavenMavenReadRequest {
	return DistributionsMavenApiDistributionsMavenMavenReadRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenDistributionHref: mavenMavenDistributionHref,
	}
}

// Execute executes the request
//  @return MavenMavenDistributionResponse
func (a *DistributionsMavenApiService) DistributionsMavenMavenReadExecute(r DistributionsMavenApiDistributionsMavenMavenReadRequest) (*MavenMavenDistributionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MavenMavenDistributionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsMavenApiService.DistributionsMavenMavenRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{maven_maven_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_distribution_href"+"}", url.PathEscape(parameterValueToString(r.mavenMavenDistributionHref, "mavenMavenDistributionHref")), -1)
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

type DistributionsMavenApiDistributionsMavenMavenUpdateRequest struct {
	ctx context.Context
	ApiService *DistributionsMavenApiService
	mavenMavenDistributionHref string
	mavenMavenDistribution *MavenMavenDistribution
}

func (r DistributionsMavenApiDistributionsMavenMavenUpdateRequest) MavenMavenDistribution(mavenMavenDistribution MavenMavenDistribution) DistributionsMavenApiDistributionsMavenMavenUpdateRequest {
	r.mavenMavenDistribution = &mavenMavenDistribution
	return r
}

func (r DistributionsMavenApiDistributionsMavenMavenUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsMavenMavenUpdateExecute(r)
}

/*
DistributionsMavenMavenUpdate Update a maven distribution

Trigger an asynchronous update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenDistributionHref
 @return DistributionsMavenApiDistributionsMavenMavenUpdateRequest
*/
func (a *DistributionsMavenApiService) DistributionsMavenMavenUpdate(ctx context.Context, mavenMavenDistributionHref string) DistributionsMavenApiDistributionsMavenMavenUpdateRequest {
	return DistributionsMavenApiDistributionsMavenMavenUpdateRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenDistributionHref: mavenMavenDistributionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsMavenApiService) DistributionsMavenMavenUpdateExecute(r DistributionsMavenApiDistributionsMavenMavenUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsMavenApiService.DistributionsMavenMavenUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{maven_maven_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_distribution_href"+"}", url.PathEscape(parameterValueToString(r.mavenMavenDistributionHref, "mavenMavenDistributionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mavenMavenDistribution == nil {
		return localVarReturnValue, nil, reportError("mavenMavenDistribution is required and must be specified")
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
	localVarPostBody = r.mavenMavenDistribution
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
