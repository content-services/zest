# \ContentReleaseFilesApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebReleaseFilesCreate**](ContentReleaseFilesApi.md#ContentDebReleaseFilesCreate) | **Post** /pulp/api/v3/content/deb/release_files/ | Create a release file
[**ContentDebReleaseFilesList**](ContentReleaseFilesApi.md#ContentDebReleaseFilesList) | **Get** /pulp/api/v3/content/deb/release_files/ | List release files
[**ContentDebReleaseFilesRead**](ContentReleaseFilesApi.md#ContentDebReleaseFilesRead) | **Get** /{deb_release_file_href} | Inspect a release file



## ContentDebReleaseFilesCreate

> DebReleaseFileResponse ContentDebReleaseFilesCreate(ctx).DebReleaseFile(debReleaseFile).Execute()

Create a release file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/pulpGoBinding/packages/pulpGoBinding"
)

func main() {
    debReleaseFile := *openapiclient.NewDebReleaseFile(map[string]interface{}(123), "Distribution_example") // DebReleaseFile | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentReleaseFilesApi.ContentDebReleaseFilesCreate(context.Background()).DebReleaseFile(debReleaseFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleaseFilesApi.ContentDebReleaseFilesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleaseFilesCreate`: DebReleaseFileResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentReleaseFilesApi.ContentDebReleaseFilesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleaseFilesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debReleaseFile** | [**DebReleaseFile**](DebReleaseFile.md) |  | 

### Return type

[**DebReleaseFileResponse**](DebReleaseFileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebReleaseFilesList

> PaginateddebReleaseFileResponseList ContentDebReleaseFilesList(ctx).Codename(codename).Limit(limit).Offset(offset).Ordering(ordering).RelativePath(relativePath).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Suite(suite).Fields(fields).ExcludeFields(excludeFields).Execute()

List release files



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/pulpGoBinding/packages/pulpGoBinding"
)

func main() {
    codename := "codename_example" // string | Filter results where codename matches value (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    relativePath := "relativePath_example" // string | Filter results where relative_path matches value (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    sha256 := "sha256_example" // string | Filter results where sha256 matches value (optional)
    suite := "suite_example" // string | Filter results where suite matches value (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentReleaseFilesApi.ContentDebReleaseFilesList(context.Background()).Codename(codename).Limit(limit).Offset(offset).Ordering(ordering).RelativePath(relativePath).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Suite(suite).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleaseFilesApi.ContentDebReleaseFilesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleaseFilesList`: PaginateddebReleaseFileResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentReleaseFilesApi.ContentDebReleaseFilesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleaseFilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codename** | **string** | Filter results where codename matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **relativePath** | **string** | Filter results where relative_path matches value | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **sha256** | **string** | Filter results where sha256 matches value | 
 **suite** | **string** | Filter results where suite matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebReleaseFileResponseList**](PaginateddebReleaseFileResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebReleaseFilesRead

> DebReleaseFileResponse ContentDebReleaseFilesRead(ctx, debReleaseFileHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a release file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/pulpGoBinding/packages/pulpGoBinding"
)

func main() {
    debReleaseFileHref := "debReleaseFileHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentReleaseFilesApi.ContentDebReleaseFilesRead(context.Background(), debReleaseFileHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentReleaseFilesApi.ContentDebReleaseFilesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebReleaseFilesRead`: DebReleaseFileResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentReleaseFilesApi.ContentDebReleaseFilesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debReleaseFileHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebReleaseFilesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**DebReleaseFileResponse**](DebReleaseFileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

