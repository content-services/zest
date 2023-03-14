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
	"time"
	"reflect"
)


// RemotesGitApiService RemotesGitApi service
type RemotesGitApiService service

type RemotesGitApiRemotesAnsibleGitCreateRequest struct {
	ctx context.Context
	ApiService *RemotesGitApiService
	ansibleGitRemote *AnsibleGitRemote
}

func (r RemotesGitApiRemotesAnsibleGitCreateRequest) AnsibleGitRemote(ansibleGitRemote AnsibleGitRemote) RemotesGitApiRemotesAnsibleGitCreateRequest {
	r.ansibleGitRemote = &ansibleGitRemote
	return r
}

func (r RemotesGitApiRemotesAnsibleGitCreateRequest) Execute() (*AnsibleGitRemoteResponse, *http.Response, error) {
	return r.ApiService.RemotesAnsibleGitCreateExecute(r)
}

/*
RemotesAnsibleGitCreate Create a git remote

ViewSet for Ansible Remotes.

This is a tech preview feature. The functionality may change in the future.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RemotesGitApiRemotesAnsibleGitCreateRequest
*/
func (a *RemotesGitApiService) RemotesAnsibleGitCreate(ctx context.Context) RemotesGitApiRemotesAnsibleGitCreateRequest {
	return RemotesGitApiRemotesAnsibleGitCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AnsibleGitRemoteResponse
func (a *RemotesGitApiService) RemotesAnsibleGitCreateExecute(r RemotesGitApiRemotesAnsibleGitCreateRequest) (*AnsibleGitRemoteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AnsibleGitRemoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesGitApiService.RemotesAnsibleGitCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/ansible/git/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ansibleGitRemote == nil {
		return localVarReturnValue, nil, reportError("ansibleGitRemote is required and must be specified")
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
	localVarPostBody = r.ansibleGitRemote
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

type RemotesGitApiRemotesAnsibleGitDeleteRequest struct {
	ctx context.Context
	ApiService *RemotesGitApiService
	ansibleGitRemoteHref string
}

func (r RemotesGitApiRemotesAnsibleGitDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RemotesAnsibleGitDeleteExecute(r)
}

/*
RemotesAnsibleGitDelete Delete a git remote

Trigger an asynchronous delete task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleGitRemoteHref
 @return RemotesGitApiRemotesAnsibleGitDeleteRequest
*/
func (a *RemotesGitApiService) RemotesAnsibleGitDelete(ctx context.Context, ansibleGitRemoteHref string) RemotesGitApiRemotesAnsibleGitDeleteRequest {
	return RemotesGitApiRemotesAnsibleGitDeleteRequest{
		ApiService: a,
		ctx: ctx,
		ansibleGitRemoteHref: ansibleGitRemoteHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RemotesGitApiService) RemotesAnsibleGitDeleteExecute(r RemotesGitApiRemotesAnsibleGitDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesGitApiService.RemotesAnsibleGitDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ansible_git_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_git_remote_href"+"}", url.PathEscape(parameterValueToString(r.ansibleGitRemoteHref, "ansibleGitRemoteHref")), -1)
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

type RemotesGitApiRemotesAnsibleGitListRequest struct {
	ctx context.Context
	ApiService *RemotesGitApiService
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
func (r RemotesGitApiRemotesAnsibleGitListRequest) Limit(limit int32) RemotesGitApiRemotesAnsibleGitListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r RemotesGitApiRemotesAnsibleGitListRequest) Name(name string) RemotesGitApiRemotesAnsibleGitListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r RemotesGitApiRemotesAnsibleGitListRequest) NameContains(nameContains string) RemotesGitApiRemotesAnsibleGitListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r RemotesGitApiRemotesAnsibleGitListRequest) NameIcontains(nameIcontains string) RemotesGitApiRemotesAnsibleGitListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r RemotesGitApiRemotesAnsibleGitListRequest) NameIn(nameIn []string) RemotesGitApiRemotesAnsibleGitListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r RemotesGitApiRemotesAnsibleGitListRequest) NameStartswith(nameStartswith string) RemotesGitApiRemotesAnsibleGitListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r RemotesGitApiRemotesAnsibleGitListRequest) Offset(offset int32) RemotesGitApiRemotesAnsibleGitListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r RemotesGitApiRemotesAnsibleGitListRequest) Ordering(ordering []string) RemotesGitApiRemotesAnsibleGitListRequest {
	r.ordering = &ordering
	return r
}

// Filter labels by search string
func (r RemotesGitApiRemotesAnsibleGitListRequest) PulpLabelSelect(pulpLabelSelect string) RemotesGitApiRemotesAnsibleGitListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Filter results where pulp_last_updated matches value
func (r RemotesGitApiRemotesAnsibleGitListRequest) PulpLastUpdated(pulpLastUpdated time.Time) RemotesGitApiRemotesAnsibleGitListRequest {
	r.pulpLastUpdated = &pulpLastUpdated
	return r
}

// Filter results where pulp_last_updated is greater than value
func (r RemotesGitApiRemotesAnsibleGitListRequest) PulpLastUpdatedGt(pulpLastUpdatedGt time.Time) RemotesGitApiRemotesAnsibleGitListRequest {
	r.pulpLastUpdatedGt = &pulpLastUpdatedGt
	return r
}

// Filter results where pulp_last_updated is greater than or equal to value
func (r RemotesGitApiRemotesAnsibleGitListRequest) PulpLastUpdatedGte(pulpLastUpdatedGte time.Time) RemotesGitApiRemotesAnsibleGitListRequest {
	r.pulpLastUpdatedGte = &pulpLastUpdatedGte
	return r
}

// Filter results where pulp_last_updated is less than value
func (r RemotesGitApiRemotesAnsibleGitListRequest) PulpLastUpdatedLt(pulpLastUpdatedLt time.Time) RemotesGitApiRemotesAnsibleGitListRequest {
	r.pulpLastUpdatedLt = &pulpLastUpdatedLt
	return r
}

// Filter results where pulp_last_updated is less than or equal to value
func (r RemotesGitApiRemotesAnsibleGitListRequest) PulpLastUpdatedLte(pulpLastUpdatedLte time.Time) RemotesGitApiRemotesAnsibleGitListRequest {
	r.pulpLastUpdatedLte = &pulpLastUpdatedLte
	return r
}

// Filter results where pulp_last_updated is between two comma separated values
func (r RemotesGitApiRemotesAnsibleGitListRequest) PulpLastUpdatedRange(pulpLastUpdatedRange []time.Time) RemotesGitApiRemotesAnsibleGitListRequest {
	r.pulpLastUpdatedRange = &pulpLastUpdatedRange
	return r
}

// A list of fields to include in the response.
func (r RemotesGitApiRemotesAnsibleGitListRequest) Fields(fields []string) RemotesGitApiRemotesAnsibleGitListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RemotesGitApiRemotesAnsibleGitListRequest) ExcludeFields(excludeFields []string) RemotesGitApiRemotesAnsibleGitListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RemotesGitApiRemotesAnsibleGitListRequest) Execute() (*PaginatedansibleGitRemoteResponseList, *http.Response, error) {
	return r.ApiService.RemotesAnsibleGitListExecute(r)
}

/*
RemotesAnsibleGitList List git remotes

ViewSet for Ansible Remotes.

This is a tech preview feature. The functionality may change in the future.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RemotesGitApiRemotesAnsibleGitListRequest
*/
func (a *RemotesGitApiService) RemotesAnsibleGitList(ctx context.Context) RemotesGitApiRemotesAnsibleGitListRequest {
	return RemotesGitApiRemotesAnsibleGitListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedansibleGitRemoteResponseList
func (a *RemotesGitApiService) RemotesAnsibleGitListExecute(r RemotesGitApiRemotesAnsibleGitListRequest) (*PaginatedansibleGitRemoteResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedansibleGitRemoteResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesGitApiService.RemotesAnsibleGitList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/ansible/git/"
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

type RemotesGitApiRemotesAnsibleGitPartialUpdateRequest struct {
	ctx context.Context
	ApiService *RemotesGitApiService
	ansibleGitRemoteHref string
	patchedansibleGitRemote *PatchedansibleGitRemote
}

func (r RemotesGitApiRemotesAnsibleGitPartialUpdateRequest) PatchedansibleGitRemote(patchedansibleGitRemote PatchedansibleGitRemote) RemotesGitApiRemotesAnsibleGitPartialUpdateRequest {
	r.patchedansibleGitRemote = &patchedansibleGitRemote
	return r
}

func (r RemotesGitApiRemotesAnsibleGitPartialUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RemotesAnsibleGitPartialUpdateExecute(r)
}

/*
RemotesAnsibleGitPartialUpdate Update a git remote

Trigger an asynchronous partial update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleGitRemoteHref
 @return RemotesGitApiRemotesAnsibleGitPartialUpdateRequest
*/
func (a *RemotesGitApiService) RemotesAnsibleGitPartialUpdate(ctx context.Context, ansibleGitRemoteHref string) RemotesGitApiRemotesAnsibleGitPartialUpdateRequest {
	return RemotesGitApiRemotesAnsibleGitPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		ansibleGitRemoteHref: ansibleGitRemoteHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RemotesGitApiService) RemotesAnsibleGitPartialUpdateExecute(r RemotesGitApiRemotesAnsibleGitPartialUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesGitApiService.RemotesAnsibleGitPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ansible_git_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_git_remote_href"+"}", url.PathEscape(parameterValueToString(r.ansibleGitRemoteHref, "ansibleGitRemoteHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedansibleGitRemote == nil {
		return localVarReturnValue, nil, reportError("patchedansibleGitRemote is required and must be specified")
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
	localVarPostBody = r.patchedansibleGitRemote
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

type RemotesGitApiRemotesAnsibleGitReadRequest struct {
	ctx context.Context
	ApiService *RemotesGitApiService
	ansibleGitRemoteHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r RemotesGitApiRemotesAnsibleGitReadRequest) Fields(fields []string) RemotesGitApiRemotesAnsibleGitReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RemotesGitApiRemotesAnsibleGitReadRequest) ExcludeFields(excludeFields []string) RemotesGitApiRemotesAnsibleGitReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RemotesGitApiRemotesAnsibleGitReadRequest) Execute() (*AnsibleGitRemoteResponse, *http.Response, error) {
	return r.ApiService.RemotesAnsibleGitReadExecute(r)
}

/*
RemotesAnsibleGitRead Inspect a git remote

ViewSet for Ansible Remotes.

This is a tech preview feature. The functionality may change in the future.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleGitRemoteHref
 @return RemotesGitApiRemotesAnsibleGitReadRequest
*/
func (a *RemotesGitApiService) RemotesAnsibleGitRead(ctx context.Context, ansibleGitRemoteHref string) RemotesGitApiRemotesAnsibleGitReadRequest {
	return RemotesGitApiRemotesAnsibleGitReadRequest{
		ApiService: a,
		ctx: ctx,
		ansibleGitRemoteHref: ansibleGitRemoteHref,
	}
}

// Execute executes the request
//  @return AnsibleGitRemoteResponse
func (a *RemotesGitApiService) RemotesAnsibleGitReadExecute(r RemotesGitApiRemotesAnsibleGitReadRequest) (*AnsibleGitRemoteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AnsibleGitRemoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesGitApiService.RemotesAnsibleGitRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ansible_git_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_git_remote_href"+"}", url.PathEscape(parameterValueToString(r.ansibleGitRemoteHref, "ansibleGitRemoteHref")), -1)
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

type RemotesGitApiRemotesAnsibleGitUpdateRequest struct {
	ctx context.Context
	ApiService *RemotesGitApiService
	ansibleGitRemoteHref string
	ansibleGitRemote *AnsibleGitRemote
}

func (r RemotesGitApiRemotesAnsibleGitUpdateRequest) AnsibleGitRemote(ansibleGitRemote AnsibleGitRemote) RemotesGitApiRemotesAnsibleGitUpdateRequest {
	r.ansibleGitRemote = &ansibleGitRemote
	return r
}

func (r RemotesGitApiRemotesAnsibleGitUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RemotesAnsibleGitUpdateExecute(r)
}

/*
RemotesAnsibleGitUpdate Update a git remote

Trigger an asynchronous update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleGitRemoteHref
 @return RemotesGitApiRemotesAnsibleGitUpdateRequest
*/
func (a *RemotesGitApiService) RemotesAnsibleGitUpdate(ctx context.Context, ansibleGitRemoteHref string) RemotesGitApiRemotesAnsibleGitUpdateRequest {
	return RemotesGitApiRemotesAnsibleGitUpdateRequest{
		ApiService: a,
		ctx: ctx,
		ansibleGitRemoteHref: ansibleGitRemoteHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RemotesGitApiService) RemotesAnsibleGitUpdateExecute(r RemotesGitApiRemotesAnsibleGitUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesGitApiService.RemotesAnsibleGitUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ansible_git_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_git_remote_href"+"}", url.PathEscape(parameterValueToString(r.ansibleGitRemoteHref, "ansibleGitRemoteHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ansibleGitRemote == nil {
		return localVarReturnValue, nil, reportError("ansibleGitRemote is required and must be specified")
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
	localVarPostBody = r.ansibleGitRemote
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