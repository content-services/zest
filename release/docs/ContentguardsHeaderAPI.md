# \ContentguardsHeaderAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentguardsCoreHeaderAddRole**](ContentguardsHeaderAPI.md#ContentguardsCoreHeaderAddRole) | **Post** /{header_content_guard_href}add_role/ | Add a role
[**ContentguardsCoreHeaderCreate**](ContentguardsHeaderAPI.md#ContentguardsCoreHeaderCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/contentguards/core/header/ | Create a header content guard
[**ContentguardsCoreHeaderDelete**](ContentguardsHeaderAPI.md#ContentguardsCoreHeaderDelete) | **Delete** /{header_content_guard_href} | Delete a header content guard
[**ContentguardsCoreHeaderList**](ContentguardsHeaderAPI.md#ContentguardsCoreHeaderList) | **Get** /api/pulp/{pulp_domain}/api/v3/contentguards/core/header/ | List header content guards
[**ContentguardsCoreHeaderListRoles**](ContentguardsHeaderAPI.md#ContentguardsCoreHeaderListRoles) | **Get** /{header_content_guard_href}list_roles/ | List roles
[**ContentguardsCoreHeaderMyPermissions**](ContentguardsHeaderAPI.md#ContentguardsCoreHeaderMyPermissions) | **Get** /{header_content_guard_href}my_permissions/ | List user permissions
[**ContentguardsCoreHeaderPartialUpdate**](ContentguardsHeaderAPI.md#ContentguardsCoreHeaderPartialUpdate) | **Patch** /{header_content_guard_href} | Update a header content guard
[**ContentguardsCoreHeaderRead**](ContentguardsHeaderAPI.md#ContentguardsCoreHeaderRead) | **Get** /{header_content_guard_href} | Inspect a header content guard
[**ContentguardsCoreHeaderRemoveRole**](ContentguardsHeaderAPI.md#ContentguardsCoreHeaderRemoveRole) | **Post** /{header_content_guard_href}remove_role/ | Remove a role
[**ContentguardsCoreHeaderUpdate**](ContentguardsHeaderAPI.md#ContentguardsCoreHeaderUpdate) | **Put** /{header_content_guard_href} | Update a header content guard



## ContentguardsCoreHeaderAddRole

> NestedRoleResponse ContentguardsCoreHeaderAddRole(ctx, headerContentGuardHref).NestedRole(nestedRole).Execute()

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
	headerContentGuardHref := "headerContentGuardHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderAddRole(context.Background(), headerContentGuardHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsHeaderAPI.ContentguardsCoreHeaderAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreHeaderAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsHeaderAPI.ContentguardsCoreHeaderAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**headerContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreHeaderAddRoleRequest struct via the builder pattern


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


## ContentguardsCoreHeaderCreate

> HeaderContentGuardResponse ContentguardsCoreHeaderCreate(ctx, pulpDomain).HeaderContentGuard(headerContentGuard).Execute()

Create a header content guard



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
	headerContentGuard := *openapiclient.NewHeaderContentGuard("Name_example", "HeaderName_example", "HeaderValue_example") // HeaderContentGuard | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderCreate(context.Background(), pulpDomain).HeaderContentGuard(headerContentGuard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsHeaderAPI.ContentguardsCoreHeaderCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreHeaderCreate`: HeaderContentGuardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsHeaderAPI.ContentguardsCoreHeaderCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreHeaderCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **headerContentGuard** | [**HeaderContentGuard**](HeaderContentGuard.md) |  | 

### Return type

[**HeaderContentGuardResponse**](HeaderContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreHeaderDelete

> ContentguardsCoreHeaderDelete(ctx, headerContentGuardHref).Execute()

Delete a header content guard



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
	headerContentGuardHref := "headerContentGuardHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderDelete(context.Background(), headerContentGuardHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsHeaderAPI.ContentguardsCoreHeaderDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**headerContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreHeaderDeleteRequest struct via the builder pattern


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


## ContentguardsCoreHeaderList

> PaginatedHeaderContentGuardResponseList ContentguardsCoreHeaderList(ctx, pulpDomain).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()

List header content guards



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
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `name` - Name* `-name` - Name (descending)* `description` - Description* `-description` - Description (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderList(context.Background(), pulpDomain).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsHeaderAPI.ContentguardsCoreHeaderList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreHeaderList`: PaginatedHeaderContentGuardResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsHeaderAPI.ContentguardsCoreHeaderList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreHeaderListRequest struct via the builder pattern


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
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;description&#x60; - Description* &#x60;-description&#x60; - Description (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedHeaderContentGuardResponseList**](PaginatedHeaderContentGuardResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreHeaderListRoles

> ObjectRolesResponse ContentguardsCoreHeaderListRoles(ctx, headerContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	headerContentGuardHref := "headerContentGuardHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderListRoles(context.Background(), headerContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsHeaderAPI.ContentguardsCoreHeaderListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreHeaderListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsHeaderAPI.ContentguardsCoreHeaderListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**headerContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreHeaderListRolesRequest struct via the builder pattern


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


## ContentguardsCoreHeaderMyPermissions

> MyPermissionsResponse ContentguardsCoreHeaderMyPermissions(ctx, headerContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	headerContentGuardHref := "headerContentGuardHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderMyPermissions(context.Background(), headerContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsHeaderAPI.ContentguardsCoreHeaderMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreHeaderMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsHeaderAPI.ContentguardsCoreHeaderMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**headerContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreHeaderMyPermissionsRequest struct via the builder pattern


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


## ContentguardsCoreHeaderPartialUpdate

> HeaderContentGuardResponse ContentguardsCoreHeaderPartialUpdate(ctx, headerContentGuardHref).PatchedHeaderContentGuard(patchedHeaderContentGuard).Execute()

Update a header content guard



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
	headerContentGuardHref := "headerContentGuardHref_example" // string | 
	patchedHeaderContentGuard := *openapiclient.NewPatchedHeaderContentGuard() // PatchedHeaderContentGuard | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderPartialUpdate(context.Background(), headerContentGuardHref).PatchedHeaderContentGuard(patchedHeaderContentGuard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsHeaderAPI.ContentguardsCoreHeaderPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreHeaderPartialUpdate`: HeaderContentGuardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsHeaderAPI.ContentguardsCoreHeaderPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**headerContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreHeaderPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedHeaderContentGuard** | [**PatchedHeaderContentGuard**](PatchedHeaderContentGuard.md) |  | 

### Return type

[**HeaderContentGuardResponse**](HeaderContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreHeaderRead

> HeaderContentGuardResponse ContentguardsCoreHeaderRead(ctx, headerContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a header content guard



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
	headerContentGuardHref := "headerContentGuardHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderRead(context.Background(), headerContentGuardHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsHeaderAPI.ContentguardsCoreHeaderRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreHeaderRead`: HeaderContentGuardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsHeaderAPI.ContentguardsCoreHeaderRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**headerContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreHeaderReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**HeaderContentGuardResponse**](HeaderContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentguardsCoreHeaderRemoveRole

> NestedRoleResponse ContentguardsCoreHeaderRemoveRole(ctx, headerContentGuardHref).NestedRole(nestedRole).Execute()

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
	headerContentGuardHref := "headerContentGuardHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderRemoveRole(context.Background(), headerContentGuardHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsHeaderAPI.ContentguardsCoreHeaderRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreHeaderRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsHeaderAPI.ContentguardsCoreHeaderRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**headerContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreHeaderRemoveRoleRequest struct via the builder pattern


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


## ContentguardsCoreHeaderUpdate

> HeaderContentGuardResponse ContentguardsCoreHeaderUpdate(ctx, headerContentGuardHref).HeaderContentGuard(headerContentGuard).Execute()

Update a header content guard



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
	headerContentGuardHref := "headerContentGuardHref_example" // string | 
	headerContentGuard := *openapiclient.NewHeaderContentGuard("Name_example", "HeaderName_example", "HeaderValue_example") // HeaderContentGuard | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderUpdate(context.Background(), headerContentGuardHref).HeaderContentGuard(headerContentGuard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentguardsHeaderAPI.ContentguardsCoreHeaderUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentguardsCoreHeaderUpdate`: HeaderContentGuardResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentguardsHeaderAPI.ContentguardsCoreHeaderUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**headerContentGuardHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentguardsCoreHeaderUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **headerContentGuard** | [**HeaderContentGuard**](HeaderContentGuard.md) |  | 

### Return type

[**HeaderContentGuardResponse**](HeaderContentGuardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

