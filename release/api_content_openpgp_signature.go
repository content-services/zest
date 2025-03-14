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


// ContentOpenpgpSignatureAPIService ContentOpenpgpSignatureAPI service
type ContentOpenpgpSignatureAPIService service

type ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest struct {
	ctx context.Context
	ApiService *ContentOpenpgpSignatureAPIService
	pulpDomain string
	issuer *string
	limit *int32
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

// Filter results where issuer matches value
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) Issuer(issuer string) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest {
	r.issuer = &issuer
	return r
}

// Number of results to return per page.
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) Limit(limit int32) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) Offset(offset int32) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;raw_data&#x60; - Raw data* &#x60;-raw_data&#x60; - Raw data (descending)* &#x60;sha256&#x60; - Sha256* &#x60;-sha256&#x60; - Sha256 (descending)* &#x60;signature_type&#x60; - Signature type* &#x60;-signature_type&#x60; - Signature type (descending)* &#x60;created&#x60; - Created* &#x60;-created&#x60; - Created (descending)* &#x60;expiration_time&#x60; - Expiration time* &#x60;-expiration_time&#x60; - Expiration time (descending)* &#x60;key_expiration_time&#x60; - Key expiration time* &#x60;-key_expiration_time&#x60; - Key expiration time (descending)* &#x60;issuer&#x60; - Issuer* &#x60;-issuer&#x60; - Issuer (descending)* &#x60;signers_user_id&#x60; - Signers user id* &#x60;-signers_user_id&#x60; - Signers user id (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) Ordering(ordering []string) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest {
	r.ordering = &ordering
	return r
}

// Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME.
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) OrphanedFor(orphanedFor float32) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest {
	r.orphanedFor = &orphanedFor
	return r
}

// Multiple values may be separated by commas.
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) PrnIn(prnIn []string) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest {
	r.prnIn = &prnIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) PulpHrefIn(pulpHrefIn []string) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) PulpIdIn(pulpIdIn []string) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter labels by search string
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) PulpLabelSelect(pulpLabelSelect string) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Filter results by using NOT, AND and OR operations on other filters
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) Q(q string) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) RepositoryVersion(repositoryVersion string) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) Fields(fields []string) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) ExcludeFields(excludeFields []string) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) Execute() (*PaginatedOpenPGPSignatureResponseList, *http.Response, error) {
	return r.ApiService.ContentCoreOpenpgpSignatureListExecute(r)
}

/*
ContentCoreOpenpgpSignatureList List open pgp signatures

Content viewset that supports only GET by default.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest
*/
func (a *ContentOpenpgpSignatureAPIService) ContentCoreOpenpgpSignatureList(ctx context.Context, pulpDomain string) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest {
	return ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedOpenPGPSignatureResponseList
func (a *ContentOpenpgpSignatureAPIService) ContentCoreOpenpgpSignatureListExecute(r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureListRequest) (*PaginatedOpenPGPSignatureResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedOpenPGPSignatureResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentOpenpgpSignatureAPIService.ContentCoreOpenpgpSignatureList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/core/openpgp_signature/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.issuer != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "issuer", r.issuer, "form", "")
	}
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

type ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureReadRequest struct {
	ctx context.Context
	ApiService *ContentOpenpgpSignatureAPIService
	openPGPSignatureHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureReadRequest) Fields(fields []string) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureReadRequest) ExcludeFields(excludeFields []string) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureReadRequest) Execute() (*OpenPGPSignatureResponse, *http.Response, error) {
	return r.ApiService.ContentCoreOpenpgpSignatureReadExecute(r)
}

/*
ContentCoreOpenpgpSignatureRead Inspect an open pgp signature

Content viewset that supports only GET by default.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param openPGPSignatureHref
 @return ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureReadRequest
*/
func (a *ContentOpenpgpSignatureAPIService) ContentCoreOpenpgpSignatureRead(ctx context.Context, openPGPSignatureHref string) ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureReadRequest {
	return ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureReadRequest{
		ApiService: a,
		ctx: ctx,
		openPGPSignatureHref: openPGPSignatureHref,
	}
}

// Execute executes the request
//  @return OpenPGPSignatureResponse
func (a *ContentOpenpgpSignatureAPIService) ContentCoreOpenpgpSignatureReadExecute(r ContentOpenpgpSignatureAPIContentCoreOpenpgpSignatureReadRequest) (*OpenPGPSignatureResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OpenPGPSignatureResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentOpenpgpSignatureAPIService.ContentCoreOpenpgpSignatureRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{open_p_g_p_signature_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"open_p_g_p_signature_href"+"}", url.PathEscape(parameterValueToString(r.openPGPSignatureHref, "openPGPSignatureHref")), -1)
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
