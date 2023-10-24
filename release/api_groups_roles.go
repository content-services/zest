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


// GroupsRolesAPIService GroupsRolesAPI service
type GroupsRolesAPIService service

type GroupsRolesAPIGroupsRolesCreateRequest struct {
	ctx context.Context
	ApiService *GroupsRolesAPIService
	groupHref string
	groupRole *GroupRole
}

func (r GroupsRolesAPIGroupsRolesCreateRequest) GroupRole(groupRole GroupRole) GroupsRolesAPIGroupsRolesCreateRequest {
	r.groupRole = &groupRole
	return r
}

func (r GroupsRolesAPIGroupsRolesCreateRequest) Execute() (*GroupRoleResponse, *http.Response, error) {
	return r.ApiService.GroupsRolesCreateExecute(r)
}

/*
GroupsRolesCreate Create a group role

ViewSet for GroupRole.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupHref
 @return GroupsRolesAPIGroupsRolesCreateRequest
*/
func (a *GroupsRolesAPIService) GroupsRolesCreate(ctx context.Context, groupHref string) GroupsRolesAPIGroupsRolesCreateRequest {
	return GroupsRolesAPIGroupsRolesCreateRequest{
		ApiService: a,
		ctx: ctx,
		groupHref: groupHref,
	}
}

// Execute executes the request
//  @return GroupRoleResponse
func (a *GroupsRolesAPIService) GroupsRolesCreateExecute(r GroupsRolesAPIGroupsRolesCreateRequest) (*GroupRoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsRolesAPIService.GroupsRolesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{group_href}roles/"
	localVarPath = strings.Replace(localVarPath, "{"+"group_href"+"}", url.PathEscape(parameterValueToString(r.groupHref, "groupHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupRole == nil {
		return localVarReturnValue, nil, reportError("groupRole is required and must be specified")
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
	localVarPostBody = r.groupRole
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

type GroupsRolesAPIGroupsRolesDeleteRequest struct {
	ctx context.Context
	ApiService *GroupsRolesAPIService
	groupsGroupRoleHref string
}

func (r GroupsRolesAPIGroupsRolesDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.GroupsRolesDeleteExecute(r)
}

/*
GroupsRolesDelete Delete a group role

ViewSet for GroupRole.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupsGroupRoleHref
 @return GroupsRolesAPIGroupsRolesDeleteRequest
*/
func (a *GroupsRolesAPIService) GroupsRolesDelete(ctx context.Context, groupsGroupRoleHref string) GroupsRolesAPIGroupsRolesDeleteRequest {
	return GroupsRolesAPIGroupsRolesDeleteRequest{
		ApiService: a,
		ctx: ctx,
		groupsGroupRoleHref: groupsGroupRoleHref,
	}
}

// Execute executes the request
func (a *GroupsRolesAPIService) GroupsRolesDeleteExecute(r GroupsRolesAPIGroupsRolesDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsRolesAPIService.GroupsRolesDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{groups_group_role_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"groups_group_role_href"+"}", url.PathEscape(parameterValueToString(r.groupsGroupRoleHref, "groupsGroupRoleHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type GroupsRolesAPIGroupsRolesListRequest struct {
	ctx context.Context
	ApiService *GroupsRolesAPIService
	groupHref string
	contentObject *string
	domain *string
	limit *int32
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	q *string
	role *string
	roleContains *string
	roleIcontains *string
	roleIn *[]string
	roleStartswith *string
	fields *[]string
	excludeFields *[]string
}

// content_object
func (r GroupsRolesAPIGroupsRolesListRequest) ContentObject(contentObject string) GroupsRolesAPIGroupsRolesListRequest {
	r.contentObject = &contentObject
	return r
}

// Foreign Key referenced by HREF
func (r GroupsRolesAPIGroupsRolesListRequest) Domain(domain string) GroupsRolesAPIGroupsRolesListRequest {
	r.domain = &domain
	return r
}

// Number of results to return per page.
func (r GroupsRolesAPIGroupsRolesListRequest) Limit(limit int32) GroupsRolesAPIGroupsRolesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r GroupsRolesAPIGroupsRolesListRequest) Offset(offset int32) GroupsRolesAPIGroupsRolesListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;role&#x60; - Role * &#x60;-role&#x60; - Role (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r GroupsRolesAPIGroupsRolesListRequest) Ordering(ordering []string) GroupsRolesAPIGroupsRolesListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r GroupsRolesAPIGroupsRolesListRequest) PulpHrefIn(pulpHrefIn []string) GroupsRolesAPIGroupsRolesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r GroupsRolesAPIGroupsRolesListRequest) PulpIdIn(pulpIdIn []string) GroupsRolesAPIGroupsRolesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r GroupsRolesAPIGroupsRolesListRequest) Q(q string) GroupsRolesAPIGroupsRolesListRequest {
	r.q = &q
	return r
}

func (r GroupsRolesAPIGroupsRolesListRequest) Role(role string) GroupsRolesAPIGroupsRolesListRequest {
	r.role = &role
	return r
}

func (r GroupsRolesAPIGroupsRolesListRequest) RoleContains(roleContains string) GroupsRolesAPIGroupsRolesListRequest {
	r.roleContains = &roleContains
	return r
}

func (r GroupsRolesAPIGroupsRolesListRequest) RoleIcontains(roleIcontains string) GroupsRolesAPIGroupsRolesListRequest {
	r.roleIcontains = &roleIcontains
	return r
}

// Multiple values may be separated by commas.
func (r GroupsRolesAPIGroupsRolesListRequest) RoleIn(roleIn []string) GroupsRolesAPIGroupsRolesListRequest {
	r.roleIn = &roleIn
	return r
}

func (r GroupsRolesAPIGroupsRolesListRequest) RoleStartswith(roleStartswith string) GroupsRolesAPIGroupsRolesListRequest {
	r.roleStartswith = &roleStartswith
	return r
}

// A list of fields to include in the response.
func (r GroupsRolesAPIGroupsRolesListRequest) Fields(fields []string) GroupsRolesAPIGroupsRolesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r GroupsRolesAPIGroupsRolesListRequest) ExcludeFields(excludeFields []string) GroupsRolesAPIGroupsRolesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r GroupsRolesAPIGroupsRolesListRequest) Execute() (*PaginatedGroupRoleResponseList, *http.Response, error) {
	return r.ApiService.GroupsRolesListExecute(r)
}

/*
GroupsRolesList List group roles

ViewSet for GroupRole.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupHref
 @return GroupsRolesAPIGroupsRolesListRequest
*/
func (a *GroupsRolesAPIService) GroupsRolesList(ctx context.Context, groupHref string) GroupsRolesAPIGroupsRolesListRequest {
	return GroupsRolesAPIGroupsRolesListRequest{
		ApiService: a,
		ctx: ctx,
		groupHref: groupHref,
	}
}

// Execute executes the request
//  @return PaginatedGroupRoleResponseList
func (a *GroupsRolesAPIService) GroupsRolesListExecute(r GroupsRolesAPIGroupsRolesListRequest) (*PaginatedGroupRoleResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedGroupRoleResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsRolesAPIService.GroupsRolesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{group_href}roles/"
	localVarPath = strings.Replace(localVarPath, "{"+"group_href"+"}", url.PathEscape(parameterValueToString(r.groupHref, "groupHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.contentObject != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content_object", r.contentObject, "")
	}
	if r.domain != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "domain", r.domain, "")
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
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
	}
	if r.role != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "role", r.role, "")
	}
	if r.roleContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "role__contains", r.roleContains, "")
	}
	if r.roleIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "role__icontains", r.roleIcontains, "")
	}
	if r.roleIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "role__in", r.roleIn, "csv")
	}
	if r.roleStartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "role__startswith", r.roleStartswith, "")
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

type GroupsRolesAPIGroupsRolesReadRequest struct {
	ctx context.Context
	ApiService *GroupsRolesAPIService
	groupsGroupRoleHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r GroupsRolesAPIGroupsRolesReadRequest) Fields(fields []string) GroupsRolesAPIGroupsRolesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r GroupsRolesAPIGroupsRolesReadRequest) ExcludeFields(excludeFields []string) GroupsRolesAPIGroupsRolesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r GroupsRolesAPIGroupsRolesReadRequest) Execute() (*GroupRoleResponse, *http.Response, error) {
	return r.ApiService.GroupsRolesReadExecute(r)
}

/*
GroupsRolesRead Inspect a group role

ViewSet for GroupRole.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupsGroupRoleHref
 @return GroupsRolesAPIGroupsRolesReadRequest
*/
func (a *GroupsRolesAPIService) GroupsRolesRead(ctx context.Context, groupsGroupRoleHref string) GroupsRolesAPIGroupsRolesReadRequest {
	return GroupsRolesAPIGroupsRolesReadRequest{
		ApiService: a,
		ctx: ctx,
		groupsGroupRoleHref: groupsGroupRoleHref,
	}
}

// Execute executes the request
//  @return GroupRoleResponse
func (a *GroupsRolesAPIService) GroupsRolesReadExecute(r GroupsRolesAPIGroupsRolesReadRequest) (*GroupRoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsRolesAPIService.GroupsRolesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{groups_group_role_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"groups_group_role_href"+"}", url.PathEscape(parameterValueToString(r.groupsGroupRoleHref, "groupsGroupRoleHref")), -1)
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
