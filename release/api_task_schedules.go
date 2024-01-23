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


// TaskSchedulesAPIService TaskSchedulesAPI service
type TaskSchedulesAPIService service

type TaskSchedulesAPITaskSchedulesAddRoleRequest struct {
	ctx context.Context
	ApiService *TaskSchedulesAPIService
	taskScheduleHref string
	nestedRole *NestedRole
}

func (r TaskSchedulesAPITaskSchedulesAddRoleRequest) NestedRole(nestedRole NestedRole) TaskSchedulesAPITaskSchedulesAddRoleRequest {
	r.nestedRole = &nestedRole
	return r
}

func (r TaskSchedulesAPITaskSchedulesAddRoleRequest) Execute() (*NestedRoleResponse, *http.Response, error) {
	return r.ApiService.TaskSchedulesAddRoleExecute(r)
}

/*
TaskSchedulesAddRole Add a role

Add a role for this object to users/groups.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskScheduleHref
 @return TaskSchedulesAPITaskSchedulesAddRoleRequest
*/
func (a *TaskSchedulesAPIService) TaskSchedulesAddRole(ctx context.Context, taskScheduleHref string) TaskSchedulesAPITaskSchedulesAddRoleRequest {
	return TaskSchedulesAPITaskSchedulesAddRoleRequest{
		ApiService: a,
		ctx: ctx,
		taskScheduleHref: taskScheduleHref,
	}
}

// Execute executes the request
//  @return NestedRoleResponse
func (a *TaskSchedulesAPIService) TaskSchedulesAddRoleExecute(r TaskSchedulesAPITaskSchedulesAddRoleRequest) (*NestedRoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NestedRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesAPIService.TaskSchedulesAddRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{task_schedule_href}add_role/"
	localVarPath = strings.Replace(localVarPath, "{"+"task_schedule_href"+"}", url.PathEscape(parameterValueToString(r.taskScheduleHref, "taskScheduleHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.nestedRole == nil {
		return localVarReturnValue, nil, reportError("nestedRole is required and must be specified")
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
	localVarPostBody = r.nestedRole
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

type TaskSchedulesAPITaskSchedulesListRequest struct {
	ctx context.Context
	ApiService *TaskSchedulesAPIService
	pulpDomain string
	limit *int32
	name *string
	nameContains *string
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	q *string
	taskName *string
	taskNameContains *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r TaskSchedulesAPITaskSchedulesListRequest) Limit(limit int32) TaskSchedulesAPITaskSchedulesListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r TaskSchedulesAPITaskSchedulesListRequest) Name(name string) TaskSchedulesAPITaskSchedulesListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r TaskSchedulesAPITaskSchedulesListRequest) NameContains(nameContains string) TaskSchedulesAPITaskSchedulesListRequest {
	r.nameContains = &nameContains
	return r
}

// The initial index from which to return the results.
func (r TaskSchedulesAPITaskSchedulesListRequest) Offset(offset int32) TaskSchedulesAPITaskSchedulesListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;next_dispatch&#x60; - Next dispatch * &#x60;-next_dispatch&#x60; - Next dispatch (descending) * &#x60;dispatch_interval&#x60; - Dispatch interval * &#x60;-dispatch_interval&#x60; - Dispatch interval (descending) * &#x60;task_name&#x60; - Task name * &#x60;-task_name&#x60; - Task name (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r TaskSchedulesAPITaskSchedulesListRequest) Ordering(ordering []string) TaskSchedulesAPITaskSchedulesListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r TaskSchedulesAPITaskSchedulesListRequest) PulpHrefIn(pulpHrefIn []string) TaskSchedulesAPITaskSchedulesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r TaskSchedulesAPITaskSchedulesListRequest) PulpIdIn(pulpIdIn []string) TaskSchedulesAPITaskSchedulesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r TaskSchedulesAPITaskSchedulesListRequest) Q(q string) TaskSchedulesAPITaskSchedulesListRequest {
	r.q = &q
	return r
}

// Filter results where task_name matches value
func (r TaskSchedulesAPITaskSchedulesListRequest) TaskName(taskName string) TaskSchedulesAPITaskSchedulesListRequest {
	r.taskName = &taskName
	return r
}

// Filter results where task_name contains value
func (r TaskSchedulesAPITaskSchedulesListRequest) TaskNameContains(taskNameContains string) TaskSchedulesAPITaskSchedulesListRequest {
	r.taskNameContains = &taskNameContains
	return r
}

// A list of fields to include in the response.
func (r TaskSchedulesAPITaskSchedulesListRequest) Fields(fields []string) TaskSchedulesAPITaskSchedulesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r TaskSchedulesAPITaskSchedulesListRequest) ExcludeFields(excludeFields []string) TaskSchedulesAPITaskSchedulesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r TaskSchedulesAPITaskSchedulesListRequest) Execute() (*PaginatedTaskScheduleResponseList, *http.Response, error) {
	return r.ApiService.TaskSchedulesListExecute(r)
}

/*
TaskSchedulesList List task schedules

ViewSet to monitor task schedules.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return TaskSchedulesAPITaskSchedulesListRequest
*/
func (a *TaskSchedulesAPIService) TaskSchedulesList(ctx context.Context, pulpDomain string) TaskSchedulesAPITaskSchedulesListRequest {
	return TaskSchedulesAPITaskSchedulesListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedTaskScheduleResponseList
func (a *TaskSchedulesAPIService) TaskSchedulesListExecute(r TaskSchedulesAPITaskSchedulesListRequest) (*PaginatedTaskScheduleResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedTaskScheduleResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesAPIService.TaskSchedulesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/task-schedules/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

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
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
	}
	if r.taskName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "task_name", r.taskName, "")
	}
	if r.taskNameContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "task_name__contains", r.taskNameContains, "")
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

type TaskSchedulesAPITaskSchedulesListRolesRequest struct {
	ctx context.Context
	ApiService *TaskSchedulesAPIService
	taskScheduleHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r TaskSchedulesAPITaskSchedulesListRolesRequest) Fields(fields []string) TaskSchedulesAPITaskSchedulesListRolesRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r TaskSchedulesAPITaskSchedulesListRolesRequest) ExcludeFields(excludeFields []string) TaskSchedulesAPITaskSchedulesListRolesRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r TaskSchedulesAPITaskSchedulesListRolesRequest) Execute() (*ObjectRolesResponse, *http.Response, error) {
	return r.ApiService.TaskSchedulesListRolesExecute(r)
}

/*
TaskSchedulesListRoles List roles

List roles assigned to this object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskScheduleHref
 @return TaskSchedulesAPITaskSchedulesListRolesRequest
*/
func (a *TaskSchedulesAPIService) TaskSchedulesListRoles(ctx context.Context, taskScheduleHref string) TaskSchedulesAPITaskSchedulesListRolesRequest {
	return TaskSchedulesAPITaskSchedulesListRolesRequest{
		ApiService: a,
		ctx: ctx,
		taskScheduleHref: taskScheduleHref,
	}
}

// Execute executes the request
//  @return ObjectRolesResponse
func (a *TaskSchedulesAPIService) TaskSchedulesListRolesExecute(r TaskSchedulesAPITaskSchedulesListRolesRequest) (*ObjectRolesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObjectRolesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesAPIService.TaskSchedulesListRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{task_schedule_href}list_roles/"
	localVarPath = strings.Replace(localVarPath, "{"+"task_schedule_href"+"}", url.PathEscape(parameterValueToString(r.taskScheduleHref, "taskScheduleHref")), -1)
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

type TaskSchedulesAPITaskSchedulesMyPermissionsRequest struct {
	ctx context.Context
	ApiService *TaskSchedulesAPIService
	taskScheduleHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r TaskSchedulesAPITaskSchedulesMyPermissionsRequest) Fields(fields []string) TaskSchedulesAPITaskSchedulesMyPermissionsRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r TaskSchedulesAPITaskSchedulesMyPermissionsRequest) ExcludeFields(excludeFields []string) TaskSchedulesAPITaskSchedulesMyPermissionsRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r TaskSchedulesAPITaskSchedulesMyPermissionsRequest) Execute() (*MyPermissionsResponse, *http.Response, error) {
	return r.ApiService.TaskSchedulesMyPermissionsExecute(r)
}

/*
TaskSchedulesMyPermissions List user permissions

List permissions available to the current user on this object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskScheduleHref
 @return TaskSchedulesAPITaskSchedulesMyPermissionsRequest
*/
func (a *TaskSchedulesAPIService) TaskSchedulesMyPermissions(ctx context.Context, taskScheduleHref string) TaskSchedulesAPITaskSchedulesMyPermissionsRequest {
	return TaskSchedulesAPITaskSchedulesMyPermissionsRequest{
		ApiService: a,
		ctx: ctx,
		taskScheduleHref: taskScheduleHref,
	}
}

// Execute executes the request
//  @return MyPermissionsResponse
func (a *TaskSchedulesAPIService) TaskSchedulesMyPermissionsExecute(r TaskSchedulesAPITaskSchedulesMyPermissionsRequest) (*MyPermissionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MyPermissionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesAPIService.TaskSchedulesMyPermissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{task_schedule_href}my_permissions/"
	localVarPath = strings.Replace(localVarPath, "{"+"task_schedule_href"+"}", url.PathEscape(parameterValueToString(r.taskScheduleHref, "taskScheduleHref")), -1)
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

type TaskSchedulesAPITaskSchedulesReadRequest struct {
	ctx context.Context
	ApiService *TaskSchedulesAPIService
	taskScheduleHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r TaskSchedulesAPITaskSchedulesReadRequest) Fields(fields []string) TaskSchedulesAPITaskSchedulesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r TaskSchedulesAPITaskSchedulesReadRequest) ExcludeFields(excludeFields []string) TaskSchedulesAPITaskSchedulesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r TaskSchedulesAPITaskSchedulesReadRequest) Execute() (*TaskScheduleResponse, *http.Response, error) {
	return r.ApiService.TaskSchedulesReadExecute(r)
}

/*
TaskSchedulesRead Inspect a task schedule

ViewSet to monitor task schedules.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskScheduleHref
 @return TaskSchedulesAPITaskSchedulesReadRequest
*/
func (a *TaskSchedulesAPIService) TaskSchedulesRead(ctx context.Context, taskScheduleHref string) TaskSchedulesAPITaskSchedulesReadRequest {
	return TaskSchedulesAPITaskSchedulesReadRequest{
		ApiService: a,
		ctx: ctx,
		taskScheduleHref: taskScheduleHref,
	}
}

// Execute executes the request
//  @return TaskScheduleResponse
func (a *TaskSchedulesAPIService) TaskSchedulesReadExecute(r TaskSchedulesAPITaskSchedulesReadRequest) (*TaskScheduleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskScheduleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesAPIService.TaskSchedulesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{task_schedule_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"task_schedule_href"+"}", url.PathEscape(parameterValueToString(r.taskScheduleHref, "taskScheduleHref")), -1)
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

type TaskSchedulesAPITaskSchedulesRemoveRoleRequest struct {
	ctx context.Context
	ApiService *TaskSchedulesAPIService
	taskScheduleHref string
	nestedRole *NestedRole
}

func (r TaskSchedulesAPITaskSchedulesRemoveRoleRequest) NestedRole(nestedRole NestedRole) TaskSchedulesAPITaskSchedulesRemoveRoleRequest {
	r.nestedRole = &nestedRole
	return r
}

func (r TaskSchedulesAPITaskSchedulesRemoveRoleRequest) Execute() (*NestedRoleResponse, *http.Response, error) {
	return r.ApiService.TaskSchedulesRemoveRoleExecute(r)
}

/*
TaskSchedulesRemoveRole Remove a role

Remove a role for this object from users/groups.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskScheduleHref
 @return TaskSchedulesAPITaskSchedulesRemoveRoleRequest
*/
func (a *TaskSchedulesAPIService) TaskSchedulesRemoveRole(ctx context.Context, taskScheduleHref string) TaskSchedulesAPITaskSchedulesRemoveRoleRequest {
	return TaskSchedulesAPITaskSchedulesRemoveRoleRequest{
		ApiService: a,
		ctx: ctx,
		taskScheduleHref: taskScheduleHref,
	}
}

// Execute executes the request
//  @return NestedRoleResponse
func (a *TaskSchedulesAPIService) TaskSchedulesRemoveRoleExecute(r TaskSchedulesAPITaskSchedulesRemoveRoleRequest) (*NestedRoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NestedRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesAPIService.TaskSchedulesRemoveRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{task_schedule_href}remove_role/"
	localVarPath = strings.Replace(localVarPath, "{"+"task_schedule_href"+"}", url.PathEscape(parameterValueToString(r.taskScheduleHref, "taskScheduleHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.nestedRole == nil {
		return localVarReturnValue, nil, reportError("nestedRole is required and must be specified")
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
	localVarPostBody = r.nestedRole
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
