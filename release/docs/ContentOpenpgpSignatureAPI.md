# \ContentOpenpgpSignatureAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentCoreOpenpgpSignatureList**](ContentOpenpgpSignatureAPI.md#ContentCoreOpenpgpSignatureList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/core/openpgp_signature/ | List open pgp signatures
[**ContentCoreOpenpgpSignatureRead**](ContentOpenpgpSignatureAPI.md#ContentCoreOpenpgpSignatureRead) | **Get** /{open_p_g_p_signature_href} | Inspect an open pgp signature
[**ContentCoreOpenpgpSignatureSetLabel**](ContentOpenpgpSignatureAPI.md#ContentCoreOpenpgpSignatureSetLabel) | **Post** /{open_p_g_p_signature_href}set_label/ | Set a label
[**ContentCoreOpenpgpSignatureUnsetLabel**](ContentOpenpgpSignatureAPI.md#ContentCoreOpenpgpSignatureUnsetLabel) | **Post** /{open_p_g_p_signature_href}unset_label/ | Unset a label



## ContentCoreOpenpgpSignatureList

> PaginatedOpenPGPSignatureResponseList ContentCoreOpenpgpSignatureList(ctx, pulpDomain).Issuer(issuer).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List open pgp signatures



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
	issuer := "issuer_example" // string | Filter results where issuer matches value (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `raw_data` - Raw data* `-raw_data` - Raw data (descending)* `sha256` - Sha256* `-sha256` - Sha256 (descending)* `signature_type` - Signature type* `-signature_type` - Signature type (descending)* `created` - Created* `-created` - Created (descending)* `expiration_time` - Expiration time* `-expiration_time` - Expiration time (descending)* `key_expiration_time` - Key expiration time* `-key_expiration_time` - Key expiration time (descending)* `issuer` - Issuer* `-issuer` - Issuer (descending)* `signers_user_id` - Signers user id* `-signers_user_id` - Signers user id (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
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
	resp, r, err := apiClient.ContentOpenpgpSignatureAPI.ContentCoreOpenpgpSignatureList(context.Background(), pulpDomain).Issuer(issuer).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentOpenpgpSignatureAPI.ContentCoreOpenpgpSignatureList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentCoreOpenpgpSignatureList`: PaginatedOpenPGPSignatureResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentOpenpgpSignatureAPI.ContentCoreOpenpgpSignatureList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentCoreOpenpgpSignatureListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issuer** | **string** | Filter results where issuer matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;raw_data&#x60; - Raw data* &#x60;-raw_data&#x60; - Raw data (descending)* &#x60;sha256&#x60; - Sha256* &#x60;-sha256&#x60; - Sha256 (descending)* &#x60;signature_type&#x60; - Signature type* &#x60;-signature_type&#x60; - Signature type (descending)* &#x60;created&#x60; - Created* &#x60;-created&#x60; - Created (descending)* &#x60;expiration_time&#x60; - Expiration time* &#x60;-expiration_time&#x60; - Expiration time (descending)* &#x60;key_expiration_time&#x60; - Key expiration time* &#x60;-key_expiration_time&#x60; - Key expiration time (descending)* &#x60;issuer&#x60; - Issuer* &#x60;-issuer&#x60; - Issuer (descending)* &#x60;signers_user_id&#x60; - Signers user id* &#x60;-signers_user_id&#x60; - Signers user id (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
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

[**PaginatedOpenPGPSignatureResponseList**](PaginatedOpenPGPSignatureResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentCoreOpenpgpSignatureRead

> OpenPGPSignatureResponse ContentCoreOpenpgpSignatureRead(ctx, openPGPSignatureHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an open pgp signature



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
	openPGPSignatureHref := "openPGPSignatureHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentOpenpgpSignatureAPI.ContentCoreOpenpgpSignatureRead(context.Background(), openPGPSignatureHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentOpenpgpSignatureAPI.ContentCoreOpenpgpSignatureRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentCoreOpenpgpSignatureRead`: OpenPGPSignatureResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentOpenpgpSignatureAPI.ContentCoreOpenpgpSignatureRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPSignatureHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentCoreOpenpgpSignatureReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**OpenPGPSignatureResponse**](OpenPGPSignatureResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentCoreOpenpgpSignatureSetLabel

> SetLabelResponse ContentCoreOpenpgpSignatureSetLabel(ctx, openPGPSignatureHref).SetLabel(setLabel).Execute()

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
	openPGPSignatureHref := "openPGPSignatureHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentOpenpgpSignatureAPI.ContentCoreOpenpgpSignatureSetLabel(context.Background(), openPGPSignatureHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentOpenpgpSignatureAPI.ContentCoreOpenpgpSignatureSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentCoreOpenpgpSignatureSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentOpenpgpSignatureAPI.ContentCoreOpenpgpSignatureSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPSignatureHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentCoreOpenpgpSignatureSetLabelRequest struct via the builder pattern


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


## ContentCoreOpenpgpSignatureUnsetLabel

> UnsetLabelResponse ContentCoreOpenpgpSignatureUnsetLabel(ctx, openPGPSignatureHref).UnsetLabel(unsetLabel).Execute()

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
	openPGPSignatureHref := "openPGPSignatureHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentOpenpgpSignatureAPI.ContentCoreOpenpgpSignatureUnsetLabel(context.Background(), openPGPSignatureHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentOpenpgpSignatureAPI.ContentCoreOpenpgpSignatureUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentCoreOpenpgpSignatureUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentOpenpgpSignatureAPI.ContentCoreOpenpgpSignatureUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPSignatureHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentCoreOpenpgpSignatureUnsetLabelRequest struct via the builder pattern


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

