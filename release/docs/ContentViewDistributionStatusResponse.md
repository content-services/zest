# ContentViewDistributionStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distribution** | Pointer to **string** | The distribution this status entry describes. | [optional] [readonly] 
**Domain** | **string** | The name of the domain the distribution belongs to. | 
**Status** | [**StatusEnum**](StatusEnum.md) | &#39;ok&#39; if the distribution currently resolves to a repository version the caller can search; &#39;no_domain_access&#39; if the caller does not (or no longer) have read access to the distribution&#39;s domain; &#39;no_version&#39; if the distribution or the repository version/publication it pointed to has been deleted.* &#x60;ok&#x60; - ok* &#x60;no_domain_access&#x60; - no_domain_access* &#x60;no_version&#x60; - no_version | 
**RepositoryVersion** | Pointer to **NullableString** | The repository version currently resolved for this distribution, if any. | [optional] [readonly] 

## Methods

### NewContentViewDistributionStatusResponse

`func NewContentViewDistributionStatusResponse(domain string, status StatusEnum, ) *ContentViewDistributionStatusResponse`

NewContentViewDistributionStatusResponse instantiates a new ContentViewDistributionStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentViewDistributionStatusResponseWithDefaults

`func NewContentViewDistributionStatusResponseWithDefaults() *ContentViewDistributionStatusResponse`

NewContentViewDistributionStatusResponseWithDefaults instantiates a new ContentViewDistributionStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistribution

`func (o *ContentViewDistributionStatusResponse) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *ContentViewDistributionStatusResponse) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *ContentViewDistributionStatusResponse) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *ContentViewDistributionStatusResponse) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetDomain

`func (o *ContentViewDistributionStatusResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ContentViewDistributionStatusResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ContentViewDistributionStatusResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetStatus

`func (o *ContentViewDistributionStatusResponse) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContentViewDistributionStatusResponse) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContentViewDistributionStatusResponse) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### GetRepositoryVersion

`func (o *ContentViewDistributionStatusResponse) GetRepositoryVersion() string`

GetRepositoryVersion returns the RepositoryVersion field if non-nil, zero value otherwise.

### GetRepositoryVersionOk

`func (o *ContentViewDistributionStatusResponse) GetRepositoryVersionOk() (*string, bool)`

GetRepositoryVersionOk returns a tuple with the RepositoryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryVersion

`func (o *ContentViewDistributionStatusResponse) SetRepositoryVersion(v string)`

SetRepositoryVersion sets RepositoryVersion field to given value.

### HasRepositoryVersion

`func (o *ContentViewDistributionStatusResponse) HasRepositoryVersion() bool`

HasRepositoryVersion returns a boolean if a field has been set.

### SetRepositoryVersionNil

`func (o *ContentViewDistributionStatusResponse) SetRepositoryVersionNil(b bool)`

 SetRepositoryVersionNil sets the value for RepositoryVersion to be an explicit nil

### UnsetRepositoryVersion
`func (o *ContentViewDistributionStatusResponse) UnsetRepositoryVersion()`

UnsetRepositoryVersion ensures that no value is present for RepositoryVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


