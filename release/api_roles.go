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
	"strings"
	"reflect"
)


// RolesApiService RolesApi service
type RolesApiService service

type RolesApiRolesCreateRequest struct {
	ctx context.Context
	ApiService *RolesApiService
	role *Role
}

func (r RolesApiRolesCreateRequest) Role(role Role) RolesApiRolesCreateRequest {
	r.role = &role
	return r
}

func (r RolesApiRolesCreateRequest) Execute() (*RoleResponse, *http.Response, error) {
	return r.ApiService.RolesCreateExecute(r)
}

/*
RolesCreate Create a role

ViewSet for Role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RolesApiRolesCreateRequest
*/
func (a *RolesApiService) RolesCreate(ctx context.Context) RolesApiRolesCreateRequest {
	return RolesApiRolesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RoleResponse
func (a *RolesApiService) RolesCreateExecute(r RolesApiRolesCreateRequest) (*RoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesApiService.RolesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/roles/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.role == nil {
		return localVarReturnValue, nil, reportError("role is required and must be specified")
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
	localVarPostBody = r.role
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

type RolesApiRolesDeleteRequest struct {
	ctx context.Context
	ApiService *RolesApiService
	roleHref string
}

func (r RolesApiRolesDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.RolesDeleteExecute(r)
}

/*
RolesDelete Delete a role

ViewSet for Role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleHref
 @return RolesApiRolesDeleteRequest
*/
func (a *RolesApiService) RolesDelete(ctx context.Context, roleHref string) RolesApiRolesDeleteRequest {
	return RolesApiRolesDeleteRequest{
		ApiService: a,
		ctx: ctx,
		roleHref: roleHref,
	}
}

// Execute executes the request
func (a *RolesApiService) RolesDeleteExecute(r RolesApiRolesDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesApiService.RolesDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{role_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"role_href"+"}", url.PathEscape(parameterValueToString(r.roleHref, "roleHref")), -1)
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

type RolesApiRolesListRequest struct {
	ctx context.Context
	ApiService *RolesApiService
	containsPermission *[]string
	description *string
	descriptionContains *string
	descriptionIcontains *string
	descriptionIexact *string
	forObjectType *string
	limit *int32
	locked *bool
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
	offset *int32
	ordering *[]string
	fields *[]string
	excludeFields *[]string
}

// Filter roles that have any of the permissions in the list.
func (r RolesApiRolesListRequest) ContainsPermission(containsPermission []string) RolesApiRolesListRequest {
	r.containsPermission = &containsPermission
	return r
}

// Filter results where description matches value
func (r RolesApiRolesListRequest) Description(description string) RolesApiRolesListRequest {
	r.description = &description
	return r
}

// Filter results where description contains value
func (r RolesApiRolesListRequest) DescriptionContains(descriptionContains string) RolesApiRolesListRequest {
	r.descriptionContains = &descriptionContains
	return r
}

// Filter results where description contains value
func (r RolesApiRolesListRequest) DescriptionIcontains(descriptionIcontains string) RolesApiRolesListRequest {
	r.descriptionIcontains = &descriptionIcontains
	return r
}

// Filter results where description matches value
func (r RolesApiRolesListRequest) DescriptionIexact(descriptionIexact string) RolesApiRolesListRequest {
	r.descriptionIexact = &descriptionIexact
	return r
}

// Filter roles that only have permissions for the specified object HREF.
func (r RolesApiRolesListRequest) ForObjectType(forObjectType string) RolesApiRolesListRequest {
	r.forObjectType = &forObjectType
	return r
}

// Number of results to return per page.
func (r RolesApiRolesListRequest) Limit(limit int32) RolesApiRolesListRequest {
	r.limit = &limit
	return r
}

// Filter results where locked matches value
func (r RolesApiRolesListRequest) Locked(locked bool) RolesApiRolesListRequest {
	r.locked = &locked
	return r
}

// Filter results where name matches value
func (r RolesApiRolesListRequest) Name(name string) RolesApiRolesListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r RolesApiRolesListRequest) NameContains(nameContains string) RolesApiRolesListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r RolesApiRolesListRequest) NameIcontains(nameIcontains string) RolesApiRolesListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r RolesApiRolesListRequest) NameIn(nameIn []string) RolesApiRolesListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r RolesApiRolesListRequest) NameStartswith(nameStartswith string) RolesApiRolesListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r RolesApiRolesListRequest) Offset(offset int32) RolesApiRolesListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r RolesApiRolesListRequest) Ordering(ordering []string) RolesApiRolesListRequest {
	r.ordering = &ordering
	return r
}

// A list of fields to include in the response.
func (r RolesApiRolesListRequest) Fields(fields []string) RolesApiRolesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RolesApiRolesListRequest) ExcludeFields(excludeFields []string) RolesApiRolesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RolesApiRolesListRequest) Execute() (*PaginatedRoleResponseList, *http.Response, error) {
	return r.ApiService.RolesListExecute(r)
}

/*
RolesList List roles

ViewSet for Role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RolesApiRolesListRequest
*/
func (a *RolesApiService) RolesList(ctx context.Context) RolesApiRolesListRequest {
	return RolesApiRolesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedRoleResponseList
func (a *RolesApiService) RolesListExecute(r RolesApiRolesListRequest) (*PaginatedRoleResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRoleResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesApiService.RolesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/roles/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.containsPermission != nil {
		t := *r.containsPermission
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "contains_permission", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "contains_permission", t, "multi")
		}
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "description", r.description, "")
	}
	if r.descriptionContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "description__contains", r.descriptionContains, "")
	}
	if r.descriptionIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "description__icontains", r.descriptionIcontains, "")
	}
	if r.descriptionIexact != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "description__iexact", r.descriptionIexact, "")
	}
	if r.forObjectType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "for_object_type", r.forObjectType, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.locked != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "locked", r.locked, "")
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

type RolesApiRolesPartialUpdateRequest struct {
	ctx context.Context
	ApiService *RolesApiService
	roleHref string
	patchedRole *PatchedRole
}

func (r RolesApiRolesPartialUpdateRequest) PatchedRole(patchedRole PatchedRole) RolesApiRolesPartialUpdateRequest {
	r.patchedRole = &patchedRole
	return r
}

func (r RolesApiRolesPartialUpdateRequest) Execute() (*RoleResponse, *http.Response, error) {
	return r.ApiService.RolesPartialUpdateExecute(r)
}

/*
RolesPartialUpdate Update a role

ViewSet for Role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleHref
 @return RolesApiRolesPartialUpdateRequest
*/
func (a *RolesApiService) RolesPartialUpdate(ctx context.Context, roleHref string) RolesApiRolesPartialUpdateRequest {
	return RolesApiRolesPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		roleHref: roleHref,
	}
}

// Execute executes the request
//  @return RoleResponse
func (a *RolesApiService) RolesPartialUpdateExecute(r RolesApiRolesPartialUpdateRequest) (*RoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesApiService.RolesPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{role_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"role_href"+"}", url.PathEscape(parameterValueToString(r.roleHref, "roleHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedRole == nil {
		return localVarReturnValue, nil, reportError("patchedRole is required and must be specified")
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
	localVarPostBody = r.patchedRole
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

type RolesApiRolesReadRequest struct {
	ctx context.Context
	ApiService *RolesApiService
	roleHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r RolesApiRolesReadRequest) Fields(fields []string) RolesApiRolesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RolesApiRolesReadRequest) ExcludeFields(excludeFields []string) RolesApiRolesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RolesApiRolesReadRequest) Execute() (*RoleResponse, *http.Response, error) {
	return r.ApiService.RolesReadExecute(r)
}

/*
RolesRead Inspect a role

ViewSet for Role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleHref
 @return RolesApiRolesReadRequest
*/
func (a *RolesApiService) RolesRead(ctx context.Context, roleHref string) RolesApiRolesReadRequest {
	return RolesApiRolesReadRequest{
		ApiService: a,
		ctx: ctx,
		roleHref: roleHref,
	}
}

// Execute executes the request
//  @return RoleResponse
func (a *RolesApiService) RolesReadExecute(r RolesApiRolesReadRequest) (*RoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesApiService.RolesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{role_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"role_href"+"}", url.PathEscape(parameterValueToString(r.roleHref, "roleHref")), -1)
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

type RolesApiRolesUpdateRequest struct {
	ctx context.Context
	ApiService *RolesApiService
	roleHref string
	role *Role
}

func (r RolesApiRolesUpdateRequest) Role(role Role) RolesApiRolesUpdateRequest {
	r.role = &role
	return r
}

func (r RolesApiRolesUpdateRequest) Execute() (*RoleResponse, *http.Response, error) {
	return r.ApiService.RolesUpdateExecute(r)
}

/*
RolesUpdate Update a role

ViewSet for Role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleHref
 @return RolesApiRolesUpdateRequest
*/
func (a *RolesApiService) RolesUpdate(ctx context.Context, roleHref string) RolesApiRolesUpdateRequest {
	return RolesApiRolesUpdateRequest{
		ApiService: a,
		ctx: ctx,
		roleHref: roleHref,
	}
}

// Execute executes the request
//  @return RoleResponse
func (a *RolesApiService) RolesUpdateExecute(r RolesApiRolesUpdateRequest) (*RoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesApiService.RolesUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{role_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"role_href"+"}", url.PathEscape(parameterValueToString(r.roleHref, "roleHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.role == nil {
		return localVarReturnValue, nil, reportError("role is required and must be specified")
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
	localVarPostBody = r.role
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
