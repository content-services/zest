# \RpmErrataAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RpmContentSearchErrataList**](RpmErrataAPI.md#RpmContentSearchErrataList) | **Get** /api/pulp/{pulp_domain}/api/v3/rpm/content-search/errata/ | List content view search scopes



## RpmContentSearchErrataList

> PaginatedContentViewErrataResponseList RpmContentSearchErrataList(ctx, pulpDomain).RepositoryVersion(repositoryVersion).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Search(search).Severity(severity).SortBy(sortBy).Type_(type_).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	pulpDomain := "pulpDomain_example" // string | 
	repositoryVersion := []string{"Inner_example"} // []string | Href or PRN of a repository version to search. Repeat the parameter to search multiple repository versions in a single request; they may span any domain the caller has read access to (e.g. the caller's own domain plus shared domains such as redhat or community). At least one is required. A version in a domain the caller can't read is silently excluded from results rather than causing an error.
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
	resp, r, err := apiClient.RpmErrataAPI.RpmContentSearchErrataList(context.Background(), pulpDomain).RepositoryVersion(repositoryVersion).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Search(search).Severity(severity).SortBy(sortBy).Type_(type_).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RpmErrataAPI.RpmContentSearchErrataList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RpmContentSearchErrataList`: PaginatedContentViewErrataResponseList
	fmt.Fprintf(os.Stdout, "Response from `RpmErrataAPI.RpmContentSearchErrataList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRpmContentSearchErrataListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositoryVersion** | **[]string** | Href or PRN of a repository version to search. Repeat the parameter to search multiple repository versions in a single request; they may span any domain the caller has read access to (e.g. the caller&#39;s own domain plus shared domains such as redhat or community). At least one is required. A version in a domain the caller can&#39;t read is silently excluded from results rather than causing an error. | 
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

