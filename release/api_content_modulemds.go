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


// ContentModulemdsAPIService ContentModulemdsAPI service
type ContentModulemdsAPIService service

type ContentModulemdsAPIContentRpmModulemdsCreateRequest struct {
	ctx context.Context
	ApiService *ContentModulemdsAPIService
	pulpDomain string
	rpmModulemd *RpmModulemd
}

func (r ContentModulemdsAPIContentRpmModulemdsCreateRequest) RpmModulemd(rpmModulemd RpmModulemd) ContentModulemdsAPIContentRpmModulemdsCreateRequest {
	r.rpmModulemd = &rpmModulemd
	return r
}

func (r ContentModulemdsAPIContentRpmModulemdsCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentRpmModulemdsCreateExecute(r)
}

/*
ContentRpmModulemdsCreate Create a modulemd

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentModulemdsAPIContentRpmModulemdsCreateRequest
*/
func (a *ContentModulemdsAPIService) ContentRpmModulemdsCreate(ctx context.Context, pulpDomain string) ContentModulemdsAPIContentRpmModulemdsCreateRequest {
	return ContentModulemdsAPIContentRpmModulemdsCreateRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentModulemdsAPIService) ContentRpmModulemdsCreateExecute(r ContentModulemdsAPIContentRpmModulemdsCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentModulemdsAPIService.ContentRpmModulemdsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/rpm/modulemds/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rpmModulemd == nil {
		return localVarReturnValue, nil, reportError("rpmModulemd is required and must be specified")
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
	localVarPostBody = r.rpmModulemd
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

type ContentModulemdsAPIContentRpmModulemdsListRequest struct {
	ctx context.Context
	ApiService *ContentModulemdsAPIService
	pulpDomain string
	arch *string
	archIn *[]string
	context *string
	contextIn *[]string
	limit *int32
	name *string
	nameIn *[]string
	offset *int32
	ordering *[]string
	orphanedFor *float32
	pulpHrefIn *[]string
	pulpIdIn *[]string
	q *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	sha256 *string
	stream *string
	streamIn *[]string
	version *string
	versionIn *[]string
	fields *[]string
	excludeFields *[]string
}

// Filter results where arch matches value
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) Arch(arch string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.arch = &arch
	return r
}

// Filter results where arch is in a comma-separated list of values
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) ArchIn(archIn []string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.archIn = &archIn
	return r
}

// Filter results where context matches value
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) Context(context string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.context = &context
	return r
}

// Filter results where context is in a comma-separated list of values
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) ContextIn(contextIn []string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.contextIn = &contextIn
	return r
}

// Number of results to return per page.
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) Limit(limit int32) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) Name(name string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.name = &name
	return r
}

// Filter results where name is in a comma-separated list of values
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) NameIn(nameIn []string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.nameIn = &nameIn
	return r
}

// The initial index from which to return the results.
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) Offset(offset int32) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;stream&#x60; - Stream* &#x60;-stream&#x60; - Stream (descending)* &#x60;version&#x60; - Version* &#x60;-version&#x60; - Version (descending)* &#x60;context&#x60; - Context* &#x60;-context&#x60; - Context (descending)* &#x60;arch&#x60; - Arch* &#x60;-arch&#x60; - Arch (descending)* &#x60;static_context&#x60; - Static context* &#x60;-static_context&#x60; - Static context (descending)* &#x60;dependencies&#x60; - Dependencies* &#x60;-dependencies&#x60; - Dependencies (descending)* &#x60;artifacts&#x60; - Artifacts* &#x60;-artifacts&#x60; - Artifacts (descending)* &#x60;profiles&#x60; - Profiles* &#x60;-profiles&#x60; - Profiles (descending)* &#x60;description&#x60; - Description* &#x60;-description&#x60; - Description (descending)* &#x60;digest&#x60; - Digest* &#x60;-digest&#x60; - Digest (descending)* &#x60;snippet&#x60; - Snippet* &#x60;-snippet&#x60; - Snippet (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) Ordering(ordering []string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.ordering = &ordering
	return r
}

// Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME.
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) OrphanedFor(orphanedFor float32) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.orphanedFor = &orphanedFor
	return r
}

// Multiple values may be separated by commas.
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) PulpHrefIn(pulpHrefIn []string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) PulpIdIn(pulpIdIn []string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r ContentModulemdsAPIContentRpmModulemdsListRequest) Q(q string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) RepositoryVersion(repositoryVersion string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

func (r ContentModulemdsAPIContentRpmModulemdsListRequest) Sha256(sha256 string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.sha256 = &sha256
	return r
}

// Filter results where stream matches value
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) Stream(stream string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.stream = &stream
	return r
}

// Filter results where stream is in a comma-separated list of values
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) StreamIn(streamIn []string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.streamIn = &streamIn
	return r
}

// Filter results where version matches value
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) Version(version string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.version = &version
	return r
}

// Filter results where version is in a comma-separated list of values
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) VersionIn(versionIn []string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.versionIn = &versionIn
	return r
}

// A list of fields to include in the response.
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) Fields(fields []string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentModulemdsAPIContentRpmModulemdsListRequest) ExcludeFields(excludeFields []string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentModulemdsAPIContentRpmModulemdsListRequest) Execute() (*PaginatedrpmModulemdResponseList, *http.Response, error) {
	return r.ApiService.ContentRpmModulemdsListExecute(r)
}

/*
ContentRpmModulemdsList List modulemds

ViewSet for Modulemd.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentModulemdsAPIContentRpmModulemdsListRequest
*/
func (a *ContentModulemdsAPIService) ContentRpmModulemdsList(ctx context.Context, pulpDomain string) ContentModulemdsAPIContentRpmModulemdsListRequest {
	return ContentModulemdsAPIContentRpmModulemdsListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedrpmModulemdResponseList
func (a *ContentModulemdsAPIService) ContentRpmModulemdsListExecute(r ContentModulemdsAPIContentRpmModulemdsListRequest) (*PaginatedrpmModulemdResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedrpmModulemdResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentModulemdsAPIService.ContentRpmModulemdsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/rpm/modulemds/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.arch != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "arch", r.arch, "form", "")
	}
	if r.archIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "arch__in", r.archIn, "form", "csv")
	}
	if r.context != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "context", r.context, "form", "")
	}
	if r.contextIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "context__in", r.contextIn, "form", "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.nameIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__in", r.nameIn, "form", "csv")
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
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "form", "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "form", "csv")
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
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha256", r.sha256, "form", "")
	}
	if r.stream != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stream", r.stream, "form", "")
	}
	if r.streamIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stream__in", r.streamIn, "form", "csv")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "form", "")
	}
	if r.versionIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version__in", r.versionIn, "form", "csv")
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

type ContentModulemdsAPIContentRpmModulemdsReadRequest struct {
	ctx context.Context
	ApiService *ContentModulemdsAPIService
	rpmModulemdHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentModulemdsAPIContentRpmModulemdsReadRequest) Fields(fields []string) ContentModulemdsAPIContentRpmModulemdsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentModulemdsAPIContentRpmModulemdsReadRequest) ExcludeFields(excludeFields []string) ContentModulemdsAPIContentRpmModulemdsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentModulemdsAPIContentRpmModulemdsReadRequest) Execute() (*RpmModulemdResponse, *http.Response, error) {
	return r.ApiService.ContentRpmModulemdsReadExecute(r)
}

/*
ContentRpmModulemdsRead Inspect a modulemd

ViewSet for Modulemd.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmModulemdHref
 @return ContentModulemdsAPIContentRpmModulemdsReadRequest
*/
func (a *ContentModulemdsAPIService) ContentRpmModulemdsRead(ctx context.Context, rpmModulemdHref string) ContentModulemdsAPIContentRpmModulemdsReadRequest {
	return ContentModulemdsAPIContentRpmModulemdsReadRequest{
		ApiService: a,
		ctx: ctx,
		rpmModulemdHref: rpmModulemdHref,
	}
}

// Execute executes the request
//  @return RpmModulemdResponse
func (a *ContentModulemdsAPIService) ContentRpmModulemdsReadExecute(r ContentModulemdsAPIContentRpmModulemdsReadRequest) (*RpmModulemdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RpmModulemdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentModulemdsAPIService.ContentRpmModulemdsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{rpm_modulemd_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_modulemd_href"+"}", url.PathEscape(parameterValueToString(r.rpmModulemdHref, "rpmModulemdHref")), -1)
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
