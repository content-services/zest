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
	"os"
	"time"
	"reflect"
)


// RemotesPythonApiService RemotesPythonApi service
type RemotesPythonApiService service

type RemotesPythonApiRemotesPythonPythonCreateRequest struct {
	ctx context.Context
	ApiService *RemotesPythonApiService
	pythonPythonRemote *PythonPythonRemote
}

func (r RemotesPythonApiRemotesPythonPythonCreateRequest) PythonPythonRemote(pythonPythonRemote PythonPythonRemote) RemotesPythonApiRemotesPythonPythonCreateRequest {
	r.pythonPythonRemote = &pythonPythonRemote
	return r
}

func (r RemotesPythonApiRemotesPythonPythonCreateRequest) Execute() (*PythonPythonRemoteResponse, *http.Response, error) {
	return r.ApiService.RemotesPythonPythonCreateExecute(r)
}

/*
RemotesPythonPythonCreate Create a python remote


Python Remotes are representations of an external repository of Python content, eg.
PyPI.  Fields include upstream repository config. Python Remotes are also used to `sync` from
upstream repositories, and contains sync settings.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RemotesPythonApiRemotesPythonPythonCreateRequest
*/
func (a *RemotesPythonApiService) RemotesPythonPythonCreate(ctx context.Context) RemotesPythonApiRemotesPythonPythonCreateRequest {
	return RemotesPythonApiRemotesPythonPythonCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PythonPythonRemoteResponse
func (a *RemotesPythonApiService) RemotesPythonPythonCreateExecute(r RemotesPythonApiRemotesPythonPythonCreateRequest) (*PythonPythonRemoteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PythonPythonRemoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesPythonApiService.RemotesPythonPythonCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/python/python/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pythonPythonRemote == nil {
		return localVarReturnValue, nil, reportError("pythonPythonRemote is required and must be specified")
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
	localVarPostBody = r.pythonPythonRemote
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

type RemotesPythonApiRemotesPythonPythonDeleteRequest struct {
	ctx context.Context
	ApiService *RemotesPythonApiService
	pythonPythonRemoteHref string
}

func (r RemotesPythonApiRemotesPythonPythonDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RemotesPythonPythonDeleteExecute(r)
}

/*
RemotesPythonPythonDelete Delete a python remote

Trigger an asynchronous delete task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pythonPythonRemoteHref
 @return RemotesPythonApiRemotesPythonPythonDeleteRequest
*/
func (a *RemotesPythonApiService) RemotesPythonPythonDelete(ctx context.Context, pythonPythonRemoteHref string) RemotesPythonApiRemotesPythonPythonDeleteRequest {
	return RemotesPythonApiRemotesPythonPythonDeleteRequest{
		ApiService: a,
		ctx: ctx,
		pythonPythonRemoteHref: pythonPythonRemoteHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RemotesPythonApiService) RemotesPythonPythonDeleteExecute(r RemotesPythonApiRemotesPythonPythonDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesPythonApiService.RemotesPythonPythonDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{python_python_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"python_python_remote_href"+"}", url.PathEscape(parameterValueToString(r.pythonPythonRemoteHref, "pythonPythonRemoteHref")), -1)
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

type RemotesPythonApiRemotesPythonPythonFromBandersnatchRequest struct {
	ctx context.Context
	ApiService *RemotesPythonApiService
	config *os.File
	name *string
	policy *Policy762Enum
}

// A Bandersnatch config that may be used to construct a Python Remote.
func (r RemotesPythonApiRemotesPythonPythonFromBandersnatchRequest) Config(config *os.File) RemotesPythonApiRemotesPythonPythonFromBandersnatchRequest {
	r.config = config
	return r
}

// A unique name for this remote
func (r RemotesPythonApiRemotesPythonPythonFromBandersnatchRequest) Name(name string) RemotesPythonApiRemotesPythonPythonFromBandersnatchRequest {
	r.name = &name
	return r
}

func (r RemotesPythonApiRemotesPythonPythonFromBandersnatchRequest) Policy(policy Policy762Enum) RemotesPythonApiRemotesPythonPythonFromBandersnatchRequest {
	r.policy = &policy
	return r
}

func (r RemotesPythonApiRemotesPythonPythonFromBandersnatchRequest) Execute() (*PythonPythonRemoteResponse, *http.Response, error) {
	return r.ApiService.RemotesPythonPythonFromBandersnatchExecute(r)
}

/*
RemotesPythonPythonFromBandersnatch Create from Bandersnatch


Takes the fields specified in the Bandersnatch config and creates a Python Remote from it.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RemotesPythonApiRemotesPythonPythonFromBandersnatchRequest
*/
func (a *RemotesPythonApiService) RemotesPythonPythonFromBandersnatch(ctx context.Context) RemotesPythonApiRemotesPythonPythonFromBandersnatchRequest {
	return RemotesPythonApiRemotesPythonPythonFromBandersnatchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PythonPythonRemoteResponse
func (a *RemotesPythonApiService) RemotesPythonPythonFromBandersnatchExecute(r RemotesPythonApiRemotesPythonPythonFromBandersnatchRequest) (*PythonPythonRemoteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PythonPythonRemoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesPythonApiService.RemotesPythonPythonFromBandersnatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/python/python/from_bandersnatch/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.config == nil {
		return localVarReturnValue, nil, reportError("config is required and must be specified")
	}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}
	if strlen(*r.name) < 1 {
		return localVarReturnValue, nil, reportError("name must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data", "application/x-www-form-urlencoded"}

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
	var configLocalVarFormFileName string
	var configLocalVarFileName     string
	var configLocalVarFileBytes    []byte

	configLocalVarFormFileName = "config"


	configLocalVarFile := r.config

	if configLocalVarFile != nil {
		fbs, _ := io.ReadAll(configLocalVarFile)

		configLocalVarFileBytes = fbs
		configLocalVarFileName = configLocalVarFile.Name()
		configLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: configLocalVarFileBytes, fileName: configLocalVarFileName, formFileName: configLocalVarFormFileName})
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "")
	if r.policy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "policy", r.policy, "")
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

type RemotesPythonApiRemotesPythonPythonListRequest struct {
	ctx context.Context
	ApiService *RemotesPythonApiService
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
func (r RemotesPythonApiRemotesPythonPythonListRequest) Limit(limit int32) RemotesPythonApiRemotesPythonPythonListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r RemotesPythonApiRemotesPythonPythonListRequest) Name(name string) RemotesPythonApiRemotesPythonPythonListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r RemotesPythonApiRemotesPythonPythonListRequest) NameContains(nameContains string) RemotesPythonApiRemotesPythonPythonListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r RemotesPythonApiRemotesPythonPythonListRequest) NameIcontains(nameIcontains string) RemotesPythonApiRemotesPythonPythonListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r RemotesPythonApiRemotesPythonPythonListRequest) NameIn(nameIn []string) RemotesPythonApiRemotesPythonPythonListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r RemotesPythonApiRemotesPythonPythonListRequest) NameStartswith(nameStartswith string) RemotesPythonApiRemotesPythonPythonListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r RemotesPythonApiRemotesPythonPythonListRequest) Offset(offset int32) RemotesPythonApiRemotesPythonPythonListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r RemotesPythonApiRemotesPythonPythonListRequest) Ordering(ordering []string) RemotesPythonApiRemotesPythonPythonListRequest {
	r.ordering = &ordering
	return r
}

// Filter labels by search string
func (r RemotesPythonApiRemotesPythonPythonListRequest) PulpLabelSelect(pulpLabelSelect string) RemotesPythonApiRemotesPythonPythonListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Filter results where pulp_last_updated matches value
func (r RemotesPythonApiRemotesPythonPythonListRequest) PulpLastUpdated(pulpLastUpdated time.Time) RemotesPythonApiRemotesPythonPythonListRequest {
	r.pulpLastUpdated = &pulpLastUpdated
	return r
}

// Filter results where pulp_last_updated is greater than value
func (r RemotesPythonApiRemotesPythonPythonListRequest) PulpLastUpdatedGt(pulpLastUpdatedGt time.Time) RemotesPythonApiRemotesPythonPythonListRequest {
	r.pulpLastUpdatedGt = &pulpLastUpdatedGt
	return r
}

// Filter results where pulp_last_updated is greater than or equal to value
func (r RemotesPythonApiRemotesPythonPythonListRequest) PulpLastUpdatedGte(pulpLastUpdatedGte time.Time) RemotesPythonApiRemotesPythonPythonListRequest {
	r.pulpLastUpdatedGte = &pulpLastUpdatedGte
	return r
}

// Filter results where pulp_last_updated is less than value
func (r RemotesPythonApiRemotesPythonPythonListRequest) PulpLastUpdatedLt(pulpLastUpdatedLt time.Time) RemotesPythonApiRemotesPythonPythonListRequest {
	r.pulpLastUpdatedLt = &pulpLastUpdatedLt
	return r
}

// Filter results where pulp_last_updated is less than or equal to value
func (r RemotesPythonApiRemotesPythonPythonListRequest) PulpLastUpdatedLte(pulpLastUpdatedLte time.Time) RemotesPythonApiRemotesPythonPythonListRequest {
	r.pulpLastUpdatedLte = &pulpLastUpdatedLte
	return r
}

// Filter results where pulp_last_updated is between two comma separated values
func (r RemotesPythonApiRemotesPythonPythonListRequest) PulpLastUpdatedRange(pulpLastUpdatedRange []time.Time) RemotesPythonApiRemotesPythonPythonListRequest {
	r.pulpLastUpdatedRange = &pulpLastUpdatedRange
	return r
}

// A list of fields to include in the response.
func (r RemotesPythonApiRemotesPythonPythonListRequest) Fields(fields []string) RemotesPythonApiRemotesPythonPythonListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RemotesPythonApiRemotesPythonPythonListRequest) ExcludeFields(excludeFields []string) RemotesPythonApiRemotesPythonPythonListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RemotesPythonApiRemotesPythonPythonListRequest) Execute() (*PaginatedpythonPythonRemoteResponseList, *http.Response, error) {
	return r.ApiService.RemotesPythonPythonListExecute(r)
}

/*
RemotesPythonPythonList List python remotes


Python Remotes are representations of an external repository of Python content, eg.
PyPI.  Fields include upstream repository config. Python Remotes are also used to `sync` from
upstream repositories, and contains sync settings.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RemotesPythonApiRemotesPythonPythonListRequest
*/
func (a *RemotesPythonApiService) RemotesPythonPythonList(ctx context.Context) RemotesPythonApiRemotesPythonPythonListRequest {
	return RemotesPythonApiRemotesPythonPythonListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedpythonPythonRemoteResponseList
func (a *RemotesPythonApiService) RemotesPythonPythonListExecute(r RemotesPythonApiRemotesPythonPythonListRequest) (*PaginatedpythonPythonRemoteResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedpythonPythonRemoteResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesPythonApiService.RemotesPythonPythonList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/remotes/python/python/"
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

type RemotesPythonApiRemotesPythonPythonPartialUpdateRequest struct {
	ctx context.Context
	ApiService *RemotesPythonApiService
	pythonPythonRemoteHref string
	patchedpythonPythonRemote *PatchedpythonPythonRemote
}

func (r RemotesPythonApiRemotesPythonPythonPartialUpdateRequest) PatchedpythonPythonRemote(patchedpythonPythonRemote PatchedpythonPythonRemote) RemotesPythonApiRemotesPythonPythonPartialUpdateRequest {
	r.patchedpythonPythonRemote = &patchedpythonPythonRemote
	return r
}

func (r RemotesPythonApiRemotesPythonPythonPartialUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RemotesPythonPythonPartialUpdateExecute(r)
}

/*
RemotesPythonPythonPartialUpdate Update a python remote

Trigger an asynchronous partial update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pythonPythonRemoteHref
 @return RemotesPythonApiRemotesPythonPythonPartialUpdateRequest
*/
func (a *RemotesPythonApiService) RemotesPythonPythonPartialUpdate(ctx context.Context, pythonPythonRemoteHref string) RemotesPythonApiRemotesPythonPythonPartialUpdateRequest {
	return RemotesPythonApiRemotesPythonPythonPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		pythonPythonRemoteHref: pythonPythonRemoteHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RemotesPythonApiService) RemotesPythonPythonPartialUpdateExecute(r RemotesPythonApiRemotesPythonPythonPartialUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesPythonApiService.RemotesPythonPythonPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{python_python_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"python_python_remote_href"+"}", url.PathEscape(parameterValueToString(r.pythonPythonRemoteHref, "pythonPythonRemoteHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedpythonPythonRemote == nil {
		return localVarReturnValue, nil, reportError("patchedpythonPythonRemote is required and must be specified")
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
	localVarPostBody = r.patchedpythonPythonRemote
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

type RemotesPythonApiRemotesPythonPythonReadRequest struct {
	ctx context.Context
	ApiService *RemotesPythonApiService
	pythonPythonRemoteHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r RemotesPythonApiRemotesPythonPythonReadRequest) Fields(fields []string) RemotesPythonApiRemotesPythonPythonReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RemotesPythonApiRemotesPythonPythonReadRequest) ExcludeFields(excludeFields []string) RemotesPythonApiRemotesPythonPythonReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RemotesPythonApiRemotesPythonPythonReadRequest) Execute() (*PythonPythonRemoteResponse, *http.Response, error) {
	return r.ApiService.RemotesPythonPythonReadExecute(r)
}

/*
RemotesPythonPythonRead Inspect a python remote


Python Remotes are representations of an external repository of Python content, eg.
PyPI.  Fields include upstream repository config. Python Remotes are also used to `sync` from
upstream repositories, and contains sync settings.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pythonPythonRemoteHref
 @return RemotesPythonApiRemotesPythonPythonReadRequest
*/
func (a *RemotesPythonApiService) RemotesPythonPythonRead(ctx context.Context, pythonPythonRemoteHref string) RemotesPythonApiRemotesPythonPythonReadRequest {
	return RemotesPythonApiRemotesPythonPythonReadRequest{
		ApiService: a,
		ctx: ctx,
		pythonPythonRemoteHref: pythonPythonRemoteHref,
	}
}

// Execute executes the request
//  @return PythonPythonRemoteResponse
func (a *RemotesPythonApiService) RemotesPythonPythonReadExecute(r RemotesPythonApiRemotesPythonPythonReadRequest) (*PythonPythonRemoteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PythonPythonRemoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesPythonApiService.RemotesPythonPythonRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{python_python_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"python_python_remote_href"+"}", url.PathEscape(parameterValueToString(r.pythonPythonRemoteHref, "pythonPythonRemoteHref")), -1)
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

type RemotesPythonApiRemotesPythonPythonUpdateRequest struct {
	ctx context.Context
	ApiService *RemotesPythonApiService
	pythonPythonRemoteHref string
	pythonPythonRemote *PythonPythonRemote
}

func (r RemotesPythonApiRemotesPythonPythonUpdateRequest) PythonPythonRemote(pythonPythonRemote PythonPythonRemote) RemotesPythonApiRemotesPythonPythonUpdateRequest {
	r.pythonPythonRemote = &pythonPythonRemote
	return r
}

func (r RemotesPythonApiRemotesPythonPythonUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RemotesPythonPythonUpdateExecute(r)
}

/*
RemotesPythonPythonUpdate Update a python remote

Trigger an asynchronous update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pythonPythonRemoteHref
 @return RemotesPythonApiRemotesPythonPythonUpdateRequest
*/
func (a *RemotesPythonApiService) RemotesPythonPythonUpdate(ctx context.Context, pythonPythonRemoteHref string) RemotesPythonApiRemotesPythonPythonUpdateRequest {
	return RemotesPythonApiRemotesPythonPythonUpdateRequest{
		ApiService: a,
		ctx: ctx,
		pythonPythonRemoteHref: pythonPythonRemoteHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RemotesPythonApiService) RemotesPythonPythonUpdateExecute(r RemotesPythonApiRemotesPythonPythonUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesPythonApiService.RemotesPythonPythonUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{python_python_remote_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"python_python_remote_href"+"}", url.PathEscape(parameterValueToString(r.pythonPythonRemoteHref, "pythonPythonRemoteHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pythonPythonRemote == nil {
		return localVarReturnValue, nil, reportError("pythonPythonRemote is required and must be specified")
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
	localVarPostBody = r.pythonPythonRemote
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
