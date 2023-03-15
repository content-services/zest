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
	"os"
	"reflect"
)


// ContentGenericContentsApiService ContentGenericContentsApi service
type ContentGenericContentsApiService service

type ContentGenericContentsApiContentDebGenericContentsCreateRequest struct {
	ctx context.Context
	ApiService *ContentGenericContentsApiService
	relativePath *string
	artifact *string
	file *os.File
	repository *string
	upload *string
}

// Path where the artifact is located relative to distributions base_path
func (r ContentGenericContentsApiContentDebGenericContentsCreateRequest) RelativePath(relativePath string) ContentGenericContentsApiContentDebGenericContentsCreateRequest {
	r.relativePath = &relativePath
	return r
}

// Artifact file representing the physical content
func (r ContentGenericContentsApiContentDebGenericContentsCreateRequest) Artifact(artifact string) ContentGenericContentsApiContentDebGenericContentsCreateRequest {
	r.artifact = &artifact
	return r
}

// An uploaded file that may be turned into the artifact of the content unit.
func (r ContentGenericContentsApiContentDebGenericContentsCreateRequest) File(file *os.File) ContentGenericContentsApiContentDebGenericContentsCreateRequest {
	r.file = file
	return r
}

// A URI of a repository the new content unit should be associated with.
func (r ContentGenericContentsApiContentDebGenericContentsCreateRequest) Repository(repository string) ContentGenericContentsApiContentDebGenericContentsCreateRequest {
	r.repository = &repository
	return r
}

// An uncommitted upload that may be turned into the artifact of the content unit.
func (r ContentGenericContentsApiContentDebGenericContentsCreateRequest) Upload(upload string) ContentGenericContentsApiContentDebGenericContentsCreateRequest {
	r.upload = &upload
	return r
}

func (r ContentGenericContentsApiContentDebGenericContentsCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentDebGenericContentsCreateExecute(r)
}

/*
ContentDebGenericContentsCreate Create a generic content

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentGenericContentsApiContentDebGenericContentsCreateRequest
*/
func (a *ContentGenericContentsApiService) ContentDebGenericContentsCreate(ctx context.Context) ContentGenericContentsApiContentDebGenericContentsCreateRequest {
	return ContentGenericContentsApiContentDebGenericContentsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentGenericContentsApiService) ContentDebGenericContentsCreateExecute(r ContentGenericContentsApiContentDebGenericContentsCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentGenericContentsApiService.ContentDebGenericContentsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/generic_contents/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.relativePath == nil {
		return localVarReturnValue, nil, reportError("relativePath is required and must be specified")
	}
	if strlen(*r.relativePath) < 1 {
		return localVarReturnValue, nil, reportError("relativePath must have at least 1 elements")
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
	if r.artifact != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "artifact", r.artifact, "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "relative_path", r.relativePath, "")
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "repository", r.repository, "")
	}
	if r.upload != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "upload", r.upload, "")
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

type ContentGenericContentsApiContentDebGenericContentsListRequest struct {
	ctx context.Context
	ApiService *ContentGenericContentsApiService
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

// Number of results to return per page.
func (r ContentGenericContentsApiContentDebGenericContentsListRequest) Limit(limit int32) ContentGenericContentsApiContentDebGenericContentsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentGenericContentsApiContentDebGenericContentsListRequest) Offset(offset int32) ContentGenericContentsApiContentDebGenericContentsListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ContentGenericContentsApiContentDebGenericContentsListRequest) Ordering(ordering []string) ContentGenericContentsApiContentDebGenericContentsListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where relative_path matches value
func (r ContentGenericContentsApiContentDebGenericContentsListRequest) RelativePath(relativePath string) ContentGenericContentsApiContentDebGenericContentsListRequest {
	r.relativePath = &relativePath
	return r
}

// Repository Version referenced by HREF
func (r ContentGenericContentsApiContentDebGenericContentsListRequest) RepositoryVersion(repositoryVersion string) ContentGenericContentsApiContentDebGenericContentsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentGenericContentsApiContentDebGenericContentsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentGenericContentsApiContentDebGenericContentsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentGenericContentsApiContentDebGenericContentsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentGenericContentsApiContentDebGenericContentsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter results where sha256 matches value
func (r ContentGenericContentsApiContentDebGenericContentsListRequest) Sha256(sha256 string) ContentGenericContentsApiContentDebGenericContentsListRequest {
	r.sha256 = &sha256
	return r
}

// A list of fields to include in the response.
func (r ContentGenericContentsApiContentDebGenericContentsListRequest) Fields(fields []string) ContentGenericContentsApiContentDebGenericContentsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentGenericContentsApiContentDebGenericContentsListRequest) ExcludeFields(excludeFields []string) ContentGenericContentsApiContentDebGenericContentsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentGenericContentsApiContentDebGenericContentsListRequest) Execute() (*PaginateddebGenericContentResponseList, *http.Response, error) {
	return r.ApiService.ContentDebGenericContentsListExecute(r)
}

/*
ContentDebGenericContentsList List generic contents

GenericContent is a catch all category for storing files not covered by any other type.

Associated artifacts: Exactly one arbitrary file that does not match any other type.

This is needed to store arbitrary files for use with the verbatim publisher. If you are not
using the verbatim publisher, you may ignore this type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentGenericContentsApiContentDebGenericContentsListRequest
*/
func (a *ContentGenericContentsApiService) ContentDebGenericContentsList(ctx context.Context) ContentGenericContentsApiContentDebGenericContentsListRequest {
	return ContentGenericContentsApiContentDebGenericContentsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginateddebGenericContentResponseList
func (a *ContentGenericContentsApiService) ContentDebGenericContentsListExecute(r ContentGenericContentsApiContentDebGenericContentsListRequest) (*PaginateddebGenericContentResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginateddebGenericContentResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentGenericContentsApiService.ContentDebGenericContentsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/generic_contents/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ContentGenericContentsApiContentDebGenericContentsReadRequest struct {
	ctx context.Context
	ApiService *ContentGenericContentsApiService
	debGenericContentHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentGenericContentsApiContentDebGenericContentsReadRequest) Fields(fields []string) ContentGenericContentsApiContentDebGenericContentsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentGenericContentsApiContentDebGenericContentsReadRequest) ExcludeFields(excludeFields []string) ContentGenericContentsApiContentDebGenericContentsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentGenericContentsApiContentDebGenericContentsReadRequest) Execute() (*DebGenericContentResponse, *http.Response, error) {
	return r.ApiService.ContentDebGenericContentsReadExecute(r)
}

/*
ContentDebGenericContentsRead Inspect a generic content

GenericContent is a catch all category for storing files not covered by any other type.

Associated artifacts: Exactly one arbitrary file that does not match any other type.

This is needed to store arbitrary files for use with the verbatim publisher. If you are not
using the verbatim publisher, you may ignore this type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debGenericContentHref
 @return ContentGenericContentsApiContentDebGenericContentsReadRequest
*/
func (a *ContentGenericContentsApiService) ContentDebGenericContentsRead(ctx context.Context, debGenericContentHref string) ContentGenericContentsApiContentDebGenericContentsReadRequest {
	return ContentGenericContentsApiContentDebGenericContentsReadRequest{
		ApiService: a,
		ctx: ctx,
		debGenericContentHref: debGenericContentHref,
	}
}

// Execute executes the request
//  @return DebGenericContentResponse
func (a *ContentGenericContentsApiService) ContentDebGenericContentsReadExecute(r ContentGenericContentsApiContentDebGenericContentsReadRequest) (*DebGenericContentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebGenericContentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentGenericContentsApiService.ContentDebGenericContentsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_generic_content_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_generic_content_href"+"}", url.PathEscape(parameterValueToString(r.debGenericContentHref, "debGenericContentHref")), -1)
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
