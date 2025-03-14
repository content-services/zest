# OstreeOstreeCommitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**Artifact** | **string** | Artifact file representing the physical content | 
**RelativePath** | **string** | Path where the artifact is located relative to distributions base_path | 
**ParentCommit** | Pointer to **NullableString** |  | [optional] 
**Checksum** | **string** |  | 
**Objs** | **[]string** |  | 

## Methods

### NewOstreeOstreeCommitResponse

`func NewOstreeOstreeCommitResponse(artifact string, relativePath string, checksum string, objs []string, ) *OstreeOstreeCommitResponse`

NewOstreeOstreeCommitResponse instantiates a new OstreeOstreeCommitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOstreeOstreeCommitResponseWithDefaults

`func NewOstreeOstreeCommitResponseWithDefaults() *OstreeOstreeCommitResponse`

NewOstreeOstreeCommitResponseWithDefaults instantiates a new OstreeOstreeCommitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *OstreeOstreeCommitResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *OstreeOstreeCommitResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *OstreeOstreeCommitResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *OstreeOstreeCommitResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *OstreeOstreeCommitResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *OstreeOstreeCommitResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *OstreeOstreeCommitResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *OstreeOstreeCommitResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *OstreeOstreeCommitResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *OstreeOstreeCommitResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *OstreeOstreeCommitResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *OstreeOstreeCommitResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *OstreeOstreeCommitResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *OstreeOstreeCommitResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *OstreeOstreeCommitResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *OstreeOstreeCommitResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetPulpLabels

`func (o *OstreeOstreeCommitResponse) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *OstreeOstreeCommitResponse) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *OstreeOstreeCommitResponse) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *OstreeOstreeCommitResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetArtifact

`func (o *OstreeOstreeCommitResponse) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *OstreeOstreeCommitResponse) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *OstreeOstreeCommitResponse) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.


### GetRelativePath

`func (o *OstreeOstreeCommitResponse) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *OstreeOstreeCommitResponse) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *OstreeOstreeCommitResponse) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.


### GetParentCommit

`func (o *OstreeOstreeCommitResponse) GetParentCommit() string`

GetParentCommit returns the ParentCommit field if non-nil, zero value otherwise.

### GetParentCommitOk

`func (o *OstreeOstreeCommitResponse) GetParentCommitOk() (*string, bool)`

GetParentCommitOk returns a tuple with the ParentCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommit

`func (o *OstreeOstreeCommitResponse) SetParentCommit(v string)`

SetParentCommit sets ParentCommit field to given value.

### HasParentCommit

`func (o *OstreeOstreeCommitResponse) HasParentCommit() bool`

HasParentCommit returns a boolean if a field has been set.

### SetParentCommitNil

`func (o *OstreeOstreeCommitResponse) SetParentCommitNil(b bool)`

 SetParentCommitNil sets the value for ParentCommit to be an explicit nil

### UnsetParentCommit
`func (o *OstreeOstreeCommitResponse) UnsetParentCommit()`

UnsetParentCommit ensures that no value is present for ParentCommit, not even an explicit nil
### GetChecksum

`func (o *OstreeOstreeCommitResponse) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *OstreeOstreeCommitResponse) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *OstreeOstreeCommitResponse) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.


### GetObjs

`func (o *OstreeOstreeCommitResponse) GetObjs() []string`

GetObjs returns the Objs field if non-nil, zero value otherwise.

### GetObjsOk

`func (o *OstreeOstreeCommitResponse) GetObjsOk() (*[]string, bool)`

GetObjsOk returns a tuple with the Objs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjs

`func (o *OstreeOstreeCommitResponse) SetObjs(v []string)`

SetObjs sets Objs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


