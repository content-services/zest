# \ContentModulemdDefaultsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentRpmModulemdDefaultsCreate**](ContentModulemdDefaultsApi.md#ContentRpmModulemdDefaultsCreate) | **Post** /pulp/api/v3/content/rpm/modulemd_defaults/ | Create a modulemd defaults
[**ContentRpmModulemdDefaultsList**](ContentModulemdDefaultsApi.md#ContentRpmModulemdDefaultsList) | **Get** /pulp/api/v3/content/rpm/modulemd_defaults/ | List modulemd defaultss
[**ContentRpmModulemdDefaultsRead**](ContentModulemdDefaultsApi.md#ContentRpmModulemdDefaultsRead) | **Get** /{rpm_modulemd_defaults_href} | Inspect a modulemd defaults



## ContentRpmModulemdDefaultsCreate

> AsyncOperationResponse ContentRpmModulemdDefaultsCreate(ctx).RpmModulemdDefaults(rpmModulemdDefaults).Execute()

Create a modulemd defaults



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
    rpmModulemdDefaults := *openapiclient.NewRpmModulemdDefaults("Module_example", "Stream_example", map[string]interface{}(123), "Snippet_example") // RpmModulemdDefaults | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentModulemdDefaultsApi.ContentRpmModulemdDefaultsCreate(context.Background()).RpmModulemdDefaults(rpmModulemdDefaults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentModulemdDefaultsApi.ContentRpmModulemdDefaultsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmModulemdDefaultsCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentModulemdDefaultsApi.ContentRpmModulemdDefaultsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmModulemdDefaultsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rpmModulemdDefaults** | [**RpmModulemdDefaults**](RpmModulemdDefaults.md) |  | 

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


## ContentRpmModulemdDefaultsList

> PaginatedrpmModulemdDefaultsResponseList ContentRpmModulemdDefaultsList(ctx).Limit(limit).Module(module).ModuleIn(moduleIn).Offset(offset).Ordering(ordering).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Stream(stream).StreamIn(streamIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List modulemd defaultss



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
    module := "module_example" // string | Filter results where module matches value (optional)
    moduleIn := []string{"Inner_example"} // []string | Filter results where module is in a comma-separated list of values (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    sha256 := "sha256_example" // string |  (optional)
    stream := "stream_example" // string | Filter results where stream matches value (optional)
    streamIn := []string{"Inner_example"} // []string | Filter results where stream is in a comma-separated list of values (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentModulemdDefaultsApi.ContentRpmModulemdDefaultsList(context.Background()).Limit(limit).Module(module).ModuleIn(moduleIn).Offset(offset).Ordering(ordering).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Stream(stream).StreamIn(streamIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentModulemdDefaultsApi.ContentRpmModulemdDefaultsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmModulemdDefaultsList`: PaginatedrpmModulemdDefaultsResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentModulemdDefaultsApi.ContentRpmModulemdDefaultsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmModulemdDefaultsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **module** | **string** | Filter results where module matches value | 
 **moduleIn** | **[]string** | Filter results where module is in a comma-separated list of values | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **sha256** | **string** |  | 
 **stream** | **string** | Filter results where stream matches value | 
 **streamIn** | **[]string** | Filter results where stream is in a comma-separated list of values | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedrpmModulemdDefaultsResponseList**](PaginatedrpmModulemdDefaultsResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentRpmModulemdDefaultsRead

> RpmModulemdDefaultsResponse ContentRpmModulemdDefaultsRead(ctx, rpmModulemdDefaultsHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a modulemd defaults



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
    rpmModulemdDefaultsHref := "rpmModulemdDefaultsHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentModulemdDefaultsApi.ContentRpmModulemdDefaultsRead(context.Background(), rpmModulemdDefaultsHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentModulemdDefaultsApi.ContentRpmModulemdDefaultsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmModulemdDefaultsRead`: RpmModulemdDefaultsResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentModulemdDefaultsApi.ContentRpmModulemdDefaultsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmModulemdDefaultsHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmModulemdDefaultsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RpmModulemdDefaultsResponse**](RpmModulemdDefaultsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

