# ServicePyPIYankMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Repository** | Pointer to **NullableString** |  | [optional] 
**RepositoryVersion** | Pointer to **NullableString** |  | [optional] 
**LastChecked** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewServicePyPIYankMonitor

`func NewServicePyPIYankMonitor(name string, ) *ServicePyPIYankMonitor`

NewServicePyPIYankMonitor instantiates a new ServicePyPIYankMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePyPIYankMonitorWithDefaults

`func NewServicePyPIYankMonitorWithDefaults() *ServicePyPIYankMonitor`

NewServicePyPIYankMonitorWithDefaults instantiates a new ServicePyPIYankMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServicePyPIYankMonitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicePyPIYankMonitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicePyPIYankMonitor) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ServicePyPIYankMonitor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicePyPIYankMonitor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicePyPIYankMonitor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServicePyPIYankMonitor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ServicePyPIYankMonitor) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ServicePyPIYankMonitor) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPulpLabels

`func (o *ServicePyPIYankMonitor) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *ServicePyPIYankMonitor) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *ServicePyPIYankMonitor) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *ServicePyPIYankMonitor) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetRepository

`func (o *ServicePyPIYankMonitor) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ServicePyPIYankMonitor) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ServicePyPIYankMonitor) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ServicePyPIYankMonitor) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *ServicePyPIYankMonitor) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *ServicePyPIYankMonitor) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetRepositoryVersion

`func (o *ServicePyPIYankMonitor) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *ServicePyPIYankMonitor) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *ServicePyPIYankMonitor) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *ServicePyPIYankMonitor) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### SetRepositoryVersionNil

`func (o *ServicePyPIYankMonitor) SetRepositoryVersionNil(b bool)`

 SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil

### UnsetRepositoryVersion
`func (o *ServicePyPIYankMonitor) UnsetRepositoryVersion()`

UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil
### GetLastChecked

`func (o *ServicePyPIYankMonitor) GetLastChecked() time.Time`

GetLastChecked returns the LastChecked field if non-nil, zero value otherwise.

### GetLastCheckedOk

`func (o *ServicePyPIYankMonitor) GetLastCheckedOk() (*time.Time, bool)`

GetLastCheckedOk returns a tuple with the LastChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChecked

`func (o *ServicePyPIYankMonitor) SetLastChecked(v time.Time)`

SetLastChecked sets LastChecked field to given value.

### HasLastChecked

`func (o *ServicePyPIYankMonitor) HasLastChecked() bool`

HasLastChecked returns a boolean if a field has been set.

### SetLastCheckedNil

`func (o *ServicePyPIYankMonitor) SetLastCheckedNil(b bool)`

 SetLastCheckedNil sets the value for LastChecked to be an explicit nil

### UnsetLastChecked
`func (o *ServicePyPIYankMonitor) UnsetLastChecked()`

UnsetLastChecked ensures that no value is present for LastChecked, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


