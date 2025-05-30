# PatchedrpmRpmRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** | A unique name for this repository. | [optional] 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt64** | Retain X versions of the repository. Default is null which retains all versions. | [optional] 
**Remote** | Pointer to **NullableString** | An optional remote to use by default when syncing. | [optional] 
**Autopublish** | Pointer to **bool** | Whether to automatically create publications for new repository versions, and update any distributions pointing to this repository. | [optional] [default to false]
**MetadataSigningService** | Pointer to **NullableString** | A reference to an associated signing service. | [optional] 
**PackageSigningService** | Pointer to **NullableString** | A reference to an associated package signing service. | [optional] 
**PackageSigningFingerprint** | Pointer to **string** | The pubkey V4 fingerprint (160 bits) to be passed to the package signing service.The signing service will use that on signing operations related to this repository. | [optional] [default to ""]
**RetainPackageVersions** | Pointer to **int64** | The number of versions of each package to keep in the repository; older versions will be purged. The default is &#39;0&#39;, which will disable this feature and keep all versions of each package. | [optional] 
**ChecksumType** | Pointer to [**NullablePackageChecksumTypeEnum**](PackageChecksumTypeEnum.md) | The preferred checksum type during repo publish.* &#x60;unknown&#x60; - unknown* &#x60;md5&#x60; - md5* &#x60;sha1&#x60; - sha1* &#x60;sha224&#x60; - sha224* &#x60;sha256&#x60; - sha256* &#x60;sha384&#x60; - sha384* &#x60;sha512&#x60; - sha512 | [optional] 
**RepoConfig** | Pointer to **interface{}** | A JSON document describing the config.repo file Pulp should generate for this repo | [optional] 
**CompressionType** | Pointer to [**NullableCompressionTypeEnum**](CompressionTypeEnum.md) | The compression type to use for metadata files.* &#x60;zstd&#x60; - zstd* &#x60;gz&#x60; - gz | [optional] 
**Layout** | Pointer to [**NullableLayoutEnum**](LayoutEnum.md) | How to layout the packages within the published repository.* &#x60;nested_alphabetically&#x60; - nested_alphabetically* &#x60;flat&#x60; - flat | [optional] 

## Methods

### NewPatchedrpmRpmRepository

`func NewPatchedrpmRpmRepository() *PatchedrpmRpmRepository`

NewPatchedrpmRpmRepository instantiates a new PatchedrpmRpmRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedrpmRpmRepositoryWithDefaults

`func NewPatchedrpmRpmRepositoryWithDefaults() *PatchedrpmRpmRepository`

NewPatchedrpmRpmRepositoryWithDefaults instantiates a new PatchedrpmRpmRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpLabels

`func (o *PatchedrpmRpmRepository) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PatchedrpmRpmRepository) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PatchedrpmRpmRepository) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PatchedrpmRpmRepository) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *PatchedrpmRpmRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedrpmRpmRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedrpmRpmRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedrpmRpmRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedrpmRpmRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedrpmRpmRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedrpmRpmRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedrpmRpmRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedrpmRpmRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedrpmRpmRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *PatchedrpmRpmRepository) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *PatchedrpmRpmRepository) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *PatchedrpmRpmRepository) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *PatchedrpmRpmRepository) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *PatchedrpmRpmRepository) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *PatchedrpmRpmRepository) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *PatchedrpmRpmRepository) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *PatchedrpmRpmRepository) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *PatchedrpmRpmRepository) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *PatchedrpmRpmRepository) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *PatchedrpmRpmRepository) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *PatchedrpmRpmRepository) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetAutopublish

`func (o *PatchedrpmRpmRepository) GetAutopublish() bool`

GetAutopublish returns the Autopublish field if non-nil, zero value otherwise.

### GetAutopublishOk

`func (o *PatchedrpmRpmRepository) GetAutopublishOk() (*bool, bool)`

GetAutopublishOk returns a tuple with the Autopublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutopublish

`func (o *PatchedrpmRpmRepository) SetAutopublish(v bool)`

SetAutopublish sets Autopublish field to given value.

### HasAutopublish

`func (o *PatchedrpmRpmRepository) HasAutopublish() bool`

HasAutopublish returns a boolean if a field has been set.

### GetMetadataSigningService

`func (o *PatchedrpmRpmRepository) GetMetadataSigningService() string`

GetMetadataSigningService returns the MetadataSigningService field if non-nil, zero value otherwise.

### GetMetadataSigningServiceOk

`func (o *PatchedrpmRpmRepository) GetMetadataSigningServiceOk() (*string, bool)`

GetMetadataSigningServiceOk returns a tuple with the MetadataSigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSigningService

`func (o *PatchedrpmRpmRepository) SetMetadataSigningService(v string)`

SetMetadataSigningService sets MetadataSigningService field to given value.

### HasMetadataSigningService

`func (o *PatchedrpmRpmRepository) HasMetadataSigningService() bool`

HasMetadataSigningService returns a boolean if a field has been set.

### SetMetadataSigningServiceNil

`func (o *PatchedrpmRpmRepository) SetMetadataSigningServiceNil(b bool)`

 SetMetadataSigningServiceNil sets the value for MetadataSigningService to be an explicit nil

### UnsetMetadataSigningService
`func (o *PatchedrpmRpmRepository) UnsetMetadataSigningService()`

UnsetMetadataSigningService ensures that no value is present for MetadataSigningService, not even an explicit nil
### GetPackageSigningService

`func (o *PatchedrpmRpmRepository) GetPackageSigningService() string`

GetPackageSigningService returns the PackageSigningService field if non-nil, zero value otherwise.

### GetPackageSigningServiceOk

`func (o *PatchedrpmRpmRepository) GetPackageSigningServiceOk() (*string, bool)`

GetPackageSigningServiceOk returns a tuple with the PackageSigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSigningService

`func (o *PatchedrpmRpmRepository) SetPackageSigningService(v string)`

SetPackageSigningService sets PackageSigningService field to given value.

### HasPackageSigningService

`func (o *PatchedrpmRpmRepository) HasPackageSigningService() bool`

HasPackageSigningService returns a boolean if a field has been set.

### SetPackageSigningServiceNil

`func (o *PatchedrpmRpmRepository) SetPackageSigningServiceNil(b bool)`

 SetPackageSigningServiceNil sets the value for PackageSigningService to be an explicit nil

### UnsetPackageSigningService
`func (o *PatchedrpmRpmRepository) UnsetPackageSigningService()`

UnsetPackageSigningService ensures that no value is present for PackageSigningService, not even an explicit nil
### GetPackageSigningFingerprint

`func (o *PatchedrpmRpmRepository) GetPackageSigningFingerprint() string`

GetPackageSigningFingerprint returns the PackageSigningFingerprint field if non-nil, zero value otherwise.

### GetPackageSigningFingerprintOk

`func (o *PatchedrpmRpmRepository) GetPackageSigningFingerprintOk() (*string, bool)`

GetPackageSigningFingerprintOk returns a tuple with the PackageSigningFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSigningFingerprint

`func (o *PatchedrpmRpmRepository) SetPackageSigningFingerprint(v string)`

SetPackageSigningFingerprint sets PackageSigningFingerprint field to given value.

### HasPackageSigningFingerprint

`func (o *PatchedrpmRpmRepository) HasPackageSigningFingerprint() bool`

HasPackageSigningFingerprint returns a boolean if a field has been set.

### GetRetainPackageVersions

`func (o *PatchedrpmRpmRepository) GetRetainPackageVersions() int64`

GetRetainPackageVersions returns the RetainPackageVersions field if non-nil, zero value otherwise.

### GetRetainPackageVersionsOk

`func (o *PatchedrpmRpmRepository) GetRetainPackageVersionsOk() (*int64, bool)`

GetRetainPackageVersionsOk returns a tuple with the RetainPackageVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPackageVersions

`func (o *PatchedrpmRpmRepository) SetRetainPackageVersions(v int64)`

SetRetainPackageVersions sets RetainPackageVersions field to given value.

### HasRetainPackageVersions

`func (o *PatchedrpmRpmRepository) HasRetainPackageVersions() bool`

HasRetainPackageVersions returns a boolean if a field has been set.

### GetChecksumType

`func (o *PatchedrpmRpmRepository) GetChecksumType() PackageChecksumTypeEnum`

GetChecksumType returns the ChecksumType field if non-nil, zero value otherwise.

### GetChecksumTypeOk

`func (o *PatchedrpmRpmRepository) GetChecksumTypeOk() (*PackageChecksumTypeEnum, bool)`

GetChecksumTypeOk returns a tuple with the ChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumType

`func (o *PatchedrpmRpmRepository) SetChecksumType(v PackageChecksumTypeEnum)`

SetChecksumType sets ChecksumType field to given value.

### HasChecksumType

`func (o *PatchedrpmRpmRepository) HasChecksumType() bool`

HasChecksumType returns a boolean if a field has been set.

### SetChecksumTypeNil

`func (o *PatchedrpmRpmRepository) SetChecksumTypeNil(b bool)`

 SetChecksumTypeNil sets the value for ChecksumType to be an explicit nil

### UnsetChecksumType
`func (o *PatchedrpmRpmRepository) UnsetChecksumType()`

UnsetChecksumType ensures that no value is present for ChecksumType, not even an explicit nil
### GetRepoConfig

`func (o *PatchedrpmRpmRepository) GetRepoConfig() interface{}`

GetRepoConfig returns the RepoConfig field if non-nil, zero value otherwise.

### GetRepoConfigOk

`func (o *PatchedrpmRpmRepository) GetRepoConfigOk() (*interface{}, bool)`

GetRepoConfigOk returns a tuple with the RepoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoConfig

`func (o *PatchedrpmRpmRepository) SetRepoConfig(v interface{})`

SetRepoConfig sets RepoConfig field to given value.

### HasRepoConfig

`func (o *PatchedrpmRpmRepository) HasRepoConfig() bool`

HasRepoConfig returns a boolean if a field has been set.

### SetRepoConfigNil

`func (o *PatchedrpmRpmRepository) SetRepoConfigNil(b bool)`

 SetRepoConfigNil sets the value for RepoConfig to be an explicit nil

### UnsetRepoConfig
`func (o *PatchedrpmRpmRepository) UnsetRepoConfig()`

UnsetRepoConfig ensures that no value is present for RepoConfig, not even an explicit nil
### GetCompressionType

`func (o *PatchedrpmRpmRepository) GetCompressionType() CompressionTypeEnum`

GetCompressionType returns the CompressionType field if non-nil, zero value otherwise.

### GetCompressionTypeOk

`func (o *PatchedrpmRpmRepository) GetCompressionTypeOk() (*CompressionTypeEnum, bool)`

GetCompressionTypeOk returns a tuple with the CompressionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionType

`func (o *PatchedrpmRpmRepository) SetCompressionType(v CompressionTypeEnum)`

SetCompressionType sets CompressionType field to given value.

### HasCompressionType

`func (o *PatchedrpmRpmRepository) HasCompressionType() bool`

HasCompressionType returns a boolean if a field has been set.

### SetCompressionTypeNil

`func (o *PatchedrpmRpmRepository) SetCompressionTypeNil(b bool)`

 SetCompressionTypeNil sets the value for CompressionType to be an explicit nil

### UnsetCompressionType
`func (o *PatchedrpmRpmRepository) UnsetCompressionType()`

UnsetCompressionType ensures that no value is present for CompressionType, not even an explicit nil
### GetLayout

`func (o *PatchedrpmRpmRepository) GetLayout() LayoutEnum`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *PatchedrpmRpmRepository) GetLayoutOk() (*LayoutEnum, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *PatchedrpmRpmRepository) SetLayout(v LayoutEnum)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *PatchedrpmRpmRepository) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### SetLayoutNil

`func (o *PatchedrpmRpmRepository) SetLayoutNil(b bool)`

 SetLayoutNil sets the value for Layout to be an explicit nil

### UnsetLayout
`func (o *PatchedrpmRpmRepository) UnsetLayout()`

UnsetLayout ensures that no value is present for Layout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


