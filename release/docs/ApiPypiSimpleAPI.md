# \ApiPypiSimpleAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPulpPypiSimpleCreate**](ApiPypiSimpleAPI.md#ApiPulpPypiSimpleCreate) | **Post** /api/pulp/pypi/{pulp_domain}/{path}/simple/ | Upload a package
[**ApiPulpPypiSimpleRead**](ApiPypiSimpleAPI.md#ApiPulpPypiSimpleRead) | **Get** /api/pulp/pypi/{pulp_domain}/{path}/simple/ | Get index simple page
[**PypiSimplePackageRead**](ApiPypiSimpleAPI.md#PypiSimplePackageRead) | **Get** /api/pulp/pypi/{pulp_domain}/{path}/simple/{package}/ | Get package simple page



## ApiPulpPypiSimpleCreate

> PackageUploadTaskResponse ApiPulpPypiSimpleCreate(ctx, path, pulpDomain).Content(content).Sha256Digest(sha256Digest).Action(action).Execute()

Upload a package



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
	path := "path_example" // string | 
	pulpDomain := "pulpDomain_example" // string | 
	content := os.NewFile(1234, "some_file") // *os.File | A Python package release file to upload to the index.
	sha256Digest := "sha256Digest_example" // string | SHA256 of package to validate upload integrity.
	action := "action_example" // string | Defaults to `file_upload`, don't change it or request will fail! (optional) (default to "file_upload")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiPypiSimpleAPI.ApiPulpPypiSimpleCreate(context.Background(), path, pulpDomain).Content(content).Sha256Digest(sha256Digest).Action(action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiPypiSimpleAPI.ApiPulpPypiSimpleCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPulpPypiSimpleCreate`: PackageUploadTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiPypiSimpleAPI.ApiPulpPypiSimpleCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPulpPypiSimpleCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **content** | ***os.File** | A Python package release file to upload to the index. | 
 **sha256Digest** | **string** | SHA256 of package to validate upload integrity. | 
 **action** | **string** | Defaults to &#x60;file_upload&#x60;, don&#39;t change it or request will fail! | [default to &quot;file_upload&quot;]

### Return type

[**PackageUploadTaskResponse**](PackageUploadTaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPulpPypiSimpleRead

> ApiPulpPypiSimpleRead(ctx, path, pulpDomain).Fields(fields).ExcludeFields(excludeFields).Execute()

Get index simple page



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
	path := "path_example" // string | 
	pulpDomain := "pulpDomain_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiPypiSimpleAPI.ApiPulpPypiSimpleRead(context.Background(), path, pulpDomain).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiPypiSimpleAPI.ApiPulpPypiSimpleRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPulpPypiSimpleReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

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


## PypiSimplePackageRead

> PypiSimplePackageRead(ctx, package_, path, pulpDomain).Fields(fields).ExcludeFields(excludeFields).Execute()

Get package simple page



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
	package_ := "package__example" // string | 
	path := "path_example" // string | 
	pulpDomain := "pulpDomain_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApiPypiSimpleAPI.PypiSimplePackageRead(context.Background(), package_, path, pulpDomain).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiPypiSimpleAPI.PypiSimplePackageRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**package_** | **string** |  | 
**path** | **string** |  | 
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPypiSimplePackageReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

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

