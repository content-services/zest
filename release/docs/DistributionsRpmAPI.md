# \DistributionsRpmAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DistributionsRpmRpmAddRole**](DistributionsRpmAPI.md#DistributionsRpmRpmAddRole) | **Post** /{rpm_rpm_distribution_href}add_role/ | Add a role
[**DistributionsRpmRpmCreate**](DistributionsRpmAPI.md#DistributionsRpmRpmCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/distributions/rpm/rpm/ | Create a rpm distribution
[**DistributionsRpmRpmDelete**](DistributionsRpmAPI.md#DistributionsRpmRpmDelete) | **Delete** /{rpm_rpm_distribution_href} | Delete a rpm distribution
[**DistributionsRpmRpmList**](DistributionsRpmAPI.md#DistributionsRpmRpmList) | **Get** /api/pulp/{pulp_domain}/api/v3/distributions/rpm/rpm/ | List rpm distributions
[**DistributionsRpmRpmListRoles**](DistributionsRpmAPI.md#DistributionsRpmRpmListRoles) | **Get** /{rpm_rpm_distribution_href}list_roles/ | List roles
[**DistributionsRpmRpmMyPermissions**](DistributionsRpmAPI.md#DistributionsRpmRpmMyPermissions) | **Get** /{rpm_rpm_distribution_href}my_permissions/ | List user permissions
[**DistributionsRpmRpmPartialUpdate**](DistributionsRpmAPI.md#DistributionsRpmRpmPartialUpdate) | **Patch** /{rpm_rpm_distribution_href} | Update a rpm distribution
[**DistributionsRpmRpmRead**](DistributionsRpmAPI.md#DistributionsRpmRpmRead) | **Get** /{rpm_rpm_distribution_href} | Inspect a rpm distribution
[**DistributionsRpmRpmRemoveRole**](DistributionsRpmAPI.md#DistributionsRpmRpmRemoveRole) | **Post** /{rpm_rpm_distribution_href}remove_role/ | Remove a role
[**DistributionsRpmRpmSetLabel**](DistributionsRpmAPI.md#DistributionsRpmRpmSetLabel) | **Post** /{rpm_rpm_distribution_href}set_label/ | Set a label
[**DistributionsRpmRpmUnsetLabel**](DistributionsRpmAPI.md#DistributionsRpmRpmUnsetLabel) | **Post** /{rpm_rpm_distribution_href}unset_label/ | Unset a label
[**DistributionsRpmRpmUpdate**](DistributionsRpmAPI.md#DistributionsRpmRpmUpdate) | **Put** /{rpm_rpm_distribution_href} | Update a rpm distribution



## DistributionsRpmRpmAddRole

> NestedRoleResponse DistributionsRpmRpmAddRole(ctx, rpmRpmDistributionHref).NestedRole(nestedRole).Execute()

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
	rpmRpmDistributionHref := "rpmRpmDistributionHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsRpmAPI.DistributionsRpmRpmAddRole(context.Background(), rpmRpmDistributionHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsRpmAPI.DistributionsRpmRpmAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsRpmRpmAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsRpmAPI.DistributionsRpmRpmAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsRpmRpmAddRoleRequest struct via the builder pattern


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


## DistributionsRpmRpmCreate

> AsyncOperationResponse DistributionsRpmRpmCreate(ctx, pulpDomain).RpmRpmDistribution(rpmRpmDistribution).Execute()

Create a rpm distribution



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
	rpmRpmDistribution := *openapiclient.NewRpmRpmDistribution("BasePath_example", "Name_example") // RpmRpmDistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsRpmAPI.DistributionsRpmRpmCreate(context.Background(), pulpDomain).RpmRpmDistribution(rpmRpmDistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsRpmAPI.DistributionsRpmRpmCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsRpmRpmCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsRpmAPI.DistributionsRpmRpmCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsRpmRpmCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rpmRpmDistribution** | [**RpmRpmDistribution**](RpmRpmDistribution.md) |  | 

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


## DistributionsRpmRpmDelete

> AsyncOperationResponse DistributionsRpmRpmDelete(ctx, rpmRpmDistributionHref).Execute()

Delete a rpm distribution



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
	rpmRpmDistributionHref := "rpmRpmDistributionHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsRpmAPI.DistributionsRpmRpmDelete(context.Background(), rpmRpmDistributionHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsRpmAPI.DistributionsRpmRpmDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsRpmRpmDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsRpmAPI.DistributionsRpmRpmDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsRpmRpmDeleteRequest struct via the builder pattern


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


## DistributionsRpmRpmList

> PaginatedrpmRpmDistributionResponseList DistributionsRpmRpmList(ctx, pulpDomain).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Checkpoint(checkpoint).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Repository(repository).RepositoryIn(repositoryIn).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()

List rpm distributions



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
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `name` - Name* `-name` - Name (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `base_path` - Base path* `-base_path` - Base path (descending)* `hidden` - Hidden* `-hidden` - Hidden (descending)* `checkpoint` - Checkpoint* `-checkpoint` - Checkpoint (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
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
	resp, r, err := apiClient.DistributionsRpmAPI.DistributionsRpmRpmList(context.Background(), pulpDomain).BasePath(basePath).BasePathContains(basePathContains).BasePathIcontains(basePathIcontains).BasePathIn(basePathIn).Checkpoint(checkpoint).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Repository(repository).RepositoryIn(repositoryIn).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsRpmAPI.DistributionsRpmRpmList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsRpmRpmList`: PaginatedrpmRpmDistributionResponseList
	fmt.Fprintf(os.Stdout, "Response from `DistributionsRpmAPI.DistributionsRpmRpmList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsRpmRpmListRequest struct via the builder pattern


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
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;base_path&#x60; - Base path* &#x60;-base_path&#x60; - Base path (descending)* &#x60;hidden&#x60; - Hidden* &#x60;-hidden&#x60; - Hidden (descending)* &#x60;checkpoint&#x60; - Checkpoint* &#x60;-checkpoint&#x60; - Checkpoint (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
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

[**PaginatedrpmRpmDistributionResponseList**](PaginatedrpmRpmDistributionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsRpmRpmListRoles

> ObjectRolesResponse DistributionsRpmRpmListRoles(ctx, rpmRpmDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	rpmRpmDistributionHref := "rpmRpmDistributionHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsRpmAPI.DistributionsRpmRpmListRoles(context.Background(), rpmRpmDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsRpmAPI.DistributionsRpmRpmListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsRpmRpmListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsRpmAPI.DistributionsRpmRpmListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsRpmRpmListRolesRequest struct via the builder pattern


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


## DistributionsRpmRpmMyPermissions

> MyPermissionsResponse DistributionsRpmRpmMyPermissions(ctx, rpmRpmDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	rpmRpmDistributionHref := "rpmRpmDistributionHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsRpmAPI.DistributionsRpmRpmMyPermissions(context.Background(), rpmRpmDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsRpmAPI.DistributionsRpmRpmMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsRpmRpmMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsRpmAPI.DistributionsRpmRpmMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsRpmRpmMyPermissionsRequest struct via the builder pattern


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


## DistributionsRpmRpmPartialUpdate

> AsyncOperationResponse DistributionsRpmRpmPartialUpdate(ctx, rpmRpmDistributionHref).PatchedrpmRpmDistribution(patchedrpmRpmDistribution).Execute()

Update a rpm distribution



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
	rpmRpmDistributionHref := "rpmRpmDistributionHref_example" // string | 
	patchedrpmRpmDistribution := *openapiclient.NewPatchedrpmRpmDistribution() // PatchedrpmRpmDistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsRpmAPI.DistributionsRpmRpmPartialUpdate(context.Background(), rpmRpmDistributionHref).PatchedrpmRpmDistribution(patchedrpmRpmDistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsRpmAPI.DistributionsRpmRpmPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsRpmRpmPartialUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsRpmAPI.DistributionsRpmRpmPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsRpmRpmPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedrpmRpmDistribution** | [**PatchedrpmRpmDistribution**](PatchedrpmRpmDistribution.md) |  | 

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


## DistributionsRpmRpmRead

> RpmRpmDistributionResponse DistributionsRpmRpmRead(ctx, rpmRpmDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a rpm distribution



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
	rpmRpmDistributionHref := "rpmRpmDistributionHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsRpmAPI.DistributionsRpmRpmRead(context.Background(), rpmRpmDistributionHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsRpmAPI.DistributionsRpmRpmRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsRpmRpmRead`: RpmRpmDistributionResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsRpmAPI.DistributionsRpmRpmRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsRpmRpmReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RpmRpmDistributionResponse**](RpmRpmDistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistributionsRpmRpmRemoveRole

> NestedRoleResponse DistributionsRpmRpmRemoveRole(ctx, rpmRpmDistributionHref).NestedRole(nestedRole).Execute()

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
	rpmRpmDistributionHref := "rpmRpmDistributionHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsRpmAPI.DistributionsRpmRpmRemoveRole(context.Background(), rpmRpmDistributionHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsRpmAPI.DistributionsRpmRpmRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsRpmRpmRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsRpmAPI.DistributionsRpmRpmRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsRpmRpmRemoveRoleRequest struct via the builder pattern


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


## DistributionsRpmRpmSetLabel

> SetLabelResponse DistributionsRpmRpmSetLabel(ctx, rpmRpmDistributionHref).SetLabel(setLabel).Execute()

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
	rpmRpmDistributionHref := "rpmRpmDistributionHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsRpmAPI.DistributionsRpmRpmSetLabel(context.Background(), rpmRpmDistributionHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsRpmAPI.DistributionsRpmRpmSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsRpmRpmSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsRpmAPI.DistributionsRpmRpmSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsRpmRpmSetLabelRequest struct via the builder pattern


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


## DistributionsRpmRpmUnsetLabel

> UnsetLabelResponse DistributionsRpmRpmUnsetLabel(ctx, rpmRpmDistributionHref).UnsetLabel(unsetLabel).Execute()

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
	rpmRpmDistributionHref := "rpmRpmDistributionHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsRpmAPI.DistributionsRpmRpmUnsetLabel(context.Background(), rpmRpmDistributionHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsRpmAPI.DistributionsRpmRpmUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsRpmRpmUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsRpmAPI.DistributionsRpmRpmUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsRpmRpmUnsetLabelRequest struct via the builder pattern


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


## DistributionsRpmRpmUpdate

> AsyncOperationResponse DistributionsRpmRpmUpdate(ctx, rpmRpmDistributionHref).RpmRpmDistribution(rpmRpmDistribution).Execute()

Update a rpm distribution



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
	rpmRpmDistributionHref := "rpmRpmDistributionHref_example" // string | 
	rpmRpmDistribution := *openapiclient.NewRpmRpmDistribution("BasePath_example", "Name_example") // RpmRpmDistribution | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionsRpmAPI.DistributionsRpmRpmUpdate(context.Background(), rpmRpmDistributionHref).RpmRpmDistribution(rpmRpmDistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionsRpmAPI.DistributionsRpmRpmUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistributionsRpmRpmUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionsRpmAPI.DistributionsRpmRpmUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmDistributionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistributionsRpmRpmUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rpmRpmDistribution** | [**RpmRpmDistribution**](RpmRpmDistribution.md) |  | 

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

