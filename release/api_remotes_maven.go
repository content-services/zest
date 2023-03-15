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
	"time"
	"reflect"
)


// RemotesMavenApiService RemotesMavenApi service
type RemotesMavenApiService service

type RemotesMavenApiRemotesMavenMavenCreateRequest struct {
	ctx context.Context
	ApiService *RemotesMavenApiService
	mavenMavenRemote *MavenMavenRemote
}

func (r RemotesMavenApiRemotesMavenMavenCreateRequest) MavenMavenRemote(mavenMavenRemote MavenMavenRemote) RemotesMavenApiRemotesMavenMavenCreateRequest {
	r.mavenMavenRemote = &mavenMavenRemote
	return r
}

func (r RemotesMavenApiRemotesMavenMavenCreateRequest) Execute() (*MavenMavenRemoteResponse, *http.Response, error) {
	return r.ApiService.RemotesMavenMavenCreateExecute(r)
}

/*
RemotesMavenMavenCreate Create a maven remote

A ViewSet for MavenRemote.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RemotesMavenApiRemotesMavenMavenCreateRequest
*/
func (a *RemotesMavenApiService) RemotesMavenMavenCreate(ctx context.Context) RemotesMavenApiRemotesMavenMavenCreateRequest {
	return RemotesMavenApiRemotesMavenMavenCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MavenMavenRemoteResponse
func (a *RemotesMavenApiService) RemotesMavenMavenCreateExecute(r RemotesMavenApiRemotesMavenMavenCreateRequest) (*MavenMavenRemoteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MavenMavenRemoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesMavenApiService.RemotesMavenMavenCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/maven/maven/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mavenMavenRemote == nil {
		return localVarReturnValue, nil, reportError("mavenMavenRemote is required and must be specified")
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
	localVarPostBody = r.mavenMavenRemote
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

type RemotesMavenApiRemotesMavenMavenDeleteRequest struct {
	ctx context.Context
	ApiService *RemotesMavenApiService
	mavenMavenRemoteHref string
}

func (r RemotesMavenApiRemotesMavenMavenDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RemotesMavenMavenDeleteExecute(r)
}

/*
RemotesMavenMavenDelete Delete a maven remote

Trigger an asynchronous delete task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenRemoteHref
 @return RemotesMavenApiRemotesMavenMavenDeleteRequest
*/
func (a *RemotesMavenApiService) RemotesMavenMavenDelete(ctx context.Context, mavenMavenRemoteHref string) RemotesMavenApiRemotesMavenMavenDeleteRequest {
	return RemotesMavenApiRemotesMavenMavenDeleteRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenRemoteHref: mavenMavenRemoteHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RemotesMavenApiService) RemotesMavenMavenDeleteExecute(r RemotesMavenApiRemotesMavenMavenDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesMavenApiService.RemotesMavenMavenDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{maven_maven_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_remote_href"+"}", url.PathEscape(parameterValueToString(r.mavenMavenRemoteHref, "mavenMavenRemoteHref")), -1)
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

type RemotesMavenApiRemotesMavenMavenListRequest struct {
	ctx context.Context
	ApiService *RemotesMavenApiService
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
	offset *int32
	ordering *[]string
	pulpLabelSelect *string
	pulpLastUpdated *time.Time
	pulpLastUpdatedGt *time.Time
	pulpLastUpdatedGte *time.Time
	pulpLastUpdatedLt *time.Time
	pulpLastUpdatedLte *time.Time
	pulpLastUpdatedRange *[]time.Time
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r RemotesMavenApiRemotesMavenMavenListRequest) Limit(limit int32) RemotesMavenApiRemotesMavenMavenListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r RemotesMavenApiRemotesMavenMavenListRequest) Name(name string) RemotesMavenApiRemotesMavenMavenListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r RemotesMavenApiRemotesMavenMavenListRequest) NameContains(nameContains string) RemotesMavenApiRemotesMavenMavenListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r RemotesMavenApiRemotesMavenMavenListRequest) NameIcontains(nameIcontains string) RemotesMavenApiRemotesMavenMavenListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r RemotesMavenApiRemotesMavenMavenListRequest) NameIn(nameIn []string) RemotesMavenApiRemotesMavenMavenListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r RemotesMavenApiRemotesMavenMavenListRequest) NameStartswith(nameStartswith string) RemotesMavenApiRemotesMavenMavenListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r RemotesMavenApiRemotesMavenMavenListRequest) Offset(offset int32) RemotesMavenApiRemotesMavenMavenListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r RemotesMavenApiRemotesMavenMavenListRequest) Ordering(ordering []string) RemotesMavenApiRemotesMavenMavenListRequest {
	r.ordering = &ordering
	return r
}

// Filter labels by search string
func (r RemotesMavenApiRemotesMavenMavenListRequest) PulpLabelSelect(pulpLabelSelect string) RemotesMavenApiRemotesMavenMavenListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Filter results where pulp_last_updated matches value
func (r RemotesMavenApiRemotesMavenMavenListRequest) PulpLastUpdated(pulpLastUpdated time.Time) RemotesMavenApiRemotesMavenMavenListRequest {
	r.pulpLastUpdated = &pulpLastUpdated
	return r
}

// Filter results where pulp_last_updated is greater than value
func (r RemotesMavenApiRemotesMavenMavenListRequest) PulpLastUpdatedGt(pulpLastUpdatedGt time.Time) RemotesMavenApiRemotesMavenMavenListRequest {
	r.pulpLastUpdatedGt = &pulpLastUpdatedGt
	return r
}

// Filter results where pulp_last_updated is greater than or equal to value
func (r RemotesMavenApiRemotesMavenMavenListRequest) PulpLastUpdatedGte(pulpLastUpdatedGte time.Time) RemotesMavenApiRemotesMavenMavenListRequest {
	r.pulpLastUpdatedGte = &pulpLastUpdatedGte
	return r
}

// Filter results where pulp_last_updated is less than value
func (r RemotesMavenApiRemotesMavenMavenListRequest) PulpLastUpdatedLt(pulpLastUpdatedLt time.Time) RemotesMavenApiRemotesMavenMavenListRequest {
	r.pulpLastUpdatedLt = &pulpLastUpdatedLt
	return r
}

// Filter results where pulp_last_updated is less than or equal to value
func (r RemotesMavenApiRemotesMavenMavenListRequest) PulpLastUpdatedLte(pulpLastUpdatedLte time.Time) RemotesMavenApiRemotesMavenMavenListRequest {
	r.pulpLastUpdatedLte = &pulpLastUpdatedLte
	return r
}

// Filter results where pulp_last_updated is between two comma separated values
func (r RemotesMavenApiRemotesMavenMavenListRequest) PulpLastUpdatedRange(pulpLastUpdatedRange []time.Time) RemotesMavenApiRemotesMavenMavenListRequest {
	r.pulpLastUpdatedRange = &pulpLastUpdatedRange
	return r
}

// A list of fields to include in the response.
func (r RemotesMavenApiRemotesMavenMavenListRequest) Fields(fields []string) RemotesMavenApiRemotesMavenMavenListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RemotesMavenApiRemotesMavenMavenListRequest) ExcludeFields(excludeFields []string) RemotesMavenApiRemotesMavenMavenListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RemotesMavenApiRemotesMavenMavenListRequest) Execute() (*PaginatedmavenMavenRemoteResponseList, *http.Response, error) {
	return r.ApiService.RemotesMavenMavenListExecute(r)
}

/*
RemotesMavenMavenList List maven remotes

A ViewSet for MavenRemote.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RemotesMavenApiRemotesMavenMavenListRequest
*/
func (a *RemotesMavenApiService) RemotesMavenMavenList(ctx context.Context) RemotesMavenApiRemotesMavenMavenListRequest {
	return RemotesMavenApiRemotesMavenMavenListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedmavenMavenRemoteResponseList
func (a *RemotesMavenApiService) RemotesMavenMavenListExecute(r RemotesMavenApiRemotesMavenMavenListRequest) (*PaginatedmavenMavenRemoteResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedmavenMavenRemoteResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesMavenApiService.RemotesMavenMavenList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/maven/maven/"
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
	if r.pulpLastUpdated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated", r.pulpLastUpdated, "")
	}
	if r.pulpLastUpdatedGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__gt", r.pulpLastUpdatedGt, "")
	}
	if r.pulpLastUpdatedGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__gte", r.pulpLastUpdatedGte, "")
	}
	if r.pulpLastUpdatedLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__lt", r.pulpLastUpdatedLt, "")
	}
	if r.pulpLastUpdatedLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__lte", r.pulpLastUpdatedLte, "")
	}
	if r.pulpLastUpdatedRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__range", r.pulpLastUpdatedRange, "csv")
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

type RemotesMavenApiRemotesMavenMavenPartialUpdateRequest struct {
	ctx context.Context
	ApiService *RemotesMavenApiService
	mavenMavenRemoteHref string
	patchedmavenMavenRemote *PatchedmavenMavenRemote
}

func (r RemotesMavenApiRemotesMavenMavenPartialUpdateRequest) PatchedmavenMavenRemote(patchedmavenMavenRemote PatchedmavenMavenRemote) RemotesMavenApiRemotesMavenMavenPartialUpdateRequest {
	r.patchedmavenMavenRemote = &patchedmavenMavenRemote
	return r
}

func (r RemotesMavenApiRemotesMavenMavenPartialUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RemotesMavenMavenPartialUpdateExecute(r)
}

/*
RemotesMavenMavenPartialUpdate Update a maven remote

Trigger an asynchronous partial update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenRemoteHref
 @return RemotesMavenApiRemotesMavenMavenPartialUpdateRequest
*/
func (a *RemotesMavenApiService) RemotesMavenMavenPartialUpdate(ctx context.Context, mavenMavenRemoteHref string) RemotesMavenApiRemotesMavenMavenPartialUpdateRequest {
	return RemotesMavenApiRemotesMavenMavenPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenRemoteHref: mavenMavenRemoteHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RemotesMavenApiService) RemotesMavenMavenPartialUpdateExecute(r RemotesMavenApiRemotesMavenMavenPartialUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesMavenApiService.RemotesMavenMavenPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{maven_maven_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_remote_href"+"}", url.PathEscape(parameterValueToString(r.mavenMavenRemoteHref, "mavenMavenRemoteHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedmavenMavenRemote == nil {
		return localVarReturnValue, nil, reportError("patchedmavenMavenRemote is required and must be specified")
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
	localVarPostBody = r.patchedmavenMavenRemote
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

type RemotesMavenApiRemotesMavenMavenReadRequest struct {
	ctx context.Context
	ApiService *RemotesMavenApiService
	mavenMavenRemoteHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r RemotesMavenApiRemotesMavenMavenReadRequest) Fields(fields []string) RemotesMavenApiRemotesMavenMavenReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RemotesMavenApiRemotesMavenMavenReadRequest) ExcludeFields(excludeFields []string) RemotesMavenApiRemotesMavenMavenReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RemotesMavenApiRemotesMavenMavenReadRequest) Execute() (*MavenMavenRemoteResponse, *http.Response, error) {
	return r.ApiService.RemotesMavenMavenReadExecute(r)
}

/*
RemotesMavenMavenRead Inspect a maven remote

A ViewSet for MavenRemote.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenRemoteHref
 @return RemotesMavenApiRemotesMavenMavenReadRequest
*/
func (a *RemotesMavenApiService) RemotesMavenMavenRead(ctx context.Context, mavenMavenRemoteHref string) RemotesMavenApiRemotesMavenMavenReadRequest {
	return RemotesMavenApiRemotesMavenMavenReadRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenRemoteHref: mavenMavenRemoteHref,
	}
}

// Execute executes the request
//  @return MavenMavenRemoteResponse
func (a *RemotesMavenApiService) RemotesMavenMavenReadExecute(r RemotesMavenApiRemotesMavenMavenReadRequest) (*MavenMavenRemoteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MavenMavenRemoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesMavenApiService.RemotesMavenMavenRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{maven_maven_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_remote_href"+"}", url.PathEscape(parameterValueToString(r.mavenMavenRemoteHref, "mavenMavenRemoteHref")), -1)
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

type RemotesMavenApiRemotesMavenMavenUpdateRequest struct {
	ctx context.Context
	ApiService *RemotesMavenApiService
	mavenMavenRemoteHref string
	mavenMavenRemote *MavenMavenRemote
}

func (r RemotesMavenApiRemotesMavenMavenUpdateRequest) MavenMavenRemote(mavenMavenRemote MavenMavenRemote) RemotesMavenApiRemotesMavenMavenUpdateRequest {
	r.mavenMavenRemote = &mavenMavenRemote
	return r
}

func (r RemotesMavenApiRemotesMavenMavenUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RemotesMavenMavenUpdateExecute(r)
}

/*
RemotesMavenMavenUpdate Update a maven remote

Trigger an asynchronous update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenRemoteHref
 @return RemotesMavenApiRemotesMavenMavenUpdateRequest
*/
func (a *RemotesMavenApiService) RemotesMavenMavenUpdate(ctx context.Context, mavenMavenRemoteHref string) RemotesMavenApiRemotesMavenMavenUpdateRequest {
	return RemotesMavenApiRemotesMavenMavenUpdateRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenRemoteHref: mavenMavenRemoteHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RemotesMavenApiService) RemotesMavenMavenUpdateExecute(r RemotesMavenApiRemotesMavenMavenUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesMavenApiService.RemotesMavenMavenUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{maven_maven_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_remote_href"+"}", url.PathEscape(parameterValueToString(r.mavenMavenRemoteHref, "mavenMavenRemoteHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mavenMavenRemote == nil {
		return localVarReturnValue, nil, reportError("mavenMavenRemote is required and must be specified")
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
	localVarPostBody = r.mavenMavenRemote
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
