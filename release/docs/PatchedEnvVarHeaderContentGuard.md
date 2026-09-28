# PatchedEnvVarHeaderContentGuard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The unique name. | [optional] 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**HeaderName** | Pointer to **string** | The header name the guard will check on. | [optional] 
**EnvVar** | Pointer to **string** | Name of a content-app environment variable holding the expected secret (plaintext UTF-8). Must be listed in ENVVAR_HEADER_CONTENT_GUARD_ALLOWED_VARS. The request header must send that value Base64-encoded. The value is never stored in or returned by the API. | [optional] 

## Methods

### NewPatchedEnvVarHeaderContentGuard

`func NewPatchedEnvVarHeaderContentGuard() *PatchedEnvVarHeaderContentGuard`

NewPatchedEnvVarHeaderContentGuard instantiates a new PatchedEnvVarHeaderContentGuard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEnvVarHeaderContentGuardWithDefaults

`func NewPatchedEnvVarHeaderContentGuardWithDefaults() *PatchedEnvVarHeaderContentGuard`

NewPatchedEnvVarHeaderContentGuardWithDefaults instantiates a new PatchedEnvVarHeaderContentGuard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEnvVarHeaderContentGuard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEnvVarHeaderContentGuard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEnvVarHeaderContentGuard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEnvVarHeaderContentGuard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedEnvVarHeaderContentGuard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedEnvVarHeaderContentGuard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedEnvVarHeaderContentGuard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedEnvVarHeaderContentGuard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedEnvVarHeaderContentGuard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedEnvVarHeaderContentGuard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHeaderName

`func (o *PatchedEnvVarHeaderContentGuard) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *PatchedEnvVarHeaderContentGuard) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *PatchedEnvVarHeaderContentGuard) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.

### HasHeaderName

`func (o *PatchedEnvVarHeaderContentGuard) HasHeaderName() bool`

HasHeaderName returns a boolean if a field has been set.

### GetEnvVar

`func (o *PatchedEnvVarHeaderContentGuard) GetEnvVar() string`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *PatchedEnvVarHeaderContentGuard) GetEnvVarOk() (*string, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *PatchedEnvVarHeaderContentGuard) SetEnvVar(v string)`

SetEnvVar sets EnvVar field to given value.

### HasEnvVar

`func (o *PatchedEnvVarHeaderContentGuard) HasEnvVar() bool`

HasEnvVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


