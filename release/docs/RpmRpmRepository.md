# RpmRpmRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** | A unique name for this repository. | 
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

### NewRpmRpmRepository

`func NewRpmRpmRepository(name string, ) *RpmRpmRepository`

NewRpmRpmRepository instantiates a new RpmRpmRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmRpmRepositoryWithDefaults

`func NewRpmRpmRepositoryWithDefaults() *RpmRpmRepository`

NewRpmRpmRepositoryWithDefaults instantiates a new RpmRpmRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpLabels

`func (o *RpmRpmRepository) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *RpmRpmRepository) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *RpmRpmRepository) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *RpmRpmRepository) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *RpmRpmRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RpmRpmRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RpmRpmRepository) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RpmRpmRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RpmRpmRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RpmRpmRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RpmRpmRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RpmRpmRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RpmRpmRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *RpmRpmRepository) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *RpmRpmRepository) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *RpmRpmRepository) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *RpmRpmRepository) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *RpmRpmRepository) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *RpmRpmRepository) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *RpmRpmRepository) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *RpmRpmRepository) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *RpmRpmRepository) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *RpmRpmRepository) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *RpmRpmRepository) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *RpmRpmRepository) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetAutopublish

`func (o *RpmRpmRepository) GetAutopublish() bool`

GetAutopublish returns the Autopublish field if non-nil, zero value otherwise.

### GetAutopublishOk

`func (o *RpmRpmRepository) GetAutopublishOk() (*bool, bool)`

GetAutopublishOk returns a tuple with the Autopublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutopublish

`func (o *RpmRpmRepository) SetAutopublish(v bool)`

SetAutopublish sets Autopublish field to given value.

### HasAutopublish

`func (o *RpmRpmRepository) HasAutopublish() bool`

HasAutopublish returns a boolean if a field has been set.

### GetMetadataSigningService

`func (o *RpmRpmRepository) GetMetadataSigningService() string`

GetMetadataSigningService returns the MetadataSigningService field if non-nil, zero value otherwise.

### GetMetadataSigningServiceOk

`func (o *RpmRpmRepository) GetMetadataSigningServiceOk() (*string, bool)`

GetMetadataSigningServiceOk returns a tuple with the MetadataSigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSigningService

`func (o *RpmRpmRepository) SetMetadataSigningService(v string)`

SetMetadataSigningService sets MetadataSigningService field to given value.

### HasMetadataSigningService

`func (o *RpmRpmRepository) HasMetadataSigningService() bool`

HasMetadataSigningService returns a boolean if a field has been set.

### SetMetadataSigningServiceNil

`func (o *RpmRpmRepository) SetMetadataSigningServiceNil(b bool)`

 SetMetadataSigningServiceNil sets the value for MetadataSigningService to be an explicit nil

### UnsetMetadataSigningService
`func (o *RpmRpmRepository) UnsetMetadataSigningService()`

UnsetMetadataSigningService ensures that no value is present for MetadataSigningService, not even an explicit nil
### GetPackageSigningService

`func (o *RpmRpmRepository) GetPackageSigningService() string`

GetPackageSigningService returns the PackageSigningService field if non-nil, zero value otherwise.

### GetPackageSigningServiceOk

`func (o *RpmRpmRepository) GetPackageSigningServiceOk() (*string, bool)`

GetPackageSigningServiceOk returns a tuple with the PackageSigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSigningService

`func (o *RpmRpmRepository) SetPackageSigningService(v string)`

SetPackageSigningService sets PackageSigningService field to given value.

### HasPackageSigningService

`func (o *RpmRpmRepository) HasPackageSigningService() bool`

HasPackageSigningService returns a boolean if a field has been set.

### SetPackageSigningServiceNil

`func (o *RpmRpmRepository) SetPackageSigningServiceNil(b bool)`

 SetPackageSigningServiceNil sets the value for PackageSigningService to be an explicit nil

### UnsetPackageSigningService
`func (o *RpmRpmRepository) UnsetPackageSigningService()`

UnsetPackageSigningService ensures that no value is present for PackageSigningService, not even an explicit nil
### GetPackageSigningFingerprint

`func (o *RpmRpmRepository) GetPackageSigningFingerprint() string`

GetPackageSigningFingerprint returns the PackageSigningFingerprint field if non-nil, zero value otherwise.

### GetPackageSigningFingerprintOk

`func (o *RpmRpmRepository) GetPackageSigningFingerprintOk() (*string, bool)`

GetPackageSigningFingerprintOk returns a tuple with the PackageSigningFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSigningFingerprint

`func (o *RpmRpmRepository) SetPackageSigningFingerprint(v string)`

SetPackageSigningFingerprint sets PackageSigningFingerprint field to given value.

### HasPackageSigningFingerprint

`func (o *RpmRpmRepository) HasPackageSigningFingerprint() bool`

HasPackageSigningFingerprint returns a boolean if a field has been set.

### GetRetainPackageVersions

`func (o *RpmRpmRepository) GetRetainPackageVersions() int64`

GetRetainPackageVersions returns the RetainPackageVersions field if non-nil, zero value otherwise.

### GetRetainPackageVersionsOk

`func (o *RpmRpmRepository) GetRetainPackageVersionsOk() (*int64, bool)`

GetRetainPackageVersionsOk returns a tuple with the RetainPackageVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPackageVersions

`func (o *RpmRpmRepository) SetRetainPackageVersions(v int64)`

SetRetainPackageVersions sets RetainPackageVersions field to given value.

### HasRetainPackageVersions

`func (o *RpmRpmRepository) HasRetainPackageVersions() bool`

HasRetainPackageVersions returns a boolean if a field has been set.

### GetChecksumType

`func (o *RpmRpmRepository) GetChecksumType() PackageChecksumTypeEnum`

GetChecksumType returns the ChecksumType field if non-nil, zero value otherwise.

### GetChecksumTypeOk

`func (o *RpmRpmRepository) GetChecksumTypeOk() (*PackageChecksumTypeEnum, bool)`

GetChecksumTypeOk returns a tuple with the ChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumType

`func (o *RpmRpmRepository) SetChecksumType(v PackageChecksumTypeEnum)`

SetChecksumType sets ChecksumType field to given value.

### HasChecksumType

`func (o *RpmRpmRepository) HasChecksumType() bool`

HasChecksumType returns a boolean if a field has been set.

### SetChecksumTypeNil

`func (o *RpmRpmRepository) SetChecksumTypeNil(b bool)`

 SetChecksumTypeNil sets the value for ChecksumType to be an explicit nil

### UnsetChecksumType
`func (o *RpmRpmRepository) UnsetChecksumType()`

UnsetChecksumType ensures that no value is present for ChecksumType, not even an explicit nil
### GetRepoConfig

`func (o *RpmRpmRepository) GetRepoConfig() interface{}`

GetRepoConfig returns the RepoConfig field if non-nil, zero value otherwise.

### GetRepoConfigOk

`func (o *RpmRpmRepository) GetRepoConfigOk() (*interface{}, bool)`

GetRepoConfigOk returns a tuple with the RepoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoConfig

`func (o *RpmRpmRepository) SetRepoConfig(v interface{})`

SetRepoConfig sets RepoConfig field to given value.

### HasRepoConfig

`func (o *RpmRpmRepository) HasRepoConfig() bool`

HasRepoConfig returns a boolean if a field has been set.

### SetRepoConfigNil

`func (o *RpmRpmRepository) SetRepoConfigNil(b bool)`

 SetRepoConfigNil sets the value for RepoConfig to be an explicit nil

### UnsetRepoConfig
`func (o *RpmRpmRepository) UnsetRepoConfig()`

UnsetRepoConfig ensures that no value is present for RepoConfig, not even an explicit nil
### GetCompressionType

`func (o *RpmRpmRepository) GetCompressionType() CompressionTypeEnum`

GetCompressionType returns the CompressionType field if non-nil, zero value otherwise.

### GetCompressionTypeOk

`func (o *RpmRpmRepository) GetCompressionTypeOk() (*CompressionTypeEnum, bool)`

GetCompressionTypeOk returns a tuple with the CompressionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionType

`func (o *RpmRpmRepository) SetCompressionType(v CompressionTypeEnum)`

SetCompressionType sets CompressionType field to given value.

### HasCompressionType

`func (o *RpmRpmRepository) HasCompressionType() bool`

HasCompressionType returns a boolean if a field has been set.

### SetCompressionTypeNil

`func (o *RpmRpmRepository) SetCompressionTypeNil(b bool)`

 SetCompressionTypeNil sets the value for CompressionType to be an explicit nil

### UnsetCompressionType
`func (o *RpmRpmRepository) UnsetCompressionType()`

UnsetCompressionType ensures that no value is present for CompressionType, not even an explicit nil
### GetLayout

`func (o *RpmRpmRepository) GetLayout() LayoutEnum`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *RpmRpmRepository) GetLayoutOk() (*LayoutEnum, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *RpmRpmRepository) SetLayout(v LayoutEnum)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *RpmRpmRepository) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### SetLayoutNil

`func (o *RpmRpmRepository) SetLayoutNil(b bool)`

 SetLayoutNil sets the value for Layout to be an explicit nil

### UnsetLayout
`func (o *RpmRpmRepository) UnsetLayout()`

UnsetLayout ensures that no value is present for Layout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


