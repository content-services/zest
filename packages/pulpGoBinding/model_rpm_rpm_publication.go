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

// checks if the RpmRpmPublication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmRpmPublication{}

// RpmRpmPublication A Serializer for RpmPublication.
type RpmRpmPublication struct {
	RepositoryVersion *string `json:"repository_version,omitempty"`
	// A URI of the repository to be published.
	Repository *string `json:"repository,omitempty"`
	// The checksum type for metadata.
	MetadataChecksumType *MetadataChecksumTypeEnum `json:"metadata_checksum_type,omitempty"`
	// The checksum type for packages.
	PackageChecksumType *PackageChecksumTypeEnum `json:"package_checksum_type,omitempty"`
	// An option specifying whether a client should perform a GPG signature check on packages.
	Gpgcheck *int32 `json:"gpgcheck,omitempty"`
	// An option specifying whether a client should perform a GPG signature check on the repodata.
	RepoGpgcheck *int32 `json:"repo_gpgcheck,omitempty"`
	// DEPRECATED: An option specifying whether Pulp should generate SQLite metadata.
	SqliteMetadata *bool `json:"sqlite_metadata,omitempty"`
}

// NewRpmRpmPublication instantiates a new RpmRpmPublication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmRpmPublication() *RpmRpmPublication {
	this := RpmRpmPublication{}
	var sqliteMetadata bool = false
	this.SqliteMetadata = &sqliteMetadata
	return &this
}

// NewRpmRpmPublicationWithDefaults instantiates a new RpmRpmPublication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmRpmPublicationWithDefaults() *RpmRpmPublication {
	this := RpmRpmPublication{}
	var sqliteMetadata bool = false
	this.SqliteMetadata = &sqliteMetadata
	return &this
}

// GetRepositoryVersion returns the RepositoryVersion field value if set, zero value otherwise.
func (o *RpmRpmPublication) GetRepositoryVersion() string {
	if o == nil || IsNil(o.RepositoryVersion) {
		var ret string
		return ret
	}
	return *o.RepositoryVersion
}

// GetRepositoryVersionOk returns a tuple with the RepositoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmPublication) GetRepositoryVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryVersion) {
		return nil, false
	}
	return o.RepositoryVersion, true
}

// HasRepositoryVersion returns a boolean if a field has been set.
func (o *RpmRpmPublication) HasRepositoryVersion() bool {
	if o != nil && !IsNil(o.RepositoryVersion) {
		return true
	}

	return false
}

// SetRepositoryVersion gets a reference to the given string and assigns it to the RepositoryVersion field.
func (o *RpmRpmPublication) SetRepositoryVersion(v string) {
	o.RepositoryVersion = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *RpmRpmPublication) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmPublication) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *RpmRpmPublication) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *RpmRpmPublication) SetRepository(v string) {
	o.Repository = &v
}

// GetMetadataChecksumType returns the MetadataChecksumType field value if set, zero value otherwise.
func (o *RpmRpmPublication) GetMetadataChecksumType() MetadataChecksumTypeEnum {
	if o == nil || IsNil(o.MetadataChecksumType) {
		var ret MetadataChecksumTypeEnum
		return ret
	}
	return *o.MetadataChecksumType
}

// GetMetadataChecksumTypeOk returns a tuple with the MetadataChecksumType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmPublication) GetMetadataChecksumTypeOk() (*MetadataChecksumTypeEnum, bool) {
	if o == nil || IsNil(o.MetadataChecksumType) {
		return nil, false
	}
	return o.MetadataChecksumType, true
}

// HasMetadataChecksumType returns a boolean if a field has been set.
func (o *RpmRpmPublication) HasMetadataChecksumType() bool {
	if o != nil && !IsNil(o.MetadataChecksumType) {
		return true
	}

	return false
}

// SetMetadataChecksumType gets a reference to the given MetadataChecksumTypeEnum and assigns it to the MetadataChecksumType field.
func (o *RpmRpmPublication) SetMetadataChecksumType(v MetadataChecksumTypeEnum) {
	o.MetadataChecksumType = &v
}

// GetPackageChecksumType returns the PackageChecksumType field value if set, zero value otherwise.
func (o *RpmRpmPublication) GetPackageChecksumType() PackageChecksumTypeEnum {
	if o == nil || IsNil(o.PackageChecksumType) {
		var ret PackageChecksumTypeEnum
		return ret
	}
	return *o.PackageChecksumType
}

// GetPackageChecksumTypeOk returns a tuple with the PackageChecksumType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmPublication) GetPackageChecksumTypeOk() (*PackageChecksumTypeEnum, bool) {
	if o == nil || IsNil(o.PackageChecksumType) {
		return nil, false
	}
	return o.PackageChecksumType, true
}

// HasPackageChecksumType returns a boolean if a field has been set.
func (o *RpmRpmPublication) HasPackageChecksumType() bool {
	if o != nil && !IsNil(o.PackageChecksumType) {
		return true
	}

	return false
}

// SetPackageChecksumType gets a reference to the given PackageChecksumTypeEnum and assigns it to the PackageChecksumType field.
func (o *RpmRpmPublication) SetPackageChecksumType(v PackageChecksumTypeEnum) {
	o.PackageChecksumType = &v
}

// GetGpgcheck returns the Gpgcheck field value if set, zero value otherwise.
func (o *RpmRpmPublication) GetGpgcheck() int32 {
	if o == nil || IsNil(o.Gpgcheck) {
		var ret int32
		return ret
	}
	return *o.Gpgcheck
}

// GetGpgcheckOk returns a tuple with the Gpgcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmPublication) GetGpgcheckOk() (*int32, bool) {
	if o == nil || IsNil(o.Gpgcheck) {
		return nil, false
	}
	return o.Gpgcheck, true
}

// HasGpgcheck returns a boolean if a field has been set.
func (o *RpmRpmPublication) HasGpgcheck() bool {
	if o != nil && !IsNil(o.Gpgcheck) {
		return true
	}

	return false
}

// SetGpgcheck gets a reference to the given int32 and assigns it to the Gpgcheck field.
func (o *RpmRpmPublication) SetGpgcheck(v int32) {
	o.Gpgcheck = &v
}

// GetRepoGpgcheck returns the RepoGpgcheck field value if set, zero value otherwise.
func (o *RpmRpmPublication) GetRepoGpgcheck() int32 {
	if o == nil || IsNil(o.RepoGpgcheck) {
		var ret int32
		return ret
	}
	return *o.RepoGpgcheck
}

// GetRepoGpgcheckOk returns a tuple with the RepoGpgcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmPublication) GetRepoGpgcheckOk() (*int32, bool) {
	if o == nil || IsNil(o.RepoGpgcheck) {
		return nil, false
	}
	return o.RepoGpgcheck, true
}

// HasRepoGpgcheck returns a boolean if a field has been set.
func (o *RpmRpmPublication) HasRepoGpgcheck() bool {
	if o != nil && !IsNil(o.RepoGpgcheck) {
		return true
	}

	return false
}

// SetRepoGpgcheck gets a reference to the given int32 and assigns it to the RepoGpgcheck field.
func (o *RpmRpmPublication) SetRepoGpgcheck(v int32) {
	o.RepoGpgcheck = &v
}

// GetSqliteMetadata returns the SqliteMetadata field value if set, zero value otherwise.
func (o *RpmRpmPublication) GetSqliteMetadata() bool {
	if o == nil || IsNil(o.SqliteMetadata) {
		var ret bool
		return ret
	}
	return *o.SqliteMetadata
}

// GetSqliteMetadataOk returns a tuple with the SqliteMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmPublication) GetSqliteMetadataOk() (*bool, bool) {
	if o == nil || IsNil(o.SqliteMetadata) {
		return nil, false
	}
	return o.SqliteMetadata, true
}

// HasSqliteMetadata returns a boolean if a field has been set.
func (o *RpmRpmPublication) HasSqliteMetadata() bool {
	if o != nil && !IsNil(o.SqliteMetadata) {
		return true
	}

	return false
}

// SetSqliteMetadata gets a reference to the given bool and assigns it to the SqliteMetadata field.
func (o *RpmRpmPublication) SetSqliteMetadata(v bool) {
	o.SqliteMetadata = &v
}

func (o RpmRpmPublication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmRpmPublication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RepositoryVersion) {
		toSerialize["repository_version"] = o.RepositoryVersion
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.MetadataChecksumType) {
		toSerialize["metadata_checksum_type"] = o.MetadataChecksumType
	}
	if !IsNil(o.PackageChecksumType) {
		toSerialize["package_checksum_type"] = o.PackageChecksumType
	}
	if !IsNil(o.Gpgcheck) {
		toSerialize["gpgcheck"] = o.Gpgcheck
	}
	if !IsNil(o.RepoGpgcheck) {
		toSerialize["repo_gpgcheck"] = o.RepoGpgcheck
	}
	if !IsNil(o.SqliteMetadata) {
		toSerialize["sqlite_metadata"] = o.SqliteMetadata
	}
	return toSerialize, nil
}

type NullableRpmRpmPublication struct {
	value *RpmRpmPublication
	isSet bool
}

func (v NullableRpmRpmPublication) Get() *RpmRpmPublication {
	return v.value
}

func (v *NullableRpmRpmPublication) Set(val *RpmRpmPublication) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmRpmPublication) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmRpmPublication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmRpmPublication(val *RpmRpmPublication) *NullableRpmRpmPublication {
	return &NullableRpmRpmPublication{value: val, isSet: true}
}

func (v NullableRpmRpmPublication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmRpmPublication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


