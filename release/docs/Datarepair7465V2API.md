# \Datarepair7465V2API

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Datarepair7465V2Post**](Datarepair7465V2API.md#Datarepair7465V2Post) | **Post** /api/pulp/{pulp_domain}/api/v3/datarepair/7465-v2/ | Rebuild content_ids cache (Issue #7465)



## Datarepair7465V2Post

> AsyncOperationResponse Datarepair7465V2Post(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Execute()

Rebuild content_ids cache (Issue #7465)



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
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Datarepair7465V2API.Datarepair7465V2Post(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Datarepair7465V2API.Datarepair7465V2Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Datarepair7465V2Post`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `Datarepair7465V2API.Datarepair7465V2Post`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatarepair7465V2PostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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

