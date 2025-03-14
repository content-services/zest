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


// ContentGemAPIService ContentGemAPI service
type ContentGemAPIService service

type ContentGemAPIContentGemGemCreateRequest struct {
	ctx context.Context
	ApiService *ContentGemAPIService
	pulpDomain string
	repository *string
	pulpLabels *map[string]string
	artifact *string
	file *os.File
}

// A URI of a repository the new content unit should be associated with.
func (r ContentGemAPIContentGemGemCreateRequest) Repository(repository string) ContentGemAPIContentGemGemCreateRequest {
	r.repository = &repository
	return r
}

// A dictionary of arbitrary key/value pairs used to describe a specific Content instance.
func (r ContentGemAPIContentGemGemCreateRequest) PulpLabels(pulpLabels map[string]string) ContentGemAPIContentGemGemCreateRequest {
	r.pulpLabels = &pulpLabels
	return r
}

// Artifact file representing the physical content
func (r ContentGemAPIContentGemGemCreateRequest) Artifact(artifact string) ContentGemAPIContentGemGemCreateRequest {
	r.artifact = &artifact
	return r
}

// An uploaded file that should be turned into the artifact of the content unit.
func (r ContentGemAPIContentGemGemCreateRequest) File(file *os.File) ContentGemAPIContentGemGemCreateRequest {
	r.file = file
	return r
}

func (r ContentGemAPIContentGemGemCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentGemGemCreateExecute(r)
}

/*
ContentGemGemCreate Create a gem content

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentGemAPIContentGemGemCreateRequest
*/
func (a *ContentGemAPIService) ContentGemGemCreate(ctx context.Context, pulpDomain string) ContentGemAPIContentGemGemCreateRequest {
	return ContentGemAPIContentGemGemCreateRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentGemAPIService) ContentGemGemCreateExecute(r ContentGemAPIContentGemGemCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentGemAPIService.ContentGemGemCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/gem/gem/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

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
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "repository", r.repository, "", "")
	}
	if r.pulpLabels != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pulp_labels", r.pulpLabels, "", "")
	}
	if r.artifact != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "artifact", r.artifact, "", "")
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

type ContentGemAPIContentGemGemListRequest struct {
	ctx context.Context
	ApiService *ContentGemAPIService
	pulpDomain string
	checksum *string
	limit *int32
	name *string
	offset *int32
	ordering *[]string
	orphanedFor *float32
	prerelease *bool
	prnIn *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	pulpLabelSelect *string
	q *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	version *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where checksum matches value
func (r ContentGemAPIContentGemGemListRequest) Checksum(checksum string) ContentGemAPIContentGemGemListRequest {
	r.checksum = &checksum
	return r
}

// Number of results to return per page.
func (r ContentGemAPIContentGemGemListRequest) Limit(limit int32) ContentGemAPIContentGemGemListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r ContentGemAPIContentGemGemListRequest) Name(name string) ContentGemAPIContentGemGemListRequest {
	r.name = &name
	return r
}

// The initial index from which to return the results.
func (r ContentGemAPIContentGemGemListRequest) Offset(offset int32) ContentGemAPIContentGemGemListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;version&#x60; - Version* &#x60;-version&#x60; - Version (descending)* &#x60;platform&#x60; - Platform* &#x60;-platform&#x60; - Platform (descending)* &#x60;checksum&#x60; - Checksum* &#x60;-checksum&#x60; - Checksum (descending)* &#x60;prerelease&#x60; - Prerelease* &#x60;-prerelease&#x60; - Prerelease (descending)* &#x60;dependencies&#x60; - Dependencies* &#x60;-dependencies&#x60; - Dependencies (descending)* &#x60;required_ruby_version&#x60; - Required ruby version* &#x60;-required_ruby_version&#x60; - Required ruby version (descending)* &#x60;required_rubygems_version&#x60; - Required rubygems version* &#x60;-required_rubygems_version&#x60; - Required rubygems version (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r ContentGemAPIContentGemGemListRequest) Ordering(ordering []string) ContentGemAPIContentGemGemListRequest {
	r.ordering = &ordering
	return r
}

// Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME.
func (r ContentGemAPIContentGemGemListRequest) OrphanedFor(orphanedFor float32) ContentGemAPIContentGemGemListRequest {
	r.orphanedFor = &orphanedFor
	return r
}

// Filter results where prerelease matches value
func (r ContentGemAPIContentGemGemListRequest) Prerelease(prerelease bool) ContentGemAPIContentGemGemListRequest {
	r.prerelease = &prerelease
	return r
}

// Multiple values may be separated by commas.
func (r ContentGemAPIContentGemGemListRequest) PrnIn(prnIn []string) ContentGemAPIContentGemGemListRequest {
	r.prnIn = &prnIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentGemAPIContentGemGemListRequest) PulpHrefIn(pulpHrefIn []string) ContentGemAPIContentGemGemListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentGemAPIContentGemGemListRequest) PulpIdIn(pulpIdIn []string) ContentGemAPIContentGemGemListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter labels by search string
func (r ContentGemAPIContentGemGemListRequest) PulpLabelSelect(pulpLabelSelect string) ContentGemAPIContentGemGemListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Filter results by using NOT, AND and OR operations on other filters
func (r ContentGemAPIContentGemGemListRequest) Q(q string) ContentGemAPIContentGemGemListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentGemAPIContentGemGemListRequest) RepositoryVersion(repositoryVersion string) ContentGemAPIContentGemGemListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentGemAPIContentGemGemListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentGemAPIContentGemGemListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentGemAPIContentGemGemListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentGemAPIContentGemGemListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter results where version matches value
func (r ContentGemAPIContentGemGemListRequest) Version(version string) ContentGemAPIContentGemGemListRequest {
	r.version = &version
	return r
}

// A list of fields to include in the response.
func (r ContentGemAPIContentGemGemListRequest) Fields(fields []string) ContentGemAPIContentGemGemListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentGemAPIContentGemGemListRequest) ExcludeFields(excludeFields []string) ContentGemAPIContentGemGemListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentGemAPIContentGemGemListRequest) Execute() (*PaginatedgemGemContentResponseList, *http.Response, error) {
	return r.ApiService.ContentGemGemListExecute(r)
}

/*
ContentGemGemList List gem contents

A ViewSet for GemContent.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentGemAPIContentGemGemListRequest
*/
func (a *ContentGemAPIService) ContentGemGemList(ctx context.Context, pulpDomain string) ContentGemAPIContentGemGemListRequest {
	return ContentGemAPIContentGemGemListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedgemGemContentResponseList
func (a *ContentGemAPIService) ContentGemGemListExecute(r ContentGemAPIContentGemGemListRequest) (*PaginatedgemGemContentResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedgemGemContentResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentGemAPIService.ContentGemGemList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/gem/gem/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.checksum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "checksum", r.checksum, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
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
	if r.prerelease != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prerelease", r.prerelease, "form", "")
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
	if r.pulpLabelSelect != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_label_select", r.pulpLabelSelect, "form", "")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
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
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "form", "")
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

type ContentGemAPIContentGemGemReadRequest struct {
	ctx context.Context
	ApiService *ContentGemAPIService
	gemGemContentHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentGemAPIContentGemGemReadRequest) Fields(fields []string) ContentGemAPIContentGemGemReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentGemAPIContentGemGemReadRequest) ExcludeFields(excludeFields []string) ContentGemAPIContentGemGemReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentGemAPIContentGemGemReadRequest) Execute() (*GemGemContentResponse, *http.Response, error) {
	return r.ApiService.ContentGemGemReadExecute(r)
}

/*
ContentGemGemRead Inspect a gem content

A ViewSet for GemContent.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gemGemContentHref
 @return ContentGemAPIContentGemGemReadRequest
*/
func (a *ContentGemAPIService) ContentGemGemRead(ctx context.Context, gemGemContentHref string) ContentGemAPIContentGemGemReadRequest {
	return ContentGemAPIContentGemGemReadRequest{
		ApiService: a,
		ctx: ctx,
		gemGemContentHref: gemGemContentHref,
	}
}

// Execute executes the request
//  @return GemGemContentResponse
func (a *ContentGemAPIService) ContentGemGemReadExecute(r ContentGemAPIContentGemGemReadRequest) (*GemGemContentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GemGemContentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentGemAPIService.ContentGemGemRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{gem_gem_content_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"gem_gem_content_href"+"}", url.PathEscape(parameterValueToString(r.gemGemContentHref, "gemGemContentHref")), -1)
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

type ContentGemAPIContentGemGemSetLabelRequest struct {
	ctx context.Context
	ApiService *ContentGemAPIService
	gemGemContentHref string
	setLabel *SetLabel
}

func (r ContentGemAPIContentGemGemSetLabelRequest) SetLabel(setLabel SetLabel) ContentGemAPIContentGemGemSetLabelRequest {
	r.setLabel = &setLabel
	return r
}

func (r ContentGemAPIContentGemGemSetLabelRequest) Execute() (*SetLabelResponse, *http.Response, error) {
	return r.ApiService.ContentGemGemSetLabelExecute(r)
}

/*
ContentGemGemSetLabel Set a label

Set a single pulp_label on the object to a specific value or null.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gemGemContentHref
 @return ContentGemAPIContentGemGemSetLabelRequest
*/
func (a *ContentGemAPIService) ContentGemGemSetLabel(ctx context.Context, gemGemContentHref string) ContentGemAPIContentGemGemSetLabelRequest {
	return ContentGemAPIContentGemGemSetLabelRequest{
		ApiService: a,
		ctx: ctx,
		gemGemContentHref: gemGemContentHref,
	}
}

// Execute executes the request
//  @return SetLabelResponse
func (a *ContentGemAPIService) ContentGemGemSetLabelExecute(r ContentGemAPIContentGemGemSetLabelRequest) (*SetLabelResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SetLabelResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentGemAPIService.ContentGemGemSetLabel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{gem_gem_content_href}set_label/"
	localVarPath = strings.Replace(localVarPath, "{"+"gem_gem_content_href"+"}", url.PathEscape(parameterValueToString(r.gemGemContentHref, "gemGemContentHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.setLabel == nil {
		return localVarReturnValue, nil, reportError("setLabel is required and must be specified")
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
	localVarPostBody = r.setLabel
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

type ContentGemAPIContentGemGemUnsetLabelRequest struct {
	ctx context.Context
	ApiService *ContentGemAPIService
	gemGemContentHref string
	unsetLabel *UnsetLabel
}

func (r ContentGemAPIContentGemGemUnsetLabelRequest) UnsetLabel(unsetLabel UnsetLabel) ContentGemAPIContentGemGemUnsetLabelRequest {
	r.unsetLabel = &unsetLabel
	return r
}

func (r ContentGemAPIContentGemGemUnsetLabelRequest) Execute() (*UnsetLabelResponse, *http.Response, error) {
	return r.ApiService.ContentGemGemUnsetLabelExecute(r)
}

/*
ContentGemGemUnsetLabel Unset a label

Unset a single pulp_label on the object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gemGemContentHref
 @return ContentGemAPIContentGemGemUnsetLabelRequest
*/
func (a *ContentGemAPIService) ContentGemGemUnsetLabel(ctx context.Context, gemGemContentHref string) ContentGemAPIContentGemGemUnsetLabelRequest {
	return ContentGemAPIContentGemGemUnsetLabelRequest{
		ApiService: a,
		ctx: ctx,
		gemGemContentHref: gemGemContentHref,
	}
}

// Execute executes the request
//  @return UnsetLabelResponse
func (a *ContentGemAPIService) ContentGemGemUnsetLabelExecute(r ContentGemAPIContentGemGemUnsetLabelRequest) (*UnsetLabelResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UnsetLabelResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentGemAPIService.ContentGemGemUnsetLabel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{gem_gem_content_href}unset_label/"
	localVarPath = strings.Replace(localVarPath, "{"+"gem_gem_content_href"+"}", url.PathEscape(parameterValueToString(r.gemGemContentHref, "gemGemContentHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unsetLabel == nil {
		return localVarReturnValue, nil, reportError("unsetLabel is required and must be specified")
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
	localVarPostBody = r.unsetLabel
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
