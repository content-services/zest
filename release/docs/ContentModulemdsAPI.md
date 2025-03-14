# \ContentModulemdsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentRpmModulemdsCreate**](ContentModulemdsAPI.md#ContentRpmModulemdsCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/content/rpm/modulemds/ | Create a modulemd
[**ContentRpmModulemdsList**](ContentModulemdsAPI.md#ContentRpmModulemdsList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/rpm/modulemds/ | List modulemds
[**ContentRpmModulemdsRead**](ContentModulemdsAPI.md#ContentRpmModulemdsRead) | **Get** /{rpm_modulemd_href} | Inspect a modulemd
[**ContentRpmModulemdsSetLabel**](ContentModulemdsAPI.md#ContentRpmModulemdsSetLabel) | **Post** /{rpm_modulemd_href}set_label/ | Set a label
[**ContentRpmModulemdsUnsetLabel**](ContentModulemdsAPI.md#ContentRpmModulemdsUnsetLabel) | **Post** /{rpm_modulemd_href}unset_label/ | Unset a label



## ContentRpmModulemdsCreate

> AsyncOperationResponse ContentRpmModulemdsCreate(ctx, pulpDomain).RpmModulemd(rpmModulemd).Execute()

Create a modulemd



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
	rpmModulemd := *openapiclient.NewRpmModulemd("Name_example", "Stream_example", "Version_example", "Context_example", "Arch_example", interface{}(123), interface{}(123), "Snippet_example", interface{}(123), "Description_example") // RpmModulemd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentModulemdsAPI.ContentRpmModulemdsCreate(context.Background(), pulpDomain).RpmModulemd(rpmModulemd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentModulemdsAPI.ContentRpmModulemdsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentRpmModulemdsCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentModulemdsAPI.ContentRpmModulemdsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmModulemdsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rpmModulemd** | [**RpmModulemd**](RpmModulemd.md) |  | 

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


## ContentRpmModulemdsList

> PaginatedrpmModulemdResponseList ContentRpmModulemdsList(ctx, pulpDomain).Arch(arch).ArchIn(archIn).Context(context).ContextIn(contextIn).Limit(limit).Name(name).NameIn(nameIn).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Stream(stream).StreamIn(streamIn).Version(version).VersionIn(versionIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List modulemds



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
	arch := "arch_example" // string | Filter results where arch matches value (optional)
	archIn := []string{"Inner_example"} // []string | Filter results where arch is in a comma-separated list of values (optional)
	context := "context_example" // string | Filter results where context matches value (optional)
	contextIn := []string{"Inner_example"} // []string | Filter results where context is in a comma-separated list of values (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string | Filter results where name matches value (optional)
	nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `name` - Name* `-name` - Name (descending)* `stream` - Stream* `-stream` - Stream (descending)* `version` - Version* `-version` - Version (descending)* `context` - Context* `-context` - Context (descending)* `arch` - Arch* `-arch` - Arch (descending)* `static_context` - Static context* `-static_context` - Static context (descending)* `dependencies` - Dependencies* `-dependencies` - Dependencies (descending)* `artifacts` - Artifacts* `-artifacts` - Artifacts (descending)* `profiles` - Profiles* `-profiles` - Profiles (descending)* `description` - Description* `-description` - Description (descending)* `digest` - Digest* `-digest` - Digest (descending)* `snippet` - Snippet* `-snippet` - Snippet (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
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
	version := "version_example" // string | Filter results where version matches value (optional)
	versionIn := []string{"Inner_example"} // []string | Filter results where version is in a comma-separated list of values (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentModulemdsAPI.ContentRpmModulemdsList(context.Background(), pulpDomain).Arch(arch).ArchIn(archIn).Context(context).ContextIn(contextIn).Limit(limit).Name(name).NameIn(nameIn).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Stream(stream).StreamIn(streamIn).Version(version).VersionIn(versionIn).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentModulemdsAPI.ContentRpmModulemdsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentRpmModulemdsList`: PaginatedrpmModulemdResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentModulemdsAPI.ContentRpmModulemdsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmModulemdsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **arch** | **string** | Filter results where arch matches value | 
 **archIn** | **[]string** | Filter results where arch is in a comma-separated list of values | 
 **context** | **string** | Filter results where context matches value | 
 **contextIn** | **[]string** | Filter results where context is in a comma-separated list of values | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;stream&#x60; - Stream* &#x60;-stream&#x60; - Stream (descending)* &#x60;version&#x60; - Version* &#x60;-version&#x60; - Version (descending)* &#x60;context&#x60; - Context* &#x60;-context&#x60; - Context (descending)* &#x60;arch&#x60; - Arch* &#x60;-arch&#x60; - Arch (descending)* &#x60;static_context&#x60; - Static context* &#x60;-static_context&#x60; - Static context (descending)* &#x60;dependencies&#x60; - Dependencies* &#x60;-dependencies&#x60; - Dependencies (descending)* &#x60;artifacts&#x60; - Artifacts* &#x60;-artifacts&#x60; - Artifacts (descending)* &#x60;profiles&#x60; - Profiles* &#x60;-profiles&#x60; - Profiles (descending)* &#x60;description&#x60; - Description* &#x60;-description&#x60; - Description (descending)* &#x60;digest&#x60; - Digest* &#x60;-digest&#x60; - Digest (descending)* &#x60;snippet&#x60; - Snippet* &#x60;-snippet&#x60; - Snippet (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
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
 **version** | **string** | Filter results where version matches value | 
 **versionIn** | **[]string** | Filter results where version is in a comma-separated list of values | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedrpmModulemdResponseList**](PaginatedrpmModulemdResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentRpmModulemdsRead

> RpmModulemdResponse ContentRpmModulemdsRead(ctx, rpmModulemdHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a modulemd



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
	rpmModulemdHref := "rpmModulemdHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentModulemdsAPI.ContentRpmModulemdsRead(context.Background(), rpmModulemdHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentModulemdsAPI.ContentRpmModulemdsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentRpmModulemdsRead`: RpmModulemdResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentModulemdsAPI.ContentRpmModulemdsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmModulemdHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmModulemdsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RpmModulemdResponse**](RpmModulemdResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentRpmModulemdsSetLabel

> SetLabelResponse ContentRpmModulemdsSetLabel(ctx, rpmModulemdHref).SetLabel(setLabel).Execute()

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
	rpmModulemdHref := "rpmModulemdHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentModulemdsAPI.ContentRpmModulemdsSetLabel(context.Background(), rpmModulemdHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentModulemdsAPI.ContentRpmModulemdsSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentRpmModulemdsSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentModulemdsAPI.ContentRpmModulemdsSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmModulemdHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmModulemdsSetLabelRequest struct via the builder pattern


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


## ContentRpmModulemdsUnsetLabel

> UnsetLabelResponse ContentRpmModulemdsUnsetLabel(ctx, rpmModulemdHref).UnsetLabel(unsetLabel).Execute()

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
	rpmModulemdHref := "rpmModulemdHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentModulemdsAPI.ContentRpmModulemdsUnsetLabel(context.Background(), rpmModulemdHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentModulemdsAPI.ContentRpmModulemdsUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentRpmModulemdsUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentModulemdsAPI.ContentRpmModulemdsUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmModulemdHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmModulemdsUnsetLabelRequest struct via the builder pattern


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

