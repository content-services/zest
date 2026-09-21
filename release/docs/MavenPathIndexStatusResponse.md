# MavenPathIndexStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Ready** | **bool** |  | 
**LatestVersion** | **NullableString** |  | 
**ManifestDigest** | **NullableString** |  | 
**Error** | **string** |  | 

## Methods

### NewMavenPathIndexStatusResponse

`func NewMavenPathIndexStatusResponse(enabled bool, ready bool, latestVersion NullableString, manifestDigest NullableString, error_ string, ) *MavenPathIndexStatusResponse`

NewMavenPathIndexStatusResponse instantiates a new MavenPathIndexStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenPathIndexStatusResponseWithDefaults

`func NewMavenPathIndexStatusResponseWithDefaults() *MavenPathIndexStatusResponse`

NewMavenPathIndexStatusResponseWithDefaults instantiates a new MavenPathIndexStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MavenPathIndexStatusResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MavenPathIndexStatusResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MavenPathIndexStatusResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetReady

`func (o *MavenPathIndexStatusResponse) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *MavenPathIndexStatusResponse) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *MavenPathIndexStatusResponse) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetLatestVersion

`func (o *MavenPathIndexStatusResponse) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *MavenPathIndexStatusResponse) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *MavenPathIndexStatusResponse) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.


### SetLatestVersionNil

`func (o *MavenPathIndexStatusResponse) SetLatestVersionNil(b bool)`

 SetLatestVersionNil sets the value for LatestVersion to be an explicit nil

### UnsetLatestVersion
`func (o *MavenPathIndexStatusResponse) UnsetLatestVersion()`

UnsetLatestVersion ensures that no value is present for LatestVersion, not even an explicit nil
### GetManifestDigest

`func (o *MavenPathIndexStatusResponse) GetManifestDigest() string`

GetManifestDigest returns the ManifestDigest field if non-nil, zero value otherwise.

### GetManifestDigestOk

`func (o *MavenPathIndexStatusResponse) GetManifestDigestOk() (*string, bool)`

GetManifestDigestOk returns a tuple with the ManifestDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestDigest

`func (o *MavenPathIndexStatusResponse) SetManifestDigest(v string)`

SetManifestDigest sets ManifestDigest field to given value.


### SetManifestDigestNil

`func (o *MavenPathIndexStatusResponse) SetManifestDigestNil(b bool)`

 SetManifestDigestNil sets the value for ManifestDigest to be an explicit nil

### UnsetManifestDigest
`func (o *MavenPathIndexStatusResponse) UnsetManifestDigest()`

UnsetManifestDigest ensures that no value is present for ManifestDigest, not even an explicit nil
### GetError

`func (o *MavenPathIndexStatusResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MavenPathIndexStatusResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MavenPathIndexStatusResponse) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


