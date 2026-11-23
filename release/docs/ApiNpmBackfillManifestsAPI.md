# \ApiNpmBackfillManifestsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPulpNpmBackfillManifestsPost**](ApiNpmBackfillManifestsAPI.md#ApiPulpNpmBackfillManifestsPost) | **Post** /api/pulp/npm/backfill-manifests/ | Backfill npm package manifests



## ApiPulpNpmBackfillManifestsPost

> AsyncOperationResponse ApiPulpNpmBackfillManifestsPost(ctx).NpmManifestBackfill(npmManifestBackfill).XTaskDiagnostics(xTaskDiagnostics).Execute()

Backfill npm package manifests



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
	npmManifestBackfill := *openapiclient.NewNpmManifestBackfill("Repository_example") // NpmManifestBackfill | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiNpmBackfillManifestsAPI.ApiPulpNpmBackfillManifestsPost(context.Background()).NpmManifestBackfill(npmManifestBackfill).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiNpmBackfillManifestsAPI.ApiPulpNpmBackfillManifestsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPulpNpmBackfillManifestsPost`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiNpmBackfillManifestsAPI.ApiPulpNpmBackfillManifestsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPulpNpmBackfillManifestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **npmManifestBackfill** | [**NpmManifestBackfill**](NpmManifestBackfill.md) |  | 
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

