# PaginatedMavenRepositoryFlatPackageResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]MavenRepositoryFlatPackageResponse**](MavenRepositoryFlatPackageResponse.md) |  | 

## Methods

### NewPaginatedMavenRepositoryFlatPackageResponseList

`func NewPaginatedMavenRepositoryFlatPackageResponseList(count int32, results []MavenRepositoryFlatPackageResponse, ) *PaginatedMavenRepositoryFlatPackageResponseList`

NewPaginatedMavenRepositoryFlatPackageResponseList instantiates a new PaginatedMavenRepositoryFlatPackageResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedMavenRepositoryFlatPackageResponseListWithDefaults

`func NewPaginatedMavenRepositoryFlatPackageResponseListWithDefaults() *PaginatedMavenRepositoryFlatPackageResponseList`

NewPaginatedMavenRepositoryFlatPackageResponseListWithDefaults instantiates a new PaginatedMavenRepositoryFlatPackageResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedMavenRepositoryFlatPackageResponseList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedMavenRepositoryFlatPackageResponseList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedMavenRepositoryFlatPackageResponseList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedMavenRepositoryFlatPackageResponseList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedMavenRepositoryFlatPackageResponseList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedMavenRepositoryFlatPackageResponseList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedMavenRepositoryFlatPackageResponseList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedMavenRepositoryFlatPackageResponseList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedMavenRepositoryFlatPackageResponseList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedMavenRepositoryFlatPackageResponseList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedMavenRepositoryFlatPackageResponseList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedMavenRepositoryFlatPackageResponseList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedMavenRepositoryFlatPackageResponseList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedMavenRepositoryFlatPackageResponseList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedMavenRepositoryFlatPackageResponseList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedMavenRepositoryFlatPackageResponseList) GetResults() []MavenRepositoryFlatPackageResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedMavenRepositoryFlatPackageResponseList) GetResultsOk() (*[]MavenRepositoryFlatPackageResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedMavenRepositoryFlatPackageResponseList) SetResults(v []MavenRepositoryFlatPackageResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


