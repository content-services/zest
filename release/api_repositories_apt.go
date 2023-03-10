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


// RepositoriesAptApiService RepositoriesAptApi service
type RepositoriesAptApiService service

type RepositoriesAptApiRepositoriesDebAptCreateRequest struct {
	ctx context.Context
	ApiService *RepositoriesAptApiService
	debAptRepository *DebAptRepository
}

func (r RepositoriesAptApiRepositoriesDebAptCreateRequest) DebAptRepository(debAptRepository DebAptRepository) RepositoriesAptApiRepositoriesDebAptCreateRequest {
	r.debAptRepository = &debAptRepository
	return r
}

func (r RepositoriesAptApiRepositoriesDebAptCreateRequest) Execute() (*DebAptRepositoryResponse, *http.Response, error) {
	return r.ApiService.RepositoriesDebAptCreateExecute(r)
}

/*
RepositoriesDebAptCreate Create an apt repository

An AptRepository is the locally stored, Pulp-internal representation of a APT repository.

It may be filled with content via synchronization or content upload to create an
AptRepositoryVersion.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RepositoriesAptApiRepositoriesDebAptCreateRequest
*/
func (a *RepositoriesAptApiService) RepositoriesDebAptCreate(ctx context.Context) RepositoriesAptApiRepositoriesDebAptCreateRequest {
	return RepositoriesAptApiRepositoriesDebAptCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DebAptRepositoryResponse
func (a *RepositoriesAptApiService) RepositoriesDebAptCreateExecute(r RepositoriesAptApiRepositoriesDebAptCreateRequest) (*DebAptRepositoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebAptRepositoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesAptApiService.RepositoriesDebAptCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/repositories/deb/apt/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.debAptRepository == nil {
		return localVarReturnValue, nil, reportError("debAptRepository is required and must be specified")
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
	localVarPostBody = r.debAptRepository
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

type RepositoriesAptApiRepositoriesDebAptDeleteRequest struct {
	ctx context.Context
	ApiService *RepositoriesAptApiService
	debAptRepositoryHref string
}

func (r RepositoriesAptApiRepositoriesDebAptDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesDebAptDeleteExecute(r)
}

/*
RepositoriesDebAptDelete Delete an apt repository

Trigger an asynchronous delete task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptRepositoryHref
 @return RepositoriesAptApiRepositoriesDebAptDeleteRequest
*/
func (a *RepositoriesAptApiService) RepositoriesDebAptDelete(ctx context.Context, debAptRepositoryHref string) RepositoriesAptApiRepositoriesDebAptDeleteRequest {
	return RepositoriesAptApiRepositoriesDebAptDeleteRequest{
		ApiService: a,
		ctx: ctx,
		debAptRepositoryHref: debAptRepositoryHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesAptApiService) RepositoriesDebAptDeleteExecute(r RepositoriesAptApiRepositoriesDebAptDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesAptApiService.RepositoriesDebAptDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_apt_repository_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_repository_href"+"}", url.PathEscape(parameterValueToString(r.debAptRepositoryHref, "debAptRepositoryHref")), -1)
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

type RepositoriesAptApiRepositoriesDebAptListRequest struct {
	ctx context.Context
	ApiService *RepositoriesAptApiService
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
	offset *int32
	ordering *[]string
	pulpLabelSelect *string
	remote *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r RepositoriesAptApiRepositoriesDebAptListRequest) Limit(limit int32) RepositoriesAptApiRepositoriesDebAptListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r RepositoriesAptApiRepositoriesDebAptListRequest) Name(name string) RepositoriesAptApiRepositoriesDebAptListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r RepositoriesAptApiRepositoriesDebAptListRequest) NameContains(nameContains string) RepositoriesAptApiRepositoriesDebAptListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r RepositoriesAptApiRepositoriesDebAptListRequest) NameIcontains(nameIcontains string) RepositoriesAptApiRepositoriesDebAptListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r RepositoriesAptApiRepositoriesDebAptListRequest) NameIn(nameIn []string) RepositoriesAptApiRepositoriesDebAptListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r RepositoriesAptApiRepositoriesDebAptListRequest) NameStartswith(nameStartswith string) RepositoriesAptApiRepositoriesDebAptListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r RepositoriesAptApiRepositoriesDebAptListRequest) Offset(offset int32) RepositoriesAptApiRepositoriesDebAptListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r RepositoriesAptApiRepositoriesDebAptListRequest) Ordering(ordering []string) RepositoriesAptApiRepositoriesDebAptListRequest {
	r.ordering = &ordering
	return r
}

// Filter labels by search string
func (r RepositoriesAptApiRepositoriesDebAptListRequest) PulpLabelSelect(pulpLabelSelect string) RepositoriesAptApiRepositoriesDebAptListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Foreign Key referenced by HREF
func (r RepositoriesAptApiRepositoriesDebAptListRequest) Remote(remote string) RepositoriesAptApiRepositoriesDebAptListRequest {
	r.remote = &remote
	return r
}

// A list of fields to include in the response.
func (r RepositoriesAptApiRepositoriesDebAptListRequest) Fields(fields []string) RepositoriesAptApiRepositoriesDebAptListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesAptApiRepositoriesDebAptListRequest) ExcludeFields(excludeFields []string) RepositoriesAptApiRepositoriesDebAptListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesAptApiRepositoriesDebAptListRequest) Execute() (*PaginateddebAptRepositoryResponseList, *http.Response, error) {
	return r.ApiService.RepositoriesDebAptListExecute(r)
}

/*
RepositoriesDebAptList List apt repositorys

An AptRepository is the locally stored, Pulp-internal representation of a APT repository.

It may be filled with content via synchronization or content upload to create an
AptRepositoryVersion.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RepositoriesAptApiRepositoriesDebAptListRequest
*/
func (a *RepositoriesAptApiService) RepositoriesDebAptList(ctx context.Context) RepositoriesAptApiRepositoriesDebAptListRequest {
	return RepositoriesAptApiRepositoriesDebAptListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginateddebAptRepositoryResponseList
func (a *RepositoriesAptApiService) RepositoriesDebAptListExecute(r RepositoriesAptApiRepositoriesDebAptListRequest) (*PaginateddebAptRepositoryResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginateddebAptRepositoryResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesAptApiService.RepositoriesDebAptList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/repositories/deb/apt/"
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
	if r.remote != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "remote", r.remote, "")
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

type RepositoriesAptApiRepositoriesDebAptModifyRequest struct {
	ctx context.Context
	ApiService *RepositoriesAptApiService
	debAptRepositoryHref string
	repositoryAddRemoveContent *RepositoryAddRemoveContent
}

func (r RepositoriesAptApiRepositoriesDebAptModifyRequest) RepositoryAddRemoveContent(repositoryAddRemoveContent RepositoryAddRemoveContent) RepositoriesAptApiRepositoriesDebAptModifyRequest {
	r.repositoryAddRemoveContent = &repositoryAddRemoveContent
	return r
}

func (r RepositoriesAptApiRepositoriesDebAptModifyRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesDebAptModifyExecute(r)
}

/*
RepositoriesDebAptModify Modify Repository Content

Trigger an asynchronous task to create a new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptRepositoryHref
 @return RepositoriesAptApiRepositoriesDebAptModifyRequest
*/
func (a *RepositoriesAptApiService) RepositoriesDebAptModify(ctx context.Context, debAptRepositoryHref string) RepositoriesAptApiRepositoriesDebAptModifyRequest {
	return RepositoriesAptApiRepositoriesDebAptModifyRequest{
		ApiService: a,
		ctx: ctx,
		debAptRepositoryHref: debAptRepositoryHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesAptApiService) RepositoriesDebAptModifyExecute(r RepositoriesAptApiRepositoriesDebAptModifyRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesAptApiService.RepositoriesDebAptModify")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_apt_repository_href}modify/"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_repository_href"+"}", url.PathEscape(parameterValueToString(r.debAptRepositoryHref, "debAptRepositoryHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.repositoryAddRemoveContent == nil {
		return localVarReturnValue, nil, reportError("repositoryAddRemoveContent is required and must be specified")
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
	localVarPostBody = r.repositoryAddRemoveContent
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

type RepositoriesAptApiRepositoriesDebAptPartialUpdateRequest struct {
	ctx context.Context
	ApiService *RepositoriesAptApiService
	debAptRepositoryHref string
	patcheddebAptRepository *PatcheddebAptRepository
}

func (r RepositoriesAptApiRepositoriesDebAptPartialUpdateRequest) PatcheddebAptRepository(patcheddebAptRepository PatcheddebAptRepository) RepositoriesAptApiRepositoriesDebAptPartialUpdateRequest {
	r.patcheddebAptRepository = &patcheddebAptRepository
	return r
}

func (r RepositoriesAptApiRepositoriesDebAptPartialUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesDebAptPartialUpdateExecute(r)
}

/*
RepositoriesDebAptPartialUpdate Update an apt repository

Trigger an asynchronous partial update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptRepositoryHref
 @return RepositoriesAptApiRepositoriesDebAptPartialUpdateRequest
*/
func (a *RepositoriesAptApiService) RepositoriesDebAptPartialUpdate(ctx context.Context, debAptRepositoryHref string) RepositoriesAptApiRepositoriesDebAptPartialUpdateRequest {
	return RepositoriesAptApiRepositoriesDebAptPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		debAptRepositoryHref: debAptRepositoryHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesAptApiService) RepositoriesDebAptPartialUpdateExecute(r RepositoriesAptApiRepositoriesDebAptPartialUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesAptApiService.RepositoriesDebAptPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_apt_repository_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_repository_href"+"}", url.PathEscape(parameterValueToString(r.debAptRepositoryHref, "debAptRepositoryHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patcheddebAptRepository == nil {
		return localVarReturnValue, nil, reportError("patcheddebAptRepository is required and must be specified")
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
	localVarPostBody = r.patcheddebAptRepository
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

type RepositoriesAptApiRepositoriesDebAptReadRequest struct {
	ctx context.Context
	ApiService *RepositoriesAptApiService
	debAptRepositoryHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r RepositoriesAptApiRepositoriesDebAptReadRequest) Fields(fields []string) RepositoriesAptApiRepositoriesDebAptReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesAptApiRepositoriesDebAptReadRequest) ExcludeFields(excludeFields []string) RepositoriesAptApiRepositoriesDebAptReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesAptApiRepositoriesDebAptReadRequest) Execute() (*DebAptRepositoryResponse, *http.Response, error) {
	return r.ApiService.RepositoriesDebAptReadExecute(r)
}

/*
RepositoriesDebAptRead Inspect an apt repository

An AptRepository is the locally stored, Pulp-internal representation of a APT repository.

It may be filled with content via synchronization or content upload to create an
AptRepositoryVersion.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptRepositoryHref
 @return RepositoriesAptApiRepositoriesDebAptReadRequest
*/
func (a *RepositoriesAptApiService) RepositoriesDebAptRead(ctx context.Context, debAptRepositoryHref string) RepositoriesAptApiRepositoriesDebAptReadRequest {
	return RepositoriesAptApiRepositoriesDebAptReadRequest{
		ApiService: a,
		ctx: ctx,
		debAptRepositoryHref: debAptRepositoryHref,
	}
}

// Execute executes the request
//  @return DebAptRepositoryResponse
func (a *RepositoriesAptApiService) RepositoriesDebAptReadExecute(r RepositoriesAptApiRepositoriesDebAptReadRequest) (*DebAptRepositoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebAptRepositoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesAptApiService.RepositoriesDebAptRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_apt_repository_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_repository_href"+"}", url.PathEscape(parameterValueToString(r.debAptRepositoryHref, "debAptRepositoryHref")), -1)
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

type RepositoriesAptApiRepositoriesDebAptSyncRequest struct {
	ctx context.Context
	ApiService *RepositoriesAptApiService
	debAptRepositoryHref string
	aptRepositorySyncURL *AptRepositorySyncURL
}

func (r RepositoriesAptApiRepositoriesDebAptSyncRequest) AptRepositorySyncURL(aptRepositorySyncURL AptRepositorySyncURL) RepositoriesAptApiRepositoriesDebAptSyncRequest {
	r.aptRepositorySyncURL = &aptRepositorySyncURL
	return r
}

func (r RepositoriesAptApiRepositoriesDebAptSyncRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesDebAptSyncExecute(r)
}

/*
RepositoriesDebAptSync Sync from remote

Trigger an asynchronous task to sync content

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptRepositoryHref
 @return RepositoriesAptApiRepositoriesDebAptSyncRequest
*/
func (a *RepositoriesAptApiService) RepositoriesDebAptSync(ctx context.Context, debAptRepositoryHref string) RepositoriesAptApiRepositoriesDebAptSyncRequest {
	return RepositoriesAptApiRepositoriesDebAptSyncRequest{
		ApiService: a,
		ctx: ctx,
		debAptRepositoryHref: debAptRepositoryHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesAptApiService) RepositoriesDebAptSyncExecute(r RepositoriesAptApiRepositoriesDebAptSyncRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesAptApiService.RepositoriesDebAptSync")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_apt_repository_href}sync/"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_repository_href"+"}", url.PathEscape(parameterValueToString(r.debAptRepositoryHref, "debAptRepositoryHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aptRepositorySyncURL == nil {
		return localVarReturnValue, nil, reportError("aptRepositorySyncURL is required and must be specified")
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
	localVarPostBody = r.aptRepositorySyncURL
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

type RepositoriesAptApiRepositoriesDebAptUpdateRequest struct {
	ctx context.Context
	ApiService *RepositoriesAptApiService
	debAptRepositoryHref string
	debAptRepository *DebAptRepository
}

func (r RepositoriesAptApiRepositoriesDebAptUpdateRequest) DebAptRepository(debAptRepository DebAptRepository) RepositoriesAptApiRepositoriesDebAptUpdateRequest {
	r.debAptRepository = &debAptRepository
	return r
}

func (r RepositoriesAptApiRepositoriesDebAptUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesDebAptUpdateExecute(r)
}

/*
RepositoriesDebAptUpdate Update an apt repository

Trigger an asynchronous update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptRepositoryHref
 @return RepositoriesAptApiRepositoriesDebAptUpdateRequest
*/
func (a *RepositoriesAptApiService) RepositoriesDebAptUpdate(ctx context.Context, debAptRepositoryHref string) RepositoriesAptApiRepositoriesDebAptUpdateRequest {
	return RepositoriesAptApiRepositoriesDebAptUpdateRequest{
		ApiService: a,
		ctx: ctx,
		debAptRepositoryHref: debAptRepositoryHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesAptApiService) RepositoriesDebAptUpdateExecute(r RepositoriesAptApiRepositoriesDebAptUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesAptApiService.RepositoriesDebAptUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_apt_repository_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_repository_href"+"}", url.PathEscape(parameterValueToString(r.debAptRepositoryHref, "debAptRepositoryHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.debAptRepository == nil {
		return localVarReturnValue, nil, reportError("debAptRepository is required and must be specified")
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
	localVarPostBody = r.debAptRepository
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
