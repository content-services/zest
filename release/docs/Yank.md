# Yank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the package to yank or unyank. | 
**Version** | **string** | The version of the package to yank or unyank. | 
**YankedReason** | Pointer to **string** | The reason for yanking the package version. | [optional] [default to ""]

## Methods

### NewYank

`func NewYank(name string, version string, ) *Yank`

NewYank instantiates a new Yank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYankWithDefaults

`func NewYankWithDefaults() *Yank`

NewYankWithDefaults instantiates a new Yank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Yank) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Yank) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Yank) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *Yank) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Yank) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Yank) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetYankedReason

`func (o *Yank) GetYankedReason() string`

GetYankedReason returns the YankedReason field if non-nil, zero value otherwise.

### GetYankedReasonOk

`func (o *Yank) GetYankedReasonOk() (*string, bool)`

GetYankedReasonOk returns a tuple with the YankedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYankedReason

`func (o *Yank) SetYankedReason(v string)`

SetYankedReason sets YankedReason field to given value.

### HasYankedReason

`func (o *Yank) HasYankedReason() bool`

HasYankedReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


