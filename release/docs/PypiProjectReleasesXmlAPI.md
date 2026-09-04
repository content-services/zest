# \PypiProjectReleasesXmlAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PypiRssProjectReleasesXmlProjectReleases**](PypiProjectReleasesXmlAPI.md#PypiRssProjectReleasesXmlProjectReleases) | **Get** /pypi/{pulp_domain}/{path}/rss/project/{package}/releases.xml | Get project releases RSS feed



## PypiRssProjectReleasesXmlProjectReleases

> PypiRssProjectReleasesXmlProjectReleases(ctx, package_, path, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Get project releases RSS feed



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
	package_ := "package__example" // string | 
	path := "path_example" // string | 
	pulpDomain := "pulpDomain_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PypiProjectReleasesXmlAPI.PypiRssProjectReleasesXmlProjectReleases(context.Background(), package_, path, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PypiProjectReleasesXmlAPI.PypiRssProjectReleasesXmlProjectReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**package_** | **string** |  | 
**path** | **string** |  | 
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPypiRssProjectReleasesXmlProjectReleasesRequest struct via the builder pattern


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

