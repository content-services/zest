# \TaskGroupsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaskGroupsCancel**](TaskGroupsAPI.md#TaskGroupsCancel) | **Patch** /{task_group_href} | Cancel a task group
[**TaskGroupsList**](TaskGroupsAPI.md#TaskGroupsList) | **Get** /api/pulp/{pulp_domain}/api/v3/task-groups/ | List task groups
[**TaskGroupsRead**](TaskGroupsAPI.md#TaskGroupsRead) | **Get** /{task_group_href} | Inspect a task group



## TaskGroupsCancel

> TaskGroupResponse TaskGroupsCancel(ctx, taskGroupHref).PatchedTaskCancel(patchedTaskCancel).Execute()

Cancel a task group



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
	taskGroupHref := "taskGroupHref_example" // string | 
	patchedTaskCancel := *openapiclient.NewPatchedTaskCancel() // PatchedTaskCancel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskGroupsAPI.TaskGroupsCancel(context.Background(), taskGroupHref).PatchedTaskCancel(patchedTaskCancel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskGroupsAPI.TaskGroupsCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskGroupsCancel`: TaskGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `TaskGroupsAPI.TaskGroupsCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskGroupHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskGroupsCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedTaskCancel** | [**PatchedTaskCancel**](PatchedTaskCancel.md) |  | 

### Return type

[**TaskGroupResponse**](TaskGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskGroupsList

> PaginatedTaskGroupResponseList TaskGroupsList(ctx, pulpDomain).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List task groups

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
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskGroupsAPI.TaskGroupsList(context.Background(), pulpDomain).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskGroupsAPI.TaskGroupsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskGroupsList`: PaginatedTaskGroupResponseList
	fmt.Fprintf(os.Stdout, "Response from `TaskGroupsAPI.TaskGroupsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedTaskGroupResponseList**](PaginatedTaskGroupResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskGroupsRead

> TaskGroupResponse TaskGroupsRead(ctx, taskGroupHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a task group

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
	taskGroupHref := "taskGroupHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskGroupsAPI.TaskGroupsRead(context.Background(), taskGroupHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskGroupsAPI.TaskGroupsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskGroupsRead`: TaskGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `TaskGroupsAPI.TaskGroupsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskGroupHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskGroupsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**TaskGroupResponse**](TaskGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

