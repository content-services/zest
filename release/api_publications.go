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
	"time"
	"reflect"
)


// PublicationsAPIService PublicationsAPI service
type PublicationsAPIService service

type PublicationsAPIPublicationsListRequest struct {
	ctx context.Context
	ApiService *PublicationsAPIService
	pulpDomain string
	content *string
	contentIn *string
	limit *int32
	offset *int32
	ordering *[]string
	pulpCreated *time.Time
	pulpCreatedGt *time.Time
	pulpCreatedGte *time.Time
	pulpCreatedLt *time.Time
	pulpCreatedLte *time.Time
	pulpCreatedRange *[]time.Time
	pulpHrefIn *[]string
	pulpIdIn *[]string
	pulpTypeIn *[]string
	repository *string
	repositoryVersion *string
	fields *[]string
	excludeFields *[]string
}

// Content Unit referenced by HREF
func (r PublicationsAPIPublicationsListRequest) Content(content string) PublicationsAPIPublicationsListRequest {
	r.content = &content
	return r
}

// Content Unit referenced by HREF
func (r PublicationsAPIPublicationsListRequest) ContentIn(contentIn string) PublicationsAPIPublicationsListRequest {
	r.contentIn = &contentIn
	return r
}

// Number of results to return per page.
func (r PublicationsAPIPublicationsListRequest) Limit(limit int32) PublicationsAPIPublicationsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r PublicationsAPIPublicationsListRequest) Offset(offset int32) PublicationsAPIPublicationsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;complete&#x60; - Complete * &#x60;-complete&#x60; - Complete (descending) * &#x60;pass_through&#x60; - Pass through * &#x60;-pass_through&#x60; - Pass through (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r PublicationsAPIPublicationsListRequest) Ordering(ordering []string) PublicationsAPIPublicationsListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where pulp_created matches value
func (r PublicationsAPIPublicationsListRequest) PulpCreated(pulpCreated time.Time) PublicationsAPIPublicationsListRequest {
	r.pulpCreated = &pulpCreated
	return r
}

// Filter results where pulp_created is greater than value
func (r PublicationsAPIPublicationsListRequest) PulpCreatedGt(pulpCreatedGt time.Time) PublicationsAPIPublicationsListRequest {
	r.pulpCreatedGt = &pulpCreatedGt
	return r
}

// Filter results where pulp_created is greater than or equal to value
func (r PublicationsAPIPublicationsListRequest) PulpCreatedGte(pulpCreatedGte time.Time) PublicationsAPIPublicationsListRequest {
	r.pulpCreatedGte = &pulpCreatedGte
	return r
}

// Filter results where pulp_created is less than value
func (r PublicationsAPIPublicationsListRequest) PulpCreatedLt(pulpCreatedLt time.Time) PublicationsAPIPublicationsListRequest {
	r.pulpCreatedLt = &pulpCreatedLt
	return r
}

// Filter results where pulp_created is less than or equal to value
func (r PublicationsAPIPublicationsListRequest) PulpCreatedLte(pulpCreatedLte time.Time) PublicationsAPIPublicationsListRequest {
	r.pulpCreatedLte = &pulpCreatedLte
	return r
}

// Filter results where pulp_created is between two comma separated values
func (r PublicationsAPIPublicationsListRequest) PulpCreatedRange(pulpCreatedRange []time.Time) PublicationsAPIPublicationsListRequest {
	r.pulpCreatedRange = &pulpCreatedRange
	return r
}

// Multiple values may be separated by commas.
func (r PublicationsAPIPublicationsListRequest) PulpHrefIn(pulpHrefIn []string) PublicationsAPIPublicationsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r PublicationsAPIPublicationsListRequest) PulpIdIn(pulpIdIn []string) PublicationsAPIPublicationsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Pulp type is in  * &#x60;rpm.rpm&#x60; - rpm.rpm
func (r PublicationsAPIPublicationsListRequest) PulpTypeIn(pulpTypeIn []string) PublicationsAPIPublicationsListRequest {
	r.pulpTypeIn = &pulpTypeIn
	return r
}

// Repository referenced by HREF
func (r PublicationsAPIPublicationsListRequest) Repository(repository string) PublicationsAPIPublicationsListRequest {
	r.repository = &repository
	return r
}

// Repository Version referenced by HREF
func (r PublicationsAPIPublicationsListRequest) RepositoryVersion(repositoryVersion string) PublicationsAPIPublicationsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// A list of fields to include in the response.
func (r PublicationsAPIPublicationsListRequest) Fields(fields []string) PublicationsAPIPublicationsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PublicationsAPIPublicationsListRequest) ExcludeFields(excludeFields []string) PublicationsAPIPublicationsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PublicationsAPIPublicationsListRequest) Execute() (*PaginatedPublicationResponseList, *http.Response, error) {
	return r.ApiService.PublicationsListExecute(r)
}

/*
PublicationsList List publications

A base class for any publication viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return PublicationsAPIPublicationsListRequest
*/
func (a *PublicationsAPIService) PublicationsList(ctx context.Context, pulpDomain string) PublicationsAPIPublicationsListRequest {
	return PublicationsAPIPublicationsListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedPublicationResponseList
func (a *PublicationsAPIService) PublicationsListExecute(r PublicationsAPIPublicationsListRequest) (*PaginatedPublicationResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedPublicationResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsAPIService.PublicationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/{pulp_domain}/api/v3/publications/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.content != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content", r.content, "")
	}
	if r.contentIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content__in", r.contentIn, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.pulpCreated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created", r.pulpCreated, "")
	}
	if r.pulpCreatedGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__gt", r.pulpCreatedGt, "")
	}
	if r.pulpCreatedGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__gte", r.pulpCreatedGte, "")
	}
	if r.pulpCreatedLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__lt", r.pulpCreatedLt, "")
	}
	if r.pulpCreatedLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__lte", r.pulpCreatedLte, "")
	}
	if r.pulpCreatedRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__range", r.pulpCreatedRange, "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.pulpTypeIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_type__in", r.pulpTypeIn, "csv")
	}
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository", r.repository, "")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
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
