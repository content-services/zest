# \RemotesMavenAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemotesMavenMavenAddRole**](RemotesMavenAPI.md#RemotesMavenMavenAddRole) | **Post** /{maven_maven_remote_href}add_role/ | Add a role
[**RemotesMavenMavenCreate**](RemotesMavenAPI.md#RemotesMavenMavenCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/remotes/maven/maven/ | Create a maven remote
[**RemotesMavenMavenDelete**](RemotesMavenAPI.md#RemotesMavenMavenDelete) | **Delete** /{maven_maven_remote_href} | Delete a maven remote
[**RemotesMavenMavenList**](RemotesMavenAPI.md#RemotesMavenMavenList) | **Get** /api/pulp/{pulp_domain}/api/v3/remotes/maven/maven/ | List maven remotes
[**RemotesMavenMavenListRoles**](RemotesMavenAPI.md#RemotesMavenMavenListRoles) | **Get** /{maven_maven_remote_href}list_roles/ | List roles
[**RemotesMavenMavenMyPermissions**](RemotesMavenAPI.md#RemotesMavenMavenMyPermissions) | **Get** /{maven_maven_remote_href}my_permissions/ | List user permissions
[**RemotesMavenMavenPartialUpdate**](RemotesMavenAPI.md#RemotesMavenMavenPartialUpdate) | **Patch** /{maven_maven_remote_href} | Update a maven remote
[**RemotesMavenMavenRead**](RemotesMavenAPI.md#RemotesMavenMavenRead) | **Get** /{maven_maven_remote_href} | Inspect a maven remote
[**RemotesMavenMavenRemoveRole**](RemotesMavenAPI.md#RemotesMavenMavenRemoveRole) | **Post** /{maven_maven_remote_href}remove_role/ | Remove a role
[**RemotesMavenMavenSetLabel**](RemotesMavenAPI.md#RemotesMavenMavenSetLabel) | **Post** /{maven_maven_remote_href}set_label/ | Set a label
[**RemotesMavenMavenUnsetLabel**](RemotesMavenAPI.md#RemotesMavenMavenUnsetLabel) | **Post** /{maven_maven_remote_href}unset_label/ | Unset a label
[**RemotesMavenMavenUpdate**](RemotesMavenAPI.md#RemotesMavenMavenUpdate) | **Put** /{maven_maven_remote_href} | Update a maven remote



## RemotesMavenMavenAddRole

> NestedRoleResponse RemotesMavenMavenAddRole(ctx, mavenMavenRemoteHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()

Add a role



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
	mavenMavenRemoteHref := "mavenMavenRemoteHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesMavenAPI.RemotesMavenMavenAddRole(context.Background(), mavenMavenRemoteHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesMavenAPI.RemotesMavenMavenAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesMavenMavenAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesMavenAPI.RemotesMavenMavenAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesMavenMavenAddRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nestedRole** | [**NestedRole**](NestedRole.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## RemotesMavenMavenCreate

> MavenMavenRemoteResponse RemotesMavenMavenCreate(ctx, pulpDomain).MavenMavenRemote(mavenMavenRemote).XTaskDiagnostics(xTaskDiagnostics).Execute()

Create a maven remote



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
	mavenMavenRemote := *openapiclient.NewMavenMavenRemote("Name_example", "Url_example") // MavenMavenRemote | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesMavenAPI.RemotesMavenMavenCreate(context.Background(), pulpDomain).MavenMavenRemote(mavenMavenRemote).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesMavenAPI.RemotesMavenMavenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesMavenMavenCreate`: MavenMavenRemoteResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesMavenAPI.RemotesMavenMavenCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesMavenMavenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mavenMavenRemote** | [**MavenMavenRemote**](MavenMavenRemote.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**MavenMavenRemoteResponse**](MavenMavenRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesMavenMavenDelete

> AsyncOperationResponse RemotesMavenMavenDelete(ctx, mavenMavenRemoteHref).XTaskDiagnostics(xTaskDiagnostics).Execute()

Delete a maven remote



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
	mavenMavenRemoteHref := "mavenMavenRemoteHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesMavenAPI.RemotesMavenMavenDelete(context.Background(), mavenMavenRemoteHref).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesMavenAPI.RemotesMavenMavenDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesMavenMavenDelete`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesMavenAPI.RemotesMavenMavenDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesMavenMavenDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## RemotesMavenMavenList

> PaginatedmavenMavenRemoteResponseList RemotesMavenMavenList(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedIsnull(pulpLastUpdatedIsnull).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()

List maven remotes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/content-services/zest/release/v2026"
)

func main() {
	pulpDomain := "pulpDomain_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
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
	resp, r, err := apiClient.RemotesMavenAPI.RemotesMavenMavenList(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpLastUpdated(pulpLastUpdated).PulpLastUpdatedGt(pulpLastUpdatedGt).PulpLastUpdatedGte(pulpLastUpdatedGte).PulpLastUpdatedIsnull(pulpLastUpdatedIsnull).PulpLastUpdatedLt(pulpLastUpdatedLt).PulpLastUpdatedLte(pulpLastUpdatedLte).PulpLastUpdatedRange(pulpLastUpdatedRange).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesMavenAPI.RemotesMavenMavenList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesMavenMavenList`: PaginatedmavenMavenRemoteResponseList
	fmt.Fprintf(os.Stdout, "Response from `RemotesMavenAPI.RemotesMavenMavenList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesMavenMavenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
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

[**PaginatedmavenMavenRemoteResponseList**](PaginatedmavenMavenRemoteResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesMavenMavenListRoles

> ObjectRolesResponse RemotesMavenMavenListRoles(ctx, mavenMavenRemoteHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

List roles



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
	mavenMavenRemoteHref := "mavenMavenRemoteHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesMavenAPI.RemotesMavenMavenListRoles(context.Background(), mavenMavenRemoteHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesMavenAPI.RemotesMavenMavenListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesMavenMavenListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesMavenAPI.RemotesMavenMavenListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesMavenMavenListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
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


## RemotesMavenMavenMyPermissions

> MyPermissionsResponse RemotesMavenMavenMyPermissions(ctx, mavenMavenRemoteHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

List user permissions



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
	mavenMavenRemoteHref := "mavenMavenRemoteHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesMavenAPI.RemotesMavenMavenMyPermissions(context.Background(), mavenMavenRemoteHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesMavenAPI.RemotesMavenMavenMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesMavenMavenMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesMavenAPI.RemotesMavenMavenMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesMavenMavenMyPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
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


## RemotesMavenMavenPartialUpdate

> MavenMavenRemoteResponse RemotesMavenMavenPartialUpdate(ctx, mavenMavenRemoteHref).PatchedmavenMavenRemote(patchedmavenMavenRemote).XTaskDiagnostics(xTaskDiagnostics).Execute()

Update a maven remote



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
	mavenMavenRemoteHref := "mavenMavenRemoteHref_example" // string | 
	patchedmavenMavenRemote := *openapiclient.NewPatchedmavenMavenRemote() // PatchedmavenMavenRemote | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesMavenAPI.RemotesMavenMavenPartialUpdate(context.Background(), mavenMavenRemoteHref).PatchedmavenMavenRemote(patchedmavenMavenRemote).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesMavenAPI.RemotesMavenMavenPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesMavenMavenPartialUpdate`: MavenMavenRemoteResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesMavenAPI.RemotesMavenMavenPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesMavenMavenPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedmavenMavenRemote** | [**PatchedmavenMavenRemote**](PatchedmavenMavenRemote.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**MavenMavenRemoteResponse**](MavenMavenRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesMavenMavenRead

> MavenMavenRemoteResponse RemotesMavenMavenRead(ctx, mavenMavenRemoteHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a maven remote



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
	mavenMavenRemoteHref := "mavenMavenRemoteHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesMavenAPI.RemotesMavenMavenRead(context.Background(), mavenMavenRemoteHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesMavenAPI.RemotesMavenMavenRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesMavenMavenRead`: MavenMavenRemoteResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesMavenAPI.RemotesMavenMavenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesMavenMavenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**MavenMavenRemoteResponse**](MavenMavenRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemotesMavenMavenRemoveRole

> NestedRoleResponse RemotesMavenMavenRemoveRole(ctx, mavenMavenRemoteHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()

Remove a role



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
	mavenMavenRemoteHref := "mavenMavenRemoteHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesMavenAPI.RemotesMavenMavenRemoveRole(context.Background(), mavenMavenRemoteHref).NestedRole(nestedRole).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesMavenAPI.RemotesMavenMavenRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesMavenMavenRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesMavenAPI.RemotesMavenMavenRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesMavenMavenRemoveRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nestedRole** | [**NestedRole**](NestedRole.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## RemotesMavenMavenSetLabel

> SetLabelResponse RemotesMavenMavenSetLabel(ctx, mavenMavenRemoteHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	mavenMavenRemoteHref := "mavenMavenRemoteHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesMavenAPI.RemotesMavenMavenSetLabel(context.Background(), mavenMavenRemoteHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesMavenAPI.RemotesMavenMavenSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesMavenMavenSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesMavenAPI.RemotesMavenMavenSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesMavenMavenSetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setLabel** | [**SetLabel**](SetLabel.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## RemotesMavenMavenUnsetLabel

> UnsetLabelResponse RemotesMavenMavenUnsetLabel(ctx, mavenMavenRemoteHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	mavenMavenRemoteHref := "mavenMavenRemoteHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesMavenAPI.RemotesMavenMavenUnsetLabel(context.Background(), mavenMavenRemoteHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesMavenAPI.RemotesMavenMavenUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesMavenMavenUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesMavenAPI.RemotesMavenMavenUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesMavenMavenUnsetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unsetLabel** | [**UnsetLabel**](UnsetLabel.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

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


## RemotesMavenMavenUpdate

> MavenMavenRemoteResponse RemotesMavenMavenUpdate(ctx, mavenMavenRemoteHref).MavenMavenRemote(mavenMavenRemote).XTaskDiagnostics(xTaskDiagnostics).Execute()

Update a maven remote



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
	mavenMavenRemoteHref := "mavenMavenRemoteHref_example" // string | 
	mavenMavenRemote := *openapiclient.NewMavenMavenRemote("Name_example", "Url_example") // MavenMavenRemote | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotesMavenAPI.RemotesMavenMavenUpdate(context.Background(), mavenMavenRemoteHref).MavenMavenRemote(mavenMavenRemote).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotesMavenAPI.RemotesMavenMavenUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemotesMavenMavenUpdate`: MavenMavenRemoteResponse
	fmt.Fprintf(os.Stdout, "Response from `RemotesMavenAPI.RemotesMavenMavenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenRemoteHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemotesMavenMavenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mavenMavenRemote** | [**MavenMavenRemote**](MavenMavenRemote.md) |  | 
 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 

### Return type

[**MavenMavenRemoteResponse**](MavenMavenRemoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

