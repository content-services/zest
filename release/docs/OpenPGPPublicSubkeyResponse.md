# OpenPGPPublicSubkeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**Fingerprint** | **string** |  | 
**Created** | **time.Time** |  | 
**Signatures** | Pointer to [**[]NestedOpenPGPSignatureResponse**](NestedOpenPGPSignatureResponse.md) |  | [optional] [readonly] 
**PublicKey** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewOpenPGPPublicSubkeyResponse

`func NewOpenPGPPublicSubkeyResponse(fingerprint string, created time.Time, ) *OpenPGPPublicSubkeyResponse`

NewOpenPGPPublicSubkeyResponse instantiates a new OpenPGPPublicSubkeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenPGPPublicSubkeyResponseWithDefaults

`func NewOpenPGPPublicSubkeyResponseWithDefaults() *OpenPGPPublicSubkeyResponse`

NewOpenPGPPublicSubkeyResponseWithDefaults instantiates a new OpenPGPPublicSubkeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *OpenPGPPublicSubkeyResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *OpenPGPPublicSubkeyResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *OpenPGPPublicSubkeyResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *OpenPGPPublicSubkeyResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *OpenPGPPublicSubkeyResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *OpenPGPPublicSubkeyResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *OpenPGPPublicSubkeyResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *OpenPGPPublicSubkeyResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *OpenPGPPublicSubkeyResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *OpenPGPPublicSubkeyResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *OpenPGPPublicSubkeyResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *OpenPGPPublicSubkeyResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *OpenPGPPublicSubkeyResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *OpenPGPPublicSubkeyResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *OpenPGPPublicSubkeyResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *OpenPGPPublicSubkeyResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetFingerprint

`func (o *OpenPGPPublicSubkeyResponse) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *OpenPGPPublicSubkeyResponse) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *OpenPGPPublicSubkeyResponse) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetCreated

`func (o *OpenPGPPublicSubkeyResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OpenPGPPublicSubkeyResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OpenPGPPublicSubkeyResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetSignatures

`func (o *OpenPGPPublicSubkeyResponse) GetSignatures() []NestedOpenPGPSignatureResponse`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *OpenPGPPublicSubkeyResponse) GetSignaturesOk() (*[]NestedOpenPGPSignatureResponse, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *OpenPGPPublicSubkeyResponse) SetSignatures(v []NestedOpenPGPSignatureResponse)`

SetSignatures sets Signatures field to given value.

### HasSignatures

`func (o *OpenPGPPublicSubkeyResponse) HasSignatures() bool`

HasSignatures returns a boolean if a field has been set.

### GetPublicKey

`func (o *OpenPGPPublicSubkeyResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *OpenPGPPublicSubkeyResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *OpenPGPPublicSubkeyResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *OpenPGPPublicSubkeyResponse) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


