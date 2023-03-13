# \PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsDelete**](PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApi.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsDelete) | **Delete** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/versions/{version}/ | 
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsList**](PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApi.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsList) | **Get** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/versions/ | 
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsRead**](PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApi.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsRead) | **Get** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/versions/{version}/ | 



## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsDelete

> AsyncOperationResponse PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsDelete(ctx, distroBasePath, name, namespace, version).Execute()





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
    distroBasePath := "distroBasePath_example" // string | 
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsDelete(context.Background(), distroBasePath, name, namespace, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**name** | **string** |  | 
**namespace** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsDeleteRequest struct via the builder pattern


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


## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsList

> PaginatedCollectionVersionListResponseList PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsList(ctx, distroBasePath, name, namespace).IsHighest(isHighest).Limit(limit).Name2(name2).Namespace2(namespace2).Offset(offset).Ordering(ordering).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Tags(tags).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    distroBasePath := "distroBasePath_example" // string | 
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
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
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsList(context.Background(), distroBasePath, name, namespace).IsHighest(isHighest).Limit(limit).Name2(name2).Namespace2(namespace2).Offset(offset).Ordering(ordering).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Tags(tags).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsList`: PaginatedCollectionVersionListResponseList
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsListRequest struct via the builder pattern


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


## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsRead

> CollectionVersionResponse PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsRead(ctx, distroBasePath, name, namespace, version).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    distroBasePath := "distroBasePath_example" // string | 
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    version := "version_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsRead(context.Background(), distroBasePath, name, namespace, version).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsRead`: CollectionVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**name** | **string** |  | 
**namespace** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsReadRequest struct via the builder pattern


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

