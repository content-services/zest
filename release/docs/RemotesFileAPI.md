# \RemotesFileAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemotesFileFileAddRole**](RemotesFileAPI.md#RemotesFileFileAddRole) | **Post** /{file_file_remote_href}add_role/ | Add a role
[**RemotesFileFileCreate**](RemotesFileAPI.md#RemotesFileFileCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/remotes/file/file/ | Create a file remote
[**RemotesFileFileDelete**](RemotesFileAPI.md#RemotesFileFileDelete) | **Delete** /{file_file_remote_href} | Delete a file remote
[**RemotesFileFileList**](RemotesFileAPI.md#RemotesFileFileList) | **Get** /api/pulp/{pulp_domain}/api/v3/remotes/file/file/ | List file remotes
[**RemotesFileFileListRoles**](RemotesFileAPI.md#RemotesFileFileListRoles) | **Get** /{file_file_remote_href}list_roles/ | List roles
[**RemotesFileFileMyPermissions**](RemotesFileAPI.md#RemotesFileFileMyPermissions) | **Get** /{file_file_remote_href}my_permissions/ | List user permissions
[**RemotesFileFilePartialUpdate**](RemotesFileAPI.md#RemotesFileFilePartialUpdate) | **Patch** /{file_file_remote_href} | Update a file remote
[**RemotesFileFileRead**](RemotesFileAPI.md#RemotesFileFileRead) | **Get** /{file_file_remote_href} | Inspect a file remote
[**RemotesFileFileRemoveRole**](RemotesFileAPI.md#RemotesFileFileRemoveRole) | **Post** /{file_file_remote_href}remove_role/ | Remove a role
[**RemotesFileFileSetLabel**](RemotesFileAPI.md#RemotesFileFileSetLabel) | **Post** /{file_file_remote_href}set_label/ | Set a label
[**RemotesFileFileUnsetLabel**](RemotesFileAPI.md#RemotesFileFileUnsetLabel) | **Post** /{file_file_remote_href}unset_label/ | Unset a label
[**RemotesFileFileUpdate**](RemotesFileAPI.md#RemotesFileFileUpdate) | **Put** /{file_file_remote_href} | Update a file remote



## RemotesFileFileAddRole

> NestedRoleResponse RemotesFileFileAddRole(ctx, fileFileRemoteHref).NestedRole(nestedRole).Execute()

Add a role



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
	fileFileRemoteHref := "fileFileRemoteHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesFileAPI.RemotesFileFileAddRole(context.Background(), fileFileRemoteHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesFileAPI.RemotesFileFileAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesFileFileAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesFileAPI.RemotesFileFileAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesFileFileAddRoleRequest struct via the builder pattern


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


## RemotesFileFileCreate

> FileFileRemoteResponse RemotesFileFileCreate(ctx, pulpDomain).FileFileRemote(fileFileRemote).Execute()

Create a file remote



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
	fileFileRemote := *openapiclient.NewFileFileRemote("Name_example", "Url_example") // FileFileRemote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesFileAPI.RemotesFileFileCreate(context.Background(), pulpDomain).FileFileRemote(fileFileRemote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesFileAPI.RemotesFileFileCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesFileFileCreate`: FileFileRemoteResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesFileAPI.RemotesFileFileCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesFileFileCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileFileRemote** | [**FileFileRemote**](FileFileRemote.md) |  | 

### Return type

[**FileFileRemoteResponse**](FileFileRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesFileFileDelete

> AsyncOperationResponse RemotesFileFileDelete(ctx, fileFileRemoteHref).Execute()

Delete a file remote



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
	fileFileRemoteHref := "fileFileRemoteHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesFileAPI.RemotesFileFileDelete(context.Background(), fileFileRemoteHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesFileAPI.RemotesFileFileDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesFileFileDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesFileAPI.RemotesFileFileDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesFileFileDeleteRequest struct via the builder pattern


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


## RemotesFileFileList

> PaginatedfileFileRemoteResponseList RemotesFileFileList(ctx, pulpDomain).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedIsnull(pulpLastUpdatedIsnull).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()

List file remotes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/content-services/zest/release/v2025"
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
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	pulpLastUpdated := time.Now() // time.Time | Filter results where pulp_last_updated matches value (optional)
	pulpLastUpdatedGt := time.Now() // time.Time | Filter results where pulp_last_updated is greater than value (optional)
	pulpLastUpdatedGte := time.Now() // time.Time | Filter results where pulp_last_updated is greater than or equal to value (optional)
	pulpLastUpdatedIsnull := true // bool | Filter results where pulp_last_updated has a null value (optional)
	pulpLastUpdatedLt := time.Now() // time.Time | Filter results where pulp_last_updated is less than value (optional)
	pulpLastUpdatedLte := time.Now() // time.Time | Filter results where pulp_last_updated is less than or equal to value (optional)
	pulpLastUpdatedRange := []time.Time{time.Now()} // []time.Time | Filter results where pulp_last_updated is between two comma separated values (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesFileAPI.RemotesFileFileList(context.Background(), pulpDomain).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedIsnull(pulpLastUpdatedIsnull).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesFileAPI.RemotesFileFileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesFileFileList`: PaginatedfileFileRemoteResponseList
	fmt.Fprintf(os.Stdout, "Response from `RemotesFileAPI.RemotesFileFileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesFileFileListRequest struct via the builder pattern


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
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **pulpLastUpdated** | **time.Time** | Filter results where pulp_last_updated matches value | 
 **pulpLastUpdatedGt** | **time.Time** | Filter results where pulp_last_updated is greater than value | 
 **pulpLastUpdatedGte** | **time.Time** | Filter results where pulp_last_updated is greater than or equal to value | 
 **pulpLastUpdatedIsnull** | **bool** | Filter results where pulp_last_updated has a null value | 
 **pulpLastUpdatedLt** | **time.Time** | Filter results where pulp_last_updated is less than value | 
 **pulpLastUpdatedLte** | **time.Time** | Filter results where pulp_last_updated is less than or equal to value | 
 **pulpLastUpdatedRange** | [**[]time.Time**](time.Time.md) | Filter results where pulp_last_updated is between two comma separated values | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedfileFileRemoteResponseList**](PaginatedfileFileRemoteResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesFileFileListRoles

> ObjectRolesResponse RemotesFileFileListRoles(ctx, fileFileRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()

List roles



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
	fileFileRemoteHref := "fileFileRemoteHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesFileAPI.RemotesFileFileListRoles(context.Background(), fileFileRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesFileAPI.RemotesFileFileListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesFileFileListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesFileAPI.RemotesFileFileListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesFileFileListRolesRequest struct via the builder pattern


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


## RemotesFileFileMyPermissions

> MyPermissionsResponse RemotesFileFileMyPermissions(ctx, fileFileRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()

List user permissions



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
	fileFileRemoteHref := "fileFileRemoteHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesFileAPI.RemotesFileFileMyPermissions(context.Background(), fileFileRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesFileAPI.RemotesFileFileMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesFileFileMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesFileAPI.RemotesFileFileMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesFileFileMyPermissionsRequest struct via the builder pattern


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


## RemotesFileFilePartialUpdate

> AsyncOperationResponse RemotesFileFilePartialUpdate(ctx, fileFileRemoteHref).PatchedfileFileRemote(patchedfileFileRemote).Execute()

Update a file remote



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
	fileFileRemoteHref := "fileFileRemoteHref_example" // string | 
	patchedfileFileRemote := *openapiclient.NewPatchedfileFileRemote() // PatchedfileFileRemote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesFileAPI.RemotesFileFilePartialUpdate(context.Background(), fileFileRemoteHref).PatchedfileFileRemote(patchedfileFileRemote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesFileAPI.RemotesFileFilePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesFileFilePartialUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesFileAPI.RemotesFileFilePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesFileFilePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedfileFileRemote** | [**PatchedfileFileRemote**](PatchedfileFileRemote.md) |  | 

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


## RemotesFileFileRead

> FileFileRemoteResponse RemotesFileFileRead(ctx, fileFileRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a file remote



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
	fileFileRemoteHref := "fileFileRemoteHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesFileAPI.RemotesFileFileRead(context.Background(), fileFileRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesFileAPI.RemotesFileFileRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesFileFileRead`: FileFileRemoteResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesFileAPI.RemotesFileFileRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesFileFileReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**FileFileRemoteResponse**](FileFileRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesFileFileRemoveRole

> NestedRoleResponse RemotesFileFileRemoveRole(ctx, fileFileRemoteHref).NestedRole(nestedRole).Execute()

Remove a role



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
	fileFileRemoteHref := "fileFileRemoteHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesFileAPI.RemotesFileFileRemoveRole(context.Background(), fileFileRemoteHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesFileAPI.RemotesFileFileRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesFileFileRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesFileAPI.RemotesFileFileRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesFileFileRemoveRoleRequest struct via the builder pattern


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


## RemotesFileFileSetLabel

> SetLabelResponse RemotesFileFileSetLabel(ctx, fileFileRemoteHref).SetLabel(setLabel).Execute()

Set a label



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
	fileFileRemoteHref := "fileFileRemoteHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesFileAPI.RemotesFileFileSetLabel(context.Background(), fileFileRemoteHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesFileAPI.RemotesFileFileSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesFileFileSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesFileAPI.RemotesFileFileSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesFileFileSetLabelRequest struct via the builder pattern


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


## RemotesFileFileUnsetLabel

> UnsetLabelResponse RemotesFileFileUnsetLabel(ctx, fileFileRemoteHref).UnsetLabel(unsetLabel).Execute()

Unset a label



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
	fileFileRemoteHref := "fileFileRemoteHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesFileAPI.RemotesFileFileUnsetLabel(context.Background(), fileFileRemoteHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesFileAPI.RemotesFileFileUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesFileFileUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesFileAPI.RemotesFileFileUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesFileFileUnsetLabelRequest struct via the builder pattern


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


## RemotesFileFileUpdate

> AsyncOperationResponse RemotesFileFileUpdate(ctx, fileFileRemoteHref).FileFileRemote(fileFileRemote).Execute()

Update a file remote



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
	fileFileRemoteHref := "fileFileRemoteHref_example" // string | 
	fileFileRemote := *openapiclient.NewFileFileRemote("Name_example", "Url_example") // FileFileRemote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesFileAPI.RemotesFileFileUpdate(context.Background(), fileFileRemoteHref).FileFileRemote(fileFileRemote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesFileAPI.RemotesFileFileUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesFileFileUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesFileAPI.RemotesFileFileUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileFileRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesFileFileUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileFileRemote** | [**FileFileRemote**](FileFileRemote.md) |  | 

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

