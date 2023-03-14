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


// ContentInstallerFileIndicesApiService ContentInstallerFileIndicesApi service
type ContentInstallerFileIndicesApiService service

type ContentInstallerFileIndicesApiContentDebInstallerFileIndicesCreateRequest struct {
	ctx context.Context
	ApiService *ContentInstallerFileIndicesApiService
	debInstallerFileIndex *DebInstallerFileIndex
}

func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesCreateRequest) DebInstallerFileIndex(debInstallerFileIndex DebInstallerFileIndex) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesCreateRequest {
	r.debInstallerFileIndex = &debInstallerFileIndex
	return r
}

func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesCreateRequest) Execute() (*DebInstallerFileIndexResponse, *http.Response, error) {
	return r.ApiService.ContentDebInstallerFileIndicesCreateExecute(r)
}

/*
ContentDebInstallerFileIndicesCreate Create an installer file index

An InstallerFileIndex represents the indices for a set of installer files.

Associated artifacts: Exactly one 'SHA256SUMS' and/or 'MD5SUMS' file.

Each InstallerFileIndes is associated with a single component-architecture combination within
a single Release. Note that installer files are currently used exclusively for verbatim
publications. The APT publisher (both simple and structured mode) does not make use of installer
content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentInstallerFileIndicesApiContentDebInstallerFileIndicesCreateRequest
*/
func (a *ContentInstallerFileIndicesApiService) ContentDebInstallerFileIndicesCreate(ctx context.Context) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesCreateRequest {
	return ContentInstallerFileIndicesApiContentDebInstallerFileIndicesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DebInstallerFileIndexResponse
func (a *ContentInstallerFileIndicesApiService) ContentDebInstallerFileIndicesCreateExecute(r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesCreateRequest) (*DebInstallerFileIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebInstallerFileIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentInstallerFileIndicesApiService.ContentDebInstallerFileIndicesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/installer_file_indices/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.debInstallerFileIndex == nil {
		return localVarReturnValue, nil, reportError("debInstallerFileIndex is required and must be specified")
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
	localVarPostBody = r.debInstallerFileIndex
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

type ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest struct {
	ctx context.Context
	ApiService *ContentInstallerFileIndicesApiService
	architecture *string
	component *string
	limit *int32
	offset *int32
	ordering *[]string
	relativePath *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	sha256 *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where architecture matches value
func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest) Architecture(architecture string) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest {
	r.architecture = &architecture
	return r
}

// Filter results where component matches value
func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest) Component(component string) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest {
	r.component = &component
	return r
}

// Number of results to return per page.
func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest) Limit(limit int32) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest) Offset(offset int32) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest) Ordering(ordering []string) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where relative_path matches value
func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest) RelativePath(relativePath string) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest {
	r.relativePath = &relativePath
	return r
}

// Repository Version referenced by HREF
func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest) RepositoryVersion(repositoryVersion string) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter results where sha256 matches value
func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest) Sha256(sha256 string) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest {
	r.sha256 = &sha256
	return r
}

// A list of fields to include in the response.
func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest) Fields(fields []string) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest) ExcludeFields(excludeFields []string) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest) Execute() (*PaginateddebInstallerFileIndexResponseList, *http.Response, error) {
	return r.ApiService.ContentDebInstallerFileIndicesListExecute(r)
}

/*
ContentDebInstallerFileIndicesList List InstallerFileIndices

An InstallerFileIndex represents the indices for a set of installer files.

Associated artifacts: Exactly one 'SHA256SUMS' and/or 'MD5SUMS' file.

Each InstallerFileIndes is associated with a single component-architecture combination within
a single Release. Note that installer files are currently used exclusively for verbatim
publications. The APT publisher (both simple and structured mode) does not make use of installer
content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest
*/
func (a *ContentInstallerFileIndicesApiService) ContentDebInstallerFileIndicesList(ctx context.Context) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest {
	return ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginateddebInstallerFileIndexResponseList
func (a *ContentInstallerFileIndicesApiService) ContentDebInstallerFileIndicesListExecute(r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesListRequest) (*PaginateddebInstallerFileIndexResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginateddebInstallerFileIndexResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentInstallerFileIndicesApiService.ContentDebInstallerFileIndicesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/installer_file_indices/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.architecture != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "architecture", r.architecture, "")
	}
	if r.component != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "component", r.component, "")
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
	if r.relativePath != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "relative_path", r.relativePath, "")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha256", r.sha256, "")
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

type ContentInstallerFileIndicesApiContentDebInstallerFileIndicesReadRequest struct {
	ctx context.Context
	ApiService *ContentInstallerFileIndicesApiService
	debInstallerFileIndexHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesReadRequest) Fields(fields []string) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesReadRequest) ExcludeFields(excludeFields []string) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesReadRequest) Execute() (*DebInstallerFileIndexResponse, *http.Response, error) {
	return r.ApiService.ContentDebInstallerFileIndicesReadExecute(r)
}

/*
ContentDebInstallerFileIndicesRead Inspect an installer file index

An InstallerFileIndex represents the indices for a set of installer files.

Associated artifacts: Exactly one 'SHA256SUMS' and/or 'MD5SUMS' file.

Each InstallerFileIndes is associated with a single component-architecture combination within
a single Release. Note that installer files are currently used exclusively for verbatim
publications. The APT publisher (both simple and structured mode) does not make use of installer
content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debInstallerFileIndexHref
 @return ContentInstallerFileIndicesApiContentDebInstallerFileIndicesReadRequest
*/
func (a *ContentInstallerFileIndicesApiService) ContentDebInstallerFileIndicesRead(ctx context.Context, debInstallerFileIndexHref string) ContentInstallerFileIndicesApiContentDebInstallerFileIndicesReadRequest {
	return ContentInstallerFileIndicesApiContentDebInstallerFileIndicesReadRequest{
		ApiService: a,
		ctx: ctx,
		debInstallerFileIndexHref: debInstallerFileIndexHref,
	}
}

// Execute executes the request
//  @return DebInstallerFileIndexResponse
func (a *ContentInstallerFileIndicesApiService) ContentDebInstallerFileIndicesReadExecute(r ContentInstallerFileIndicesApiContentDebInstallerFileIndicesReadRequest) (*DebInstallerFileIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebInstallerFileIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentInstallerFileIndicesApiService.ContentDebInstallerFileIndicesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_installer_file_index_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_installer_file_index_href"+"}", url.PathEscape(parameterValueToString(r.debInstallerFileIndexHref, "debInstallerFileIndexHref")), -1)
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