# PaginatedOpenPGPPublicKeyResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]OpenPGPPublicKeyResponse**](OpenPGPPublicKeyResponse.md) |  | 

## Methods

### NewPaginatedOpenPGPPublicKeyResponseList

`func NewPaginatedOpenPGPPublicKeyResponseList(count int32, results []OpenPGPPublicKeyResponse, ) *PaginatedOpenPGPPublicKeyResponseList`

NewPaginatedOpenPGPPublicKeyResponseList instantiates a new PaginatedOpenPGPPublicKeyResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedOpenPGPPublicKeyResponseListWithDefaults

`func NewPaginatedOpenPGPPublicKeyResponseListWithDefaults() *PaginatedOpenPGPPublicKeyResponseList`

NewPaginatedOpenPGPPublicKeyResponseListWithDefaults instantiates a new PaginatedOpenPGPPublicKeyResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedOpenPGPPublicKeyResponseList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedOpenPGPPublicKeyResponseList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedOpenPGPPublicKeyResponseList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedOpenPGPPublicKeyResponseList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedOpenPGPPublicKeyResponseList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedOpenPGPPublicKeyResponseList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedOpenPGPPublicKeyResponseList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedOpenPGPPublicKeyResponseList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedOpenPGPPublicKeyResponseList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedOpenPGPPublicKeyResponseList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedOpenPGPPublicKeyResponseList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedOpenPGPPublicKeyResponseList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedOpenPGPPublicKeyResponseList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedOpenPGPPublicKeyResponseList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedOpenPGPPublicKeyResponseList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedOpenPGPPublicKeyResponseList) GetResults() []OpenPGPPublicKeyResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedOpenPGPPublicKeyResponseList) GetResultsOk() (*[]OpenPGPPublicKeyResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedOpenPGPPublicKeyResponseList) SetResults(v []OpenPGPPublicKeyResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


