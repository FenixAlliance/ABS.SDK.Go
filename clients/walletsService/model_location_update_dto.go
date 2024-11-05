/*
WalletsService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.1.4089
Contact: support@fenix-alliance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the LocationUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationUpdateDto{}

// LocationUpdateDto struct for LocationUpdateDto
type LocationUpdateDto struct {
	Title NullableString `json:"title,omitempty"`
	Email NullableString `json:"email,omitempty"`
	Phone NullableString `json:"phone,omitempty"`
	Fax NullableString `json:"fax,omitempty"`
	Address1 NullableString `json:"address1,omitempty"`
	Address2 NullableString `json:"address2,omitempty"`
	Address3 NullableString `json:"address3,omitempty"`
	Unit NullableString `json:"unit,omitempty"`
	CityId NullableString `json:"cityId,omitempty"`
	StateId NullableString `json:"stateId,omitempty"`
	PostalCode NullableString `json:"postalCode,omitempty"`
	CountryId NullableString `json:"countryId,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	IsRoutable *bool `json:"isRoutable,omitempty"`
	IsGlobalPrimary *bool `json:"isGlobalPrimary,omitempty"`
	IsCountryPrimary *bool `json:"isCountryPrimary,omitempty"`
	CanGenerateLabels *bool `json:"canGenerateLabels,omitempty"`
	IsDefaultSenderAddress *bool `json:"isDefaultSenderAddress,omitempty"`
	IsDefaultReturnAddress *bool `json:"isDefaultReturnAddress,omitempty"`
	IsDefaultSuppingLocation *bool `json:"isDefaultSuppingLocation,omitempty"`
}

// NewLocationUpdateDto instantiates a new LocationUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationUpdateDto() *LocationUpdateDto {
	this := LocationUpdateDto{}
	return &this
}

// NewLocationUpdateDtoWithDefaults instantiates a new LocationUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationUpdateDtoWithDefaults() *LocationUpdateDto {
	this := LocationUpdateDto{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationUpdateDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationUpdateDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *LocationUpdateDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *LocationUpdateDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *LocationUpdateDto) UnsetTitle() {
	o.Title.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationUpdateDto) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationUpdateDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *LocationUpdateDto) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *LocationUpdateDto) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *LocationUpdateDto) UnsetEmail() {
	o.Email.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationUpdateDto) GetPhone() string {
	if o == nil || IsNil(o.Phone.Get()) {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationUpdateDto) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *LocationUpdateDto) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *LocationUpdateDto) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *LocationUpdateDto) UnsetPhone() {
	o.Phone.Unset()
}

// GetFax returns the Fax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationUpdateDto) GetFax() string {
	if o == nil || IsNil(o.Fax.Get()) {
		var ret string
		return ret
	}
	return *o.Fax.Get()
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationUpdateDto) GetFaxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fax.Get(), o.Fax.IsSet()
}

// HasFax returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasFax() bool {
	if o != nil && o.Fax.IsSet() {
		return true
	}

	return false
}

// SetFax gets a reference to the given NullableString and assigns it to the Fax field.
func (o *LocationUpdateDto) SetFax(v string) {
	o.Fax.Set(&v)
}
// SetFaxNil sets the value for Fax to be an explicit nil
func (o *LocationUpdateDto) SetFaxNil() {
	o.Fax.Set(nil)
}

// UnsetFax ensures that no value is present for Fax, not even an explicit nil
func (o *LocationUpdateDto) UnsetFax() {
	o.Fax.Unset()
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationUpdateDto) GetAddress1() string {
	if o == nil || IsNil(o.Address1.Get()) {
		var ret string
		return ret
	}
	return *o.Address1.Get()
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationUpdateDto) GetAddress1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address1.Get(), o.Address1.IsSet()
}

// HasAddress1 returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasAddress1() bool {
	if o != nil && o.Address1.IsSet() {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given NullableString and assigns it to the Address1 field.
func (o *LocationUpdateDto) SetAddress1(v string) {
	o.Address1.Set(&v)
}
// SetAddress1Nil sets the value for Address1 to be an explicit nil
func (o *LocationUpdateDto) SetAddress1Nil() {
	o.Address1.Set(nil)
}

// UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
func (o *LocationUpdateDto) UnsetAddress1() {
	o.Address1.Unset()
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationUpdateDto) GetAddress2() string {
	if o == nil || IsNil(o.Address2.Get()) {
		var ret string
		return ret
	}
	return *o.Address2.Get()
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationUpdateDto) GetAddress2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address2.Get(), o.Address2.IsSet()
}

// HasAddress2 returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasAddress2() bool {
	if o != nil && o.Address2.IsSet() {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given NullableString and assigns it to the Address2 field.
func (o *LocationUpdateDto) SetAddress2(v string) {
	o.Address2.Set(&v)
}
// SetAddress2Nil sets the value for Address2 to be an explicit nil
func (o *LocationUpdateDto) SetAddress2Nil() {
	o.Address2.Set(nil)
}

// UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
func (o *LocationUpdateDto) UnsetAddress2() {
	o.Address2.Unset()
}

// GetAddress3 returns the Address3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationUpdateDto) GetAddress3() string {
	if o == nil || IsNil(o.Address3.Get()) {
		var ret string
		return ret
	}
	return *o.Address3.Get()
}

// GetAddress3Ok returns a tuple with the Address3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationUpdateDto) GetAddress3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address3.Get(), o.Address3.IsSet()
}

// HasAddress3 returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasAddress3() bool {
	if o != nil && o.Address3.IsSet() {
		return true
	}

	return false
}

// SetAddress3 gets a reference to the given NullableString and assigns it to the Address3 field.
func (o *LocationUpdateDto) SetAddress3(v string) {
	o.Address3.Set(&v)
}
// SetAddress3Nil sets the value for Address3 to be an explicit nil
func (o *LocationUpdateDto) SetAddress3Nil() {
	o.Address3.Set(nil)
}

// UnsetAddress3 ensures that no value is present for Address3, not even an explicit nil
func (o *LocationUpdateDto) UnsetAddress3() {
	o.Address3.Unset()
}

// GetUnit returns the Unit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationUpdateDto) GetUnit() string {
	if o == nil || IsNil(o.Unit.Get()) {
		var ret string
		return ret
	}
	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationUpdateDto) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// HasUnit returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasUnit() bool {
	if o != nil && o.Unit.IsSet() {
		return true
	}

	return false
}

// SetUnit gets a reference to the given NullableString and assigns it to the Unit field.
func (o *LocationUpdateDto) SetUnit(v string) {
	o.Unit.Set(&v)
}
// SetUnitNil sets the value for Unit to be an explicit nil
func (o *LocationUpdateDto) SetUnitNil() {
	o.Unit.Set(nil)
}

// UnsetUnit ensures that no value is present for Unit, not even an explicit nil
func (o *LocationUpdateDto) UnsetUnit() {
	o.Unit.Unset()
}

// GetCityId returns the CityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationUpdateDto) GetCityId() string {
	if o == nil || IsNil(o.CityId.Get()) {
		var ret string
		return ret
	}
	return *o.CityId.Get()
}

// GetCityIdOk returns a tuple with the CityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationUpdateDto) GetCityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CityId.Get(), o.CityId.IsSet()
}

// HasCityId returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasCityId() bool {
	if o != nil && o.CityId.IsSet() {
		return true
	}

	return false
}

// SetCityId gets a reference to the given NullableString and assigns it to the CityId field.
func (o *LocationUpdateDto) SetCityId(v string) {
	o.CityId.Set(&v)
}
// SetCityIdNil sets the value for CityId to be an explicit nil
func (o *LocationUpdateDto) SetCityIdNil() {
	o.CityId.Set(nil)
}

// UnsetCityId ensures that no value is present for CityId, not even an explicit nil
func (o *LocationUpdateDto) UnsetCityId() {
	o.CityId.Unset()
}

// GetStateId returns the StateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationUpdateDto) GetStateId() string {
	if o == nil || IsNil(o.StateId.Get()) {
		var ret string
		return ret
	}
	return *o.StateId.Get()
}

// GetStateIdOk returns a tuple with the StateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationUpdateDto) GetStateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StateId.Get(), o.StateId.IsSet()
}

// HasStateId returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasStateId() bool {
	if o != nil && o.StateId.IsSet() {
		return true
	}

	return false
}

// SetStateId gets a reference to the given NullableString and assigns it to the StateId field.
func (o *LocationUpdateDto) SetStateId(v string) {
	o.StateId.Set(&v)
}
// SetStateIdNil sets the value for StateId to be an explicit nil
func (o *LocationUpdateDto) SetStateIdNil() {
	o.StateId.Set(nil)
}

// UnsetStateId ensures that no value is present for StateId, not even an explicit nil
func (o *LocationUpdateDto) UnsetStateId() {
	o.StateId.Unset()
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationUpdateDto) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode.Get()) {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationUpdateDto) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *LocationUpdateDto) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *LocationUpdateDto) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *LocationUpdateDto) UnsetPostalCode() {
	o.PostalCode.Unset()
}

// GetCountryId returns the CountryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationUpdateDto) GetCountryId() string {
	if o == nil || IsNil(o.CountryId.Get()) {
		var ret string
		return ret
	}
	return *o.CountryId.Get()
}

// GetCountryIdOk returns a tuple with the CountryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationUpdateDto) GetCountryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryId.Get(), o.CountryId.IsSet()
}

// HasCountryId returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasCountryId() bool {
	if o != nil && o.CountryId.IsSet() {
		return true
	}

	return false
}

// SetCountryId gets a reference to the given NullableString and assigns it to the CountryId field.
func (o *LocationUpdateDto) SetCountryId(v string) {
	o.CountryId.Set(&v)
}
// SetCountryIdNil sets the value for CountryId to be an explicit nil
func (o *LocationUpdateDto) SetCountryIdNil() {
	o.CountryId.Set(nil)
}

// UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
func (o *LocationUpdateDto) UnsetCountryId() {
	o.CountryId.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationUpdateDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationUpdateDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *LocationUpdateDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *LocationUpdateDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *LocationUpdateDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *LocationUpdateDto) GetLongitude() float64 {
	if o == nil || IsNil(o.Longitude) {
		var ret float64
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationUpdateDto) GetLongitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float64 and assigns it to the Longitude field.
func (o *LocationUpdateDto) SetLongitude(v float64) {
	o.Longitude = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *LocationUpdateDto) GetLatitude() float64 {
	if o == nil || IsNil(o.Latitude) {
		var ret float64
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationUpdateDto) GetLatitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float64 and assigns it to the Latitude field.
func (o *LocationUpdateDto) SetLatitude(v float64) {
	o.Latitude = &v
}

// GetIsRoutable returns the IsRoutable field value if set, zero value otherwise.
func (o *LocationUpdateDto) GetIsRoutable() bool {
	if o == nil || IsNil(o.IsRoutable) {
		var ret bool
		return ret
	}
	return *o.IsRoutable
}

// GetIsRoutableOk returns a tuple with the IsRoutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationUpdateDto) GetIsRoutableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRoutable) {
		return nil, false
	}
	return o.IsRoutable, true
}

// HasIsRoutable returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasIsRoutable() bool {
	if o != nil && !IsNil(o.IsRoutable) {
		return true
	}

	return false
}

// SetIsRoutable gets a reference to the given bool and assigns it to the IsRoutable field.
func (o *LocationUpdateDto) SetIsRoutable(v bool) {
	o.IsRoutable = &v
}

// GetIsGlobalPrimary returns the IsGlobalPrimary field value if set, zero value otherwise.
func (o *LocationUpdateDto) GetIsGlobalPrimary() bool {
	if o == nil || IsNil(o.IsGlobalPrimary) {
		var ret bool
		return ret
	}
	return *o.IsGlobalPrimary
}

// GetIsGlobalPrimaryOk returns a tuple with the IsGlobalPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationUpdateDto) GetIsGlobalPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGlobalPrimary) {
		return nil, false
	}
	return o.IsGlobalPrimary, true
}

// HasIsGlobalPrimary returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasIsGlobalPrimary() bool {
	if o != nil && !IsNil(o.IsGlobalPrimary) {
		return true
	}

	return false
}

// SetIsGlobalPrimary gets a reference to the given bool and assigns it to the IsGlobalPrimary field.
func (o *LocationUpdateDto) SetIsGlobalPrimary(v bool) {
	o.IsGlobalPrimary = &v
}

// GetIsCountryPrimary returns the IsCountryPrimary field value if set, zero value otherwise.
func (o *LocationUpdateDto) GetIsCountryPrimary() bool {
	if o == nil || IsNil(o.IsCountryPrimary) {
		var ret bool
		return ret
	}
	return *o.IsCountryPrimary
}

// GetIsCountryPrimaryOk returns a tuple with the IsCountryPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationUpdateDto) GetIsCountryPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCountryPrimary) {
		return nil, false
	}
	return o.IsCountryPrimary, true
}

// HasIsCountryPrimary returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasIsCountryPrimary() bool {
	if o != nil && !IsNil(o.IsCountryPrimary) {
		return true
	}

	return false
}

// SetIsCountryPrimary gets a reference to the given bool and assigns it to the IsCountryPrimary field.
func (o *LocationUpdateDto) SetIsCountryPrimary(v bool) {
	o.IsCountryPrimary = &v
}

// GetCanGenerateLabels returns the CanGenerateLabels field value if set, zero value otherwise.
func (o *LocationUpdateDto) GetCanGenerateLabels() bool {
	if o == nil || IsNil(o.CanGenerateLabels) {
		var ret bool
		return ret
	}
	return *o.CanGenerateLabels
}

// GetCanGenerateLabelsOk returns a tuple with the CanGenerateLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationUpdateDto) GetCanGenerateLabelsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanGenerateLabels) {
		return nil, false
	}
	return o.CanGenerateLabels, true
}

// HasCanGenerateLabels returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasCanGenerateLabels() bool {
	if o != nil && !IsNil(o.CanGenerateLabels) {
		return true
	}

	return false
}

// SetCanGenerateLabels gets a reference to the given bool and assigns it to the CanGenerateLabels field.
func (o *LocationUpdateDto) SetCanGenerateLabels(v bool) {
	o.CanGenerateLabels = &v
}

// GetIsDefaultSenderAddress returns the IsDefaultSenderAddress field value if set, zero value otherwise.
func (o *LocationUpdateDto) GetIsDefaultSenderAddress() bool {
	if o == nil || IsNil(o.IsDefaultSenderAddress) {
		var ret bool
		return ret
	}
	return *o.IsDefaultSenderAddress
}

// GetIsDefaultSenderAddressOk returns a tuple with the IsDefaultSenderAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationUpdateDto) GetIsDefaultSenderAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultSenderAddress) {
		return nil, false
	}
	return o.IsDefaultSenderAddress, true
}

// HasIsDefaultSenderAddress returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasIsDefaultSenderAddress() bool {
	if o != nil && !IsNil(o.IsDefaultSenderAddress) {
		return true
	}

	return false
}

// SetIsDefaultSenderAddress gets a reference to the given bool and assigns it to the IsDefaultSenderAddress field.
func (o *LocationUpdateDto) SetIsDefaultSenderAddress(v bool) {
	o.IsDefaultSenderAddress = &v
}

// GetIsDefaultReturnAddress returns the IsDefaultReturnAddress field value if set, zero value otherwise.
func (o *LocationUpdateDto) GetIsDefaultReturnAddress() bool {
	if o == nil || IsNil(o.IsDefaultReturnAddress) {
		var ret bool
		return ret
	}
	return *o.IsDefaultReturnAddress
}

// GetIsDefaultReturnAddressOk returns a tuple with the IsDefaultReturnAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationUpdateDto) GetIsDefaultReturnAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultReturnAddress) {
		return nil, false
	}
	return o.IsDefaultReturnAddress, true
}

// HasIsDefaultReturnAddress returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasIsDefaultReturnAddress() bool {
	if o != nil && !IsNil(o.IsDefaultReturnAddress) {
		return true
	}

	return false
}

// SetIsDefaultReturnAddress gets a reference to the given bool and assigns it to the IsDefaultReturnAddress field.
func (o *LocationUpdateDto) SetIsDefaultReturnAddress(v bool) {
	o.IsDefaultReturnAddress = &v
}

// GetIsDefaultSuppingLocation returns the IsDefaultSuppingLocation field value if set, zero value otherwise.
func (o *LocationUpdateDto) GetIsDefaultSuppingLocation() bool {
	if o == nil || IsNil(o.IsDefaultSuppingLocation) {
		var ret bool
		return ret
	}
	return *o.IsDefaultSuppingLocation
}

// GetIsDefaultSuppingLocationOk returns a tuple with the IsDefaultSuppingLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationUpdateDto) GetIsDefaultSuppingLocationOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultSuppingLocation) {
		return nil, false
	}
	return o.IsDefaultSuppingLocation, true
}

// HasIsDefaultSuppingLocation returns a boolean if a field has been set.
func (o *LocationUpdateDto) HasIsDefaultSuppingLocation() bool {
	if o != nil && !IsNil(o.IsDefaultSuppingLocation) {
		return true
	}

	return false
}

// SetIsDefaultSuppingLocation gets a reference to the given bool and assigns it to the IsDefaultSuppingLocation field.
func (o *LocationUpdateDto) SetIsDefaultSuppingLocation(v bool) {
	o.IsDefaultSuppingLocation = &v
}

func (o LocationUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if o.Fax.IsSet() {
		toSerialize["fax"] = o.Fax.Get()
	}
	if o.Address1.IsSet() {
		toSerialize["address1"] = o.Address1.Get()
	}
	if o.Address2.IsSet() {
		toSerialize["address2"] = o.Address2.Get()
	}
	if o.Address3.IsSet() {
		toSerialize["address3"] = o.Address3.Get()
	}
	if o.Unit.IsSet() {
		toSerialize["unit"] = o.Unit.Get()
	}
	if o.CityId.IsSet() {
		toSerialize["cityId"] = o.CityId.Get()
	}
	if o.StateId.IsSet() {
		toSerialize["stateId"] = o.StateId.Get()
	}
	if o.PostalCode.IsSet() {
		toSerialize["postalCode"] = o.PostalCode.Get()
	}
	if o.CountryId.IsSet() {
		toSerialize["countryId"] = o.CountryId.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.IsRoutable) {
		toSerialize["isRoutable"] = o.IsRoutable
	}
	if !IsNil(o.IsGlobalPrimary) {
		toSerialize["isGlobalPrimary"] = o.IsGlobalPrimary
	}
	if !IsNil(o.IsCountryPrimary) {
		toSerialize["isCountryPrimary"] = o.IsCountryPrimary
	}
	if !IsNil(o.CanGenerateLabels) {
		toSerialize["canGenerateLabels"] = o.CanGenerateLabels
	}
	if !IsNil(o.IsDefaultSenderAddress) {
		toSerialize["isDefaultSenderAddress"] = o.IsDefaultSenderAddress
	}
	if !IsNil(o.IsDefaultReturnAddress) {
		toSerialize["isDefaultReturnAddress"] = o.IsDefaultReturnAddress
	}
	if !IsNil(o.IsDefaultSuppingLocation) {
		toSerialize["isDefaultSuppingLocation"] = o.IsDefaultSuppingLocation
	}
	return toSerialize, nil
}

type NullableLocationUpdateDto struct {
	value *LocationUpdateDto
	isSet bool
}

func (v NullableLocationUpdateDto) Get() *LocationUpdateDto {
	return v.value
}

func (v *NullableLocationUpdateDto) Set(val *LocationUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationUpdateDto(val *LocationUpdateDto) *NullableLocationUpdateDto {
	return &NullableLocationUpdateDto{value: val, isSet: true}
}

func (v NullableLocationUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


