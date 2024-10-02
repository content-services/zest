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


// RepositoriesGemVersionsAPIService RepositoriesGemVersionsAPI service
type RepositoriesGemVersionsAPIService service

type RepositoriesGemVersionsAPIRepositoriesGemGemVersionsDeleteRequest struct {
	ctx context.Context
	ApiService *RepositoriesGemVersionsAPIService
	gemGemRepositoryVersionHref string
}

func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesGemGemVersionsDeleteExecute(r)
}

/*
RepositoriesGemGemVersionsDelete Delete a repository version

Trigger an asynchronous task to delete a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gemGemRepositoryVersionHref
 @return RepositoriesGemVersionsAPIRepositoriesGemGemVersionsDeleteRequest
*/
func (a *RepositoriesGemVersionsAPIService) RepositoriesGemGemVersionsDelete(ctx context.Context, gemGemRepositoryVersionHref string) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsDeleteRequest {
	return RepositoriesGemVersionsAPIRepositoriesGemGemVersionsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		gemGemRepositoryVersionHref: gemGemRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesGemVersionsAPIService) RepositoriesGemGemVersionsDeleteExecute(r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesGemVersionsAPIService.RepositoriesGemGemVersionsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{gem_gem_repository_version_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"gem_gem_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.gemGemRepositoryVersionHref, "gemGemRepositoryVersionHref")), -1)
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

type RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest struct {
	ctx context.Context
	ApiService *RepositoriesGemVersionsAPIService
	gemGemRepositoryHref string
	content *string
	contentIn *[]string
	limit *int32
	number *int32
	numberGt *int32
	numberGte *int32
	numberLt *int32
	numberLte *int32
	numberRange *[]int32
	offset *int32
	ordering *[]string
	prnIn *[]string
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

// Content Unit referenced by HREF/PRN
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) Content(content string) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.content = &content
	return r
}

// Multiple values may be separated by commas.
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) ContentIn(contentIn []string) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.contentIn = &contentIn
	return r
}

// Number of results to return per page.
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) Limit(limit int32) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.limit = &limit
	return r
}

// Filter results where number matches value
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) Number(number int32) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.number = &number
	return r
}

// Filter results where number is greater than value
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) NumberGt(numberGt int32) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.numberGt = &numberGt
	return r
}

// Filter results where number is greater than or equal to value
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) NumberGte(numberGte int32) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.numberGte = &numberGte
	return r
}

// Filter results where number is less than value
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) NumberLt(numberLt int32) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.numberLt = &numberLt
	return r
}

// Filter results where number is less than or equal to value
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) NumberLte(numberLte int32) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.numberLte = &numberLte
	return r
}

// Filter results where number is between two comma separated values
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) NumberRange(numberRange []int32) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.numberRange = &numberRange
	return r
}

// The initial index from which to return the results.
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) Offset(offset int32) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;number&#x60; - Number* &#x60;-number&#x60; - Number (descending)* &#x60;complete&#x60; - Complete* &#x60;-complete&#x60; - Complete (descending)* &#x60;info&#x60; - Info* &#x60;-info&#x60; - Info (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) Ordering(ordering []string) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) PrnIn(prnIn []string) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.prnIn = &prnIn
	return r
}

// Filter results where pulp_created matches value
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) PulpCreated(pulpCreated time.Time) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.pulpCreated = &pulpCreated
	return r
}

// Filter results where pulp_created is greater than value
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) PulpCreatedGt(pulpCreatedGt time.Time) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.pulpCreatedGt = &pulpCreatedGt
	return r
}

// Filter results where pulp_created is greater than or equal to value
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) PulpCreatedGte(pulpCreatedGte time.Time) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.pulpCreatedGte = &pulpCreatedGte
	return r
}

// Filter results where pulp_created is less than value
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) PulpCreatedLt(pulpCreatedLt time.Time) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.pulpCreatedLt = &pulpCreatedLt
	return r
}

// Filter results where pulp_created is less than or equal to value
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) PulpCreatedLte(pulpCreatedLte time.Time) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.pulpCreatedLte = &pulpCreatedLte
	return r
}

// Filter results where pulp_created is between two comma separated values
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) PulpCreatedRange(pulpCreatedRange []time.Time) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.pulpCreatedRange = &pulpCreatedRange
	return r
}

// Multiple values may be separated by commas.
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) PulpHrefIn(pulpHrefIn []string) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Filter results by using NOT, AND and OR operations on other filters
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) Q(q string) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.q = &q
	return r
}

// A list of fields to include in the response.
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) Fields(fields []string) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) ExcludeFields(excludeFields []string) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) Execute() (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	return r.ApiService.RepositoriesGemGemVersionsListExecute(r)
}

/*
RepositoriesGemGemVersionsList List repository versions

A ViewSet for a GemRepositoryVersion represents a single Gem repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gemGemRepositoryHref
 @return RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest
*/
func (a *RepositoriesGemVersionsAPIService) RepositoriesGemGemVersionsList(ctx context.Context, gemGemRepositoryHref string) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest {
	return RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest{
		ApiService: a,
		ctx: ctx,
		gemGemRepositoryHref: gemGemRepositoryHref,
	}
}

// Execute executes the request
//  @return PaginatedRepositoryVersionResponseList
func (a *RepositoriesGemVersionsAPIService) RepositoriesGemGemVersionsListExecute(r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsListRequest) (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRepositoryVersionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesGemVersionsAPIService.RepositoriesGemGemVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{gem_gem_repository_href}versions/"
	localVarPath = strings.Replace(localVarPath, "{"+"gem_gem_repository_href"+"}", url.PathEscape(parameterValueToString(r.gemGemRepositoryHref, "gemGemRepositoryHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.content != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content", r.content, "form", "")
	}
	if r.contentIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content__in", r.contentIn, "form", "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	}
	if r.numberGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__gt", r.numberGt, "form", "")
	}
	if r.numberGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__gte", r.numberGte, "form", "")
	}
	if r.numberLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__lt", r.numberLt, "form", "")
	}
	if r.numberLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__lte", r.numberLte, "form", "")
	}
	if r.numberRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__range", r.numberRange, "form", "csv")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "form", "csv")
	}
	if r.prnIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prn__in", r.prnIn, "form", "csv")
	}
	if r.pulpCreated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created", r.pulpCreated, "form", "")
	}
	if r.pulpCreatedGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__gt", r.pulpCreatedGt, "form", "")
	}
	if r.pulpCreatedGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__gte", r.pulpCreatedGte, "form", "")
	}
	if r.pulpCreatedLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__lt", r.pulpCreatedLt, "form", "")
	}
	if r.pulpCreatedLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__lte", r.pulpCreatedLte, "form", "")
	}
	if r.pulpCreatedRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__range", r.pulpCreatedRange, "form", "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "form", "csv")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
                               parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "form", "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
                               parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "form", "multi")
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

type RepositoriesGemVersionsAPIRepositoriesGemGemVersionsReadRequest struct {
	ctx context.Context
	ApiService *RepositoriesGemVersionsAPIService
	gemGemRepositoryVersionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsReadRequest) Fields(fields []string) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsReadRequest) ExcludeFields(excludeFields []string) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsReadRequest) Execute() (*RepositoryVersionResponse, *http.Response, error) {
	return r.ApiService.RepositoriesGemGemVersionsReadExecute(r)
}

/*
RepositoriesGemGemVersionsRead Inspect a repository version

A ViewSet for a GemRepositoryVersion represents a single Gem repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gemGemRepositoryVersionHref
 @return RepositoriesGemVersionsAPIRepositoriesGemGemVersionsReadRequest
*/
func (a *RepositoriesGemVersionsAPIService) RepositoriesGemGemVersionsRead(ctx context.Context, gemGemRepositoryVersionHref string) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsReadRequest {
	return RepositoriesGemVersionsAPIRepositoriesGemGemVersionsReadRequest{
		ApiService: a,
		ctx: ctx,
		gemGemRepositoryVersionHref: gemGemRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return RepositoryVersionResponse
func (a *RepositoriesGemVersionsAPIService) RepositoriesGemGemVersionsReadExecute(r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsReadRequest) (*RepositoryVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepositoryVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesGemVersionsAPIService.RepositoriesGemGemVersionsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{gem_gem_repository_version_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"gem_gem_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.gemGemRepositoryVersionHref, "gemGemRepositoryVersionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
                               parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "form", "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
                               parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "form", "multi")
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

type RepositoriesGemVersionsAPIRepositoriesGemGemVersionsRepairRequest struct {
	ctx context.Context
	ApiService *RepositoriesGemVersionsAPIService
	gemGemRepositoryVersionHref string
	repair *Repair
}

func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsRepairRequest) Repair(repair Repair) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsRepairRequest {
	r.repair = &repair
	return r
}

func (r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsRepairRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesGemGemVersionsRepairExecute(r)
}

/*
RepositoriesGemGemVersionsRepair Method for RepositoriesGemGemVersionsRepair

Trigger an asynchronous task to repair a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gemGemRepositoryVersionHref
 @return RepositoriesGemVersionsAPIRepositoriesGemGemVersionsRepairRequest
*/
func (a *RepositoriesGemVersionsAPIService) RepositoriesGemGemVersionsRepair(ctx context.Context, gemGemRepositoryVersionHref string) RepositoriesGemVersionsAPIRepositoriesGemGemVersionsRepairRequest {
	return RepositoriesGemVersionsAPIRepositoriesGemGemVersionsRepairRequest{
		ApiService: a,
		ctx: ctx,
		gemGemRepositoryVersionHref: gemGemRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesGemVersionsAPIService) RepositoriesGemGemVersionsRepairExecute(r RepositoriesGemVersionsAPIRepositoriesGemGemVersionsRepairRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesGemVersionsAPIService.RepositoriesGemGemVersionsRepair")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{gem_gem_repository_version_href}repair/"
	localVarPath = strings.Replace(localVarPath, "{"+"gem_gem_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.gemGemRepositoryVersionHref, "gemGemRepositoryVersionHref")), -1)
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
