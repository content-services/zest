# \ApiDebugFlushContentCacheAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPulpDebugFlushContentCachePost**](ApiDebugFlushContentCacheAPI.md#ApiPulpDebugFlushContentCachePost) | **Post** /api/pulp/debug/flush-content-cache/ | Flush the content app&#39;s distribution cache (preserves locks)



## ApiPulpDebugFlushContentCachePost

> AsyncOperationResponse ApiPulpDebugFlushContentCachePost(ctx).XTaskDiagnostics(xTaskDiagnostics).Execute()

Flush the content app's distribution cache (preserves locks)



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
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiDebugFlushContentCacheAPI.ApiPulpDebugFlushContentCachePost(context.Background()).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiDebugFlushContentCacheAPI.ApiPulpDebugFlushContentCachePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPulpDebugFlushContentCachePost`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiDebugFlushContentCacheAPI.ApiPulpDebugFlushContentCachePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPulpDebugFlushContentCachePostRequest struct via the builder pattern


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

