# HuggingFaceHuggingFaceRemoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Prn** | Pointer to **string** | The Pulp Resource Name (PRN). | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**PulpLastUpdated** | Pointer to **time.Time** | Timestamp of the most recent update of the remote. | [optional] [readonly] 
**Name** | **string** | A unique name for this remote. | 
**Url** | **string** | The URL of an external content source. | 
**PulpLabels** | Pointer to **map[string]string** |  | [optional] 
**Policy** | Pointer to [**Policy692Enum**](Policy692Enum.md) | The policy to use when downloading content. The possible values include: &#39;immediate&#39;, &#39;on_demand&#39;, and &#39;streamed&#39;. &#39;on_demand&#39; enables pull-through caching.* &#x60;immediate&#x60; - When syncing, download all metadata and content now.* &#x60;on_demand&#x60; - When syncing, download metadata, but do not download content now. Instead, download content as clients request it, and save it in Pulp to be served for future client requests.* &#x60;streamed&#x60; - When syncing, download metadata, but do not download content now. Instead,download content as clients request it, but never save it in Pulp. This causes future requests for that same content to have to be downloaded again. | [optional] [default to POLICY692ENUM_ON_DEMAND]
**HiddenFields** | Pointer to [**[]GenericRemoteResponseHiddenFieldsInner**](GenericRemoteResponseHiddenFieldsInner.md) | List of hidden (write only) fields | [optional] [readonly] 
**CaCert** | Pointer to **NullableString** | A PEM encoded CA certificate used to validate the server certificate presented by the remote server. | [optional] 
**ClientCert** | Pointer to **NullableString** | A PEM encoded client certificate used for authentication. | [optional] 
**TlsValidation** | Pointer to **bool** | If True, TLS peer validation must be performed. | [optional] 
**ProxyUrl** | Pointer to **NullableString** | The proxy URL. Format: scheme://host:port | [optional] 
**MaxRetries** | Pointer to **NullableInt64** | Maximum number of retry attempts after a download failure. If not set then the default value (3) will be used. | [optional] 
**TotalTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.total (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**ConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockConnectTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**SockReadTimeout** | Pointer to **NullableFloat64** | aiohttp.ClientTimeout.sock_read (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used. | [optional] 
**Headers** | Pointer to **[]map[string]interface{}** | Headers for aiohttp.Clientsession | [optional] 
**DownloadConcurrency** | Pointer to **NullableInt64** | Total number of simultaneous connections. If not set then the default value will be used. | [optional] 
**RateLimit** | Pointer to **NullableInt64** | Limits requests per second for each concurrent downloader | [optional] 
**HfHubUrl** | Pointer to **string** | Base URL for Hugging Face Hub | [optional] [default to "https://huggingface.co"]
**HfToken** | Pointer to **NullableString** | Hugging Face authentication token for private repositories | [optional] 

## Methods

### NewHuggingFaceHuggingFaceRemoteResponse

`func NewHuggingFaceHuggingFaceRemoteResponse(name string, url string, ) *HuggingFaceHuggingFaceRemoteResponse`

NewHuggingFaceHuggingFaceRemoteResponse instantiates a new HuggingFaceHuggingFaceRemoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHuggingFaceHuggingFaceRemoteResponseWithDefaults

`func NewHuggingFaceHuggingFaceRemoteResponseWithDefaults() *HuggingFaceHuggingFaceRemoteResponse`

NewHuggingFaceHuggingFaceRemoteResponseWithDefaults instantiates a new HuggingFaceHuggingFaceRemoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPrn

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetPrn() string`

GetPrn returns the Prn field if non-nil, zero value otherwise.

### GetPrnOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetPrnOk() (*string, bool)`

GetPrnOk returns a tuple with the Prn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrn

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetPrn(v string)`

SetPrn sets Prn field to given value.

### HasPrn

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasPrn() bool`

HasPrn returns a boolean if a field has been set.

### GetPulpCreated

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetPulpLastUpdated

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetPulpLastUpdated() time.Time`

GetPulpLastUpdated returns the PulpLastUpdated field if non-nil, zero value otherwise.

### GetPulpLastUpdatedOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetPulpLastUpdatedOk() (*time.Time, bool)`

GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLastUpdated

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetPulpLastUpdated(v time.Time)`

SetPulpLastUpdated sets PulpLastUpdated field to given value.

### HasPulpLastUpdated

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasPulpLastUpdated() bool`

HasPulpLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPulpLabels

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetPulpLabels() map[string]*string`

GetPulpLabels returns the PulpLabels field if non-nil, zero value otherwise.

### GetPulpLabelsOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetPulpLabelsOk() (*map[string]*string, bool)`

GetPulpLabelsOk returns a tuple with the PulpLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabels

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetPulpLabels(v map[string]*string)`

SetPulpLabels sets PulpLabels field to given value.

### HasPulpLabels

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasPulpLabels() bool`

HasPulpLabels returns a boolean if a field has been set.

### GetPolicy

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetPolicy() Policy692Enum`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetPolicyOk() (*Policy692Enum, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetPolicy(v Policy692Enum)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetHiddenFields

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetHiddenFields() []GenericRemoteResponseHiddenFieldsInner`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetHiddenFieldsOk() (*[]GenericRemoteResponseHiddenFieldsInner, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetHiddenFields(v []GenericRemoteResponseHiddenFieldsInner)`

SetHiddenFields sets HiddenFields field to given value.

### HasHiddenFields

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasHiddenFields() bool`

HasHiddenFields returns a boolean if a field has been set.

### GetCaCert

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *HuggingFaceHuggingFaceRemoteResponse) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetClientCert

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### SetClientCertNil

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetClientCertNil(b bool)`

 SetClientCertNil sets the value for ClientCert to be an explicit nil

### UnsetClientCert
`func (o *HuggingFaceHuggingFaceRemoteResponse) UnsetClientCert()`

UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
### GetTlsValidation

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetTlsValidation() bool`

GetTlsValidation returns the TlsValidation field if non-nil, zero value otherwise.

### GetTlsValidationOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetTlsValidationOk() (*bool, bool)`

GetTlsValidationOk returns a tuple with the TlsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsValidation

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetTlsValidation(v bool)`

SetTlsValidation sets TlsValidation field to given value.

### HasTlsValidation

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasTlsValidation() bool`

HasTlsValidation returns a boolean if a field has been set.

### GetProxyUrl

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

### SetProxyUrlNil

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetProxyUrlNil(b bool)`

 SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil

### UnsetProxyUrl
`func (o *HuggingFaceHuggingFaceRemoteResponse) UnsetProxyUrl()`

UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
### GetMaxRetries

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *HuggingFaceHuggingFaceRemoteResponse) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetTotalTimeout

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetTotalTimeout() float64`

GetTotalTimeout returns the TotalTimeout field if non-nil, zero value otherwise.

### GetTotalTimeoutOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetTotalTimeoutOk() (*float64, bool)`

GetTotalTimeoutOk returns a tuple with the TotalTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeout

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetTotalTimeout(v float64)`

SetTotalTimeout sets TotalTimeout field to given value.

### HasTotalTimeout

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasTotalTimeout() bool`

HasTotalTimeout returns a boolean if a field has been set.

### SetTotalTimeoutNil

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetTotalTimeoutNil(b bool)`

 SetTotalTimeoutNil sets the value for TotalTimeout to be an explicit nil

### UnsetTotalTimeout
`func (o *HuggingFaceHuggingFaceRemoteResponse) UnsetTotalTimeout()`

UnsetTotalTimeout ensures that no value is present for TotalTimeout, not even an explicit nil
### GetConnectTimeout

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetConnectTimeout() float64`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetConnectTimeoutOk() (*float64, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetConnectTimeout(v float64)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### SetConnectTimeoutNil

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetConnectTimeoutNil(b bool)`

 SetConnectTimeoutNil sets the value for ConnectTimeout to be an explicit nil

### UnsetConnectTimeout
`func (o *HuggingFaceHuggingFaceRemoteResponse) UnsetConnectTimeout()`

UnsetConnectTimeout ensures that no value is present for ConnectTimeout, not even an explicit nil
### GetSockConnectTimeout

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetSockConnectTimeout() float64`

GetSockConnectTimeout returns the SockConnectTimeout field if non-nil, zero value otherwise.

### GetSockConnectTimeoutOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetSockConnectTimeoutOk() (*float64, bool)`

GetSockConnectTimeoutOk returns a tuple with the SockConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockConnectTimeout

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetSockConnectTimeout(v float64)`

SetSockConnectTimeout sets SockConnectTimeout field to given value.

### HasSockConnectTimeout

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasSockConnectTimeout() bool`

HasSockConnectTimeout returns a boolean if a field has been set.

### SetSockConnectTimeoutNil

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetSockConnectTimeoutNil(b bool)`

 SetSockConnectTimeoutNil sets the value for SockConnectTimeout to be an explicit nil

### UnsetSockConnectTimeout
`func (o *HuggingFaceHuggingFaceRemoteResponse) UnsetSockConnectTimeout()`

UnsetSockConnectTimeout ensures that no value is present for SockConnectTimeout, not even an explicit nil
### GetSockReadTimeout

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetSockReadTimeout() float64`

GetSockReadTimeout returns the SockReadTimeout field if non-nil, zero value otherwise.

### GetSockReadTimeoutOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetSockReadTimeoutOk() (*float64, bool)`

GetSockReadTimeoutOk returns a tuple with the SockReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockReadTimeout

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetSockReadTimeout(v float64)`

SetSockReadTimeout sets SockReadTimeout field to given value.

### HasSockReadTimeout

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasSockReadTimeout() bool`

HasSockReadTimeout returns a boolean if a field has been set.

### SetSockReadTimeoutNil

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetSockReadTimeoutNil(b bool)`

 SetSockReadTimeoutNil sets the value for SockReadTimeout to be an explicit nil

### UnsetSockReadTimeout
`func (o *HuggingFaceHuggingFaceRemoteResponse) UnsetSockReadTimeout()`

UnsetSockReadTimeout ensures that no value is present for SockReadTimeout, not even an explicit nil
### GetHeaders

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetHeaders() []map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetHeadersOk() (*[]map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetHeaders(v []map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetDownloadConcurrency

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetDownloadConcurrency() int64`

GetDownloadConcurrency returns the DownloadConcurrency field if non-nil, zero value otherwise.

### GetDownloadConcurrencyOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetDownloadConcurrencyOk() (*int64, bool)`

GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadConcurrency

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetDownloadConcurrency(v int64)`

SetDownloadConcurrency sets DownloadConcurrency field to given value.

### HasDownloadConcurrency

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasDownloadConcurrency() bool`

HasDownloadConcurrency returns a boolean if a field has been set.

### SetDownloadConcurrencyNil

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetDownloadConcurrencyNil(b bool)`

 SetDownloadConcurrencyNil sets the value for DownloadConcurrency to be an explicit nil

### UnsetDownloadConcurrency
`func (o *HuggingFaceHuggingFaceRemoteResponse) UnsetDownloadConcurrency()`

UnsetDownloadConcurrency ensures that no value is present for DownloadConcurrency, not even an explicit nil
### GetRateLimit

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetRateLimit() int64`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetRateLimitOk() (*int64, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetRateLimit(v int64)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### SetRateLimitNil

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetRateLimitNil(b bool)`

 SetRateLimitNil sets the value for RateLimit to be an explicit nil

### UnsetRateLimit
`func (o *HuggingFaceHuggingFaceRemoteResponse) UnsetRateLimit()`

UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil
### GetHfHubUrl

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetHfHubUrl() string`

GetHfHubUrl returns the HfHubUrl field if non-nil, zero value otherwise.

### GetHfHubUrlOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetHfHubUrlOk() (*string, bool)`

GetHfHubUrlOk returns a tuple with the HfHubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHfHubUrl

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetHfHubUrl(v string)`

SetHfHubUrl sets HfHubUrl field to given value.

### HasHfHubUrl

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasHfHubUrl() bool`

HasHfHubUrl returns a boolean if a field has been set.

### GetHfToken

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetHfToken() string`

GetHfToken returns the HfToken field if non-nil, zero value otherwise.

### GetHfTokenOk

`func (o *HuggingFaceHuggingFaceRemoteResponse) GetHfTokenOk() (*string, bool)`

GetHfTokenOk returns a tuple with the HfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHfToken

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetHfToken(v string)`

SetHfToken sets HfToken field to given value.

### HasHfToken

`func (o *HuggingFaceHuggingFaceRemoteResponse) HasHfToken() bool`

HasHfToken returns a boolean if a field has been set.

### SetHfTokenNil

`func (o *HuggingFaceHuggingFaceRemoteResponse) SetHfTokenNil(b bool)`

 SetHfTokenNil sets the value for HfToken to be an explicit nil

### UnsetHfToken
`func (o *HuggingFaceHuggingFaceRemoteResponse) UnsetHfToken()`

UnsetHfToken ensures that no value is present for HfToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


