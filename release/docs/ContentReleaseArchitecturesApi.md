# \ContentReleaseArchitecturesApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebReleaseArchitecturesCreate**](ContentReleaseArchitecturesApi.md#ContentDebReleaseArchitecturesCreate) | **Post** /pulp/api/v3/content/deb/release_architectures/ | Create a release architecture
[**ContentDebReleaseArchitecturesList**](ContentReleaseArchitecturesApi.md#ContentDebReleaseArchitecturesList) | **Get** /pulp/api/v3/content/deb/release_architectures/ | List release architectures
[**ContentDebReleaseArchitecturesRead**](ContentReleaseArchitecturesApi.md#ContentDebReleaseArchitecturesRead) | **Get** /{deb_release_architecture_href} | Inspect a release architecture



## ContentDebReleaseArchitecturesCreate

> DebReleaseArchitectureResponse ContentDebReleaseArchitecturesCreate(ctx).DebReleaseArchitecture(debReleaseArchitecture).Execute()

Create a release architecture



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release/v3"
)

func main() {
    debReleaseArchitecture := *openapiclient.NewDebReleaseArchitecture("Architecture_example", "Release_example") // DebReleaseArchitecture | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentReleaseArchitecturesApi.ContentDebReleaseArchitecturesCreate(context.Background()).DebReleaseArchitecture(debReleaseArchitecture).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleaseArchitecturesApi.ContentDebReleaseArchitecturesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleaseArchitecturesCreate`: DebReleaseArchitectureResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentReleaseArchitecturesApi.ContentDebReleaseArchitecturesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleaseArchitecturesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debReleaseArchitecture** | [**DebReleaseArchitecture**](DebReleaseArchitecture.md) |  | 

### Return type

[**DebReleaseArchitectureResponse**](DebReleaseArchitectureResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebReleaseArchitecturesList

> PaginateddebReleaseArchitectureResponseList ContentDebReleaseArchitecturesList(ctx).Architecture(architecture).Limit(limit).Offset(offset).Ordering(ordering).Release(release).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List release architectures



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release/v3"
)

func main() {
    architecture := "architecture_example" // string | Filter results where architecture matches value (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    release := "release_example" // string | Filter results where release matches value (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentReleaseArchitecturesApi.ContentDebReleaseArchitecturesList(context.Background()).Architecture(architecture).Limit(limit).Offset(offset).Ordering(ordering).Release(release).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleaseArchitecturesApi.ContentDebReleaseArchitecturesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleaseArchitecturesList`: PaginateddebReleaseArchitectureResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentReleaseArchitecturesApi.ContentDebReleaseArchitecturesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleaseArchitecturesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **architecture** | **string** | Filter results where architecture matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **release** | **string** | Filter results where release matches value | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebReleaseArchitectureResponseList**](PaginateddebReleaseArchitectureResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebReleaseArchitecturesRead

> DebReleaseArchitectureResponse ContentDebReleaseArchitecturesRead(ctx, debReleaseArchitectureHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a release architecture



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release/v3"
)

func main() {
    debReleaseArchitectureHref := "debReleaseArchitectureHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentReleaseArchitecturesApi.ContentDebReleaseArchitecturesRead(context.Background(), debReleaseArchitectureHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleaseArchitecturesApi.ContentDebReleaseArchitecturesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleaseArchitecturesRead`: DebReleaseArchitectureResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentReleaseArchitecturesApi.ContentDebReleaseArchitecturesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debReleaseArchitectureHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleaseArchitecturesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**DebReleaseArchitectureResponse**](DebReleaseArchitectureResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

