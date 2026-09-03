# \RepositoriesMavenAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesMavenMavenAddCachedContent**](RepositoriesMavenAPI.md#RepositoriesMavenMavenAddCachedContent) | **Post** /{maven_maven_repository_href}add_cached_content/ | Add cached content
[**RepositoriesMavenMavenAddRole**](RepositoriesMavenAPI.md#RepositoriesMavenMavenAddRole) | **Post** /{maven_maven_repository_href}add_role/ | Add a role
[**RepositoriesMavenMavenCreate**](RepositoriesMavenAPI.md#RepositoriesMavenMavenCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/repositories/maven/maven/ | Create a maven repository
[**RepositoriesMavenMavenDelete**](RepositoriesMavenAPI.md#RepositoriesMavenMavenDelete) | **Delete** /{maven_maven_repository_href} | Delete a maven repository
[**RepositoriesMavenMavenList**](RepositoriesMavenAPI.md#RepositoriesMavenMavenList) | **Get** /api/pulp/{pulp_domain}/api/v3/repositories/maven/maven/ | List maven repositorys
[**RepositoriesMavenMavenListRoles**](RepositoriesMavenAPI.md#RepositoriesMavenMavenListRoles) | **Get** /{maven_maven_repository_href}list_roles/ | List roles
[**RepositoriesMavenMavenMetrics**](RepositoriesMavenAPI.md#RepositoriesMavenMavenMetrics) | **Get** /{maven_maven_repository_href}metrics/ | Repository metrics
[**RepositoriesMavenMavenModify**](RepositoriesMavenAPI.md#RepositoriesMavenMavenModify) | **Post** /{maven_maven_repository_href}modify/ | Modify Repository Content
[**RepositoriesMavenMavenMyPermissions**](RepositoriesMavenAPI.md#RepositoriesMavenMavenMyPermissions) | **Get** /{maven_maven_repository_href}my_permissions/ | List user permissions
[**RepositoriesMavenMavenPackages**](RepositoriesMavenAPI.md#RepositoriesMavenMavenPackages) | **Get** /{maven_maven_repository_href}packages/ | List packages
[**RepositoriesMavenMavenPartialUpdate**](RepositoriesMavenAPI.md#RepositoriesMavenMavenPartialUpdate) | **Patch** /{maven_maven_repository_href} | Update a maven repository
[**RepositoriesMavenMavenRead**](RepositoriesMavenAPI.md#RepositoriesMavenMavenRead) | **Get** /{maven_maven_repository_href} | Inspect a maven repository
[**RepositoriesMavenMavenRemoveRole**](RepositoriesMavenAPI.md#RepositoriesMavenMavenRemoveRole) | **Post** /{maven_maven_repository_href}remove_role/ | Remove a role
[**RepositoriesMavenMavenRepairMetadata**](RepositoriesMavenAPI.md#RepositoriesMavenMavenRepairMetadata) | **Post** /{maven_maven_repository_href}repair_metadata/ | Repair metadata
[**RepositoriesMavenMavenSetLabel**](RepositoriesMavenAPI.md#RepositoriesMavenMavenSetLabel) | **Post** /{maven_maven_repository_href}set_label/ | Set a label
[**RepositoriesMavenMavenUnsetLabel**](RepositoriesMavenAPI.md#RepositoriesMavenMavenUnsetLabel) | **Post** /{maven_maven_repository_href}unset_label/ | Unset a label
[**RepositoriesMavenMavenUpdate**](RepositoriesMavenAPI.md#RepositoriesMavenMavenUpdate) | **Put** /{maven_maven_repository_href} | Update a maven repository



## RepositoriesMavenMavenAddCachedContent

> AsyncOperationResponse RepositoriesMavenMavenAddCachedContent(ctx, mavenMavenRepositoryHref).RepositoryAddCachedContent(repositoryAddCachedContent).XTaskDiagnostics(xTaskDiagnostics).Execute()

Add cached content



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
	mavenMavenRepositoryHref := "mavenMavenRepositoryHref_example" // string | 
	repositoryAddCachedContent := *openapiclient.NewRepositoryAddCachedContent() // RepositoryAddCachedContent | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenAddCachedContent(context.Background(), mavenMavenRepositoryHref).RepositoryAddCachedContent(repositoryAddCachedContent).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenAddCachedContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenAddCachedContent`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenAddCachedContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenAddCachedContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositoryAddCachedContent** | [**RepositoryAddCachedContent**](RepositoryAddCachedContent.md) |  | 
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


## RepositoriesMavenMavenAddRole

> NestedRoleResponse RepositoriesMavenMavenAddRole(ctx, mavenMavenRepositoryHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	mavenMavenRepositoryHref := "mavenMavenRepositoryHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenAddRole(context.Background(), mavenMavenRepositoryHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenAddRoleRequest struct via the builder pattern


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


## RepositoriesMavenMavenCreate

> MavenMavenRepositoryResponse RepositoriesMavenMavenCreate(ctx, pulpDomain).MavenMavenRepository(mavenMavenRepository).XTaskDiagnostics(xTaskDiagnostics).Execute()

Create a maven repository



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
	mavenMavenRepository := *openapiclient.NewMavenMavenRepository("Name_example") // MavenMavenRepository | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenCreate(context.Background(), pulpDomain).MavenMavenRepository(mavenMavenRepository).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenCreate`: MavenMavenRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mavenMavenRepository** | [**MavenMavenRepository**](MavenMavenRepository.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**MavenMavenRepositoryResponse**](MavenMavenRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesMavenMavenDelete

> AsyncOperationResponse RepositoriesMavenMavenDelete(ctx, mavenMavenRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Execute()

Delete a maven repository



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
	mavenMavenRepositoryHref := "mavenMavenRepositoryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenDelete(context.Background(), mavenMavenRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenDeleteRequest struct via the builder pattern


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


## RepositoriesMavenMavenList

> PaginatedmavenMavenRepositoryResponseList RepositoriesMavenMavenList(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).LatestWithContent(latestWithContent).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Remote(remote).RetainCheckpoints(retainCheckpoints).RetainCheckpointsGt(retainCheckpointsGt).RetainCheckpointsGte(retainCheckpointsGte).RetainCheckpointsIsnull(retainCheckpointsIsnull).RetainCheckpointsLt(retainCheckpointsLt).RetainCheckpointsLte(retainCheckpointsLte).RetainCheckpointsNe(retainCheckpointsNe).RetainCheckpointsRange(retainCheckpointsRange).RetainRepoVersions(retainRepoVersions).RetainRepoVersionsGt(retainRepoVersionsGt).RetainRepoVersionsGte(retainRepoVersionsGte).RetainRepoVersionsIsnull(retainRepoVersionsIsnull).RetainRepoVersionsLt(retainRepoVersionsLt).RetainRepoVersionsLte(retainRepoVersionsLte).RetainRepoVersionsNe(retainRepoVersionsNe).RetainRepoVersionsRange(retainRepoVersionsRange).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()

List maven repositorys



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
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenList(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).LatestWithContent(latestWithContent).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Remote(remote).RetainCheckpoints(retainCheckpoints).RetainCheckpointsGt(retainCheckpointsGt).RetainCheckpointsGte(retainCheckpointsGte).RetainCheckpointsIsnull(retainCheckpointsIsnull).RetainCheckpointsLt(retainCheckpointsLt).RetainCheckpointsLte(retainCheckpointsLte).RetainCheckpointsNe(retainCheckpointsNe).RetainCheckpointsRange(retainCheckpointsRange).RetainRepoVersions(retainRepoVersions).RetainRepoVersionsGt(retainRepoVersionsGt).RetainRepoVersionsGte(retainRepoVersionsGte).RetainRepoVersionsIsnull(retainRepoVersionsIsnull).RetainRepoVersionsLt(retainRepoVersionsLt).RetainRepoVersionsLte(retainRepoVersionsLte).RetainRepoVersionsNe(retainRepoVersionsNe).RetainRepoVersionsRange(retainRepoVersionsRange).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenList`: PaginatedmavenMavenRepositoryResponseList
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenListRequest struct via the builder pattern


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

[**PaginatedmavenMavenRepositoryResponseList**](PaginatedmavenMavenRepositoryResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesMavenMavenListRoles

> ObjectRolesResponse RepositoriesMavenMavenListRoles(ctx, mavenMavenRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	mavenMavenRepositoryHref := "mavenMavenRepositoryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenListRoles(context.Background(), mavenMavenRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenListRolesRequest struct via the builder pattern


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


## RepositoriesMavenMavenMetrics

> MavenRepositoryMetricsResponse RepositoriesMavenMavenMetrics(ctx, mavenMavenRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).RepositoryVersion(repositoryVersion).Fields(fields).ExcludeFields(excludeFields).Execute()

Repository metrics



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
	mavenMavenRepositoryHref := "mavenMavenRepositoryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	repositoryVersion := "repositoryVersion_example" // string | HREF or PRN of a version of this repository. Defaults to the latest complete version. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenMetrics(context.Background(), mavenMavenRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).RepositoryVersion(repositoryVersion).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenMetrics`: MavenRepositoryMetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **repositoryVersion** | **string** | HREF or PRN of a version of this repository. Defaults to the latest complete version. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**MavenRepositoryMetricsResponse**](MavenRepositoryMetricsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesMavenMavenModify

> AsyncOperationResponse RepositoriesMavenMavenModify(ctx, mavenMavenRepositoryHref).RepositoryAddRemoveContent(repositoryAddRemoveContent).XTaskDiagnostics(xTaskDiagnostics).Execute()

Modify Repository Content



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
	mavenMavenRepositoryHref := "mavenMavenRepositoryHref_example" // string | 
	repositoryAddRemoveContent := *openapiclient.NewRepositoryAddRemoveContent() // RepositoryAddRemoveContent | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenModify(context.Background(), mavenMavenRepositoryHref).RepositoryAddRemoveContent(repositoryAddRemoveContent).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenModify`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositoryAddRemoveContent** | [**RepositoryAddRemoveContent**](RepositoryAddRemoveContent.md) |  | 
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


## RepositoriesMavenMavenMyPermissions

> MyPermissionsResponse RepositoriesMavenMavenMyPermissions(ctx, mavenMavenRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	mavenMavenRepositoryHref := "mavenMavenRepositoryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenMyPermissions(context.Background(), mavenMavenRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenMyPermissionsRequest struct via the builder pattern


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


## RepositoriesMavenMavenPackages

> PaginatedMavenRepositoryPackageListResponse RepositoriesMavenMavenPackages(ctx, mavenMavenRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).ArtifactIdIstartswith(artifactIdIstartswith).GroupIdIstartswith(groupIdIstartswith).RepositoryVersion(repositoryVersion).Fields(fields).ExcludeFields(excludeFields).Execute()

List packages



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
	mavenMavenRepositoryHref := "mavenMavenRepositoryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	artifactIdIstartswith := "artifactIdIstartswith_example" // string | Case-insensitive prefix on artifact_id. (optional)
	groupIdIstartswith := "groupIdIstartswith_example" // string | Case-insensitive prefix on group_id. (optional)
	repositoryVersion := "repositoryVersion_example" // string | HREF or PRN of a version of this repository. Defaults to the latest complete version. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenPackages(context.Background(), mavenMavenRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).ArtifactIdIstartswith(artifactIdIstartswith).GroupIdIstartswith(groupIdIstartswith).RepositoryVersion(repositoryVersion).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenPackages`: PaginatedMavenRepositoryPackageListResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenPackages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **artifactIdIstartswith** | **string** | Case-insensitive prefix on artifact_id. | 
 **groupIdIstartswith** | **string** | Case-insensitive prefix on group_id. | 
 **repositoryVersion** | **string** | HREF or PRN of a version of this repository. Defaults to the latest complete version. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedMavenRepositoryPackageListResponse**](PaginatedMavenRepositoryPackageListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesMavenMavenPartialUpdate

> MavenMavenRepositoryResponse RepositoriesMavenMavenPartialUpdate(ctx, mavenMavenRepositoryHref).PatchedmavenMavenRepository(patchedmavenMavenRepository).XTaskDiagnostics(xTaskDiagnostics).Execute()

Update a maven repository



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
	mavenMavenRepositoryHref := "mavenMavenRepositoryHref_example" // string | 
	patchedmavenMavenRepository := *openapiclient.NewPatchedmavenMavenRepository() // PatchedmavenMavenRepository | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenPartialUpdate(context.Background(), mavenMavenRepositoryHref).PatchedmavenMavenRepository(patchedmavenMavenRepository).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenPartialUpdate`: MavenMavenRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedmavenMavenRepository** | [**PatchedmavenMavenRepository**](PatchedmavenMavenRepository.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**MavenMavenRepositoryResponse**](MavenMavenRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesMavenMavenRead

> MavenMavenRepositoryResponse RepositoriesMavenMavenRead(ctx, mavenMavenRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a maven repository



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
	mavenMavenRepositoryHref := "mavenMavenRepositoryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenRead(context.Background(), mavenMavenRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenRead`: MavenMavenRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**MavenMavenRepositoryResponse**](MavenMavenRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesMavenMavenRemoveRole

> NestedRoleResponse RepositoriesMavenMavenRemoveRole(ctx, mavenMavenRepositoryHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	mavenMavenRepositoryHref := "mavenMavenRepositoryHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenRemoveRole(context.Background(), mavenMavenRepositoryHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenRemoveRoleRequest struct via the builder pattern


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


## RepositoriesMavenMavenRepairMetadata

> AsyncOperationResponse RepositoriesMavenMavenRepairMetadata(ctx, mavenMavenRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Execute()

Repair metadata



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
	mavenMavenRepositoryHref := "mavenMavenRepositoryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenRepairMetadata(context.Background(), mavenMavenRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenRepairMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenRepairMetadata`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenRepairMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenRepairMetadataRequest struct via the builder pattern


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


## RepositoriesMavenMavenSetLabel

> SetLabelResponse RepositoriesMavenMavenSetLabel(ctx, mavenMavenRepositoryHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

Set a label



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
	mavenMavenRepositoryHref := "mavenMavenRepositoryHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenSetLabel(context.Background(), mavenMavenRepositoryHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenSetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setLabel** | [**SetLabel**](SetLabel.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## RepositoriesMavenMavenUnsetLabel

> UnsetLabelResponse RepositoriesMavenMavenUnsetLabel(ctx, mavenMavenRepositoryHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

Unset a label



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
	mavenMavenRepositoryHref := "mavenMavenRepositoryHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenUnsetLabel(context.Background(), mavenMavenRepositoryHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenUnsetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unsetLabel** | [**UnsetLabel**](UnsetLabel.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## RepositoriesMavenMavenUpdate

> MavenMavenRepositoryResponse RepositoriesMavenMavenUpdate(ctx, mavenMavenRepositoryHref).MavenMavenRepository(mavenMavenRepository).XTaskDiagnostics(xTaskDiagnostics).Execute()

Update a maven repository



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
	mavenMavenRepositoryHref := "mavenMavenRepositoryHref_example" // string | 
	mavenMavenRepository := *openapiclient.NewMavenMavenRepository("Name_example") // MavenMavenRepository | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesMavenAPI.RepositoriesMavenMavenUpdate(context.Background(), mavenMavenRepositoryHref).MavenMavenRepository(mavenMavenRepository).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesMavenAPI.RepositoriesMavenMavenUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesMavenMavenUpdate`: MavenMavenRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesMavenAPI.RepositoriesMavenMavenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesMavenMavenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mavenMavenRepository** | [**MavenMavenRepository**](MavenMavenRepository.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**MavenMavenRepositoryResponse**](MavenMavenRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

