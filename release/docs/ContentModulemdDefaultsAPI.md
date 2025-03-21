# \ContentModulemdDefaultsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentRpmModulemdDefaultsCreate**](ContentModulemdDefaultsAPI.md#ContentRpmModulemdDefaultsCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/content/rpm/modulemd_defaults/ | Create a modulemd defaults
[**ContentRpmModulemdDefaultsList**](ContentModulemdDefaultsAPI.md#ContentRpmModulemdDefaultsList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/rpm/modulemd_defaults/ | List modulemd defaultss
[**ContentRpmModulemdDefaultsRead**](ContentModulemdDefaultsAPI.md#ContentRpmModulemdDefaultsRead) | **Get** /{rpm_modulemd_defaults_href} | Inspect a modulemd defaults
[**ContentRpmModulemdDefaultsSetLabel**](ContentModulemdDefaultsAPI.md#ContentRpmModulemdDefaultsSetLabel) | **Post** /{rpm_modulemd_defaults_href}set_label/ | Set a label
[**ContentRpmModulemdDefaultsUnsetLabel**](ContentModulemdDefaultsAPI.md#ContentRpmModulemdDefaultsUnsetLabel) | **Post** /{rpm_modulemd_defaults_href}unset_label/ | Unset a label



## ContentRpmModulemdDefaultsCreate

> AsyncOperationResponse ContentRpmModulemdDefaultsCreate(ctx, pulpDomain).RpmModulemdDefaults(rpmModulemdDefaults).Execute()

Create a modulemd defaults



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	rpmModulemdDefaults := *openapiclient.NewRpmModulemdDefaults("Module_example", "Stream_example", interface{}(123), "Snippet_example") // RpmModulemdDefaults | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentModulemdDefaultsAPI.ContentRpmModulemdDefaultsCreate(context.Background(), pulpDomain).RpmModulemdDefaults(rpmModulemdDefaults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentModulemdDefaultsAPI.ContentRpmModulemdDefaultsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentRpmModulemdDefaultsCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentModulemdDefaultsAPI.ContentRpmModulemdDefaultsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmModulemdDefaultsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rpmModulemdDefaults** | [**RpmModulemdDefaults**](RpmModulemdDefaults.md) |  | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentRpmModulemdDefaultsList

> PaginatedrpmModulemdDefaultsResponseList ContentRpmModulemdDefaultsList(ctx, pulpDomain).Limit(limit).Module(module).ModuleIn(moduleIn).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Stream(stream).StreamIn(streamIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List modulemd defaultss



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	module := "module_example" // string | Filter results where module matches value (optional)
	moduleIn := []string{"Inner_example"} // []string | Filter results where module is in a comma-separated list of values (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `module` - Module* `-module` - Module (descending)* `stream` - Stream* `-stream` - Stream (descending)* `profiles` - Profiles* `-profiles` - Profiles (descending)* `digest` - Digest* `-digest` - Digest (descending)* `snippet` - Snippet* `-snippet` - Snippet (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF/PRN (optional)
	sha256 := "sha256_example" // string |  (optional)
	stream := "stream_example" // string | Filter results where stream matches value (optional)
	streamIn := []string{"Inner_example"} // []string | Filter results where stream is in a comma-separated list of values (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentModulemdDefaultsAPI.ContentRpmModulemdDefaultsList(context.Background(), pulpDomain).Limit(limit).Module(module).ModuleIn(moduleIn).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Stream(stream).StreamIn(streamIn).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentModulemdDefaultsAPI.ContentRpmModulemdDefaultsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentRpmModulemdDefaultsList`: PaginatedrpmModulemdDefaultsResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentModulemdDefaultsAPI.ContentRpmModulemdDefaultsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmModulemdDefaultsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **module** | **string** | Filter results where module matches value | 
 **moduleIn** | **[]string** | Filter results where module is in a comma-separated list of values | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;module&#x60; - Module* &#x60;-module&#x60; - Module (descending)* &#x60;stream&#x60; - Stream* &#x60;-stream&#x60; - Stream (descending)* &#x60;profiles&#x60; - Profiles* &#x60;-profiles&#x60; - Profiles (descending)* &#x60;digest&#x60; - Digest* &#x60;-digest&#x60; - Digest (descending)* &#x60;snippet&#x60; - Snippet* &#x60;-snippet&#x60; - Snippet (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF/PRN | 
 **sha256** | **string** |  | 
 **stream** | **string** | Filter results where stream matches value | 
 **streamIn** | **[]string** | Filter results where stream is in a comma-separated list of values | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedrpmModulemdDefaultsResponseList**](PaginatedrpmModulemdDefaultsResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentRpmModulemdDefaultsRead

> RpmModulemdDefaultsResponse ContentRpmModulemdDefaultsRead(ctx, rpmModulemdDefaultsHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a modulemd defaults



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	rpmModulemdDefaultsHref := "rpmModulemdDefaultsHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentModulemdDefaultsAPI.ContentRpmModulemdDefaultsRead(context.Background(), rpmModulemdDefaultsHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentModulemdDefaultsAPI.ContentRpmModulemdDefaultsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentRpmModulemdDefaultsRead`: RpmModulemdDefaultsResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentModulemdDefaultsAPI.ContentRpmModulemdDefaultsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmModulemdDefaultsHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmModulemdDefaultsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RpmModulemdDefaultsResponse**](RpmModulemdDefaultsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentRpmModulemdDefaultsSetLabel

> SetLabelResponse ContentRpmModulemdDefaultsSetLabel(ctx, rpmModulemdDefaultsHref).SetLabel(setLabel).Execute()

Set a label



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	rpmModulemdDefaultsHref := "rpmModulemdDefaultsHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentModulemdDefaultsAPI.ContentRpmModulemdDefaultsSetLabel(context.Background(), rpmModulemdDefaultsHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentModulemdDefaultsAPI.ContentRpmModulemdDefaultsSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentRpmModulemdDefaultsSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentModulemdDefaultsAPI.ContentRpmModulemdDefaultsSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmModulemdDefaultsHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmModulemdDefaultsSetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setLabel** | [**SetLabel**](SetLabel.md) |  | 

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


## ContentRpmModulemdDefaultsUnsetLabel

> UnsetLabelResponse ContentRpmModulemdDefaultsUnsetLabel(ctx, rpmModulemdDefaultsHref).UnsetLabel(unsetLabel).Execute()

Unset a label



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	rpmModulemdDefaultsHref := "rpmModulemdDefaultsHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentModulemdDefaultsAPI.ContentRpmModulemdDefaultsUnsetLabel(context.Background(), rpmModulemdDefaultsHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentModulemdDefaultsAPI.ContentRpmModulemdDefaultsUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentRpmModulemdDefaultsUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentModulemdDefaultsAPI.ContentRpmModulemdDefaultsUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmModulemdDefaultsHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmModulemdDefaultsUnsetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unsetLabel** | [**UnsetLabel**](UnsetLabel.md) |  | 

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

