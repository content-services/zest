# \ContentOpenpgpUseridAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentCoreOpenpgpUseridList**](ContentOpenpgpUseridAPI.md#ContentCoreOpenpgpUseridList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/core/openpgp_userid/ | List open pgp user ids
[**ContentCoreOpenpgpUseridRead**](ContentOpenpgpUseridAPI.md#ContentCoreOpenpgpUseridRead) | **Get** /{open_p_g_p_user_i_d_href} | Inspect an open pgp user id
[**ContentCoreOpenpgpUseridSetLabel**](ContentOpenpgpUseridAPI.md#ContentCoreOpenpgpUseridSetLabel) | **Post** /{open_p_g_p_user_i_d_href}set_label/ | Set a label
[**ContentCoreOpenpgpUseridUnsetLabel**](ContentOpenpgpUseridAPI.md#ContentCoreOpenpgpUseridUnsetLabel) | **Post** /{open_p_g_p_user_i_d_href}unset_label/ | Unset a label



## ContentCoreOpenpgpUseridList

> PaginatedOpenPGPUserIDResponseList ContentCoreOpenpgpUseridList(ctx, pulpDomain).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).UserId(userId).UserIdContains(userIdContains).UserIdIcontains(userIdIcontains).UserIdIexact(userIdIexact).UserIdIn(userIdIn).UserIdIregex(userIdIregex).UserIdIstartswith(userIdIstartswith).UserIdRegex(userIdRegex).UserIdStartswith(userIdStartswith).Fields(fields).ExcludeFields(excludeFields).Execute()

List open pgp user ids



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
	pulpDomain := "pulpDomain_example" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `raw_data` - Raw data* `-raw_data` - Raw data (descending)* `user_id` - User id* `-user_id` - User id (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF/PRN (optional)
	userId := "userId_example" // string | Filter results where user_id matches value (optional)
	userIdContains := "userIdContains_example" // string | Filter results where user_id contains value (optional)
	userIdIcontains := "userIdIcontains_example" // string | Filter results where user_id contains value (optional)
	userIdIexact := "userIdIexact_example" // string | Filter results where user_id matches value (optional)
	userIdIn := []string{"Inner_example"} // []string | Filter results where user_id is in a comma-separated list of values (optional)
	userIdIregex := "userIdIregex_example" // string | Filter results where user_id matches regex value (optional)
	userIdIstartswith := "userIdIstartswith_example" // string | Filter results where user_id starts with value (optional)
	userIdRegex := "userIdRegex_example" // string | Filter results where user_id matches regex value (optional)
	userIdStartswith := "userIdStartswith_example" // string | Filter results where user_id starts with value (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentOpenpgpUseridAPI.ContentCoreOpenpgpUseridList(context.Background(), pulpDomain).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).UserId(userId).UserIdContains(userIdContains).UserIdIcontains(userIdIcontains).UserIdIexact(userIdIexact).UserIdIn(userIdIn).UserIdIregex(userIdIregex).UserIdIstartswith(userIdIstartswith).UserIdRegex(userIdRegex).UserIdStartswith(userIdStartswith).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentOpenpgpUseridAPI.ContentCoreOpenpgpUseridList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentCoreOpenpgpUseridList`: PaginatedOpenPGPUserIDResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentOpenpgpUseridAPI.ContentCoreOpenpgpUseridList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentCoreOpenpgpUseridListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;raw_data&#x60; - Raw data* &#x60;-raw_data&#x60; - Raw data (descending)* &#x60;user_id&#x60; - User id* &#x60;-user_id&#x60; - User id (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF/PRN | 
 **userId** | **string** | Filter results where user_id matches value | 
 **userIdContains** | **string** | Filter results where user_id contains value | 
 **userIdIcontains** | **string** | Filter results where user_id contains value | 
 **userIdIexact** | **string** | Filter results where user_id matches value | 
 **userIdIn** | **[]string** | Filter results where user_id is in a comma-separated list of values | 
 **userIdIregex** | **string** | Filter results where user_id matches regex value | 
 **userIdIstartswith** | **string** | Filter results where user_id starts with value | 
 **userIdRegex** | **string** | Filter results where user_id matches regex value | 
 **userIdStartswith** | **string** | Filter results where user_id starts with value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedOpenPGPUserIDResponseList**](PaginatedOpenPGPUserIDResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentCoreOpenpgpUseridRead

> OpenPGPUserIDResponse ContentCoreOpenpgpUseridRead(ctx, openPGPUserIDHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an open pgp user id



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
	openPGPUserIDHref := "openPGPUserIDHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentOpenpgpUseridAPI.ContentCoreOpenpgpUseridRead(context.Background(), openPGPUserIDHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentOpenpgpUseridAPI.ContentCoreOpenpgpUseridRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentCoreOpenpgpUseridRead`: OpenPGPUserIDResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentOpenpgpUseridAPI.ContentCoreOpenpgpUseridRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPUserIDHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentCoreOpenpgpUseridReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**OpenPGPUserIDResponse**](OpenPGPUserIDResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentCoreOpenpgpUseridSetLabel

> SetLabelResponse ContentCoreOpenpgpUseridSetLabel(ctx, openPGPUserIDHref).SetLabel(setLabel).Execute()

Set a label



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
	openPGPUserIDHref := "openPGPUserIDHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentOpenpgpUseridAPI.ContentCoreOpenpgpUseridSetLabel(context.Background(), openPGPUserIDHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentOpenpgpUseridAPI.ContentCoreOpenpgpUseridSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentCoreOpenpgpUseridSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentOpenpgpUseridAPI.ContentCoreOpenpgpUseridSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPUserIDHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentCoreOpenpgpUseridSetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setLabel** | [**SetLabel**](SetLabel.md) |  | 

### Return type

[**SetLabelResponse**](SetLabelResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentCoreOpenpgpUseridUnsetLabel

> UnsetLabelResponse ContentCoreOpenpgpUseridUnsetLabel(ctx, openPGPUserIDHref).UnsetLabel(unsetLabel).Execute()

Unset a label



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
	openPGPUserIDHref := "openPGPUserIDHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentOpenpgpUseridAPI.ContentCoreOpenpgpUseridUnsetLabel(context.Background(), openPGPUserIDHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentOpenpgpUseridAPI.ContentCoreOpenpgpUseridUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentCoreOpenpgpUseridUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentOpenpgpUseridAPI.ContentCoreOpenpgpUseridUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**openPGPUserIDHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentCoreOpenpgpUseridUnsetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unsetLabel** | [**UnsetLabel**](UnsetLabel.md) |  | 

### Return type

[**UnsetLabelResponse**](UnsetLabelResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

