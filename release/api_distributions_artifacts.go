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


// DistributionsArtifactsAPIService DistributionsArtifactsAPI service
type DistributionsArtifactsAPIService service

type DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest struct {
	ctx context.Context
	ApiService *DistributionsArtifactsAPIService
	pulpDomain string
	basePath *string
	basePathContains *string
	basePathIcontains *string
	basePathIn *[]string
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	pulpLabelSelect *string
	repository *string
	repositoryIn *[]string
	withContent *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where base_path matches value
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) BasePath(basePath string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.basePath = &basePath
	return r
}

// Filter results where base_path contains value
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) BasePathContains(basePathContains string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.basePathContains = &basePathContains
	return r
}

// Filter results where base_path contains value
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) BasePathIcontains(basePathIcontains string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.basePathIcontains = &basePathIcontains
	return r
}

// Filter results where base_path is in a comma-separated list of values
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) BasePathIn(basePathIn []string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.basePathIn = &basePathIn
	return r
}

// Number of results to return per page.
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) Limit(limit int32) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) Name(name string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) NameContains(nameContains string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) NameIcontains(nameIcontains string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) NameIn(nameIn []string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) NameStartswith(nameStartswith string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) Offset(offset int32) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;pulp_labels&#x60; - Pulp labels * &#x60;-pulp_labels&#x60; - Pulp labels (descending) * &#x60;base_path&#x60; - Base path * &#x60;-base_path&#x60; - Base path (descending) * &#x60;hidden&#x60; - Hidden * &#x60;-hidden&#x60; - Hidden (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) Ordering(ordering []string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) PulpHrefIn(pulpHrefIn []string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) PulpIdIn(pulpIdIn []string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter labels by search string
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) PulpLabelSelect(pulpLabelSelect string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Filter results where repository matches value
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) Repository(repository string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.repository = &repository
	return r
}

// Filter results where repository is in a comma-separated list of values
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) RepositoryIn(repositoryIn []string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.repositoryIn = &repositoryIn
	return r
}

// Filter distributions based on the content served by them
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) WithContent(withContent string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.withContent = &withContent
	return r
}

// A list of fields to include in the response.
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) Fields(fields []string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) ExcludeFields(excludeFields []string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) Execute() (*PaginatedArtifactDistributionResponseList, *http.Response, error) {
	return r.ApiService.DistributionsCoreArtifactsListExecute(r)
}

/*
DistributionsCoreArtifactsList List artifact distributions

ViewSet for ArtifactDistribution.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest
*/
func (a *DistributionsArtifactsAPIService) DistributionsCoreArtifactsList(ctx context.Context, pulpDomain string) DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest {
	return DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedArtifactDistributionResponseList
func (a *DistributionsArtifactsAPIService) DistributionsCoreArtifactsListExecute(r DistributionsArtifactsAPIDistributionsCoreArtifactsListRequest) (*PaginatedArtifactDistributionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedArtifactDistributionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsArtifactsAPIService.DistributionsCoreArtifactsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/{pulp_domain}/api/v3/distributions/core/artifacts/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.basePath != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path", r.basePath, "")
	}
	if r.basePathContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path__contains", r.basePathContains, "")
	}
	if r.basePathIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path__icontains", r.basePathIcontains, "")
	}
	if r.basePathIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path__in", r.basePathIn, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.nameContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__contains", r.nameContains, "")
	}
	if r.nameIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__icontains", r.nameIcontains, "")
	}
	if r.nameIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__in", r.nameIn, "csv")
	}
	if r.nameStartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__startswith", r.nameStartswith, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.pulpLabelSelect != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_label_select", r.pulpLabelSelect, "")
	}
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository", r.repository, "")
	}
	if r.repositoryIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository__in", r.repositoryIn, "csv")
	}
	if r.withContent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "with_content", r.withContent, "")
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

type DistributionsArtifactsAPIDistributionsCoreArtifactsReadRequest struct {
	ctx context.Context
	ApiService *DistributionsArtifactsAPIService
	artifactDistributionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsReadRequest) Fields(fields []string) DistributionsArtifactsAPIDistributionsCoreArtifactsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r DistributionsArtifactsAPIDistributionsCoreArtifactsReadRequest) ExcludeFields(excludeFields []string) DistributionsArtifactsAPIDistributionsCoreArtifactsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r DistributionsArtifactsAPIDistributionsCoreArtifactsReadRequest) Execute() (*ArtifactDistributionResponse, *http.Response, error) {
	return r.ApiService.DistributionsCoreArtifactsReadExecute(r)
}

/*
DistributionsCoreArtifactsRead Inspect an artifact distribution

ViewSet for ArtifactDistribution.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param artifactDistributionHref
 @return DistributionsArtifactsAPIDistributionsCoreArtifactsReadRequest
*/
func (a *DistributionsArtifactsAPIService) DistributionsCoreArtifactsRead(ctx context.Context, artifactDistributionHref string) DistributionsArtifactsAPIDistributionsCoreArtifactsReadRequest {
	return DistributionsArtifactsAPIDistributionsCoreArtifactsReadRequest{
		ApiService: a,
		ctx: ctx,
		artifactDistributionHref: artifactDistributionHref,
	}
}

// Execute executes the request
//  @return ArtifactDistributionResponse
func (a *DistributionsArtifactsAPIService) DistributionsCoreArtifactsReadExecute(r DistributionsArtifactsAPIDistributionsCoreArtifactsReadRequest) (*ArtifactDistributionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactDistributionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsArtifactsAPIService.DistributionsCoreArtifactsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{artifact_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"artifact_distribution_href"+"}", url.PathEscape(parameterValueToString(r.artifactDistributionHref, "artifactDistributionHref")), -1)
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
