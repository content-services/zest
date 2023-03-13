# \AcsRpmApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcsRpmRpmAddRole**](AcsRpmApi.md#AcsRpmRpmAddRole) | **Post** /{rpm_rpm_alternate_content_source_href}add_role/ | 
[**AcsRpmRpmCreate**](AcsRpmApi.md#AcsRpmRpmCreate) | **Post** /pulp/api/v3/acs/rpm/rpm/ | Create a rpm alternate content source
[**AcsRpmRpmDelete**](AcsRpmApi.md#AcsRpmRpmDelete) | **Delete** /{rpm_rpm_alternate_content_source_href} | Delete a rpm alternate content source
[**AcsRpmRpmList**](AcsRpmApi.md#AcsRpmRpmList) | **Get** /pulp/api/v3/acs/rpm/rpm/ | List rpm alternate content sources
[**AcsRpmRpmListRoles**](AcsRpmApi.md#AcsRpmRpmListRoles) | **Get** /{rpm_rpm_alternate_content_source_href}list_roles/ | 
[**AcsRpmRpmMyPermissions**](AcsRpmApi.md#AcsRpmRpmMyPermissions) | **Get** /{rpm_rpm_alternate_content_source_href}my_permissions/ | 
[**AcsRpmRpmPartialUpdate**](AcsRpmApi.md#AcsRpmRpmPartialUpdate) | **Patch** /{rpm_rpm_alternate_content_source_href} | Update a rpm alternate content source
[**AcsRpmRpmRead**](AcsRpmApi.md#AcsRpmRpmRead) | **Get** /{rpm_rpm_alternate_content_source_href} | Inspect a rpm alternate content source
[**AcsRpmRpmRefresh**](AcsRpmApi.md#AcsRpmRpmRefresh) | **Post** /{rpm_rpm_alternate_content_source_href}refresh/ | 
[**AcsRpmRpmRemoveRole**](AcsRpmApi.md#AcsRpmRpmRemoveRole) | **Post** /{rpm_rpm_alternate_content_source_href}remove_role/ | 
[**AcsRpmRpmUpdate**](AcsRpmApi.md#AcsRpmRpmUpdate) | **Put** /{rpm_rpm_alternate_content_source_href} | Update a rpm alternate content source



## AcsRpmRpmAddRole

> NestedRoleResponse AcsRpmRpmAddRole(ctx, rpmRpmAlternateContentSourceHref).NestedRole(nestedRole).Execute()





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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmApi.AcsRpmRpmAddRole(context.Background(), rpmRpmAlternateContentSourceHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmApi.AcsRpmRpmAddRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmAddRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmApi.AcsRpmRpmAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmAddRoleRequest struct via the builder pattern


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


## AcsRpmRpmCreate

> RpmRpmAlternateContentSourceResponse AcsRpmRpmCreate(ctx).RpmRpmAlternateContentSource(rpmRpmAlternateContentSource).Execute()

Create a rpm alternate content source



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
    rpmRpmAlternateContentSource := *openapiclient.NewRpmRpmAlternateContentSource("Name_example", "Remote_example") // RpmRpmAlternateContentSource | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmApi.AcsRpmRpmCreate(context.Background()).RpmRpmAlternateContentSource(rpmRpmAlternateContentSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmApi.AcsRpmRpmCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmCreate`: RpmRpmAlternateContentSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmApi.AcsRpmRpmCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rpmRpmAlternateContentSource** | [**RpmRpmAlternateContentSource**](RpmRpmAlternateContentSource.md) |  | 

### Return type

[**RpmRpmAlternateContentSourceResponse**](RpmRpmAlternateContentSourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcsRpmRpmDelete

> AsyncOperationResponse AcsRpmRpmDelete(ctx, rpmRpmAlternateContentSourceHref).Execute()

Delete a rpm alternate content source



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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmApi.AcsRpmRpmDelete(context.Background(), rpmRpmAlternateContentSourceHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmApi.AcsRpmRpmDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmApi.AcsRpmRpmDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmDeleteRequest struct via the builder pattern


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


## AcsRpmRpmList

> PaginatedrpmRpmAlternateContentSourceResponseList AcsRpmRpmList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()

List rpm alternate content sources



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
    resp, r, err := apiClient.AcsRpmApi.AcsRpmRpmList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmApi.AcsRpmRpmList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmList`: PaginatedrpmRpmAlternateContentSourceResponseList
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmApi.AcsRpmRpmList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmListRequest struct via the builder pattern


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

[**PaginatedrpmRpmAlternateContentSourceResponseList**](PaginatedrpmRpmAlternateContentSourceResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcsRpmRpmListRoles

> ObjectRolesResponse AcsRpmRpmListRoles(ctx, rpmRpmAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmApi.AcsRpmRpmListRoles(context.Background(), rpmRpmAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmApi.AcsRpmRpmListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmListRoles`: ObjectRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmApi.AcsRpmRpmListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmListRolesRequest struct via the builder pattern


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


## AcsRpmRpmMyPermissions

> MyPermissionsResponse AcsRpmRpmMyPermissions(ctx, rpmRpmAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmApi.AcsRpmRpmMyPermissions(context.Background(), rpmRpmAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmApi.AcsRpmRpmMyPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmMyPermissions`: MyPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmApi.AcsRpmRpmMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmMyPermissionsRequest struct via the builder pattern


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


## AcsRpmRpmPartialUpdate

> AsyncOperationResponse AcsRpmRpmPartialUpdate(ctx, rpmRpmAlternateContentSourceHref).PatchedrpmRpmAlternateContentSource(patchedrpmRpmAlternateContentSource).Execute()

Update a rpm alternate content source



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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 
    patchedrpmRpmAlternateContentSource := *openapiclient.NewPatchedrpmRpmAlternateContentSource() // PatchedrpmRpmAlternateContentSource | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmApi.AcsRpmRpmPartialUpdate(context.Background(), rpmRpmAlternateContentSourceHref).PatchedrpmRpmAlternateContentSource(patchedrpmRpmAlternateContentSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmApi.AcsRpmRpmPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmApi.AcsRpmRpmPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedrpmRpmAlternateContentSource** | [**PatchedrpmRpmAlternateContentSource**](PatchedrpmRpmAlternateContentSource.md) |  | 

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


## AcsRpmRpmRead

> RpmRpmAlternateContentSourceResponse AcsRpmRpmRead(ctx, rpmRpmAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a rpm alternate content source



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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmApi.AcsRpmRpmRead(context.Background(), rpmRpmAlternateContentSourceHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmApi.AcsRpmRpmRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmRead`: RpmRpmAlternateContentSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmApi.AcsRpmRpmRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RpmRpmAlternateContentSourceResponse**](RpmRpmAlternateContentSourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcsRpmRpmRefresh

> TaskGroupOperationResponse AcsRpmRpmRefresh(ctx, rpmRpmAlternateContentSourceHref).Execute()





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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmApi.AcsRpmRpmRefresh(context.Background(), rpmRpmAlternateContentSourceHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmApi.AcsRpmRpmRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmRefresh`: TaskGroupOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmApi.AcsRpmRpmRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskGroupOperationResponse**](TaskGroupOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcsRpmRpmRemoveRole

> NestedRoleResponse AcsRpmRpmRemoveRole(ctx, rpmRpmAlternateContentSourceHref).NestedRole(nestedRole).Execute()





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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 
    nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmApi.AcsRpmRpmRemoveRole(context.Background(), rpmRpmAlternateContentSourceHref).NestedRole(nestedRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmApi.AcsRpmRpmRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmRemoveRole`: NestedRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmApi.AcsRpmRpmRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmRemoveRoleRequest struct via the builder pattern


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


## AcsRpmRpmUpdate

> AsyncOperationResponse AcsRpmRpmUpdate(ctx, rpmRpmAlternateContentSourceHref).RpmRpmAlternateContentSource(rpmRpmAlternateContentSource).Execute()

Update a rpm alternate content source



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
    rpmRpmAlternateContentSourceHref := "rpmRpmAlternateContentSourceHref_example" // string | 
    rpmRpmAlternateContentSource := *openapiclient.NewRpmRpmAlternateContentSource("Name_example", "Remote_example") // RpmRpmAlternateContentSource | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcsRpmApi.AcsRpmRpmUpdate(context.Background(), rpmRpmAlternateContentSourceHref).RpmRpmAlternateContentSource(rpmRpmAlternateContentSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsRpmApi.AcsRpmRpmUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcsRpmRpmUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `AcsRpmApi.AcsRpmRpmUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmAlternateContentSourceHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcsRpmRpmUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rpmRpmAlternateContentSource** | [**RpmRpmAlternateContentSource**](RpmRpmAlternateContentSource.md) |  | 

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

