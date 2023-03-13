# \ExportersPulpApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportersCorePulpCreate**](ExportersPulpApi.md#ExportersCorePulpCreate) | **Post** /pulp/api/v3/exporters/core/pulp/ | Create a pulp exporter
[**ExportersCorePulpDelete**](ExportersPulpApi.md#ExportersCorePulpDelete) | **Delete** /{pulp_exporter_href} | Delete a pulp exporter
[**ExportersCorePulpList**](ExportersPulpApi.md#ExportersCorePulpList) | **Get** /pulp/api/v3/exporters/core/pulp/ | List pulp exporters
[**ExportersCorePulpPartialUpdate**](ExportersPulpApi.md#ExportersCorePulpPartialUpdate) | **Patch** /{pulp_exporter_href} | Update a pulp exporter
[**ExportersCorePulpRead**](ExportersPulpApi.md#ExportersCorePulpRead) | **Get** /{pulp_exporter_href} | Inspect a pulp exporter
[**ExportersCorePulpUpdate**](ExportersPulpApi.md#ExportersCorePulpUpdate) | **Put** /{pulp_exporter_href} | Update a pulp exporter



## ExportersCorePulpCreate

> PulpExporterResponse ExportersCorePulpCreate(ctx).PulpExporter(pulpExporter).Execute()

Create a pulp exporter



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/packages/pulpGoBinding"
)

func main() {
    pulpExporter := *openapiclient.NewPulpExporter("Name_example", "Path_example", []string{"Repositories_example"}) // PulpExporter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersPulpApi.ExportersCorePulpCreate(context.Background()).PulpExporter(pulpExporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpApi.ExportersCorePulpCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpCreate`: PulpExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpApi.ExportersCorePulpCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pulpExporter** | [**PulpExporter**](PulpExporter.md) |  | 

### Return type

[**PulpExporterResponse**](PulpExporterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCorePulpDelete

> AsyncOperationResponse ExportersCorePulpDelete(ctx, pulpExporterHref).Execute()

Delete a pulp exporter



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/packages/pulpGoBinding"
)

func main() {
    pulpExporterHref := "pulpExporterHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersPulpApi.ExportersCorePulpDelete(context.Background(), pulpExporterHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpApi.ExportersCorePulpDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpApi.ExportersCorePulpDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpExporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ExportersCorePulpList

> PaginatedPulpExporterResponseList ExportersCorePulpList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()

List pulp exporters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/packages/pulpGoBinding"
)

func main() {
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersPulpApi.ExportersCorePulpList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpApi.ExportersCorePulpList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpList`: PaginatedPulpExporterResponseList
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpApi.ExportersCorePulpList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedPulpExporterResponseList**](PaginatedPulpExporterResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCorePulpPartialUpdate

> AsyncOperationResponse ExportersCorePulpPartialUpdate(ctx, pulpExporterHref).PatchedPulpExporter(patchedPulpExporter).Execute()

Update a pulp exporter



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/packages/pulpGoBinding"
)

func main() {
    pulpExporterHref := "pulpExporterHref_example" // string | 
    patchedPulpExporter := *openapiclient.NewPatchedPulpExporter() // PatchedPulpExporter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersPulpApi.ExportersCorePulpPartialUpdate(context.Background(), pulpExporterHref).PatchedPulpExporter(patchedPulpExporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpApi.ExportersCorePulpPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpApi.ExportersCorePulpPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpExporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPulpExporter** | [**PatchedPulpExporter**](PatchedPulpExporter.md) |  | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCorePulpRead

> PulpExporterResponse ExportersCorePulpRead(ctx, pulpExporterHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a pulp exporter



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/packages/pulpGoBinding"
)

func main() {
    pulpExporterHref := "pulpExporterHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersPulpApi.ExportersCorePulpRead(context.Background(), pulpExporterHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpApi.ExportersCorePulpRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpRead`: PulpExporterResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpApi.ExportersCorePulpRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpExporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PulpExporterResponse**](PulpExporterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportersCorePulpUpdate

> AsyncOperationResponse ExportersCorePulpUpdate(ctx, pulpExporterHref).PulpExporter(pulpExporter).Execute()

Update a pulp exporter



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/packages/pulpGoBinding"
)

func main() {
    pulpExporterHref := "pulpExporterHref_example" // string | 
    pulpExporter := *openapiclient.NewPulpExporter("Name_example", "Path_example", []string{"Repositories_example"}) // PulpExporter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportersPulpApi.ExportersCorePulpUpdate(context.Background(), pulpExporterHref).PulpExporter(pulpExporter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportersPulpApi.ExportersCorePulpUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportersCorePulpUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ExportersPulpApi.ExportersCorePulpUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpExporterHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportersCorePulpUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pulpExporter** | [**PulpExporter**](PulpExporter.md) |  | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

