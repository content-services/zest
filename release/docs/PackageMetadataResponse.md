# PackageMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSerial** | **int64** | Cache value from last PyPI sync | 
**Info** | **interface{}** | Core metadata of the package | 
**Releases** | **interface{}** | List of all the releases of the package | 
**Urls** | **interface{}** |  | 

## Methods

### NewPackageMetadataResponse

`func NewPackageMetadataResponse(lastSerial int64, info interface{}, releases interface{}, urls interface{}, ) *PackageMetadataResponse`

NewPackageMetadataResponse instantiates a new PackageMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageMetadataResponseWithDefaults

`func NewPackageMetadataResponseWithDefaults() *PackageMetadataResponse`

NewPackageMetadataResponseWithDefaults instantiates a new PackageMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastSerial

`func (o *PackageMetadataResponse) GetLastSerial() int64`

GetLastSerial returns the LastSerial field if non-nil, zero value otherwise.

### GetLastSerialOk

`func (o *PackageMetadataResponse) GetLastSerialOk() (*int64, bool)`

GetLastSerialOk returns a tuple with the LastSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSerial

`func (o *PackageMetadataResponse) SetLastSerial(v int64)`

SetLastSerial sets LastSerial field to given value.


### GetInfo

`func (o *PackageMetadataResponse) GetInfo() interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PackageMetadataResponse) GetInfoOk() (*interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PackageMetadataResponse) SetInfo(v interface{})`

SetInfo sets Info field to given value.


### SetInfoNil

`func (o *PackageMetadataResponse) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *PackageMetadataResponse) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil
### GetReleases

`func (o *PackageMetadataResponse) GetReleases() interface{}`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *PackageMetadataResponse) GetReleasesOk() (*interface{}, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *PackageMetadataResponse) SetReleases(v interface{})`

SetReleases sets Releases field to given value.


### SetReleasesNil

`func (o *PackageMetadataResponse) SetReleasesNil(b bool)`

 SetReleasesNil sets the value for Releases to be an explicit nil

### UnsetReleases
`func (o *PackageMetadataResponse) UnsetReleases()`

UnsetReleases ensures that no value is present for Releases, not even an explicit nil
### GetUrls

`func (o *PackageMetadataResponse) GetUrls() interface{}`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *PackageMetadataResponse) GetUrlsOk() (*interface{}, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *PackageMetadataResponse) SetUrls(v interface{})`

SetUrls sets Urls field to given value.


### SetUrlsNil

`func (o *PackageMetadataResponse) SetUrlsNil(b bool)`

 SetUrlsNil sets the value for Urls to be an explicit nil

### UnsetUrls
`func (o *PackageMetadataResponse) UnsetUrls()`

UnsetUrls ensures that no value is present for Urls, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


