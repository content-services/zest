# ContainerManifestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**Digest** | **string** | sha256 of the Manifest file | 
**SchemaVersion** | **int64** | Manifest schema version | 
**MediaType** | **string** | Manifest media type of the file | 
**ListedManifests** | **[]string** | Manifests that are referenced by this Manifest List | 
**ConfigBlob** | Pointer to **string** | Blob that contains configuration for this Manifest | [optional] 
**Blobs** | **[]string** | Blobs that are referenced by this Manifest | 
**Annotations** | Pointer to **map[string]interface{}** | Property that contains arbitrary metadata stored inside the image manifest. | [optional] [readonly] 
**Labels** | Pointer to **map[string]interface{}** | Property describing metadata stored inside the image configuration | [optional] [readonly] 
**IsBootable** | Pointer to **bool** | A boolean determining whether users can boot from an image or not.[deprecated] check type field instead | [optional] [default to false]
**IsFlatpak** | Pointer to **bool** | A boolean determining whether the image bundles a Flatpak application.[deprecated] check type field instead | [optional] [default to false]
**Type** | Pointer to **string** | Manifest type (flatpak, bootable, signature, etc.). | [optional] 
**Architecture** | Pointer to **string** | The CPU architecture which the binaries in this image are built to run on. | [optional] 
**Os** | Pointer to **string** | The name of the operating system which the image is built to run on. | [optional] 
**CompressedImageSize** | Pointer to **int64** | Specifies the sum of the sizes, in bytes, of all compressed layers | [optional] 

## Methods

### NewContainerManifestResponse

`func NewContainerManifestResponse(digest string, schemaVersion int64, mediaType string, listedManifests []string, blobs []string, ) *ContainerManifestResponse`

NewContainerManifestResponse instantiates a new ContainerManifestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerManifestResponseWithDefaults

`func NewContainerManifestResponseWithDefaults() *ContainerManifestResponse`

NewContainerManifestResponseWithDefaults instantiates a new ContainerManifestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *ContainerManifestResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ContainerManifestResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ContainerManifestResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ContainerManifestResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *ContainerManifestResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *ContainerManifestResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *ContainerManifestResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *ContainerManifestResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *ContainerManifestResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *ContainerManifestResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *ContainerManifestResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *ContainerManifestResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *ContainerManifestResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *ContainerManifestResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *ContainerManifestResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *ContainerManifestResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetPulpLabels

`func (o *ContainerManifestResponse) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *ContainerManifestResponse) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *ContainerManifestResponse) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *ContainerManifestResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetDigest

`func (o *ContainerManifestResponse) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ContainerManifestResponse) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ContainerManifestResponse) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetSchemaVersion

`func (o *ContainerManifestResponse) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ContainerManifestResponse) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ContainerManifestResponse) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetMediaType

`func (o *ContainerManifestResponse) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *ContainerManifestResponse) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *ContainerManifestResponse) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetListedManifests

`func (o *ContainerManifestResponse) GetListedManifests() []string`

GetListedManifests returns the ListedManifests field if non-nil, zero value otherwise.

### GetListedManifestsOk

`func (o *ContainerManifestResponse) GetListedManifestsOk() (*[]string, bool)`

GetListedManifestsOk returns a tuple with the ListedManifests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedManifests

`func (o *ContainerManifestResponse) SetListedManifests(v []string)`

SetListedManifests sets ListedManifests field to given value.


### GetConfigBlob

`func (o *ContainerManifestResponse) GetConfigBlob() string`

GetConfigBlob returns the ConfigBlob field if non-nil, zero value otherwise.

### GetConfigBlobOk

`func (o *ContainerManifestResponse) GetConfigBlobOk() (*string, bool)`

GetConfigBlobOk returns a tuple with the ConfigBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigBlob

`func (o *ContainerManifestResponse) SetConfigBlob(v string)`

SetConfigBlob sets ConfigBlob field to given value.

### HasConfigBlob

`func (o *ContainerManifestResponse) HasConfigBlob() bool`

HasConfigBlob returns a boolean if a field has been set.

### GetBlobs

`func (o *ContainerManifestResponse) GetBlobs() []string`

GetBlobs returns the Blobs field if non-nil, zero value otherwise.

### GetBlobsOk

`func (o *ContainerManifestResponse) GetBlobsOk() (*[]string, bool)`

GetBlobsOk returns a tuple with the Blobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobs

`func (o *ContainerManifestResponse) SetBlobs(v []string)`

SetBlobs sets Blobs field to given value.


### GetAnnotations

`func (o *ContainerManifestResponse) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ContainerManifestResponse) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ContainerManifestResponse) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ContainerManifestResponse) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetLabels

`func (o *ContainerManifestResponse) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ContainerManifestResponse) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ContainerManifestResponse) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ContainerManifestResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetIsBootable

`func (o *ContainerManifestResponse) GetIsBootable() bool`

GetIsBootable returns the IsBootable field if non-nil, zero value otherwise.

### GetIsBootableOk

`func (o *ContainerManifestResponse) GetIsBootableOk() (*bool, bool)`

GetIsBootableOk returns a tuple with the IsBootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootable

`func (o *ContainerManifestResponse) SetIsBootable(v bool)`

SetIsBootable sets IsBootable field to given value.

### HasIsBootable

`func (o *ContainerManifestResponse) HasIsBootable() bool`

HasIsBootable returns a boolean if a field has been set.

### GetIsFlatpak

`func (o *ContainerManifestResponse) GetIsFlatpak() bool`

GetIsFlatpak returns the IsFlatpak field if non-nil, zero value otherwise.

### GetIsFlatpakOk

`func (o *ContainerManifestResponse) GetIsFlatpakOk() (*bool, bool)`

GetIsFlatpakOk returns a tuple with the IsFlatpak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlatpak

`func (o *ContainerManifestResponse) SetIsFlatpak(v bool)`

SetIsFlatpak sets IsFlatpak field to given value.

### HasIsFlatpak

`func (o *ContainerManifestResponse) HasIsFlatpak() bool`

HasIsFlatpak returns a boolean if a field has been set.

### GetType

`func (o *ContainerManifestResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContainerManifestResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContainerManifestResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContainerManifestResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetArchitecture

`func (o *ContainerManifestResponse) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ContainerManifestResponse) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ContainerManifestResponse) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *ContainerManifestResponse) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetOs

`func (o *ContainerManifestResponse) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ContainerManifestResponse) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ContainerManifestResponse) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *ContainerManifestResponse) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetCompressedImageSize

`func (o *ContainerManifestResponse) GetCompressedImageSize() int64`

GetCompressedImageSize returns the CompressedImageSize field if non-nil, zero value otherwise.

### GetCompressedImageSizeOk

`func (o *ContainerManifestResponse) GetCompressedImageSizeOk() (*int64, bool)`

GetCompressedImageSizeOk returns a tuple with the CompressedImageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedImageSize

`func (o *ContainerManifestResponse) SetCompressedImageSize(v int64)`

SetCompressedImageSize sets CompressedImageSize field to given value.

### HasCompressedImageSize

`func (o *ContainerManifestResponse) HasCompressedImageSize() bool`

HasCompressedImageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


