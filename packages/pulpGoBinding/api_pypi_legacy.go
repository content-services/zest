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
	"strings"
	"os"
)


// PypiLegacyApiService PypiLegacyApi service
type PypiLegacyApiService service

type PypiLegacyApiPypiLegacyCreateRequest struct {
	ctx context.Context
	ApiService *PypiLegacyApiService
	path string
	content *os.File
	sha256Digest *string
	action *string
}

// A Python package release file to upload to the index.
func (r PypiLegacyApiPypiLegacyCreateRequest) Content(content *os.File) PypiLegacyApiPypiLegacyCreateRequest {
	r.content = content
	return r
}

// SHA256 of package to validate upload integrity.
func (r PypiLegacyApiPypiLegacyCreateRequest) Sha256Digest(sha256Digest string) PypiLegacyApiPypiLegacyCreateRequest {
	r.sha256Digest = &sha256Digest
	return r
}

// Defaults to &#x60;file_upload&#x60;, don&#39;t change it or request will fail!
func (r PypiLegacyApiPypiLegacyCreateRequest) Action(action string) PypiLegacyApiPypiLegacyCreateRequest {
	r.action = &action
	return r
}

func (r PypiLegacyApiPypiLegacyCreateRequest) Execute() (*PackageUploadTaskResponse, *http.Response, error) {
	return r.ApiService.PypiLegacyCreateExecute(r)
}

/*
PypiLegacyCreate Upload a package

Upload package to the index.

This is the endpoint that tools like Twine and Poetry use for their upload commands.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param path
 @return PypiLegacyApiPypiLegacyCreateRequest
*/
func (a *PypiLegacyApiService) PypiLegacyCreate(ctx context.Context, path string) PypiLegacyApiPypiLegacyCreateRequest {
	return PypiLegacyApiPypiLegacyCreateRequest{
		ApiService: a,
		ctx: ctx,
		path: path,
	}
}

// Execute executes the request
//  @return PackageUploadTaskResponse
func (a *PypiLegacyApiService) PypiLegacyCreateExecute(r PypiLegacyApiPypiLegacyCreateRequest) (*PackageUploadTaskResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PackageUploadTaskResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PypiLegacyApiService.PypiLegacyCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pypi/{path}/legacy/"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterValueToString(r.path, "path")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.content == nil {
		return localVarReturnValue, nil, reportError("content is required and must be specified")
	}
	if r.sha256Digest == nil {
		return localVarReturnValue, nil, reportError("sha256Digest is required and must be specified")
	}
	if strlen(*r.sha256Digest) < 64 {
		return localVarReturnValue, nil, reportError("sha256Digest must have at least 64 elements")
	}
	if strlen(*r.sha256Digest) > 64 {
		return localVarReturnValue, nil, reportError("sha256Digest must have less than 64 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data", "application/x-www-form-urlencoded"}

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
	var contentLocalVarFormFileName string
	var contentLocalVarFileName     string
	var contentLocalVarFileBytes    []byte

	contentLocalVarFormFileName = "content"


	contentLocalVarFile := r.content

	if contentLocalVarFile != nil {
		fbs, _ := io.ReadAll(contentLocalVarFile)

		contentLocalVarFileBytes = fbs
		contentLocalVarFileName = contentLocalVarFile.Name()
		contentLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: contentLocalVarFileBytes, fileName: contentLocalVarFileName, formFileName: contentLocalVarFormFileName})
	}
	if r.action != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "action", r.action, "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "sha256_digest", r.sha256Digest, "")
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