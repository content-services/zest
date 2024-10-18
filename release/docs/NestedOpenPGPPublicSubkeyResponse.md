# NestedOpenPGPPublicSubkeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | **string** |  | 
**Created** | **time.Time** |  | 
**Signatures** | Pointer to [**[]NestedOpenPGPSignatureResponse**](NestedOpenPGPSignatureResponse.md) |  | [optional] [readonly] 

## Methods

### NewNestedOpenPGPPublicSubkeyResponse

`func NewNestedOpenPGPPublicSubkeyResponse(fingerprint string, created time.Time, ) *NestedOpenPGPPublicSubkeyResponse`

NewNestedOpenPGPPublicSubkeyResponse instantiates a new NestedOpenPGPPublicSubkeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedOpenPGPPublicSubkeyResponseWithDefaults

`func NewNestedOpenPGPPublicSubkeyResponseWithDefaults() *NestedOpenPGPPublicSubkeyResponse`

NewNestedOpenPGPPublicSubkeyResponseWithDefaults instantiates a new NestedOpenPGPPublicSubkeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprint

`func (o *NestedOpenPGPPublicSubkeyResponse) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *NestedOpenPGPPublicSubkeyResponse) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *NestedOpenPGPPublicSubkeyResponse) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetCreated

`func (o *NestedOpenPGPPublicSubkeyResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NestedOpenPGPPublicSubkeyResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NestedOpenPGPPublicSubkeyResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetSignatures

`func (o *NestedOpenPGPPublicSubkeyResponse) GetSignatures() []NestedOpenPGPSignatureResponse`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *NestedOpenPGPPublicSubkeyResponse) GetSignaturesOk() (*[]NestedOpenPGPSignatureResponse, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *NestedOpenPGPPublicSubkeyResponse) SetSignatures(v []NestedOpenPGPSignatureResponse)`

SetSignatures sets Signatures field to given value.

### HasSignatures

`func (o *NestedOpenPGPPublicSubkeyResponse) HasSignatures() bool`

HasSignatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


