# \RemotesGemAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemotesGemGemAddRole**](RemotesGemAPI.md#RemotesGemGemAddRole) | **Post** /{gem_gem_remote_href}add_role/ | Add a role
[**RemotesGemGemCreate**](RemotesGemAPI.md#RemotesGemGemCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/remotes/gem/gem/ | Create a gem remote
[**RemotesGemGemDelete**](RemotesGemAPI.md#RemotesGemGemDelete) | **Delete** /{gem_gem_remote_href} | Delete a gem remote
[**RemotesGemGemList**](RemotesGemAPI.md#RemotesGemGemList) | **Get** /api/pulp/{pulp_domain}/api/v3/remotes/gem/gem/ | List gem remotes
[**RemotesGemGemListRoles**](RemotesGemAPI.md#RemotesGemGemListRoles) | **Get** /{gem_gem_remote_href}list_roles/ | List roles
[**RemotesGemGemMyPermissions**](RemotesGemAPI.md#RemotesGemGemMyPermissions) | **Get** /{gem_gem_remote_href}my_permissions/ | List user permissions
[**RemotesGemGemPartialUpdate**](RemotesGemAPI.md#RemotesGemGemPartialUpdate) | **Patch** /{gem_gem_remote_href} | Update a gem remote
[**RemotesGemGemRead**](RemotesGemAPI.md#RemotesGemGemRead) | **Get** /{gem_gem_remote_href} | Inspect a gem remote
[**RemotesGemGemRemoveRole**](RemotesGemAPI.md#RemotesGemGemRemoveRole) | **Post** /{gem_gem_remote_href}remove_role/ | Remove a role
[**RemotesGemGemSetLabel**](RemotesGemAPI.md#RemotesGemGemSetLabel) | **Post** /{gem_gem_remote_href}set_label/ | Set a label
[**RemotesGemGemUnsetLabel**](RemotesGemAPI.md#RemotesGemGemUnsetLabel) | **Post** /{gem_gem_remote_href}unset_label/ | Unset a label
[**RemotesGemGemUpdate**](RemotesGemAPI.md#RemotesGemGemUpdate) | **Put** /{gem_gem_remote_href} | Update a gem remote



## RemotesGemGemAddRole

> NestedRoleResponse RemotesGemGemAddRole(ctx, gemGemRemoteHref).NestedRole(nestedRole).Execute()

Add a role



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
	gemGemRemoteHref := "gemGemRemoteHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesGemAPI.RemotesGemGemAddRole(context.Background(), gemGemRemoteHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesGemAPI.RemotesGemGemAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesGemGemAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesGemAPI.RemotesGemGemAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesGemGemAddRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nestedRole** | [**NestedRole**](NestedRole.md) |  | 

### Return type

[**NestedRoleResponse**](NestedRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesGemGemCreate

> GemGemRemoteResponse RemotesGemGemCreate(ctx, pulpDomain).GemGemRemote(gemGemRemote).Execute()

Create a gem remote



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
	pulpDomain := "pulpDomain_example" // string | 
	gemGemRemote := *openapiclient.NewGemGemRemote("Name_example", "Url_example") // GemGemRemote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesGemAPI.RemotesGemGemCreate(context.Background(), pulpDomain).GemGemRemote(gemGemRemote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesGemAPI.RemotesGemGemCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesGemGemCreate`: GemGemRemoteResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesGemAPI.RemotesGemGemCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesGemGemCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gemGemRemote** | [**GemGemRemote**](GemGemRemote.md) |  | 

### Return type

[**GemGemRemoteResponse**](GemGemRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesGemGemDelete

> AsyncOperationResponse RemotesGemGemDelete(ctx, gemGemRemoteHref).Execute()

Delete a gem remote



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
	gemGemRemoteHref := "gemGemRemoteHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesGemAPI.RemotesGemGemDelete(context.Background(), gemGemRemoteHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesGemAPI.RemotesGemGemDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesGemGemDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesGemAPI.RemotesGemGemDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesGemGemDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesGemGemList

> PaginatedgemGemRemoteResponseList RemotesGemGemList(ctx, pulpDomain).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()

List gem remotes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string | Filter results where name matches value (optional)
	nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
	nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
	nameIexact := "nameIexact_example" // string | Filter results where name matches value (optional)
	nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
	nameIregex := "nameIregex_example" // string | Filter results where name matches regex value (optional)
	nameIstartswith := "nameIstartswith_example" // string | Filter results where name starts with value (optional)
	nameRegex := "nameRegex_example" // string | Filter results where name matches regex value (optional)
	nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `name` - Name* `-name` - Name (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `url` - Url* `-url` - Url (descending)* `ca_cert` - Ca cert* `-ca_cert` - Ca cert (descending)* `client_cert` - Client cert* `-client_cert` - Client cert (descending)* `client_key` - Client key* `-client_key` - Client key (descending)* `tls_validation` - Tls validation* `-tls_validation` - Tls validation (descending)* `username` - Username* `-username` - Username (descending)* `password` - Password* `-password` - Password (descending)* `proxy_url` - Proxy url* `-proxy_url` - Proxy url (descending)* `proxy_username` - Proxy username* `-proxy_username` - Proxy username (descending)* `proxy_password` - Proxy password* `-proxy_password` - Proxy password (descending)* `download_concurrency` - Download concurrency* `-download_concurrency` - Download concurrency (descending)* `max_retries` - Max retries* `-max_retries` - Max retries (descending)* `policy` - Policy* `-policy` - Policy (descending)* `total_timeout` - Total timeout* `-total_timeout` - Total timeout (descending)* `connect_timeout` - Connect timeout* `-connect_timeout` - Connect timeout (descending)* `sock_connect_timeout` - Sock connect timeout* `-sock_connect_timeout` - Sock connect timeout (descending)* `sock_read_timeout` - Sock read timeout* `-sock_read_timeout` - Sock read timeout (descending)* `headers` - Headers* `-headers` - Headers (descending)* `rate_limit` - Rate limit* `-rate_limit` - Rate limit (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	pulpLastUpdated := time.Now() // time.Time | Filter results where pulp_last_updated matches value (optional)
	pulpLastUpdatedGt := time.Now() // time.Time | Filter results where pulp_last_updated is greater than value (optional)
	pulpLastUpdatedGte := time.Now() // time.Time | Filter results where pulp_last_updated is greater than or equal to value (optional)
	pulpLastUpdatedLt := time.Now() // time.Time | Filter results where pulp_last_updated is less than value (optional)
	pulpLastUpdatedLte := time.Now() // time.Time | Filter results where pulp_last_updated is less than or equal to value (optional)
	pulpLastUpdatedRange := []time.Time{time.Now()} // []time.Time | Filter results where pulp_last_updated is between two comma separated values (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesGemAPI.RemotesGemGemList(context.Background(), pulpDomain).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesGemAPI.RemotesGemGemList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesGemGemList`: PaginatedgemGemRemoteResponseList
	fmt.Fprintf(os.Stdout, "Response from `RemotesGemAPI.RemotesGemGemList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesGemGemListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIexact** | **string** | Filter results where name matches value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameIregex** | **string** | Filter results where name matches regex value | 
 **nameIstartswith** | **string** | Filter results where name starts with value | 
 **nameRegex** | **string** | Filter results where name matches regex value | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;url&#x60; - Url* &#x60;-url&#x60; - Url (descending)* &#x60;ca_cert&#x60; - Ca cert* &#x60;-ca_cert&#x60; - Ca cert (descending)* &#x60;client_cert&#x60; - Client cert* &#x60;-client_cert&#x60; - Client cert (descending)* &#x60;client_key&#x60; - Client key* &#x60;-client_key&#x60; - Client key (descending)* &#x60;tls_validation&#x60; - Tls validation* &#x60;-tls_validation&#x60; - Tls validation (descending)* &#x60;username&#x60; - Username* &#x60;-username&#x60; - Username (descending)* &#x60;password&#x60; - Password* &#x60;-password&#x60; - Password (descending)* &#x60;proxy_url&#x60; - Proxy url* &#x60;-proxy_url&#x60; - Proxy url (descending)* &#x60;proxy_username&#x60; - Proxy username* &#x60;-proxy_username&#x60; - Proxy username (descending)* &#x60;proxy_password&#x60; - Proxy password* &#x60;-proxy_password&#x60; - Proxy password (descending)* &#x60;download_concurrency&#x60; - Download concurrency* &#x60;-download_concurrency&#x60; - Download concurrency (descending)* &#x60;max_retries&#x60; - Max retries* &#x60;-max_retries&#x60; - Max retries (descending)* &#x60;policy&#x60; - Policy* &#x60;-policy&#x60; - Policy (descending)* &#x60;total_timeout&#x60; - Total timeout* &#x60;-total_timeout&#x60; - Total timeout (descending)* &#x60;connect_timeout&#x60; - Connect timeout* &#x60;-connect_timeout&#x60; - Connect timeout (descending)* &#x60;sock_connect_timeout&#x60; - Sock connect timeout* &#x60;-sock_connect_timeout&#x60; - Sock connect timeout (descending)* &#x60;sock_read_timeout&#x60; - Sock read timeout* &#x60;-sock_read_timeout&#x60; - Sock read timeout (descending)* &#x60;headers&#x60; - Headers* &#x60;-headers&#x60; - Headers (descending)* &#x60;rate_limit&#x60; - Rate limit* &#x60;-rate_limit&#x60; - Rate limit (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **pulpLastUpdated** | **time.Time** | Filter results where pulp_last_updated matches value | 
 **pulpLastUpdatedGt** | **time.Time** | Filter results where pulp_last_updated is greater than value | 
 **pulpLastUpdatedGte** | **time.Time** | Filter results where pulp_last_updated is greater than or equal to value | 
 **pulpLastUpdatedLt** | **time.Time** | Filter results where pulp_last_updated is less than value | 
 **pulpLastUpdatedLte** | **time.Time** | Filter results where pulp_last_updated is less than or equal to value | 
 **pulpLastUpdatedRange** | [**[]time.Time**](time.Time.md) | Filter results where pulp_last_updated is between two comma separated values | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedgemGemRemoteResponseList**](PaginatedgemGemRemoteResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesGemGemListRoles

> ObjectRolesResponse RemotesGemGemListRoles(ctx, gemGemRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()

List roles



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
	gemGemRemoteHref := "gemGemRemoteHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesGemAPI.RemotesGemGemListRoles(context.Background(), gemGemRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesGemAPI.RemotesGemGemListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesGemGemListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesGemAPI.RemotesGemGemListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesGemGemListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ObjectRolesResponse**](ObjectRolesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesGemGemMyPermissions

> MyPermissionsResponse RemotesGemGemMyPermissions(ctx, gemGemRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()

List user permissions



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
	gemGemRemoteHref := "gemGemRemoteHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesGemAPI.RemotesGemGemMyPermissions(context.Background(), gemGemRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesGemAPI.RemotesGemGemMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesGemGemMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesGemAPI.RemotesGemGemMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesGemGemMyPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**MyPermissionsResponse**](MyPermissionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesGemGemPartialUpdate

> AsyncOperationResponse RemotesGemGemPartialUpdate(ctx, gemGemRemoteHref).PatchedgemGemRemote(patchedgemGemRemote).Execute()

Update a gem remote



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
	gemGemRemoteHref := "gemGemRemoteHref_example" // string | 
	patchedgemGemRemote := *openapiclient.NewPatchedgemGemRemote() // PatchedgemGemRemote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesGemAPI.RemotesGemGemPartialUpdate(context.Background(), gemGemRemoteHref).PatchedgemGemRemote(patchedgemGemRemote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesGemAPI.RemotesGemGemPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesGemGemPartialUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesGemAPI.RemotesGemGemPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesGemGemPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedgemGemRemote** | [**PatchedgemGemRemote**](PatchedgemGemRemote.md) |  | 

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


## RemotesGemGemRead

> GemGemRemoteResponse RemotesGemGemRead(ctx, gemGemRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a gem remote



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
	gemGemRemoteHref := "gemGemRemoteHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesGemAPI.RemotesGemGemRead(context.Background(), gemGemRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesGemAPI.RemotesGemGemRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesGemGemRead`: GemGemRemoteResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesGemAPI.RemotesGemGemRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesGemGemReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**GemGemRemoteResponse**](GemGemRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesGemGemRemoveRole

> NestedRoleResponse RemotesGemGemRemoveRole(ctx, gemGemRemoteHref).NestedRole(nestedRole).Execute()

Remove a role



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
	gemGemRemoteHref := "gemGemRemoteHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesGemAPI.RemotesGemGemRemoveRole(context.Background(), gemGemRemoteHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesGemAPI.RemotesGemGemRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesGemGemRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesGemAPI.RemotesGemGemRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesGemGemRemoveRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nestedRole** | [**NestedRole**](NestedRole.md) |  | 

### Return type

[**NestedRoleResponse**](NestedRoleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesGemGemSetLabel

> SetLabelResponse RemotesGemGemSetLabel(ctx, gemGemRemoteHref).SetLabel(setLabel).Execute()

Set a label



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
	gemGemRemoteHref := "gemGemRemoteHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesGemAPI.RemotesGemGemSetLabel(context.Background(), gemGemRemoteHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesGemAPI.RemotesGemGemSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesGemGemSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesGemAPI.RemotesGemGemSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesGemGemSetLabelRequest struct via the builder pattern


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


## RemotesGemGemUnsetLabel

> UnsetLabelResponse RemotesGemGemUnsetLabel(ctx, gemGemRemoteHref).UnsetLabel(unsetLabel).Execute()

Unset a label



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
	gemGemRemoteHref := "gemGemRemoteHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesGemAPI.RemotesGemGemUnsetLabel(context.Background(), gemGemRemoteHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesGemAPI.RemotesGemGemUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesGemGemUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesGemAPI.RemotesGemGemUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesGemGemUnsetLabelRequest struct via the builder pattern


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


## RemotesGemGemUpdate

> AsyncOperationResponse RemotesGemGemUpdate(ctx, gemGemRemoteHref).GemGemRemote(gemGemRemote).Execute()

Update a gem remote



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
	gemGemRemoteHref := "gemGemRemoteHref_example" // string | 
	gemGemRemote := *openapiclient.NewGemGemRemote("Name_example", "Url_example") // GemGemRemote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesGemAPI.RemotesGemGemUpdate(context.Background(), gemGemRemoteHref).GemGemRemote(gemGemRemote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesGemAPI.RemotesGemGemUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesGemGemUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesGemAPI.RemotesGemGemUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gemGemRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesGemGemUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gemGemRemote** | [**GemGemRemote**](GemGemRemote.md) |  | 

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

