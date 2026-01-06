# \UpstreamPulpsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpstreamPulpsAddRole**](UpstreamPulpsAPI.md#UpstreamPulpsAddRole) | **Post** /{upstream_pulp_href}add_role/ | Add a role
[**UpstreamPulpsCreate**](UpstreamPulpsAPI.md#UpstreamPulpsCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/upstream-pulps/ | Create an upstream pulp
[**UpstreamPulpsDelete**](UpstreamPulpsAPI.md#UpstreamPulpsDelete) | **Delete** /{upstream_pulp_href} | Delete an upstream pulp
[**UpstreamPulpsList**](UpstreamPulpsAPI.md#UpstreamPulpsList) | **Get** /api/pulp/{pulp_domain}/api/v3/upstream-pulps/ | List upstream pulps
[**UpstreamPulpsListRoles**](UpstreamPulpsAPI.md#UpstreamPulpsListRoles) | **Get** /{upstream_pulp_href}list_roles/ | List roles
[**UpstreamPulpsMyPermissions**](UpstreamPulpsAPI.md#UpstreamPulpsMyPermissions) | **Get** /{upstream_pulp_href}my_permissions/ | List user permissions
[**UpstreamPulpsPartialUpdate**](UpstreamPulpsAPI.md#UpstreamPulpsPartialUpdate) | **Patch** /{upstream_pulp_href} | Update an upstream pulp
[**UpstreamPulpsRead**](UpstreamPulpsAPI.md#UpstreamPulpsRead) | **Get** /{upstream_pulp_href} | Inspect an upstream pulp
[**UpstreamPulpsRemoveRole**](UpstreamPulpsAPI.md#UpstreamPulpsRemoveRole) | **Post** /{upstream_pulp_href}remove_role/ | Remove a role
[**UpstreamPulpsReplicate**](UpstreamPulpsAPI.md#UpstreamPulpsReplicate) | **Post** /{upstream_pulp_href}replicate/ | Replicate
[**UpstreamPulpsUpdate**](UpstreamPulpsAPI.md#UpstreamPulpsUpdate) | **Put** /{upstream_pulp_href} | Update an upstream pulp



## UpstreamPulpsAddRole

> NestedRoleResponse UpstreamPulpsAddRole(ctx, upstreamPulpHref).NestedRole(nestedRole).Execute()

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
	upstreamPulpHref := "upstreamPulpHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsAddRole(context.Background(), upstreamPulpHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsAddRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsAddRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsAddRoleRequest struct via the builder pattern


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


## UpstreamPulpsCreate

> UpstreamPulpResponse UpstreamPulpsCreate(ctx, pulpDomain).UpstreamPulp(upstreamPulp).Execute()

Create an upstream pulp



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
	upstreamPulp := *openapiclient.NewUpstreamPulp("Name_example", "BaseUrl_example", "ApiRoot_example") // UpstreamPulp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsCreate(context.Background(), pulpDomain).UpstreamPulp(upstreamPulp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsCreate`: UpstreamPulpResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upstreamPulp** | [**UpstreamPulp**](UpstreamPulp.md) |  | 

### Return type

[**UpstreamPulpResponse**](UpstreamPulpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsDelete

> UpstreamPulpsDelete(ctx, upstreamPulpHref).Execute()

Delete an upstream pulp



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
	upstreamPulpHref := "upstreamPulpHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsDelete(context.Background(), upstreamPulpHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsList

> PaginatedUpstreamPulpResponseList UpstreamPulpsList(ctx, pulpDomain).BaseUrl(baseUrl).BaseUrlContains(baseUrlContains).BaseUrlIcontains(baseUrlIcontains).BaseUrlIexact(baseUrlIexact).BaseUrlIn(baseUrlIn).BaseUrlIregex(baseUrlIregex).BaseUrlIstartswith(baseUrlIstartswith).BaseUrlRegex(baseUrlRegex).BaseUrlStartswith(baseUrlStartswith).LastReplication(lastReplication).LastReplicationGt(lastReplicationGt).LastReplicationGte(lastReplicationGte).LastReplicationIsnull(lastReplicationIsnull).LastReplicationLt(lastReplicationLt).LastReplicationLte(lastReplicationLte).LastReplicationRange(lastReplicationRange).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()

List upstream pulps



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
	baseUrl := "baseUrl_example" // string | Filter results where base_url matches value (optional)
	baseUrlContains := "baseUrlContains_example" // string | Filter results where base_url contains value (optional)
	baseUrlIcontains := "baseUrlIcontains_example" // string | Filter results where base_url contains value (optional)
	baseUrlIexact := "baseUrlIexact_example" // string | Filter results where base_url matches value (optional)
	baseUrlIn := []string{"Inner_example"} // []string | Filter results where base_url is in a comma-separated list of values (optional)
	baseUrlIregex := "baseUrlIregex_example" // string | Filter results where base_url matches regex value (optional)
	baseUrlIstartswith := "baseUrlIstartswith_example" // string | Filter results where base_url starts with value (optional)
	baseUrlRegex := "baseUrlRegex_example" // string | Filter results where base_url matches regex value (optional)
	baseUrlStartswith := "baseUrlStartswith_example" // string | Filter results where base_url starts with value (optional)
	lastReplication := time.Now() // time.Time | Filter results where last_replication matches value (optional)
	lastReplicationGt := time.Now() // time.Time | Filter results where last_replication is greater than value (optional)
	lastReplicationGte := time.Now() // time.Time | Filter results where last_replication is greater than or equal to value (optional)
	lastReplicationIsnull := true // bool | Filter results where last_replication has a null value (optional)
	lastReplicationLt := time.Now() // time.Time | Filter results where last_replication is less than value (optional)
	lastReplicationLte := time.Now() // time.Time | Filter results where last_replication is less than or equal to value (optional)
	lastReplicationRange := []time.Time{time.Now()} // []time.Time | Filter results where last_replication is between two comma separated values (optional)
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
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `name` - Name* `-name` - Name (descending)* `base_url` - Base url* `-base_url` - Base url (descending)* `api_root` - Api root* `-api_root` - Api root (descending)* `domain` - Domain* `-domain` - Domain (descending)* `ca_cert` - Ca cert* `-ca_cert` - Ca cert (descending)* `client_cert` - Client cert* `-client_cert` - Client cert (descending)* `client_key` - Client key* `-client_key` - Client key (descending)* `tls_validation` - Tls validation* `-tls_validation` - Tls validation (descending)* `username` - Username* `-username` - Username (descending)* `password` - Password* `-password` - Password (descending)* `q_select` - Q select* `-q_select` - Q select (descending)* `policy` - Policy* `-policy` - Policy (descending)* `last_replication` - Last replication* `-last_replication` - Last replication (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsList(context.Background(), pulpDomain).BaseUrl(baseUrl).BaseUrlContains(baseUrlContains).BaseUrlIcontains(baseUrlIcontains).BaseUrlIexact(baseUrlIexact).BaseUrlIn(baseUrlIn).BaseUrlIregex(baseUrlIregex).BaseUrlIstartswith(baseUrlIstartswith).BaseUrlRegex(baseUrlRegex).BaseUrlStartswith(baseUrlStartswith).LastReplication(lastReplication).LastReplicationGt(lastReplicationGt).LastReplicationGte(lastReplicationGte).LastReplicationIsnull(lastReplicationIsnull).LastReplicationLt(lastReplicationLt).LastReplicationLte(lastReplicationLte).LastReplicationRange(lastReplicationRange).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIexact(nameIexact).NameIn(nameIn).NameIregex(nameIregex).NameIstartswith(nameIstartswith).NameRegex(nameRegex).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Q(q).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsList`: PaginatedUpstreamPulpResponseList
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseUrl** | **string** | Filter results where base_url matches value | 
 **baseUrlContains** | **string** | Filter results where base_url contains value | 
 **baseUrlIcontains** | **string** | Filter results where base_url contains value | 
 **baseUrlIexact** | **string** | Filter results where base_url matches value | 
 **baseUrlIn** | **[]string** | Filter results where base_url is in a comma-separated list of values | 
 **baseUrlIregex** | **string** | Filter results where base_url matches regex value | 
 **baseUrlIstartswith** | **string** | Filter results where base_url starts with value | 
 **baseUrlRegex** | **string** | Filter results where base_url matches regex value | 
 **baseUrlStartswith** | **string** | Filter results where base_url starts with value | 
 **lastReplication** | **time.Time** | Filter results where last_replication matches value | 
 **lastReplicationGt** | **time.Time** | Filter results where last_replication is greater than value | 
 **lastReplicationGte** | **time.Time** | Filter results where last_replication is greater than or equal to value | 
 **lastReplicationIsnull** | **bool** | Filter results where last_replication has a null value | 
 **lastReplicationLt** | **time.Time** | Filter results where last_replication is less than value | 
 **lastReplicationLte** | **time.Time** | Filter results where last_replication is less than or equal to value | 
 **lastReplicationRange** | [**[]time.Time**](time.Time.md) | Filter results where last_replication is between two comma separated values | 
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
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;base_url&#x60; - Base url* &#x60;-base_url&#x60; - Base url (descending)* &#x60;api_root&#x60; - Api root* &#x60;-api_root&#x60; - Api root (descending)* &#x60;domain&#x60; - Domain* &#x60;-domain&#x60; - Domain (descending)* &#x60;ca_cert&#x60; - Ca cert* &#x60;-ca_cert&#x60; - Ca cert (descending)* &#x60;client_cert&#x60; - Client cert* &#x60;-client_cert&#x60; - Client cert (descending)* &#x60;client_key&#x60; - Client key* &#x60;-client_key&#x60; - Client key (descending)* &#x60;tls_validation&#x60; - Tls validation* &#x60;-tls_validation&#x60; - Tls validation (descending)* &#x60;username&#x60; - Username* &#x60;-username&#x60; - Username (descending)* &#x60;password&#x60; - Password* &#x60;-password&#x60; - Password (descending)* &#x60;q_select&#x60; - Q select* &#x60;-q_select&#x60; - Q select (descending)* &#x60;policy&#x60; - Policy* &#x60;-policy&#x60; - Policy (descending)* &#x60;last_replication&#x60; - Last replication* &#x60;-last_replication&#x60; - Last replication (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedUpstreamPulpResponseList**](PaginatedUpstreamPulpResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsListRoles

> ObjectRolesResponse UpstreamPulpsListRoles(ctx, upstreamPulpHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	upstreamPulpHref := "upstreamPulpHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsListRoles(context.Background(), upstreamPulpHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsListRoles`: ObjectRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsListRolesRequest struct via the builder pattern


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


## UpstreamPulpsMyPermissions

> MyPermissionsResponse UpstreamPulpsMyPermissions(ctx, upstreamPulpHref).Fields(fields).ExcludeFields(excludeFields).Execute()

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
	upstreamPulpHref := "upstreamPulpHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsMyPermissions(context.Background(), upstreamPulpHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsMyPermissions`: MyPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsMyPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsMyPermissionsRequest struct via the builder pattern


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


## UpstreamPulpsPartialUpdate

> UpstreamPulpResponse UpstreamPulpsPartialUpdate(ctx, upstreamPulpHref).PatchedUpstreamPulp(patchedUpstreamPulp).Execute()

Update an upstream pulp



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
	upstreamPulpHref := "upstreamPulpHref_example" // string | 
	patchedUpstreamPulp := *openapiclient.NewPatchedUpstreamPulp() // PatchedUpstreamPulp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsPartialUpdate(context.Background(), upstreamPulpHref).PatchedUpstreamPulp(patchedUpstreamPulp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsPartialUpdate`: UpstreamPulpResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedUpstreamPulp** | [**PatchedUpstreamPulp**](PatchedUpstreamPulp.md) |  | 

### Return type

[**UpstreamPulpResponse**](UpstreamPulpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsRead

> UpstreamPulpResponse UpstreamPulpsRead(ctx, upstreamPulpHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an upstream pulp



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
	upstreamPulpHref := "upstreamPulpHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsRead(context.Background(), upstreamPulpHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsRead`: UpstreamPulpResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**UpstreamPulpResponse**](UpstreamPulpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsRemoveRole

> NestedRoleResponse UpstreamPulpsRemoveRole(ctx, upstreamPulpHref).NestedRole(nestedRole).Execute()

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
	upstreamPulpHref := "upstreamPulpHref_example" // string | 
	nestedRole := *openapiclient.NewNestedRole("Role_example") // NestedRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsRemoveRole(context.Background(), upstreamPulpHref).NestedRole(nestedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsRemoveRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsRemoveRole`: NestedRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsRemoveRoleRequest struct via the builder pattern


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


## UpstreamPulpsReplicate

> TaskGroupOperationResponse UpstreamPulpsReplicate(ctx, upstreamPulpHref).Execute()

Replicate



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
	upstreamPulpHref := "upstreamPulpHref_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsReplicate(context.Background(), upstreamPulpHref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsReplicate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsReplicate`: TaskGroupOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsReplicate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsReplicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskGroupOperationResponse**](TaskGroupOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsUpdate

> UpstreamPulpResponse UpstreamPulpsUpdate(ctx, upstreamPulpHref).UpstreamPulp(upstreamPulp).Execute()

Update an upstream pulp



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
	upstreamPulpHref := "upstreamPulpHref_example" // string | 
	upstreamPulp := *openapiclient.NewUpstreamPulp("Name_example", "BaseUrl_example", "ApiRoot_example") // UpstreamPulp | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsUpdate(context.Background(), upstreamPulpHref).UpstreamPulp(upstreamPulp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsAPI.UpstreamPulpsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsUpdate`: UpstreamPulpResponse
	fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsAPI.UpstreamPulpsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upstreamPulp** | [**UpstreamPulp**](UpstreamPulp.md) |  | 

### Return type

[**UpstreamPulpResponse**](UpstreamPulpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

