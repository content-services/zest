# \RepositoriesAnsibleApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesAnsibleAnsibleCreate**](RepositoriesAnsibleApi.md#RepositoriesAnsibleAnsibleCreate) | **Post** /pulp/api/v3/repositories/ansible/ansible/ | Create an ansible repository
[**RepositoriesAnsibleAnsibleDelete**](RepositoriesAnsibleApi.md#RepositoriesAnsibleAnsibleDelete) | **Delete** /{ansible_ansible_repository_href} | Delete an ansible repository
[**RepositoriesAnsibleAnsibleList**](RepositoriesAnsibleApi.md#RepositoriesAnsibleAnsibleList) | **Get** /pulp/api/v3/repositories/ansible/ansible/ | List ansible repositorys
[**RepositoriesAnsibleAnsibleModify**](RepositoriesAnsibleApi.md#RepositoriesAnsibleAnsibleModify) | **Post** /{ansible_ansible_repository_href}modify/ | Modify Repository Content
[**RepositoriesAnsibleAnsiblePartialUpdate**](RepositoriesAnsibleApi.md#RepositoriesAnsibleAnsiblePartialUpdate) | **Patch** /{ansible_ansible_repository_href} | Update an ansible repository
[**RepositoriesAnsibleAnsibleRead**](RepositoriesAnsibleApi.md#RepositoriesAnsibleAnsibleRead) | **Get** /{ansible_ansible_repository_href} | Inspect an ansible repository
[**RepositoriesAnsibleAnsibleRebuildMetadata**](RepositoriesAnsibleApi.md#RepositoriesAnsibleAnsibleRebuildMetadata) | **Post** /{ansible_ansible_repository_href}rebuild_metadata/ | 
[**RepositoriesAnsibleAnsibleSign**](RepositoriesAnsibleApi.md#RepositoriesAnsibleAnsibleSign) | **Post** /{ansible_ansible_repository_href}sign/ | 
[**RepositoriesAnsibleAnsibleSync**](RepositoriesAnsibleApi.md#RepositoriesAnsibleAnsibleSync) | **Post** /{ansible_ansible_repository_href}sync/ | 
[**RepositoriesAnsibleAnsibleUpdate**](RepositoriesAnsibleApi.md#RepositoriesAnsibleAnsibleUpdate) | **Put** /{ansible_ansible_repository_href} | Update an ansible repository



## RepositoriesAnsibleAnsibleCreate

> AnsibleAnsibleRepositoryResponse RepositoriesAnsibleAnsibleCreate(ctx).AnsibleAnsibleRepository(ansibleAnsibleRepository).Execute()

Create an ansible repository



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
    ansibleAnsibleRepository := *openapiclient.NewAnsibleAnsibleRepository("Name_example") // AnsibleAnsibleRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleCreate(context.Background()).AnsibleAnsibleRepository(ansibleAnsibleRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleCreate`: AnsibleAnsibleRepositoryResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ansibleAnsibleRepository** | [**AnsibleAnsibleRepository**](AnsibleAnsibleRepository.md) |  | 

### Return type

[**AnsibleAnsibleRepositoryResponse**](AnsibleAnsibleRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesAnsibleAnsibleDelete

> AsyncOperationResponse RepositoriesAnsibleAnsibleDelete(ctx, ansibleAnsibleRepositoryHref).Execute()

Delete an ansible repository



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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleDelete(context.Background(), ansibleAnsibleRepositoryHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleDeleteRequest struct via the builder pattern


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


## RepositoriesAnsibleAnsibleList

> PaginatedansibleAnsibleRepositoryResponseList RepositoriesAnsibleAnsibleList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).Remote(remote).Fields(fields).ExcludeFields(excludeFields).Execute()

List ansible repositorys



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
    pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
    remote := "remote_example" // string | Foreign Key referenced by HREF (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).Remote(remote).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleList`: PaginatedansibleAnsibleRepositoryResponseList
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleListRequest struct via the builder pattern


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
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **remote** | **string** | Foreign Key referenced by HREF | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedansibleAnsibleRepositoryResponseList**](PaginatedansibleAnsibleRepositoryResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesAnsibleAnsibleModify

> AsyncOperationResponse RepositoriesAnsibleAnsibleModify(ctx, ansibleAnsibleRepositoryHref).RepositoryAddRemoveContent(repositoryAddRemoveContent).Execute()

Modify Repository Content



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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    repositoryAddRemoveContent := *openapiclient.NewRepositoryAddRemoveContent() // RepositoryAddRemoveContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleModify(context.Background(), ansibleAnsibleRepositoryHref).RepositoryAddRemoveContent(repositoryAddRemoveContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleModify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleModify`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositoryAddRemoveContent** | [**RepositoryAddRemoveContent**](RepositoryAddRemoveContent.md) |  | 

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


## RepositoriesAnsibleAnsiblePartialUpdate

> AsyncOperationResponse RepositoriesAnsibleAnsiblePartialUpdate(ctx, ansibleAnsibleRepositoryHref).PatchedansibleAnsibleRepository(patchedansibleAnsibleRepository).Execute()

Update an ansible repository



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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    patchedansibleAnsibleRepository := *openapiclient.NewPatchedansibleAnsibleRepository() // PatchedansibleAnsibleRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleApi.RepositoriesAnsibleAnsiblePartialUpdate(context.Background(), ansibleAnsibleRepositoryHref).PatchedansibleAnsibleRepository(patchedansibleAnsibleRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleApi.RepositoriesAnsibleAnsiblePartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsiblePartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleApi.RepositoriesAnsibleAnsiblePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsiblePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedansibleAnsibleRepository** | [**PatchedansibleAnsibleRepository**](PatchedansibleAnsibleRepository.md) |  | 

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


## RepositoriesAnsibleAnsibleRead

> AnsibleAnsibleRepositoryResponse RepositoriesAnsibleAnsibleRead(ctx, ansibleAnsibleRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an ansible repository



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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleRead(context.Background(), ansibleAnsibleRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleRead`: AnsibleAnsibleRepositoryResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**AnsibleAnsibleRepositoryResponse**](AnsibleAnsibleRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesAnsibleAnsibleRebuildMetadata

> AsyncOperationResponse RepositoriesAnsibleAnsibleRebuildMetadata(ctx, ansibleAnsibleRepositoryHref).AnsibleRepositoryRebuild(ansibleRepositoryRebuild).Execute()





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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    ansibleRepositoryRebuild := *openapiclient.NewAnsibleRepositoryRebuild() // AnsibleRepositoryRebuild | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleRebuildMetadata(context.Background(), ansibleAnsibleRepositoryHref).AnsibleRepositoryRebuild(ansibleRepositoryRebuild).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleRebuildMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleRebuildMetadata`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleRebuildMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleRebuildMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ansibleRepositoryRebuild** | [**AnsibleRepositoryRebuild**](AnsibleRepositoryRebuild.md) |  | 

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


## RepositoriesAnsibleAnsibleSign

> AsyncOperationResponse RepositoriesAnsibleAnsibleSign(ctx, ansibleAnsibleRepositoryHref).AnsibleRepositorySignature(ansibleRepositorySignature).Execute()





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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    ansibleRepositorySignature := *openapiclient.NewAnsibleRepositorySignature([]interface{}{nil}, "SigningService_example") // AnsibleRepositorySignature | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleSign(context.Background(), ansibleAnsibleRepositoryHref).AnsibleRepositorySignature(ansibleRepositorySignature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleSign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleSign`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleSign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleSignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ansibleRepositorySignature** | [**AnsibleRepositorySignature**](AnsibleRepositorySignature.md) |  | 

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


## RepositoriesAnsibleAnsibleSync

> AsyncOperationResponse RepositoriesAnsibleAnsibleSync(ctx, ansibleAnsibleRepositoryHref).AnsibleRepositorySyncURL(ansibleRepositorySyncURL).Execute()





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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    ansibleRepositorySyncURL := *openapiclient.NewAnsibleRepositorySyncURL() // AnsibleRepositorySyncURL | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleSync(context.Background(), ansibleAnsibleRepositoryHref).AnsibleRepositorySyncURL(ansibleRepositorySyncURL).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleSync`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ansibleRepositorySyncURL** | [**AnsibleRepositorySyncURL**](AnsibleRepositorySyncURL.md) |  | 

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


## RepositoriesAnsibleAnsibleUpdate

> AsyncOperationResponse RepositoriesAnsibleAnsibleUpdate(ctx, ansibleAnsibleRepositoryHref).AnsibleAnsibleRepository(ansibleAnsibleRepository).Execute()

Update an ansible repository



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
    ansibleAnsibleRepositoryHref := "ansibleAnsibleRepositoryHref_example" // string | 
    ansibleAnsibleRepository := *openapiclient.NewAnsibleAnsibleRepository("Name_example") // AnsibleAnsibleRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleUpdate(context.Background(), ansibleAnsibleRepositoryHref).AnsibleAnsibleRepository(ansibleAnsibleRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesAnsibleAnsibleUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesAnsibleApi.RepositoriesAnsibleAnsibleUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleAnsibleRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesAnsibleAnsibleUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ansibleAnsibleRepository** | [**AnsibleAnsibleRepository**](AnsibleAnsibleRepository.md) |  | 

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

