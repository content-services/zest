# RpmModulemdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same. | [optional] [readonly] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**Name** | **string** | Modulemd name. | 
**Stream** | **string** | Stream name. | 
**Version** | **string** | Modulemd version. | 
**StaticContext** | Pointer to **bool** | Modulemd static-context flag. | [optional] 
**Context** | **string** | Modulemd context. | 
**Arch** | **string** | Modulemd architecture. | 
**Artifacts** | **interface{}** | Modulemd artifacts. | 
**Dependencies** | **interface{}** | Modulemd dependencies. | 
**Packages** | Pointer to **[]string** | Modulemd artifacts&#39; packages. | [optional] 
**Profiles** | **interface{}** | Modulemd profiles. | 
**Description** | **string** | Description of module. | 

## Methods

### NewRpmModulemdResponse

`func NewRpmModulemdResponse(name string, stream string, version string, context string, arch string, artifacts interface{}, dependencies interface{}, profiles interface{}, description string, ) *RpmModulemdResponse`

NewRpmModulemdResponse instantiates a new RpmModulemdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmModulemdResponseWithDefaults

`func NewRpmModulemdResponseWithDefaults() *RpmModulemdResponse`

NewRpmModulemdResponseWithDefaults instantiates a new RpmModulemdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *RpmModulemdResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *RpmModulemdResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *RpmModulemdResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *RpmModulemdResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *RpmModulemdResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *RpmModulemdResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *RpmModulemdResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *RpmModulemdResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *RpmModulemdResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *RpmModulemdResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *RpmModulemdResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *RpmModulemdResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *RpmModulemdResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *RpmModulemdResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *RpmModulemdResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *RpmModulemdResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetPulpLabels

`func (o *RpmModulemdResponse) GetPulpLabels() map[string]string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *RpmModulemdResponse) GetPulpLabelsOk() (*map[string]string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *RpmModulemdResponse) SetPulpLabels(v map[string]string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *RpmModulemdResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *RpmModulemdResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RpmModulemdResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RpmModulemdResponse) SetName(v string)`

SetName sets Name field to given value.


### GetStream

`func (o *RpmModulemdResponse) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *RpmModulemdResponse) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *RpmModulemdResponse) SetStream(v string)`

SetStream sets Stream field to given value.


### GetVersion

`func (o *RpmModulemdResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RpmModulemdResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RpmModulemdResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetStaticContext

`func (o *RpmModulemdResponse) GetStaticContext() bool`

GetStaticContext returns the StaticContext field if non-nil, zero value otherwise.

### GetStaticContextOk

`func (o *RpmModulemdResponse) GetStaticContextOk() (*bool, bool)`

GetStaticContextOk returns a tuple with the StaticContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticContext

`func (o *RpmModulemdResponse) SetStaticContext(v bool)`

SetStaticContext sets StaticContext field to given value.

### HasStaticContext

`func (o *RpmModulemdResponse) HasStaticContext() bool`

HasStaticContext returns a boolean if a field has been set.

### GetContext

`func (o *RpmModulemdResponse) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RpmModulemdResponse) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RpmModulemdResponse) SetContext(v string)`

SetContext sets Context field to given value.


### GetArch

`func (o *RpmModulemdResponse) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *RpmModulemdResponse) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *RpmModulemdResponse) SetArch(v string)`

SetArch sets Arch field to given value.


### GetArtifacts

`func (o *RpmModulemdResponse) GetArtifacts() interface{}`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *RpmModulemdResponse) GetArtifactsOk() (*interface{}, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *RpmModulemdResponse) SetArtifacts(v interface{})`

SetArtifacts sets Artifacts field to given value.


### SetArtifactsNil

`func (o *RpmModulemdResponse) SetArtifactsNil(b bool)`

 SetArtifactsNil sets the value for Artifacts to be an explicit nil

### UnsetArtifacts
`func (o *RpmModulemdResponse) UnsetArtifacts()`

UnsetArtifacts ensures that no value is present for Artifacts, not even an explicit nil
### GetDependencies

`func (o *RpmModulemdResponse) GetDependencies() interface{}`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *RpmModulemdResponse) GetDependenciesOk() (*interface{}, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *RpmModulemdResponse) SetDependencies(v interface{})`

SetDependencies sets Dependencies field to given value.


### SetDependenciesNil

`func (o *RpmModulemdResponse) SetDependenciesNil(b bool)`

 SetDependenciesNil sets the value for Dependencies to be an explicit nil

### UnsetDependencies
`func (o *RpmModulemdResponse) UnsetDependencies()`

UnsetDependencies ensures that no value is present for Dependencies, not even an explicit nil
### GetPackages

`func (o *RpmModulemdResponse) GetPackages() []*string`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *RpmModulemdResponse) GetPackagesOk() (*[]*string, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *RpmModulemdResponse) SetPackages(v []*string)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *RpmModulemdResponse) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetProfiles

`func (o *RpmModulemdResponse) GetProfiles() interface{}`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *RpmModulemdResponse) GetProfilesOk() (*interface{}, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *RpmModulemdResponse) SetProfiles(v interface{})`

SetProfiles sets Profiles field to given value.


### SetProfilesNil

`func (o *RpmModulemdResponse) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *RpmModulemdResponse) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil
### GetDescription

`func (o *RpmModulemdResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RpmModulemdResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RpmModulemdResponse) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


