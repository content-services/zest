# MavenMavenArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Overwrite** | Pointer to **bool** | When set to true, existing content in the repository with the same unique key will be silently overwritten. When set to false, the task will fail if content would be overwritten. Only used when &#39;repository&#39; is specified. Defaults to true. | [optional] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**Artifact** | Pointer to **string** | Artifact file representing the physical content | [optional] 
**RelativePath** | **string** | Path where the artifact is located relative to distributions base_path | 
**File** | Pointer to ***os.File** | An uploaded file that may be turned into the content unit. | [optional] 
**Upload** | Pointer to **string** | An uncommitted upload that may be turned into the content unit. | [optional] 
**FileUrl** | Pointer to **string** | A url that Pulp can download and turn into the content unit. | [optional] 
**DownloaderConfig** | Pointer to [**RemoteNetworkConfig**](RemoteNetworkConfig.md) | Configuration for the download process (e.g., proxies, auth, timeouts). Only applicable when providing a &#39;file_url. | [optional] 

## Methods

### NewMavenMavenArtifact

`func NewMavenMavenArtifact(relativePath string, ) *MavenMavenArtifact`

NewMavenMavenArtifact instantiates a new MavenMavenArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenMavenArtifactWithDefaults

`func NewMavenMavenArtifactWithDefaults() *MavenMavenArtifact`

NewMavenMavenArtifactWithDefaults instantiates a new MavenMavenArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *MavenMavenArtifact) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *MavenMavenArtifact) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *MavenMavenArtifact) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *MavenMavenArtifact) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetOverwrite

`func (o *MavenMavenArtifact) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *MavenMavenArtifact) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *MavenMavenArtifact) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *MavenMavenArtifact) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetPulpLabels

`func (o *MavenMavenArtifact) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *MavenMavenArtifact) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *MavenMavenArtifact) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *MavenMavenArtifact) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetArtifact

`func (o *MavenMavenArtifact) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *MavenMavenArtifact) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *MavenMavenArtifact) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *MavenMavenArtifact) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetRelativePath

`func (o *MavenMavenArtifact) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *MavenMavenArtifact) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *MavenMavenArtifact) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.


### GetFile

`func (o *MavenMavenArtifact) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *MavenMavenArtifact) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *MavenMavenArtifact) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *MavenMavenArtifact) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetUpload

`func (o *MavenMavenArtifact) GetUpload() string`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *MavenMavenArtifact) GetUploadOk() (*string, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *MavenMavenArtifact) SetUpload(v string)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *MavenMavenArtifact) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetFileUrl

`func (o *MavenMavenArtifact) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *MavenMavenArtifact) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *MavenMavenArtifact) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *MavenMavenArtifact) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetDownloaderConfig

`func (o *MavenMavenArtifact) GetDownloaderConfig() RemoteNetworkConfig`

GetDownloaderConfig returns the DownloaderConfig field if non-nil, zero value otherwise.

### GetDownloaderConfigOk

`func (o *MavenMavenArtifact) GetDownloaderConfigOk() (*RemoteNetworkConfig, bool)`

GetDownloaderConfigOk returns a tuple with the DownloaderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloaderConfig

`func (o *MavenMavenArtifact) SetDownloaderConfig(v RemoteNetworkConfig)`

SetDownloaderConfig sets DownloaderConfig field to given value.

### HasDownloaderConfig

`func (o *MavenMavenArtifact) HasDownloaderConfig() bool`

HasDownloaderConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


