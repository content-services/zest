# \PulpAnsibleApiV3CollectionsAllApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyApiV3CollectionsAllList**](PulpAnsibleApiV3CollectionsAllApi.md#PulpAnsibleGalaxyApiV3CollectionsAllList) | **Get** /pulp_ansible/galaxy/{path}/api/v3/collections/all/ | 



## PulpAnsibleGalaxyApiV3CollectionsAllList

> []CollectionResponse PulpAnsibleGalaxyApiV3CollectionsAllList(ctx, path).Deprecated(deprecated).Name(name).Namespace(namespace).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    path := "path_example" // string | 
    deprecated := true // bool |  (optional)
    name := "name_example" // string |  (optional)
    namespace := "namespace_example" // string |  (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleApiV3CollectionsAllApi.PulpAnsibleGalaxyApiV3CollectionsAllList(context.Background(), path).Deprecated(deprecated).Name(name).Namespace(namespace).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleApiV3CollectionsAllApi.PulpAnsibleGalaxyApiV3CollectionsAllList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyApiV3CollectionsAllList`: []CollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleApiV3CollectionsAllApi.PulpAnsibleGalaxyApiV3CollectionsAllList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyApiV3CollectionsAllListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deprecated** | **bool** |  | 
 **name** | **string** |  | 
 **namespace** | **string** |  | 
 **ordering** | **[]string** | Ordering | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**[]CollectionResponse**](CollectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

