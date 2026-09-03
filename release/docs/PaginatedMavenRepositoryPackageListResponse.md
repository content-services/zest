# PaginatedMavenRepositoryPackageListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** |  | 
**Next** | **NullableString** |  | 
**Previous** | **NullableString** |  | 
**Results** | [**[]MavenRepositoryPackageResponse**](MavenRepositoryPackageResponse.md) |  | 

## Methods

### NewPaginatedMavenRepositoryPackageListResponse

`func NewPaginatedMavenRepositoryPackageListResponse(count int64, next NullableString, previous NullableString, results []MavenRepositoryPackageResponse, ) *PaginatedMavenRepositoryPackageListResponse`

NewPaginatedMavenRepositoryPackageListResponse instantiates a new PaginatedMavenRepositoryPackageListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedMavenRepositoryPackageListResponseWithDefaults

`func NewPaginatedMavenRepositoryPackageListResponseWithDefaults() *PaginatedMavenRepositoryPackageListResponse`

NewPaginatedMavenRepositoryPackageListResponseWithDefaults instantiates a new PaginatedMavenRepositoryPackageListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedMavenRepositoryPackageListResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedMavenRepositoryPackageListResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedMavenRepositoryPackageListResponse) SetCount(v int64)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedMavenRepositoryPackageListResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedMavenRepositoryPackageListResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedMavenRepositoryPackageListResponse) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *PaginatedMavenRepositoryPackageListResponse) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedMavenRepositoryPackageListResponse) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedMavenRepositoryPackageListResponse) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedMavenRepositoryPackageListResponse) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedMavenRepositoryPackageListResponse) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### SetPreviousNil

`func (o *PaginatedMavenRepositoryPackageListResponse) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedMavenRepositoryPackageListResponse) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedMavenRepositoryPackageListResponse) GetResults() []MavenRepositoryPackageResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedMavenRepositoryPackageListResponse) GetResultsOk() (*[]MavenRepositoryPackageResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedMavenRepositoryPackageListResponse) SetResults(v []MavenRepositoryPackageResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


