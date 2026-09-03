# \ContentPackageAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentMavenPackageList**](ContentPackageAPI.md#ContentMavenPackageList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/maven/package/ | List maven packages
[**ContentMavenPackageRead**](ContentPackageAPI.md#ContentMavenPackageRead) | **Get** /{maven_maven_package_href} | Inspect a maven package
[**ContentMavenPackageSetLabel**](ContentPackageAPI.md#ContentMavenPackageSetLabel) | **Post** /{maven_maven_package_href}set_label/ | Set a label
[**ContentMavenPackageUnsetLabel**](ContentPackageAPI.md#ContentMavenPackageUnsetLabel) | **Post** /{maven_maven_package_href}unset_label/ | Unset a label



## ContentMavenPackageList

> PaginatedmavenMavenPackageResponseList ContentMavenPackageList(ctx, pulpDomain).XTaskDiagnostics(xTaskDiagnostics).ArtifactId(artifactId).BaseVersion(baseVersion).CollapseBuilds(collapseBuilds).GroupId(groupId).Limit(limit).Name(name).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).Packaging(packaging).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Version(version).VersionStartswith(versionStartswith).Fields(fields).ExcludeFields(excludeFields).Execute()

List maven packages



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
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	artifactId := "artifactId_example" // string | Filter results where artifact_id matches value (optional)
	baseVersion := "baseVersion_example" // string | Match units whose version strips to this logical version (same suffix as collapse_builds: \\.[a-zA-Z]+-\\d+$). 5.3.18 matches 5.3.18 and 5.3.18.rhlw-00003, but not 5.3.180. (optional)
	collapseBuilds := true // bool | When true, collapse rebuilds of the same logical version: strip a trailing suffix matching \\.[a-zA-Z]+-\\d+$ from version, then keep one MavenPackage per (group_id, artifact_id, base_version) with the latest pulp_created. Default false. (optional)
	groupId := "groupId_example" // string | Filter results where group_id matches value (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string | Filter results where name matches value (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `group_id` - Group id* `-group_id` - Group id (descending)* `artifact_id` - Artifact id* `-artifact_id` - Artifact id (descending)* `version` - Version* `-version` - Version (descending)* `name` - Name* `-name` - Name (descending)* `description` - Description* `-description` - Description (descending)* `packaging` - Packaging* `-packaging` - Packaging (descending)* `url` - Url* `-url` - Url (descending)* `licenses` - Licenses* `-licenses` - Licenses (descending)* `dependencies` - Dependencies* `-dependencies` - Dependencies (descending)* `scm_url` - Scm url* `-scm_url` - Scm url (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	packaging := "packaging_example" // string | Filter results where packaging matches value (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repositoryVersion := "repositoryVersion_example" // string |  (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string |  (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string |  (optional)
	version := "version_example" // string | Filter results where version matches value (optional)
	versionStartswith := "versionStartswith_example" // string | Filter results where version starts with value (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentPackageAPI.ContentMavenPackageList(context.Background(), pulpDomain).XTaskDiagnostics(xTaskDiagnostics).ArtifactId(artifactId).BaseVersion(baseVersion).CollapseBuilds(collapseBuilds).GroupId(groupId).Limit(limit).Name(name).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).Packaging(packaging).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Version(version).VersionStartswith(versionStartswith).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentPackageAPI.ContentMavenPackageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentMavenPackageList`: PaginatedmavenMavenPackageResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentPackageAPI.ContentMavenPackageList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentMavenPackageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **artifactId** | **string** | Filter results where artifact_id matches value | 
 **baseVersion** | **string** | Match units whose version strips to this logical version (same suffix as collapse_builds: \\.[a-zA-Z]+-\\d+$). 5.3.18 matches 5.3.18 and 5.3.18.rhlw-00003, but not 5.3.180. | 
 **collapseBuilds** | **bool** | When true, collapse rebuilds of the same logical version: strip a trailing suffix matching \\.[a-zA-Z]+-\\d+$ from version, then keep one MavenPackage per (group_id, artifact_id, base_version) with the latest pulp_created. Default false. | 
 **groupId** | **string** | Filter results where group_id matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;group_id&#x60; - Group id* &#x60;-group_id&#x60; - Group id (descending)* &#x60;artifact_id&#x60; - Artifact id* &#x60;-artifact_id&#x60; - Artifact id (descending)* &#x60;version&#x60; - Version* &#x60;-version&#x60; - Version (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;description&#x60; - Description* &#x60;-description&#x60; - Description (descending)* &#x60;packaging&#x60; - Packaging* &#x60;-packaging&#x60; - Packaging (descending)* &#x60;url&#x60; - Url* &#x60;-url&#x60; - Url (descending)* &#x60;licenses&#x60; - Licenses* &#x60;-licenses&#x60; - Licenses (descending)* &#x60;dependencies&#x60; - Dependencies* &#x60;-dependencies&#x60; - Dependencies (descending)* &#x60;scm_url&#x60; - Scm url* &#x60;-scm_url&#x60; - Scm url (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **packaging** | **string** | Filter results where packaging matches value | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repositoryVersion** | **string** |  | 
 **repositoryVersionAdded** | **string** |  | 
 **repositoryVersionRemoved** | **string** |  | 
 **version** | **string** | Filter results where version matches value | 
 **versionStartswith** | **string** | Filter results where version starts with value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedmavenMavenPackageResponseList**](PaginatedmavenMavenPackageResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentMavenPackageRead

> MavenMavenPackageResponse ContentMavenPackageRead(ctx, mavenMavenPackageHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a maven package



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
	mavenMavenPackageHref := "mavenMavenPackageHref_example" // string | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentPackageAPI.ContentMavenPackageRead(context.Background(), mavenMavenPackageHref).XTaskDiagnostics(xTaskDiagnostics).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentPackageAPI.ContentMavenPackageRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentMavenPackageRead`: MavenMavenPackageResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentPackageAPI.ContentMavenPackageRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenPackageHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentMavenPackageReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTaskDiagnostics** | **[]string** | List of profilers to use on tasks. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**MavenMavenPackageResponse**](MavenMavenPackageResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentMavenPackageSetLabel

> SetLabelResponse ContentMavenPackageSetLabel(ctx, mavenMavenPackageHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	mavenMavenPackageHref := "mavenMavenPackageHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentPackageAPI.ContentMavenPackageSetLabel(context.Background(), mavenMavenPackageHref).SetLabel(setLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentPackageAPI.ContentMavenPackageSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentMavenPackageSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentPackageAPI.ContentMavenPackageSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenPackageHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentMavenPackageSetLabelRequest struct via the builder pattern


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


## ContentMavenPackageUnsetLabel

> UnsetLabelResponse ContentMavenPackageUnsetLabel(ctx, mavenMavenPackageHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()

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
	mavenMavenPackageHref := "mavenMavenPackageHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 
	xTaskDiagnostics := []string{"Inner_example"} // []string | List of profilers to use on tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentPackageAPI.ContentMavenPackageUnsetLabel(context.Background(), mavenMavenPackageHref).UnsetLabel(unsetLabel).XTaskDiagnostics(xTaskDiagnostics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentPackageAPI.ContentMavenPackageUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentMavenPackageUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentPackageAPI.ContentMavenPackageUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mavenMavenPackageHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentMavenPackageUnsetLabelRequest struct via the builder pattern


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

