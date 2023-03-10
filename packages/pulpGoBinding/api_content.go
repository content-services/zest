/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpGoBinding

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
)


// ContentApiService ContentApi service
type ContentApiService service

type ContentApiContentListRequest struct {
	ctx context.Context
	ApiService *ContentApiService
	limit *int32
	offset *int32
	ordering *[]string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ContentApiContentListRequest) Limit(limit int32) ContentApiContentListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentApiContentListRequest) Offset(offset int32) ContentApiContentListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ContentApiContentListRequest) Ordering(ordering []string) ContentApiContentListRequest {
	r.ordering = &ordering
	return r
}

// Repository Version referenced by HREF
func (r ContentApiContentListRequest) RepositoryVersion(repositoryVersion string) ContentApiContentListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentApiContentListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentApiContentListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentApiContentListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentApiContentListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentApiContentListRequest) Fields(fields []string) ContentApiContentListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentApiContentListRequest) ExcludeFields(excludeFields []string) ContentApiContentListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentApiContentListRequest) Execute() (*PaginatedMultipleArtifactContentResponseList, *http.Response, error) {
	return r.ApiService.ContentListExecute(r)
}

/*
ContentList List content

Endpoint to list all content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentApiContentListRequest
*/
func (a *ContentApiService) ContentList(ctx context.Context) ContentApiContentListRequest {
	return ContentApiContentListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedMultipleArtifactContentResponseList
func (a *ContentApiService) ContentListExecute(r ContentApiContentListRequest) (*PaginatedMultipleArtifactContentResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedMultipleArtifactContentResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentApiService.ContentList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
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
