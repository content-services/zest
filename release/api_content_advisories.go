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
	"os"
	"reflect"
)


// ContentAdvisoriesApiService ContentAdvisoriesApi service
type ContentAdvisoriesApiService service

type ContentAdvisoriesApiContentRpmAdvisoriesCreateRequest struct {
	ctx context.Context
	ApiService *ContentAdvisoriesApiService
	file *os.File
	repository *string
}

// An uploaded file that may be turned into the artifact of the content unit.
func (r ContentAdvisoriesApiContentRpmAdvisoriesCreateRequest) File(file *os.File) ContentAdvisoriesApiContentRpmAdvisoriesCreateRequest {
	r.file = file
	return r
}

// A URI of a repository the new content unit should be associated with.
func (r ContentAdvisoriesApiContentRpmAdvisoriesCreateRequest) Repository(repository string) ContentAdvisoriesApiContentRpmAdvisoriesCreateRequest {
	r.repository = &repository
	return r
}

func (r ContentAdvisoriesApiContentRpmAdvisoriesCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentRpmAdvisoriesCreateExecute(r)
}

/*
ContentRpmAdvisoriesCreate Create an update record

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentAdvisoriesApiContentRpmAdvisoriesCreateRequest
*/
func (a *ContentAdvisoriesApiService) ContentRpmAdvisoriesCreate(ctx context.Context) ContentAdvisoriesApiContentRpmAdvisoriesCreateRequest {
	return ContentAdvisoriesApiContentRpmAdvisoriesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentAdvisoriesApiService) ContentRpmAdvisoriesCreateExecute(r ContentAdvisoriesApiContentRpmAdvisoriesCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentAdvisoriesApiService.ContentRpmAdvisoriesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/rpm/advisories/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data", "application/x-www-form-urlencoded"}

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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "repository", r.repository, "")
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

type ContentAdvisoriesApiContentRpmAdvisoriesListRequest struct {
	ctx context.Context
	ApiService *ContentAdvisoriesApiService
	id *string
	idIn *[]string
	limit *int32
	offset *int32
	ordering *[]string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	severity *string
	severityIn *[]string
	severityNe *string
	status *string
	statusIn *[]string
	statusNe *string
	type_ *string
	typeIn *[]string
	typeNe *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where id matches value
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) Id(id string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.id = &id
	return r
}

// Filter results where id is in a comma-separated list of values
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) IdIn(idIn []string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.idIn = &idIn
	return r
}

// Number of results to return per page.
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) Limit(limit int32) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) Offset(offset int32) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) Ordering(ordering []string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.ordering = &ordering
	return r
}

// Repository Version referenced by HREF
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) RepositoryVersion(repositoryVersion string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter results where severity matches value
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) Severity(severity string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.severity = &severity
	return r
}

// Filter results where severity is in a comma-separated list of values
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) SeverityIn(severityIn []string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.severityIn = &severityIn
	return r
}

// Filter results where severity not equal to value
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) SeverityNe(severityNe string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.severityNe = &severityNe
	return r
}

// Filter results where status matches value
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) Status(status string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.status = &status
	return r
}

// Filter results where status is in a comma-separated list of values
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) StatusIn(statusIn []string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.statusIn = &statusIn
	return r
}

// Filter results where status not equal to value
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) StatusNe(statusNe string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.statusNe = &statusNe
	return r
}

// Filter results where type matches value
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) Type_(type_ string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.type_ = &type_
	return r
}

// Filter results where type is in a comma-separated list of values
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) TypeIn(typeIn []string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.typeIn = &typeIn
	return r
}

// Filter results where type not equal to value
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) TypeNe(typeNe string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.typeNe = &typeNe
	return r
}

// A list of fields to include in the response.
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) Fields(fields []string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) ExcludeFields(excludeFields []string) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) Execute() (*PaginatedrpmUpdateRecordResponseList, *http.Response, error) {
	return r.ApiService.ContentRpmAdvisoriesListExecute(r)
}

/*
ContentRpmAdvisoriesList List update records

A ViewSet for UpdateRecord.

Define endpoint name which will appear in the API endpoint for this content type.
For example::
    http://pulp.example.com/pulp/api/v3/content/rpm/advisories/

Also specify queryset and serializer for UpdateRecord.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentAdvisoriesApiContentRpmAdvisoriesListRequest
*/
func (a *ContentAdvisoriesApiService) ContentRpmAdvisoriesList(ctx context.Context) ContentAdvisoriesApiContentRpmAdvisoriesListRequest {
	return ContentAdvisoriesApiContentRpmAdvisoriesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedrpmUpdateRecordResponseList
func (a *ContentAdvisoriesApiService) ContentRpmAdvisoriesListExecute(r ContentAdvisoriesApiContentRpmAdvisoriesListRequest) (*PaginatedrpmUpdateRecordResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedrpmUpdateRecordResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentAdvisoriesApiService.ContentRpmAdvisoriesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/rpm/advisories/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "")
	}
	if r.idIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id__in", r.idIn, "csv")
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
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
	}
	if r.severity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "severity", r.severity, "")
	}
	if r.severityIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "severity__in", r.severityIn, "csv")
	}
	if r.severityNe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "severity__ne", r.severityNe, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.statusIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status__in", r.statusIn, "csv")
	}
	if r.statusNe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status__ne", r.statusNe, "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	if r.typeIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type__in", r.typeIn, "csv")
	}
	if r.typeNe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type__ne", r.typeNe, "")
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

type ContentAdvisoriesApiContentRpmAdvisoriesReadRequest struct {
	ctx context.Context
	ApiService *ContentAdvisoriesApiService
	rpmUpdateRecordHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentAdvisoriesApiContentRpmAdvisoriesReadRequest) Fields(fields []string) ContentAdvisoriesApiContentRpmAdvisoriesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentAdvisoriesApiContentRpmAdvisoriesReadRequest) ExcludeFields(excludeFields []string) ContentAdvisoriesApiContentRpmAdvisoriesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentAdvisoriesApiContentRpmAdvisoriesReadRequest) Execute() (*RpmUpdateRecordResponse, *http.Response, error) {
	return r.ApiService.ContentRpmAdvisoriesReadExecute(r)
}

/*
ContentRpmAdvisoriesRead Inspect an update record

A ViewSet for UpdateRecord.

Define endpoint name which will appear in the API endpoint for this content type.
For example::
    http://pulp.example.com/pulp/api/v3/content/rpm/advisories/

Also specify queryset and serializer for UpdateRecord.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmUpdateRecordHref
 @return ContentAdvisoriesApiContentRpmAdvisoriesReadRequest
*/
func (a *ContentAdvisoriesApiService) ContentRpmAdvisoriesRead(ctx context.Context, rpmUpdateRecordHref string) ContentAdvisoriesApiContentRpmAdvisoriesReadRequest {
	return ContentAdvisoriesApiContentRpmAdvisoriesReadRequest{
		ApiService: a,
		ctx: ctx,
		rpmUpdateRecordHref: rpmUpdateRecordHref,
	}
}

// Execute executes the request
//  @return RpmUpdateRecordResponse
func (a *ContentAdvisoriesApiService) ContentRpmAdvisoriesReadExecute(r ContentAdvisoriesApiContentRpmAdvisoriesReadRequest) (*RpmUpdateRecordResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RpmUpdateRecordResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentAdvisoriesApiService.ContentRpmAdvisoriesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{rpm_update_record_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_update_record_href"+"}", url.PathEscape(parameterValueToString(r.rpmUpdateRecordHref, "rpmUpdateRecordHref")), -1)
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
