# \DistributionsMavenAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DistributionsMavenMavenCreate**](DistributionsMavenAPI.md#DistributionsMavenMavenCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/distributions/maven/maven/ | Create a maven distribution
[**DistributionsMavenMavenDelete**](DistributionsMavenAPI.md#DistributionsMavenMavenDelete) | **Delete** /{maven_maven_distribution_href} | Delete a maven distribution
[**DistributionsMavenMavenList**](DistributionsMavenAPI.md#DistributionsMavenMavenList) | **Get** /api/pulp/{pulp_domain}/api/v3/distributions/maven/maven/ | List maven distributions
[**DistributionsMavenMavenPartialUpdate**](DistributionsMavenAPI.md#DistributionsMavenMavenPartialUpdate) | **Patch** /{maven_maven_distribution_href} | Update a maven distribution
[**DistributionsMavenMavenRead**](DistributionsMavenAPI.md#DistributionsMavenMavenRead) | **Get** /{maven_maven_distribution_href} | Inspect a maven distribution
[**DistributionsMavenMavenSetLabel**](DistributionsMavenAPI.md#DistributionsMavenMavenSetLabel) | **Post** /{maven_maven_distribution_href}set_label/ | Set a label
[**DistributionsMavenMavenUnsetLabel**](DistributionsMavenAPI.md#DistributionsMavenMavenUnsetLabel) | **Post** /{maven_maven_distribution_href}unset_label/ | Unset a label
[**DistributionsMavenMavenUpdate**](DistributionsMavenAPI.md#DistributionsMavenMavenUpdate) | **Put** /{maven_maven_distribution_href} | Update a maven distribution



## DistributionsMavenMavenCreate

> AsyncOperationResponse DistributionsMavenMavenCreate(ctx, pulpDomain).MavenMavenDistribution(mavenMavenDistribution).Execute()

Create a maven distribution



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
	mavenMavenDistribution := *openapiclient.NewMavenMavenDistribution("BasePath_example", "Name_example") // MavenMavenDistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsMavenAPI.DistributionsMavenMavenCreate(context.Background(), pulpDomain).MavenMavenDistribution(mavenMavenDistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsMavenAPI.DistributionsMavenMavenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsMavenMavenCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsMavenAPI.DistributionsMavenMavenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsMavenMavenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mavenMavenDistribution** | [**MavenMavenDistribution**](MavenMavenDistribution.md) |  | 

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


## DistributionsMavenMavenDelete

> AsyncOperationResponse DistributionsMavenMavenDelete(ctx, mavenMavenDistributionHref).Execute()

Delete a maven distribution



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
	mavenMavenDistributionHref := "mavenMavenDistributionHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsMavenAPI.DistributionsMavenMavenDelete(context.Background(), mavenMavenDistributionHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsMavenAPI.DistributionsMavenMavenDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsMavenMavenDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsMavenAPI.DistributionsMavenMavenDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsMavenMavenDeleteRequest struct via the builder pattern


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


## DistributionsMavenMavenList

> PaginatedmavenMavenDistributionResponseList DistributionsMavenMavenList(ctx, pulpDomain).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Checkpoint(checkpoint).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Repository(repository).RepositoryIn(repositoryIn).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()

List maven distributions



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
	repository := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter results where repository matches value (optional)
	repositoryIn := []string{"Inner_example"} // []string | Filter results where repository is in a comma-separated list of values (optional)
	withContent := "withContent_example" // string | Filter distributions based on the content served by them (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsMavenAPI.DistributionsMavenMavenList(context.Background(), pulpDomain).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Checkpoint(checkpoint).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Repository(repository).RepositoryIn(repositoryIn).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsMavenAPI.DistributionsMavenMavenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsMavenMavenList`: PaginatedmavenMavenDistributionResponseList
	fmt.Fprintf(os.Stdout, "Response from `DistributionsMavenAPI.DistributionsMavenMavenList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsMavenMavenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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
 **repository** | **string** | Filter results where repository matches value | 
 **repositoryIn** | **[]string** | Filter results where repository is in a comma-separated list of values | 
 **withContent** | **string** | Filter distributions based on the content served by them | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedmavenMavenDistributionResponseList**](PaginatedmavenMavenDistributionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsMavenMavenPartialUpdate

> AsyncOperationResponse DistributionsMavenMavenPartialUpdate(ctx, mavenMavenDistributionHref).PatchedmavenMavenDistribution(patchedmavenMavenDistribution).Execute()

Update a maven distribution



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
	mavenMavenDistributionHref := "mavenMavenDistributionHref_example" // string | 
	patchedmavenMavenDistribution := *openapiclient.NewPatchedmavenMavenDistribution() // PatchedmavenMavenDistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsMavenAPI.DistributionsMavenMavenPartialUpdate(context.Background(), mavenMavenDistributionHref).PatchedmavenMavenDistribution(patchedmavenMavenDistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsMavenAPI.DistributionsMavenMavenPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsMavenMavenPartialUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsMavenAPI.DistributionsMavenMavenPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsMavenMavenPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedmavenMavenDistribution** | [**PatchedmavenMavenDistribution**](PatchedmavenMavenDistribution.md) |  | 

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


## DistributionsMavenMavenRead

> MavenMavenDistributionResponse DistributionsMavenMavenRead(ctx, mavenMavenDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a maven distribution



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
	mavenMavenDistributionHref := "mavenMavenDistributionHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsMavenAPI.DistributionsMavenMavenRead(context.Background(), mavenMavenDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsMavenAPI.DistributionsMavenMavenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsMavenMavenRead`: MavenMavenDistributionResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsMavenAPI.DistributionsMavenMavenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsMavenMavenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**MavenMavenDistributionResponse**](MavenMavenDistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsMavenMavenSetLabel

> SetLabelResponse DistributionsMavenMavenSetLabel(ctx, mavenMavenDistributionHref).SetLabel(setLabel).Execute()

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
	mavenMavenDistributionHref := "mavenMavenDistributionHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsMavenAPI.DistributionsMavenMavenSetLabel(context.Background(), mavenMavenDistributionHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsMavenAPI.DistributionsMavenMavenSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsMavenMavenSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsMavenAPI.DistributionsMavenMavenSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsMavenMavenSetLabelRequest struct via the builder pattern


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


## DistributionsMavenMavenUnsetLabel

> UnsetLabelResponse DistributionsMavenMavenUnsetLabel(ctx, mavenMavenDistributionHref).UnsetLabel(unsetLabel).Execute()

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
	mavenMavenDistributionHref := "mavenMavenDistributionHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsMavenAPI.DistributionsMavenMavenUnsetLabel(context.Background(), mavenMavenDistributionHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsMavenAPI.DistributionsMavenMavenUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsMavenMavenUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsMavenAPI.DistributionsMavenMavenUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsMavenMavenUnsetLabelRequest struct via the builder pattern


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


## DistributionsMavenMavenUpdate

> AsyncOperationResponse DistributionsMavenMavenUpdate(ctx, mavenMavenDistributionHref).MavenMavenDistribution(mavenMavenDistribution).Execute()

Update a maven distribution



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
	mavenMavenDistributionHref := "mavenMavenDistributionHref_example" // string | 
	mavenMavenDistribution := *openapiclient.NewMavenMavenDistribution("BasePath_example", "Name_example") // MavenMavenDistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsMavenAPI.DistributionsMavenMavenUpdate(context.Background(), mavenMavenDistributionHref).MavenMavenDistribution(mavenMavenDistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsMavenAPI.DistributionsMavenMavenUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsMavenMavenUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsMavenAPI.DistributionsMavenMavenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsMavenMavenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mavenMavenDistribution** | [**MavenMavenDistribution**](MavenMavenDistribution.md) |  | 

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

