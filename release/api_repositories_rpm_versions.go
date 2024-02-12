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


// RepositoriesRpmVersionsAPIService RepositoriesRpmVersionsAPI service
type RepositoriesRpmVersionsAPIService service

type RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsDeleteRequest struct {
	ctx context.Context
	ApiService *RepositoriesRpmVersionsAPIService
	rpmRpmRepositoryVersionHref string
}

func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesRpmRpmVersionsDeleteExecute(r)
}

/*
RepositoriesRpmRpmVersionsDelete Delete a repository version

Trigger an asynchronous task to delete a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmRpmRepositoryVersionHref
 @return RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsDeleteRequest
*/
func (a *RepositoriesRpmVersionsAPIService) RepositoriesRpmRpmVersionsDelete(ctx context.Context, rpmRpmRepositoryVersionHref string) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsDeleteRequest {
	return RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		rpmRpmRepositoryVersionHref: rpmRpmRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesRpmVersionsAPIService) RepositoriesRpmRpmVersionsDeleteExecute(r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesRpmVersionsAPIService.RepositoriesRpmRpmVersionsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{rpm_rpm_repository_version_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_rpm_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.rpmRpmRepositoryVersionHref, "rpmRpmRepositoryVersionHref")), -1)
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

type RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest struct {
	ctx context.Context
	ApiService *RepositoriesRpmVersionsAPIService
	rpmRpmRepositoryHref string
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
	pulpHrefIn *[]string
	q *string
	fields *[]string
	excludeFields *[]string
}

// Content Unit referenced by HREF
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) Content(content string) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.content = &content
	return r
}

// Content Unit referenced by HREF
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) ContentIn(contentIn string) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.contentIn = &contentIn
	return r
}

// Number of results to return per page.
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) Limit(limit int32) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.limit = &limit
	return r
}

// Filter results where number matches value
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) Number(number int32) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.number = &number
	return r
}

// Filter results where number is greater than value
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) NumberGt(numberGt int32) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.numberGt = &numberGt
	return r
}

// Filter results where number is greater than or equal to value
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) NumberGte(numberGte int32) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.numberGte = &numberGte
	return r
}

// Filter results where number is less than value
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) NumberLt(numberLt int32) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.numberLt = &numberLt
	return r
}

// Filter results where number is less than or equal to value
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) NumberLte(numberLte int32) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.numberLte = &numberLte
	return r
}

// Filter results where number is between two comma separated values
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) NumberRange(numberRange []int32) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.numberRange = &numberRange
	return r
}

// The initial index from which to return the results.
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) Offset(offset int32) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;number&#x60; - Number* &#x60;-number&#x60; - Number (descending)* &#x60;complete&#x60; - Complete* &#x60;-complete&#x60; - Complete (descending)* &#x60;info&#x60; - Info* &#x60;-info&#x60; - Info (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) Ordering(ordering []string) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where pulp_created matches value
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) PulpCreated(pulpCreated time.Time) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.pulpCreated = &pulpCreated
	return r
}

// Filter results where pulp_created is greater than value
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) PulpCreatedGt(pulpCreatedGt time.Time) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.pulpCreatedGt = &pulpCreatedGt
	return r
}

// Filter results where pulp_created is greater than or equal to value
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) PulpCreatedGte(pulpCreatedGte time.Time) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.pulpCreatedGte = &pulpCreatedGte
	return r
}

// Filter results where pulp_created is less than value
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) PulpCreatedLt(pulpCreatedLt time.Time) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.pulpCreatedLt = &pulpCreatedLt
	return r
}

// Filter results where pulp_created is less than or equal to value
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) PulpCreatedLte(pulpCreatedLte time.Time) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.pulpCreatedLte = &pulpCreatedLte
	return r
}

// Filter results where pulp_created is between two comma separated values
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) PulpCreatedRange(pulpCreatedRange []time.Time) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.pulpCreatedRange = &pulpCreatedRange
	return r
}

// Multiple values may be separated by commas.
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) PulpHrefIn(pulpHrefIn []string) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) Q(q string) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.q = &q
	return r
}

// A list of fields to include in the response.
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) Fields(fields []string) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) ExcludeFields(excludeFields []string) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) Execute() (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	return r.ApiService.RepositoriesRpmRpmVersionsListExecute(r)
}

/*
RepositoriesRpmRpmVersionsList List repository versions

RpmRepositoryVersion represents a single rpm repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmRpmRepositoryHref
 @return RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest
*/
func (a *RepositoriesRpmVersionsAPIService) RepositoriesRpmRpmVersionsList(ctx context.Context, rpmRpmRepositoryHref string) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest {
	return RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest{
		ApiService: a,
		ctx: ctx,
		rpmRpmRepositoryHref: rpmRpmRepositoryHref,
	}
}

// Execute executes the request
//  @return PaginatedRepositoryVersionResponseList
func (a *RepositoriesRpmVersionsAPIService) RepositoriesRpmRpmVersionsListExecute(r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsListRequest) (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRepositoryVersionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesRpmVersionsAPIService.RepositoriesRpmRpmVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{rpm_rpm_repository_href}versions/"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_rpm_repository_href"+"}", url.PathEscape(parameterValueToString(r.rpmRpmRepositoryHref, "rpmRpmRepositoryHref")), -1)
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
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
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

type RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsReadRequest struct {
	ctx context.Context
	ApiService *RepositoriesRpmVersionsAPIService
	rpmRpmRepositoryVersionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsReadRequest) Fields(fields []string) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsReadRequest) ExcludeFields(excludeFields []string) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsReadRequest) Execute() (*RepositoryVersionResponse, *http.Response, error) {
	return r.ApiService.RepositoriesRpmRpmVersionsReadExecute(r)
}

/*
RepositoriesRpmRpmVersionsRead Inspect a repository version

RpmRepositoryVersion represents a single rpm repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmRpmRepositoryVersionHref
 @return RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsReadRequest
*/
func (a *RepositoriesRpmVersionsAPIService) RepositoriesRpmRpmVersionsRead(ctx context.Context, rpmRpmRepositoryVersionHref string) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsReadRequest {
	return RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsReadRequest{
		ApiService: a,
		ctx: ctx,
		rpmRpmRepositoryVersionHref: rpmRpmRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return RepositoryVersionResponse
func (a *RepositoriesRpmVersionsAPIService) RepositoriesRpmRpmVersionsReadExecute(r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsReadRequest) (*RepositoryVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepositoryVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesRpmVersionsAPIService.RepositoriesRpmRpmVersionsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{rpm_rpm_repository_version_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_rpm_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.rpmRpmRepositoryVersionHref, "rpmRpmRepositoryVersionHref")), -1)
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

type RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsRepairRequest struct {
	ctx context.Context
	ApiService *RepositoriesRpmVersionsAPIService
	rpmRpmRepositoryVersionHref string
	repair *Repair
}

func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsRepairRequest) Repair(repair Repair) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsRepairRequest {
	r.repair = &repair
	return r
}

func (r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsRepairRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesRpmRpmVersionsRepairExecute(r)
}

/*
RepositoriesRpmRpmVersionsRepair Method for RepositoriesRpmRpmVersionsRepair

Trigger an asynchronous task to repair a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmRpmRepositoryVersionHref
 @return RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsRepairRequest
*/
func (a *RepositoriesRpmVersionsAPIService) RepositoriesRpmRpmVersionsRepair(ctx context.Context, rpmRpmRepositoryVersionHref string) RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsRepairRequest {
	return RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsRepairRequest{
		ApiService: a,
		ctx: ctx,
		rpmRpmRepositoryVersionHref: rpmRpmRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesRpmVersionsAPIService) RepositoriesRpmRpmVersionsRepairExecute(r RepositoriesRpmVersionsAPIRepositoriesRpmRpmVersionsRepairRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesRpmVersionsAPIService.RepositoriesRpmRpmVersionsRepair")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{rpm_rpm_repository_version_href}repair/"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_rpm_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.rpmRpmRepositoryVersionHref, "rpmRpmRepositoryVersionHref")), -1)
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
