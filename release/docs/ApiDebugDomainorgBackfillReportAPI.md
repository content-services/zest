# \ApiDebugDomainorgBackfillReportAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPulpDebugDomainorgBackfillReportGet**](ApiDebugDomainorgBackfillReportAPI.md#ApiPulpDebugDomainorgBackfillReportGet) | **Get** /api/pulp/debug/domainorg-backfill-report/ | 
[**ApiPulpDebugDomainorgBackfillReportPost**](ApiDebugDomainorgBackfillReportAPI.md#ApiPulpDebugDomainorgBackfillReportPost) | **Post** /api/pulp/debug/domainorg-backfill-report/ | Dispatch DomainOrg backfill report



## ApiPulpDebugDomainorgBackfillReportGet

> ApiPulpDebugDomainorgBackfillReportGet(ctx).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()





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
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiDebugDomainorgBackfillReportAPI.ApiPulpDebugDomainorgBackfillReportGet(context.Background()).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiDebugDomainorgBackfillReportAPI.ApiPulpDebugDomainorgBackfillReportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPulpDebugDomainorgBackfillReportGetRequest struct via the builder pattern


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


## ApiPulpDebugDomainorgBackfillReportPost

> AsyncOperationResponse ApiPulpDebugDomainorgBackfillReportPost(ctx).XTaskDiagnostics(xTaskDiagnostics).Execute()

Dispatch DomainOrg backfill report



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
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiDebugDomainorgBackfillReportAPI.ApiPulpDebugDomainorgBackfillReportPost(context.Background()).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiDebugDomainorgBackfillReportAPI.ApiPulpDebugDomainorgBackfillReportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPulpDebugDomainorgBackfillReportPost`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiDebugDomainorgBackfillReportAPI.ApiPulpDebugDomainorgBackfillReportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPulpDebugDomainorgBackfillReportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

