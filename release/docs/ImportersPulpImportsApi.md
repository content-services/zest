# \ImportersPulpImportsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportersCorePulpImportsCreate**](ImportersPulpImportsApi.md#ImportersCorePulpImportsCreate) | **Post** /{pulp_importer_href}imports/ | Create a pulp import
[**ImportersCorePulpImportsDelete**](ImportersPulpImportsApi.md#ImportersCorePulpImportsDelete) | **Delete** /{pulp_pulp_import_href} | Delete a pulp import
[**ImportersCorePulpImportsList**](ImportersPulpImportsApi.md#ImportersCorePulpImportsList) | **Get** /{pulp_importer_href}imports/ | List pulp imports
[**ImportersCorePulpImportsRead**](ImportersPulpImportsApi.md#ImportersCorePulpImportsRead) | **Get** /{pulp_pulp_import_href} | Inspect a pulp import



## ImportersCorePulpImportsCreate

> TaskGroupOperationResponse ImportersCorePulpImportsCreate(ctx, pulpImporterHref).PulpImport(pulpImport).Execute()

Create a pulp import



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release"
)

func main() {
    pulpImporterHref := "pulpImporterHref_example" // string | 
    pulpImport := *openapiclient.NewPulpImport() // PulpImport | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportersPulpImportsApi.ImportersCorePulpImportsCreate(context.Background(), pulpImporterHref).PulpImport(pulpImport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpImportsApi.ImportersCorePulpImportsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportersCorePulpImportsCreate`: TaskGroupOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportersPulpImportsApi.ImportersCorePulpImportsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpImporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportersCorePulpImportsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pulpImport** | [**PulpImport**](PulpImport.md) |  | 

### Return type

[**TaskGroupOperationResponse**](TaskGroupOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportersCorePulpImportsDelete

> ImportersCorePulpImportsDelete(ctx, pulpPulpImportHref).Execute()

Delete a pulp import



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release"
)

func main() {
    pulpPulpImportHref := "pulpPulpImportHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImportersPulpImportsApi.ImportersCorePulpImportsDelete(context.Background(), pulpPulpImportHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpImportsApi.ImportersCorePulpImportsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpPulpImportHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportersCorePulpImportsDeleteRequest struct via the builder pattern


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


## ImportersCorePulpImportsList

> PaginatedImportResponseList ImportersCorePulpImportsList(ctx, pulpImporterHref).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List pulp imports



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release"
)

func main() {
    pulpImporterHref := "pulpImporterHref_example" // string | 
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportersPulpImportsApi.ImportersCorePulpImportsList(context.Background(), pulpImporterHref).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpImportsApi.ImportersCorePulpImportsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportersCorePulpImportsList`: PaginatedImportResponseList
    fmt.Fprintf(os.Stdout, "Response from `ImportersPulpImportsApi.ImportersCorePulpImportsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpImporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportersCorePulpImportsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedImportResponseList**](PaginatedImportResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportersCorePulpImportsRead

> ImportResponse ImportersCorePulpImportsRead(ctx, pulpPulpImportHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a pulp import



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release"
)

func main() {
    pulpPulpImportHref := "pulpPulpImportHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportersPulpImportsApi.ImportersCorePulpImportsRead(context.Background(), pulpPulpImportHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportersPulpImportsApi.ImportersCorePulpImportsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportersCorePulpImportsRead`: ImportResponse
    fmt.Fprintf(os.Stdout, "Response from `ImportersPulpImportsApi.ImportersCorePulpImportsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpPulpImportHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportersCorePulpImportsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ImportResponse**](ImportResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

