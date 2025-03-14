# RpmPackageEnvironmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**Id** | **string** | Environment id. | 
**Name** | **string** | Environment name. | 
**Description** | **string** | Environment description. | 
**DisplayOrder** | **NullableInt64** | Environment display order. | 
**GroupIds** | **interface{}** | Environment group list. | 
**OptionIds** | **interface{}** | Environment option ids | 
**DescByLang** | **interface{}** | Environment description by language. | 
**NameByLang** | **interface{}** | Environment name by language. | 
**Digest** | **string** | Environment digest. | 

## Methods

### NewRpmPackageEnvironmentResponse

`func NewRpmPackageEnvironmentResponse(id string, name string, description string, displayOrder NullableInt64, groupIds interface{}, optionIds interface{}, descByLang interface{}, nameByLang interface{}, digest string, ) *RpmPackageEnvironmentResponse`

NewRpmPackageEnvironmentResponse instantiates a new RpmPackageEnvironmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmPackageEnvironmentResponseWithDefaults

`func NewRpmPackageEnvironmentResponseWithDefaults() *RpmPackageEnvironmentResponse`

NewRpmPackageEnvironmentResponseWithDefaults instantiates a new RpmPackageEnvironmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *RpmPackageEnvironmentResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *RpmPackageEnvironmentResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *RpmPackageEnvironmentResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *RpmPackageEnvironmentResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *RpmPackageEnvironmentResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *RpmPackageEnvironmentResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *RpmPackageEnvironmentResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *RpmPackageEnvironmentResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *RpmPackageEnvironmentResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *RpmPackageEnvironmentResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *RpmPackageEnvironmentResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *RpmPackageEnvironmentResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *RpmPackageEnvironmentResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *RpmPackageEnvironmentResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *RpmPackageEnvironmentResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *RpmPackageEnvironmentResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetPulpLabels

`func (o *RpmPackageEnvironmentResponse) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *RpmPackageEnvironmentResponse) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *RpmPackageEnvironmentResponse) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *RpmPackageEnvironmentResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetId

`func (o *RpmPackageEnvironmentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RpmPackageEnvironmentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RpmPackageEnvironmentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RpmPackageEnvironmentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RpmPackageEnvironmentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RpmPackageEnvironmentResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RpmPackageEnvironmentResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RpmPackageEnvironmentResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RpmPackageEnvironmentResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayOrder

`func (o *RpmPackageEnvironmentResponse) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *RpmPackageEnvironmentResponse) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *RpmPackageEnvironmentResponse) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.


### SetDisplayOrderNil

`func (o *RpmPackageEnvironmentResponse) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *RpmPackageEnvironmentResponse) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil
### GetGroupIds

`func (o *RpmPackageEnvironmentResponse) GetGroupIds() interface{}`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *RpmPackageEnvironmentResponse) GetGroupIdsOk() (*interface{}, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *RpmPackageEnvironmentResponse) SetGroupIds(v interface{})`

SetGroupIds sets GroupIds field to given value.


### SetGroupIdsNil

`func (o *RpmPackageEnvironmentResponse) SetGroupIdsNil(b bool)`

 SetGroupIdsNil sets the value for GroupIds to be an explicit nil

### UnsetGroupIds
`func (o *RpmPackageEnvironmentResponse) UnsetGroupIds()`

UnsetGroupIds ensures that no value is present for GroupIds, not even an explicit nil
### GetOptionIds

`func (o *RpmPackageEnvironmentResponse) GetOptionIds() interface{}`

GetOptionIds returns the OptionIds field if non-nil, zero value otherwise.

### GetOptionIdsOk

`func (o *RpmPackageEnvironmentResponse) GetOptionIdsOk() (*interface{}, bool)`

GetOptionIdsOk returns a tuple with the OptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionIds

`func (o *RpmPackageEnvironmentResponse) SetOptionIds(v interface{})`

SetOptionIds sets OptionIds field to given value.


### SetOptionIdsNil

`func (o *RpmPackageEnvironmentResponse) SetOptionIdsNil(b bool)`

 SetOptionIdsNil sets the value for OptionIds to be an explicit nil

### UnsetOptionIds
`func (o *RpmPackageEnvironmentResponse) UnsetOptionIds()`

UnsetOptionIds ensures that no value is present for OptionIds, not even an explicit nil
### GetDescByLang

`func (o *RpmPackageEnvironmentResponse) GetDescByLang() interface{}`

GetDescByLang returns the DescByLang field if non-nil, zero value otherwise.

### GetDescByLangOk

`func (o *RpmPackageEnvironmentResponse) GetDescByLangOk() (*interface{}, bool)`

GetDescByLangOk returns a tuple with the DescByLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescByLang

`func (o *RpmPackageEnvironmentResponse) SetDescByLang(v interface{})`

SetDescByLang sets DescByLang field to given value.


### SetDescByLangNil

`func (o *RpmPackageEnvironmentResponse) SetDescByLangNil(b bool)`

 SetDescByLangNil sets the value for DescByLang to be an explicit nil

### UnsetDescByLang
`func (o *RpmPackageEnvironmentResponse) UnsetDescByLang()`

UnsetDescByLang ensures that no value is present for DescByLang, not even an explicit nil
### GetNameByLang

`func (o *RpmPackageEnvironmentResponse) GetNameByLang() interface{}`

GetNameByLang returns the NameByLang field if non-nil, zero value otherwise.

### GetNameByLangOk

`func (o *RpmPackageEnvironmentResponse) GetNameByLangOk() (*interface{}, bool)`

GetNameByLangOk returns a tuple with the NameByLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameByLang

`func (o *RpmPackageEnvironmentResponse) SetNameByLang(v interface{})`

SetNameByLang sets NameByLang field to given value.


### SetNameByLangNil

`func (o *RpmPackageEnvironmentResponse) SetNameByLangNil(b bool)`

 SetNameByLangNil sets the value for NameByLang to be an explicit nil

### UnsetNameByLang
`func (o *RpmPackageEnvironmentResponse) UnsetNameByLang()`

UnsetNameByLang ensures that no value is present for NameByLang, not even an explicit nil
### GetDigest

`func (o *RpmPackageEnvironmentResponse) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *RpmPackageEnvironmentResponse) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *RpmPackageEnvironmentResponse) SetDigest(v string)`

SetDigest sets Digest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


