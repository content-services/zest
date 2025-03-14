# \ContentManifestsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentContainerManifestsList**](ContentManifestsAPI.md#ContentContainerManifestsList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/container/manifests/ | List manifests
[**ContentContainerManifestsRead**](ContentManifestsAPI.md#ContentContainerManifestsRead) | **Get** /{container_manifest_href} | Inspect a manifest



## ContentContainerManifestsList

> PaginatedcontainerManifestResponseList ContentContainerManifestsList(ctx, pulpDomain).Digest(digest).DigestIn(digestIn).IsBootable(isBootable).IsFlatpak(isFlatpak).Limit(limit).MediaType(mediaType).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List manifests



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
	digest := "digest_example" // string | Filter results where digest matches value (optional)
	digestIn := []string{"Inner_example"} // []string | Filter results where digest is in a comma-separated list of values (optional)
	isBootable := true // bool | Filter results where is_bootable matches value (optional)
	isFlatpak := true // bool | Filter results where is_flatpak matches value (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	mediaType := []string{"MediaType_example"} // []string | * `application/vnd.docker.distribution.manifest.v1+json` - application/vnd.docker.distribution.manifest.v1+json* `application/vnd.docker.distribution.manifest.v2+json` - application/vnd.docker.distribution.manifest.v2+json* `application/vnd.docker.distribution.manifest.list.v2+json` - application/vnd.docker.distribution.manifest.list.v2+json* `application/vnd.oci.image.manifest.v1+json` - application/vnd.oci.image.manifest.v1+json* `application/vnd.oci.image.index.v1+json` - application/vnd.oci.image.index.v1+json (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pulp_id` - Pulp id* `-pulp_id` - Pulp id (descending)* `pulp_created` - Pulp created* `-pulp_created` - Pulp created (descending)* `pulp_last_updated` - Pulp last updated* `-pulp_last_updated` - Pulp last updated (descending)* `pulp_type` - Pulp type* `-pulp_type` - Pulp type (descending)* `upstream_id` - Upstream id* `-upstream_id` - Upstream id (descending)* `pulp_labels` - Pulp labels* `-pulp_labels` - Pulp labels (descending)* `timestamp_of_interest` - Timestamp of interest* `-timestamp_of_interest` - Timestamp of interest (descending)* `digest` - Digest* `-digest` - Digest (descending)* `schema_version` - Schema version* `-schema_version` - Schema version (descending)* `media_type` - Media type* `-media_type` - Media type (descending)* `type` - Type* `-type` - Type (descending)* `data` - Data* `-data` - Data (descending)* `annotations` - Annotations* `-annotations` - Annotations (descending)* `labels` - Labels* `-labels` - Labels (descending)* `architecture` - Architecture* `-architecture` - Architecture (descending)* `os` - Os* `-os` - Os (descending)* `compressed_image_size` - Compressed image size* `-compressed_image_size` - Compressed image size (descending)* `is_bootable` - Is bootable* `-is_bootable` - Is bootable (descending)* `is_flatpak` - Is flatpak* `-is_flatpak` - Is flatpak (descending)* `pk` - Pk* `-pk` - Pk (descending) (optional)
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
	resp, r, err := apiClient.ContentManifestsAPI.ContentContainerManifestsList(context.Background(), pulpDomain).Digest(digest).DigestIn(digestIn).IsBootable(isBootable).IsFlatpak(isFlatpak).Limit(limit).MediaType(mediaType).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentManifestsAPI.ContentContainerManifestsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentContainerManifestsList`: PaginatedcontainerManifestResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentManifestsAPI.ContentContainerManifestsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentContainerManifestsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **digest** | **string** | Filter results where digest matches value | 
 **digestIn** | **[]string** | Filter results where digest is in a comma-separated list of values | 
 **isBootable** | **bool** | Filter results where is_bootable matches value | 
 **isFlatpak** | **bool** | Filter results where is_flatpak matches value | 
 **limit** | **int32** | Number of results to return per page. | 
 **mediaType** | **[]string** | * &#x60;application/vnd.docker.distribution.manifest.v1+json&#x60; - application/vnd.docker.distribution.manifest.v1+json* &#x60;application/vnd.docker.distribution.manifest.v2+json&#x60; - application/vnd.docker.distribution.manifest.v2+json* &#x60;application/vnd.docker.distribution.manifest.list.v2+json&#x60; - application/vnd.docker.distribution.manifest.list.v2+json* &#x60;application/vnd.oci.image.manifest.v1+json&#x60; - application/vnd.oci.image.manifest.v1+json* &#x60;application/vnd.oci.image.index.v1+json&#x60; - application/vnd.oci.image.index.v1+json | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pulp_id&#x60; - Pulp id* &#x60;-pulp_id&#x60; - Pulp id (descending)* &#x60;pulp_created&#x60; - Pulp created* &#x60;-pulp_created&#x60; - Pulp created (descending)* &#x60;pulp_last_updated&#x60; - Pulp last updated* &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending)* &#x60;pulp_type&#x60; - Pulp type* &#x60;-pulp_type&#x60; - Pulp type (descending)* &#x60;upstream_id&#x60; - Upstream id* &#x60;-upstream_id&#x60; - Upstream id (descending)* &#x60;pulp_labels&#x60; - Pulp labels* &#x60;-pulp_labels&#x60; - Pulp labels (descending)* &#x60;timestamp_of_interest&#x60; - Timestamp of interest* &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending)* &#x60;digest&#x60; - Digest* &#x60;-digest&#x60; - Digest (descending)* &#x60;schema_version&#x60; - Schema version* &#x60;-schema_version&#x60; - Schema version (descending)* &#x60;media_type&#x60; - Media type* &#x60;-media_type&#x60; - Media type (descending)* &#x60;type&#x60; - Type* &#x60;-type&#x60; - Type (descending)* &#x60;data&#x60; - Data* &#x60;-data&#x60; - Data (descending)* &#x60;annotations&#x60; - Annotations* &#x60;-annotations&#x60; - Annotations (descending)* &#x60;labels&#x60; - Labels* &#x60;-labels&#x60; - Labels (descending)* &#x60;architecture&#x60; - Architecture* &#x60;-architecture&#x60; - Architecture (descending)* &#x60;os&#x60; - Os* &#x60;-os&#x60; - Os (descending)* &#x60;compressed_image_size&#x60; - Compressed image size* &#x60;-compressed_image_size&#x60; - Compressed image size (descending)* &#x60;is_bootable&#x60; - Is bootable* &#x60;-is_bootable&#x60; - Is bootable (descending)* &#x60;is_flatpak&#x60; - Is flatpak* &#x60;-is_flatpak&#x60; - Is flatpak (descending)* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
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

[**PaginatedcontainerManifestResponseList**](PaginatedcontainerManifestResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentContainerManifestsRead

> ContainerManifestResponse ContentContainerManifestsRead(ctx, containerManifestHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect a manifest



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
	containerManifestHref := "containerManifestHref_example" // string | 
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentManifestsAPI.ContentContainerManifestsRead(context.Background(), containerManifestHref).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentManifestsAPI.ContentContainerManifestsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentContainerManifestsRead`: ContainerManifestResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentManifestsAPI.ContentContainerManifestsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerManifestHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentContainerManifestsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**ContainerManifestResponse**](ContainerManifestResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

