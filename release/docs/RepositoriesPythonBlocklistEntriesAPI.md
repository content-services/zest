# \RepositoriesPythonBlocklistEntriesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesPythonPythonBlocklistEntriesCreate**](RepositoriesPythonBlocklistEntriesAPI.md#RepositoriesPythonPythonBlocklistEntriesCreate) | **Post** /{python_python_repository_href}blocklist_entries/ | Create a python blocklist entry
[**RepositoriesPythonPythonBlocklistEntriesDelete**](RepositoriesPythonBlocklistEntriesAPI.md#RepositoriesPythonPythonBlocklistEntriesDelete) | **Delete** /{python_python_python_blocklist_entry_href} | Delete a python blocklist entry
[**RepositoriesPythonPythonBlocklistEntriesList**](RepositoriesPythonBlocklistEntriesAPI.md#RepositoriesPythonPythonBlocklistEntriesList) | **Get** /{python_python_repository_href}blocklist_entries/ | List python blocklist entrys
[**RepositoriesPythonPythonBlocklistEntriesRead**](RepositoriesPythonBlocklistEntriesAPI.md#RepositoriesPythonPythonBlocklistEntriesRead) | **Get** /{python_python_python_blocklist_entry_href} | Inspect a python blocklist entry



## RepositoriesPythonPythonBlocklistEntriesCreate

> PythonPythonBlocklistEntryResponse RepositoriesPythonPythonBlocklistEntriesCreate(ctx, pythonPythonRepositoryHref).PythonPythonBlocklistEntry(pythonPythonBlocklistEntry).XTaskDiagnostics(xTaskDiagnostics).Execute()

Create a python blocklist entry



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
	pythonPythonRepositoryHref := "pythonPythonRepositoryHref_example" // string | 
	pythonPythonBlocklistEntry := *openapiclient.NewPythonPythonBlocklistEntry() // PythonPythonBlocklistEntry | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesPythonBlocklistEntriesAPI.RepositoriesPythonPythonBlocklistEntriesCreate(context.Background(), pythonPythonRepositoryHref).PythonPythonBlocklistEntry(pythonPythonBlocklistEntry).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesPythonBlocklistEntriesAPI.RepositoriesPythonPythonBlocklistEntriesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesPythonPythonBlocklistEntriesCreate`: PythonPythonBlocklistEntryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesPythonBlocklistEntriesAPI.RepositoriesPythonPythonBlocklistEntriesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pythonPythonRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesPythonPythonBlocklistEntriesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pythonPythonBlocklistEntry** | [**PythonPythonBlocklistEntry**](PythonPythonBlocklistEntry.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**PythonPythonBlocklistEntryResponse**](PythonPythonBlocklistEntryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesPythonPythonBlocklistEntriesDelete

> RepositoriesPythonPythonBlocklistEntriesDelete(ctx, pythonPythonPythonBlocklistEntryHref).XTaskDiagnostics(xTaskDiagnostics).Execute()

Delete a python blocklist entry



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
	pythonPythonPythonBlocklistEntryHref := "pythonPythonPythonBlocklistEntryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoriesPythonBlocklistEntriesAPI.RepositoriesPythonPythonBlocklistEntriesDelete(context.Background(), pythonPythonPythonBlocklistEntryHref).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesPythonBlocklistEntriesAPI.RepositoriesPythonPythonBlocklistEntriesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pythonPythonPythonBlocklistEntryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesPythonPythonBlocklistEntriesDeleteRequest struct via the builder pattern


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


## RepositoriesPythonPythonBlocklistEntriesList

> PaginatedpythonPythonBlocklistEntryResponseList RepositoriesPythonPythonBlocklistEntriesList(ctx, pythonPythonRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Filename(filename).Limit(limit).Name(name).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).Version(version).VersionIsnull(versionIsnull).Fields(fields).ExcludeFields(excludeFields).Execute()

List python blocklist entrys



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
	pythonPythonRepositoryHref := "pythonPythonRepositoryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	filename := "filename_example" // string | Filter results where filename matches value (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string | Filter results where name matches value (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `name` - Name* `-name` - Name (descending)* `version` - Version* `-version` - Version (descending)* `filename` - Filename* `-filename` - Filename (descending)* `added_by` - Added by* `-added_by` - Added by (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	version := "version_example" // string | Filter results where version matches value (optional)
	versionIsnull := true // bool | Filter results where version has a null value (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesPythonBlocklistEntriesAPI.RepositoriesPythonPythonBlocklistEntriesList(context.Background(), pythonPythonRepositoryHref).XTaskDiagnostics(xTaskDiagnostics).Filename(filename).Limit(limit).Name(name).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).Version(version).VersionIsnull(versionIsnull).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesPythonBlocklistEntriesAPI.RepositoriesPythonPythonBlocklistEntriesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesPythonPythonBlocklistEntriesList`: PaginatedpythonPythonBlocklistEntryResponseList
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesPythonBlocklistEntriesAPI.RepositoriesPythonPythonBlocklistEntriesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pythonPythonRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesPythonPythonBlocklistEntriesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **filename** | **string** | Filter results where filename matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;version&#x60; - Version* &#x60;-version&#x60; - Version (descending)* &#x60;filename&#x60; - Filename* &#x60;-filename&#x60; - Filename (descending)* &#x60;added_by&#x60; - Added by* &#x60;-added_by&#x60; - Added by (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **version** | **string** | Filter results where version matches value | 
 **versionIsnull** | **bool** | Filter results where version has a null value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedpythonPythonBlocklistEntryResponseList**](PaginatedpythonPythonBlocklistEntryResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesPythonPythonBlocklistEntriesRead

> PythonPythonBlocklistEntryResponse RepositoriesPythonPythonBlocklistEntriesRead(ctx, pythonPythonPythonBlocklistEntryHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a python blocklist entry



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
	pythonPythonPythonBlocklistEntryHref := "pythonPythonPythonBlocklistEntryHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesPythonBlocklistEntriesAPI.RepositoriesPythonPythonBlocklistEntriesRead(context.Background(), pythonPythonPythonBlocklistEntryHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesPythonBlocklistEntriesAPI.RepositoriesPythonPythonBlocklistEntriesRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesPythonPythonBlocklistEntriesRead`: PythonPythonBlocklistEntryResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesPythonBlocklistEntriesAPI.RepositoriesPythonPythonBlocklistEntriesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pythonPythonPythonBlocklistEntryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesPythonPythonBlocklistEntriesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PythonPythonBlocklistEntryResponse**](PythonPythonBlocklistEntryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

