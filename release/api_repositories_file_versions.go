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


// RepositoriesFileVersionsApiService RepositoriesFileVersionsApi service
type RepositoriesFileVersionsApiService service

type RepositoriesFileVersionsApiRepositoriesFileFileVersionsDeleteRequest struct {
	ctx context.Context
	ApiService *RepositoriesFileVersionsApiService
	fileFileRepositoryVersionHref string
}

func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesFileFileVersionsDeleteExecute(r)
}

/*
RepositoriesFileFileVersionsDelete Delete a repository version

Trigger an asynchronous task to delete a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileFileRepositoryVersionHref
 @return RepositoriesFileVersionsApiRepositoriesFileFileVersionsDeleteRequest
*/
func (a *RepositoriesFileVersionsApiService) RepositoriesFileFileVersionsDelete(ctx context.Context, fileFileRepositoryVersionHref string) RepositoriesFileVersionsApiRepositoriesFileFileVersionsDeleteRequest {
	return RepositoriesFileVersionsApiRepositoriesFileFileVersionsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		fileFileRepositoryVersionHref: fileFileRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesFileVersionsApiService) RepositoriesFileFileVersionsDeleteExecute(r RepositoriesFileVersionsApiRepositoriesFileFileVersionsDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesFileVersionsApiService.RepositoriesFileFileVersionsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{file_file_repository_version_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"file_file_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.fileFileRepositoryVersionHref, "fileFileRepositoryVersionHref")), -1)
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

type RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest struct {
	ctx context.Context
	ApiService *RepositoriesFileVersionsApiService
	fileFileRepositoryHref string
	content *string
	contentIn *string
	limit *int32
	number *int32
	numberGt *int32
	numberGte *int32
	numberLt *int32
	numberLte *int32
	numberRange *[]int32
	offset *int32
	ordering *[]string
	pulpCreated *time.Time
	pulpCreatedGt *time.Time
	pulpCreatedGte *time.Time
	pulpCreatedLt *time.Time
	pulpCreatedLte *time.Time
	pulpCreatedRange *[]time.Time
	fields *[]string
	excludeFields *[]string
}

// Content Unit referenced by HREF
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) Content(content string) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.content = &content
	return r
}

// Content Unit referenced by HREF
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) ContentIn(contentIn string) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.contentIn = &contentIn
	return r
}

// Number of results to return per page.
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) Limit(limit int32) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.limit = &limit
	return r
}

// Filter results where number matches value
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) Number(number int32) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.number = &number
	return r
}

// Filter results where number is greater than value
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) NumberGt(numberGt int32) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.numberGt = &numberGt
	return r
}

// Filter results where number is greater than or equal to value
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) NumberGte(numberGte int32) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.numberGte = &numberGte
	return r
}

// Filter results where number is less than value
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) NumberLt(numberLt int32) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.numberLt = &numberLt
	return r
}

// Filter results where number is less than or equal to value
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) NumberLte(numberLte int32) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.numberLte = &numberLte
	return r
}

// Filter results where number is between two comma separated values
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) NumberRange(numberRange []int32) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.numberRange = &numberRange
	return r
}

// The initial index from which to return the results.
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) Offset(offset int32) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) Ordering(ordering []string) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where pulp_created matches value
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) PulpCreated(pulpCreated time.Time) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.pulpCreated = &pulpCreated
	return r
}

// Filter results where pulp_created is greater than value
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) PulpCreatedGt(pulpCreatedGt time.Time) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.pulpCreatedGt = &pulpCreatedGt
	return r
}

// Filter results where pulp_created is greater than or equal to value
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) PulpCreatedGte(pulpCreatedGte time.Time) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.pulpCreatedGte = &pulpCreatedGte
	return r
}

// Filter results where pulp_created is less than value
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) PulpCreatedLt(pulpCreatedLt time.Time) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.pulpCreatedLt = &pulpCreatedLt
	return r
}

// Filter results where pulp_created is less than or equal to value
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) PulpCreatedLte(pulpCreatedLte time.Time) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.pulpCreatedLte = &pulpCreatedLte
	return r
}

// Filter results where pulp_created is between two comma separated values
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) PulpCreatedRange(pulpCreatedRange []time.Time) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.pulpCreatedRange = &pulpCreatedRange
	return r
}

// A list of fields to include in the response.
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) Fields(fields []string) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) ExcludeFields(excludeFields []string) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) Execute() (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	return r.ApiService.RepositoriesFileFileVersionsListExecute(r)
}

/*
RepositoriesFileFileVersionsList List repository versions


FileRepositoryVersion represents a single file repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileFileRepositoryHref
 @return RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest
*/
func (a *RepositoriesFileVersionsApiService) RepositoriesFileFileVersionsList(ctx context.Context, fileFileRepositoryHref string) RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest {
	return RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest{
		ApiService: a,
		ctx: ctx,
		fileFileRepositoryHref: fileFileRepositoryHref,
	}
}

// Execute executes the request
//  @return PaginatedRepositoryVersionResponseList
func (a *RepositoriesFileVersionsApiService) RepositoriesFileFileVersionsListExecute(r RepositoriesFileVersionsApiRepositoriesFileFileVersionsListRequest) (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRepositoryVersionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesFileVersionsApiService.RepositoriesFileFileVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{file_file_repository_href}versions/"
	localVarPath = strings.Replace(localVarPath, "{"+"file_file_repository_href"+"}", url.PathEscape(parameterValueToString(r.fileFileRepositoryHref, "fileFileRepositoryHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.content != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content", r.content, "")
	}
	if r.contentIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content__in", r.contentIn, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "")
	}
	if r.numberGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__gt", r.numberGt, "")
	}
	if r.numberGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__gte", r.numberGte, "")
	}
	if r.numberLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__lt", r.numberLt, "")
	}
	if r.numberLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__lte", r.numberLte, "")
	}
	if r.numberRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__range", r.numberRange, "csv")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.pulpCreated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created", r.pulpCreated, "")
	}
	if r.pulpCreatedGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__gt", r.pulpCreatedGt, "")
	}
	if r.pulpCreatedGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__gte", r.pulpCreatedGte, "")
	}
	if r.pulpCreatedLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__lt", r.pulpCreatedLt, "")
	}
	if r.pulpCreatedLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__lte", r.pulpCreatedLte, "")
	}
	if r.pulpCreatedRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__range", r.pulpCreatedRange, "csv")
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

type RepositoriesFileVersionsApiRepositoriesFileFileVersionsReadRequest struct {
	ctx context.Context
	ApiService *RepositoriesFileVersionsApiService
	fileFileRepositoryVersionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsReadRequest) Fields(fields []string) RepositoriesFileVersionsApiRepositoriesFileFileVersionsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsReadRequest) ExcludeFields(excludeFields []string) RepositoriesFileVersionsApiRepositoriesFileFileVersionsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsReadRequest) Execute() (*RepositoryVersionResponse, *http.Response, error) {
	return r.ApiService.RepositoriesFileFileVersionsReadExecute(r)
}

/*
RepositoriesFileFileVersionsRead Inspect a repository version


FileRepositoryVersion represents a single file repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileFileRepositoryVersionHref
 @return RepositoriesFileVersionsApiRepositoriesFileFileVersionsReadRequest
*/
func (a *RepositoriesFileVersionsApiService) RepositoriesFileFileVersionsRead(ctx context.Context, fileFileRepositoryVersionHref string) RepositoriesFileVersionsApiRepositoriesFileFileVersionsReadRequest {
	return RepositoriesFileVersionsApiRepositoriesFileFileVersionsReadRequest{
		ApiService: a,
		ctx: ctx,
		fileFileRepositoryVersionHref: fileFileRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return RepositoryVersionResponse
func (a *RepositoriesFileVersionsApiService) RepositoriesFileFileVersionsReadExecute(r RepositoriesFileVersionsApiRepositoriesFileFileVersionsReadRequest) (*RepositoryVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepositoryVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesFileVersionsApiService.RepositoriesFileFileVersionsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{file_file_repository_version_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"file_file_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.fileFileRepositoryVersionHref, "fileFileRepositoryVersionHref")), -1)
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

type RepositoriesFileVersionsApiRepositoriesFileFileVersionsRepairRequest struct {
	ctx context.Context
	ApiService *RepositoriesFileVersionsApiService
	fileFileRepositoryVersionHref string
	repair *Repair
}

func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsRepairRequest) Repair(repair Repair) RepositoriesFileVersionsApiRepositoriesFileFileVersionsRepairRequest {
	r.repair = &repair
	return r
}

func (r RepositoriesFileVersionsApiRepositoriesFileFileVersionsRepairRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesFileFileVersionsRepairExecute(r)
}

/*
RepositoriesFileFileVersionsRepair Method for RepositoriesFileFileVersionsRepair

Trigger an asynchronous task to repair a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileFileRepositoryVersionHref
 @return RepositoriesFileVersionsApiRepositoriesFileFileVersionsRepairRequest
*/
func (a *RepositoriesFileVersionsApiService) RepositoriesFileFileVersionsRepair(ctx context.Context, fileFileRepositoryVersionHref string) RepositoriesFileVersionsApiRepositoriesFileFileVersionsRepairRequest {
	return RepositoriesFileVersionsApiRepositoriesFileFileVersionsRepairRequest{
		ApiService: a,
		ctx: ctx,
		fileFileRepositoryVersionHref: fileFileRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesFileVersionsApiService) RepositoriesFileFileVersionsRepairExecute(r RepositoriesFileVersionsApiRepositoriesFileFileVersionsRepairRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesFileVersionsApiService.RepositoriesFileFileVersionsRepair")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{file_file_repository_version_href}repair/"
	localVarPath = strings.Replace(localVarPath, "{"+"file_file_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.fileFileRepositoryVersionHref, "fileFileRepositoryVersionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.repair == nil {
		return localVarReturnValue, nil, reportError("repair is required and must be specified")
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
	localVarPostBody = r.repair
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
