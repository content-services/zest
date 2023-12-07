# CompositeContentGuard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**Guards** | Pointer to **[]string** | List of ContentGuards to ask for access-permission. | [optional] 

## Methods

### NewCompositeContentGuard

`func NewCompositeContentGuard(name string, ) *CompositeContentGuard`

NewCompositeContentGuard instantiates a new CompositeContentGuard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompositeContentGuardWithDefaults

`func NewCompositeContentGuardWithDefaults() *CompositeContentGuard`

NewCompositeContentGuardWithDefaults instantiates a new CompositeContentGuard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CompositeContentGuard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompositeContentGuard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompositeContentGuard) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CompositeContentGuard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CompositeContentGuard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CompositeContentGuard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CompositeContentGuard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CompositeContentGuard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CompositeContentGuard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetGuards

`func (o *CompositeContentGuard) GetGuards() []*string`

GetGuards returns the Guards field if non-nil, zero value otherwise.

### GetGuardsOk

`func (o *CompositeContentGuard) GetGuardsOk() (*[]*string, bool)`

GetGuardsOk returns a tuple with the Guards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuards

`func (o *CompositeContentGuard) SetGuards(v []*string)`

SetGuards sets Guards field to given value.

### HasGuards

`func (o *CompositeContentGuard) HasGuards() bool`

HasGuards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


