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


// ContentAPIService ContentAPI service
type ContentAPIService service

type ContentAPIContentListRequest struct {
	ctx context.Context
	ApiService *ContentAPIService
	pulpDomain string
	limit *int32
	offset *int32
	ordering *[]string
	orphanedFor *float32
	prnIn *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	pulpLabelSelect *string
	pulpType *string
	pulpTypeIn *[]string
	q *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ContentAPIContentListRequest) Limit(limit int32) ContentAPIContentListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentAPIContentListRequest) Offset(offset int32) ContentAPIContentListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r ContentAPIContentListRequest) Ordering(ordering []string) ContentAPIContentListRequest {
	r.ordering = &ordering
	return r
}

// Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME.
func (r ContentAPIContentListRequest) OrphanedFor(orphanedFor float32) ContentAPIContentListRequest {
	r.orphanedFor = &orphanedFor
	return r
}

// Multiple values may be separated by commas.
func (r ContentAPIContentListRequest) PrnIn(prnIn []string) ContentAPIContentListRequest {
	r.prnIn = &prnIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentAPIContentListRequest) PulpHrefIn(pulpHrefIn []string) ContentAPIContentListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentAPIContentListRequest) PulpIdIn(pulpIdIn []string) ContentAPIContentListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter labels by search string
func (r ContentAPIContentListRequest) PulpLabelSelect(pulpLabelSelect string) ContentAPIContentListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Pulp type* &#x60;core.publishedmetadata&#x60; - core.publishedmetadata* &#x60;core.openpgp_publickey&#x60; - core.openpgp_publickey* &#x60;core.openpgp_publicsubkey&#x60; - core.openpgp_publicsubkey* &#x60;core.openpgp_userid&#x60; - core.openpgp_userid* &#x60;core.openpgp_userattribute&#x60; - core.openpgp_userattribute* &#x60;core.openpgp_signature&#x60; - core.openpgp_signature* &#x60;container.blob&#x60; - container.blob* &#x60;container.manifest&#x60; - container.manifest* &#x60;container.tag&#x60; - container.tag* &#x60;container.signature&#x60; - container.signature* &#x60;ostree.object&#x60; - ostree.object* &#x60;ostree.commit&#x60; - ostree.commit* &#x60;ostree.refs&#x60; - ostree.refs* &#x60;ostree.content&#x60; - ostree.content* &#x60;ostree.config&#x60; - ostree.config* &#x60;ostree.summary&#x60; - ostree.summary* &#x60;rpm.advisory&#x60; - rpm.advisory* &#x60;rpm.packagegroup&#x60; - rpm.packagegroup* &#x60;rpm.packagecategory&#x60; - rpm.packagecategory* &#x60;rpm.packageenvironment&#x60; - rpm.packageenvironment* &#x60;rpm.packagelangpacks&#x60; - rpm.packagelangpacks* &#x60;rpm.repo_metadata_file&#x60; - rpm.repo_metadata_file* &#x60;rpm.distribution_tree&#x60; - rpm.distribution_tree* &#x60;rpm.package&#x60; - rpm.package* &#x60;rpm.modulemd&#x60; - rpm.modulemd* &#x60;rpm.modulemd_defaults&#x60; - rpm.modulemd_defaults* &#x60;rpm.modulemd_obsolete&#x60; - rpm.modulemd_obsolete* &#x60;gem.gem&#x60; - gem.gem* &#x60;npm.package&#x60; - npm.package* &#x60;file.file&#x60; - file.file* &#x60;python.python&#x60; - python.python
func (r ContentAPIContentListRequest) PulpType(pulpType string) ContentAPIContentListRequest {
	r.pulpType = &pulpType
	return r
}

// Multiple values may be separated by commas.* &#x60;core.publishedmetadata&#x60; - core.publishedmetadata* &#x60;core.openpgp_publickey&#x60; - core.openpgp_publickey* &#x60;core.openpgp_publicsubkey&#x60; - core.openpgp_publicsubkey* &#x60;core.openpgp_userid&#x60; - core.openpgp_userid* &#x60;core.openpgp_userattribute&#x60; - core.openpgp_userattribute* &#x60;core.openpgp_signature&#x60; - core.openpgp_signature* &#x60;container.blob&#x60; - container.blob* &#x60;container.manifest&#x60; - container.manifest* &#x60;container.tag&#x60; - container.tag* &#x60;container.signature&#x60; - container.signature* &#x60;ostree.object&#x60; - ostree.object* &#x60;ostree.commit&#x60; - ostree.commit* &#x60;ostree.refs&#x60; - ostree.refs* &#x60;ostree.content&#x60; - ostree.content* &#x60;ostree.config&#x60; - ostree.config* &#x60;ostree.summary&#x60; - ostree.summary* &#x60;rpm.advisory&#x60; - rpm.advisory* &#x60;rpm.packagegroup&#x60; - rpm.packagegroup* &#x60;rpm.packagecategory&#x60; - rpm.packagecategory* &#x60;rpm.packageenvironment&#x60; - rpm.packageenvironment* &#x60;rpm.packagelangpacks&#x60; - rpm.packagelangpacks* &#x60;rpm.repo_metadata_file&#x60; - rpm.repo_metadata_file* &#x60;rpm.distribution_tree&#x60; - rpm.distribution_tree* &#x60;rpm.package&#x60; - rpm.package* &#x60;rpm.modulemd&#x60; - rpm.modulemd* &#x60;rpm.modulemd_defaults&#x60; - rpm.modulemd_defaults* &#x60;rpm.modulemd_obsolete&#x60; - rpm.modulemd_obsolete* &#x60;gem.gem&#x60; - gem.gem* &#x60;npm.package&#x60; - npm.package* &#x60;file.file&#x60; - file.file* &#x60;python.python&#x60; - python.python
func (r ContentAPIContentListRequest) PulpTypeIn(pulpTypeIn []string) ContentAPIContentListRequest {
	r.pulpTypeIn = &pulpTypeIn
	return r
}

// Filter results by using NOT, AND and OR operations on other filters
func (r ContentAPIContentListRequest) Q(q string) ContentAPIContentListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentAPIContentListRequest) RepositoryVersion(repositoryVersion string) ContentAPIContentListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentAPIContentListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentAPIContentListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF/PRN
func (r ContentAPIContentListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentAPIContentListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentAPIContentListRequest) Fields(fields []string) ContentAPIContentListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentAPIContentListRequest) ExcludeFields(excludeFields []string) ContentAPIContentListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentAPIContentListRequest) Execute() (*PaginatedMultipleArtifactContentResponseList, *http.Response, error) {
	return r.ApiService.ContentListExecute(r)
}

/*
ContentList List content

Endpoint to list all content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentAPIContentListRequest
*/
func (a *ContentAPIService) ContentList(ctx context.Context, pulpDomain string) ContentAPIContentListRequest {
	return ContentAPIContentListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedMultipleArtifactContentResponseList
func (a *ContentAPIService) ContentListExecute(r ContentAPIContentListRequest) (*PaginatedMultipleArtifactContentResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedMultipleArtifactContentResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentAPIService.ContentList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/"
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
