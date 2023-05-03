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


// PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiService PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApi service
type PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiService service

type PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiService
	path string
	dependency *string
	deprecated *bool
	distribution *[]string
	distributionBasePath *[]string
	highest *bool
	isDeprecated *bool
	isHighest *bool
	isSigned *bool
	keywords *string
	limit *int32
	name *string
	namespace *string
	offset *int32
	orderBy *[]string
	q *string
	repository *[]string
	repositoryLabel *string
	repositoryName *[]string
	repositoryVersion *string
	signed *bool
	tags *string
	version *string
	versionRange *string
	fields *[]string
	excludeFields *[]string
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) Dependency(dependency string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.dependency = &dependency
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) Deprecated(deprecated bool) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.deprecated = &deprecated
	return r
}

// Filter collectionversions that are in these distrubtion ids.
func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) Distribution(distribution []string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.distribution = &distribution
	return r
}

// Filter collectionversions that are in these base paths.
func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) DistributionBasePath(distributionBasePath []string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.distributionBasePath = &distributionBasePath
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) Highest(highest bool) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.highest = &highest
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) IsDeprecated(isDeprecated bool) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.isDeprecated = &isDeprecated
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) IsHighest(isHighest bool) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.isHighest = &isHighest
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) IsSigned(isSigned bool) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.isSigned = &isSigned
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) Keywords(keywords string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.keywords = &keywords
	return r
}

// Number of results to return per page.
func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) Limit(limit int32) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.limit = &limit
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) Name(name string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.name = &name
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) Namespace(namespace string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.namespace = &namespace
	return r
}

// The initial index from which to return the results.
func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) Offset(offset int32) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_created&#x60; - by CV created * &#x60;-pulp_created&#x60; - by CV created (descending) * &#x60;namespace&#x60; - by CV namespace * &#x60;-namespace&#x60; - by CV namespace (descending) * &#x60;name&#x60; - by CV name * &#x60;-name&#x60; - by CV name (descending) * &#x60;version&#x60; - by CV version * &#x60;-version&#x60; - by CV version (descending)
func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) OrderBy(orderBy []string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.orderBy = &orderBy
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) Q(q string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.q = &q
	return r
}

// Filter collectionversions that are in these repository ids.
func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) Repository(repository []string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.repository = &repository
	return r
}

// Filter labels by search string
func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) RepositoryLabel(repositoryLabel string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.repositoryLabel = &repositoryLabel
	return r
}

// Filter collectionversions that are in these repositories.
func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) RepositoryName(repositoryName []string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.repositoryName = &repositoryName
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) RepositoryVersion(repositoryVersion string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) Signed(signed bool) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.signed = &signed
	return r
}

// Filter by comma separate list of tags that must all be matched
func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) Tags(tags string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.tags = &tags
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) Version(version string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.version = &version
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) VersionRange(versionRange string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.versionRange = &versionRange
	return r
}

// A list of fields to include in the response.
func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) Fields(fields []string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) ExcludeFields(excludeFields []string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) Execute() (*PaginatedCollectionVersionSearchListResponseList, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListExecute(r)
}

/*
PulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsList Method for PulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsList

A viewset for cross-repo searches.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param path
 @return PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest
*/
func (a *PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiService) PulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsList(ctx context.Context, path string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest {
	return PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest{
		ApiService: a,
		ctx: ctx,
		path: path,
	}
}

// Execute executes the request
//  @return PaginatedCollectionVersionSearchListResponseList
func (a *PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiService) PulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListExecute(r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsListRequest) (*PaginatedCollectionVersionSearchListResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedCollectionVersionSearchListResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiService.PulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/search/collection-versions/"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterValueToString(r.path, "path")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.dependency != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dependency", r.dependency, "")
	}
	if r.deprecated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deprecated", r.deprecated, "")
	}
	if r.distribution != nil {
		t := *r.distribution
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "distribution", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "distribution", t, "multi")
		}
	}
	if r.distributionBasePath != nil {
		t := *r.distributionBasePath
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "distribution_base_path", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "distribution_base_path", t, "multi")
		}
	}
	if r.highest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "highest", r.highest, "")
	}
	if r.isDeprecated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_deprecated", r.isDeprecated, "")
	}
	if r.isHighest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_highest", r.isHighest, "")
	}
	if r.isSigned != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_signed", r.isSigned, "")
	}
	if r.keywords != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "keywords", r.keywords, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.namespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "csv")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
	}
	if r.repository != nil {
		t := *r.repository
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "repository", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "repository", t, "multi")
		}
	}
	if r.repositoryLabel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_label", r.repositoryLabel, "")
	}
	if r.repositoryName != nil {
		t := *r.repositoryName
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "repository_name", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "repository_name", t, "multi")
		}
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.signed != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "signed", r.signed, "")
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tags", r.tags, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
	if r.versionRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version_range", r.versionRange, "")
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

type PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsRebuildRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiService
	path string
	collectionVersionSearchList *CollectionVersionSearchList
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsRebuildRequest) CollectionVersionSearchList(collectionVersionSearchList CollectionVersionSearchList) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsRebuildRequest {
	r.collectionVersionSearchList = &collectionVersionSearchList
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsRebuildRequest) Execute() (*CollectionVersionSearchListResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsRebuildExecute(r)
}

/*
PulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsRebuild Method for PulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsRebuild

A viewset for cross-repo searches.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param path
 @return PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsRebuildRequest
*/
func (a *PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiService) PulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsRebuild(ctx context.Context, path string) PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsRebuildRequest {
	return PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsRebuildRequest{
		ApiService: a,
		ctx: ctx,
		path: path,
	}
}

// Execute executes the request
//  @return CollectionVersionSearchListResponse
func (a *PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiService) PulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsRebuildExecute(r PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiPulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsRebuildRequest) (*CollectionVersionSearchListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionVersionSearchListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3PluginAnsibleSearchCollectionVersionsApiService.PulpAnsibleGalaxyApiV3PluginAnsibleSearchCollectionVersionsRebuild")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/search/collection-versions/"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterValueToString(r.path, "path")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.collectionVersionSearchList == nil {
		return localVarReturnValue, nil, reportError("collectionVersionSearchList is required and must be specified")
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
	localVarPostBody = r.collectionVersionSearchList
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