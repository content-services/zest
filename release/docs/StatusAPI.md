# \StatusAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusRead1**](StatusAPI.md#StatusRead1) | **Get** /pulp/api/v3/status/ | Inspect status of Pulp



## StatusRead1

> StatusResponse StatusRead1(ctx).Execute()

Inspect status of Pulp



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release/v2023"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatusAPI.StatusRead1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.StatusRead1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatusRead1`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `StatusAPI.StatusRead1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatusRead1Request struct via the builder pattern


### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

