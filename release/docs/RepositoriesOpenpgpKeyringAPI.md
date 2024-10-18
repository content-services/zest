# \RepositoriesOpenpgpKeyringAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesCoreOpenpgpKeyringAddRole**](RepositoriesOpenpgpKeyringAPI.md#RepositoriesCoreOpenpgpKeyringAddRole) | **Post** /{open_p_g_p_keyring_href}add_role/ | Add a role
[**RepositoriesCoreOpenpgpKeyringCreate**](RepositoriesOpenpgpKeyringAPI.md#RepositoriesCoreOpenpgpKeyringCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/repositories/core/openpgp_keyring/ | Create an open pgp keyring
[**RepositoriesCoreOpenpgpKeyringDelete**](RepositoriesOpenpgpKeyringAPI.md#RepositoriesCoreOpenpgpKeyringDelete) | **Delete** /{open_p_g_p_keyring_href} | Delete an open pgp keyring
[**RepositoriesCoreOpenpgpKeyringList**](RepositoriesOpenpgpKeyringAPI.md#RepositoriesCoreOpenpgpKeyringList) | **Get** /api/pulp/{pulp_domain}/api/v3/repositories/core/openpgp_keyring/ | List open pgp keyrings
[**RepositoriesCoreOpenpgpKeyringListRoles**](RepositoriesOpenpgpKeyringAPI.md#RepositoriesCoreOpenpgpKeyringListRoles) | **Get** /{open_p_g_p_keyring_href}list_roles/ | List roles
[**RepositoriesCoreOpenpgpKeyringModify**](RepositoriesOpenpgpKeyringAPI.md#RepositoriesCoreOpenpgpKeyringModify) | **Post** /{open_p_g_p_keyring_href}modify/ | Modify Repository Content
[**RepositoriesCoreOpenpgpKeyringMyPermissions**](RepositoriesOpenpgpKeyringAPI.md#RepositoriesCoreOpenpgpKeyringMyPermissions) | **Get** /{open_p_g_p_keyring_href}my_permissions/ | List user permissions
[**RepositoriesCoreOpenpgpKeyringPartialUpdate**](RepositoriesOpenpgpKeyringAPI.md#RepositoriesCoreOpenpgpKeyringPartialUpdate) | **Patch** /{open_p_g_p_keyring_href} | Update an open pgp keyring
[**RepositoriesCoreOpenpgpKeyringRead**](RepositoriesOpenpgpKeyringAPI.md#RepositoriesCoreOpenpgpKeyringRead) | **Get** /{open_p_g_p_keyring_href} | Inspect an open pgp keyring
[**RepositoriesCoreOpenpgpKeyringRemoveRole**](RepositoriesOpenpgpKeyringAPI.md#RepositoriesCoreOpenpgpKeyringRemoveRole) | **Post** /{open_p_g_p_keyring_href}remove_role/ | Remove a role
[**RepositoriesCoreOpenpgpKeyringSetLabel**](RepositoriesOpenpgpKeyringAPI.md#RepositoriesCoreOpenpgpKeyringSetLabel) | **Post** /{open_p_g_p_keyring_href}set_label/ | Set a label
[**RepositoriesCoreOpenpgpKeyringUnsetLabel**](RepositoriesOpenpgpKeyringAPI.md#RepositoriesCoreOpenpgpKeyringUnsetLabel) | **Post** /{open_p_g_p_keyring_href}unset_label/ | Unset a label
[**RepositoriesCoreOpenpgpKeyringUpdate**](RepositoriesOpenpgpKeyringAPI.md#RepositoriesCoreOpenpgpKeyringUpdate) | **Put** /{open_p_g_p_keyring_href} | Update an open pgp keyring



## RepositoriesCoreOpenpgpKeyringAddRole

> NestedRoleResponse RepositoriesCoreOpenpgpKeyringAddRole(ctx, openPGPKeyringHref).NestedRole(nestedRole).Execute()

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
	openPGPKeyringHref := "openPGPKeyringHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringAddRole(context.Background(), openPGPKeyringHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesCoreOpenpgpKeyringAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPKeyringHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesCoreOpenpgpKeyringAddRoleRequest struct via the builder pattern


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


## RepositoriesCoreOpenpgpKeyringCreate

> OpenPGPKeyringResponse RepositoriesCoreOpenpgpKeyringCreate(ctx, pulpDomain).OpenPGPKeyring(openPGPKeyring).Execute()

Create an open pgp keyring



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
	openPGPKeyring := *openapiclient.NewOpenPGPKeyring("Name_example") // OpenPGPKeyring | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringCreate(context.Background(), pulpDomain).OpenPGPKeyring(openPGPKeyring).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesCoreOpenpgpKeyringCreate`: OpenPGPKeyringResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesCoreOpenpgpKeyringCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openPGPKeyring** | [**OpenPGPKeyring**](OpenPGPKeyring.md) |  | 

### Return type

[**OpenPGPKeyringResponse**](OpenPGPKeyringResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesCoreOpenpgpKeyringDelete

> AsyncOperationResponse RepositoriesCoreOpenpgpKeyringDelete(ctx, openPGPKeyringHref).Execute()

Delete an open pgp keyring



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
	openPGPKeyringHref := "openPGPKeyringHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringDelete(context.Background(), openPGPKeyringHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesCoreOpenpgpKeyringDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPKeyringHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesCoreOpenpgpKeyringDeleteRequest struct via the builder pattern


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


## RepositoriesCoreOpenpgpKeyringList

> PaginatedOpenPGPKeyringResponseList RepositoriesCoreOpenpgpKeyringList(ctx, pulpDomain).LatestWithContent(latestWithContent).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Remote(remote).RetainRepoVersions(retainRepoVersions).RetainRepoVersionsGt(retainRepoVersionsGt).RetainRepoVersionsGte(retainRepoVersionsGte).RetainRepoVersionsIsnull(retainRepoVersionsIsnull).RetainRepoVersionsLt(retainRepoVersionsLt).RetainRepoVersionsLte(retainRepoVersionsLte).RetainRepoVersionsNe(retainRepoVersionsNe).RetainRepoVersionsRange(retainRepoVersionsRange).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()

List open pgp keyrings



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
	resp, r, err := apiClient.RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringList(context.Background(), pulpDomain).LatestWithContent(latestWithContent).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Remote(remote).RetainRepoVersions(retainRepoVersions).RetainRepoVersionsGt(retainRepoVersionsGt).RetainRepoVersionsGte(retainRepoVersionsGte).RetainRepoVersionsIsnull(retainRepoVersionsIsnull).RetainRepoVersionsLt(retainRepoVersionsLt).RetainRepoVersionsLte(retainRepoVersionsLte).RetainRepoVersionsNe(retainRepoVersionsNe).RetainRepoVersionsRange(retainRepoVersionsRange).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesCoreOpenpgpKeyringList`: PaginatedOpenPGPKeyringResponseList
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesCoreOpenpgpKeyringListRequest struct via the builder pattern


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

[**PaginatedOpenPGPKeyringResponseList**](PaginatedOpenPGPKeyringResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesCoreOpenpgpKeyringListRoles

> ObjectRolesResponse RepositoriesCoreOpenpgpKeyringListRoles(ctx, openPGPKeyringHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	openPGPKeyringHref := "openPGPKeyringHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringListRoles(context.Background(), openPGPKeyringHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesCoreOpenpgpKeyringListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPKeyringHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesCoreOpenpgpKeyringListRolesRequest struct via the builder pattern


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


## RepositoriesCoreOpenpgpKeyringModify

> AsyncOperationResponse RepositoriesCoreOpenpgpKeyringModify(ctx, openPGPKeyringHref).RepositoryAddRemoveContent(repositoryAddRemoveContent).Execute()

Modify Repository Content



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
	openPGPKeyringHref := "openPGPKeyringHref_example" // string | 
	repositoryAddRemoveContent := *openapiclient.NewRepositoryAddRemoveContent() // RepositoryAddRemoveContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringModify(context.Background(), openPGPKeyringHref).RepositoryAddRemoveContent(repositoryAddRemoveContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesCoreOpenpgpKeyringModify`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPKeyringHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesCoreOpenpgpKeyringModifyRequest struct via the builder pattern


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


## RepositoriesCoreOpenpgpKeyringMyPermissions

> MyPermissionsResponse RepositoriesCoreOpenpgpKeyringMyPermissions(ctx, openPGPKeyringHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	openPGPKeyringHref := "openPGPKeyringHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringMyPermissions(context.Background(), openPGPKeyringHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesCoreOpenpgpKeyringMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPKeyringHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesCoreOpenpgpKeyringMyPermissionsRequest struct via the builder pattern


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


## RepositoriesCoreOpenpgpKeyringPartialUpdate

> AsyncOperationResponse RepositoriesCoreOpenpgpKeyringPartialUpdate(ctx, openPGPKeyringHref).PatchedOpenPGPKeyring(patchedOpenPGPKeyring).Execute()

Update an open pgp keyring



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
	openPGPKeyringHref := "openPGPKeyringHref_example" // string | 
	patchedOpenPGPKeyring := *openapiclient.NewPatchedOpenPGPKeyring() // PatchedOpenPGPKeyring | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringPartialUpdate(context.Background(), openPGPKeyringHref).PatchedOpenPGPKeyring(patchedOpenPGPKeyring).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesCoreOpenpgpKeyringPartialUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPKeyringHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesCoreOpenpgpKeyringPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedOpenPGPKeyring** | [**PatchedOpenPGPKeyring**](PatchedOpenPGPKeyring.md) |  | 

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


## RepositoriesCoreOpenpgpKeyringRead

> OpenPGPKeyringResponse RepositoriesCoreOpenpgpKeyringRead(ctx, openPGPKeyringHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an open pgp keyring



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
	openPGPKeyringHref := "openPGPKeyringHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringRead(context.Background(), openPGPKeyringHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesCoreOpenpgpKeyringRead`: OpenPGPKeyringResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPKeyringHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesCoreOpenpgpKeyringReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**OpenPGPKeyringResponse**](OpenPGPKeyringResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesCoreOpenpgpKeyringRemoveRole

> NestedRoleResponse RepositoriesCoreOpenpgpKeyringRemoveRole(ctx, openPGPKeyringHref).NestedRole(nestedRole).Execute()

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
	openPGPKeyringHref := "openPGPKeyringHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringRemoveRole(context.Background(), openPGPKeyringHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesCoreOpenpgpKeyringRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPKeyringHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesCoreOpenpgpKeyringRemoveRoleRequest struct via the builder pattern


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


## RepositoriesCoreOpenpgpKeyringSetLabel

> SetLabelResponse RepositoriesCoreOpenpgpKeyringSetLabel(ctx, openPGPKeyringHref).SetLabel(setLabel).Execute()

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
	openPGPKeyringHref := "openPGPKeyringHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringSetLabel(context.Background(), openPGPKeyringHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesCoreOpenpgpKeyringSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPKeyringHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesCoreOpenpgpKeyringSetLabelRequest struct via the builder pattern


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


## RepositoriesCoreOpenpgpKeyringUnsetLabel

> UnsetLabelResponse RepositoriesCoreOpenpgpKeyringUnsetLabel(ctx, openPGPKeyringHref).UnsetLabel(unsetLabel).Execute()

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
	openPGPKeyringHref := "openPGPKeyringHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringUnsetLabel(context.Background(), openPGPKeyringHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesCoreOpenpgpKeyringUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPKeyringHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesCoreOpenpgpKeyringUnsetLabelRequest struct via the builder pattern


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


## RepositoriesCoreOpenpgpKeyringUpdate

> AsyncOperationResponse RepositoriesCoreOpenpgpKeyringUpdate(ctx, openPGPKeyringHref).OpenPGPKeyring(openPGPKeyring).Execute()

Update an open pgp keyring



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
	openPGPKeyringHref := "openPGPKeyringHref_example" // string | 
	openPGPKeyring := *openapiclient.NewOpenPGPKeyring("Name_example") // OpenPGPKeyring | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringUpdate(context.Background(), openPGPKeyringHref).OpenPGPKeyring(openPGPKeyring).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesCoreOpenpgpKeyringUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesOpenpgpKeyringAPI.RepositoriesCoreOpenpgpKeyringUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPKeyringHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesCoreOpenpgpKeyringUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openPGPKeyring** | [**OpenPGPKeyring**](OpenPGPKeyring.md) |  | 

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

