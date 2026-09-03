# PythonPackageContentUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overwrite** | Pointer to **bool** | When set to true, existing content in the repository with the same unique key will be silently overwritten. When set to false, the task will fail if content would be overwritten. Only used when &#39;repository&#39; is specified. Defaults to true. | [optional] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**Artifact** | Pointer to **string** | Artifact file representing the physical content | [optional] 
**File** | Pointer to ***os.File** | An uploaded file that may be turned into the content unit. | [optional] 
**Upload** | Pointer to **string** | An uncommitted upload that may be turned into the content unit. | [optional] 
**FileUrl** | Pointer to **string** | A url that Pulp can download and turn into the content unit. | [optional] 
**DownloaderConfig** | Pointer to [**RemoteNetworkConfig**](RemoteNetworkConfig.md) | Configuration for the download process (e.g., proxies, auth, timeouts). Only applicable when providing a &#39;file_url. | [optional] 
**Author** | Pointer to **string** | Text containing the author&#39;s name. Contact information can also be added, separated with newlines. | [optional] 
**AuthorEmail** | Pointer to **string** | The author&#39;s e-mail address.  | [optional] 
**Description** | Pointer to **string** | A longer description of the package that can run to several paragraphs. | [optional] 
**HomePage** | Pointer to **string** | The URL for the package&#39;s home page. | [optional] 
**Keywords** | Pointer to **string** | Additional keywords to be used to assist searching for the package in a larger catalog. | [optional] 
**License** | Pointer to **string** | Text indicating the license covering the distribution | [optional] 
**Platform** | Pointer to **string** | A comma-separated list of platform specifications, summarizing the operating systems supported by the package. | [optional] 
**Summary** | Pointer to **string** | A one-line summary of what the package does. | [optional] 
**Classifiers** | Pointer to **interface{}** | A JSON list containing classification values for a Python package. | [optional] 
**DownloadUrl** | Pointer to **string** | Legacy field denoting the URL from which this package can be downloaded. | [optional] 
**SupportedPlatform** | Pointer to **string** | Field to specify the OS and CPU for which the binary package was compiled.  | [optional] 
**Maintainer** | Pointer to **string** | The maintainer&#39;s name at a minimum; additional contact information may be provided. | [optional] 
**MaintainerEmail** | Pointer to **string** | The maintainer&#39;s e-mail address. | [optional] 
**ObsoletesDist** | Pointer to **interface{}** | A JSON list containing names of a distutils project&#39;s distribution which this distribution renders obsolete, meaning that the two projects should not be installed at the same time. | [optional] 
**ProjectUrl** | Pointer to **string** | A browsable URL for the project and a label for it, separated by a comma. | [optional] 
**ProjectUrls** | Pointer to **interface{}** | A dictionary of labels and URLs for the project. | [optional] 
**ProvidesDist** | Pointer to **interface{}** | A JSON list containing names of a Distutils project which is contained within this distribution. | [optional] 
**RequiresExternal** | Pointer to **interface{}** | A JSON list containing some dependency in the system that the distribution is to be used. | [optional] 
**RequiresDist** | Pointer to **interface{}** | A JSON list containing names of some other distutils project required by this distribution. | [optional] 
**RequiresPython** | Pointer to **string** | The Python version(s) that the distribution is guaranteed to be compatible with. | [optional] 
**DescriptionContentType** | Pointer to **string** | A string stating the markup syntax (if any) used in the distribution&#39;s description, so that tools can intelligently render the description. | [optional] 
**ProvidesExtras** | Pointer to **interface{}** | A JSON list containing names of optional features provided by the package. | [optional] 
**Dynamic** | Pointer to **interface{}** | A JSON list containing names of other core metadata fields which are permitted to vary between sdist and bdist packages. Fields NOT marked dynamic MUST be the same between bdist and sdist. | [optional] 
**LicenseExpression** | Pointer to **string** | Text string that is a valid SPDX license expression. | [optional] 
**LicenseFile** | Pointer to **interface{}** | A JSON list containing names of the paths to license-related files. | [optional] 
**Sha256** | Pointer to **string** | The SHA256 digest of this package. | [optional] [default to ""]
**MetadataSha256** | Pointer to **NullableString** | The SHA256 digest of the package&#39;s METADATA file. | [optional] 
**Attestations** | Pointer to **interface{}** | A JSON list containing attestations for the package. | [optional] 

## Methods

### NewPythonPackageContentUpload

`func NewPythonPackageContentUpload() *PythonPackageContentUpload`

NewPythonPackageContentUpload instantiates a new PythonPackageContentUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPythonPackageContentUploadWithDefaults

`func NewPythonPackageContentUploadWithDefaults() *PythonPackageContentUpload`

NewPythonPackageContentUploadWithDefaults instantiates a new PythonPackageContentUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverwrite

`func (o *PythonPackageContentUpload) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *PythonPackageContentUpload) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *PythonPackageContentUpload) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *PythonPackageContentUpload) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetPulpLabels

`func (o *PythonPackageContentUpload) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PythonPackageContentUpload) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PythonPackageContentUpload) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PythonPackageContentUpload) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetArtifact

`func (o *PythonPackageContentUpload) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *PythonPackageContentUpload) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *PythonPackageContentUpload) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *PythonPackageContentUpload) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetFile

`func (o *PythonPackageContentUpload) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *PythonPackageContentUpload) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *PythonPackageContentUpload) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *PythonPackageContentUpload) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetUpload

`func (o *PythonPackageContentUpload) GetUpload() string`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *PythonPackageContentUpload) GetUploadOk() (*string, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *PythonPackageContentUpload) SetUpload(v string)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *PythonPackageContentUpload) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetFileUrl

`func (o *PythonPackageContentUpload) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *PythonPackageContentUpload) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *PythonPackageContentUpload) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *PythonPackageContentUpload) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetDownloaderConfig

`func (o *PythonPackageContentUpload) GetDownloaderConfig() RemoteNetworkConfig`

GetDownloaderConfig returns the DownloaderConfig field if non-nil, zero value otherwise.

### GetDownloaderConfigOk

`func (o *PythonPackageContentUpload) GetDownloaderConfigOk() (*RemoteNetworkConfig, bool)`

GetDownloaderConfigOk returns a tuple with the DownloaderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloaderConfig

`func (o *PythonPackageContentUpload) SetDownloaderConfig(v RemoteNetworkConfig)`

SetDownloaderConfig sets DownloaderConfig field to given value.

### HasDownloaderConfig

`func (o *PythonPackageContentUpload) HasDownloaderConfig() bool`

HasDownloaderConfig returns a boolean if a field has been set.

### GetAuthor

`func (o *PythonPackageContentUpload) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PythonPackageContentUpload) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PythonPackageContentUpload) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *PythonPackageContentUpload) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetAuthorEmail

`func (o *PythonPackageContentUpload) GetAuthorEmail() string`

GetAuthorEmail returns the AuthorEmail field if non-nil, zero value otherwise.

### GetAuthorEmailOk

`func (o *PythonPackageContentUpload) GetAuthorEmailOk() (*string, bool)`

GetAuthorEmailOk returns a tuple with the AuthorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorEmail

`func (o *PythonPackageContentUpload) SetAuthorEmail(v string)`

SetAuthorEmail sets AuthorEmail field to given value.

### HasAuthorEmail

`func (o *PythonPackageContentUpload) HasAuthorEmail() bool`

HasAuthorEmail returns a boolean if a field has been set.

### GetDescription

`func (o *PythonPackageContentUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PythonPackageContentUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PythonPackageContentUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PythonPackageContentUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHomePage

`func (o *PythonPackageContentUpload) GetHomePage() string`

GetHomePage returns the HomePage field if non-nil, zero value otherwise.

### GetHomePageOk

`func (o *PythonPackageContentUpload) GetHomePageOk() (*string, bool)`

GetHomePageOk returns a tuple with the HomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePage

`func (o *PythonPackageContentUpload) SetHomePage(v string)`

SetHomePage sets HomePage field to given value.

### HasHomePage

`func (o *PythonPackageContentUpload) HasHomePage() bool`

HasHomePage returns a boolean if a field has been set.

### GetKeywords

`func (o *PythonPackageContentUpload) GetKeywords() string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *PythonPackageContentUpload) GetKeywordsOk() (*string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *PythonPackageContentUpload) SetKeywords(v string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *PythonPackageContentUpload) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetLicense

`func (o *PythonPackageContentUpload) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *PythonPackageContentUpload) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *PythonPackageContentUpload) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *PythonPackageContentUpload) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetPlatform

`func (o *PythonPackageContentUpload) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PythonPackageContentUpload) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PythonPackageContentUpload) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PythonPackageContentUpload) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSummary

`func (o *PythonPackageContentUpload) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PythonPackageContentUpload) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PythonPackageContentUpload) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PythonPackageContentUpload) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetClassifiers

`func (o *PythonPackageContentUpload) GetClassifiers() interface{}`

GetClassifiers returns the Classifiers field if non-nil, zero value otherwise.

### GetClassifiersOk

`func (o *PythonPackageContentUpload) GetClassifiersOk() (*interface{}, bool)`

GetClassifiersOk returns a tuple with the Classifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifiers

`func (o *PythonPackageContentUpload) SetClassifiers(v interface{})`

SetClassifiers sets Classifiers field to given value.

### HasClassifiers

`func (o *PythonPackageContentUpload) HasClassifiers() bool`

HasClassifiers returns a boolean if a field has been set.

### SetClassifiersNil

`func (o *PythonPackageContentUpload) SetClassifiersNil(b bool)`

 SetClassifiersNil sets the value for Classifiers to be an explicit nil

### UnsetClassifiers
`func (o *PythonPackageContentUpload) UnsetClassifiers()`

UnsetClassifiers ensures that no value is present for Classifiers, not even an explicit nil
### GetDownloadUrl

`func (o *PythonPackageContentUpload) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *PythonPackageContentUpload) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *PythonPackageContentUpload) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *PythonPackageContentUpload) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetSupportedPlatform

`func (o *PythonPackageContentUpload) GetSupportedPlatform() string`

GetSupportedPlatform returns the SupportedPlatform field if non-nil, zero value otherwise.

### GetSupportedPlatformOk

`func (o *PythonPackageContentUpload) GetSupportedPlatformOk() (*string, bool)`

GetSupportedPlatformOk returns a tuple with the SupportedPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPlatform

`func (o *PythonPackageContentUpload) SetSupportedPlatform(v string)`

SetSupportedPlatform sets SupportedPlatform field to given value.

### HasSupportedPlatform

`func (o *PythonPackageContentUpload) HasSupportedPlatform() bool`

HasSupportedPlatform returns a boolean if a field has been set.

### GetMaintainer

`func (o *PythonPackageContentUpload) GetMaintainer() string`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *PythonPackageContentUpload) GetMaintainerOk() (*string, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *PythonPackageContentUpload) SetMaintainer(v string)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *PythonPackageContentUpload) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetMaintainerEmail

`func (o *PythonPackageContentUpload) GetMaintainerEmail() string`

GetMaintainerEmail returns the MaintainerEmail field if non-nil, zero value otherwise.

### GetMaintainerEmailOk

`func (o *PythonPackageContentUpload) GetMaintainerEmailOk() (*string, bool)`

GetMaintainerEmailOk returns a tuple with the MaintainerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerEmail

`func (o *PythonPackageContentUpload) SetMaintainerEmail(v string)`

SetMaintainerEmail sets MaintainerEmail field to given value.

### HasMaintainerEmail

`func (o *PythonPackageContentUpload) HasMaintainerEmail() bool`

HasMaintainerEmail returns a boolean if a field has been set.

### GetObsoletesDist

`func (o *PythonPackageContentUpload) GetObsoletesDist() interface{}`

GetObsoletesDist returns the ObsoletesDist field if non-nil, zero value otherwise.

### GetObsoletesDistOk

`func (o *PythonPackageContentUpload) GetObsoletesDistOk() (*interface{}, bool)`

GetObsoletesDistOk returns a tuple with the ObsoletesDist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoletesDist

`func (o *PythonPackageContentUpload) SetObsoletesDist(v interface{})`

SetObsoletesDist sets ObsoletesDist field to given value.

### HasObsoletesDist

`func (o *PythonPackageContentUpload) HasObsoletesDist() bool`

HasObsoletesDist returns a boolean if a field has been set.

### SetObsoletesDistNil

`func (o *PythonPackageContentUpload) SetObsoletesDistNil(b bool)`

 SetObsoletesDistNil sets the value for ObsoletesDist to be an explicit nil

### UnsetObsoletesDist
`func (o *PythonPackageContentUpload) UnsetObsoletesDist()`

UnsetObsoletesDist ensures that no value is present for ObsoletesDist, not even an explicit nil
### GetProjectUrl

`func (o *PythonPackageContentUpload) GetProjectUrl() string`

GetProjectUrl returns the ProjectUrl field if non-nil, zero value otherwise.

### GetProjectUrlOk

`func (o *PythonPackageContentUpload) GetProjectUrlOk() (*string, bool)`

GetProjectUrlOk returns a tuple with the ProjectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUrl

`func (o *PythonPackageContentUpload) SetProjectUrl(v string)`

SetProjectUrl sets ProjectUrl field to given value.

### HasProjectUrl

`func (o *PythonPackageContentUpload) HasProjectUrl() bool`

HasProjectUrl returns a boolean if a field has been set.

### GetProjectUrls

`func (o *PythonPackageContentUpload) GetProjectUrls() interface{}`

GetProjectUrls returns the ProjectUrls field if non-nil, zero value otherwise.

### GetProjectUrlsOk

`func (o *PythonPackageContentUpload) GetProjectUrlsOk() (*interface{}, bool)`

GetProjectUrlsOk returns a tuple with the ProjectUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUrls

`func (o *PythonPackageContentUpload) SetProjectUrls(v interface{})`

SetProjectUrls sets ProjectUrls field to given value.

### HasProjectUrls

`func (o *PythonPackageContentUpload) HasProjectUrls() bool`

HasProjectUrls returns a boolean if a field has been set.

### SetProjectUrlsNil

`func (o *PythonPackageContentUpload) SetProjectUrlsNil(b bool)`

 SetProjectUrlsNil sets the value for ProjectUrls to be an explicit nil

### UnsetProjectUrls
`func (o *PythonPackageContentUpload) UnsetProjectUrls()`

UnsetProjectUrls ensures that no value is present for ProjectUrls, not even an explicit nil
### GetProvidesDist

`func (o *PythonPackageContentUpload) GetProvidesDist() interface{}`

GetProvidesDist returns the ProvidesDist field if non-nil, zero value otherwise.

### GetProvidesDistOk

`func (o *PythonPackageContentUpload) GetProvidesDistOk() (*interface{}, bool)`

GetProvidesDistOk returns a tuple with the ProvidesDist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidesDist

`func (o *PythonPackageContentUpload) SetProvidesDist(v interface{})`

SetProvidesDist sets ProvidesDist field to given value.

### HasProvidesDist

`func (o *PythonPackageContentUpload) HasProvidesDist() bool`

HasProvidesDist returns a boolean if a field has been set.

### SetProvidesDistNil

`func (o *PythonPackageContentUpload) SetProvidesDistNil(b bool)`

 SetProvidesDistNil sets the value for ProvidesDist to be an explicit nil

### UnsetProvidesDist
`func (o *PythonPackageContentUpload) UnsetProvidesDist()`

UnsetProvidesDist ensures that no value is present for ProvidesDist, not even an explicit nil
### GetRequiresExternal

`func (o *PythonPackageContentUpload) GetRequiresExternal() interface{}`

GetRequiresExternal returns the RequiresExternal field if non-nil, zero value otherwise.

### GetRequiresExternalOk

`func (o *PythonPackageContentUpload) GetRequiresExternalOk() (*interface{}, bool)`

GetRequiresExternalOk returns a tuple with the RequiresExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresExternal

`func (o *PythonPackageContentUpload) SetRequiresExternal(v interface{})`

SetRequiresExternal sets RequiresExternal field to given value.

### HasRequiresExternal

`func (o *PythonPackageContentUpload) HasRequiresExternal() bool`

HasRequiresExternal returns a boolean if a field has been set.

### SetRequiresExternalNil

`func (o *PythonPackageContentUpload) SetRequiresExternalNil(b bool)`

 SetRequiresExternalNil sets the value for RequiresExternal to be an explicit nil

### UnsetRequiresExternal
`func (o *PythonPackageContentUpload) UnsetRequiresExternal()`

UnsetRequiresExternal ensures that no value is present for RequiresExternal, not even an explicit nil
### GetRequiresDist

`func (o *PythonPackageContentUpload) GetRequiresDist() interface{}`

GetRequiresDist returns the RequiresDist field if non-nil, zero value otherwise.

### GetRequiresDistOk

`func (o *PythonPackageContentUpload) GetRequiresDistOk() (*interface{}, bool)`

GetRequiresDistOk returns a tuple with the RequiresDist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresDist

`func (o *PythonPackageContentUpload) SetRequiresDist(v interface{})`

SetRequiresDist sets RequiresDist field to given value.

### HasRequiresDist

`func (o *PythonPackageContentUpload) HasRequiresDist() bool`

HasRequiresDist returns a boolean if a field has been set.

### SetRequiresDistNil

`func (o *PythonPackageContentUpload) SetRequiresDistNil(b bool)`

 SetRequiresDistNil sets the value for RequiresDist to be an explicit nil

### UnsetRequiresDist
`func (o *PythonPackageContentUpload) UnsetRequiresDist()`

UnsetRequiresDist ensures that no value is present for RequiresDist, not even an explicit nil
### GetRequiresPython

`func (o *PythonPackageContentUpload) GetRequiresPython() string`

GetRequiresPython returns the RequiresPython field if non-nil, zero value otherwise.

### GetRequiresPythonOk

`func (o *PythonPackageContentUpload) GetRequiresPythonOk() (*string, bool)`

GetRequiresPythonOk returns a tuple with the RequiresPython field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPython

`func (o *PythonPackageContentUpload) SetRequiresPython(v string)`

SetRequiresPython sets RequiresPython field to given value.

### HasRequiresPython

`func (o *PythonPackageContentUpload) HasRequiresPython() bool`

HasRequiresPython returns a boolean if a field has been set.

### GetDescriptionContentType

`func (o *PythonPackageContentUpload) GetDescriptionContentType() string`

GetDescriptionContentType returns the DescriptionContentType field if non-nil, zero value otherwise.

### GetDescriptionContentTypeOk

`func (o *PythonPackageContentUpload) GetDescriptionContentTypeOk() (*string, bool)`

GetDescriptionContentTypeOk returns a tuple with the DescriptionContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionContentType

`func (o *PythonPackageContentUpload) SetDescriptionContentType(v string)`

SetDescriptionContentType sets DescriptionContentType field to given value.

### HasDescriptionContentType

`func (o *PythonPackageContentUpload) HasDescriptionContentType() bool`

HasDescriptionContentType returns a boolean if a field has been set.

### GetProvidesExtras

`func (o *PythonPackageContentUpload) GetProvidesExtras() interface{}`

GetProvidesExtras returns the ProvidesExtras field if non-nil, zero value otherwise.

### GetProvidesExtrasOk

`func (o *PythonPackageContentUpload) GetProvidesExtrasOk() (*interface{}, bool)`

GetProvidesExtrasOk returns a tuple with the ProvidesExtras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidesExtras

`func (o *PythonPackageContentUpload) SetProvidesExtras(v interface{})`

SetProvidesExtras sets ProvidesExtras field to given value.

### HasProvidesExtras

`func (o *PythonPackageContentUpload) HasProvidesExtras() bool`

HasProvidesExtras returns a boolean if a field has been set.

### SetProvidesExtrasNil

`func (o *PythonPackageContentUpload) SetProvidesExtrasNil(b bool)`

 SetProvidesExtrasNil sets the value for ProvidesExtras to be an explicit nil

### UnsetProvidesExtras
`func (o *PythonPackageContentUpload) UnsetProvidesExtras()`

UnsetProvidesExtras ensures that no value is present for ProvidesExtras, not even an explicit nil
### GetDynamic

`func (o *PythonPackageContentUpload) GetDynamic() interface{}`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *PythonPackageContentUpload) GetDynamicOk() (*interface{}, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *PythonPackageContentUpload) SetDynamic(v interface{})`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *PythonPackageContentUpload) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### SetDynamicNil

`func (o *PythonPackageContentUpload) SetDynamicNil(b bool)`

 SetDynamicNil sets the value for Dynamic to be an explicit nil

### UnsetDynamic
`func (o *PythonPackageContentUpload) UnsetDynamic()`

UnsetDynamic ensures that no value is present for Dynamic, not even an explicit nil
### GetLicenseExpression

`func (o *PythonPackageContentUpload) GetLicenseExpression() string`

GetLicenseExpression returns the LicenseExpression field if non-nil, zero value otherwise.

### GetLicenseExpressionOk

`func (o *PythonPackageContentUpload) GetLicenseExpressionOk() (*string, bool)`

GetLicenseExpressionOk returns a tuple with the LicenseExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpression

`func (o *PythonPackageContentUpload) SetLicenseExpression(v string)`

SetLicenseExpression sets LicenseExpression field to given value.

### HasLicenseExpression

`func (o *PythonPackageContentUpload) HasLicenseExpression() bool`

HasLicenseExpression returns a boolean if a field has been set.

### GetLicenseFile

`func (o *PythonPackageContentUpload) GetLicenseFile() interface{}`

GetLicenseFile returns the LicenseFile field if non-nil, zero value otherwise.

### GetLicenseFileOk

`func (o *PythonPackageContentUpload) GetLicenseFileOk() (*interface{}, bool)`

GetLicenseFileOk returns a tuple with the LicenseFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseFile

`func (o *PythonPackageContentUpload) SetLicenseFile(v interface{})`

SetLicenseFile sets LicenseFile field to given value.

### HasLicenseFile

`func (o *PythonPackageContentUpload) HasLicenseFile() bool`

HasLicenseFile returns a boolean if a field has been set.

### SetLicenseFileNil

`func (o *PythonPackageContentUpload) SetLicenseFileNil(b bool)`

 SetLicenseFileNil sets the value for LicenseFile to be an explicit nil

### UnsetLicenseFile
`func (o *PythonPackageContentUpload) UnsetLicenseFile()`

UnsetLicenseFile ensures that no value is present for LicenseFile, not even an explicit nil
### GetSha256

`func (o *PythonPackageContentUpload) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *PythonPackageContentUpload) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *PythonPackageContentUpload) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *PythonPackageContentUpload) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetMetadataSha256

`func (o *PythonPackageContentUpload) GetMetadataSha256() string`

GetMetadataSha256 returns the MetadataSha256 field if non-nil, zero value otherwise.

### GetMetadataSha256Ok

`func (o *PythonPackageContentUpload) GetMetadataSha256Ok() (*string, bool)`

GetMetadataSha256Ok returns a tuple with the MetadataSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSha256

`func (o *PythonPackageContentUpload) SetMetadataSha256(v string)`

SetMetadataSha256 sets MetadataSha256 field to given value.

### HasMetadataSha256

`func (o *PythonPackageContentUpload) HasMetadataSha256() bool`

HasMetadataSha256 returns a boolean if a field has been set.

### SetMetadataSha256Nil

`func (o *PythonPackageContentUpload) SetMetadataSha256Nil(b bool)`

 SetMetadataSha256Nil sets the value for MetadataSha256 to be an explicit nil

### UnsetMetadataSha256
`func (o *PythonPackageContentUpload) UnsetMetadataSha256()`

UnsetMetadataSha256 ensures that no value is present for MetadataSha256, not even an explicit nil
### GetAttestations

`func (o *PythonPackageContentUpload) GetAttestations() interface{}`

GetAttestations returns the Attestations field if non-nil, zero value otherwise.

### GetAttestationsOk

`func (o *PythonPackageContentUpload) GetAttestationsOk() (*interface{}, bool)`

GetAttestationsOk returns a tuple with the Attestations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestations

`func (o *PythonPackageContentUpload) SetAttestations(v interface{})`

SetAttestations sets Attestations field to given value.

### HasAttestations

`func (o *PythonPackageContentUpload) HasAttestations() bool`

HasAttestations returns a boolean if a field has been set.

### SetAttestationsNil

`func (o *PythonPackageContentUpload) SetAttestationsNil(b bool)`

 SetAttestationsNil sets the value for Attestations to be an explicit nil

### UnsetAttestations
`func (o *PythonPackageContentUpload) UnsetAttestations()`

UnsetAttestations ensures that no value is present for Attestations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


