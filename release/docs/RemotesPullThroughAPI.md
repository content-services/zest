# \RemotesPullThroughAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemotesContainerPullThroughAddRole**](RemotesPullThroughAPI.md#RemotesContainerPullThroughAddRole) | **Post** /{container_container_pull_through_remote_href}add_role/ | Add a role
[**RemotesContainerPullThroughCreate**](RemotesPullThroughAPI.md#RemotesContainerPullThroughCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/remotes/container/pull-through/ | Create a container pull through remote
[**RemotesContainerPullThroughDelete**](RemotesPullThroughAPI.md#RemotesContainerPullThroughDelete) | **Delete** /{container_container_pull_through_remote_href} | Delete a container pull through remote
[**RemotesContainerPullThroughList**](RemotesPullThroughAPI.md#RemotesContainerPullThroughList) | **Get** /api/pulp/{pulp_domain}/api/v3/remotes/container/pull-through/ | List container pull through remotes
[**RemotesContainerPullThroughListRoles**](RemotesPullThroughAPI.md#RemotesContainerPullThroughListRoles) | **Get** /{container_container_pull_through_remote_href}list_roles/ | List roles
[**RemotesContainerPullThroughMyPermissions**](RemotesPullThroughAPI.md#RemotesContainerPullThroughMyPermissions) | **Get** /{container_container_pull_through_remote_href}my_permissions/ | List user permissions
[**RemotesContainerPullThroughPartialUpdate**](RemotesPullThroughAPI.md#RemotesContainerPullThroughPartialUpdate) | **Patch** /{container_container_pull_through_remote_href} | Update a container pull through remote
[**RemotesContainerPullThroughRead**](RemotesPullThroughAPI.md#RemotesContainerPullThroughRead) | **Get** /{container_container_pull_through_remote_href} | Inspect a container pull through remote
[**RemotesContainerPullThroughRemoveRole**](RemotesPullThroughAPI.md#RemotesContainerPullThroughRemoveRole) | **Post** /{container_container_pull_through_remote_href}remove_role/ | Remove a role
[**RemotesContainerPullThroughSetLabel**](RemotesPullThroughAPI.md#RemotesContainerPullThroughSetLabel) | **Post** /{container_container_pull_through_remote_href}set_label/ | Set a label
[**RemotesContainerPullThroughUnsetLabel**](RemotesPullThroughAPI.md#RemotesContainerPullThroughUnsetLabel) | **Post** /{container_container_pull_through_remote_href}unset_label/ | Unset a label
[**RemotesContainerPullThroughUpdate**](RemotesPullThroughAPI.md#RemotesContainerPullThroughUpdate) | **Put** /{container_container_pull_through_remote_href} | Update a container pull through remote



## RemotesContainerPullThroughAddRole

> NestedRoleResponse RemotesContainerPullThroughAddRole(ctx, containerContainerPullThroughRemoteHref).NestedRole(nestedRole).Execute()

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
	containerContainerPullThroughRemoteHref := "containerContainerPullThroughRemoteHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesPullThroughAPI.RemotesContainerPullThroughAddRole(context.Background(), containerContainerPullThroughRemoteHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesPullThroughAPI.RemotesContainerPullThroughAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesContainerPullThroughAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesPullThroughAPI.RemotesContainerPullThroughAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPullThroughRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerPullThroughAddRoleRequest struct via the builder pattern


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


## RemotesContainerPullThroughCreate

> ContainerContainerPullThroughRemoteResponse RemotesContainerPullThroughCreate(ctx, pulpDomain).ContainerContainerPullThroughRemote(containerContainerPullThroughRemote).Execute()

Create a container pull through remote



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
	containerContainerPullThroughRemote := *openapiclient.NewContainerContainerPullThroughRemote("Name_example", "Url_example") // ContainerContainerPullThroughRemote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesPullThroughAPI.RemotesContainerPullThroughCreate(context.Background(), pulpDomain).ContainerContainerPullThroughRemote(containerContainerPullThroughRemote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesPullThroughAPI.RemotesContainerPullThroughCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesContainerPullThroughCreate`: ContainerContainerPullThroughRemoteResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesPullThroughAPI.RemotesContainerPullThroughCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerPullThroughCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerContainerPullThroughRemote** | [**ContainerContainerPullThroughRemote**](ContainerContainerPullThroughRemote.md) |  | 

### Return type

[**ContainerContainerPullThroughRemoteResponse**](ContainerContainerPullThroughRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesContainerPullThroughDelete

> AsyncOperationResponse RemotesContainerPullThroughDelete(ctx, containerContainerPullThroughRemoteHref).Execute()

Delete a container pull through remote



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
	containerContainerPullThroughRemoteHref := "containerContainerPullThroughRemoteHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesPullThroughAPI.RemotesContainerPullThroughDelete(context.Background(), containerContainerPullThroughRemoteHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesPullThroughAPI.RemotesContainerPullThroughDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesContainerPullThroughDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesPullThroughAPI.RemotesContainerPullThroughDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPullThroughRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerPullThroughDeleteRequest struct via the builder pattern


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


## RemotesContainerPullThroughList

> PaginatedcontainerContainerPullThroughRemoteResponseList RemotesContainerPullThroughList(ctx, pulpDomain).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedIsnull(pulpLastUpdatedIsnull).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()

List container pull through remotes



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
	resp, r, err := apiClient.RemotesPullThroughAPI.RemotesContainerPullThroughList(context.Background(), pulpDomain).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedIsnull(pulpLastUpdatedIsnull).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesPullThroughAPI.RemotesContainerPullThroughList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesContainerPullThroughList`: PaginatedcontainerContainerPullThroughRemoteResponseList
	fmt.Fprintf(os.Stdout, "Response from `RemotesPullThroughAPI.RemotesContainerPullThroughList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerPullThroughListRequest struct via the builder pattern


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

[**PaginatedcontainerContainerPullThroughRemoteResponseList**](PaginatedcontainerContainerPullThroughRemoteResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesContainerPullThroughListRoles

> ObjectRolesResponse RemotesContainerPullThroughListRoles(ctx, containerContainerPullThroughRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	containerContainerPullThroughRemoteHref := "containerContainerPullThroughRemoteHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesPullThroughAPI.RemotesContainerPullThroughListRoles(context.Background(), containerContainerPullThroughRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesPullThroughAPI.RemotesContainerPullThroughListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesContainerPullThroughListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesPullThroughAPI.RemotesContainerPullThroughListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPullThroughRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerPullThroughListRolesRequest struct via the builder pattern


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


## RemotesContainerPullThroughMyPermissions

> MyPermissionsResponse RemotesContainerPullThroughMyPermissions(ctx, containerContainerPullThroughRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	containerContainerPullThroughRemoteHref := "containerContainerPullThroughRemoteHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesPullThroughAPI.RemotesContainerPullThroughMyPermissions(context.Background(), containerContainerPullThroughRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesPullThroughAPI.RemotesContainerPullThroughMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesContainerPullThroughMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesPullThroughAPI.RemotesContainerPullThroughMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPullThroughRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerPullThroughMyPermissionsRequest struct via the builder pattern


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


## RemotesContainerPullThroughPartialUpdate

> AsyncOperationResponse RemotesContainerPullThroughPartialUpdate(ctx, containerContainerPullThroughRemoteHref).PatchedcontainerContainerPullThroughRemote(patchedcontainerContainerPullThroughRemote).Execute()

Update a container pull through remote



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
	containerContainerPullThroughRemoteHref := "containerContainerPullThroughRemoteHref_example" // string | 
	patchedcontainerContainerPullThroughRemote := *openapiclient.NewPatchedcontainerContainerPullThroughRemote() // PatchedcontainerContainerPullThroughRemote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesPullThroughAPI.RemotesContainerPullThroughPartialUpdate(context.Background(), containerContainerPullThroughRemoteHref).PatchedcontainerContainerPullThroughRemote(patchedcontainerContainerPullThroughRemote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesPullThroughAPI.RemotesContainerPullThroughPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesContainerPullThroughPartialUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesPullThroughAPI.RemotesContainerPullThroughPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPullThroughRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerPullThroughPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedcontainerContainerPullThroughRemote** | [**PatchedcontainerContainerPullThroughRemote**](PatchedcontainerContainerPullThroughRemote.md) |  | 

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


## RemotesContainerPullThroughRead

> ContainerContainerPullThroughRemoteResponse RemotesContainerPullThroughRead(ctx, containerContainerPullThroughRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a container pull through remote



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
	containerContainerPullThroughRemoteHref := "containerContainerPullThroughRemoteHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesPullThroughAPI.RemotesContainerPullThroughRead(context.Background(), containerContainerPullThroughRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesPullThroughAPI.RemotesContainerPullThroughRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesContainerPullThroughRead`: ContainerContainerPullThroughRemoteResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesPullThroughAPI.RemotesContainerPullThroughRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPullThroughRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerPullThroughReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ContainerContainerPullThroughRemoteResponse**](ContainerContainerPullThroughRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesContainerPullThroughRemoveRole

> NestedRoleResponse RemotesContainerPullThroughRemoveRole(ctx, containerContainerPullThroughRemoteHref).NestedRole(nestedRole).Execute()

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
	containerContainerPullThroughRemoteHref := "containerContainerPullThroughRemoteHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesPullThroughAPI.RemotesContainerPullThroughRemoveRole(context.Background(), containerContainerPullThroughRemoteHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesPullThroughAPI.RemotesContainerPullThroughRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesContainerPullThroughRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesPullThroughAPI.RemotesContainerPullThroughRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPullThroughRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerPullThroughRemoveRoleRequest struct via the builder pattern


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


## RemotesContainerPullThroughSetLabel

> SetLabelResponse RemotesContainerPullThroughSetLabel(ctx, containerContainerPullThroughRemoteHref).SetLabel(setLabel).Execute()

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
	containerContainerPullThroughRemoteHref := "containerContainerPullThroughRemoteHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesPullThroughAPI.RemotesContainerPullThroughSetLabel(context.Background(), containerContainerPullThroughRemoteHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesPullThroughAPI.RemotesContainerPullThroughSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesContainerPullThroughSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesPullThroughAPI.RemotesContainerPullThroughSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPullThroughRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerPullThroughSetLabelRequest struct via the builder pattern


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


## RemotesContainerPullThroughUnsetLabel

> UnsetLabelResponse RemotesContainerPullThroughUnsetLabel(ctx, containerContainerPullThroughRemoteHref).UnsetLabel(unsetLabel).Execute()

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
	containerContainerPullThroughRemoteHref := "containerContainerPullThroughRemoteHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesPullThroughAPI.RemotesContainerPullThroughUnsetLabel(context.Background(), containerContainerPullThroughRemoteHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesPullThroughAPI.RemotesContainerPullThroughUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesContainerPullThroughUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesPullThroughAPI.RemotesContainerPullThroughUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPullThroughRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerPullThroughUnsetLabelRequest struct via the builder pattern


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


## RemotesContainerPullThroughUpdate

> AsyncOperationResponse RemotesContainerPullThroughUpdate(ctx, containerContainerPullThroughRemoteHref).ContainerContainerPullThroughRemote(containerContainerPullThroughRemote).Execute()

Update a container pull through remote



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
	containerContainerPullThroughRemoteHref := "containerContainerPullThroughRemoteHref_example" // string | 
	containerContainerPullThroughRemote := *openapiclient.NewContainerContainerPullThroughRemote("Name_example", "Url_example") // ContainerContainerPullThroughRemote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesPullThroughAPI.RemotesContainerPullThroughUpdate(context.Background(), containerContainerPullThroughRemoteHref).ContainerContainerPullThroughRemote(containerContainerPullThroughRemote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesPullThroughAPI.RemotesContainerPullThroughUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesContainerPullThroughUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesPullThroughAPI.RemotesContainerPullThroughUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerContainerPullThroughRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesContainerPullThroughUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerContainerPullThroughRemote** | [**ContainerContainerPullThroughRemote**](ContainerContainerPullThroughRemote.md) |  | 

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

