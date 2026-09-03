# PythonPackageProvenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Overwrite** | Pointer to **bool** | When set to true, existing content in the repository with the same unique key will be silently overwritten. When set to false, the task will fail if content would be overwritten. Only used when &#39;repository&#39; is specified. Defaults to true. | [optional] 
**PulpLabels** | Pointer to **map[string]string** | A dictionary of arbitrary key/value pairs used to describe a specific Content instance. | [optional] 
**File** | Pointer to ***os.File** | An uploaded file that may be turned into the content unit. | [optional] 
**Upload** | Pointer to **string** | An uncommitted upload that may be turned into the content unit. | [optional] 
**FileUrl** | Pointer to **string** | A url that Pulp can download and turn into the content unit. | [optional] 
**DownloaderConfig** | Pointer to [**RemoteNetworkConfig**](RemoteNetworkConfig.md) | Configuration for the download process (e.g., proxies, auth, timeouts). Only applicable when providing a &#39;file_url. | [optional] 
**Package** | **string** | The package that the provenance is for. | 
**Verify** | Pointer to **bool** | Verify each attestation in the provenance. | [optional] [default to true]

## Methods

### NewPythonPackageProvenance

`func NewPythonPackageProvenance(package_ string, ) *PythonPackageProvenance`

NewPythonPackageProvenance instantiates a new PythonPackageProvenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPythonPackageProvenanceWithDefaults

`func NewPythonPackageProvenanceWithDefaults() *PythonPackageProvenance`

NewPythonPackageProvenanceWithDefaults instantiates a new PythonPackageProvenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *PythonPackageProvenance) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PythonPackageProvenance) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PythonPackageProvenance) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PythonPackageProvenance) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetOverwrite

`func (o *PythonPackageProvenance) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *PythonPackageProvenance) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *PythonPackageProvenance) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *PythonPackageProvenance) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetPulpLabels

`func (o *PythonPackageProvenance) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *PythonPackageProvenance) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *PythonPackageProvenance) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *PythonPackageProvenance) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetFile

`func (o *PythonPackageProvenance) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *PythonPackageProvenance) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *PythonPackageProvenance) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *PythonPackageProvenance) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetUpload

`func (o *PythonPackageProvenance) GetUpload() string`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *PythonPackageProvenance) GetUploadOk() (*string, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *PythonPackageProvenance) SetUpload(v string)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *PythonPackageProvenance) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetFileUrl

`func (o *PythonPackageProvenance) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *PythonPackageProvenance) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *PythonPackageProvenance) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *PythonPackageProvenance) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetDownloaderConfig

`func (o *PythonPackageProvenance) GetDownloaderConfig() RemoteNetworkConfig`

GetDownloaderConfig returns the DownloaderConfig field if non-nil, zero value otherwise.

### GetDownloaderConfigOk

`func (o *PythonPackageProvenance) GetDownloaderConfigOk() (*RemoteNetworkConfig, bool)`

GetDownloaderConfigOk returns a tuple with the DownloaderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloaderConfig

`func (o *PythonPackageProvenance) SetDownloaderConfig(v RemoteNetworkConfig)`

SetDownloaderConfig sets DownloaderConfig field to given value.

### HasDownloaderConfig

`func (o *PythonPackageProvenance) HasDownloaderConfig() bool`

HasDownloaderConfig returns a boolean if a field has been set.

### GetPackage

`func (o *PythonPackageProvenance) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *PythonPackageProvenance) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *PythonPackageProvenance) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetVerify

`func (o *PythonPackageProvenance) GetVerify() bool`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *PythonPackageProvenance) GetVerifyOk() (*bool, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *PythonPackageProvenance) SetVerify(v bool)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *PythonPackageProvenance) HasVerify() bool`

HasVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


