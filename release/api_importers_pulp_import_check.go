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


// ImportersPulpImportCheckAPIService ImportersPulpImportCheckAPI service
type ImportersPulpImportCheckAPIService service

type ImportersPulpImportCheckAPIPulpImportCheckPostRequest struct {
	ctx context.Context
	ApiService *ImportersPulpImportCheckAPIService
	pulpImportCheck *PulpImportCheck
}

func (r ImportersPulpImportCheckAPIPulpImportCheckPostRequest) PulpImportCheck(pulpImportCheck PulpImportCheck) ImportersPulpImportCheckAPIPulpImportCheckPostRequest {
	r.pulpImportCheck = &pulpImportCheck
	return r
}

func (r ImportersPulpImportCheckAPIPulpImportCheckPostRequest) Execute() (*PulpImportCheckResponse, *http.Response, error) {
	return r.ApiService.PulpImportCheckPostExecute(r)
}

/*
PulpImportCheckPost Validate the parameters to be used for a PulpImport call

Evaluates validity of proposed PulpImport parameters 'toc', 'path', and 'repo_mapping'.

* Checks that toc, path are in ALLOWED_IMPORT_PATHS
* if ALLOWED:
  * Checks that toc, path exist and are readable
  * If toc specified, checks that containing dir is writeable
* Checks that repo_mapping is valid JSON

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ImportersPulpImportCheckAPIPulpImportCheckPostRequest
*/
func (a *ImportersPulpImportCheckAPIService) PulpImportCheckPost(ctx context.Context) ImportersPulpImportCheckAPIPulpImportCheckPostRequest {
	return ImportersPulpImportCheckAPIPulpImportCheckPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PulpImportCheckResponse
func (a *ImportersPulpImportCheckAPIService) PulpImportCheckPostExecute(r ImportersPulpImportCheckAPIPulpImportCheckPostRequest) (*PulpImportCheckResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PulpImportCheckResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportersPulpImportCheckAPIService.PulpImportCheckPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/importers/core/pulp/import-check/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pulpImportCheck == nil {
		return localVarReturnValue, nil, reportError("pulpImportCheck is required and must be specified")
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
	localVarPostBody = r.pulpImportCheck
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
