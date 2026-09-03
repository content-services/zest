# ContentViewModuleStreamResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** | Name of the modulemd | 
**Stream** | **string** | The modulemd&#39;s stream | 
**Version** | **string** | The version of the modulemd | 
**Context** | **string** | The modulemd&#39;s context flag | 
**Arch** | **string** | Module artifact architecture | 
**Description** | **string** | A verbose description of the module | 
**Profiles** | **interface{}** | Package lists of installable profiles | 
**Packages** | Pointer to **string** | Names of packages in this module | [optional] [readonly] 

## Methods

### NewContentViewModuleStreamResponse

`func NewContentViewModuleStreamResponse(name string, stream string, version string, context string, arch string, description string, profiles interface{}, ) *ContentViewModuleStreamResponse`

NewContentViewModuleStreamResponse instantiates a new ContentViewModuleStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentViewModuleStreamResponseWithDefaults

`func NewContentViewModuleStreamResponseWithDefaults() *ContentViewModuleStreamResponse`

NewContentViewModuleStreamResponseWithDefaults instantiates a new ContentViewModuleStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *ContentViewModuleStreamResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ContentViewModuleStreamResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ContentViewModuleStreamResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ContentViewModuleStreamResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetName

`func (o *ContentViewModuleStreamResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentViewModuleStreamResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentViewModuleStreamResponse) SetName(v string)`

SetName sets Name field to given value.


### GetStream

`func (o *ContentViewModuleStreamResponse) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ContentViewModuleStreamResponse) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ContentViewModuleStreamResponse) SetStream(v string)`

SetStream sets Stream field to given value.


### GetVersion

`func (o *ContentViewModuleStreamResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ContentViewModuleStreamResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ContentViewModuleStreamResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetContext

`func (o *ContentViewModuleStreamResponse) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ContentViewModuleStreamResponse) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ContentViewModuleStreamResponse) SetContext(v string)`

SetContext sets Context field to given value.


### GetArch

`func (o *ContentViewModuleStreamResponse) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *ContentViewModuleStreamResponse) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *ContentViewModuleStreamResponse) SetArch(v string)`

SetArch sets Arch field to given value.


### GetDescription

`func (o *ContentViewModuleStreamResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContentViewModuleStreamResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContentViewModuleStreamResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProfiles

`func (o *ContentViewModuleStreamResponse) GetProfiles() interface{}`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *ContentViewModuleStreamResponse) GetProfilesOk() (*interface{}, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *ContentViewModuleStreamResponse) SetProfiles(v interface{})`

SetProfiles sets Profiles field to given value.


### SetProfilesNil

`func (o *ContentViewModuleStreamResponse) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *ContentViewModuleStreamResponse) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil
### GetPackages

`func (o *ContentViewModuleStreamResponse) GetPackages() string`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *ContentViewModuleStreamResponse) GetPackagesOk() (*string, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *ContentViewModuleStreamResponse) SetPackages(v string)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *ContentViewModuleStreamResponse) HasPackages() bool`

HasPackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


