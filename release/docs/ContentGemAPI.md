# \ContentGemAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentGemGemCreate**](ContentGemAPI.md#ContentGemGemCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/content/gem/gem/ | Create a gem content
[**ContentGemGemList**](ContentGemAPI.md#ContentGemGemList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/gem/gem/ | List gem contents
[**ContentGemGemRead**](ContentGemAPI.md#ContentGemGemRead) | **Get** /{gem_gem_content_href} | Inspect a gem content
[**ContentGemGemSetLabel**](ContentGemAPI.md#ContentGemGemSetLabel) | **Post** /{gem_gem_content_href}set_label/ | Set a label
[**ContentGemGemUnsetLabel**](ContentGemAPI.md#ContentGemGemUnsetLabel) | **Post** /{gem_gem_content_href}unset_label/ | Unset a label



## ContentGemGemCreate

> AsyncOperationResponse ContentGemGemCreate(ctx, pulpDomain).Repository(repository).PulpLabels(pulpLabels).Artifact(artifact).File(file).Execute()

Create a gem content



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
	repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
	pulpLabels := map[string]string{"key": "Inner_example"} // map[string]string | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. (optional)
	artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
	file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that should be turned into the artifact of the content unit. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentGemAPI.ContentGemGemCreate(context.Background(), pulpDomain).Repository(repository).PulpLabels(pulpLabels).Artifact(artifact).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentGemAPI.ContentGemGemCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentGemGemCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentGemAPI.ContentGemGemCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentGemGemCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **pulpLabels** | **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | 
 **artifact** | **string** | Artifact file representing the physical content | 
 **file** | ***os.File** | An uploaded file that should be turned into the artifact of the content unit. | 

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


## ContentGemGemList

> PaginatedgemGemContentResponseList ContentGemGemList(ctx, pulpDomain).Checksum(checksum).Limit(limit).Name(name).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).Prerelease(prerelease).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()

List gem contents



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
	checksum := "checksum_example" // string | Filter results where checksum matches value (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string | Filter results where name matches value (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `name` - Name* `-name` - Name (descending)* `version` - Version* `-version` - Version (descending)* `platform` - Platform* `-platform` - Platform (descending)* `checksum` - Checksum* `-checksum` - Checksum (descending)* `prerelease` - Prerelease* `-prerelease` - Prerelease (descending)* `dependencies` - Dependencies* `-dependencies` - Dependencies (descending)* `required_ruby_version` - Required ruby version* `-required_ruby_version` - Required ruby version (descending)* `required_rubygems_version` - Required rubygems version* `-required_rubygems_version` - Required rubygems version (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	prerelease := true // bool | Filter results where prerelease matches value (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF/PRN (optional)
	version := "version_example" // string | Filter results where version matches value (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentGemAPI.ContentGemGemList(context.Background(), pulpDomain).Checksum(checksum).Limit(limit).Name(name).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).Prerelease(prerelease).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentGemAPI.ContentGemGemList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentGemGemList`: PaginatedgemGemContentResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentGemAPI.ContentGemGemList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentGemGemListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checksum** | **string** | Filter results where checksum matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;version&#x60; - Version* &#x60;-version&#x60; - Version (descending)* &#x60;platform&#x60; - Platform* &#x60;-platform&#x60; - Platform (descending)* &#x60;checksum&#x60; - Checksum* &#x60;-checksum&#x60; - Checksum (descending)* &#x60;prerelease&#x60; - Prerelease* &#x60;-prerelease&#x60; - Prerelease (descending)* &#x60;dependencies&#x60; - Dependencies* &#x60;-dependencies&#x60; - Dependencies (descending)* &#x60;required_ruby_version&#x60; - Required ruby version* &#x60;-required_ruby_version&#x60; - Required ruby version (descending)* &#x60;required_rubygems_version&#x60; - Required rubygems version* &#x60;-required_rubygems_version&#x60; - Required rubygems version (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **prerelease** | **bool** | Filter results where prerelease matches value | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF/PRN | 
 **version** | **string** | Filter results where version matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedgemGemContentResponseList**](PaginatedgemGemContentResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentGemGemRead

> GemGemContentResponse ContentGemGemRead(ctx, gemGemContentHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a gem content



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
	gemGemContentHref := "gemGemContentHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentGemAPI.ContentGemGemRead(context.Background(), gemGemContentHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentGemAPI.ContentGemGemRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentGemGemRead`: GemGemContentResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentGemAPI.ContentGemGemRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemContentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentGemGemReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**GemGemContentResponse**](GemGemContentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentGemGemSetLabel

> SetLabelResponse ContentGemGemSetLabel(ctx, gemGemContentHref).SetLabel(setLabel).Execute()

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
	gemGemContentHref := "gemGemContentHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentGemAPI.ContentGemGemSetLabel(context.Background(), gemGemContentHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentGemAPI.ContentGemGemSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentGemGemSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentGemAPI.ContentGemGemSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemContentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentGemGemSetLabelRequest struct via the builder pattern


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


## ContentGemGemUnsetLabel

> UnsetLabelResponse ContentGemGemUnsetLabel(ctx, gemGemContentHref).UnsetLabel(unsetLabel).Execute()

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
	gemGemContentHref := "gemGemContentHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentGemAPI.ContentGemGemUnsetLabel(context.Background(), gemGemContentHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentGemAPI.ContentGemGemUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentGemGemUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentGemAPI.ContentGemGemUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemContentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentGemGemUnsetLabelRequest struct via the builder pattern


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

