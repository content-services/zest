# HeaderContentGuard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**HeaderName** | **string** | The header name the guard will check on. | 
**HeaderValue** | **string** | The value that will authorize the request. | 
**JqFilter** | Pointer to **NullableString** | A JQ syntax compatible filter. If jq_filter is not set, then the value willonly be Base64 decoded and checked as an explicit string match. | [optional] 

## Methods

### NewHeaderContentGuard

`func NewHeaderContentGuard(name string, headerName string, headerValue string, ) *HeaderContentGuard`

NewHeaderContentGuard instantiates a new HeaderContentGuard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderContentGuardWithDefaults

`func NewHeaderContentGuardWithDefaults() *HeaderContentGuard`

NewHeaderContentGuardWithDefaults instantiates a new HeaderContentGuard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HeaderContentGuard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeaderContentGuard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeaderContentGuard) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *HeaderContentGuard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HeaderContentGuard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HeaderContentGuard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HeaderContentGuard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HeaderContentGuard) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HeaderContentGuard) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHeaderName

`func (o *HeaderContentGuard) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *HeaderContentGuard) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *HeaderContentGuard) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.


### GetHeaderValue

`func (o *HeaderContentGuard) GetHeaderValue() string`

GetHeaderValue returns the HeaderValue field if non-nil, zero value otherwise.

### GetHeaderValueOk

`func (o *HeaderContentGuard) GetHeaderValueOk() (*string, bool)`

GetHeaderValueOk returns a tuple with the HeaderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderValue

`func (o *HeaderContentGuard) SetHeaderValue(v string)`

SetHeaderValue sets HeaderValue field to given value.


### GetJqFilter

`func (o *HeaderContentGuard) GetJqFilter() string`

GetJqFilter returns the JqFilter field if non-nil, zero value otherwise.

### GetJqFilterOk

`func (o *HeaderContentGuard) GetJqFilterOk() (*string, bool)`

GetJqFilterOk returns a tuple with the JqFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJqFilter

`func (o *HeaderContentGuard) SetJqFilter(v string)`

SetJqFilter sets JqFilter field to given value.

### HasJqFilter

`func (o *HeaderContentGuard) HasJqFilter() bool`

HasJqFilter returns a boolean if a field has been set.

### SetJqFilterNil

`func (o *HeaderContentGuard) SetJqFilterNil(b bool)`

 SetJqFilterNil sets the value for JqFilter to be an explicit nil

### UnsetJqFilter
`func (o *HeaderContentGuard) UnsetJqFilter()`

UnsetJqFilter ensures that no value is present for JqFilter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


