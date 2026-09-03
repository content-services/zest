# \RepositoriesHuggingFaceAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesHuggingFaceHuggingFaceCreate**](RepositoriesHuggingFaceAPI.md#RepositoriesHuggingFaceHuggingFaceCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/repositories/hugging_face/hugging-face/ | Create a hugging face repository
[**RepositoriesHuggingFaceHuggingFaceDelete**](RepositoriesHuggingFaceAPI.md#RepositoriesHuggingFaceHuggingFaceDelete) | **Delete** /{hugging_face_hugging_face_repository_href} | Delete a hugging face repository
[**RepositoriesHuggingFaceHuggingFaceList**](RepositoriesHuggingFaceAPI.md#RepositoriesHuggingFaceHuggingFaceList) | **Get** /api/pulp/{pulp_domain}/api/v3/repositories/hugging_face/hugging-face/ | List hugging face repositorys
[**RepositoriesHuggingFaceHuggingFaceModify**](RepositoriesHuggingFaceAPI.md#RepositoriesHuggingFaceHuggingFaceModify) | **Post** /{hugging_face_hugging_face_repository_href}modify/ | Modify Repository Content
[**RepositoriesHuggingFaceHuggingFacePartialUpdate**](RepositoriesHuggingFaceAPI.md#RepositoriesHuggingFaceHuggingFacePartialUpdate) | **Patch** /{hugging_face_hugging_face_repository_href} | Update a hugging face repository
[**RepositoriesHuggingFaceHuggingFaceRead**](RepositoriesHuggingFaceAPI.md#RepositoriesHuggingFaceHuggingFaceRead) | **Get** /{hugging_face_hugging_face_repository_href} | Inspect a hugging face repository
[**RepositoriesHuggingFaceHuggingFaceSetLabel**](RepositoriesHuggingFaceAPI.md#RepositoriesHuggingFaceHuggingFaceSetLabel) | **Post** /{hugging_face_hugging_face_repository_href}set_label/ | Set a label
[**RepositoriesHuggingFaceHuggingFaceSync**](RepositoriesHuggingFaceAPI.md#RepositoriesHuggingFaceHuggingFaceSync) | **Post** /{hugging_face_hugging_face_repository_href}sync/ | Sync from remote
[**RepositoriesHuggingFaceHuggingFaceUnsetLabel**](RepositoriesHuggingFaceAPI.md#RepositoriesHuggingFaceHuggingFaceUnsetLabel) | **Post** /{hugging_face_hugging_face_repository_href}unset_label/ | Unset a label
[**RepositoriesHuggingFaceHuggingFaceUpdate**](RepositoriesHuggingFaceAPI.md#RepositoriesHuggingFaceHuggingFaceUpdate) | **Put** /{hugging_face_hugging_face_repository_href} | Update a hugging face repository



## RepositoriesHuggingFaceHuggingFaceCreate

> HuggingFaceHuggingFaceRepositoryResponse RepositoriesHuggingFaceHuggingFaceCreate(ctx, pulpDomain).HuggingFaceHuggingFaceRepository(huggingFaceHuggingFaceRepository).XTaskDiagnostics(xTaskDiagnostics).Execute()

Create a hugging face repository



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
	huggingFaceHuggingFaceRepository := *openapiclient.NewHuggingFaceHuggingFaceRepository("Name_example") // HuggingFaceHuggingFaceRepository | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceCreate(context.Background(), pulpDomain).HuggingFaceHuggingFaceRepository(huggingFaceHuggingFaceRepository).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesHuggingFaceHuggingFaceCreate`: HuggingFaceHuggingFaceRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesHuggingFaceHuggingFaceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **huggingFaceHuggingFaceRepository** | [**HuggingFaceHuggingFaceRepository**](HuggingFaceHuggingFaceRepository.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**HuggingFaceHuggingFaceRepositoryResponse**](HuggingFaceHuggingFaceRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesHuggingFaceHuggingFaceDelete

> AsyncOperationResponse RepositoriesHuggingFaceHuggingFaceDelete(ctx, huggingFaceHuggingFaceRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Execute()

Delete a hugging face repository



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
	huggingFaceHuggingFaceRepositoryHref := "huggingFaceHuggingFaceRepositoryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceDelete(context.Background(), huggingFaceHuggingFaceRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesHuggingFaceHuggingFaceDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesHuggingFaceHuggingFaceDeleteRequest struct via the builder pattern


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


## RepositoriesHuggingFaceHuggingFaceList

> PaginatedhuggingFaceHuggingFaceRepositoryResponseList RepositoriesHuggingFaceHuggingFaceList(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).LatestWithContent(latestWithContent).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Remote(remote).RetainCheckpoints(retainCheckpoints).RetainCheckpointsGt(retainCheckpointsGt).RetainCheckpointsGte(retainCheckpointsGte).RetainCheckpointsIsnull(retainCheckpointsIsnull).RetainCheckpointsLt(retainCheckpointsLt).RetainCheckpointsLte(retainCheckpointsLte).RetainCheckpointsNe(retainCheckpointsNe).RetainCheckpointsRange(retainCheckpointsRange).RetainRepoVersions(retainRepoVersions).RetainRepoVersionsGt(retainRepoVersionsGt).RetainRepoVersionsGte(retainRepoVersionsGte).RetainRepoVersionsIsnull(retainRepoVersionsIsnull).RetainRepoVersionsLt(retainRepoVersionsLt).RetainRepoVersionsLte(retainRepoVersionsLte).RetainRepoVersionsNe(retainRepoVersionsNe).RetainRepoVersionsRange(retainRepoVersionsRange).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()

List hugging face repositorys



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
	resp, r, err := apiClient.RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceList(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).LatestWithContent(latestWithContent).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Remote(remote).RetainCheckpoints(retainCheckpoints).RetainCheckpointsGt(retainCheckpointsGt).RetainCheckpointsGte(retainCheckpointsGte).RetainCheckpointsIsnull(retainCheckpointsIsnull).RetainCheckpointsLt(retainCheckpointsLt).RetainCheckpointsLte(retainCheckpointsLte).RetainCheckpointsNe(retainCheckpointsNe).RetainCheckpointsRange(retainCheckpointsRange).RetainRepoVersions(retainRepoVersions).RetainRepoVersionsGt(retainRepoVersionsGt).RetainRepoVersionsGte(retainRepoVersionsGte).RetainRepoVersionsIsnull(retainRepoVersionsIsnull).RetainRepoVersionsLt(retainRepoVersionsLt).RetainRepoVersionsLte(retainRepoVersionsLte).RetainRepoVersionsNe(retainRepoVersionsNe).RetainRepoVersionsRange(retainRepoVersionsRange).WithContent(withContent).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesHuggingFaceHuggingFaceList`: PaginatedhuggingFaceHuggingFaceRepositoryResponseList
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesHuggingFaceHuggingFaceListRequest struct via the builder pattern


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

[**PaginatedhuggingFaceHuggingFaceRepositoryResponseList**](PaginatedhuggingFaceHuggingFaceRepositoryResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesHuggingFaceHuggingFaceModify

> AsyncOperationResponse RepositoriesHuggingFaceHuggingFaceModify(ctx, huggingFaceHuggingFaceRepositoryHref).RepositoryAddRemoveContent(repositoryAddRemoveContent).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	huggingFaceHuggingFaceRepositoryHref := "huggingFaceHuggingFaceRepositoryHref_example" // string | 
	repositoryAddRemoveContent := *openapiclient.NewRepositoryAddRemoveContent() // RepositoryAddRemoveContent | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceModify(context.Background(), huggingFaceHuggingFaceRepositoryHref).RepositoryAddRemoveContent(repositoryAddRemoveContent).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesHuggingFaceHuggingFaceModify`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesHuggingFaceHuggingFaceModifyRequest struct via the builder pattern


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


## RepositoriesHuggingFaceHuggingFacePartialUpdate

> HuggingFaceHuggingFaceRepositoryResponse RepositoriesHuggingFaceHuggingFacePartialUpdate(ctx, huggingFaceHuggingFaceRepositoryHref).PatchedhuggingFaceHuggingFaceRepository(patchedhuggingFaceHuggingFaceRepository).XTaskDiagnostics(xTaskDiagnostics).Execute()

Update a hugging face repository



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
	huggingFaceHuggingFaceRepositoryHref := "huggingFaceHuggingFaceRepositoryHref_example" // string | 
	patchedhuggingFaceHuggingFaceRepository := *openapiclient.NewPatchedhuggingFaceHuggingFaceRepository() // PatchedhuggingFaceHuggingFaceRepository | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFacePartialUpdate(context.Background(), huggingFaceHuggingFaceRepositoryHref).PatchedhuggingFaceHuggingFaceRepository(patchedhuggingFaceHuggingFaceRepository).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFacePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesHuggingFaceHuggingFacePartialUpdate`: HuggingFaceHuggingFaceRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFacePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesHuggingFaceHuggingFacePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedhuggingFaceHuggingFaceRepository** | [**PatchedhuggingFaceHuggingFaceRepository**](PatchedhuggingFaceHuggingFaceRepository.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**HuggingFaceHuggingFaceRepositoryResponse**](HuggingFaceHuggingFaceRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesHuggingFaceHuggingFaceRead

> HuggingFaceHuggingFaceRepositoryResponse RepositoriesHuggingFaceHuggingFaceRead(ctx, huggingFaceHuggingFaceRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a hugging face repository



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
	huggingFaceHuggingFaceRepositoryHref := "huggingFaceHuggingFaceRepositoryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceRead(context.Background(), huggingFaceHuggingFaceRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesHuggingFaceHuggingFaceRead`: HuggingFaceHuggingFaceRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesHuggingFaceHuggingFaceReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**HuggingFaceHuggingFaceRepositoryResponse**](HuggingFaceHuggingFaceRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesHuggingFaceHuggingFaceSetLabel

> SetLabelResponse RepositoriesHuggingFaceHuggingFaceSetLabel(ctx, huggingFaceHuggingFaceRepositoryHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	huggingFaceHuggingFaceRepositoryHref := "huggingFaceHuggingFaceRepositoryHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceSetLabel(context.Background(), huggingFaceHuggingFaceRepositoryHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesHuggingFaceHuggingFaceSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesHuggingFaceHuggingFaceSetLabelRequest struct via the builder pattern


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


## RepositoriesHuggingFaceHuggingFaceSync

> AsyncOperationResponse RepositoriesHuggingFaceHuggingFaceSync(ctx, huggingFaceHuggingFaceRepositoryHref).RepositorySyncURL(repositorySyncURL).XTaskDiagnostics(xTaskDiagnostics).Execute()

Sync from remote



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
	huggingFaceHuggingFaceRepositoryHref := "huggingFaceHuggingFaceRepositoryHref_example" // string | 
	repositorySyncURL := *openapiclient.NewRepositorySyncURL() // RepositorySyncURL | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceSync(context.Background(), huggingFaceHuggingFaceRepositoryHref).RepositorySyncURL(repositorySyncURL).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesHuggingFaceHuggingFaceSync`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesHuggingFaceHuggingFaceSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositorySyncURL** | [**RepositorySyncURL**](RepositorySyncURL.md) |  | 
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


## RepositoriesHuggingFaceHuggingFaceUnsetLabel

> UnsetLabelResponse RepositoriesHuggingFaceHuggingFaceUnsetLabel(ctx, huggingFaceHuggingFaceRepositoryHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	huggingFaceHuggingFaceRepositoryHref := "huggingFaceHuggingFaceRepositoryHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceUnsetLabel(context.Background(), huggingFaceHuggingFaceRepositoryHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesHuggingFaceHuggingFaceUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesHuggingFaceHuggingFaceUnsetLabelRequest struct via the builder pattern


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


## RepositoriesHuggingFaceHuggingFaceUpdate

> HuggingFaceHuggingFaceRepositoryResponse RepositoriesHuggingFaceHuggingFaceUpdate(ctx, huggingFaceHuggingFaceRepositoryHref).HuggingFaceHuggingFaceRepository(huggingFaceHuggingFaceRepository).XTaskDiagnostics(xTaskDiagnostics).Execute()

Update a hugging face repository



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
	huggingFaceHuggingFaceRepositoryHref := "huggingFaceHuggingFaceRepositoryHref_example" // string | 
	huggingFaceHuggingFaceRepository := *openapiclient.NewHuggingFaceHuggingFaceRepository("Name_example") // HuggingFaceHuggingFaceRepository | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceUpdate(context.Background(), huggingFaceHuggingFaceRepositoryHref).HuggingFaceHuggingFaceRepository(huggingFaceHuggingFaceRepository).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesHuggingFaceHuggingFaceUpdate`: HuggingFaceHuggingFaceRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesHuggingFaceAPI.RepositoriesHuggingFaceHuggingFaceUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesHuggingFaceHuggingFaceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **huggingFaceHuggingFaceRepository** | [**HuggingFaceHuggingFaceRepository**](HuggingFaceHuggingFaceRepository.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**HuggingFaceHuggingFaceRepositoryResponse**](HuggingFaceHuggingFaceRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

