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
	"reflect"
)


// ContentPackageIndicesApiService ContentPackageIndicesApi service
type ContentPackageIndicesApiService service

type ContentPackageIndicesApiContentDebPackageIndicesCreateRequest struct {
	ctx context.Context
	ApiService *ContentPackageIndicesApiService
	debPackageIndex *DebPackageIndex
}

func (r ContentPackageIndicesApiContentDebPackageIndicesCreateRequest) DebPackageIndex(debPackageIndex DebPackageIndex) ContentPackageIndicesApiContentDebPackageIndicesCreateRequest {
	r.debPackageIndex = &debPackageIndex
	return r
}

func (r ContentPackageIndicesApiContentDebPackageIndicesCreateRequest) Execute() (*DebPackageIndexResponse, *http.Response, error) {
	return r.ApiService.ContentDebPackageIndicesCreateExecute(r)
}

/*
ContentDebPackageIndicesCreate Create a package index

A PackageIndex represents the package indices of a single component-architecture combination.

Associated artifacts: Exactly one 'Packages' file. May optionally include one or more of
'Packages.gz', 'Packages.xz', 'Release'. If included, the 'Release' file is a legacy
per-component-and-architecture Release file.

Note: The verbatim publisher will republish all associated artifacts, while the APT publisher
(both simple and structured mode) will generate any 'Packages' files it needs when creating the
publication. It does not make use of PackageIndex content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentPackageIndicesApiContentDebPackageIndicesCreateRequest
*/
func (a *ContentPackageIndicesApiService) ContentDebPackageIndicesCreate(ctx context.Context) ContentPackageIndicesApiContentDebPackageIndicesCreateRequest {
	return ContentPackageIndicesApiContentDebPackageIndicesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DebPackageIndexResponse
func (a *ContentPackageIndicesApiService) ContentDebPackageIndicesCreateExecute(r ContentPackageIndicesApiContentDebPackageIndicesCreateRequest) (*DebPackageIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebPackageIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackageIndicesApiService.ContentDebPackageIndicesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/package_indices/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.debPackageIndex == nil {
		return localVarReturnValue, nil, reportError("debPackageIndex is required and must be specified")
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
	localVarPostBody = r.debPackageIndex
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

type ContentPackageIndicesApiContentDebPackageIndicesListRequest struct {
	ctx context.Context
	ApiService *ContentPackageIndicesApiService
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
func (r ContentPackageIndicesApiContentDebPackageIndicesListRequest) Architecture(architecture string) ContentPackageIndicesApiContentDebPackageIndicesListRequest {
	r.architecture = &architecture
	return r
}

// Filter results where component matches value
func (r ContentPackageIndicesApiContentDebPackageIndicesListRequest) Component(component string) ContentPackageIndicesApiContentDebPackageIndicesListRequest {
	r.component = &component
	return r
}

// Number of results to return per page.
func (r ContentPackageIndicesApiContentDebPackageIndicesListRequest) Limit(limit int32) ContentPackageIndicesApiContentDebPackageIndicesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentPackageIndicesApiContentDebPackageIndicesListRequest) Offset(offset int32) ContentPackageIndicesApiContentDebPackageIndicesListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ContentPackageIndicesApiContentDebPackageIndicesListRequest) Ordering(ordering []string) ContentPackageIndicesApiContentDebPackageIndicesListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where relative_path matches value
func (r ContentPackageIndicesApiContentDebPackageIndicesListRequest) RelativePath(relativePath string) ContentPackageIndicesApiContentDebPackageIndicesListRequest {
	r.relativePath = &relativePath
	return r
}

// Repository Version referenced by HREF
func (r ContentPackageIndicesApiContentDebPackageIndicesListRequest) RepositoryVersion(repositoryVersion string) ContentPackageIndicesApiContentDebPackageIndicesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentPackageIndicesApiContentDebPackageIndicesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentPackageIndicesApiContentDebPackageIndicesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentPackageIndicesApiContentDebPackageIndicesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentPackageIndicesApiContentDebPackageIndicesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter results where sha256 matches value
func (r ContentPackageIndicesApiContentDebPackageIndicesListRequest) Sha256(sha256 string) ContentPackageIndicesApiContentDebPackageIndicesListRequest {
	r.sha256 = &sha256
	return r
}

// A list of fields to include in the response.
func (r ContentPackageIndicesApiContentDebPackageIndicesListRequest) Fields(fields []string) ContentPackageIndicesApiContentDebPackageIndicesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackageIndicesApiContentDebPackageIndicesListRequest) ExcludeFields(excludeFields []string) ContentPackageIndicesApiContentDebPackageIndicesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackageIndicesApiContentDebPackageIndicesListRequest) Execute() (*PaginateddebPackageIndexResponseList, *http.Response, error) {
	return r.ApiService.ContentDebPackageIndicesListExecute(r)
}

/*
ContentDebPackageIndicesList List PackageIndices

A PackageIndex represents the package indices of a single component-architecture combination.

Associated artifacts: Exactly one 'Packages' file. May optionally include one or more of
'Packages.gz', 'Packages.xz', 'Release'. If included, the 'Release' file is a legacy
per-component-and-architecture Release file.

Note: The verbatim publisher will republish all associated artifacts, while the APT publisher
(both simple and structured mode) will generate any 'Packages' files it needs when creating the
publication. It does not make use of PackageIndex content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentPackageIndicesApiContentDebPackageIndicesListRequest
*/
func (a *ContentPackageIndicesApiService) ContentDebPackageIndicesList(ctx context.Context) ContentPackageIndicesApiContentDebPackageIndicesListRequest {
	return ContentPackageIndicesApiContentDebPackageIndicesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginateddebPackageIndexResponseList
func (a *ContentPackageIndicesApiService) ContentDebPackageIndicesListExecute(r ContentPackageIndicesApiContentDebPackageIndicesListRequest) (*PaginateddebPackageIndexResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginateddebPackageIndexResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackageIndicesApiService.ContentDebPackageIndicesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/package_indices/"
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

type ContentPackageIndicesApiContentDebPackageIndicesReadRequest struct {
	ctx context.Context
	ApiService *ContentPackageIndicesApiService
	debPackageIndexHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentPackageIndicesApiContentDebPackageIndicesReadRequest) Fields(fields []string) ContentPackageIndicesApiContentDebPackageIndicesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackageIndicesApiContentDebPackageIndicesReadRequest) ExcludeFields(excludeFields []string) ContentPackageIndicesApiContentDebPackageIndicesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackageIndicesApiContentDebPackageIndicesReadRequest) Execute() (*DebPackageIndexResponse, *http.Response, error) {
	return r.ApiService.ContentDebPackageIndicesReadExecute(r)
}

/*
ContentDebPackageIndicesRead Inspect a package index

A PackageIndex represents the package indices of a single component-architecture combination.

Associated artifacts: Exactly one 'Packages' file. May optionally include one or more of
'Packages.gz', 'Packages.xz', 'Release'. If included, the 'Release' file is a legacy
per-component-and-architecture Release file.

Note: The verbatim publisher will republish all associated artifacts, while the APT publisher
(both simple and structured mode) will generate any 'Packages' files it needs when creating the
publication. It does not make use of PackageIndex content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debPackageIndexHref
 @return ContentPackageIndicesApiContentDebPackageIndicesReadRequest
*/
func (a *ContentPackageIndicesApiService) ContentDebPackageIndicesRead(ctx context.Context, debPackageIndexHref string) ContentPackageIndicesApiContentDebPackageIndicesReadRequest {
	return ContentPackageIndicesApiContentDebPackageIndicesReadRequest{
		ApiService: a,
		ctx: ctx,
		debPackageIndexHref: debPackageIndexHref,
	}
}

// Execute executes the request
//  @return DebPackageIndexResponse
func (a *ContentPackageIndicesApiService) ContentDebPackageIndicesReadExecute(r ContentPackageIndicesApiContentDebPackageIndicesReadRequest) (*DebPackageIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebPackageIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackageIndicesApiService.ContentDebPackageIndicesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_package_index_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_package_index_href"+"}", url.PathEscape(parameterValueToString(r.debPackageIndexHref, "debPackageIndexHref")), -1)
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
