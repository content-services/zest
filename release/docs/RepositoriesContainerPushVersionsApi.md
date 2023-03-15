# \RepositoriesContainerPushVersionsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesContainerContainerPushVersionsDelete**](RepositoriesContainerPushVersionsApi.md#RepositoriesContainerContainerPushVersionsDelete) | **Delete** /{container_container_push_repository_version_href} | Delete a repository version
[**RepositoriesContainerContainerPushVersionsList**](RepositoriesContainerPushVersionsApi.md#RepositoriesContainerContainerPushVersionsList) | **Get** /{container_container_push_repository_href}versions/ | List repository versions
[**RepositoriesContainerContainerPushVersionsRead**](RepositoriesContainerPushVersionsApi.md#RepositoriesContainerContainerPushVersionsRead) | **Get** /{container_container_push_repository_version_href} | Inspect a repository version
[**RepositoriesContainerContainerPushVersionsRepair**](RepositoriesContainerPushVersionsApi.md#RepositoriesContainerContainerPushVersionsRepair) | **Post** /{container_container_push_repository_version_href}repair/ | 



## RepositoriesContainerContainerPushVersionsDelete

> AsyncOperationResponse RepositoriesContainerContainerPushVersionsDelete(ctx, containerContainerPushRepositoryVersionHref).Execute()

Delete a repository version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release"
)

func main() {
    containerContainerPushRepositoryVersionHref := "containerContainerPushRepositoryVersionHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushVersionsApi.RepositoriesContainerContainerPushVersionsDelete(context.Background(), containerContainerPushRepositoryVersionHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushVersionsApi.RepositoriesContainerContainerPushVersionsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushVersionsDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushVersionsApi.RepositoriesContainerContainerPushVersionsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryVersionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushVersionsDeleteRequest struct via the builder pattern


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


## RepositoriesContainerContainerPushVersionsList

> PaginatedRepositoryVersionResponseList RepositoriesContainerContainerPushVersionsList(ctx, containerContainerPushRepositoryHref).Content(content).ContentIn(contentIn).Limit(limit).Number(number).NumberGt(numberGt).NumberGte(numberGte).NumberLt(numberLt).NumberLte(numberLte).NumberRange(numberRange).Offset(offset).Ordering(ordering).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).Fields(fields).ExcludeFields(excludeFields).Execute()

List repository versions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/content-services/zest/v3/release"
)

func main() {
    containerContainerPushRepositoryHref := "containerContainerPushRepositoryHref_example" // string | 
    content := "content_example" // string | Content Unit referenced by HREF (optional)
    contentIn := "contentIn_example" // string | Content Unit referenced by HREF (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    number := int32(56) // int32 | Filter results where number matches value (optional)
    numberGt := int32(56) // int32 | Filter results where number is greater than value (optional)
    numberGte := int32(56) // int32 | Filter results where number is greater than or equal to value (optional)
    numberLt := int32(56) // int32 | Filter results where number is less than value (optional)
    numberLte := int32(56) // int32 | Filter results where number is less than or equal to value (optional)
    numberRange := []int32{int32(123)} // []int32 | Filter results where number is between two comma separated values (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    pulpCreated := time.Now() // time.Time | Filter results where pulp_created matches value (optional)
    pulpCreatedGt := time.Now() // time.Time | Filter results where pulp_created is greater than value (optional)
    pulpCreatedGte := time.Now() // time.Time | Filter results where pulp_created is greater than or equal to value (optional)
    pulpCreatedLt := time.Now() // time.Time | Filter results where pulp_created is less than value (optional)
    pulpCreatedLte := time.Now() // time.Time | Filter results where pulp_created is less than or equal to value (optional)
    pulpCreatedRange := []time.Time{time.Now()} // []time.Time | Filter results where pulp_created is between two comma separated values (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushVersionsApi.RepositoriesContainerContainerPushVersionsList(context.Background(), containerContainerPushRepositoryHref).Content(content).ContentIn(contentIn).Limit(limit).Number(number).NumberGt(numberGt).NumberGte(numberGte).NumberLt(numberLt).NumberLte(numberLte).NumberRange(numberRange).Offset(offset).Ordering(ordering).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushVersionsApi.RepositoriesContainerContainerPushVersionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushVersionsList`: PaginatedRepositoryVersionResponseList
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushVersionsApi.RepositoriesContainerContainerPushVersionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushVersionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **content** | **string** | Content Unit referenced by HREF | 
 **contentIn** | **string** | Content Unit referenced by HREF | 
 **limit** | **int32** | Number of results to return per page. | 
 **number** | **int32** | Filter results where number matches value | 
 **numberGt** | **int32** | Filter results where number is greater than value | 
 **numberGte** | **int32** | Filter results where number is greater than or equal to value | 
 **numberLt** | **int32** | Filter results where number is less than value | 
 **numberLte** | **int32** | Filter results where number is less than or equal to value | 
 **numberRange** | **[]int32** | Filter results where number is between two comma separated values | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **pulpCreated** | **time.Time** | Filter results where pulp_created matches value | 
 **pulpCreatedGt** | **time.Time** | Filter results where pulp_created is greater than value | 
 **pulpCreatedGte** | **time.Time** | Filter results where pulp_created is greater than or equal to value | 
 **pulpCreatedLt** | **time.Time** | Filter results where pulp_created is less than value | 
 **pulpCreatedLte** | **time.Time** | Filter results where pulp_created is less than or equal to value | 
 **pulpCreatedRange** | [**[]time.Time**](time.Time.md) | Filter results where pulp_created is between two comma separated values | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedRepositoryVersionResponseList**](PaginatedRepositoryVersionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesContainerContainerPushVersionsRead

> RepositoryVersionResponse RepositoriesContainerContainerPushVersionsRead(ctx, containerContainerPushRepositoryVersionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a repository version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release"
)

func main() {
    containerContainerPushRepositoryVersionHref := "containerContainerPushRepositoryVersionHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushVersionsApi.RepositoriesContainerContainerPushVersionsRead(context.Background(), containerContainerPushRepositoryVersionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushVersionsApi.RepositoriesContainerContainerPushVersionsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushVersionsRead`: RepositoryVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushVersionsApi.RepositoriesContainerContainerPushVersionsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryVersionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushVersionsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RepositoryVersionResponse**](RepositoryVersionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesContainerContainerPushVersionsRepair

> AsyncOperationResponse RepositoriesContainerContainerPushVersionsRepair(ctx, containerContainerPushRepositoryVersionHref).Repair(repair).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release"
)

func main() {
    containerContainerPushRepositoryVersionHref := "containerContainerPushRepositoryVersionHref_example" // string | 
    repair := *openapiclient.NewRepair() // Repair | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesContainerPushVersionsApi.RepositoriesContainerContainerPushVersionsRepair(context.Background(), containerContainerPushRepositoryVersionHref).Repair(repair).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushVersionsApi.RepositoriesContainerContainerPushVersionsRepair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesContainerContainerPushVersionsRepair`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushVersionsApi.RepositoriesContainerContainerPushVersionsRepair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryVersionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushVersionsRepairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repair** | [**Repair**](Repair.md) |  | 

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

