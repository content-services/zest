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
)


// RepositoriesReclaimSpaceAPIService RepositoriesReclaimSpaceAPI service
type RepositoriesReclaimSpaceAPIService service

type RepositoriesReclaimSpaceAPIRepositoriesReclaimSpaceReclaimRequest struct {
	ctx context.Context
	ApiService *RepositoriesReclaimSpaceAPIService
	pulpDomain string
	reclaimSpace *ReclaimSpace
}

func (r RepositoriesReclaimSpaceAPIRepositoriesReclaimSpaceReclaimRequest) ReclaimSpace(reclaimSpace ReclaimSpace) RepositoriesReclaimSpaceAPIRepositoriesReclaimSpaceReclaimRequest {
	r.reclaimSpace = &reclaimSpace
	return r
}

func (r RepositoriesReclaimSpaceAPIRepositoriesReclaimSpaceReclaimRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesReclaimSpaceReclaimExecute(r)
}

/*
RepositoriesReclaimSpaceReclaim Method for RepositoriesReclaimSpaceReclaim

Trigger an asynchronous space reclaim operation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return RepositoriesReclaimSpaceAPIRepositoriesReclaimSpaceReclaimRequest
*/
func (a *RepositoriesReclaimSpaceAPIService) RepositoriesReclaimSpaceReclaim(ctx context.Context, pulpDomain string) RepositoriesReclaimSpaceAPIRepositoriesReclaimSpaceReclaimRequest {
	return RepositoriesReclaimSpaceAPIRepositoriesReclaimSpaceReclaimRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesReclaimSpaceAPIService) RepositoriesReclaimSpaceReclaimExecute(r RepositoriesReclaimSpaceAPIRepositoriesReclaimSpaceReclaimRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesReclaimSpaceAPIService.RepositoriesReclaimSpaceReclaim")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/{pulp_domain}/api/v3/repositories/reclaim_space/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.reclaimSpace == nil {
		return localVarReturnValue, nil, reportError("reclaimSpace is required and must be specified")
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
	localVarPostBody = r.reclaimSpace
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
