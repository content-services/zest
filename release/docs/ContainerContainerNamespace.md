# ContainerContainerNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewContainerContainerNamespace

`func NewContainerContainerNamespace(name string, ) *ContainerContainerNamespace`

NewContainerContainerNamespace instantiates a new ContainerContainerNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerContainerNamespaceWithDefaults

`func NewContainerContainerNamespaceWithDefaults() *ContainerContainerNamespace`

NewContainerContainerNamespaceWithDefaults instantiates a new ContainerContainerNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContainerContainerNamespace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerContainerNamespace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerContainerNamespace) SetName(v string)`

SetName sets Name field to given value.


### GetPulpLabels

`func (o *ContainerContainerNamespace) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *ContainerContainerNamespace) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *ContainerContainerNamespace) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *ContainerContainerNamespace) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


