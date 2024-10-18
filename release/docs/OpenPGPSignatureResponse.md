# OpenPGPSignatureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**Issuer** | Pointer to **NullableString** |  | [optional] 
**Created** | **time.Time** |  | 
**ExpirationTime** | Pointer to **NullableString** |  | [optional] 
**SignersUserId** | Pointer to **NullableString** |  | [optional] 
**KeyExpirationTime** | Pointer to **NullableString** |  | [optional] 
**Expired** | **bool** |  | 
**KeyExpired** | Pointer to **string** |  | [optional] [readonly] 
**SignedContent** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewOpenPGPSignatureResponse

`func NewOpenPGPSignatureResponse(created time.Time, expired bool, ) *OpenPGPSignatureResponse`

NewOpenPGPSignatureResponse instantiates a new OpenPGPSignatureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenPGPSignatureResponseWithDefaults

`func NewOpenPGPSignatureResponseWithDefaults() *OpenPGPSignatureResponse`

NewOpenPGPSignatureResponseWithDefaults instantiates a new OpenPGPSignatureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *OpenPGPSignatureResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *OpenPGPSignatureResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *OpenPGPSignatureResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *OpenPGPSignatureResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *OpenPGPSignatureResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *OpenPGPSignatureResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *OpenPGPSignatureResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *OpenPGPSignatureResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *OpenPGPSignatureResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *OpenPGPSignatureResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *OpenPGPSignatureResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *OpenPGPSignatureResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *OpenPGPSignatureResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *OpenPGPSignatureResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *OpenPGPSignatureResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *OpenPGPSignatureResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetIssuer

`func (o *OpenPGPSignatureResponse) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OpenPGPSignatureResponse) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OpenPGPSignatureResponse) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OpenPGPSignatureResponse) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *OpenPGPSignatureResponse) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *OpenPGPSignatureResponse) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetCreated

`func (o *OpenPGPSignatureResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OpenPGPSignatureResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OpenPGPSignatureResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetExpirationTime

`func (o *OpenPGPSignatureResponse) GetExpirationTime() string`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *OpenPGPSignatureResponse) GetExpirationTimeOk() (*string, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *OpenPGPSignatureResponse) SetExpirationTime(v string)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *OpenPGPSignatureResponse) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### SetExpirationTimeNil

`func (o *OpenPGPSignatureResponse) SetExpirationTimeNil(b bool)`

 SetExpirationTimeNil sets the value for ExpirationTime to be an explicit nil

### UnsetExpirationTime
`func (o *OpenPGPSignatureResponse) UnsetExpirationTime()`

UnsetExpirationTime ensures that no value is present for ExpirationTime, not even an explicit nil
### GetSignersUserId

`func (o *OpenPGPSignatureResponse) GetSignersUserId() string`

GetSignersUserId returns the SignersUserId field if non-nil, zero value otherwise.

### GetSignersUserIdOk

`func (o *OpenPGPSignatureResponse) GetSignersUserIdOk() (*string, bool)`

GetSignersUserIdOk returns a tuple with the SignersUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignersUserId

`func (o *OpenPGPSignatureResponse) SetSignersUserId(v string)`

SetSignersUserId sets SignersUserId field to given value.

### HasSignersUserId

`func (o *OpenPGPSignatureResponse) HasSignersUserId() bool`

HasSignersUserId returns a boolean if a field has been set.

### SetSignersUserIdNil

`func (o *OpenPGPSignatureResponse) SetSignersUserIdNil(b bool)`

 SetSignersUserIdNil sets the value for SignersUserId to be an explicit nil

### UnsetSignersUserId
`func (o *OpenPGPSignatureResponse) UnsetSignersUserId()`

UnsetSignersUserId ensures that no value is present for SignersUserId, not even an explicit nil
### GetKeyExpirationTime

`func (o *OpenPGPSignatureResponse) GetKeyExpirationTime() string`

GetKeyExpirationTime returns the KeyExpirationTime field if non-nil, zero value otherwise.

### GetKeyExpirationTimeOk

`func (o *OpenPGPSignatureResponse) GetKeyExpirationTimeOk() (*string, bool)`

GetKeyExpirationTimeOk returns a tuple with the KeyExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExpirationTime

`func (o *OpenPGPSignatureResponse) SetKeyExpirationTime(v string)`

SetKeyExpirationTime sets KeyExpirationTime field to given value.

### HasKeyExpirationTime

`func (o *OpenPGPSignatureResponse) HasKeyExpirationTime() bool`

HasKeyExpirationTime returns a boolean if a field has been set.

### SetKeyExpirationTimeNil

`func (o *OpenPGPSignatureResponse) SetKeyExpirationTimeNil(b bool)`

 SetKeyExpirationTimeNil sets the value for KeyExpirationTime to be an explicit nil

### UnsetKeyExpirationTime
`func (o *OpenPGPSignatureResponse) UnsetKeyExpirationTime()`

UnsetKeyExpirationTime ensures that no value is present for KeyExpirationTime, not even an explicit nil
### GetExpired

`func (o *OpenPGPSignatureResponse) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *OpenPGPSignatureResponse) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *OpenPGPSignatureResponse) SetExpired(v bool)`

SetExpired sets Expired field to given value.


### GetKeyExpired

`func (o *OpenPGPSignatureResponse) GetKeyExpired() string`

GetKeyExpired returns the KeyExpired field if non-nil, zero value otherwise.

### GetKeyExpiredOk

`func (o *OpenPGPSignatureResponse) GetKeyExpiredOk() (*string, bool)`

GetKeyExpiredOk returns a tuple with the KeyExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExpired

`func (o *OpenPGPSignatureResponse) SetKeyExpired(v string)`

SetKeyExpired sets KeyExpired field to given value.

### HasKeyExpired

`func (o *OpenPGPSignatureResponse) HasKeyExpired() bool`

HasKeyExpired returns a boolean if a field has been set.

### GetSignedContent

`func (o *OpenPGPSignatureResponse) GetSignedContent() string`

GetSignedContent returns the SignedContent field if non-nil, zero value otherwise.

### GetSignedContentOk

`func (o *OpenPGPSignatureResponse) GetSignedContentOk() (*string, bool)`

GetSignedContentOk returns a tuple with the SignedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedContent

`func (o *OpenPGPSignatureResponse) SetSignedContent(v string)`

SetSignedContent sets SignedContent field to given value.

### HasSignedContent

`func (o *OpenPGPSignatureResponse) HasSignedContent() bool`

HasSignedContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


