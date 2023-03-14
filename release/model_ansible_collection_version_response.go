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
	"time"
)

// checks if the AnsibleCollectionVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleCollectionVersionResponse{}

// AnsibleCollectionVersionResponse A serializer for CollectionVersion Content.
type AnsibleCollectionVersionResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Artifact file representing the physical content
	Artifact *string `json:"artifact,omitempty"`
	// The SHA-256 checksum if available.
	Sha256 *string `json:"sha256,omitempty"`
	// The MD5 checksum if available.
	Md5 *string `json:"md5,omitempty"`
	// The SHA-1 checksum if available.
	Sha1 *string `json:"sha1,omitempty"`
	// The SHA-224 checksum if available.
	Sha224 *string `json:"sha224,omitempty"`
	// The SHA-384 checksum if available.
	Sha384 *string `json:"sha384,omitempty"`
	// The SHA-512 checksum if available.
	Sha512 *string `json:"sha512,omitempty"`
	// A collection identifier.
	Id *string `json:"id,omitempty"`
	// A list of the CollectionVersion content's authors.
	Authors []string `json:"authors,omitempty"`
	// A JSON field with data about the contents.
	Contents map[string]interface{} `json:"contents,omitempty"`
	// A dict declaring Collections that this collection requires to be installed for it to be usable.
	Dependencies map[string]interface{} `json:"dependencies,omitempty"`
	// A short summary description of the collection.
	Description *string `json:"description,omitempty"`
	// A JSON field holding the various documentation blobs in the collection.
	DocsBlob map[string]interface{} `json:"docs_blob,omitempty"`
	// A JSON field holding MANIFEST.json data.
	Manifest map[string]interface{} `json:"manifest,omitempty"`
	// A JSON field holding FILES.json data.
	Files map[string]interface{} `json:"files,omitempty"`
	// The URL to any online docs.
	Documentation *string `json:"documentation,omitempty"`
	// The URL to the homepage of the collection/project.
	Homepage *string `json:"homepage,omitempty"`
	// The URL to the collection issue tracker.
	Issues *string `json:"issues,omitempty"`
	// A list of licenses for content inside of a collection.
	License []string `json:"license,omitempty"`
	// The name of the collection.
	Name *string `json:"name,omitempty"`
	// The namespace of the collection.
	Namespace *string `json:"namespace,omitempty"`
	// The URL of the originating SCM repository.
	OriginRepository *string `json:"origin_repository,omitempty"`
	Tags []AnsibleTagResponse `json:"tags,omitempty"`
	// The version of the collection.
	Version *string `json:"version,omitempty"`
	// The version of Ansible required to use the collection. Multiple versions can be separated with a comma.
	RequiresAnsible NullableString `json:"requires_ansible,omitempty"`
}

// NewAnsibleCollectionVersionResponse instantiates a new AnsibleCollectionVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleCollectionVersionResponse() *AnsibleCollectionVersionResponse {
	this := AnsibleCollectionVersionResponse{}
	return &this
}

// NewAnsibleCollectionVersionResponseWithDefaults instantiates a new AnsibleCollectionVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleCollectionVersionResponseWithDefaults() *AnsibleCollectionVersionResponse {
	this := AnsibleCollectionVersionResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *AnsibleCollectionVersionResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *AnsibleCollectionVersionResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetArtifact returns the Artifact field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetArtifact() string {
	if o == nil || IsNil(o.Artifact) {
		var ret string
		return ret
	}
	return *o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetArtifactOk() (*string, bool) {
	if o == nil || IsNil(o.Artifact) {
		return nil, false
	}
	return o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasArtifact() bool {
	if o != nil && !IsNil(o.Artifact) {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given string and assigns it to the Artifact field.
func (o *AnsibleCollectionVersionResponse) SetArtifact(v string) {
	o.Artifact = &v
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetSha256() string {
	if o == nil || IsNil(o.Sha256) {
		var ret string
		return ret
	}
	return *o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetSha256Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha256) {
		return nil, false
	}
	return o.Sha256, true
}

// HasSha256 returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasSha256() bool {
	if o != nil && !IsNil(o.Sha256) {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given string and assigns it to the Sha256 field.
func (o *AnsibleCollectionVersionResponse) SetSha256(v string) {
	o.Sha256 = &v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetMd5() string {
	if o == nil || IsNil(o.Md5) {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetMd5Ok() (*string, bool) {
	if o == nil || IsNil(o.Md5) {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasMd5() bool {
	if o != nil && !IsNil(o.Md5) {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *AnsibleCollectionVersionResponse) SetMd5(v string) {
	o.Md5 = &v
}

// GetSha1 returns the Sha1 field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetSha1() string {
	if o == nil || IsNil(o.Sha1) {
		var ret string
		return ret
	}
	return *o.Sha1
}

// GetSha1Ok returns a tuple with the Sha1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetSha1Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha1) {
		return nil, false
	}
	return o.Sha1, true
}

// HasSha1 returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasSha1() bool {
	if o != nil && !IsNil(o.Sha1) {
		return true
	}

	return false
}

// SetSha1 gets a reference to the given string and assigns it to the Sha1 field.
func (o *AnsibleCollectionVersionResponse) SetSha1(v string) {
	o.Sha1 = &v
}

// GetSha224 returns the Sha224 field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetSha224() string {
	if o == nil || IsNil(o.Sha224) {
		var ret string
		return ret
	}
	return *o.Sha224
}

// GetSha224Ok returns a tuple with the Sha224 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetSha224Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha224) {
		return nil, false
	}
	return o.Sha224, true
}

// HasSha224 returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasSha224() bool {
	if o != nil && !IsNil(o.Sha224) {
		return true
	}

	return false
}

// SetSha224 gets a reference to the given string and assigns it to the Sha224 field.
func (o *AnsibleCollectionVersionResponse) SetSha224(v string) {
	o.Sha224 = &v
}

// GetSha384 returns the Sha384 field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetSha384() string {
	if o == nil || IsNil(o.Sha384) {
		var ret string
		return ret
	}
	return *o.Sha384
}

// GetSha384Ok returns a tuple with the Sha384 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetSha384Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha384) {
		return nil, false
	}
	return o.Sha384, true
}

// HasSha384 returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasSha384() bool {
	if o != nil && !IsNil(o.Sha384) {
		return true
	}

	return false
}

// SetSha384 gets a reference to the given string and assigns it to the Sha384 field.
func (o *AnsibleCollectionVersionResponse) SetSha384(v string) {
	o.Sha384 = &v
}

// GetSha512 returns the Sha512 field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetSha512() string {
	if o == nil || IsNil(o.Sha512) {
		var ret string
		return ret
	}
	return *o.Sha512
}

// GetSha512Ok returns a tuple with the Sha512 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetSha512Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha512) {
		return nil, false
	}
	return o.Sha512, true
}

// HasSha512 returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasSha512() bool {
	if o != nil && !IsNil(o.Sha512) {
		return true
	}

	return false
}

// SetSha512 gets a reference to the given string and assigns it to the Sha512 field.
func (o *AnsibleCollectionVersionResponse) SetSha512(v string) {
	o.Sha512 = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AnsibleCollectionVersionResponse) SetId(v string) {
	o.Id = &v
}

// GetAuthors returns the Authors field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetAuthors() []string {
	if o == nil || IsNil(o.Authors) {
		var ret []string
		return ret
	}
	return o.Authors
}

// GetAuthorsOk returns a tuple with the Authors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetAuthorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Authors) {
		return nil, false
	}
	return o.Authors, true
}

// HasAuthors returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasAuthors() bool {
	if o != nil && !IsNil(o.Authors) {
		return true
	}

	return false
}

// SetAuthors gets a reference to the given []string and assigns it to the Authors field.
func (o *AnsibleCollectionVersionResponse) SetAuthors(v []string) {
	o.Authors = v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetContents() map[string]interface{} {
	if o == nil || IsNil(o.Contents) {
		var ret map[string]interface{}
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetContentsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Contents) {
		return map[string]interface{}{}, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given map[string]interface{} and assigns it to the Contents field.
func (o *AnsibleCollectionVersionResponse) SetContents(v map[string]interface{}) {
	o.Contents = v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetDependencies() map[string]interface{} {
	if o == nil || IsNil(o.Dependencies) {
		var ret map[string]interface{}
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetDependenciesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Dependencies) {
		return map[string]interface{}{}, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasDependencies() bool {
	if o != nil && !IsNil(o.Dependencies) {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given map[string]interface{} and assigns it to the Dependencies field.
func (o *AnsibleCollectionVersionResponse) SetDependencies(v map[string]interface{}) {
	o.Dependencies = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AnsibleCollectionVersionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDocsBlob returns the DocsBlob field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetDocsBlob() map[string]interface{} {
	if o == nil || IsNil(o.DocsBlob) {
		var ret map[string]interface{}
		return ret
	}
	return o.DocsBlob
}

// GetDocsBlobOk returns a tuple with the DocsBlob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetDocsBlobOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DocsBlob) {
		return map[string]interface{}{}, false
	}
	return o.DocsBlob, true
}

// HasDocsBlob returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasDocsBlob() bool {
	if o != nil && !IsNil(o.DocsBlob) {
		return true
	}

	return false
}

// SetDocsBlob gets a reference to the given map[string]interface{} and assigns it to the DocsBlob field.
func (o *AnsibleCollectionVersionResponse) SetDocsBlob(v map[string]interface{}) {
	o.DocsBlob = v
}

// GetManifest returns the Manifest field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetManifest() map[string]interface{} {
	if o == nil || IsNil(o.Manifest) {
		var ret map[string]interface{}
		return ret
	}
	return o.Manifest
}

// GetManifestOk returns a tuple with the Manifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetManifestOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Manifest) {
		return map[string]interface{}{}, false
	}
	return o.Manifest, true
}

// HasManifest returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasManifest() bool {
	if o != nil && !IsNil(o.Manifest) {
		return true
	}

	return false
}

// SetManifest gets a reference to the given map[string]interface{} and assigns it to the Manifest field.
func (o *AnsibleCollectionVersionResponse) SetManifest(v map[string]interface{}) {
	o.Manifest = v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetFiles() map[string]interface{} {
	if o == nil || IsNil(o.Files) {
		var ret map[string]interface{}
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetFilesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Files) {
		return map[string]interface{}{}, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given map[string]interface{} and assigns it to the Files field.
func (o *AnsibleCollectionVersionResponse) SetFiles(v map[string]interface{}) {
	o.Files = v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetDocumentation() string {
	if o == nil || IsNil(o.Documentation) {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetDocumentationOk() (*string, bool) {
	if o == nil || IsNil(o.Documentation) {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasDocumentation() bool {
	if o != nil && !IsNil(o.Documentation) {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *AnsibleCollectionVersionResponse) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetHomepage returns the Homepage field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetHomepage() string {
	if o == nil || IsNil(o.Homepage) {
		var ret string
		return ret
	}
	return *o.Homepage
}

// GetHomepageOk returns a tuple with the Homepage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetHomepageOk() (*string, bool) {
	if o == nil || IsNil(o.Homepage) {
		return nil, false
	}
	return o.Homepage, true
}

// HasHomepage returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasHomepage() bool {
	if o != nil && !IsNil(o.Homepage) {
		return true
	}

	return false
}

// SetHomepage gets a reference to the given string and assigns it to the Homepage field.
func (o *AnsibleCollectionVersionResponse) SetHomepage(v string) {
	o.Homepage = &v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetIssues() string {
	if o == nil || IsNil(o.Issues) {
		var ret string
		return ret
	}
	return *o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetIssuesOk() (*string, bool) {
	if o == nil || IsNil(o.Issues) {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasIssues() bool {
	if o != nil && !IsNil(o.Issues) {
		return true
	}

	return false
}

// SetIssues gets a reference to the given string and assigns it to the Issues field.
func (o *AnsibleCollectionVersionResponse) SetIssues(v string) {
	o.Issues = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetLicense() []string {
	if o == nil || IsNil(o.License) {
		var ret []string
		return ret
	}
	return o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetLicenseOk() ([]string, bool) {
	if o == nil || IsNil(o.License) {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasLicense() bool {
	if o != nil && !IsNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given []string and assigns it to the License field.
func (o *AnsibleCollectionVersionResponse) SetLicense(v []string) {
	o.License = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AnsibleCollectionVersionResponse) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *AnsibleCollectionVersionResponse) SetNamespace(v string) {
	o.Namespace = &v
}

// GetOriginRepository returns the OriginRepository field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetOriginRepository() string {
	if o == nil || IsNil(o.OriginRepository) {
		var ret string
		return ret
	}
	return *o.OriginRepository
}

// GetOriginRepositoryOk returns a tuple with the OriginRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetOriginRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.OriginRepository) {
		return nil, false
	}
	return o.OriginRepository, true
}

// HasOriginRepository returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasOriginRepository() bool {
	if o != nil && !IsNil(o.OriginRepository) {
		return true
	}

	return false
}

// SetOriginRepository gets a reference to the given string and assigns it to the OriginRepository field.
func (o *AnsibleCollectionVersionResponse) SetOriginRepository(v string) {
	o.OriginRepository = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetTags() []AnsibleTagResponse {
	if o == nil || IsNil(o.Tags) {
		var ret []AnsibleTagResponse
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetTagsOk() ([]AnsibleTagResponse, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []AnsibleTagResponse and assigns it to the Tags field.
func (o *AnsibleCollectionVersionResponse) SetTags(v []AnsibleTagResponse) {
	o.Tags = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AnsibleCollectionVersionResponse) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionResponse) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AnsibleCollectionVersionResponse) SetVersion(v string) {
	o.Version = &v
}

// GetRequiresAnsible returns the RequiresAnsible field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleCollectionVersionResponse) GetRequiresAnsible() string {
	if o == nil || IsNil(o.RequiresAnsible.Get()) {
		var ret string
		return ret
	}
	return *o.RequiresAnsible.Get()
}

// GetRequiresAnsibleOk returns a tuple with the RequiresAnsible field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleCollectionVersionResponse) GetRequiresAnsibleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiresAnsible.Get(), o.RequiresAnsible.IsSet()
}

// HasRequiresAnsible returns a boolean if a field has been set.
func (o *AnsibleCollectionVersionResponse) HasRequiresAnsible() bool {
	if o != nil && o.RequiresAnsible.IsSet() {
		return true
	}

	return false
}

// SetRequiresAnsible gets a reference to the given NullableString and assigns it to the RequiresAnsible field.
func (o *AnsibleCollectionVersionResponse) SetRequiresAnsible(v string) {
	o.RequiresAnsible.Set(&v)
}
// SetRequiresAnsibleNil sets the value for RequiresAnsible to be an explicit nil
func (o *AnsibleCollectionVersionResponse) SetRequiresAnsibleNil() {
	o.RequiresAnsible.Set(nil)
}

// UnsetRequiresAnsible ensures that no value is present for RequiresAnsible, not even an explicit nil
func (o *AnsibleCollectionVersionResponse) UnsetRequiresAnsible() {
	o.RequiresAnsible.Unset()
}

func (o AnsibleCollectionVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleCollectionVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	if !IsNil(o.Artifact) {
		toSerialize["artifact"] = o.Artifact
	}
	// skip: sha256 is readOnly
	// skip: md5 is readOnly
	// skip: sha1 is readOnly
	// skip: sha224 is readOnly
	// skip: sha384 is readOnly
	// skip: sha512 is readOnly
	// skip: id is readOnly
	// skip: authors is readOnly
	// skip: contents is readOnly
	// skip: dependencies is readOnly
	// skip: description is readOnly
	// skip: docs_blob is readOnly
	// skip: manifest is readOnly
	// skip: files is readOnly
	// skip: documentation is readOnly
	// skip: homepage is readOnly
	// skip: issues is readOnly
	// skip: license is readOnly
	// skip: name is readOnly
	// skip: namespace is readOnly
	// skip: origin_repository is readOnly
	// skip: tags is readOnly
	// skip: version is readOnly
	if o.RequiresAnsible.IsSet() {
		toSerialize["requires_ansible"] = o.RequiresAnsible.Get()
	}
	return toSerialize, nil
}

type NullableAnsibleCollectionVersionResponse struct {
	value *AnsibleCollectionVersionResponse
	isSet bool
}

func (v NullableAnsibleCollectionVersionResponse) Get() *AnsibleCollectionVersionResponse {
	return v.value
}

func (v *NullableAnsibleCollectionVersionResponse) Set(val *AnsibleCollectionVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleCollectionVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleCollectionVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleCollectionVersionResponse(val *AnsibleCollectionVersionResponse) *NullableAnsibleCollectionVersionResponse {
	return &NullableAnsibleCollectionVersionResponse{value: val, isSet: true}
}

func (v NullableAnsibleCollectionVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleCollectionVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

