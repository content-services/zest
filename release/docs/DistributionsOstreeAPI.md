# \DistributionsOstreeAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DistributionsOstreeOstreeAddRole**](DistributionsOstreeAPI.md#DistributionsOstreeOstreeAddRole) | **Post** /{ostree_ostree_distribution_href}add_role/ | Add a role
[**DistributionsOstreeOstreeCreate**](DistributionsOstreeAPI.md#DistributionsOstreeOstreeCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/distributions/ostree/ostree/ | Create an ostree distribution
[**DistributionsOstreeOstreeDelete**](DistributionsOstreeAPI.md#DistributionsOstreeOstreeDelete) | **Delete** /{ostree_ostree_distribution_href} | Delete an ostree distribution
[**DistributionsOstreeOstreeList**](DistributionsOstreeAPI.md#DistributionsOstreeOstreeList) | **Get** /api/pulp/{pulp_domain}/api/v3/distributions/ostree/ostree/ | List ostree distributions
[**DistributionsOstreeOstreeListRoles**](DistributionsOstreeAPI.md#DistributionsOstreeOstreeListRoles) | **Get** /{ostree_ostree_distribution_href}list_roles/ | List roles
[**DistributionsOstreeOstreeMyPermissions**](DistributionsOstreeAPI.md#DistributionsOstreeOstreeMyPermissions) | **Get** /{ostree_ostree_distribution_href}my_permissions/ | List user permissions
[**DistributionsOstreeOstreePartialUpdate**](DistributionsOstreeAPI.md#DistributionsOstreeOstreePartialUpdate) | **Patch** /{ostree_ostree_distribution_href} | Update an ostree distribution
[**DistributionsOstreeOstreeRead**](DistributionsOstreeAPI.md#DistributionsOstreeOstreeRead) | **Get** /{ostree_ostree_distribution_href} | Inspect an ostree distribution
[**DistributionsOstreeOstreeRemoveRole**](DistributionsOstreeAPI.md#DistributionsOstreeOstreeRemoveRole) | **Post** /{ostree_ostree_distribution_href}remove_role/ | Remove a role
[**DistributionsOstreeOstreeSetLabel**](DistributionsOstreeAPI.md#DistributionsOstreeOstreeSetLabel) | **Post** /{ostree_ostree_distribution_href}set_label/ | Set a label
[**DistributionsOstreeOstreeUnsetLabel**](DistributionsOstreeAPI.md#DistributionsOstreeOstreeUnsetLabel) | **Post** /{ostree_ostree_distribution_href}unset_label/ | Unset a label
[**DistributionsOstreeOstreeUpdate**](DistributionsOstreeAPI.md#DistributionsOstreeOstreeUpdate) | **Put** /{ostree_ostree_distribution_href} | Update an ostree distribution



## DistributionsOstreeOstreeAddRole

> NestedRoleResponse DistributionsOstreeOstreeAddRole(ctx, ostreeOstreeDistributionHref).NestedRole(nestedRole).Execute()

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
	ostreeOstreeDistributionHref := "ostreeOstreeDistributionHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOstreeAPI.DistributionsOstreeOstreeAddRole(context.Background(), ostreeOstreeDistributionHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOstreeAPI.DistributionsOstreeOstreeAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsOstreeOstreeAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOstreeAPI.DistributionsOstreeOstreeAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsOstreeOstreeAddRoleRequest struct via the builder pattern


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


## DistributionsOstreeOstreeCreate

> AsyncOperationResponse DistributionsOstreeOstreeCreate(ctx, pulpDomain).OstreeOstreeDistribution(ostreeOstreeDistribution).Execute()

Create an ostree distribution



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
	ostreeOstreeDistribution := *openapiclient.NewOstreeOstreeDistribution("BasePath_example", "Name_example") // OstreeOstreeDistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOstreeAPI.DistributionsOstreeOstreeCreate(context.Background(), pulpDomain).OstreeOstreeDistribution(ostreeOstreeDistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOstreeAPI.DistributionsOstreeOstreeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsOstreeOstreeCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOstreeAPI.DistributionsOstreeOstreeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsOstreeOstreeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ostreeOstreeDistribution** | [**OstreeOstreeDistribution**](OstreeOstreeDistribution.md) |  | 

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


## DistributionsOstreeOstreeDelete

> AsyncOperationResponse DistributionsOstreeOstreeDelete(ctx, ostreeOstreeDistributionHref).Execute()

Delete an ostree distribution



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
	ostreeOstreeDistributionHref := "ostreeOstreeDistributionHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOstreeAPI.DistributionsOstreeOstreeDelete(context.Background(), ostreeOstreeDistributionHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOstreeAPI.DistributionsOstreeOstreeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsOstreeOstreeDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOstreeAPI.DistributionsOstreeOstreeDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsOstreeOstreeDeleteRequest struct via the builder pattern


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


## DistributionsOstreeOstreeList

> PaginatedostreeOstreeDistributionResponseList DistributionsOstreeOstreeList(ctx, pulpDomain).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Repository(repository).RepositoryIn(repositoryIn).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()

List ostree distributions



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
	basePath := "basePath_example" // string | Filter results where base_path matches value (optional)
	basePathContains := "basePathContains_example" // string | Filter results where base_path contains value (optional)
	basePathIcontains := "basePathIcontains_example" // string | Filter results where base_path contains value (optional)
	basePathIn := []string{"Inner_example"} // []string | Filter results where base_path is in a comma-separated list of values (optional)
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
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `name` - Name* `-name` - Name (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `base_path` - Base path* `-base_path` - Base path (descending)* `hidden` - Hidden* `-hidden` - Hidden (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string |  (optional)
	repository := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter results where repository matches value (optional)
	repositoryIn := []string{"Inner_example"} // []string | Filter results where repository is in a comma-separated list of values (optional)
	withContent := "withContent_example" // string | Filter distributions based on the content served by them (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOstreeAPI.DistributionsOstreeOstreeList(context.Background(), pulpDomain).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Repository(repository).RepositoryIn(repositoryIn).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOstreeAPI.DistributionsOstreeOstreeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsOstreeOstreeList`: PaginatedostreeOstreeDistributionResponseList
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOstreeAPI.DistributionsOstreeOstreeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsOstreeOstreeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **basePath** | **string** | Filter results where base_path matches value | 
 **basePathContains** | **string** | Filter results where base_path contains value | 
 **basePathIcontains** | **string** | Filter results where base_path contains value | 
 **basePathIn** | **[]string** | Filter results where base_path is in a comma-separated list of values | 
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
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;base_path&#x60; - Base path* &#x60;-base_path&#x60; - Base path (descending)* &#x60;hidden&#x60; - Hidden* &#x60;-hidden&#x60; - Hidden (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** |  | 
 **repository** | **string** | Filter results where repository matches value | 
 **repositoryIn** | **[]string** | Filter results where repository is in a comma-separated list of values | 
 **withContent** | **string** | Filter distributions based on the content served by them | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedostreeOstreeDistributionResponseList**](PaginatedostreeOstreeDistributionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsOstreeOstreeListRoles

> ObjectRolesResponse DistributionsOstreeOstreeListRoles(ctx, ostreeOstreeDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	ostreeOstreeDistributionHref := "ostreeOstreeDistributionHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOstreeAPI.DistributionsOstreeOstreeListRoles(context.Background(), ostreeOstreeDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOstreeAPI.DistributionsOstreeOstreeListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsOstreeOstreeListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOstreeAPI.DistributionsOstreeOstreeListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsOstreeOstreeListRolesRequest struct via the builder pattern


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


## DistributionsOstreeOstreeMyPermissions

> MyPermissionsResponse DistributionsOstreeOstreeMyPermissions(ctx, ostreeOstreeDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	ostreeOstreeDistributionHref := "ostreeOstreeDistributionHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOstreeAPI.DistributionsOstreeOstreeMyPermissions(context.Background(), ostreeOstreeDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOstreeAPI.DistributionsOstreeOstreeMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsOstreeOstreeMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOstreeAPI.DistributionsOstreeOstreeMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsOstreeOstreeMyPermissionsRequest struct via the builder pattern


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


## DistributionsOstreeOstreePartialUpdate

> AsyncOperationResponse DistributionsOstreeOstreePartialUpdate(ctx, ostreeOstreeDistributionHref).PatchedostreeOstreeDistribution(patchedostreeOstreeDistribution).Execute()

Update an ostree distribution



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
	ostreeOstreeDistributionHref := "ostreeOstreeDistributionHref_example" // string | 
	patchedostreeOstreeDistribution := *openapiclient.NewPatchedostreeOstreeDistribution() // PatchedostreeOstreeDistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOstreeAPI.DistributionsOstreeOstreePartialUpdate(context.Background(), ostreeOstreeDistributionHref).PatchedostreeOstreeDistribution(patchedostreeOstreeDistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOstreeAPI.DistributionsOstreeOstreePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsOstreeOstreePartialUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOstreeAPI.DistributionsOstreeOstreePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsOstreeOstreePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedostreeOstreeDistribution** | [**PatchedostreeOstreeDistribution**](PatchedostreeOstreeDistribution.md) |  | 

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


## DistributionsOstreeOstreeRead

> OstreeOstreeDistributionResponse DistributionsOstreeOstreeRead(ctx, ostreeOstreeDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an ostree distribution



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
	ostreeOstreeDistributionHref := "ostreeOstreeDistributionHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOstreeAPI.DistributionsOstreeOstreeRead(context.Background(), ostreeOstreeDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOstreeAPI.DistributionsOstreeOstreeRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsOstreeOstreeRead`: OstreeOstreeDistributionResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOstreeAPI.DistributionsOstreeOstreeRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsOstreeOstreeReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**OstreeOstreeDistributionResponse**](OstreeOstreeDistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsOstreeOstreeRemoveRole

> NestedRoleResponse DistributionsOstreeOstreeRemoveRole(ctx, ostreeOstreeDistributionHref).NestedRole(nestedRole).Execute()

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
	ostreeOstreeDistributionHref := "ostreeOstreeDistributionHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOstreeAPI.DistributionsOstreeOstreeRemoveRole(context.Background(), ostreeOstreeDistributionHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOstreeAPI.DistributionsOstreeOstreeRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsOstreeOstreeRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOstreeAPI.DistributionsOstreeOstreeRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsOstreeOstreeRemoveRoleRequest struct via the builder pattern


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


## DistributionsOstreeOstreeSetLabel

> SetLabelResponse DistributionsOstreeOstreeSetLabel(ctx, ostreeOstreeDistributionHref).SetLabel(setLabel).Execute()

Set a label



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
	ostreeOstreeDistributionHref := "ostreeOstreeDistributionHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOstreeAPI.DistributionsOstreeOstreeSetLabel(context.Background(), ostreeOstreeDistributionHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOstreeAPI.DistributionsOstreeOstreeSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsOstreeOstreeSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOstreeAPI.DistributionsOstreeOstreeSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsOstreeOstreeSetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setLabel** | [**SetLabel**](SetLabel.md) |  | 

### Return type

[**SetLabelResponse**](SetLabelResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsOstreeOstreeUnsetLabel

> UnsetLabelResponse DistributionsOstreeOstreeUnsetLabel(ctx, ostreeOstreeDistributionHref).UnsetLabel(unsetLabel).Execute()

Unset a label



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
	ostreeOstreeDistributionHref := "ostreeOstreeDistributionHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOstreeAPI.DistributionsOstreeOstreeUnsetLabel(context.Background(), ostreeOstreeDistributionHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOstreeAPI.DistributionsOstreeOstreeUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsOstreeOstreeUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOstreeAPI.DistributionsOstreeOstreeUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsOstreeOstreeUnsetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unsetLabel** | [**UnsetLabel**](UnsetLabel.md) |  | 

### Return type

[**UnsetLabelResponse**](UnsetLabelResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsOstreeOstreeUpdate

> AsyncOperationResponse DistributionsOstreeOstreeUpdate(ctx, ostreeOstreeDistributionHref).OstreeOstreeDistribution(ostreeOstreeDistribution).Execute()

Update an ostree distribution



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
	ostreeOstreeDistributionHref := "ostreeOstreeDistributionHref_example" // string | 
	ostreeOstreeDistribution := *openapiclient.NewOstreeOstreeDistribution("BasePath_example", "Name_example") // OstreeOstreeDistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsOstreeAPI.DistributionsOstreeOstreeUpdate(context.Background(), ostreeOstreeDistributionHref).OstreeOstreeDistribution(ostreeOstreeDistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsOstreeAPI.DistributionsOstreeOstreeUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsOstreeOstreeUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsOstreeAPI.DistributionsOstreeOstreeUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsOstreeOstreeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ostreeOstreeDistribution** | [**OstreeOstreeDistribution**](OstreeOstreeDistribution.md) |  | 

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

