# \ContentCollectionSignaturesApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentAnsibleCollectionSignaturesCreate**](ContentCollectionSignaturesApi.md#ContentAnsibleCollectionSignaturesCreate) | **Post** /pulp/api/v3/content/ansible/collection_signatures/ | Create a collection version signature
[**ContentAnsibleCollectionSignaturesList**](ContentCollectionSignaturesApi.md#ContentAnsibleCollectionSignaturesList) | **Get** /pulp/api/v3/content/ansible/collection_signatures/ | List collection version signatures
[**ContentAnsibleCollectionSignaturesRead**](ContentCollectionSignaturesApi.md#ContentAnsibleCollectionSignaturesRead) | **Get** /{ansible_collection_version_signature_href} | Inspect a collection version signature



## ContentAnsibleCollectionSignaturesCreate

> AsyncOperationResponse ContentAnsibleCollectionSignaturesCreate(ctx).File(file).SignedCollection(signedCollection).Repository(repository).Execute()

Create a collection version signature



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
    file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the artifact of the content unit.
    signedCollection := "signedCollection_example" // string | The content this signature is pointing to.
    repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentCollectionSignaturesApi.ContentAnsibleCollectionSignaturesCreate(context.Background()).File(file).SignedCollection(signedCollection).Repository(repository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionSignaturesApi.ContentAnsibleCollectionSignaturesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentAnsibleCollectionSignaturesCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentCollectionSignaturesApi.ContentAnsibleCollectionSignaturesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionSignaturesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | 
 **signedCollection** | **string** | The content this signature is pointing to. | 
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


## ContentAnsibleCollectionSignaturesList

> PaginatedansibleCollectionVersionSignatureResponseList ContentAnsibleCollectionSignaturesList(ctx).Limit(limit).Offset(offset).Ordering(ordering).PubkeyFingerprint(pubkeyFingerprint).PubkeyFingerprintIn(pubkeyFingerprintIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).SignedCollection(signedCollection).SigningService(signingService).Fields(fields).ExcludeFields(excludeFields).Execute()

List collection version signatures



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
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    pubkeyFingerprint := "pubkeyFingerprint_example" // string | Filter results where pubkey_fingerprint matches value (optional)
    pubkeyFingerprintIn := []string{"Inner_example"} // []string | Filter results where pubkey_fingerprint is in a comma-separated list of values (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    signedCollection := "signedCollection_example" // string | Filter signatures for collection version (optional)
    signingService := "signingService_example" // string | Filter signatures produced by signature service (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentCollectionSignaturesApi.ContentAnsibleCollectionSignaturesList(context.Background()).Limit(limit).Offset(offset).Ordering(ordering).PubkeyFingerprint(pubkeyFingerprint).PubkeyFingerprintIn(pubkeyFingerprintIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).SignedCollection(signedCollection).SigningService(signingService).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionSignaturesApi.ContentAnsibleCollectionSignaturesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentAnsibleCollectionSignaturesList`: PaginatedansibleCollectionVersionSignatureResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentCollectionSignaturesApi.ContentAnsibleCollectionSignaturesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionSignaturesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **pubkeyFingerprint** | **string** | Filter results where pubkey_fingerprint matches value | 
 **pubkeyFingerprintIn** | **[]string** | Filter results where pubkey_fingerprint is in a comma-separated list of values | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **signedCollection** | **string** | Filter signatures for collection version | 
 **signingService** | **string** | Filter signatures produced by signature service | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedansibleCollectionVersionSignatureResponseList**](PaginatedansibleCollectionVersionSignatureResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentAnsibleCollectionSignaturesRead

> AnsibleCollectionVersionSignatureResponse ContentAnsibleCollectionSignaturesRead(ctx, ansibleCollectionVersionSignatureHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a collection version signature



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
    ansibleCollectionVersionSignatureHref := "ansibleCollectionVersionSignatureHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentCollectionSignaturesApi.ContentAnsibleCollectionSignaturesRead(context.Background(), ansibleCollectionVersionSignatureHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentCollectionSignaturesApi.ContentAnsibleCollectionSignaturesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentAnsibleCollectionSignaturesRead`: AnsibleCollectionVersionSignatureResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentCollectionSignaturesApi.ContentAnsibleCollectionSignaturesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ansibleCollectionVersionSignatureHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentAnsibleCollectionSignaturesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**AnsibleCollectionVersionSignatureResponse**](AnsibleCollectionVersionSignatureResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

