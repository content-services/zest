/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpGoBinding

import (
	"encoding/json"
)

// checks if the PatchedmavenMavenRemote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedmavenMavenRemote{}

// PatchedmavenMavenRemote A Serializer for MavenRemote.  Add any new fields if defined on MavenRemote. Similar to the example above, in MavenArtifactSerializer. Additional validators can be added to the parent validators list  For example::  class Meta:     validators = platform.RemoteSerializer.Meta.validators + [myValidator1, myValidator2]
type PatchedmavenMavenRemote struct {
	// A unique name for this remote.
	Name *string `json:"name,omitempty"`
	// The URL of an external content source.
	Url *string `json:"url,omitempty"`
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
	DownloadConcurrency NullableInt32 `json:"download_concurrency,omitempty"`
	// Maximum number of retry attempts after a download failure. If not set then the default value (3) will be used.
	MaxRetries NullableInt32 `json:"max_retries,omitempty"`
	Policy *PolicyDb6Enum `json:"policy,omitempty"`
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
	RateLimit NullableInt32 `json:"rate_limit,omitempty"`
}

// NewPatchedmavenMavenRemote instantiates a new PatchedmavenMavenRemote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedmavenMavenRemote() *PatchedmavenMavenRemote {
	this := PatchedmavenMavenRemote{}
	return &this
}

// NewPatchedmavenMavenRemoteWithDefaults instantiates a new PatchedmavenMavenRemote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedmavenMavenRemoteWithDefaults() *PatchedmavenMavenRemote {
	this := PatchedmavenMavenRemote{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedmavenMavenRemote) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedmavenMavenRemote) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedmavenMavenRemote) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedmavenMavenRemote) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedmavenMavenRemote) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedmavenMavenRemote) SetUrl(v string) {
	o.Url = &v
}

// GetCaCert returns the CaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedmavenMavenRemote) GetCaCert() string {
	if o == nil || IsNil(o.CaCert.Get()) {
		var ret string
		return ret
	}
	return *o.CaCert.Get()
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedmavenMavenRemote) GetCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaCert.Get(), o.CaCert.IsSet()
}

// HasCaCert returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasCaCert() bool {
	if o != nil && o.CaCert.IsSet() {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given NullableString and assigns it to the CaCert field.
func (o *PatchedmavenMavenRemote) SetCaCert(v string) {
	o.CaCert.Set(&v)
}
// SetCaCertNil sets the value for CaCert to be an explicit nil
func (o *PatchedmavenMavenRemote) SetCaCertNil() {
	o.CaCert.Set(nil)
}

// UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
func (o *PatchedmavenMavenRemote) UnsetCaCert() {
	o.CaCert.Unset()
}

// GetClientCert returns the ClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedmavenMavenRemote) GetClientCert() string {
	if o == nil || IsNil(o.ClientCert.Get()) {
		var ret string
		return ret
	}
	return *o.ClientCert.Get()
}

// GetClientCertOk returns a tuple with the ClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedmavenMavenRemote) GetClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientCert.Get(), o.ClientCert.IsSet()
}

// HasClientCert returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasClientCert() bool {
	if o != nil && o.ClientCert.IsSet() {
		return true
	}

	return false
}

// SetClientCert gets a reference to the given NullableString and assigns it to the ClientCert field.
func (o *PatchedmavenMavenRemote) SetClientCert(v string) {
	o.ClientCert.Set(&v)
}
// SetClientCertNil sets the value for ClientCert to be an explicit nil
func (o *PatchedmavenMavenRemote) SetClientCertNil() {
	o.ClientCert.Set(nil)
}

// UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
func (o *PatchedmavenMavenRemote) UnsetClientCert() {
	o.ClientCert.Unset()
}

// GetClientKey returns the ClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedmavenMavenRemote) GetClientKey() string {
	if o == nil || IsNil(o.ClientKey.Get()) {
		var ret string
		return ret
	}
	return *o.ClientKey.Get()
}

// GetClientKeyOk returns a tuple with the ClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedmavenMavenRemote) GetClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientKey.Get(), o.ClientKey.IsSet()
}

// HasClientKey returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasClientKey() bool {
	if o != nil && o.ClientKey.IsSet() {
		return true
	}

	return false
}

// SetClientKey gets a reference to the given NullableString and assigns it to the ClientKey field.
func (o *PatchedmavenMavenRemote) SetClientKey(v string) {
	o.ClientKey.Set(&v)
}
// SetClientKeyNil sets the value for ClientKey to be an explicit nil
func (o *PatchedmavenMavenRemote) SetClientKeyNil() {
	o.ClientKey.Set(nil)
}

// UnsetClientKey ensures that no value is present for ClientKey, not even an explicit nil
func (o *PatchedmavenMavenRemote) UnsetClientKey() {
	o.ClientKey.Unset()
}

// GetTlsValidation returns the TlsValidation field value if set, zero value otherwise.
func (o *PatchedmavenMavenRemote) GetTlsValidation() bool {
	if o == nil || IsNil(o.TlsValidation) {
		var ret bool
		return ret
	}
	return *o.TlsValidation
}

// GetTlsValidationOk returns a tuple with the TlsValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedmavenMavenRemote) GetTlsValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsValidation) {
		return nil, false
	}
	return o.TlsValidation, true
}

// HasTlsValidation returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasTlsValidation() bool {
	if o != nil && !IsNil(o.TlsValidation) {
		return true
	}

	return false
}

// SetTlsValidation gets a reference to the given bool and assigns it to the TlsValidation field.
func (o *PatchedmavenMavenRemote) SetTlsValidation(v bool) {
	o.TlsValidation = &v
}

// GetProxyUrl returns the ProxyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedmavenMavenRemote) GetProxyUrl() string {
	if o == nil || IsNil(o.ProxyUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyUrl.Get()
}

// GetProxyUrlOk returns a tuple with the ProxyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedmavenMavenRemote) GetProxyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyUrl.Get(), o.ProxyUrl.IsSet()
}

// HasProxyUrl returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasProxyUrl() bool {
	if o != nil && o.ProxyUrl.IsSet() {
		return true
	}

	return false
}

// SetProxyUrl gets a reference to the given NullableString and assigns it to the ProxyUrl field.
func (o *PatchedmavenMavenRemote) SetProxyUrl(v string) {
	o.ProxyUrl.Set(&v)
}
// SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil
func (o *PatchedmavenMavenRemote) SetProxyUrlNil() {
	o.ProxyUrl.Set(nil)
}

// UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
func (o *PatchedmavenMavenRemote) UnsetProxyUrl() {
	o.ProxyUrl.Unset()
}

// GetProxyUsername returns the ProxyUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedmavenMavenRemote) GetProxyUsername() string {
	if o == nil || IsNil(o.ProxyUsername.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyUsername.Get()
}

// GetProxyUsernameOk returns a tuple with the ProxyUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedmavenMavenRemote) GetProxyUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyUsername.Get(), o.ProxyUsername.IsSet()
}

// HasProxyUsername returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasProxyUsername() bool {
	if o != nil && o.ProxyUsername.IsSet() {
		return true
	}

	return false
}

// SetProxyUsername gets a reference to the given NullableString and assigns it to the ProxyUsername field.
func (o *PatchedmavenMavenRemote) SetProxyUsername(v string) {
	o.ProxyUsername.Set(&v)
}
// SetProxyUsernameNil sets the value for ProxyUsername to be an explicit nil
func (o *PatchedmavenMavenRemote) SetProxyUsernameNil() {
	o.ProxyUsername.Set(nil)
}

// UnsetProxyUsername ensures that no value is present for ProxyUsername, not even an explicit nil
func (o *PatchedmavenMavenRemote) UnsetProxyUsername() {
	o.ProxyUsername.Unset()
}

// GetProxyPassword returns the ProxyPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedmavenMavenRemote) GetProxyPassword() string {
	if o == nil || IsNil(o.ProxyPassword.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyPassword.Get()
}

// GetProxyPasswordOk returns a tuple with the ProxyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedmavenMavenRemote) GetProxyPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyPassword.Get(), o.ProxyPassword.IsSet()
}

// HasProxyPassword returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasProxyPassword() bool {
	if o != nil && o.ProxyPassword.IsSet() {
		return true
	}

	return false
}

// SetProxyPassword gets a reference to the given NullableString and assigns it to the ProxyPassword field.
func (o *PatchedmavenMavenRemote) SetProxyPassword(v string) {
	o.ProxyPassword.Set(&v)
}
// SetProxyPasswordNil sets the value for ProxyPassword to be an explicit nil
func (o *PatchedmavenMavenRemote) SetProxyPasswordNil() {
	o.ProxyPassword.Set(nil)
}

// UnsetProxyPassword ensures that no value is present for ProxyPassword, not even an explicit nil
func (o *PatchedmavenMavenRemote) UnsetProxyPassword() {
	o.ProxyPassword.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedmavenMavenRemote) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedmavenMavenRemote) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *PatchedmavenMavenRemote) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *PatchedmavenMavenRemote) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *PatchedmavenMavenRemote) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedmavenMavenRemote) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedmavenMavenRemote) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *PatchedmavenMavenRemote) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *PatchedmavenMavenRemote) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *PatchedmavenMavenRemote) UnsetPassword() {
	o.Password.Unset()
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *PatchedmavenMavenRemote) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedmavenMavenRemote) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *PatchedmavenMavenRemote) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetDownloadConcurrency returns the DownloadConcurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedmavenMavenRemote) GetDownloadConcurrency() int32 {
	if o == nil || IsNil(o.DownloadConcurrency.Get()) {
		var ret int32
		return ret
	}
	return *o.DownloadConcurrency.Get()
}

// GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedmavenMavenRemote) GetDownloadConcurrencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadConcurrency.Get(), o.DownloadConcurrency.IsSet()
}

// HasDownloadConcurrency returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasDownloadConcurrency() bool {
	if o != nil && o.DownloadConcurrency.IsSet() {
		return true
	}

	return false
}

// SetDownloadConcurrency gets a reference to the given NullableInt32 and assigns it to the DownloadConcurrency field.
func (o *PatchedmavenMavenRemote) SetDownloadConcurrency(v int32) {
	o.DownloadConcurrency.Set(&v)
}
// SetDownloadConcurrencyNil sets the value for DownloadConcurrency to be an explicit nil
func (o *PatchedmavenMavenRemote) SetDownloadConcurrencyNil() {
	o.DownloadConcurrency.Set(nil)
}

// UnsetDownloadConcurrency ensures that no value is present for DownloadConcurrency, not even an explicit nil
func (o *PatchedmavenMavenRemote) UnsetDownloadConcurrency() {
	o.DownloadConcurrency.Unset()
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedmavenMavenRemote) GetMaxRetries() int32 {
	if o == nil || IsNil(o.MaxRetries.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxRetries.Get()
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedmavenMavenRemote) GetMaxRetriesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRetries.Get(), o.MaxRetries.IsSet()
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasMaxRetries() bool {
	if o != nil && o.MaxRetries.IsSet() {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given NullableInt32 and assigns it to the MaxRetries field.
func (o *PatchedmavenMavenRemote) SetMaxRetries(v int32) {
	o.MaxRetries.Set(&v)
}
// SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil
func (o *PatchedmavenMavenRemote) SetMaxRetriesNil() {
	o.MaxRetries.Set(nil)
}

// UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
func (o *PatchedmavenMavenRemote) UnsetMaxRetries() {
	o.MaxRetries.Unset()
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *PatchedmavenMavenRemote) GetPolicy() PolicyDb6Enum {
	if o == nil || IsNil(o.Policy) {
		var ret PolicyDb6Enum
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedmavenMavenRemote) GetPolicyOk() (*PolicyDb6Enum, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given PolicyDb6Enum and assigns it to the Policy field.
func (o *PatchedmavenMavenRemote) SetPolicy(v PolicyDb6Enum) {
	o.Policy = &v
}

// GetTotalTimeout returns the TotalTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedmavenMavenRemote) GetTotalTimeout() float64 {
	if o == nil || IsNil(o.TotalTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.TotalTimeout.Get()
}

// GetTotalTimeoutOk returns a tuple with the TotalTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedmavenMavenRemote) GetTotalTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalTimeout.Get(), o.TotalTimeout.IsSet()
}

// HasTotalTimeout returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasTotalTimeout() bool {
	if o != nil && o.TotalTimeout.IsSet() {
		return true
	}

	return false
}

// SetTotalTimeout gets a reference to the given NullableFloat64 and assigns it to the TotalTimeout field.
func (o *PatchedmavenMavenRemote) SetTotalTimeout(v float64) {
	o.TotalTimeout.Set(&v)
}
// SetTotalTimeoutNil sets the value for TotalTimeout to be an explicit nil
func (o *PatchedmavenMavenRemote) SetTotalTimeoutNil() {
	o.TotalTimeout.Set(nil)
}

// UnsetTotalTimeout ensures that no value is present for TotalTimeout, not even an explicit nil
func (o *PatchedmavenMavenRemote) UnsetTotalTimeout() {
	o.TotalTimeout.Unset()
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedmavenMavenRemote) GetConnectTimeout() float64 {
	if o == nil || IsNil(o.ConnectTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.ConnectTimeout.Get()
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedmavenMavenRemote) GetConnectTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectTimeout.Get(), o.ConnectTimeout.IsSet()
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasConnectTimeout() bool {
	if o != nil && o.ConnectTimeout.IsSet() {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given NullableFloat64 and assigns it to the ConnectTimeout field.
func (o *PatchedmavenMavenRemote) SetConnectTimeout(v float64) {
	o.ConnectTimeout.Set(&v)
}
// SetConnectTimeoutNil sets the value for ConnectTimeout to be an explicit nil
func (o *PatchedmavenMavenRemote) SetConnectTimeoutNil() {
	o.ConnectTimeout.Set(nil)
}

// UnsetConnectTimeout ensures that no value is present for ConnectTimeout, not even an explicit nil
func (o *PatchedmavenMavenRemote) UnsetConnectTimeout() {
	o.ConnectTimeout.Unset()
}

// GetSockConnectTimeout returns the SockConnectTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedmavenMavenRemote) GetSockConnectTimeout() float64 {
	if o == nil || IsNil(o.SockConnectTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.SockConnectTimeout.Get()
}

// GetSockConnectTimeoutOk returns a tuple with the SockConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedmavenMavenRemote) GetSockConnectTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SockConnectTimeout.Get(), o.SockConnectTimeout.IsSet()
}

// HasSockConnectTimeout returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasSockConnectTimeout() bool {
	if o != nil && o.SockConnectTimeout.IsSet() {
		return true
	}

	return false
}

// SetSockConnectTimeout gets a reference to the given NullableFloat64 and assigns it to the SockConnectTimeout field.
func (o *PatchedmavenMavenRemote) SetSockConnectTimeout(v float64) {
	o.SockConnectTimeout.Set(&v)
}
// SetSockConnectTimeoutNil sets the value for SockConnectTimeout to be an explicit nil
func (o *PatchedmavenMavenRemote) SetSockConnectTimeoutNil() {
	o.SockConnectTimeout.Set(nil)
}

// UnsetSockConnectTimeout ensures that no value is present for SockConnectTimeout, not even an explicit nil
func (o *PatchedmavenMavenRemote) UnsetSockConnectTimeout() {
	o.SockConnectTimeout.Unset()
}

// GetSockReadTimeout returns the SockReadTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedmavenMavenRemote) GetSockReadTimeout() float64 {
	if o == nil || IsNil(o.SockReadTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.SockReadTimeout.Get()
}

// GetSockReadTimeoutOk returns a tuple with the SockReadTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedmavenMavenRemote) GetSockReadTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SockReadTimeout.Get(), o.SockReadTimeout.IsSet()
}

// HasSockReadTimeout returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasSockReadTimeout() bool {
	if o != nil && o.SockReadTimeout.IsSet() {
		return true
	}

	return false
}

// SetSockReadTimeout gets a reference to the given NullableFloat64 and assigns it to the SockReadTimeout field.
func (o *PatchedmavenMavenRemote) SetSockReadTimeout(v float64) {
	o.SockReadTimeout.Set(&v)
}
// SetSockReadTimeoutNil sets the value for SockReadTimeout to be an explicit nil
func (o *PatchedmavenMavenRemote) SetSockReadTimeoutNil() {
	o.SockReadTimeout.Set(nil)
}

// UnsetSockReadTimeout ensures that no value is present for SockReadTimeout, not even an explicit nil
func (o *PatchedmavenMavenRemote) UnsetSockReadTimeout() {
	o.SockReadTimeout.Unset()
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *PatchedmavenMavenRemote) GetHeaders() []map[string]interface{} {
	if o == nil || IsNil(o.Headers) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedmavenMavenRemote) GetHeadersOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []map[string]interface{} and assigns it to the Headers field.
func (o *PatchedmavenMavenRemote) SetHeaders(v []map[string]interface{}) {
	o.Headers = v
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedmavenMavenRemote) GetRateLimit() int32 {
	if o == nil || IsNil(o.RateLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.RateLimit.Get()
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedmavenMavenRemote) GetRateLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RateLimit.Get(), o.RateLimit.IsSet()
}

// HasRateLimit returns a boolean if a field has been set.
func (o *PatchedmavenMavenRemote) HasRateLimit() bool {
	if o != nil && o.RateLimit.IsSet() {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given NullableInt32 and assigns it to the RateLimit field.
func (o *PatchedmavenMavenRemote) SetRateLimit(v int32) {
	o.RateLimit.Set(&v)
}
// SetRateLimitNil sets the value for RateLimit to be an explicit nil
func (o *PatchedmavenMavenRemote) SetRateLimitNil() {
	o.RateLimit.Set(nil)
}

// UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil
func (o *PatchedmavenMavenRemote) UnsetRateLimit() {
	o.RateLimit.Unset()
}

func (o PatchedmavenMavenRemote) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedmavenMavenRemote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
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
	return toSerialize, nil
}

type NullablePatchedmavenMavenRemote struct {
	value *PatchedmavenMavenRemote
	isSet bool
}

func (v NullablePatchedmavenMavenRemote) Get() *PatchedmavenMavenRemote {
	return v.value
}

func (v *NullablePatchedmavenMavenRemote) Set(val *PatchedmavenMavenRemote) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedmavenMavenRemote) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedmavenMavenRemote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedmavenMavenRemote(val *PatchedmavenMavenRemote) *NullablePatchedmavenMavenRemote {
	return &NullablePatchedmavenMavenRemote{value: val, isSet: true}
}

func (v NullablePatchedmavenMavenRemote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedmavenMavenRemote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

