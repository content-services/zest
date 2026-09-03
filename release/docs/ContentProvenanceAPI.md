# \ContentProvenanceAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentPythonProvenanceCreate**](ContentProvenanceAPI.md#ContentPythonProvenanceCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/content/python/provenance/ | Create a package provenance
[**ContentPythonProvenanceList**](ContentProvenanceAPI.md#ContentPythonProvenanceList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/python/provenance/ | List package provenances
[**ContentPythonProvenanceRead**](ContentProvenanceAPI.md#ContentPythonProvenanceRead) | **Get** /{python_package_provenance_href} | Inspect a package provenance
[**ContentPythonProvenanceSetLabel**](ContentProvenanceAPI.md#ContentPythonProvenanceSetLabel) | **Post** /{python_package_provenance_href}set_label/ | Set a label
[**ContentPythonProvenanceUnsetLabel**](ContentProvenanceAPI.md#ContentPythonProvenanceUnsetLabel) | **Post** /{python_package_provenance_href}unset_label/ | Unset a label



## ContentPythonProvenanceCreate

> AsyncOperationResponse ContentPythonProvenanceCreate(ctx, pulpDomain).Package_(package_).XTaskDiagnostics(xTaskDiagnostics).Repository(repository).Overwrite(overwrite).PulpLabels(pulpLabels).File(file).Upload(upload).FileUrl(fileUrl).DownloaderConfig(downloaderConfig).Verify(verify).Execute()

Create a package provenance



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
	package_ := "package__example" // string | The package that the provenance is for.
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
	overwrite := true // bool | When set to true, existing content in the repository with the same unique key will be silently overwritten. When set to false, the task will fail if content would be overwritten. Only used when 'repository' is specified. Defaults to true. (optional)
	pulpLabels := map[string]*string{"key": "Inner_example"} // map[string]*string | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the content unit. (optional)
	upload := "upload_example" // string | An uncommitted upload that may be turned into the content unit. (optional)
	fileUrl := "fileUrl_example" // string | A url that Pulp can download and turn into the content unit. (optional)
	downloaderConfig := *openapiclient.NewRemoteNetworkConfig() // RemoteNetworkConfig | Configuration for the download process (e.g., proxies, auth, timeouts). Only applicable when providing a 'file_url. (optional)
	verify := true // bool | Verify each attestation in the provenance. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentProvenanceAPI.ContentPythonProvenanceCreate(context.Background(), pulpDomain).Package_(package_).XTaskDiagnostics(xTaskDiagnostics).Repository(repository).Overwrite(overwrite).PulpLabels(pulpLabels).File(file).Upload(upload).FileUrl(fileUrl).DownloaderConfig(downloaderConfig).Verify(verify).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentProvenanceAPI.ContentPythonProvenanceCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentPythonProvenanceCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentProvenanceAPI.ContentPythonProvenanceCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentPythonProvenanceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **package_** | **string** | The package that the provenance is for. | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **overwrite** | **bool** | When set to true, existing content in the repository with the same unique key will be silently overwritten. When set to false, the task will fail if content would be overwritten. Only used when &#39;repository&#39; is specified. Defaults to true. | 
 **pulpLabels** | **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | 
 **file** | ***os.File** | An uploaded file that may be turned into the content unit. | 
 **upload** | **string** | An uncommitted upload that may be turned into the content unit. | 
 **fileUrl** | **string** | A url that Pulp can download and turn into the content unit. | 
 **downloaderConfig** | [**RemoteNetworkConfig**](RemoteNetworkConfig.md) | Configuration for the download process (e.g., proxies, auth, timeouts). Only applicable when providing a &#39;file_url. | 
 **verify** | **bool** | Verify each attestation in the provenance. | [default to true]

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


## ContentPythonProvenanceList

> PaginatedpythonPackageProvenanceResponseList ContentPythonProvenanceList(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).Package_(package_).PackageIn(packageIn).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Sha256In(sha256In).Fields(fields).ExcludeFields(excludeFields).Execute()

List package provenances



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
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `provenance` - Provenance* `-provenance` - Provenance (descending)* `sha256` - Sha256* `-sha256` - Sha256 (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	package_ := "package__example" // string |  (optional)
	packageIn := []string{"Inner_example"} // []string |  (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repositoryVersion := "repositoryVersion_example" // string |  (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string |  (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string |  (optional)
	sha256 := "sha256_example" // string | Filter results where sha256 matches value (optional)
	sha256In := []string{"Inner_example"} // []string | Filter results where sha256 is in a comma-separated list of values (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentProvenanceAPI.ContentPythonProvenanceList(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).Package_(package_).PackageIn(packageIn).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Sha256In(sha256In).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentProvenanceAPI.ContentPythonProvenanceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentPythonProvenanceList`: PaginatedpythonPackageProvenanceResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentProvenanceAPI.ContentPythonProvenanceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentPythonProvenanceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;provenance&#x60; - Provenance* &#x60;-provenance&#x60; - Provenance (descending)* &#x60;sha256&#x60; - Sha256* &#x60;-sha256&#x60; - Sha256 (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **package_** | **string** |  | 
 **packageIn** | **[]string** |  | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repositoryVersion** | **string** |  | 
 **repositoryVersionAdded** | **string** |  | 
 **repositoryVersionRemoved** | **string** |  | 
 **sha256** | **string** | Filter results where sha256 matches value | 
 **sha256In** | **[]string** | Filter results where sha256 is in a comma-separated list of values | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedpythonPackageProvenanceResponseList**](PaginatedpythonPackageProvenanceResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentPythonProvenanceRead

> PythonPackageProvenanceResponse ContentPythonProvenanceRead(ctx, pythonPackageProvenanceHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a package provenance



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
	pythonPackageProvenanceHref := "pythonPackageProvenanceHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentProvenanceAPI.ContentPythonProvenanceRead(context.Background(), pythonPackageProvenanceHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentProvenanceAPI.ContentPythonProvenanceRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentPythonProvenanceRead`: PythonPackageProvenanceResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentProvenanceAPI.ContentPythonProvenanceRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pythonPackageProvenanceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentPythonProvenanceReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PythonPackageProvenanceResponse**](PythonPackageProvenanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentPythonProvenanceSetLabel

> SetLabelResponse ContentPythonProvenanceSetLabel(ctx, pythonPackageProvenanceHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	pythonPackageProvenanceHref := "pythonPackageProvenanceHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentProvenanceAPI.ContentPythonProvenanceSetLabel(context.Background(), pythonPackageProvenanceHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentProvenanceAPI.ContentPythonProvenanceSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentPythonProvenanceSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentProvenanceAPI.ContentPythonProvenanceSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pythonPackageProvenanceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentPythonProvenanceSetLabelRequest struct via the builder pattern


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


## ContentPythonProvenanceUnsetLabel

> UnsetLabelResponse ContentPythonProvenanceUnsetLabel(ctx, pythonPackageProvenanceHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	pythonPackageProvenanceHref := "pythonPackageProvenanceHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentProvenanceAPI.ContentPythonProvenanceUnsetLabel(context.Background(), pythonPackageProvenanceHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentProvenanceAPI.ContentPythonProvenanceUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentPythonProvenanceUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentProvenanceAPI.ContentPythonProvenanceUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pythonPackageProvenanceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentPythonProvenanceUnsetLabelRequest struct via the builder pattern


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

