# RpmModulemdDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**Module** | **string** | Modulemd name. | 
**Stream** | **string** | Modulemd default stream. | 
**Profiles** | **interface{}** | Default profiles for modulemd streams. | 
**Snippet** | **string** | Modulemd default snippet | 

## Methods

### NewRpmModulemdDefaults

`func NewRpmModulemdDefaults(module string, stream string, profiles interface{}, snippet string, ) *RpmModulemdDefaults`

NewRpmModulemdDefaults instantiates a new RpmModulemdDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmModulemdDefaultsWithDefaults

`func NewRpmModulemdDefaultsWithDefaults() *RpmModulemdDefaults`

NewRpmModulemdDefaultsWithDefaults instantiates a new RpmModulemdDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *RpmModulemdDefaults) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RpmModulemdDefaults) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RpmModulemdDefaults) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *RpmModulemdDefaults) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetPulpLabels

`func (o *RpmModulemdDefaults) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *RpmModulemdDefaults) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *RpmModulemdDefaults) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *RpmModulemdDefaults) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetModule

`func (o *RpmModulemdDefaults) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *RpmModulemdDefaults) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *RpmModulemdDefaults) SetModule(v string)`

SetModule sets Module field to given value.


### GetStream

`func (o *RpmModulemdDefaults) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *RpmModulemdDefaults) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *RpmModulemdDefaults) SetStream(v string)`

SetStream sets Stream field to given value.


### GetProfiles

`func (o *RpmModulemdDefaults) GetProfiles() interface{}`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *RpmModulemdDefaults) GetProfilesOk() (*interface{}, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *RpmModulemdDefaults) SetProfiles(v interface{})`

SetProfiles sets Profiles field to given value.


### SetProfilesNil

`func (o *RpmModulemdDefaults) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *RpmModulemdDefaults) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil
### GetSnippet

`func (o *RpmModulemdDefaults) GetSnippet() string`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *RpmModulemdDefaults) GetSnippetOk() (*string, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *RpmModulemdDefaults) SetSnippet(v string)`

SetSnippet sets Snippet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


