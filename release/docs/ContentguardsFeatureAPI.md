# \ContentguardsFeatureAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentguardsServiceFeatureAddRole**](ContentguardsFeatureAPI.md#ContentguardsServiceFeatureAddRole) | **Post** /{service_feature_content_guard_href}add_role/ | Add a role
[**ContentguardsServiceFeatureCreate**](ContentguardsFeatureAPI.md#ContentguardsServiceFeatureCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/contentguards/service/feature/ | Create a feature content guard
[**ContentguardsServiceFeatureDelete**](ContentguardsFeatureAPI.md#ContentguardsServiceFeatureDelete) | **Delete** /{service_feature_content_guard_href} | Delete a feature content guard
[**ContentguardsServiceFeatureList**](ContentguardsFeatureAPI.md#ContentguardsServiceFeatureList) | **Get** /api/pulp/{pulp_domain}/api/v3/contentguards/service/feature/ | List feature content guards
[**ContentguardsServiceFeatureListRoles**](ContentguardsFeatureAPI.md#ContentguardsServiceFeatureListRoles) | **Get** /{service_feature_content_guard_href}list_roles/ | List roles
[**ContentguardsServiceFeatureMyPermissions**](ContentguardsFeatureAPI.md#ContentguardsServiceFeatureMyPermissions) | **Get** /{service_feature_content_guard_href}my_permissions/ | List user permissions
[**ContentguardsServiceFeaturePartialUpdate**](ContentguardsFeatureAPI.md#ContentguardsServiceFeaturePartialUpdate) | **Patch** /{service_feature_content_guard_href} | Update a feature content guard
[**ContentguardsServiceFeatureRead**](ContentguardsFeatureAPI.md#ContentguardsServiceFeatureRead) | **Get** /{service_feature_content_guard_href} | Inspect a feature content guard
[**ContentguardsServiceFeatureRemoveRole**](ContentguardsFeatureAPI.md#ContentguardsServiceFeatureRemoveRole) | **Post** /{service_feature_content_guard_href}remove_role/ | Remove a role
[**ContentguardsServiceFeatureUpdate**](ContentguardsFeatureAPI.md#ContentguardsServiceFeatureUpdate) | **Put** /{service_feature_content_guard_href} | Update a feature content guard



## ContentguardsServiceFeatureAddRole

> NestedRoleResponse ContentguardsServiceFeatureAddRole(ctx, serviceFeatureContentGuardHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	serviceFeatureContentGuardHref := "serviceFeatureContentGuardHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsFeatureAPI.ContentguardsServiceFeatureAddRole(context.Background(), serviceFeatureContentGuardHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsFeatureAPI.ContentguardsServiceFeatureAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsServiceFeatureAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsFeatureAPI.ContentguardsServiceFeatureAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceFeatureContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsServiceFeatureAddRoleRequest struct via the builder pattern


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


## ContentguardsServiceFeatureCreate

> ServiceFeatureContentGuardResponse ContentguardsServiceFeatureCreate(ctx, pulpDomain).ServiceFeatureContentGuard(serviceFeatureContentGuard).XTaskDiagnostics(xTaskDiagnostics).Execute()

Create a feature content guard



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
	serviceFeatureContentGuard := *openapiclient.NewServiceFeatureContentGuard("Name_example", "HeaderName_example", []string{"Features_example"}) // ServiceFeatureContentGuard | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsFeatureAPI.ContentguardsServiceFeatureCreate(context.Background(), pulpDomain).ServiceFeatureContentGuard(serviceFeatureContentGuard).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsFeatureAPI.ContentguardsServiceFeatureCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsServiceFeatureCreate`: ServiceFeatureContentGuardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsFeatureAPI.ContentguardsServiceFeatureCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsServiceFeatureCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceFeatureContentGuard** | [**ServiceFeatureContentGuard**](ServiceFeatureContentGuard.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**ServiceFeatureContentGuardResponse**](ServiceFeatureContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsServiceFeatureDelete

> ContentguardsServiceFeatureDelete(ctx, serviceFeatureContentGuardHref).XTaskDiagnostics(xTaskDiagnostics).Execute()

Delete a feature content guard



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
	serviceFeatureContentGuardHref := "serviceFeatureContentGuardHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContentguardsFeatureAPI.ContentguardsServiceFeatureDelete(context.Background(), serviceFeatureContentGuardHref).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsFeatureAPI.ContentguardsServiceFeatureDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceFeatureContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsServiceFeatureDeleteRequest struct via the builder pattern


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


## ContentguardsServiceFeatureList

> PaginatedserviceFeatureContentGuardResponseList ContentguardsServiceFeatureList(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()

List feature content guards



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
	resp, r, err := apiClient.ContentguardsFeatureAPI.ContentguardsServiceFeatureList(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsFeatureAPI.ContentguardsServiceFeatureList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsServiceFeatureList`: PaginatedserviceFeatureContentGuardResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsFeatureAPI.ContentguardsServiceFeatureList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsServiceFeatureListRequest struct via the builder pattern


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

[**PaginatedserviceFeatureContentGuardResponseList**](PaginatedserviceFeatureContentGuardResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsServiceFeatureListRoles

> ObjectRolesResponse ContentguardsServiceFeatureListRoles(ctx, serviceFeatureContentGuardHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	serviceFeatureContentGuardHref := "serviceFeatureContentGuardHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsFeatureAPI.ContentguardsServiceFeatureListRoles(context.Background(), serviceFeatureContentGuardHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsFeatureAPI.ContentguardsServiceFeatureListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsServiceFeatureListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsFeatureAPI.ContentguardsServiceFeatureListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceFeatureContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsServiceFeatureListRolesRequest struct via the builder pattern


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


## ContentguardsServiceFeatureMyPermissions

> MyPermissionsResponse ContentguardsServiceFeatureMyPermissions(ctx, serviceFeatureContentGuardHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	serviceFeatureContentGuardHref := "serviceFeatureContentGuardHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsFeatureAPI.ContentguardsServiceFeatureMyPermissions(context.Background(), serviceFeatureContentGuardHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsFeatureAPI.ContentguardsServiceFeatureMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsServiceFeatureMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsFeatureAPI.ContentguardsServiceFeatureMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceFeatureContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsServiceFeatureMyPermissionsRequest struct via the builder pattern


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


## ContentguardsServiceFeaturePartialUpdate

> ServiceFeatureContentGuardResponse ContentguardsServiceFeaturePartialUpdate(ctx, serviceFeatureContentGuardHref).PatchedserviceFeatureContentGuard(patchedserviceFeatureContentGuard).XTaskDiagnostics(xTaskDiagnostics).Execute()

Update a feature content guard



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
	serviceFeatureContentGuardHref := "serviceFeatureContentGuardHref_example" // string | 
	patchedserviceFeatureContentGuard := *openapiclient.NewPatchedserviceFeatureContentGuard() // PatchedserviceFeatureContentGuard | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsFeatureAPI.ContentguardsServiceFeaturePartialUpdate(context.Background(), serviceFeatureContentGuardHref).PatchedserviceFeatureContentGuard(patchedserviceFeatureContentGuard).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsFeatureAPI.ContentguardsServiceFeaturePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsServiceFeaturePartialUpdate`: ServiceFeatureContentGuardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsFeatureAPI.ContentguardsServiceFeaturePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceFeatureContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsServiceFeaturePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedserviceFeatureContentGuard** | [**PatchedserviceFeatureContentGuard**](PatchedserviceFeatureContentGuard.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**ServiceFeatureContentGuardResponse**](ServiceFeatureContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsServiceFeatureRead

> ServiceFeatureContentGuardResponse ContentguardsServiceFeatureRead(ctx, serviceFeatureContentGuardHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a feature content guard



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
	serviceFeatureContentGuardHref := "serviceFeatureContentGuardHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsFeatureAPI.ContentguardsServiceFeatureRead(context.Background(), serviceFeatureContentGuardHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsFeatureAPI.ContentguardsServiceFeatureRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsServiceFeatureRead`: ServiceFeatureContentGuardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsFeatureAPI.ContentguardsServiceFeatureRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceFeatureContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsServiceFeatureReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ServiceFeatureContentGuardResponse**](ServiceFeatureContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsServiceFeatureRemoveRole

> NestedRoleResponse ContentguardsServiceFeatureRemoveRole(ctx, serviceFeatureContentGuardHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	serviceFeatureContentGuardHref := "serviceFeatureContentGuardHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsFeatureAPI.ContentguardsServiceFeatureRemoveRole(context.Background(), serviceFeatureContentGuardHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsFeatureAPI.ContentguardsServiceFeatureRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsServiceFeatureRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsFeatureAPI.ContentguardsServiceFeatureRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceFeatureContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsServiceFeatureRemoveRoleRequest struct via the builder pattern


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


## ContentguardsServiceFeatureUpdate

> ServiceFeatureContentGuardResponse ContentguardsServiceFeatureUpdate(ctx, serviceFeatureContentGuardHref).ServiceFeatureContentGuard(serviceFeatureContentGuard).XTaskDiagnostics(xTaskDiagnostics).Execute()

Update a feature content guard



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
	serviceFeatureContentGuardHref := "serviceFeatureContentGuardHref_example" // string | 
	serviceFeatureContentGuard := *openapiclient.NewServiceFeatureContentGuard("Name_example", "HeaderName_example", []string{"Features_example"}) // ServiceFeatureContentGuard | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsFeatureAPI.ContentguardsServiceFeatureUpdate(context.Background(), serviceFeatureContentGuardHref).ServiceFeatureContentGuard(serviceFeatureContentGuard).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsFeatureAPI.ContentguardsServiceFeatureUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsServiceFeatureUpdate`: ServiceFeatureContentGuardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsFeatureAPI.ContentguardsServiceFeatureUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceFeatureContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsServiceFeatureUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceFeatureContentGuard** | [**ServiceFeatureContentGuard**](ServiceFeatureContentGuard.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**ServiceFeatureContentGuardResponse**](ServiceFeatureContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

