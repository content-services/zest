# PatchedhuggingFaceHuggingFaceDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasePath** | Pointer to **string** | The base (relative) path component of the published url. Avoid paths that                     overlap with other distribution base paths (e.g. \&quot;foo\&quot; and \&quot;foo/bar\&quot;) | [optional] 
**ContentGuard** | Pointer to **NullableString** | An optional content-guard. | [optional] 
**Hidden** | Pointer to **bool** | Whether this distribution should be shown in the content app. | [optional] [default to false]
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** | A unique name. Ex, &#x60;rawhide&#x60; and &#x60;stable&#x60;. | [optional] 
**Repository** | Pointer to **NullableString** | The latest RepositoryVersion for this Repository will be served. | [optional] 
**RepositoryVersion** | Pointer to **NullableString** | RepositoryVersion to be served | [optional] 
**Publication** | Pointer to **NullableString** | Publication to be served | [optional] 
**Remote** | Pointer to **NullableString** | Remote that can be used to fetch content when using pull-through caching. | [optional] 

## Methods

### NewPatchedhuggingFaceHuggingFaceDistribution

`func NewPatchedhuggingFaceHuggingFaceDistribution() *PatchedhuggingFaceHuggingFaceDistribution`

NewPatchedhuggingFaceHuggingFaceDistribution instantiates a new PatchedhuggingFaceHuggingFaceDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedhuggingFaceHuggingFaceDistributionWithDefaults

`func NewPatchedhuggingFaceHuggingFaceDistributionWithDefaults() *PatchedhuggingFaceHuggingFaceDistribution`

NewPatchedhuggingFaceHuggingFaceDistributionWithDefaults instantiates a new PatchedhuggingFaceHuggingFaceDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasePath

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *PatchedhuggingFaceHuggingFaceDistribution) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.

### HasBasePath

`func (o *PatchedhuggingFaceHuggingFaceDistribution) HasBasePath() bool`

HasBasePath returns a boolean if a field has been set.

### GetContentGuard

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetContentGuard() string`

GetContentGuard returns the ContentGuard field if non-nil, zero value otherwise.

### GetContentGuardOk

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetContentGuardOk() (*string, bool)`

GetContentGuardOk returns a tuple with the ContentGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentGuard

`func (o *PatchedhuggingFaceHuggingFaceDistribution) SetContentGuard(v string)`

SetContentGuard sets ContentGuard field to given value.

### HasContentGuard

`func (o *PatchedhuggingFaceHuggingFaceDistribution) HasContentGuard() bool`

HasContentGuard returns a boolean if a field has been set.

### SetContentGuardNil

`func (o *PatchedhuggingFaceHuggingFaceDistribution) SetContentGuardNil(b bool)`

 SetContentGuardNil sets the value for ContentGuard to be an explicit nil

### UnsetContentGuard
`func (o *PatchedhuggingFaceHuggingFaceDistribution) UnsetContentGuard()`

UnsetContentGuard ensures that no value is present for ContentGuard, not even an explicit nil
### GetHidden

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *PatchedhuggingFaceHuggingFaceDistribution) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *PatchedhuggingFaceHuggingFaceDistribution) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetPulpLabels

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PatchedhuggingFaceHuggingFaceDistribution) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PatchedhuggingFaceHuggingFaceDistribution) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetName

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedhuggingFaceHuggingFaceDistribution) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedhuggingFaceHuggingFaceDistribution) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepository

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PatchedhuggingFaceHuggingFaceDistribution) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PatchedhuggingFaceHuggingFaceDistribution) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *PatchedhuggingFaceHuggingFaceDistribution) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *PatchedhuggingFaceHuggingFaceDistribution) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetRepositoryVersion

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *PatchedhuggingFaceHuggingFaceDistribution) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *PatchedhuggingFaceHuggingFaceDistribution) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### SetRepositoryVersionNil

`func (o *PatchedhuggingFaceHuggingFaceDistribution) SetRepositoryVersionNil(b bool)`

 SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil

### UnsetRepositoryVersion
`func (o *PatchedhuggingFaceHuggingFaceDistribution) UnsetRepositoryVersion()`

UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil
### GetPublication

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetPublication() string`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetPublicationOk() (*string, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublication

`func (o *PatchedhuggingFaceHuggingFaceDistribution) SetPublication(v string)`

SetPublication sets Publication field to given value.

### HasPublication

`func (o *PatchedhuggingFaceHuggingFaceDistribution) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### SetPublicationNil

`func (o *PatchedhuggingFaceHuggingFaceDistribution) SetPublicationNil(b bool)`

 SetPublicationNil sets the value for Publication to be an explicit nil

### UnsetPublication
`func (o *PatchedhuggingFaceHuggingFaceDistribution) UnsetPublication()`

UnsetPublication ensures that no value is present for Publication, not even an explicit nil
### GetRemote

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *PatchedhuggingFaceHuggingFaceDistribution) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *PatchedhuggingFaceHuggingFaceDistribution) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *PatchedhuggingFaceHuggingFaceDistribution) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### SetRemoteNil

`func (o *PatchedhuggingFaceHuggingFaceDistribution) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *PatchedhuggingFaceHuggingFaceDistribution) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


