# \ApiCreateDomainAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPulpCreateDomainCreate**](ApiCreateDomainAPI.md#ApiPulpCreateDomainCreate) | **Post** /api/pulp/create-domain/ | Create domain



## ApiPulpCreateDomainCreate

> DomainResponse ApiPulpCreateDomainCreate(ctx).Domain(domain).XTaskDiagnostics(xTaskDiagnostics).Execute()

Create domain



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
	domain := *openapiclient.NewDomain("Name_example", openapiclient.StorageClassEnum("pulpcore.app.models.storage.FileSystem"), map[string]interface{}(123)) // Domain | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiCreateDomainAPI.ApiPulpCreateDomainCreate(context.Background()).Domain(domain).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiCreateDomainAPI.ApiPulpCreateDomainCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPulpCreateDomainCreate`: DomainResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiCreateDomainAPI.ApiPulpCreateDomainCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPulpCreateDomainCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | [**Domain**](Domain.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

