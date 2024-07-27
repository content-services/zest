# Copy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | **interface{}** | A JSON document describing sources, destinations, and content to be copied | 
**DependencySolving** | Pointer to **bool** | Also copy dependencies of the content being copied. | [optional] [default to true]

## Methods

### NewCopy

`func NewCopy(config interface{}, ) *Copy`

NewCopy instantiates a new Copy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyWithDefaults

`func NewCopyWithDefaults() *Copy`

NewCopyWithDefaults instantiates a new Copy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *Copy) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Copy) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Copy) SetConfig(v interface{})`

SetConfig sets Config field to given value.


### SetConfigNil

`func (o *Copy) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *Copy) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetDependencySolving

`func (o *Copy) GetDependencySolving() bool`

GetDependencySolving returns the DependencySolving field if non-nil, zero value otherwise.

### GetDependencySolvingOk

`func (o *Copy) GetDependencySolvingOk() (*bool, bool)`

GetDependencySolvingOk returns a tuple with the DependencySolving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencySolving

`func (o *Copy) SetDependencySolving(v bool)`

SetDependencySolving sets DependencySolving field to given value.

### HasDependencySolving

`func (o *Copy) HasDependencySolving() bool`

HasDependencySolving returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


