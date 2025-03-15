# \ContentOpenpgpPublicsubkeyAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentCoreOpenpgpPublicsubkeyList**](ContentOpenpgpPublicsubkeyAPI.md#ContentCoreOpenpgpPublicsubkeyList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/core/openpgp_publicsubkey/ | List open pgp public subkeys
[**ContentCoreOpenpgpPublicsubkeyRead**](ContentOpenpgpPublicsubkeyAPI.md#ContentCoreOpenpgpPublicsubkeyRead) | **Get** /{open_p_g_p_public_subkey_href} | Inspect an open pgp public subkey
[**ContentCoreOpenpgpPublicsubkeySetLabel**](ContentOpenpgpPublicsubkeyAPI.md#ContentCoreOpenpgpPublicsubkeySetLabel) | **Post** /{open_p_g_p_public_subkey_href}set_label/ | Set a label
[**ContentCoreOpenpgpPublicsubkeyUnsetLabel**](ContentOpenpgpPublicsubkeyAPI.md#ContentCoreOpenpgpPublicsubkeyUnsetLabel) | **Post** /{open_p_g_p_public_subkey_href}unset_label/ | Unset a label



## ContentCoreOpenpgpPublicsubkeyList

> PaginatedOpenPGPPublicSubkeyResponseList ContentCoreOpenpgpPublicsubkeyList(ctx, pulpDomain).Fingerprint(fingerprint).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List open pgp public subkeys



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
	fingerprint := "fingerprint_example" // string | Filter results where fingerprint matches value (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `raw_data` - Raw data* `-raw_data` - Raw data (descending)* `fingerprint` - Fingerprint* `-fingerprint` - Fingerprint (descending)* `created` - Created* `-created` - Created (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF/PRN (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentOpenpgpPublicsubkeyAPI.ContentCoreOpenpgpPublicsubkeyList(context.Background(), pulpDomain).Fingerprint(fingerprint).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentOpenpgpPublicsubkeyAPI.ContentCoreOpenpgpPublicsubkeyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentCoreOpenpgpPublicsubkeyList`: PaginatedOpenPGPPublicSubkeyResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentOpenpgpPublicsubkeyAPI.ContentCoreOpenpgpPublicsubkeyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentCoreOpenpgpPublicsubkeyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fingerprint** | **string** | Filter results where fingerprint matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;raw_data&#x60; - Raw data* &#x60;-raw_data&#x60; - Raw data (descending)* &#x60;fingerprint&#x60; - Fingerprint* &#x60;-fingerprint&#x60; - Fingerprint (descending)* &#x60;created&#x60; - Created* &#x60;-created&#x60; - Created (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF/PRN | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedOpenPGPPublicSubkeyResponseList**](PaginatedOpenPGPPublicSubkeyResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentCoreOpenpgpPublicsubkeyRead

> OpenPGPPublicSubkeyResponse ContentCoreOpenpgpPublicsubkeyRead(ctx, openPGPPublicSubkeyHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an open pgp public subkey



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
	openPGPPublicSubkeyHref := "openPGPPublicSubkeyHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentOpenpgpPublicsubkeyAPI.ContentCoreOpenpgpPublicsubkeyRead(context.Background(), openPGPPublicSubkeyHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentOpenpgpPublicsubkeyAPI.ContentCoreOpenpgpPublicsubkeyRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentCoreOpenpgpPublicsubkeyRead`: OpenPGPPublicSubkeyResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentOpenpgpPublicsubkeyAPI.ContentCoreOpenpgpPublicsubkeyRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPPublicSubkeyHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentCoreOpenpgpPublicsubkeyReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**OpenPGPPublicSubkeyResponse**](OpenPGPPublicSubkeyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentCoreOpenpgpPublicsubkeySetLabel

> SetLabelResponse ContentCoreOpenpgpPublicsubkeySetLabel(ctx, openPGPPublicSubkeyHref).SetLabel(setLabel).Execute()

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
	openPGPPublicSubkeyHref := "openPGPPublicSubkeyHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentOpenpgpPublicsubkeyAPI.ContentCoreOpenpgpPublicsubkeySetLabel(context.Background(), openPGPPublicSubkeyHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentOpenpgpPublicsubkeyAPI.ContentCoreOpenpgpPublicsubkeySetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentCoreOpenpgpPublicsubkeySetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentOpenpgpPublicsubkeyAPI.ContentCoreOpenpgpPublicsubkeySetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPPublicSubkeyHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentCoreOpenpgpPublicsubkeySetLabelRequest struct via the builder pattern


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


## ContentCoreOpenpgpPublicsubkeyUnsetLabel

> UnsetLabelResponse ContentCoreOpenpgpPublicsubkeyUnsetLabel(ctx, openPGPPublicSubkeyHref).UnsetLabel(unsetLabel).Execute()

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
	openPGPPublicSubkeyHref := "openPGPPublicSubkeyHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentOpenpgpPublicsubkeyAPI.ContentCoreOpenpgpPublicsubkeyUnsetLabel(context.Background(), openPGPPublicSubkeyHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentOpenpgpPublicsubkeyAPI.ContentCoreOpenpgpPublicsubkeyUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentCoreOpenpgpPublicsubkeyUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentOpenpgpPublicsubkeyAPI.ContentCoreOpenpgpPublicsubkeyUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPPublicSubkeyHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentCoreOpenpgpPublicsubkeyUnsetLabelRequest struct via the builder pattern


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

