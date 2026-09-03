# \ContentArtifactAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentMavenArtifactCreate**](ContentArtifactAPI.md#ContentMavenArtifactCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/content/maven/artifact/ | Create a maven artifact
[**ContentMavenArtifactList**](ContentArtifactAPI.md#ContentMavenArtifactList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/maven/artifact/ | List maven artifacts
[**ContentMavenArtifactRead**](ContentArtifactAPI.md#ContentMavenArtifactRead) | **Get** /{maven_maven_artifact_href} | Inspect a maven artifact
[**ContentMavenArtifactSetLabel**](ContentArtifactAPI.md#ContentMavenArtifactSetLabel) | **Post** /{maven_maven_artifact_href}set_label/ | Set a label
[**ContentMavenArtifactUnsetLabel**](ContentArtifactAPI.md#ContentMavenArtifactUnsetLabel) | **Post** /{maven_maven_artifact_href}unset_label/ | Unset a label
[**ContentMavenArtifactUpload**](ContentArtifactAPI.md#ContentMavenArtifactUpload) | **Post** /api/pulp/{pulp_domain}/api/v3/content/maven/artifact/upload/ | Upload a Maven artifact synchronously.



## ContentMavenArtifactCreate

> AsyncOperationResponse ContentMavenArtifactCreate(ctx, pulpDomain).RelativePath(relativePath).XTaskDiagnostics(xTaskDiagnostics).Repository(repository).Overwrite(overwrite).PulpLabels(pulpLabels).Artifact(artifact).File(file).Upload(upload).FileUrl(fileUrl).DownloaderConfig(downloaderConfig).Execute()

Create a maven artifact



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
	resp, r, err := apiClient.ContentArtifactAPI.ContentMavenArtifactCreate(context.Background(), pulpDomain).RelativePath(relativePath).XTaskDiagnostics(xTaskDiagnostics).Repository(repository).Overwrite(overwrite).PulpLabels(pulpLabels).Artifact(artifact).File(file).Upload(upload).FileUrl(fileUrl).DownloaderConfig(downloaderConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentArtifactAPI.ContentMavenArtifactCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentMavenArtifactCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentArtifactAPI.ContentMavenArtifactCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentMavenArtifactCreateRequest struct via the builder pattern


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


## ContentMavenArtifactList

> PaginatedmavenMavenArtifactResponseList ContentMavenArtifactList(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).ArtifactId(artifactId).Filename(filename).GroupId(groupId).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()

List maven artifacts



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
	artifactId := "artifactId_example" // string | Filter results where artifact_id matches value (optional)
	filename := "filename_example" // string | Filter results where filename matches value (optional)
	groupId := "groupId_example" // string | Filter results where group_id matches value (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `group_id` - Group id* `-group_id` - Group id (descending)* `artifact_id` - Artifact id* `-artifact_id` - Artifact id (descending)* `version` - Version* `-version` - Version (descending)* `filename` - Filename* `-filename` - Filename (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repositoryVersion := "repositoryVersion_example" // string |  (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string |  (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string |  (optional)
	version := "version_example" // string | Filter results where version matches value (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentArtifactAPI.ContentMavenArtifactList(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).ArtifactId(artifactId).Filename(filename).GroupId(groupId).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentArtifactAPI.ContentMavenArtifactList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentMavenArtifactList`: PaginatedmavenMavenArtifactResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentArtifactAPI.ContentMavenArtifactList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentMavenArtifactListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **artifactId** | **string** | Filter results where artifact_id matches value | 
 **filename** | **string** | Filter results where filename matches value | 
 **groupId** | **string** | Filter results where group_id matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;group_id&#x60; - Group id* &#x60;-group_id&#x60; - Group id (descending)* &#x60;artifact_id&#x60; - Artifact id* &#x60;-artifact_id&#x60; - Artifact id (descending)* &#x60;version&#x60; - Version* &#x60;-version&#x60; - Version (descending)* &#x60;filename&#x60; - Filename* &#x60;-filename&#x60; - Filename (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repositoryVersion** | **string** |  | 
 **repositoryVersionAdded** | **string** |  | 
 **repositoryVersionRemoved** | **string** |  | 
 **version** | **string** | Filter results where version matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedmavenMavenArtifactResponseList**](PaginatedmavenMavenArtifactResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentMavenArtifactRead

> MavenMavenArtifactResponse ContentMavenArtifactRead(ctx, mavenMavenArtifactHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a maven artifact



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
	mavenMavenArtifactHref := "mavenMavenArtifactHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentArtifactAPI.ContentMavenArtifactRead(context.Background(), mavenMavenArtifactHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentArtifactAPI.ContentMavenArtifactRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentMavenArtifactRead`: MavenMavenArtifactResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentArtifactAPI.ContentMavenArtifactRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenArtifactHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentMavenArtifactReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**MavenMavenArtifactResponse**](MavenMavenArtifactResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentMavenArtifactSetLabel

> SetLabelResponse ContentMavenArtifactSetLabel(ctx, mavenMavenArtifactHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	mavenMavenArtifactHref := "mavenMavenArtifactHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentArtifactAPI.ContentMavenArtifactSetLabel(context.Background(), mavenMavenArtifactHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentArtifactAPI.ContentMavenArtifactSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentMavenArtifactSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentArtifactAPI.ContentMavenArtifactSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenArtifactHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentMavenArtifactSetLabelRequest struct via the builder pattern


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


## ContentMavenArtifactUnsetLabel

> UnsetLabelResponse ContentMavenArtifactUnsetLabel(ctx, mavenMavenArtifactHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	mavenMavenArtifactHref := "mavenMavenArtifactHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentArtifactAPI.ContentMavenArtifactUnsetLabel(context.Background(), mavenMavenArtifactHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentArtifactAPI.ContentMavenArtifactUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentMavenArtifactUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentArtifactAPI.ContentMavenArtifactUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenArtifactHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentMavenArtifactUnsetLabelRequest struct via the builder pattern


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


## ContentMavenArtifactUpload

> MavenMavenArtifactResponse ContentMavenArtifactUpload(ctx, pulpDomain).RelativePath(relativePath).XTaskDiagnostics(xTaskDiagnostics).Repository(repository).Overwrite(overwrite).PulpLabels(pulpLabels).Artifact(artifact).File(file).Upload(upload).FileUrl(fileUrl).DownloaderConfig(downloaderConfig).Execute()

Upload a Maven artifact synchronously.



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
	resp, r, err := apiClient.ContentArtifactAPI.ContentMavenArtifactUpload(context.Background(), pulpDomain).RelativePath(relativePath).XTaskDiagnostics(xTaskDiagnostics).Repository(repository).Overwrite(overwrite).PulpLabels(pulpLabels).Artifact(artifact).File(file).Upload(upload).FileUrl(fileUrl).DownloaderConfig(downloaderConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentArtifactAPI.ContentMavenArtifactUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentMavenArtifactUpload`: MavenMavenArtifactResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentArtifactAPI.ContentMavenArtifactUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentMavenArtifactUploadRequest struct via the builder pattern


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

[**MavenMavenArtifactResponse**](MavenMavenArtifactResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

