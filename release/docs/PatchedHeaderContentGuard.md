# PatchedHeaderContentGuard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The unique name. | [optional] 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**HeaderName** | Pointer to **string** | The header name the guard will check on. | [optional] 
**HeaderValue** | Pointer to **string** | The value that will authorize the request. | [optional] 
**JqFilter** | Pointer to **NullableString** | A JQ syntax compatible filter. If jq_filter is not set, then the value willonly be Base64 decoded and checked as an explicit string match. | [optional] 

## Methods

### NewPatchedHeaderContentGuard

`func NewPatchedHeaderContentGuard() *PatchedHeaderContentGuard`

NewPatchedHeaderContentGuard instantiates a new PatchedHeaderContentGuard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedHeaderContentGuardWithDefaults

`func NewPatchedHeaderContentGuardWithDefaults() *PatchedHeaderContentGuard`

NewPatchedHeaderContentGuardWithDefaults instantiates a new PatchedHeaderContentGuard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedHeaderContentGuard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedHeaderContentGuard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedHeaderContentGuard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedHeaderContentGuard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedHeaderContentGuard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedHeaderContentGuard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedHeaderContentGuard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedHeaderContentGuard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedHeaderContentGuard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedHeaderContentGuard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHeaderName

`func (o *PatchedHeaderContentGuard) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *PatchedHeaderContentGuard) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *PatchedHeaderContentGuard) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.

### HasHeaderName

`func (o *PatchedHeaderContentGuard) HasHeaderName() bool`

HasHeaderName returns a boolean if a field has been set.

### GetHeaderValue

`func (o *PatchedHeaderContentGuard) GetHeaderValue() string`

GetHeaderValue returns the HeaderValue field if non-nil, zero value otherwise.

### GetHeaderValueOk

`func (o *PatchedHeaderContentGuard) GetHeaderValueOk() (*string, bool)`

GetHeaderValueOk returns a tuple with the HeaderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderValue

`func (o *PatchedHeaderContentGuard) SetHeaderValue(v string)`

SetHeaderValue sets HeaderValue field to given value.

### HasHeaderValue

`func (o *PatchedHeaderContentGuard) HasHeaderValue() bool`

HasHeaderValue returns a boolean if a field has been set.

### GetJqFilter

`func (o *PatchedHeaderContentGuard) GetJqFilter() string`

GetJqFilter returns the JqFilter field if non-nil, zero value otherwise.

### GetJqFilterOk

`func (o *PatchedHeaderContentGuard) GetJqFilterOk() (*string, bool)`

GetJqFilterOk returns a tuple with the JqFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJqFilter

`func (o *PatchedHeaderContentGuard) SetJqFilter(v string)`

SetJqFilter sets JqFilter field to given value.

### HasJqFilter

`func (o *PatchedHeaderContentGuard) HasJqFilter() bool`

HasJqFilter returns a boolean if a field has been set.

### SetJqFilterNil

`func (o *PatchedHeaderContentGuard) SetJqFilterNil(b bool)`

 SetJqFilterNil sets the value for JqFilter to be an explicit nil

### UnsetJqFilter
`func (o *PatchedHeaderContentGuard) UnsetJqFilter()`

UnsetJqFilter ensures that no value is present for JqFilter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


