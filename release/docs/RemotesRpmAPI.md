# \RemotesRpmAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemotesRpmRpmAddRole**](RemotesRpmAPI.md#RemotesRpmRpmAddRole) | **Post** /{rpm_rpm_remote_href}add_role/ | Add a role
[**RemotesRpmRpmCreate**](RemotesRpmAPI.md#RemotesRpmRpmCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/remotes/rpm/rpm/ | Create a rpm remote
[**RemotesRpmRpmDelete**](RemotesRpmAPI.md#RemotesRpmRpmDelete) | **Delete** /{rpm_rpm_remote_href} | Delete a rpm remote
[**RemotesRpmRpmList**](RemotesRpmAPI.md#RemotesRpmRpmList) | **Get** /api/pulp/{pulp_domain}/api/v3/remotes/rpm/rpm/ | List rpm remotes
[**RemotesRpmRpmListRoles**](RemotesRpmAPI.md#RemotesRpmRpmListRoles) | **Get** /{rpm_rpm_remote_href}list_roles/ | List roles
[**RemotesRpmRpmMyPermissions**](RemotesRpmAPI.md#RemotesRpmRpmMyPermissions) | **Get** /{rpm_rpm_remote_href}my_permissions/ | List user permissions
[**RemotesRpmRpmPartialUpdate**](RemotesRpmAPI.md#RemotesRpmRpmPartialUpdate) | **Patch** /{rpm_rpm_remote_href} | Update a rpm remote
[**RemotesRpmRpmRead**](RemotesRpmAPI.md#RemotesRpmRpmRead) | **Get** /{rpm_rpm_remote_href} | Inspect a rpm remote
[**RemotesRpmRpmRemoveRole**](RemotesRpmAPI.md#RemotesRpmRpmRemoveRole) | **Post** /{rpm_rpm_remote_href}remove_role/ | Remove a role
[**RemotesRpmRpmSetLabel**](RemotesRpmAPI.md#RemotesRpmRpmSetLabel) | **Post** /{rpm_rpm_remote_href}set_label/ | Set a label
[**RemotesRpmRpmUnsetLabel**](RemotesRpmAPI.md#RemotesRpmRpmUnsetLabel) | **Post** /{rpm_rpm_remote_href}unset_label/ | Unset a label
[**RemotesRpmRpmUpdate**](RemotesRpmAPI.md#RemotesRpmRpmUpdate) | **Put** /{rpm_rpm_remote_href} | Update a rpm remote



## RemotesRpmRpmAddRole

> NestedRoleResponse RemotesRpmRpmAddRole(ctx, rpmRpmRemoteHref).NestedRole(nestedRole).Execute()

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
	rpmRpmRemoteHref := "rpmRpmRemoteHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesRpmAPI.RemotesRpmRpmAddRole(context.Background(), rpmRpmRemoteHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesRpmAPI.RemotesRpmRpmAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesRpmRpmAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesRpmAPI.RemotesRpmRpmAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesRpmRpmAddRoleRequest struct via the builder pattern


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


## RemotesRpmRpmCreate

> RpmRpmRemoteResponse RemotesRpmRpmCreate(ctx, pulpDomain).RpmRpmRemote(rpmRpmRemote).Execute()

Create a rpm remote



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
	rpmRpmRemote := *openapiclient.NewRpmRpmRemote("Name_example", "Url_example") // RpmRpmRemote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesRpmAPI.RemotesRpmRpmCreate(context.Background(), pulpDomain).RpmRpmRemote(rpmRpmRemote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesRpmAPI.RemotesRpmRpmCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesRpmRpmCreate`: RpmRpmRemoteResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesRpmAPI.RemotesRpmRpmCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesRpmRpmCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rpmRpmRemote** | [**RpmRpmRemote**](RpmRpmRemote.md) |  | 

### Return type

[**RpmRpmRemoteResponse**](RpmRpmRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesRpmRpmDelete

> AsyncOperationResponse RemotesRpmRpmDelete(ctx, rpmRpmRemoteHref).Execute()

Delete a rpm remote



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
	rpmRpmRemoteHref := "rpmRpmRemoteHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesRpmAPI.RemotesRpmRpmDelete(context.Background(), rpmRpmRemoteHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesRpmAPI.RemotesRpmRpmDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesRpmRpmDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesRpmAPI.RemotesRpmRpmDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesRpmRpmDeleteRequest struct via the builder pattern


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


## RemotesRpmRpmList

> PaginatedrpmRpmRemoteResponseList RemotesRpmRpmList(ctx, pulpDomain).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedIsnull(pulpLastUpdatedIsnull).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()

List rpm remotes



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
	resp, r, err := apiClient.RemotesRpmAPI.RemotesRpmRpmList(context.Background(), pulpDomain).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedIsnull(pulpLastUpdatedIsnull).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesRpmAPI.RemotesRpmRpmList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesRpmRpmList`: PaginatedrpmRpmRemoteResponseList
	fmt.Fprintf(os.Stdout, "Response from `RemotesRpmAPI.RemotesRpmRpmList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesRpmRpmListRequest struct via the builder pattern


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

[**PaginatedrpmRpmRemoteResponseList**](PaginatedrpmRpmRemoteResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesRpmRpmListRoles

> ObjectRolesResponse RemotesRpmRpmListRoles(ctx, rpmRpmRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	rpmRpmRemoteHref := "rpmRpmRemoteHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesRpmAPI.RemotesRpmRpmListRoles(context.Background(), rpmRpmRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesRpmAPI.RemotesRpmRpmListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesRpmRpmListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesRpmAPI.RemotesRpmRpmListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesRpmRpmListRolesRequest struct via the builder pattern


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


## RemotesRpmRpmMyPermissions

> MyPermissionsResponse RemotesRpmRpmMyPermissions(ctx, rpmRpmRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	rpmRpmRemoteHref := "rpmRpmRemoteHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesRpmAPI.RemotesRpmRpmMyPermissions(context.Background(), rpmRpmRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesRpmAPI.RemotesRpmRpmMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesRpmRpmMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesRpmAPI.RemotesRpmRpmMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesRpmRpmMyPermissionsRequest struct via the builder pattern


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


## RemotesRpmRpmPartialUpdate

> AsyncOperationResponse RemotesRpmRpmPartialUpdate(ctx, rpmRpmRemoteHref).PatchedrpmRpmRemote(patchedrpmRpmRemote).Execute()

Update a rpm remote



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
	rpmRpmRemoteHref := "rpmRpmRemoteHref_example" // string | 
	patchedrpmRpmRemote := *openapiclient.NewPatchedrpmRpmRemote() // PatchedrpmRpmRemote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesRpmAPI.RemotesRpmRpmPartialUpdate(context.Background(), rpmRpmRemoteHref).PatchedrpmRpmRemote(patchedrpmRpmRemote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesRpmAPI.RemotesRpmRpmPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesRpmRpmPartialUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesRpmAPI.RemotesRpmRpmPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesRpmRpmPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedrpmRpmRemote** | [**PatchedrpmRpmRemote**](PatchedrpmRpmRemote.md) |  | 

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


## RemotesRpmRpmRead

> RpmRpmRemoteResponse RemotesRpmRpmRead(ctx, rpmRpmRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a rpm remote



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
	rpmRpmRemoteHref := "rpmRpmRemoteHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesRpmAPI.RemotesRpmRpmRead(context.Background(), rpmRpmRemoteHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesRpmAPI.RemotesRpmRpmRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesRpmRpmRead`: RpmRpmRemoteResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesRpmAPI.RemotesRpmRpmRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesRpmRpmReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RpmRpmRemoteResponse**](RpmRpmRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesRpmRpmRemoveRole

> NestedRoleResponse RemotesRpmRpmRemoveRole(ctx, rpmRpmRemoteHref).NestedRole(nestedRole).Execute()

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
	rpmRpmRemoteHref := "rpmRpmRemoteHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesRpmAPI.RemotesRpmRpmRemoveRole(context.Background(), rpmRpmRemoteHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesRpmAPI.RemotesRpmRpmRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesRpmRpmRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesRpmAPI.RemotesRpmRpmRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesRpmRpmRemoveRoleRequest struct via the builder pattern


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


## RemotesRpmRpmSetLabel

> SetLabelResponse RemotesRpmRpmSetLabel(ctx, rpmRpmRemoteHref).SetLabel(setLabel).Execute()

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
	rpmRpmRemoteHref := "rpmRpmRemoteHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesRpmAPI.RemotesRpmRpmSetLabel(context.Background(), rpmRpmRemoteHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesRpmAPI.RemotesRpmRpmSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesRpmRpmSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesRpmAPI.RemotesRpmRpmSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesRpmRpmSetLabelRequest struct via the builder pattern


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


## RemotesRpmRpmUnsetLabel

> UnsetLabelResponse RemotesRpmRpmUnsetLabel(ctx, rpmRpmRemoteHref).UnsetLabel(unsetLabel).Execute()

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
	rpmRpmRemoteHref := "rpmRpmRemoteHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesRpmAPI.RemotesRpmRpmUnsetLabel(context.Background(), rpmRpmRemoteHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesRpmAPI.RemotesRpmRpmUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesRpmRpmUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesRpmAPI.RemotesRpmRpmUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesRpmRpmUnsetLabelRequest struct via the builder pattern


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


## RemotesRpmRpmUpdate

> AsyncOperationResponse RemotesRpmRpmUpdate(ctx, rpmRpmRemoteHref).RpmRpmRemote(rpmRpmRemote).Execute()

Update a rpm remote



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
	rpmRpmRemoteHref := "rpmRpmRemoteHref_example" // string | 
	rpmRpmRemote := *openapiclient.NewRpmRpmRemote("Name_example", "Url_example") // RpmRpmRemote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesRpmAPI.RemotesRpmRpmUpdate(context.Background(), rpmRpmRemoteHref).RpmRpmRemote(rpmRpmRemote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesRpmAPI.RemotesRpmRpmUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesRpmRpmUpdate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesRpmAPI.RemotesRpmRpmUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmRpmRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesRpmRpmUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rpmRpmRemote** | [**RpmRpmRemote**](RpmRpmRemote.md) |  | 

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

