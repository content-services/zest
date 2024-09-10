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
	"os"
	"reflect"
)


// ContentAdvisoriesAPIService ContentAdvisoriesAPI service
type ContentAdvisoriesAPIService service

type ContentAdvisoriesAPIContentRpmAdvisoriesCreateRequest struct {
	ctx context.Context
	ApiService *ContentAdvisoriesAPIService
	pulpDomain string
	repository *string
	file *os.File
	upload *string
	fileUrl *string
}

// A URI of a repository the new content unit should be associated with.
func (r ContentAdvisoriesAPIContentRpmAdvisoriesCreateRequest) Repository(repository string) ContentAdvisoriesAPIContentRpmAdvisoriesCreateRequest {
	r.repository = &repository
	return r
}

// An uploaded file that may be turned into the content unit.
func (r ContentAdvisoriesAPIContentRpmAdvisoriesCreateRequest) File(file *os.File) ContentAdvisoriesAPIContentRpmAdvisoriesCreateRequest {
	r.file = file
	return r
}

// An uncommitted upload that may be turned into the content unit.
func (r ContentAdvisoriesAPIContentRpmAdvisoriesCreateRequest) Upload(upload string) ContentAdvisoriesAPIContentRpmAdvisoriesCreateRequest {
	r.upload = &upload
	return r
}

// A url that Pulp can download and turn into the content unit.
func (r ContentAdvisoriesAPIContentRpmAdvisoriesCreateRequest) FileUrl(fileUrl string) ContentAdvisoriesAPIContentRpmAdvisoriesCreateRequest {
	r.fileUrl = &fileUrl
	return r
}

func (r ContentAdvisoriesAPIContentRpmAdvisoriesCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentRpmAdvisoriesCreateExecute(r)
}

/*
ContentRpmAdvisoriesCreate Create an update record

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentAdvisoriesAPIContentRpmAdvisoriesCreateRequest
*/
func (a *ContentAdvisoriesAPIService) ContentRpmAdvisoriesCreate(ctx context.Context, pulpDomain string) ContentAdvisoriesAPIContentRpmAdvisoriesCreateRequest {
	return ContentAdvisoriesAPIContentRpmAdvisoriesCreateRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentAdvisoriesAPIService) ContentRpmAdvisoriesCreateExecute(r ContentAdvisoriesAPIContentRpmAdvisoriesCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentAdvisoriesAPIService.ContentRpmAdvisoriesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/rpm/advisories/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

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
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "repository", r.repository, "", "")
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
	if r.upload != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "upload", r.upload, "", "")
	}
	if r.fileUrl != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "file_url", r.fileUrl, "", "")
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

type ContentAdvisoriesAPIContentRpmAdvisoriesListRequest struct {
	ctx context.Context
	ApiService *ContentAdvisoriesAPIService
	pulpDomain string
	id *string
	idIn *[]string
	limit *int32
	offset *int32
	ordering *[]string
	orphanedFor *float32
	pulpHrefIn *[]string
	pulpIdIn *[]string
	q *string
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
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) Id(id string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.id = &id
	return r
}

// Filter results where id is in a comma-separated list of values
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) IdIn(idIn []string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.idIn = &idIn
	return r
}

// Number of results to return per page.
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) Limit(limit int32) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) Offset(offset int32) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;id&#x60; - Id* &#x60;-id&#x60; - Id (descending)* &#x60;updated_date&#x60; - Updated date* &#x60;-updated_date&#x60; - Updated date (descending)* &#x60;description&#x60; - Description* &#x60;-description&#x60; - Description (descending)* &#x60;issued_date&#x60; - Issued date* &#x60;-issued_date&#x60; - Issued date (descending)* &#x60;fromstr&#x60; - Fromstr* &#x60;-fromstr&#x60; - Fromstr (descending)* &#x60;status&#x60; - Status* &#x60;-status&#x60; - Status (descending)* &#x60;title&#x60; - Title* &#x60;-title&#x60; - Title (descending)* &#x60;summary&#x60; - Summary* &#x60;-summary&#x60; - Summary (descending)* &#x60;version&#x60; - Version* &#x60;-version&#x60; - Version (descending)* &#x60;type&#x60; - Type* &#x60;-type&#x60; - Type (descending)* &#x60;severity&#x60; - Severity* &#x60;-severity&#x60; - Severity (descending)* &#x60;solution&#x60; - Solution* &#x60;-solution&#x60; - Solution (descending)* &#x60;release&#x60; - Release* &#x60;-release&#x60; - Release (descending)* &#x60;rights&#x60; - Rights* &#x60;-rights&#x60; - Rights (descending)* &#x60;reboot_suggested&#x60; - Reboot suggested* &#x60;-reboot_suggested&#x60; - Reboot suggested (descending)* &#x60;pushcount&#x60; - Pushcount* &#x60;-pushcount&#x60; - Pushcount (descending)* &#x60;digest&#x60; - Digest* &#x60;-digest&#x60; - Digest (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) Ordering(ordering []string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.ordering = &ordering
	return r
}

// Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME.
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) OrphanedFor(orphanedFor float32) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.orphanedFor = &orphanedFor
	return r
}

// Multiple values may be separated by commas.
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) PulpHrefIn(pulpHrefIn []string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) PulpIdIn(pulpIdIn []string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) Q(q string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) RepositoryVersion(repositoryVersion string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter results where severity matches value
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) Severity(severity string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.severity = &severity
	return r
}

// Filter results where severity is in a comma-separated list of values
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) SeverityIn(severityIn []string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.severityIn = &severityIn
	return r
}

// Filter results where severity not equal to value
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) SeverityNe(severityNe string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.severityNe = &severityNe
	return r
}

// Filter results where status matches value
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) Status(status string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.status = &status
	return r
}

// Filter results where status is in a comma-separated list of values
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) StatusIn(statusIn []string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.statusIn = &statusIn
	return r
}

// Filter results where status not equal to value
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) StatusNe(statusNe string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.statusNe = &statusNe
	return r
}

// Filter results where type matches value
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) Type_(type_ string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.type_ = &type_
	return r
}

// Filter results where type is in a comma-separated list of values
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) TypeIn(typeIn []string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.typeIn = &typeIn
	return r
}

// Filter results where type not equal to value
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) TypeNe(typeNe string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.typeNe = &typeNe
	return r
}

// A list of fields to include in the response.
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) Fields(fields []string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) ExcludeFields(excludeFields []string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) Execute() (*PaginatedrpmUpdateRecordResponseList, *http.Response, error) {
	return r.ApiService.ContentRpmAdvisoriesListExecute(r)
}

/*
ContentRpmAdvisoriesList List update records

A ViewSet for UpdateRecord.Define endpoint name which will appear in the API endpoint for this content type.For example::    http://pulp.example.com/pulp/api/v3/content/rpm/advisories/Also specify queryset and serializer for UpdateRecord.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentAdvisoriesAPIContentRpmAdvisoriesListRequest
*/
func (a *ContentAdvisoriesAPIService) ContentRpmAdvisoriesList(ctx context.Context, pulpDomain string) ContentAdvisoriesAPIContentRpmAdvisoriesListRequest {
	return ContentAdvisoriesAPIContentRpmAdvisoriesListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedrpmUpdateRecordResponseList
func (a *ContentAdvisoriesAPIService) ContentRpmAdvisoriesListExecute(r ContentAdvisoriesAPIContentRpmAdvisoriesListRequest) (*PaginatedrpmUpdateRecordResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedrpmUpdateRecordResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentAdvisoriesAPIService.ContentRpmAdvisoriesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/content/rpm/advisories/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
	}
	if r.idIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id__in", r.idIn, "form", "csv")
	}
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
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "form", "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "form", "csv")
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
	if r.severity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "severity", r.severity, "form", "")
	}
	if r.severityIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "severity__in", r.severityIn, "form", "csv")
	}
	if r.severityNe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "severity__ne", r.severityNe, "form", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.statusIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status__in", r.statusIn, "form", "csv")
	}
	if r.statusNe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status__ne", r.statusNe, "form", "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.typeIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type__in", r.typeIn, "form", "csv")
	}
	if r.typeNe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type__ne", r.typeNe, "form", "")
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

type ContentAdvisoriesAPIContentRpmAdvisoriesReadRequest struct {
	ctx context.Context
	ApiService *ContentAdvisoriesAPIService
	rpmUpdateRecordHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentAdvisoriesAPIContentRpmAdvisoriesReadRequest) Fields(fields []string) ContentAdvisoriesAPIContentRpmAdvisoriesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentAdvisoriesAPIContentRpmAdvisoriesReadRequest) ExcludeFields(excludeFields []string) ContentAdvisoriesAPIContentRpmAdvisoriesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentAdvisoriesAPIContentRpmAdvisoriesReadRequest) Execute() (*RpmUpdateRecordResponse, *http.Response, error) {
	return r.ApiService.ContentRpmAdvisoriesReadExecute(r)
}

/*
ContentRpmAdvisoriesRead Inspect an update record

A ViewSet for UpdateRecord.Define endpoint name which will appear in the API endpoint for this content type.For example::    http://pulp.example.com/pulp/api/v3/content/rpm/advisories/Also specify queryset and serializer for UpdateRecord.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmUpdateRecordHref
 @return ContentAdvisoriesAPIContentRpmAdvisoriesReadRequest
*/
func (a *ContentAdvisoriesAPIService) ContentRpmAdvisoriesRead(ctx context.Context, rpmUpdateRecordHref string) ContentAdvisoriesAPIContentRpmAdvisoriesReadRequest {
	return ContentAdvisoriesAPIContentRpmAdvisoriesReadRequest{
		ApiService: a,
		ctx: ctx,
		rpmUpdateRecordHref: rpmUpdateRecordHref,
	}
}

// Execute executes the request
//  @return RpmUpdateRecordResponse
func (a *ContentAdvisoriesAPIService) ContentRpmAdvisoriesReadExecute(r ContentAdvisoriesAPIContentRpmAdvisoriesReadRequest) (*RpmUpdateRecordResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RpmUpdateRecordResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentAdvisoriesAPIService.ContentRpmAdvisoriesRead")
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
