# \ContentAdvisoriesApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentRpmAdvisoriesCreate**](ContentAdvisoriesApi.md#ContentRpmAdvisoriesCreate) | **Post** /pulp/api/v3/content/rpm/advisories/ | Create an update record
[**ContentRpmAdvisoriesList**](ContentAdvisoriesApi.md#ContentRpmAdvisoriesList) | **Get** /pulp/api/v3/content/rpm/advisories/ | List update records
[**ContentRpmAdvisoriesRead**](ContentAdvisoriesApi.md#ContentRpmAdvisoriesRead) | **Get** /{rpm_update_record_href} | Inspect an update record



## ContentRpmAdvisoriesCreate

> AsyncOperationResponse ContentRpmAdvisoriesCreate(ctx).File(file).Repository(repository).Execute()

Create an update record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the artifact of the content unit. (optional)
    repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentAdvisoriesApi.ContentRpmAdvisoriesCreate(context.Background()).File(file).Repository(repository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentAdvisoriesApi.ContentRpmAdvisoriesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmAdvisoriesCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentAdvisoriesApi.ContentRpmAdvisoriesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmAdvisoriesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | 
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentRpmAdvisoriesList

> PaginatedrpmUpdateRecordResponseList ContentRpmAdvisoriesList(ctx).Id(id).IdIn(idIn).Limit(limit).Offset(offset).Ordering(ordering).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Severity(severity).SeverityIn(severityIn).SeverityNe(severityNe).Status(status).StatusIn(statusIn).StatusNe(statusNe).Type_(type_).TypeIn(typeIn).TypeNe(typeNe).Fields(fields).ExcludeFields(excludeFields).Execute()

List update records



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release"
)

func main() {
    id := "id_example" // string | Filter results where id matches value (optional)
    idIn := []string{"Inner_example"} // []string | Filter results where id is in a comma-separated list of values (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    severity := "severity_example" // string | Filter results where severity matches value (optional)
    severityIn := []string{"Inner_example"} // []string | Filter results where severity is in a comma-separated list of values (optional)
    severityNe := "severityNe_example" // string | Filter results where severity not equal to value (optional)
    status := "status_example" // string | Filter results where status matches value (optional)
    statusIn := []string{"Inner_example"} // []string | Filter results where status is in a comma-separated list of values (optional)
    statusNe := "statusNe_example" // string | Filter results where status not equal to value (optional)
    type_ := "type__example" // string | Filter results where type matches value (optional)
    typeIn := []string{"Inner_example"} // []string | Filter results where type is in a comma-separated list of values (optional)
    typeNe := "typeNe_example" // string | Filter results where type not equal to value (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentAdvisoriesApi.ContentRpmAdvisoriesList(context.Background()).Id(id).IdIn(idIn).Limit(limit).Offset(offset).Ordering(ordering).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Severity(severity).SeverityIn(severityIn).SeverityNe(severityNe).Status(status).StatusIn(statusIn).StatusNe(statusNe).Type_(type_).TypeIn(typeIn).TypeNe(typeNe).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentAdvisoriesApi.ContentRpmAdvisoriesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmAdvisoriesList`: PaginatedrpmUpdateRecordResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentAdvisoriesApi.ContentRpmAdvisoriesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmAdvisoriesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Filter results where id matches value | 
 **idIn** | **[]string** | Filter results where id is in a comma-separated list of values | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **severity** | **string** | Filter results where severity matches value | 
 **severityIn** | **[]string** | Filter results where severity is in a comma-separated list of values | 
 **severityNe** | **string** | Filter results where severity not equal to value | 
 **status** | **string** | Filter results where status matches value | 
 **statusIn** | **[]string** | Filter results where status is in a comma-separated list of values | 
 **statusNe** | **string** | Filter results where status not equal to value | 
 **type_** | **string** | Filter results where type matches value | 
 **typeIn** | **[]string** | Filter results where type is in a comma-separated list of values | 
 **typeNe** | **string** | Filter results where type not equal to value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedrpmUpdateRecordResponseList**](PaginatedrpmUpdateRecordResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentRpmAdvisoriesRead

> RpmUpdateRecordResponse ContentRpmAdvisoriesRead(ctx, rpmUpdateRecordHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an update record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release"
)

func main() {
    rpmUpdateRecordHref := "rpmUpdateRecordHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentAdvisoriesApi.ContentRpmAdvisoriesRead(context.Background(), rpmUpdateRecordHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentAdvisoriesApi.ContentRpmAdvisoriesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmAdvisoriesRead`: RpmUpdateRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentAdvisoriesApi.ContentRpmAdvisoriesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmUpdateRecordHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmAdvisoriesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RpmUpdateRecordResponse**](RpmUpdateRecordResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

