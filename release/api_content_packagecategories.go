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


// ContentPackagecategoriesAPIService ContentPackagecategoriesAPI service
type ContentPackagecategoriesAPIService service

type ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest struct {
	ctx context.Context
	ApiService *ContentPackagecategoriesAPIService
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
func (r ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest) Limit(limit int32) ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest) Offset(offset int32) ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest) Ordering(ordering []string) ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest {
	r.ordering = &ordering
	return r
}

// Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME.
func (r ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest) OrphanedFor(orphanedFor float32) ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest {
	r.orphanedFor = &orphanedFor
	return r
}

// Multiple values may be separated by commas.
func (r ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest) PulpHrefIn(pulpHrefIn []string) ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest) PulpIdIn(pulpIdIn []string) ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest) Q(q string) ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest) RepositoryVersion(repositoryVersion string) ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest) Fields(fields []string) ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest) ExcludeFields(excludeFields []string) ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest) Execute() (*PaginatedrpmPackageCategoryResponseList, *http.Response, error) {
	return r.ApiService.ContentRpmPackagecategoriesListExecute(r)
}

/*
ContentRpmPackagecategoriesList List package categorys

PackageCategory ViewSet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest
*/
func (a *ContentPackagecategoriesAPIService) ContentRpmPackagecategoriesList(ctx context.Context, pulpDomain string) ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest {
	return ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedrpmPackageCategoryResponseList
func (a *ContentPackagecategoriesAPIService) ContentRpmPackagecategoriesListExecute(r ContentPackagecategoriesAPIContentRpmPackagecategoriesListRequest) (*PaginatedrpmPackageCategoryResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedrpmPackageCategoryResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackagecategoriesAPIService.ContentRpmPackagecategoriesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/rpm/packagecategories/"
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

type ContentPackagecategoriesAPIContentRpmPackagecategoriesReadRequest struct {
	ctx context.Context
	ApiService *ContentPackagecategoriesAPIService
	rpmPackageCategoryHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentPackagecategoriesAPIContentRpmPackagecategoriesReadRequest) Fields(fields []string) ContentPackagecategoriesAPIContentRpmPackagecategoriesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackagecategoriesAPIContentRpmPackagecategoriesReadRequest) ExcludeFields(excludeFields []string) ContentPackagecategoriesAPIContentRpmPackagecategoriesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackagecategoriesAPIContentRpmPackagecategoriesReadRequest) Execute() (*RpmPackageCategoryResponse, *http.Response, error) {
	return r.ApiService.ContentRpmPackagecategoriesReadExecute(r)
}

/*
ContentRpmPackagecategoriesRead Inspect a package category

PackageCategory ViewSet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmPackageCategoryHref
 @return ContentPackagecategoriesAPIContentRpmPackagecategoriesReadRequest
*/
func (a *ContentPackagecategoriesAPIService) ContentRpmPackagecategoriesRead(ctx context.Context, rpmPackageCategoryHref string) ContentPackagecategoriesAPIContentRpmPackagecategoriesReadRequest {
	return ContentPackagecategoriesAPIContentRpmPackagecategoriesReadRequest{
		ApiService: a,
		ctx: ctx,
		rpmPackageCategoryHref: rpmPackageCategoryHref,
	}
}

// Execute executes the request
//  @return RpmPackageCategoryResponse
func (a *ContentPackagecategoriesAPIService) ContentRpmPackagecategoriesReadExecute(r ContentPackagecategoriesAPIContentRpmPackagecategoriesReadRequest) (*RpmPackageCategoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RpmPackageCategoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackagecategoriesAPIService.ContentRpmPackagecategoriesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{rpm_package_category_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_package_category_href"+"}", url.PathEscape(parameterValueToString(r.rpmPackageCategoryHref, "rpmPackageCategoryHref")), -1)
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
