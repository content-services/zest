# PaginatedcontainerContainerRemoteResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]ContainerContainerRemoteResponse**](ContainerContainerRemoteResponse.md) |  | 

## Methods

### NewPaginatedcontainerContainerRemoteResponseList

`func NewPaginatedcontainerContainerRemoteResponseList(count int32, results []ContainerContainerRemoteResponse, ) *PaginatedcontainerContainerRemoteResponseList`

NewPaginatedcontainerContainerRemoteResponseList instantiates a new PaginatedcontainerContainerRemoteResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedcontainerContainerRemoteResponseListWithDefaults

`func NewPaginatedcontainerContainerRemoteResponseListWithDefaults() *PaginatedcontainerContainerRemoteResponseList`

NewPaginatedcontainerContainerRemoteResponseListWithDefaults instantiates a new PaginatedcontainerContainerRemoteResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedcontainerContainerRemoteResponseList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedcontainerContainerRemoteResponseList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedcontainerContainerRemoteResponseList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedcontainerContainerRemoteResponseList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedcontainerContainerRemoteResponseList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedcontainerContainerRemoteResponseList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedcontainerContainerRemoteResponseList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedcontainerContainerRemoteResponseList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedcontainerContainerRemoteResponseList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedcontainerContainerRemoteResponseList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedcontainerContainerRemoteResponseList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedcontainerContainerRemoteResponseList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedcontainerContainerRemoteResponseList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedcontainerContainerRemoteResponseList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedcontainerContainerRemoteResponseList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedcontainerContainerRemoteResponseList) GetResults() []ContainerContainerRemoteResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedcontainerContainerRemoteResponseList) GetResultsOk() (*[]ContainerContainerRemoteResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedcontainerContainerRemoteResponseList) SetResults(v []ContainerContainerRemoteResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


