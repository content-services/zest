# \ApiPypiLegacyAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPulpPypiLegacyCreate**](ApiPypiLegacyAPI.md#ApiPulpPypiLegacyCreate) | **Post** /api/pulp/pypi/{pulp_domain}/{path}/legacy/ | Upload a package



## ApiPulpPypiLegacyCreate

> PackageUploadTaskResponse ApiPulpPypiLegacyCreate(ctx, path, pulpDomain).Content(content).Sha256Digest(sha256Digest).Action(action).Execute()

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
	resp, r, err := apiClient.ApiPypiLegacyAPI.ApiPulpPypiLegacyCreate(context.Background(), path, pulpDomain).Content(content).Sha256Digest(sha256Digest).Action(action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiPypiLegacyAPI.ApiPulpPypiLegacyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPulpPypiLegacyCreate`: PackageUploadTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiPypiLegacyAPI.ApiPulpPypiLegacyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPulpPypiLegacyCreateRequest struct via the builder pattern


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

