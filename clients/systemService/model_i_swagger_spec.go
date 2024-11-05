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

// checks if the ISwaggerSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ISwaggerSpec{}

// ISwaggerSpec struct for ISwaggerSpec
type ISwaggerSpec struct {
	Enable *bool `json:"enable,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Version NullableString `json:"version,omitempty"`
	Description NullableString `json:"description,omitempty"`
	TermsOfService NullableString `json:"termsOfService,omitempty"`
	SwaggerEndpoint *ISwaggerEndpoint `json:"swaggerEndpoint,omitempty"`
	OpenApiContact *ISwaggerContact `json:"openApiContact,omitempty"`
	License *ISwaggerLicense `json:"license,omitempty"`
}

// NewISwaggerSpec instantiates a new ISwaggerSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewISwaggerSpec() *ISwaggerSpec {
	this := ISwaggerSpec{}
	return &this
}

// NewISwaggerSpecWithDefaults instantiates a new ISwaggerSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewISwaggerSpecWithDefaults() *ISwaggerSpec {
	this := ISwaggerSpec{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *ISwaggerSpec) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISwaggerSpec) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *ISwaggerSpec) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *ISwaggerSpec) SetEnable(v bool) {
	o.Enable = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ISwaggerSpec) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ISwaggerSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ISwaggerSpec) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ISwaggerSpec) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ISwaggerSpec) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ISwaggerSpec) UnsetName() {
	o.Name.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ISwaggerSpec) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ISwaggerSpec) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ISwaggerSpec) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ISwaggerSpec) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ISwaggerSpec) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ISwaggerSpec) UnsetTitle() {
	o.Title.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ISwaggerSpec) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ISwaggerSpec) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *ISwaggerSpec) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *ISwaggerSpec) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *ISwaggerSpec) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *ISwaggerSpec) UnsetVersion() {
	o.Version.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ISwaggerSpec) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ISwaggerSpec) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ISwaggerSpec) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ISwaggerSpec) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ISwaggerSpec) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ISwaggerSpec) UnsetDescription() {
	o.Description.Unset()
}

// GetTermsOfService returns the TermsOfService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ISwaggerSpec) GetTermsOfService() string {
	if o == nil || IsNil(o.TermsOfService.Get()) {
		var ret string
		return ret
	}
	return *o.TermsOfService.Get()
}

// GetTermsOfServiceOk returns a tuple with the TermsOfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ISwaggerSpec) GetTermsOfServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TermsOfService.Get(), o.TermsOfService.IsSet()
}

// HasTermsOfService returns a boolean if a field has been set.
func (o *ISwaggerSpec) HasTermsOfService() bool {
	if o != nil && o.TermsOfService.IsSet() {
		return true
	}

	return false
}

// SetTermsOfService gets a reference to the given NullableString and assigns it to the TermsOfService field.
func (o *ISwaggerSpec) SetTermsOfService(v string) {
	o.TermsOfService.Set(&v)
}
// SetTermsOfServiceNil sets the value for TermsOfService to be an explicit nil
func (o *ISwaggerSpec) SetTermsOfServiceNil() {
	o.TermsOfService.Set(nil)
}

// UnsetTermsOfService ensures that no value is present for TermsOfService, not even an explicit nil
func (o *ISwaggerSpec) UnsetTermsOfService() {
	o.TermsOfService.Unset()
}

// GetSwaggerEndpoint returns the SwaggerEndpoint field value if set, zero value otherwise.
func (o *ISwaggerSpec) GetSwaggerEndpoint() ISwaggerEndpoint {
	if o == nil || IsNil(o.SwaggerEndpoint) {
		var ret ISwaggerEndpoint
		return ret
	}
	return *o.SwaggerEndpoint
}

// GetSwaggerEndpointOk returns a tuple with the SwaggerEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISwaggerSpec) GetSwaggerEndpointOk() (*ISwaggerEndpoint, bool) {
	if o == nil || IsNil(o.SwaggerEndpoint) {
		return nil, false
	}
	return o.SwaggerEndpoint, true
}

// HasSwaggerEndpoint returns a boolean if a field has been set.
func (o *ISwaggerSpec) HasSwaggerEndpoint() bool {
	if o != nil && !IsNil(o.SwaggerEndpoint) {
		return true
	}

	return false
}

// SetSwaggerEndpoint gets a reference to the given ISwaggerEndpoint and assigns it to the SwaggerEndpoint field.
func (o *ISwaggerSpec) SetSwaggerEndpoint(v ISwaggerEndpoint) {
	o.SwaggerEndpoint = &v
}

// GetOpenApiContact returns the OpenApiContact field value if set, zero value otherwise.
func (o *ISwaggerSpec) GetOpenApiContact() ISwaggerContact {
	if o == nil || IsNil(o.OpenApiContact) {
		var ret ISwaggerContact
		return ret
	}
	return *o.OpenApiContact
}

// GetOpenApiContactOk returns a tuple with the OpenApiContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISwaggerSpec) GetOpenApiContactOk() (*ISwaggerContact, bool) {
	if o == nil || IsNil(o.OpenApiContact) {
		return nil, false
	}
	return o.OpenApiContact, true
}

// HasOpenApiContact returns a boolean if a field has been set.
func (o *ISwaggerSpec) HasOpenApiContact() bool {
	if o != nil && !IsNil(o.OpenApiContact) {
		return true
	}

	return false
}

// SetOpenApiContact gets a reference to the given ISwaggerContact and assigns it to the OpenApiContact field.
func (o *ISwaggerSpec) SetOpenApiContact(v ISwaggerContact) {
	o.OpenApiContact = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *ISwaggerSpec) GetLicense() ISwaggerLicense {
	if o == nil || IsNil(o.License) {
		var ret ISwaggerLicense
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISwaggerSpec) GetLicenseOk() (*ISwaggerLicense, bool) {
	if o == nil || IsNil(o.License) {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *ISwaggerSpec) HasLicense() bool {
	if o != nil && !IsNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given ISwaggerLicense and assigns it to the License field.
func (o *ISwaggerSpec) SetLicense(v ISwaggerLicense) {
	o.License = &v
}

func (o ISwaggerSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ISwaggerSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.TermsOfService.IsSet() {
		toSerialize["termsOfService"] = o.TermsOfService.Get()
	}
	if !IsNil(o.SwaggerEndpoint) {
		toSerialize["swaggerEndpoint"] = o.SwaggerEndpoint
	}
	if !IsNil(o.OpenApiContact) {
		toSerialize["openApiContact"] = o.OpenApiContact
	}
	if !IsNil(o.License) {
		toSerialize["license"] = o.License
	}
	return toSerialize, nil
}

type NullableISwaggerSpec struct {
	value *ISwaggerSpec
	isSet bool
}

func (v NullableISwaggerSpec) Get() *ISwaggerSpec {
	return v.value
}

func (v *NullableISwaggerSpec) Set(val *ISwaggerSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableISwaggerSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableISwaggerSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableISwaggerSpec(val *ISwaggerSpec) *NullableISwaggerSpec {
	return &NullableISwaggerSpec{value: val, isSet: true}
}

func (v NullableISwaggerSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableISwaggerSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


