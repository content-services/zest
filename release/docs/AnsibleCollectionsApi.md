# \AnsibleCollectionsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnsibleCollectionsList**](AnsibleCollectionsApi.md#AnsibleCollectionsList) | **Get** /pulp/api/v3/ansible/collections/ | List collections
[**UploadCollection**](AnsibleCollectionsApi.md#UploadCollection) | **Post** /ansible/collections/ | Upload a collection



## AnsibleCollectionsList

> PaginatedansibleCollectionResponseList AnsibleCollectionsList(ctx).Limit(limit).Name(name).Namespace(namespace).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()

List collections



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
    name := "name_example" // string |  (optional)
    namespace := "namespace_example" // string |  (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnsibleCollectionsApi.AnsibleCollectionsList(context.Background()).Limit(limit).Name(name).Namespace(namespace).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnsibleCollectionsApi.AnsibleCollectionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AnsibleCollectionsList`: PaginatedansibleCollectionResponseList
    fmt.Fprintf(os.Stdout, "Response from `AnsibleCollectionsApi.AnsibleCollectionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnsibleCollectionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **namespace** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedansibleCollectionResponseList**](PaginatedansibleCollectionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadCollection

> AsyncOperationResponse UploadCollection(ctx).File(file).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()

Upload a collection



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
    file := os.NewFile(1234, "some_file") // *os.File | The Collection tarball.
    sha256 := "sha256_example" // string | An optional sha256 checksum of the uploaded file. (optional)
    expectedNamespace := "expectedNamespace_example" // string | The expected 'namespace' of the Collection to be verified against the metadata during import. (optional)
    expectedName := "expectedName_example" // string | The expected 'name' of the Collection to be verified against the metadata during import. (optional)
    expectedVersion := "expectedVersion_example" // string | The expected version of the Collection to be verified against the metadata during import. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnsibleCollectionsApi.UploadCollection(context.Background()).File(file).Sha256(sha256).ExpectedNamespace(expectedNamespace).ExpectedName(expectedName).ExpectedVersion(expectedVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnsibleCollectionsApi.UploadCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadCollection`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AnsibleCollectionsApi.UploadCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The Collection tarball. | 
 **sha256** | **string** | An optional sha256 checksum of the uploaded file. | 
 **expectedNamespace** | **string** | The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import. | 
 **expectedName** | **string** | The expected &#39;name&#39; of the Collection to be verified against the metadata during import. | 
 **expectedVersion** | **string** | The expected version of the Collection to be verified against the metadata during import. | 

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

