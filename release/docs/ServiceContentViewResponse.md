# ServiceContentViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**Name** | **string** | A unique name for this content view. | 
**Description** | Pointer to **NullableString** | An optional description of this content view. | [optional] 
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Distributions** | Pointer to **[]string** | Distributions this content view searches across. May reference distributions belonging to any domain the user has read access to, not just this content view&#39;s own domain. | [optional] 
**DistributionsStatus** | Pointer to **string** | Per-distribution resolution status: whether each linked distribution&#39;s domain is currently accessible and whether it resolves to a repository version. | [optional] [readonly] 

## Methods

### NewServiceContentViewResponse

`func NewServiceContentViewResponse(name string, ) *ServiceContentViewResponse`

NewServiceContentViewResponse instantiates a new ServiceContentViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceContentViewResponseWithDefaults

`func NewServiceContentViewResponseWithDefaults() *ServiceContentViewResponse`

NewServiceContentViewResponseWithDefaults instantiates a new ServiceContentViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *ServiceContentViewResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ServiceContentViewResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ServiceContentViewResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ServiceContentViewResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *ServiceContentViewResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *ServiceContentViewResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *ServiceContentViewResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *ServiceContentViewResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *ServiceContentViewResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *ServiceContentViewResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *ServiceContentViewResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *ServiceContentViewResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *ServiceContentViewResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *ServiceContentViewResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *ServiceContentViewResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *ServiceContentViewResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *ServiceContentViewResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceContentViewResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceContentViewResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ServiceContentViewResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceContentViewResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceContentViewResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceContentViewResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ServiceContentViewResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ServiceContentViewResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPulpLabels

`func (o *ServiceContentViewResponse) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *ServiceContentViewResponse) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *ServiceContentViewResponse) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *ServiceContentViewResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetDistributions

`func (o *ServiceContentViewResponse) GetDistributions() []string`

GetDistributions returns the Distributions field if non-nil, zero value otherwise.

### GetDistributionsOk

`func (o *ServiceContentViewResponse) GetDistributionsOk() (*[]string, bool)`

GetDistributionsOk returns a tuple with the Distributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributions

`func (o *ServiceContentViewResponse) SetDistributions(v []string)`

SetDistributions sets Distributions field to given value.

### HasDistributions

`func (o *ServiceContentViewResponse) HasDistributions() bool`

HasDistributions returns a boolean if a field has been set.

### GetDistributionsStatus

`func (o *ServiceContentViewResponse) GetDistributionsStatus() string`

GetDistributionsStatus returns the DistributionsStatus field if non-nil, zero value otherwise.

### GetDistributionsStatusOk

`func (o *ServiceContentViewResponse) GetDistributionsStatusOk() (*string, bool)`

GetDistributionsStatusOk returns a tuple with the DistributionsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionsStatus

`func (o *ServiceContentViewResponse) SetDistributionsStatus(v string)`

SetDistributionsStatus sets DistributionsStatus field to given value.

### HasDistributionsStatus

`func (o *ServiceContentViewResponse) HasDistributionsStatus() bool`

HasDistributionsStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


