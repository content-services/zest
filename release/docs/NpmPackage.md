# NpmPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Overwrite** | Pointer to **bool** | When set to true, existing content in the repository with the same unique key will be silently overwritten. When set to false, the task will fail if content would be overwritten. Only used when &#39;repository&#39; is specified. Defaults to true. | [optional] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**Artifact** | Pointer to **string** | Artifact file representing the physical content | [optional] 
**RelativePath** | Pointer to **string** | Path where the artifact is located relative to distributions base_path. If not provided, it will be computed from name and version. | [optional] 
**File** | Pointer to ***os.File** | An uploaded file that may be turned into the content unit. | [optional] 
**Upload** | Pointer to **string** | An uncommitted upload that may be turned into the content unit. | [optional] 
**FileUrl** | Pointer to **string** | A url that Pulp can download and turn into the content unit. | [optional] 
**DownloaderConfig** | Pointer to [**RemoteNetworkConfig**](RemoteNetworkConfig.md) | Configuration for the download process (e.g., proxies, auth, timeouts). Only applicable when providing a &#39;file_url. | [optional] 
**Name** | Pointer to **string** | The name of the npm package. | [optional] 
**Version** | Pointer to **string** | The version of the npm package. | [optional] 

## Methods

### NewNpmPackage

`func NewNpmPackage() *NpmPackage`

NewNpmPackage instantiates a new NpmPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNpmPackageWithDefaults

`func NewNpmPackageWithDefaults() *NpmPackage`

NewNpmPackageWithDefaults instantiates a new NpmPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *NpmPackage) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *NpmPackage) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *NpmPackage) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *NpmPackage) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetOverwrite

`func (o *NpmPackage) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *NpmPackage) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *NpmPackage) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *NpmPackage) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetPulpLabels

`func (o *NpmPackage) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *NpmPackage) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *NpmPackage) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *NpmPackage) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetArtifact

`func (o *NpmPackage) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *NpmPackage) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *NpmPackage) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *NpmPackage) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetRelativePath

`func (o *NpmPackage) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *NpmPackage) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *NpmPackage) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *NpmPackage) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### GetFile

`func (o *NpmPackage) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *NpmPackage) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *NpmPackage) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *NpmPackage) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetUpload

`func (o *NpmPackage) GetUpload() string`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *NpmPackage) GetUploadOk() (*string, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *NpmPackage) SetUpload(v string)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *NpmPackage) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetFileUrl

`func (o *NpmPackage) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *NpmPackage) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *NpmPackage) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *NpmPackage) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetDownloaderConfig

`func (o *NpmPackage) GetDownloaderConfig() RemoteNetworkConfig`

GetDownloaderConfig returns the DownloaderConfig field if non-nil, zero value otherwise.

### GetDownloaderConfigOk

`func (o *NpmPackage) GetDownloaderConfigOk() (*RemoteNetworkConfig, bool)`

GetDownloaderConfigOk returns a tuple with the DownloaderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloaderConfig

`func (o *NpmPackage) SetDownloaderConfig(v RemoteNetworkConfig)`

SetDownloaderConfig sets DownloaderConfig field to given value.

### HasDownloaderConfig

`func (o *NpmPackage) HasDownloaderConfig() bool`

HasDownloaderConfig returns a boolean if a field has been set.

### GetName

`func (o *NpmPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NpmPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NpmPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NpmPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *NpmPackage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NpmPackage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NpmPackage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NpmPackage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


