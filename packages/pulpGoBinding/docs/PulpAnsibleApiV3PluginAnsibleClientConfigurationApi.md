# \PulpAnsibleApiV3PluginAnsibleClientConfigurationApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationGet**](PulpAnsibleApiV3PluginAnsibleClientConfigurationApi.md#PulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationGet) | **Get** /pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/client-configuration/ | 



## PulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationGet

> ClientConfigurationResponse PulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationGet(ctx, path).Fields(fields).ExcludeFields(excludeFields).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/pulpGoBinding/packages/pulpGoBinding"
)

func main() {
    path := "path_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleApiV3PluginAnsibleClientConfigurationApi.PulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationGet(context.Background(), path).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleApiV3PluginAnsibleClientConfigurationApi.PulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationGet`: ClientConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleApiV3PluginAnsibleClientConfigurationApi.PulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyApiV3PluginAnsibleClientConfigurationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ClientConfigurationResponse**](ClientConfigurationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
