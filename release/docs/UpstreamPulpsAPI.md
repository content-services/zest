# \UpstreamPulpsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpstreamPulpsAddRole**](UpstreamPulpsAPI.md#UpstreamPulpsAddRole) | **Post** /{upstream_pulp_href}add_role/ | Add a role
[**UpstreamPulpsCreate**](UpstreamPulpsAPI.md#UpstreamPulpsCreate) | **Post** /pulp/{pulp_domain}/api/v3/upstream-pulps/ | Create an upstream pulp
[**UpstreamPulpsDelete**](UpstreamPulpsAPI.md#UpstreamPulpsDelete) | **Delete** /{upstream_pulp_href} | Delete an upstream pulp
[**UpstreamPulpsList**](UpstreamPulpsAPI.md#UpstreamPulpsList) | **Get** /pulp/{pulp_domain}/api/v3/upstream-pulps/ | List upstream pulps
[**UpstreamPulpsListRoles**](UpstreamPulpsAPI.md#UpstreamPulpsListRoles) | **Get** /{upstream_pulp_href}list_roles/ | List roles
[**UpstreamPulpsMyPermissions**](UpstreamPulpsAPI.md#UpstreamPulpsMyPermissions) | **Get** /{upstream_pulp_href}my_permissions/ | List user permissions
[**UpstreamPulpsPartialUpdate**](UpstreamPulpsAPI.md#UpstreamPulpsPartialUpdate) | **Patch** /{upstream_pulp_href} | Update an upstream pulp
[**UpstreamPulpsRead**](UpstreamPulpsAPI.md#UpstreamPulpsRead) | **Get** /{upstream_pulp_href} | Inspect an upstream pulp
[**UpstreamPulpsRemoveRole**](UpstreamPulpsAPI.md#UpstreamPulpsRemoveRole) | **Post** /{upstream_pulp_href}remove_role/ | Remove a role
[**UpstreamPulpsReplicate**](UpstreamPulpsAPI.md#UpstreamPulpsReplicate) | **Post** /{upstream_pulp_href}replicate/ | Replicate
[**UpstreamPulpsUpdate**](UpstreamPulpsAPI.md#UpstreamPulpsUpdate) | **Put** /{upstream_pulp_href} | Update an upstream pulp



## UpstreamPulpsAddRole

> NestedRoleResponse UpstreamPulpsAddRole(ctx, upstreamPulpHref).NestedRole(nestedRole).Execute()

Add a role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2023"
)

func main() {
	upstreamPulpHref := "upstreamPulpHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsAddRole(context.Background(), upstreamPulpHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsAddRoleRequest struct via the builder pattern


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


## UpstreamPulpsCreate

> UpstreamPulpResponse UpstreamPulpsCreate(ctx, pulpDomain).UpstreamPulp(upstreamPulp).Execute()

Create an upstream pulp



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2023"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	upstreamPulp := *openapiclient.NewUpstreamPulp("Name_example", "BaseUrl_example", "ApiRoot_example") // UpstreamPulp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsCreate(context.Background(), pulpDomain).UpstreamPulp(upstreamPulp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsCreate`: UpstreamPulpResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upstreamPulp** | [**UpstreamPulp**](UpstreamPulp.md) |  | 

### Return type

[**UpstreamPulpResponse**](UpstreamPulpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsDelete

> UpstreamPulpsDelete(ctx, upstreamPulpHref).Execute()

Delete an upstream pulp



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2023"
)

func main() {
	upstreamPulpHref := "upstreamPulpHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsDelete(context.Background(), upstreamPulpHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsDeleteRequest struct via the builder pattern


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


## UpstreamPulpsList

> PaginatedUpstreamPulpResponseList UpstreamPulpsList(ctx, pulpDomain).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List upstream pulps



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2023"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsList(context.Background(), pulpDomain).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsList`: PaginatedUpstreamPulpResponseList
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedUpstreamPulpResponseList**](PaginatedUpstreamPulpResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsListRoles

> ObjectRolesResponse UpstreamPulpsListRoles(ctx, upstreamPulpHref).Fields(fields).ExcludeFields(excludeFields).Execute()

List roles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2023"
)

func main() {
	upstreamPulpHref := "upstreamPulpHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsListRoles(context.Background(), upstreamPulpHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsListRolesRequest struct via the builder pattern


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


## UpstreamPulpsMyPermissions

> MyPermissionsResponse UpstreamPulpsMyPermissions(ctx, upstreamPulpHref).Fields(fields).ExcludeFields(excludeFields).Execute()

List user permissions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2023"
)

func main() {
	upstreamPulpHref := "upstreamPulpHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsMyPermissions(context.Background(), upstreamPulpHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsMyPermissionsRequest struct via the builder pattern


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


## UpstreamPulpsPartialUpdate

> UpstreamPulpResponse UpstreamPulpsPartialUpdate(ctx, upstreamPulpHref).PatchedUpstreamPulp(patchedUpstreamPulp).Execute()

Update an upstream pulp



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2023"
)

func main() {
	upstreamPulpHref := "upstreamPulpHref_example" // string | 
	patchedUpstreamPulp := *openapiclient.NewPatchedUpstreamPulp() // PatchedUpstreamPulp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsPartialUpdate(context.Background(), upstreamPulpHref).PatchedUpstreamPulp(patchedUpstreamPulp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsPartialUpdate`: UpstreamPulpResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedUpstreamPulp** | [**PatchedUpstreamPulp**](PatchedUpstreamPulp.md) |  | 

### Return type

[**UpstreamPulpResponse**](UpstreamPulpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsRead

> UpstreamPulpResponse UpstreamPulpsRead(ctx, upstreamPulpHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an upstream pulp



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2023"
)

func main() {
	upstreamPulpHref := "upstreamPulpHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsRead(context.Background(), upstreamPulpHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsRead`: UpstreamPulpResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**UpstreamPulpResponse**](UpstreamPulpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsRemoveRole

> NestedRoleResponse UpstreamPulpsRemoveRole(ctx, upstreamPulpHref).NestedRole(nestedRole).Execute()

Remove a role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2023"
)

func main() {
	upstreamPulpHref := "upstreamPulpHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsRemoveRole(context.Background(), upstreamPulpHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsRemoveRoleRequest struct via the builder pattern


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


## UpstreamPulpsReplicate

> TaskGroupOperationResponse UpstreamPulpsReplicate(ctx, upstreamPulpHref).Execute()

Replicate



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2023"
)

func main() {
	upstreamPulpHref := "upstreamPulpHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsReplicate(context.Background(), upstreamPulpHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsReplicate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsReplicate`: TaskGroupOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsReplicate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsReplicateRequest struct via the builder pattern


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


## UpstreamPulpsUpdate

> UpstreamPulpResponse UpstreamPulpsUpdate(ctx, upstreamPulpHref).UpstreamPulp(upstreamPulp).Execute()

Update an upstream pulp



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2023"
)

func main() {
	upstreamPulpHref := "upstreamPulpHref_example" // string | 
	upstreamPulp := *openapiclient.NewUpstreamPulp("Name_example", "BaseUrl_example", "ApiRoot_example") // UpstreamPulp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsUpdate(context.Background(), upstreamPulpHref).UpstreamPulp(upstreamPulp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsUpdate`: UpstreamPulpResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upstreamPulp** | [**UpstreamPulp**](UpstreamPulp.md) |  | 

### Return type

[**UpstreamPulpResponse**](UpstreamPulpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

