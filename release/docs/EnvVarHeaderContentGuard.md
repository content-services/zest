# EnvVarHeaderContentGuard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**HeaderName** | **string** | The header name the guard will check on. | 
**EnvVar** | **string** | Name of a content-app environment variable holding the expected secret (plaintext UTF-8). Must be listed in ENVVAR_HEADER_CONTENT_GUARD_ALLOWED_VARS. The request header must send that value Base64-encoded. The value is never stored in or returned by the API. | 

## Methods

### NewEnvVarHeaderContentGuard

`func NewEnvVarHeaderContentGuard(name string, headerName string, envVar string, ) *EnvVarHeaderContentGuard`

NewEnvVarHeaderContentGuard instantiates a new EnvVarHeaderContentGuard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvVarHeaderContentGuardWithDefaults

`func NewEnvVarHeaderContentGuardWithDefaults() *EnvVarHeaderContentGuard`

NewEnvVarHeaderContentGuardWithDefaults instantiates a new EnvVarHeaderContentGuard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvVarHeaderContentGuard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvVarHeaderContentGuard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvVarHeaderContentGuard) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EnvVarHeaderContentGuard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvVarHeaderContentGuard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvVarHeaderContentGuard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvVarHeaderContentGuard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EnvVarHeaderContentGuard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EnvVarHeaderContentGuard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHeaderName

`func (o *EnvVarHeaderContentGuard) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *EnvVarHeaderContentGuard) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *EnvVarHeaderContentGuard) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.


### GetEnvVar

`func (o *EnvVarHeaderContentGuard) GetEnvVar() string`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *EnvVarHeaderContentGuard) GetEnvVarOk() (*string, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *EnvVarHeaderContentGuard) SetEnvVar(v string)`

SetEnvVar sets EnvVar field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


