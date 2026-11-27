# \ContentViewsRpmEnvironmentsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentViewsSearchRpmEnvironmentsList**](ContentViewsRpmEnvironmentsAPI.md#ContentViewsSearchRpmEnvironmentsList) | **Get** /{service_content_view_href}search/rpm/environments/ | List content view search scopes



## ContentViewsSearchRpmEnvironmentsList

> PaginatedContentViewPackageEnvironmentResponseList ContentViewsSearchRpmEnvironmentsList(ctx, serviceContentViewHref).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Search(search).Fields(fields).ExcludeFields(excludeFields).Execute()

List content view search scopes



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
	serviceContentViewHref := "serviceContentViewHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	limit := int32(56) // int32 | Maximum number of results to return (default 100, maximum 1000). (optional)
	offset := int32(56) // int32 | Number of results to skip. (optional)
	search := "search_example" // string | Case-insensitive substring of the name. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentViewsRpmEnvironmentsAPI.ContentViewsSearchRpmEnvironmentsList(context.Background(), serviceContentViewHref).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Search(search).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentViewsRpmEnvironmentsAPI.ContentViewsSearchRpmEnvironmentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentViewsSearchRpmEnvironmentsList`: PaginatedContentViewPackageEnvironmentResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentViewsRpmEnvironmentsAPI.ContentViewsSearchRpmEnvironmentsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceContentViewHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentViewsSearchRpmEnvironmentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **limit** | **int32** | Maximum number of results to return (default 100, maximum 1000). | 
 **offset** | **int32** | Number of results to skip. | 
 **search** | **string** | Case-insensitive substring of the name. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedContentViewPackageEnvironmentResponseList**](PaginatedContentViewPackageEnvironmentResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

