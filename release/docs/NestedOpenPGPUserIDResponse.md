# NestedOpenPGPUserIDResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**Signatures** | Pointer to [**[]NestedOpenPGPSignatureResponse**](NestedOpenPGPSignatureResponse.md) |  | [optional] [readonly] 

## Methods

### NewNestedOpenPGPUserIDResponse

`func NewNestedOpenPGPUserIDResponse(userId string, ) *NestedOpenPGPUserIDResponse`

NewNestedOpenPGPUserIDResponse instantiates a new NestedOpenPGPUserIDResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedOpenPGPUserIDResponseWithDefaults

`func NewNestedOpenPGPUserIDResponseWithDefaults() *NestedOpenPGPUserIDResponse`

NewNestedOpenPGPUserIDResponseWithDefaults instantiates a new NestedOpenPGPUserIDResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *NestedOpenPGPUserIDResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *NestedOpenPGPUserIDResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *NestedOpenPGPUserIDResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetSignatures

`func (o *NestedOpenPGPUserIDResponse) GetSignatures() []NestedOpenPGPSignatureResponse`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *NestedOpenPGPUserIDResponse) GetSignaturesOk() (*[]NestedOpenPGPSignatureResponse, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *NestedOpenPGPUserIDResponse) SetSignatures(v []NestedOpenPGPSignatureResponse)`

SetSignatures sets Signatures field to given value.

### HasSignatures

`func (o *NestedOpenPGPUserIDResponse) HasSignatures() bool`

HasSignatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


