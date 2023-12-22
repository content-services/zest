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
	"bytes"
	"fmt"
)

// checks if the RpmRpmRemote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmRpmRemote{}

// RpmRpmRemote A Serializer for RpmRemote.
type RpmRpmRemote struct {
	// A unique name for this remote.
	Name string `json:"name"`
	// The URL of an external content source.
	Url string `json:"url"`
	// A PEM encoded CA certificate used to validate the server certificate presented by the remote server.
	CaCert NullableString `json:"ca_cert,omitempty"`
	// A PEM encoded client certificate used for authentication.
	ClientCert NullableString `json:"client_cert,omitempty"`
	// A PEM encoded private key used for authentication.
	ClientKey NullableString `json:"client_key,omitempty"`
	// If True, TLS peer validation must be performed.
	TlsValidation *bool `json:"tls_validation,omitempty"`
	// The proxy URL. Format: scheme://host:port
	ProxyUrl NullableString `json:"proxy_url,omitempty"`
	// The username to authenticte to the proxy.
	ProxyUsername NullableString `json:"proxy_username,omitempty"`
	// The password to authenticate to the proxy. Extra leading and trailing whitespace characters are not trimmed.
	ProxyPassword NullableString `json:"proxy_password,omitempty"`
	// The username to be used for authentication when syncing.
	Username NullableString `json:"username,omitempty"`
	// The password to be used for authentication when syncing. Extra leading and trailing whitespace characters are not trimmed.
	Password NullableString `json:"password,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// Total number of simultaneous connections. If not set then the default value will be used.
	DownloadConcurrency NullableInt64 `json:"download_concurrency,omitempty"`
	// Maximum number of retry attempts after a download failure. If not set then the default value (3) will be used.
	MaxRetries NullableInt64 `json:"max_retries,omitempty"`
	Policy *Policy762Enum `json:"policy,omitempty"`
	// aiohttp.ClientTimeout.total (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	TotalTimeout NullableFloat64 `json:"total_timeout,omitempty"`
	// aiohttp.ClientTimeout.connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	ConnectTimeout NullableFloat64 `json:"connect_timeout,omitempty"`
	// aiohttp.ClientTimeout.sock_connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	SockConnectTimeout NullableFloat64 `json:"sock_connect_timeout,omitempty"`
	// aiohttp.ClientTimeout.sock_read (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	SockReadTimeout NullableFloat64 `json:"sock_read_timeout,omitempty"`
	// Headers for aiohttp.Clientsession
	Headers []map[string]interface{} `json:"headers,omitempty"`
	// Limits requests per second for each concurrent downloader
	RateLimit NullableInt64 `json:"rate_limit,omitempty"`
	// Authentication token for SLES repositories.
	SlesAuthToken NullableString `json:"sles_auth_token,omitempty"`
}

type _RpmRpmRemote RpmRpmRemote

// NewRpmRpmRemote instantiates a new RpmRpmRemote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmRpmRemote(name string, url string) *RpmRpmRemote {
	this := RpmRpmRemote{}
	this.Name = name
	this.Url = url
	var policy Policy762Enum = POLICY762ENUM_IMMEDIATE
	this.Policy = &policy
	return &this
}

// NewRpmRpmRemoteWithDefaults instantiates a new RpmRpmRemote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmRpmRemoteWithDefaults() *RpmRpmRemote {
	this := RpmRpmRemote{}
	var policy Policy762Enum = POLICY762ENUM_IMMEDIATE
	this.Policy = &policy
	return &this
}

// GetName returns the Name field value
func (o *RpmRpmRemote) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RpmRpmRemote) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RpmRpmRemote) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *RpmRpmRemote) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RpmRpmRemote) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *RpmRpmRemote) SetUrl(v string) {
	o.Url = v
}

// GetCaCert returns the CaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRemote) GetCaCert() string {
	if o == nil || IsNil(o.CaCert.Get()) {
		var ret string
		return ret
	}
	return *o.CaCert.Get()
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRemote) GetCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaCert.Get(), o.CaCert.IsSet()
}

// HasCaCert returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasCaCert() bool {
	if o != nil && o.CaCert.IsSet() {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given NullableString and assigns it to the CaCert field.
func (o *RpmRpmRemote) SetCaCert(v string) {
	o.CaCert.Set(&v)
}
// SetCaCertNil sets the value for CaCert to be an explicit nil
func (o *RpmRpmRemote) SetCaCertNil() {
	o.CaCert.Set(nil)
}

// UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
func (o *RpmRpmRemote) UnsetCaCert() {
	o.CaCert.Unset()
}

// GetClientCert returns the ClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRemote) GetClientCert() string {
	if o == nil || IsNil(o.ClientCert.Get()) {
		var ret string
		return ret
	}
	return *o.ClientCert.Get()
}

// GetClientCertOk returns a tuple with the ClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRemote) GetClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientCert.Get(), o.ClientCert.IsSet()
}

// HasClientCert returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasClientCert() bool {
	if o != nil && o.ClientCert.IsSet() {
		return true
	}

	return false
}

// SetClientCert gets a reference to the given NullableString and assigns it to the ClientCert field.
func (o *RpmRpmRemote) SetClientCert(v string) {
	o.ClientCert.Set(&v)
}
// SetClientCertNil sets the value for ClientCert to be an explicit nil
func (o *RpmRpmRemote) SetClientCertNil() {
	o.ClientCert.Set(nil)
}

// UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
func (o *RpmRpmRemote) UnsetClientCert() {
	o.ClientCert.Unset()
}

// GetClientKey returns the ClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRemote) GetClientKey() string {
	if o == nil || IsNil(o.ClientKey.Get()) {
		var ret string
		return ret
	}
	return *o.ClientKey.Get()
}

// GetClientKeyOk returns a tuple with the ClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRemote) GetClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientKey.Get(), o.ClientKey.IsSet()
}

// HasClientKey returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasClientKey() bool {
	if o != nil && o.ClientKey.IsSet() {
		return true
	}

	return false
}

// SetClientKey gets a reference to the given NullableString and assigns it to the ClientKey field.
func (o *RpmRpmRemote) SetClientKey(v string) {
	o.ClientKey.Set(&v)
}
// SetClientKeyNil sets the value for ClientKey to be an explicit nil
func (o *RpmRpmRemote) SetClientKeyNil() {
	o.ClientKey.Set(nil)
}

// UnsetClientKey ensures that no value is present for ClientKey, not even an explicit nil
func (o *RpmRpmRemote) UnsetClientKey() {
	o.ClientKey.Unset()
}

// GetTlsValidation returns the TlsValidation field value if set, zero value otherwise.
func (o *RpmRpmRemote) GetTlsValidation() bool {
	if o == nil || IsNil(o.TlsValidation) {
		var ret bool
		return ret
	}
	return *o.TlsValidation
}

// GetTlsValidationOk returns a tuple with the TlsValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmRemote) GetTlsValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsValidation) {
		return nil, false
	}
	return o.TlsValidation, true
}

// HasTlsValidation returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasTlsValidation() bool {
	if o != nil && !IsNil(o.TlsValidation) {
		return true
	}

	return false
}

// SetTlsValidation gets a reference to the given bool and assigns it to the TlsValidation field.
func (o *RpmRpmRemote) SetTlsValidation(v bool) {
	o.TlsValidation = &v
}

// GetProxyUrl returns the ProxyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRemote) GetProxyUrl() string {
	if o == nil || IsNil(o.ProxyUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyUrl.Get()
}

// GetProxyUrlOk returns a tuple with the ProxyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRemote) GetProxyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyUrl.Get(), o.ProxyUrl.IsSet()
}

// HasProxyUrl returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasProxyUrl() bool {
	if o != nil && o.ProxyUrl.IsSet() {
		return true
	}

	return false
}

// SetProxyUrl gets a reference to the given NullableString and assigns it to the ProxyUrl field.
func (o *RpmRpmRemote) SetProxyUrl(v string) {
	o.ProxyUrl.Set(&v)
}
// SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil
func (o *RpmRpmRemote) SetProxyUrlNil() {
	o.ProxyUrl.Set(nil)
}

// UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
func (o *RpmRpmRemote) UnsetProxyUrl() {
	o.ProxyUrl.Unset()
}

// GetProxyUsername returns the ProxyUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRemote) GetProxyUsername() string {
	if o == nil || IsNil(o.ProxyUsername.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyUsername.Get()
}

// GetProxyUsernameOk returns a tuple with the ProxyUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRemote) GetProxyUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyUsername.Get(), o.ProxyUsername.IsSet()
}

// HasProxyUsername returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasProxyUsername() bool {
	if o != nil && o.ProxyUsername.IsSet() {
		return true
	}

	return false
}

// SetProxyUsername gets a reference to the given NullableString and assigns it to the ProxyUsername field.
func (o *RpmRpmRemote) SetProxyUsername(v string) {
	o.ProxyUsername.Set(&v)
}
// SetProxyUsernameNil sets the value for ProxyUsername to be an explicit nil
func (o *RpmRpmRemote) SetProxyUsernameNil() {
	o.ProxyUsername.Set(nil)
}

// UnsetProxyUsername ensures that no value is present for ProxyUsername, not even an explicit nil
func (o *RpmRpmRemote) UnsetProxyUsername() {
	o.ProxyUsername.Unset()
}

// GetProxyPassword returns the ProxyPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRemote) GetProxyPassword() string {
	if o == nil || IsNil(o.ProxyPassword.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyPassword.Get()
}

// GetProxyPasswordOk returns a tuple with the ProxyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRemote) GetProxyPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyPassword.Get(), o.ProxyPassword.IsSet()
}

// HasProxyPassword returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasProxyPassword() bool {
	if o != nil && o.ProxyPassword.IsSet() {
		return true
	}

	return false
}

// SetProxyPassword gets a reference to the given NullableString and assigns it to the ProxyPassword field.
func (o *RpmRpmRemote) SetProxyPassword(v string) {
	o.ProxyPassword.Set(&v)
}
// SetProxyPasswordNil sets the value for ProxyPassword to be an explicit nil
func (o *RpmRpmRemote) SetProxyPasswordNil() {
	o.ProxyPassword.Set(nil)
}

// UnsetProxyPassword ensures that no value is present for ProxyPassword, not even an explicit nil
func (o *RpmRpmRemote) UnsetProxyPassword() {
	o.ProxyPassword.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRemote) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRemote) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *RpmRpmRemote) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *RpmRpmRemote) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *RpmRpmRemote) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRemote) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRemote) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *RpmRpmRemote) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *RpmRpmRemote) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *RpmRpmRemote) UnsetPassword() {
	o.Password.Unset()
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *RpmRpmRemote) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmRemote) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *RpmRpmRemote) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetDownloadConcurrency returns the DownloadConcurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRemote) GetDownloadConcurrency() int64 {
	if o == nil || IsNil(o.DownloadConcurrency.Get()) {
		var ret int64
		return ret
	}
	return *o.DownloadConcurrency.Get()
}

// GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRemote) GetDownloadConcurrencyOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadConcurrency.Get(), o.DownloadConcurrency.IsSet()
}

// HasDownloadConcurrency returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasDownloadConcurrency() bool {
	if o != nil && o.DownloadConcurrency.IsSet() {
		return true
	}

	return false
}

// SetDownloadConcurrency gets a reference to the given NullableInt64 and assigns it to the DownloadConcurrency field.
func (o *RpmRpmRemote) SetDownloadConcurrency(v int64) {
	o.DownloadConcurrency.Set(&v)
}
// SetDownloadConcurrencyNil sets the value for DownloadConcurrency to be an explicit nil
func (o *RpmRpmRemote) SetDownloadConcurrencyNil() {
	o.DownloadConcurrency.Set(nil)
}

// UnsetDownloadConcurrency ensures that no value is present for DownloadConcurrency, not even an explicit nil
func (o *RpmRpmRemote) UnsetDownloadConcurrency() {
	o.DownloadConcurrency.Unset()
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRemote) GetMaxRetries() int64 {
	if o == nil || IsNil(o.MaxRetries.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxRetries.Get()
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRemote) GetMaxRetriesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRetries.Get(), o.MaxRetries.IsSet()
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasMaxRetries() bool {
	if o != nil && o.MaxRetries.IsSet() {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given NullableInt64 and assigns it to the MaxRetries field.
func (o *RpmRpmRemote) SetMaxRetries(v int64) {
	o.MaxRetries.Set(&v)
}
// SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil
func (o *RpmRpmRemote) SetMaxRetriesNil() {
	o.MaxRetries.Set(nil)
}

// UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
func (o *RpmRpmRemote) UnsetMaxRetries() {
	o.MaxRetries.Unset()
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *RpmRpmRemote) GetPolicy() Policy762Enum {
	if o == nil || IsNil(o.Policy) {
		var ret Policy762Enum
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmRemote) GetPolicyOk() (*Policy762Enum, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given Policy762Enum and assigns it to the Policy field.
func (o *RpmRpmRemote) SetPolicy(v Policy762Enum) {
	o.Policy = &v
}

// GetTotalTimeout returns the TotalTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRemote) GetTotalTimeout() float64 {
	if o == nil || IsNil(o.TotalTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.TotalTimeout.Get()
}

// GetTotalTimeoutOk returns a tuple with the TotalTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRemote) GetTotalTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalTimeout.Get(), o.TotalTimeout.IsSet()
}

// HasTotalTimeout returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasTotalTimeout() bool {
	if o != nil && o.TotalTimeout.IsSet() {
		return true
	}

	return false
}

// SetTotalTimeout gets a reference to the given NullableFloat64 and assigns it to the TotalTimeout field.
func (o *RpmRpmRemote) SetTotalTimeout(v float64) {
	o.TotalTimeout.Set(&v)
}
// SetTotalTimeoutNil sets the value for TotalTimeout to be an explicit nil
func (o *RpmRpmRemote) SetTotalTimeoutNil() {
	o.TotalTimeout.Set(nil)
}

// UnsetTotalTimeout ensures that no value is present for TotalTimeout, not even an explicit nil
func (o *RpmRpmRemote) UnsetTotalTimeout() {
	o.TotalTimeout.Unset()
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRemote) GetConnectTimeout() float64 {
	if o == nil || IsNil(o.ConnectTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.ConnectTimeout.Get()
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRemote) GetConnectTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectTimeout.Get(), o.ConnectTimeout.IsSet()
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasConnectTimeout() bool {
	if o != nil && o.ConnectTimeout.IsSet() {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given NullableFloat64 and assigns it to the ConnectTimeout field.
func (o *RpmRpmRemote) SetConnectTimeout(v float64) {
	o.ConnectTimeout.Set(&v)
}
// SetConnectTimeoutNil sets the value for ConnectTimeout to be an explicit nil
func (o *RpmRpmRemote) SetConnectTimeoutNil() {
	o.ConnectTimeout.Set(nil)
}

// UnsetConnectTimeout ensures that no value is present for ConnectTimeout, not even an explicit nil
func (o *RpmRpmRemote) UnsetConnectTimeout() {
	o.ConnectTimeout.Unset()
}

// GetSockConnectTimeout returns the SockConnectTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRemote) GetSockConnectTimeout() float64 {
	if o == nil || IsNil(o.SockConnectTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.SockConnectTimeout.Get()
}

// GetSockConnectTimeoutOk returns a tuple with the SockConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRemote) GetSockConnectTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SockConnectTimeout.Get(), o.SockConnectTimeout.IsSet()
}

// HasSockConnectTimeout returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasSockConnectTimeout() bool {
	if o != nil && o.SockConnectTimeout.IsSet() {
		return true
	}

	return false
}

// SetSockConnectTimeout gets a reference to the given NullableFloat64 and assigns it to the SockConnectTimeout field.
func (o *RpmRpmRemote) SetSockConnectTimeout(v float64) {
	o.SockConnectTimeout.Set(&v)
}
// SetSockConnectTimeoutNil sets the value for SockConnectTimeout to be an explicit nil
func (o *RpmRpmRemote) SetSockConnectTimeoutNil() {
	o.SockConnectTimeout.Set(nil)
}

// UnsetSockConnectTimeout ensures that no value is present for SockConnectTimeout, not even an explicit nil
func (o *RpmRpmRemote) UnsetSockConnectTimeout() {
	o.SockConnectTimeout.Unset()
}

// GetSockReadTimeout returns the SockReadTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRemote) GetSockReadTimeout() float64 {
	if o == nil || IsNil(o.SockReadTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.SockReadTimeout.Get()
}

// GetSockReadTimeoutOk returns a tuple with the SockReadTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRemote) GetSockReadTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SockReadTimeout.Get(), o.SockReadTimeout.IsSet()
}

// HasSockReadTimeout returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasSockReadTimeout() bool {
	if o != nil && o.SockReadTimeout.IsSet() {
		return true
	}

	return false
}

// SetSockReadTimeout gets a reference to the given NullableFloat64 and assigns it to the SockReadTimeout field.
func (o *RpmRpmRemote) SetSockReadTimeout(v float64) {
	o.SockReadTimeout.Set(&v)
}
// SetSockReadTimeoutNil sets the value for SockReadTimeout to be an explicit nil
func (o *RpmRpmRemote) SetSockReadTimeoutNil() {
	o.SockReadTimeout.Set(nil)
}

// UnsetSockReadTimeout ensures that no value is present for SockReadTimeout, not even an explicit nil
func (o *RpmRpmRemote) UnsetSockReadTimeout() {
	o.SockReadTimeout.Unset()
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *RpmRpmRemote) GetHeaders() []map[string]interface{} {
	if o == nil || IsNil(o.Headers) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmRemote) GetHeadersOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []map[string]interface{} and assigns it to the Headers field.
func (o *RpmRpmRemote) SetHeaders(v []map[string]interface{}) {
	o.Headers = v
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRemote) GetRateLimit() int64 {
	if o == nil || IsNil(o.RateLimit.Get()) {
		var ret int64
		return ret
	}
	return *o.RateLimit.Get()
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRemote) GetRateLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RateLimit.Get(), o.RateLimit.IsSet()
}

// HasRateLimit returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasRateLimit() bool {
	if o != nil && o.RateLimit.IsSet() {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given NullableInt64 and assigns it to the RateLimit field.
func (o *RpmRpmRemote) SetRateLimit(v int64) {
	o.RateLimit.Set(&v)
}
// SetRateLimitNil sets the value for RateLimit to be an explicit nil
func (o *RpmRpmRemote) SetRateLimitNil() {
	o.RateLimit.Set(nil)
}

// UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil
func (o *RpmRpmRemote) UnsetRateLimit() {
	o.RateLimit.Unset()
}

// GetSlesAuthToken returns the SlesAuthToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmRemote) GetSlesAuthToken() string {
	if o == nil || IsNil(o.SlesAuthToken.Get()) {
		var ret string
		return ret
	}
	return *o.SlesAuthToken.Get()
}

// GetSlesAuthTokenOk returns a tuple with the SlesAuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmRemote) GetSlesAuthTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlesAuthToken.Get(), o.SlesAuthToken.IsSet()
}

// HasSlesAuthToken returns a boolean if a field has been set.
func (o *RpmRpmRemote) HasSlesAuthToken() bool {
	if o != nil && o.SlesAuthToken.IsSet() {
		return true
	}

	return false
}

// SetSlesAuthToken gets a reference to the given NullableString and assigns it to the SlesAuthToken field.
func (o *RpmRpmRemote) SetSlesAuthToken(v string) {
	o.SlesAuthToken.Set(&v)
}
// SetSlesAuthTokenNil sets the value for SlesAuthToken to be an explicit nil
func (o *RpmRpmRemote) SetSlesAuthTokenNil() {
	o.SlesAuthToken.Set(nil)
}

// UnsetSlesAuthToken ensures that no value is present for SlesAuthToken, not even an explicit nil
func (o *RpmRpmRemote) UnsetSlesAuthToken() {
	o.SlesAuthToken.Unset()
}

func (o RpmRpmRemote) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmRpmRemote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	if o.CaCert.IsSet() {
		toSerialize["ca_cert"] = o.CaCert.Get()
	}
	if o.ClientCert.IsSet() {
		toSerialize["client_cert"] = o.ClientCert.Get()
	}
	if o.ClientKey.IsSet() {
		toSerialize["client_key"] = o.ClientKey.Get()
	}
	if !IsNil(o.TlsValidation) {
		toSerialize["tls_validation"] = o.TlsValidation
	}
	if o.ProxyUrl.IsSet() {
		toSerialize["proxy_url"] = o.ProxyUrl.Get()
	}
	if o.ProxyUsername.IsSet() {
		toSerialize["proxy_username"] = o.ProxyUsername.Get()
	}
	if o.ProxyPassword.IsSet() {
		toSerialize["proxy_password"] = o.ProxyPassword.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	if o.DownloadConcurrency.IsSet() {
		toSerialize["download_concurrency"] = o.DownloadConcurrency.Get()
	}
	if o.MaxRetries.IsSet() {
		toSerialize["max_retries"] = o.MaxRetries.Get()
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if o.TotalTimeout.IsSet() {
		toSerialize["total_timeout"] = o.TotalTimeout.Get()
	}
	if o.ConnectTimeout.IsSet() {
		toSerialize["connect_timeout"] = o.ConnectTimeout.Get()
	}
	if o.SockConnectTimeout.IsSet() {
		toSerialize["sock_connect_timeout"] = o.SockConnectTimeout.Get()
	}
	if o.SockReadTimeout.IsSet() {
		toSerialize["sock_read_timeout"] = o.SockReadTimeout.Get()
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if o.RateLimit.IsSet() {
		toSerialize["rate_limit"] = o.RateLimit.Get()
	}
	if o.SlesAuthToken.IsSet() {
		toSerialize["sles_auth_token"] = o.SlesAuthToken.Get()
	}
	return toSerialize, nil
}

func (o *RpmRpmRemote) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRpmRpmRemote := _RpmRpmRemote{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRpmRpmRemote)

	if err != nil {
		return err
	}

	*o = RpmRpmRemote(varRpmRpmRemote)

	return err
}

type NullableRpmRpmRemote struct {
	value *RpmRpmRemote
	isSet bool
}

func (v NullableRpmRpmRemote) Get() *RpmRpmRemote {
	return v.value
}

func (v *NullableRpmRpmRemote) Set(val *RpmRpmRemote) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmRpmRemote) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmRpmRemote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmRpmRemote(val *RpmRpmRemote) *NullableRpmRpmRemote {
	return &NullableRpmRpmRemote{value: val, isSet: true}
}

func (v NullableRpmRpmRemote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmRpmRemote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


