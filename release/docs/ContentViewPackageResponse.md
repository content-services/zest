# ContentViewPackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** | Name of the package | 
**Epoch** | **string** | The package&#39;s epoch | 
**Version** | **string** | The version of the package | 
**Release** | **string** | The release of the package | 
**Arch** | **string** | The target architecture for the package | 
**Summary** | **string** | Short description of the packaged software | 
**Description** | **string** | In-depth description of the package | 
**ChecksumType** | **string** | Type of checksum, e.g. &#39;sha256&#39; | 
**PkgId** | **string** | Checksum of the package file | 
**Url** | **string** | URL with more information about the package | 
**LocationHref** | **string** | Relative location of the package | 
**IsModular** | **bool** | Whether the package is modular | 

## Methods

### NewContentViewPackageResponse

`func NewContentViewPackageResponse(name string, epoch string, version string, release string, arch string, summary string, description string, checksumType string, pkgId string, url string, locationHref string, isModular bool, ) *ContentViewPackageResponse`

NewContentViewPackageResponse instantiates a new ContentViewPackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentViewPackageResponseWithDefaults

`func NewContentViewPackageResponseWithDefaults() *ContentViewPackageResponse`

NewContentViewPackageResponseWithDefaults instantiates a new ContentViewPackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *ContentViewPackageResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *ContentViewPackageResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *ContentViewPackageResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *ContentViewPackageResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetName

`func (o *ContentViewPackageResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentViewPackageResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentViewPackageResponse) SetName(v string)`

SetName sets Name field to given value.


### GetEpoch

`func (o *ContentViewPackageResponse) GetEpoch() string`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *ContentViewPackageResponse) GetEpochOk() (*string, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *ContentViewPackageResponse) SetEpoch(v string)`

SetEpoch sets Epoch field to given value.


### GetVersion

`func (o *ContentViewPackageResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ContentViewPackageResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ContentViewPackageResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetRelease

`func (o *ContentViewPackageResponse) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *ContentViewPackageResponse) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *ContentViewPackageResponse) SetRelease(v string)`

SetRelease sets Release field to given value.


### GetArch

`func (o *ContentViewPackageResponse) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *ContentViewPackageResponse) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *ContentViewPackageResponse) SetArch(v string)`

SetArch sets Arch field to given value.


### GetSummary

`func (o *ContentViewPackageResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ContentViewPackageResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ContentViewPackageResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetDescription

`func (o *ContentViewPackageResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContentViewPackageResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContentViewPackageResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetChecksumType

`func (o *ContentViewPackageResponse) GetChecksumType() string`

GetChecksumType returns the ChecksumType field if non-nil, zero value otherwise.

### GetChecksumTypeOk

`func (o *ContentViewPackageResponse) GetChecksumTypeOk() (*string, bool)`

GetChecksumTypeOk returns a tuple with the ChecksumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumType

`func (o *ContentViewPackageResponse) SetChecksumType(v string)`

SetChecksumType sets ChecksumType field to given value.


### GetPkgId

`func (o *ContentViewPackageResponse) GetPkgId() string`

GetPkgId returns the PkgId field if non-nil, zero value otherwise.

### GetPkgIdOk

`func (o *ContentViewPackageResponse) GetPkgIdOk() (*string, bool)`

GetPkgIdOk returns a tuple with the PkgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgId

`func (o *ContentViewPackageResponse) SetPkgId(v string)`

SetPkgId sets PkgId field to given value.


### GetUrl

`func (o *ContentViewPackageResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ContentViewPackageResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ContentViewPackageResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLocationHref

`func (o *ContentViewPackageResponse) GetLocationHref() string`

GetLocationHref returns the LocationHref field if non-nil, zero value otherwise.

### GetLocationHrefOk

`func (o *ContentViewPackageResponse) GetLocationHrefOk() (*string, bool)`

GetLocationHrefOk returns a tuple with the LocationHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationHref

`func (o *ContentViewPackageResponse) SetLocationHref(v string)`

SetLocationHref sets LocationHref field to given value.


### GetIsModular

`func (o *ContentViewPackageResponse) GetIsModular() bool`

GetIsModular returns the IsModular field if non-nil, zero value otherwise.

### GetIsModularOk

`func (o *ContentViewPackageResponse) GetIsModularOk() (*bool, bool)`

GetIsModularOk returns a tuple with the IsModular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsModular

`func (o *ContentViewPackageResponse) SetIsModular(v bool)`

SetIsModular sets IsModular field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


