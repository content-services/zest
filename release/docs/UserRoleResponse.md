# UserRoleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**Role** | **string** |  | 
**ContentObject** | Pointer to **NullableString** | pulp_href of the object for which role permissions should be asserted. If set to &#39;null&#39;, permissions will act on either domain or model-level. | [optional] 
**ContentObjectPrn** | Pointer to **NullableString** | prn of the object for which role permissions should be asserted. If set to &#39;null&#39;, permissions will act on either domain or model-level. | [optional] 
**Description** | Pointer to **string** |  | [optional] [readonly] 
**Permissions** | Pointer to **[]string** |  | [optional] [readonly] 
**Domain** | Pointer to **NullableString** | Domain this role should be applied on, mutually exclusive with content_object. | [optional] 

## Methods

### NewUserRoleResponse

`func NewUserRoleResponse(role string, ) *UserRoleResponse`

NewUserRoleResponse instantiates a new UserRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRoleResponseWithDefaults

`func NewUserRoleResponseWithDefaults() *UserRoleResponse`

NewUserRoleResponseWithDefaults instantiates a new UserRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *UserRoleResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *UserRoleResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *UserRoleResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *UserRoleResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *UserRoleResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *UserRoleResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *UserRoleResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *UserRoleResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *UserRoleResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *UserRoleResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *UserRoleResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *UserRoleResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *UserRoleResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *UserRoleResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *UserRoleResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *UserRoleResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetRole

`func (o *UserRoleResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserRoleResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserRoleResponse) SetRole(v string)`

SetRole sets Role field to given value.


### GetContentObject

`func (o *UserRoleResponse) GetContentObject() string`

GetContentObject returns the ContentObject field if non-nil, zero value otherwise.

### GetContentObjectOk

`func (o *UserRoleResponse) GetContentObjectOk() (*string, bool)`

GetContentObjectOk returns a tuple with the ContentObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentObject

`func (o *UserRoleResponse) SetContentObject(v string)`

SetContentObject sets ContentObject field to given value.

### HasContentObject

`func (o *UserRoleResponse) HasContentObject() bool`

HasContentObject returns a boolean if a field has been set.

### SetContentObjectNil

`func (o *UserRoleResponse) SetContentObjectNil(b bool)`

 SetContentObjectNil sets the value for ContentObject to be an explicit nil

### UnsetContentObject
`func (o *UserRoleResponse) UnsetContentObject()`

UnsetContentObject ensures that no value is present for ContentObject, not even an explicit nil
### GetContentObjectPrn

`func (o *UserRoleResponse) GetContentObjectPrn() string`

GetContentObjectPrn returns the ContentObjectPrn field if non-nil, zero value otherwise.

### GetContentObjectPrnOk

`func (o *UserRoleResponse) GetContentObjectPrnOk() (*string, bool)`

GetContentObjectPrnOk returns a tuple with the ContentObjectPrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentObjectPrn

`func (o *UserRoleResponse) SetContentObjectPrn(v string)`

SetContentObjectPrn sets ContentObjectPrn field to given value.

### HasContentObjectPrn

`func (o *UserRoleResponse) HasContentObjectPrn() bool`

HasContentObjectPrn returns a boolean if a field has been set.

### SetContentObjectPrnNil

`func (o *UserRoleResponse) SetContentObjectPrnNil(b bool)`

 SetContentObjectPrnNil sets the value for ContentObjectPrn to be an explicit nil

### UnsetContentObjectPrn
`func (o *UserRoleResponse) UnsetContentObjectPrn()`

UnsetContentObjectPrn ensures that no value is present for ContentObjectPrn, not even an explicit nil
### GetDescription

`func (o *UserRoleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserRoleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserRoleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserRoleResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermissions

`func (o *UserRoleResponse) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserRoleResponse) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserRoleResponse) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UserRoleResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetDomain

`func (o *UserRoleResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UserRoleResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UserRoleResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UserRoleResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *UserRoleResponse) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *UserRoleResponse) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


