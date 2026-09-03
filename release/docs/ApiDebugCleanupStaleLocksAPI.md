# \ApiDebugCleanupStaleLocksAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPulpDebugCleanupStaleLocksGet**](ApiDebugCleanupStaleLocksAPI.md#ApiPulpDebugCleanupStaleLocksGet) | **Get** /api/pulp/debug/cleanup-stale-locks/ | 
[**ApiPulpDebugCleanupStaleLocksPost**](ApiDebugCleanupStaleLocksAPI.md#ApiPulpDebugCleanupStaleLocksPost) | **Post** /api/pulp/debug/cleanup-stale-locks/ | Dispatch stale lock cleanup



## ApiPulpDebugCleanupStaleLocksGet

> ApiPulpDebugCleanupStaleLocksGet(ctx).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()





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
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiDebugCleanupStaleLocksAPI.ApiPulpDebugCleanupStaleLocksGet(context.Background()).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiDebugCleanupStaleLocksAPI.ApiPulpDebugCleanupStaleLocksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPulpDebugCleanupStaleLocksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPulpDebugCleanupStaleLocksPost

> AsyncOperationResponse ApiPulpDebugCleanupStaleLocksPost(ctx).XTaskDiagnostics(xTaskDiagnostics).Execute()

Dispatch stale lock cleanup



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
	resp, r, err := apiClient.ApiDebugCleanupStaleLocksAPI.ApiPulpDebugCleanupStaleLocksPost(context.Background()).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiDebugCleanupStaleLocksAPI.ApiPulpDebugCleanupStaleLocksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPulpDebugCleanupStaleLocksPost`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiDebugCleanupStaleLocksAPI.ApiPulpDebugCleanupStaleLocksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPulpDebugCleanupStaleLocksPostRequest struct via the builder pattern


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

