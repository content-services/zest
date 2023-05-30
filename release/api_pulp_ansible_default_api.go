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
)


// PulpAnsibleDefaultApiAPIService PulpAnsibleDefaultApiAPI service
type PulpAnsibleDefaultApiAPIService service

type PulpAnsibleDefaultApiAPIPulpAnsibleGalaxyDefaultApiGetRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleDefaultApiAPIService
}

func (r PulpAnsibleDefaultApiAPIPulpAnsibleGalaxyDefaultApiGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyDefaultApiGetExecute(r)
}

/*
PulpAnsibleGalaxyDefaultApiGet Method for PulpAnsibleGalaxyDefaultApiGet

Return a response to the "GET" action.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PulpAnsibleDefaultApiAPIPulpAnsibleGalaxyDefaultApiGetRequest
*/
func (a *PulpAnsibleDefaultApiAPIService) PulpAnsibleGalaxyDefaultApiGet(ctx context.Context) PulpAnsibleDefaultApiAPIPulpAnsibleGalaxyDefaultApiGetRequest {
	return PulpAnsibleDefaultApiAPIPulpAnsibleGalaxyDefaultApiGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PulpAnsibleDefaultApiAPIService) PulpAnsibleGalaxyDefaultApiGetExecute(r PulpAnsibleDefaultApiAPIPulpAnsibleGalaxyDefaultApiGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleDefaultApiAPIService.PulpAnsibleGalaxyDefaultApiGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/default/api/"
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
