# \ContentHuggingFaceAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentHuggingFaceHuggingFaceCreate**](ContentHuggingFaceAPI.md#ContentHuggingFaceHuggingFaceCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/content/hugging_face/hugging-face/ | Create a hugging face content
[**ContentHuggingFaceHuggingFaceList**](ContentHuggingFaceAPI.md#ContentHuggingFaceHuggingFaceList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/hugging_face/hugging-face/ | List hugging face contents
[**ContentHuggingFaceHuggingFaceRead**](ContentHuggingFaceAPI.md#ContentHuggingFaceHuggingFaceRead) | **Get** /{hugging_face_hugging_face_content_href} | Inspect a hugging face content
[**ContentHuggingFaceHuggingFaceSetLabel**](ContentHuggingFaceAPI.md#ContentHuggingFaceHuggingFaceSetLabel) | **Post** /{hugging_face_hugging_face_content_href}set_label/ | Set a label
[**ContentHuggingFaceHuggingFaceUnsetLabel**](ContentHuggingFaceAPI.md#ContentHuggingFaceHuggingFaceUnsetLabel) | **Post** /{hugging_face_hugging_face_content_href}unset_label/ | Unset a label



## ContentHuggingFaceHuggingFaceCreate

> HuggingFaceHuggingFaceContentResponse ContentHuggingFaceHuggingFaceCreate(ctx, pulpDomain).HuggingFaceHuggingFaceContent(huggingFaceHuggingFaceContent).XTaskDiagnostics(xTaskDiagnostics).Execute()

Create a hugging face content



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
	huggingFaceHuggingFaceContent := *openapiclient.NewHuggingFaceHuggingFaceContent("Artifact_example", "RelativePath_example", "RepoId_example") // HuggingFaceHuggingFaceContent | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentHuggingFaceAPI.ContentHuggingFaceHuggingFaceCreate(context.Background(), pulpDomain).HuggingFaceHuggingFaceContent(huggingFaceHuggingFaceContent).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentHuggingFaceAPI.ContentHuggingFaceHuggingFaceCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentHuggingFaceHuggingFaceCreate`: HuggingFaceHuggingFaceContentResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentHuggingFaceAPI.ContentHuggingFaceHuggingFaceCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentHuggingFaceHuggingFaceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **huggingFaceHuggingFaceContent** | [**HuggingFaceHuggingFaceContent**](HuggingFaceHuggingFaceContent.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**HuggingFaceHuggingFaceContentResponse**](HuggingFaceHuggingFaceContentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentHuggingFaceHuggingFaceList

> PaginatedhuggingFaceHuggingFaceContentResponseList ContentHuggingFaceHuggingFaceList(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RelativePath(relativePath).RepoId(repoId).RepoType(repoType).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Revision(revision).Fields(fields).ExcludeFields(excludeFields).Execute()

List hugging face contents



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
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `repo_id` - Repo id* `-repo_id` - Repo id (descending)* `repo_type` - Repo type* `-repo_type` - Repo type (descending)* `relative_path` - Relative path* `-relative_path` - Relative path (descending)* `revision` - Revision* `-revision` - Revision (descending)* `size` - Size* `-size` - Size (descending)* `etag` - Etag* `-etag` - Etag (descending)* `last_modified` - Last modified* `-last_modified` - Last modified (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	relativePath := "relativePath_example" // string | Filter results where relative_path matches value (optional)
	repoId := "repoId_example" // string | Filter results where repo_id matches value (optional)
	repoType := "repoType_example" // string | Filter results where repo_type matches value* `models` - Models* `datasets` - Datasets* `spaces` - Spaces (optional)
	repositoryVersion := "repositoryVersion_example" // string |  (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string |  (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string |  (optional)
	revision := "revision_example" // string | Filter results where revision matches value (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentHuggingFaceAPI.ContentHuggingFaceHuggingFaceList(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RelativePath(relativePath).RepoId(repoId).RepoType(repoType).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Revision(revision).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentHuggingFaceAPI.ContentHuggingFaceHuggingFaceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentHuggingFaceHuggingFaceList`: PaginatedhuggingFaceHuggingFaceContentResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentHuggingFaceAPI.ContentHuggingFaceHuggingFaceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentHuggingFaceHuggingFaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;repo_id&#x60; - Repo id* &#x60;-repo_id&#x60; - Repo id (descending)* &#x60;repo_type&#x60; - Repo type* &#x60;-repo_type&#x60; - Repo type (descending)* &#x60;relative_path&#x60; - Relative path* &#x60;-relative_path&#x60; - Relative path (descending)* &#x60;revision&#x60; - Revision* &#x60;-revision&#x60; - Revision (descending)* &#x60;size&#x60; - Size* &#x60;-size&#x60; - Size (descending)* &#x60;etag&#x60; - Etag* &#x60;-etag&#x60; - Etag (descending)* &#x60;last_modified&#x60; - Last modified* &#x60;-last_modified&#x60; - Last modified (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **relativePath** | **string** | Filter results where relative_path matches value | 
 **repoId** | **string** | Filter results where repo_id matches value | 
 **repoType** | **string** | Filter results where repo_type matches value* &#x60;models&#x60; - Models* &#x60;datasets&#x60; - Datasets* &#x60;spaces&#x60; - Spaces | 
 **repositoryVersion** | **string** |  | 
 **repositoryVersionAdded** | **string** |  | 
 **repositoryVersionRemoved** | **string** |  | 
 **revision** | **string** | Filter results where revision matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedhuggingFaceHuggingFaceContentResponseList**](PaginatedhuggingFaceHuggingFaceContentResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentHuggingFaceHuggingFaceRead

> HuggingFaceHuggingFaceContentResponse ContentHuggingFaceHuggingFaceRead(ctx, huggingFaceHuggingFaceContentHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a hugging face content



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
	huggingFaceHuggingFaceContentHref := "huggingFaceHuggingFaceContentHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentHuggingFaceAPI.ContentHuggingFaceHuggingFaceRead(context.Background(), huggingFaceHuggingFaceContentHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentHuggingFaceAPI.ContentHuggingFaceHuggingFaceRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentHuggingFaceHuggingFaceRead`: HuggingFaceHuggingFaceContentResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentHuggingFaceAPI.ContentHuggingFaceHuggingFaceRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceContentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentHuggingFaceHuggingFaceReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**HuggingFaceHuggingFaceContentResponse**](HuggingFaceHuggingFaceContentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentHuggingFaceHuggingFaceSetLabel

> SetLabelResponse ContentHuggingFaceHuggingFaceSetLabel(ctx, huggingFaceHuggingFaceContentHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	huggingFaceHuggingFaceContentHref := "huggingFaceHuggingFaceContentHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentHuggingFaceAPI.ContentHuggingFaceHuggingFaceSetLabel(context.Background(), huggingFaceHuggingFaceContentHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentHuggingFaceAPI.ContentHuggingFaceHuggingFaceSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentHuggingFaceHuggingFaceSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentHuggingFaceAPI.ContentHuggingFaceHuggingFaceSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceContentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentHuggingFaceHuggingFaceSetLabelRequest struct via the builder pattern


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


## ContentHuggingFaceHuggingFaceUnsetLabel

> UnsetLabelResponse ContentHuggingFaceHuggingFaceUnsetLabel(ctx, huggingFaceHuggingFaceContentHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	huggingFaceHuggingFaceContentHref := "huggingFaceHuggingFaceContentHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentHuggingFaceAPI.ContentHuggingFaceHuggingFaceUnsetLabel(context.Background(), huggingFaceHuggingFaceContentHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentHuggingFaceAPI.ContentHuggingFaceHuggingFaceUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentHuggingFaceHuggingFaceUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentHuggingFaceAPI.ContentHuggingFaceHuggingFaceUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceContentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentHuggingFaceHuggingFaceUnsetLabelRequest struct via the builder pattern


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

