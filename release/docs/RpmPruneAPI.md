# \RpmPruneAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RpmPrunePrunePackages**](RpmPruneAPI.md#RpmPrunePrunePackages) | **Post** /api/pulp/{pulp_domain}/api/v3/rpm/prune/ | 



## RpmPrunePrunePackages

> TaskGroupOperationResponse RpmPrunePrunePackages(ctx, pulpDomain).PrunePackages(prunePackages).Execute()





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
	pulpDomain := "pulpDomain_example" // string | 
	prunePackages := *openapiclient.NewPrunePackages([]string{"RepoHrefs_example"}) // PrunePackages | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RpmPruneAPI.RpmPrunePrunePackages(context.Background(), pulpDomain).PrunePackages(prunePackages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RpmPruneAPI.RpmPrunePrunePackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RpmPrunePrunePackages`: TaskGroupOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RpmPruneAPI.RpmPrunePrunePackages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRpmPrunePrunePackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prunePackages** | [**PrunePackages**](PrunePackages.md) |  | 

### Return type

[**TaskGroupOperationResponse**](TaskGroupOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

