# \GroupsRolesApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsRolesCreate**](GroupsRolesApi.md#GroupsRolesCreate) | **Post** /{group_href}roles/ | Create a group role
[**GroupsRolesDelete**](GroupsRolesApi.md#GroupsRolesDelete) | **Delete** /{groups_group_role_href} | Delete a group role
[**GroupsRolesList**](GroupsRolesApi.md#GroupsRolesList) | **Get** /{group_href}roles/ | List group roles
[**GroupsRolesRead**](GroupsRolesApi.md#GroupsRolesRead) | **Get** /{groups_group_role_href} | Inspect a group role



## GroupsRolesCreate

> GroupRoleResponse GroupsRolesCreate(ctx, groupHref).GroupRole(groupRole).Execute()

Create a group role



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
    groupHref := "groupHref_example" // string | 
    groupRole := *openapiclient.NewGroupRole("Role_example", "ContentObject_example") // GroupRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsRolesApi.GroupsRolesCreate(context.Background(), groupHref).GroupRole(groupRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsRolesApi.GroupsRolesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsRolesCreate`: GroupRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsRolesApi.GroupsRolesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsRolesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupRole** | [**GroupRole**](GroupRole.md) |  | 

### Return type

[**GroupRoleResponse**](GroupRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsRolesDelete

> GroupsRolesDelete(ctx, groupsGroupRoleHref).Execute()

Delete a group role



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
    groupsGroupRoleHref := "groupsGroupRoleHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GroupsRolesApi.GroupsRolesDelete(context.Background(), groupsGroupRoleHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsRolesApi.GroupsRolesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupsGroupRoleHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsRolesDeleteRequest struct via the builder pattern


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


## GroupsRolesList

> PaginatedGroupRoleResponseList GroupsRolesList(ctx, groupHref).ContentObject(contentObject).Limit(limit).Offset(offset).Ordering(ordering).Role(role).RoleContains(roleContains).RoleIcontains(roleIcontains).RoleIn(roleIn).RoleStartswith(roleStartswith).Fields(fields).ExcludeFields(excludeFields).Execute()

List group roles



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
    groupHref := "groupHref_example" // string | 
    contentObject := "contentObject_example" // string | content_object (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    role := "role_example" // string |  (optional)
    roleContains := "roleContains_example" // string |  (optional)
    roleIcontains := "roleIcontains_example" // string |  (optional)
    roleIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    roleStartswith := "roleStartswith_example" // string |  (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsRolesApi.GroupsRolesList(context.Background(), groupHref).ContentObject(contentObject).Limit(limit).Offset(offset).Ordering(ordering).Role(role).RoleContains(roleContains).RoleIcontains(roleIcontains).RoleIn(roleIn).RoleStartswith(roleStartswith).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsRolesApi.GroupsRolesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsRolesList`: PaginatedGroupRoleResponseList
    fmt.Fprintf(os.Stdout, "Response from `GroupsRolesApi.GroupsRolesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsRolesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentObject** | **string** | content_object | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **role** | **string** |  | 
 **roleContains** | **string** |  | 
 **roleIcontains** | **string** |  | 
 **roleIn** | **[]string** | Multiple values may be separated by commas. | 
 **roleStartswith** | **string** |  | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedGroupRoleResponseList**](PaginatedGroupRoleResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsRolesRead

> GroupRoleResponse GroupsRolesRead(ctx, groupsGroupRoleHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a group role



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
    groupsGroupRoleHref := "groupsGroupRoleHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsRolesApi.GroupsRolesRead(context.Background(), groupsGroupRoleHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsRolesApi.GroupsRolesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsRolesRead`: GroupRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsRolesApi.GroupsRolesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupsGroupRoleHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsRolesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**GroupRoleResponse**](GroupRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

