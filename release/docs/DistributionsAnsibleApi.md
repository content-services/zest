# \DistributionsAnsibleApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DistributionsAnsibleAnsibleCreate**](DistributionsAnsibleApi.md#DistributionsAnsibleAnsibleCreate) | **Post** /pulp/api/v3/distributions/ansible/ansible/ | Create an ansible distribution
[**DistributionsAnsibleAnsibleDelete**](DistributionsAnsibleApi.md#DistributionsAnsibleAnsibleDelete) | **Delete** /{ansible_ansible_distribution_href} | Delete an ansible distribution
[**DistributionsAnsibleAnsibleList**](DistributionsAnsibleApi.md#DistributionsAnsibleAnsibleList) | **Get** /pulp/api/v3/distributions/ansible/ansible/ | List ansible distributions
[**DistributionsAnsibleAnsiblePartialUpdate**](DistributionsAnsibleApi.md#DistributionsAnsibleAnsiblePartialUpdate) | **Patch** /{ansible_ansible_distribution_href} | Update an ansible distribution
[**DistributionsAnsibleAnsibleRead**](DistributionsAnsibleApi.md#DistributionsAnsibleAnsibleRead) | **Get** /{ansible_ansible_distribution_href} | Inspect an ansible distribution
[**DistributionsAnsibleAnsibleUpdate**](DistributionsAnsibleApi.md#DistributionsAnsibleAnsibleUpdate) | **Put** /{ansible_ansible_distribution_href} | Update an ansible distribution



## DistributionsAnsibleAnsibleCreate

> AsyncOperationResponse DistributionsAnsibleAnsibleCreate(ctx).AnsibleAnsibleDistribution(ansibleAnsibleDistribution).Execute()

Create an ansible distribution



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
    ansibleAnsibleDistribution := *openapiclient.NewAnsibleAnsibleDistribution("BasePath_example", "Name_example") // AnsibleAnsibleDistribution | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAnsibleApi.DistributionsAnsibleAnsibleCreate(context.Background()).AnsibleAnsibleDistribution(ansibleAnsibleDistribution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAnsibleApi.DistributionsAnsibleAnsibleCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsAnsibleAnsibleCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAnsibleApi.DistributionsAnsibleAnsibleCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsAnsibleAnsibleCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ansibleAnsibleDistribution** | [**AnsibleAnsibleDistribution**](AnsibleAnsibleDistribution.md) |  | 

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


## DistributionsAnsibleAnsibleDelete

> AsyncOperationResponse DistributionsAnsibleAnsibleDelete(ctx, ansibleAnsibleDistributionHref).Execute()

Delete an ansible distribution



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
    ansibleAnsibleDistributionHref := "ansibleAnsibleDistributionHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAnsibleApi.DistributionsAnsibleAnsibleDelete(context.Background(), ansibleAnsibleDistributionHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAnsibleApi.DistributionsAnsibleAnsibleDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsAnsibleAnsibleDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAnsibleApi.DistributionsAnsibleAnsibleDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsAnsibleAnsibleDeleteRequest struct via the builder pattern


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


## DistributionsAnsibleAnsibleList

> PaginatedansibleAnsibleDistributionResponseList DistributionsAnsibleAnsibleList(ctx).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()

List ansible distributions



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
    withContent := "withContent_example" // string | Filter distributions based on the content served by them (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAnsibleApi.DistributionsAnsibleAnsibleList(context.Background()).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAnsibleApi.DistributionsAnsibleAnsibleList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsAnsibleAnsibleList`: PaginatedansibleAnsibleDistributionResponseList
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAnsibleApi.DistributionsAnsibleAnsibleList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsAnsibleAnsibleListRequest struct via the builder pattern


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
 **withContent** | **string** | Filter distributions based on the content served by them | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedansibleAnsibleDistributionResponseList**](PaginatedansibleAnsibleDistributionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsAnsibleAnsiblePartialUpdate

> AsyncOperationResponse DistributionsAnsibleAnsiblePartialUpdate(ctx, ansibleAnsibleDistributionHref).PatchedansibleAnsibleDistribution(patchedansibleAnsibleDistribution).Execute()

Update an ansible distribution



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
    ansibleAnsibleDistributionHref := "ansibleAnsibleDistributionHref_example" // string | 
    patchedansibleAnsibleDistribution := *openapiclient.NewPatchedansibleAnsibleDistribution() // PatchedansibleAnsibleDistribution | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAnsibleApi.DistributionsAnsibleAnsiblePartialUpdate(context.Background(), ansibleAnsibleDistributionHref).PatchedansibleAnsibleDistribution(patchedansibleAnsibleDistribution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAnsibleApi.DistributionsAnsibleAnsiblePartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsAnsibleAnsiblePartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAnsibleApi.DistributionsAnsibleAnsiblePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsAnsibleAnsiblePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedansibleAnsibleDistribution** | [**PatchedansibleAnsibleDistribution**](PatchedansibleAnsibleDistribution.md) |  | 

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


## DistributionsAnsibleAnsibleRead

> AnsibleAnsibleDistributionResponse DistributionsAnsibleAnsibleRead(ctx, ansibleAnsibleDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an ansible distribution



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
    ansibleAnsibleDistributionHref := "ansibleAnsibleDistributionHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAnsibleApi.DistributionsAnsibleAnsibleRead(context.Background(), ansibleAnsibleDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAnsibleApi.DistributionsAnsibleAnsibleRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsAnsibleAnsibleRead`: AnsibleAnsibleDistributionResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAnsibleApi.DistributionsAnsibleAnsibleRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsAnsibleAnsibleReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**AnsibleAnsibleDistributionResponse**](AnsibleAnsibleDistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsAnsibleAnsibleUpdate

> AsyncOperationResponse DistributionsAnsibleAnsibleUpdate(ctx, ansibleAnsibleDistributionHref).AnsibleAnsibleDistribution(ansibleAnsibleDistribution).Execute()

Update an ansible distribution



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
    ansibleAnsibleDistributionHref := "ansibleAnsibleDistributionHref_example" // string | 
    ansibleAnsibleDistribution := *openapiclient.NewAnsibleAnsibleDistribution("BasePath_example", "Name_example") // AnsibleAnsibleDistribution | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistributionsAnsibleApi.DistributionsAnsibleAnsibleUpdate(context.Background(), ansibleAnsibleDistributionHref).AnsibleAnsibleDistribution(ansibleAnsibleDistribution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsAnsibleApi.DistributionsAnsibleAnsibleUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistributionsAnsibleAnsibleUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributionsAnsibleApi.DistributionsAnsibleAnsibleUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsAnsibleAnsibleUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ansibleAnsibleDistribution** | [**AnsibleAnsibleDistribution**](AnsibleAnsibleDistribution.md) |  | 

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

