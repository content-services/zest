# ServicePyPIYankMonitorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Repository** | Pointer to **NullableString** |  | [optional] 
**RepositoryVersion** | Pointer to **NullableString** |  | [optional] 
**LastChecked** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewServicePyPIYankMonitorResponse

`func NewServicePyPIYankMonitorResponse(name string, ) *ServicePyPIYankMonitorResponse`

NewServicePyPIYankMonitorResponse instantiates a new ServicePyPIYankMonitorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePyPIYankMonitorResponseWithDefaults

`func NewServicePyPIYankMonitorResponseWithDefaults() *ServicePyPIYankMonitorResponse`

NewServicePyPIYankMonitorResponseWithDefaults instantiates a new ServicePyPIYankMonitorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *ServicePyPIYankMonitorResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ServicePyPIYankMonitorResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ServicePyPIYankMonitorResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ServicePyPIYankMonitorResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *ServicePyPIYankMonitorResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *ServicePyPIYankMonitorResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *ServicePyPIYankMonitorResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *ServicePyPIYankMonitorResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *ServicePyPIYankMonitorResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *ServicePyPIYankMonitorResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *ServicePyPIYankMonitorResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *ServicePyPIYankMonitorResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *ServicePyPIYankMonitorResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *ServicePyPIYankMonitorResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *ServicePyPIYankMonitorResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *ServicePyPIYankMonitorResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *ServicePyPIYankMonitorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicePyPIYankMonitorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicePyPIYankMonitorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ServicePyPIYankMonitorResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicePyPIYankMonitorResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicePyPIYankMonitorResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServicePyPIYankMonitorResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ServicePyPIYankMonitorResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ServicePyPIYankMonitorResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPulpLabels

`func (o *ServicePyPIYankMonitorResponse) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *ServicePyPIYankMonitorResponse) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *ServicePyPIYankMonitorResponse) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *ServicePyPIYankMonitorResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetRepository

`func (o *ServicePyPIYankMonitorResponse) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ServicePyPIYankMonitorResponse) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ServicePyPIYankMonitorResponse) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ServicePyPIYankMonitorResponse) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *ServicePyPIYankMonitorResponse) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *ServicePyPIYankMonitorResponse) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetRepositoryVersion

`func (o *ServicePyPIYankMonitorResponse) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *ServicePyPIYankMonitorResponse) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *ServicePyPIYankMonitorResponse) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *ServicePyPIYankMonitorResponse) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### SetRepositoryVersionNil

`func (o *ServicePyPIYankMonitorResponse) SetRepositoryVersionNil(b bool)`

 SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil

### UnsetRepositoryVersion
`func (o *ServicePyPIYankMonitorResponse) UnsetRepositoryVersion()`

UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil
### GetLastChecked

`func (o *ServicePyPIYankMonitorResponse) GetLastChecked() time.Time`

GetLastChecked returns the LastChecked field if non-nil, zero value otherwise.

### GetLastCheckedOk

`func (o *ServicePyPIYankMonitorResponse) GetLastCheckedOk() (*time.Time, bool)`

GetLastCheckedOk returns a tuple with the LastChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChecked

`func (o *ServicePyPIYankMonitorResponse) SetLastChecked(v time.Time)`

SetLastChecked sets LastChecked field to given value.

### HasLastChecked

`func (o *ServicePyPIYankMonitorResponse) HasLastChecked() bool`

HasLastChecked returns a boolean if a field has been set.

### SetLastCheckedNil

`func (o *ServicePyPIYankMonitorResponse) SetLastCheckedNil(b bool)`

 SetLastCheckedNil sets the value for LastChecked to be an explicit nil

### UnsetLastChecked
`func (o *ServicePyPIYankMonitorResponse) UnsetLastChecked()`

UnsetLastChecked ensures that no value is present for LastChecked, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


