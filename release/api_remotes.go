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


// RemotesAPIService RemotesAPI service
type RemotesAPIService service

type RemotesAPIRemotesListRequest struct {
	ctx context.Context
	ApiService *RemotesAPIService
	pulpDomain string
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
	pulpLastUpdated *time.Time
	pulpLastUpdatedGt *time.Time
	pulpLastUpdatedGte *time.Time
	pulpLastUpdatedLt *time.Time
	pulpLastUpdatedLte *time.Time
	pulpLastUpdatedRange *[]time.Time
	pulpType *string
	pulpTypeIn *[]string
	q *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r RemotesAPIRemotesListRequest) Limit(limit int32) RemotesAPIRemotesListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r RemotesAPIRemotesListRequest) Name(name string) RemotesAPIRemotesListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r RemotesAPIRemotesListRequest) NameContains(nameContains string) RemotesAPIRemotesListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r RemotesAPIRemotesListRequest) NameIcontains(nameIcontains string) RemotesAPIRemotesListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name matches value
func (r RemotesAPIRemotesListRequest) NameIexact(nameIexact string) RemotesAPIRemotesListRequest {
	r.nameIexact = &nameIexact
	return r
}

// Filter results where name is in a comma-separated list of values
func (r RemotesAPIRemotesListRequest) NameIn(nameIn []string) RemotesAPIRemotesListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name matches regex value
func (r RemotesAPIRemotesListRequest) NameIregex(nameIregex string) RemotesAPIRemotesListRequest {
	r.nameIregex = &nameIregex
	return r
}

// Filter results where name starts with value
func (r RemotesAPIRemotesListRequest) NameIstartswith(nameIstartswith string) RemotesAPIRemotesListRequest {
	r.nameIstartswith = &nameIstartswith
	return r
}

// Filter results where name matches regex value
func (r RemotesAPIRemotesListRequest) NameRegex(nameRegex string) RemotesAPIRemotesListRequest {
	r.nameRegex = &nameRegex
	return r
}

// Filter results where name starts with value
func (r RemotesAPIRemotesListRequest) NameStartswith(nameStartswith string) RemotesAPIRemotesListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r RemotesAPIRemotesListRequest) Offset(offset int32) RemotesAPIRemotesListRequest {
	r.offset = &offset
	return r
}

// Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;url&#x60; - Url* &#x60;-url&#x60; - Url (descending)* &#x60;ca_cert&#x60; - Ca cert* &#x60;-ca_cert&#x60; - Ca cert (descending)* &#x60;client_cert&#x60; - Client cert* &#x60;-client_cert&#x60; - Client cert (descending)* &#x60;client_key&#x60; - Client key* &#x60;-client_key&#x60; - Client key (descending)* &#x60;tls_validation&#x60; - Tls validation* &#x60;-tls_validation&#x60; - Tls validation (descending)* &#x60;username&#x60; - Username* &#x60;-username&#x60; - Username (descending)* &#x60;password&#x60; - Password* &#x60;-password&#x60; - Password (descending)* &#x60;proxy_url&#x60; - Proxy url* &#x60;-proxy_url&#x60; - Proxy url (descending)* &#x60;proxy_username&#x60; - Proxy username* &#x60;-proxy_username&#x60; - Proxy username (descending)* &#x60;proxy_password&#x60; - Proxy password* &#x60;-proxy_password&#x60; - Proxy password (descending)* &#x60;download_concurrency&#x60; - Download concurrency* &#x60;-download_concurrency&#x60; - Download concurrency (descending)* &#x60;max_retries&#x60; - Max retries* &#x60;-max_retries&#x60; - Max retries (descending)* &#x60;policy&#x60; - Policy* &#x60;-policy&#x60; - Policy (descending)* &#x60;total_timeout&#x60; - Total timeout* &#x60;-total_timeout&#x60; - Total timeout (descending)* &#x60;connect_timeout&#x60; - Connect timeout* &#x60;-connect_timeout&#x60; - Connect timeout (descending)* &#x60;sock_connect_timeout&#x60; - Sock connect timeout* &#x60;-sock_connect_timeout&#x60; - Sock connect timeout (descending)* &#x60;sock_read_timeout&#x60; - Sock read timeout* &#x60;-sock_read_timeout&#x60; - Sock read timeout (descending)* &#x60;headers&#x60; - Headers* &#x60;-headers&#x60; - Headers (descending)* &#x60;rate_limit&#x60; - Rate limit* &#x60;-rate_limit&#x60; - Rate limit (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending)
func (r RemotesAPIRemotesListRequest) Ordering(ordering []string) RemotesAPIRemotesListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r RemotesAPIRemotesListRequest) PrnIn(prnIn []string) RemotesAPIRemotesListRequest {
	r.prnIn = &prnIn
	return r
}

// Multiple values may be separated by commas.
func (r RemotesAPIRemotesListRequest) PulpHrefIn(pulpHrefIn []string) RemotesAPIRemotesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r RemotesAPIRemotesListRequest) PulpIdIn(pulpIdIn []string) RemotesAPIRemotesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter labels by search string
func (r RemotesAPIRemotesListRequest) PulpLabelSelect(pulpLabelSelect string) RemotesAPIRemotesListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Filter results where pulp_last_updated matches value
func (r RemotesAPIRemotesListRequest) PulpLastUpdated(pulpLastUpdated time.Time) RemotesAPIRemotesListRequest {
	r.pulpLastUpdated = &pulpLastUpdated
	return r
}

// Filter results where pulp_last_updated is greater than value
func (r RemotesAPIRemotesListRequest) PulpLastUpdatedGt(pulpLastUpdatedGt time.Time) RemotesAPIRemotesListRequest {
	r.pulpLastUpdatedGt = &pulpLastUpdatedGt
	return r
}

// Filter results where pulp_last_updated is greater than or equal to value
func (r RemotesAPIRemotesListRequest) PulpLastUpdatedGte(pulpLastUpdatedGte time.Time) RemotesAPIRemotesListRequest {
	r.pulpLastUpdatedGte = &pulpLastUpdatedGte
	return r
}

// Filter results where pulp_last_updated is less than value
func (r RemotesAPIRemotesListRequest) PulpLastUpdatedLt(pulpLastUpdatedLt time.Time) RemotesAPIRemotesListRequest {
	r.pulpLastUpdatedLt = &pulpLastUpdatedLt
	return r
}

// Filter results where pulp_last_updated is less than or equal to value
func (r RemotesAPIRemotesListRequest) PulpLastUpdatedLte(pulpLastUpdatedLte time.Time) RemotesAPIRemotesListRequest {
	r.pulpLastUpdatedLte = &pulpLastUpdatedLte
	return r
}

// Filter results where pulp_last_updated is between two comma separated values
func (r RemotesAPIRemotesListRequest) PulpLastUpdatedRange(pulpLastUpdatedRange []time.Time) RemotesAPIRemotesListRequest {
	r.pulpLastUpdatedRange = &pulpLastUpdatedRange
	return r
}

// Pulp type* &#x60;rpm.rpm&#x60; - rpm.rpm* &#x60;rpm.uln&#x60; - rpm.uln* &#x60;gem.gem&#x60; - gem.gem* &#x60;python.python&#x60; - python.python* &#x60;ostree.ostree&#x60; - ostree.ostree* &#x60;file.file&#x60; - file.file
func (r RemotesAPIRemotesListRequest) PulpType(pulpType string) RemotesAPIRemotesListRequest {
	r.pulpType = &pulpType
	return r
}

// Multiple values may be separated by commas.* &#x60;rpm.rpm&#x60; - rpm.rpm* &#x60;rpm.uln&#x60; - rpm.uln* &#x60;gem.gem&#x60; - gem.gem* &#x60;python.python&#x60; - python.python* &#x60;ostree.ostree&#x60; - ostree.ostree* &#x60;file.file&#x60; - file.file
func (r RemotesAPIRemotesListRequest) PulpTypeIn(pulpTypeIn []string) RemotesAPIRemotesListRequest {
	r.pulpTypeIn = &pulpTypeIn
	return r
}

// Filter results by using NOT, AND and OR operations on other filters
func (r RemotesAPIRemotesListRequest) Q(q string) RemotesAPIRemotesListRequest {
	r.q = &q
	return r
}

// A list of fields to include in the response.
func (r RemotesAPIRemotesListRequest) Fields(fields []string) RemotesAPIRemotesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RemotesAPIRemotesListRequest) ExcludeFields(excludeFields []string) RemotesAPIRemotesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RemotesAPIRemotesListRequest) Execute() (*PaginatedRemoteResponseList, *http.Response, error) {
	return r.ApiService.RemotesListExecute(r)
}

/*
RemotesList List remotes

A customized named ModelViewSet that knows how to register itself with the Pulp API router.This viewset is discoverable by its name."Normal" Django Models and Master/Detail models are supported by the ``register_with`` method.Attributes:    lookup_field (str): The name of the field by which an object should be looked up, in        addition to any parent lookups if this ViewSet is nested. Defaults to 'pk'    endpoint_name (str): The name of the final path segment that should identify the ViewSet's        collection endpoint.    nest_prefix (str): Optional prefix under which this ViewSet should be nested. This must        correspond to the "parent_prefix" of a router with rest_framework_nested.NestedMixin.        None indicates this ViewSet should not be nested.    parent_lookup_kwargs (dict): Optional mapping of key names that would appear in self.kwargs        to django model filter expressions that can be used with the corresponding value from        self.kwargs, used only by a nested ViewSet to filter based on the parent object's        identity.    schema (DefaultSchema): The schema class to use by default in a viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return RemotesAPIRemotesListRequest
*/
func (a *RemotesAPIService) RemotesList(ctx context.Context, pulpDomain string) RemotesAPIRemotesListRequest {
	return RemotesAPIRemotesListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedRemoteResponseList
func (a *RemotesAPIService) RemotesListExecute(r RemotesAPIRemotesListRequest) (*PaginatedRemoteResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRemoteResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RemotesAPIService.RemotesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/pulp/{pulp_domain}/api/v3/remotes/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.pulpLastUpdated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated", r.pulpLastUpdated, "form", "")
	}
	if r.pulpLastUpdatedGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__gt", r.pulpLastUpdatedGt, "form", "")
	}
	if r.pulpLastUpdatedGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__gte", r.pulpLastUpdatedGte, "form", "")
	}
	if r.pulpLastUpdatedLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__lt", r.pulpLastUpdatedLt, "form", "")
	}
	if r.pulpLastUpdatedLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__lte", r.pulpLastUpdatedLte, "form", "")
	}
	if r.pulpLastUpdatedRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_last_updated__range", r.pulpLastUpdatedRange, "form", "csv")
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
