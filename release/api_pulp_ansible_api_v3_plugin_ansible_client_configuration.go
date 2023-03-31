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


// PulpAnsibleApiV3PluginAnsibleClientConfigurationApiService PulpAnsibleApiV3PluginAnsibleClientConfigurationApi service
type PulpAnsibleApiV3PluginAnsibleClientConfigurationApiService service

type PulpAnsibleApiV3PluginAnsibleClientConfigurationApiPulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationReadRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3PluginAnsibleClientConfigurationApiService
	path string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PulpAnsibleApiV3PluginAnsibleClientConfigurationApiPulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationReadRequest) Fields(fields []string) PulpAnsibleApiV3PluginAnsibleClientConfigurationApiPulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleApiV3PluginAnsibleClientConfigurationApiPulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationReadRequest) ExcludeFields(excludeFields []string) PulpAnsibleApiV3PluginAnsibleClientConfigurationApiPulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleClientConfigurationApiPulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationReadRequest) Execute() (*ClientConfigurationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationReadExecute(r)
}

/*
PulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationRead Method for PulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationRead

Return configurations for the ansible-galaxy client.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param path
 @return PulpAnsibleApiV3PluginAnsibleClientConfigurationApiPulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationReadRequest
*/
func (a *PulpAnsibleApiV3PluginAnsibleClientConfigurationApiService) PulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationRead(ctx context.Context, path string) PulpAnsibleApiV3PluginAnsibleClientConfigurationApiPulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationReadRequest {
	return PulpAnsibleApiV3PluginAnsibleClientConfigurationApiPulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationReadRequest{
		ApiService: a,
		ctx: ctx,
		path: path,
	}
}

// Execute executes the request
//  @return ClientConfigurationResponse
func (a *PulpAnsibleApiV3PluginAnsibleClientConfigurationApiService) PulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationReadExecute(r PulpAnsibleApiV3PluginAnsibleClientConfigurationApiPulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationReadRequest) (*ClientConfigurationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClientConfigurationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3PluginAnsibleClientConfigurationApiService.PulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/client-configuration/"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterValueToString(r.path, "path")), -1)
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
