# \RpmCopyAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyContent**](RpmCopyAPI.md#CopyContent) | **Post** /api/pulp/{pulp_domain}/api/v3/rpm/copy/ | Copy content



## CopyContent

> AsyncOperationResponse CopyContent(ctx, pulpDomain).Copy(copy).Execute()

Copy content



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	copy := *openapiclient.NewCopy(map[string]interface{}(123)) // Copy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RpmCopyAPI.CopyContent(context.Background(), pulpDomain).Copy(copy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RpmCopyAPI.CopyContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopyContent`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RpmCopyAPI.CopyContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **copy** | [**Copy**](Copy.md) |  | 

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

