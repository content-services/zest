# \ContentAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentList**](ContentAPI.md#ContentList) | **Get** /api/pulp/{pulp_domain}/api/v3/content/ | List content



## ContentList

> PaginatedMultipleArtifactContentResponseList ContentList(ctx, pulpDomain).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpType(pulpType).PulpTypeIn(pulpTypeIn).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()

List content



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
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := []string{"Ordering_example"} // []string | Ordering* `pk` - Pk* `-pk` - Pk (descending) (optional)
	orphanedFor := float32(8.14) // float32 | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. (optional)
	prnIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpHrefIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpIdIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
	pulpType := "pulpType_example" // string | Pulp type* `core.publishedmetadata` - core.publishedmetadata* `core.openpgp_publickey` - core.openpgp_publickey* `core.openpgp_publicsubkey` - core.openpgp_publicsubkey* `core.openpgp_userid` - core.openpgp_userid* `core.openpgp_userattribute` - core.openpgp_userattribute* `core.openpgp_signature` - core.openpgp_signature* `python.python` - python.python* `gem.gem` - gem.gem* `npm.package` - npm.package* `ostree.object` - ostree.object* `ostree.commit` - ostree.commit* `ostree.refs` - ostree.refs* `ostree.content` - ostree.content* `ostree.config` - ostree.config* `ostree.summary` - ostree.summary* `container.blob` - container.blob* `container.manifest` - container.manifest* `container.tag` - container.tag* `container.signature` - container.signature* `file.file` - file.file* `maven.artifact` - maven.artifact* `maven.metadata` - maven.metadata* `rpm.advisory` - rpm.advisory* `rpm.packagegroup` - rpm.packagegroup* `rpm.packagecategory` - rpm.packagecategory* `rpm.packageenvironment` - rpm.packageenvironment* `rpm.packagelangpacks` - rpm.packagelangpacks* `rpm.repo_metadata_file` - rpm.repo_metadata_file* `rpm.distribution_tree` - rpm.distribution_tree* `rpm.package` - rpm.package* `rpm.modulemd` - rpm.modulemd* `rpm.modulemd_defaults` - rpm.modulemd_defaults* `rpm.modulemd_obsolete` - rpm.modulemd_obsolete (optional)
	pulpTypeIn := []string{"PulpTypeIn_example"} // []string | Multiple values may be separated by commas.* `core.publishedmetadata` - core.publishedmetadata* `core.openpgp_publickey` - core.openpgp_publickey* `core.openpgp_publicsubkey` - core.openpgp_publicsubkey* `core.openpgp_userid` - core.openpgp_userid* `core.openpgp_userattribute` - core.openpgp_userattribute* `core.openpgp_signature` - core.openpgp_signature* `python.python` - python.python* `gem.gem` - gem.gem* `npm.package` - npm.package* `ostree.object` - ostree.object* `ostree.commit` - ostree.commit* `ostree.refs` - ostree.refs* `ostree.content` - ostree.content* `ostree.config` - ostree.config* `ostree.summary` - ostree.summary* `container.blob` - container.blob* `container.manifest` - container.manifest* `container.tag` - container.tag* `container.signature` - container.signature* `file.file` - file.file* `maven.artifact` - maven.artifact* `maven.metadata` - maven.metadata* `rpm.advisory` - rpm.advisory* `rpm.packagegroup` - rpm.packagegroup* `rpm.packagecategory` - rpm.packagecategory* `rpm.packageenvironment` - rpm.packageenvironment* `rpm.packagelangpacks` - rpm.packagelangpacks* `rpm.repo_metadata_file` - rpm.repo_metadata_file* `rpm.distribution_tree` - rpm.distribution_tree* `rpm.package` - rpm.package* `rpm.modulemd` - rpm.modulemd* `rpm.modulemd_defaults` - rpm.modulemd_defaults* `rpm.modulemd_obsolete` - rpm.modulemd_obsolete (optional)
	q := "q_example" // string | Filter results by using NOT, AND and OR operations on other filters (optional)
	repositoryVersion := "repositoryVersion_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionAdded := "repositoryVersionAdded_example" // string | Repository Version referenced by HREF/PRN (optional)
	repositoryVersionRemoved := "repositoryVersionRemoved_example" // string | Repository Version referenced by HREF/PRN (optional)
	fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
	excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.ContentList(context.Background(), pulpDomain).Limit(limit).Offset(offset).Ordering(ordering).OrphanedFor(orphanedFor).PrnIn(prnIn).PulpHrefIn(pulpHrefIn).PulpIdIn(pulpIdIn).PulpLabelSelect(pulpLabelSelect).PulpType(pulpType).PulpTypeIn(pulpTypeIn).Q(q).RepositoryVersion(repositoryVersion).RepositoryVersionAdded(repositoryVersionAdded).RepositoryVersionRemoved(repositoryVersionRemoved).Fields(fields).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.ContentList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContentList`: PaginatedMultipleArtifactContentResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.ContentList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pulpDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering* &#x60;pk&#x60; - Pk* &#x60;-pk&#x60; - Pk (descending) | 
 **orphanedFor** | **float32** | Minutes Content has been orphaned for. -1 uses ORPHAN_PROTECTION_TIME. | 
 **prnIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpHrefIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpIdIn** | **[]string** | Multiple values may be separated by commas. | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **pulpType** | **string** | Pulp type* &#x60;core.publishedmetadata&#x60; - core.publishedmetadata* &#x60;core.openpgp_publickey&#x60; - core.openpgp_publickey* &#x60;core.openpgp_publicsubkey&#x60; - core.openpgp_publicsubkey* &#x60;core.openpgp_userid&#x60; - core.openpgp_userid* &#x60;core.openpgp_userattribute&#x60; - core.openpgp_userattribute* &#x60;core.openpgp_signature&#x60; - core.openpgp_signature* &#x60;python.python&#x60; - python.python* &#x60;gem.gem&#x60; - gem.gem* &#x60;npm.package&#x60; - npm.package* &#x60;ostree.object&#x60; - ostree.object* &#x60;ostree.commit&#x60; - ostree.commit* &#x60;ostree.refs&#x60; - ostree.refs* &#x60;ostree.content&#x60; - ostree.content* &#x60;ostree.config&#x60; - ostree.config* &#x60;ostree.summary&#x60; - ostree.summary* &#x60;container.blob&#x60; - container.blob* &#x60;container.manifest&#x60; - container.manifest* &#x60;container.tag&#x60; - container.tag* &#x60;container.signature&#x60; - container.signature* &#x60;file.file&#x60; - file.file* &#x60;maven.artifact&#x60; - maven.artifact* &#x60;maven.metadata&#x60; - maven.metadata* &#x60;rpm.advisory&#x60; - rpm.advisory* &#x60;rpm.packagegroup&#x60; - rpm.packagegroup* &#x60;rpm.packagecategory&#x60; - rpm.packagecategory* &#x60;rpm.packageenvironment&#x60; - rpm.packageenvironment* &#x60;rpm.packagelangpacks&#x60; - rpm.packagelangpacks* &#x60;rpm.repo_metadata_file&#x60; - rpm.repo_metadata_file* &#x60;rpm.distribution_tree&#x60; - rpm.distribution_tree* &#x60;rpm.package&#x60; - rpm.package* &#x60;rpm.modulemd&#x60; - rpm.modulemd* &#x60;rpm.modulemd_defaults&#x60; - rpm.modulemd_defaults* &#x60;rpm.modulemd_obsolete&#x60; - rpm.modulemd_obsolete | 
 **pulpTypeIn** | **[]string** | Multiple values may be separated by commas.* &#x60;core.publishedmetadata&#x60; - core.publishedmetadata* &#x60;core.openpgp_publickey&#x60; - core.openpgp_publickey* &#x60;core.openpgp_publicsubkey&#x60; - core.openpgp_publicsubkey* &#x60;core.openpgp_userid&#x60; - core.openpgp_userid* &#x60;core.openpgp_userattribute&#x60; - core.openpgp_userattribute* &#x60;core.openpgp_signature&#x60; - core.openpgp_signature* &#x60;python.python&#x60; - python.python* &#x60;gem.gem&#x60; - gem.gem* &#x60;npm.package&#x60; - npm.package* &#x60;ostree.object&#x60; - ostree.object* &#x60;ostree.commit&#x60; - ostree.commit* &#x60;ostree.refs&#x60; - ostree.refs* &#x60;ostree.content&#x60; - ostree.content* &#x60;ostree.config&#x60; - ostree.config* &#x60;ostree.summary&#x60; - ostree.summary* &#x60;container.blob&#x60; - container.blob* &#x60;container.manifest&#x60; - container.manifest* &#x60;container.tag&#x60; - container.tag* &#x60;container.signature&#x60; - container.signature* &#x60;file.file&#x60; - file.file* &#x60;maven.artifact&#x60; - maven.artifact* &#x60;maven.metadata&#x60; - maven.metadata* &#x60;rpm.advisory&#x60; - rpm.advisory* &#x60;rpm.packagegroup&#x60; - rpm.packagegroup* &#x60;rpm.packagecategory&#x60; - rpm.packagecategory* &#x60;rpm.packageenvironment&#x60; - rpm.packageenvironment* &#x60;rpm.packagelangpacks&#x60; - rpm.packagelangpacks* &#x60;rpm.repo_metadata_file&#x60; - rpm.repo_metadata_file* &#x60;rpm.distribution_tree&#x60; - rpm.distribution_tree* &#x60;rpm.package&#x60; - rpm.package* &#x60;rpm.modulemd&#x60; - rpm.modulemd* &#x60;rpm.modulemd_defaults&#x60; - rpm.modulemd_defaults* &#x60;rpm.modulemd_obsolete&#x60; - rpm.modulemd_obsolete | 
 **q** | **string** | Filter results by using NOT, AND and OR operations on other filters | 
 **repositoryVersion** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionAdded** | **string** | Repository Version referenced by HREF/PRN | 
 **repositoryVersionRemoved** | **string** | Repository Version referenced by HREF/PRN | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedMultipleArtifactContentResponseList**](PaginatedMultipleArtifactContentResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

