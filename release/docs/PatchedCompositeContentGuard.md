# PatchedCompositeContentGuard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The unique name. | [optional] 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**Guards** | Pointer to **[]string** | List of ContentGuards to ask for access-permission. | [optional] 

## Methods

### NewPatchedCompositeContentGuard

`func NewPatchedCompositeContentGuard() *PatchedCompositeContentGuard`

NewPatchedCompositeContentGuard instantiates a new PatchedCompositeContentGuard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCompositeContentGuardWithDefaults

`func NewPatchedCompositeContentGuardWithDefaults() *PatchedCompositeContentGuard`

NewPatchedCompositeContentGuardWithDefaults instantiates a new PatchedCompositeContentGuard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedCompositeContentGuard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedCompositeContentGuard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedCompositeContentGuard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedCompositeContentGuard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedCompositeContentGuard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedCompositeContentGuard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedCompositeContentGuard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedCompositeContentGuard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedCompositeContentGuard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedCompositeContentGuard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetGuards

`func (o *PatchedCompositeContentGuard) GetGuards() []*string`

GetGuards returns the Guards field if non-nil, zero value otherwise.

### GetGuardsOk

`func (o *PatchedCompositeContentGuard) GetGuardsOk() (*[]*string, bool)`

GetGuardsOk returns a tuple with the Guards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuards

`func (o *PatchedCompositeContentGuard) SetGuards(v []*string)`

SetGuards sets Guards field to given value.

### HasGuards

`func (o *PatchedCompositeContentGuard) HasGuards() bool`

HasGuards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


