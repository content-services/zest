# NestedOpenPGPUserAttributeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha256** | **string** |  | 
**Signatures** | Pointer to [**[]NestedOpenPGPSignatureResponse**](NestedOpenPGPSignatureResponse.md) |  | [optional] [readonly] 

## Methods

### NewNestedOpenPGPUserAttributeResponse

`func NewNestedOpenPGPUserAttributeResponse(sha256 string, ) *NestedOpenPGPUserAttributeResponse`

NewNestedOpenPGPUserAttributeResponse instantiates a new NestedOpenPGPUserAttributeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedOpenPGPUserAttributeResponseWithDefaults

`func NewNestedOpenPGPUserAttributeResponseWithDefaults() *NestedOpenPGPUserAttributeResponse`

NewNestedOpenPGPUserAttributeResponseWithDefaults instantiates a new NestedOpenPGPUserAttributeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha256

`func (o *NestedOpenPGPUserAttributeResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *NestedOpenPGPUserAttributeResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *NestedOpenPGPUserAttributeResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### GetSignatures

`func (o *NestedOpenPGPUserAttributeResponse) GetSignatures() []NestedOpenPGPSignatureResponse`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *NestedOpenPGPUserAttributeResponse) GetSignaturesOk() (*[]NestedOpenPGPSignatureResponse, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *NestedOpenPGPUserAttributeResponse) SetSignatures(v []NestedOpenPGPSignatureResponse)`

SetSignatures sets Signatures field to given value.

### HasSignatures

`func (o *NestedOpenPGPUserAttributeResponse) HasSignatures() bool`

HasSignatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


