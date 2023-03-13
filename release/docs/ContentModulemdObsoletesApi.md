# \ContentModulemdObsoletesApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentRpmModulemdObsoletesCreate**](ContentModulemdObsoletesApi.md#ContentRpmModulemdObsoletesCreate) | **Post** /pulp/api/v3/content/rpm/modulemd_obsoletes/ | Create a modulemd obsolete
[**ContentRpmModulemdObsoletesList**](ContentModulemdObsoletesApi.md#ContentRpmModulemdObsoletesList) | **Get** /pulp/api/v3/content/rpm/modulemd_obsoletes/ | List modulemd obsoletes
[**ContentRpmModulemdObsoletesRead**](ContentModulemdObsoletesApi.md#ContentRpmModulemdObsoletesRead) | **Get** /{rpm_modulemd_obsolete_href} | Inspect a modulemd obsolete



## ContentRpmModulemdObsoletesCreate

> AsyncOperationResponse ContentRpmModulemdObsoletesCreate(ctx).RpmModulemdObsolete(rpmModulemdObsolete).Execute()

Create a modulemd obsolete



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
    rpmModulemdObsolete := *openapiclient.NewRpmModulemdObsolete("Modified_example", "ModuleName_example", "ModuleStream_example", "Message_example", "OverridePrevious_example", "ModuleContext_example", "EolDate_example", "ObsoletedByModuleName_example", "ObsoletedByModuleStream_example", "Snippet_example") // RpmModulemdObsolete | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentModulemdObsoletesApi.ContentRpmModulemdObsoletesCreate(context.Background()).RpmModulemdObsolete(rpmModulemdObsolete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentModulemdObsoletesApi.ContentRpmModulemdObsoletesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmModulemdObsoletesCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentModulemdObsoletesApi.ContentRpmModulemdObsoletesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmModulemdObsoletesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rpmModulemdObsolete** | [**RpmModulemdObsolete**](RpmModulemdObsolete.md) |  | 

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


## ContentRpmModulemdObsoletesList

> PaginatedrpmModulemdObsoleteResponseList ContentRpmModulemdObsoletesList(ctx).Limit(limit).Offset(offset).Ordering(ordering).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List modulemd obsoletes



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
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentModulemdObsoletesApi.ContentRpmModulemdObsoletesList(context.Background()).Limit(limit).Offset(offset).Ordering(ordering).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentModulemdObsoletesApi.ContentRpmModulemdObsoletesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmModulemdObsoletesList`: PaginatedrpmModulemdObsoleteResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentModulemdObsoletesApi.ContentRpmModulemdObsoletesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmModulemdObsoletesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedrpmModulemdObsoleteResponseList**](PaginatedrpmModulemdObsoleteResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentRpmModulemdObsoletesRead

> RpmModulemdObsoleteResponse ContentRpmModulemdObsoletesRead(ctx, rpmModulemdObsoleteHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a modulemd obsolete



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
    rpmModulemdObsoleteHref := "rpmModulemdObsoleteHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentModulemdObsoletesApi.ContentRpmModulemdObsoletesRead(context.Background(), rpmModulemdObsoleteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentModulemdObsoletesApi.ContentRpmModulemdObsoletesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmModulemdObsoletesRead`: RpmModulemdObsoleteResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentModulemdObsoletesApi.ContentRpmModulemdObsoletesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmModulemdObsoleteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmModulemdObsoletesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RpmModulemdObsoleteResponse**](RpmModulemdObsoleteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

