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


// ContentPackagelangpacksAPIService ContentPackagelangpacksAPI service
type ContentPackagelangpacksAPIService service

type ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest struct {
	ctx context.Context
	ApiService *ContentPackagelangpacksAPIService
	pulpDomain string
	limit *int32
	offset *int32
	ordering *[]string
	orphanedFor *float32
	pulpHrefIn *[]string
	pulpIdIn *[]string
	q *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest) Limit(limit int32) ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest) Offset(offset int32) ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest) Ordering(ordering []string) ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest {
	r.ordering = &ordering
	return r
}

// Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME.
func (r ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest) OrphanedFor(orphanedFor float32) ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest {
	r.orphanedFor = &orphanedFor
	return r
}

// Multiple values may be separated by commas.
func (r ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest) PulpHrefIn(pulpHrefIn []string) ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest) PulpIdIn(pulpIdIn []string) ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest) Q(q string) ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest) RepositoryVersion(repositoryVersion string) ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest) Fields(fields []string) ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest) ExcludeFields(excludeFields []string) ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest) Execute() (*PaginatedrpmPackageLangpacksResponseList, *http.Response, error) {
	return r.ApiService.ContentRpmPackagelangpacksListExecute(r)
}

/*
ContentRpmPackagelangpacksList List package langpackss

PackageLangpacks ViewSet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest
*/
func (a *ContentPackagelangpacksAPIService) ContentRpmPackagelangpacksList(ctx context.Context, pulpDomain string) ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest {
	return ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedrpmPackageLangpacksResponseList
func (a *ContentPackagelangpacksAPIService) ContentRpmPackagelangpacksListExecute(r ContentPackagelangpacksAPIContentRpmPackagelangpacksListRequest) (*PaginatedrpmPackageLangpacksResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedrpmPackageLangpacksResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackagelangpacksAPIService.ContentRpmPackagelangpacksList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/rpm/packagelangpacks/"
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

type ContentPackagelangpacksAPIContentRpmPackagelangpacksReadRequest struct {
	ctx context.Context
	ApiService *ContentPackagelangpacksAPIService
	rpmPackageLangpacksHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentPackagelangpacksAPIContentRpmPackagelangpacksReadRequest) Fields(fields []string) ContentPackagelangpacksAPIContentRpmPackagelangpacksReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackagelangpacksAPIContentRpmPackagelangpacksReadRequest) ExcludeFields(excludeFields []string) ContentPackagelangpacksAPIContentRpmPackagelangpacksReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackagelangpacksAPIContentRpmPackagelangpacksReadRequest) Execute() (*RpmPackageLangpacksResponse, *http.Response, error) {
	return r.ApiService.ContentRpmPackagelangpacksReadExecute(r)
}

/*
ContentRpmPackagelangpacksRead Inspect a package langpacks

PackageLangpacks ViewSet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmPackageLangpacksHref
 @return ContentPackagelangpacksAPIContentRpmPackagelangpacksReadRequest
*/
func (a *ContentPackagelangpacksAPIService) ContentRpmPackagelangpacksRead(ctx context.Context, rpmPackageLangpacksHref string) ContentPackagelangpacksAPIContentRpmPackagelangpacksReadRequest {
	return ContentPackagelangpacksAPIContentRpmPackagelangpacksReadRequest{
		ApiService: a,
		ctx: ctx,
		rpmPackageLangpacksHref: rpmPackageLangpacksHref,
	}
}

// Execute executes the request
//  @return RpmPackageLangpacksResponse
func (a *ContentPackagelangpacksAPIService) ContentRpmPackagelangpacksReadExecute(r ContentPackagelangpacksAPIContentRpmPackagelangpacksReadRequest) (*RpmPackageLangpacksResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RpmPackageLangpacksResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackagelangpacksAPIService.ContentRpmPackagelangpacksRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{rpm_package_langpacks_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_package_langpacks_href"+"}", url.PathEscape(parameterValueToString(r.rpmPackageLangpacksHref, "rpmPackageLangpacksHref")), -1)
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
