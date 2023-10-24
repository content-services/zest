# StorageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **NullableInt64** | Total number of bytes | 
**Used** | **NullableInt64** | Number of bytes in use | 
**Free** | **NullableInt64** | Number of free bytes | 

## Methods

### NewStorageResponse

`func NewStorageResponse(total NullableInt64, used NullableInt64, free NullableInt64, ) *StorageResponse`

NewStorageResponse instantiates a new StorageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageResponseWithDefaults

`func NewStorageResponseWithDefaults() *StorageResponse`

NewStorageResponseWithDefaults instantiates a new StorageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *StorageResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *StorageResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *StorageResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.


### SetTotalNil

`func (o *StorageResponse) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *StorageResponse) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetUsed

`func (o *StorageResponse) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *StorageResponse) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *StorageResponse) SetUsed(v int64)`

SetUsed sets Used field to given value.


### SetUsedNil

`func (o *StorageResponse) SetUsedNil(b bool)`

 SetUsedNil sets the value for Used to be an explicit nil

### UnsetUsed
`func (o *StorageResponse) UnsetUsed()`

UnsetUsed ensures that no value is present for Used, not even an explicit nil
### GetFree

`func (o *StorageResponse) GetFree() int64`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *StorageResponse) GetFreeOk() (*int64, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *StorageResponse) SetFree(v int64)`

SetFree sets Free field to given value.


### SetFreeNil

`func (o *StorageResponse) SetFreeNil(b bool)`

 SetFreeNil sets the value for Free to be an explicit nil

### UnsetFree
`func (o *StorageResponse) UnsetFree()`

UnsetFree ensures that no value is present for Free, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


