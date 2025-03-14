# \RepositoriesRpmAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesRpmRpmAddRole**](RepositoriesRpmAPI.md#RepositoriesRpmRpmAddRole) | **Post** /{rpm_rpm_repository_href}add_role/ | Add a role
[**RepositoriesRpmRpmCreate**](RepositoriesRpmAPI.md#RepositoriesRpmRpmCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/repositories/rpm/rpm/ | Create a rpm repository
[**RepositoriesRpmRpmDelete**](RepositoriesRpmAPI.md#RepositoriesRpmRpmDelete) | **Delete** /{rpm_rpm_repository_href} | Delete a rpm repository
[**RepositoriesRpmRpmList**](RepositoriesRpmAPI.md#RepositoriesRpmRpmList) | **Get** /api/pulp/{pulp_domain}/api/v3/repositories/rpm/rpm/ | List rpm repositorys
[**RepositoriesRpmRpmListRoles**](RepositoriesRpmAPI.md#RepositoriesRpmRpmListRoles) | **Get** /{rpm_rpm_repository_href}list_roles/ | List roles
[**RepositoriesRpmRpmModify**](RepositoriesRpmAPI.md#RepositoriesRpmRpmModify) | **Post** /{rpm_rpm_repository_href}modify/ | Modify Repository Content
[**RepositoriesRpmRpmMyPermissions**](RepositoriesRpmAPI.md#RepositoriesRpmRpmMyPermissions) | **Get** /{rpm_rpm_repository_href}my_permissions/ | List user permissions
[**RepositoriesRpmRpmPartialUpdate**](RepositoriesRpmAPI.md#RepositoriesRpmRpmPartialUpdate) | **Patch** /{rpm_rpm_repository_href} | Update a rpm repository
[**RepositoriesRpmRpmRead**](RepositoriesRpmAPI.md#RepositoriesRpmRpmRead) | **Get** /{rpm_rpm_repository_href} | Inspect a rpm repository
[**RepositoriesRpmRpmRemoveRole**](RepositoriesRpmAPI.md#RepositoriesRpmRpmRemoveRole) | **Post** /{rpm_rpm_repository_href}remove_role/ | Remove a role
[**RepositoriesRpmRpmSetLabel**](RepositoriesRpmAPI.md#RepositoriesRpmRpmSetLabel) | **Post** /{rpm_rpm_repository_href}set_label/ | Set a label
[**RepositoriesRpmRpmSync**](RepositoriesRpmAPI.md#RepositoriesRpmRpmSync) | **Post** /{rpm_rpm_repository_href}sync/ | Sync from remote
[**RepositoriesRpmRpmUnsetLabel**](RepositoriesRpmAPI.md#RepositoriesRpmRpmUnsetLabel) | **Post** /{rpm_rpm_repository_href}unset_label/ | Unset a label
[**RepositoriesRpmRpmUpdate**](RepositoriesRpmAPI.md#RepositoriesRpmRpmUpdate) | **Put** /{rpm_rpm_repository_href} | Update a rpm repository



## RepositoriesRpmRpmAddRole

> NestedRoleResponse RepositoriesRpmRpmAddRole(ctx, rpmRpmRepositoryHref).NestedRole(nestedRole).Execute()

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
	rpmRpmRepositoryHref := "rpmRpmRepositoryHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesRpmAPI.RepositoriesRpmRpmAddRole(context.Background(), rpmRpmRepositoryHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesRpmAPI.RepositoriesRpmRpmAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesRpmRpmAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesRpmAPI.RepositoriesRpmRpmAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesRpmRpmAddRoleRequest struct via the builder pattern


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


## RepositoriesRpmRpmCreate

> RpmRpmRepositoryResponse RepositoriesRpmRpmCreate(ctx, pulpDomain).RpmRpmRepository(rpmRpmRepository).Execute()

Create a rpm repository



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
	rpmRpmRepository := *openapiclient.NewRpmRpmRepository("Name_example") // RpmRpmRepository | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesRpmAPI.RepositoriesRpmRpmCreate(context.Background(), pulpDomain).RpmRpmRepository(rpmRpmRepository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesRpmAPI.RepositoriesRpmRpmCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesRpmRpmCreate`: RpmRpmRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesRpmAPI.RepositoriesRpmRpmCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesRpmRpmCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rpmRpmRepository** | [**RpmRpmRepository**](RpmRpmRepository.md) |  | 

### Return type

[**RpmRpmRepositoryResponse**](RpmRpmRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesRpmRpmDelete

> AsyncOperationResponse RepositoriesRpmRpmDelete(ctx, rpmRpmRepositoryHref).Execute()

Delete a rpm repository



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
	rpmRpmRepositoryHref := "rpmRpmRepositoryHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesRpmAPI.RepositoriesRpmRpmDelete(context.Background(), rpmRpmRepositoryHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesRpmAPI.RepositoriesRpmRpmDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesRpmRpmDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesRpmAPI.RepositoriesRpmRpmDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesRpmRpmDeleteRequest struct via the builder pattern


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


## RepositoriesRpmRpmList

> PaginatedrpmRpmRepositoryResponseList RepositoriesRpmRpmList(ctx, pulpDomain).LatestWithContent(latestWithContent).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Remote(remote).RetainRepoVersions(retainRepoVersions).RetainRepoVersionsGt(retainRepoVersionsGt).RetainRepoVersionsGte(retainRepoVersionsGte).RetainRepoVersionsIsnull(retainRepoVersionsIsnull).RetainRepoVersionsLt(retainRepoVersionsLt).RetainRepoVersionsLte(retainRepoVersionsLte).RetainRepoVersionsNe(retainRepoVersionsNe).RetainRepoVersionsRange(retainRepoVersionsRange).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()

List rpm repositorys



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
	latestWithContent := "latestWithContent_example" // string | Content Unit referenced by HREF/PRN (optional)
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
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `name` - Name* `-name` - Name (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `description` - Description* `-description` - Description (descending)* `next_version` - Next version* `-next_version` - Next version (descending)* `retain_repo_versions` - Retain repo versions* `-retain_repo_versions` - Retain repo versions (descending)* `user_hidden` - User hidden* `-user_hidden` - User hidden (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	remote := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Foreign Key referenced by HREF (optional)
	retainRepoVersions := int32(56) // int32 | Filter results where retain_repo_versions matches value (optional)
	retainRepoVersionsGt := int32(56) // int32 | Filter results where retain_repo_versions is greater than value (optional)
	retainRepoVersionsGte := int32(56) // int32 | Filter results where retain_repo_versions is greater than or equal to value (optional)
	retainRepoVersionsIsnull := true // bool | Filter results where retain_repo_versions has a null value (optional)
	retainRepoVersionsLt := int32(56) // int32 | Filter results where retain_repo_versions is less than value (optional)
	retainRepoVersionsLte := int32(56) // int32 | Filter results where retain_repo_versions is less than or equal to value (optional)
	retainRepoVersionsNe := int32(56) // int32 | Filter results where retain_repo_versions not equal to value (optional)
	retainRepoVersionsRange := []int32{int32(123)} // []int32 | Filter results where retain_repo_versions is between two comma separated values (optional)
	withContent := "withContent_example" // string | Content Unit referenced by HREF/PRN (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesRpmAPI.RepositoriesRpmRpmList(context.Background(), pulpDomain).LatestWithContent(latestWithContent).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Remote(remote).RetainRepoVersions(retainRepoVersions).RetainRepoVersionsGt(retainRepoVersionsGt).RetainRepoVersionsGte(retainRepoVersionsGte).RetainRepoVersionsIsnull(retainRepoVersionsIsnull).RetainRepoVersionsLt(retainRepoVersionsLt).RetainRepoVersionsLte(retainRepoVersionsLte).RetainRepoVersionsNe(retainRepoVersionsNe).RetainRepoVersionsRange(retainRepoVersionsRange).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesRpmAPI.RepositoriesRpmRpmList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesRpmRpmList`: PaginatedrpmRpmRepositoryResponseList
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesRpmAPI.RepositoriesRpmRpmList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesRpmRpmListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **latestWithContent** | **string** | Content Unit referenced by HREF/PRN | 
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
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;description&#x60; - Description* &#x60;-description&#x60; - Description (descending)* &#x60;next_version&#x60; - Next version* &#x60;-next_version&#x60; - Next version (descending)* &#x60;retain_repo_versions&#x60; - Retain repo versions* &#x60;-retain_repo_versions&#x60; - Retain repo versions (descending)* &#x60;user_hidden&#x60; - User hidden* &#x60;-user_hidden&#x60; - User hidden (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **remote** | **string** | Foreign Key referenced by HREF | 
 **retainRepoVersions** | **int32** | Filter results where retain_repo_versions matches value | 
 **retainRepoVersionsGt** | **int32** | Filter results where retain_repo_versions is greater than value | 
 **retainRepoVersionsGte** | **int32** | Filter results where retain_repo_versions is greater than or equal to value | 
 **retainRepoVersionsIsnull** | **bool** | Filter results where retain_repo_versions has a null value | 
 **retainRepoVersionsLt** | **int32** | Filter results where retain_repo_versions is less than value | 
 **retainRepoVersionsLte** | **int32** | Filter results where retain_repo_versions is less than or equal to value | 
 **retainRepoVersionsNe** | **int32** | Filter results where retain_repo_versions not equal to value | 
 **retainRepoVersionsRange** | **[]int32** | Filter results where retain_repo_versions is between two comma separated values | 
 **withContent** | **string** | Content Unit referenced by HREF/PRN | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedrpmRpmRepositoryResponseList**](PaginatedrpmRpmRepositoryResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesRpmRpmListRoles

> ObjectRolesResponse RepositoriesRpmRpmListRoles(ctx, rpmRpmRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	rpmRpmRepositoryHref := "rpmRpmRepositoryHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesRpmAPI.RepositoriesRpmRpmListRoles(context.Background(), rpmRpmRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesRpmAPI.RepositoriesRpmRpmListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesRpmRpmListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesRpmAPI.RepositoriesRpmRpmListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesRpmRpmListRolesRequest struct via the builder pattern


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


## RepositoriesRpmRpmModify

> AsyncOperationResponse RepositoriesRpmRpmModify(ctx, rpmRpmRepositoryHref).RepositoryAddRemoveContent(repositoryAddRemoveContent).Execute()

Modify Repository Content



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
	rpmRpmRepositoryHref := "rpmRpmRepositoryHref_example" // string | 
	repositoryAddRemoveContent := *openapiclient.NewRepositoryAddRemoveContent() // RepositoryAddRemoveContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesRpmAPI.RepositoriesRpmRpmModify(context.Background(), rpmRpmRepositoryHref).RepositoryAddRemoveContent(repositoryAddRemoveContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesRpmAPI.RepositoriesRpmRpmModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesRpmRpmModify`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesRpmAPI.RepositoriesRpmRpmModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesRpmRpmModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositoryAddRemoveContent** | [**RepositoryAddRemoveContent**](RepositoryAddRemoveContent.md) |  | 

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


## RepositoriesRpmRpmMyPermissions

> MyPermissionsResponse RepositoriesRpmRpmMyPermissions(ctx, rpmRpmRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	rpmRpmRepositoryHref := "rpmRpmRepositoryHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesRpmAPI.RepositoriesRpmRpmMyPermissions(context.Background(), rpmRpmRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesRpmAPI.RepositoriesRpmRpmMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesRpmRpmMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesRpmAPI.RepositoriesRpmRpmMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesRpmRpmMyPermissionsRequest struct via the builder pattern


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


## RepositoriesRpmRpmPartialUpdate

> AsyncOperationResponse RepositoriesRpmRpmPartialUpdate(ctx, rpmRpmRepositoryHref).PatchedrpmRpmRepository(patchedrpmRpmRepository).Execute()

Update a rpm repository



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
	rpmRpmRepositoryHref := "rpmRpmRepositoryHref_example" // string | 
	patchedrpmRpmRepository := *openapiclient.NewPatchedrpmRpmRepository() // PatchedrpmRpmRepository | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesRpmAPI.RepositoriesRpmRpmPartialUpdate(context.Background(), rpmRpmRepositoryHref).PatchedrpmRpmRepository(patchedrpmRpmRepository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesRpmAPI.RepositoriesRpmRpmPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesRpmRpmPartialUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesRpmAPI.RepositoriesRpmRpmPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesRpmRpmPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedrpmRpmRepository** | [**PatchedrpmRpmRepository**](PatchedrpmRpmRepository.md) |  | 

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


## RepositoriesRpmRpmRead

> RpmRpmRepositoryResponse RepositoriesRpmRpmRead(ctx, rpmRpmRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a rpm repository



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
	rpmRpmRepositoryHref := "rpmRpmRepositoryHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesRpmAPI.RepositoriesRpmRpmRead(context.Background(), rpmRpmRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesRpmAPI.RepositoriesRpmRpmRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesRpmRpmRead`: RpmRpmRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesRpmAPI.RepositoriesRpmRpmRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesRpmRpmReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RpmRpmRepositoryResponse**](RpmRpmRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesRpmRpmRemoveRole

> NestedRoleResponse RepositoriesRpmRpmRemoveRole(ctx, rpmRpmRepositoryHref).NestedRole(nestedRole).Execute()

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
	rpmRpmRepositoryHref := "rpmRpmRepositoryHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesRpmAPI.RepositoriesRpmRpmRemoveRole(context.Background(), rpmRpmRepositoryHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesRpmAPI.RepositoriesRpmRpmRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesRpmRpmRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesRpmAPI.RepositoriesRpmRpmRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesRpmRpmRemoveRoleRequest struct via the builder pattern


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


## RepositoriesRpmRpmSetLabel

> SetLabelResponse RepositoriesRpmRpmSetLabel(ctx, rpmRpmRepositoryHref).SetLabel(setLabel).Execute()

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
	rpmRpmRepositoryHref := "rpmRpmRepositoryHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesRpmAPI.RepositoriesRpmRpmSetLabel(context.Background(), rpmRpmRepositoryHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesRpmAPI.RepositoriesRpmRpmSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesRpmRpmSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesRpmAPI.RepositoriesRpmRpmSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesRpmRpmSetLabelRequest struct via the builder pattern


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


## RepositoriesRpmRpmSync

> AsyncOperationResponse RepositoriesRpmRpmSync(ctx, rpmRpmRepositoryHref).RpmRepositorySyncURL(rpmRepositorySyncURL).Execute()

Sync from remote



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
	rpmRpmRepositoryHref := "rpmRpmRepositoryHref_example" // string | 
	rpmRepositorySyncURL := *openapiclient.NewRpmRepositorySyncURL() // RpmRepositorySyncURL | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesRpmAPI.RepositoriesRpmRpmSync(context.Background(), rpmRpmRepositoryHref).RpmRepositorySyncURL(rpmRepositorySyncURL).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesRpmAPI.RepositoriesRpmRpmSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesRpmRpmSync`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesRpmAPI.RepositoriesRpmRpmSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesRpmRpmSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rpmRepositorySyncURL** | [**RpmRepositorySyncURL**](RpmRepositorySyncURL.md) |  | 

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


## RepositoriesRpmRpmUnsetLabel

> UnsetLabelResponse RepositoriesRpmRpmUnsetLabel(ctx, rpmRpmRepositoryHref).UnsetLabel(unsetLabel).Execute()

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
	rpmRpmRepositoryHref := "rpmRpmRepositoryHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesRpmAPI.RepositoriesRpmRpmUnsetLabel(context.Background(), rpmRpmRepositoryHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesRpmAPI.RepositoriesRpmRpmUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesRpmRpmUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesRpmAPI.RepositoriesRpmRpmUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesRpmRpmUnsetLabelRequest struct via the builder pattern


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


## RepositoriesRpmRpmUpdate

> AsyncOperationResponse RepositoriesRpmRpmUpdate(ctx, rpmRpmRepositoryHref).RpmRpmRepository(rpmRpmRepository).Execute()

Update a rpm repository



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
	rpmRpmRepositoryHref := "rpmRpmRepositoryHref_example" // string | 
	rpmRpmRepository := *openapiclient.NewRpmRpmRepository("Name_example") // RpmRpmRepository | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesRpmAPI.RepositoriesRpmRpmUpdate(context.Background(), rpmRpmRepositoryHref).RpmRpmRepository(rpmRpmRepository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesRpmAPI.RepositoriesRpmRpmUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesRpmRpmUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesRpmAPI.RepositoriesRpmRpmUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesRpmRpmUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rpmRpmRepository** | [**RpmRpmRepository**](RpmRpmRepository.md) |  | 

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

