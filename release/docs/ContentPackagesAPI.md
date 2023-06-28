# \ContentPackagesAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentRpmPackagesCreate**](ContentPackagesAPI.md#ContentRpmPackagesCreate) | **Post** /pulp/api/v3/content/rpm/packages/ | Create a package
[**ContentRpmPackagesList**](ContentPackagesAPI.md#ContentRpmPackagesList) | **Get** /pulp/api/v3/content/rpm/packages/ | List packages
[**ContentRpmPackagesRead**](ContentPackagesAPI.md#ContentRpmPackagesRead) | **Get** /{rpm_package_href} | Inspect a package



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
    openapiclient "github.com/content-services/zest/release/v2023"
)

func main() {
    artifact := "artifact_example" // string | Artifact file representing the physical content (optional)
    relativePath := "relativePath_example" // string | Path where the artifact is located relative to distributions base_path (optional)
    file := os.NewFile(1234, "some_file") // *os.File | An uploaded file that may be turned into the artifact of the content unit. (optional)
    repository := "repository_example" // string | A URI of a repository the new content unit should be associated with. (optional)
    upload := "upload_example" // string | An uncommitted upload that may be turned into the artifact of the content unit. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPackagesAPI.ContentRpmPackagesCreate(context.Background()).Artifact(artifact).RelativePath(relativePath).File(file).Repository(repository).Upload(upload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesAPI.ContentRpmPackagesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmPackagesCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentPackagesAPI.ContentRpmPackagesCreate`: %v\n", resp)
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

> PaginatedrpmPackageResponseList ContentRpmPackagesList(ctx).Arch(arch).ArchContains(archContains).ArchIn(archIn).ArchNe(archNe).ArchStartswith(archStartswith).ChecksumType(checksumType).ChecksumTypeIn(checksumTypeIn).ChecksumTypeNe(checksumTypeNe).Epoch(epoch).EpochIn(epochIn).EpochNe(epochNe).Limit(limit).Name(name).NameContains(nameContains).NameIn(nameIn).NameNe(nameNe).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PkgId(pkgId).PkgIdIn(pkgIdIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Release(release).ReleaseContains(releaseContains).ReleaseIn(releaseIn).ReleaseNe(releaseNe).ReleaseStartswith(releaseStartswith).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Version(version).VersionIn(versionIn).VersionNe(versionNe).Fields(fields).ExcludeFields(excludeFields).Execute()

List packages



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release/v2023"
)

func main() {
    arch := "arch_example" // string | Filter results where arch matches value (optional)
    archContains := "archContains_example" // string | Filter results where arch contains value (optional)
    archIn := []string{"Inner_example"} // []string | Filter results where arch is in a comma-separated list of values (optional)
    archNe := "archNe_example" // string | Filter results where arch not equal to value (optional)
    archStartswith := "archStartswith_example" // string | Filter results where arch starts with value (optional)
    checksumType := "checksumType_example" // string | Filter results where checksum_type matches value  * `unknown` - unknown * `md5` - md5 * `sha1` - sha1 * `sha1` - sha1 * `sha224` - sha224 * `sha256` - sha256 * `sha384` - sha384 * `sha512` - sha512 (optional)
    checksumTypeIn := []string{"Inner_example"} // []string | Filter results where checksum_type is in a comma-separated list of values (optional)
    checksumTypeNe := "checksumTypeNe_example" // string | Filter results where checksum_type not equal to value (optional)
    epoch := "epoch_example" // string | Filter results where epoch matches value (optional)
    epochIn := []string{"Inner_example"} // []string | Filter results where epoch is in a comma-separated list of values (optional)
    epochNe := "epochNe_example" // string | Filter results where epoch not equal to value (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameNe := "nameNe_example" // string | Filter results where name not equal to value (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering  * `pulp_id` - Pulp id * `-pulp_id` - Pulp id (descending) * `pulp_created` - Pulp created * `-pulp_created` - Pulp created (descending) * `pulp_last_updated` - Pulp last updated * `-pulp_last_updated` - Pulp last updated (descending) * `pulp_type` - Pulp type * `-pulp_type` - Pulp type (descending) * `upstream_id` - Upstream id * `-upstream_id` - Upstream id (descending) * `timestamp_of_interest` - Timestamp of interest * `-timestamp_of_interest` - Timestamp of interest (descending) * `name` - Name * `-name` - Name (descending) * `epoch` - Epoch * `-epoch` - Epoch (descending) * `version` - Version * `-version` - Version (descending) * `release` - Release * `-release` - Release (descending) * `arch` - Arch * `-arch` - Arch (descending) * `evr` - Evr * `-evr` - Evr (descending) * `pkgId` - Pkgid * `-pkgId` - Pkgid (descending) * `checksum_type` - Checksum type * `-checksum_type` - Checksum type (descending) * `summary` - Summary * `-summary` - Summary (descending) * `description` - Description * `-description` - Description (descending) * `url` - Url * `-url` - Url (descending) * `changelogs` - Changelogs * `-changelogs` - Changelogs (descending) * `files` - Files * `-files` - Files (descending) * `requires` - Requires * `-requires` - Requires (descending) * `provides` - Provides * `-provides` - Provides (descending) * `conflicts` - Conflicts * `-conflicts` - Conflicts (descending) * `obsoletes` - Obsoletes * `-obsoletes` - Obsoletes (descending) * `suggests` - Suggests * `-suggests` - Suggests (descending) * `enhances` - Enhances * `-enhances` - Enhances (descending) * `recommends` - Recommends * `-recommends` - Recommends (descending) * `supplements` - Supplements * `-supplements` - Supplements (descending) * `location_base` - Location base * `-location_base` - Location base (descending) * `location_href` - Location href * `-location_href` - Location href (descending) * `rpm_buildhost` - Rpm buildhost * `-rpm_buildhost` - Rpm buildhost (descending) * `rpm_group` - Rpm group * `-rpm_group` - Rpm group (descending) * `rpm_license` - Rpm license * `-rpm_license` - Rpm license (descending) * `rpm_packager` - Rpm packager * `-rpm_packager` - Rpm packager (descending) * `rpm_sourcerpm` - Rpm sourcerpm * `-rpm_sourcerpm` - Rpm sourcerpm (descending) * `rpm_vendor` - Rpm vendor * `-rpm_vendor` - Rpm vendor (descending) * `rpm_header_start` - Rpm header start * `-rpm_header_start` - Rpm header start (descending) * `rpm_header_end` - Rpm header end * `-rpm_header_end` - Rpm header end (descending) * `size_archive` - Size archive * `-size_archive` - Size archive (descending) * `size_installed` - Size installed * `-size_installed` - Size installed (descending) * `size_package` - Size package * `-size_package` - Size package (descending) * `time_build` - Time build * `-time_build` - Time build (descending) * `time_file` - Time file * `-time_file` - Time file (descending) * `is_modular` - Is modular * `-is_modular` - Is modular (descending) * `pk` - Pk * `-pk` - Pk (descending) (optional)
    pkgId := "pkgId_example" // string | Filter results where pkgId matches value (optional)
    pkgIdIn := []string{"Inner_example"} // []string | Filter results where pkgId is in a comma-separated list of values (optional)
    pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    release := "release_example" // string | Filter results where release matches value (optional)
    releaseContains := "releaseContains_example" // string | Filter results where release contains value (optional)
    releaseIn := []string{"Inner_example"} // []string | Filter results where release is in a comma-separated list of values (optional)
    releaseNe := "releaseNe_example" // string | Filter results where release not equal to value (optional)
    releaseStartswith := "releaseStartswith_example" // string | Filter results where release starts with value (optional)
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
    resp, r, err := apiClient.ContentPackagesAPI.ContentRpmPackagesList(context.Background()).Arch(arch).ArchContains(archContains).ArchIn(archIn).ArchNe(archNe).ArchStartswith(archStartswith).ChecksumType(checksumType).ChecksumTypeIn(checksumTypeIn).ChecksumTypeNe(checksumTypeNe).Epoch(epoch).EpochIn(epochIn).EpochNe(epochNe).Limit(limit).Name(name).NameContains(nameContains).NameIn(nameIn).NameNe(nameNe).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PkgId(pkgId).PkgIdIn(pkgIdIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).Release(release).ReleaseContains(releaseContains).ReleaseIn(releaseIn).ReleaseNe(releaseNe).ReleaseStartswith(releaseStartswith).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Sha256(sha256).Version(version).VersionIn(versionIn).VersionNe(versionNe).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesAPI.ContentRpmPackagesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmPackagesList`: PaginatedrpmPackageResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContentPackagesAPI.ContentRpmPackagesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContentRpmPackagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arch** | **string** | Filter results where arch matches value | 
 **archContains** | **string** | Filter results where arch contains value | 
 **archIn** | **[]string** | Filter results where arch is in a comma-separated list of values | 
 **archNe** | **string** | Filter results where arch not equal to value | 
 **archStartswith** | **string** | Filter results where arch starts with value | 
 **checksumType** | **string** | Filter results where checksum_type matches value  * &#x60;unknown&#x60; - unknown * &#x60;md5&#x60; - md5 * &#x60;sha1&#x60; - sha1 * &#x60;sha1&#x60; - sha1 * &#x60;sha224&#x60; - sha224 * &#x60;sha256&#x60; - sha256 * &#x60;sha384&#x60; - sha384 * &#x60;sha512&#x60; - sha512 | 
 **checksumTypeIn** | **[]string** | Filter results where checksum_type is in a comma-separated list of values | 
 **checksumTypeNe** | **string** | Filter results where checksum_type not equal to value | 
 **epoch** | **string** | Filter results where epoch matches value | 
 **epochIn** | **[]string** | Filter results where epoch is in a comma-separated list of values | 
 **epochNe** | **string** | Filter results where epoch not equal to value | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameNe** | **string** | Filter results where name not equal to value | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;epoch&#x60; - Epoch * &#x60;-epoch&#x60; - Epoch (descending) * &#x60;version&#x60; - Version * &#x60;-version&#x60; - Version (descending) * &#x60;release&#x60; - Release * &#x60;-release&#x60; - Release (descending) * &#x60;arch&#x60; - Arch * &#x60;-arch&#x60; - Arch (descending) * &#x60;evr&#x60; - Evr * &#x60;-evr&#x60; - Evr (descending) * &#x60;pkgId&#x60; - Pkgid * &#x60;-pkgId&#x60; - Pkgid (descending) * &#x60;checksum_type&#x60; - Checksum type * &#x60;-checksum_type&#x60; - Checksum type (descending) * &#x60;summary&#x60; - Summary * &#x60;-summary&#x60; - Summary (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;url&#x60; - Url * &#x60;-url&#x60; - Url (descending) * &#x60;changelogs&#x60; - Changelogs * &#x60;-changelogs&#x60; - Changelogs (descending) * &#x60;files&#x60; - Files * &#x60;-files&#x60; - Files (descending) * &#x60;requires&#x60; - Requires * &#x60;-requires&#x60; - Requires (descending) * &#x60;provides&#x60; - Provides * &#x60;-provides&#x60; - Provides (descending) * &#x60;conflicts&#x60; - Conflicts * &#x60;-conflicts&#x60; - Conflicts (descending) * &#x60;obsoletes&#x60; - Obsoletes * &#x60;-obsoletes&#x60; - Obsoletes (descending) * &#x60;suggests&#x60; - Suggests * &#x60;-suggests&#x60; - Suggests (descending) * &#x60;enhances&#x60; - Enhances * &#x60;-enhances&#x60; - Enhances (descending) * &#x60;recommends&#x60; - Recommends * &#x60;-recommends&#x60; - Recommends (descending) * &#x60;supplements&#x60; - Supplements * &#x60;-supplements&#x60; - Supplements (descending) * &#x60;location_base&#x60; - Location base * &#x60;-location_base&#x60; - Location base (descending) * &#x60;location_href&#x60; - Location href * &#x60;-location_href&#x60; - Location href (descending) * &#x60;rpm_buildhost&#x60; - Rpm buildhost * &#x60;-rpm_buildhost&#x60; - Rpm buildhost (descending) * &#x60;rpm_group&#x60; - Rpm group * &#x60;-rpm_group&#x60; - Rpm group (descending) * &#x60;rpm_license&#x60; - Rpm license * &#x60;-rpm_license&#x60; - Rpm license (descending) * &#x60;rpm_packager&#x60; - Rpm packager * &#x60;-rpm_packager&#x60; - Rpm packager (descending) * &#x60;rpm_sourcerpm&#x60; - Rpm sourcerpm * &#x60;-rpm_sourcerpm&#x60; - Rpm sourcerpm (descending) * &#x60;rpm_vendor&#x60; - Rpm vendor * &#x60;-rpm_vendor&#x60; - Rpm vendor (descending) * &#x60;rpm_header_start&#x60; - Rpm header start * &#x60;-rpm_header_start&#x60; - Rpm header start (descending) * &#x60;rpm_header_end&#x60; - Rpm header end * &#x60;-rpm_header_end&#x60; - Rpm header end (descending) * &#x60;size_archive&#x60; - Size archive * &#x60;-size_archive&#x60; - Size archive (descending) * &#x60;size_installed&#x60; - Size installed * &#x60;-size_installed&#x60; - Size installed (descending) * &#x60;size_package&#x60; - Size package * &#x60;-size_package&#x60; - Size package (descending) * &#x60;time_build&#x60; - Time build * &#x60;-time_build&#x60; - Time build (descending) * &#x60;time_file&#x60; - Time file * &#x60;-time_file&#x60; - Time file (descending) * &#x60;is_modular&#x60; - Is modular * &#x60;-is_modular&#x60; - Is modular (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending) | 
 **pkgId** | **string** | Filter results where pkgId matches value | 
 **pkgIdIn** | **[]string** | Filter results where pkgId is in a comma-separated list of values | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **release** | **string** | Filter results where release matches value | 
 **releaseContains** | **string** | Filter results where release contains value | 
 **releaseIn** | **[]string** | Filter results where release is in a comma-separated list of values | 
 **releaseNe** | **string** | Filter results where release not equal to value | 
 **releaseStartswith** | **string** | Filter results where release starts with value | 
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
    openapiclient "github.com/content-services/zest/release/v2023"
)

func main() {
    rpmPackageHref := "rpmPackageHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentPackagesAPI.ContentRpmPackagesRead(context.Background(), rpmPackageHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPackagesAPI.ContentRpmPackagesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContentRpmPackagesRead`: RpmPackageResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentPackagesAPI.ContentRpmPackagesRead`: %v\n", resp)
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

