# \DistributionsContainerAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DistributionsContainerContainerAddRole**](DistributionsContainerAPI.md#DistributionsContainerContainerAddRole) | **Post** /{container_container_distribution_href}add_role/ | Add a role
[**DistributionsContainerContainerCreate**](DistributionsContainerAPI.md#DistributionsContainerContainerCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/distributions/container/container/ | Create a container distribution
[**DistributionsContainerContainerDelete**](DistributionsContainerAPI.md#DistributionsContainerContainerDelete) | **Delete** /{container_container_distribution_href} | Delete a container distribution
[**DistributionsContainerContainerList**](DistributionsContainerAPI.md#DistributionsContainerContainerList) | **Get** /api/pulp/{pulp_domain}/api/v3/distributions/container/container/ | List container distributions
[**DistributionsContainerContainerListRoles**](DistributionsContainerAPI.md#DistributionsContainerContainerListRoles) | **Get** /{container_container_distribution_href}list_roles/ | List roles
[**DistributionsContainerContainerMyPermissions**](DistributionsContainerAPI.md#DistributionsContainerContainerMyPermissions) | **Get** /{container_container_distribution_href}my_permissions/ | List user permissions
[**DistributionsContainerContainerPartialUpdate**](DistributionsContainerAPI.md#DistributionsContainerContainerPartialUpdate) | **Patch** /{container_container_distribution_href} | Update a container distribution
[**DistributionsContainerContainerRead**](DistributionsContainerAPI.md#DistributionsContainerContainerRead) | **Get** /{container_container_distribution_href} | Inspect a container distribution
[**DistributionsContainerContainerRemoveRole**](DistributionsContainerAPI.md#DistributionsContainerContainerRemoveRole) | **Post** /{container_container_distribution_href}remove_role/ | Remove a role
[**DistributionsContainerContainerSetLabel**](DistributionsContainerAPI.md#DistributionsContainerContainerSetLabel) | **Post** /{container_container_distribution_href}set_label/ | Set a label
[**DistributionsContainerContainerUnsetLabel**](DistributionsContainerAPI.md#DistributionsContainerContainerUnsetLabel) | **Post** /{container_container_distribution_href}unset_label/ | Unset a label
[**DistributionsContainerContainerUpdate**](DistributionsContainerAPI.md#DistributionsContainerContainerUpdate) | **Put** /{container_container_distribution_href} | Update a container distribution



## DistributionsContainerContainerAddRole

> NestedRoleResponse DistributionsContainerContainerAddRole(ctx, containerContainerDistributionHref).NestedRole(nestedRole).Execute()

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
	containerContainerDistributionHref := "containerContainerDistributionHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerAddRole(context.Background(), containerContainerDistributionHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerAPI.DistributionsContainerContainerAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsContainerContainerAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerAPI.DistributionsContainerContainerAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerAddRoleRequest struct via the builder pattern


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


## DistributionsContainerContainerCreate

> AsyncOperationResponse DistributionsContainerContainerCreate(ctx, pulpDomain).ContainerContainerDistribution(containerContainerDistribution).Execute()

Create a container distribution



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
	containerContainerDistribution := *openapiclient.NewContainerContainerDistribution("BasePath_example", "Name_example") // ContainerContainerDistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerCreate(context.Background(), pulpDomain).ContainerContainerDistribution(containerContainerDistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerAPI.DistributionsContainerContainerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsContainerContainerCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerAPI.DistributionsContainerContainerCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerContainerDistribution** | [**ContainerContainerDistribution**](ContainerContainerDistribution.md) |  | 

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


## DistributionsContainerContainerDelete

> AsyncOperationResponse DistributionsContainerContainerDelete(ctx, containerContainerDistributionHref).Execute()

Delete a container distribution



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
	containerContainerDistributionHref := "containerContainerDistributionHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerDelete(context.Background(), containerContainerDistributionHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerAPI.DistributionsContainerContainerDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsContainerContainerDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerAPI.DistributionsContainerContainerDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerDeleteRequest struct via the builder pattern


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


## DistributionsContainerContainerList

> PaginatedcontainerContainerDistributionResponseList DistributionsContainerContainerList(ctx, pulpDomain).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Checkpoint(checkpoint).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).NamespaceName(namespaceName).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Repository(repository).RepositoryIn(repositoryIn).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()

List container distributions



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
	basePath := "basePath_example" // string | Filter results where base_path matches value (optional)
	basePathContains := "basePathContains_example" // string | Filter results where base_path contains value (optional)
	basePathIcontains := "basePathIcontains_example" // string | Filter results where base_path contains value (optional)
	basePathIn := []string{"Inner_example"} // []string | Filter results where base_path is in a comma-separated list of values (optional)
	checkpoint := true // bool | Filter results where checkpoint matches value (optional)
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
	namespaceName := "namespaceName_example" // string |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `name` - Name* `-name` - Name (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `base_path` - Base path* `-base_path` - Base path (descending)* `hidden` - Hidden* `-hidden` - Hidden (descending)* `checkpoint` - Checkpoint* `-checkpoint` - Checkpoint (descending)* `private` - Private* `-private` - Private (descending)* `description` - Description* `-description` - Description (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repository := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter results where repository matches value (optional)
	repositoryIn := []string{"Inner_example"} // []string | Filter results where repository is in a comma-separated list of values (optional)
	withContent := "withContent_example" // string | Filter distributions based on the content served by them (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerList(context.Background(), pulpDomain).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Checkpoint(checkpoint).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).NamespaceName(namespaceName).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Repository(repository).RepositoryIn(repositoryIn).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerAPI.DistributionsContainerContainerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsContainerContainerList`: PaginatedcontainerContainerDistributionResponseList
	fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerAPI.DistributionsContainerContainerList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **basePath** | **string** | Filter results where base_path matches value | 
 **basePathContains** | **string** | Filter results where base_path contains value | 
 **basePathIcontains** | **string** | Filter results where base_path contains value | 
 **basePathIn** | **[]string** | Filter results where base_path is in a comma-separated list of values | 
 **checkpoint** | **bool** | Filter results where checkpoint matches value | 
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
 **namespaceName** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;base_path&#x60; - Base path* &#x60;-base_path&#x60; - Base path (descending)* &#x60;hidden&#x60; - Hidden* &#x60;-hidden&#x60; - Hidden (descending)* &#x60;checkpoint&#x60; - Checkpoint* &#x60;-checkpoint&#x60; - Checkpoint (descending)* &#x60;private&#x60; - Private* &#x60;-private&#x60; - Private (descending)* &#x60;description&#x60; - Description* &#x60;-description&#x60; - Description (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repository** | **string** | Filter results where repository matches value | 
 **repositoryIn** | **[]string** | Filter results where repository is in a comma-separated list of values | 
 **withContent** | **string** | Filter distributions based on the content served by them | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedcontainerContainerDistributionResponseList**](PaginatedcontainerContainerDistributionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsContainerContainerListRoles

> ObjectRolesResponse DistributionsContainerContainerListRoles(ctx, containerContainerDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	containerContainerDistributionHref := "containerContainerDistributionHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerListRoles(context.Background(), containerContainerDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerAPI.DistributionsContainerContainerListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsContainerContainerListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerAPI.DistributionsContainerContainerListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerListRolesRequest struct via the builder pattern


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


## DistributionsContainerContainerMyPermissions

> MyPermissionsResponse DistributionsContainerContainerMyPermissions(ctx, containerContainerDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	containerContainerDistributionHref := "containerContainerDistributionHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerMyPermissions(context.Background(), containerContainerDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerAPI.DistributionsContainerContainerMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsContainerContainerMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerAPI.DistributionsContainerContainerMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerMyPermissionsRequest struct via the builder pattern


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


## DistributionsContainerContainerPartialUpdate

> AsyncOperationResponse DistributionsContainerContainerPartialUpdate(ctx, containerContainerDistributionHref).PatchedcontainerContainerDistribution(patchedcontainerContainerDistribution).Execute()

Update a container distribution



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
	containerContainerDistributionHref := "containerContainerDistributionHref_example" // string | 
	patchedcontainerContainerDistribution := *openapiclient.NewPatchedcontainerContainerDistribution() // PatchedcontainerContainerDistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerPartialUpdate(context.Background(), containerContainerDistributionHref).PatchedcontainerContainerDistribution(patchedcontainerContainerDistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerAPI.DistributionsContainerContainerPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsContainerContainerPartialUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerAPI.DistributionsContainerContainerPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedcontainerContainerDistribution** | [**PatchedcontainerContainerDistribution**](PatchedcontainerContainerDistribution.md) |  | 

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


## DistributionsContainerContainerRead

> ContainerContainerDistributionResponse DistributionsContainerContainerRead(ctx, containerContainerDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a container distribution



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
	containerContainerDistributionHref := "containerContainerDistributionHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerRead(context.Background(), containerContainerDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerAPI.DistributionsContainerContainerRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsContainerContainerRead`: ContainerContainerDistributionResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerAPI.DistributionsContainerContainerRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ContainerContainerDistributionResponse**](ContainerContainerDistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsContainerContainerRemoveRole

> NestedRoleResponse DistributionsContainerContainerRemoveRole(ctx, containerContainerDistributionHref).NestedRole(nestedRole).Execute()

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
	containerContainerDistributionHref := "containerContainerDistributionHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerRemoveRole(context.Background(), containerContainerDistributionHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerAPI.DistributionsContainerContainerRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsContainerContainerRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerAPI.DistributionsContainerContainerRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerRemoveRoleRequest struct via the builder pattern


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


## DistributionsContainerContainerSetLabel

> SetLabelResponse DistributionsContainerContainerSetLabel(ctx, containerContainerDistributionHref).SetLabel(setLabel).Execute()

Set a label



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
	containerContainerDistributionHref := "containerContainerDistributionHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerSetLabel(context.Background(), containerContainerDistributionHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerAPI.DistributionsContainerContainerSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsContainerContainerSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerAPI.DistributionsContainerContainerSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerSetLabelRequest struct via the builder pattern


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


## DistributionsContainerContainerUnsetLabel

> UnsetLabelResponse DistributionsContainerContainerUnsetLabel(ctx, containerContainerDistributionHref).UnsetLabel(unsetLabel).Execute()

Unset a label



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
	containerContainerDistributionHref := "containerContainerDistributionHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerUnsetLabel(context.Background(), containerContainerDistributionHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerAPI.DistributionsContainerContainerUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsContainerContainerUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerAPI.DistributionsContainerContainerUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerUnsetLabelRequest struct via the builder pattern


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


## DistributionsContainerContainerUpdate

> AsyncOperationResponse DistributionsContainerContainerUpdate(ctx, containerContainerDistributionHref).ContainerContainerDistribution(containerContainerDistribution).Execute()

Update a container distribution



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
	containerContainerDistributionHref := "containerContainerDistributionHref_example" // string | 
	containerContainerDistribution := *openapiclient.NewContainerContainerDistribution("BasePath_example", "Name_example") // ContainerContainerDistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerUpdate(context.Background(), containerContainerDistributionHref).ContainerContainerDistribution(containerContainerDistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsContainerAPI.DistributionsContainerContainerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsContainerContainerUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsContainerAPI.DistributionsContainerContainerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsContainerContainerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerContainerDistribution** | [**ContainerContainerDistribution**](ContainerContainerDistribution.md) |  | 

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

