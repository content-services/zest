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
	"reflect"
)


// DistributionsArtifactsApiService DistributionsArtifactsApi service
type DistributionsArtifactsApiService service

type DistributionsArtifactsApiDistributionsCoreArtifactsListRequest struct {
	ctx context.Context
	ApiService *DistributionsArtifactsApiService
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
	pulpLabelSelect *string
	repository *string
	repositoryIn *[]string
	withContent *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where base_path matches value
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) BasePath(basePath string) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.basePath = &basePath
	return r
}

// Filter results where base_path contains value
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) BasePathContains(basePathContains string) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.basePathContains = &basePathContains
	return r
}

// Filter results where base_path contains value
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) BasePathIcontains(basePathIcontains string) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.basePathIcontains = &basePathIcontains
	return r
}

// Filter results where base_path is in a comma-separated list of values
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) BasePathIn(basePathIn []string) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.basePathIn = &basePathIn
	return r
}

// Number of results to return per page.
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) Limit(limit int32) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) Name(name string) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) NameContains(nameContains string) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) NameIcontains(nameIcontains string) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) NameIn(nameIn []string) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) NameStartswith(nameStartswith string) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) Offset(offset int32) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) Ordering(ordering []string) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.ordering = &ordering
	return r
}

// Filter labels by search string
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) PulpLabelSelect(pulpLabelSelect string) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Filter results where repository matches value
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) Repository(repository string) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.repository = &repository
	return r
}

// Filter results where repository is in a comma-separated list of values
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) RepositoryIn(repositoryIn []string) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.repositoryIn = &repositoryIn
	return r
}

// Filter distributions based on the content served by them
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) WithContent(withContent string) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.withContent = &withContent
	return r
}

// A list of fields to include in the response.
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) Fields(fields []string) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) ExcludeFields(excludeFields []string) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) Execute() (*PaginatedArtifactDistributionResponseList, *http.Response, error) {
	return r.ApiService.DistributionsCoreArtifactsListExecute(r)
}

/*
DistributionsCoreArtifactsList List artifact distributions

ViewSet for ArtifactDistribution.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DistributionsArtifactsApiDistributionsCoreArtifactsListRequest
*/
func (a *DistributionsArtifactsApiService) DistributionsCoreArtifactsList(ctx context.Context) DistributionsArtifactsApiDistributionsCoreArtifactsListRequest {
	return DistributionsArtifactsApiDistributionsCoreArtifactsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedArtifactDistributionResponseList
func (a *DistributionsArtifactsApiService) DistributionsCoreArtifactsListExecute(r DistributionsArtifactsApiDistributionsCoreArtifactsListRequest) (*PaginatedArtifactDistributionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedArtifactDistributionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsArtifactsApiService.DistributionsCoreArtifactsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/distributions/core/artifacts/"
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

type DistributionsArtifactsApiDistributionsCoreArtifactsReadRequest struct {
	ctx context.Context
	ApiService *DistributionsArtifactsApiService
	artifactDistributionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r DistributionsArtifactsApiDistributionsCoreArtifactsReadRequest) Fields(fields []string) DistributionsArtifactsApiDistributionsCoreArtifactsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r DistributionsArtifactsApiDistributionsCoreArtifactsReadRequest) ExcludeFields(excludeFields []string) DistributionsArtifactsApiDistributionsCoreArtifactsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r DistributionsArtifactsApiDistributionsCoreArtifactsReadRequest) Execute() (*ArtifactDistributionResponse, *http.Response, error) {
	return r.ApiService.DistributionsCoreArtifactsReadExecute(r)
}

/*
DistributionsCoreArtifactsRead Inspect an artifact distribution

ViewSet for ArtifactDistribution.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param artifactDistributionHref
 @return DistributionsArtifactsApiDistributionsCoreArtifactsReadRequest
*/
func (a *DistributionsArtifactsApiService) DistributionsCoreArtifactsRead(ctx context.Context, artifactDistributionHref string) DistributionsArtifactsApiDistributionsCoreArtifactsReadRequest {
	return DistributionsArtifactsApiDistributionsCoreArtifactsReadRequest{
		ApiService: a,
		ctx: ctx,
		artifactDistributionHref: artifactDistributionHref,
	}
}

// Execute executes the request
//  @return ArtifactDistributionResponse
func (a *DistributionsArtifactsApiService) DistributionsCoreArtifactsReadExecute(r DistributionsArtifactsApiDistributionsCoreArtifactsReadRequest) (*ArtifactDistributionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactDistributionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsArtifactsApiService.DistributionsCoreArtifactsRead")
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
