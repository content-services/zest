# \VulnReportAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VulnReportCreate**](VulnReportAPI.md#VulnReportCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/vuln_report/ | Create a vulnerability report
[**VulnReportDelete**](VulnReportAPI.md#VulnReportDelete) | **Delete** /{service_vulnerability_report_href} | Delete a vulnerability report
[**VulnReportList**](VulnReportAPI.md#VulnReportList) | **Get** /api/pulp/{pulp_domain}/api/v3/vuln_report/ | List vulnerability reports
[**VulnReportRead**](VulnReportAPI.md#VulnReportRead) | **Get** /{service_vulnerability_report_href} | Inspect a vulnerability report



## VulnReportCreate

> ServiceVulnerabilityReportResponse VulnReportCreate(ctx, pulpDomain).ServiceVulnerabilityReport(serviceVulnerabilityReport).Execute()

Create a vulnerability report



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	serviceVulnerabilityReport := *openapiclient.NewServiceVulnerabilityReport(interface{}(123)) // ServiceVulnerabilityReport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnReportAPI.VulnReportCreate(context.Background(), pulpDomain).ServiceVulnerabilityReport(serviceVulnerabilityReport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnReportAPI.VulnReportCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VulnReportCreate`: ServiceVulnerabilityReportResponse
	fmt.Fprintf(os.Stdout, "Response from `VulnReportAPI.VulnReportCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnReportCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceVulnerabilityReport** | [**ServiceVulnerabilityReport**](ServiceVulnerabilityReport.md) |  | 

### Return type

[**ServiceVulnerabilityReportResponse**](ServiceVulnerabilityReportResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VulnReportDelete

> VulnReportDelete(ctx, serviceVulnerabilityReportHref).Execute()

Delete a vulnerability report



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	serviceVulnerabilityReportHref := "serviceVulnerabilityReportHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VulnReportAPI.VulnReportDelete(context.Background(), serviceVulnerabilityReportHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnReportAPI.VulnReportDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceVulnerabilityReportHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnReportDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## VulnReportList

> PaginatedserviceVulnerabilityReportResponseList VulnReportList(ctx, pulpDomain).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List vulnerability reports



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnReportAPI.VulnReportList(context.Background(), pulpDomain).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnReportAPI.VulnReportList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VulnReportList`: PaginatedserviceVulnerabilityReportResponseList
	fmt.Fprintf(os.Stdout, "Response from `VulnReportAPI.VulnReportList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnReportListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedserviceVulnerabilityReportResponseList**](PaginatedserviceVulnerabilityReportResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VulnReportRead

> ServiceVulnerabilityReportResponse VulnReportRead(ctx, serviceVulnerabilityReportHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a vulnerability report



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	serviceVulnerabilityReportHref := "serviceVulnerabilityReportHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnReportAPI.VulnReportRead(context.Background(), serviceVulnerabilityReportHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnReportAPI.VulnReportRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VulnReportRead`: ServiceVulnerabilityReportResponse
	fmt.Fprintf(os.Stdout, "Response from `VulnReportAPI.VulnReportRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceVulnerabilityReportHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnReportReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ServiceVulnerabilityReportResponse**](ServiceVulnerabilityReportResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

