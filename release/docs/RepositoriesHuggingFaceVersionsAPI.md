# \RepositoriesHuggingFaceVersionsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesHuggingFaceHuggingFaceVersionsDelete**](RepositoriesHuggingFaceVersionsAPI.md#RepositoriesHuggingFaceHuggingFaceVersionsDelete) | **Delete** /{hugging_face_hugging_face_repository_version_href} | Delete a repository version
[**RepositoriesHuggingFaceHuggingFaceVersionsList**](RepositoriesHuggingFaceVersionsAPI.md#RepositoriesHuggingFaceHuggingFaceVersionsList) | **Get** /{hugging_face_hugging_face_repository_href}versions/ | List repository versions
[**RepositoriesHuggingFaceHuggingFaceVersionsRead**](RepositoriesHuggingFaceVersionsAPI.md#RepositoriesHuggingFaceHuggingFaceVersionsRead) | **Get** /{hugging_face_hugging_face_repository_version_href} | Inspect a repository version
[**RepositoriesHuggingFaceHuggingFaceVersionsRepair**](RepositoriesHuggingFaceVersionsAPI.md#RepositoriesHuggingFaceHuggingFaceVersionsRepair) | **Post** /{hugging_face_hugging_face_repository_version_href}repair/ | 



## RepositoriesHuggingFaceHuggingFaceVersionsDelete

> AsyncOperationResponse RepositoriesHuggingFaceHuggingFaceVersionsDelete(ctx, huggingFaceHuggingFaceRepositoryVersionHref).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	huggingFaceHuggingFaceRepositoryVersionHref := "huggingFaceHuggingFaceRepositoryVersionHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesHuggingFaceVersionsAPI.RepositoriesHuggingFaceHuggingFaceVersionsDelete(context.Background(), huggingFaceHuggingFaceRepositoryVersionHref).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesHuggingFaceVersionsAPI.RepositoriesHuggingFaceHuggingFaceVersionsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesHuggingFaceHuggingFaceVersionsDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesHuggingFaceVersionsAPI.RepositoriesHuggingFaceHuggingFaceVersionsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceRepositoryVersionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesHuggingFaceHuggingFaceVersionsDeleteRequest struct via the builder pattern


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


## RepositoriesHuggingFaceHuggingFaceVersionsList

> PaginatedRepositoryVersionResponseList RepositoriesHuggingFaceHuggingFaceVersionsList(ctx, huggingFaceHuggingFaceRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Content(content).ContentIn(contentIn).Limit(limit).Number(number).NumberGt(numberGt).NumberGte(numberGte).NumberLt(numberLt).NumberLte(numberLte).NumberRange(numberRange).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedIsnull(pulpCreatedIsnull).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).PulpHrefIn(pulpHrefIn).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	huggingFaceHuggingFaceRepositoryHref := "huggingFaceHuggingFaceRepositoryHref_example" // string | 
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
	resp, r, err := apiClient.RepositoriesHuggingFaceVersionsAPI.RepositoriesHuggingFaceHuggingFaceVersionsList(context.Background(), huggingFaceHuggingFaceRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Content(content).ContentIn(contentIn).Limit(limit).Number(number).NumberGt(numberGt).NumberGte(numberGte).NumberLt(numberLt).NumberLte(numberLte).NumberRange(numberRange).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedIsnull(pulpCreatedIsnull).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).PulpHrefIn(pulpHrefIn).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesHuggingFaceVersionsAPI.RepositoriesHuggingFaceHuggingFaceVersionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesHuggingFaceHuggingFaceVersionsList`: PaginatedRepositoryVersionResponseList
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesHuggingFaceVersionsAPI.RepositoriesHuggingFaceHuggingFaceVersionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesHuggingFaceHuggingFaceVersionsListRequest struct via the builder pattern


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


## RepositoriesHuggingFaceHuggingFaceVersionsRead

> RepositoryVersionResponse RepositoriesHuggingFaceHuggingFaceVersionsRead(ctx, huggingFaceHuggingFaceRepositoryVersionHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	huggingFaceHuggingFaceRepositoryVersionHref := "huggingFaceHuggingFaceRepositoryVersionHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesHuggingFaceVersionsAPI.RepositoriesHuggingFaceHuggingFaceVersionsRead(context.Background(), huggingFaceHuggingFaceRepositoryVersionHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesHuggingFaceVersionsAPI.RepositoriesHuggingFaceHuggingFaceVersionsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesHuggingFaceHuggingFaceVersionsRead`: RepositoryVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesHuggingFaceVersionsAPI.RepositoriesHuggingFaceHuggingFaceVersionsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceRepositoryVersionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesHuggingFaceHuggingFaceVersionsReadRequest struct via the builder pattern


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


## RepositoriesHuggingFaceHuggingFaceVersionsRepair

> AsyncOperationResponse RepositoriesHuggingFaceHuggingFaceVersionsRepair(ctx, huggingFaceHuggingFaceRepositoryVersionHref).Repair(repair).XTaskDiagnostics(xTaskDiagnostics).Execute()





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
	huggingFaceHuggingFaceRepositoryVersionHref := "huggingFaceHuggingFaceRepositoryVersionHref_example" // string | 
	repair := *openapiclient.NewRepair() // Repair | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesHuggingFaceVersionsAPI.RepositoriesHuggingFaceHuggingFaceVersionsRepair(context.Background(), huggingFaceHuggingFaceRepositoryVersionHref).Repair(repair).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesHuggingFaceVersionsAPI.RepositoriesHuggingFaceHuggingFaceVersionsRepair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesHuggingFaceHuggingFaceVersionsRepair`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesHuggingFaceVersionsAPI.RepositoriesHuggingFaceHuggingFaceVersionsRepair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFaceRepositoryVersionHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesHuggingFaceHuggingFaceVersionsRepairRequest struct via the builder pattern


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

