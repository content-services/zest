# \ContentguardsRhsmApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentguardsCertguardRhsmCreate**](ContentguardsRhsmApi.md#ContentguardsCertguardRhsmCreate) | **Post** /pulp/api/v3/contentguards/certguard/rhsm/ | Create a rhsm cert guard
[**ContentguardsCertguardRhsmDelete**](ContentguardsRhsmApi.md#ContentguardsCertguardRhsmDelete) | **Delete** /{certguard_r_h_s_m_cert_guard_href} | Delete a rhsm cert guard
[**ContentguardsCertguardRhsmList**](ContentguardsRhsmApi.md#ContentguardsCertguardRhsmList) | **Get** /pulp/api/v3/contentguards/certguard/rhsm/ | List rhsm cert guards
[**ContentguardsCertguardRhsmPartialUpdate**](ContentguardsRhsmApi.md#ContentguardsCertguardRhsmPartialUpdate) | **Patch** /{certguard_r_h_s_m_cert_guard_href} | Update a rhsm cert guard
[**ContentguardsCertguardRhsmRead**](ContentguardsRhsmApi.md#ContentguardsCertguardRhsmRead) | **Get** /{certguard_r_h_s_m_cert_guard_href} | Inspect a rhsm cert guard
[**ContentguardsCertguardRhsmUpdate**](ContentguardsRhsmApi.md#ContentguardsCertguardRhsmUpdate) | **Put** /{certguard_r_h_s_m_cert_guard_href} | Update a rhsm cert guard



## ContentguardsCertguardRhsmCreate

> CertguardRHSMCertGuardResponse ContentguardsCertguardRhsmCreate(ctx).CertguardRHSMCertGuard(certguardRHSMCertGuard).Execute()

Create a rhsm cert guard



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release"
)

func main() {
    certguardRHSMCertGuard := *openapiclient.NewCertguardRHSMCertGuard("Name_example", "CaCertificate_example") // CertguardRHSMCertGuard | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsRhsmApi.ContentguardsCertguardRhsmCreate(context.Background()).CertguardRHSMCertGuard(certguardRHSMCertGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRhsmApi.ContentguardsCertguardRhsmCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCertguardRhsmCreate`: CertguardRHSMCertGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRhsmApi.ContentguardsCertguardRhsmCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCertguardRhsmCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certguardRHSMCertGuard** | [**CertguardRHSMCertGuard**](CertguardRHSMCertGuard.md) |  | 

### Return type

[**CertguardRHSMCertGuardResponse**](CertguardRHSMCertGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCertguardRhsmDelete

> ContentguardsCertguardRhsmDelete(ctx, certguardRHSMCertGuardHref).Execute()

Delete a rhsm cert guard



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release"
)

func main() {
    certguardRHSMCertGuardHref := "certguardRHSMCertGuardHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContentguardsRhsmApi.ContentguardsCertguardRhsmDelete(context.Background(), certguardRHSMCertGuardHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRhsmApi.ContentguardsCertguardRhsmDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certguardRHSMCertGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCertguardRhsmDeleteRequest struct via the builder pattern


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


## ContentguardsCertguardRhsmList

> PaginatedcertguardRHSMCertGuardResponseList ContentguardsCertguardRhsmList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()

List rhsm cert guards



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release"
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
    resp, r, err := apiClient.ContentguardsRhsmApi.ContentguardsCertguardRhsmList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRhsmApi.ContentguardsCertguardRhsmList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCertguardRhsmList`: PaginatedcertguardRHSMCertGuardResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRhsmApi.ContentguardsCertguardRhsmList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCertguardRhsmListRequest struct via the builder pattern


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

[**PaginatedcertguardRHSMCertGuardResponseList**](PaginatedcertguardRHSMCertGuardResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCertguardRhsmPartialUpdate

> CertguardRHSMCertGuardResponse ContentguardsCertguardRhsmPartialUpdate(ctx, certguardRHSMCertGuardHref).PatchedcertguardRHSMCertGuard(patchedcertguardRHSMCertGuard).Execute()

Update a rhsm cert guard



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release"
)

func main() {
    certguardRHSMCertGuardHref := "certguardRHSMCertGuardHref_example" // string | 
    patchedcertguardRHSMCertGuard := *openapiclient.NewPatchedcertguardRHSMCertGuard() // PatchedcertguardRHSMCertGuard | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsRhsmApi.ContentguardsCertguardRhsmPartialUpdate(context.Background(), certguardRHSMCertGuardHref).PatchedcertguardRHSMCertGuard(patchedcertguardRHSMCertGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRhsmApi.ContentguardsCertguardRhsmPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCertguardRhsmPartialUpdate`: CertguardRHSMCertGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRhsmApi.ContentguardsCertguardRhsmPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certguardRHSMCertGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCertguardRhsmPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedcertguardRHSMCertGuard** | [**PatchedcertguardRHSMCertGuard**](PatchedcertguardRHSMCertGuard.md) |  | 

### Return type

[**CertguardRHSMCertGuardResponse**](CertguardRHSMCertGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCertguardRhsmRead

> CertguardRHSMCertGuardResponse ContentguardsCertguardRhsmRead(ctx, certguardRHSMCertGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a rhsm cert guard



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release"
)

func main() {
    certguardRHSMCertGuardHref := "certguardRHSMCertGuardHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsRhsmApi.ContentguardsCertguardRhsmRead(context.Background(), certguardRHSMCertGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRhsmApi.ContentguardsCertguardRhsmRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCertguardRhsmRead`: CertguardRHSMCertGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRhsmApi.ContentguardsCertguardRhsmRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certguardRHSMCertGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCertguardRhsmReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**CertguardRHSMCertGuardResponse**](CertguardRHSMCertGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCertguardRhsmUpdate

> CertguardRHSMCertGuardResponse ContentguardsCertguardRhsmUpdate(ctx, certguardRHSMCertGuardHref).CertguardRHSMCertGuard(certguardRHSMCertGuard).Execute()

Update a rhsm cert guard



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release"
)

func main() {
    certguardRHSMCertGuardHref := "certguardRHSMCertGuardHref_example" // string | 
    certguardRHSMCertGuard := *openapiclient.NewCertguardRHSMCertGuard("Name_example", "CaCertificate_example") // CertguardRHSMCertGuard | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsRhsmApi.ContentguardsCertguardRhsmUpdate(context.Background(), certguardRHSMCertGuardHref).CertguardRHSMCertGuard(certguardRHSMCertGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsRhsmApi.ContentguardsCertguardRhsmUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCertguardRhsmUpdate`: CertguardRHSMCertGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsRhsmApi.ContentguardsCertguardRhsmUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certguardRHSMCertGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCertguardRhsmUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certguardRHSMCertGuard** | [**CertguardRHSMCertGuard**](CertguardRHSMCertGuard.md) |  | 

### Return type

[**CertguardRHSMCertGuardResponse**](CertguardRHSMCertGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

