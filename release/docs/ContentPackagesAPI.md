# \ContentPackagesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentNpmPackagesCreate**](ContentPackagesAPI.md#ContentNpmPackagesCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/content/npm/packages/ | Create a package
[**ContentNpmPackagesList**](ContentPackagesAPI.md#ContentNpmPackagesList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/npm/packages/ | List packages
[**ContentNpmPackagesRead**](ContentPackagesAPI.md#ContentNpmPackagesRead) | **Get** /{npm_package_href} | Inspect a package
[**ContentNpmPackagesSetLabel**](ContentPackagesAPI.md#ContentNpmPackagesSetLabel) | **Post** /{npm_package_href}set_label/ | Set a label
[**ContentNpmPackagesUnsetLabel**](ContentPackagesAPI.md#ContentNpmPackagesUnsetLabel) | **Post** /{npm_package_href}unset_label/ | Unset a label
[**ContentPythonPackagesCreate**](ContentPackagesAPI.md#ContentPythonPackagesCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/content/python/packages/ | Create a python package content
[**ContentPythonPackagesList**](ContentPackagesAPI.md#ContentPythonPackagesList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/python/packages/ | List python package contents
[**ContentPythonPackagesRead**](ContentPackagesAPI.md#ContentPythonPackagesRead) | **Get** /{python_python_package_content_href} | Inspect a python package content
[**ContentPythonPackagesSetLabel**](ContentPackagesAPI.md#ContentPythonPackagesSetLabel) | **Post** /{python_python_package_content_href}set_label/ | Set a label
[**ContentPythonPackagesUnsetLabel**](ContentPackagesAPI.md#ContentPythonPackagesUnsetLabel) | **Post** /{python_python_package_content_href}unset_label/ | Unset a label
[**ContentRpmPackagesCreate**](ContentPackagesAPI.md#ContentRpmPackagesCreate) | **Post** /api/pulp/{pulp_domain}/api/v3/content/rpm/packages/ | Create a package
[**ContentRpmPackagesList**](ContentPackagesAPI.md#ContentRpmPackagesList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/rpm/packages/ | List packages



## ContentNpmPackagesCreate

> NpmPackageResponse ContentNpmPackagesCreate(ctx, pulpDomain).RelativePath(relativePath).Name(name).Version(version).Repository(repository).PulpLabels(pulpLabels).Artifact(artifact).File(file).Upload(upload).FileUrl(fileUrl).Execute()

Create a package



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
	relativePath := "relativePath_example" // string | 
	name := "name_example" // string | 
	version := "version_example" // string | 
	repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
	pulpLabels := map[string]string{"key": "Inner_example"} // map[string]string | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. (optional)
	artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
	file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the content unit. (optional)
	upload := "upload_example" // string | An uncommitted upload that may be turned into the content unit. (optional)
	fileUrl := "fileUrl_example" // string | A url that Pulp can download and turn into the content unit. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentPackagesAPI.ContentNpmPackagesCreate(context.Background(), pulpDomain).RelativePath(relativePath).Name(name).Version(version).Repository(repository).PulpLabels(pulpLabels).Artifact(artifact).File(file).Upload(upload).FileUrl(fileUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesAPI.ContentNpmPackagesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentNpmPackagesCreate`: NpmPackageResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentPackagesAPI.ContentNpmPackagesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentNpmPackagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **relativePath** | **string** |  | 
 **name** | **string** |  | 
 **version** | **string** |  | 
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **pulpLabels** | **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | 
 **artifact** | **string** | Artifact file representing the physical content | 
 **file** | ***os.File** | An uploaded file that may be turned into the content unit. | 
 **upload** | **string** | An uncommitted upload that may be turned into the content unit. | 
 **fileUrl** | **string** | A url that Pulp can download and turn into the content unit. | 

### Return type

[**NpmPackageResponse**](NpmPackageResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentNpmPackagesList

> PaginatednpmPackageResponseList ContentNpmPackagesList(ctx, pulpDomain).Limit(limit).Name(name).NameIn(nameIn).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List packages



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
	name := "name_example" // string | Filter results where name matches value (optional)
	nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `name` - Name* `-name` - Name (descending)* `version` - Version* `-version` - Version (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF/PRN (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentPackagesAPI.ContentNpmPackagesList(context.Background(), pulpDomain).Limit(limit).Name(name).NameIn(nameIn).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesAPI.ContentNpmPackagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentNpmPackagesList`: PaginatednpmPackageResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentPackagesAPI.ContentNpmPackagesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentNpmPackagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;version&#x60; - Version* &#x60;-version&#x60; - Version (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF/PRN | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatednpmPackageResponseList**](PaginatednpmPackageResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentNpmPackagesRead

> NpmPackageResponse ContentNpmPackagesRead(ctx, npmPackageHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a package



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
	npmPackageHref := "npmPackageHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentPackagesAPI.ContentNpmPackagesRead(context.Background(), npmPackageHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesAPI.ContentNpmPackagesRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentNpmPackagesRead`: NpmPackageResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentPackagesAPI.ContentNpmPackagesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**npmPackageHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentNpmPackagesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**NpmPackageResponse**](NpmPackageResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentNpmPackagesSetLabel

> SetLabelResponse ContentNpmPackagesSetLabel(ctx, npmPackageHref).SetLabel(setLabel).Execute()

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
	npmPackageHref := "npmPackageHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentPackagesAPI.ContentNpmPackagesSetLabel(context.Background(), npmPackageHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesAPI.ContentNpmPackagesSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentNpmPackagesSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentPackagesAPI.ContentNpmPackagesSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**npmPackageHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentNpmPackagesSetLabelRequest struct via the builder pattern


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


## ContentNpmPackagesUnsetLabel

> UnsetLabelResponse ContentNpmPackagesUnsetLabel(ctx, npmPackageHref).UnsetLabel(unsetLabel).Execute()

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
	npmPackageHref := "npmPackageHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentPackagesAPI.ContentNpmPackagesUnsetLabel(context.Background(), npmPackageHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesAPI.ContentNpmPackagesUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentNpmPackagesUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentPackagesAPI.ContentNpmPackagesUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**npmPackageHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentNpmPackagesUnsetLabelRequest struct via the builder pattern


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


## ContentPythonPackagesCreate

> AsyncOperationResponse ContentPythonPackagesCreate(ctx, pulpDomain).RelativePath(relativePath).Repository(repository).PulpLabels(pulpLabels).Artifact(artifact).File(file).Upload(upload).FileUrl(fileUrl).Sha256(sha256).Summary(summary).Description(description).DescriptionContentType(descriptionContentType).Keywords(keywords).HomePage(homePage).DownloadUrl(downloadUrl).Author(author).AuthorEmail(authorEmail).Maintainer(maintainer).MaintainerEmail(maintainerEmail).License(license).RequiresPython(requiresPython).ProjectUrl(projectUrl).ProjectUrls(projectUrls).Platform(platform).SupportedPlatform(supportedPlatform).RequiresDist(requiresDist).ProvidesDist(providesDist).ObsoletesDist(obsoletesDist).RequiresExternal(requiresExternal).Classifiers(classifiers).Execute()

Create a python package content



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
	relativePath := "relativePath_example" // string | Path where the artifact is located relative to distributions base_path
	repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
	pulpLabels := map[string]string{"key": "Inner_example"} // map[string]string | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. (optional)
	artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
	file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the content unit. (optional)
	upload := "upload_example" // string | An uncommitted upload that may be turned into the content unit. (optional)
	fileUrl := "fileUrl_example" // string | A url that Pulp can download and turn into the content unit. (optional)
	sha256 := "sha256_example" // string | The SHA256 digest of this package. (optional) (default to "")
	summary := "summary_example" // string | A one-line summary of what the package does. (optional)
	description := "description_example" // string | A longer description of the package that can run to several paragraphs. (optional)
	descriptionContentType := "descriptionContentType_example" // string | A string stating the markup syntax (if any) used in the distribution’s description, so that tools can intelligently render the description. (optional)
	keywords := "keywords_example" // string | Additional keywords to be used to assist searching for the package in a larger catalog. (optional)
	homePage := "homePage_example" // string | The URL for the package's home page. (optional)
	downloadUrl := "downloadUrl_example" // string | Legacy field denoting the URL from which this package can be downloaded. (optional)
	author := "author_example" // string | Text containing the author's name. Contact information can also be added, separated with newlines. (optional)
	authorEmail := "authorEmail_example" // string | The author's e-mail address.  (optional)
	maintainer := "maintainer_example" // string | The maintainer's name at a minimum; additional contact information may be provided. (optional)
	maintainerEmail := "maintainerEmail_example" // string | The maintainer's e-mail address. (optional)
	license := "license_example" // string | Text indicating the license covering the distribution (optional)
	requiresPython := "requiresPython_example" // string | The Python version(s) that the distribution is guaranteed to be compatible with. (optional)
	projectUrl := "projectUrl_example" // string | A browsable URL for the project and a label for it, separated by a comma. (optional)
	projectUrls := TODO // interface{} | A dictionary of labels and URLs for the project. (optional)
	platform := "platform_example" // string | A comma-separated list of platform specifications, summarizing the operating systems supported by the package. (optional)
	supportedPlatform := "supportedPlatform_example" // string | Field to specify the OS and CPU for which the binary package was compiled.  (optional)
	requiresDist := TODO // interface{} | A JSON list containing names of some other distutils project required by this distribution. (optional)
	providesDist := TODO // interface{} | A JSON list containing names of a Distutils project which is contained within this distribution. (optional)
	obsoletesDist := TODO // interface{} | A JSON list containing names of a distutils project's distribution which this distribution renders obsolete, meaning that the two projects should not be installed at the same time. (optional)
	requiresExternal := TODO // interface{} | A JSON list containing some dependency in the system that the distribution is to be used. (optional)
	classifiers := TODO // interface{} | A JSON list containing classification values for a Python package. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentPackagesAPI.ContentPythonPackagesCreate(context.Background(), pulpDomain).RelativePath(relativePath).Repository(repository).PulpLabels(pulpLabels).Artifact(artifact).File(file).Upload(upload).FileUrl(fileUrl).Sha256(sha256).Summary(summary).Description(description).DescriptionContentType(descriptionContentType).Keywords(keywords).HomePage(homePage).DownloadUrl(downloadUrl).Author(author).AuthorEmail(authorEmail).Maintainer(maintainer).MaintainerEmail(maintainerEmail).License(license).RequiresPython(requiresPython).ProjectUrl(projectUrl).ProjectUrls(projectUrls).Platform(platform).SupportedPlatform(supportedPlatform).RequiresDist(requiresDist).ProvidesDist(providesDist).ObsoletesDist(obsoletesDist).RequiresExternal(requiresExternal).Classifiers(classifiers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesAPI.ContentPythonPackagesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentPythonPackagesCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentPackagesAPI.ContentPythonPackagesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentPythonPackagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **relativePath** | **string** | Path where the artifact is located relative to distributions base_path | 
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **pulpLabels** | **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | 
 **artifact** | **string** | Artifact file representing the physical content | 
 **file** | ***os.File** | An uploaded file that may be turned into the content unit. | 
 **upload** | **string** | An uncommitted upload that may be turned into the content unit. | 
 **fileUrl** | **string** | A url that Pulp can download and turn into the content unit. | 
 **sha256** | **string** | The SHA256 digest of this package. | [default to &quot;&quot;]
 **summary** | **string** | A one-line summary of what the package does. | 
 **description** | **string** | A longer description of the package that can run to several paragraphs. | 
 **descriptionContentType** | **string** | A string stating the markup syntax (if any) used in the distribution’s description, so that tools can intelligently render the description. | 
 **keywords** | **string** | Additional keywords to be used to assist searching for the package in a larger catalog. | 
 **homePage** | **string** | The URL for the package&#39;s home page. | 
 **downloadUrl** | **string** | Legacy field denoting the URL from which this package can be downloaded. | 
 **author** | **string** | Text containing the author&#39;s name. Contact information can also be added, separated with newlines. | 
 **authorEmail** | **string** | The author&#39;s e-mail address.  | 
 **maintainer** | **string** | The maintainer&#39;s name at a minimum; additional contact information may be provided. | 
 **maintainerEmail** | **string** | The maintainer&#39;s e-mail address. | 
 **license** | **string** | Text indicating the license covering the distribution | 
 **requiresPython** | **string** | The Python version(s) that the distribution is guaranteed to be compatible with. | 
 **projectUrl** | **string** | A browsable URL for the project and a label for it, separated by a comma. | 
 **projectUrls** | [**interface{}**](interface{}.md) | A dictionary of labels and URLs for the project. | 
 **platform** | **string** | A comma-separated list of platform specifications, summarizing the operating systems supported by the package. | 
 **supportedPlatform** | **string** | Field to specify the OS and CPU for which the binary package was compiled.  | 
 **requiresDist** | [**interface{}**](interface{}.md) | A JSON list containing names of some other distutils project required by this distribution. | 
 **providesDist** | [**interface{}**](interface{}.md) | A JSON list containing names of a Distutils project which is contained within this distribution. | 
 **obsoletesDist** | [**interface{}**](interface{}.md) | A JSON list containing names of a distutils project&#39;s distribution which this distribution renders obsolete, meaning that the two projects should not be installed at the same time. | 
 **requiresExternal** | [**interface{}**](interface{}.md) | A JSON list containing some dependency in the system that the distribution is to be used. | 
 **classifiers** | [**interface{}**](interface{}.md) | A JSON list containing classification values for a Python package. | 

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


## ContentPythonPackagesList

> PaginatedpythonPythonPackageContentResponseList ContentPythonPackagesList(ctx, pulpDomain).Author(author).AuthorIn(authorIn).Filename(filename).FilenameContains(filenameContains).FilenameIn(filenameIn).KeywordsContains(keywordsContains).KeywordsIn(keywordsIn).Limit(limit).Name(name).NameIn(nameIn).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).Packagetype(packagetype).PackagetypeIn(packagetypeIn).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).RequiresPython(requiresPython).RequiresPythonContains(requiresPythonContains).RequiresPythonIn(requiresPythonIn).Sha256(sha256).Sha256In(sha256In).Version(version).VersionGt(versionGt).VersionGte(versionGte).VersionLt(versionLt).VersionLte(versionLte).Fields(fields).ExcludeFields(excludeFields).Execute()

List python package contents



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
	author := "author_example" // string | Filter results where author matches value (optional)
	authorIn := []string{"Inner_example"} // []string | Filter results where author is in a comma-separated list of values (optional)
	filename := "filename_example" // string | Filter results where filename matches value (optional)
	filenameContains := "filenameContains_example" // string | Filter results where filename contains value (optional)
	filenameIn := []string{"Inner_example"} // []string | Filter results where filename is in a comma-separated list of values (optional)
	keywordsContains := "keywordsContains_example" // string | Filter results where keywords contains value (optional)
	keywordsIn := []string{"Inner_example"} // []string | Filter results where keywords is in a comma-separated list of values (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string | Filter results where name matches value (optional)
	nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `filename` - Filename* `-filename` - Filename (descending)* `packagetype` - Packagetype* `-packagetype` - Packagetype (descending)* `name` - Name* `-name` - Name (descending)* `version` - Version* `-version` - Version (descending)* `sha256` - Sha256* `-sha256` - Sha256 (descending)* `python_version` - Python version* `-python_version` - Python version (descending)* `metadata_version` - Metadata version* `-metadata_version` - Metadata version (descending)* `summary` - Summary* `-summary` - Summary (descending)* `description` - Description* `-description` - Description (descending)* `keywords` - Keywords* `-keywords` - Keywords (descending)* `home_page` - Home page* `-home_page` - Home page (descending)* `download_url` - Download url* `-download_url` - Download url (descending)* `author` - Author* `-author` - Author (descending)* `author_email` - Author email* `-author_email` - Author email (descending)* `maintainer` - Maintainer* `-maintainer` - Maintainer (descending)* `maintainer_email` - Maintainer email* `-maintainer_email` - Maintainer email (descending)* `license` - License* `-license` - License (descending)* `requires_python` - Requires python* `-requires_python` - Requires python (descending)* `project_url` - Project url* `-project_url` - Project url (descending)* `platform` - Platform* `-platform` - Platform (descending)* `supported_platform` - Supported platform* `-supported_platform` - Supported platform (descending)* `requires_dist` - Requires dist* `-requires_dist` - Requires dist (descending)* `provides_dist` - Provides dist* `-provides_dist` - Provides dist (descending)* `obsoletes_dist` - Obsoletes dist* `-obsoletes_dist` - Obsoletes dist (descending)* `requires_external` - Requires external* `-requires_external` - Requires external (descending)* `classifiers` - Classifiers* `-classifiers` - Classifiers (descending)* `project_urls` - Project urls* `-project_urls` - Project urls (descending)* `description_content_type` - Description content type* `-description_content_type` - Description content type (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	packagetype := "packagetype_example" // string | Filter results where packagetype matches value* `bdist_dmg` - bdist_dmg* `bdist_dumb` - bdist_dumb* `bdist_egg` - bdist_egg* `bdist_msi` - bdist_msi* `bdist_rpm` - bdist_rpm* `bdist_wheel` - bdist_wheel* `bdist_wininst` - bdist_wininst* `sdist` - sdist (optional)
	packagetypeIn := []string{"Inner_example"} // []string | Filter results where packagetype is in a comma-separated list of values (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF/PRN (optional)
	requiresPython := "requiresPython_example" // string | Filter results where requires_python matches value (optional)
	requiresPythonContains := "requiresPythonContains_example" // string | Filter results where requires_python contains value (optional)
	requiresPythonIn := []string{"Inner_example"} // []string | Filter results where requires_python is in a comma-separated list of values (optional)
	sha256 := "sha256_example" // string | Filter results where sha256 matches value (optional)
	sha256In := []string{"Inner_example"} // []string | Filter results where sha256 is in a comma-separated list of values (optional)
	version := "version_example" // string | Filter results where version matches value (optional)
	versionGt := "versionGt_example" // string | Filter results where version is greater than value (optional)
	versionGte := "versionGte_example" // string | Filter results where version is greater than or equal to value (optional)
	versionLt := "versionLt_example" // string | Filter results where version is less than value (optional)
	versionLte := "versionLte_example" // string | Filter results where version is less than or equal to value (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentPackagesAPI.ContentPythonPackagesList(context.Background(), pulpDomain).Author(author).AuthorIn(authorIn).Filename(filename).FilenameContains(filenameContains).FilenameIn(filenameIn).KeywordsContains(keywordsContains).KeywordsIn(keywordsIn).Limit(limit).Name(name).NameIn(nameIn).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).Packagetype(packagetype).PackagetypeIn(packagetypeIn).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).RequiresPython(requiresPython).RequiresPythonContains(requiresPythonContains).RequiresPythonIn(requiresPythonIn).Sha256(sha256).Sha256In(sha256In).Version(version).VersionGt(versionGt).VersionGte(versionGte).VersionLt(versionLt).VersionLte(versionLte).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesAPI.ContentPythonPackagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentPythonPackagesList`: PaginatedpythonPythonPackageContentResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentPackagesAPI.ContentPythonPackagesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentPythonPackagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **author** | **string** | Filter results where author matches value | 
 **authorIn** | **[]string** | Filter results where author is in a comma-separated list of values | 
 **filename** | **string** | Filter results where filename matches value | 
 **filenameContains** | **string** | Filter results where filename contains value | 
 **filenameIn** | **[]string** | Filter results where filename is in a comma-separated list of values | 
 **keywordsContains** | **string** | Filter results where keywords contains value | 
 **keywordsIn** | **[]string** | Filter results where keywords is in a comma-separated list of values | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;filename&#x60; - Filename* &#x60;-filename&#x60; - Filename (descending)* &#x60;packagetype&#x60; - Packagetype* &#x60;-packagetype&#x60; - Packagetype (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;version&#x60; - Version* &#x60;-version&#x60; - Version (descending)* &#x60;sha256&#x60; - Sha256* &#x60;-sha256&#x60; - Sha256 (descending)* &#x60;python_version&#x60; - Python version* &#x60;-python_version&#x60; - Python version (descending)* &#x60;metadata_version&#x60; - Metadata version* &#x60;-metadata_version&#x60; - Metadata version (descending)* &#x60;summary&#x60; - Summary* &#x60;-summary&#x60; - Summary (descending)* &#x60;description&#x60; - Description* &#x60;-description&#x60; - Description (descending)* &#x60;keywords&#x60; - Keywords* &#x60;-keywords&#x60; - Keywords (descending)* &#x60;home_page&#x60; - Home page* &#x60;-home_page&#x60; - Home page (descending)* &#x60;download_url&#x60; - Download url* &#x60;-download_url&#x60; - Download url (descending)* &#x60;author&#x60; - Author* &#x60;-author&#x60; - Author (descending)* &#x60;author_email&#x60; - Author email* &#x60;-author_email&#x60; - Author email (descending)* &#x60;maintainer&#x60; - Maintainer* &#x60;-maintainer&#x60; - Maintainer (descending)* &#x60;maintainer_email&#x60; - Maintainer email* &#x60;-maintainer_email&#x60; - Maintainer email (descending)* &#x60;license&#x60; - License* &#x60;-license&#x60; - License (descending)* &#x60;requires_python&#x60; - Requires python* &#x60;-requires_python&#x60; - Requires python (descending)* &#x60;project_url&#x60; - Project url* &#x60;-project_url&#x60; - Project url (descending)* &#x60;platform&#x60; - Platform* &#x60;-platform&#x60; - Platform (descending)* &#x60;supported_platform&#x60; - Supported platform* &#x60;-supported_platform&#x60; - Supported platform (descending)* &#x60;requires_dist&#x60; - Requires dist* &#x60;-requires_dist&#x60; - Requires dist (descending)* &#x60;provides_dist&#x60; - Provides dist* &#x60;-provides_dist&#x60; - Provides dist (descending)* &#x60;obsoletes_dist&#x60; - Obsoletes dist* &#x60;-obsoletes_dist&#x60; - Obsoletes dist (descending)* &#x60;requires_external&#x60; - Requires external* &#x60;-requires_external&#x60; - Requires external (descending)* &#x60;classifiers&#x60; - Classifiers* &#x60;-classifiers&#x60; - Classifiers (descending)* &#x60;project_urls&#x60; - Project urls* &#x60;-project_urls&#x60; - Project urls (descending)* &#x60;description_content_type&#x60; - Description content type* &#x60;-description_content_type&#x60; - Description content type (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **packagetype** | **string** | Filter results where packagetype matches value* &#x60;bdist_dmg&#x60; - bdist_dmg* &#x60;bdist_dumb&#x60; - bdist_dumb* &#x60;bdist_egg&#x60; - bdist_egg* &#x60;bdist_msi&#x60; - bdist_msi* &#x60;bdist_rpm&#x60; - bdist_rpm* &#x60;bdist_wheel&#x60; - bdist_wheel* &#x60;bdist_wininst&#x60; - bdist_wininst* &#x60;sdist&#x60; - sdist | 
 **packagetypeIn** | **[]string** | Filter results where packagetype is in a comma-separated list of values | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF/PRN | 
 **requiresPython** | **string** | Filter results where requires_python matches value | 
 **requiresPythonContains** | **string** | Filter results where requires_python contains value | 
 **requiresPythonIn** | **[]string** | Filter results where requires_python is in a comma-separated list of values | 
 **sha256** | **string** | Filter results where sha256 matches value | 
 **sha256In** | **[]string** | Filter results where sha256 is in a comma-separated list of values | 
 **version** | **string** | Filter results where version matches value | 
 **versionGt** | **string** | Filter results where version is greater than value | 
 **versionGte** | **string** | Filter results where version is greater than or equal to value | 
 **versionLt** | **string** | Filter results where version is less than value | 
 **versionLte** | **string** | Filter results where version is less than or equal to value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedpythonPythonPackageContentResponseList**](PaginatedpythonPythonPackageContentResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentPythonPackagesRead

> PythonPythonPackageContentResponse ContentPythonPackagesRead(ctx, pythonPythonPackageContentHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a python package content



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
	pythonPythonPackageContentHref := "pythonPythonPackageContentHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentPackagesAPI.ContentPythonPackagesRead(context.Background(), pythonPythonPackageContentHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesAPI.ContentPythonPackagesRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentPythonPackagesRead`: PythonPythonPackageContentResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentPackagesAPI.ContentPythonPackagesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pythonPythonPackageContentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentPythonPackagesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PythonPythonPackageContentResponse**](PythonPythonPackageContentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentPythonPackagesSetLabel

> SetLabelResponse ContentPythonPackagesSetLabel(ctx, pythonPythonPackageContentHref).SetLabel(setLabel).Execute()

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
	pythonPythonPackageContentHref := "pythonPythonPackageContentHref_example" // string | 
	setLabel := *openapiclient.NewSetLabel("Key_example", "Value_example") // SetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentPackagesAPI.ContentPythonPackagesSetLabel(context.Background(), pythonPythonPackageContentHref).SetLabel(setLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesAPI.ContentPythonPackagesSetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentPythonPackagesSetLabel`: SetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentPackagesAPI.ContentPythonPackagesSetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pythonPythonPackageContentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentPythonPackagesSetLabelRequest struct via the builder pattern


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


## ContentPythonPackagesUnsetLabel

> UnsetLabelResponse ContentPythonPackagesUnsetLabel(ctx, pythonPythonPackageContentHref).UnsetLabel(unsetLabel).Execute()

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
	pythonPythonPackageContentHref := "pythonPythonPackageContentHref_example" // string | 
	unsetLabel := *openapiclient.NewUnsetLabel("Key_example") // UnsetLabel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentPackagesAPI.ContentPythonPackagesUnsetLabel(context.Background(), pythonPythonPackageContentHref).UnsetLabel(unsetLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesAPI.ContentPythonPackagesUnsetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentPythonPackagesUnsetLabel`: UnsetLabelResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentPackagesAPI.ContentPythonPackagesUnsetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pythonPythonPackageContentHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentPythonPackagesUnsetLabelRequest struct via the builder pattern


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


## ContentRpmPackagesCreate

> AsyncOperationResponse ContentRpmPackagesCreate(ctx, pulpDomain).Repository(repository).PulpLabels(pulpLabels).Artifact(artifact).RelativePath(relativePath).File(file).Upload(upload).FileUrl(fileUrl).Execute()

Create a package



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
	repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
	pulpLabels := map[string]string{"key": "Inner_example"} // map[string]string | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. (optional)
	artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
	relativePath := "relativePath_example" // string | Path where the artifact is located relative to distributions base_path (optional)
	file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the content unit. (optional)
	upload := "upload_example" // string | An uncommitted upload that may be turned into the content unit. (optional)
	fileUrl := "fileUrl_example" // string | A url that Pulp can download and turn into the content unit. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentPackagesAPI.ContentRpmPackagesCreate(context.Background(), pulpDomain).Repository(repository).PulpLabels(pulpLabels).Artifact(artifact).RelativePath(relativePath).File(file).Upload(upload).FileUrl(fileUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesAPI.ContentRpmPackagesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentRpmPackagesCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentPackagesAPI.ContentRpmPackagesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmPackagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **pulpLabels** | **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | 
 **artifact** | **string** | Artifact file representing the physical content | 
 **relativePath** | **string** | Path where the artifact is located relative to distributions base_path | 
 **file** | ***os.File** | An uploaded file that may be turned into the content unit. | 
 **upload** | **string** | An uncommitted upload that may be turned into the content unit. | 
 **fileUrl** | **string** | A url that Pulp can download and turn into the content unit. | 

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


## ContentRpmPackagesList

> PaginatedrpmPackageResponseList ContentRpmPackagesList(ctx, pulpDomain).Arch(arch).ArchContains(archContains).ArchIn(archIn).ArchNe(archNe).ArchStartswith(archStartswith).ChecksumType(checksumType).ChecksumTypeIn(checksumTypeIn).ChecksumTypeNe(checksumTypeNe).Epoch(epoch).EpochIn(epochIn).EpochNe(epochNe).Filename(filename).Limit(limit).Name(name).NameContains(nameContains).NameIn(nameIn).NameNe(nameNe).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PkgId(pkgId).PkgIdIn(pkgIdIn).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Release(release).ReleaseContains(releaseContains).ReleaseIn(releaseIn).ReleaseNe(releaseNe).ReleaseStartswith(releaseStartswith).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Version(version).VersionIn(versionIn).VersionNe(versionNe).Fields(fields).ExcludeFields(excludeFields).Execute()

List packages



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
	arch := "arch_example" // string | Filter results where arch matches value (optional)
	archContains := "archContains_example" // string | Filter results where arch contains value (optional)
	archIn := []string{"Inner_example"} // []string | Filter results where arch is in a comma-separated list of values (optional)
	archNe := "archNe_example" // string | Filter results where arch not equal to value (optional)
	archStartswith := "archStartswith_example" // string | Filter results where arch starts with value (optional)
	checksumType := "checksumType_example" // string | Filter results where checksum_type matches value* `unknown` - unknown* `md5` - md5* `sha1` - sha1* `sha1` - sha1* `sha224` - sha224* `sha256` - sha256* `sha384` - sha384* `sha512` - sha512 (optional)
	checksumTypeIn := []string{"Inner_example"} // []string | Filter results where checksum_type is in a comma-separated list of values (optional)
	checksumTypeNe := "checksumTypeNe_example" // string | Filter results where checksum_type not equal to value (optional)
	epoch := "epoch_example" // string | Filter results where epoch matches value (optional)
	epochIn := []string{"Inner_example"} // []string | Filter results where epoch is in a comma-separated list of values (optional)
	epochNe := "epochNe_example" // string | Filter results where epoch not equal to value (optional)
	filename := "filename_example" // string |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string | Filter results where name matches value (optional)
	nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
	nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
	nameNe := "nameNe_example" // string | Filter results where name not equal to value (optional)
	nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `name` - Name* `-name` - Name (descending)* `epoch` - Epoch* `-epoch` - Epoch (descending)* `version` - Version* `-version` - Version (descending)* `release` - Release* `-release` - Release (descending)* `arch` - Arch* `-arch` - Arch (descending)* `evr` - Evr* `-evr` - Evr (descending)* `pkgId` - Pkgid* `-pkgId` - Pkgid (descending)* `checksum_type` - Checksum type* `-checksum_type` - Checksum type (descending)* `summary` - Summary* `-summary` - Summary (descending)* `description` - Description* `-description` - Description (descending)* `url` - Url* `-url` - Url (descending)* `changelogs` - Changelogs* `-changelogs` - Changelogs (descending)* `files` - Files* `-files` - Files (descending)* `requires` - Requires* `-requires` - Requires (descending)* `provides` - Provides* `-provides` - Provides (descending)* `conflicts` - Conflicts* `-conflicts` - Conflicts (descending)* `obsoletes` - Obsoletes* `-obsoletes` - Obsoletes (descending)* `suggests` - Suggests* `-suggests` - Suggests (descending)* `enhances` - Enhances* `-enhances` - Enhances (descending)* `recommends` - Recommends* `-recommends` - Recommends (descending)* `supplements` - Supplements* `-supplements` - Supplements (descending)* `location_base` - Location base* `-location_base` - Location base (descending)* `location_href` - Location href* `-location_href` - Location href (descending)* `rpm_buildhost` - Rpm buildhost* `-rpm_buildhost` - Rpm buildhost (descending)* `rpm_group` - Rpm group* `-rpm_group` - Rpm group (descending)* `rpm_license` - Rpm license* `-rpm_license` - Rpm license (descending)* `rpm_packager` - Rpm packager* `-rpm_packager` - Rpm packager (descending)* `rpm_sourcerpm` - Rpm sourcerpm* `-rpm_sourcerpm` - Rpm sourcerpm (descending)* `rpm_vendor` - Rpm vendor* `-rpm_vendor` - Rpm vendor (descending)* `rpm_header_start` - Rpm header start* `-rpm_header_start` - Rpm header start (descending)* `rpm_header_end` - Rpm header end* `-rpm_header_end` - Rpm header end (descending)* `size_archive` - Size archive* `-size_archive` - Size archive (descending)* `size_installed` - Size installed* `-size_installed` - Size installed (descending)* `size_package` - Size package* `-size_package` - Size package (descending)* `time_build` - Time build* `-time_build` - Time build (descending)* `time_file` - Time file* `-time_file` - Time file (descending)* `is_modular` - Is modular* `-is_modular` - Is modular (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	pkgId := "pkgId_example" // string | Filter results where pkgId matches value (optional)
	pkgIdIn := []string{"Inner_example"} // []string | Filter results where pkgId is in a comma-separated list of values (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	release := "release_example" // string | Filter results where release matches value (optional)
	releaseContains := "releaseContains_example" // string | Filter results where release contains value (optional)
	releaseIn := []string{"Inner_example"} // []string | Filter results where release is in a comma-separated list of values (optional)
	releaseNe := "releaseNe_example" // string | Filter results where release not equal to value (optional)
	releaseStartswith := "releaseStartswith_example" // string | Filter results where release starts with value (optional)
	repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF/PRN (optional)
	sha256 := "sha256_example" // string |  (optional)
	version := "version_example" // string | Filter results where version matches value (optional)
	versionIn := []string{"Inner_example"} // []string | Filter results where version is in a comma-separated list of values (optional)
	versionNe := "versionNe_example" // string | Filter results where version not equal to value (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentPackagesAPI.ContentRpmPackagesList(context.Background(), pulpDomain).Arch(arch).ArchContains(archContains).ArchIn(archIn).ArchNe(archNe).ArchStartswith(archStartswith).ChecksumType(checksumType).ChecksumTypeIn(checksumTypeIn).ChecksumTypeNe(checksumTypeNe).Epoch(epoch).EpochIn(epochIn).EpochNe(epochNe).Filename(filename).Limit(limit).Name(name).NameContains(nameContains).NameIn(nameIn).NameNe(nameNe).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PkgId(pkgId).PkgIdIn(pkgIdIn).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).Release(release).ReleaseContains(releaseContains).ReleaseIn(releaseIn).ReleaseNe(releaseNe).ReleaseStartswith(releaseStartswith).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Version(version).VersionIn(versionIn).VersionNe(versionNe).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesAPI.ContentRpmPackagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentRpmPackagesList`: PaginatedrpmPackageResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentPackagesAPI.ContentRpmPackagesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmPackagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **arch** | **string** | Filter results where arch matches value | 
 **archContains** | **string** | Filter results where arch contains value | 
 **archIn** | **[]string** | Filter results where arch is in a comma-separated list of values | 
 **archNe** | **string** | Filter results where arch not equal to value | 
 **archStartswith** | **string** | Filter results where arch starts with value | 
 **checksumType** | **string** | Filter results where checksum_type matches value* &#x60;unknown&#x60; - unknown* &#x60;md5&#x60; - md5* &#x60;sha1&#x60; - sha1* &#x60;sha1&#x60; - sha1* &#x60;sha224&#x60; - sha224* &#x60;sha256&#x60; - sha256* &#x60;sha384&#x60; - sha384* &#x60;sha512&#x60; - sha512 | 
 **checksumTypeIn** | **[]string** | Filter results where checksum_type is in a comma-separated list of values | 
 **checksumTypeNe** | **string** | Filter results where checksum_type not equal to value | 
 **epoch** | **string** | Filter results where epoch matches value | 
 **epochIn** | **[]string** | Filter results where epoch is in a comma-separated list of values | 
 **epochNe** | **string** | Filter results where epoch not equal to value | 
 **filename** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameNe** | **string** | Filter results where name not equal to value | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;name&#x60; - Name* &#x60;-name&#x60; - Name (descending)* &#x60;epoch&#x60; - Epoch* &#x60;-epoch&#x60; - Epoch (descending)* &#x60;version&#x60; - Version* &#x60;-version&#x60; - Version (descending)* &#x60;release&#x60; - Release* &#x60;-release&#x60; - Release (descending)* &#x60;arch&#x60; - Arch* &#x60;-arch&#x60; - Arch (descending)* &#x60;evr&#x60; - Evr* &#x60;-evr&#x60; - Evr (descending)* &#x60;pkgId&#x60; - Pkgid* &#x60;-pkgId&#x60; - Pkgid (descending)* &#x60;checksum_type&#x60; - Checksum type* &#x60;-checksum_type&#x60; - Checksum type (descending)* &#x60;summary&#x60; - Summary* &#x60;-summary&#x60; - Summary (descending)* &#x60;description&#x60; - Description* &#x60;-description&#x60; - Description (descending)* &#x60;url&#x60; - Url* &#x60;-url&#x60; - Url (descending)* &#x60;changelogs&#x60; - Changelogs* &#x60;-changelogs&#x60; - Changelogs (descending)* &#x60;files&#x60; - Files* &#x60;-files&#x60; - Files (descending)* &#x60;requires&#x60; - Requires* &#x60;-requires&#x60; - Requires (descending)* &#x60;provides&#x60; - Provides* &#x60;-provides&#x60; - Provides (descending)* &#x60;conflicts&#x60; - Conflicts* &#x60;-conflicts&#x60; - Conflicts (descending)* &#x60;obsoletes&#x60; - Obsoletes* &#x60;-obsoletes&#x60; - Obsoletes (descending)* &#x60;suggests&#x60; - Suggests* &#x60;-suggests&#x60; - Suggests (descending)* &#x60;enhances&#x60; - Enhances* &#x60;-enhances&#x60; - Enhances (descending)* &#x60;recommends&#x60; - Recommends* &#x60;-recommends&#x60; - Recommends (descending)* &#x60;supplements&#x60; - Supplements* &#x60;-supplements&#x60; - Supplements (descending)* &#x60;location_base&#x60; - Location base* &#x60;-location_base&#x60; - Location base (descending)* &#x60;location_href&#x60; - Location href* &#x60;-location_href&#x60; - Location href (descending)* &#x60;rpm_buildhost&#x60; - Rpm buildhost* &#x60;-rpm_buildhost&#x60; - Rpm buildhost (descending)* &#x60;rpm_group&#x60; - Rpm group* &#x60;-rpm_group&#x60; - Rpm group (descending)* &#x60;rpm_license&#x60; - Rpm license* &#x60;-rpm_license&#x60; - Rpm license (descending)* &#x60;rpm_packager&#x60; - Rpm packager* &#x60;-rpm_packager&#x60; - Rpm packager (descending)* &#x60;rpm_sourcerpm&#x60; - Rpm sourcerpm* &#x60;-rpm_sourcerpm&#x60; - Rpm sourcerpm (descending)* &#x60;rpm_vendor&#x60; - Rpm vendor* &#x60;-rpm_vendor&#x60; - Rpm vendor (descending)* &#x60;rpm_header_start&#x60; - Rpm header start* &#x60;-rpm_header_start&#x60; - Rpm header start (descending)* &#x60;rpm_header_end&#x60; - Rpm header end* &#x60;-rpm_header_end&#x60; - Rpm header end (descending)* &#x60;size_archive&#x60; - Size archive* &#x60;-size_archive&#x60; - Size archive (descending)* &#x60;size_installed&#x60; - Size installed* &#x60;-size_installed&#x60; - Size installed (descending)* &#x60;size_package&#x60; - Size package* &#x60;-size_package&#x60; - Size package (descending)* &#x60;time_build&#x60; - Time build* &#x60;-time_build&#x60; - Time build (descending)* &#x60;time_file&#x60; - Time file* &#x60;-time_file&#x60; - Time file (descending)* &#x60;is_modular&#x60; - Is modular* &#x60;-is_modular&#x60; - Is modular (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **pkgId** | **string** | Filter results where pkgId matches value | 
 **pkgIdIn** | **[]string** | Filter results where pkgId is in a comma-separated list of values | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **release** | **string** | Filter results where release matches value | 
 **releaseContains** | **string** | Filter results where release contains value | 
 **releaseIn** | **[]string** | Filter results where release is in a comma-separated list of values | 
 **releaseNe** | **string** | Filter results where release not equal to value | 
 **releaseStartswith** | **string** | Filter results where release starts with value | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF/PRN | 
 **sha256** | **string** |  | 
 **version** | **string** | Filter results where version matches value | 
 **versionIn** | **[]string** | Filter results where version is in a comma-separated list of values | 
 **versionNe** | **string** | Filter results where version not equal to value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedrpmPackageResponseList**](PaginatedrpmPackageResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

