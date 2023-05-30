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


// PulpAnsibleApiV2CollectionsVersionsAPIService PulpAnsibleApiV2CollectionsVersionsAPI service
type PulpAnsibleApiV2CollectionsVersionsAPIService service

type PulpAnsibleApiV2CollectionsVersionsAPIPulpAnsibleGalaxyApiV2CollectionsVersionsGetRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV2CollectionsVersionsAPIService
	name string
	namespace string
	path string
	version string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PulpAnsibleApiV2CollectionsVersionsAPIPulpAnsibleGalaxyApiV2CollectionsVersionsGetRequest) Fields(fields []string) PulpAnsibleApiV2CollectionsVersionsAPIPulpAnsibleGalaxyApiV2CollectionsVersionsGetRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleApiV2CollectionsVersionsAPIPulpAnsibleGalaxyApiV2CollectionsVersionsGetRequest) ExcludeFields(excludeFields []string) PulpAnsibleApiV2CollectionsVersionsAPIPulpAnsibleGalaxyApiV2CollectionsVersionsGetRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleApiV2CollectionsVersionsAPIPulpAnsibleGalaxyApiV2CollectionsVersionsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV2CollectionsVersionsGetExecute(r)
}

/*
PulpAnsibleGalaxyApiV2CollectionsVersionsGet Method for PulpAnsibleGalaxyApiV2CollectionsVersionsGet

Return a response to the "GET" action.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @param namespace
 @param path
 @param version
 @return PulpAnsibleApiV2CollectionsVersionsAPIPulpAnsibleGalaxyApiV2CollectionsVersionsGetRequest
*/
func (a *PulpAnsibleApiV2CollectionsVersionsAPIService) PulpAnsibleGalaxyApiV2CollectionsVersionsGet(ctx context.Context, name string, namespace string, path string, version string) PulpAnsibleApiV2CollectionsVersionsAPIPulpAnsibleGalaxyApiV2CollectionsVersionsGetRequest {
	return PulpAnsibleApiV2CollectionsVersionsAPIPulpAnsibleGalaxyApiV2CollectionsVersionsGetRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
		namespace: namespace,
		path: path,
		version: version,
	}
}

// Execute executes the request
func (a *PulpAnsibleApiV2CollectionsVersionsAPIService) PulpAnsibleGalaxyApiV2CollectionsVersionsGetExecute(r PulpAnsibleApiV2CollectionsVersionsAPIPulpAnsibleGalaxyApiV2CollectionsVersionsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV2CollectionsVersionsAPIService.PulpAnsibleGalaxyApiV2CollectionsVersionsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v2/collections/{namespace}/{name}/versions/{version}/"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(parameterValueToString(r.namespace, "namespace")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterValueToString(r.path, "path")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", url.PathEscape(parameterValueToString(r.version, "version")), -1)
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
