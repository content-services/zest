# \DistributionsAptApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DistributionsDebAptCreate**](DistributionsAptApi.md#DistributionsDebAptCreate) | **Post** /pulp/api/v3/distributions/deb/apt/ | Create an apt distribution
[**DistributionsDebAptDelete**](DistributionsAptApi.md#DistributionsDebAptDelete) | **Delete** /{deb_apt_distribution_href} | Delete an apt distribution
[**DistributionsDebAptList**](DistributionsAptApi.md#DistributionsDebAptList) | **Get** /pulp/api/v3/distributions/deb/apt/ | List apt distributions
[**DistributionsDebAptPartialUpdate**](DistributionsAptApi.md#DistributionsDebAptPartialUpdate) | **Patch** /{deb_apt_distribution_href} | Update an apt distribution
[**DistributionsDebAptRead**](DistributionsAptApi.md#DistributionsDebAptRead) | **Get** /{deb_apt_distribution_href} | Inspect an apt distribution
[**DistributionsDebAptUpdate**](DistributionsAptApi.md#DistributionsDebAptUpdate) | **Put** /{deb_apt_distribution_href} | Update an apt distribution



## DistributionsDebAptCreate

> AsyncOperationResponse DistributionsDebAptCreate(ctx).DebAptDistribution(debAptDistribution).Execute()

Create an apt distribution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release"
)

func main() {
    debAptDistribution := *openapiclient.NewDebAptDistribution("BasePath_example", "Name_example") // DebAptDistribution | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAptApi.DistributionsDebAptCreate(context.Background()).DebAptDistribution(debAptDistribution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAptApi.DistributionsDebAptCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsDebAptCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAptApi.DistributionsDebAptCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsDebAptCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debAptDistribution** | [**DebAptDistribution**](DebAptDistribution.md) |  | 

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


## DistributionsDebAptDelete

> AsyncOperationResponse DistributionsDebAptDelete(ctx, debAptDistributionHref).Execute()

Delete an apt distribution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release"
)

func main() {
    debAptDistributionHref := "debAptDistributionHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAptApi.DistributionsDebAptDelete(context.Background(), debAptDistributionHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAptApi.DistributionsDebAptDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsDebAptDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAptApi.DistributionsDebAptDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debAptDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsDebAptDeleteRequest struct via the builder pattern


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


## DistributionsDebAptList

> PaginateddebAptDistributionResponseList DistributionsDebAptList(ctx).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).Repository(repository).RepositoryIn(repositoryIn).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()

List apt distributions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release"
)

func main() {
    basePath := "basePath_example" // string | Filter results where base_path matches value (optional)
    basePathContains := "basePathContains_example" // string | Filter results where base_path contains value (optional)
    basePathIcontains := "basePathIcontains_example" // string | Filter results where base_path contains value (optional)
    basePathIn := []string{"Inner_example"} // []string | Filter results where base_path is in a comma-separated list of values (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
    repository := "repository_example" // string | Filter results where repository matches value (optional)
    repositoryIn := []string{"Inner_example"} // []string | Filter results where repository is in a comma-separated list of values (optional)
    withContent := "withContent_example" // string | Filter distributions based on the content served by them (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAptApi.DistributionsDebAptList(context.Background()).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).Repository(repository).RepositoryIn(repositoryIn).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAptApi.DistributionsDebAptList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsDebAptList`: PaginateddebAptDistributionResponseList
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAptApi.DistributionsDebAptList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsDebAptListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **basePath** | **string** | Filter results where base_path matches value | 
 **basePathContains** | **string** | Filter results where base_path contains value | 
 **basePathIcontains** | **string** | Filter results where base_path contains value | 
 **basePathIn** | **[]string** | Filter results where base_path is in a comma-separated list of values | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **repository** | **string** | Filter results where repository matches value | 
 **repositoryIn** | **[]string** | Filter results where repository is in a comma-separated list of values | 
 **withContent** | **string** | Filter distributions based on the content served by them | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebAptDistributionResponseList**](PaginateddebAptDistributionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsDebAptPartialUpdate

> AsyncOperationResponse DistributionsDebAptPartialUpdate(ctx, debAptDistributionHref).PatcheddebAptDistribution(patcheddebAptDistribution).Execute()

Update an apt distribution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release"
)

func main() {
    debAptDistributionHref := "debAptDistributionHref_example" // string | 
    patcheddebAptDistribution := *openapiclient.NewPatcheddebAptDistribution() // PatcheddebAptDistribution | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAptApi.DistributionsDebAptPartialUpdate(context.Background(), debAptDistributionHref).PatcheddebAptDistribution(patcheddebAptDistribution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAptApi.DistributionsDebAptPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsDebAptPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAptApi.DistributionsDebAptPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debAptDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsDebAptPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patcheddebAptDistribution** | [**PatcheddebAptDistribution**](PatcheddebAptDistribution.md) |  | 

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


## DistributionsDebAptRead

> DebAptDistributionResponse DistributionsDebAptRead(ctx, debAptDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an apt distribution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release"
)

func main() {
    debAptDistributionHref := "debAptDistributionHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAptApi.DistributionsDebAptRead(context.Background(), debAptDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAptApi.DistributionsDebAptRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsDebAptRead`: DebAptDistributionResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAptApi.DistributionsDebAptRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debAptDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsDebAptReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**DebAptDistributionResponse**](DebAptDistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsDebAptUpdate

> AsyncOperationResponse DistributionsDebAptUpdate(ctx, debAptDistributionHref).DebAptDistribution(debAptDistribution).Execute()

Update an apt distribution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release"
)

func main() {
    debAptDistributionHref := "debAptDistributionHref_example" // string | 
    debAptDistribution := *openapiclient.NewDebAptDistribution("BasePath_example", "Name_example") // DebAptDistribution | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAptApi.DistributionsDebAptUpdate(context.Background(), debAptDistributionHref).DebAptDistribution(debAptDistribution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAptApi.DistributionsDebAptUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsDebAptUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAptApi.DistributionsDebAptUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debAptDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsDebAptUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **debAptDistribution** | [**DebAptDistribution**](DebAptDistribution.md) |  | 

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

