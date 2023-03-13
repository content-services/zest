/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpGoBinding

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// RepositoriesReclaimSpaceApiService RepositoriesReclaimSpaceApi service
type RepositoriesReclaimSpaceApiService service

type RepositoriesReclaimSpaceApiRepositoriesReclaimSpaceReclaimRequest struct {
	ctx context.Context
	ApiService *RepositoriesReclaimSpaceApiService
	reclaimSpace *ReclaimSpace
}

func (r RepositoriesReclaimSpaceApiRepositoriesReclaimSpaceReclaimRequest) ReclaimSpace(reclaimSpace ReclaimSpace) RepositoriesReclaimSpaceApiRepositoriesReclaimSpaceReclaimRequest {
	r.reclaimSpace = &reclaimSpace
	return r
}

func (r RepositoriesReclaimSpaceApiRepositoriesReclaimSpaceReclaimRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesReclaimSpaceReclaimExecute(r)
}

/*
RepositoriesReclaimSpaceReclaim Method for RepositoriesReclaimSpaceReclaim

Trigger an asynchronous space reclaim operation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RepositoriesReclaimSpaceApiRepositoriesReclaimSpaceReclaimRequest
*/
func (a *RepositoriesReclaimSpaceApiService) RepositoriesReclaimSpaceReclaim(ctx context.Context) RepositoriesReclaimSpaceApiRepositoriesReclaimSpaceReclaimRequest {
	return RepositoriesReclaimSpaceApiRepositoriesReclaimSpaceReclaimRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesReclaimSpaceApiService) RepositoriesReclaimSpaceReclaimExecute(r RepositoriesReclaimSpaceApiRepositoriesReclaimSpaceReclaimRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesReclaimSpaceApiService.RepositoriesReclaimSpaceReclaim")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/repositories/reclaim_space/"
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