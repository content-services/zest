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


// DistributionsAPIService DistributionsAPI service
type DistributionsAPIService service

type DistributionsAPIDistributionsListRequest struct {
	ctx context.Context
	ApiService *DistributionsAPIService
	pulpDomain string
	basePath *string
	basePathContains *string
	basePathIcontains *string
	basePathIn *[]string
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
	prnIn *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	pulpLabelSelect *string
	pulpType *string
	pulpTypeIn *[]string
	q *string
	repository *string
	repositoryIn *[]string
	withContent *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where base_path matches value
func (r DistributionsAPIDistributionsListRequest) BasePath(basePath string) DistributionsAPIDistributionsListRequest {
	r.basePath = &basePath
	return r
}

// Filter results where base_path contains value
func (r DistributionsAPIDistributionsListRequest) BasePathContains(basePathContains string) DistributionsAPIDistributionsListRequest {
	r.basePathContains = &basePathContains
	return r
}

// Filter results where base_path contains value
func (r DistributionsAPIDistributionsListRequest) BasePathIcontains(basePathIcontains string) DistributionsAPIDistributionsListRequest {
	r.basePathIcontains = &basePathIcontains
	return r
}

// Filter results where base_path is in a comma-separated list of values
func (r DistributionsAPIDistributionsListRequest) BasePathIn(basePathIn []string) DistributionsAPIDistributionsListRequest {
	r.basePathIn = &basePathIn
	return r
}

// Number of results to return per page.
func (r DistributionsAPIDistributionsListRequest) Limit(limit int32) DistributionsAPIDistributionsListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r DistributionsAPIDistributionsListRequest) Name(name string) DistributionsAPIDistributionsListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r DistributionsAPIDistributionsListRequest) NameContains(nameContains string) DistributionsAPIDistributionsListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r DistributionsAPIDistributionsListRequest) NameIcontains(nameIcontains string) DistributionsAPIDistributionsListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name matches value
func (r DistributionsAPIDistributionsListRequest) NameIexact(nameIexact string) DistributionsAPIDistributionsListRequest {
	r.nameIexact = &nameIexact
	return r
}

// Filter results where name is in a comma-separated list of values
func (r DistributionsAPIDistributionsListRequest) NameIn(nameIn []string) DistributionsAPIDistributionsListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name matches regex value
func (r DistributionsAPIDistributionsListRequest) NameIregex(nameIregex string) DistributionsAPIDistributionsListRequest {
	r.nameIregex = &nameIregex
	return r
}

// Filter results where name starts with value
func (r DistributionsAPIDistributionsListRequest) NameIstartswith(nameIstartswith string) DistributionsAPIDistributionsListRequest {
	r.nameIstartswith = &nameIstartswith
	return r
}

// Filter results where name matches regex value
func (r DistributionsAPIDistributionsListRequest) NameRegex(nameRegex string) DistributionsAPIDistributionsListRequest {
	r.nameRegex = &nameRegex
	return r
}

// Filter results where name starts with value
func (r DistributionsAPIDistributionsListRequest) NameStartswith(nameStartswith string) DistributionsAPIDistributionsListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r DistributionsAPIDistributionsListRequest) Offset(offset int32) DistributionsAPIDistributionsListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;base_path&#x60; - Base path* &#x60;-base_path&#x60; - Base path (descending)* &#x60;hidden&#x60; - Hidden* &#x60;-hidden&#x60; - Hidden (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r DistributionsAPIDistributionsListRequest) Ordering(ordering []string) DistributionsAPIDistributionsListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r DistributionsAPIDistributionsListRequest) PrnIn(prnIn []string) DistributionsAPIDistributionsListRequest {
	r.prnIn = &prnIn
	return r
}

// Multiple values may be separated by commas.
func (r DistributionsAPIDistributionsListRequest) PulpHrefIn(pulpHrefIn []string) DistributionsAPIDistributionsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r DistributionsAPIDistributionsListRequest) PulpIdIn(pulpIdIn []string) DistributionsAPIDistributionsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter labels by search string
func (r DistributionsAPIDistributionsListRequest) PulpLabelSelect(pulpLabelSelect string) DistributionsAPIDistributionsListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Pulp type* &#x60;core.artifact&#x60; - core.artifact* &#x60;core.openpgp&#x60; - core.openpgp* &#x60;rpm.rpm&#x60; - rpm.rpm* &#x60;file.file&#x60; - file.file* &#x60;ostree.ostree&#x60; - ostree.ostree* &#x60;python.python&#x60; - python.python* &#x60;gem.gem&#x60; - gem.gem
func (r DistributionsAPIDistributionsListRequest) PulpType(pulpType string) DistributionsAPIDistributionsListRequest {
	r.pulpType = &pulpType
	return r
}

// Multiple values may be separated by commas.* &#x60;core.artifact&#x60; - core.artifact* &#x60;core.openpgp&#x60; - core.openpgp* &#x60;rpm.rpm&#x60; - rpm.rpm* &#x60;file.file&#x60; - file.file* &#x60;ostree.ostree&#x60; - ostree.ostree* &#x60;python.python&#x60; - python.python* &#x60;gem.gem&#x60; - gem.gem
func (r DistributionsAPIDistributionsListRequest) PulpTypeIn(pulpTypeIn []string) DistributionsAPIDistributionsListRequest {
	r.pulpTypeIn = &pulpTypeIn
	return r
}

// Filter results by using NOT, AND and OR operations on other filters
func (r DistributionsAPIDistributionsListRequest) Q(q string) DistributionsAPIDistributionsListRequest {
	r.q = &q
	return r
}

// Filter results where repository matches value
func (r DistributionsAPIDistributionsListRequest) Repository(repository string) DistributionsAPIDistributionsListRequest {
	r.repository = &repository
	return r
}

// Filter results where repository is in a comma-separated list of values
func (r DistributionsAPIDistributionsListRequest) RepositoryIn(repositoryIn []string) DistributionsAPIDistributionsListRequest {
	r.repositoryIn = &repositoryIn
	return r
}

// Filter distributions based on the content served by them
func (r DistributionsAPIDistributionsListRequest) WithContent(withContent string) DistributionsAPIDistributionsListRequest {
	r.withContent = &withContent
	return r
}

// A list of fields to include in the response.
func (r DistributionsAPIDistributionsListRequest) Fields(fields []string) DistributionsAPIDistributionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r DistributionsAPIDistributionsListRequest) ExcludeFields(excludeFields []string) DistributionsAPIDistributionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r DistributionsAPIDistributionsListRequest) Execute() (*PaginatedDistributionResponseList, *http.Response, error) {
	return r.ApiService.DistributionsListExecute(r)
}

/*
DistributionsList List distributions

Provides base viewset for Distributions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return DistributionsAPIDistributionsListRequest
*/
func (a *DistributionsAPIService) DistributionsList(ctx context.Context, pulpDomain string) DistributionsAPIDistributionsListRequest {
	return DistributionsAPIDistributionsListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedDistributionResponseList
func (a *DistributionsAPIService) DistributionsListExecute(r DistributionsAPIDistributionsListRequest) (*PaginatedDistributionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDistributionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsAPIService.DistributionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/distributions/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.basePath != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path", r.basePath, "form", "")
	}
	if r.basePathContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path__contains", r.basePathContains, "form", "")
	}
	if r.basePathIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path__icontains", r.basePathIcontains, "form", "")
	}
	if r.basePathIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path__in", r.basePathIn, "form", "csv")
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
	if r.pulpType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_type", r.pulpType, "form", "")
	}
	if r.pulpTypeIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_type__in", r.pulpTypeIn, "form", "csv")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
	}
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository", r.repository, "form", "")
	}
	if r.repositoryIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository__in", r.repositoryIn, "form", "csv")
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
