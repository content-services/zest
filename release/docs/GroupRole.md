# GroupRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** |  | 
**ContentObject** | Pointer to **NullableString** | pulp_href of the object for which role permissions should be asserted. If set to &#39;null&#39;, permissions will act on the model-level. | [optional] 
**ContentObjectPrn** | Pointer to **NullableString** | prn of the object for which role permissions should be asserted. If set to &#39;null&#39;, permissions will act on either domain or model-level. | [optional] 
**Domain** | Pointer to **NullableString** | Domain this role should be applied on, mutually exclusive with content_object. | [optional] 

## Methods

### NewGroupRole

`func NewGroupRole(role string, ) *GroupRole`

NewGroupRole instantiates a new GroupRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRoleWithDefaults

`func NewGroupRoleWithDefaults() *GroupRole`

NewGroupRoleWithDefaults instantiates a new GroupRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *GroupRole) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GroupRole) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GroupRole) SetRole(v string)`

SetRole sets Role field to given value.


### GetContentObject

`func (o *GroupRole) GetContentObject() string`

GetContentObject returns the ContentObject field if non-nil, zero value otherwise.

### GetContentObjectOk

`func (o *GroupRole) GetContentObjectOk() (*string, bool)`

GetContentObjectOk returns a tuple with the ContentObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentObject

`func (o *GroupRole) SetContentObject(v string)`

SetContentObject sets ContentObject field to given value.

### HasContentObject

`func (o *GroupRole) HasContentObject() bool`

HasContentObject returns a boolean if a field has been set.

### SetContentObjectNil

`func (o *GroupRole) SetContentObjectNil(b bool)`

 SetContentObjectNil sets the value for ContentObject to be an explicit nil

### UnsetContentObject
`func (o *GroupRole) UnsetContentObject()`

UnsetContentObject ensures that no value is present for ContentObject, not even an explicit nil
### GetContentObjectPrn

`func (o *GroupRole) GetContentObjectPrn() string`

GetContentObjectPrn returns the ContentObjectPrn field if non-nil, zero value otherwise.

### GetContentObjectPrnOk

`func (o *GroupRole) GetContentObjectPrnOk() (*string, bool)`

GetContentObjectPrnOk returns a tuple with the ContentObjectPrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentObjectPrn

`func (o *GroupRole) SetContentObjectPrn(v string)`

SetContentObjectPrn sets ContentObjectPrn field to given value.

### HasContentObjectPrn

`func (o *GroupRole) HasContentObjectPrn() bool`

HasContentObjectPrn returns a boolean if a field has been set.

### SetContentObjectPrnNil

`func (o *GroupRole) SetContentObjectPrnNil(b bool)`

 SetContentObjectPrnNil sets the value for ContentObjectPrn to be an explicit nil

### UnsetContentObjectPrn
`func (o *GroupRole) UnsetContentObjectPrn()`

UnsetContentObjectPrn ensures that no value is present for ContentObjectPrn, not even an explicit nil
### GetDomain

`func (o *GroupRole) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GroupRole) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GroupRole) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GroupRole) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *GroupRole) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *GroupRole) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


