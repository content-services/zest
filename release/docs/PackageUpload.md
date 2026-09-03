# PackageUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | ***os.File** | A Python package release file to upload to the index. | 
**Action** | Pointer to **string** | Defaults to &#x60;file_upload&#x60;, don&#39;t change it or request will fail! | [optional] [default to "file_upload"]
**Sha256Digest** | **string** | SHA256 of package to validate upload integrity. | 
**ProtocolVersion** | Pointer to [**ProtocolVersionEnum**](ProtocolVersionEnum.md) | Protocol version to use for the upload. Only version 1 is supported.* &#x60;1&#x60; - 1 | [optional] [default to PROTOCOLVERSIONENUM__1]
**Filetype** | Pointer to [**FiletypeEnum**](FiletypeEnum.md) | Type of artifact to upload.* &#x60;bdist_wheel&#x60; - bdist_wheel* &#x60;sdist&#x60; - sdist | [optional] 
**MetadataVersion** | Pointer to [**MetadataVersionEnum**](MetadataVersionEnum.md) | Metadata version of the uploaded package.* &#x60;1.0&#x60; - 1.0* &#x60;1.1&#x60; - 1.1* &#x60;1.2&#x60; - 1.2* &#x60;2.0&#x60; - 2.0* &#x60;2.1&#x60; - 2.1* &#x60;2.2&#x60; - 2.2* &#x60;2.3&#x60; - 2.3* &#x60;2.4&#x60; - 2.4* &#x60;2.5&#x60; - 2.5 | [optional] 
**Attestations** | Pointer to **interface{}** | A JSON list containing attestations for the package. | [optional] 

## Methods

### NewPackageUpload

`func NewPackageUpload(content *os.File, sha256Digest string, ) *PackageUpload`

NewPackageUpload instantiates a new PackageUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageUploadWithDefaults

`func NewPackageUploadWithDefaults() *PackageUpload`

NewPackageUploadWithDefaults instantiates a new PackageUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PackageUpload) GetContent() *os.File`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PackageUpload) GetContentOk() (**os.File, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PackageUpload) SetContent(v *os.File)`

SetContent sets Content field to given value.


### GetAction

`func (o *PackageUpload) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PackageUpload) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PackageUpload) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PackageUpload) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSha256Digest

`func (o *PackageUpload) GetSha256Digest() string`

GetSha256Digest returns the Sha256Digest field if non-nil, zero value otherwise.

### GetSha256DigestOk

`func (o *PackageUpload) GetSha256DigestOk() (*string, bool)`

GetSha256DigestOk returns a tuple with the Sha256Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Digest

`func (o *PackageUpload) SetSha256Digest(v string)`

SetSha256Digest sets Sha256Digest field to given value.


### GetProtocolVersion

`func (o *PackageUpload) GetProtocolVersion() ProtocolVersionEnum`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *PackageUpload) GetProtocolVersionOk() (*ProtocolVersionEnum, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *PackageUpload) SetProtocolVersion(v ProtocolVersionEnum)`

SetProtocolVersion sets ProtocolVersion field to given value.

### HasProtocolVersion

`func (o *PackageUpload) HasProtocolVersion() bool`

HasProtocolVersion returns a boolean if a field has been set.

### GetFiletype

`func (o *PackageUpload) GetFiletype() FiletypeEnum`

GetFiletype returns the Filetype field if non-nil, zero value otherwise.

### GetFiletypeOk

`func (o *PackageUpload) GetFiletypeOk() (*FiletypeEnum, bool)`

GetFiletypeOk returns a tuple with the Filetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiletype

`func (o *PackageUpload) SetFiletype(v FiletypeEnum)`

SetFiletype sets Filetype field to given value.

### HasFiletype

`func (o *PackageUpload) HasFiletype() bool`

HasFiletype returns a boolean if a field has been set.

### GetMetadataVersion

`func (o *PackageUpload) GetMetadataVersion() MetadataVersionEnum`

GetMetadataVersion returns the MetadataVersion field if non-nil, zero value otherwise.

### GetMetadataVersionOk

`func (o *PackageUpload) GetMetadataVersionOk() (*MetadataVersionEnum, bool)`

GetMetadataVersionOk returns a tuple with the MetadataVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataVersion

`func (o *PackageUpload) SetMetadataVersion(v MetadataVersionEnum)`

SetMetadataVersion sets MetadataVersion field to given value.

### HasMetadataVersion

`func (o *PackageUpload) HasMetadataVersion() bool`

HasMetadataVersion returns a boolean if a field has been set.

### GetAttestations

`func (o *PackageUpload) GetAttestations() interface{}`

GetAttestations returns the Attestations field if non-nil, zero value otherwise.

### GetAttestationsOk

`func (o *PackageUpload) GetAttestationsOk() (*interface{}, bool)`

GetAttestationsOk returns a tuple with the Attestations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestations

`func (o *PackageUpload) SetAttestations(v interface{})`

SetAttestations sets Attestations field to given value.

### HasAttestations

`func (o *PackageUpload) HasAttestations() bool`

HasAttestations returns a boolean if a field has been set.

### SetAttestationsNil

`func (o *PackageUpload) SetAttestationsNil(b bool)`

 SetAttestationsNil sets the value for Attestations to be an explicit nil

### UnsetAttestations
`func (o *PackageUpload) UnsetAttestations()`

UnsetAttestations ensures that no value is present for Attestations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


