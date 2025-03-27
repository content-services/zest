# MavenMavenArtifactResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**Artifact** | **string** | Artifact file representing the physical content | 
**GroupId** | Pointer to **string** | Group Id of the artifact&#39;s package. | [optional] [readonly] 
**ArtifactId** | Pointer to **string** | Artifact Id of the artifact&#39;s package. | [optional] [readonly] 
**Version** | Pointer to **string** | Version of the artifact&#39;s package. | [optional] [readonly] 
**Filename** | Pointer to **string** | Filename of the artifact. | [optional] [readonly] 

## Methods

### NewMavenMavenArtifactResponse

`func NewMavenMavenArtifactResponse(artifact string, ) *MavenMavenArtifactResponse`

NewMavenMavenArtifactResponse instantiates a new MavenMavenArtifactResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenMavenArtifactResponseWithDefaults

`func NewMavenMavenArtifactResponseWithDefaults() *MavenMavenArtifactResponse`

NewMavenMavenArtifactResponseWithDefaults instantiates a new MavenMavenArtifactResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *MavenMavenArtifactResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *MavenMavenArtifactResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *MavenMavenArtifactResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *MavenMavenArtifactResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *MavenMavenArtifactResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *MavenMavenArtifactResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *MavenMavenArtifactResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *MavenMavenArtifactResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *MavenMavenArtifactResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *MavenMavenArtifactResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *MavenMavenArtifactResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *MavenMavenArtifactResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *MavenMavenArtifactResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *MavenMavenArtifactResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *MavenMavenArtifactResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *MavenMavenArtifactResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetPulpLabels

`func (o *MavenMavenArtifactResponse) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *MavenMavenArtifactResponse) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *MavenMavenArtifactResponse) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *MavenMavenArtifactResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetArtifact

`func (o *MavenMavenArtifactResponse) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *MavenMavenArtifactResponse) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *MavenMavenArtifactResponse) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.


### GetGroupId

`func (o *MavenMavenArtifactResponse) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *MavenMavenArtifactResponse) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *MavenMavenArtifactResponse) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *MavenMavenArtifactResponse) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetArtifactId

`func (o *MavenMavenArtifactResponse) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *MavenMavenArtifactResponse) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *MavenMavenArtifactResponse) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *MavenMavenArtifactResponse) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### GetVersion

`func (o *MavenMavenArtifactResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MavenMavenArtifactResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MavenMavenArtifactResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MavenMavenArtifactResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFilename

`func (o *MavenMavenArtifactResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *MavenMavenArtifactResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *MavenMavenArtifactResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *MavenMavenArtifactResponse) HasFilename() bool`

HasFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


