# RpmRpmPublication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepositoryVersion** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** | A URI of the repository to be published. | [optional] 
**Checkpoint** | Pointer to **bool** |  | [optional] 
**ChecksumType** | Pointer to [**PackageChecksumTypeEnum**](PackageChecksumTypeEnum.md) | The preferred checksum type used during repo publishes.* &#x60;unknown&#x60; - unknown* &#x60;md5&#x60; - md5* &#x60;sha1&#x60; - sha1* &#x60;sha224&#x60; - sha224* &#x60;sha256&#x60; - sha256* &#x60;sha384&#x60; - sha384* &#x60;sha512&#x60; - sha512 | [optional] 
**RepoConfig** | Pointer to **interface{}** | A JSON document describing the config.repo file Pulp should generate for this repo | [optional] 
**CompressionType** | Pointer to [**CompressionTypeEnum**](CompressionTypeEnum.md) | The compression type to use for metadata files.* &#x60;zstd&#x60; - zstd* &#x60;gz&#x60; - gz | [optional] 
**Layout** | Pointer to [**NullableLayoutEnum**](LayoutEnum.md) | How to layout the packages within the published repository.* &#x60;nested_alphabetically&#x60; - nested_alphabetically* &#x60;flat&#x60; - flat | [optional] 

## Methods

### NewRpmRpmPublication

`func NewRpmRpmPublication() *RpmRpmPublication`

NewRpmRpmPublication instantiates a new RpmRpmPublication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmRpmPublicationWithDefaults

`func NewRpmRpmPublicationWithDefaults() *RpmRpmPublication`

NewRpmRpmPublicationWithDefaults instantiates a new RpmRpmPublication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositoryVersion

`func (o *RpmRpmPublication) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *RpmRpmPublication) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *RpmRpmPublication) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *RpmRpmPublication) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### GetRepository

`func (o *RpmRpmPublication) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RpmRpmPublication) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RpmRpmPublication) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *RpmRpmPublication) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetCheckpoint

`func (o *RpmRpmPublication) GetCheckpoint() bool`

GetCheckpoint returns the Checkpoint field if non-nil, zero value otherwise.

### GetCheckpointOk

`func (o *RpmRpmPublication) GetCheckpointOk() (*bool, bool)`

GetCheckpointOk returns a tuple with the Checkpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpoint

`func (o *RpmRpmPublication) SetCheckpoint(v bool)`

SetCheckpoint sets Checkpoint field to given value.

### HasCheckpoint

`func (o *RpmRpmPublication) HasCheckpoint() bool`

HasCheckpoint returns a boolean if a field has been set.

### GetChecksumType

`func (o *RpmRpmPublication) GetChecksumType() PackageChecksumTypeEnum`

GetChecksumType returns the ChecksumType field if non-nil, zero value otherwise.

### GetChecksumTypeOk

`func (o *RpmRpmPublication) GetChecksumTypeOk() (*PackageChecksumTypeEnum, bool)`

GetChecksumTypeOk returns a tuple with the ChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumType

`func (o *RpmRpmPublication) SetChecksumType(v PackageChecksumTypeEnum)`

SetChecksumType sets ChecksumType field to given value.

### HasChecksumType

`func (o *RpmRpmPublication) HasChecksumType() bool`

HasChecksumType returns a boolean if a field has been set.

### GetRepoConfig

`func (o *RpmRpmPublication) GetRepoConfig() interface{}`

GetRepoConfig returns the RepoConfig field if non-nil, zero value otherwise.

### GetRepoConfigOk

`func (o *RpmRpmPublication) GetRepoConfigOk() (*interface{}, bool)`

GetRepoConfigOk returns a tuple with the RepoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoConfig

`func (o *RpmRpmPublication) SetRepoConfig(v interface{})`

SetRepoConfig sets RepoConfig field to given value.

### HasRepoConfig

`func (o *RpmRpmPublication) HasRepoConfig() bool`

HasRepoConfig returns a boolean if a field has been set.

### SetRepoConfigNil

`func (o *RpmRpmPublication) SetRepoConfigNil(b bool)`

 SetRepoConfigNil sets the value for RepoConfig to be an explicit nil

### UnsetRepoConfig
`func (o *RpmRpmPublication) UnsetRepoConfig()`

UnsetRepoConfig ensures that no value is present for RepoConfig, not even an explicit nil
### GetCompressionType

`func (o *RpmRpmPublication) GetCompressionType() CompressionTypeEnum`

GetCompressionType returns the CompressionType field if non-nil, zero value otherwise.

### GetCompressionTypeOk

`func (o *RpmRpmPublication) GetCompressionTypeOk() (*CompressionTypeEnum, bool)`

GetCompressionTypeOk returns a tuple with the CompressionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionType

`func (o *RpmRpmPublication) SetCompressionType(v CompressionTypeEnum)`

SetCompressionType sets CompressionType field to given value.

### HasCompressionType

`func (o *RpmRpmPublication) HasCompressionType() bool`

HasCompressionType returns a boolean if a field has been set.

### GetLayout

`func (o *RpmRpmPublication) GetLayout() LayoutEnum`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *RpmRpmPublication) GetLayoutOk() (*LayoutEnum, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *RpmRpmPublication) SetLayout(v LayoutEnum)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *RpmRpmPublication) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### SetLayoutNil

`func (o *RpmRpmPublication) SetLayoutNil(b bool)`

 SetLayoutNil sets the value for Layout to be an explicit nil

### UnsetLayout
`func (o *RpmRpmPublication) UnsetLayout()`

UnsetLayout ensures that no value is present for Layout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


