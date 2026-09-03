# \PypiUnyankAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PypiUnyankUnyank**](PypiUnyankAPI.md#PypiUnyankUnyank) | **Post** /pypi/{pulp_domain}/{path}/unyank/ | Unyank a package version



## PypiUnyankUnyank

> AsyncOperationResponse PypiUnyankUnyank(ctx, path, pulpDomain).Yank(yank).XTaskDiagnostics(xTaskDiagnostics).Execute()

Unyank a package version



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
	path := "path_example" // string | 
	pulpDomain := "pulpDomain_example" // string | 
	yank := *openapiclient.NewYank("Name_example", "Version_example") // Yank | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PypiUnyankAPI.PypiUnyankUnyank(context.Background(), path, pulpDomain).Yank(yank).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PypiUnyankAPI.PypiUnyankUnyank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PypiUnyankUnyank`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `PypiUnyankAPI.PypiUnyankUnyank`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPypiUnyankUnyankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **yank** | [**Yank**](Yank.md) |  | 
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

