# HeaderContentGuardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**Name** | **string** | The unique name. | 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**HeaderName** | **string** | The header name the guard will check on. | 
**HeaderValue** | **string** | The value that will authorize the request. | 
**JqFilter** | Pointer to **NullableString** | A JQ syntax compatible filter. If jq_filter is not set, then the value willonly be Base64 decoded and checked as an explicit string match. | [optional] 

## Methods

### NewHeaderContentGuardResponse

`func NewHeaderContentGuardResponse(name string, headerName string, headerValue string, ) *HeaderContentGuardResponse`

NewHeaderContentGuardResponse instantiates a new HeaderContentGuardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderContentGuardResponseWithDefaults

`func NewHeaderContentGuardResponseWithDefaults() *HeaderContentGuardResponse`

NewHeaderContentGuardResponseWithDefaults instantiates a new HeaderContentGuardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *HeaderContentGuardResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *HeaderContentGuardResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *HeaderContentGuardResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *HeaderContentGuardResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *HeaderContentGuardResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *HeaderContentGuardResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *HeaderContentGuardResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *HeaderContentGuardResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *HeaderContentGuardResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *HeaderContentGuardResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *HeaderContentGuardResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *HeaderContentGuardResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *HeaderContentGuardResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeaderContentGuardResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeaderContentGuardResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *HeaderContentGuardResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HeaderContentGuardResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HeaderContentGuardResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HeaderContentGuardResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HeaderContentGuardResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HeaderContentGuardResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHeaderName

`func (o *HeaderContentGuardResponse) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *HeaderContentGuardResponse) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *HeaderContentGuardResponse) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.


### GetHeaderValue

`func (o *HeaderContentGuardResponse) GetHeaderValue() string`

GetHeaderValue returns the HeaderValue field if non-nil, zero value otherwise.

### GetHeaderValueOk

`func (o *HeaderContentGuardResponse) GetHeaderValueOk() (*string, bool)`

GetHeaderValueOk returns a tuple with the HeaderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderValue

`func (o *HeaderContentGuardResponse) SetHeaderValue(v string)`

SetHeaderValue sets HeaderValue field to given value.


### GetJqFilter

`func (o *HeaderContentGuardResponse) GetJqFilter() string`

GetJqFilter returns the JqFilter field if non-nil, zero value otherwise.

### GetJqFilterOk

`func (o *HeaderContentGuardResponse) GetJqFilterOk() (*string, bool)`

GetJqFilterOk returns a tuple with the JqFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJqFilter

`func (o *HeaderContentGuardResponse) SetJqFilter(v string)`

SetJqFilter sets JqFilter field to given value.

### HasJqFilter

`func (o *HeaderContentGuardResponse) HasJqFilter() bool`

HasJqFilter returns a boolean if a field has been set.

### SetJqFilterNil

`func (o *HeaderContentGuardResponse) SetJqFilterNil(b bool)`

 SetJqFilterNil sets the value for JqFilter to be an explicit nil

### UnsetJqFilter
`func (o *HeaderContentGuardResponse) UnsetJqFilter()`

UnsetJqFilter ensures that no value is present for JqFilter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


