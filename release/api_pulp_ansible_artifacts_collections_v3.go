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
	"strings"
	"os"
)


// PulpAnsibleArtifactsCollectionsV3ApiService PulpAnsibleArtifactsCollectionsV3Api service
type PulpAnsibleArtifactsCollectionsV3ApiService service

type PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleArtifactsCollectionsV3ApiService
	path string
	file *os.File
	sha256 *string
	expectedNamespace *string
	expectedName *string
	expectedVersion *string
}

// The Collection tarball.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest) File(file *os.File) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest {
	r.file = file
	return r
}

// An optional sha256 checksum of the uploaded file.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest) Sha256(sha256 string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest {
	r.sha256 = &sha256
	return r
}

// The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest) ExpectedNamespace(expectedNamespace string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest {
	r.expectedNamespace = &expectedNamespace
	return r
}

// The expected &#39;name&#39; of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest) ExpectedName(expectedName string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest {
	r.expectedName = &expectedName
	return r
}

// The expected version of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest) ExpectedVersion(expectedVersion string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest {
	r.expectedVersion = &expectedVersion
	return r
}

func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateExecute(r)
}

/*
PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreate Upload a collection

Create an artifact and trigger an asynchronous task to create Collection content from it.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param path
 @return PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest

Deprecated
*/
func (a *PulpAnsibleArtifactsCollectionsV3ApiService) PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreate(ctx context.Context, path string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest {
	return PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest{
		ApiService: a,
		ctx: ctx,
		path: path,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
// Deprecated
func (a *PulpAnsibleArtifactsCollectionsV3ApiService) PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateExecute(r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3ArtifactsCollectionsCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleArtifactsCollectionsV3ApiService.PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/artifacts/collections/"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterValueToString(r.path, "path")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha256", r.sha256, "")
	}
	if r.expectedNamespace != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_namespace", r.expectedNamespace, "")
	}
	if r.expectedName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_name", r.expectedName, "")
	}
	if r.expectedVersion != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_version", r.expectedVersion, "")
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

type PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleArtifactsCollectionsV3ApiService
	distroBasePath string
	path string
	file *os.File
	sha256 *string
	expectedNamespace *string
	expectedName *string
	expectedVersion *string
}

// The Collection tarball.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) File(file *os.File) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.file = file
	return r
}

// An optional sha256 checksum of the uploaded file.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) Sha256(sha256 string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.sha256 = &sha256
	return r
}

// The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) ExpectedNamespace(expectedNamespace string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.expectedNamespace = &expectedNamespace
	return r
}

// The expected &#39;name&#39; of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) ExpectedName(expectedName string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.expectedName = &expectedName
	return r
}

// The expected version of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) ExpectedVersion(expectedVersion string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.expectedVersion = &expectedVersion
	return r
}

func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateExecute(r)
}

/*
PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreate Upload a collection

Create an artifact and trigger an asynchronous task to create Collection content from it.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distroBasePath
 @param path
 @return PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest
*/
func (a *PulpAnsibleArtifactsCollectionsV3ApiService) PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreate(ctx context.Context, distroBasePath string, path string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	return PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest{
		ApiService: a,
		ctx: ctx,
		distroBasePath: distroBasePath,
		path: path,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *PulpAnsibleArtifactsCollectionsV3ApiService) PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateExecute(r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleArtifactsCollectionsV3ApiService.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/content/{distro_base_path}/collections/artifacts/"
	localVarPath = strings.Replace(localVarPath, "{"+"distro_base_path"+"}", url.PathEscape(parameterValueToString(r.distroBasePath, "distroBasePath")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterValueToString(r.path, "path")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha256", r.sha256, "")
	}
	if r.expectedNamespace != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_namespace", r.expectedNamespace, "")
	}
	if r.expectedName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_name", r.expectedName, "")
	}
	if r.expectedVersion != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_version", r.expectedVersion, "")
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

type PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleArtifactsCollectionsV3ApiService
	file *os.File
	sha256 *string
	expectedNamespace *string
	expectedName *string
	expectedVersion *string
}

// The Collection tarball.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest) File(file *os.File) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest {
	r.file = file
	return r
}

// An optional sha256 checksum of the uploaded file.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest) Sha256(sha256 string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest {
	r.sha256 = &sha256
	return r
}

// The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest) ExpectedNamespace(expectedNamespace string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest {
	r.expectedNamespace = &expectedNamespace
	return r
}

// The expected &#39;name&#39; of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest) ExpectedName(expectedName string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest {
	r.expectedName = &expectedName
	return r
}

// The expected version of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest) ExpectedVersion(expectedVersion string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest {
	r.expectedVersion = &expectedVersion
	return r
}

func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateExecute(r)
}

/*
PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreate Upload a collection

Create an artifact and trigger an asynchronous task to create Collection content from it.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest

Deprecated
*/
func (a *PulpAnsibleArtifactsCollectionsV3ApiService) PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreate(ctx context.Context) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest {
	return PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
// Deprecated
func (a *PulpAnsibleArtifactsCollectionsV3ApiService) PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateExecute(r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleArtifactsCollectionsV3ApiService.PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/default/api/v3/artifacts/collections/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha256", r.sha256, "")
	}
	if r.expectedNamespace != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_namespace", r.expectedNamespace, "")
	}
	if r.expectedName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_name", r.expectedName, "")
	}
	if r.expectedVersion != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_version", r.expectedVersion, "")
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

type PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleArtifactsCollectionsV3ApiService
	distroBasePath string
	file *os.File
	sha256 *string
	expectedNamespace *string
	expectedName *string
	expectedVersion *string
}

// The Collection tarball.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) File(file *os.File) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.file = file
	return r
}

// An optional sha256 checksum of the uploaded file.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) Sha256(sha256 string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.sha256 = &sha256
	return r
}

// The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) ExpectedNamespace(expectedNamespace string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.expectedNamespace = &expectedNamespace
	return r
}

// The expected &#39;name&#39; of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) ExpectedName(expectedName string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.expectedName = &expectedName
	return r
}

// The expected version of the Collection to be verified against the metadata during import.
func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) ExpectedVersion(expectedVersion string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	r.expectedVersion = &expectedVersion
	return r
}

func (r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateExecute(r)
}

/*
PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreate Upload a collection

Create an artifact and trigger an asynchronous task to create Collection content from it.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distroBasePath
 @return PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest
*/
func (a *PulpAnsibleArtifactsCollectionsV3ApiService) PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreate(ctx context.Context, distroBasePath string) PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest {
	return PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest{
		ApiService: a,
		ctx: ctx,
		distroBasePath: distroBasePath,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *PulpAnsibleArtifactsCollectionsV3ApiService) PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateExecute(r PulpAnsibleArtifactsCollectionsV3ApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleArtifactsCollectionsV3ApiService.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/artifacts/"
	localVarPath = strings.Replace(localVarPath, "{"+"distro_base_path"+"}", url.PathEscape(parameterValueToString(r.distroBasePath, "distroBasePath")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha256", r.sha256, "")
	}
	if r.expectedNamespace != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_namespace", r.expectedNamespace, "")
	}
	if r.expectedName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_name", r.expectedName, "")
	}
	if r.expectedVersion != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_version", r.expectedVersion, "")
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
