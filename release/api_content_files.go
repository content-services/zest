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
	"os"
	"reflect"
)


// ContentFilesAPIService ContentFilesAPI service
type ContentFilesAPIService service

type ContentFilesAPIContentFileFilesCreateRequest struct {
	ctx context.Context
	ApiService *ContentFilesAPIService
	pulpDomain string
	relativePath *string
	repository *string
	artifact *string
	file *os.File
	upload *string
	fileUrl *string
}

// Path where the artifact is located relative to distributions base_path
func (r ContentFilesAPIContentFileFilesCreateRequest) RelativePath(relativePath string) ContentFilesAPIContentFileFilesCreateRequest {
	r.relativePath = &relativePath
	return r
}

// A URI of a repository the new content unit should be associated with.
func (r ContentFilesAPIContentFileFilesCreateRequest) Repository(repository string) ContentFilesAPIContentFileFilesCreateRequest {
	r.repository = &repository
	return r
}

// Artifact file representing the physical content
func (r ContentFilesAPIContentFileFilesCreateRequest) Artifact(artifact string) ContentFilesAPIContentFileFilesCreateRequest {
	r.artifact = &artifact
	return r
}

// An uploaded file that may be turned into the content unit.
func (r ContentFilesAPIContentFileFilesCreateRequest) File(file *os.File) ContentFilesAPIContentFileFilesCreateRequest {
	r.file = file
	return r
}

// An uncommitted upload that may be turned into the content unit.
func (r ContentFilesAPIContentFileFilesCreateRequest) Upload(upload string) ContentFilesAPIContentFileFilesCreateRequest {
	r.upload = &upload
	return r
}

// A url that Pulp can download and turn into the content unit.
func (r ContentFilesAPIContentFileFilesCreateRequest) FileUrl(fileUrl string) ContentFilesAPIContentFileFilesCreateRequest {
	r.fileUrl = &fileUrl
	return r
}

func (r ContentFilesAPIContentFileFilesCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentFileFilesCreateExecute(r)
}

/*
ContentFileFilesCreate Create a file content

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentFilesAPIContentFileFilesCreateRequest
*/
func (a *ContentFilesAPIService) ContentFileFilesCreate(ctx context.Context, pulpDomain string) ContentFilesAPIContentFileFilesCreateRequest {
	return ContentFilesAPIContentFileFilesCreateRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentFilesAPIService) ContentFileFilesCreateExecute(r ContentFilesAPIContentFileFilesCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentFilesAPIService.ContentFileFilesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/file/files/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

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
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "repository", r.repository, "", "")
	}
	if r.artifact != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "artifact", r.artifact, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "relative_path", r.relativePath, "", "")
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
	if r.upload != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "upload", r.upload, "", "")
	}
	if r.fileUrl != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "file_url", r.fileUrl, "", "")
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

type ContentFilesAPIContentFileFilesListRequest struct {
	ctx context.Context
	ApiService *ContentFilesAPIService
	pulpDomain string
	limit *int32
	offset *int32
	ordering *[]string
	orphanedFor *float32
	prnIn *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	q *string
	relativePath *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	sha256 *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ContentFilesAPIContentFileFilesListRequest) Limit(limit int32) ContentFilesAPIContentFileFilesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentFilesAPIContentFileFilesListRequest) Offset(offset int32) ContentFilesAPIContentFileFilesListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;relative_path&#x60; - Relative path* &#x60;-relative_path&#x60; - Relative path (descending)* &#x60;digest&#x60; - Digest* &#x60;-digest&#x60; - Digest (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r ContentFilesAPIContentFileFilesListRequest) Ordering(ordering []string) ContentFilesAPIContentFileFilesListRequest {
	r.ordering = &ordering
	return r
}

// Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME.
func (r ContentFilesAPIContentFileFilesListRequest) OrphanedFor(orphanedFor float32) ContentFilesAPIContentFileFilesListRequest {
	r.orphanedFor = &orphanedFor
	return r
}

// Multiple values may be separated by commas.
func (r ContentFilesAPIContentFileFilesListRequest) PrnIn(prnIn []string) ContentFilesAPIContentFileFilesListRequest {
	r.prnIn = &prnIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentFilesAPIContentFileFilesListRequest) PulpHrefIn(pulpHrefIn []string) ContentFilesAPIContentFileFilesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentFilesAPIContentFileFilesListRequest) PulpIdIn(pulpIdIn []string) ContentFilesAPIContentFileFilesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter results by using NOT, AND and OR operations on other filters
func (r ContentFilesAPIContentFileFilesListRequest) Q(q string) ContentFilesAPIContentFileFilesListRequest {
	r.q = &q
	return r
}

// Filter results where relative_path matches value
func (r ContentFilesAPIContentFileFilesListRequest) RelativePath(relativePath string) ContentFilesAPIContentFileFilesListRequest {
	r.relativePath = &relativePath
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentFilesAPIContentFileFilesListRequest) RepositoryVersion(repositoryVersion string) ContentFilesAPIContentFileFilesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentFilesAPIContentFileFilesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentFilesAPIContentFileFilesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentFilesAPIContentFileFilesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentFilesAPIContentFileFilesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

func (r ContentFilesAPIContentFileFilesListRequest) Sha256(sha256 string) ContentFilesAPIContentFileFilesListRequest {
	r.sha256 = &sha256
	return r
}

// A list of fields to include in the response.
func (r ContentFilesAPIContentFileFilesListRequest) Fields(fields []string) ContentFilesAPIContentFileFilesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentFilesAPIContentFileFilesListRequest) ExcludeFields(excludeFields []string) ContentFilesAPIContentFileFilesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentFilesAPIContentFileFilesListRequest) Execute() (*PaginatedfileFileContentResponseList, *http.Response, error) {
	return r.ApiService.ContentFileFilesListExecute(r)
}

/*
ContentFileFilesList List file contents

FileContent represents a single file and its metadata, which can be added and removed fromrepositories.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentFilesAPIContentFileFilesListRequest
*/
func (a *ContentFilesAPIService) ContentFileFilesList(ctx context.Context, pulpDomain string) ContentFilesAPIContentFileFilesListRequest {
	return ContentFilesAPIContentFileFilesListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedfileFileContentResponseList
func (a *ContentFilesAPIService) ContentFileFilesListExecute(r ContentFilesAPIContentFileFilesListRequest) (*PaginatedfileFileContentResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedfileFileContentResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentFilesAPIService.ContentFileFilesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/file/files/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "form", "csv")
	}
	if r.orphanedFor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orphaned_for", r.orphanedFor, "form", "")
	}
	if r.prnIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prn__in", r.prnIn, "form", "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "form", "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "form", "csv")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
	}
	if r.relativePath != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "relative_path", r.relativePath, "form", "")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "form", "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "form", "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "form", "")
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha256", r.sha256, "form", "")
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

type ContentFilesAPIContentFileFilesReadRequest struct {
	ctx context.Context
	ApiService *ContentFilesAPIService
	fileFileContentHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentFilesAPIContentFileFilesReadRequest) Fields(fields []string) ContentFilesAPIContentFileFilesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentFilesAPIContentFileFilesReadRequest) ExcludeFields(excludeFields []string) ContentFilesAPIContentFileFilesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentFilesAPIContentFileFilesReadRequest) Execute() (*FileFileContentResponse, *http.Response, error) {
	return r.ApiService.ContentFileFilesReadExecute(r)
}

/*
ContentFileFilesRead Inspect a file content

FileContent represents a single file and its metadata, which can be added and removed fromrepositories.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileFileContentHref
 @return ContentFilesAPIContentFileFilesReadRequest
*/
func (a *ContentFilesAPIService) ContentFileFilesRead(ctx context.Context, fileFileContentHref string) ContentFilesAPIContentFileFilesReadRequest {
	return ContentFilesAPIContentFileFilesReadRequest{
		ApiService: a,
		ctx: ctx,
		fileFileContentHref: fileFileContentHref,
	}
}

// Execute executes the request
//  @return FileFileContentResponse
func (a *ContentFilesAPIService) ContentFileFilesReadExecute(r ContentFilesAPIContentFileFilesReadRequest) (*FileFileContentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileFileContentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentFilesAPIService.ContentFileFilesRead")
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
