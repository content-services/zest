# \ContentViewsRpmErrataAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentViewsSearchRpmErrataList**](ContentViewsRpmErrataAPI.md#ContentViewsSearchRpmErrataList) | **Get** /{service_content_view_href}search/rpm/errata/ | List content view search scopes



## ContentViewsSearchRpmErrataList

> PaginatedContentViewErrataResponseList ContentViewsSearchRpmErrataList(ctx, serviceContentViewHref).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Search(search).Severity(severity).SortBy(sortBy).Type_(type_).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	search := "search_example" // string | Case-insensitive substring of the advisory id, title or summary. (optional)
	severity := "severity_example" // string | Comma-separated list of errata severities to include, matched exactly (case-sensitive; Pulp values are Critical, Important, Moderate, Low). The special value 'Unknown' also matches every severity outside that list (including empty and 'None'). (optional)
	sortBy := "sortBy_example" // string | Field to sort by; prefix with '-' for descending order (default '-issued_date'). Ties are broken deterministically, so paging is stable. (optional) (default to "-issued_date")
	type_ := "type__example" // string | Comma-separated list of errata types to include, matched exactly. The special value 'other' also matches every type other than security, bugfix, enhancement (e.g. 'newpackage'). (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentViewsRpmErrataAPI.ContentViewsSearchRpmErrataList(context.Background(), serviceContentViewHref).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Search(search).Severity(severity).SortBy(sortBy).Type_(type_).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentViewsRpmErrataAPI.ContentViewsSearchRpmErrataList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentViewsSearchRpmErrataList`: PaginatedContentViewErrataResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentViewsRpmErrataAPI.ContentViewsSearchRpmErrataList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceContentViewHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentViewsSearchRpmErrataListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **limit** | **int32** | Maximum number of results to return (default 100, maximum 1000). | 
 **offset** | **int32** | Number of results to skip. | 
 **search** | **string** | Case-insensitive substring of the advisory id, title or summary. | 
 **severity** | **string** | Comma-separated list of errata severities to include, matched exactly (case-sensitive; Pulp values are Critical, Important, Moderate, Low). The special value &#39;Unknown&#39; also matches every severity outside that list (including empty and &#39;None&#39;). | 
 **sortBy** | **string** | Field to sort by; prefix with &#39;-&#39; for descending order (default &#39;-issued_date&#39;). Ties are broken deterministically, so paging is stable. | [default to &quot;-issued_date&quot;]
 **type_** | **string** | Comma-separated list of errata types to include, matched exactly. The special value &#39;other&#39; also matches every type other than security, bugfix, enhancement (e.g. &#39;newpackage&#39;). | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedContentViewErrataResponseList**](PaginatedContentViewErrataResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

