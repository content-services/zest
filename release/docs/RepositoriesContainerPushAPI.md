# \RepositoriesContainerPushAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesContainerContainerPushAddRole**](RepositoriesContainerPushAPI.md#RepositoriesContainerContainerPushAddRole) | **Post** /{container_container_push_repository_href}add_role/ | Add a role
[**RepositoriesContainerContainerPushList**](RepositoriesContainerPushAPI.md#RepositoriesContainerContainerPushList) | **Get** /api/pulp/{pulp_domain}/api/v3/repositories/container/container-push/ | List container push repositorys
[**RepositoriesContainerContainerPushListRoles**](RepositoriesContainerPushAPI.md#RepositoriesContainerContainerPushListRoles) | **Get** /{container_container_push_repository_href}list_roles/ | List roles
[**RepositoriesContainerContainerPushMigrate**](RepositoriesContainerPushAPI.md#RepositoriesContainerContainerPushMigrate) | **Post** /{container_container_push_repository_href}migrate/ | Migrate push repository to container repository
[**RepositoriesContainerContainerPushMyPermissions**](RepositoriesContainerPushAPI.md#RepositoriesContainerContainerPushMyPermissions) | **Get** /{container_container_push_repository_href}my_permissions/ | List user permissions
[**RepositoriesContainerContainerPushPartialUpdate**](RepositoriesContainerPushAPI.md#RepositoriesContainerContainerPushPartialUpdate) | **Patch** /{container_container_push_repository_href} | Update a container push repository
[**RepositoriesContainerContainerPushRead**](RepositoriesContainerPushAPI.md#RepositoriesContainerContainerPushRead) | **Get** /{container_container_push_repository_href} | Inspect a container push repository
[**RepositoriesContainerContainerPushRemoveImage**](RepositoriesContainerPushAPI.md#RepositoriesContainerContainerPushRemoveImage) | **Post** /{container_container_push_repository_href}remove_image/ | Delete an image from a repository
[**RepositoriesContainerContainerPushRemoveRole**](RepositoriesContainerPushAPI.md#RepositoriesContainerContainerPushRemoveRole) | **Post** /{container_container_push_repository_href}remove_role/ | Remove a role
[**RepositoriesContainerContainerPushRemoveSignatures**](RepositoriesContainerPushAPI.md#RepositoriesContainerContainerPushRemoveSignatures) | **Post** /{container_container_push_repository_href}remove_signatures/ | 
[**RepositoriesContainerContainerPushSign**](RepositoriesContainerPushAPI.md#RepositoriesContainerContainerPushSign) | **Post** /{container_container_push_repository_href}sign/ | Sign images in the repo
[**RepositoriesContainerContainerPushTag**](RepositoriesContainerPushAPI.md#RepositoriesContainerContainerPushTag) | **Post** /{container_container_push_repository_href}tag/ | Create a Tag
[**RepositoriesContainerContainerPushUntag**](RepositoriesContainerPushAPI.md#RepositoriesContainerContainerPushUntag) | **Post** /{container_container_push_repository_href}untag/ | Delete a tag
[**RepositoriesContainerContainerPushUpdate**](RepositoriesContainerPushAPI.md#RepositoriesContainerContainerPushUpdate) | **Put** /{container_container_push_repository_href} | Update a container push repository



## RepositoriesContainerContainerPushAddRole

> NestedRoleResponse RepositoriesContainerContainerPushAddRole(ctx, containerContainerPushRepositoryHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	containerContainerPushRepositoryHref := "containerContainerPushRepositoryHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesContainerPushAPI.RepositoriesContainerContainerPushAddRole(context.Background(), containerContainerPushRepositoryHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesContainerContainerPushAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushAddRoleRequest struct via the builder pattern


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


## RepositoriesContainerContainerPushList

> PaginatedcontainerContainerPushRepositoryResponseList RepositoriesContainerContainerPushList(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).LatestWithContent(latestWithContent).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Remote(remote).RetainCheckpoints(retainCheckpoints).RetainCheckpointsGt(retainCheckpointsGt).RetainCheckpointsGte(retainCheckpointsGte).RetainCheckpointsIsnull(retainCheckpointsIsnull).RetainCheckpointsLt(retainCheckpointsLt).RetainCheckpointsLte(retainCheckpointsLte).RetainCheckpointsNe(retainCheckpointsNe).RetainCheckpointsRange(retainCheckpointsRange).RetainRepoVersions(retainRepoVersions).RetainRepoVersionsGt(retainRepoVersionsGt).RetainRepoVersionsGte(retainRepoVersionsGte).RetainRepoVersionsIsnull(retainRepoVersionsIsnull).RetainRepoVersionsLt(retainRepoVersionsLt).RetainRepoVersionsLte(retainRepoVersionsLte).RetainRepoVersionsNe(retainRepoVersionsNe).RetainRepoVersionsRange(retainRepoVersionsRange).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()

List container push repositorys



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
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `name` - Name* `-name` - Name (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `description` - Description* `-description` - Description (descending)* `next_version` - Next version* `-next_version` - Next version (descending)* `retain_repo_versions` - Retain repo versions* `-retain_repo_versions` - Retain repo versions (descending)* `retain_checkpoints` - Retain checkpoints* `-retain_checkpoints` - Retain checkpoints (descending)* `user_hidden` - User hidden* `-user_hidden` - User hidden (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	remote := "remote_example" // string |  (optional)
	retainCheckpoints := int32(56) // int32 | Filter results where retain_checkpoints matches value (optional)
	retainCheckpointsGt := int32(56) // int32 | Filter results where retain_checkpoints is greater than value (optional)
	retainCheckpointsGte := int32(56) // int32 | Filter results where retain_checkpoints is greater than or equal to value (optional)
	retainCheckpointsIsnull := true // bool | Filter results where retain_checkpoints has a null value (optional)
	retainCheckpointsLt := int32(56) // int32 | Filter results where retain_checkpoints is less than value (optional)
	retainCheckpointsLte := int32(56) // int32 | Filter results where retain_checkpoints is less than or equal to value (optional)
	retainCheckpointsNe := int32(56) // int32 | Filter results where retain_checkpoints not equal to value (optional)
	retainCheckpointsRange := []int32{int32(123)} // []int32 | Filter results where retain_checkpoints is between two comma separated values (optional)
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
	resp, r, err := apiClient.RepositoriesContainerPushAPI.RepositoriesContainerContainerPushList(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).LatestWithContent(latestWithContent).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Remote(remote).RetainCheckpoints(retainCheckpoints).RetainCheckpointsGt(retainCheckpointsGt).RetainCheckpointsGte(retainCheckpointsGte).RetainCheckpointsIsnull(retainCheckpointsIsnull).RetainCheckpointsLt(retainCheckpointsLt).RetainCheckpointsLte(retainCheckpointsLte).RetainCheckpointsNe(retainCheckpointsNe).RetainCheckpointsRange(retainCheckpointsRange).RetainRepoVersions(retainRepoVersions).RetainRepoVersionsGt(retainRepoVersionsGt).RetainRepoVersionsGte(retainRepoVersionsGte).RetainRepoVersionsIsnull(retainRepoVersionsIsnull).RetainRepoVersionsLt(retainRepoVersionsLt).RetainRepoVersionsLte(retainRepoVersionsLte).RetainRepoVersionsNe(retainRepoVersionsNe).RetainRepoVersionsRange(retainRepoVersionsRange).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesContainerContainerPushList`: PaginatedcontainerContainerPushRepositoryResponseList
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
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
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;description&#x60; - Description* &#x60;-description&#x60; - Description (descending)* &#x60;next_version&#x60; - Next version* &#x60;-next_version&#x60; - Next version (descending)* &#x60;retain_repo_versions&#x60; - Retain repo versions* &#x60;-retain_repo_versions&#x60; - Retain repo versions (descending)* &#x60;retain_checkpoints&#x60; - Retain checkpoints* &#x60;-retain_checkpoints&#x60; - Retain checkpoints (descending)* &#x60;user_hidden&#x60; - User hidden* &#x60;-user_hidden&#x60; - User hidden (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **remote** | **string** |  | 
 **retainCheckpoints** | **int32** | Filter results where retain_checkpoints matches value | 
 **retainCheckpointsGt** | **int32** | Filter results where retain_checkpoints is greater than value | 
 **retainCheckpointsGte** | **int32** | Filter results where retain_checkpoints is greater than or equal to value | 
 **retainCheckpointsIsnull** | **bool** | Filter results where retain_checkpoints has a null value | 
 **retainCheckpointsLt** | **int32** | Filter results where retain_checkpoints is less than value | 
 **retainCheckpointsLte** | **int32** | Filter results where retain_checkpoints is less than or equal to value | 
 **retainCheckpointsNe** | **int32** | Filter results where retain_checkpoints not equal to value | 
 **retainCheckpointsRange** | **[]int32** | Filter results where retain_checkpoints is between two comma separated values | 
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

[**PaginatedcontainerContainerPushRepositoryResponseList**](PaginatedcontainerContainerPushRepositoryResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesContainerContainerPushListRoles

> ObjectRolesResponse RepositoriesContainerContainerPushListRoles(ctx, containerContainerPushRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	containerContainerPushRepositoryHref := "containerContainerPushRepositoryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesContainerPushAPI.RepositoriesContainerContainerPushListRoles(context.Background(), containerContainerPushRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesContainerContainerPushListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushListRolesRequest struct via the builder pattern


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


## RepositoriesContainerContainerPushMigrate

> AsyncOperationResponse RepositoriesContainerContainerPushMigrate(ctx, containerContainerPushRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Execute()

Migrate push repository to container repository



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
	containerContainerPushRepositoryHref := "containerContainerPushRepositoryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesContainerPushAPI.RepositoriesContainerContainerPushMigrate(context.Background(), containerContainerPushRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushMigrate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesContainerContainerPushMigrate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushMigrate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushMigrateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## RepositoriesContainerContainerPushMyPermissions

> MyPermissionsResponse RepositoriesContainerContainerPushMyPermissions(ctx, containerContainerPushRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	containerContainerPushRepositoryHref := "containerContainerPushRepositoryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesContainerPushAPI.RepositoriesContainerContainerPushMyPermissions(context.Background(), containerContainerPushRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesContainerContainerPushMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushMyPermissionsRequest struct via the builder pattern


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


## RepositoriesContainerContainerPushPartialUpdate

> ContainerContainerPushRepositoryResponse RepositoriesContainerContainerPushPartialUpdate(ctx, containerContainerPushRepositoryHref).PatchedcontainerContainerPushRepository(patchedcontainerContainerPushRepository).XTaskDiagnostics(xTaskDiagnostics).Execute()

Update a container push repository



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
	containerContainerPushRepositoryHref := "containerContainerPushRepositoryHref_example" // string | 
	patchedcontainerContainerPushRepository := *openapiclient.NewPatchedcontainerContainerPushRepository() // PatchedcontainerContainerPushRepository | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesContainerPushAPI.RepositoriesContainerContainerPushPartialUpdate(context.Background(), containerContainerPushRepositoryHref).PatchedcontainerContainerPushRepository(patchedcontainerContainerPushRepository).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesContainerContainerPushPartialUpdate`: ContainerContainerPushRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedcontainerContainerPushRepository** | [**PatchedcontainerContainerPushRepository**](PatchedcontainerContainerPushRepository.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**ContainerContainerPushRepositoryResponse**](ContainerContainerPushRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesContainerContainerPushRead

> ContainerContainerPushRepositoryResponse RepositoriesContainerContainerPushRead(ctx, containerContainerPushRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a container push repository



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
	containerContainerPushRepositoryHref := "containerContainerPushRepositoryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesContainerPushAPI.RepositoriesContainerContainerPushRead(context.Background(), containerContainerPushRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesContainerContainerPushRead`: ContainerContainerPushRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ContainerContainerPushRepositoryResponse**](ContainerContainerPushRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesContainerContainerPushRemoveImage

> AsyncOperationResponse RepositoriesContainerContainerPushRemoveImage(ctx, containerContainerPushRepositoryHref).RemoveImage(removeImage).XTaskDiagnostics(xTaskDiagnostics).Execute()

Delete an image from a repository



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
	containerContainerPushRepositoryHref := "containerContainerPushRepositoryHref_example" // string | 
	removeImage := *openapiclient.NewRemoveImage("Digest_example") // RemoveImage | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesContainerPushAPI.RepositoriesContainerContainerPushRemoveImage(context.Background(), containerContainerPushRepositoryHref).RemoveImage(removeImage).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushRemoveImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesContainerContainerPushRemoveImage`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushRemoveImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushRemoveImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeImage** | [**RemoveImage**](RemoveImage.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## RepositoriesContainerContainerPushRemoveRole

> NestedRoleResponse RepositoriesContainerContainerPushRemoveRole(ctx, containerContainerPushRepositoryHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	containerContainerPushRepositoryHref := "containerContainerPushRepositoryHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesContainerPushAPI.RepositoriesContainerContainerPushRemoveRole(context.Background(), containerContainerPushRepositoryHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesContainerContainerPushRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushRemoveRoleRequest struct via the builder pattern


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


## RepositoriesContainerContainerPushRemoveSignatures

> RemoveSignaturesResponse RepositoriesContainerContainerPushRemoveSignatures(ctx, containerContainerPushRepositoryHref).RemoveSignatures(removeSignatures).XTaskDiagnostics(xTaskDiagnostics).Execute()





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
	containerContainerPushRepositoryHref := "containerContainerPushRepositoryHref_example" // string | 
	removeSignatures := *openapiclient.NewRemoveSignatures("SignedWithKeyId_example") // RemoveSignatures | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesContainerPushAPI.RepositoriesContainerContainerPushRemoveSignatures(context.Background(), containerContainerPushRepositoryHref).RemoveSignatures(removeSignatures).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushRemoveSignatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesContainerContainerPushRemoveSignatures`: RemoveSignaturesResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushRemoveSignatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushRemoveSignaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeSignatures** | [**RemoveSignatures**](RemoveSignatures.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**RemoveSignaturesResponse**](RemoveSignaturesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesContainerContainerPushSign

> AsyncOperationResponse RepositoriesContainerContainerPushSign(ctx, containerContainerPushRepositoryHref).RepositorySign(repositorySign).XTaskDiagnostics(xTaskDiagnostics).Execute()

Sign images in the repo



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
	containerContainerPushRepositoryHref := "containerContainerPushRepositoryHref_example" // string | 
	repositorySign := *openapiclient.NewRepositorySign() // RepositorySign | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesContainerPushAPI.RepositoriesContainerContainerPushSign(context.Background(), containerContainerPushRepositoryHref).RepositorySign(repositorySign).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushSign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesContainerContainerPushSign`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushSign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushSignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositorySign** | [**RepositorySign**](RepositorySign.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## RepositoriesContainerContainerPushTag

> AsyncOperationResponse RepositoriesContainerContainerPushTag(ctx, containerContainerPushRepositoryHref).TagImage(tagImage).XTaskDiagnostics(xTaskDiagnostics).Execute()

Create a Tag



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
	containerContainerPushRepositoryHref := "containerContainerPushRepositoryHref_example" // string | 
	tagImage := *openapiclient.NewTagImage("Tag_example", "Digest_example") // TagImage | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesContainerPushAPI.RepositoriesContainerContainerPushTag(context.Background(), containerContainerPushRepositoryHref).TagImage(tagImage).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesContainerContainerPushTag`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagImage** | [**TagImage**](TagImage.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## RepositoriesContainerContainerPushUntag

> AsyncOperationResponse RepositoriesContainerContainerPushUntag(ctx, containerContainerPushRepositoryHref).UnTagImage(unTagImage).XTaskDiagnostics(xTaskDiagnostics).Execute()

Delete a tag



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
	containerContainerPushRepositoryHref := "containerContainerPushRepositoryHref_example" // string | 
	unTagImage := *openapiclient.NewUnTagImage("Tag_example") // UnTagImage | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesContainerPushAPI.RepositoriesContainerContainerPushUntag(context.Background(), containerContainerPushRepositoryHref).UnTagImage(unTagImage).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushUntag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesContainerContainerPushUntag`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushUntag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushUntagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unTagImage** | [**UnTagImage**](UnTagImage.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## RepositoriesContainerContainerPushUpdate

> ContainerContainerPushRepositoryResponse RepositoriesContainerContainerPushUpdate(ctx, containerContainerPushRepositoryHref).ContainerContainerPushRepository(containerContainerPushRepository).XTaskDiagnostics(xTaskDiagnostics).Execute()

Update a container push repository



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
	containerContainerPushRepositoryHref := "containerContainerPushRepositoryHref_example" // string | 
	containerContainerPushRepository := *openapiclient.NewContainerContainerPushRepository("Name_example") // ContainerContainerPushRepository | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesContainerPushAPI.RepositoriesContainerContainerPushUpdate(context.Background(), containerContainerPushRepositoryHref).ContainerContainerPushRepository(containerContainerPushRepository).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesContainerContainerPushUpdate`: ContainerContainerPushRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesContainerPushAPI.RepositoriesContainerContainerPushUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPushRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerPushUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerContainerPushRepository** | [**ContainerContainerPushRepository**](ContainerContainerPushRepository.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**ContainerContainerPushRepositoryResponse**](ContainerContainerPushRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

