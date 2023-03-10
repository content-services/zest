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


// ContentCollectionVersionsApiService ContentCollectionVersionsApi service
type ContentCollectionVersionsApiService service

type ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest struct {
	ctx context.Context
	ApiService *ContentCollectionVersionsApiService
	artifact *string
	upload *string
	file *os.File
	repository *string
	expectedName *string
	expectedNamespace *string
	expectedVersion *string
}

// Artifact file representing the physical content
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest) Artifact(artifact string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest {
	r.artifact = &artifact
	return r
}

// An uncommitted upload that may be turned into the artifact of the content unit.
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest) Upload(upload string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest {
	r.upload = &upload
	return r
}

// An uploaded file that may be turned into the artifact of the content unit.
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest) File(file *os.File) ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest {
	r.file = file
	return r
}

// A URI of a repository the new content unit should be associated with.
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest) Repository(repository string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest {
	r.repository = &repository
	return r
}

// The name of the collection.
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest) ExpectedName(expectedName string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest {
	r.expectedName = &expectedName
	return r
}

// The namespace of the collection.
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest) ExpectedNamespace(expectedNamespace string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest {
	r.expectedNamespace = &expectedNamespace
	return r
}

// The version of the collection.
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest) ExpectedVersion(expectedVersion string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest {
	r.expectedVersion = &expectedVersion
	return r
}

func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentAnsibleCollectionVersionsCreateExecute(r)
}

/*
ContentAnsibleCollectionVersionsCreate Create a collection version

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest
*/
func (a *ContentCollectionVersionsApiService) ContentAnsibleCollectionVersionsCreate(ctx context.Context) ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest {
	return ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentCollectionVersionsApiService) ContentAnsibleCollectionVersionsCreateExecute(r ContentCollectionVersionsApiContentAnsibleCollectionVersionsCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCollectionVersionsApiService.ContentAnsibleCollectionVersionsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/ansible/collection_versions/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.upload != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "upload", r.upload, "")
	}
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
	if r.expectedName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_name", r.expectedName, "")
	}
	if r.expectedNamespace != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_namespace", r.expectedNamespace, "")
	}
	if r.expectedVersion != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_version", r.expectedVersion, "")
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

type ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest struct {
	ctx context.Context
	ApiService *ContentCollectionVersionsApiService
	isHighest *bool
	limit *int32
	name *string
	namespace *string
	offset *int32
	ordering *[]string
	q *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	tags *string
	version *string
	fields *[]string
	excludeFields *[]string
}

func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest) IsHighest(isHighest bool) ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest {
	r.isHighest = &isHighest
	return r
}

// Number of results to return per page.
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest) Limit(limit int32) ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest {
	r.limit = &limit
	return r
}

func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest) Name(name string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest {
	r.name = &name
	return r
}

func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest) Namespace(namespace string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest {
	r.namespace = &namespace
	return r
}

// The initial index from which to return the results.
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest) Offset(offset int32) ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest) Ordering(ordering []string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest {
	r.ordering = &ordering
	return r
}

func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest) Q(q string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest) RepositoryVersion(repositoryVersion string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter by comma separate list of tags that must all be matched
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest) Tags(tags string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest {
	r.tags = &tags
	return r
}

// Filter results where version matches value
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest) Version(version string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest {
	r.version = &version
	return r
}

// A list of fields to include in the response.
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest) Fields(fields []string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest) ExcludeFields(excludeFields []string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest) Execute() (*PaginatedansibleCollectionVersionResponseList, *http.Response, error) {
	return r.ApiService.ContentAnsibleCollectionVersionsListExecute(r)
}

/*
ContentAnsibleCollectionVersionsList List collection versions

ViewSet for Ansible Collection.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest
*/
func (a *ContentCollectionVersionsApiService) ContentAnsibleCollectionVersionsList(ctx context.Context) ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest {
	return ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedansibleCollectionVersionResponseList
func (a *ContentCollectionVersionsApiService) ContentAnsibleCollectionVersionsListExecute(r ContentCollectionVersionsApiContentAnsibleCollectionVersionsListRequest) (*PaginatedansibleCollectionVersionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedansibleCollectionVersionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCollectionVersionsApiService.ContentAnsibleCollectionVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/ansible/collection_versions/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isHighest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_highest", r.isHighest, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.namespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
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
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tags", r.tags, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
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

type ContentCollectionVersionsApiContentAnsibleCollectionVersionsReadRequest struct {
	ctx context.Context
	ApiService *ContentCollectionVersionsApiService
	ansibleCollectionVersionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsReadRequest) Fields(fields []string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsReadRequest) ExcludeFields(excludeFields []string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentCollectionVersionsApiContentAnsibleCollectionVersionsReadRequest) Execute() (*AnsibleCollectionVersionResponse, *http.Response, error) {
	return r.ApiService.ContentAnsibleCollectionVersionsReadExecute(r)
}

/*
ContentAnsibleCollectionVersionsRead Inspect a collection version

ViewSet for Ansible Collection.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleCollectionVersionHref
 @return ContentCollectionVersionsApiContentAnsibleCollectionVersionsReadRequest
*/
func (a *ContentCollectionVersionsApiService) ContentAnsibleCollectionVersionsRead(ctx context.Context, ansibleCollectionVersionHref string) ContentCollectionVersionsApiContentAnsibleCollectionVersionsReadRequest {
	return ContentCollectionVersionsApiContentAnsibleCollectionVersionsReadRequest{
		ApiService: a,
		ctx: ctx,
		ansibleCollectionVersionHref: ansibleCollectionVersionHref,
	}
}

// Execute executes the request
//  @return AnsibleCollectionVersionResponse
func (a *ContentCollectionVersionsApiService) ContentAnsibleCollectionVersionsReadExecute(r ContentCollectionVersionsApiContentAnsibleCollectionVersionsReadRequest) (*AnsibleCollectionVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AnsibleCollectionVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCollectionVersionsApiService.ContentAnsibleCollectionVersionsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ansible_collection_version_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_collection_version_href"+"}", url.PathEscape(parameterValueToString(r.ansibleCollectionVersionHref, "ansibleCollectionVersionHref")), -1)
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
