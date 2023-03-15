/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zest/release/v3

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiService PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApi service
type PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiService service

type PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiService
	distroBasePath string
	filename string
	path string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest) Fields(fields []string) PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest) ExcludeFields(excludeFields []string) PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest) Execute() (*http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownloadExecute(r)
}

/*
PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownload Method for PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownload

Collection download endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distroBasePath
 @param filename
 @param path
 @return PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest
*/
func (a *PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiService) PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownload(ctx context.Context, distroBasePath string, filename string, path string) PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest {
	return PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest{
		ApiService: a,
		ctx: ctx,
		distroBasePath: distroBasePath,
		filename: filename,
		path: path,
	}
}

// Execute executes the request
func (a *PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiService) PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownloadExecute(r PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiService.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownload")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/content/{distro_base_path}/collections/artifacts/{filename}"
	localVarPath = strings.Replace(localVarPath, "{"+"distro_base_path"+"}", url.PathEscape(parameterValueToString(r.distroBasePath, "distroBasePath")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"filename"+"}", url.PathEscape(parameterValueToString(r.filename, "filename")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

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
