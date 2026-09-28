# \ContentguardsEnvvarHeaderAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentguardsCoreEnvvarHeaderAddRole**](ContentguardsEnvvarHeaderAPI.md#ContentguardsCoreEnvvarHeaderAddRole) | **Post** /{env_var_header_content_guard_href}add_role/ | Add a role
[**ContentguardsCoreEnvvarHeaderCreate**](ContentguardsEnvvarHeaderAPI.md#ContentguardsCoreEnvvarHeaderCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/contentguards/core/envvar_header/ | Create an env var header content guard
[**ContentguardsCoreEnvvarHeaderDelete**](ContentguardsEnvvarHeaderAPI.md#ContentguardsCoreEnvvarHeaderDelete) | **Delete** /{env_var_header_content_guard_href} | Delete an env var header content guard
[**ContentguardsCoreEnvvarHeaderList**](ContentguardsEnvvarHeaderAPI.md#ContentguardsCoreEnvvarHeaderList) | **Get** /api/pulp/{pulp_domain}/api/v3/contentguards/core/envvar_header/ | List env var header content guards
[**ContentguardsCoreEnvvarHeaderListRoles**](ContentguardsEnvvarHeaderAPI.md#ContentguardsCoreEnvvarHeaderListRoles) | **Get** /{env_var_header_content_guard_href}list_roles/ | List roles
[**ContentguardsCoreEnvvarHeaderMyPermissions**](ContentguardsEnvvarHeaderAPI.md#ContentguardsCoreEnvvarHeaderMyPermissions) | **Get** /{env_var_header_content_guard_href}my_permissions/ | List user permissions
[**ContentguardsCoreEnvvarHeaderPartialUpdate**](ContentguardsEnvvarHeaderAPI.md#ContentguardsCoreEnvvarHeaderPartialUpdate) | **Patch** /{env_var_header_content_guard_href} | Update an env var header content guard
[**ContentguardsCoreEnvvarHeaderRead**](ContentguardsEnvvarHeaderAPI.md#ContentguardsCoreEnvvarHeaderRead) | **Get** /{env_var_header_content_guard_href} | Inspect an env var header content guard
[**ContentguardsCoreEnvvarHeaderRemoveRole**](ContentguardsEnvvarHeaderAPI.md#ContentguardsCoreEnvvarHeaderRemoveRole) | **Post** /{env_var_header_content_guard_href}remove_role/ | Remove a role
[**ContentguardsCoreEnvvarHeaderUpdate**](ContentguardsEnvvarHeaderAPI.md#ContentguardsCoreEnvvarHeaderUpdate) | **Put** /{env_var_header_content_guard_href} | Update an env var header content guard



## ContentguardsCoreEnvvarHeaderAddRole

> NestedRoleResponse ContentguardsCoreEnvvarHeaderAddRole(ctx, envVarHeaderContentGuardHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()

Add a role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	envVarHeaderContentGuardHref := "envVarHeaderContentGuardHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderAddRole(context.Background(), envVarHeaderContentGuardHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreEnvvarHeaderAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envVarHeaderContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreEnvvarHeaderAddRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nestedRole** | [**NestedRole**](NestedRole.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## ContentguardsCoreEnvvarHeaderCreate

> EnvVarHeaderContentGuardResponse ContentguardsCoreEnvvarHeaderCreate(ctx, pulpDomain).EnvVarHeaderContentGuard(envVarHeaderContentGuard).XTaskDiagnostics(xTaskDiagnostics).Execute()

Create an env var header content guard



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	envVarHeaderContentGuard := *openapiclient.NewEnvVarHeaderContentGuard("Name_example", "HeaderName_example", "EnvVar_example") // EnvVarHeaderContentGuard | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderCreate(context.Background(), pulpDomain).EnvVarHeaderContentGuard(envVarHeaderContentGuard).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreEnvvarHeaderCreate`: EnvVarHeaderContentGuardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreEnvvarHeaderCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envVarHeaderContentGuard** | [**EnvVarHeaderContentGuard**](EnvVarHeaderContentGuard.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**EnvVarHeaderContentGuardResponse**](EnvVarHeaderContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreEnvvarHeaderDelete

> ContentguardsCoreEnvvarHeaderDelete(ctx, envVarHeaderContentGuardHref).XTaskDiagnostics(xTaskDiagnostics).Execute()

Delete an env var header content guard



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	envVarHeaderContentGuardHref := "envVarHeaderContentGuardHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderDelete(context.Background(), envVarHeaderContentGuardHref).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envVarHeaderContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreEnvvarHeaderDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## ContentguardsCoreEnvvarHeaderList

> PaginatedEnvVarHeaderContentGuardResponseList ContentguardsCoreEnvvarHeaderList(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()

List env var header content guards



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string | Filter results where name matches value (optional)
	nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
	nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
	nameIexact := "nameIexact_example" // string | Filter results where name matches value (optional)
	nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
	nameIregex := "nameIregex_example" // string | Filter results where name matches regex value (optional)
	nameIstartswith := "nameIstartswith_example" // string | Filter results where name starts with value (optional)
	nameRegex := "nameRegex_example" // string | Filter results where name matches regex value (optional)
	nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `name` - Name* `-name` - Name (descending)* `description` - Description* `-description` - Description (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderList(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreEnvvarHeaderList`: PaginatedEnvVarHeaderContentGuardResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreEnvvarHeaderListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIexact** | **string** | Filter results where name matches value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameIregex** | **string** | Filter results where name matches regex value | 
 **nameIstartswith** | **string** | Filter results where name starts with value | 
 **nameRegex** | **string** | Filter results where name matches regex value | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;description&#x60; - Description* &#x60;-description&#x60; - Description (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedEnvVarHeaderContentGuardResponseList**](PaginatedEnvVarHeaderContentGuardResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreEnvvarHeaderListRoles

> ObjectRolesResponse ContentguardsCoreEnvvarHeaderListRoles(ctx, envVarHeaderContentGuardHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

List roles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	envVarHeaderContentGuardHref := "envVarHeaderContentGuardHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderListRoles(context.Background(), envVarHeaderContentGuardHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreEnvvarHeaderListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envVarHeaderContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreEnvvarHeaderListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
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


## ContentguardsCoreEnvvarHeaderMyPermissions

> MyPermissionsResponse ContentguardsCoreEnvvarHeaderMyPermissions(ctx, envVarHeaderContentGuardHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

List user permissions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	envVarHeaderContentGuardHref := "envVarHeaderContentGuardHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderMyPermissions(context.Background(), envVarHeaderContentGuardHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreEnvvarHeaderMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envVarHeaderContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreEnvvarHeaderMyPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
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


## ContentguardsCoreEnvvarHeaderPartialUpdate

> EnvVarHeaderContentGuardResponse ContentguardsCoreEnvvarHeaderPartialUpdate(ctx, envVarHeaderContentGuardHref).PatchedEnvVarHeaderContentGuard(patchedEnvVarHeaderContentGuard).XTaskDiagnostics(xTaskDiagnostics).Execute()

Update an env var header content guard



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	envVarHeaderContentGuardHref := "envVarHeaderContentGuardHref_example" // string | 
	patchedEnvVarHeaderContentGuard := *openapiclient.NewPatchedEnvVarHeaderContentGuard() // PatchedEnvVarHeaderContentGuard | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderPartialUpdate(context.Background(), envVarHeaderContentGuardHref).PatchedEnvVarHeaderContentGuard(patchedEnvVarHeaderContentGuard).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreEnvvarHeaderPartialUpdate`: EnvVarHeaderContentGuardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envVarHeaderContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreEnvvarHeaderPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEnvVarHeaderContentGuard** | [**PatchedEnvVarHeaderContentGuard**](PatchedEnvVarHeaderContentGuard.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**EnvVarHeaderContentGuardResponse**](EnvVarHeaderContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreEnvvarHeaderRead

> EnvVarHeaderContentGuardResponse ContentguardsCoreEnvvarHeaderRead(ctx, envVarHeaderContentGuardHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an env var header content guard



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	envVarHeaderContentGuardHref := "envVarHeaderContentGuardHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderRead(context.Background(), envVarHeaderContentGuardHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreEnvvarHeaderRead`: EnvVarHeaderContentGuardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envVarHeaderContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreEnvvarHeaderReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**EnvVarHeaderContentGuardResponse**](EnvVarHeaderContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreEnvvarHeaderRemoveRole

> NestedRoleResponse ContentguardsCoreEnvvarHeaderRemoveRole(ctx, envVarHeaderContentGuardHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()

Remove a role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	envVarHeaderContentGuardHref := "envVarHeaderContentGuardHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderRemoveRole(context.Background(), envVarHeaderContentGuardHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreEnvvarHeaderRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envVarHeaderContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreEnvvarHeaderRemoveRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nestedRole** | [**NestedRole**](NestedRole.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## ContentguardsCoreEnvvarHeaderUpdate

> EnvVarHeaderContentGuardResponse ContentguardsCoreEnvvarHeaderUpdate(ctx, envVarHeaderContentGuardHref).EnvVarHeaderContentGuard(envVarHeaderContentGuard).XTaskDiagnostics(xTaskDiagnostics).Execute()

Update an env var header content guard



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	envVarHeaderContentGuardHref := "envVarHeaderContentGuardHref_example" // string | 
	envVarHeaderContentGuard := *openapiclient.NewEnvVarHeaderContentGuard("Name_example", "HeaderName_example", "EnvVar_example") // EnvVarHeaderContentGuard | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderUpdate(context.Background(), envVarHeaderContentGuardHref).EnvVarHeaderContentGuard(envVarHeaderContentGuard).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreEnvvarHeaderUpdate`: EnvVarHeaderContentGuardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsEnvvarHeaderAPI.ContentguardsCoreEnvvarHeaderUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envVarHeaderContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreEnvvarHeaderUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envVarHeaderContentGuard** | [**EnvVarHeaderContentGuard**](EnvVarHeaderContentGuard.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**EnvVarHeaderContentGuardResponse**](EnvVarHeaderContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

