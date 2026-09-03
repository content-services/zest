# ContentViewPackageGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Id** | **string** | ID of the group | 
**Name** | **string** | Name of the group | 
**Description** | **string** | Description of the group | 
**Packages** | **interface{}** | The list of packages in this group | 

## Methods

### NewContentViewPackageGroupResponse

`func NewContentViewPackageGroupResponse(id string, name string, description string, packages interface{}, ) *ContentViewPackageGroupResponse`

NewContentViewPackageGroupResponse instantiates a new ContentViewPackageGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentViewPackageGroupResponseWithDefaults

`func NewContentViewPackageGroupResponseWithDefaults() *ContentViewPackageGroupResponse`

NewContentViewPackageGroupResponseWithDefaults instantiates a new ContentViewPackageGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *ContentViewPackageGroupResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ContentViewPackageGroupResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ContentViewPackageGroupResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ContentViewPackageGroupResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetId

`func (o *ContentViewPackageGroupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentViewPackageGroupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentViewPackageGroupResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ContentViewPackageGroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentViewPackageGroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentViewPackageGroupResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ContentViewPackageGroupResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContentViewPackageGroupResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContentViewPackageGroupResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPackages

`func (o *ContentViewPackageGroupResponse) GetPackages() interface{}`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *ContentViewPackageGroupResponse) GetPackagesOk() (*interface{}, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *ContentViewPackageGroupResponse) SetPackages(v interface{})`

SetPackages sets Packages field to given value.


### SetPackagesNil

`func (o *ContentViewPackageGroupResponse) SetPackagesNil(b bool)`

 SetPackagesNil sets the value for Packages to be an explicit nil

### UnsetPackages
`func (o *ContentViewPackageGroupResponse) UnsetPackages()`

UnsetPackages ensures that no value is present for Packages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


