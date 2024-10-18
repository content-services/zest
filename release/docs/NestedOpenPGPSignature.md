# NestedOpenPGPSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to **NullableString** |  | [optional] 
**Created** | **time.Time** |  | 
**ExpirationTime** | Pointer to **NullableString** |  | [optional] 
**SignersUserId** | Pointer to **NullableString** |  | [optional] 
**KeyExpirationTime** | Pointer to **NullableString** |  | [optional] 
**Expired** | **bool** |  | 

## Methods

### NewNestedOpenPGPSignature

`func NewNestedOpenPGPSignature(created time.Time, expired bool, ) *NestedOpenPGPSignature`

NewNestedOpenPGPSignature instantiates a new NestedOpenPGPSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedOpenPGPSignatureWithDefaults

`func NewNestedOpenPGPSignatureWithDefaults() *NestedOpenPGPSignature`

NewNestedOpenPGPSignatureWithDefaults instantiates a new NestedOpenPGPSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *NestedOpenPGPSignature) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *NestedOpenPGPSignature) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *NestedOpenPGPSignature) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *NestedOpenPGPSignature) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *NestedOpenPGPSignature) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *NestedOpenPGPSignature) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetCreated

`func (o *NestedOpenPGPSignature) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NestedOpenPGPSignature) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NestedOpenPGPSignature) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetExpirationTime

`func (o *NestedOpenPGPSignature) GetExpirationTime() string`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *NestedOpenPGPSignature) GetExpirationTimeOk() (*string, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *NestedOpenPGPSignature) SetExpirationTime(v string)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *NestedOpenPGPSignature) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### SetExpirationTimeNil

`func (o *NestedOpenPGPSignature) SetExpirationTimeNil(b bool)`

 SetExpirationTimeNil sets the value for ExpirationTime to be an explicit nil

### UnsetExpirationTime
`func (o *NestedOpenPGPSignature) UnsetExpirationTime()`

UnsetExpirationTime ensures that no value is present for ExpirationTime, not even an explicit nil
### GetSignersUserId

`func (o *NestedOpenPGPSignature) GetSignersUserId() string`

GetSignersUserId returns the SignersUserId field if non-nil, zero value otherwise.

### GetSignersUserIdOk

`func (o *NestedOpenPGPSignature) GetSignersUserIdOk() (*string, bool)`

GetSignersUserIdOk returns a tuple with the SignersUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignersUserId

`func (o *NestedOpenPGPSignature) SetSignersUserId(v string)`

SetSignersUserId sets SignersUserId field to given value.

### HasSignersUserId

`func (o *NestedOpenPGPSignature) HasSignersUserId() bool`

HasSignersUserId returns a boolean if a field has been set.

### SetSignersUserIdNil

`func (o *NestedOpenPGPSignature) SetSignersUserIdNil(b bool)`

 SetSignersUserIdNil sets the value for SignersUserId to be an explicit nil

### UnsetSignersUserId
`func (o *NestedOpenPGPSignature) UnsetSignersUserId()`

UnsetSignersUserId ensures that no value is present for SignersUserId, not even an explicit nil
### GetKeyExpirationTime

`func (o *NestedOpenPGPSignature) GetKeyExpirationTime() string`

GetKeyExpirationTime returns the KeyExpirationTime field if non-nil, zero value otherwise.

### GetKeyExpirationTimeOk

`func (o *NestedOpenPGPSignature) GetKeyExpirationTimeOk() (*string, bool)`

GetKeyExpirationTimeOk returns a tuple with the KeyExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExpirationTime

`func (o *NestedOpenPGPSignature) SetKeyExpirationTime(v string)`

SetKeyExpirationTime sets KeyExpirationTime field to given value.

### HasKeyExpirationTime

`func (o *NestedOpenPGPSignature) HasKeyExpirationTime() bool`

HasKeyExpirationTime returns a boolean if a field has been set.

### SetKeyExpirationTimeNil

`func (o *NestedOpenPGPSignature) SetKeyExpirationTimeNil(b bool)`

 SetKeyExpirationTimeNil sets the value for KeyExpirationTime to be an explicit nil

### UnsetKeyExpirationTime
`func (o *NestedOpenPGPSignature) UnsetKeyExpirationTime()`

UnsetKeyExpirationTime ensures that no value is present for KeyExpirationTime, not even an explicit nil
### GetExpired

`func (o *NestedOpenPGPSignature) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *NestedOpenPGPSignature) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *NestedOpenPGPSignature) SetExpired(v bool)`

SetExpired sets Expired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


