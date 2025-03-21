# \PublicationsGemAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicationsGemGemAddRole**](PublicationsGemAPI.md#PublicationsGemGemAddRole) | **Post** /{gem_gem_publication_href}add_role/ | Add a role
[**PublicationsGemGemCreate**](PublicationsGemAPI.md#PublicationsGemGemCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/publications/gem/gem/ | Create a gem publication
[**PublicationsGemGemDelete**](PublicationsGemAPI.md#PublicationsGemGemDelete) | **Delete** /{gem_gem_publication_href} | Delete a gem publication
[**PublicationsGemGemList**](PublicationsGemAPI.md#PublicationsGemGemList) | **Get** /api/pulp/{pulp_domain}/api/v3/publications/gem/gem/ | List gem publications
[**PublicationsGemGemListRoles**](PublicationsGemAPI.md#PublicationsGemGemListRoles) | **Get** /{gem_gem_publication_href}list_roles/ | List roles
[**PublicationsGemGemMyPermissions**](PublicationsGemAPI.md#PublicationsGemGemMyPermissions) | **Get** /{gem_gem_publication_href}my_permissions/ | List user permissions
[**PublicationsGemGemRead**](PublicationsGemAPI.md#PublicationsGemGemRead) | **Get** /{gem_gem_publication_href} | Inspect a gem publication
[**PublicationsGemGemRemoveRole**](PublicationsGemAPI.md#PublicationsGemGemRemoveRole) | **Post** /{gem_gem_publication_href}remove_role/ | Remove a role



## PublicationsGemGemAddRole

> NestedRoleResponse PublicationsGemGemAddRole(ctx, gemGemPublicationHref).NestedRole(nestedRole).Execute()

Add a role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	gemGemPublicationHref := "gemGemPublicationHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicationsGemAPI.PublicationsGemGemAddRole(context.Background(), gemGemPublicationHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicationsGemAPI.PublicationsGemGemAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicationsGemGemAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicationsGemAPI.PublicationsGemGemAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemPublicationHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicationsGemGemAddRoleRequest struct via the builder pattern


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


## PublicationsGemGemCreate

> AsyncOperationResponse PublicationsGemGemCreate(ctx, pulpDomain).GemGemPublication(gemGemPublication).Execute()

Create a gem publication



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	gemGemPublication := *openapiclient.NewGemGemPublication() // GemGemPublication | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicationsGemAPI.PublicationsGemGemCreate(context.Background(), pulpDomain).GemGemPublication(gemGemPublication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicationsGemAPI.PublicationsGemGemCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicationsGemGemCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicationsGemAPI.PublicationsGemGemCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicationsGemGemCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gemGemPublication** | [**GemGemPublication**](GemGemPublication.md) |  | 

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


## PublicationsGemGemDelete

> PublicationsGemGemDelete(ctx, gemGemPublicationHref).Execute()

Delete a gem publication



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	gemGemPublicationHref := "gemGemPublicationHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicationsGemAPI.PublicationsGemGemDelete(context.Background(), gemGemPublicationHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicationsGemAPI.PublicationsGemGemDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemPublicationHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicationsGemGemDeleteRequest struct via the builder pattern


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


## PublicationsGemGemList

> PaginatedgemGemPublicationResponseList PublicationsGemGemList(ctx, pulpDomain).Checkpoint(checkpoint).Content(content).ContentIn(contentIn).Limit(limit).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedIsnull(pulpCreatedIsnull).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).Repository(repository).RepositoryVersion(repositoryVersion).Fields(fields).ExcludeFields(excludeFields).Execute()

List gem publications



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	checkpoint := true // bool | Filter results where checkpoint matches value (optional)
	content := "content_example" // string | Content Unit referenced by HREF/PRN (optional)
	contentIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `complete` - Complete* `-complete` - Complete (descending)* `pass_through` - Pass through* `-pass_through` - Pass through (descending)* `checkpoint` - Checkpoint* `-checkpoint` - Checkpoint (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpCreated := time.Now() // time.Time | Filter results where pulp_created matches value (optional)
	pulpCreatedGt := time.Now() // time.Time | Filter results where pulp_created is greater than value (optional)
	pulpCreatedGte := time.Now() // time.Time | Filter results where pulp_created is greater than or equal to value (optional)
	pulpCreatedIsnull := true // bool | Filter results where pulp_created has a null value (optional)
	pulpCreatedLt := time.Now() // time.Time | Filter results where pulp_created is less than value (optional)
	pulpCreatedLte := time.Now() // time.Time | Filter results where pulp_created is less than or equal to value (optional)
	pulpCreatedRange := []time.Time{time.Now()} // []time.Time | Filter results where pulp_created is between two comma separated values (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repository := "repository_example" // string | Repository referenced by HREF/PRN (optional)
	repositoryVersion := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Repository Version referenced by HREF/PRN (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicationsGemAPI.PublicationsGemGemList(context.Background(), pulpDomain).Checkpoint(checkpoint).Content(content).ContentIn(contentIn).Limit(limit).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedIsnull(pulpCreatedIsnull).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).Repository(repository).RepositoryVersion(repositoryVersion).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicationsGemAPI.PublicationsGemGemList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicationsGemGemList`: PaginatedgemGemPublicationResponseList
	fmt.Fprintf(os.Stdout, "Response from `PublicationsGemAPI.PublicationsGemGemList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicationsGemGemListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkpoint** | **bool** | Filter results where checkpoint matches value | 
 **content** | **string** | Content Unit referenced by HREF/PRN | 
 **contentIn** | **[]string** | Multiple values may be separated by commas. | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;complete&#x60; - Complete* &#x60;-complete&#x60; - Complete (descending)* &#x60;pass_through&#x60; - Pass through* &#x60;-pass_through&#x60; - Pass through (descending)* &#x60;checkpoint&#x60; - Checkpoint* &#x60;-checkpoint&#x60; - Checkpoint (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpCreated** | **time.Time** | Filter results where pulp_created matches value | 
 **pulpCreatedGt** | **time.Time** | Filter results where pulp_created is greater than value | 
 **pulpCreatedGte** | **time.Time** | Filter results where pulp_created is greater than or equal to value | 
 **pulpCreatedIsnull** | **bool** | Filter results where pulp_created has a null value | 
 **pulpCreatedLt** | **time.Time** | Filter results where pulp_created is less than value | 
 **pulpCreatedLte** | **time.Time** | Filter results where pulp_created is less than or equal to value | 
 **pulpCreatedRange** | [**[]time.Time**](time.Time.md) | Filter results where pulp_created is between two comma separated values | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repository** | **string** | Repository referenced by HREF/PRN | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF/PRN | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedgemGemPublicationResponseList**](PaginatedgemGemPublicationResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicationsGemGemListRoles

> ObjectRolesResponse PublicationsGemGemListRoles(ctx, gemGemPublicationHref).Fields(fields).ExcludeFields(excludeFields).Execute()

List roles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	gemGemPublicationHref := "gemGemPublicationHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicationsGemAPI.PublicationsGemGemListRoles(context.Background(), gemGemPublicationHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicationsGemAPI.PublicationsGemGemListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicationsGemGemListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicationsGemAPI.PublicationsGemGemListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemPublicationHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicationsGemGemListRolesRequest struct via the builder pattern


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


## PublicationsGemGemMyPermissions

> MyPermissionsResponse PublicationsGemGemMyPermissions(ctx, gemGemPublicationHref).Fields(fields).ExcludeFields(excludeFields).Execute()

List user permissions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	gemGemPublicationHref := "gemGemPublicationHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicationsGemAPI.PublicationsGemGemMyPermissions(context.Background(), gemGemPublicationHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicationsGemAPI.PublicationsGemGemMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicationsGemGemMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicationsGemAPI.PublicationsGemGemMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemPublicationHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicationsGemGemMyPermissionsRequest struct via the builder pattern


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


## PublicationsGemGemRead

> GemGemPublicationResponse PublicationsGemGemRead(ctx, gemGemPublicationHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a gem publication



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	gemGemPublicationHref := "gemGemPublicationHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicationsGemAPI.PublicationsGemGemRead(context.Background(), gemGemPublicationHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicationsGemAPI.PublicationsGemGemRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicationsGemGemRead`: GemGemPublicationResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicationsGemAPI.PublicationsGemGemRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemPublicationHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicationsGemGemReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**GemGemPublicationResponse**](GemGemPublicationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicationsGemGemRemoveRole

> NestedRoleResponse PublicationsGemGemRemoveRole(ctx, gemGemPublicationHref).NestedRole(nestedRole).Execute()

Remove a role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func main() {
	gemGemPublicationHref := "gemGemPublicationHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicationsGemAPI.PublicationsGemGemRemoveRole(context.Background(), gemGemPublicationHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicationsGemAPI.PublicationsGemGemRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicationsGemGemRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicationsGemAPI.PublicationsGemGemRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemPublicationHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicationsGemGemRemoveRoleRequest struct via the builder pattern


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

