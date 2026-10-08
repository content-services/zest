# PaginatedMavenRepositoryFlatPackageListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** |  | 
**Next** | **NullableString** |  | 
**Previous** | **NullableString** |  | 
**Results** | [**[]MavenRepositoryFlatPackageResponse**](MavenRepositoryFlatPackageResponse.md) |  | 

## Methods

### NewPaginatedMavenRepositoryFlatPackageListResponse

`func NewPaginatedMavenRepositoryFlatPackageListResponse(count int64, next NullableString, previous NullableString, results []MavenRepositoryFlatPackageResponse, ) *PaginatedMavenRepositoryFlatPackageListResponse`

NewPaginatedMavenRepositoryFlatPackageListResponse instantiates a new PaginatedMavenRepositoryFlatPackageListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedMavenRepositoryFlatPackageListResponseWithDefaults

`func NewPaginatedMavenRepositoryFlatPackageListResponseWithDefaults() *PaginatedMavenRepositoryFlatPackageListResponse`

NewPaginatedMavenRepositoryFlatPackageListResponseWithDefaults instantiates a new PaginatedMavenRepositoryFlatPackageListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedMavenRepositoryFlatPackageListResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedMavenRepositoryFlatPackageListResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedMavenRepositoryFlatPackageListResponse) SetCount(v int64)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedMavenRepositoryFlatPackageListResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedMavenRepositoryFlatPackageListResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedMavenRepositoryFlatPackageListResponse) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *PaginatedMavenRepositoryFlatPackageListResponse) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedMavenRepositoryFlatPackageListResponse) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedMavenRepositoryFlatPackageListResponse) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedMavenRepositoryFlatPackageListResponse) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedMavenRepositoryFlatPackageListResponse) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### SetPreviousNil

`func (o *PaginatedMavenRepositoryFlatPackageListResponse) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedMavenRepositoryFlatPackageListResponse) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedMavenRepositoryFlatPackageListResponse) GetResults() []MavenRepositoryFlatPackageResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedMavenRepositoryFlatPackageListResponse) GetResultsOk() (*[]MavenRepositoryFlatPackageResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedMavenRepositoryFlatPackageListResponse) SetResults(v []MavenRepositoryFlatPackageResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


