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


// RolesAPIService RolesAPI service
type RolesAPIService service

type RolesAPIRolesCreateRequest struct {
	ctx context.Context
	ApiService *RolesAPIService
	pulpDomain string
	role *Role
}

func (r RolesAPIRolesCreateRequest) Role(role Role) RolesAPIRolesCreateRequest {
	r.role = &role
	return r
}

func (r RolesAPIRolesCreateRequest) Execute() (*RoleResponse, *http.Response, error) {
	return r.ApiService.RolesCreateExecute(r)
}

/*
RolesCreate Create a role

ViewSet for Role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return RolesAPIRolesCreateRequest
*/
func (a *RolesAPIService) RolesCreate(ctx context.Context, pulpDomain string) RolesAPIRolesCreateRequest {
	return RolesAPIRolesCreateRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return RoleResponse
func (a *RolesAPIService) RolesCreateExecute(r RolesAPIRolesCreateRequest) (*RoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesAPIService.RolesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/{pulp_domain}/api/v3/roles/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
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

type RolesAPIRolesDeleteRequest struct {
	ctx context.Context
	ApiService *RolesAPIService
	roleHref string
}

func (r RolesAPIRolesDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.RolesDeleteExecute(r)
}

/*
RolesDelete Delete a role

ViewSet for Role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleHref
 @return RolesAPIRolesDeleteRequest
*/
func (a *RolesAPIService) RolesDelete(ctx context.Context, roleHref string) RolesAPIRolesDeleteRequest {
	return RolesAPIRolesDeleteRequest{
		ApiService: a,
		ctx: ctx,
		roleHref: roleHref,
	}
}

// Execute executes the request
func (a *RolesAPIService) RolesDeleteExecute(r RolesAPIRolesDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesAPIService.RolesDelete")
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

type RolesAPIRolesListRequest struct {
	ctx context.Context
	ApiService *RolesAPIService
	pulpDomain string
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
	q *string
	fields *[]string
	excludeFields *[]string
}

// Filter roles that have any of the permissions in the list.
func (r RolesAPIRolesListRequest) ContainsPermission(containsPermission []string) RolesAPIRolesListRequest {
	r.containsPermission = &containsPermission
	return r
}

// Filter results where description matches value
func (r RolesAPIRolesListRequest) Description(description string) RolesAPIRolesListRequest {
	r.description = &description
	return r
}

// Filter results where description contains value
func (r RolesAPIRolesListRequest) DescriptionContains(descriptionContains string) RolesAPIRolesListRequest {
	r.descriptionContains = &descriptionContains
	return r
}

// Filter results where description contains value
func (r RolesAPIRolesListRequest) DescriptionIcontains(descriptionIcontains string) RolesAPIRolesListRequest {
	r.descriptionIcontains = &descriptionIcontains
	return r
}

// Filter results where description matches value
func (r RolesAPIRolesListRequest) DescriptionIexact(descriptionIexact string) RolesAPIRolesListRequest {
	r.descriptionIexact = &descriptionIexact
	return r
}

// Filter roles that only have permissions for the specified object HREF.
func (r RolesAPIRolesListRequest) ForObjectType(forObjectType string) RolesAPIRolesListRequest {
	r.forObjectType = &forObjectType
	return r
}

// Number of results to return per page.
func (r RolesAPIRolesListRequest) Limit(limit int32) RolesAPIRolesListRequest {
	r.limit = &limit
	return r
}

// Filter results where locked matches value
func (r RolesAPIRolesListRequest) Locked(locked bool) RolesAPIRolesListRequest {
	r.locked = &locked
	return r
}

// Filter results where name matches value
func (r RolesAPIRolesListRequest) Name(name string) RolesAPIRolesListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r RolesAPIRolesListRequest) NameContains(nameContains string) RolesAPIRolesListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r RolesAPIRolesListRequest) NameIcontains(nameIcontains string) RolesAPIRolesListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name matches value
func (r RolesAPIRolesListRequest) NameIexact(nameIexact string) RolesAPIRolesListRequest {
	r.nameIexact = &nameIexact
	return r
}

// Filter results where name is in a comma-separated list of values
func (r RolesAPIRolesListRequest) NameIn(nameIn []string) RolesAPIRolesListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name matches regex value
func (r RolesAPIRolesListRequest) NameIregex(nameIregex string) RolesAPIRolesListRequest {
	r.nameIregex = &nameIregex
	return r
}

// Filter results where name starts with value
func (r RolesAPIRolesListRequest) NameIstartswith(nameIstartswith string) RolesAPIRolesListRequest {
	r.nameIstartswith = &nameIstartswith
	return r
}

// Filter results where name matches regex value
func (r RolesAPIRolesListRequest) NameRegex(nameRegex string) RolesAPIRolesListRequest {
	r.nameRegex = &nameRegex
	return r
}

// Filter results where name starts with value
func (r RolesAPIRolesListRequest) NameStartswith(nameStartswith string) RolesAPIRolesListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r RolesAPIRolesListRequest) Offset(offset int32) RolesAPIRolesListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;locked&#x60; - Locked * &#x60;-locked&#x60; - Locked (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r RolesAPIRolesListRequest) Ordering(ordering []string) RolesAPIRolesListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r RolesAPIRolesListRequest) PulpHrefIn(pulpHrefIn []string) RolesAPIRolesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r RolesAPIRolesListRequest) PulpIdIn(pulpIdIn []string) RolesAPIRolesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r RolesAPIRolesListRequest) Q(q string) RolesAPIRolesListRequest {
	r.q = &q
	return r
}

// A list of fields to include in the response.
func (r RolesAPIRolesListRequest) Fields(fields []string) RolesAPIRolesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RolesAPIRolesListRequest) ExcludeFields(excludeFields []string) RolesAPIRolesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RolesAPIRolesListRequest) Execute() (*PaginatedRoleResponseList, *http.Response, error) {
	return r.ApiService.RolesListExecute(r)
}

/*
RolesList List roles

ViewSet for Role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return RolesAPIRolesListRequest
*/
func (a *RolesAPIService) RolesList(ctx context.Context, pulpDomain string) RolesAPIRolesListRequest {
	return RolesAPIRolesListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedRoleResponseList
func (a *RolesAPIService) RolesListExecute(r RolesAPIRolesListRequest) (*PaginatedRoleResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRoleResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesAPIService.RolesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/{pulp_domain}/api/v3/roles/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

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
	if r.nameIexact != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__iexact", r.nameIexact, "")
	}
	if r.nameIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__in", r.nameIn, "csv")
	}
	if r.nameIregex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__iregex", r.nameIregex, "")
	}
	if r.nameIstartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__istartswith", r.nameIstartswith, "")
	}
	if r.nameRegex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__regex", r.nameRegex, "")
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
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
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

type RolesAPIRolesPartialUpdateRequest struct {
	ctx context.Context
	ApiService *RolesAPIService
	roleHref string
	patchedRole *PatchedRole
}

func (r RolesAPIRolesPartialUpdateRequest) PatchedRole(patchedRole PatchedRole) RolesAPIRolesPartialUpdateRequest {
	r.patchedRole = &patchedRole
	return r
}

func (r RolesAPIRolesPartialUpdateRequest) Execute() (*RoleResponse, *http.Response, error) {
	return r.ApiService.RolesPartialUpdateExecute(r)
}

/*
RolesPartialUpdate Update a role

ViewSet for Role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleHref
 @return RolesAPIRolesPartialUpdateRequest
*/
func (a *RolesAPIService) RolesPartialUpdate(ctx context.Context, roleHref string) RolesAPIRolesPartialUpdateRequest {
	return RolesAPIRolesPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		roleHref: roleHref,
	}
}

// Execute executes the request
//  @return RoleResponse
func (a *RolesAPIService) RolesPartialUpdateExecute(r RolesAPIRolesPartialUpdateRequest) (*RoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesAPIService.RolesPartialUpdate")
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

type RolesAPIRolesReadRequest struct {
	ctx context.Context
	ApiService *RolesAPIService
	roleHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r RolesAPIRolesReadRequest) Fields(fields []string) RolesAPIRolesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RolesAPIRolesReadRequest) ExcludeFields(excludeFields []string) RolesAPIRolesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RolesAPIRolesReadRequest) Execute() (*RoleResponse, *http.Response, error) {
	return r.ApiService.RolesReadExecute(r)
}

/*
RolesRead Inspect a role

ViewSet for Role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleHref
 @return RolesAPIRolesReadRequest
*/
func (a *RolesAPIService) RolesRead(ctx context.Context, roleHref string) RolesAPIRolesReadRequest {
	return RolesAPIRolesReadRequest{
		ApiService: a,
		ctx: ctx,
		roleHref: roleHref,
	}
}

// Execute executes the request
//  @return RoleResponse
func (a *RolesAPIService) RolesReadExecute(r RolesAPIRolesReadRequest) (*RoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesAPIService.RolesRead")
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

type RolesAPIRolesUpdateRequest struct {
	ctx context.Context
	ApiService *RolesAPIService
	roleHref string
	role *Role
}

func (r RolesAPIRolesUpdateRequest) Role(role Role) RolesAPIRolesUpdateRequest {
	r.role = &role
	return r
}

func (r RolesAPIRolesUpdateRequest) Execute() (*RoleResponse, *http.Response, error) {
	return r.ApiService.RolesUpdateExecute(r)
}

/*
RolesUpdate Update a role

ViewSet for Role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleHref
 @return RolesAPIRolesUpdateRequest
*/
func (a *RolesAPIService) RolesUpdate(ctx context.Context, roleHref string) RolesAPIRolesUpdateRequest {
	return RolesAPIRolesUpdateRequest{
		ApiService: a,
		ctx: ctx,
		roleHref: roleHref,
	}
}

// Execute executes the request
//  @return RoleResponse
func (a *RolesAPIService) RolesUpdateExecute(r RolesAPIRolesUpdateRequest) (*RoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesAPIService.RolesUpdate")
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
