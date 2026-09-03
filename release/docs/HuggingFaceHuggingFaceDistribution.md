# HuggingFaceHuggingFaceDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasePath** | **string** | The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \&quot;foo\&quot; and \&quot;foo/bar\&quot;) | 
**ContentGuard** | Pointer to **NullableString** | An optional content-guard. | [optional] 
**Hidden** | Pointer to **bool** | Whether this distribution should be shown in the content app. | [optional] [default to false]
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** | A unique name. Ex, &#x60;rawhide&#x60; and &#x60;stable&#x60;. | 
**Repository** | Pointer to **NullableString** | The latest RepositoryVersion for this Repository will be served. | [optional] 
**RepositoryVersion** | Pointer to **NullableString** | RepositoryVersion to be served | [optional] 
**Publication** | Pointer to **NullableString** | Publication to be served | [optional] 
**Remote** | Pointer to **NullableString** | Remote that can be used to fetch content when using pull-through caching. | [optional] 

## Methods

### NewHuggingFaceHuggingFaceDistribution

`func NewHuggingFaceHuggingFaceDistribution(basePath string, name string, ) *HuggingFaceHuggingFaceDistribution`

NewHuggingFaceHuggingFaceDistribution instantiates a new HuggingFaceHuggingFaceDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHuggingFaceHuggingFaceDistributionWithDefaults

`func NewHuggingFaceHuggingFaceDistributionWithDefaults() *HuggingFaceHuggingFaceDistribution`

NewHuggingFaceHuggingFaceDistributionWithDefaults instantiates a new HuggingFaceHuggingFaceDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasePath

`func (o *HuggingFaceHuggingFaceDistribution) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *HuggingFaceHuggingFaceDistribution) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *HuggingFaceHuggingFaceDistribution) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.


### GetContentGuard

`func (o *HuggingFaceHuggingFaceDistribution) GetContentGuard() string`

GetContentGuard returns the ContentGuard field if non-nil, zero value otherwise.

### GetContentGuardOk

`func (o *HuggingFaceHuggingFaceDistribution) GetContentGuardOk() (*string, bool)`

GetContentGuardOk returns a tuple with the ContentGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuard

`func (o *HuggingFaceHuggingFaceDistribution) SetContentGuard(v string)`

SetContentGuard sets ContentGuard field to given value.

### HasContentGuard

`func (o *HuggingFaceHuggingFaceDistribution) HasContentGuard() bool`

HasContentGuard returns a boolean if a field has been set.

### SetContentGuardNil

`func (o *HuggingFaceHuggingFaceDistribution) SetContentGuardNil(b bool)`

 SetContentGuardNil sets the value for ContentGuard to be an explicit nil

### UnsetContentGuard
`func (o *HuggingFaceHuggingFaceDistribution) UnsetContentGuard()`

UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
### GetHidden

`func (o *HuggingFaceHuggingFaceDistribution) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *HuggingFaceHuggingFaceDistribution) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *HuggingFaceHuggingFaceDistribution) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *HuggingFaceHuggingFaceDistribution) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetPulpLabels

`func (o *HuggingFaceHuggingFaceDistribution) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *HuggingFaceHuggingFaceDistribution) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *HuggingFaceHuggingFaceDistribution) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *HuggingFaceHuggingFaceDistribution) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *HuggingFaceHuggingFaceDistribution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HuggingFaceHuggingFaceDistribution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HuggingFaceHuggingFaceDistribution) SetName(v string)`

SetName sets Name field to given value.


### GetRepository

`func (o *HuggingFaceHuggingFaceDistribution) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *HuggingFaceHuggingFaceDistribution) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *HuggingFaceHuggingFaceDistribution) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *HuggingFaceHuggingFaceDistribution) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *HuggingFaceHuggingFaceDistribution) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *HuggingFaceHuggingFaceDistribution) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetRepositoryVersion

`func (o *HuggingFaceHuggingFaceDistribution) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *HuggingFaceHuggingFaceDistribution) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *HuggingFaceHuggingFaceDistribution) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *HuggingFaceHuggingFaceDistribution) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### SetRepositoryVersionNil

`func (o *HuggingFaceHuggingFaceDistribution) SetRepositoryVersionNil(b bool)`

 SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil

### UnsetRepositoryVersion
`func (o *HuggingFaceHuggingFaceDistribution) UnsetRepositoryVersion()`

UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil
### GetPublication

`func (o *HuggingFaceHuggingFaceDistribution) GetPublication() string`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *HuggingFaceHuggingFaceDistribution) GetPublicationOk() (*string, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublication

`func (o *HuggingFaceHuggingFaceDistribution) SetPublication(v string)`

SetPublication sets Publication field to given value.

### HasPublication

`func (o *HuggingFaceHuggingFaceDistribution) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### SetPublicationNil

`func (o *HuggingFaceHuggingFaceDistribution) SetPublicationNil(b bool)`

 SetPublicationNil sets the value for Publication to be an explicit nil

### UnsetPublication
`func (o *HuggingFaceHuggingFaceDistribution) UnsetPublication()`

UnsetPublication ensures that no value is present for Publication, not even an explicit nil
### GetRemote

`func (o *HuggingFaceHuggingFaceDistribution) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *HuggingFaceHuggingFaceDistribution) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *HuggingFaceHuggingFaceDistribution) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *HuggingFaceHuggingFaceDistribution) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *HuggingFaceHuggingFaceDistribution) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *HuggingFaceHuggingFaceDistribution) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


