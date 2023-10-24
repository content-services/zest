# SetLabelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Value** | **NullableString** |  | 

## Methods

### NewSetLabelResponse

`func NewSetLabelResponse(key string, value NullableString, ) *SetLabelResponse`

NewSetLabelResponse instantiates a new SetLabelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetLabelResponseWithDefaults

`func NewSetLabelResponseWithDefaults() *SetLabelResponse`

NewSetLabelResponseWithDefaults instantiates a new SetLabelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SetLabelResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SetLabelResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SetLabelResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *SetLabelResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SetLabelResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SetLabelResponse) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *SetLabelResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *SetLabelResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


