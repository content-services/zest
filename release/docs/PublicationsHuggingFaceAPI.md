# \PublicationsHuggingFaceAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicationsHuggingFaceHuggingFaceCreate**](PublicationsHuggingFaceAPI.md#PublicationsHuggingFaceHuggingFaceCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/publications/hugging_face/hugging-face/ | Create a hugging face publication
[**PublicationsHuggingFaceHuggingFaceDelete**](PublicationsHuggingFaceAPI.md#PublicationsHuggingFaceHuggingFaceDelete) | **Delete** /{hugging_face_hugging_face_publication_href} | Delete a hugging face publication
[**PublicationsHuggingFaceHuggingFaceList**](PublicationsHuggingFaceAPI.md#PublicationsHuggingFaceHuggingFaceList) | **Get** /api/pulp/{pulp_domain}/api/v3/publications/hugging_face/hugging-face/ | List hugging face publications
[**PublicationsHuggingFaceHuggingFaceRead**](PublicationsHuggingFaceAPI.md#PublicationsHuggingFaceHuggingFaceRead) | **Get** /{hugging_face_hugging_face_publication_href} | Inspect a hugging face publication



## PublicationsHuggingFaceHuggingFaceCreate

> AsyncOperationResponse PublicationsHuggingFaceHuggingFaceCreate(ctx, pulpDomain).HuggingFaceHuggingFacePublication(huggingFaceHuggingFacePublication).XTaskDiagnostics(xTaskDiagnostics).Execute()

Create a hugging face publication



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
	huggingFaceHuggingFacePublication := *openapiclient.NewHuggingFaceHuggingFacePublication() // HuggingFaceHuggingFacePublication | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicationsHuggingFaceAPI.PublicationsHuggingFaceHuggingFaceCreate(context.Background(), pulpDomain).HuggingFaceHuggingFacePublication(huggingFaceHuggingFacePublication).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicationsHuggingFaceAPI.PublicationsHuggingFaceHuggingFaceCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicationsHuggingFaceHuggingFaceCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicationsHuggingFaceAPI.PublicationsHuggingFaceHuggingFaceCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicationsHuggingFaceHuggingFaceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **huggingFaceHuggingFacePublication** | [**HuggingFaceHuggingFacePublication**](HuggingFaceHuggingFacePublication.md) |  | 
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


## PublicationsHuggingFaceHuggingFaceDelete

> PublicationsHuggingFaceHuggingFaceDelete(ctx, huggingFaceHuggingFacePublicationHref).XTaskDiagnostics(xTaskDiagnostics).Execute()

Delete a hugging face publication



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
	huggingFaceHuggingFacePublicationHref := "huggingFaceHuggingFacePublicationHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicationsHuggingFaceAPI.PublicationsHuggingFaceHuggingFaceDelete(context.Background(), huggingFaceHuggingFacePublicationHref).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicationsHuggingFaceAPI.PublicationsHuggingFaceHuggingFaceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFacePublicationHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicationsHuggingFaceHuggingFaceDeleteRequest struct via the builder pattern


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


## PublicationsHuggingFaceHuggingFaceList

> PaginatedhuggingFaceHuggingFacePublicationResponseList PublicationsHuggingFaceHuggingFaceList(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Checkpoint(checkpoint).Content(content).ContentIn(contentIn).Limit(limit).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedIsnull(pulpCreatedIsnull).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).Repository(repository).RepositoryVersion(repositoryVersion).Fields(fields).ExcludeFields(excludeFields).Execute()

List hugging face publications



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
	checkpoint := true // bool | Filter results where checkpoint matches value (optional)
	content := "content_example" // string | Content Unit referenced by HREF/PRN (optional)
	contentIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `complete` - Complete* `-complete` - Complete (descending)* `pass_through` - Pass through* `-pass_through` - Pass through (descending)* `checkpoint` - Checkpoint* `-checkpoint` - Checkpoint (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
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
	repository := "repository_example" // string | Repository referenced by HREF/PRN (optional)
	repositoryVersion := "repositoryVersion_example" // string |  (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicationsHuggingFaceAPI.PublicationsHuggingFaceHuggingFaceList(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Checkpoint(checkpoint).Content(content).ContentIn(contentIn).Limit(limit).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpCreated(pulpCreated).PulpCreatedGt(pulpCreatedGt).PulpCreatedGte(pulpCreatedGte).PulpCreatedIsnull(pulpCreatedIsnull).PulpCreatedLt(pulpCreatedLt).PulpCreatedLte(pulpCreatedLte).PulpCreatedRange(pulpCreatedRange).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).Repository(repository).RepositoryVersion(repositoryVersion).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicationsHuggingFaceAPI.PublicationsHuggingFaceHuggingFaceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicationsHuggingFaceHuggingFaceList`: PaginatedhuggingFaceHuggingFacePublicationResponseList
	fmt.Fprintf(os.Stdout, "Response from `PublicationsHuggingFaceAPI.PublicationsHuggingFaceHuggingFaceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicationsHuggingFaceHuggingFaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **checkpoint** | **bool** | Filter results where checkpoint matches value | 
 **content** | **string** | Content Unit referenced by HREF/PRN | 
 **contentIn** | **[]string** | Multiple values may be separated by commas. | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;complete&#x60; - Complete* &#x60;-complete&#x60; - Complete (descending)* &#x60;pass_through&#x60; - Pass through* &#x60;-pass_through&#x60; - Pass through (descending)* &#x60;checkpoint&#x60; - Checkpoint* &#x60;-checkpoint&#x60; - Checkpoint (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
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
 **repository** | **string** | Repository referenced by HREF/PRN | 
 **repositoryVersion** | **string** |  | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedhuggingFaceHuggingFacePublicationResponseList**](PaginatedhuggingFaceHuggingFacePublicationResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicationsHuggingFaceHuggingFaceRead

> HuggingFaceHuggingFacePublicationResponse PublicationsHuggingFaceHuggingFaceRead(ctx, huggingFaceHuggingFacePublicationHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a hugging face publication



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
	huggingFaceHuggingFacePublicationHref := "huggingFaceHuggingFacePublicationHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicationsHuggingFaceAPI.PublicationsHuggingFaceHuggingFaceRead(context.Background(), huggingFaceHuggingFacePublicationHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicationsHuggingFaceAPI.PublicationsHuggingFaceHuggingFaceRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicationsHuggingFaceHuggingFaceRead`: HuggingFaceHuggingFacePublicationResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicationsHuggingFaceAPI.PublicationsHuggingFaceHuggingFaceRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**huggingFaceHuggingFacePublicationHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicationsHuggingFaceHuggingFaceReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**HuggingFaceHuggingFacePublicationResponse**](HuggingFaceHuggingFacePublicationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

