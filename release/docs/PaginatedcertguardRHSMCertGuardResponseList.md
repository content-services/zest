# PaginatedcertguardRHSMCertGuardResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]CertguardRHSMCertGuardResponse**](CertguardRHSMCertGuardResponse.md) |  | 

## Methods

### NewPaginatedcertguardRHSMCertGuardResponseList

`func NewPaginatedcertguardRHSMCertGuardResponseList(count int32, results []CertguardRHSMCertGuardResponse, ) *PaginatedcertguardRHSMCertGuardResponseList`

NewPaginatedcertguardRHSMCertGuardResponseList instantiates a new PaginatedcertguardRHSMCertGuardResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedcertguardRHSMCertGuardResponseListWithDefaults

`func NewPaginatedcertguardRHSMCertGuardResponseListWithDefaults() *PaginatedcertguardRHSMCertGuardResponseList`

NewPaginatedcertguardRHSMCertGuardResponseListWithDefaults instantiates a new PaginatedcertguardRHSMCertGuardResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedcertguardRHSMCertGuardResponseList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedcertguardRHSMCertGuardResponseList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedcertguardRHSMCertGuardResponseList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedcertguardRHSMCertGuardResponseList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedcertguardRHSMCertGuardResponseList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedcertguardRHSMCertGuardResponseList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedcertguardRHSMCertGuardResponseList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedcertguardRHSMCertGuardResponseList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedcertguardRHSMCertGuardResponseList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedcertguardRHSMCertGuardResponseList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedcertguardRHSMCertGuardResponseList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedcertguardRHSMCertGuardResponseList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedcertguardRHSMCertGuardResponseList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedcertguardRHSMCertGuardResponseList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedcertguardRHSMCertGuardResponseList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedcertguardRHSMCertGuardResponseList) GetResults() []CertguardRHSMCertGuardResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedcertguardRHSMCertGuardResponseList) GetResultsOk() (*[]CertguardRHSMCertGuardResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedcertguardRHSMCertGuardResponseList) SetResults(v []CertguardRHSMCertGuardResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


