# \ContentguardsCompositeAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentguardsCoreCompositeAddRole**](ContentguardsCompositeAPI.md#ContentguardsCoreCompositeAddRole) | **Post** /{composite_content_guard_href}add_role/ | Add a role
[**ContentguardsCoreCompositeCreate**](ContentguardsCompositeAPI.md#ContentguardsCoreCompositeCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/contentguards/core/composite/ | Create a composite content guard
[**ContentguardsCoreCompositeDelete**](ContentguardsCompositeAPI.md#ContentguardsCoreCompositeDelete) | **Delete** /{composite_content_guard_href} | Delete a composite content guard
[**ContentguardsCoreCompositeList**](ContentguardsCompositeAPI.md#ContentguardsCoreCompositeList) | **Get** /api/pulp/{pulp_domain}/api/v3/contentguards/core/composite/ | List composite content guards
[**ContentguardsCoreCompositeListRoles**](ContentguardsCompositeAPI.md#ContentguardsCoreCompositeListRoles) | **Get** /{composite_content_guard_href}list_roles/ | List roles
[**ContentguardsCoreCompositeMyPermissions**](ContentguardsCompositeAPI.md#ContentguardsCoreCompositeMyPermissions) | **Get** /{composite_content_guard_href}my_permissions/ | List user permissions
[**ContentguardsCoreCompositePartialUpdate**](ContentguardsCompositeAPI.md#ContentguardsCoreCompositePartialUpdate) | **Patch** /{composite_content_guard_href} | Update a composite content guard
[**ContentguardsCoreCompositeRead**](ContentguardsCompositeAPI.md#ContentguardsCoreCompositeRead) | **Get** /{composite_content_guard_href} | Inspect a composite content guard
[**ContentguardsCoreCompositeRemoveRole**](ContentguardsCompositeAPI.md#ContentguardsCoreCompositeRemoveRole) | **Post** /{composite_content_guard_href}remove_role/ | Remove a role
[**ContentguardsCoreCompositeUpdate**](ContentguardsCompositeAPI.md#ContentguardsCoreCompositeUpdate) | **Put** /{composite_content_guard_href} | Update a composite content guard



## ContentguardsCoreCompositeAddRole

> NestedRoleResponse ContentguardsCoreCompositeAddRole(ctx, compositeContentGuardHref).NestedRole(nestedRole).Execute()

Add a role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	compositeContentGuardHref := "compositeContentGuardHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsCompositeAPI.ContentguardsCoreCompositeAddRole(context.Background(), compositeContentGuardHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsCompositeAPI.ContentguardsCoreCompositeAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreCompositeAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsCompositeAPI.ContentguardsCoreCompositeAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**compositeContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreCompositeAddRoleRequest struct via the builder pattern


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


## ContentguardsCoreCompositeCreate

> CompositeContentGuardResponse ContentguardsCoreCompositeCreate(ctx, pulpDomain).CompositeContentGuard(compositeContentGuard).Execute()

Create a composite content guard



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	compositeContentGuard := *openapiclient.NewCompositeContentGuard("Name_example") // CompositeContentGuard | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsCompositeAPI.ContentguardsCoreCompositeCreate(context.Background(), pulpDomain).CompositeContentGuard(compositeContentGuard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsCompositeAPI.ContentguardsCoreCompositeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreCompositeCreate`: CompositeContentGuardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsCompositeAPI.ContentguardsCoreCompositeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreCompositeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **compositeContentGuard** | [**CompositeContentGuard**](CompositeContentGuard.md) |  | 

### Return type

[**CompositeContentGuardResponse**](CompositeContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreCompositeDelete

> ContentguardsCoreCompositeDelete(ctx, compositeContentGuardHref).Execute()

Delete a composite content guard



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	compositeContentGuardHref := "compositeContentGuardHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContentguardsCompositeAPI.ContentguardsCoreCompositeDelete(context.Background(), compositeContentGuardHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsCompositeAPI.ContentguardsCoreCompositeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**compositeContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreCompositeDeleteRequest struct via the builder pattern


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


## ContentguardsCoreCompositeList

> PaginatedCompositeContentGuardResponseList ContentguardsCoreCompositeList(ctx, pulpDomain).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()

List composite content guards



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
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
	ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `name` - Name * `-name` - Name (descending) * `description` - Description * `-description` - Description (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	q := "q_example" // string |  (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsCompositeAPI.ContentguardsCoreCompositeList(context.Background(), pulpDomain).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsCompositeAPI.ContentguardsCoreCompositeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreCompositeList`: PaginatedCompositeContentGuardResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsCompositeAPI.ContentguardsCoreCompositeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreCompositeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **q** | **string** |  | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedCompositeContentGuardResponseList**](PaginatedCompositeContentGuardResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreCompositeListRoles

> ObjectRolesResponse ContentguardsCoreCompositeListRoles(ctx, compositeContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()

List roles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	compositeContentGuardHref := "compositeContentGuardHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsCompositeAPI.ContentguardsCoreCompositeListRoles(context.Background(), compositeContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsCompositeAPI.ContentguardsCoreCompositeListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreCompositeListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsCompositeAPI.ContentguardsCoreCompositeListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**compositeContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreCompositeListRolesRequest struct via the builder pattern


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


## ContentguardsCoreCompositeMyPermissions

> MyPermissionsResponse ContentguardsCoreCompositeMyPermissions(ctx, compositeContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()

List user permissions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	compositeContentGuardHref := "compositeContentGuardHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsCompositeAPI.ContentguardsCoreCompositeMyPermissions(context.Background(), compositeContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsCompositeAPI.ContentguardsCoreCompositeMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreCompositeMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsCompositeAPI.ContentguardsCoreCompositeMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**compositeContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreCompositeMyPermissionsRequest struct via the builder pattern


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


## ContentguardsCoreCompositePartialUpdate

> CompositeContentGuardResponse ContentguardsCoreCompositePartialUpdate(ctx, compositeContentGuardHref).PatchedCompositeContentGuard(patchedCompositeContentGuard).Execute()

Update a composite content guard



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	compositeContentGuardHref := "compositeContentGuardHref_example" // string | 
	patchedCompositeContentGuard := *openapiclient.NewPatchedCompositeContentGuard() // PatchedCompositeContentGuard | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsCompositeAPI.ContentguardsCoreCompositePartialUpdate(context.Background(), compositeContentGuardHref).PatchedCompositeContentGuard(patchedCompositeContentGuard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsCompositeAPI.ContentguardsCoreCompositePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreCompositePartialUpdate`: CompositeContentGuardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsCompositeAPI.ContentguardsCoreCompositePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**compositeContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreCompositePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedCompositeContentGuard** | [**PatchedCompositeContentGuard**](PatchedCompositeContentGuard.md) |  | 

### Return type

[**CompositeContentGuardResponse**](CompositeContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreCompositeRead

> CompositeContentGuardResponse ContentguardsCoreCompositeRead(ctx, compositeContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a composite content guard



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	compositeContentGuardHref := "compositeContentGuardHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsCompositeAPI.ContentguardsCoreCompositeRead(context.Background(), compositeContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsCompositeAPI.ContentguardsCoreCompositeRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreCompositeRead`: CompositeContentGuardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsCompositeAPI.ContentguardsCoreCompositeRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**compositeContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreCompositeReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**CompositeContentGuardResponse**](CompositeContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreCompositeRemoveRole

> NestedRoleResponse ContentguardsCoreCompositeRemoveRole(ctx, compositeContentGuardHref).NestedRole(nestedRole).Execute()

Remove a role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	compositeContentGuardHref := "compositeContentGuardHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsCompositeAPI.ContentguardsCoreCompositeRemoveRole(context.Background(), compositeContentGuardHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsCompositeAPI.ContentguardsCoreCompositeRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreCompositeRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsCompositeAPI.ContentguardsCoreCompositeRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**compositeContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreCompositeRemoveRoleRequest struct via the builder pattern


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


## ContentguardsCoreCompositeUpdate

> CompositeContentGuardResponse ContentguardsCoreCompositeUpdate(ctx, compositeContentGuardHref).CompositeContentGuard(compositeContentGuard).Execute()

Update a composite content guard



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	compositeContentGuardHref := "compositeContentGuardHref_example" // string | 
	compositeContentGuard := *openapiclient.NewCompositeContentGuard("Name_example") // CompositeContentGuard | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsCompositeAPI.ContentguardsCoreCompositeUpdate(context.Background(), compositeContentGuardHref).CompositeContentGuard(compositeContentGuard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsCompositeAPI.ContentguardsCoreCompositeUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreCompositeUpdate`: CompositeContentGuardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsCompositeAPI.ContentguardsCoreCompositeUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**compositeContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreCompositeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **compositeContentGuard** | [**CompositeContentGuard**](CompositeContentGuard.md) |  | 

### Return type

[**CompositeContentGuardResponse**](CompositeContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

