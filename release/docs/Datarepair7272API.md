# \Datarepair7272API

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Datarepair7272Post**](Datarepair7272API.md#Datarepair7272Post) | **Post** /api/pulp/{pulp_domain}/api/v3/datarepair/7272/ | Repair Repository Version Data (Issue #7272)



## Datarepair7272Post

> AsyncOperationResponse Datarepair7272Post(ctx, pulpDomain).DataRepair7272(dataRepair7272).XTaskDiagnostics(xTaskDiagnostics).Execute()

Repair Repository Version Data (Issue #7272)



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
	dataRepair7272 := *openapiclient.NewDataRepair7272() // DataRepair7272 | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Datarepair7272API.Datarepair7272Post(context.Background(), pulpDomain).DataRepair7272(dataRepair7272).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Datarepair7272API.Datarepair7272Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Datarepair7272Post`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `Datarepair7272API.Datarepair7272Post`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatarepair7272PostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataRepair7272** | [**DataRepair7272**](DataRepair7272.md) |  | 
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

