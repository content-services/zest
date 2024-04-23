# \RepositoriesOstreeAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesOstreeOstreeAddRole**](RepositoriesOstreeAPI.md#RepositoriesOstreeOstreeAddRole) | **Post** /{ostree_ostree_repository_href}add_role/ | Add a role
[**RepositoriesOstreeOstreeCreate**](RepositoriesOstreeAPI.md#RepositoriesOstreeOstreeCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/repositories/ostree/ostree/ | Create an ostree repository
[**RepositoriesOstreeOstreeDelete**](RepositoriesOstreeAPI.md#RepositoriesOstreeOstreeDelete) | **Delete** /{ostree_ostree_repository_href} | Delete an ostree repository
[**RepositoriesOstreeOstreeImportAll**](RepositoriesOstreeAPI.md#RepositoriesOstreeOstreeImportAll) | **Post** /{ostree_ostree_repository_href}import_all/ | Import refs and commits to a repository
[**RepositoriesOstreeOstreeImportCommits**](RepositoriesOstreeAPI.md#RepositoriesOstreeOstreeImportCommits) | **Post** /{ostree_ostree_repository_href}import_commits/ | Append child commits to a repository
[**RepositoriesOstreeOstreeList**](RepositoriesOstreeAPI.md#RepositoriesOstreeOstreeList) | **Get** /api/pulp/{pulp_domain}/api/v3/repositories/ostree/ostree/ | List ostree repositorys
[**RepositoriesOstreeOstreeListRoles**](RepositoriesOstreeAPI.md#RepositoriesOstreeOstreeListRoles) | **Get** /{ostree_ostree_repository_href}list_roles/ | List roles
[**RepositoriesOstreeOstreeModify**](RepositoriesOstreeAPI.md#RepositoriesOstreeOstreeModify) | **Post** /{ostree_ostree_repository_href}modify/ | Modify repository
[**RepositoriesOstreeOstreeMyPermissions**](RepositoriesOstreeAPI.md#RepositoriesOstreeOstreeMyPermissions) | **Get** /{ostree_ostree_repository_href}my_permissions/ | List user permissions
[**RepositoriesOstreeOstreePartialUpdate**](RepositoriesOstreeAPI.md#RepositoriesOstreeOstreePartialUpdate) | **Patch** /{ostree_ostree_repository_href} | Update an ostree repository
[**RepositoriesOstreeOstreeRead**](RepositoriesOstreeAPI.md#RepositoriesOstreeOstreeRead) | **Get** /{ostree_ostree_repository_href} | Inspect an ostree repository
[**RepositoriesOstreeOstreeRemoveRole**](RepositoriesOstreeAPI.md#RepositoriesOstreeOstreeRemoveRole) | **Post** /{ostree_ostree_repository_href}remove_role/ | Remove a role
[**RepositoriesOstreeOstreeSetLabel**](RepositoriesOstreeAPI.md#RepositoriesOstreeOstreeSetLabel) | **Post** /{ostree_ostree_repository_href}set_label/ | Set a label
[**RepositoriesOstreeOstreeSync**](RepositoriesOstreeAPI.md#RepositoriesOstreeOstreeSync) | **Post** /{ostree_ostree_repository_href}sync/ | Sync from remote
[**RepositoriesOstreeOstreeUnsetLabel**](RepositoriesOstreeAPI.md#RepositoriesOstreeOstreeUnsetLabel) | **Post** /{ostree_ostree_repository_href}unset_label/ | Unset a label
[**RepositoriesOstreeOstreeUpdate**](RepositoriesOstreeAPI.md#RepositoriesOstreeOstreeUpdate) | **Put** /{ostree_ostree_repository_href} | Update an ostree repository



## RepositoriesOstreeOstreeAddRole

> NestedRoleResponse RepositoriesOstreeOstreeAddRole(ctx, ostreeOstreeRepositoryHref).NestedRole(nestedRole).Execute()

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
	ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeAddRole(context.Background(), ostreeOstreeRepositoryHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeAPI.RepositoriesOstreeOstreeAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesOstreeOstreeAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeAPI.RepositoriesOstreeOstreeAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeAddRoleRequest struct via the builder pattern


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


## RepositoriesOstreeOstreeCreate

> OstreeOstreeRepositoryResponse RepositoriesOstreeOstreeCreate(ctx, pulpDomain).OstreeOstreeRepository(ostreeOstreeRepository).Execute()

Create an ostree repository



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
	ostreeOstreeRepository := *openapiclient.NewOstreeOstreeRepository("Name_example") // OstreeOstreeRepository | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeCreate(context.Background(), pulpDomain).OstreeOstreeRepository(ostreeOstreeRepository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeAPI.RepositoriesOstreeOstreeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesOstreeOstreeCreate`: OstreeOstreeRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeAPI.RepositoriesOstreeOstreeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ostreeOstreeRepository** | [**OstreeOstreeRepository**](OstreeOstreeRepository.md) |  | 

### Return type

[**OstreeOstreeRepositoryResponse**](OstreeOstreeRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesOstreeOstreeDelete

> AsyncOperationResponse RepositoriesOstreeOstreeDelete(ctx, ostreeOstreeRepositoryHref).Execute()

Delete an ostree repository



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
	ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeDelete(context.Background(), ostreeOstreeRepositoryHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeAPI.RepositoriesOstreeOstreeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesOstreeOstreeDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeAPI.RepositoriesOstreeOstreeDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeDeleteRequest struct via the builder pattern


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


## RepositoriesOstreeOstreeImportAll

> AsyncOperationResponse RepositoriesOstreeOstreeImportAll(ctx, ostreeOstreeRepositoryHref).OstreeImportAll(ostreeImportAll).Execute()

Import refs and commits to a repository



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
	ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
	ostreeImportAll := *openapiclient.NewOstreeImportAll("Artifact_example", "RepositoryName_example") // OstreeImportAll | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeImportAll(context.Background(), ostreeOstreeRepositoryHref).OstreeImportAll(ostreeImportAll).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeAPI.RepositoriesOstreeOstreeImportAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesOstreeOstreeImportAll`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeAPI.RepositoriesOstreeOstreeImportAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeImportAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ostreeImportAll** | [**OstreeImportAll**](OstreeImportAll.md) |  | 

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


## RepositoriesOstreeOstreeImportCommits

> AsyncOperationResponse RepositoriesOstreeOstreeImportCommits(ctx, ostreeOstreeRepositoryHref).OstreeImportCommitsToRef(ostreeImportCommitsToRef).Execute()

Append child commits to a repository



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
	ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
	ostreeImportCommitsToRef := *openapiclient.NewOstreeImportCommitsToRef("Artifact_example", "RepositoryName_example", "Ref_example") // OstreeImportCommitsToRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeImportCommits(context.Background(), ostreeOstreeRepositoryHref).OstreeImportCommitsToRef(ostreeImportCommitsToRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeAPI.RepositoriesOstreeOstreeImportCommits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesOstreeOstreeImportCommits`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeAPI.RepositoriesOstreeOstreeImportCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeImportCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ostreeImportCommitsToRef** | [**OstreeImportCommitsToRef**](OstreeImportCommitsToRef.md) |  | 

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


## RepositoriesOstreeOstreeList

> PaginatedostreeOstreeRepositoryResponseList RepositoriesOstreeOstreeList(ctx, pulpDomain).LatestWithContent(latestWithContent).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Remote(remote).RetainRepoVersions(retainRepoVersions).RetainRepoVersionsGt(retainRepoVersionsGt).RetainRepoVersionsGte(retainRepoVersionsGte).RetainRepoVersionsIsnull(retainRepoVersionsIsnull).RetainRepoVersionsLt(retainRepoVersionsLt).RetainRepoVersionsLte(retainRepoVersionsLte).RetainRepoVersionsNe(retainRepoVersionsNe).RetainRepoVersionsRange(retainRepoVersionsRange).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()

List ostree repositorys



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
	latestWithContent := "latestWithContent_example" // string | Content Unit referenced by HREF (optional)
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
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string |  (optional)
	remote := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Foreign Key referenced by HREF (optional)
	retainRepoVersions := int32(56) // int32 | Filter results where retain_repo_versions matches value (optional)
	retainRepoVersionsGt := int32(56) // int32 | Filter results where retain_repo_versions is greater than value (optional)
	retainRepoVersionsGte := int32(56) // int32 | Filter results where retain_repo_versions is greater than or equal to value (optional)
	retainRepoVersionsIsnull := true // bool | Filter results where retain_repo_versions has a null value (optional)
	retainRepoVersionsLt := int32(56) // int32 | Filter results where retain_repo_versions is less than value (optional)
	retainRepoVersionsLte := int32(56) // int32 | Filter results where retain_repo_versions is less than or equal to value (optional)
	retainRepoVersionsNe := int32(56) // int32 | Filter results where retain_repo_versions not equal to value (optional)
	retainRepoVersionsRange := []int32{int32(123)} // []int32 | Filter results where retain_repo_versions is between two comma separated values (optional)
	withContent := "withContent_example" // string | Content Unit referenced by HREF (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeList(context.Background(), pulpDomain).LatestWithContent(latestWithContent).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Remote(remote).RetainRepoVersions(retainRepoVersions).RetainRepoVersionsGt(retainRepoVersionsGt).RetainRepoVersionsGte(retainRepoVersionsGte).RetainRepoVersionsIsnull(retainRepoVersionsIsnull).RetainRepoVersionsLt(retainRepoVersionsLt).RetainRepoVersionsLte(retainRepoVersionsLte).RetainRepoVersionsNe(retainRepoVersionsNe).RetainRepoVersionsRange(retainRepoVersionsRange).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeAPI.RepositoriesOstreeOstreeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesOstreeOstreeList`: PaginatedostreeOstreeRepositoryResponseList
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeAPI.RepositoriesOstreeOstreeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **latestWithContent** | **string** | Content Unit referenced by HREF | 
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
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** |  | 
 **remote** | **string** | Foreign Key referenced by HREF | 
 **retainRepoVersions** | **int32** | Filter results where retain_repo_versions matches value | 
 **retainRepoVersionsGt** | **int32** | Filter results where retain_repo_versions is greater than value | 
 **retainRepoVersionsGte** | **int32** | Filter results where retain_repo_versions is greater than or equal to value | 
 **retainRepoVersionsIsnull** | **bool** | Filter results where retain_repo_versions has a null value | 
 **retainRepoVersionsLt** | **int32** | Filter results where retain_repo_versions is less than value | 
 **retainRepoVersionsLte** | **int32** | Filter results where retain_repo_versions is less than or equal to value | 
 **retainRepoVersionsNe** | **int32** | Filter results where retain_repo_versions not equal to value | 
 **retainRepoVersionsRange** | **[]int32** | Filter results where retain_repo_versions is between two comma separated values | 
 **withContent** | **string** | Content Unit referenced by HREF | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedostreeOstreeRepositoryResponseList**](PaginatedostreeOstreeRepositoryResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesOstreeOstreeListRoles

> ObjectRolesResponse RepositoriesOstreeOstreeListRoles(ctx, ostreeOstreeRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeListRoles(context.Background(), ostreeOstreeRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeAPI.RepositoriesOstreeOstreeListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesOstreeOstreeListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeAPI.RepositoriesOstreeOstreeListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeListRolesRequest struct via the builder pattern


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


## RepositoriesOstreeOstreeModify

> AsyncOperationResponse RepositoriesOstreeOstreeModify(ctx, ostreeOstreeRepositoryHref).RepositoryAddRemoveContent(repositoryAddRemoveContent).Execute()

Modify repository



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
	ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
	repositoryAddRemoveContent := *openapiclient.NewRepositoryAddRemoveContent() // RepositoryAddRemoveContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeModify(context.Background(), ostreeOstreeRepositoryHref).RepositoryAddRemoveContent(repositoryAddRemoveContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeAPI.RepositoriesOstreeOstreeModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesOstreeOstreeModify`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeAPI.RepositoriesOstreeOstreeModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeModifyRequest struct via the builder pattern


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


## RepositoriesOstreeOstreeMyPermissions

> MyPermissionsResponse RepositoriesOstreeOstreeMyPermissions(ctx, ostreeOstreeRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeMyPermissions(context.Background(), ostreeOstreeRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeAPI.RepositoriesOstreeOstreeMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesOstreeOstreeMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeAPI.RepositoriesOstreeOstreeMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeMyPermissionsRequest struct via the builder pattern


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


## RepositoriesOstreeOstreePartialUpdate

> AsyncOperationResponse RepositoriesOstreeOstreePartialUpdate(ctx, ostreeOstreeRepositoryHref).PatchedostreeOstreeRepository(patchedostreeOstreeRepository).Execute()

Update an ostree repository



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
	ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
	patchedostreeOstreeRepository := *openapiclient.NewPatchedostreeOstreeRepository() // PatchedostreeOstreeRepository | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreePartialUpdate(context.Background(), ostreeOstreeRepositoryHref).PatchedostreeOstreeRepository(patchedostreeOstreeRepository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeAPI.RepositoriesOstreeOstreePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesOstreeOstreePartialUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeAPI.RepositoriesOstreeOstreePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedostreeOstreeRepository** | [**PatchedostreeOstreeRepository**](PatchedostreeOstreeRepository.md) |  | 

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


## RepositoriesOstreeOstreeRead

> OstreeOstreeRepositoryResponse RepositoriesOstreeOstreeRead(ctx, ostreeOstreeRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an ostree repository



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
	ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeRead(context.Background(), ostreeOstreeRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeAPI.RepositoriesOstreeOstreeRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesOstreeOstreeRead`: OstreeOstreeRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeAPI.RepositoriesOstreeOstreeRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**OstreeOstreeRepositoryResponse**](OstreeOstreeRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesOstreeOstreeRemoveRole

> NestedRoleResponse RepositoriesOstreeOstreeRemoveRole(ctx, ostreeOstreeRepositoryHref).NestedRole(nestedRole).Execute()

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
	ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeRemoveRole(context.Background(), ostreeOstreeRepositoryHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeAPI.RepositoriesOstreeOstreeRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesOstreeOstreeRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeAPI.RepositoriesOstreeOstreeRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeRemoveRoleRequest struct via the builder pattern


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


## RepositoriesOstreeOstreeSetLabel

> SetLabelResponse RepositoriesOstreeOstreeSetLabel(ctx, ostreeOstreeRepositoryHref).SetLabel(setLabel).Execute()

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
	ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeSetLabel(context.Background(), ostreeOstreeRepositoryHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeAPI.RepositoriesOstreeOstreeSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesOstreeOstreeSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeAPI.RepositoriesOstreeOstreeSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeSetLabelRequest struct via the builder pattern


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


## RepositoriesOstreeOstreeSync

> AsyncOperationResponse RepositoriesOstreeOstreeSync(ctx, ostreeOstreeRepositoryHref).RepositorySyncURL(repositorySyncURL).Execute()

Sync from remote



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
	ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
	repositorySyncURL := *openapiclient.NewRepositorySyncURL() // RepositorySyncURL | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeSync(context.Background(), ostreeOstreeRepositoryHref).RepositorySyncURL(repositorySyncURL).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeAPI.RepositoriesOstreeOstreeSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesOstreeOstreeSync`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeAPI.RepositoriesOstreeOstreeSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositorySyncURL** | [**RepositorySyncURL**](RepositorySyncURL.md) |  | 

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


## RepositoriesOstreeOstreeUnsetLabel

> UnsetLabelResponse RepositoriesOstreeOstreeUnsetLabel(ctx, ostreeOstreeRepositoryHref).UnsetLabel(unsetLabel).Execute()

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
	ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeUnsetLabel(context.Background(), ostreeOstreeRepositoryHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeAPI.RepositoriesOstreeOstreeUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesOstreeOstreeUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeAPI.RepositoriesOstreeOstreeUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeUnsetLabelRequest struct via the builder pattern


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


## RepositoriesOstreeOstreeUpdate

> AsyncOperationResponse RepositoriesOstreeOstreeUpdate(ctx, ostreeOstreeRepositoryHref).OstreeOstreeRepository(ostreeOstreeRepository).Execute()

Update an ostree repository



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
	ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
	ostreeOstreeRepository := *openapiclient.NewOstreeOstreeRepository("Name_example") // OstreeOstreeRepository | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeUpdate(context.Background(), ostreeOstreeRepositoryHref).OstreeOstreeRepository(ostreeOstreeRepository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeAPI.RepositoriesOstreeOstreeUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesOstreeOstreeUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeAPI.RepositoriesOstreeOstreeUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ostreeOstreeRepository** | [**OstreeOstreeRepository**](OstreeOstreeRepository.md) |  | 

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

