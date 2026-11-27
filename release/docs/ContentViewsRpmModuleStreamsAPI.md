# \ContentViewsRpmModuleStreamsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentViewsSearchRpmModuleStreamsList**](ContentViewsRpmModuleStreamsAPI.md#ContentViewsSearchRpmModuleStreamsList) | **Get** /{service_content_view_href}search/rpm/module-streams/ | List content view search scopes



## ContentViewsSearchRpmModuleStreamsList

> PaginatedContentViewModuleStreamResponseList ContentViewsSearchRpmModuleStreamsList(ctx, serviceContentViewHref).XTaskDiagnostics(xTaskDiagnostics).HasPackages(hasPackages).Limit(limit).Offset(offset).RpmNames(rpmNames).Search(search).SortBy(sortBy).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	hasPackages := true // bool | If true, only module streams that contain packages are returned; if false, only module streams without packages. If omitted, all module streams are returned. (optional)
	limit := int32(56) // int32 | Maximum number of results to return (default 5000, maximum 5000). (optional)
	offset := int32(56) // int32 | Number of results to skip. (optional)
	rpmNames := "rpmNames_example" // string | Comma-separated list of RPM package names; only module streams containing at least one of these packages are returned. (optional)
	search := "search_example" // string | Case-insensitive substring of the name. (optional)
	sortBy := "sortBy_example" // string | Field to sort by; prefix with '-' for descending order (default 'name'). Ties are broken deterministically, so paging is stable. (optional) (default to "name")
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentViewsRpmModuleStreamsAPI.ContentViewsSearchRpmModuleStreamsList(context.Background(), serviceContentViewHref).XTaskDiagnostics(xTaskDiagnostics).HasPackages(hasPackages).Limit(limit).Offset(offset).RpmNames(rpmNames).Search(search).SortBy(sortBy).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentViewsRpmModuleStreamsAPI.ContentViewsSearchRpmModuleStreamsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentViewsSearchRpmModuleStreamsList`: PaginatedContentViewModuleStreamResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentViewsRpmModuleStreamsAPI.ContentViewsSearchRpmModuleStreamsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceContentViewHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentViewsSearchRpmModuleStreamsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **hasPackages** | **bool** | If true, only module streams that contain packages are returned; if false, only module streams without packages. If omitted, all module streams are returned. | 
 **limit** | **int32** | Maximum number of results to return (default 5000, maximum 5000). | 
 **offset** | **int32** | Number of results to skip. | 
 **rpmNames** | **string** | Comma-separated list of RPM package names; only module streams containing at least one of these packages are returned. | 
 **search** | **string** | Case-insensitive substring of the name. | 
 **sortBy** | **string** | Field to sort by; prefix with &#39;-&#39; for descending order (default &#39;name&#39;). Ties are broken deterministically, so paging is stable. | [default to &quot;name&quot;]
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedContentViewModuleStreamResponseList**](PaginatedContentViewModuleStreamResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

