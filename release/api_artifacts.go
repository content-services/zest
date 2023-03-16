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
	"os"
	"reflect"
)


// ArtifactsApiService ArtifactsApi service
type ArtifactsApiService service

type ArtifactsApiArtifactsCreateRequest struct {
	ctx context.Context
	ApiService *ArtifactsApiService
	file *os.File
	size *int64
	md5 *string
	sha1 *string
	sha224 *string
	sha256 *string
	sha384 *string
	sha512 *string
}

// The stored file.
func (r ArtifactsApiArtifactsCreateRequest) File(file *os.File) ArtifactsApiArtifactsCreateRequest {
	r.file = file
	return r
}

// The size of the file in bytes.
func (r ArtifactsApiArtifactsCreateRequest) Size(size int64) ArtifactsApiArtifactsCreateRequest {
	r.size = &size
	return r
}

// The MD5 checksum of the file if available.
func (r ArtifactsApiArtifactsCreateRequest) Md5(md5 string) ArtifactsApiArtifactsCreateRequest {
	r.md5 = &md5
	return r
}

// The SHA-1 checksum of the file if available.
func (r ArtifactsApiArtifactsCreateRequest) Sha1(sha1 string) ArtifactsApiArtifactsCreateRequest {
	r.sha1 = &sha1
	return r
}

// The SHA-224 checksum of the file if available.
func (r ArtifactsApiArtifactsCreateRequest) Sha224(sha224 string) ArtifactsApiArtifactsCreateRequest {
	r.sha224 = &sha224
	return r
}

// The SHA-256 checksum of the file if available.
func (r ArtifactsApiArtifactsCreateRequest) Sha256(sha256 string) ArtifactsApiArtifactsCreateRequest {
	r.sha256 = &sha256
	return r
}

// The SHA-384 checksum of the file if available.
func (r ArtifactsApiArtifactsCreateRequest) Sha384(sha384 string) ArtifactsApiArtifactsCreateRequest {
	r.sha384 = &sha384
	return r
}

// The SHA-512 checksum of the file if available.
func (r ArtifactsApiArtifactsCreateRequest) Sha512(sha512 string) ArtifactsApiArtifactsCreateRequest {
	r.sha512 = &sha512
	return r
}

func (r ArtifactsApiArtifactsCreateRequest) Execute() (*ArtifactResponse, *http.Response, error) {
	return r.ApiService.ArtifactsCreateExecute(r)
}

/*
ArtifactsCreate Create an artifact

A customized named ModelViewSet that knows how to register itself with the Pulp API router.

This viewset is discoverable by its name.
"Normal" Django Models and Master/Detail models are supported by the ``register_with`` method.

Attributes:
    lookup_field (str): The name of the field by which an object should be looked up, in
        addition to any parent lookups if this ViewSet is nested. Defaults to 'pk'
    endpoint_name (str): The name of the final path segment that should identify the ViewSet's
        collection endpoint.
    nest_prefix (str): Optional prefix under which this ViewSet should be nested. This must
        correspond to the "parent_prefix" of a router with rest_framework_nested.NestedMixin.
        None indicates this ViewSet should not be nested.
    parent_lookup_kwargs (dict): Optional mapping of key names that would appear in self.kwargs
        to django model filter expressions that can be used with the corresponding value from
        self.kwargs, used only by a nested ViewSet to filter based on the parent object's
        identity.
    schema (DefaultSchema): The schema class to use by default in a viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ArtifactsApiArtifactsCreateRequest
*/
func (a *ArtifactsApiService) ArtifactsCreate(ctx context.Context) ArtifactsApiArtifactsCreateRequest {
	return ArtifactsApiArtifactsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ArtifactResponse
func (a *ArtifactsApiService) ArtifactsCreateExecute(r ArtifactsApiArtifactsCreateRequest) (*ArtifactResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactsApiService.ArtifactsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/artifacts/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
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
	if r.size != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "size", r.size, "")
	}
	if r.md5 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "md5", r.md5, "")
	}
	if r.sha1 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha1", r.sha1, "")
	}
	if r.sha224 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha224", r.sha224, "")
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha256", r.sha256, "")
	}
	if r.sha384 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha384", r.sha384, "")
	}
	if r.sha512 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha512", r.sha512, "")
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

type ArtifactsApiArtifactsDeleteRequest struct {
	ctx context.Context
	ApiService *ArtifactsApiService
	artifactHref string
}

func (r ArtifactsApiArtifactsDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ArtifactsDeleteExecute(r)
}

/*
ArtifactsDelete Delete an artifact

Remove Artifact only if it is not associated with any Content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param artifactHref
 @return ArtifactsApiArtifactsDeleteRequest
*/
func (a *ArtifactsApiService) ArtifactsDelete(ctx context.Context, artifactHref string) ArtifactsApiArtifactsDeleteRequest {
	return ArtifactsApiArtifactsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		artifactHref: artifactHref,
	}
}

// Execute executes the request
func (a *ArtifactsApiService) ArtifactsDeleteExecute(r ArtifactsApiArtifactsDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactsApiService.ArtifactsDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{artifact_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"artifact_href"+"}", url.PathEscape(parameterValueToString(r.artifactHref, "artifactHref")), -1)
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

type ArtifactsApiArtifactsListRequest struct {
	ctx context.Context
	ApiService *ArtifactsApiService
	limit *int32
	md5 *string
	offset *int32
	ordering *[]string
	repositoryVersion *string
	sha1 *string
	sha224 *string
	sha256 *string
	sha384 *string
	sha512 *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ArtifactsApiArtifactsListRequest) Limit(limit int32) ArtifactsApiArtifactsListRequest {
	r.limit = &limit
	return r
}

// Filter results where md5 matches value
func (r ArtifactsApiArtifactsListRequest) Md5(md5 string) ArtifactsApiArtifactsListRequest {
	r.md5 = &md5
	return r
}

// The initial index from which to return the results.
func (r ArtifactsApiArtifactsListRequest) Offset(offset int32) ArtifactsApiArtifactsListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ArtifactsApiArtifactsListRequest) Ordering(ordering []string) ArtifactsApiArtifactsListRequest {
	r.ordering = &ordering
	return r
}

// Repository Version referenced by HREF
func (r ArtifactsApiArtifactsListRequest) RepositoryVersion(repositoryVersion string) ArtifactsApiArtifactsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Filter results where sha1 matches value
func (r ArtifactsApiArtifactsListRequest) Sha1(sha1 string) ArtifactsApiArtifactsListRequest {
	r.sha1 = &sha1
	return r
}

// Filter results where sha224 matches value
func (r ArtifactsApiArtifactsListRequest) Sha224(sha224 string) ArtifactsApiArtifactsListRequest {
	r.sha224 = &sha224
	return r
}

// Filter results where sha256 matches value
func (r ArtifactsApiArtifactsListRequest) Sha256(sha256 string) ArtifactsApiArtifactsListRequest {
	r.sha256 = &sha256
	return r
}

// Filter results where sha384 matches value
func (r ArtifactsApiArtifactsListRequest) Sha384(sha384 string) ArtifactsApiArtifactsListRequest {
	r.sha384 = &sha384
	return r
}

// Filter results where sha512 matches value
func (r ArtifactsApiArtifactsListRequest) Sha512(sha512 string) ArtifactsApiArtifactsListRequest {
	r.sha512 = &sha512
	return r
}

// A list of fields to include in the response.
func (r ArtifactsApiArtifactsListRequest) Fields(fields []string) ArtifactsApiArtifactsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ArtifactsApiArtifactsListRequest) ExcludeFields(excludeFields []string) ArtifactsApiArtifactsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ArtifactsApiArtifactsListRequest) Execute() (*PaginatedArtifactResponseList, *http.Response, error) {
	return r.ApiService.ArtifactsListExecute(r)
}

/*
ArtifactsList List artifacts

A customized named ModelViewSet that knows how to register itself with the Pulp API router.

This viewset is discoverable by its name.
"Normal" Django Models and Master/Detail models are supported by the ``register_with`` method.

Attributes:
    lookup_field (str): The name of the field by which an object should be looked up, in
        addition to any parent lookups if this ViewSet is nested. Defaults to 'pk'
    endpoint_name (str): The name of the final path segment that should identify the ViewSet's
        collection endpoint.
    nest_prefix (str): Optional prefix under which this ViewSet should be nested. This must
        correspond to the "parent_prefix" of a router with rest_framework_nested.NestedMixin.
        None indicates this ViewSet should not be nested.
    parent_lookup_kwargs (dict): Optional mapping of key names that would appear in self.kwargs
        to django model filter expressions that can be used with the corresponding value from
        self.kwargs, used only by a nested ViewSet to filter based on the parent object's
        identity.
    schema (DefaultSchema): The schema class to use by default in a viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ArtifactsApiArtifactsListRequest
*/
func (a *ArtifactsApiService) ArtifactsList(ctx context.Context) ArtifactsApiArtifactsListRequest {
	return ArtifactsApiArtifactsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedArtifactResponseList
func (a *ArtifactsApiService) ArtifactsListExecute(r ArtifactsApiArtifactsListRequest) (*PaginatedArtifactResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedArtifactResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactsApiService.ArtifactsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/artifacts/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.md5 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "md5", r.md5, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.sha1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha1", r.sha1, "")
	}
	if r.sha224 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha224", r.sha224, "")
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha256", r.sha256, "")
	}
	if r.sha384 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha384", r.sha384, "")
	}
	if r.sha512 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha512", r.sha512, "")
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

type ArtifactsApiArtifactsReadRequest struct {
	ctx context.Context
	ApiService *ArtifactsApiService
	artifactHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ArtifactsApiArtifactsReadRequest) Fields(fields []string) ArtifactsApiArtifactsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ArtifactsApiArtifactsReadRequest) ExcludeFields(excludeFields []string) ArtifactsApiArtifactsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ArtifactsApiArtifactsReadRequest) Execute() (*ArtifactResponse, *http.Response, error) {
	return r.ApiService.ArtifactsReadExecute(r)
}

/*
ArtifactsRead Inspect an artifact

A customized named ModelViewSet that knows how to register itself with the Pulp API router.

This viewset is discoverable by its name.
"Normal" Django Models and Master/Detail models are supported by the ``register_with`` method.

Attributes:
    lookup_field (str): The name of the field by which an object should be looked up, in
        addition to any parent lookups if this ViewSet is nested. Defaults to 'pk'
    endpoint_name (str): The name of the final path segment that should identify the ViewSet's
        collection endpoint.
    nest_prefix (str): Optional prefix under which this ViewSet should be nested. This must
        correspond to the "parent_prefix" of a router with rest_framework_nested.NestedMixin.
        None indicates this ViewSet should not be nested.
    parent_lookup_kwargs (dict): Optional mapping of key names that would appear in self.kwargs
        to django model filter expressions that can be used with the corresponding value from
        self.kwargs, used only by a nested ViewSet to filter based on the parent object's
        identity.
    schema (DefaultSchema): The schema class to use by default in a viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param artifactHref
 @return ArtifactsApiArtifactsReadRequest
*/
func (a *ArtifactsApiService) ArtifactsRead(ctx context.Context, artifactHref string) ArtifactsApiArtifactsReadRequest {
	return ArtifactsApiArtifactsReadRequest{
		ApiService: a,
		ctx: ctx,
		artifactHref: artifactHref,
	}
}

// Execute executes the request
//  @return ArtifactResponse
func (a *ArtifactsApiService) ArtifactsReadExecute(r ArtifactsApiArtifactsReadRequest) (*ArtifactResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtifactsApiService.ArtifactsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{artifact_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"artifact_href"+"}", url.PathEscape(parameterValueToString(r.artifactHref, "artifactHref")), -1)
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
