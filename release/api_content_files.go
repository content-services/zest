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


// ContentFilesApiService ContentFilesApi service
type ContentFilesApiService service

type ContentFilesApiContentFileFilesCreateRequest struct {
	ctx context.Context
	ApiService *ContentFilesApiService
	relativePath *string
	artifact *string
	file *os.File
	repository *string
	upload *string
}

// Path where the artifact is located relative to distributions base_path
func (r ContentFilesApiContentFileFilesCreateRequest) RelativePath(relativePath string) ContentFilesApiContentFileFilesCreateRequest {
	r.relativePath = &relativePath
	return r
}

// Artifact file representing the physical content
func (r ContentFilesApiContentFileFilesCreateRequest) Artifact(artifact string) ContentFilesApiContentFileFilesCreateRequest {
	r.artifact = &artifact
	return r
}

// An uploaded file that may be turned into the artifact of the content unit.
func (r ContentFilesApiContentFileFilesCreateRequest) File(file *os.File) ContentFilesApiContentFileFilesCreateRequest {
	r.file = file
	return r
}

// A URI of a repository the new content unit should be associated with.
func (r ContentFilesApiContentFileFilesCreateRequest) Repository(repository string) ContentFilesApiContentFileFilesCreateRequest {
	r.repository = &repository
	return r
}

// An uncommitted upload that may be turned into the artifact of the content unit.
func (r ContentFilesApiContentFileFilesCreateRequest) Upload(upload string) ContentFilesApiContentFileFilesCreateRequest {
	r.upload = &upload
	return r
}

func (r ContentFilesApiContentFileFilesCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentFileFilesCreateExecute(r)
}

/*
ContentFileFilesCreate Create a file content

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentFilesApiContentFileFilesCreateRequest
*/
func (a *ContentFilesApiService) ContentFileFilesCreate(ctx context.Context) ContentFilesApiContentFileFilesCreateRequest {
	return ContentFilesApiContentFileFilesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentFilesApiService) ContentFileFilesCreateExecute(r ContentFilesApiContentFileFilesCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentFilesApiService.ContentFileFilesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/file/files/"
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

type ContentFilesApiContentFileFilesListRequest struct {
	ctx context.Context
	ApiService *ContentFilesApiService
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
func (r ContentFilesApiContentFileFilesListRequest) Limit(limit int32) ContentFilesApiContentFileFilesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentFilesApiContentFileFilesListRequest) Offset(offset int32) ContentFilesApiContentFileFilesListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ContentFilesApiContentFileFilesListRequest) Ordering(ordering []string) ContentFilesApiContentFileFilesListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where relative_path matches value
func (r ContentFilesApiContentFileFilesListRequest) RelativePath(relativePath string) ContentFilesApiContentFileFilesListRequest {
	r.relativePath = &relativePath
	return r
}

// Repository Version referenced by HREF
func (r ContentFilesApiContentFileFilesListRequest) RepositoryVersion(repositoryVersion string) ContentFilesApiContentFileFilesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentFilesApiContentFileFilesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentFilesApiContentFileFilesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentFilesApiContentFileFilesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentFilesApiContentFileFilesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

func (r ContentFilesApiContentFileFilesListRequest) Sha256(sha256 string) ContentFilesApiContentFileFilesListRequest {
	r.sha256 = &sha256
	return r
}

// A list of fields to include in the response.
func (r ContentFilesApiContentFileFilesListRequest) Fields(fields []string) ContentFilesApiContentFileFilesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentFilesApiContentFileFilesListRequest) ExcludeFields(excludeFields []string) ContentFilesApiContentFileFilesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentFilesApiContentFileFilesListRequest) Execute() (*PaginatedfileFileContentResponseList, *http.Response, error) {
	return r.ApiService.ContentFileFilesListExecute(r)
}

/*
ContentFileFilesList List file contents


FileContent represents a single file and its metadata, which can be added and removed from
repositories.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentFilesApiContentFileFilesListRequest
*/
func (a *ContentFilesApiService) ContentFileFilesList(ctx context.Context) ContentFilesApiContentFileFilesListRequest {
	return ContentFilesApiContentFileFilesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedfileFileContentResponseList
func (a *ContentFilesApiService) ContentFileFilesListExecute(r ContentFilesApiContentFileFilesListRequest) (*PaginatedfileFileContentResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedfileFileContentResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentFilesApiService.ContentFileFilesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/file/files/"
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

type ContentFilesApiContentFileFilesReadRequest struct {
	ctx context.Context
	ApiService *ContentFilesApiService
	fileFileContentHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentFilesApiContentFileFilesReadRequest) Fields(fields []string) ContentFilesApiContentFileFilesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentFilesApiContentFileFilesReadRequest) ExcludeFields(excludeFields []string) ContentFilesApiContentFileFilesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentFilesApiContentFileFilesReadRequest) Execute() (*FileFileContentResponse, *http.Response, error) {
	return r.ApiService.ContentFileFilesReadExecute(r)
}

/*
ContentFileFilesRead Inspect a file content


FileContent represents a single file and its metadata, which can be added and removed from
repositories.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileFileContentHref
 @return ContentFilesApiContentFileFilesReadRequest
*/
func (a *ContentFilesApiService) ContentFileFilesRead(ctx context.Context, fileFileContentHref string) ContentFilesApiContentFileFilesReadRequest {
	return ContentFilesApiContentFileFilesReadRequest{
		ApiService: a,
		ctx: ctx,
		fileFileContentHref: fileFileContentHref,
	}
}

// Execute executes the request
//  @return FileFileContentResponse
func (a *ContentFilesApiService) ContentFileFilesReadExecute(r ContentFilesApiContentFileFilesReadRequest) (*FileFileContentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileFileContentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentFilesApiService.ContentFileFilesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{file_file_content_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"file_file_content_href"+"}", url.PathEscape(parameterValueToString(r.fileFileContentHref, "fileFileContentHref")), -1)
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
