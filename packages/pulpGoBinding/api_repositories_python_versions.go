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
	"time"
	"reflect"
)


// RepositoriesPythonVersionsApiService RepositoriesPythonVersionsApi service
type RepositoriesPythonVersionsApiService service

type RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsDeleteRequest struct {
	ctx context.Context
	ApiService *RepositoriesPythonVersionsApiService
	pythonPythonRepositoryVersionHref string
}

func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesPythonPythonVersionsDeleteExecute(r)
}

/*
RepositoriesPythonPythonVersionsDelete Delete a repository version

Trigger an asynchronous task to delete a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pythonPythonRepositoryVersionHref
 @return RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsDeleteRequest
*/
func (a *RepositoriesPythonVersionsApiService) RepositoriesPythonPythonVersionsDelete(ctx context.Context, pythonPythonRepositoryVersionHref string) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsDeleteRequest {
	return RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		pythonPythonRepositoryVersionHref: pythonPythonRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesPythonVersionsApiService) RepositoriesPythonPythonVersionsDeleteExecute(r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesPythonVersionsApiService.RepositoriesPythonPythonVersionsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{python_python_repository_version_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"python_python_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.pythonPythonRepositoryVersionHref, "pythonPythonRepositoryVersionHref")), -1)
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

type RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest struct {
	ctx context.Context
	ApiService *RepositoriesPythonVersionsApiService
	pythonPythonRepositoryHref string
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
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) Content(content string) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.content = &content
	return r
}

// Content Unit referenced by HREF
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) ContentIn(contentIn string) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.contentIn = &contentIn
	return r
}

// Number of results to return per page.
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) Limit(limit int32) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.limit = &limit
	return r
}

// Filter results where number matches value
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) Number(number int32) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.number = &number
	return r
}

// Filter results where number is greater than value
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) NumberGt(numberGt int32) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.numberGt = &numberGt
	return r
}

// Filter results where number is greater than or equal to value
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) NumberGte(numberGte int32) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.numberGte = &numberGte
	return r
}

// Filter results where number is less than value
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) NumberLt(numberLt int32) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.numberLt = &numberLt
	return r
}

// Filter results where number is less than or equal to value
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) NumberLte(numberLte int32) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.numberLte = &numberLte
	return r
}

// Filter results where number is between two comma separated values
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) NumberRange(numberRange []int32) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.numberRange = &numberRange
	return r
}

// The initial index from which to return the results.
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) Offset(offset int32) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) Ordering(ordering []string) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where pulp_created matches value
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) PulpCreated(pulpCreated time.Time) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.pulpCreated = &pulpCreated
	return r
}

// Filter results where pulp_created is greater than value
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) PulpCreatedGt(pulpCreatedGt time.Time) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.pulpCreatedGt = &pulpCreatedGt
	return r
}

// Filter results where pulp_created is greater than or equal to value
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) PulpCreatedGte(pulpCreatedGte time.Time) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.pulpCreatedGte = &pulpCreatedGte
	return r
}

// Filter results where pulp_created is less than value
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) PulpCreatedLt(pulpCreatedLt time.Time) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.pulpCreatedLt = &pulpCreatedLt
	return r
}

// Filter results where pulp_created is less than or equal to value
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) PulpCreatedLte(pulpCreatedLte time.Time) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.pulpCreatedLte = &pulpCreatedLte
	return r
}

// Filter results where pulp_created is between two comma separated values
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) PulpCreatedRange(pulpCreatedRange []time.Time) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.pulpCreatedRange = &pulpCreatedRange
	return r
}

// A list of fields to include in the response.
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) Fields(fields []string) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) ExcludeFields(excludeFields []string) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) Execute() (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	return r.ApiService.RepositoriesPythonPythonVersionsListExecute(r)
}

/*
RepositoriesPythonPythonVersionsList List repository versions

PythonRepositoryVersion represents a single Python repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pythonPythonRepositoryHref
 @return RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest
*/
func (a *RepositoriesPythonVersionsApiService) RepositoriesPythonPythonVersionsList(ctx context.Context, pythonPythonRepositoryHref string) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest {
	return RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest{
		ApiService: a,
		ctx: ctx,
		pythonPythonRepositoryHref: pythonPythonRepositoryHref,
	}
}

// Execute executes the request
//  @return PaginatedRepositoryVersionResponseList
func (a *RepositoriesPythonVersionsApiService) RepositoriesPythonPythonVersionsListExecute(r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsListRequest) (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRepositoryVersionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesPythonVersionsApiService.RepositoriesPythonPythonVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{python_python_repository_href}versions/"
	localVarPath = strings.Replace(localVarPath, "{"+"python_python_repository_href"+"}", url.PathEscape(parameterValueToString(r.pythonPythonRepositoryHref, "pythonPythonRepositoryHref")), -1)
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

type RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsReadRequest struct {
	ctx context.Context
	ApiService *RepositoriesPythonVersionsApiService
	pythonPythonRepositoryVersionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsReadRequest) Fields(fields []string) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsReadRequest) ExcludeFields(excludeFields []string) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsReadRequest) Execute() (*RepositoryVersionResponse, *http.Response, error) {
	return r.ApiService.RepositoriesPythonPythonVersionsReadExecute(r)
}

/*
RepositoriesPythonPythonVersionsRead Inspect a repository version

PythonRepositoryVersion represents a single Python repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pythonPythonRepositoryVersionHref
 @return RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsReadRequest
*/
func (a *RepositoriesPythonVersionsApiService) RepositoriesPythonPythonVersionsRead(ctx context.Context, pythonPythonRepositoryVersionHref string) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsReadRequest {
	return RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsReadRequest{
		ApiService: a,
		ctx: ctx,
		pythonPythonRepositoryVersionHref: pythonPythonRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return RepositoryVersionResponse
func (a *RepositoriesPythonVersionsApiService) RepositoriesPythonPythonVersionsReadExecute(r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsReadRequest) (*RepositoryVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepositoryVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesPythonVersionsApiService.RepositoriesPythonPythonVersionsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{python_python_repository_version_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"python_python_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.pythonPythonRepositoryVersionHref, "pythonPythonRepositoryVersionHref")), -1)
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

type RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsRepairRequest struct {
	ctx context.Context
	ApiService *RepositoriesPythonVersionsApiService
	pythonPythonRepositoryVersionHref string
	repair *Repair
}

func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsRepairRequest) Repair(repair Repair) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsRepairRequest {
	r.repair = &repair
	return r
}

func (r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsRepairRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesPythonPythonVersionsRepairExecute(r)
}

/*
RepositoriesPythonPythonVersionsRepair Method for RepositoriesPythonPythonVersionsRepair

Trigger an asynchronous task to repair a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pythonPythonRepositoryVersionHref
 @return RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsRepairRequest
*/
func (a *RepositoriesPythonVersionsApiService) RepositoriesPythonPythonVersionsRepair(ctx context.Context, pythonPythonRepositoryVersionHref string) RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsRepairRequest {
	return RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsRepairRequest{
		ApiService: a,
		ctx: ctx,
		pythonPythonRepositoryVersionHref: pythonPythonRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesPythonVersionsApiService) RepositoriesPythonPythonVersionsRepairExecute(r RepositoriesPythonVersionsApiRepositoriesPythonPythonVersionsRepairRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesPythonVersionsApiService.RepositoriesPythonPythonVersionsRepair")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{python_python_repository_version_href}repair/"
	localVarPath = strings.Replace(localVarPath, "{"+"python_python_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.pythonPythonRepositoryVersionHref, "pythonPythonRepositoryVersionHref")), -1)
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
