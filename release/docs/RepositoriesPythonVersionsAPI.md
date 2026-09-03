# \RepositoriesPythonVersionsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesPythonPythonVersionsDelete**](RepositoriesPythonVersionsAPI.md#RepositoriesPythonPythonVersionsDelete) | **Delete** /{python_python_repository_version_href} | Delete a repository version
[**RepositoriesPythonPythonVersionsList**](RepositoriesPythonVersionsAPI.md#RepositoriesPythonPythonVersionsList) | **Get** /{python_python_repository_href}versions/ | List repository versions
[**RepositoriesPythonPythonVersionsRead**](RepositoriesPythonVersionsAPI.md#RepositoriesPythonPythonVersionsRead) | **Get** /{python_python_repository_version_href} | Inspect a repository version
[**RepositoriesPythonPythonVersionsRepair**](RepositoriesPythonVersionsAPI.md#RepositoriesPythonPythonVersionsRepair) | **Post** /{python_python_repository_version_href}repair/ | 
[**RepositoriesPythonPythonVersionsScan**](RepositoriesPythonVersionsAPI.md#RepositoriesPythonPythonVersionsScan) | **Post** /{python_python_repository_version_href}scan/ | Generate vulnerability report



## RepositoriesPythonPythonVersionsDelete

> AsyncOperationResponse RepositoriesPythonPythonVersionsDelete(ctx, pythonPythonRepositoryVersionHref).XTaskDiagnostics(xTaskDiagnostics).Execute()

Delete a repository version



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
	pythonPythonRepositoryVersionHref := "pythonPythonRepositoryVersionHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsDelete(context.Background(), pythonPythonRepositoryVersionHref).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesPythonPythonVersionsDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pythonPythonRepositoryVersionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesPythonPythonVersionsDeleteRequest struct via the builder pattern


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


## RepositoriesPythonPythonVersionsList

> PaginatedRepositoryVersionResponseList RepositoriesPythonPythonVersionsList(ctx, pythonPythonRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Content(content).ContentIn(contentIn).Limit(limit).Number(number).NumberGt(numberGt).NumberGte(numberGte).NumberLt(numberLt).NumberLte(numberLte).NumberRange(numberRange).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedIsnull(pulpCreatedIsnull).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).PulpHrefIn(pulpHrefIn).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()

List repository versions



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
	pythonPythonRepositoryHref := "pythonPythonRepositoryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	content := "content_example" // string | Content Unit referenced by HREF/PRN (optional)
	contentIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	number := int32(56) // int32 | Filter results where number matches value (optional)
	numberGt := int32(56) // int32 | Filter results where number is greater than value (optional)
	numberGte := int32(56) // int32 | Filter results where number is greater than or equal to value (optional)
	numberLt := int32(56) // int32 | Filter results where number is less than value (optional)
	numberLte := int32(56) // int32 | Filter results where number is less than or equal to value (optional)
	numberRange := []int32{int32(123)} // []int32 | Filter results where number is between two comma separated values (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `number` - Number* `-number` - Number (descending)* `complete` - Complete* `-complete` - Complete (descending)* `info` - Info* `-info` - Info (descending)* `content_ids` - Content ids* `-content_ids` - Content ids (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpCreated := time.Now() // time.Time | Filter results where pulp_created matches value (optional)
	pulpCreatedGt := time.Now() // time.Time | Filter results where pulp_created is greater than value (optional)
	pulpCreatedGte := time.Now() // time.Time | Filter results where pulp_created is greater than or equal to value (optional)
	pulpCreatedIsnull := true // bool | Filter results where pulp_created has a null value (optional)
	pulpCreatedLt := time.Now() // time.Time | Filter results where pulp_created is less than value (optional)
	pulpCreatedLte := time.Now() // time.Time | Filter results where pulp_created is less than or equal to value (optional)
	pulpCreatedRange := []time.Time{time.Now()} // []time.Time | Filter results where pulp_created is between two comma separated values (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsList(context.Background(), pythonPythonRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Content(content).ContentIn(contentIn).Limit(limit).Number(number).NumberGt(numberGt).NumberGte(numberGte).NumberLt(numberLt).NumberLte(numberLte).NumberRange(numberRange).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedIsnull(pulpCreatedIsnull).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).PulpHrefIn(pulpHrefIn).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesPythonPythonVersionsList`: PaginatedRepositoryVersionResponseList
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pythonPythonRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesPythonPythonVersionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **content** | **string** | Content Unit referenced by HREF/PRN | 
 **contentIn** | **[]string** | Multiple values may be separated by commas. | 
 **limit** | **int32** | Number of results to return per page. | 
 **number** | **int32** | Filter results where number matches value | 
 **numberGt** | **int32** | Filter results where number is greater than value | 
 **numberGte** | **int32** | Filter results where number is greater than or equal to value | 
 **numberLt** | **int32** | Filter results where number is less than value | 
 **numberLte** | **int32** | Filter results where number is less than or equal to value | 
 **numberRange** | **[]int32** | Filter results where number is between two comma separated values | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;number&#x60; - Number* &#x60;-number&#x60; - Number (descending)* &#x60;complete&#x60; - Complete* &#x60;-complete&#x60; - Complete (descending)* &#x60;info&#x60; - Info* &#x60;-info&#x60; - Info (descending)* &#x60;content_ids&#x60; - Content ids* &#x60;-content_ids&#x60; - Content ids (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpCreated** | **time.Time** | Filter results where pulp_created matches value | 
 **pulpCreatedGt** | **time.Time** | Filter results where pulp_created is greater than value | 
 **pulpCreatedGte** | **time.Time** | Filter results where pulp_created is greater than or equal to value | 
 **pulpCreatedIsnull** | **bool** | Filter results where pulp_created has a null value | 
 **pulpCreatedLt** | **time.Time** | Filter results where pulp_created is less than value | 
 **pulpCreatedLte** | **time.Time** | Filter results where pulp_created is less than or equal to value | 
 **pulpCreatedRange** | [**[]time.Time**](time.Time.md) | Filter results where pulp_created is between two comma separated values | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedRepositoryVersionResponseList**](PaginatedRepositoryVersionResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesPythonPythonVersionsRead

> RepositoryVersionResponse RepositoriesPythonPythonVersionsRead(ctx, pythonPythonRepositoryVersionHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a repository version



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
	pythonPythonRepositoryVersionHref := "pythonPythonRepositoryVersionHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsRead(context.Background(), pythonPythonRepositoryVersionHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesPythonPythonVersionsRead`: RepositoryVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pythonPythonRepositoryVersionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesPythonPythonVersionsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RepositoryVersionResponse**](RepositoryVersionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesPythonPythonVersionsRepair

> AsyncOperationResponse RepositoriesPythonPythonVersionsRepair(ctx, pythonPythonRepositoryVersionHref).Repair(repair).XTaskDiagnostics(xTaskDiagnostics).Execute()





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
	pythonPythonRepositoryVersionHref := "pythonPythonRepositoryVersionHref_example" // string | 
	repair := *openapiclient.NewRepair() // Repair | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsRepair(context.Background(), pythonPythonRepositoryVersionHref).Repair(repair).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsRepair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesPythonPythonVersionsRepair`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsRepair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pythonPythonRepositoryVersionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesPythonPythonVersionsRepairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repair** | [**Repair**](Repair.md) |  | 
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


## RepositoriesPythonPythonVersionsScan

> AsyncOperationResponse RepositoriesPythonPythonVersionsScan(ctx, pythonPythonRepositoryVersionHref).XTaskDiagnostics(xTaskDiagnostics).Execute()

Generate vulnerability report



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
	pythonPythonRepositoryVersionHref := "pythonPythonRepositoryVersionHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsScan(context.Background(), pythonPythonRepositoryVersionHref).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesPythonPythonVersionsScan`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pythonPythonRepositoryVersionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesPythonPythonVersionsScanRequest struct via the builder pattern


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

