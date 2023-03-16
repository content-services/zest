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
	"time"
	"reflect"
)


// PublicationsVerbatimApiService PublicationsVerbatimApi service
type PublicationsVerbatimApiService service

type PublicationsVerbatimApiPublicationsDebVerbatimCreateRequest struct {
	ctx context.Context
	ApiService *PublicationsVerbatimApiService
	debVerbatimPublication *DebVerbatimPublication
}

func (r PublicationsVerbatimApiPublicationsDebVerbatimCreateRequest) DebVerbatimPublication(debVerbatimPublication DebVerbatimPublication) PublicationsVerbatimApiPublicationsDebVerbatimCreateRequest {
	r.debVerbatimPublication = &debVerbatimPublication
	return r
}

func (r PublicationsVerbatimApiPublicationsDebVerbatimCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PublicationsDebVerbatimCreateExecute(r)
}

/*
PublicationsDebVerbatimCreate Create a verbatim publication

Trigger an asynchronous task to publish content

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PublicationsVerbatimApiPublicationsDebVerbatimCreateRequest
*/
func (a *PublicationsVerbatimApiService) PublicationsDebVerbatimCreate(ctx context.Context) PublicationsVerbatimApiPublicationsDebVerbatimCreateRequest {
	return PublicationsVerbatimApiPublicationsDebVerbatimCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *PublicationsVerbatimApiService) PublicationsDebVerbatimCreateExecute(r PublicationsVerbatimApiPublicationsDebVerbatimCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsVerbatimApiService.PublicationsDebVerbatimCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/publications/deb/verbatim/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.debVerbatimPublication == nil {
		return localVarReturnValue, nil, reportError("debVerbatimPublication is required and must be specified")
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
	localVarPostBody = r.debVerbatimPublication
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

type PublicationsVerbatimApiPublicationsDebVerbatimDeleteRequest struct {
	ctx context.Context
	ApiService *PublicationsVerbatimApiService
	debVerbatimPublicationHref string
}

func (r PublicationsVerbatimApiPublicationsDebVerbatimDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.PublicationsDebVerbatimDeleteExecute(r)
}

/*
PublicationsDebVerbatimDelete Delete a verbatim publication

An VerbatimPublication is the Pulp-internal representation of a "mirrored" AptRepositoryVersion.

In other words, the verbatim publisher will recreate the synced subset of some a APT
repository using the exact same metadata files and signatures as used by the upstream original.
Once a Pulp publication has been created, it can be served by creating a Pulp distribution (in
a near atomic action).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debVerbatimPublicationHref
 @return PublicationsVerbatimApiPublicationsDebVerbatimDeleteRequest
*/
func (a *PublicationsVerbatimApiService) PublicationsDebVerbatimDelete(ctx context.Context, debVerbatimPublicationHref string) PublicationsVerbatimApiPublicationsDebVerbatimDeleteRequest {
	return PublicationsVerbatimApiPublicationsDebVerbatimDeleteRequest{
		ApiService: a,
		ctx: ctx,
		debVerbatimPublicationHref: debVerbatimPublicationHref,
	}
}

// Execute executes the request
func (a *PublicationsVerbatimApiService) PublicationsDebVerbatimDeleteExecute(r PublicationsVerbatimApiPublicationsDebVerbatimDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsVerbatimApiService.PublicationsDebVerbatimDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_verbatim_publication_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_verbatim_publication_href"+"}", url.PathEscape(parameterValueToString(r.debVerbatimPublicationHref, "debVerbatimPublicationHref")), -1)
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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type PublicationsVerbatimApiPublicationsDebVerbatimListRequest struct {
	ctx context.Context
	ApiService *PublicationsVerbatimApiService
	content *string
	contentIn *string
	limit *int32
	offset *int32
	ordering *[]string
	pulpCreated *time.Time
	pulpCreatedGt *time.Time
	pulpCreatedGte *time.Time
	pulpCreatedLt *time.Time
	pulpCreatedLte *time.Time
	pulpCreatedRange *[]time.Time
	repository *string
	repositoryVersion *string
	fields *[]string
	excludeFields *[]string
}

// Content Unit referenced by HREF
func (r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) Content(content string) PublicationsVerbatimApiPublicationsDebVerbatimListRequest {
	r.content = &content
	return r
}

// Content Unit referenced by HREF
func (r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) ContentIn(contentIn string) PublicationsVerbatimApiPublicationsDebVerbatimListRequest {
	r.contentIn = &contentIn
	return r
}

// Number of results to return per page.
func (r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) Limit(limit int32) PublicationsVerbatimApiPublicationsDebVerbatimListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) Offset(offset int32) PublicationsVerbatimApiPublicationsDebVerbatimListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) Ordering(ordering []string) PublicationsVerbatimApiPublicationsDebVerbatimListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where pulp_created matches value
func (r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) PulpCreated(pulpCreated time.Time) PublicationsVerbatimApiPublicationsDebVerbatimListRequest {
	r.pulpCreated = &pulpCreated
	return r
}

// Filter results where pulp_created is greater than value
func (r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) PulpCreatedGt(pulpCreatedGt time.Time) PublicationsVerbatimApiPublicationsDebVerbatimListRequest {
	r.pulpCreatedGt = &pulpCreatedGt
	return r
}

// Filter results where pulp_created is greater than or equal to value
func (r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) PulpCreatedGte(pulpCreatedGte time.Time) PublicationsVerbatimApiPublicationsDebVerbatimListRequest {
	r.pulpCreatedGte = &pulpCreatedGte
	return r
}

// Filter results where pulp_created is less than value
func (r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) PulpCreatedLt(pulpCreatedLt time.Time) PublicationsVerbatimApiPublicationsDebVerbatimListRequest {
	r.pulpCreatedLt = &pulpCreatedLt
	return r
}

// Filter results where pulp_created is less than or equal to value
func (r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) PulpCreatedLte(pulpCreatedLte time.Time) PublicationsVerbatimApiPublicationsDebVerbatimListRequest {
	r.pulpCreatedLte = &pulpCreatedLte
	return r
}

// Filter results where pulp_created is between two comma separated values
func (r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) PulpCreatedRange(pulpCreatedRange []time.Time) PublicationsVerbatimApiPublicationsDebVerbatimListRequest {
	r.pulpCreatedRange = &pulpCreatedRange
	return r
}

// Repository referenced by HREF
func (r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) Repository(repository string) PublicationsVerbatimApiPublicationsDebVerbatimListRequest {
	r.repository = &repository
	return r
}

// Repository Version referenced by HREF
func (r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) RepositoryVersion(repositoryVersion string) PublicationsVerbatimApiPublicationsDebVerbatimListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// A list of fields to include in the response.
func (r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) Fields(fields []string) PublicationsVerbatimApiPublicationsDebVerbatimListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) ExcludeFields(excludeFields []string) PublicationsVerbatimApiPublicationsDebVerbatimListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) Execute() (*PaginateddebVerbatimPublicationResponseList, *http.Response, error) {
	return r.ApiService.PublicationsDebVerbatimListExecute(r)
}

/*
PublicationsDebVerbatimList List verbatim publications

An VerbatimPublication is the Pulp-internal representation of a "mirrored" AptRepositoryVersion.

In other words, the verbatim publisher will recreate the synced subset of some a APT
repository using the exact same metadata files and signatures as used by the upstream original.
Once a Pulp publication has been created, it can be served by creating a Pulp distribution (in
a near atomic action).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PublicationsVerbatimApiPublicationsDebVerbatimListRequest
*/
func (a *PublicationsVerbatimApiService) PublicationsDebVerbatimList(ctx context.Context) PublicationsVerbatimApiPublicationsDebVerbatimListRequest {
	return PublicationsVerbatimApiPublicationsDebVerbatimListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginateddebVerbatimPublicationResponseList
func (a *PublicationsVerbatimApiService) PublicationsDebVerbatimListExecute(r PublicationsVerbatimApiPublicationsDebVerbatimListRequest) (*PaginateddebVerbatimPublicationResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginateddebVerbatimPublicationResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsVerbatimApiService.PublicationsDebVerbatimList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/publications/deb/verbatim/"
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
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository", r.repository, "")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
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

type PublicationsVerbatimApiPublicationsDebVerbatimReadRequest struct {
	ctx context.Context
	ApiService *PublicationsVerbatimApiService
	debVerbatimPublicationHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PublicationsVerbatimApiPublicationsDebVerbatimReadRequest) Fields(fields []string) PublicationsVerbatimApiPublicationsDebVerbatimReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PublicationsVerbatimApiPublicationsDebVerbatimReadRequest) ExcludeFields(excludeFields []string) PublicationsVerbatimApiPublicationsDebVerbatimReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PublicationsVerbatimApiPublicationsDebVerbatimReadRequest) Execute() (*DebVerbatimPublicationResponse, *http.Response, error) {
	return r.ApiService.PublicationsDebVerbatimReadExecute(r)
}

/*
PublicationsDebVerbatimRead Inspect a verbatim publication

An VerbatimPublication is the Pulp-internal representation of a "mirrored" AptRepositoryVersion.

In other words, the verbatim publisher will recreate the synced subset of some a APT
repository using the exact same metadata files and signatures as used by the upstream original.
Once a Pulp publication has been created, it can be served by creating a Pulp distribution (in
a near atomic action).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debVerbatimPublicationHref
 @return PublicationsVerbatimApiPublicationsDebVerbatimReadRequest
*/
func (a *PublicationsVerbatimApiService) PublicationsDebVerbatimRead(ctx context.Context, debVerbatimPublicationHref string) PublicationsVerbatimApiPublicationsDebVerbatimReadRequest {
	return PublicationsVerbatimApiPublicationsDebVerbatimReadRequest{
		ApiService: a,
		ctx: ctx,
		debVerbatimPublicationHref: debVerbatimPublicationHref,
	}
}

// Execute executes the request
//  @return DebVerbatimPublicationResponse
func (a *PublicationsVerbatimApiService) PublicationsDebVerbatimReadExecute(r PublicationsVerbatimApiPublicationsDebVerbatimReadRequest) (*DebVerbatimPublicationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebVerbatimPublicationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsVerbatimApiService.PublicationsDebVerbatimRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_verbatim_publication_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_verbatim_publication_href"+"}", url.PathEscape(parameterValueToString(r.debVerbatimPublicationHref, "debVerbatimPublicationHref")), -1)
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
