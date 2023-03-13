# \ContentCollectionVersionsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentAnsibleCollectionVersionsCreate**](ContentCollectionVersionsApi.md#ContentAnsibleCollectionVersionsCreate) | **Post** /pulp/api/v3/content/ansible/collection_versions/ | Create a collection version
[**ContentAnsibleCollectionVersionsList**](ContentCollectionVersionsApi.md#ContentAnsibleCollectionVersionsList) | **Get** /pulp/api/v3/content/ansible/collection_versions/ | List collection versions
[**ContentAnsibleCollectionVersionsRead**](ContentCollectionVersionsApi.md#ContentAnsibleCollectionVersionsRead) | **Get** /{ansible_collection_version_href} | Inspect a collection version



## ContentAnsibleCollectionVersionsCreate

> AsyncOperationResponse ContentAnsibleCollectionVersionsCreate(ctx).Artifact(artifact).Upload(upload).File(file).Repository(repository).ExpectedName(expectedName).ExpectedNamespace(expectedNamespace).ExpectedVersion(expectedVersion).Execute()

Create a collection version



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
    artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
    upload := "upload_example" // string | An uncommitted upload that may be turned into the artifact of the content unit. (optional)
    file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the artifact of the content unit. (optional)
    repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
    expectedName := "expectedName_example" // string | The name of the collection. (optional)
    expectedNamespace := "expectedNamespace_example" // string | The namespace of the collection. (optional)
    expectedVersion := "expectedVersion_example" // string | The version of the collection. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentCollectionVersionsApi.ContentAnsibleCollectionVersionsCreate(context.Background()).Artifact(artifact).Upload(upload).File(file).Repository(repository).ExpectedName(expectedName).ExpectedNamespace(expectedNamespace).ExpectedVersion(expectedVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionVersionsApi.ContentAnsibleCollectionVersionsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentAnsibleCollectionVersionsCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentCollectionVersionsApi.ContentAnsibleCollectionVersionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionVersionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifact** | **string** | Artifact file representing the physical content | 
 **upload** | **string** | An uncommitted upload that may be turned into the artifact of the content unit. | 
 **file** | ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | 
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **expectedName** | **string** | The name of the collection. | 
 **expectedNamespace** | **string** | The namespace of the collection. | 
 **expectedVersion** | **string** | The version of the collection. | 

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


## ContentAnsibleCollectionVersionsList

> PaginatedansibleCollectionVersionResponseList ContentAnsibleCollectionVersionsList(ctx).IsHighest(isHighest).Limit(limit).Name(name).Namespace(namespace).Offset(offset).Ordering(ordering).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Tags(tags).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()

List collection versions



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
    isHighest := true // bool |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string |  (optional)
    namespace := "namespace_example" // string |  (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    q := "q_example" // string |  (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    tags := "tags_example" // string | Filter by comma separate list of tags that must all be matched (optional)
    version := "version_example" // string | Filter results where version matches value (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentCollectionVersionsApi.ContentAnsibleCollectionVersionsList(context.Background()).IsHighest(isHighest).Limit(limit).Name(name).Namespace(namespace).Offset(offset).Ordering(ordering).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Tags(tags).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionVersionsApi.ContentAnsibleCollectionVersionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentAnsibleCollectionVersionsList`: PaginatedansibleCollectionVersionResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentCollectionVersionsApi.ContentAnsibleCollectionVersionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionVersionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isHighest** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **namespace** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **q** | **string** |  | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **tags** | **string** | Filter by comma separate list of tags that must all be matched | 
 **version** | **string** | Filter results where version matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedansibleCollectionVersionResponseList**](PaginatedansibleCollectionVersionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentAnsibleCollectionVersionsRead

> AnsibleCollectionVersionResponse ContentAnsibleCollectionVersionsRead(ctx, ansibleCollectionVersionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a collection version



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
    ansibleCollectionVersionHref := "ansibleCollectionVersionHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentCollectionVersionsApi.ContentAnsibleCollectionVersionsRead(context.Background(), ansibleCollectionVersionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionVersionsApi.ContentAnsibleCollectionVersionsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentAnsibleCollectionVersionsRead`: AnsibleCollectionVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentCollectionVersionsApi.ContentAnsibleCollectionVersionsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionVersionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionVersionsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**AnsibleCollectionVersionResponse**](AnsibleCollectionVersionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

