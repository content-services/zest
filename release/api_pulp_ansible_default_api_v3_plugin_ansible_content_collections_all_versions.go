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


// PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPIService PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPI service
type PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPIService service

type PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsListRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPIService
	distroBasePath string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsListRequest) Fields(fields []string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsListRequest) ExcludeFields(excludeFields []string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsListRequest) Execute() ([]UnpaginatedCollectionVersionResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsListExecute(r)
}

/*
PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsList Method for PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsList

Returns paginated CollectionVersions list.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distroBasePath
 @return PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsListRequest
*/
func (a *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPIService) PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsList(ctx context.Context, distroBasePath string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsListRequest {
	return PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsListRequest{
		ApiService: a,
		ctx: ctx,
		distroBasePath: distroBasePath,
	}
}

// Execute executes the request
//  @return []UnpaginatedCollectionVersionResponse
func (a *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPIService) PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsListExecute(r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPIPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsListRequest) ([]UnpaginatedCollectionVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []UnpaginatedCollectionVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsAPIService.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/all-versions/"
	localVarPath = strings.Replace(localVarPath, "{"+"distro_base_path"+"}", url.PathEscape(parameterValueToString(r.distroBasePath, "distroBasePath")), -1)
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
