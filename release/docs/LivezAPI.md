# \LivezAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LivezRead**](LivezAPI.md#LivezRead) | **Get** /api/pulp/api/v3/livez/ | Inspect liveness of Pulp&#39;s REST API.



## LivezRead

> LivezRead(ctx).Execute()

Inspect liveness of Pulp's REST API.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LivezAPI.LivezRead(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LivezAPI.LivezRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLivezReadRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

