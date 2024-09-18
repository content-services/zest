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


// RepositoriesAPIService RepositoriesAPI service
type RepositoriesAPIService service

type RepositoriesAPIRepositoriesListRequest struct {
	ctx context.Context
	ApiService *RepositoriesAPIService
	pulpDomain string
	latestWithContent *string
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
	pulpHrefIn *[]string
	pulpIdIn *[]string
	pulpLabelSelect *string
	pulpType *string
	pulpTypeIn *[]string
	q *string
	remote *string
	retainRepoVersions *int32
	retainRepoVersionsGt *int32
	retainRepoVersionsGte *int32
	retainRepoVersionsIsnull *bool
	retainRepoVersionsLt *int32
	retainRepoVersionsLte *int32
	retainRepoVersionsNe *int32
	retainRepoVersionsRange *[]int32
	withContent *string
	fields *[]string
	excludeFields *[]string
}

// Content Unit referenced by HREF
func (r RepositoriesAPIRepositoriesListRequest) LatestWithContent(latestWithContent string) RepositoriesAPIRepositoriesListRequest {
	r.latestWithContent = &latestWithContent
	return r
}

// Number of results to return per page.
func (r RepositoriesAPIRepositoriesListRequest) Limit(limit int32) RepositoriesAPIRepositoriesListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r RepositoriesAPIRepositoriesListRequest) Name(name string) RepositoriesAPIRepositoriesListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r RepositoriesAPIRepositoriesListRequest) NameContains(nameContains string) RepositoriesAPIRepositoriesListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r RepositoriesAPIRepositoriesListRequest) NameIcontains(nameIcontains string) RepositoriesAPIRepositoriesListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name matches value
func (r RepositoriesAPIRepositoriesListRequest) NameIexact(nameIexact string) RepositoriesAPIRepositoriesListRequest {
	r.nameIexact = &nameIexact
	return r
}

// Filter results where name is in a comma-separated list of values
func (r RepositoriesAPIRepositoriesListRequest) NameIn(nameIn []string) RepositoriesAPIRepositoriesListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name matches regex value
func (r RepositoriesAPIRepositoriesListRequest) NameIregex(nameIregex string) RepositoriesAPIRepositoriesListRequest {
	r.nameIregex = &nameIregex
	return r
}

// Filter results where name starts with value
func (r RepositoriesAPIRepositoriesListRequest) NameIstartswith(nameIstartswith string) RepositoriesAPIRepositoriesListRequest {
	r.nameIstartswith = &nameIstartswith
	return r
}

// Filter results where name matches regex value
func (r RepositoriesAPIRepositoriesListRequest) NameRegex(nameRegex string) RepositoriesAPIRepositoriesListRequest {
	r.nameRegex = &nameRegex
	return r
}

// Filter results where name starts with value
func (r RepositoriesAPIRepositoriesListRequest) NameStartswith(nameStartswith string) RepositoriesAPIRepositoriesListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r RepositoriesAPIRepositoriesListRequest) Offset(offset int32) RepositoriesAPIRepositoriesListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;description&#x60; - Description* &#x60;-description&#x60; - Description (descending)* &#x60;next_version&#x60; - Next version* &#x60;-next_version&#x60; - Next version (descending)* &#x60;retain_repo_versions&#x60; - Retain repo versions* &#x60;-retain_repo_versions&#x60; - Retain repo versions (descending)* &#x60;user_hidden&#x60; - User hidden* &#x60;-user_hidden&#x60; - User hidden (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r RepositoriesAPIRepositoriesListRequest) Ordering(ordering []string) RepositoriesAPIRepositoriesListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r RepositoriesAPIRepositoriesListRequest) PulpHrefIn(pulpHrefIn []string) RepositoriesAPIRepositoriesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r RepositoriesAPIRepositoriesListRequest) PulpIdIn(pulpIdIn []string) RepositoriesAPIRepositoriesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter labels by search string
func (r RepositoriesAPIRepositoriesListRequest) PulpLabelSelect(pulpLabelSelect string) RepositoriesAPIRepositoriesListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Pulp type* &#x60;rpm.rpm&#x60; - rpm.rpm* &#x60;file.file&#x60; - file.file* &#x60;gem.gem&#x60; - gem.gem* &#x60;python.python&#x60; - python.python* &#x60;ostree.ostree&#x60; - ostree.ostree
func (r RepositoriesAPIRepositoriesListRequest) PulpType(pulpType string) RepositoriesAPIRepositoriesListRequest {
	r.pulpType = &pulpType
	return r
}

// Multiple values may be separated by commas.* &#x60;rpm.rpm&#x60; - rpm.rpm* &#x60;file.file&#x60; - file.file* &#x60;gem.gem&#x60; - gem.gem* &#x60;python.python&#x60; - python.python* &#x60;ostree.ostree&#x60; - ostree.ostree
func (r RepositoriesAPIRepositoriesListRequest) PulpTypeIn(pulpTypeIn []string) RepositoriesAPIRepositoriesListRequest {
	r.pulpTypeIn = &pulpTypeIn
	return r
}

// Filter results by using NOT, AND and OR operations on other filters
func (r RepositoriesAPIRepositoriesListRequest) Q(q string) RepositoriesAPIRepositoriesListRequest {
	r.q = &q
	return r
}

// Foreign Key referenced by HREF
func (r RepositoriesAPIRepositoriesListRequest) Remote(remote string) RepositoriesAPIRepositoriesListRequest {
	r.remote = &remote
	return r
}

// Filter results where retain_repo_versions matches value
func (r RepositoriesAPIRepositoriesListRequest) RetainRepoVersions(retainRepoVersions int32) RepositoriesAPIRepositoriesListRequest {
	r.retainRepoVersions = &retainRepoVersions
	return r
}

// Filter results where retain_repo_versions is greater than value
func (r RepositoriesAPIRepositoriesListRequest) RetainRepoVersionsGt(retainRepoVersionsGt int32) RepositoriesAPIRepositoriesListRequest {
	r.retainRepoVersionsGt = &retainRepoVersionsGt
	return r
}

// Filter results where retain_repo_versions is greater than or equal to value
func (r RepositoriesAPIRepositoriesListRequest) RetainRepoVersionsGte(retainRepoVersionsGte int32) RepositoriesAPIRepositoriesListRequest {
	r.retainRepoVersionsGte = &retainRepoVersionsGte
	return r
}

// Filter results where retain_repo_versions has a null value
func (r RepositoriesAPIRepositoriesListRequest) RetainRepoVersionsIsnull(retainRepoVersionsIsnull bool) RepositoriesAPIRepositoriesListRequest {
	r.retainRepoVersionsIsnull = &retainRepoVersionsIsnull
	return r
}

// Filter results where retain_repo_versions is less than value
func (r RepositoriesAPIRepositoriesListRequest) RetainRepoVersionsLt(retainRepoVersionsLt int32) RepositoriesAPIRepositoriesListRequest {
	r.retainRepoVersionsLt = &retainRepoVersionsLt
	return r
}

// Filter results where retain_repo_versions is less than or equal to value
func (r RepositoriesAPIRepositoriesListRequest) RetainRepoVersionsLte(retainRepoVersionsLte int32) RepositoriesAPIRepositoriesListRequest {
	r.retainRepoVersionsLte = &retainRepoVersionsLte
	return r
}

// Filter results where retain_repo_versions not equal to value
func (r RepositoriesAPIRepositoriesListRequest) RetainRepoVersionsNe(retainRepoVersionsNe int32) RepositoriesAPIRepositoriesListRequest {
	r.retainRepoVersionsNe = &retainRepoVersionsNe
	return r
}

// Filter results where retain_repo_versions is between two comma separated values
func (r RepositoriesAPIRepositoriesListRequest) RetainRepoVersionsRange(retainRepoVersionsRange []int32) RepositoriesAPIRepositoriesListRequest {
	r.retainRepoVersionsRange = &retainRepoVersionsRange
	return r
}

// Content Unit referenced by HREF
func (r RepositoriesAPIRepositoriesListRequest) WithContent(withContent string) RepositoriesAPIRepositoriesListRequest {
	r.withContent = &withContent
	return r
}

// A list of fields to include in the response.
func (r RepositoriesAPIRepositoriesListRequest) Fields(fields []string) RepositoriesAPIRepositoriesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesAPIRepositoriesListRequest) ExcludeFields(excludeFields []string) RepositoriesAPIRepositoriesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesAPIRepositoriesListRequest) Execute() (*PaginatedRepositoryResponseList, *http.Response, error) {
	return r.ApiService.RepositoriesListExecute(r)
}

/*
RepositoriesList List repositories

Endpoint to list all repositories.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return RepositoriesAPIRepositoriesListRequest
*/
func (a *RepositoriesAPIService) RepositoriesList(ctx context.Context, pulpDomain string) RepositoriesAPIRepositoriesListRequest {
	return RepositoriesAPIRepositoriesListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedRepositoryResponseList
func (a *RepositoriesAPIService) RepositoriesListExecute(r RepositoriesAPIRepositoriesListRequest) (*PaginatedRepositoryResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRepositoryResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesAPIService.RepositoriesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/repositories/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.latestWithContent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "latest_with_content", r.latestWithContent, "form", "")
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
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "form", "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "form", "csv")
	}
	if r.pulpLabelSelect != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_label_select", r.pulpLabelSelect, "form", "")
	}
	if r.pulpType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_type", r.pulpType, "form", "")
	}
	if r.pulpTypeIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_type__in", r.pulpTypeIn, "form", "csv")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
	}
	if r.remote != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "remote", r.remote, "form", "")
	}
	if r.retainRepoVersions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retain_repo_versions", r.retainRepoVersions, "form", "")
	}
	if r.retainRepoVersionsGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retain_repo_versions__gt", r.retainRepoVersionsGt, "form", "")
	}
	if r.retainRepoVersionsGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retain_repo_versions__gte", r.retainRepoVersionsGte, "form", "")
	}
	if r.retainRepoVersionsIsnull != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retain_repo_versions__isnull", r.retainRepoVersionsIsnull, "form", "")
	}
	if r.retainRepoVersionsLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retain_repo_versions__lt", r.retainRepoVersionsLt, "form", "")
	}
	if r.retainRepoVersionsLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retain_repo_versions__lte", r.retainRepoVersionsLte, "form", "")
	}
	if r.retainRepoVersionsNe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retain_repo_versions__ne", r.retainRepoVersionsNe, "form", "")
	}
	if r.retainRepoVersionsRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retain_repo_versions__range", r.retainRepoVersionsRange, "form", "csv")
	}
	if r.withContent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "with_content", r.withContent, "form", "")
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
