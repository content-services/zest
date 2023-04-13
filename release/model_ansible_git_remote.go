/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zest

import (
	"encoding/json"
)

// checks if the AnsibleGitRemote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleGitRemote{}

// AnsibleGitRemote A serializer for Git Collection Remotes.
type AnsibleGitRemote struct {
	// A PEM encoded CA certificate used to validate the server certificate presented by the remote server.
	CaCert NullableString `json:"ca_cert,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// aiohttp.ClientTimeout.sock_read (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	SockReadTimeout NullableFloat64 `json:"sock_read_timeout,omitempty"`
	// Maximum number of retry attempts after a download failure. If not set then the default value (3) will be used.
	MaxRetries NullableInt64 `json:"max_retries,omitempty"`
	// Limits requests per second for each concurrent downloader
	RateLimit NullableInt64 `json:"rate_limit,omitempty"`
	// The proxy URL. Format: scheme://host:port
	ProxyUrl NullableString `json:"proxy_url,omitempty"`
	// Headers for aiohttp.Clientsession
	Headers []map[string]interface{} `json:"headers,omitempty"`
	// The username to authenticte to the proxy.
	ProxyUsername NullableString `json:"proxy_username,omitempty"`
	// The URL of an external content source.
	Url string `json:"url"`
	// aiohttp.ClientTimeout.total (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	TotalTimeout NullableFloat64 `json:"total_timeout,omitempty"`
	// The password to authenticate to the proxy. Extra leading and trailing whitespace characters are not trimmed.
	ProxyPassword NullableString `json:"proxy_password,omitempty"`
	// A PEM encoded client certificate used for authentication.
	ClientCert NullableString `json:"client_cert,omitempty"`
	// aiohttp.ClientTimeout.sock_connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	SockConnectTimeout NullableFloat64 `json:"sock_connect_timeout,omitempty"`
	// A PEM encoded private key used for authentication.
	ClientKey NullableString `json:"client_key,omitempty"`
	// aiohttp.ClientTimeout.connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	ConnectTimeout NullableFloat64 `json:"connect_timeout,omitempty"`
	// If True, TLS peer validation must be performed.
	TlsValidation *bool `json:"tls_validation,omitempty"`
	// A unique name for this remote.
	Name string `json:"name"`
	// Total number of simultaneous connections. If not set then the default value will be used.
	DownloadConcurrency NullableInt64 `json:"download_concurrency,omitempty"`
	// The username to be used for authentication when syncing.
	Username NullableString `json:"username,omitempty"`
	// The password to be used for authentication when syncing. Extra leading and trailing whitespace characters are not trimmed.
	Password NullableString `json:"password,omitempty"`
	// If True, only metadata about the content will be stored in Pulp. Clients will retrieve content from the remote URL.
	MetadataOnly *bool `json:"metadata_only,omitempty"`
	// A git ref. e.g.: branch, tag, or commit sha.
	GitRef *string `json:"git_ref,omitempty"`
}

// NewAnsibleGitRemote instantiates a new AnsibleGitRemote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleGitRemote(url string, name string) *AnsibleGitRemote {
	this := AnsibleGitRemote{}
	this.Url = url
	this.Name = name
	return &this
}

// NewAnsibleGitRemoteWithDefaults instantiates a new AnsibleGitRemote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleGitRemoteWithDefaults() *AnsibleGitRemote {
	this := AnsibleGitRemote{}
	return &this
}

// GetCaCert returns the CaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemote) GetCaCert() string {
	if o == nil || IsNil(o.CaCert.Get()) {
		var ret string
		return ret
	}
	return *o.CaCert.Get()
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemote) GetCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaCert.Get(), o.CaCert.IsSet()
}

// HasCaCert returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasCaCert() bool {
	if o != nil && o.CaCert.IsSet() {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given NullableString and assigns it to the CaCert field.
func (o *AnsibleGitRemote) SetCaCert(v string) {
	o.CaCert.Set(&v)
}
// SetCaCertNil sets the value for CaCert to be an explicit nil
func (o *AnsibleGitRemote) SetCaCertNil() {
	o.CaCert.Set(nil)
}

// UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
func (o *AnsibleGitRemote) UnsetCaCert() {
	o.CaCert.Unset()
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *AnsibleGitRemote) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemote) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *AnsibleGitRemote) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetSockReadTimeout returns the SockReadTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemote) GetSockReadTimeout() float64 {
	if o == nil || IsNil(o.SockReadTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.SockReadTimeout.Get()
}

// GetSockReadTimeoutOk returns a tuple with the SockReadTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemote) GetSockReadTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SockReadTimeout.Get(), o.SockReadTimeout.IsSet()
}

// HasSockReadTimeout returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasSockReadTimeout() bool {
	if o != nil && o.SockReadTimeout.IsSet() {
		return true
	}

	return false
}

// SetSockReadTimeout gets a reference to the given NullableFloat64 and assigns it to the SockReadTimeout field.
func (o *AnsibleGitRemote) SetSockReadTimeout(v float64) {
	o.SockReadTimeout.Set(&v)
}
// SetSockReadTimeoutNil sets the value for SockReadTimeout to be an explicit nil
func (o *AnsibleGitRemote) SetSockReadTimeoutNil() {
	o.SockReadTimeout.Set(nil)
}

// UnsetSockReadTimeout ensures that no value is present for SockReadTimeout, not even an explicit nil
func (o *AnsibleGitRemote) UnsetSockReadTimeout() {
	o.SockReadTimeout.Unset()
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemote) GetMaxRetries() int64 {
	if o == nil || IsNil(o.MaxRetries.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxRetries.Get()
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemote) GetMaxRetriesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRetries.Get(), o.MaxRetries.IsSet()
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasMaxRetries() bool {
	if o != nil && o.MaxRetries.IsSet() {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given NullableInt64 and assigns it to the MaxRetries field.
func (o *AnsibleGitRemote) SetMaxRetries(v int64) {
	o.MaxRetries.Set(&v)
}
// SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil
func (o *AnsibleGitRemote) SetMaxRetriesNil() {
	o.MaxRetries.Set(nil)
}

// UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
func (o *AnsibleGitRemote) UnsetMaxRetries() {
	o.MaxRetries.Unset()
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemote) GetRateLimit() int64 {
	if o == nil || IsNil(o.RateLimit.Get()) {
		var ret int64
		return ret
	}
	return *o.RateLimit.Get()
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemote) GetRateLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RateLimit.Get(), o.RateLimit.IsSet()
}

// HasRateLimit returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasRateLimit() bool {
	if o != nil && o.RateLimit.IsSet() {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given NullableInt64 and assigns it to the RateLimit field.
func (o *AnsibleGitRemote) SetRateLimit(v int64) {
	o.RateLimit.Set(&v)
}
// SetRateLimitNil sets the value for RateLimit to be an explicit nil
func (o *AnsibleGitRemote) SetRateLimitNil() {
	o.RateLimit.Set(nil)
}

// UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil
func (o *AnsibleGitRemote) UnsetRateLimit() {
	o.RateLimit.Unset()
}

// GetProxyUrl returns the ProxyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemote) GetProxyUrl() string {
	if o == nil || IsNil(o.ProxyUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyUrl.Get()
}

// GetProxyUrlOk returns a tuple with the ProxyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemote) GetProxyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyUrl.Get(), o.ProxyUrl.IsSet()
}

// HasProxyUrl returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasProxyUrl() bool {
	if o != nil && o.ProxyUrl.IsSet() {
		return true
	}

	return false
}

// SetProxyUrl gets a reference to the given NullableString and assigns it to the ProxyUrl field.
func (o *AnsibleGitRemote) SetProxyUrl(v string) {
	o.ProxyUrl.Set(&v)
}
// SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil
func (o *AnsibleGitRemote) SetProxyUrlNil() {
	o.ProxyUrl.Set(nil)
}

// UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
func (o *AnsibleGitRemote) UnsetProxyUrl() {
	o.ProxyUrl.Unset()
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *AnsibleGitRemote) GetHeaders() []map[string]interface{} {
	if o == nil || IsNil(o.Headers) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemote) GetHeadersOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []map[string]interface{} and assigns it to the Headers field.
func (o *AnsibleGitRemote) SetHeaders(v []map[string]interface{}) {
	o.Headers = v
}

// GetProxyUsername returns the ProxyUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemote) GetProxyUsername() string {
	if o == nil || IsNil(o.ProxyUsername.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyUsername.Get()
}

// GetProxyUsernameOk returns a tuple with the ProxyUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemote) GetProxyUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyUsername.Get(), o.ProxyUsername.IsSet()
}

// HasProxyUsername returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasProxyUsername() bool {
	if o != nil && o.ProxyUsername.IsSet() {
		return true
	}

	return false
}

// SetProxyUsername gets a reference to the given NullableString and assigns it to the ProxyUsername field.
func (o *AnsibleGitRemote) SetProxyUsername(v string) {
	o.ProxyUsername.Set(&v)
}
// SetProxyUsernameNil sets the value for ProxyUsername to be an explicit nil
func (o *AnsibleGitRemote) SetProxyUsernameNil() {
	o.ProxyUsername.Set(nil)
}

// UnsetProxyUsername ensures that no value is present for ProxyUsername, not even an explicit nil
func (o *AnsibleGitRemote) UnsetProxyUsername() {
	o.ProxyUsername.Unset()
}

// GetUrl returns the Url field value
func (o *AnsibleGitRemote) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemote) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AnsibleGitRemote) SetUrl(v string) {
	o.Url = v
}

// GetTotalTimeout returns the TotalTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemote) GetTotalTimeout() float64 {
	if o == nil || IsNil(o.TotalTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.TotalTimeout.Get()
}

// GetTotalTimeoutOk returns a tuple with the TotalTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemote) GetTotalTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalTimeout.Get(), o.TotalTimeout.IsSet()
}

// HasTotalTimeout returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasTotalTimeout() bool {
	if o != nil && o.TotalTimeout.IsSet() {
		return true
	}

	return false
}

// SetTotalTimeout gets a reference to the given NullableFloat64 and assigns it to the TotalTimeout field.
func (o *AnsibleGitRemote) SetTotalTimeout(v float64) {
	o.TotalTimeout.Set(&v)
}
// SetTotalTimeoutNil sets the value for TotalTimeout to be an explicit nil
func (o *AnsibleGitRemote) SetTotalTimeoutNil() {
	o.TotalTimeout.Set(nil)
}

// UnsetTotalTimeout ensures that no value is present for TotalTimeout, not even an explicit nil
func (o *AnsibleGitRemote) UnsetTotalTimeout() {
	o.TotalTimeout.Unset()
}

// GetProxyPassword returns the ProxyPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemote) GetProxyPassword() string {
	if o == nil || IsNil(o.ProxyPassword.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyPassword.Get()
}

// GetProxyPasswordOk returns a tuple with the ProxyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemote) GetProxyPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyPassword.Get(), o.ProxyPassword.IsSet()
}

// HasProxyPassword returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasProxyPassword() bool {
	if o != nil && o.ProxyPassword.IsSet() {
		return true
	}

	return false
}

// SetProxyPassword gets a reference to the given NullableString and assigns it to the ProxyPassword field.
func (o *AnsibleGitRemote) SetProxyPassword(v string) {
	o.ProxyPassword.Set(&v)
}
// SetProxyPasswordNil sets the value for ProxyPassword to be an explicit nil
func (o *AnsibleGitRemote) SetProxyPasswordNil() {
	o.ProxyPassword.Set(nil)
}

// UnsetProxyPassword ensures that no value is present for ProxyPassword, not even an explicit nil
func (o *AnsibleGitRemote) UnsetProxyPassword() {
	o.ProxyPassword.Unset()
}

// GetClientCert returns the ClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemote) GetClientCert() string {
	if o == nil || IsNil(o.ClientCert.Get()) {
		var ret string
		return ret
	}
	return *o.ClientCert.Get()
}

// GetClientCertOk returns a tuple with the ClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemote) GetClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientCert.Get(), o.ClientCert.IsSet()
}

// HasClientCert returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasClientCert() bool {
	if o != nil && o.ClientCert.IsSet() {
		return true
	}

	return false
}

// SetClientCert gets a reference to the given NullableString and assigns it to the ClientCert field.
func (o *AnsibleGitRemote) SetClientCert(v string) {
	o.ClientCert.Set(&v)
}
// SetClientCertNil sets the value for ClientCert to be an explicit nil
func (o *AnsibleGitRemote) SetClientCertNil() {
	o.ClientCert.Set(nil)
}

// UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
func (o *AnsibleGitRemote) UnsetClientCert() {
	o.ClientCert.Unset()
}

// GetSockConnectTimeout returns the SockConnectTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemote) GetSockConnectTimeout() float64 {
	if o == nil || IsNil(o.SockConnectTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.SockConnectTimeout.Get()
}

// GetSockConnectTimeoutOk returns a tuple with the SockConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemote) GetSockConnectTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SockConnectTimeout.Get(), o.SockConnectTimeout.IsSet()
}

// HasSockConnectTimeout returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasSockConnectTimeout() bool {
	if o != nil && o.SockConnectTimeout.IsSet() {
		return true
	}

	return false
}

// SetSockConnectTimeout gets a reference to the given NullableFloat64 and assigns it to the SockConnectTimeout field.
func (o *AnsibleGitRemote) SetSockConnectTimeout(v float64) {
	o.SockConnectTimeout.Set(&v)
}
// SetSockConnectTimeoutNil sets the value for SockConnectTimeout to be an explicit nil
func (o *AnsibleGitRemote) SetSockConnectTimeoutNil() {
	o.SockConnectTimeout.Set(nil)
}

// UnsetSockConnectTimeout ensures that no value is present for SockConnectTimeout, not even an explicit nil
func (o *AnsibleGitRemote) UnsetSockConnectTimeout() {
	o.SockConnectTimeout.Unset()
}

// GetClientKey returns the ClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemote) GetClientKey() string {
	if o == nil || IsNil(o.ClientKey.Get()) {
		var ret string
		return ret
	}
	return *o.ClientKey.Get()
}

// GetClientKeyOk returns a tuple with the ClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemote) GetClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientKey.Get(), o.ClientKey.IsSet()
}

// HasClientKey returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasClientKey() bool {
	if o != nil && o.ClientKey.IsSet() {
		return true
	}

	return false
}

// SetClientKey gets a reference to the given NullableString and assigns it to the ClientKey field.
func (o *AnsibleGitRemote) SetClientKey(v string) {
	o.ClientKey.Set(&v)
}
// SetClientKeyNil sets the value for ClientKey to be an explicit nil
func (o *AnsibleGitRemote) SetClientKeyNil() {
	o.ClientKey.Set(nil)
}

// UnsetClientKey ensures that no value is present for ClientKey, not even an explicit nil
func (o *AnsibleGitRemote) UnsetClientKey() {
	o.ClientKey.Unset()
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemote) GetConnectTimeout() float64 {
	if o == nil || IsNil(o.ConnectTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.ConnectTimeout.Get()
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemote) GetConnectTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectTimeout.Get(), o.ConnectTimeout.IsSet()
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasConnectTimeout() bool {
	if o != nil && o.ConnectTimeout.IsSet() {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given NullableFloat64 and assigns it to the ConnectTimeout field.
func (o *AnsibleGitRemote) SetConnectTimeout(v float64) {
	o.ConnectTimeout.Set(&v)
}
// SetConnectTimeoutNil sets the value for ConnectTimeout to be an explicit nil
func (o *AnsibleGitRemote) SetConnectTimeoutNil() {
	o.ConnectTimeout.Set(nil)
}

// UnsetConnectTimeout ensures that no value is present for ConnectTimeout, not even an explicit nil
func (o *AnsibleGitRemote) UnsetConnectTimeout() {
	o.ConnectTimeout.Unset()
}

// GetTlsValidation returns the TlsValidation field value if set, zero value otherwise.
func (o *AnsibleGitRemote) GetTlsValidation() bool {
	if o == nil || IsNil(o.TlsValidation) {
		var ret bool
		return ret
	}
	return *o.TlsValidation
}

// GetTlsValidationOk returns a tuple with the TlsValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemote) GetTlsValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsValidation) {
		return nil, false
	}
	return o.TlsValidation, true
}

// HasTlsValidation returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasTlsValidation() bool {
	if o != nil && !IsNil(o.TlsValidation) {
		return true
	}

	return false
}

// SetTlsValidation gets a reference to the given bool and assigns it to the TlsValidation field.
func (o *AnsibleGitRemote) SetTlsValidation(v bool) {
	o.TlsValidation = &v
}

// GetName returns the Name field value
func (o *AnsibleGitRemote) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemote) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AnsibleGitRemote) SetName(v string) {
	o.Name = v
}

// GetDownloadConcurrency returns the DownloadConcurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemote) GetDownloadConcurrency() int64 {
	if o == nil || IsNil(o.DownloadConcurrency.Get()) {
		var ret int64
		return ret
	}
	return *o.DownloadConcurrency.Get()
}

// GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemote) GetDownloadConcurrencyOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadConcurrency.Get(), o.DownloadConcurrency.IsSet()
}

// HasDownloadConcurrency returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasDownloadConcurrency() bool {
	if o != nil && o.DownloadConcurrency.IsSet() {
		return true
	}

	return false
}

// SetDownloadConcurrency gets a reference to the given NullableInt64 and assigns it to the DownloadConcurrency field.
func (o *AnsibleGitRemote) SetDownloadConcurrency(v int64) {
	o.DownloadConcurrency.Set(&v)
}
// SetDownloadConcurrencyNil sets the value for DownloadConcurrency to be an explicit nil
func (o *AnsibleGitRemote) SetDownloadConcurrencyNil() {
	o.DownloadConcurrency.Set(nil)
}

// UnsetDownloadConcurrency ensures that no value is present for DownloadConcurrency, not even an explicit nil
func (o *AnsibleGitRemote) UnsetDownloadConcurrency() {
	o.DownloadConcurrency.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemote) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemote) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *AnsibleGitRemote) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *AnsibleGitRemote) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *AnsibleGitRemote) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleGitRemote) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleGitRemote) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *AnsibleGitRemote) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *AnsibleGitRemote) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *AnsibleGitRemote) UnsetPassword() {
	o.Password.Unset()
}

// GetMetadataOnly returns the MetadataOnly field value if set, zero value otherwise.
func (o *AnsibleGitRemote) GetMetadataOnly() bool {
	if o == nil || IsNil(o.MetadataOnly) {
		var ret bool
		return ret
	}
	return *o.MetadataOnly
}

// GetMetadataOnlyOk returns a tuple with the MetadataOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemote) GetMetadataOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.MetadataOnly) {
		return nil, false
	}
	return o.MetadataOnly, true
}

// HasMetadataOnly returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasMetadataOnly() bool {
	if o != nil && !IsNil(o.MetadataOnly) {
		return true
	}

	return false
}

// SetMetadataOnly gets a reference to the given bool and assigns it to the MetadataOnly field.
func (o *AnsibleGitRemote) SetMetadataOnly(v bool) {
	o.MetadataOnly = &v
}

// GetGitRef returns the GitRef field value if set, zero value otherwise.
func (o *AnsibleGitRemote) GetGitRef() string {
	if o == nil || IsNil(o.GitRef) {
		var ret string
		return ret
	}
	return *o.GitRef
}

// GetGitRefOk returns a tuple with the GitRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleGitRemote) GetGitRefOk() (*string, bool) {
	if o == nil || IsNil(o.GitRef) {
		return nil, false
	}
	return o.GitRef, true
}

// HasGitRef returns a boolean if a field has been set.
func (o *AnsibleGitRemote) HasGitRef() bool {
	if o != nil && !IsNil(o.GitRef) {
		return true
	}

	return false
}

// SetGitRef gets a reference to the given string and assigns it to the GitRef field.
func (o *AnsibleGitRemote) SetGitRef(v string) {
	o.GitRef = &v
}

func (o AnsibleGitRemote) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleGitRemote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CaCert.IsSet() {
		toSerialize["ca_cert"] = o.CaCert.Get()
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	if o.SockReadTimeout.IsSet() {
		toSerialize["sock_read_timeout"] = o.SockReadTimeout.Get()
	}
	if o.MaxRetries.IsSet() {
		toSerialize["max_retries"] = o.MaxRetries.Get()
	}
	if o.RateLimit.IsSet() {
		toSerialize["rate_limit"] = o.RateLimit.Get()
	}
	if o.ProxyUrl.IsSet() {
		toSerialize["proxy_url"] = o.ProxyUrl.Get()
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if o.ProxyUsername.IsSet() {
		toSerialize["proxy_username"] = o.ProxyUsername.Get()
	}
	toSerialize["url"] = o.Url
	if o.TotalTimeout.IsSet() {
		toSerialize["total_timeout"] = o.TotalTimeout.Get()
	}
	if o.ProxyPassword.IsSet() {
		toSerialize["proxy_password"] = o.ProxyPassword.Get()
	}
	if o.ClientCert.IsSet() {
		toSerialize["client_cert"] = o.ClientCert.Get()
	}
	if o.SockConnectTimeout.IsSet() {
		toSerialize["sock_connect_timeout"] = o.SockConnectTimeout.Get()
	}
	if o.ClientKey.IsSet() {
		toSerialize["client_key"] = o.ClientKey.Get()
	}
	if o.ConnectTimeout.IsSet() {
		toSerialize["connect_timeout"] = o.ConnectTimeout.Get()
	}
	if !IsNil(o.TlsValidation) {
		toSerialize["tls_validation"] = o.TlsValidation
	}
	toSerialize["name"] = o.Name
	if o.DownloadConcurrency.IsSet() {
		toSerialize["download_concurrency"] = o.DownloadConcurrency.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if !IsNil(o.MetadataOnly) {
		toSerialize["metadata_only"] = o.MetadataOnly
	}
	if !IsNil(o.GitRef) {
		toSerialize["git_ref"] = o.GitRef
	}
	return toSerialize, nil
}

type NullableAnsibleGitRemote struct {
	value *AnsibleGitRemote
	isSet bool
}

func (v NullableAnsibleGitRemote) Get() *AnsibleGitRemote {
	return v.value
}

func (v *NullableAnsibleGitRemote) Set(val *AnsibleGitRemote) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleGitRemote) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleGitRemote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleGitRemote(val *AnsibleGitRemote) *NullableAnsibleGitRemote {
	return &NullableAnsibleGitRemote{value: val, isSet: true}
}

func (v NullableAnsibleGitRemote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleGitRemote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


