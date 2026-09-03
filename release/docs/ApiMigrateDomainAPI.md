# \ApiMigrateDomainAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPulpMigrateDomainCreate**](ApiMigrateDomainAPI.md#ApiPulpMigrateDomainCreate) | **Post** /api/pulp/migrate-domain/ | Migrate domain storage to S3



## ApiPulpMigrateDomainCreate

> AsyncOperationResponse ApiPulpMigrateDomainCreate(ctx).ApiPulpMigrateDomainCreateRequest(apiPulpMigrateDomainCreateRequest).XTaskDiagnostics(xTaskDiagnostics).Execute()

Migrate domain storage to S3



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
	apiPulpMigrateDomainCreateRequest := *openapiclient.NewApiPulpMigrateDomainCreateRequest("Name_example") // ApiPulpMigrateDomainCreateRequest | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiMigrateDomainAPI.ApiPulpMigrateDomainCreate(context.Background()).ApiPulpMigrateDomainCreateRequest(apiPulpMigrateDomainCreateRequest).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiMigrateDomainAPI.ApiPulpMigrateDomainCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPulpMigrateDomainCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiMigrateDomainAPI.ApiPulpMigrateDomainCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPulpMigrateDomainCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiPulpMigrateDomainCreateRequest** | [**ApiPulpMigrateDomainCreateRequest**](ApiPulpMigrateDomainCreateRequest.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

