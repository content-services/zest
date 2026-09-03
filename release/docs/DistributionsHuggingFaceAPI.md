# \DistributionsHuggingFaceAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DistributionsHuggingFaceHuggingFaceCreate**](DistributionsHuggingFaceAPI.md#DistributionsHuggingFaceHuggingFaceCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/distributions/hugging_face/hugging-face/ | Create a hugging face distribution
[**DistributionsHuggingFaceHuggingFaceDelete**](DistributionsHuggingFaceAPI.md#DistributionsHuggingFaceHuggingFaceDelete) | **Delete** /{hugging_face_hugging_face_distribution_href} | Delete a hugging face distribution
[**DistributionsHuggingFaceHuggingFaceList**](DistributionsHuggingFaceAPI.md#DistributionsHuggingFaceHuggingFaceList) | **Get** /api/pulp/{pulp_domain}/api/v3/distributions/hugging_face/hugging-face/ | List hugging face distributions
[**DistributionsHuggingFaceHuggingFacePartialUpdate**](DistributionsHuggingFaceAPI.md#DistributionsHuggingFaceHuggingFacePartialUpdate) | **Patch** /{hugging_face_hugging_face_distribution_href} | Update a hugging face distribution
[**DistributionsHuggingFaceHuggingFaceRead**](DistributionsHuggingFaceAPI.md#DistributionsHuggingFaceHuggingFaceRead) | **Get** /{hugging_face_hugging_face_distribution_href} | Inspect a hugging face distribution
[**DistributionsHuggingFaceHuggingFaceSetLabel**](DistributionsHuggingFaceAPI.md#DistributionsHuggingFaceHuggingFaceSetLabel) | **Post** /{hugging_face_hugging_face_distribution_href}set_label/ | Set a label
[**DistributionsHuggingFaceHuggingFaceUnsetLabel**](DistributionsHuggingFaceAPI.md#DistributionsHuggingFaceHuggingFaceUnsetLabel) | **Post** /{hugging_face_hugging_face_distribution_href}unset_label/ | Unset a label
[**DistributionsHuggingFaceHuggingFaceUpdate**](DistributionsHuggingFaceAPI.md#DistributionsHuggingFaceHuggingFaceUpdate) | **Put** /{hugging_face_hugging_face_distribution_href} | Update a hugging face distribution



## DistributionsHuggingFaceHuggingFaceCreate

> AsyncOperationResponse DistributionsHuggingFaceHuggingFaceCreate(ctx, pulpDomain).HuggingFaceHuggingFaceDistribution(huggingFaceHuggingFaceDistribution).XTaskDiagnostics(xTaskDiagnostics).Execute()

Create a hugging face distribution



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
	huggingFaceHuggingFaceDistribution := *openapiclient.NewHuggingFaceHuggingFaceDistribution("BasePath_example", "Name_example") // HuggingFaceHuggingFaceDistribution | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceCreate(context.Background(), pulpDomain).HuggingFaceHuggingFaceDistribution(huggingFaceHuggingFaceDistribution).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsHuggingFaceHuggingFaceCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsHuggingFaceHuggingFaceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **huggingFaceHuggingFaceDistribution** | [**HuggingFaceHuggingFaceDistribution**](HuggingFaceHuggingFaceDistribution.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## DistributionsHuggingFaceHuggingFaceDelete

> AsyncOperationResponse DistributionsHuggingFaceHuggingFaceDelete(ctx, huggingFaceHuggingFaceDistributionHref).XTaskDiagnostics(xTaskDiagnostics).Execute()

Delete a hugging face distribution



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
	huggingFaceHuggingFaceDistributionHref := "huggingFaceHuggingFaceDistributionHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceDelete(context.Background(), huggingFaceHuggingFaceDistributionHref).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsHuggingFaceHuggingFaceDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsHuggingFaceHuggingFaceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## DistributionsHuggingFaceHuggingFaceList

> PaginatedhuggingFaceHuggingFaceDistributionResponseList DistributionsHuggingFaceHuggingFaceList(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Checkpoint(checkpoint).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Repository(repository).RepositoryIn(repositoryIn).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()

List hugging face distributions



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
	basePath := "basePath_example" // string | Filter results where base_path matches value (optional)
	basePathContains := "basePathContains_example" // string | Filter results where base_path contains value (optional)
	basePathIcontains := "basePathIcontains_example" // string | Filter results where base_path contains value (optional)
	basePathIn := []string{"Inner_example"} // []string | Filter results where base_path is in a comma-separated list of values (optional)
	checkpoint := true // bool | Filter results where checkpoint matches value (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
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
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `name` - Name* `-name` - Name (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `base_path` - Base path* `-base_path` - Base path (descending)* `hidden` - Hidden* `-hidden` - Hidden (descending)* `checkpoint` - Checkpoint* `-checkpoint` - Checkpoint (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repository := "repository_example" // string |  (optional)
	repositoryIn := []string{"Inner_example"} // []string |  (optional)
	withContent := "withContent_example" // string | Filter distributions based on the content served by them (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceList(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Checkpoint(checkpoint).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Repository(repository).RepositoryIn(repositoryIn).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsHuggingFaceHuggingFaceList`: PaginatedhuggingFaceHuggingFaceDistributionResponseList
	fmt.Fprintf(os.Stdout, "Response from `DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsHuggingFaceHuggingFaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **basePath** | **string** | Filter results where base_path matches value | 
 **basePathContains** | **string** | Filter results where base_path contains value | 
 **basePathIcontains** | **string** | Filter results where base_path contains value | 
 **basePathIn** | **[]string** | Filter results where base_path is in a comma-separated list of values | 
 **checkpoint** | **bool** | Filter results where checkpoint matches value | 
 **limit** | **int32** | Number of results to return per page. | 
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
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;base_path&#x60; - Base path* &#x60;-base_path&#x60; - Base path (descending)* &#x60;hidden&#x60; - Hidden* &#x60;-hidden&#x60; - Hidden (descending)* &#x60;checkpoint&#x60; - Checkpoint* &#x60;-checkpoint&#x60; - Checkpoint (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repository** | **string** |  | 
 **repositoryIn** | **[]string** |  | 
 **withContent** | **string** | Filter distributions based on the content served by them | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedhuggingFaceHuggingFaceDistributionResponseList**](PaginatedhuggingFaceHuggingFaceDistributionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsHuggingFaceHuggingFacePartialUpdate

> HuggingFaceHuggingFaceDistributionResponse DistributionsHuggingFaceHuggingFacePartialUpdate(ctx, huggingFaceHuggingFaceDistributionHref).PatchedhuggingFaceHuggingFaceDistribution(patchedhuggingFaceHuggingFaceDistribution).XTaskDiagnostics(xTaskDiagnostics).Execute()

Update a hugging face distribution



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
	huggingFaceHuggingFaceDistributionHref := "huggingFaceHuggingFaceDistributionHref_example" // string | 
	patchedhuggingFaceHuggingFaceDistribution := *openapiclient.NewPatchedhuggingFaceHuggingFaceDistribution() // PatchedhuggingFaceHuggingFaceDistribution | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFacePartialUpdate(context.Background(), huggingFaceHuggingFaceDistributionHref).PatchedhuggingFaceHuggingFaceDistribution(patchedhuggingFaceHuggingFaceDistribution).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFacePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsHuggingFaceHuggingFacePartialUpdate`: HuggingFaceHuggingFaceDistributionResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFacePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsHuggingFaceHuggingFacePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedhuggingFaceHuggingFaceDistribution** | [**PatchedhuggingFaceHuggingFaceDistribution**](PatchedhuggingFaceHuggingFaceDistribution.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**HuggingFaceHuggingFaceDistributionResponse**](HuggingFaceHuggingFaceDistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsHuggingFaceHuggingFaceRead

> HuggingFaceHuggingFaceDistributionResponse DistributionsHuggingFaceHuggingFaceRead(ctx, huggingFaceHuggingFaceDistributionHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a hugging face distribution



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
	huggingFaceHuggingFaceDistributionHref := "huggingFaceHuggingFaceDistributionHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceRead(context.Background(), huggingFaceHuggingFaceDistributionHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsHuggingFaceHuggingFaceRead`: HuggingFaceHuggingFaceDistributionResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsHuggingFaceHuggingFaceReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**HuggingFaceHuggingFaceDistributionResponse**](HuggingFaceHuggingFaceDistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsHuggingFaceHuggingFaceSetLabel

> SetLabelResponse DistributionsHuggingFaceHuggingFaceSetLabel(ctx, huggingFaceHuggingFaceDistributionHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	huggingFaceHuggingFaceDistributionHref := "huggingFaceHuggingFaceDistributionHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceSetLabel(context.Background(), huggingFaceHuggingFaceDistributionHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsHuggingFaceHuggingFaceSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsHuggingFaceHuggingFaceSetLabelRequest struct via the builder pattern


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


## DistributionsHuggingFaceHuggingFaceUnsetLabel

> UnsetLabelResponse DistributionsHuggingFaceHuggingFaceUnsetLabel(ctx, huggingFaceHuggingFaceDistributionHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	huggingFaceHuggingFaceDistributionHref := "huggingFaceHuggingFaceDistributionHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceUnsetLabel(context.Background(), huggingFaceHuggingFaceDistributionHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsHuggingFaceHuggingFaceUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsHuggingFaceHuggingFaceUnsetLabelRequest struct via the builder pattern


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


## DistributionsHuggingFaceHuggingFaceUpdate

> HuggingFaceHuggingFaceDistributionResponse DistributionsHuggingFaceHuggingFaceUpdate(ctx, huggingFaceHuggingFaceDistributionHref).HuggingFaceHuggingFaceDistribution(huggingFaceHuggingFaceDistribution).XTaskDiagnostics(xTaskDiagnostics).Execute()

Update a hugging face distribution



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
	huggingFaceHuggingFaceDistributionHref := "huggingFaceHuggingFaceDistributionHref_example" // string | 
	huggingFaceHuggingFaceDistribution := *openapiclient.NewHuggingFaceHuggingFaceDistribution("BasePath_example", "Name_example") // HuggingFaceHuggingFaceDistribution | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceUpdate(context.Background(), huggingFaceHuggingFaceDistributionHref).HuggingFaceHuggingFaceDistribution(huggingFaceHuggingFaceDistribution).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsHuggingFaceHuggingFaceUpdate`: HuggingFaceHuggingFaceDistributionResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsHuggingFaceAPI.DistributionsHuggingFaceHuggingFaceUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsHuggingFaceHuggingFaceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **huggingFaceHuggingFaceDistribution** | [**HuggingFaceHuggingFaceDistribution**](HuggingFaceHuggingFaceDistribution.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**HuggingFaceHuggingFaceDistributionResponse**](HuggingFaceHuggingFaceDistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

