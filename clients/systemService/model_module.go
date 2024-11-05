/*
SystemService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.1.4089
Contact: support@fenix-alliance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Module type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Module{}

// Module struct for Module
type Module struct {
	Enable *bool `json:"enable,omitempty"`
	Active *bool `json:"active,omitempty"`
	Order *int32 `json:"order,omitempty"`
	Id NullableString `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	FullName NullableString `json:"fullName,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Type *int32 `json:"type,omitempty"`
	Configuration NullableString `json:"configuration,omitempty"`
	Author NullableString `json:"author,omitempty"`
	AuthorUrl NullableString `json:"authorUrl,omitempty"`
	License NullableString `json:"license,omitempty"`
	RequireLicenseAcceptance NullableBool `json:"requireLicenseAcceptance,omitempty"`
	Repository NullableString `json:"repository,omitempty"`
	Path NullableString `json:"path,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	Image NullableString `json:"image,omitempty"`
	NuSpecPath NullableString `json:"nuSpecPath,omitempty"`
	Manifest NullableString `json:"manifest,omitempty"`
	Documentation NullableString `json:"documentation,omitempty"`
	Website NullableString `json:"website,omitempty"`
	Logo NullableString `json:"logo,omitempty"`
	SwaggerSpec *ISwaggerSpec `json:"swaggerSpec,omitempty"`
	SwaggerSpecs []ISwaggerSpec `json:"swaggerSpecs,omitempty"`
	Url NullableString `json:"url,omitempty"`
	AssemblyPaths []string `json:"assemblyPaths,omitempty"`
	MarkedForDeletion *bool `json:"markedForDeletion,omitempty"`
	Version NullableString `json:"version,omitempty"`
}

// NewModule instantiates a new Module object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModule() *Module {
	this := Module{}
	return &this
}

// NewModuleWithDefaults instantiates a new Module object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleWithDefaults() *Module {
	this := Module{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *Module) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Module) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *Module) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *Module) SetEnable(v bool) {
	o.Enable = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Module) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Module) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Module) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Module) SetActive(v bool) {
	o.Active = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *Module) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Module) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *Module) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *Module) SetOrder(v int32) {
	o.Order = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Module) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *Module) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Module) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Module) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Module) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Module) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Module) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Module) UnsetName() {
	o.Name.Unset()
}

// GetFullName returns the FullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetFullName() string {
	if o == nil || IsNil(o.FullName.Get()) {
		var ret string
		return ret
	}
	return *o.FullName.Get()
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullName.Get(), o.FullName.IsSet()
}

// HasFullName returns a boolean if a field has been set.
func (o *Module) HasFullName() bool {
	if o != nil && o.FullName.IsSet() {
		return true
	}

	return false
}

// SetFullName gets a reference to the given NullableString and assigns it to the FullName field.
func (o *Module) SetFullName(v string) {
	o.FullName.Set(&v)
}
// SetFullNameNil sets the value for FullName to be an explicit nil
func (o *Module) SetFullNameNil() {
	o.FullName.Set(nil)
}

// UnsetFullName ensures that no value is present for FullName, not even an explicit nil
func (o *Module) UnsetFullName() {
	o.FullName.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Module) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Module) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Module) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Module) UnsetDescription() {
	o.Description.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Module) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Module) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Module) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *Module) SetType(v int32) {
	o.Type = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetConfiguration() string {
	if o == nil || IsNil(o.Configuration.Get()) {
		var ret string
		return ret
	}
	return *o.Configuration.Get()
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetConfigurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Configuration.Get(), o.Configuration.IsSet()
}

// HasConfiguration returns a boolean if a field has been set.
func (o *Module) HasConfiguration() bool {
	if o != nil && o.Configuration.IsSet() {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given NullableString and assigns it to the Configuration field.
func (o *Module) SetConfiguration(v string) {
	o.Configuration.Set(&v)
}
// SetConfigurationNil sets the value for Configuration to be an explicit nil
func (o *Module) SetConfigurationNil() {
	o.Configuration.Set(nil)
}

// UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
func (o *Module) UnsetConfiguration() {
	o.Configuration.Unset()
}

// GetAuthor returns the Author field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetAuthor() string {
	if o == nil || IsNil(o.Author.Get()) {
		var ret string
		return ret
	}
	return *o.Author.Get()
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Author.Get(), o.Author.IsSet()
}

// HasAuthor returns a boolean if a field has been set.
func (o *Module) HasAuthor() bool {
	if o != nil && o.Author.IsSet() {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given NullableString and assigns it to the Author field.
func (o *Module) SetAuthor(v string) {
	o.Author.Set(&v)
}
// SetAuthorNil sets the value for Author to be an explicit nil
func (o *Module) SetAuthorNil() {
	o.Author.Set(nil)
}

// UnsetAuthor ensures that no value is present for Author, not even an explicit nil
func (o *Module) UnsetAuthor() {
	o.Author.Unset()
}

// GetAuthorUrl returns the AuthorUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetAuthorUrl() string {
	if o == nil || IsNil(o.AuthorUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AuthorUrl.Get()
}

// GetAuthorUrlOk returns a tuple with the AuthorUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetAuthorUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorUrl.Get(), o.AuthorUrl.IsSet()
}

// HasAuthorUrl returns a boolean if a field has been set.
func (o *Module) HasAuthorUrl() bool {
	if o != nil && o.AuthorUrl.IsSet() {
		return true
	}

	return false
}

// SetAuthorUrl gets a reference to the given NullableString and assigns it to the AuthorUrl field.
func (o *Module) SetAuthorUrl(v string) {
	o.AuthorUrl.Set(&v)
}
// SetAuthorUrlNil sets the value for AuthorUrl to be an explicit nil
func (o *Module) SetAuthorUrlNil() {
	o.AuthorUrl.Set(nil)
}

// UnsetAuthorUrl ensures that no value is present for AuthorUrl, not even an explicit nil
func (o *Module) UnsetAuthorUrl() {
	o.AuthorUrl.Unset()
}

// GetLicense returns the License field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetLicense() string {
	if o == nil || IsNil(o.License.Get()) {
		var ret string
		return ret
	}
	return *o.License.Get()
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetLicenseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.License.Get(), o.License.IsSet()
}

// HasLicense returns a boolean if a field has been set.
func (o *Module) HasLicense() bool {
	if o != nil && o.License.IsSet() {
		return true
	}

	return false
}

// SetLicense gets a reference to the given NullableString and assigns it to the License field.
func (o *Module) SetLicense(v string) {
	o.License.Set(&v)
}
// SetLicenseNil sets the value for License to be an explicit nil
func (o *Module) SetLicenseNil() {
	o.License.Set(nil)
}

// UnsetLicense ensures that no value is present for License, not even an explicit nil
func (o *Module) UnsetLicense() {
	o.License.Unset()
}

// GetRequireLicenseAcceptance returns the RequireLicenseAcceptance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetRequireLicenseAcceptance() bool {
	if o == nil || IsNil(o.RequireLicenseAcceptance.Get()) {
		var ret bool
		return ret
	}
	return *o.RequireLicenseAcceptance.Get()
}

// GetRequireLicenseAcceptanceOk returns a tuple with the RequireLicenseAcceptance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetRequireLicenseAcceptanceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequireLicenseAcceptance.Get(), o.RequireLicenseAcceptance.IsSet()
}

// HasRequireLicenseAcceptance returns a boolean if a field has been set.
func (o *Module) HasRequireLicenseAcceptance() bool {
	if o != nil && o.RequireLicenseAcceptance.IsSet() {
		return true
	}

	return false
}

// SetRequireLicenseAcceptance gets a reference to the given NullableBool and assigns it to the RequireLicenseAcceptance field.
func (o *Module) SetRequireLicenseAcceptance(v bool) {
	o.RequireLicenseAcceptance.Set(&v)
}
// SetRequireLicenseAcceptanceNil sets the value for RequireLicenseAcceptance to be an explicit nil
func (o *Module) SetRequireLicenseAcceptanceNil() {
	o.RequireLicenseAcceptance.Set(nil)
}

// UnsetRequireLicenseAcceptance ensures that no value is present for RequireLicenseAcceptance, not even an explicit nil
func (o *Module) UnsetRequireLicenseAcceptance() {
	o.RequireLicenseAcceptance.Unset()
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetRepository() string {
	if o == nil || IsNil(o.Repository.Get()) {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *Module) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *Module) SetRepository(v string) {
	o.Repository.Set(&v)
}
// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *Module) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *Module) UnsetRepository() {
	o.Repository.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *Module) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *Module) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *Module) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *Module) UnsetPath() {
	o.Path.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *Module) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *Module) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *Module) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *Module) UnsetIcon() {
	o.Icon.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *Module) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *Module) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *Module) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *Module) UnsetImage() {
	o.Image.Unset()
}

// GetNuSpecPath returns the NuSpecPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetNuSpecPath() string {
	if o == nil || IsNil(o.NuSpecPath.Get()) {
		var ret string
		return ret
	}
	return *o.NuSpecPath.Get()
}

// GetNuSpecPathOk returns a tuple with the NuSpecPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetNuSpecPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NuSpecPath.Get(), o.NuSpecPath.IsSet()
}

// HasNuSpecPath returns a boolean if a field has been set.
func (o *Module) HasNuSpecPath() bool {
	if o != nil && o.NuSpecPath.IsSet() {
		return true
	}

	return false
}

// SetNuSpecPath gets a reference to the given NullableString and assigns it to the NuSpecPath field.
func (o *Module) SetNuSpecPath(v string) {
	o.NuSpecPath.Set(&v)
}
// SetNuSpecPathNil sets the value for NuSpecPath to be an explicit nil
func (o *Module) SetNuSpecPathNil() {
	o.NuSpecPath.Set(nil)
}

// UnsetNuSpecPath ensures that no value is present for NuSpecPath, not even an explicit nil
func (o *Module) UnsetNuSpecPath() {
	o.NuSpecPath.Unset()
}

// GetManifest returns the Manifest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetManifest() string {
	if o == nil || IsNil(o.Manifest.Get()) {
		var ret string
		return ret
	}
	return *o.Manifest.Get()
}

// GetManifestOk returns a tuple with the Manifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetManifestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manifest.Get(), o.Manifest.IsSet()
}

// HasManifest returns a boolean if a field has been set.
func (o *Module) HasManifest() bool {
	if o != nil && o.Manifest.IsSet() {
		return true
	}

	return false
}

// SetManifest gets a reference to the given NullableString and assigns it to the Manifest field.
func (o *Module) SetManifest(v string) {
	o.Manifest.Set(&v)
}
// SetManifestNil sets the value for Manifest to be an explicit nil
func (o *Module) SetManifestNil() {
	o.Manifest.Set(nil)
}

// UnsetManifest ensures that no value is present for Manifest, not even an explicit nil
func (o *Module) UnsetManifest() {
	o.Manifest.Unset()
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetDocumentation() string {
	if o == nil || IsNil(o.Documentation.Get()) {
		var ret string
		return ret
	}
	return *o.Documentation.Get()
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetDocumentationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Documentation.Get(), o.Documentation.IsSet()
}

// HasDocumentation returns a boolean if a field has been set.
func (o *Module) HasDocumentation() bool {
	if o != nil && o.Documentation.IsSet() {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given NullableString and assigns it to the Documentation field.
func (o *Module) SetDocumentation(v string) {
	o.Documentation.Set(&v)
}
// SetDocumentationNil sets the value for Documentation to be an explicit nil
func (o *Module) SetDocumentationNil() {
	o.Documentation.Set(nil)
}

// UnsetDocumentation ensures that no value is present for Documentation, not even an explicit nil
func (o *Module) UnsetDocumentation() {
	o.Documentation.Unset()
}

// GetWebsite returns the Website field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetWebsite() string {
	if o == nil || IsNil(o.Website.Get()) {
		var ret string
		return ret
	}
	return *o.Website.Get()
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetWebsiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Website.Get(), o.Website.IsSet()
}

// HasWebsite returns a boolean if a field has been set.
func (o *Module) HasWebsite() bool {
	if o != nil && o.Website.IsSet() {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given NullableString and assigns it to the Website field.
func (o *Module) SetWebsite(v string) {
	o.Website.Set(&v)
}
// SetWebsiteNil sets the value for Website to be an explicit nil
func (o *Module) SetWebsiteNil() {
	o.Website.Set(nil)
}

// UnsetWebsite ensures that no value is present for Website, not even an explicit nil
func (o *Module) UnsetWebsite() {
	o.Website.Unset()
}

// GetLogo returns the Logo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetLogo() string {
	if o == nil || IsNil(o.Logo.Get()) {
		var ret string
		return ret
	}
	return *o.Logo.Get()
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logo.Get(), o.Logo.IsSet()
}

// HasLogo returns a boolean if a field has been set.
func (o *Module) HasLogo() bool {
	if o != nil && o.Logo.IsSet() {
		return true
	}

	return false
}

// SetLogo gets a reference to the given NullableString and assigns it to the Logo field.
func (o *Module) SetLogo(v string) {
	o.Logo.Set(&v)
}
// SetLogoNil sets the value for Logo to be an explicit nil
func (o *Module) SetLogoNil() {
	o.Logo.Set(nil)
}

// UnsetLogo ensures that no value is present for Logo, not even an explicit nil
func (o *Module) UnsetLogo() {
	o.Logo.Unset()
}

// GetSwaggerSpec returns the SwaggerSpec field value if set, zero value otherwise.
func (o *Module) GetSwaggerSpec() ISwaggerSpec {
	if o == nil || IsNil(o.SwaggerSpec) {
		var ret ISwaggerSpec
		return ret
	}
	return *o.SwaggerSpec
}

// GetSwaggerSpecOk returns a tuple with the SwaggerSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Module) GetSwaggerSpecOk() (*ISwaggerSpec, bool) {
	if o == nil || IsNil(o.SwaggerSpec) {
		return nil, false
	}
	return o.SwaggerSpec, true
}

// HasSwaggerSpec returns a boolean if a field has been set.
func (o *Module) HasSwaggerSpec() bool {
	if o != nil && !IsNil(o.SwaggerSpec) {
		return true
	}

	return false
}

// SetSwaggerSpec gets a reference to the given ISwaggerSpec and assigns it to the SwaggerSpec field.
func (o *Module) SetSwaggerSpec(v ISwaggerSpec) {
	o.SwaggerSpec = &v
}

// GetSwaggerSpecs returns the SwaggerSpecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetSwaggerSpecs() []ISwaggerSpec {
	if o == nil {
		var ret []ISwaggerSpec
		return ret
	}
	return o.SwaggerSpecs
}

// GetSwaggerSpecsOk returns a tuple with the SwaggerSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetSwaggerSpecsOk() ([]ISwaggerSpec, bool) {
	if o == nil || IsNil(o.SwaggerSpecs) {
		return nil, false
	}
	return o.SwaggerSpecs, true
}

// HasSwaggerSpecs returns a boolean if a field has been set.
func (o *Module) HasSwaggerSpecs() bool {
	if o != nil && !IsNil(o.SwaggerSpecs) {
		return true
	}

	return false
}

// SetSwaggerSpecs gets a reference to the given []ISwaggerSpec and assigns it to the SwaggerSpecs field.
func (o *Module) SetSwaggerSpecs(v []ISwaggerSpec) {
	o.SwaggerSpecs = v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *Module) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *Module) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *Module) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *Module) UnsetUrl() {
	o.Url.Unset()
}

// GetAssemblyPaths returns the AssemblyPaths field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetAssemblyPaths() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AssemblyPaths
}

// GetAssemblyPathsOk returns a tuple with the AssemblyPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetAssemblyPathsOk() ([]string, bool) {
	if o == nil || IsNil(o.AssemblyPaths) {
		return nil, false
	}
	return o.AssemblyPaths, true
}

// HasAssemblyPaths returns a boolean if a field has been set.
func (o *Module) HasAssemblyPaths() bool {
	if o != nil && !IsNil(o.AssemblyPaths) {
		return true
	}

	return false
}

// SetAssemblyPaths gets a reference to the given []string and assigns it to the AssemblyPaths field.
func (o *Module) SetAssemblyPaths(v []string) {
	o.AssemblyPaths = v
}

// GetMarkedForDeletion returns the MarkedForDeletion field value if set, zero value otherwise.
func (o *Module) GetMarkedForDeletion() bool {
	if o == nil || IsNil(o.MarkedForDeletion) {
		var ret bool
		return ret
	}
	return *o.MarkedForDeletion
}

// GetMarkedForDeletionOk returns a tuple with the MarkedForDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Module) GetMarkedForDeletionOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkedForDeletion) {
		return nil, false
	}
	return o.MarkedForDeletion, true
}

// HasMarkedForDeletion returns a boolean if a field has been set.
func (o *Module) HasMarkedForDeletion() bool {
	if o != nil && !IsNil(o.MarkedForDeletion) {
		return true
	}

	return false
}

// SetMarkedForDeletion gets a reference to the given bool and assigns it to the MarkedForDeletion field.
func (o *Module) SetMarkedForDeletion(v bool) {
	o.MarkedForDeletion = &v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *Module) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *Module) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *Module) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *Module) UnsetVersion() {
	o.Version.Unset()
}

func (o Module) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Module) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.FullName.IsSet() {
		toSerialize["fullName"] = o.FullName.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Configuration.IsSet() {
		toSerialize["configuration"] = o.Configuration.Get()
	}
	if o.Author.IsSet() {
		toSerialize["author"] = o.Author.Get()
	}
	if o.AuthorUrl.IsSet() {
		toSerialize["authorUrl"] = o.AuthorUrl.Get()
	}
	if o.License.IsSet() {
		toSerialize["license"] = o.License.Get()
	}
	if o.RequireLicenseAcceptance.IsSet() {
		toSerialize["requireLicenseAcceptance"] = o.RequireLicenseAcceptance.Get()
	}
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if o.NuSpecPath.IsSet() {
		toSerialize["nuSpecPath"] = o.NuSpecPath.Get()
	}
	if o.Manifest.IsSet() {
		toSerialize["manifest"] = o.Manifest.Get()
	}
	if o.Documentation.IsSet() {
		toSerialize["documentation"] = o.Documentation.Get()
	}
	if o.Website.IsSet() {
		toSerialize["website"] = o.Website.Get()
	}
	if o.Logo.IsSet() {
		toSerialize["logo"] = o.Logo.Get()
	}
	if !IsNil(o.SwaggerSpec) {
		toSerialize["swaggerSpec"] = o.SwaggerSpec
	}
	if o.SwaggerSpecs != nil {
		toSerialize["swaggerSpecs"] = o.SwaggerSpecs
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.AssemblyPaths != nil {
		toSerialize["assemblyPaths"] = o.AssemblyPaths
	}
	if !IsNil(o.MarkedForDeletion) {
		toSerialize["markedForDeletion"] = o.MarkedForDeletion
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	return toSerialize, nil
}

type NullableModule struct {
	value *Module
	isSet bool
}

func (v NullableModule) Get() *Module {
	return v.value
}

func (v *NullableModule) Set(val *Module) {
	v.value = val
	v.isSet = true
}

func (v NullableModule) IsSet() bool {
	return v.isSet
}

func (v *NullableModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModule(val *Module) *NullableModule {
	return &NullableModule{value: val, isSet: true}
}

func (v NullableModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


