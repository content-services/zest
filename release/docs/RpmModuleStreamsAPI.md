# \RpmModuleStreamsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RpmContentSearchModuleStreamsList**](RpmModuleStreamsAPI.md#RpmContentSearchModuleStreamsList) | **Get** /api/pulp/{pulp_domain}/api/v3/rpm/content-search/module-streams/ | List content view search scopes



## RpmContentSearchModuleStreamsList

> PaginatedContentViewModuleStreamResponseList RpmContentSearchModuleStreamsList(ctx, pulpDomain).RepositoryVersion(repositoryVersion).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).RpmNames(rpmNames).Search(search).SortBy(sortBy).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	limit := int32(56) // int32 | Maximum number of results to return. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	rpmNames := "rpmNames_example" // string | Comma-separated list of RPM package names; only module streams containing at least one of these packages are returned. (optional)
	search := "search_example" // string | Case-insensitive substring/prefix search term. (optional)
	sortBy := "sortBy_example" // string | Field to sort by; prefix with '-' for descending order. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RpmModuleStreamsAPI.RpmContentSearchModuleStreamsList(context.Background(), pulpDomain).RepositoryVersion(repositoryVersion).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).RpmNames(rpmNames).Search(search).SortBy(sortBy).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RpmModuleStreamsAPI.RpmContentSearchModuleStreamsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RpmContentSearchModuleStreamsList`: PaginatedContentViewModuleStreamResponseList
	fmt.Fprintf(os.Stdout, "Response from `RpmModuleStreamsAPI.RpmContentSearchModuleStreamsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRpmContentSearchModuleStreamsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositoryVersion** | **[]string** | Href or PRN of a repository version to search. Repeat the parameter to search multiple repository versions in a single request; they may span any domain the caller has read access to (e.g. the caller&#39;s own domain plus shared domains such as redhat or community). At least one is required. A version in a domain the caller can&#39;t read is silently excluded from results rather than causing an error. | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **limit** | **int32** | Maximum number of results to return. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **rpmNames** | **string** | Comma-separated list of RPM package names; only module streams containing at least one of these packages are returned. | 
 **search** | **string** | Case-insensitive substring/prefix search term. | 
 **sortBy** | **string** | Field to sort by; prefix with &#39;-&#39; for descending order. | 
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

