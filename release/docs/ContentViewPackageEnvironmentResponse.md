# ContentViewPackageEnvironmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Id** | **string** | ID of the environment | 
**Name** | **string** | The name of the environment | 
**Description** | **string** | The description of the environment | 
**GroupIds** | **interface{}** | A list of group ids | 

## Methods

### NewContentViewPackageEnvironmentResponse

`func NewContentViewPackageEnvironmentResponse(id string, name string, description string, groupIds interface{}, ) *ContentViewPackageEnvironmentResponse`

NewContentViewPackageEnvironmentResponse instantiates a new ContentViewPackageEnvironmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentViewPackageEnvironmentResponseWithDefaults

`func NewContentViewPackageEnvironmentResponseWithDefaults() *ContentViewPackageEnvironmentResponse`

NewContentViewPackageEnvironmentResponseWithDefaults instantiates a new ContentViewPackageEnvironmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *ContentViewPackageEnvironmentResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ContentViewPackageEnvironmentResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ContentViewPackageEnvironmentResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ContentViewPackageEnvironmentResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetId

`func (o *ContentViewPackageEnvironmentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentViewPackageEnvironmentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentViewPackageEnvironmentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ContentViewPackageEnvironmentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentViewPackageEnvironmentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentViewPackageEnvironmentResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ContentViewPackageEnvironmentResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContentViewPackageEnvironmentResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContentViewPackageEnvironmentResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGroupIds

`func (o *ContentViewPackageEnvironmentResponse) GetGroupIds() interface{}`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *ContentViewPackageEnvironmentResponse) GetGroupIdsOk() (*interface{}, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *ContentViewPackageEnvironmentResponse) SetGroupIds(v interface{})`

SetGroupIds sets GroupIds field to given value.


### SetGroupIdsNil

`func (o *ContentViewPackageEnvironmentResponse) SetGroupIdsNil(b bool)`

 SetGroupIdsNil sets the value for GroupIds to be an explicit nil

### UnsetGroupIds
`func (o *ContentViewPackageEnvironmentResponse) UnsetGroupIds()`

UnsetGroupIds ensures that no value is present for GroupIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


