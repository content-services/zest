# PatchedserviceAndGateContentGuard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The unique name. | [optional] 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**Guards** | Pointer to **[]string** | List of ContentGuards; access is permitted only if ALL of them permit. | [optional] 

## Methods

### NewPatchedserviceAndGateContentGuard

`func NewPatchedserviceAndGateContentGuard() *PatchedserviceAndGateContentGuard`

NewPatchedserviceAndGateContentGuard instantiates a new PatchedserviceAndGateContentGuard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedserviceAndGateContentGuardWithDefaults

`func NewPatchedserviceAndGateContentGuardWithDefaults() *PatchedserviceAndGateContentGuard`

NewPatchedserviceAndGateContentGuardWithDefaults instantiates a new PatchedserviceAndGateContentGuard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedserviceAndGateContentGuard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedserviceAndGateContentGuard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedserviceAndGateContentGuard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedserviceAndGateContentGuard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedserviceAndGateContentGuard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedserviceAndGateContentGuard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedserviceAndGateContentGuard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedserviceAndGateContentGuard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedserviceAndGateContentGuard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedserviceAndGateContentGuard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetGuards

`func (o *PatchedserviceAndGateContentGuard) GetGuards() []string`

GetGuards returns the Guards field if non-nil, zero value otherwise.

### GetGuardsOk

`func (o *PatchedserviceAndGateContentGuard) GetGuardsOk() (*[]string, bool)`

GetGuardsOk returns a tuple with the Guards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuards

`func (o *PatchedserviceAndGateContentGuard) SetGuards(v []string)`

SetGuards sets Guards field to given value.

### HasGuards

`func (o *PatchedserviceAndGateContentGuard) HasGuards() bool`

HasGuards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


