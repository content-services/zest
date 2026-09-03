# MavenMavenMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**VulnReport** | Pointer to **string** |  | [optional] [readonly] 
**Artifact** | Pointer to **string** | Artifact file representing the physical content | [optional] 
**GroupId** | Pointer to **string** | Group Id of the metadata&#39;s package. | [optional] [readonly] 
**ArtifactId** | Pointer to **string** | Artifact Id of the metadata&#39;s package. | [optional] [readonly] 
**Version** | Pointer to **NullableString** | Version of the metadata&#39;s package. | [optional] [readonly] 
**Filename** | Pointer to **string** | Filename of the metadata. | [optional] [readonly] 
**Sha256** | Pointer to **string** | SHA256 digest of the metadata. | [optional] [readonly] 

## Methods

### NewMavenMavenMetadataResponse

`func NewMavenMavenMetadataResponse() *MavenMavenMetadataResponse`

NewMavenMavenMetadataResponse instantiates a new MavenMavenMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenMavenMetadataResponseWithDefaults

`func NewMavenMavenMetadataResponseWithDefaults() *MavenMavenMetadataResponse`

NewMavenMavenMetadataResponseWithDefaults instantiates a new MavenMavenMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *MavenMavenMetadataResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *MavenMavenMetadataResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *MavenMavenMetadataResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *MavenMavenMetadataResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *MavenMavenMetadataResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *MavenMavenMetadataResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *MavenMavenMetadataResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *MavenMavenMetadataResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *MavenMavenMetadataResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *MavenMavenMetadataResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *MavenMavenMetadataResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *MavenMavenMetadataResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *MavenMavenMetadataResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *MavenMavenMetadataResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *MavenMavenMetadataResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *MavenMavenMetadataResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetPulpLabels

`func (o *MavenMavenMetadataResponse) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *MavenMavenMetadataResponse) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *MavenMavenMetadataResponse) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *MavenMavenMetadataResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetVulnReport

`func (o *MavenMavenMetadataResponse) GetVulnReport() string`

GetVulnReport returns the VulnReport field if non-nil, zero value otherwise.

### GetVulnReportOk

`func (o *MavenMavenMetadataResponse) GetVulnReportOk() (*string, bool)`

GetVulnReportOk returns a tuple with the VulnReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnReport

`func (o *MavenMavenMetadataResponse) SetVulnReport(v string)`

SetVulnReport sets VulnReport field to given value.

### HasVulnReport

`func (o *MavenMavenMetadataResponse) HasVulnReport() bool`

HasVulnReport returns a boolean if a field has been set.

### GetArtifact

`func (o *MavenMavenMetadataResponse) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *MavenMavenMetadataResponse) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *MavenMavenMetadataResponse) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *MavenMavenMetadataResponse) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetGroupId

`func (o *MavenMavenMetadataResponse) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *MavenMavenMetadataResponse) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *MavenMavenMetadataResponse) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *MavenMavenMetadataResponse) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetArtifactId

`func (o *MavenMavenMetadataResponse) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *MavenMavenMetadataResponse) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *MavenMavenMetadataResponse) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *MavenMavenMetadataResponse) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### GetVersion

`func (o *MavenMavenMetadataResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MavenMavenMetadataResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MavenMavenMetadataResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MavenMavenMetadataResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MavenMavenMetadataResponse) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MavenMavenMetadataResponse) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetFilename

`func (o *MavenMavenMetadataResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *MavenMavenMetadataResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *MavenMavenMetadataResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *MavenMavenMetadataResponse) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetSha256

`func (o *MavenMavenMetadataResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *MavenMavenMetadataResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *MavenMavenMetadataResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *MavenMavenMetadataResponse) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


