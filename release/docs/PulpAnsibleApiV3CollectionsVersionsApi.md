# \PulpAnsibleApiV3CollectionsVersionsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyApiV3CollectionsVersionsDelete**](PulpAnsibleApiV3CollectionsVersionsApi.md#PulpAnsibleGalaxyApiV3CollectionsVersionsDelete) | **Delete** /pulp_ansible/galaxy/{path}/api/v3/collections/{namespace}/{name}/versions/{version}/ | 
[**PulpAnsibleGalaxyApiV3CollectionsVersionsList**](PulpAnsibleApiV3CollectionsVersionsApi.md#PulpAnsibleGalaxyApiV3CollectionsVersionsList) | **Get** /pulp_ansible/galaxy/{path}/api/v3/collections/{namespace}/{name}/versions/ | 
[**PulpAnsibleGalaxyApiV3CollectionsVersionsRead**](PulpAnsibleApiV3CollectionsVersionsApi.md#PulpAnsibleGalaxyApiV3CollectionsVersionsRead) | **Get** /pulp_ansible/galaxy/{path}/api/v3/collections/{namespace}/{name}/versions/{version}/ | 



## PulpAnsibleGalaxyApiV3CollectionsVersionsDelete

> AsyncOperationResponse PulpAnsibleGalaxyApiV3CollectionsVersionsDelete(ctx, name, namespace, path, version).Execute()





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
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    path := "path_example" // string | 
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleApiV3CollectionsVersionsApi.PulpAnsibleGalaxyApiV3CollectionsVersionsDelete(context.Background(), name, namespace, path, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleApiV3CollectionsVersionsApi.PulpAnsibleGalaxyApiV3CollectionsVersionsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyApiV3CollectionsVersionsDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleApiV3CollectionsVersionsApi.PulpAnsibleGalaxyApiV3CollectionsVersionsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 
**path** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyApiV3CollectionsVersionsDeleteRequest struct via the builder pattern


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


## PulpAnsibleGalaxyApiV3CollectionsVersionsList

> PaginatedCollectionVersionListResponseList PulpAnsibleGalaxyApiV3CollectionsVersionsList(ctx, name, namespace, path).IsHighest(isHighest).Limit(limit).Name2(name2).Namespace2(namespace2).Offset(offset).Ordering(ordering).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Tags(tags).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    path := "path_example" // string | 
    isHighest := true // bool |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name2 := "name_example" // string |  (optional)
    namespace2 := "namespace_example" // string |  (optional)
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
    resp, r, err := apiClient.PulpAnsibleApiV3CollectionsVersionsApi.PulpAnsibleGalaxyApiV3CollectionsVersionsList(context.Background(), name, namespace, path).IsHighest(isHighest).Limit(limit).Name2(name2).Namespace2(namespace2).Offset(offset).Ordering(ordering).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Tags(tags).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleApiV3CollectionsVersionsApi.PulpAnsibleGalaxyApiV3CollectionsVersionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyApiV3CollectionsVersionsList`: PaginatedCollectionVersionListResponseList
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleApiV3CollectionsVersionsApi.PulpAnsibleGalaxyApiV3CollectionsVersionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyApiV3CollectionsVersionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **isHighest** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name2** | **string** |  | 
 **namespace2** | **string** |  | 
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

[**PaginatedCollectionVersionListResponseList**](PaginatedCollectionVersionListResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PulpAnsibleGalaxyApiV3CollectionsVersionsRead

> CollectionVersionResponse PulpAnsibleGalaxyApiV3CollectionsVersionsRead(ctx, name, namespace, path, version).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    path := "path_example" // string | 
    version := "version_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleApiV3CollectionsVersionsApi.PulpAnsibleGalaxyApiV3CollectionsVersionsRead(context.Background(), name, namespace, path, version).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleApiV3CollectionsVersionsApi.PulpAnsibleGalaxyApiV3CollectionsVersionsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyApiV3CollectionsVersionsRead`: CollectionVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleApiV3CollectionsVersionsApi.PulpAnsibleGalaxyApiV3CollectionsVersionsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 
**path** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyApiV3CollectionsVersionsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**CollectionVersionResponse**](CollectionVersionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

