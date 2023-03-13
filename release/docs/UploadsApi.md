# \UploadsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadsAddRole**](UploadsApi.md#UploadsAddRole) | **Post** /{upload_href}add_role/ | 
[**UploadsCommit**](UploadsApi.md#UploadsCommit) | **Post** /{upload_href}commit/ | Finish an Upload
[**UploadsCreate**](UploadsApi.md#UploadsCreate) | **Post** /pulp/api/v3/uploads/ | Create an upload
[**UploadsDelete**](UploadsApi.md#UploadsDelete) | **Delete** /{upload_href} | Delete an upload
[**UploadsList**](UploadsApi.md#UploadsList) | **Get** /pulp/api/v3/uploads/ | List uploads
[**UploadsListRoles**](UploadsApi.md#UploadsListRoles) | **Get** /{upload_href}list_roles/ | 
[**UploadsMyPermissions**](UploadsApi.md#UploadsMyPermissions) | **Get** /{upload_href}my_permissions/ | 
[**UploadsRead**](UploadsApi.md#UploadsRead) | **Get** /{upload_href} | Inspect an upload
[**UploadsRemoveRole**](UploadsApi.md#UploadsRemoveRole) | **Post** /{upload_href}remove_role/ | 
[**UploadsUpdate**](UploadsApi.md#UploadsUpdate) | **Put** /{upload_href} | Upload a file chunk



## UploadsAddRole

> NestedRoleResponse UploadsAddRole(ctx, uploadHref).Upload(upload).Execute()





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
    uploadHref := "uploadHref_example" // string | 
    upload := *openapiclient.NewUpload(int32(123)) // Upload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.UploadsAddRole(context.Background(), uploadHref).Upload(upload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadsAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadsAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadsAddRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upload** | [**Upload**](Upload.md) |  | 

### Return type

[**NestedRoleResponse**](NestedRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadsCommit

> AsyncOperationResponse UploadsCommit(ctx, uploadHref).UploadCommit(uploadCommit).Execute()

Finish an Upload



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
    uploadHref := "uploadHref_example" // string | 
    uploadCommit := *openapiclient.NewUploadCommit("Sha256_example") // UploadCommit | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.UploadsCommit(context.Background(), uploadHref).UploadCommit(uploadCommit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadsCommit`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadsCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadsCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uploadCommit** | [**UploadCommit**](UploadCommit.md) |  | 

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


## UploadsCreate

> UploadResponse UploadsCreate(ctx).Upload(upload).Execute()

Create an upload



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
    upload := *openapiclient.NewUpload(int32(123)) // Upload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.UploadsCreate(context.Background()).Upload(upload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadsCreate`: UploadResponse
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upload** | [**Upload**](Upload.md) |  | 

### Return type

[**UploadResponse**](UploadResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadsDelete

> UploadsDelete(ctx, uploadHref).Execute()

Delete an upload



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
    uploadHref := "uploadHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UploadsApi.UploadsDelete(context.Background(), uploadHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadsDeleteRequest struct via the builder pattern


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


## UploadsList

> PaginatedUploadResponseList UploadsList(ctx).Limit(limit).Offset(offset).Ordering(ordering).Size(size).SizeGt(sizeGt).SizeLt(sizeLt).SizeRange(sizeRange).Fields(fields).ExcludeFields(excludeFields).Execute()

List uploads



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
    size := int32(56) // int32 | Filter results where size matches value (optional)
    sizeGt := int32(56) // int32 | Filter results where size is greater than value (optional)
    sizeLt := int32(56) // int32 | Filter results where size is less than value (optional)
    sizeRange := []int32{int32(123)} // []int32 | Filter results where size is between two comma separated values (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.UploadsList(context.Background()).Limit(limit).Offset(offset).Ordering(ordering).Size(size).SizeGt(sizeGt).SizeLt(sizeLt).SizeRange(sizeRange).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadsList`: PaginatedUploadResponseList
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **size** | **int32** | Filter results where size matches value | 
 **sizeGt** | **int32** | Filter results where size is greater than value | 
 **sizeLt** | **int32** | Filter results where size is less than value | 
 **sizeRange** | **[]int32** | Filter results where size is between two comma separated values | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedUploadResponseList**](PaginatedUploadResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadsListRoles

> ObjectRolesResponse UploadsListRoles(ctx, uploadHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    uploadHref := "uploadHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.UploadsListRoles(context.Background(), uploadHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadsListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadsListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadsListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ObjectRolesResponse**](ObjectRolesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadsMyPermissions

> MyPermissionsResponse UploadsMyPermissions(ctx, uploadHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    uploadHref := "uploadHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.UploadsMyPermissions(context.Background(), uploadHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadsMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadsMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadsMyPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**MyPermissionsResponse**](MyPermissionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadsRead

> UploadDetailResponse UploadsRead(ctx, uploadHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an upload



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
    uploadHref := "uploadHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.UploadsRead(context.Background(), uploadHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadsRead`: UploadDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**UploadDetailResponse**](UploadDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadsRemoveRole

> NestedRoleResponse UploadsRemoveRole(ctx, uploadHref).Upload(upload).Execute()





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
    uploadHref := "uploadHref_example" // string | 
    upload := *openapiclient.NewUpload(int32(123)) // Upload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.UploadsRemoveRole(context.Background(), uploadHref).Upload(upload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadsRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadsRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadsRemoveRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upload** | [**Upload**](Upload.md) |  | 

### Return type

[**NestedRoleResponse**](NestedRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadsUpdate

> UploadResponse UploadsUpdate(ctx, uploadHref).ContentRange(contentRange).File(file).Sha256(sha256).Execute()

Upload a file chunk



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
    contentRange := "contentRange_example" // string | The Content-Range header specifies the location of the file chunk within the file.
    uploadHref := "uploadHref_example" // string | 
    file := os.NewFile(1234, "some_file") // *os.File | A chunk of the uploaded file.
    sha256 := "sha256_example" // string | The SHA-256 checksum of the chunk if available. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadsApi.UploadsUpdate(context.Background(), uploadHref).ContentRange(contentRange).File(file).Sha256(sha256).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadsApi.UploadsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadsUpdate`: UploadResponse
    fmt.Fprintf(os.Stdout, "Response from `UploadsApi.UploadsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentRange** | **string** | The Content-Range header specifies the location of the file chunk within the file. | 

 **file** | ***os.File** | A chunk of the uploaded file. | 
 **sha256** | **string** | The SHA-256 checksum of the chunk if available. | 

### Return type

[**UploadResponse**](UploadResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

