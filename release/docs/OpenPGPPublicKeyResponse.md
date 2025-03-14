# OpenPGPPublicKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] [readonly] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**UserIds** | Pointer to [**[]NestedOpenPGPUserIDResponse**](NestedOpenPGPUserIDResponse.md) |  | [optional] [readonly] 
**UserAttributes** | Pointer to [**[]NestedOpenPGPUserAttributeResponse**](NestedOpenPGPUserAttributeResponse.md) |  | [optional] [readonly] 
**PublicSubkeys** | Pointer to [**[]NestedOpenPGPPublicSubkeyResponse**](NestedOpenPGPPublicSubkeyResponse.md) |  | [optional] [readonly] 

## Methods

### NewOpenPGPPublicKeyResponse

`func NewOpenPGPPublicKeyResponse() *OpenPGPPublicKeyResponse`

NewOpenPGPPublicKeyResponse instantiates a new OpenPGPPublicKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenPGPPublicKeyResponseWithDefaults

`func NewOpenPGPPublicKeyResponseWithDefaults() *OpenPGPPublicKeyResponse`

NewOpenPGPPublicKeyResponseWithDefaults instantiates a new OpenPGPPublicKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *OpenPGPPublicKeyResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *OpenPGPPublicKeyResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *OpenPGPPublicKeyResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *OpenPGPPublicKeyResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *OpenPGPPublicKeyResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *OpenPGPPublicKeyResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *OpenPGPPublicKeyResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *OpenPGPPublicKeyResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *OpenPGPPublicKeyResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *OpenPGPPublicKeyResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *OpenPGPPublicKeyResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *OpenPGPPublicKeyResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *OpenPGPPublicKeyResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *OpenPGPPublicKeyResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *OpenPGPPublicKeyResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *OpenPGPPublicKeyResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetPulpLabels

`func (o *OpenPGPPublicKeyResponse) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *OpenPGPPublicKeyResponse) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *OpenPGPPublicKeyResponse) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *OpenPGPPublicKeyResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetFingerprint

`func (o *OpenPGPPublicKeyResponse) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *OpenPGPPublicKeyResponse) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *OpenPGPPublicKeyResponse) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *OpenPGPPublicKeyResponse) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetCreated

`func (o *OpenPGPPublicKeyResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OpenPGPPublicKeyResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OpenPGPPublicKeyResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OpenPGPPublicKeyResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUserIds

`func (o *OpenPGPPublicKeyResponse) GetUserIds() []NestedOpenPGPUserIDResponse`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *OpenPGPPublicKeyResponse) GetUserIdsOk() (*[]NestedOpenPGPUserIDResponse, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *OpenPGPPublicKeyResponse) SetUserIds(v []NestedOpenPGPUserIDResponse)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *OpenPGPPublicKeyResponse) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### GetUserAttributes

`func (o *OpenPGPPublicKeyResponse) GetUserAttributes() []NestedOpenPGPUserAttributeResponse`

GetUserAttributes returns the UserAttributes field if non-nil, zero value otherwise.

### GetUserAttributesOk

`func (o *OpenPGPPublicKeyResponse) GetUserAttributesOk() (*[]NestedOpenPGPUserAttributeResponse, bool)`

GetUserAttributesOk returns a tuple with the UserAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributes

`func (o *OpenPGPPublicKeyResponse) SetUserAttributes(v []NestedOpenPGPUserAttributeResponse)`

SetUserAttributes sets UserAttributes field to given value.

### HasUserAttributes

`func (o *OpenPGPPublicKeyResponse) HasUserAttributes() bool`

HasUserAttributes returns a boolean if a field has been set.

### GetPublicSubkeys

`func (o *OpenPGPPublicKeyResponse) GetPublicSubkeys() []NestedOpenPGPPublicSubkeyResponse`

GetPublicSubkeys returns the PublicSubkeys field if non-nil, zero value otherwise.

### GetPublicSubkeysOk

`func (o *OpenPGPPublicKeyResponse) GetPublicSubkeysOk() (*[]NestedOpenPGPPublicSubkeyResponse, bool)`

GetPublicSubkeysOk returns a tuple with the PublicSubkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicSubkeys

`func (o *OpenPGPPublicKeyResponse) SetPublicSubkeys(v []NestedOpenPGPPublicSubkeyResponse)`

SetPublicSubkeys sets PublicSubkeys field to given value.

### HasPublicSubkeys

`func (o *OpenPGPPublicKeyResponse) HasPublicSubkeys() bool`

HasPublicSubkeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


