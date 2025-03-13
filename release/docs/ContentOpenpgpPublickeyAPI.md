# \ContentOpenpgpPublickeyAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentCoreOpenpgpPublickeyCreate**](ContentOpenpgpPublickeyAPI.md#ContentCoreOpenpgpPublickeyCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/content/core/openpgp_publickey/ | Create an open pgp public key
[**ContentCoreOpenpgpPublickeyList**](ContentOpenpgpPublickeyAPI.md#ContentCoreOpenpgpPublickeyList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/core/openpgp_publickey/ | List open pgp public keys
[**ContentCoreOpenpgpPublickeyRead**](ContentOpenpgpPublickeyAPI.md#ContentCoreOpenpgpPublickeyRead) | **Get** /{open_p_g_p_public_key_href} | Inspect an open pgp public key



## ContentCoreOpenpgpPublickeyCreate

> AsyncOperationResponse ContentCoreOpenpgpPublickeyCreate(ctx, pulpDomain).Repository(repository).File(file).Upload(upload).FileUrl(fileUrl).Execute()

Create an open pgp public key



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
	file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the content unit. (optional)
	upload := "upload_example" // string | An uncommitted upload that may be turned into the content unit. (optional)
	fileUrl := "fileUrl_example" // string | A url that Pulp can download and turn into the content unit. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentOpenpgpPublickeyAPI.ContentCoreOpenpgpPublickeyCreate(context.Background(), pulpDomain).Repository(repository).File(file).Upload(upload).FileUrl(fileUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentOpenpgpPublickeyAPI.ContentCoreOpenpgpPublickeyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentCoreOpenpgpPublickeyCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentOpenpgpPublickeyAPI.ContentCoreOpenpgpPublickeyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentCoreOpenpgpPublickeyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **file** | ***os.File** | An uploaded file that may be turned into the content unit. | 
 **upload** | **string** | An uncommitted upload that may be turned into the content unit. | 
 **fileUrl** | **string** | A url that Pulp can download and turn into the content unit. | 

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


## ContentCoreOpenpgpPublickeyList

> PaginatedOpenPGPPublicKeyResponseList ContentCoreOpenpgpPublickeyList(ctx, pulpDomain).Fingerprint(fingerprint).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List open pgp public keys



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
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `raw_data` - Raw data* `-raw_data` - Raw data (descending)* `fingerprint` - Fingerprint* `-fingerprint` - Fingerprint (descending)* `created` - Created* `-created` - Created (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF/PRN (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentOpenpgpPublickeyAPI.ContentCoreOpenpgpPublickeyList(context.Background(), pulpDomain).Fingerprint(fingerprint).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentOpenpgpPublickeyAPI.ContentCoreOpenpgpPublickeyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentCoreOpenpgpPublickeyList`: PaginatedOpenPGPPublicKeyResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentOpenpgpPublickeyAPI.ContentCoreOpenpgpPublickeyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentCoreOpenpgpPublickeyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fingerprint** | **string** | Filter results where fingerprint matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;raw_data&#x60; - Raw data* &#x60;-raw_data&#x60; - Raw data (descending)* &#x60;fingerprint&#x60; - Fingerprint* &#x60;-fingerprint&#x60; - Fingerprint (descending)* &#x60;created&#x60; - Created* &#x60;-created&#x60; - Created (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF/PRN | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedOpenPGPPublicKeyResponseList**](PaginatedOpenPGPPublicKeyResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentCoreOpenpgpPublickeyRead

> OpenPGPPublicKeyResponse ContentCoreOpenpgpPublickeyRead(ctx, openPGPPublicKeyHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an open pgp public key



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
	openPGPPublicKeyHref := "openPGPPublicKeyHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentOpenpgpPublickeyAPI.ContentCoreOpenpgpPublickeyRead(context.Background(), openPGPPublicKeyHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentOpenpgpPublickeyAPI.ContentCoreOpenpgpPublickeyRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentCoreOpenpgpPublickeyRead`: OpenPGPPublicKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentOpenpgpPublickeyAPI.ContentCoreOpenpgpPublickeyRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPPublicKeyHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentCoreOpenpgpPublickeyReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**OpenPGPPublicKeyResponse**](OpenPGPPublicKeyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

