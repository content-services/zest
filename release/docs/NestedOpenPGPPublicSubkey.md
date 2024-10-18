# NestedOpenPGPPublicSubkey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | **string** |  | 
**Created** | **time.Time** |  | 

## Methods

### NewNestedOpenPGPPublicSubkey

`func NewNestedOpenPGPPublicSubkey(fingerprint string, created time.Time, ) *NestedOpenPGPPublicSubkey`

NewNestedOpenPGPPublicSubkey instantiates a new NestedOpenPGPPublicSubkey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedOpenPGPPublicSubkeyWithDefaults

`func NewNestedOpenPGPPublicSubkeyWithDefaults() *NestedOpenPGPPublicSubkey`

NewNestedOpenPGPPublicSubkeyWithDefaults instantiates a new NestedOpenPGPPublicSubkey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprint

`func (o *NestedOpenPGPPublicSubkey) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *NestedOpenPGPPublicSubkey) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *NestedOpenPGPPublicSubkey) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetCreated

`func (o *NestedOpenPGPPublicSubkey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NestedOpenPGPPublicSubkey) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NestedOpenPGPPublicSubkey) SetCreated(v time.Time)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


