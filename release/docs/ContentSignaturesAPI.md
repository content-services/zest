# \ContentSignaturesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentContainerSignaturesList**](ContentSignaturesAPI.md#ContentContainerSignaturesList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/container/signatures/ | List manifest signatures
[**ContentContainerSignaturesRead**](ContentSignaturesAPI.md#ContentContainerSignaturesRead) | **Get** /{container_manifest_signature_href} | Inspect a manifest signature
[**ContentContainerSignaturesSetLabel**](ContentSignaturesAPI.md#ContentContainerSignaturesSetLabel) | **Post** /{container_manifest_signature_href}set_label/ | Set a label
[**ContentContainerSignaturesUnsetLabel**](ContentSignaturesAPI.md#ContentContainerSignaturesUnsetLabel) | **Post** /{container_manifest_signature_href}unset_label/ | Unset a label



## ContentContainerSignaturesList

> PaginatedcontainerManifestSignatureResponseList ContentContainerSignaturesList(ctx, pulpDomain).Digest(digest).DigestIn(digestIn).KeyId(keyId).KeyIdIn(keyIdIn).Limit(limit).Manifest(manifest).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List manifest signatures



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
	digest := "digest_example" // string | Filter results where digest matches value (optional)
	digestIn := []string{"Inner_example"} // []string | Filter results where digest is in a comma-separated list of values (optional)
	keyId := "keyId_example" // string | Filter results where key_id matches value (optional)
	keyIdIn := []string{"Inner_example"} // []string | Filter results where key_id is in a comma-separated list of values (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	manifest := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	name := "name_example" // string | Filter results where name matches value (optional)
	nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
	nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
	nameIexact := "nameIexact_example" // string | Filter results where name matches value (optional)
	nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
	nameIregex := "nameIregex_example" // string | Filter results where name matches regex value (optional)
	nameIstartswith := "nameIstartswith_example" // string | Filter results where name starts with value (optional)
	nameRegex := "nameRegex_example" // string | Filter results where name matches regex value (optional)
	nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `name` - Name* `-name` - Name (descending)* `digest` - Digest* `-digest` - Digest (descending)* `type` - Type* `-type` - Type (descending)* `key_id` - Key id* `-key_id` - Key id (descending)* `timestamp` - Timestamp* `-timestamp` - Timestamp (descending)* `creator` - Creator* `-creator` - Creator (descending)* `data` - Data* `-data` - Data (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
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
	resp, r, err := apiClient.ContentSignaturesAPI.ContentContainerSignaturesList(context.Background(), pulpDomain).Digest(digest).DigestIn(digestIn).KeyId(keyId).KeyIdIn(keyIdIn).Limit(limit).Manifest(manifest).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSignaturesAPI.ContentContainerSignaturesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentContainerSignaturesList`: PaginatedcontainerManifestSignatureResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentSignaturesAPI.ContentContainerSignaturesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentContainerSignaturesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **digest** | **string** | Filter results where digest matches value | 
 **digestIn** | **[]string** | Filter results where digest is in a comma-separated list of values | 
 **keyId** | **string** | Filter results where key_id matches value | 
 **keyIdIn** | **[]string** | Filter results where key_id is in a comma-separated list of values | 
 **limit** | **int32** | Number of results to return per page. | 
 **manifest** | **[]string** | Multiple values may be separated by commas. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIexact** | **string** | Filter results where name matches value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameIregex** | **string** | Filter results where name matches regex value | 
 **nameIstartswith** | **string** | Filter results where name starts with value | 
 **nameRegex** | **string** | Filter results where name matches regex value | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;digest&#x60; - Digest* &#x60;-digest&#x60; - Digest (descending)* &#x60;type&#x60; - Type* &#x60;-type&#x60; - Type (descending)* &#x60;key_id&#x60; - Key id* &#x60;-key_id&#x60; - Key id (descending)* &#x60;timestamp&#x60; - Timestamp* &#x60;-timestamp&#x60; - Timestamp (descending)* &#x60;creator&#x60; - Creator* &#x60;-creator&#x60; - Creator (descending)* &#x60;data&#x60; - Data* &#x60;-data&#x60; - Data (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
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

[**PaginatedcontainerManifestSignatureResponseList**](PaginatedcontainerManifestSignatureResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentContainerSignaturesRead

> ContainerManifestSignatureResponse ContentContainerSignaturesRead(ctx, containerManifestSignatureHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a manifest signature



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
	containerManifestSignatureHref := "containerManifestSignatureHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentSignaturesAPI.ContentContainerSignaturesRead(context.Background(), containerManifestSignatureHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSignaturesAPI.ContentContainerSignaturesRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentContainerSignaturesRead`: ContainerManifestSignatureResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentSignaturesAPI.ContentContainerSignaturesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerManifestSignatureHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentContainerSignaturesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ContainerManifestSignatureResponse**](ContainerManifestSignatureResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentContainerSignaturesSetLabel

> SetLabelResponse ContentContainerSignaturesSetLabel(ctx, containerManifestSignatureHref).SetLabel(setLabel).Execute()

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
	containerManifestSignatureHref := "containerManifestSignatureHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentSignaturesAPI.ContentContainerSignaturesSetLabel(context.Background(), containerManifestSignatureHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSignaturesAPI.ContentContainerSignaturesSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentContainerSignaturesSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentSignaturesAPI.ContentContainerSignaturesSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerManifestSignatureHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentContainerSignaturesSetLabelRequest struct via the builder pattern


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


## ContentContainerSignaturesUnsetLabel

> UnsetLabelResponse ContentContainerSignaturesUnsetLabel(ctx, containerManifestSignatureHref).UnsetLabel(unsetLabel).Execute()

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
	containerManifestSignatureHref := "containerManifestSignatureHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentSignaturesAPI.ContentContainerSignaturesUnsetLabel(context.Background(), containerManifestSignatureHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSignaturesAPI.ContentContainerSignaturesUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentContainerSignaturesUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentSignaturesAPI.ContentContainerSignaturesUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerManifestSignatureHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentContainerSignaturesUnsetLabelRequest struct via the builder pattern


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

