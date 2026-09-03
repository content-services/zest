# ServiceContentView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name for this content view. | 
**Description** | Pointer to **NullableString** | An optional description of this content view. | [optional] 
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Distributions** | Pointer to **[]string** | Distributions this content view searches across. May reference distributions belonging to any domain the user has read access to, not just this content view&#39;s own domain. | [optional] 

## Methods

### NewServiceContentView

`func NewServiceContentView(name string, ) *ServiceContentView`

NewServiceContentView instantiates a new ServiceContentView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceContentViewWithDefaults

`func NewServiceContentViewWithDefaults() *ServiceContentView`

NewServiceContentViewWithDefaults instantiates a new ServiceContentView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceContentView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceContentView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceContentView) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ServiceContentView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceContentView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceContentView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceContentView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ServiceContentView) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ServiceContentView) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPulpLabels

`func (o *ServiceContentView) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *ServiceContentView) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *ServiceContentView) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *ServiceContentView) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetDistributions

`func (o *ServiceContentView) GetDistributions() []string`

GetDistributions returns the Distributions field if non-nil, zero value otherwise.

### GetDistributionsOk

`func (o *ServiceContentView) GetDistributionsOk() (*[]string, bool)`

GetDistributionsOk returns a tuple with the Distributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributions

`func (o *ServiceContentView) SetDistributions(v []string)`

SetDistributions sets Distributions field to given value.

### HasDistributions

`func (o *ServiceContentView) HasDistributions() bool`

HasDistributions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


