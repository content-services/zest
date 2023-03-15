/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
)


// RepositoriesApiService RepositoriesApi service
type RepositoriesApiService service

type RepositoriesApiRepositoriesListRequest struct {
	ctx context.Context
	ApiService *RepositoriesApiService
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
	offset *int32
	ordering *[]string
	pulpLabelSelect *string
	remote *string
	retainRepoVersions *int32
	retainRepoVersionsGt *int32
	retainRepoVersionsGte *int32
	retainRepoVersionsIsnull *bool
	retainRepoVersionsLt *int32
	retainRepoVersionsLte *int32
	retainRepoVersionsNe *int32
	retainRepoVersionsRange *[]int32
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r RepositoriesApiRepositoriesListRequest) Limit(limit int32) RepositoriesApiRepositoriesListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r RepositoriesApiRepositoriesListRequest) Name(name string) RepositoriesApiRepositoriesListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r RepositoriesApiRepositoriesListRequest) NameContains(nameContains string) RepositoriesApiRepositoriesListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r RepositoriesApiRepositoriesListRequest) NameIcontains(nameIcontains string) RepositoriesApiRepositoriesListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r RepositoriesApiRepositoriesListRequest) NameIn(nameIn []string) RepositoriesApiRepositoriesListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r RepositoriesApiRepositoriesListRequest) NameStartswith(nameStartswith string) RepositoriesApiRepositoriesListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r RepositoriesApiRepositoriesListRequest) Offset(offset int32) RepositoriesApiRepositoriesListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r RepositoriesApiRepositoriesListRequest) Ordering(ordering []string) RepositoriesApiRepositoriesListRequest {
	r.ordering = &ordering
	return r
}

// Filter labels by search string
func (r RepositoriesApiRepositoriesListRequest) PulpLabelSelect(pulpLabelSelect string) RepositoriesApiRepositoriesListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Foreign Key referenced by HREF
func (r RepositoriesApiRepositoriesListRequest) Remote(remote string) RepositoriesApiRepositoriesListRequest {
	r.remote = &remote
	return r
}

// Filter results where retain_repo_versions matches value
func (r RepositoriesApiRepositoriesListRequest) RetainRepoVersions(retainRepoVersions int32) RepositoriesApiRepositoriesListRequest {
	r.retainRepoVersions = &retainRepoVersions
	return r
}

// Filter results where retain_repo_versions is greater than value
func (r RepositoriesApiRepositoriesListRequest) RetainRepoVersionsGt(retainRepoVersionsGt int32) RepositoriesApiRepositoriesListRequest {
	r.retainRepoVersionsGt = &retainRepoVersionsGt
	return r
}

// Filter results where retain_repo_versions is greater than or equal to value
func (r RepositoriesApiRepositoriesListRequest) RetainRepoVersionsGte(retainRepoVersionsGte int32) RepositoriesApiRepositoriesListRequest {
	r.retainRepoVersionsGte = &retainRepoVersionsGte
	return r
}

// Filter results where retain_repo_versions has a null value
func (r RepositoriesApiRepositoriesListRequest) RetainRepoVersionsIsnull(retainRepoVersionsIsnull bool) RepositoriesApiRepositoriesListRequest {
	r.retainRepoVersionsIsnull = &retainRepoVersionsIsnull
	return r
}

// Filter results where retain_repo_versions is less than value
func (r RepositoriesApiRepositoriesListRequest) RetainRepoVersionsLt(retainRepoVersionsLt int32) RepositoriesApiRepositoriesListRequest {
	r.retainRepoVersionsLt = &retainRepoVersionsLt
	return r
}

// Filter results where retain_repo_versions is less than or equal to value
func (r RepositoriesApiRepositoriesListRequest) RetainRepoVersionsLte(retainRepoVersionsLte int32) RepositoriesApiRepositoriesListRequest {
	r.retainRepoVersionsLte = &retainRepoVersionsLte
	return r
}

// Filter results where retain_repo_versions not equal to value
func (r RepositoriesApiRepositoriesListRequest) RetainRepoVersionsNe(retainRepoVersionsNe int32) RepositoriesApiRepositoriesListRequest {
	r.retainRepoVersionsNe = &retainRepoVersionsNe
	return r
}

// Filter results where retain_repo_versions is between two comma separated values
func (r RepositoriesApiRepositoriesListRequest) RetainRepoVersionsRange(retainRepoVersionsRange []int32) RepositoriesApiRepositoriesListRequest {
	r.retainRepoVersionsRange = &retainRepoVersionsRange
	return r
}

// A list of fields to include in the response.
func (r RepositoriesApiRepositoriesListRequest) Fields(fields []string) RepositoriesApiRepositoriesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesApiRepositoriesListRequest) ExcludeFields(excludeFields []string) RepositoriesApiRepositoriesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesApiRepositoriesListRequest) Execute() (*PaginatedRepositoryResponseList, *http.Response, error) {
	return r.ApiService.RepositoriesListExecute(r)
}

/*
RepositoriesList List repositories

Endpoint to list all repositories.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RepositoriesApiRepositoriesListRequest
*/
func (a *RepositoriesApiService) RepositoriesList(ctx context.Context) RepositoriesApiRepositoriesListRequest {
	return RepositoriesApiRepositoriesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedRepositoryResponseList
func (a *RepositoriesApiService) RepositoriesListExecute(r RepositoriesApiRepositoriesListRequest) (*PaginatedRepositoryResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRepositoryResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesApiService.RepositoriesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/repositories/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.remote != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "remote", r.remote, "")
	}
	if r.retainRepoVersions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retain_repo_versions", r.retainRepoVersions, "")
	}
	if r.retainRepoVersionsGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retain_repo_versions__gt", r.retainRepoVersionsGt, "")
	}
	if r.retainRepoVersionsGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retain_repo_versions__gte", r.retainRepoVersionsGte, "")
	}
	if r.retainRepoVersionsIsnull != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retain_repo_versions__isnull", r.retainRepoVersionsIsnull, "")
	}
	if r.retainRepoVersionsLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retain_repo_versions__lt", r.retainRepoVersionsLt, "")
	}
	if r.retainRepoVersionsLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retain_repo_versions__lte", r.retainRepoVersionsLte, "")
	}
	if r.retainRepoVersionsNe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retain_repo_versions__ne", r.retainRepoVersionsNe, "")
	}
	if r.retainRepoVersionsRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retain_repo_versions__range", r.retainRepoVersionsRange, "csv")
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
