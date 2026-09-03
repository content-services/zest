# \ContentFilesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentFileFilesCreate**](ContentFilesAPI.md#ContentFileFilesCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/content/file/files/ | Create a file content
[**ContentFileFilesList**](ContentFilesAPI.md#ContentFileFilesList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/file/files/ | List file contents
[**ContentFileFilesRead**](ContentFilesAPI.md#ContentFileFilesRead) | **Get** /{file_file_content_href} | Inspect a file content
[**ContentFileFilesSetLabel**](ContentFilesAPI.md#ContentFileFilesSetLabel) | **Post** /{file_file_content_href}set_label/ | Set a label
[**ContentFileFilesUnsetLabel**](ContentFilesAPI.md#ContentFileFilesUnsetLabel) | **Post** /{file_file_content_href}unset_label/ | Unset a label
[**ContentFileFilesUpload**](ContentFilesAPI.md#ContentFileFilesUpload) | **Post** /api/pulp/{pulp_domain}/api/v3/content/file/files/upload/ | Upload a File synchronously.



## ContentFileFilesCreate

> AsyncOperationResponse ContentFileFilesCreate(ctx, pulpDomain).RelativePath(relativePath).XTaskDiagnostics(xTaskDiagnostics).Repository(repository).Overwrite(overwrite).PulpLabels(pulpLabels).Artifact(artifact).File(file).Upload(upload).FileUrl(fileUrl).DownloaderConfig(downloaderConfig).Execute()

Create a file content



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	relativePath := "relativePath_example" // string | Path where the artifact is located relative to distributions base_path
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
	overwrite := true // bool | When set to true, existing content in the repository with the same unique key will be silently overwritten. When set to false, the task will fail if content would be overwritten. Only used when 'repository' is specified. Defaults to true. (optional)
	pulpLabels := map[string]*string{"key": "Inner_example"} // map[string]*string | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. (optional)
	artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
	file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the content unit. (optional)
	upload := "upload_example" // string | An uncommitted upload that may be turned into the content unit. (optional)
	fileUrl := "fileUrl_example" // string | A url that Pulp can download and turn into the content unit. (optional)
	downloaderConfig := *openapiclient.NewRemoteNetworkConfig() // RemoteNetworkConfig | Configuration for the download process (e.g., proxies, auth, timeouts). Only applicable when providing a 'file_url. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentFilesAPI.ContentFileFilesCreate(context.Background(), pulpDomain).RelativePath(relativePath).XTaskDiagnostics(xTaskDiagnostics).Repository(repository).Overwrite(overwrite).PulpLabels(pulpLabels).Artifact(artifact).File(file).Upload(upload).FileUrl(fileUrl).DownloaderConfig(downloaderConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentFilesAPI.ContentFileFilesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentFileFilesCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentFilesAPI.ContentFileFilesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentFileFilesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **relativePath** | **string** | Path where the artifact is located relative to distributions base_path | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **overwrite** | **bool** | When set to true, existing content in the repository with the same unique key will be silently overwritten. When set to false, the task will fail if content would be overwritten. Only used when &#39;repository&#39; is specified. Defaults to true. | 
 **pulpLabels** | **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | 
 **artifact** | **string** | Artifact file representing the physical content | 
 **file** | ***os.File** | An uploaded file that may be turned into the content unit. | 
 **upload** | **string** | An uncommitted upload that may be turned into the content unit. | 
 **fileUrl** | **string** | A url that Pulp can download and turn into the content unit. | 
 **downloaderConfig** | [**RemoteNetworkConfig**](RemoteNetworkConfig.md) | Configuration for the download process (e.g., proxies, auth, timeouts). Only applicable when providing a &#39;file_url. | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentFileFilesList

> PaginatedfileFileContentResponseList ContentFileFilesList(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Digest(digest).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RelativePath(relativePath).RelativePathContains(relativePathContains).RelativePathIcontains(relativePathIcontains).RelativePathIexact(relativePathIexact).RelativePathIn(relativePathIn).RelativePathIregex(relativePathIregex).RelativePathIstartswith(relativePathIstartswith).RelativePathRegex(relativePathRegex).RelativePathStartswith(relativePathStartswith).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Fields(fields).ExcludeFields(excludeFields).Execute()

List file contents



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	digest := "digest_example" // string | Filter results where digest matches value (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `relative_path` - Relative path* `-relative_path` - Relative path (descending)* `digest` - Digest* `-digest` - Digest (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	relativePath := "relativePath_example" // string | Filter results where relative_path matches value (optional)
	relativePathContains := "relativePathContains_example" // string | Filter results where relative_path contains value (optional)
	relativePathIcontains := "relativePathIcontains_example" // string | Filter results where relative_path contains value (optional)
	relativePathIexact := "relativePathIexact_example" // string | Filter results where relative_path matches value (optional)
	relativePathIn := []string{"Inner_example"} // []string | Filter results where relative_path is in a comma-separated list of values (optional)
	relativePathIregex := "relativePathIregex_example" // string | Filter results where relative_path matches regex value (optional)
	relativePathIstartswith := "relativePathIstartswith_example" // string | Filter results where relative_path starts with value (optional)
	relativePathRegex := "relativePathRegex_example" // string | Filter results where relative_path matches regex value (optional)
	relativePathStartswith := "relativePathStartswith_example" // string | Filter results where relative_path starts with value (optional)
	repositoryVersion := "repositoryVersion_example" // string |  (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string |  (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string |  (optional)
	sha256 := "sha256_example" // string |  (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentFilesAPI.ContentFileFilesList(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Digest(digest).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RelativePath(relativePath).RelativePathContains(relativePathContains).RelativePathIcontains(relativePathIcontains).RelativePathIexact(relativePathIexact).RelativePathIn(relativePathIn).RelativePathIregex(relativePathIregex).RelativePathIstartswith(relativePathIstartswith).RelativePathRegex(relativePathRegex).RelativePathStartswith(relativePathStartswith).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentFilesAPI.ContentFileFilesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentFileFilesList`: PaginatedfileFileContentResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentFilesAPI.ContentFileFilesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentFileFilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **digest** | **string** | Filter results where digest matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;relative_path&#x60; - Relative path* &#x60;-relative_path&#x60; - Relative path (descending)* &#x60;digest&#x60; - Digest* &#x60;-digest&#x60; - Digest (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **relativePath** | **string** | Filter results where relative_path matches value | 
 **relativePathContains** | **string** | Filter results where relative_path contains value | 
 **relativePathIcontains** | **string** | Filter results where relative_path contains value | 
 **relativePathIexact** | **string** | Filter results where relative_path matches value | 
 **relativePathIn** | **[]string** | Filter results where relative_path is in a comma-separated list of values | 
 **relativePathIregex** | **string** | Filter results where relative_path matches regex value | 
 **relativePathIstartswith** | **string** | Filter results where relative_path starts with value | 
 **relativePathRegex** | **string** | Filter results where relative_path matches regex value | 
 **relativePathStartswith** | **string** | Filter results where relative_path starts with value | 
 **repositoryVersion** | **string** |  | 
 **repositoryVersionAdded** | **string** |  | 
 **repositoryVersionRemoved** | **string** |  | 
 **sha256** | **string** |  | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedfileFileContentResponseList**](PaginatedfileFileContentResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentFileFilesRead

> FileFileContentResponse ContentFileFilesRead(ctx, fileFileContentHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a file content



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	fileFileContentHref := "fileFileContentHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentFilesAPI.ContentFileFilesRead(context.Background(), fileFileContentHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentFilesAPI.ContentFileFilesRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentFileFilesRead`: FileFileContentResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentFilesAPI.ContentFileFilesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileContentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentFileFilesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**FileFileContentResponse**](FileFileContentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentFileFilesSetLabel

> SetLabelResponse ContentFileFilesSetLabel(ctx, fileFileContentHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

Set a label



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	fileFileContentHref := "fileFileContentHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentFilesAPI.ContentFileFilesSetLabel(context.Background(), fileFileContentHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentFilesAPI.ContentFileFilesSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentFileFilesSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentFilesAPI.ContentFileFilesSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileContentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentFileFilesSetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setLabel** | [**SetLabel**](SetLabel.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**SetLabelResponse**](SetLabelResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentFileFilesUnsetLabel

> UnsetLabelResponse ContentFileFilesUnsetLabel(ctx, fileFileContentHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

Unset a label



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	fileFileContentHref := "fileFileContentHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentFilesAPI.ContentFileFilesUnsetLabel(context.Background(), fileFileContentHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentFilesAPI.ContentFileFilesUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentFileFilesUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentFilesAPI.ContentFileFilesUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileContentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentFileFilesUnsetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unsetLabel** | [**UnsetLabel**](UnsetLabel.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**UnsetLabelResponse**](UnsetLabelResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentFileFilesUpload

> FileContentUploadResponse ContentFileFilesUpload(ctx, pulpDomain).RelativePath(relativePath).XTaskDiagnostics(xTaskDiagnostics).Overwrite(overwrite).PulpLabels(pulpLabels).Artifact(artifact).File(file).Upload(upload).FileUrl(fileUrl).DownloaderConfig(downloaderConfig).Execute()

Upload a File synchronously.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	relativePath := "relativePath_example" // string | Path where the artifact is located relative to distributions base_path
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	overwrite := true // bool | When set to true, existing content in the repository with the same unique key will be silently overwritten. When set to false, the task will fail if content would be overwritten. Only used when 'repository' is specified. Defaults to true. (optional)
	pulpLabels := map[string]*string{"key": "Inner_example"} // map[string]*string | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. (optional)
	artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
	file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the content unit. (optional)
	upload := "upload_example" // string | An uncommitted upload that may be turned into the content unit. (optional)
	fileUrl := "fileUrl_example" // string | A url that Pulp can download and turn into the content unit. (optional)
	downloaderConfig := *openapiclient.NewRemoteNetworkConfig() // RemoteNetworkConfig | Configuration for the download process (e.g., proxies, auth, timeouts). Only applicable when providing a 'file_url. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentFilesAPI.ContentFileFilesUpload(context.Background(), pulpDomain).RelativePath(relativePath).XTaskDiagnostics(xTaskDiagnostics).Overwrite(overwrite).PulpLabels(pulpLabels).Artifact(artifact).File(file).Upload(upload).FileUrl(fileUrl).DownloaderConfig(downloaderConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentFilesAPI.ContentFileFilesUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentFileFilesUpload`: FileContentUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentFilesAPI.ContentFileFilesUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentFileFilesUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **relativePath** | **string** | Path where the artifact is located relative to distributions base_path | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **overwrite** | **bool** | When set to true, existing content in the repository with the same unique key will be silently overwritten. When set to false, the task will fail if content would be overwritten. Only used when &#39;repository&#39; is specified. Defaults to true. | 
 **pulpLabels** | **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | 
 **artifact** | **string** | Artifact file representing the physical content | 
 **file** | ***os.File** | An uploaded file that may be turned into the content unit. | 
 **upload** | **string** | An uncommitted upload that may be turned into the content unit. | 
 **fileUrl** | **string** | A url that Pulp can download and turn into the content unit. | 
 **downloaderConfig** | [**RemoteNetworkConfig**](RemoteNetworkConfig.md) | Configuration for the download process (e.g., proxies, auth, timeouts). Only applicable when providing a &#39;file_url. | 

### Return type

[**FileContentUploadResponse**](FileContentUploadResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

