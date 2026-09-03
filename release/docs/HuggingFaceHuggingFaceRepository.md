# HuggingFaceHuggingFaceRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** | A unique name for this repository. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**RetainRepoVersions** | Pointer to **NullableInt64** | Retain X versions of the repository. Default is null which retains all versions. | [optional] 
**RetainCheckpoints** | Pointer to **NullableInt64** | Retain X checkpoint publications for the repository. Default is null which retains all checkpoints. | [optional] 
**Remote** | Pointer to **NullableString** | An optional remote to use by default when syncing. | [optional] 

## Methods

### NewHuggingFaceHuggingFaceRepository

`func NewHuggingFaceHuggingFaceRepository(name string, ) *HuggingFaceHuggingFaceRepository`

NewHuggingFaceHuggingFaceRepository instantiates a new HuggingFaceHuggingFaceRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHuggingFaceHuggingFaceRepositoryWithDefaults

`func NewHuggingFaceHuggingFaceRepositoryWithDefaults() *HuggingFaceHuggingFaceRepository`

NewHuggingFaceHuggingFaceRepositoryWithDefaults instantiates a new HuggingFaceHuggingFaceRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpLabels

`func (o *HuggingFaceHuggingFaceRepository) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *HuggingFaceHuggingFaceRepository) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *HuggingFaceHuggingFaceRepository) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *HuggingFaceHuggingFaceRepository) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *HuggingFaceHuggingFaceRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HuggingFaceHuggingFaceRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HuggingFaceHuggingFaceRepository) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *HuggingFaceHuggingFaceRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HuggingFaceHuggingFaceRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HuggingFaceHuggingFaceRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HuggingFaceHuggingFaceRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HuggingFaceHuggingFaceRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HuggingFaceHuggingFaceRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRetainRepoVersions

`func (o *HuggingFaceHuggingFaceRepository) GetRetainRepoVersions() int64`

GetRetainRepoVersions returns the RetainRepoVersions field if non-nil, zero value otherwise.

### GetRetainRepoVersionsOk

`func (o *HuggingFaceHuggingFaceRepository) GetRetainRepoVersionsOk() (*int64, bool)`

GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainRepoVersions

`func (o *HuggingFaceHuggingFaceRepository) SetRetainRepoVersions(v int64)`

SetRetainRepoVersions sets RetainRepoVersions field to given value.

### HasRetainRepoVersions

`func (o *HuggingFaceHuggingFaceRepository) HasRetainRepoVersions() bool`

HasRetainRepoVersions returns a boolean if a field has been set.

### SetRetainRepoVersionsNil

`func (o *HuggingFaceHuggingFaceRepository) SetRetainRepoVersionsNil(b bool)`

 SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil

### UnsetRetainRepoVersions
`func (o *HuggingFaceHuggingFaceRepository) UnsetRetainRepoVersions()`

UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
### GetRetainCheckpoints

`func (o *HuggingFaceHuggingFaceRepository) GetRetainCheckpoints() int64`

GetRetainCheckpoints returns the RetainCheckpoints field if non-nil, zero value otherwise.

### GetRetainCheckpointsOk

`func (o *HuggingFaceHuggingFaceRepository) GetRetainCheckpointsOk() (*int64, bool)`

GetRetainCheckpointsOk returns a tuple with the RetainCheckpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainCheckpoints

`func (o *HuggingFaceHuggingFaceRepository) SetRetainCheckpoints(v int64)`

SetRetainCheckpoints sets RetainCheckpoints field to given value.

### HasRetainCheckpoints

`func (o *HuggingFaceHuggingFaceRepository) HasRetainCheckpoints() bool`

HasRetainCheckpoints returns a boolean if a field has been set.

### SetRetainCheckpointsNil

`func (o *HuggingFaceHuggingFaceRepository) SetRetainCheckpointsNil(b bool)`

 SetRetainCheckpointsNil sets the value for RetainCheckpoints to be an explicit nil

### UnsetRetainCheckpoints
`func (o *HuggingFaceHuggingFaceRepository) UnsetRetainCheckpoints()`

UnsetRetainCheckpoints ensures that no value is present for RetainCheckpoints, not even an explicit nil
### GetRemote

`func (o *HuggingFaceHuggingFaceRepository) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *HuggingFaceHuggingFaceRepository) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *HuggingFaceHuggingFaceRepository) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *HuggingFaceHuggingFaceRepository) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *HuggingFaceHuggingFaceRepository) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *HuggingFaceHuggingFaceRepository) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


