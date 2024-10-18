# \DistributionsOpenpgpAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DistributionsCoreOpenpgpCreate**](DistributionsOpenpgpAPI.md#DistributionsCoreOpenpgpCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/distributions/core/openpgp/ | Create an open pgp distribution
[**DistributionsCoreOpenpgpDelete**](DistributionsOpenpgpAPI.md#DistributionsCoreOpenpgpDelete) | **Delete** /{open_p_g_p_distribution_href} | Delete an open pgp distribution
[**DistributionsCoreOpenpgpList**](DistributionsOpenpgpAPI.md#DistributionsCoreOpenpgpList) | **Get** /api/pulp/{pulp_domain}/api/v3/distributions/core/openpgp/ | List open pgp distributions
[**DistributionsCoreOpenpgpPartialUpdate**](DistributionsOpenpgpAPI.md#DistributionsCoreOpenpgpPartialUpdate) | **Patch** /{open_p_g_p_distribution_href} | Update an open pgp distribution
[**DistributionsCoreOpenpgpRead**](DistributionsOpenpgpAPI.md#DistributionsCoreOpenpgpRead) | **Get** /{open_p_g_p_distribution_href} | Inspect an open pgp distribution
[**DistributionsCoreOpenpgpSetLabel**](DistributionsOpenpgpAPI.md#DistributionsCoreOpenpgpSetLabel) | **Post** /{open_p_g_p_distribution_href}set_label/ | Set a label
[**DistributionsCoreOpenpgpUnsetLabel**](DistributionsOpenpgpAPI.md#DistributionsCoreOpenpgpUnsetLabel) | **Post** /{open_p_g_p_distribution_href}unset_label/ | Unset a label
[**DistributionsCoreOpenpgpUpdate**](DistributionsOpenpgpAPI.md#DistributionsCoreOpenpgpUpdate) | **Put** /{open_p_g_p_distribution_href} | Update an open pgp distribution



## DistributionsCoreOpenpgpCreate

> AsyncOperationResponse DistributionsCoreOpenpgpCreate(ctx, pulpDomain).OpenPGPDistribution(openPGPDistribution).Execute()

Create an open pgp distribution



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	openPGPDistribution := *openapiclient.NewOpenPGPDistribution("BasePath_example", "Name_example") // OpenPGPDistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOpenpgpAPI.DistributionsCoreOpenpgpCreate(context.Background(), pulpDomain).OpenPGPDistribution(openPGPDistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOpenpgpAPI.DistributionsCoreOpenpgpCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsCoreOpenpgpCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOpenpgpAPI.DistributionsCoreOpenpgpCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsCoreOpenpgpCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openPGPDistribution** | [**OpenPGPDistribution**](OpenPGPDistribution.md) |  | 

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


## DistributionsCoreOpenpgpDelete

> AsyncOperationResponse DistributionsCoreOpenpgpDelete(ctx, openPGPDistributionHref).Execute()

Delete an open pgp distribution



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	openPGPDistributionHref := "openPGPDistributionHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOpenpgpAPI.DistributionsCoreOpenpgpDelete(context.Background(), openPGPDistributionHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOpenpgpAPI.DistributionsCoreOpenpgpDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsCoreOpenpgpDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOpenpgpAPI.DistributionsCoreOpenpgpDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsCoreOpenpgpDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsCoreOpenpgpList

> PaginatedOpenPGPDistributionResponseList DistributionsCoreOpenpgpList(ctx, pulpDomain).Limit(limit).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()

List open pgp distributions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `name` - Name* `-name` - Name (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `base_path` - Base path* `-base_path` - Base path (descending)* `hidden` - Hidden* `-hidden` - Hidden (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repositoryVersion := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter results where repository_version matches value (optional)
	withContent := "withContent_example" // string | Filter distributions based on the content served by them (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOpenpgpAPI.DistributionsCoreOpenpgpList(context.Background(), pulpDomain).Limit(limit).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOpenpgpAPI.DistributionsCoreOpenpgpList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsCoreOpenpgpList`: PaginatedOpenPGPDistributionResponseList
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOpenpgpAPI.DistributionsCoreOpenpgpList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsCoreOpenpgpListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;base_path&#x60; - Base path* &#x60;-base_path&#x60; - Base path (descending)* &#x60;hidden&#x60; - Hidden* &#x60;-hidden&#x60; - Hidden (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repositoryVersion** | **string** | Filter results where repository_version matches value | 
 **withContent** | **string** | Filter distributions based on the content served by them | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedOpenPGPDistributionResponseList**](PaginatedOpenPGPDistributionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsCoreOpenpgpPartialUpdate

> AsyncOperationResponse DistributionsCoreOpenpgpPartialUpdate(ctx, openPGPDistributionHref).PatchedOpenPGPDistribution(patchedOpenPGPDistribution).Execute()

Update an open pgp distribution



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	openPGPDistributionHref := "openPGPDistributionHref_example" // string | 
	patchedOpenPGPDistribution := *openapiclient.NewPatchedOpenPGPDistribution() // PatchedOpenPGPDistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOpenpgpAPI.DistributionsCoreOpenpgpPartialUpdate(context.Background(), openPGPDistributionHref).PatchedOpenPGPDistribution(patchedOpenPGPDistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOpenpgpAPI.DistributionsCoreOpenpgpPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsCoreOpenpgpPartialUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOpenpgpAPI.DistributionsCoreOpenpgpPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsCoreOpenpgpPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedOpenPGPDistribution** | [**PatchedOpenPGPDistribution**](PatchedOpenPGPDistribution.md) |  | 

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


## DistributionsCoreOpenpgpRead

> OpenPGPDistributionResponse DistributionsCoreOpenpgpRead(ctx, openPGPDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an open pgp distribution



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	openPGPDistributionHref := "openPGPDistributionHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOpenpgpAPI.DistributionsCoreOpenpgpRead(context.Background(), openPGPDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOpenpgpAPI.DistributionsCoreOpenpgpRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsCoreOpenpgpRead`: OpenPGPDistributionResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOpenpgpAPI.DistributionsCoreOpenpgpRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsCoreOpenpgpReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**OpenPGPDistributionResponse**](OpenPGPDistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsCoreOpenpgpSetLabel

> SetLabelResponse DistributionsCoreOpenpgpSetLabel(ctx, openPGPDistributionHref).SetLabel(setLabel).Execute()

Set a label



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	openPGPDistributionHref := "openPGPDistributionHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOpenpgpAPI.DistributionsCoreOpenpgpSetLabel(context.Background(), openPGPDistributionHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOpenpgpAPI.DistributionsCoreOpenpgpSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsCoreOpenpgpSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOpenpgpAPI.DistributionsCoreOpenpgpSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsCoreOpenpgpSetLabelRequest struct via the builder pattern


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


## DistributionsCoreOpenpgpUnsetLabel

> UnsetLabelResponse DistributionsCoreOpenpgpUnsetLabel(ctx, openPGPDistributionHref).UnsetLabel(unsetLabel).Execute()

Unset a label



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	openPGPDistributionHref := "openPGPDistributionHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOpenpgpAPI.DistributionsCoreOpenpgpUnsetLabel(context.Background(), openPGPDistributionHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOpenpgpAPI.DistributionsCoreOpenpgpUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsCoreOpenpgpUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOpenpgpAPI.DistributionsCoreOpenpgpUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsCoreOpenpgpUnsetLabelRequest struct via the builder pattern


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


## DistributionsCoreOpenpgpUpdate

> AsyncOperationResponse DistributionsCoreOpenpgpUpdate(ctx, openPGPDistributionHref).OpenPGPDistribution(openPGPDistribution).Execute()

Update an open pgp distribution



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	openPGPDistributionHref := "openPGPDistributionHref_example" // string | 
	openPGPDistribution := *openapiclient.NewOpenPGPDistribution("BasePath_example", "Name_example") // OpenPGPDistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOpenpgpAPI.DistributionsCoreOpenpgpUpdate(context.Background(), openPGPDistributionHref).OpenPGPDistribution(openPGPDistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOpenpgpAPI.DistributionsCoreOpenpgpUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsCoreOpenpgpUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOpenpgpAPI.DistributionsCoreOpenpgpUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsCoreOpenpgpUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openPGPDistribution** | [**OpenPGPDistribution**](OpenPGPDistribution.md) |  | 

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

