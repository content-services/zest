# \TasksAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminTasks**](TasksAPI.md#AdminTasks) | **Get** /api/pulp/admin/tasks/ | 
[**TasksAddRole**](TasksAPI.md#TasksAddRole) | **Post** /{task_href}add_role/ | Add a role
[**TasksCancel**](TasksAPI.md#TasksCancel) | **Patch** /{task_href} | Cancel a task
[**TasksDelete**](TasksAPI.md#TasksDelete) | **Delete** /{task_href} | Delete a task
[**TasksList**](TasksAPI.md#TasksList) | **Get** /api/pulp/{pulp_domain}/api/v3/tasks/ | List tasks
[**TasksListRoles**](TasksAPI.md#TasksListRoles) | **Get** /{task_href}list_roles/ | List roles
[**TasksMyPermissions**](TasksAPI.md#TasksMyPermissions) | **Get** /{task_href}my_permissions/ | List user permissions
[**TasksProfileArtifacts**](TasksAPI.md#TasksProfileArtifacts) | **Get** /{task_href}profile_artifacts/ | Fetch downloadable links for profile artifacts
[**TasksPurge**](TasksAPI.md#TasksPurge) | **Post** /api/pulp/{pulp_domain}/api/v3/tasks/purge/ | Purge Completed Tasks
[**TasksRead**](TasksAPI.md#TasksRead) | **Get** /{task_href} | Inspect a task
[**TasksRemoveRole**](TasksAPI.md#TasksRemoveRole) | **Post** /{task_href}remove_role/ | Remove a role



## AdminTasks

> PaginatedTaskResponseList AdminTasks(ctx).XTaskDiagnostics(xTaskDiagnostics).ChildTasks(childTasks).CreatedResources(createdResources).ExclusiveResources(exclusiveResources).ExclusiveResourcesIn(exclusiveResourcesIn).FinishedAt(finishedAt).FinishedAtGt(finishedAtGt).FinishedAtGte(finishedAtGte).FinishedAtIsnull(finishedAtIsnull).FinishedAtLt(finishedAtLt).FinishedAtLte(finishedAtLte).FinishedAtRange(finishedAtRange).Limit(limit).LoggingCid(loggingCid).LoggingCidContains(loggingCidContains).Name(name).NameContains(nameContains).NameIn(nameIn).NameNe(nameNe).Offset(offset).Ordering(ordering).ParentTask(parentTask).PrnIn(prnIn).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedIsnull(pulpCreatedIsnull).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).ReservedResources(reservedResources).ReservedResourcesIn(reservedResourcesIn).SharedResources(sharedResources).SharedResourcesIn(sharedResourcesIn).StartedAt(startedAt).StartedAtGt(startedAtGt).StartedAtGte(startedAtGte).StartedAtIsnull(startedAtIsnull).StartedAtLt(startedAtLt).StartedAtLte(startedAtLte).StartedAtRange(startedAtRange).State(state).StateIn(stateIn).StateNe(stateNe).TaskGroup(taskGroup).UnblockedAt(unblockedAt).UnblockedAtGt(unblockedAtGt).UnblockedAtGte(unblockedAtGte).UnblockedAtIsnull(unblockedAtIsnull).UnblockedAtLt(unblockedAtLt).UnblockedAtLte(unblockedAtLte).UnblockedAtRange(unblockedAtRange).Worker(worker).Fields(fields).ExcludeFields(excludeFields).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	childTasks := "childTasks_example" // string |  (optional)
	createdResources := "createdResources_example" // string |  (optional)
	exclusiveResources := "exclusiveResources_example" // string |  (optional)
	exclusiveResourcesIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	finishedAt := time.Now() // time.Time | Filter results where finished_at matches value (optional)
	finishedAtGt := time.Now() // time.Time | Filter results where finished_at is greater than value (optional)
	finishedAtGte := time.Now() // time.Time | Filter results where finished_at is greater than or equal to value (optional)
	finishedAtIsnull := true // bool | Filter results where finished_at has a null value (optional)
	finishedAtLt := time.Now() // time.Time | Filter results where finished_at is less than value (optional)
	finishedAtLte := time.Now() // time.Time | Filter results where finished_at is less than or equal to value (optional)
	finishedAtRange := []time.Time{time.Now()} // []time.Time | Filter results where finished_at is between two comma separated values (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	loggingCid := "loggingCid_example" // string | Filter results where logging_cid matches value (optional)
	loggingCidContains := "loggingCidContains_example" // string | Filter results where logging_cid contains value (optional)
	name := "name_example" // string | Filter results where name matches value (optional)
	nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
	nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
	nameNe := "nameNe_example" // string | Filter results where name not equal to value (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `state` - State* `-state` - State (descending)* `name` - Name* `-name` - Name (descending)* `logging_cid` - Logging cid* `-logging_cid` - Logging cid (descending)* `unblocked_at` - Unblocked at* `-unblocked_at` - Unblocked at (descending)* `started_at` - Started at* `-started_at` - Started at (descending)* `finished_at` - Finished at* `-finished_at` - Finished at (descending)* `error` - Error* `-error` - Error (descending)* `enc_args` - Enc args* `-enc_args` - Enc args (descending)* `enc_kwargs` - Enc kwargs* `-enc_kwargs` - Enc kwargs (descending)* `reserved_resources_record` - Reserved resources record* `-reserved_resources_record` - Reserved resources record (descending)* `versions` - Versions* `-versions` - Versions (descending)* `profile_options` - Profile options* `-profile_options` - Profile options (descending)* `immediate` - Immediate* `-immediate` - Immediate (descending)* `deferred` - Deferred* `-deferred` - Deferred (descending)* `result` - Result* `-result` - Result (descending)* `pulp_api_version` - Pulp api version* `-pulp_api_version` - Pulp api version (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	parentTask := "parentTask_example" // string |  (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpCreated := time.Now() // time.Time | Filter results where pulp_created matches value (optional)
	pulpCreatedGt := time.Now() // time.Time | Filter results where pulp_created is greater than value (optional)
	pulpCreatedGte := time.Now() // time.Time | Filter results where pulp_created is greater than or equal to value (optional)
	pulpCreatedIsnull := true // bool | Filter results where pulp_created has a null value (optional)
	pulpCreatedLt := time.Now() // time.Time | Filter results where pulp_created is less than value (optional)
	pulpCreatedLte := time.Now() // time.Time | Filter results where pulp_created is less than or equal to value (optional)
	pulpCreatedRange := []time.Time{time.Now()} // []time.Time | Filter results where pulp_created is between two comma separated values (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	reservedResources := "reservedResources_example" // string |  (optional)
	reservedResourcesIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	sharedResources := "sharedResources_example" // string |  (optional)
	sharedResourcesIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	startedAt := time.Now() // time.Time | Filter results where started_at matches value (optional)
	startedAtGt := time.Now() // time.Time | Filter results where started_at is greater than value (optional)
	startedAtGte := time.Now() // time.Time | Filter results where started_at is greater than or equal to value (optional)
	startedAtIsnull := true // bool | Filter results where started_at has a null value (optional)
	startedAtLt := time.Now() // time.Time | Filter results where started_at is less than value (optional)
	startedAtLte := time.Now() // time.Time | Filter results where started_at is less than or equal to value (optional)
	startedAtRange := []time.Time{time.Now()} // []time.Time | Filter results where started_at is between two comma separated values (optional)
	state := "state_example" // string | Filter results where state matches value* `waiting` - Waiting* `skipped` - Skipped* `running` - Running* `completed` - Completed* `failed` - Failed* `canceled` - Canceled* `canceling` - Canceling (optional)
	stateIn := []string{"Inner_example"} // []string | Filter results where state is in a comma-separated list of values (optional)
	stateNe := "stateNe_example" // string | Filter results where state not equal to value (optional)
	taskGroup := "taskGroup_example" // string |  (optional)
	unblockedAt := time.Now() // time.Time | Filter results where unblocked_at matches value (optional)
	unblockedAtGt := time.Now() // time.Time | Filter results where unblocked_at is greater than value (optional)
	unblockedAtGte := time.Now() // time.Time | Filter results where unblocked_at is greater than or equal to value (optional)
	unblockedAtIsnull := true // bool | Filter results where unblocked_at has a null value (optional)
	unblockedAtLt := time.Now() // time.Time | Filter results where unblocked_at is less than value (optional)
	unblockedAtLte := time.Now() // time.Time | Filter results where unblocked_at is less than or equal to value (optional)
	unblockedAtRange := []time.Time{time.Now()} // []time.Time | Filter results where unblocked_at is between two comma separated values (optional)
	worker := "worker_example" // string |  (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.AdminTasks(context.Background()).XTaskDiagnostics(xTaskDiagnostics).ChildTasks(childTasks).CreatedResources(createdResources).ExclusiveResources(exclusiveResources).ExclusiveResourcesIn(exclusiveResourcesIn).FinishedAt(finishedAt).FinishedAtGt(finishedAtGt).FinishedAtGte(finishedAtGte).FinishedAtIsnull(finishedAtIsnull).FinishedAtLt(finishedAtLt).FinishedAtLte(finishedAtLte).FinishedAtRange(finishedAtRange).Limit(limit).LoggingCid(loggingCid).LoggingCidContains(loggingCidContains).Name(name).NameContains(nameContains).NameIn(nameIn).NameNe(nameNe).Offset(offset).Ordering(ordering).ParentTask(parentTask).PrnIn(prnIn).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedIsnull(pulpCreatedIsnull).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).ReservedResources(reservedResources).ReservedResourcesIn(reservedResourcesIn).SharedResources(sharedResources).SharedResourcesIn(sharedResourcesIn).StartedAt(startedAt).StartedAtGt(startedAtGt).StartedAtGte(startedAtGte).StartedAtIsnull(startedAtIsnull).StartedAtLt(startedAtLt).StartedAtLte(startedAtLte).StartedAtRange(startedAtRange).State(state).StateIn(stateIn).StateNe(stateNe).TaskGroup(taskGroup).UnblockedAt(unblockedAt).UnblockedAtGt(unblockedAtGt).UnblockedAtGte(unblockedAtGte).UnblockedAtIsnull(unblockedAtIsnull).UnblockedAtLt(unblockedAtLt).UnblockedAtLte(unblockedAtLte).UnblockedAtRange(unblockedAtRange).Worker(worker).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.AdminTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminTasks`: PaginatedTaskResponseList
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.AdminTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **childTasks** | **string** |  | 
 **createdResources** | **string** |  | 
 **exclusiveResources** | **string** |  | 
 **exclusiveResourcesIn** | **[]string** | Multiple values may be separated by commas. | 
 **finishedAt** | **time.Time** | Filter results where finished_at matches value | 
 **finishedAtGt** | **time.Time** | Filter results where finished_at is greater than value | 
 **finishedAtGte** | **time.Time** | Filter results where finished_at is greater than or equal to value | 
 **finishedAtIsnull** | **bool** | Filter results where finished_at has a null value | 
 **finishedAtLt** | **time.Time** | Filter results where finished_at is less than value | 
 **finishedAtLte** | **time.Time** | Filter results where finished_at is less than or equal to value | 
 **finishedAtRange** | [**[]time.Time**](time.Time.md) | Filter results where finished_at is between two comma separated values | 
 **limit** | **int32** | Number of results to return per page. | 
 **loggingCid** | **string** | Filter results where logging_cid matches value | 
 **loggingCidContains** | **string** | Filter results where logging_cid contains value | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameNe** | **string** | Filter results where name not equal to value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;state&#x60; - State* &#x60;-state&#x60; - State (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;logging_cid&#x60; - Logging cid* &#x60;-logging_cid&#x60; - Logging cid (descending)* &#x60;unblocked_at&#x60; - Unblocked at* &#x60;-unblocked_at&#x60; - Unblocked at (descending)* &#x60;started_at&#x60; - Started at* &#x60;-started_at&#x60; - Started at (descending)* &#x60;finished_at&#x60; - Finished at* &#x60;-finished_at&#x60; - Finished at (descending)* &#x60;error&#x60; - Error* &#x60;-error&#x60; - Error (descending)* &#x60;enc_args&#x60; - Enc args* &#x60;-enc_args&#x60; - Enc args (descending)* &#x60;enc_kwargs&#x60; - Enc kwargs* &#x60;-enc_kwargs&#x60; - Enc kwargs (descending)* &#x60;reserved_resources_record&#x60; - Reserved resources record* &#x60;-reserved_resources_record&#x60; - Reserved resources record (descending)* &#x60;versions&#x60; - Versions* &#x60;-versions&#x60; - Versions (descending)* &#x60;profile_options&#x60; - Profile options* &#x60;-profile_options&#x60; - Profile options (descending)* &#x60;immediate&#x60; - Immediate* &#x60;-immediate&#x60; - Immediate (descending)* &#x60;deferred&#x60; - Deferred* &#x60;-deferred&#x60; - Deferred (descending)* &#x60;result&#x60; - Result* &#x60;-result&#x60; - Result (descending)* &#x60;pulp_api_version&#x60; - Pulp api version* &#x60;-pulp_api_version&#x60; - Pulp api version (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **parentTask** | **string** |  | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpCreated** | **time.Time** | Filter results where pulp_created matches value | 
 **pulpCreatedGt** | **time.Time** | Filter results where pulp_created is greater than value | 
 **pulpCreatedGte** | **time.Time** | Filter results where pulp_created is greater than or equal to value | 
 **pulpCreatedIsnull** | **bool** | Filter results where pulp_created has a null value | 
 **pulpCreatedLt** | **time.Time** | Filter results where pulp_created is less than value | 
 **pulpCreatedLte** | **time.Time** | Filter results where pulp_created is less than or equal to value | 
 **pulpCreatedRange** | [**[]time.Time**](time.Time.md) | Filter results where pulp_created is between two comma separated values | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **reservedResources** | **string** |  | 
 **reservedResourcesIn** | **[]string** | Multiple values may be separated by commas. | 
 **sharedResources** | **string** |  | 
 **sharedResourcesIn** | **[]string** | Multiple values may be separated by commas. | 
 **startedAt** | **time.Time** | Filter results where started_at matches value | 
 **startedAtGt** | **time.Time** | Filter results where started_at is greater than value | 
 **startedAtGte** | **time.Time** | Filter results where started_at is greater than or equal to value | 
 **startedAtIsnull** | **bool** | Filter results where started_at has a null value | 
 **startedAtLt** | **time.Time** | Filter results where started_at is less than value | 
 **startedAtLte** | **time.Time** | Filter results where started_at is less than or equal to value | 
 **startedAtRange** | [**[]time.Time**](time.Time.md) | Filter results where started_at is between two comma separated values | 
 **state** | **string** | Filter results where state matches value* &#x60;waiting&#x60; - Waiting* &#x60;skipped&#x60; - Skipped* &#x60;running&#x60; - Running* &#x60;completed&#x60; - Completed* &#x60;failed&#x60; - Failed* &#x60;canceled&#x60; - Canceled* &#x60;canceling&#x60; - Canceling | 
 **stateIn** | **[]string** | Filter results where state is in a comma-separated list of values | 
 **stateNe** | **string** | Filter results where state not equal to value | 
 **taskGroup** | **string** |  | 
 **unblockedAt** | **time.Time** | Filter results where unblocked_at matches value | 
 **unblockedAtGt** | **time.Time** | Filter results where unblocked_at is greater than value | 
 **unblockedAtGte** | **time.Time** | Filter results where unblocked_at is greater than or equal to value | 
 **unblockedAtIsnull** | **bool** | Filter results where unblocked_at has a null value | 
 **unblockedAtLt** | **time.Time** | Filter results where unblocked_at is less than value | 
 **unblockedAtLte** | **time.Time** | Filter results where unblocked_at is less than or equal to value | 
 **unblockedAtRange** | [**[]time.Time**](time.Time.md) | Filter results where unblocked_at is between two comma separated values | 
 **worker** | **string** |  | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedTaskResponseList**](PaginatedTaskResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksAddRole

> NestedRoleResponse TasksAddRole(ctx, taskHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	taskHref := "taskHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksAddRole(context.Background(), taskHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksAddRoleRequest struct via the builder pattern


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


## TasksCancel

> TaskResponse TasksCancel(ctx, taskHref).PatchedTaskCancel(patchedTaskCancel).XTaskDiagnostics(xTaskDiagnostics).Execute()

Cancel a task



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
	taskHref := "taskHref_example" // string | 
	patchedTaskCancel := *openapiclient.NewPatchedTaskCancel() // PatchedTaskCancel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksCancel(context.Background(), taskHref).PatchedTaskCancel(patchedTaskCancel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksCancel`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedTaskCancel** | [**PatchedTaskCancel**](PatchedTaskCancel.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksDelete

> TasksDelete(ctx, taskHref).XTaskDiagnostics(xTaskDiagnostics).Execute()

Delete a task



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
	taskHref := "taskHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TasksAPI.TasksDelete(context.Background(), taskHref).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksList

> PaginatedTaskResponseList TasksList(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).ChildTasks(childTasks).CreatedResources(createdResources).ExclusiveResources(exclusiveResources).ExclusiveResourcesIn(exclusiveResourcesIn).FinishedAt(finishedAt).FinishedAtGt(finishedAtGt).FinishedAtGte(finishedAtGte).FinishedAtIsnull(finishedAtIsnull).FinishedAtLt(finishedAtLt).FinishedAtLte(finishedAtLte).FinishedAtRange(finishedAtRange).Limit(limit).LoggingCid(loggingCid).LoggingCidContains(loggingCidContains).Name(name).NameContains(nameContains).NameIn(nameIn).NameNe(nameNe).Offset(offset).Ordering(ordering).ParentTask(parentTask).PrnIn(prnIn).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedIsnull(pulpCreatedIsnull).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).ReservedResources(reservedResources).ReservedResourcesIn(reservedResourcesIn).SharedResources(sharedResources).SharedResourcesIn(sharedResourcesIn).StartedAt(startedAt).StartedAtGt(startedAtGt).StartedAtGte(startedAtGte).StartedAtIsnull(startedAtIsnull).StartedAtLt(startedAtLt).StartedAtLte(startedAtLte).StartedAtRange(startedAtRange).State(state).StateIn(stateIn).StateNe(stateNe).TaskGroup(taskGroup).UnblockedAt(unblockedAt).UnblockedAtGt(unblockedAtGt).UnblockedAtGte(unblockedAtGte).UnblockedAtIsnull(unblockedAtIsnull).UnblockedAtLt(unblockedAtLt).UnblockedAtLte(unblockedAtLte).UnblockedAtRange(unblockedAtRange).Worker(worker).Fields(fields).ExcludeFields(excludeFields).Execute()

List tasks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	childTasks := "childTasks_example" // string |  (optional)
	createdResources := "createdResources_example" // string |  (optional)
	exclusiveResources := "exclusiveResources_example" // string |  (optional)
	exclusiveResourcesIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	finishedAt := time.Now() // time.Time | Filter results where finished_at matches value (optional)
	finishedAtGt := time.Now() // time.Time | Filter results where finished_at is greater than value (optional)
	finishedAtGte := time.Now() // time.Time | Filter results where finished_at is greater than or equal to value (optional)
	finishedAtIsnull := true // bool | Filter results where finished_at has a null value (optional)
	finishedAtLt := time.Now() // time.Time | Filter results where finished_at is less than value (optional)
	finishedAtLte := time.Now() // time.Time | Filter results where finished_at is less than or equal to value (optional)
	finishedAtRange := []time.Time{time.Now()} // []time.Time | Filter results where finished_at is between two comma separated values (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	loggingCid := "loggingCid_example" // string | Filter results where logging_cid matches value (optional)
	loggingCidContains := "loggingCidContains_example" // string | Filter results where logging_cid contains value (optional)
	name := "name_example" // string | Filter results where name matches value (optional)
	nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
	nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
	nameNe := "nameNe_example" // string | Filter results where name not equal to value (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `state` - State* `-state` - State (descending)* `name` - Name* `-name` - Name (descending)* `logging_cid` - Logging cid* `-logging_cid` - Logging cid (descending)* `unblocked_at` - Unblocked at* `-unblocked_at` - Unblocked at (descending)* `started_at` - Started at* `-started_at` - Started at (descending)* `finished_at` - Finished at* `-finished_at` - Finished at (descending)* `error` - Error* `-error` - Error (descending)* `enc_args` - Enc args* `-enc_args` - Enc args (descending)* `enc_kwargs` - Enc kwargs* `-enc_kwargs` - Enc kwargs (descending)* `reserved_resources_record` - Reserved resources record* `-reserved_resources_record` - Reserved resources record (descending)* `versions` - Versions* `-versions` - Versions (descending)* `profile_options` - Profile options* `-profile_options` - Profile options (descending)* `immediate` - Immediate* `-immediate` - Immediate (descending)* `deferred` - Deferred* `-deferred` - Deferred (descending)* `result` - Result* `-result` - Result (descending)* `pulp_api_version` - Pulp api version* `-pulp_api_version` - Pulp api version (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	parentTask := "parentTask_example" // string |  (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpCreated := time.Now() // time.Time | Filter results where pulp_created matches value (optional)
	pulpCreatedGt := time.Now() // time.Time | Filter results where pulp_created is greater than value (optional)
	pulpCreatedGte := time.Now() // time.Time | Filter results where pulp_created is greater than or equal to value (optional)
	pulpCreatedIsnull := true // bool | Filter results where pulp_created has a null value (optional)
	pulpCreatedLt := time.Now() // time.Time | Filter results where pulp_created is less than value (optional)
	pulpCreatedLte := time.Now() // time.Time | Filter results where pulp_created is less than or equal to value (optional)
	pulpCreatedRange := []time.Time{time.Now()} // []time.Time | Filter results where pulp_created is between two comma separated values (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	reservedResources := "reservedResources_example" // string |  (optional)
	reservedResourcesIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	sharedResources := "sharedResources_example" // string |  (optional)
	sharedResourcesIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	startedAt := time.Now() // time.Time | Filter results where started_at matches value (optional)
	startedAtGt := time.Now() // time.Time | Filter results where started_at is greater than value (optional)
	startedAtGte := time.Now() // time.Time | Filter results where started_at is greater than or equal to value (optional)
	startedAtIsnull := true // bool | Filter results where started_at has a null value (optional)
	startedAtLt := time.Now() // time.Time | Filter results where started_at is less than value (optional)
	startedAtLte := time.Now() // time.Time | Filter results where started_at is less than or equal to value (optional)
	startedAtRange := []time.Time{time.Now()} // []time.Time | Filter results where started_at is between two comma separated values (optional)
	state := "state_example" // string | Filter results where state matches value* `waiting` - Waiting* `skipped` - Skipped* `running` - Running* `completed` - Completed* `failed` - Failed* `canceled` - Canceled* `canceling` - Canceling (optional)
	stateIn := []string{"Inner_example"} // []string | Filter results where state is in a comma-separated list of values (optional)
	stateNe := "stateNe_example" // string | Filter results where state not equal to value (optional)
	taskGroup := "taskGroup_example" // string |  (optional)
	unblockedAt := time.Now() // time.Time | Filter results where unblocked_at matches value (optional)
	unblockedAtGt := time.Now() // time.Time | Filter results where unblocked_at is greater than value (optional)
	unblockedAtGte := time.Now() // time.Time | Filter results where unblocked_at is greater than or equal to value (optional)
	unblockedAtIsnull := true // bool | Filter results where unblocked_at has a null value (optional)
	unblockedAtLt := time.Now() // time.Time | Filter results where unblocked_at is less than value (optional)
	unblockedAtLte := time.Now() // time.Time | Filter results where unblocked_at is less than or equal to value (optional)
	unblockedAtRange := []time.Time{time.Now()} // []time.Time | Filter results where unblocked_at is between two comma separated values (optional)
	worker := "worker_example" // string |  (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksList(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).ChildTasks(childTasks).CreatedResources(createdResources).ExclusiveResources(exclusiveResources).ExclusiveResourcesIn(exclusiveResourcesIn).FinishedAt(finishedAt).FinishedAtGt(finishedAtGt).FinishedAtGte(finishedAtGte).FinishedAtIsnull(finishedAtIsnull).FinishedAtLt(finishedAtLt).FinishedAtLte(finishedAtLte).FinishedAtRange(finishedAtRange).Limit(limit).LoggingCid(loggingCid).LoggingCidContains(loggingCidContains).Name(name).NameContains(nameContains).NameIn(nameIn).NameNe(nameNe).Offset(offset).Ordering(ordering).ParentTask(parentTask).PrnIn(prnIn).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedIsnull(pulpCreatedIsnull).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).ReservedResources(reservedResources).ReservedResourcesIn(reservedResourcesIn).SharedResources(sharedResources).SharedResourcesIn(sharedResourcesIn).StartedAt(startedAt).StartedAtGt(startedAtGt).StartedAtGte(startedAtGte).StartedAtIsnull(startedAtIsnull).StartedAtLt(startedAtLt).StartedAtLte(startedAtLte).StartedAtRange(startedAtRange).State(state).StateIn(stateIn).StateNe(stateNe).TaskGroup(taskGroup).UnblockedAt(unblockedAt).UnblockedAtGt(unblockedAtGt).UnblockedAtGte(unblockedAtGte).UnblockedAtIsnull(unblockedAtIsnull).UnblockedAtLt(unblockedAtLt).UnblockedAtLte(unblockedAtLte).UnblockedAtRange(unblockedAtRange).Worker(worker).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksList`: PaginatedTaskResponseList
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **childTasks** | **string** |  | 
 **createdResources** | **string** |  | 
 **exclusiveResources** | **string** |  | 
 **exclusiveResourcesIn** | **[]string** | Multiple values may be separated by commas. | 
 **finishedAt** | **time.Time** | Filter results where finished_at matches value | 
 **finishedAtGt** | **time.Time** | Filter results where finished_at is greater than value | 
 **finishedAtGte** | **time.Time** | Filter results where finished_at is greater than or equal to value | 
 **finishedAtIsnull** | **bool** | Filter results where finished_at has a null value | 
 **finishedAtLt** | **time.Time** | Filter results where finished_at is less than value | 
 **finishedAtLte** | **time.Time** | Filter results where finished_at is less than or equal to value | 
 **finishedAtRange** | [**[]time.Time**](time.Time.md) | Filter results where finished_at is between two comma separated values | 
 **limit** | **int32** | Number of results to return per page. | 
 **loggingCid** | **string** | Filter results where logging_cid matches value | 
 **loggingCidContains** | **string** | Filter results where logging_cid contains value | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameNe** | **string** | Filter results where name not equal to value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;state&#x60; - State* &#x60;-state&#x60; - State (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;logging_cid&#x60; - Logging cid* &#x60;-logging_cid&#x60; - Logging cid (descending)* &#x60;unblocked_at&#x60; - Unblocked at* &#x60;-unblocked_at&#x60; - Unblocked at (descending)* &#x60;started_at&#x60; - Started at* &#x60;-started_at&#x60; - Started at (descending)* &#x60;finished_at&#x60; - Finished at* &#x60;-finished_at&#x60; - Finished at (descending)* &#x60;error&#x60; - Error* &#x60;-error&#x60; - Error (descending)* &#x60;enc_args&#x60; - Enc args* &#x60;-enc_args&#x60; - Enc args (descending)* &#x60;enc_kwargs&#x60; - Enc kwargs* &#x60;-enc_kwargs&#x60; - Enc kwargs (descending)* &#x60;reserved_resources_record&#x60; - Reserved resources record* &#x60;-reserved_resources_record&#x60; - Reserved resources record (descending)* &#x60;versions&#x60; - Versions* &#x60;-versions&#x60; - Versions (descending)* &#x60;profile_options&#x60; - Profile options* &#x60;-profile_options&#x60; - Profile options (descending)* &#x60;immediate&#x60; - Immediate* &#x60;-immediate&#x60; - Immediate (descending)* &#x60;deferred&#x60; - Deferred* &#x60;-deferred&#x60; - Deferred (descending)* &#x60;result&#x60; - Result* &#x60;-result&#x60; - Result (descending)* &#x60;pulp_api_version&#x60; - Pulp api version* &#x60;-pulp_api_version&#x60; - Pulp api version (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **parentTask** | **string** |  | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpCreated** | **time.Time** | Filter results where pulp_created matches value | 
 **pulpCreatedGt** | **time.Time** | Filter results where pulp_created is greater than value | 
 **pulpCreatedGte** | **time.Time** | Filter results where pulp_created is greater than or equal to value | 
 **pulpCreatedIsnull** | **bool** | Filter results where pulp_created has a null value | 
 **pulpCreatedLt** | **time.Time** | Filter results where pulp_created is less than value | 
 **pulpCreatedLte** | **time.Time** | Filter results where pulp_created is less than or equal to value | 
 **pulpCreatedRange** | [**[]time.Time**](time.Time.md) | Filter results where pulp_created is between two comma separated values | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **reservedResources** | **string** |  | 
 **reservedResourcesIn** | **[]string** | Multiple values may be separated by commas. | 
 **sharedResources** | **string** |  | 
 **sharedResourcesIn** | **[]string** | Multiple values may be separated by commas. | 
 **startedAt** | **time.Time** | Filter results where started_at matches value | 
 **startedAtGt** | **time.Time** | Filter results where started_at is greater than value | 
 **startedAtGte** | **time.Time** | Filter results where started_at is greater than or equal to value | 
 **startedAtIsnull** | **bool** | Filter results where started_at has a null value | 
 **startedAtLt** | **time.Time** | Filter results where started_at is less than value | 
 **startedAtLte** | **time.Time** | Filter results where started_at is less than or equal to value | 
 **startedAtRange** | [**[]time.Time**](time.Time.md) | Filter results where started_at is between two comma separated values | 
 **state** | **string** | Filter results where state matches value* &#x60;waiting&#x60; - Waiting* &#x60;skipped&#x60; - Skipped* &#x60;running&#x60; - Running* &#x60;completed&#x60; - Completed* &#x60;failed&#x60; - Failed* &#x60;canceled&#x60; - Canceled* &#x60;canceling&#x60; - Canceling | 
 **stateIn** | **[]string** | Filter results where state is in a comma-separated list of values | 
 **stateNe** | **string** | Filter results where state not equal to value | 
 **taskGroup** | **string** |  | 
 **unblockedAt** | **time.Time** | Filter results where unblocked_at matches value | 
 **unblockedAtGt** | **time.Time** | Filter results where unblocked_at is greater than value | 
 **unblockedAtGte** | **time.Time** | Filter results where unblocked_at is greater than or equal to value | 
 **unblockedAtIsnull** | **bool** | Filter results where unblocked_at has a null value | 
 **unblockedAtLt** | **time.Time** | Filter results where unblocked_at is less than value | 
 **unblockedAtLte** | **time.Time** | Filter results where unblocked_at is less than or equal to value | 
 **unblockedAtRange** | [**[]time.Time**](time.Time.md) | Filter results where unblocked_at is between two comma separated values | 
 **worker** | **string** |  | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedTaskResponseList**](PaginatedTaskResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksListRoles

> ObjectRolesResponse TasksListRoles(ctx, taskHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	taskHref := "taskHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksListRoles(context.Background(), taskHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksListRolesRequest struct via the builder pattern


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


## TasksMyPermissions

> MyPermissionsResponse TasksMyPermissions(ctx, taskHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	taskHref := "taskHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksMyPermissions(context.Background(), taskHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksMyPermissionsRequest struct via the builder pattern


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


## TasksProfileArtifacts

> ProfileArtifactResponse TasksProfileArtifacts(ctx, taskHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Fetch downloadable links for profile artifacts



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
	taskHref := "taskHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksProfileArtifacts(context.Background(), taskHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksProfileArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksProfileArtifacts`: ProfileArtifactResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksProfileArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksProfileArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ProfileArtifactResponse**](ProfileArtifactResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksPurge

> AsyncOperationResponse TasksPurge(ctx, pulpDomain).Purge(purge).XTaskDiagnostics(xTaskDiagnostics).Execute()

Purge Completed Tasks



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
	purge := *openapiclient.NewPurge() // Purge | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksPurge(context.Background(), pulpDomain).Purge(purge).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksPurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksPurge`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksPurge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksPurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purge** | [**Purge**](Purge.md) |  | 
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


## TasksRead

> TaskResponse TasksRead(ctx, taskHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a task



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
	taskHref := "taskHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksRead(context.Background(), taskHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksRead`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksRemoveRole

> NestedRoleResponse TasksRemoveRole(ctx, taskHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	taskHref := "taskHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksRemoveRole(context.Background(), taskHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksRemoveRoleRequest struct via the builder pattern


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

