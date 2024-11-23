# ServiceFeatureContentGuard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**HeaderName** | **string** |  | 
**JqFilter** | Pointer to **NullableString** |  | [optional] 
**Features** | **[]string** | The list of features required to access the content. | 

## Methods

### NewServiceFeatureContentGuard

`func NewServiceFeatureContentGuard(name string, headerName string, features []string, ) *ServiceFeatureContentGuard`

NewServiceFeatureContentGuard instantiates a new ServiceFeatureContentGuard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceFeatureContentGuardWithDefaults

`func NewServiceFeatureContentGuardWithDefaults() *ServiceFeatureContentGuard`

NewServiceFeatureContentGuardWithDefaults instantiates a new ServiceFeatureContentGuard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceFeatureContentGuard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceFeatureContentGuard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceFeatureContentGuard) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ServiceFeatureContentGuard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceFeatureContentGuard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceFeatureContentGuard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceFeatureContentGuard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ServiceFeatureContentGuard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ServiceFeatureContentGuard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHeaderName

`func (o *ServiceFeatureContentGuard) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *ServiceFeatureContentGuard) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *ServiceFeatureContentGuard) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.


### GetJqFilter

`func (o *ServiceFeatureContentGuard) GetJqFilter() string`

GetJqFilter returns the JqFilter field if non-nil, zero value otherwise.

### GetJqFilterOk

`func (o *ServiceFeatureContentGuard) GetJqFilterOk() (*string, bool)`

GetJqFilterOk returns a tuple with the JqFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJqFilter

`func (o *ServiceFeatureContentGuard) SetJqFilter(v string)`

SetJqFilter sets JqFilter field to given value.

### HasJqFilter

`func (o *ServiceFeatureContentGuard) HasJqFilter() bool`

HasJqFilter returns a boolean if a field has been set.

### SetJqFilterNil

`func (o *ServiceFeatureContentGuard) SetJqFilterNil(b bool)`

 SetJqFilterNil sets the value for JqFilter to be an explicit nil

### UnsetJqFilter
`func (o *ServiceFeatureContentGuard) UnsetJqFilter()`

UnsetJqFilter ensures that no value is present for JqFilter, not even an explicit nil
### GetFeatures

`func (o *ServiceFeatureContentGuard) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ServiceFeatureContentGuard) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ServiceFeatureContentGuard) SetFeatures(v []string)`

SetFeatures sets Features field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


