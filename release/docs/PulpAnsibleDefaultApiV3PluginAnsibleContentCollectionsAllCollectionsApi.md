# \PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsList**](PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsApi.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsList) | **Get** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/all-collections/ | 



## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsList

> []CollectionResponse PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsList(ctx, distroBasePath).Deprecated(deprecated).Name(name).Namespace(namespace).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    distroBasePath := "distroBasePath_example" // string | 
    deprecated := true // bool |  (optional)
    name := "name_example" // string |  (optional)
    namespace := "namespace_example" // string |  (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsList(context.Background(), distroBasePath).Deprecated(deprecated).Name(name).Namespace(namespace).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsList`: []CollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsListRequest struct via the builder pattern


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

