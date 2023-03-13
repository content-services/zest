# \ContentguardsContentRedirectApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentguardsCoreContentRedirectAddRole**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectAddRole) | **Post** /{content_redirect_content_guard_href}add_role/ | 
[**ContentguardsCoreContentRedirectCreate**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectCreate) | **Post** /pulp/api/v3/contentguards/core/content_redirect/ | Create a content redirect content guard
[**ContentguardsCoreContentRedirectDelete**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectDelete) | **Delete** /{content_redirect_content_guard_href} | Delete a content redirect content guard
[**ContentguardsCoreContentRedirectList**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectList) | **Get** /pulp/api/v3/contentguards/core/content_redirect/ | List content redirect content guards
[**ContentguardsCoreContentRedirectListRoles**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectListRoles) | **Get** /{content_redirect_content_guard_href}list_roles/ | 
[**ContentguardsCoreContentRedirectMyPermissions**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectMyPermissions) | **Get** /{content_redirect_content_guard_href}my_permissions/ | 
[**ContentguardsCoreContentRedirectPartialUpdate**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectPartialUpdate) | **Patch** /{content_redirect_content_guard_href} | Update a content redirect content guard
[**ContentguardsCoreContentRedirectRead**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectRead) | **Get** /{content_redirect_content_guard_href} | Inspect a content redirect content guard
[**ContentguardsCoreContentRedirectRemoveRole**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectRemoveRole) | **Post** /{content_redirect_content_guard_href}remove_role/ | 
[**ContentguardsCoreContentRedirectUpdate**](ContentguardsContentRedirectApi.md#ContentguardsCoreContentRedirectUpdate) | **Put** /{content_redirect_content_guard_href} | Update a content redirect content guard



## ContentguardsCoreContentRedirectAddRole

> NestedRoleResponse ContentguardsCoreContentRedirectAddRole(ctx, contentRedirectContentGuardHref).NestedRole(nestedRole).Execute()





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
    contentRedirectContentGuardHref := "contentRedirectContentGuardHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectAddRole(context.Background(), contentRedirectContentGuardHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentRedirectContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectAddRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nestedRole** | [**NestedRole**](NestedRole.md) |  | 

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


## ContentguardsCoreContentRedirectCreate

> ContentRedirectContentGuardResponse ContentguardsCoreContentRedirectCreate(ctx).ContentRedirectContentGuard(contentRedirectContentGuard).Execute()

Create a content redirect content guard



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
    contentRedirectContentGuard := *openapiclient.NewContentRedirectContentGuard("Name_example") // ContentRedirectContentGuard | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectCreate(context.Background()).ContentRedirectContentGuard(contentRedirectContentGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectCreate`: ContentRedirectContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentRedirectContentGuard** | [**ContentRedirectContentGuard**](ContentRedirectContentGuard.md) |  | 

### Return type

[**ContentRedirectContentGuardResponse**](ContentRedirectContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreContentRedirectDelete

> ContentguardsCoreContentRedirectDelete(ctx, contentRedirectContentGuardHref).Execute()

Delete a content redirect content guard



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
    contentRedirectContentGuardHref := "contentRedirectContentGuardHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectDelete(context.Background(), contentRedirectContentGuardHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentRedirectContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectDeleteRequest struct via the builder pattern


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


## ContentguardsCoreContentRedirectList

> PaginatedContentRedirectContentGuardResponseList ContentguardsCoreContentRedirectList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()

List content redirect content guards



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
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectList`: PaginatedContentRedirectContentGuardResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectListRequest struct via the builder pattern


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
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedContentRedirectContentGuardResponseList**](PaginatedContentRedirectContentGuardResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreContentRedirectListRoles

> ObjectRolesResponse ContentguardsCoreContentRedirectListRoles(ctx, contentRedirectContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    contentRedirectContentGuardHref := "contentRedirectContentGuardHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectListRoles(context.Background(), contentRedirectContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentRedirectContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectListRolesRequest struct via the builder pattern


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


## ContentguardsCoreContentRedirectMyPermissions

> MyPermissionsResponse ContentguardsCoreContentRedirectMyPermissions(ctx, contentRedirectContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    contentRedirectContentGuardHref := "contentRedirectContentGuardHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectMyPermissions(context.Background(), contentRedirectContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentRedirectContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectMyPermissionsRequest struct via the builder pattern


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


## ContentguardsCoreContentRedirectPartialUpdate

> ContentRedirectContentGuardResponse ContentguardsCoreContentRedirectPartialUpdate(ctx, contentRedirectContentGuardHref).PatchedContentRedirectContentGuard(patchedContentRedirectContentGuard).Execute()

Update a content redirect content guard



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
    contentRedirectContentGuardHref := "contentRedirectContentGuardHref_example" // string | 
    patchedContentRedirectContentGuard := *openapiclient.NewPatchedContentRedirectContentGuard() // PatchedContentRedirectContentGuard | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectPartialUpdate(context.Background(), contentRedirectContentGuardHref).PatchedContentRedirectContentGuard(patchedContentRedirectContentGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectPartialUpdate`: ContentRedirectContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentRedirectContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedContentRedirectContentGuard** | [**PatchedContentRedirectContentGuard**](PatchedContentRedirectContentGuard.md) |  | 

### Return type

[**ContentRedirectContentGuardResponse**](ContentRedirectContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreContentRedirectRead

> ContentRedirectContentGuardResponse ContentguardsCoreContentRedirectRead(ctx, contentRedirectContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a content redirect content guard



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
    contentRedirectContentGuardHref := "contentRedirectContentGuardHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectRead(context.Background(), contentRedirectContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectRead`: ContentRedirectContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentRedirectContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ContentRedirectContentGuardResponse**](ContentRedirectContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreContentRedirectRemoveRole

> NestedRoleResponse ContentguardsCoreContentRedirectRemoveRole(ctx, contentRedirectContentGuardHref).NestedRole(nestedRole).Execute()





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
    contentRedirectContentGuardHref := "contentRedirectContentGuardHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectRemoveRole(context.Background(), contentRedirectContentGuardHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentRedirectContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectRemoveRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nestedRole** | [**NestedRole**](NestedRole.md) |  | 

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


## ContentguardsCoreContentRedirectUpdate

> ContentRedirectContentGuardResponse ContentguardsCoreContentRedirectUpdate(ctx, contentRedirectContentGuardHref).ContentRedirectContentGuard(contentRedirectContentGuard).Execute()

Update a content redirect content guard



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
    contentRedirectContentGuardHref := "contentRedirectContentGuardHref_example" // string | 
    contentRedirectContentGuard := *openapiclient.NewContentRedirectContentGuard("Name_example") // ContentRedirectContentGuard | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectUpdate(context.Background(), contentRedirectContentGuardHref).ContentRedirectContentGuard(contentRedirectContentGuard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentguardsCoreContentRedirectUpdate`: ContentRedirectContentGuardResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentRedirectContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreContentRedirectUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentRedirectContentGuard** | [**ContentRedirectContentGuard**](ContentRedirectContentGuard.md) |  | 

### Return type

[**ContentRedirectContentGuardResponse**](ContentRedirectContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

