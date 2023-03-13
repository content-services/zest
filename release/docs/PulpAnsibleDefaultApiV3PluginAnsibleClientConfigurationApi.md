# \PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationGet**](PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationApi.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationGet) | **Get** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/client-configuration/ | 



## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationGet

> ClientConfigurationResponse PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationGet(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationGet`: ClientConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationGetRequest struct via the builder pattern


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

