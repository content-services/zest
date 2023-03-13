# \ContentPackagesApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentDebPackagesCreate**](ContentPackagesApi.md#ContentDebPackagesCreate) | **Post** /pulp/api/v3/content/deb/packages/ | Create a package
[**ContentDebPackagesList**](ContentPackagesApi.md#ContentDebPackagesList) | **Get** /pulp/api/v3/content/deb/packages/ | List packages
[**ContentDebPackagesRead**](ContentPackagesApi.md#ContentDebPackagesRead) | **Get** /{deb_package_href} | Inspect a package
[**ContentPythonPackagesCreate**](ContentPackagesApi.md#ContentPythonPackagesCreate) | **Post** /pulp/api/v3/content/python/packages/ | Create a python package content
[**ContentPythonPackagesList**](ContentPackagesApi.md#ContentPythonPackagesList) | **Get** /pulp/api/v3/content/python/packages/ | List python package contents
[**ContentPythonPackagesRead**](ContentPackagesApi.md#ContentPythonPackagesRead) | **Get** /{python_python_package_content_href} | Inspect a python package content
[**ContentRpmPackagesCreate**](ContentPackagesApi.md#ContentRpmPackagesCreate) | **Post** /pulp/api/v3/content/rpm/packages/ | Create a package
[**ContentRpmPackagesList**](ContentPackagesApi.md#ContentRpmPackagesList) | **Get** /pulp/api/v3/content/rpm/packages/ | List packages
[**ContentRpmPackagesRead**](ContentPackagesApi.md#ContentRpmPackagesRead) | **Get** /{rpm_package_href} | Inspect a package



## ContentDebPackagesCreate

> AsyncOperationResponse ContentDebPackagesCreate(ctx).Artifact(artifact).RelativePath(relativePath).File(file).Repository(repository).Upload(upload).Execute()

Create a package



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
    artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
    relativePath := "relativePath_example" // string | Path where the artifact is located relative to distributions base_path (optional)
    file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the artifact of the content unit. (optional)
    repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
    upload := "upload_example" // string | An uncommitted upload that may be turned into the artifact of the content unit. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPackagesApi.ContentDebPackagesCreate(context.Background()).Artifact(artifact).RelativePath(relativePath).File(file).Repository(repository).Upload(upload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesApi.ContentDebPackagesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebPackagesCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentPackagesApi.ContentDebPackagesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebPackagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifact** | **string** | Artifact file representing the physical content | 
 **relativePath** | **string** | Path where the artifact is located relative to distributions base_path | 
 **file** | ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | 
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **upload** | **string** | An uncommitted upload that may be turned into the artifact of the content unit. | 

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


## ContentDebPackagesList

> PaginateddebBasePackageResponseList ContentDebPackagesList(ctx).Architecture(architecture).AutoBuiltPackage(autoBuiltPackage).BuildEssential(buildEssential).BuiltUsing(builtUsing).Essential(essential).InstalledSize(installedSize).Limit(limit).Maintainer(maintainer).MultiArch(multiArch).Offset(offset).Ordering(ordering).Origin(origin).OriginalMaintainer(originalMaintainer).Package_(package_).Priority(priority).RelativePath(relativePath).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Section(section).Sha256(sha256).Source(source).Tag(tag).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()

List packages



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
    architecture := "architecture_example" // string | Filter results where architecture matches value (optional)
    autoBuiltPackage := "autoBuiltPackage_example" // string | Filter results where auto_built_package matches value (optional)
    buildEssential := true // bool | Filter results where build_essential matches value (optional)
    builtUsing := "builtUsing_example" // string | Filter results where built_using matches value (optional)
    essential := true // bool | Filter results where essential matches value (optional)
    installedSize := int32(56) // int32 | Filter results where installed_size matches value (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    maintainer := "maintainer_example" // string | Filter results where maintainer matches value (optional)
    multiArch := "multiArch_example" // string | Filter results where multi_arch matches value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    origin := "origin_example" // string | Filter results where origin matches value (optional)
    originalMaintainer := "originalMaintainer_example" // string | Filter results where original_maintainer matches value (optional)
    package_ := "package__example" // string | Filter results where package matches value (optional)
    priority := "priority_example" // string | Filter results where priority matches value (optional)
    relativePath := "relativePath_example" // string | Filter results where relative_path matches value (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    section := "section_example" // string | Filter results where section matches value (optional)
    sha256 := "sha256_example" // string | Filter results where sha256 matches value (optional)
    source := "source_example" // string | Filter results where source matches value (optional)
    tag := "tag_example" // string | Filter results where tag matches value (optional)
    version := "version_example" // string | Filter results where version matches value (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPackagesApi.ContentDebPackagesList(context.Background()).Architecture(architecture).AutoBuiltPackage(autoBuiltPackage).BuildEssential(buildEssential).BuiltUsing(builtUsing).Essential(essential).InstalledSize(installedSize).Limit(limit).Maintainer(maintainer).MultiArch(multiArch).Offset(offset).Ordering(ordering).Origin(origin).OriginalMaintainer(originalMaintainer).Package_(package_).Priority(priority).RelativePath(relativePath).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Section(section).Sha256(sha256).Source(source).Tag(tag).Version(version).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesApi.ContentDebPackagesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebPackagesList`: PaginateddebBasePackageResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentPackagesApi.ContentDebPackagesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentDebPackagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **architecture** | **string** | Filter results where architecture matches value | 
 **autoBuiltPackage** | **string** | Filter results where auto_built_package matches value | 
 **buildEssential** | **bool** | Filter results where build_essential matches value | 
 **builtUsing** | **string** | Filter results where built_using matches value | 
 **essential** | **bool** | Filter results where essential matches value | 
 **installedSize** | **int32** | Filter results where installed_size matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **maintainer** | **string** | Filter results where maintainer matches value | 
 **multiArch** | **string** | Filter results where multi_arch matches value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **origin** | **string** | Filter results where origin matches value | 
 **originalMaintainer** | **string** | Filter results where original_maintainer matches value | 
 **package_** | **string** | Filter results where package matches value | 
 **priority** | **string** | Filter results where priority matches value | 
 **relativePath** | **string** | Filter results where relative_path matches value | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **section** | **string** | Filter results where section matches value | 
 **sha256** | **string** | Filter results where sha256 matches value | 
 **source** | **string** | Filter results where source matches value | 
 **tag** | **string** | Filter results where tag matches value | 
 **version** | **string** | Filter results where version matches value | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginateddebBasePackageResponseList**](PaginateddebBasePackageResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentDebPackagesRead

> DebBasePackageResponse ContentDebPackagesRead(ctx, debPackageHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a package



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
    debPackageHref := "debPackageHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPackagesApi.ContentDebPackagesRead(context.Background(), debPackageHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesApi.ContentDebPackagesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentDebPackagesRead`: DebBasePackageResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentPackagesApi.ContentDebPackagesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debPackageHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentDebPackagesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**DebBasePackageResponse**](DebBasePackageResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentPythonPackagesCreate

> AsyncOperationResponse ContentPythonPackagesCreate(ctx).RelativePath(relativePath).Artifact(artifact).File(file).Repository(repository).Upload(upload).Sha256(sha256).Summary(summary).Description(description).DescriptionContentType(descriptionContentType).Keywords(keywords).HomePage(homePage).DownloadUrl(downloadUrl).Author(author).AuthorEmail(authorEmail).Maintainer(maintainer).MaintainerEmail(maintainerEmail).License(license).RequiresPython(requiresPython).ProjectUrl(projectUrl).ProjectUrls(projectUrls).Platform(platform).SupportedPlatform(supportedPlatform).RequiresDist(requiresDist).ProvidesDist(providesDist).ObsoletesDist(obsoletesDist).RequiresExternal(requiresExternal).Classifiers(classifiers).Execute()

Create a python package content



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
    relativePath := "relativePath_example" // string | Path where the artifact is located relative to distributions base_path
    artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
    file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the artifact of the content unit. (optional)
    repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
    upload := "upload_example" // string | An uncommitted upload that may be turned into the artifact of the content unit. (optional)
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
    projectUrls := map[string]interface{}{ ... } // map[string]interface{} | A dictionary of labels and URLs for the project. (optional)
    platform := "platform_example" // string | A comma-separated list of platform specifications, summarizing the operating systems supported by the package. (optional)
    supportedPlatform := "supportedPlatform_example" // string | Field to specify the OS and CPU for which the binary package was compiled.  (optional)
    requiresDist := map[string]interface{}{ ... } // map[string]interface{} | A JSON list containing names of some other distutils project required by this distribution. (optional)
    providesDist := map[string]interface{}{ ... } // map[string]interface{} | A JSON list containing names of a Distutils project which is contained within this distribution. (optional)
    obsoletesDist := map[string]interface{}{ ... } // map[string]interface{} | A JSON list containing names of a distutils project's distribution which this distribution renders obsolete, meaning that the two projects should not be installed at the same time. (optional)
    requiresExternal := map[string]interface{}{ ... } // map[string]interface{} | A JSON list containing some dependency in the system that the distribution is to be used. (optional)
    classifiers := map[string]interface{}{ ... } // map[string]interface{} | A JSON list containing classification values for a Python package. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPackagesApi.ContentPythonPackagesCreate(context.Background()).RelativePath(relativePath).Artifact(artifact).File(file).Repository(repository).Upload(upload).Sha256(sha256).Summary(summary).Description(description).DescriptionContentType(descriptionContentType).Keywords(keywords).HomePage(homePage).DownloadUrl(downloadUrl).Author(author).AuthorEmail(authorEmail).Maintainer(maintainer).MaintainerEmail(maintainerEmail).License(license).RequiresPython(requiresPython).ProjectUrl(projectUrl).ProjectUrls(projectUrls).Platform(platform).SupportedPlatform(supportedPlatform).RequiresDist(requiresDist).ProvidesDist(providesDist).ObsoletesDist(obsoletesDist).RequiresExternal(requiresExternal).Classifiers(classifiers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesApi.ContentPythonPackagesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentPythonPackagesCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentPackagesApi.ContentPythonPackagesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentPythonPackagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **relativePath** | **string** | Path where the artifact is located relative to distributions base_path | 
 **artifact** | **string** | Artifact file representing the physical content | 
 **file** | ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | 
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **upload** | **string** | An uncommitted upload that may be turned into the artifact of the content unit. | 
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
 **projectUrls** | [**map[string]interface{}**](map[string]interface{}.md) | A dictionary of labels and URLs for the project. | 
 **platform** | **string** | A comma-separated list of platform specifications, summarizing the operating systems supported by the package. | 
 **supportedPlatform** | **string** | Field to specify the OS and CPU for which the binary package was compiled.  | 
 **requiresDist** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON list containing names of some other distutils project required by this distribution. | 
 **providesDist** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON list containing names of a Distutils project which is contained within this distribution. | 
 **obsoletesDist** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON list containing names of a distutils project&#39;s distribution which this distribution renders obsolete, meaning that the two projects should not be installed at the same time. | 
 **requiresExternal** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON list containing some dependency in the system that the distribution is to be used. | 
 **classifiers** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON list containing classification values for a Python package. | 

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

> PaginatedpythonPythonPackageContentResponseList ContentPythonPackagesList(ctx).Author(author).AuthorIn(authorIn).Filename(filename).FilenameContains(filenameContains).FilenameIn(filenameIn).KeywordsContains(keywordsContains).KeywordsIn(keywordsIn).Limit(limit).Name(name).NameIn(nameIn).Offset(offset).Ordering(ordering).Packagetype(packagetype).PackagetypeIn(packagetypeIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).RequiresPython(requiresPython).RequiresPythonContains(requiresPythonContains).RequiresPythonIn(requiresPythonIn).Sha256(sha256).Sha256In(sha256In).Fields(fields).ExcludeFields(excludeFields).Execute()

List python package contents



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
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    packagetype := "packagetype_example" // string | Filter results where packagetype matches value (optional)
    packagetypeIn := []string{"Inner_example"} // []string | Filter results where packagetype is in a comma-separated list of values (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    requiresPython := "requiresPython_example" // string | Filter results where requires_python matches value (optional)
    requiresPythonContains := "requiresPythonContains_example" // string | Filter results where requires_python contains value (optional)
    requiresPythonIn := []string{"Inner_example"} // []string | Filter results where requires_python is in a comma-separated list of values (optional)
    sha256 := "sha256_example" // string | Filter results where sha256 matches value (optional)
    sha256In := []string{"Inner_example"} // []string | Filter results where sha256 is in a comma-separated list of values (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPackagesApi.ContentPythonPackagesList(context.Background()).Author(author).AuthorIn(authorIn).Filename(filename).FilenameContains(filenameContains).FilenameIn(filenameIn).KeywordsContains(keywordsContains).KeywordsIn(keywordsIn).Limit(limit).Name(name).NameIn(nameIn).Offset(offset).Ordering(ordering).Packagetype(packagetype).PackagetypeIn(packagetypeIn).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).RequiresPython(requiresPython).RequiresPythonContains(requiresPythonContains).RequiresPythonIn(requiresPythonIn).Sha256(sha256).Sha256In(sha256In).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesApi.ContentPythonPackagesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentPythonPackagesList`: PaginatedpythonPythonPackageContentResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentPackagesApi.ContentPythonPackagesList`: %v\n", resp)
}
```

### Path Parameters



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
 **ordering** | **[]string** | Ordering | 
 **packagetype** | **string** | Filter results where packagetype matches value | 
 **packagetypeIn** | **[]string** | Filter results where packagetype is in a comma-separated list of values | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
 **requiresPython** | **string** | Filter results where requires_python matches value | 
 **requiresPythonContains** | **string** | Filter results where requires_python contains value | 
 **requiresPythonIn** | **[]string** | Filter results where requires_python is in a comma-separated list of values | 
 **sha256** | **string** | Filter results where sha256 matches value | 
 **sha256In** | **[]string** | Filter results where sha256 is in a comma-separated list of values | 
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
    openapiclient "github.com/content-services/zest/release"
)

func main() {
    pythonPythonPackageContentHref := "pythonPythonPackageContentHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPackagesApi.ContentPythonPackagesRead(context.Background(), pythonPythonPackageContentHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesApi.ContentPythonPackagesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentPythonPackagesRead`: PythonPythonPackageContentResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentPackagesApi.ContentPythonPackagesRead`: %v\n", resp)
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


## ContentRpmPackagesCreate

> AsyncOperationResponse ContentRpmPackagesCreate(ctx).Artifact(artifact).RelativePath(relativePath).File(file).Repository(repository).Upload(upload).Execute()

Create a package



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
    artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
    relativePath := "relativePath_example" // string | Path where the artifact is located relative to distributions base_path (optional)
    file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the artifact of the content unit. (optional)
    repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
    upload := "upload_example" // string | An uncommitted upload that may be turned into the artifact of the content unit. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPackagesApi.ContentRpmPackagesCreate(context.Background()).Artifact(artifact).RelativePath(relativePath).File(file).Repository(repository).Upload(upload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesApi.ContentRpmPackagesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmPackagesCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentPackagesApi.ContentRpmPackagesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmPackagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifact** | **string** | Artifact file representing the physical content | 
 **relativePath** | **string** | Path where the artifact is located relative to distributions base_path | 
 **file** | ***os.File** | An uploaded file that may be turned into the artifact of the content unit. | 
 **repository** | **string** | A URI of a repository the new content unit should be associated with. | 
 **upload** | **string** | An uncommitted upload that may be turned into the artifact of the content unit. | 

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

> PaginatedrpmPackageResponseList ContentRpmPackagesList(ctx).Arch(arch).ArchIn(archIn).ArchNe(archNe).ChecksumType(checksumType).ChecksumTypeIn(checksumTypeIn).ChecksumTypeNe(checksumTypeNe).Epoch(epoch).EpochIn(epochIn).EpochNe(epochNe).Limit(limit).Name(name).NameIn(nameIn).NameNe(nameNe).Offset(offset).Ordering(ordering).PkgId(pkgId).PkgIdIn(pkgIdIn).Release(release).ReleaseIn(releaseIn).ReleaseNe(releaseNe).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Version(version).VersionIn(versionIn).VersionNe(versionNe).Fields(fields).ExcludeFields(excludeFields).Execute()

List packages



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
    arch := "arch_example" // string | Filter results where arch matches value (optional)
    archIn := []string{"Inner_example"} // []string | Filter results where arch is in a comma-separated list of values (optional)
    archNe := "archNe_example" // string | Filter results where arch not equal to value (optional)
    checksumType := "checksumType_example" // string | Filter results where checksum_type matches value (optional)
    checksumTypeIn := []string{"Inner_example"} // []string | Filter results where checksum_type is in a comma-separated list of values (optional)
    checksumTypeNe := "checksumTypeNe_example" // string | Filter results where checksum_type not equal to value (optional)
    epoch := "epoch_example" // string | Filter results where epoch matches value (optional)
    epochIn := []string{"Inner_example"} // []string | Filter results where epoch is in a comma-separated list of values (optional)
    epochNe := "epochNe_example" // string | Filter results where epoch not equal to value (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameNe := "nameNe_example" // string | Filter results where name not equal to value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    pkgId := "pkgId_example" // string | Filter results where pkgId matches value (optional)
    pkgIdIn := []string{"Inner_example"} // []string | Filter results where pkgId is in a comma-separated list of values (optional)
    release := "release_example" // string | Filter results where release matches value (optional)
    releaseIn := []string{"Inner_example"} // []string | Filter results where release is in a comma-separated list of values (optional)
    releaseNe := "releaseNe_example" // string | Filter results where release not equal to value (optional)
    repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF (optional)
    repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF (optional)
    sha256 := "sha256_example" // string |  (optional)
    version := "version_example" // string | Filter results where version matches value (optional)
    versionIn := []string{"Inner_example"} // []string | Filter results where version is in a comma-separated list of values (optional)
    versionNe := "versionNe_example" // string | Filter results where version not equal to value (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPackagesApi.ContentRpmPackagesList(context.Background()).Arch(arch).ArchIn(archIn).ArchNe(archNe).ChecksumType(checksumType).ChecksumTypeIn(checksumTypeIn).ChecksumTypeNe(checksumTypeNe).Epoch(epoch).EpochIn(epochIn).EpochNe(epochNe).Limit(limit).Name(name).NameIn(nameIn).NameNe(nameNe).Offset(offset).Ordering(ordering).PkgId(pkgId).PkgIdIn(pkgIdIn).Release(release).ReleaseIn(releaseIn).ReleaseNe(releaseNe).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Version(version).VersionIn(versionIn).VersionNe(versionNe).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesApi.ContentRpmPackagesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmPackagesList`: PaginatedrpmPackageResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentPackagesApi.ContentRpmPackagesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmPackagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arch** | **string** | Filter results where arch matches value | 
 **archIn** | **[]string** | Filter results where arch is in a comma-separated list of values | 
 **archNe** | **string** | Filter results where arch not equal to value | 
 **checksumType** | **string** | Filter results where checksum_type matches value | 
 **checksumTypeIn** | **[]string** | Filter results where checksum_type is in a comma-separated list of values | 
 **checksumTypeNe** | **string** | Filter results where checksum_type not equal to value | 
 **epoch** | **string** | Filter results where epoch matches value | 
 **epochIn** | **[]string** | Filter results where epoch is in a comma-separated list of values | 
 **epochNe** | **string** | Filter results where epoch not equal to value | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameNe** | **string** | Filter results where name not equal to value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **pkgId** | **string** | Filter results where pkgId matches value | 
 **pkgIdIn** | **[]string** | Filter results where pkgId is in a comma-separated list of values | 
 **release** | **string** | Filter results where release matches value | 
 **releaseIn** | **[]string** | Filter results where release is in a comma-separated list of values | 
 **releaseNe** | **string** | Filter results where release not equal to value | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF | 
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


## ContentRpmPackagesRead

> RpmPackageResponse ContentRpmPackagesRead(ctx, rpmPackageHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a package



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
    rpmPackageHref := "rpmPackageHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPackagesApi.ContentRpmPackagesRead(context.Background(), rpmPackageHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesApi.ContentRpmPackagesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmPackagesRead`: RpmPackageResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentPackagesApi.ContentRpmPackagesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpmPackageHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmPackagesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**RpmPackageResponse**](RpmPackageResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

