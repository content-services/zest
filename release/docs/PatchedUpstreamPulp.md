# PatchedUpstreamPulp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name for this Pulp server. | [optional] 
**BaseUrl** | Pointer to **string** | The transport, hostname, and an optional port of the Pulp server. e.g. https://example.com | [optional] 
**ApiRoot** | Pointer to **string** | The API root. Defaults to &#39;/pulp/&#39;. | [optional] 
**Domain** | Pointer to **NullableString** | The domain of the Pulp server if enabled. | [optional] 
**CaCert** | Pointer to **NullableString** | A PEM encoded CA certificate used to validate the server certificate presented by the remote server. | [optional] 
**ClientCert** | Pointer to **NullableString** | A PEM encoded client certificate used for authentication. | [optional] 
**ClientKey** | Pointer to **NullableString** | A PEM encoded private key used for authentication. | [optional] 
**TlsValidation** | Pointer to **bool** | If True, TLS peer validation must be performed. | [optional] 
**Username** | Pointer to **NullableString** | The username to be used for authentication when syncing. | [optional] 
**Password** | Pointer to **NullableString** | The password to be used for authentication when syncing. Extra leading and trailing whitespace characters are not trimmed. | [optional] 
**DownloadConcurrency** | Pointer to **NullableInt64** | Total number of simultaneous connections. If not set then the default value will be used. | [optional] 
**MaxRetries** | Pointer to **NullableInt64** | Maximum number of retry attempts after a download failure. If not set then the default value (3) will be used. | [optional] 
**TotalTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.total (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**ConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockReadTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_read (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**QSelect** | Pointer to **NullableString** | Filter distributions on the upstream Pulp using complex filtering. E.g. pulp_label_select&#x3D;\&quot;foo\&quot; OR pulp_label_select&#x3D;\&quot;key&#x3D;val\&quot; | [optional] 
**Policy** | Pointer to [**Policy357Enum**](Policy357Enum.md) | Policy for how replicate will manage the local objects within the domain.* &#x60;all&#x60; - Replicate manages ALL local objects within the domain.* &#x60;labeled&#x60; - Replicate will only manage the objects created from a previous replication, unlabled local objects will be untouched.* &#x60;nodelete&#x60; - Replicate will not delete any local object whether they were created by replication or not. | [optional] 

## Methods

### NewPatchedUpstreamPulp

`func NewPatchedUpstreamPulp() *PatchedUpstreamPulp`

NewPatchedUpstreamPulp instantiates a new PatchedUpstreamPulp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedUpstreamPulpWithDefaults

`func NewPatchedUpstreamPulpWithDefaults() *PatchedUpstreamPulp`

NewPatchedUpstreamPulpWithDefaults instantiates a new PatchedUpstreamPulp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedUpstreamPulp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedUpstreamPulp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedUpstreamPulp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedUpstreamPulp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBaseUrl

`func (o *PatchedUpstreamPulp) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *PatchedUpstreamPulp) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *PatchedUpstreamPulp) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *PatchedUpstreamPulp) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetApiRoot

`func (o *PatchedUpstreamPulp) GetApiRoot() string`

GetApiRoot returns the ApiRoot field if non-nil, zero value otherwise.

### GetApiRootOk

`func (o *PatchedUpstreamPulp) GetApiRootOk() (*string, bool)`

GetApiRootOk returns a tuple with the ApiRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRoot

`func (o *PatchedUpstreamPulp) SetApiRoot(v string)`

SetApiRoot sets ApiRoot field to given value.

### HasApiRoot

`func (o *PatchedUpstreamPulp) HasApiRoot() bool`

HasApiRoot returns a boolean if a field has been set.

### GetDomain

`func (o *PatchedUpstreamPulp) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PatchedUpstreamPulp) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PatchedUpstreamPulp) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PatchedUpstreamPulp) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *PatchedUpstreamPulp) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *PatchedUpstreamPulp) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetCaCert

`func (o *PatchedUpstreamPulp) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *PatchedUpstreamPulp) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *PatchedUpstreamPulp) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *PatchedUpstreamPulp) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *PatchedUpstreamPulp) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *PatchedUpstreamPulp) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetClientCert

`func (o *PatchedUpstreamPulp) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *PatchedUpstreamPulp) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *PatchedUpstreamPulp) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *PatchedUpstreamPulp) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### SetClientCertNil

`func (o *PatchedUpstreamPulp) SetClientCertNil(b bool)`

 SetClientCertNil sets the value for ClientCert to be an explicit nil

### UnsetClientCert
`func (o *PatchedUpstreamPulp) UnsetClientCert()`

UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
### GetClientKey

`func (o *PatchedUpstreamPulp) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *PatchedUpstreamPulp) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *PatchedUpstreamPulp) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *PatchedUpstreamPulp) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### SetClientKeyNil

`func (o *PatchedUpstreamPulp) SetClientKeyNil(b bool)`

 SetClientKeyNil sets the value for ClientKey to be an explicit nil

### UnsetClientKey
`func (o *PatchedUpstreamPulp) UnsetClientKey()`

UnsetClientKey ensures that no value is present for ClientKey, not even an explicit nil
### GetTlsValidation

`func (o *PatchedUpstreamPulp) GetTlsValidation() bool`

GetTlsValidation returns the TlsValidation field if non-nil, zero value otherwise.

### GetTlsValidationOk

`func (o *PatchedUpstreamPulp) GetTlsValidationOk() (*bool, bool)`

GetTlsValidationOk returns a tuple with the TlsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsValidation

`func (o *PatchedUpstreamPulp) SetTlsValidation(v bool)`

SetTlsValidation sets TlsValidation field to given value.

### HasTlsValidation

`func (o *PatchedUpstreamPulp) HasTlsValidation() bool`

HasTlsValidation returns a boolean if a field has been set.

### GetUsername

`func (o *PatchedUpstreamPulp) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedUpstreamPulp) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedUpstreamPulp) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedUpstreamPulp) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *PatchedUpstreamPulp) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *PatchedUpstreamPulp) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *PatchedUpstreamPulp) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchedUpstreamPulp) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchedUpstreamPulp) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchedUpstreamPulp) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PatchedUpstreamPulp) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PatchedUpstreamPulp) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDownloadConcurrency

`func (o *PatchedUpstreamPulp) GetDownloadConcurrency() int64`

GetDownloadConcurrency returns the DownloadConcurrency field if non-nil, zero value otherwise.

### GetDownloadConcurrencyOk

`func (o *PatchedUpstreamPulp) GetDownloadConcurrencyOk() (*int64, bool)`

GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadConcurrency

`func (o *PatchedUpstreamPulp) SetDownloadConcurrency(v int64)`

SetDownloadConcurrency sets DownloadConcurrency field to given value.

### HasDownloadConcurrency

`func (o *PatchedUpstreamPulp) HasDownloadConcurrency() bool`

HasDownloadConcurrency returns a boolean if a field has been set.

### SetDownloadConcurrencyNil

`func (o *PatchedUpstreamPulp) SetDownloadConcurrencyNil(b bool)`

 SetDownloadConcurrencyNil sets the value for DownloadConcurrency to be an explicit nil

### UnsetDownloadConcurrency
`func (o *PatchedUpstreamPulp) UnsetDownloadConcurrency()`

UnsetDownloadConcurrency ensures that no value is present for DownloadConcurrency, not even an explicit nil
### GetMaxRetries

`func (o *PatchedUpstreamPulp) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *PatchedUpstreamPulp) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *PatchedUpstreamPulp) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *PatchedUpstreamPulp) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *PatchedUpstreamPulp) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *PatchedUpstreamPulp) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetTotalTimeout

`func (o *PatchedUpstreamPulp) GetTotalTimeout() float64`

GetTotalTimeout returns the TotalTimeout field if non-nil, zero value otherwise.

### GetTotalTimeoutOk

`func (o *PatchedUpstreamPulp) GetTotalTimeoutOk() (*float64, bool)`

GetTotalTimeoutOk returns a tuple with the TotalTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeout

`func (o *PatchedUpstreamPulp) SetTotalTimeout(v float64)`

SetTotalTimeout sets TotalTimeout field to given value.

### HasTotalTimeout

`func (o *PatchedUpstreamPulp) HasTotalTimeout() bool`

HasTotalTimeout returns a boolean if a field has been set.

### SetTotalTimeoutNil

`func (o *PatchedUpstreamPulp) SetTotalTimeoutNil(b bool)`

 SetTotalTimeoutNil sets the value for TotalTimeout to be an explicit nil

### UnsetTotalTimeout
`func (o *PatchedUpstreamPulp) UnsetTotalTimeout()`

UnsetTotalTimeout ensures that no value is present for TotalTimeout, not even an explicit nil
### GetConnectTimeout

`func (o *PatchedUpstreamPulp) GetConnectTimeout() float64`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *PatchedUpstreamPulp) GetConnectTimeoutOk() (*float64, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *PatchedUpstreamPulp) SetConnectTimeout(v float64)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *PatchedUpstreamPulp) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### SetConnectTimeoutNil

`func (o *PatchedUpstreamPulp) SetConnectTimeoutNil(b bool)`

 SetConnectTimeoutNil sets the value for ConnectTimeout to be an explicit nil

### UnsetConnectTimeout
`func (o *PatchedUpstreamPulp) UnsetConnectTimeout()`

UnsetConnectTimeout ensures that no value is present for ConnectTimeout, not even an explicit nil
### GetSockConnectTimeout

`func (o *PatchedUpstreamPulp) GetSockConnectTimeout() float64`

GetSockConnectTimeout returns the SockConnectTimeout field if non-nil, zero value otherwise.

### GetSockConnectTimeoutOk

`func (o *PatchedUpstreamPulp) GetSockConnectTimeoutOk() (*float64, bool)`

GetSockConnectTimeoutOk returns a tuple with the SockConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockConnectTimeout

`func (o *PatchedUpstreamPulp) SetSockConnectTimeout(v float64)`

SetSockConnectTimeout sets SockConnectTimeout field to given value.

### HasSockConnectTimeout

`func (o *PatchedUpstreamPulp) HasSockConnectTimeout() bool`

HasSockConnectTimeout returns a boolean if a field has been set.

### SetSockConnectTimeoutNil

`func (o *PatchedUpstreamPulp) SetSockConnectTimeoutNil(b bool)`

 SetSockConnectTimeoutNil sets the value for SockConnectTimeout to be an explicit nil

### UnsetSockConnectTimeout
`func (o *PatchedUpstreamPulp) UnsetSockConnectTimeout()`

UnsetSockConnectTimeout ensures that no value is present for SockConnectTimeout, not even an explicit nil
### GetSockReadTimeout

`func (o *PatchedUpstreamPulp) GetSockReadTimeout() float64`

GetSockReadTimeout returns the SockReadTimeout field if non-nil, zero value otherwise.

### GetSockReadTimeoutOk

`func (o *PatchedUpstreamPulp) GetSockReadTimeoutOk() (*float64, bool)`

GetSockReadTimeoutOk returns a tuple with the SockReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockReadTimeout

`func (o *PatchedUpstreamPulp) SetSockReadTimeout(v float64)`

SetSockReadTimeout sets SockReadTimeout field to given value.

### HasSockReadTimeout

`func (o *PatchedUpstreamPulp) HasSockReadTimeout() bool`

HasSockReadTimeout returns a boolean if a field has been set.

### SetSockReadTimeoutNil

`func (o *PatchedUpstreamPulp) SetSockReadTimeoutNil(b bool)`

 SetSockReadTimeoutNil sets the value for SockReadTimeout to be an explicit nil

### UnsetSockReadTimeout
`func (o *PatchedUpstreamPulp) UnsetSockReadTimeout()`

UnsetSockReadTimeout ensures that no value is present for SockReadTimeout, not even an explicit nil
### GetQSelect

`func (o *PatchedUpstreamPulp) GetQSelect() string`

GetQSelect returns the QSelect field if non-nil, zero value otherwise.

### GetQSelectOk

`func (o *PatchedUpstreamPulp) GetQSelectOk() (*string, bool)`

GetQSelectOk returns a tuple with the QSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQSelect

`func (o *PatchedUpstreamPulp) SetQSelect(v string)`

SetQSelect sets QSelect field to given value.

### HasQSelect

`func (o *PatchedUpstreamPulp) HasQSelect() bool`

HasQSelect returns a boolean if a field has been set.

### SetQSelectNil

`func (o *PatchedUpstreamPulp) SetQSelectNil(b bool)`

 SetQSelectNil sets the value for QSelect to be an explicit nil

### UnsetQSelect
`func (o *PatchedUpstreamPulp) UnsetQSelect()`

UnsetQSelect ensures that no value is present for QSelect, not even an explicit nil
### GetPolicy

`func (o *PatchedUpstreamPulp) GetPolicy() Policy357Enum`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PatchedUpstreamPulp) GetPolicyOk() (*Policy357Enum, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PatchedUpstreamPulp) SetPolicy(v Policy357Enum)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PatchedUpstreamPulp) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


