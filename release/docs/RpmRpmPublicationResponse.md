# RpmRpmPublicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**RepositoryVersion** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** | A URI of the repository to be published. | [optional] 
**ChecksumType** | Pointer to [**PackageChecksumTypeEnum**](PackageChecksumTypeEnum.md) | The preferred checksum type used during repo publishes.* &#x60;unknown&#x60; - unknown* &#x60;md5&#x60; - md5* &#x60;sha1&#x60; - sha1* &#x60;sha224&#x60; - sha224* &#x60;sha256&#x60; - sha256* &#x60;sha384&#x60; - sha384* &#x60;sha512&#x60; - sha512 | [optional] 
**MetadataChecksumType** | Pointer to [**PackageChecksumTypeEnum**](PackageChecksumTypeEnum.md) | DEPRECATED: The checksum type for metadata.* &#x60;unknown&#x60; - unknown* &#x60;md5&#x60; - md5* &#x60;sha1&#x60; - sha1* &#x60;sha224&#x60; - sha224* &#x60;sha256&#x60; - sha256* &#x60;sha384&#x60; - sha384* &#x60;sha512&#x60; - sha512 | [optional] 
**PackageChecksumType** | Pointer to [**PackageChecksumTypeEnum**](PackageChecksumTypeEnum.md) | DEPRECATED: The checksum type for packages.* &#x60;unknown&#x60; - unknown* &#x60;md5&#x60; - md5* &#x60;sha1&#x60; - sha1* &#x60;sha224&#x60; - sha224* &#x60;sha256&#x60; - sha256* &#x60;sha384&#x60; - sha384* &#x60;sha512&#x60; - sha512 | [optional] 
**Gpgcheck** | Pointer to **NullableInt64** | DEPRECATED: An option specifying whether a client should perform a GPG signature check on packages. | [optional] 
**RepoGpgcheck** | Pointer to **NullableInt64** | DEPRECATED: An option specifying whether a client should perform a GPG signature check on the repodata. | [optional] 
**SqliteMetadata** | Pointer to **bool** | REMOVED: An option specifying whether Pulp should generate SQLite metadata. Not operation since pulp_rpm 3.25.0 release | [optional] [readonly] [default to false]
**RepoConfig** | Pointer to **map[string]interface{}** | A JSON document describing config.repo file | [optional] 
**CompressionType** | Pointer to [**CompressionTypeEnum**](CompressionTypeEnum.md) | The compression type to use for metadata files.* &#x60;zstd&#x60; - zstd* &#x60;gz&#x60; - gz | [optional] 

## Methods

### NewRpmRpmPublicationResponse

`func NewRpmRpmPublicationResponse() *RpmRpmPublicationResponse`

NewRpmRpmPublicationResponse instantiates a new RpmRpmPublicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmRpmPublicationResponseWithDefaults

`func NewRpmRpmPublicationResponseWithDefaults() *RpmRpmPublicationResponse`

NewRpmRpmPublicationResponseWithDefaults instantiates a new RpmRpmPublicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *RpmRpmPublicationResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *RpmRpmPublicationResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *RpmRpmPublicationResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *RpmRpmPublicationResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *RpmRpmPublicationResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *RpmRpmPublicationResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *RpmRpmPublicationResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *RpmRpmPublicationResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *RpmRpmPublicationResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *RpmRpmPublicationResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *RpmRpmPublicationResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *RpmRpmPublicationResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *RpmRpmPublicationResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *RpmRpmPublicationResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *RpmRpmPublicationResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *RpmRpmPublicationResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetRepositoryVersion

`func (o *RpmRpmPublicationResponse) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *RpmRpmPublicationResponse) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *RpmRpmPublicationResponse) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *RpmRpmPublicationResponse) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### GetRepository

`func (o *RpmRpmPublicationResponse) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RpmRpmPublicationResponse) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RpmRpmPublicationResponse) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *RpmRpmPublicationResponse) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetChecksumType

`func (o *RpmRpmPublicationResponse) GetChecksumType() PackageChecksumTypeEnum`

GetChecksumType returns the ChecksumType field if non-nil, zero value otherwise.

### GetChecksumTypeOk

`func (o *RpmRpmPublicationResponse) GetChecksumTypeOk() (*PackageChecksumTypeEnum, bool)`

GetChecksumTypeOk returns a tuple with the ChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumType

`func (o *RpmRpmPublicationResponse) SetChecksumType(v PackageChecksumTypeEnum)`

SetChecksumType sets ChecksumType field to given value.

### HasChecksumType

`func (o *RpmRpmPublicationResponse) HasChecksumType() bool`

HasChecksumType returns a boolean if a field has been set.

### GetMetadataChecksumType

`func (o *RpmRpmPublicationResponse) GetMetadataChecksumType() PackageChecksumTypeEnum`

GetMetadataChecksumType returns the MetadataChecksumType field if non-nil, zero value otherwise.

### GetMetadataChecksumTypeOk

`func (o *RpmRpmPublicationResponse) GetMetadataChecksumTypeOk() (*PackageChecksumTypeEnum, bool)`

GetMetadataChecksumTypeOk returns a tuple with the MetadataChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataChecksumType

`func (o *RpmRpmPublicationResponse) SetMetadataChecksumType(v PackageChecksumTypeEnum)`

SetMetadataChecksumType sets MetadataChecksumType field to given value.

### HasMetadataChecksumType

`func (o *RpmRpmPublicationResponse) HasMetadataChecksumType() bool`

HasMetadataChecksumType returns a boolean if a field has been set.

### GetPackageChecksumType

`func (o *RpmRpmPublicationResponse) GetPackageChecksumType() PackageChecksumTypeEnum`

GetPackageChecksumType returns the PackageChecksumType field if non-nil, zero value otherwise.

### GetPackageChecksumTypeOk

`func (o *RpmRpmPublicationResponse) GetPackageChecksumTypeOk() (*PackageChecksumTypeEnum, bool)`

GetPackageChecksumTypeOk returns a tuple with the PackageChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageChecksumType

`func (o *RpmRpmPublicationResponse) SetPackageChecksumType(v PackageChecksumTypeEnum)`

SetPackageChecksumType sets PackageChecksumType field to given value.

### HasPackageChecksumType

`func (o *RpmRpmPublicationResponse) HasPackageChecksumType() bool`

HasPackageChecksumType returns a boolean if a field has been set.

### GetGpgcheck

`func (o *RpmRpmPublicationResponse) GetGpgcheck() int64`

GetGpgcheck returns the Gpgcheck field if non-nil, zero value otherwise.

### GetGpgcheckOk

`func (o *RpmRpmPublicationResponse) GetGpgcheckOk() (*int64, bool)`

GetGpgcheckOk returns a tuple with the Gpgcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgcheck

`func (o *RpmRpmPublicationResponse) SetGpgcheck(v int64)`

SetGpgcheck sets Gpgcheck field to given value.

### HasGpgcheck

`func (o *RpmRpmPublicationResponse) HasGpgcheck() bool`

HasGpgcheck returns a boolean if a field has been set.

### SetGpgcheckNil

`func (o *RpmRpmPublicationResponse) SetGpgcheckNil(b bool)`

 SetGpgcheckNil sets the value for Gpgcheck to be an explicit nil

### UnsetGpgcheck
`func (o *RpmRpmPublicationResponse) UnsetGpgcheck()`

UnsetGpgcheck ensures that no value is present for Gpgcheck, not even an explicit nil
### GetRepoGpgcheck

`func (o *RpmRpmPublicationResponse) GetRepoGpgcheck() int64`

GetRepoGpgcheck returns the RepoGpgcheck field if non-nil, zero value otherwise.

### GetRepoGpgcheckOk

`func (o *RpmRpmPublicationResponse) GetRepoGpgcheckOk() (*int64, bool)`

GetRepoGpgcheckOk returns a tuple with the RepoGpgcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoGpgcheck

`func (o *RpmRpmPublicationResponse) SetRepoGpgcheck(v int64)`

SetRepoGpgcheck sets RepoGpgcheck field to given value.

### HasRepoGpgcheck

`func (o *RpmRpmPublicationResponse) HasRepoGpgcheck() bool`

HasRepoGpgcheck returns a boolean if a field has been set.

### SetRepoGpgcheckNil

`func (o *RpmRpmPublicationResponse) SetRepoGpgcheckNil(b bool)`

 SetRepoGpgcheckNil sets the value for RepoGpgcheck to be an explicit nil

### UnsetRepoGpgcheck
`func (o *RpmRpmPublicationResponse) UnsetRepoGpgcheck()`

UnsetRepoGpgcheck ensures that no value is present for RepoGpgcheck, not even an explicit nil
### GetSqliteMetadata

`func (o *RpmRpmPublicationResponse) GetSqliteMetadata() bool`

GetSqliteMetadata returns the SqliteMetadata field if non-nil, zero value otherwise.

### GetSqliteMetadataOk

`func (o *RpmRpmPublicationResponse) GetSqliteMetadataOk() (*bool, bool)`

GetSqliteMetadataOk returns a tuple with the SqliteMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqliteMetadata

`func (o *RpmRpmPublicationResponse) SetSqliteMetadata(v bool)`

SetSqliteMetadata sets SqliteMetadata field to given value.

### HasSqliteMetadata

`func (o *RpmRpmPublicationResponse) HasSqliteMetadata() bool`

HasSqliteMetadata returns a boolean if a field has been set.

### GetRepoConfig

`func (o *RpmRpmPublicationResponse) GetRepoConfig() map[string]interface{}`

GetRepoConfig returns the RepoConfig field if non-nil, zero value otherwise.

### GetRepoConfigOk

`func (o *RpmRpmPublicationResponse) GetRepoConfigOk() (*map[string]interface{}, bool)`

GetRepoConfigOk returns a tuple with the RepoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoConfig

`func (o *RpmRpmPublicationResponse) SetRepoConfig(v map[string]interface{})`

SetRepoConfig sets RepoConfig field to given value.

### HasRepoConfig

`func (o *RpmRpmPublicationResponse) HasRepoConfig() bool`

HasRepoConfig returns a boolean if a field has been set.

### GetCompressionType

`func (o *RpmRpmPublicationResponse) GetCompressionType() CompressionTypeEnum`

GetCompressionType returns the CompressionType field if non-nil, zero value otherwise.

### GetCompressionTypeOk

`func (o *RpmRpmPublicationResponse) GetCompressionTypeOk() (*CompressionTypeEnum, bool)`

GetCompressionTypeOk returns a tuple with the CompressionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionType

`func (o *RpmRpmPublicationResponse) SetCompressionType(v CompressionTypeEnum)`

SetCompressionType sets CompressionType field to given value.

### HasCompressionType

`func (o *RpmRpmPublicationResponse) HasCompressionType() bool`

HasCompressionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


