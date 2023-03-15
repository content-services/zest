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
)


// PulpAnsibleDefaultApiV3ApiService PulpAnsibleDefaultApiV3Api service
type PulpAnsibleDefaultApiV3ApiService service

type PulpAnsibleDefaultApiV3ApiPulpAnsibleGalaxyDefaultApiV3ReadRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleDefaultApiV3ApiService
}

func (r PulpAnsibleDefaultApiV3ApiPulpAnsibleGalaxyDefaultApiV3ReadRequest) Execute() (*RepoMetadataResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyDefaultApiV3ReadExecute(r)
}

/*
PulpAnsibleGalaxyDefaultApiV3Read Method for PulpAnsibleGalaxyDefaultApiV3Read

Legacy v3 endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PulpAnsibleDefaultApiV3ApiPulpAnsibleGalaxyDefaultApiV3ReadRequest

Deprecated
*/
func (a *PulpAnsibleDefaultApiV3ApiService) PulpAnsibleGalaxyDefaultApiV3Read(ctx context.Context) PulpAnsibleDefaultApiV3ApiPulpAnsibleGalaxyDefaultApiV3ReadRequest {
	return PulpAnsibleDefaultApiV3ApiPulpAnsibleGalaxyDefaultApiV3ReadRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RepoMetadataResponse
// Deprecated
func (a *PulpAnsibleDefaultApiV3ApiService) PulpAnsibleGalaxyDefaultApiV3ReadExecute(r PulpAnsibleDefaultApiV3ApiPulpAnsibleGalaxyDefaultApiV3ReadRequest) (*RepoMetadataResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepoMetadataResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleDefaultApiV3ApiService.PulpAnsibleGalaxyDefaultApiV3Read")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/default/api/v3/"
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
