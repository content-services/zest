# \ContentViewsRpmPackagesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentViewsSearchRpmPackagesList**](ContentViewsRpmPackagesAPI.md#ContentViewsSearchRpmPackagesList) | **Get** /{service_content_view_href}search/rpm/packages/ | List content view search scopes



## ContentViewsSearchRpmPackagesList

> PaginatedContentViewPackageResponseList ContentViewsSearchRpmPackagesList(ctx, serviceContentViewHref).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Search(search).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	limit := int32(56) // int32 | Maximum number of results to return. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	search := "search_example" // string | Case-insensitive substring/prefix search term. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentViewsRpmPackagesAPI.ContentViewsSearchRpmPackagesList(context.Background(), serviceContentViewHref).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Search(search).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentViewsRpmPackagesAPI.ContentViewsSearchRpmPackagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentViewsSearchRpmPackagesList`: PaginatedContentViewPackageResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentViewsRpmPackagesAPI.ContentViewsSearchRpmPackagesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceContentViewHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentViewsSearchRpmPackagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **limit** | **int32** | Maximum number of results to return. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **search** | **string** | Case-insensitive substring/prefix search term. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedContentViewPackageResponseList**](PaginatedContentViewPackageResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

