# \PypiLegacyAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PypiLegacyCreate**](PypiLegacyAPI.md#PypiLegacyCreate) | **Post** /pypi/{pulp_domain}/{path}/legacy/ | Upload a package



## PypiLegacyCreate

> PackageUploadTaskResponse PypiLegacyCreate(ctx, path, pulpDomain).Content(content).Sha256Digest(sha256Digest).XTaskDiagnostics(xTaskDiagnostics).Action(action).ProtocolVersion(protocolVersion).Filetype(filetype).MetadataVersion(metadataVersion).Attestations(attestations).Execute()

Upload a package



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
	path := "path_example" // string | 
	pulpDomain := "pulpDomain_example" // string | 
	content := os.NewFile(1234, "some_file") // *os.File | A Python package release file to upload to the index.
	sha256Digest := "sha256Digest_example" // string | SHA256 of package to validate upload integrity.
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	action := "action_example" // string | Defaults to `file_upload`, don't change it or request will fail! (optional) (default to "file_upload")
	protocolVersion := openapiclient.ProtocolVersionEnum(1) // ProtocolVersionEnum | Protocol version to use for the upload. Only version 1 is supported.* `1` - 1 (optional) (default to 1)
	filetype := openapiclient.FiletypeEnum("bdist_wheel") // FiletypeEnum | Type of artifact to upload.* `bdist_wheel` - bdist_wheel* `sdist` - sdist (optional)
	metadataVersion := openapiclient.MetadataVersionEnum("1.0") // MetadataVersionEnum | Metadata version of the uploaded package.* `1.0` - 1.0* `1.1` - 1.1* `1.2` - 1.2* `2.0` - 2.0* `2.1` - 2.1* `2.2` - 2.2* `2.3` - 2.3* `2.4` - 2.4* `2.5` - 2.5 (optional)
	attestations := TODO // interface{} | A JSON list containing attestations for the package. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PypiLegacyAPI.PypiLegacyCreate(context.Background(), path, pulpDomain).Content(content).Sha256Digest(sha256Digest).XTaskDiagnostics(xTaskDiagnostics).Action(action).ProtocolVersion(protocolVersion).Filetype(filetype).MetadataVersion(metadataVersion).Attestations(attestations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PypiLegacyAPI.PypiLegacyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PypiLegacyCreate`: PackageUploadTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `PypiLegacyAPI.PypiLegacyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPypiLegacyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **content** | ***os.File** | A Python package release file to upload to the index. | 
 **sha256Digest** | **string** | SHA256 of package to validate upload integrity. | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **action** | **string** | Defaults to &#x60;file_upload&#x60;, don&#39;t change it or request will fail! | [default to &quot;file_upload&quot;]
 **protocolVersion** | [**ProtocolVersionEnum**](ProtocolVersionEnum.md) | Protocol version to use for the upload. Only version 1 is supported.* &#x60;1&#x60; - 1 | [default to 1]
 **filetype** | [**FiletypeEnum**](FiletypeEnum.md) | Type of artifact to upload.* &#x60;bdist_wheel&#x60; - bdist_wheel* &#x60;sdist&#x60; - sdist | 
 **metadataVersion** | [**MetadataVersionEnum**](MetadataVersionEnum.md) | Metadata version of the uploaded package.* &#x60;1.0&#x60; - 1.0* &#x60;1.1&#x60; - 1.1* &#x60;1.2&#x60; - 1.2* &#x60;2.0&#x60; - 2.0* &#x60;2.1&#x60; - 2.1* &#x60;2.2&#x60; - 2.2* &#x60;2.3&#x60; - 2.3* &#x60;2.4&#x60; - 2.4* &#x60;2.5&#x60; - 2.5 | 
 **attestations** | [**interface{}**](interface{}.md) | A JSON list containing attestations for the package. | 

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

