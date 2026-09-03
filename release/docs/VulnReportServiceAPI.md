# \VulnReportServiceAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VulnReportServiceAddRole**](VulnReportServiceAPI.md#VulnReportServiceAddRole) | **Post** /{service_vulnerability_report_href}add_role/ | Add a role
[**VulnReportServiceCreate**](VulnReportServiceAPI.md#VulnReportServiceCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/vuln_report_service/ | Generate vulnerability report
[**VulnReportServiceDelete**](VulnReportServiceAPI.md#VulnReportServiceDelete) | **Delete** /{service_vulnerability_report_href} | Delete a vulnerability report
[**VulnReportServiceList**](VulnReportServiceAPI.md#VulnReportServiceList) | **Get** /api/pulp/{pulp_domain}/api/v3/vuln_report_service/ | List vulnerability reports
[**VulnReportServiceListRoles**](VulnReportServiceAPI.md#VulnReportServiceListRoles) | **Get** /{service_vulnerability_report_href}list_roles/ | List roles
[**VulnReportServiceMyPermissions**](VulnReportServiceAPI.md#VulnReportServiceMyPermissions) | **Get** /{service_vulnerability_report_href}my_permissions/ | List user permissions
[**VulnReportServiceRead**](VulnReportServiceAPI.md#VulnReportServiceRead) | **Get** /{service_vulnerability_report_href} | Inspect a vulnerability report
[**VulnReportServiceRemoveRole**](VulnReportServiceAPI.md#VulnReportServiceRemoveRole) | **Post** /{service_vulnerability_report_href}remove_role/ | Remove a role



## VulnReportServiceAddRole

> NestedRoleResponse VulnReportServiceAddRole(ctx, serviceVulnerabilityReportHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()

Add a role



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
	serviceVulnerabilityReportHref := "serviceVulnerabilityReportHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnReportServiceAPI.VulnReportServiceAddRole(context.Background(), serviceVulnerabilityReportHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnReportServiceAPI.VulnReportServiceAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VulnReportServiceAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `VulnReportServiceAPI.VulnReportServiceAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceVulnerabilityReportHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnReportServiceAddRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nestedRole** | [**NestedRole**](NestedRole.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**NestedRoleResponse**](NestedRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VulnReportServiceCreate

> AsyncOperationResponse VulnReportServiceCreate(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).RepoVersion(repoVersion).PackageJson(packageJson).Execute()

Generate vulnerability report



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
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	repoVersion := "repoVersion_example" // string | RepositoryVersion HREF with the packages to be checked. (optional)
	packageJson := os.NewFile(1234, "some_file") // *os.File | package-lock.json file with the definition of dependencies to be checked. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnReportServiceAPI.VulnReportServiceCreate(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).RepoVersion(repoVersion).PackageJson(packageJson).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnReportServiceAPI.VulnReportServiceCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VulnReportServiceCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `VulnReportServiceAPI.VulnReportServiceCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnReportServiceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **repoVersion** | **string** | RepositoryVersion HREF with the packages to be checked. | 
 **packageJson** | ***os.File** | package-lock.json file with the definition of dependencies to be checked. | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VulnReportServiceDelete

> VulnReportServiceDelete(ctx, serviceVulnerabilityReportHref).XTaskDiagnostics(xTaskDiagnostics).Execute()

Delete a vulnerability report



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
	serviceVulnerabilityReportHref := "serviceVulnerabilityReportHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VulnReportServiceAPI.VulnReportServiceDelete(context.Background(), serviceVulnerabilityReportHref).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnReportServiceAPI.VulnReportServiceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiVulnReportServiceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## VulnReportServiceList

> PaginatedserviceVulnerabilityReportResponseList VulnReportServiceList(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List vulnerability reports



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
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnReportServiceAPI.VulnReportServiceList(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnReportServiceAPI.VulnReportServiceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VulnReportServiceList`: PaginatedserviceVulnerabilityReportResponseList
	fmt.Fprintf(os.Stdout, "Response from `VulnReportServiceAPI.VulnReportServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnReportServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
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


## VulnReportServiceListRoles

> ObjectRolesResponse VulnReportServiceListRoles(ctx, serviceVulnerabilityReportHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

List roles



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
	serviceVulnerabilityReportHref := "serviceVulnerabilityReportHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnReportServiceAPI.VulnReportServiceListRoles(context.Background(), serviceVulnerabilityReportHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnReportServiceAPI.VulnReportServiceListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VulnReportServiceListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `VulnReportServiceAPI.VulnReportServiceListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceVulnerabilityReportHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnReportServiceListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ObjectRolesResponse**](ObjectRolesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VulnReportServiceMyPermissions

> MyPermissionsResponse VulnReportServiceMyPermissions(ctx, serviceVulnerabilityReportHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

List user permissions



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
	serviceVulnerabilityReportHref := "serviceVulnerabilityReportHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnReportServiceAPI.VulnReportServiceMyPermissions(context.Background(), serviceVulnerabilityReportHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnReportServiceAPI.VulnReportServiceMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VulnReportServiceMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `VulnReportServiceAPI.VulnReportServiceMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceVulnerabilityReportHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnReportServiceMyPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**MyPermissionsResponse**](MyPermissionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VulnReportServiceRead

> ServiceVulnerabilityReportResponse VulnReportServiceRead(ctx, serviceVulnerabilityReportHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a vulnerability report



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
	serviceVulnerabilityReportHref := "serviceVulnerabilityReportHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnReportServiceAPI.VulnReportServiceRead(context.Background(), serviceVulnerabilityReportHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnReportServiceAPI.VulnReportServiceRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VulnReportServiceRead`: ServiceVulnerabilityReportResponse
	fmt.Fprintf(os.Stdout, "Response from `VulnReportServiceAPI.VulnReportServiceRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceVulnerabilityReportHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnReportServiceReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
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


## VulnReportServiceRemoveRole

> NestedRoleResponse VulnReportServiceRemoveRole(ctx, serviceVulnerabilityReportHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()

Remove a role



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
	serviceVulnerabilityReportHref := "serviceVulnerabilityReportHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnReportServiceAPI.VulnReportServiceRemoveRole(context.Background(), serviceVulnerabilityReportHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnReportServiceAPI.VulnReportServiceRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VulnReportServiceRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `VulnReportServiceAPI.VulnReportServiceRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceVulnerabilityReportHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnReportServiceRemoveRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nestedRole** | [**NestedRole**](NestedRole.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**NestedRoleResponse**](NestedRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

