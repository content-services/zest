# \UsersApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreate**](UsersApi.md#UsersCreate) | **Post** /pulp/api/v3/users/ | Create an user
[**UsersDelete**](UsersApi.md#UsersDelete) | **Delete** /{auth_user_href} | Delete an user
[**UsersList**](UsersApi.md#UsersList) | **Get** /pulp/api/v3/users/ | List users
[**UsersPartialUpdate**](UsersApi.md#UsersPartialUpdate) | **Patch** /{auth_user_href} | Update an user
[**UsersRead**](UsersApi.md#UsersRead) | **Get** /{auth_user_href} | Inspect an user
[**UsersUpdate**](UsersApi.md#UsersUpdate) | **Put** /{auth_user_href} | Update an user



## UsersCreate

> UserResponse UsersCreate(ctx).User(user).Execute()

Create an user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func main() {
    user := *openapiclient.NewUser("Username_example") // User | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersCreate(context.Background()).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreate`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md) |  | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersDelete

> UsersDelete(ctx, authUserHref).Execute()

Delete an user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func main() {
    authUserHref := "authUserHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersApi.UsersDelete(context.Background(), authUserHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authUserHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteRequest struct via the builder pattern


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


## UsersList

> PaginatedUserResponseList UsersList(ctx).Email(email).EmailContains(emailContains).EmailIcontains(emailIcontains).EmailIexact(emailIexact).EmailIn(emailIn).FirstName(firstName).FirstNameContains(firstNameContains).FirstNameIcontains(firstNameIcontains).FirstNameIexact(firstNameIexact).FirstNameIn(firstNameIn).IsActive(isActive).IsStaff(isStaff).LastName(lastName).LastNameContains(lastNameContains).LastNameIcontains(lastNameIcontains).LastNameIexact(lastNameIexact).LastNameIn(lastNameIn).Limit(limit).Offset(offset).Ordering(ordering).Username(username).UsernameContains(usernameContains).UsernameIcontains(usernameIcontains).UsernameIexact(usernameIexact).UsernameIn(usernameIn).Fields(fields).ExcludeFields(excludeFields).Execute()

List users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func main() {
    email := "email_example" // string | Filter results where email matches value (optional)
    emailContains := "emailContains_example" // string | Filter results where email contains value (optional)
    emailIcontains := "emailIcontains_example" // string | Filter results where email contains value (optional)
    emailIexact := "emailIexact_example" // string | Filter results where email matches value (optional)
    emailIn := []string{"Inner_example"} // []string | Filter results where email is in a comma-separated list of values (optional)
    firstName := "firstName_example" // string | Filter results where first_name matches value (optional)
    firstNameContains := "firstNameContains_example" // string | Filter results where first_name contains value (optional)
    firstNameIcontains := "firstNameIcontains_example" // string | Filter results where first_name contains value (optional)
    firstNameIexact := "firstNameIexact_example" // string | Filter results where first_name matches value (optional)
    firstNameIn := []string{"Inner_example"} // []string | Filter results where first_name is in a comma-separated list of values (optional)
    isActive := true // bool | Filter results where is_active matches value (optional)
    isStaff := true // bool | Filter results where is_staff matches value (optional)
    lastName := "lastName_example" // string | Filter results where last_name matches value (optional)
    lastNameContains := "lastNameContains_example" // string | Filter results where last_name contains value (optional)
    lastNameIcontains := "lastNameIcontains_example" // string | Filter results where last_name contains value (optional)
    lastNameIexact := "lastNameIexact_example" // string | Filter results where last_name matches value (optional)
    lastNameIn := []string{"Inner_example"} // []string | Filter results where last_name is in a comma-separated list of values (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    username := "username_example" // string | Filter results where username matches value (optional)
    usernameContains := "usernameContains_example" // string | Filter results where username contains value (optional)
    usernameIcontains := "usernameIcontains_example" // string | Filter results where username contains value (optional)
    usernameIexact := "usernameIexact_example" // string | Filter results where username matches value (optional)
    usernameIn := []string{"Inner_example"} // []string | Filter results where username is in a comma-separated list of values (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersList(context.Background()).Email(email).EmailContains(emailContains).EmailIcontains(emailIcontains).EmailIexact(emailIexact).EmailIn(emailIn).FirstName(firstName).FirstNameContains(firstNameContains).FirstNameIcontains(firstNameIcontains).FirstNameIexact(firstNameIexact).FirstNameIn(firstNameIn).IsActive(isActive).IsStaff(isStaff).LastName(lastName).LastNameContains(lastNameContains).LastNameIcontains(lastNameIcontains).LastNameIexact(lastNameIexact).LastNameIn(lastNameIn).Limit(limit).Offset(offset).Ordering(ordering).Username(username).UsernameContains(usernameContains).UsernameIcontains(usernameIcontains).UsernameIexact(usernameIexact).UsernameIn(usernameIn).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersList`: PaginatedUserResponseList
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Filter results where email matches value | 
 **emailContains** | **string** | Filter results where email contains value | 
 **emailIcontains** | **string** | Filter results where email contains value | 
 **emailIexact** | **string** | Filter results where email matches value | 
 **emailIn** | **[]string** | Filter results where email is in a comma-separated list of values | 
 **firstName** | **string** | Filter results where first_name matches value | 
 **firstNameContains** | **string** | Filter results where first_name contains value | 
 **firstNameIcontains** | **string** | Filter results where first_name contains value | 
 **firstNameIexact** | **string** | Filter results where first_name matches value | 
 **firstNameIn** | **[]string** | Filter results where first_name is in a comma-separated list of values | 
 **isActive** | **bool** | Filter results where is_active matches value | 
 **isStaff** | **bool** | Filter results where is_staff matches value | 
 **lastName** | **string** | Filter results where last_name matches value | 
 **lastNameContains** | **string** | Filter results where last_name contains value | 
 **lastNameIcontains** | **string** | Filter results where last_name contains value | 
 **lastNameIexact** | **string** | Filter results where last_name matches value | 
 **lastNameIn** | **[]string** | Filter results where last_name is in a comma-separated list of values | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **username** | **string** | Filter results where username matches value | 
 **usernameContains** | **string** | Filter results where username contains value | 
 **usernameIcontains** | **string** | Filter results where username contains value | 
 **usernameIexact** | **string** | Filter results where username matches value | 
 **usernameIn** | **[]string** | Filter results where username is in a comma-separated list of values | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedUserResponseList**](PaginatedUserResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPartialUpdate

> UserResponse UsersPartialUpdate(ctx, authUserHref).PatchedUser(patchedUser).Execute()

Update an user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func main() {
    authUserHref := "authUserHref_example" // string | 
    patchedUser := *openapiclient.NewPatchedUser() // PatchedUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPartialUpdate(context.Background(), authUserHref).PatchedUser(patchedUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPartialUpdate`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authUserHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedUser** | [**PatchedUser**](PatchedUser.md) |  | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersRead

> UserResponse UsersRead(ctx, authUserHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func main() {
    authUserHref := "authUserHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersRead(context.Background(), authUserHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersRead`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authUserHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdate

> UserResponse UsersUpdate(ctx, authUserHref).User(user).Execute()

Update an user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func main() {
    authUserHref := "authUserHref_example" // string | 
    user := *openapiclient.NewUser("Username_example") // User | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUpdate(context.Background(), authUserHref).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUpdate`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authUserHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

