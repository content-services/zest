# UpstreamPulpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the most recent update of the remote. | [optional] [readonly] 
**Name** | **string** | A unique name for this Pulp server. | 
**BaseUrl** | **string** | The transport, hostname, and an optional port of the Pulp server. e.g. https://example.com | 
**ApiRoot** | **string** | The API root. Defaults to &#39;/pulp/&#39;. | 
**Domain** | Pointer to **NullableString** | The domain of the Pulp server if enabled. | [optional] 
**CaCert** | Pointer to **NullableString** | A PEM encoded CA certificate used to validate the server certificate presented by the remote server. | [optional] 
**ClientCert** | Pointer to **NullableString** | A PEM encoded client certificate used for authentication. | [optional] 
**TlsValidation** | Pointer to **bool** | If True, TLS peer validation must be performed. | [optional] 
**DownloadConcurrency** | Pointer to **NullableInt64** | Total number of simultaneous connections. If not set then the default value will be used. | [optional] 
**MaxRetries** | Pointer to **NullableInt64** | Maximum number of retry attempts after a download failure. If not set then the default value (3) will be used. | [optional] 
**TotalTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.total (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**ConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockReadTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_read (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**HiddenFields** | Pointer to [**[]GenericRemoteResponseHiddenFieldsInner**](GenericRemoteResponseHiddenFieldsInner.md) | List of hidden (write only) fields | [optional] [readonly] 
**QSelect** | Pointer to **NullableString** | Filter distributions on the upstream Pulp using complex filtering. E.g. pulp_label_select&#x3D;\&quot;foo\&quot; OR pulp_label_select&#x3D;\&quot;key&#x3D;val\&quot; | [optional] 
**LastReplication** | Pointer to **time.Time** | Timestamp of the last replication that occurred. Equals to &#39;null&#39; if no replication task has been executed. | [optional] [readonly] 
**Policy** | Pointer to [**Policy357Enum**](Policy357Enum.md) | Policy for how replicate will manage the local objects within the domain.* &#x60;all&#x60; - Replicate manages ALL local objects within the domain.* &#x60;labeled&#x60; - Replicate will only manage the objects created from a previous replication, unlabled local objects will be untouched.* &#x60;nodelete&#x60; - Replicate will not delete any local object whether they were created by replication or not. | [optional] 

## Methods

### NewUpstreamPulpResponse

`func NewUpstreamPulpResponse(name string, baseUrl string, apiRoot string, ) *UpstreamPulpResponse`

NewUpstreamPulpResponse instantiates a new UpstreamPulpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamPulpResponseWithDefaults

`func NewUpstreamPulpResponseWithDefaults() *UpstreamPulpResponse`

NewUpstreamPulpResponseWithDefaults instantiates a new UpstreamPulpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *UpstreamPulpResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *UpstreamPulpResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *UpstreamPulpResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *UpstreamPulpResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *UpstreamPulpResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *UpstreamPulpResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *UpstreamPulpResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *UpstreamPulpResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *UpstreamPulpResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *UpstreamPulpResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *UpstreamPulpResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *UpstreamPulpResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *UpstreamPulpResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *UpstreamPulpResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *UpstreamPulpResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *UpstreamPulpResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *UpstreamPulpResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpstreamPulpResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpstreamPulpResponse) SetName(v string)`

SetName sets Name field to given value.


### GetBaseUrl

`func (o *UpstreamPulpResponse) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *UpstreamPulpResponse) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *UpstreamPulpResponse) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetApiRoot

`func (o *UpstreamPulpResponse) GetApiRoot() string`

GetApiRoot returns the ApiRoot field if non-nil, zero value otherwise.

### GetApiRootOk

`func (o *UpstreamPulpResponse) GetApiRootOk() (*string, bool)`

GetApiRootOk returns a tuple with the ApiRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRoot

`func (o *UpstreamPulpResponse) SetApiRoot(v string)`

SetApiRoot sets ApiRoot field to given value.


### GetDomain

`func (o *UpstreamPulpResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UpstreamPulpResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UpstreamPulpResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UpstreamPulpResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *UpstreamPulpResponse) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *UpstreamPulpResponse) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetCaCert

`func (o *UpstreamPulpResponse) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *UpstreamPulpResponse) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *UpstreamPulpResponse) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *UpstreamPulpResponse) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *UpstreamPulpResponse) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *UpstreamPulpResponse) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetClientCert

`func (o *UpstreamPulpResponse) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *UpstreamPulpResponse) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *UpstreamPulpResponse) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *UpstreamPulpResponse) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### SetClientCertNil

`func (o *UpstreamPulpResponse) SetClientCertNil(b bool)`

 SetClientCertNil sets the value for ClientCert to be an explicit nil

### UnsetClientCert
`func (o *UpstreamPulpResponse) UnsetClientCert()`

UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
### GetTlsValidation

`func (o *UpstreamPulpResponse) GetTlsValidation() bool`

GetTlsValidation returns the TlsValidation field if non-nil, zero value otherwise.

### GetTlsValidationOk

`func (o *UpstreamPulpResponse) GetTlsValidationOk() (*bool, bool)`

GetTlsValidationOk returns a tuple with the TlsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsValidation

`func (o *UpstreamPulpResponse) SetTlsValidation(v bool)`

SetTlsValidation sets TlsValidation field to given value.

### HasTlsValidation

`func (o *UpstreamPulpResponse) HasTlsValidation() bool`

HasTlsValidation returns a boolean if a field has been set.

### GetDownloadConcurrency

`func (o *UpstreamPulpResponse) GetDownloadConcurrency() int64`

GetDownloadConcurrency returns the DownloadConcurrency field if non-nil, zero value otherwise.

### GetDownloadConcurrencyOk

`func (o *UpstreamPulpResponse) GetDownloadConcurrencyOk() (*int64, bool)`

GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadConcurrency

`func (o *UpstreamPulpResponse) SetDownloadConcurrency(v int64)`

SetDownloadConcurrency sets DownloadConcurrency field to given value.

### HasDownloadConcurrency

`func (o *UpstreamPulpResponse) HasDownloadConcurrency() bool`

HasDownloadConcurrency returns a boolean if a field has been set.

### SetDownloadConcurrencyNil

`func (o *UpstreamPulpResponse) SetDownloadConcurrencyNil(b bool)`

 SetDownloadConcurrencyNil sets the value for DownloadConcurrency to be an explicit nil

### UnsetDownloadConcurrency
`func (o *UpstreamPulpResponse) UnsetDownloadConcurrency()`

UnsetDownloadConcurrency ensures that no value is present for DownloadConcurrency, not even an explicit nil
### GetMaxRetries

`func (o *UpstreamPulpResponse) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *UpstreamPulpResponse) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *UpstreamPulpResponse) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *UpstreamPulpResponse) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *UpstreamPulpResponse) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *UpstreamPulpResponse) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetTotalTimeout

`func (o *UpstreamPulpResponse) GetTotalTimeout() float64`

GetTotalTimeout returns the TotalTimeout field if non-nil, zero value otherwise.

### GetTotalTimeoutOk

`func (o *UpstreamPulpResponse) GetTotalTimeoutOk() (*float64, bool)`

GetTotalTimeoutOk returns a tuple with the TotalTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeout

`func (o *UpstreamPulpResponse) SetTotalTimeout(v float64)`

SetTotalTimeout sets TotalTimeout field to given value.

### HasTotalTimeout

`func (o *UpstreamPulpResponse) HasTotalTimeout() bool`

HasTotalTimeout returns a boolean if a field has been set.

### SetTotalTimeoutNil

`func (o *UpstreamPulpResponse) SetTotalTimeoutNil(b bool)`

 SetTotalTimeoutNil sets the value for TotalTimeout to be an explicit nil

### UnsetTotalTimeout
`func (o *UpstreamPulpResponse) UnsetTotalTimeout()`

UnsetTotalTimeout ensures that no value is present for TotalTimeout, not even an explicit nil
### GetConnectTimeout

`func (o *UpstreamPulpResponse) GetConnectTimeout() float64`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *UpstreamPulpResponse) GetConnectTimeoutOk() (*float64, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *UpstreamPulpResponse) SetConnectTimeout(v float64)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *UpstreamPulpResponse) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### SetConnectTimeoutNil

`func (o *UpstreamPulpResponse) SetConnectTimeoutNil(b bool)`

 SetConnectTimeoutNil sets the value for ConnectTimeout to be an explicit nil

### UnsetConnectTimeout
`func (o *UpstreamPulpResponse) UnsetConnectTimeout()`

UnsetConnectTimeout ensures that no value is present for ConnectTimeout, not even an explicit nil
### GetSockConnectTimeout

`func (o *UpstreamPulpResponse) GetSockConnectTimeout() float64`

GetSockConnectTimeout returns the SockConnectTimeout field if non-nil, zero value otherwise.

### GetSockConnectTimeoutOk

`func (o *UpstreamPulpResponse) GetSockConnectTimeoutOk() (*float64, bool)`

GetSockConnectTimeoutOk returns a tuple with the SockConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockConnectTimeout

`func (o *UpstreamPulpResponse) SetSockConnectTimeout(v float64)`

SetSockConnectTimeout sets SockConnectTimeout field to given value.

### HasSockConnectTimeout

`func (o *UpstreamPulpResponse) HasSockConnectTimeout() bool`

HasSockConnectTimeout returns a boolean if a field has been set.

### SetSockConnectTimeoutNil

`func (o *UpstreamPulpResponse) SetSockConnectTimeoutNil(b bool)`

 SetSockConnectTimeoutNil sets the value for SockConnectTimeout to be an explicit nil

### UnsetSockConnectTimeout
`func (o *UpstreamPulpResponse) UnsetSockConnectTimeout()`

UnsetSockConnectTimeout ensures that no value is present for SockConnectTimeout, not even an explicit nil
### GetSockReadTimeout

`func (o *UpstreamPulpResponse) GetSockReadTimeout() float64`

GetSockReadTimeout returns the SockReadTimeout field if non-nil, zero value otherwise.

### GetSockReadTimeoutOk

`func (o *UpstreamPulpResponse) GetSockReadTimeoutOk() (*float64, bool)`

GetSockReadTimeoutOk returns a tuple with the SockReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockReadTimeout

`func (o *UpstreamPulpResponse) SetSockReadTimeout(v float64)`

SetSockReadTimeout sets SockReadTimeout field to given value.

### HasSockReadTimeout

`func (o *UpstreamPulpResponse) HasSockReadTimeout() bool`

HasSockReadTimeout returns a boolean if a field has been set.

### SetSockReadTimeoutNil

`func (o *UpstreamPulpResponse) SetSockReadTimeoutNil(b bool)`

 SetSockReadTimeoutNil sets the value for SockReadTimeout to be an explicit nil

### UnsetSockReadTimeout
`func (o *UpstreamPulpResponse) UnsetSockReadTimeout()`

UnsetSockReadTimeout ensures that no value is present for SockReadTimeout, not even an explicit nil
### GetHiddenFields

`func (o *UpstreamPulpResponse) GetHiddenFields() []GenericRemoteResponseHiddenFieldsInner`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *UpstreamPulpResponse) GetHiddenFieldsOk() (*[]GenericRemoteResponseHiddenFieldsInner, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *UpstreamPulpResponse) SetHiddenFields(v []GenericRemoteResponseHiddenFieldsInner)`

SetHiddenFields sets HiddenFields field to given value.

### HasHiddenFields

`func (o *UpstreamPulpResponse) HasHiddenFields() bool`

HasHiddenFields returns a boolean if a field has been set.

### GetQSelect

`func (o *UpstreamPulpResponse) GetQSelect() string`

GetQSelect returns the QSelect field if non-nil, zero value otherwise.

### GetQSelectOk

`func (o *UpstreamPulpResponse) GetQSelectOk() (*string, bool)`

GetQSelectOk returns a tuple with the QSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQSelect

`func (o *UpstreamPulpResponse) SetQSelect(v string)`

SetQSelect sets QSelect field to given value.

### HasQSelect

`func (o *UpstreamPulpResponse) HasQSelect() bool`

HasQSelect returns a boolean if a field has been set.

### SetQSelectNil

`func (o *UpstreamPulpResponse) SetQSelectNil(b bool)`

 SetQSelectNil sets the value for QSelect to be an explicit nil

### UnsetQSelect
`func (o *UpstreamPulpResponse) UnsetQSelect()`

UnsetQSelect ensures that no value is present for QSelect, not even an explicit nil
### GetLastReplication

`func (o *UpstreamPulpResponse) GetLastReplication() time.Time`

GetLastReplication returns the LastReplication field if non-nil, zero value otherwise.

### GetLastReplicationOk

`func (o *UpstreamPulpResponse) GetLastReplicationOk() (*time.Time, bool)`

GetLastReplicationOk returns a tuple with the LastReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplication

`func (o *UpstreamPulpResponse) SetLastReplication(v time.Time)`

SetLastReplication sets LastReplication field to given value.

### HasLastReplication

`func (o *UpstreamPulpResponse) HasLastReplication() bool`

HasLastReplication returns a boolean if a field has been set.

### GetPolicy

`func (o *UpstreamPulpResponse) GetPolicy() Policy357Enum`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UpstreamPulpResponse) GetPolicyOk() (*Policy357Enum, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UpstreamPulpResponse) SetPolicy(v Policy357Enum)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *UpstreamPulpResponse) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


