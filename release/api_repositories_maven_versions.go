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


// RepositoriesMavenVersionsApiService RepositoriesMavenVersionsApi service
type RepositoriesMavenVersionsApiService service

type RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsDeleteRequest struct {
	ctx context.Context
	ApiService *RepositoriesMavenVersionsApiService
	mavenMavenRepositoryVersionHref string
}

func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesMavenMavenVersionsDeleteExecute(r)
}

/*
RepositoriesMavenMavenVersionsDelete Delete a repository version

Trigger an asynchronous task to delete a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenRepositoryVersionHref
 @return RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsDeleteRequest
*/
func (a *RepositoriesMavenVersionsApiService) RepositoriesMavenMavenVersionsDelete(ctx context.Context, mavenMavenRepositoryVersionHref string) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsDeleteRequest {
	return RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenRepositoryVersionHref: mavenMavenRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesMavenVersionsApiService) RepositoriesMavenMavenVersionsDeleteExecute(r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesMavenVersionsApiService.RepositoriesMavenMavenVersionsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{maven_maven_repository_version_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.mavenMavenRepositoryVersionHref, "mavenMavenRepositoryVersionHref")), -1)
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

type RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest struct {
	ctx context.Context
	ApiService *RepositoriesMavenVersionsApiService
	mavenMavenRepositoryHref string
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
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) Content(content string) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.content = &content
	return r
}

// Content Unit referenced by HREF
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) ContentIn(contentIn string) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.contentIn = &contentIn
	return r
}

// Number of results to return per page.
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) Limit(limit int32) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.limit = &limit
	return r
}

// Filter results where number matches value
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) Number(number int32) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.number = &number
	return r
}

// Filter results where number is greater than value
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) NumberGt(numberGt int32) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.numberGt = &numberGt
	return r
}

// Filter results where number is greater than or equal to value
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) NumberGte(numberGte int32) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.numberGte = &numberGte
	return r
}

// Filter results where number is less than value
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) NumberLt(numberLt int32) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.numberLt = &numberLt
	return r
}

// Filter results where number is less than or equal to value
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) NumberLte(numberLte int32) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.numberLte = &numberLte
	return r
}

// Filter results where number is between two comma separated values
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) NumberRange(numberRange []int32) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.numberRange = &numberRange
	return r
}

// The initial index from which to return the results.
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) Offset(offset int32) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) Ordering(ordering []string) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where pulp_created matches value
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) PulpCreated(pulpCreated time.Time) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.pulpCreated = &pulpCreated
	return r
}

// Filter results where pulp_created is greater than value
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) PulpCreatedGt(pulpCreatedGt time.Time) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.pulpCreatedGt = &pulpCreatedGt
	return r
}

// Filter results where pulp_created is greater than or equal to value
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) PulpCreatedGte(pulpCreatedGte time.Time) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.pulpCreatedGte = &pulpCreatedGte
	return r
}

// Filter results where pulp_created is less than value
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) PulpCreatedLt(pulpCreatedLt time.Time) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.pulpCreatedLt = &pulpCreatedLt
	return r
}

// Filter results where pulp_created is less than or equal to value
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) PulpCreatedLte(pulpCreatedLte time.Time) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.pulpCreatedLte = &pulpCreatedLte
	return r
}

// Filter results where pulp_created is between two comma separated values
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) PulpCreatedRange(pulpCreatedRange []time.Time) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.pulpCreatedRange = &pulpCreatedRange
	return r
}

// A list of fields to include in the response.
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) Fields(fields []string) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) ExcludeFields(excludeFields []string) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) Execute() (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	return r.ApiService.RepositoriesMavenMavenVersionsListExecute(r)
}

/*
RepositoriesMavenMavenVersionsList List repository versions

MavenRepositoryVersion represents a single Maven repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenRepositoryHref
 @return RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest
*/
func (a *RepositoriesMavenVersionsApiService) RepositoriesMavenMavenVersionsList(ctx context.Context, mavenMavenRepositoryHref string) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest {
	return RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenRepositoryHref: mavenMavenRepositoryHref,
	}
}

// Execute executes the request
//  @return PaginatedRepositoryVersionResponseList
func (a *RepositoriesMavenVersionsApiService) RepositoriesMavenMavenVersionsListExecute(r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsListRequest) (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRepositoryVersionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesMavenVersionsApiService.RepositoriesMavenMavenVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{maven_maven_repository_href}versions/"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_repository_href"+"}", url.PathEscape(parameterValueToString(r.mavenMavenRepositoryHref, "mavenMavenRepositoryHref")), -1)
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

type RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsReadRequest struct {
	ctx context.Context
	ApiService *RepositoriesMavenVersionsApiService
	mavenMavenRepositoryVersionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsReadRequest) Fields(fields []string) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsReadRequest) ExcludeFields(excludeFields []string) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsReadRequest) Execute() (*RepositoryVersionResponse, *http.Response, error) {
	return r.ApiService.RepositoriesMavenMavenVersionsReadExecute(r)
}

/*
RepositoriesMavenMavenVersionsRead Inspect a repository version

MavenRepositoryVersion represents a single Maven repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenRepositoryVersionHref
 @return RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsReadRequest
*/
func (a *RepositoriesMavenVersionsApiService) RepositoriesMavenMavenVersionsRead(ctx context.Context, mavenMavenRepositoryVersionHref string) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsReadRequest {
	return RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsReadRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenRepositoryVersionHref: mavenMavenRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return RepositoryVersionResponse
func (a *RepositoriesMavenVersionsApiService) RepositoriesMavenMavenVersionsReadExecute(r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsReadRequest) (*RepositoryVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepositoryVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesMavenVersionsApiService.RepositoriesMavenMavenVersionsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{maven_maven_repository_version_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.mavenMavenRepositoryVersionHref, "mavenMavenRepositoryVersionHref")), -1)
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

type RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsRepairRequest struct {
	ctx context.Context
	ApiService *RepositoriesMavenVersionsApiService
	mavenMavenRepositoryVersionHref string
	repair *Repair
}

func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsRepairRequest) Repair(repair Repair) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsRepairRequest {
	r.repair = &repair
	return r
}

func (r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsRepairRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesMavenMavenVersionsRepairExecute(r)
}

/*
RepositoriesMavenMavenVersionsRepair Method for RepositoriesMavenMavenVersionsRepair

Trigger an asynchronous task to repair a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenRepositoryVersionHref
 @return RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsRepairRequest
*/
func (a *RepositoriesMavenVersionsApiService) RepositoriesMavenMavenVersionsRepair(ctx context.Context, mavenMavenRepositoryVersionHref string) RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsRepairRequest {
	return RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsRepairRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenRepositoryVersionHref: mavenMavenRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesMavenVersionsApiService) RepositoriesMavenMavenVersionsRepairExecute(r RepositoriesMavenVersionsApiRepositoriesMavenMavenVersionsRepairRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesMavenVersionsApiService.RepositoriesMavenMavenVersionsRepair")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{maven_maven_repository_version_href}repair/"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.mavenMavenRepositoryVersionHref, "mavenMavenRepositoryVersionHref")), -1)
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
