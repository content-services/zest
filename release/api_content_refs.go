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


// ContentRefsAPIService ContentRefsAPI service
type ContentRefsAPIService service

type ContentRefsAPIContentOstreeRefsListRequest struct {
	ctx context.Context
	ApiService *ContentRefsAPIService
	pulpDomain string
	checksum *string
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIexact *string
	nameIn *[]string
	nameIregex *string
	nameIstartswith *string
	nameRegex *string
	nameStartswith *string
	offset *int32
	ordering *[]string
	orphanedFor *float32
	prnIn *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	pulpLabelSelect *string
	q *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	fields *[]string
	excludeFields *[]string
}

func (r ContentRefsAPIContentOstreeRefsListRequest) Checksum(checksum string) ContentRefsAPIContentOstreeRefsListRequest {
	r.checksum = &checksum
	return r
}

// Number of results to return per page.
func (r ContentRefsAPIContentOstreeRefsListRequest) Limit(limit int32) ContentRefsAPIContentOstreeRefsListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r ContentRefsAPIContentOstreeRefsListRequest) Name(name string) ContentRefsAPIContentOstreeRefsListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r ContentRefsAPIContentOstreeRefsListRequest) NameContains(nameContains string) ContentRefsAPIContentOstreeRefsListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r ContentRefsAPIContentOstreeRefsListRequest) NameIcontains(nameIcontains string) ContentRefsAPIContentOstreeRefsListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name matches value
func (r ContentRefsAPIContentOstreeRefsListRequest) NameIexact(nameIexact string) ContentRefsAPIContentOstreeRefsListRequest {
	r.nameIexact = &nameIexact
	return r
}

// Filter results where name is in a comma-separated list of values
func (r ContentRefsAPIContentOstreeRefsListRequest) NameIn(nameIn []string) ContentRefsAPIContentOstreeRefsListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name matches regex value
func (r ContentRefsAPIContentOstreeRefsListRequest) NameIregex(nameIregex string) ContentRefsAPIContentOstreeRefsListRequest {
	r.nameIregex = &nameIregex
	return r
}

// Filter results where name starts with value
func (r ContentRefsAPIContentOstreeRefsListRequest) NameIstartswith(nameIstartswith string) ContentRefsAPIContentOstreeRefsListRequest {
	r.nameIstartswith = &nameIstartswith
	return r
}

// Filter results where name matches regex value
func (r ContentRefsAPIContentOstreeRefsListRequest) NameRegex(nameRegex string) ContentRefsAPIContentOstreeRefsListRequest {
	r.nameRegex = &nameRegex
	return r
}

// Filter results where name starts with value
func (r ContentRefsAPIContentOstreeRefsListRequest) NameStartswith(nameStartswith string) ContentRefsAPIContentOstreeRefsListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r ContentRefsAPIContentOstreeRefsListRequest) Offset(offset int32) ContentRefsAPIContentOstreeRefsListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;relative_path&#x60; - Relative path* &#x60;-relative_path&#x60; - Relative path (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r ContentRefsAPIContentOstreeRefsListRequest) Ordering(ordering []string) ContentRefsAPIContentOstreeRefsListRequest {
	r.ordering = &ordering
	return r
}

// Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME.
func (r ContentRefsAPIContentOstreeRefsListRequest) OrphanedFor(orphanedFor float32) ContentRefsAPIContentOstreeRefsListRequest {
	r.orphanedFor = &orphanedFor
	return r
}

// Multiple values may be separated by commas.
func (r ContentRefsAPIContentOstreeRefsListRequest) PrnIn(prnIn []string) ContentRefsAPIContentOstreeRefsListRequest {
	r.prnIn = &prnIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentRefsAPIContentOstreeRefsListRequest) PulpHrefIn(pulpHrefIn []string) ContentRefsAPIContentOstreeRefsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentRefsAPIContentOstreeRefsListRequest) PulpIdIn(pulpIdIn []string) ContentRefsAPIContentOstreeRefsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter labels by search string
func (r ContentRefsAPIContentOstreeRefsListRequest) PulpLabelSelect(pulpLabelSelect string) ContentRefsAPIContentOstreeRefsListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Filter results by using NOT, AND and OR operations on other filters
func (r ContentRefsAPIContentOstreeRefsListRequest) Q(q string) ContentRefsAPIContentOstreeRefsListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentRefsAPIContentOstreeRefsListRequest) RepositoryVersion(repositoryVersion string) ContentRefsAPIContentOstreeRefsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentRefsAPIContentOstreeRefsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentRefsAPIContentOstreeRefsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentRefsAPIContentOstreeRefsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentRefsAPIContentOstreeRefsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentRefsAPIContentOstreeRefsListRequest) Fields(fields []string) ContentRefsAPIContentOstreeRefsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentRefsAPIContentOstreeRefsListRequest) ExcludeFields(excludeFields []string) ContentRefsAPIContentOstreeRefsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentRefsAPIContentOstreeRefsListRequest) Execute() (*PaginatedostreeOstreeRefResponseList, *http.Response, error) {
	return r.ApiService.ContentOstreeRefsListExecute(r)
}

/*
ContentOstreeRefsList List ostree refs

A ViewSet class for OSTree head commits.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentRefsAPIContentOstreeRefsListRequest
*/
func (a *ContentRefsAPIService) ContentOstreeRefsList(ctx context.Context, pulpDomain string) ContentRefsAPIContentOstreeRefsListRequest {
	return ContentRefsAPIContentOstreeRefsListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedostreeOstreeRefResponseList
func (a *ContentRefsAPIService) ContentOstreeRefsListExecute(r ContentRefsAPIContentOstreeRefsListRequest) (*PaginatedostreeOstreeRefResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedostreeOstreeRefResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentRefsAPIService.ContentOstreeRefsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/ostree/refs/"
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
	if r.nameContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__contains", r.nameContains, "form", "")
	}
	if r.nameIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__icontains", r.nameIcontains, "form", "")
	}
	if r.nameIexact != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__iexact", r.nameIexact, "form", "")
	}
	if r.nameIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__in", r.nameIn, "form", "csv")
	}
	if r.nameIregex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__iregex", r.nameIregex, "form", "")
	}
	if r.nameIstartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__istartswith", r.nameIstartswith, "form", "")
	}
	if r.nameRegex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__regex", r.nameRegex, "form", "")
	}
	if r.nameStartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__startswith", r.nameStartswith, "form", "")
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

type ContentRefsAPIContentOstreeRefsReadRequest struct {
	ctx context.Context
	ApiService *ContentRefsAPIService
	ostreeOstreeRefHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentRefsAPIContentOstreeRefsReadRequest) Fields(fields []string) ContentRefsAPIContentOstreeRefsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentRefsAPIContentOstreeRefsReadRequest) ExcludeFields(excludeFields []string) ContentRefsAPIContentOstreeRefsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentRefsAPIContentOstreeRefsReadRequest) Execute() (*OstreeOstreeRefResponse, *http.Response, error) {
	return r.ApiService.ContentOstreeRefsReadExecute(r)
}

/*
ContentOstreeRefsRead Inspect an ostree ref

A ViewSet class for OSTree head commits.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ostreeOstreeRefHref
 @return ContentRefsAPIContentOstreeRefsReadRequest
*/
func (a *ContentRefsAPIService) ContentOstreeRefsRead(ctx context.Context, ostreeOstreeRefHref string) ContentRefsAPIContentOstreeRefsReadRequest {
	return ContentRefsAPIContentOstreeRefsReadRequest{
		ApiService: a,
		ctx: ctx,
		ostreeOstreeRefHref: ostreeOstreeRefHref,
	}
}

// Execute executes the request
//  @return OstreeOstreeRefResponse
func (a *ContentRefsAPIService) ContentOstreeRefsReadExecute(r ContentRefsAPIContentOstreeRefsReadRequest) (*OstreeOstreeRefResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OstreeOstreeRefResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentRefsAPIService.ContentOstreeRefsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ostree_ostree_ref_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ostree_ostree_ref_href"+"}", url.PathEscape(parameterValueToString(r.ostreeOstreeRefHref, "ostreeOstreeRefHref")), -1)
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

type ContentRefsAPIContentOstreeRefsSetLabelRequest struct {
	ctx context.Context
	ApiService *ContentRefsAPIService
	ostreeOstreeRefHref string
	setLabel *SetLabel
}

func (r ContentRefsAPIContentOstreeRefsSetLabelRequest) SetLabel(setLabel SetLabel) ContentRefsAPIContentOstreeRefsSetLabelRequest {
	r.setLabel = &setLabel
	return r
}

func (r ContentRefsAPIContentOstreeRefsSetLabelRequest) Execute() (*SetLabelResponse, *http.Response, error) {
	return r.ApiService.ContentOstreeRefsSetLabelExecute(r)
}

/*
ContentOstreeRefsSetLabel Set a label

Set a single pulp_label on the object to a specific value or null.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ostreeOstreeRefHref
 @return ContentRefsAPIContentOstreeRefsSetLabelRequest
*/
func (a *ContentRefsAPIService) ContentOstreeRefsSetLabel(ctx context.Context, ostreeOstreeRefHref string) ContentRefsAPIContentOstreeRefsSetLabelRequest {
	return ContentRefsAPIContentOstreeRefsSetLabelRequest{
		ApiService: a,
		ctx: ctx,
		ostreeOstreeRefHref: ostreeOstreeRefHref,
	}
}

// Execute executes the request
//  @return SetLabelResponse
func (a *ContentRefsAPIService) ContentOstreeRefsSetLabelExecute(r ContentRefsAPIContentOstreeRefsSetLabelRequest) (*SetLabelResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SetLabelResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentRefsAPIService.ContentOstreeRefsSetLabel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ostree_ostree_ref_href}set_label/"
	localVarPath = strings.Replace(localVarPath, "{"+"ostree_ostree_ref_href"+"}", url.PathEscape(parameterValueToString(r.ostreeOstreeRefHref, "ostreeOstreeRefHref")), -1)
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

type ContentRefsAPIContentOstreeRefsUnsetLabelRequest struct {
	ctx context.Context
	ApiService *ContentRefsAPIService
	ostreeOstreeRefHref string
	unsetLabel *UnsetLabel
}

func (r ContentRefsAPIContentOstreeRefsUnsetLabelRequest) UnsetLabel(unsetLabel UnsetLabel) ContentRefsAPIContentOstreeRefsUnsetLabelRequest {
	r.unsetLabel = &unsetLabel
	return r
}

func (r ContentRefsAPIContentOstreeRefsUnsetLabelRequest) Execute() (*UnsetLabelResponse, *http.Response, error) {
	return r.ApiService.ContentOstreeRefsUnsetLabelExecute(r)
}

/*
ContentOstreeRefsUnsetLabel Unset a label

Unset a single pulp_label on the object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ostreeOstreeRefHref
 @return ContentRefsAPIContentOstreeRefsUnsetLabelRequest
*/
func (a *ContentRefsAPIService) ContentOstreeRefsUnsetLabel(ctx context.Context, ostreeOstreeRefHref string) ContentRefsAPIContentOstreeRefsUnsetLabelRequest {
	return ContentRefsAPIContentOstreeRefsUnsetLabelRequest{
		ApiService: a,
		ctx: ctx,
		ostreeOstreeRefHref: ostreeOstreeRefHref,
	}
}

// Execute executes the request
//  @return UnsetLabelResponse
func (a *ContentRefsAPIService) ContentOstreeRefsUnsetLabelExecute(r ContentRefsAPIContentOstreeRefsUnsetLabelRequest) (*UnsetLabelResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UnsetLabelResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentRefsAPIService.ContentOstreeRefsUnsetLabel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ostree_ostree_ref_href}unset_label/"
	localVarPath = strings.Replace(localVarPath, "{"+"ostree_ostree_ref_href"+"}", url.PathEscape(parameterValueToString(r.ostreeOstreeRefHref, "ostreeOstreeRefHref")), -1)
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
