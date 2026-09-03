# HuggingFaceHuggingFaceContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Overwrite** | Pointer to **bool** | When set to true, existing content in the repository with the same unique key will be silently overwritten. When set to false, the task will fail if content would be overwritten. Only used when &#39;repository&#39; is specified. Defaults to true. | [optional] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**Artifact** | **string** | Artifact file representing the physical content | 
**RelativePath** | **string** | The relative path within the repository | 
**RepoId** | **string** | The Hugging Face repository ID (e.g., &#39;microsoft/DialoGPT-medium&#39;) | 
**RepoType** | Pointer to [**RepoTypeEnum**](RepoTypeEnum.md) | The type of Hugging Face repository* &#x60;models&#x60; - Models* &#x60;datasets&#x60; - Datasets* &#x60;spaces&#x60; - Spaces | [optional] [default to REPOTYPEENUM_MODELS]
**Revision** | Pointer to **string** | The git revision/branch/tag | [optional] [default to "main"]
**Size** | Pointer to **NullableInt64** | File size in bytes | [optional] 
**Etag** | Pointer to **NullableString** | ETag from Hugging Face | [optional] 
**LastModified** | Pointer to **NullableTime** | Last modified timestamp | [optional] 

## Methods

### NewHuggingFaceHuggingFaceContent

`func NewHuggingFaceHuggingFaceContent(artifact string, relativePath string, repoId string, ) *HuggingFaceHuggingFaceContent`

NewHuggingFaceHuggingFaceContent instantiates a new HuggingFaceHuggingFaceContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHuggingFaceHuggingFaceContentWithDefaults

`func NewHuggingFaceHuggingFaceContentWithDefaults() *HuggingFaceHuggingFaceContent`

NewHuggingFaceHuggingFaceContentWithDefaults instantiates a new HuggingFaceHuggingFaceContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *HuggingFaceHuggingFaceContent) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *HuggingFaceHuggingFaceContent) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *HuggingFaceHuggingFaceContent) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *HuggingFaceHuggingFaceContent) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetOverwrite

`func (o *HuggingFaceHuggingFaceContent) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *HuggingFaceHuggingFaceContent) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *HuggingFaceHuggingFaceContent) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *HuggingFaceHuggingFaceContent) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetPulpLabels

`func (o *HuggingFaceHuggingFaceContent) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *HuggingFaceHuggingFaceContent) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *HuggingFaceHuggingFaceContent) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *HuggingFaceHuggingFaceContent) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetArtifact

`func (o *HuggingFaceHuggingFaceContent) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *HuggingFaceHuggingFaceContent) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *HuggingFaceHuggingFaceContent) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.


### GetRelativePath

`func (o *HuggingFaceHuggingFaceContent) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *HuggingFaceHuggingFaceContent) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *HuggingFaceHuggingFaceContent) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.


### GetRepoId

`func (o *HuggingFaceHuggingFaceContent) GetRepoId() string`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *HuggingFaceHuggingFaceContent) GetRepoIdOk() (*string, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *HuggingFaceHuggingFaceContent) SetRepoId(v string)`

SetRepoId sets RepoId field to given value.


### GetRepoType

`func (o *HuggingFaceHuggingFaceContent) GetRepoType() RepoTypeEnum`

GetRepoType returns the RepoType field if non-nil, zero value otherwise.

### GetRepoTypeOk

`func (o *HuggingFaceHuggingFaceContent) GetRepoTypeOk() (*RepoTypeEnum, bool)`

GetRepoTypeOk returns a tuple with the RepoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoType

`func (o *HuggingFaceHuggingFaceContent) SetRepoType(v RepoTypeEnum)`

SetRepoType sets RepoType field to given value.

### HasRepoType

`func (o *HuggingFaceHuggingFaceContent) HasRepoType() bool`

HasRepoType returns a boolean if a field has been set.

### GetRevision

`func (o *HuggingFaceHuggingFaceContent) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *HuggingFaceHuggingFaceContent) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *HuggingFaceHuggingFaceContent) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *HuggingFaceHuggingFaceContent) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSize

`func (o *HuggingFaceHuggingFaceContent) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *HuggingFaceHuggingFaceContent) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *HuggingFaceHuggingFaceContent) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *HuggingFaceHuggingFaceContent) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *HuggingFaceHuggingFaceContent) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *HuggingFaceHuggingFaceContent) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetEtag

`func (o *HuggingFaceHuggingFaceContent) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *HuggingFaceHuggingFaceContent) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *HuggingFaceHuggingFaceContent) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *HuggingFaceHuggingFaceContent) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### SetEtagNil

`func (o *HuggingFaceHuggingFaceContent) SetEtagNil(b bool)`

 SetEtagNil sets the value for Etag to be an explicit nil

### UnsetEtag
`func (o *HuggingFaceHuggingFaceContent) UnsetEtag()`

UnsetEtag ensures that no value is present for Etag, not even an explicit nil
### GetLastModified

`func (o *HuggingFaceHuggingFaceContent) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *HuggingFaceHuggingFaceContent) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *HuggingFaceHuggingFaceContent) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *HuggingFaceHuggingFaceContent) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### SetLastModifiedNil

`func (o *HuggingFaceHuggingFaceContent) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *HuggingFaceHuggingFaceContent) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


