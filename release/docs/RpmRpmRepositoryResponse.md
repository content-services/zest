# RpmRpmRepositoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**VersionsHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**LatestVersionHref** | Pointer to **string** |  | [optional] [readonly] 
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
**MetadataChecksumType** | Pointer to [**NullablePackageChecksumTypeEnum**](PackageChecksumTypeEnum.md) | DEPRECATED: use CHECKSUM_TYPE instead.* &#x60;unknown&#x60; - unknown* &#x60;md5&#x60; - md5* &#x60;sha1&#x60; - sha1* &#x60;sha224&#x60; - sha224* &#x60;sha256&#x60; - sha256* &#x60;sha384&#x60; - sha384* &#x60;sha512&#x60; - sha512 | [optional] 
**PackageChecksumType** | Pointer to [**NullablePackageChecksumTypeEnum**](PackageChecksumTypeEnum.md) | DEPRECATED: use CHECKSUM_TYPE instead.* &#x60;unknown&#x60; - unknown* &#x60;md5&#x60; - md5* &#x60;sha1&#x60; - sha1* &#x60;sha224&#x60; - sha224* &#x60;sha256&#x60; - sha256* &#x60;sha384&#x60; - sha384* &#x60;sha512&#x60; - sha512 | [optional] 
**Gpgcheck** | Pointer to **NullableInt64** | DEPRECATED: An option specifying whether a client should perform a GPG signature check on packages. | [optional] 
**RepoGpgcheck** | Pointer to **NullableInt64** | DEPRECATED: An option specifying whether a client should perform a GPG signature check on the repodata. | [optional] 
**SqliteMetadata** | Pointer to **bool** | REMOVED: An option specifying whether Pulp should generate SQLite metadata. Not operation since pulp_rpm 3.25.0 release | [optional] [readonly] [default to false]
**RepoConfig** | Pointer to **interface{}** | A JSON document describing config.repo file | [optional] 
**CompressionType** | Pointer to [**NullableCompressionTypeEnum**](CompressionTypeEnum.md) | The compression type to use for metadata files.* &#x60;zstd&#x60; - zstd* &#x60;gz&#x60; - gz | [optional] 
**Layout** | Pointer to [**NullableLayoutEnum**](LayoutEnum.md) | How to layout the packages within the published repository.* &#x60;nested_alphabetically&#x60; - nested_alphabetically* &#x60;flat&#x60; - flat | [optional] 

## Methods

### NewRpmRpmRepositoryResponse

`func NewRpmRpmRepositoryResponse(name string, ) *RpmRpmRepositoryResponse`

NewRpmRpmRepositoryResponse instantiates a new RpmRpmRepositoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmRpmRepositoryResponseWithDefaults

`func NewRpmRpmRepositoryResponseWithDefaults() *RpmRpmRepositoryResponse`

NewRpmRpmRepositoryResponseWithDefaults instantiates a new RpmRpmRepositoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *RpmRpmRepositoryResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *RpmRpmRepositoryResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *RpmRpmRepositoryResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *RpmRpmRepositoryResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *RpmRpmRepositoryResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *RpmRpmRepositoryResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *RpmRpmRepositoryResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *RpmRpmRepositoryResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *RpmRpmRepositoryResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *RpmRpmRepositoryResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *RpmRpmRepositoryResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *RpmRpmRepositoryResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *RpmRpmRepositoryResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *RpmRpmRepositoryResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *RpmRpmRepositoryResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *RpmRpmRepositoryResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetVersionsHref

`func (o *RpmRpmRepositoryResponse) GetVersionsHref() string`

GetVersionsHref returns the VersionsHref field if non-nil, zero value otherwise.

### GetVersionsHrefOk

`func (o *RpmRpmRepositoryResponse) GetVersionsHrefOk() (*string, bool)`

GetVersionsHrefOk returns a tuple with the VersionsHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionsHref

`func (o *RpmRpmRepositoryResponse) SetVersionsHref(v string)`

SetVersionsHref sets VersionsHref field to given value.

### HasVersionsHref

`func (o *RpmRpmRepositoryResponse) HasVersionsHref() bool`

HasVersionsHref returns a boolean if a field has been set.

### GetPulpLabels

`func (o *RpmRpmRepositoryResponse) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *RpmRpmRepositoryResponse) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *RpmRpmRepositoryResponse) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *RpmRpmRepositoryResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetLatestVersionHref

`func (o *RpmRpmRepositoryResponse) GetLatestVersionHref() string`

GetLatestVersionHref returns the LatestVersionHref field if non-nil, zero value otherwise.

### GetLatestVersionHrefOk

`func (o *RpmRpmRepositoryResponse) GetLatestVersionHrefOk() (*string, bool)`

GetLatestVersionHrefOk returns a tuple with the LatestVersionHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersionHref

`func (o *RpmRpmRepositoryResponse) SetLatestVersionHref(v string)`

SetLatestVersionHref sets LatestVersionHref field to given value.

### HasLatestVersionHref

`func (o *RpmRpmRepositoryResponse) HasLatestVersionHref() bool`

HasLatestVersionHref returns a boolean if a field has been set.

### GetName

`func (o *RpmRpmRepositoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RpmRpmRepositoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RpmRpmRepositoryResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RpmRpmRepositoryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RpmRpmRepositoryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RpmRpmRepositoryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RpmRpmRepositoryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RpmRpmRepositoryResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RpmRpmRepositoryResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *RpmRpmRepositoryResponse) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *RpmRpmRepositoryResponse) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *RpmRpmRepositoryResponse) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *RpmRpmRepositoryResponse) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *RpmRpmRepositoryResponse) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *RpmRpmRepositoryResponse) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRemote

`func (o *RpmRpmRepositoryResponse) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *RpmRpmRepositoryResponse) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *RpmRpmRepositoryResponse) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *RpmRpmRepositoryResponse) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *RpmRpmRepositoryResponse) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *RpmRpmRepositoryResponse) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetAutopublish

`func (o *RpmRpmRepositoryResponse) GetAutopublish() bool`

GetAutopublish returns the Autopublish field if non-nil, zero value otherwise.

### GetAutopublishOk

`func (o *RpmRpmRepositoryResponse) GetAutopublishOk() (*bool, bool)`

GetAutopublishOk returns a tuple with the Autopublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutopublish

`func (o *RpmRpmRepositoryResponse) SetAutopublish(v bool)`

SetAutopublish sets Autopublish field to given value.

### HasAutopublish

`func (o *RpmRpmRepositoryResponse) HasAutopublish() bool`

HasAutopublish returns a boolean if a field has been set.

### GetMetadataSigningService

`func (o *RpmRpmRepositoryResponse) GetMetadataSigningService() string`

GetMetadataSigningService returns the MetadataSigningService field if non-nil, zero value otherwise.

### GetMetadataSigningServiceOk

`func (o *RpmRpmRepositoryResponse) GetMetadataSigningServiceOk() (*string, bool)`

GetMetadataSigningServiceOk returns a tuple with the MetadataSigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSigningService

`func (o *RpmRpmRepositoryResponse) SetMetadataSigningService(v string)`

SetMetadataSigningService sets MetadataSigningService field to given value.

### HasMetadataSigningService

`func (o *RpmRpmRepositoryResponse) HasMetadataSigningService() bool`

HasMetadataSigningService returns a boolean if a field has been set.

### SetMetadataSigningServiceNil

`func (o *RpmRpmRepositoryResponse) SetMetadataSigningServiceNil(b bool)`

 SetMetadataSigningServiceNil sets the value for MetadataSigningService to be an explicit nil

### UnsetMetadataSigningService
`func (o *RpmRpmRepositoryResponse) UnsetMetadataSigningService()`

UnsetMetadataSigningService ensures that no value is present for MetadataSigningService, not even an explicit nil
### GetPackageSigningService

`func (o *RpmRpmRepositoryResponse) GetPackageSigningService() string`

GetPackageSigningService returns the PackageSigningService field if non-nil, zero value otherwise.

### GetPackageSigningServiceOk

`func (o *RpmRpmRepositoryResponse) GetPackageSigningServiceOk() (*string, bool)`

GetPackageSigningServiceOk returns a tuple with the PackageSigningService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSigningService

`func (o *RpmRpmRepositoryResponse) SetPackageSigningService(v string)`

SetPackageSigningService sets PackageSigningService field to given value.

### HasPackageSigningService

`func (o *RpmRpmRepositoryResponse) HasPackageSigningService() bool`

HasPackageSigningService returns a boolean if a field has been set.

### SetPackageSigningServiceNil

`func (o *RpmRpmRepositoryResponse) SetPackageSigningServiceNil(b bool)`

 SetPackageSigningServiceNil sets the value for PackageSigningService to be an explicit nil

### UnsetPackageSigningService
`func (o *RpmRpmRepositoryResponse) UnsetPackageSigningService()`

UnsetPackageSigningService ensures that no value is present for PackageSigningService, not even an explicit nil
### GetPackageSigningFingerprint

`func (o *RpmRpmRepositoryResponse) GetPackageSigningFingerprint() string`

GetPackageSigningFingerprint returns the PackageSigningFingerprint field if non-nil, zero value otherwise.

### GetPackageSigningFingerprintOk

`func (o *RpmRpmRepositoryResponse) GetPackageSigningFingerprintOk() (*string, bool)`

GetPackageSigningFingerprintOk returns a tuple with the PackageSigningFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSigningFingerprint

`func (o *RpmRpmRepositoryResponse) SetPackageSigningFingerprint(v string)`

SetPackageSigningFingerprint sets PackageSigningFingerprint field to given value.

### HasPackageSigningFingerprint

`func (o *RpmRpmRepositoryResponse) HasPackageSigningFingerprint() bool`

HasPackageSigningFingerprint returns a boolean if a field has been set.

### GetRetainPackageVersions

`func (o *RpmRpmRepositoryResponse) GetRetainPackageVersions() int64`

GetRetainPackageVersions returns the RetainPackageVersions field if non-nil, zero value otherwise.

### GetRetainPackageVersionsOk

`func (o *RpmRpmRepositoryResponse) GetRetainPackageVersionsOk() (*int64, bool)`

GetRetainPackageVersionsOk returns a tuple with the RetainPackageVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPackageVersions

`func (o *RpmRpmRepositoryResponse) SetRetainPackageVersions(v int64)`

SetRetainPackageVersions sets RetainPackageVersions field to given value.

### HasRetainPackageVersions

`func (o *RpmRpmRepositoryResponse) HasRetainPackageVersions() bool`

HasRetainPackageVersions returns a boolean if a field has been set.

### GetChecksumType

`func (o *RpmRpmRepositoryResponse) GetChecksumType() PackageChecksumTypeEnum`

GetChecksumType returns the ChecksumType field if non-nil, zero value otherwise.

### GetChecksumTypeOk

`func (o *RpmRpmRepositoryResponse) GetChecksumTypeOk() (*PackageChecksumTypeEnum, bool)`

GetChecksumTypeOk returns a tuple with the ChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumType

`func (o *RpmRpmRepositoryResponse) SetChecksumType(v PackageChecksumTypeEnum)`

SetChecksumType sets ChecksumType field to given value.

### HasChecksumType

`func (o *RpmRpmRepositoryResponse) HasChecksumType() bool`

HasChecksumType returns a boolean if a field has been set.

### SetChecksumTypeNil

`func (o *RpmRpmRepositoryResponse) SetChecksumTypeNil(b bool)`

 SetChecksumTypeNil sets the value for ChecksumType to be an explicit nil

### UnsetChecksumType
`func (o *RpmRpmRepositoryResponse) UnsetChecksumType()`

UnsetChecksumType ensures that no value is present for ChecksumType, not even an explicit nil
### GetMetadataChecksumType

`func (o *RpmRpmRepositoryResponse) GetMetadataChecksumType() PackageChecksumTypeEnum`

GetMetadataChecksumType returns the MetadataChecksumType field if non-nil, zero value otherwise.

### GetMetadataChecksumTypeOk

`func (o *RpmRpmRepositoryResponse) GetMetadataChecksumTypeOk() (*PackageChecksumTypeEnum, bool)`

GetMetadataChecksumTypeOk returns a tuple with the MetadataChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataChecksumType

`func (o *RpmRpmRepositoryResponse) SetMetadataChecksumType(v PackageChecksumTypeEnum)`

SetMetadataChecksumType sets MetadataChecksumType field to given value.

### HasMetadataChecksumType

`func (o *RpmRpmRepositoryResponse) HasMetadataChecksumType() bool`

HasMetadataChecksumType returns a boolean if a field has been set.

### SetMetadataChecksumTypeNil

`func (o *RpmRpmRepositoryResponse) SetMetadataChecksumTypeNil(b bool)`

 SetMetadataChecksumTypeNil sets the value for MetadataChecksumType to be an explicit nil

### UnsetMetadataChecksumType
`func (o *RpmRpmRepositoryResponse) UnsetMetadataChecksumType()`

UnsetMetadataChecksumType ensures that no value is present for MetadataChecksumType, not even an explicit nil
### GetPackageChecksumType

`func (o *RpmRpmRepositoryResponse) GetPackageChecksumType() PackageChecksumTypeEnum`

GetPackageChecksumType returns the PackageChecksumType field if non-nil, zero value otherwise.

### GetPackageChecksumTypeOk

`func (o *RpmRpmRepositoryResponse) GetPackageChecksumTypeOk() (*PackageChecksumTypeEnum, bool)`

GetPackageChecksumTypeOk returns a tuple with the PackageChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageChecksumType

`func (o *RpmRpmRepositoryResponse) SetPackageChecksumType(v PackageChecksumTypeEnum)`

SetPackageChecksumType sets PackageChecksumType field to given value.

### HasPackageChecksumType

`func (o *RpmRpmRepositoryResponse) HasPackageChecksumType() bool`

HasPackageChecksumType returns a boolean if a field has been set.

### SetPackageChecksumTypeNil

`func (o *RpmRpmRepositoryResponse) SetPackageChecksumTypeNil(b bool)`

 SetPackageChecksumTypeNil sets the value for PackageChecksumType to be an explicit nil

### UnsetPackageChecksumType
`func (o *RpmRpmRepositoryResponse) UnsetPackageChecksumType()`

UnsetPackageChecksumType ensures that no value is present for PackageChecksumType, not even an explicit nil
### GetGpgcheck

`func (o *RpmRpmRepositoryResponse) GetGpgcheck() int64`

GetGpgcheck returns the Gpgcheck field if non-nil, zero value otherwise.

### GetGpgcheckOk

`func (o *RpmRpmRepositoryResponse) GetGpgcheckOk() (*int64, bool)`

GetGpgcheckOk returns a tuple with the Gpgcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgcheck

`func (o *RpmRpmRepositoryResponse) SetGpgcheck(v int64)`

SetGpgcheck sets Gpgcheck field to given value.

### HasGpgcheck

`func (o *RpmRpmRepositoryResponse) HasGpgcheck() bool`

HasGpgcheck returns a boolean if a field has been set.

### SetGpgcheckNil

`func (o *RpmRpmRepositoryResponse) SetGpgcheckNil(b bool)`

 SetGpgcheckNil sets the value for Gpgcheck to be an explicit nil

### UnsetGpgcheck
`func (o *RpmRpmRepositoryResponse) UnsetGpgcheck()`

UnsetGpgcheck ensures that no value is present for Gpgcheck, not even an explicit nil
### GetRepoGpgcheck

`func (o *RpmRpmRepositoryResponse) GetRepoGpgcheck() int64`

GetRepoGpgcheck returns the RepoGpgcheck field if non-nil, zero value otherwise.

### GetRepoGpgcheckOk

`func (o *RpmRpmRepositoryResponse) GetRepoGpgcheckOk() (*int64, bool)`

GetRepoGpgcheckOk returns a tuple with the RepoGpgcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoGpgcheck

`func (o *RpmRpmRepositoryResponse) SetRepoGpgcheck(v int64)`

SetRepoGpgcheck sets RepoGpgcheck field to given value.

### HasRepoGpgcheck

`func (o *RpmRpmRepositoryResponse) HasRepoGpgcheck() bool`

HasRepoGpgcheck returns a boolean if a field has been set.

### SetRepoGpgcheckNil

`func (o *RpmRpmRepositoryResponse) SetRepoGpgcheckNil(b bool)`

 SetRepoGpgcheckNil sets the value for RepoGpgcheck to be an explicit nil

### UnsetRepoGpgcheck
`func (o *RpmRpmRepositoryResponse) UnsetRepoGpgcheck()`

UnsetRepoGpgcheck ensures that no value is present for RepoGpgcheck, not even an explicit nil
### GetSqliteMetadata

`func (o *RpmRpmRepositoryResponse) GetSqliteMetadata() bool`

GetSqliteMetadata returns the SqliteMetadata field if non-nil, zero value otherwise.

### GetSqliteMetadataOk

`func (o *RpmRpmRepositoryResponse) GetSqliteMetadataOk() (*bool, bool)`

GetSqliteMetadataOk returns a tuple with the SqliteMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqliteMetadata

`func (o *RpmRpmRepositoryResponse) SetSqliteMetadata(v bool)`

SetSqliteMetadata sets SqliteMetadata field to given value.

### HasSqliteMetadata

`func (o *RpmRpmRepositoryResponse) HasSqliteMetadata() bool`

HasSqliteMetadata returns a boolean if a field has been set.

### GetRepoConfig

`func (o *RpmRpmRepositoryResponse) GetRepoConfig() interface{}`

GetRepoConfig returns the RepoConfig field if non-nil, zero value otherwise.

### GetRepoConfigOk

`func (o *RpmRpmRepositoryResponse) GetRepoConfigOk() (*interface{}, bool)`

GetRepoConfigOk returns a tuple with the RepoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoConfig

`func (o *RpmRpmRepositoryResponse) SetRepoConfig(v interface{})`

SetRepoConfig sets RepoConfig field to given value.

### HasRepoConfig

`func (o *RpmRpmRepositoryResponse) HasRepoConfig() bool`

HasRepoConfig returns a boolean if a field has been set.

### SetRepoConfigNil

`func (o *RpmRpmRepositoryResponse) SetRepoConfigNil(b bool)`

 SetRepoConfigNil sets the value for RepoConfig to be an explicit nil

### UnsetRepoConfig
`func (o *RpmRpmRepositoryResponse) UnsetRepoConfig()`

UnsetRepoConfig ensures that no value is present for RepoConfig, not even an explicit nil
### GetCompressionType

`func (o *RpmRpmRepositoryResponse) GetCompressionType() CompressionTypeEnum`

GetCompressionType returns the CompressionType field if non-nil, zero value otherwise.

### GetCompressionTypeOk

`func (o *RpmRpmRepositoryResponse) GetCompressionTypeOk() (*CompressionTypeEnum, bool)`

GetCompressionTypeOk returns a tuple with the CompressionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionType

`func (o *RpmRpmRepositoryResponse) SetCompressionType(v CompressionTypeEnum)`

SetCompressionType sets CompressionType field to given value.

### HasCompressionType

`func (o *RpmRpmRepositoryResponse) HasCompressionType() bool`

HasCompressionType returns a boolean if a field has been set.

### SetCompressionTypeNil

`func (o *RpmRpmRepositoryResponse) SetCompressionTypeNil(b bool)`

 SetCompressionTypeNil sets the value for CompressionType to be an explicit nil

### UnsetCompressionType
`func (o *RpmRpmRepositoryResponse) UnsetCompressionType()`

UnsetCompressionType ensures that no value is present for CompressionType, not even an explicit nil
### GetLayout

`func (o *RpmRpmRepositoryResponse) GetLayout() LayoutEnum`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *RpmRpmRepositoryResponse) GetLayoutOk() (*LayoutEnum, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *RpmRpmRepositoryResponse) SetLayout(v LayoutEnum)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *RpmRpmRepositoryResponse) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### SetLayoutNil

`func (o *RpmRpmRepositoryResponse) SetLayoutNil(b bool)`

 SetLayoutNil sets the value for Layout to be an explicit nil

### UnsetLayout
`func (o *RpmRpmRepositoryResponse) UnsetLayout()`

UnsetLayout ensures that no value is present for Layout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


