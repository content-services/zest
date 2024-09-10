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
	"time"
	"reflect"
)


// WorkersAPIService WorkersAPI service
type WorkersAPIService service

type WorkersAPIWorkersListRequest struct {
	ctx context.Context
	ApiService *WorkersAPIService
	pulpDomain string
	lastHeartbeat *time.Time
	lastHeartbeatGt *time.Time
	lastHeartbeatGte *time.Time
	lastHeartbeatLt *time.Time
	lastHeartbeatLte *time.Time
	lastHeartbeatRange *[]time.Time
	limit *int32
	missing *bool
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
	online *bool
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	q *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where last_heartbeat matches value
func (r WorkersAPIWorkersListRequest) LastHeartbeat(lastHeartbeat time.Time) WorkersAPIWorkersListRequest {
	r.lastHeartbeat = &lastHeartbeat
	return r
}

// Filter results where last_heartbeat is greater than value
func (r WorkersAPIWorkersListRequest) LastHeartbeatGt(lastHeartbeatGt time.Time) WorkersAPIWorkersListRequest {
	r.lastHeartbeatGt = &lastHeartbeatGt
	return r
}

// Filter results where last_heartbeat is greater than or equal to value
func (r WorkersAPIWorkersListRequest) LastHeartbeatGte(lastHeartbeatGte time.Time) WorkersAPIWorkersListRequest {
	r.lastHeartbeatGte = &lastHeartbeatGte
	return r
}

// Filter results where last_heartbeat is less than value
func (r WorkersAPIWorkersListRequest) LastHeartbeatLt(lastHeartbeatLt time.Time) WorkersAPIWorkersListRequest {
	r.lastHeartbeatLt = &lastHeartbeatLt
	return r
}

// Filter results where last_heartbeat is less than or equal to value
func (r WorkersAPIWorkersListRequest) LastHeartbeatLte(lastHeartbeatLte time.Time) WorkersAPIWorkersListRequest {
	r.lastHeartbeatLte = &lastHeartbeatLte
	return r
}

// Filter results where last_heartbeat is between two comma separated values
func (r WorkersAPIWorkersListRequest) LastHeartbeatRange(lastHeartbeatRange []time.Time) WorkersAPIWorkersListRequest {
	r.lastHeartbeatRange = &lastHeartbeatRange
	return r
}

// Number of results to return per page.
func (r WorkersAPIWorkersListRequest) Limit(limit int32) WorkersAPIWorkersListRequest {
	r.limit = &limit
	return r
}

func (r WorkersAPIWorkersListRequest) Missing(missing bool) WorkersAPIWorkersListRequest {
	r.missing = &missing
	return r
}

// Filter results where name matches value
func (r WorkersAPIWorkersListRequest) Name(name string) WorkersAPIWorkersListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r WorkersAPIWorkersListRequest) NameContains(nameContains string) WorkersAPIWorkersListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r WorkersAPIWorkersListRequest) NameIcontains(nameIcontains string) WorkersAPIWorkersListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name matches value
func (r WorkersAPIWorkersListRequest) NameIexact(nameIexact string) WorkersAPIWorkersListRequest {
	r.nameIexact = &nameIexact
	return r
}

// Filter results where name is in a comma-separated list of values
func (r WorkersAPIWorkersListRequest) NameIn(nameIn []string) WorkersAPIWorkersListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name matches regex value
func (r WorkersAPIWorkersListRequest) NameIregex(nameIregex string) WorkersAPIWorkersListRequest {
	r.nameIregex = &nameIregex
	return r
}

// Filter results where name starts with value
func (r WorkersAPIWorkersListRequest) NameIstartswith(nameIstartswith string) WorkersAPIWorkersListRequest {
	r.nameIstartswith = &nameIstartswith
	return r
}

// Filter results where name matches regex value
func (r WorkersAPIWorkersListRequest) NameRegex(nameRegex string) WorkersAPIWorkersListRequest {
	r.nameRegex = &nameRegex
	return r
}

// Filter results where name starts with value
func (r WorkersAPIWorkersListRequest) NameStartswith(nameStartswith string) WorkersAPIWorkersListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r WorkersAPIWorkersListRequest) Offset(offset int32) WorkersAPIWorkersListRequest {
	r.offset = &offset
	return r
}

func (r WorkersAPIWorkersListRequest) Online(online bool) WorkersAPIWorkersListRequest {
	r.online = &online
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;last_heartbeat&#x60; - Last heartbeat* &#x60;-last_heartbeat&#x60; - Last heartbeat (descending)* &#x60;versions&#x60; - Versions* &#x60;-versions&#x60; - Versions (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r WorkersAPIWorkersListRequest) Ordering(ordering []string) WorkersAPIWorkersListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r WorkersAPIWorkersListRequest) PulpHrefIn(pulpHrefIn []string) WorkersAPIWorkersListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r WorkersAPIWorkersListRequest) PulpIdIn(pulpIdIn []string) WorkersAPIWorkersListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r WorkersAPIWorkersListRequest) Q(q string) WorkersAPIWorkersListRequest {
	r.q = &q
	return r
}

// A list of fields to include in the response.
func (r WorkersAPIWorkersListRequest) Fields(fields []string) WorkersAPIWorkersListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r WorkersAPIWorkersListRequest) ExcludeFields(excludeFields []string) WorkersAPIWorkersListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r WorkersAPIWorkersListRequest) Execute() (*PaginatedWorkerResponseList, *http.Response, error) {
	return r.ApiService.WorkersListExecute(r)
}

/*
WorkersList List workers

A customized named ModelViewSet that knows how to register itself with the Pulp API router.This viewset is discoverable by its name."Normal" Django Models and Master/Detail models are supported by the ``register_with`` method.Attributes:    lookup_field (str): The name of the field by which an object should be looked up, in        addition to any parent lookups if this ViewSet is nested. Defaults to 'pk'    endpoint_name (str): The name of the final path segment that should identify the ViewSet's        collection endpoint.    nest_prefix (str): Optional prefix under which this ViewSet should be nested. This must        correspond to the "parent_prefix" of a router with rest_framework_nested.NestedMixin.        None indicates this ViewSet should not be nested.    parent_lookup_kwargs (dict): Optional mapping of key names that would appear in self.kwargs        to django model filter expressions that can be used with the corresponding value from        self.kwargs, used only by a nested ViewSet to filter based on the parent object's        identity.    schema (DefaultSchema): The schema class to use by default in a viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return WorkersAPIWorkersListRequest
*/
func (a *WorkersAPIService) WorkersList(ctx context.Context, pulpDomain string) WorkersAPIWorkersListRequest {
	return WorkersAPIWorkersListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedWorkerResponseList
func (a *WorkersAPIService) WorkersListExecute(r WorkersAPIWorkersListRequest) (*PaginatedWorkerResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedWorkerResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkersAPIService.WorkersList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/workers/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.lastHeartbeat != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_heartbeat", r.lastHeartbeat, "form", "")
	}
	if r.lastHeartbeatGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_heartbeat__gt", r.lastHeartbeatGt, "form", "")
	}
	if r.lastHeartbeatGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_heartbeat__gte", r.lastHeartbeatGte, "form", "")
	}
	if r.lastHeartbeatLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_heartbeat__lt", r.lastHeartbeatLt, "form", "")
	}
	if r.lastHeartbeatLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_heartbeat__lte", r.lastHeartbeatLte, "form", "")
	}
	if r.lastHeartbeatRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_heartbeat__range", r.lastHeartbeatRange, "form", "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.missing != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "missing", r.missing, "form", "")
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
	if r.online != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "online", r.online, "form", "")
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
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
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

type WorkersAPIWorkersReadRequest struct {
	ctx context.Context
	ApiService *WorkersAPIService
	workerHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r WorkersAPIWorkersReadRequest) Fields(fields []string) WorkersAPIWorkersReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r WorkersAPIWorkersReadRequest) ExcludeFields(excludeFields []string) WorkersAPIWorkersReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r WorkersAPIWorkersReadRequest) Execute() (*WorkerResponse, *http.Response, error) {
	return r.ApiService.WorkersReadExecute(r)
}

/*
WorkersRead Inspect a worker

A customized named ModelViewSet that knows how to register itself with the Pulp API router.This viewset is discoverable by its name."Normal" Django Models and Master/Detail models are supported by the ``register_with`` method.Attributes:    lookup_field (str): The name of the field by which an object should be looked up, in        addition to any parent lookups if this ViewSet is nested. Defaults to 'pk'    endpoint_name (str): The name of the final path segment that should identify the ViewSet's        collection endpoint.    nest_prefix (str): Optional prefix under which this ViewSet should be nested. This must        correspond to the "parent_prefix" of a router with rest_framework_nested.NestedMixin.        None indicates this ViewSet should not be nested.    parent_lookup_kwargs (dict): Optional mapping of key names that would appear in self.kwargs        to django model filter expressions that can be used with the corresponding value from        self.kwargs, used only by a nested ViewSet to filter based on the parent object's        identity.    schema (DefaultSchema): The schema class to use by default in a viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workerHref
 @return WorkersAPIWorkersReadRequest
*/
func (a *WorkersAPIService) WorkersRead(ctx context.Context, workerHref string) WorkersAPIWorkersReadRequest {
	return WorkersAPIWorkersReadRequest{
		ApiService: a,
		ctx: ctx,
		workerHref: workerHref,
	}
}

// Execute executes the request
//  @return WorkerResponse
func (a *WorkersAPIService) WorkersReadExecute(r WorkersAPIWorkersReadRequest) (*WorkerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WorkerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkersAPIService.WorkersRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{worker_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"worker_href"+"}", url.PathEscape(parameterValueToString(r.workerHref, "workerHref")), -1)
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
