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
	"time"
)

// checks if the LocationCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationCreateDto{}

// LocationCreateDto struct for LocationCreateDto
type LocationCreateDto struct {
	Id *string `json:"id,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
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

// NewLocationCreateDto instantiates a new LocationCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationCreateDto() *LocationCreateDto {
	this := LocationCreateDto{}
	return &this
}

// NewLocationCreateDtoWithDefaults instantiates a new LocationCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationCreateDtoWithDefaults() *LocationCreateDto {
	this := LocationCreateDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LocationCreateDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationCreateDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LocationCreateDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LocationCreateDto) SetId(v string) {
	o.Id = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *LocationCreateDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationCreateDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *LocationCreateDto) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *LocationCreateDto) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationCreateDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationCreateDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *LocationCreateDto) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *LocationCreateDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *LocationCreateDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *LocationCreateDto) UnsetTitle() {
	o.Title.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationCreateDto) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationCreateDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *LocationCreateDto) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *LocationCreateDto) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *LocationCreateDto) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *LocationCreateDto) UnsetEmail() {
	o.Email.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationCreateDto) GetPhone() string {
	if o == nil || IsNil(o.Phone.Get()) {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationCreateDto) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *LocationCreateDto) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *LocationCreateDto) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *LocationCreateDto) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *LocationCreateDto) UnsetPhone() {
	o.Phone.Unset()
}

// GetFax returns the Fax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationCreateDto) GetFax() string {
	if o == nil || IsNil(o.Fax.Get()) {
		var ret string
		return ret
	}
	return *o.Fax.Get()
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationCreateDto) GetFaxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fax.Get(), o.Fax.IsSet()
}

// HasFax returns a boolean if a field has been set.
func (o *LocationCreateDto) HasFax() bool {
	if o != nil && o.Fax.IsSet() {
		return true
	}

	return false
}

// SetFax gets a reference to the given NullableString and assigns it to the Fax field.
func (o *LocationCreateDto) SetFax(v string) {
	o.Fax.Set(&v)
}
// SetFaxNil sets the value for Fax to be an explicit nil
func (o *LocationCreateDto) SetFaxNil() {
	o.Fax.Set(nil)
}

// UnsetFax ensures that no value is present for Fax, not even an explicit nil
func (o *LocationCreateDto) UnsetFax() {
	o.Fax.Unset()
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationCreateDto) GetAddress1() string {
	if o == nil || IsNil(o.Address1.Get()) {
		var ret string
		return ret
	}
	return *o.Address1.Get()
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationCreateDto) GetAddress1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address1.Get(), o.Address1.IsSet()
}

// HasAddress1 returns a boolean if a field has been set.
func (o *LocationCreateDto) HasAddress1() bool {
	if o != nil && o.Address1.IsSet() {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given NullableString and assigns it to the Address1 field.
func (o *LocationCreateDto) SetAddress1(v string) {
	o.Address1.Set(&v)
}
// SetAddress1Nil sets the value for Address1 to be an explicit nil
func (o *LocationCreateDto) SetAddress1Nil() {
	o.Address1.Set(nil)
}

// UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
func (o *LocationCreateDto) UnsetAddress1() {
	o.Address1.Unset()
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationCreateDto) GetAddress2() string {
	if o == nil || IsNil(o.Address2.Get()) {
		var ret string
		return ret
	}
	return *o.Address2.Get()
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationCreateDto) GetAddress2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address2.Get(), o.Address2.IsSet()
}

// HasAddress2 returns a boolean if a field has been set.
func (o *LocationCreateDto) HasAddress2() bool {
	if o != nil && o.Address2.IsSet() {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given NullableString and assigns it to the Address2 field.
func (o *LocationCreateDto) SetAddress2(v string) {
	o.Address2.Set(&v)
}
// SetAddress2Nil sets the value for Address2 to be an explicit nil
func (o *LocationCreateDto) SetAddress2Nil() {
	o.Address2.Set(nil)
}

// UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
func (o *LocationCreateDto) UnsetAddress2() {
	o.Address2.Unset()
}

// GetAddress3 returns the Address3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationCreateDto) GetAddress3() string {
	if o == nil || IsNil(o.Address3.Get()) {
		var ret string
		return ret
	}
	return *o.Address3.Get()
}

// GetAddress3Ok returns a tuple with the Address3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationCreateDto) GetAddress3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address3.Get(), o.Address3.IsSet()
}

// HasAddress3 returns a boolean if a field has been set.
func (o *LocationCreateDto) HasAddress3() bool {
	if o != nil && o.Address3.IsSet() {
		return true
	}

	return false
}

// SetAddress3 gets a reference to the given NullableString and assigns it to the Address3 field.
func (o *LocationCreateDto) SetAddress3(v string) {
	o.Address3.Set(&v)
}
// SetAddress3Nil sets the value for Address3 to be an explicit nil
func (o *LocationCreateDto) SetAddress3Nil() {
	o.Address3.Set(nil)
}

// UnsetAddress3 ensures that no value is present for Address3, not even an explicit nil
func (o *LocationCreateDto) UnsetAddress3() {
	o.Address3.Unset()
}

// GetUnit returns the Unit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationCreateDto) GetUnit() string {
	if o == nil || IsNil(o.Unit.Get()) {
		var ret string
		return ret
	}
	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationCreateDto) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// HasUnit returns a boolean if a field has been set.
func (o *LocationCreateDto) HasUnit() bool {
	if o != nil && o.Unit.IsSet() {
		return true
	}

	return false
}

// SetUnit gets a reference to the given NullableString and assigns it to the Unit field.
func (o *LocationCreateDto) SetUnit(v string) {
	o.Unit.Set(&v)
}
// SetUnitNil sets the value for Unit to be an explicit nil
func (o *LocationCreateDto) SetUnitNil() {
	o.Unit.Set(nil)
}

// UnsetUnit ensures that no value is present for Unit, not even an explicit nil
func (o *LocationCreateDto) UnsetUnit() {
	o.Unit.Unset()
}

// GetCityId returns the CityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationCreateDto) GetCityId() string {
	if o == nil || IsNil(o.CityId.Get()) {
		var ret string
		return ret
	}
	return *o.CityId.Get()
}

// GetCityIdOk returns a tuple with the CityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationCreateDto) GetCityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CityId.Get(), o.CityId.IsSet()
}

// HasCityId returns a boolean if a field has been set.
func (o *LocationCreateDto) HasCityId() bool {
	if o != nil && o.CityId.IsSet() {
		return true
	}

	return false
}

// SetCityId gets a reference to the given NullableString and assigns it to the CityId field.
func (o *LocationCreateDto) SetCityId(v string) {
	o.CityId.Set(&v)
}
// SetCityIdNil sets the value for CityId to be an explicit nil
func (o *LocationCreateDto) SetCityIdNil() {
	o.CityId.Set(nil)
}

// UnsetCityId ensures that no value is present for CityId, not even an explicit nil
func (o *LocationCreateDto) UnsetCityId() {
	o.CityId.Unset()
}

// GetStateId returns the StateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationCreateDto) GetStateId() string {
	if o == nil || IsNil(o.StateId.Get()) {
		var ret string
		return ret
	}
	return *o.StateId.Get()
}

// GetStateIdOk returns a tuple with the StateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationCreateDto) GetStateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StateId.Get(), o.StateId.IsSet()
}

// HasStateId returns a boolean if a field has been set.
func (o *LocationCreateDto) HasStateId() bool {
	if o != nil && o.StateId.IsSet() {
		return true
	}

	return false
}

// SetStateId gets a reference to the given NullableString and assigns it to the StateId field.
func (o *LocationCreateDto) SetStateId(v string) {
	o.StateId.Set(&v)
}
// SetStateIdNil sets the value for StateId to be an explicit nil
func (o *LocationCreateDto) SetStateIdNil() {
	o.StateId.Set(nil)
}

// UnsetStateId ensures that no value is present for StateId, not even an explicit nil
func (o *LocationCreateDto) UnsetStateId() {
	o.StateId.Unset()
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationCreateDto) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode.Get()) {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationCreateDto) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *LocationCreateDto) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *LocationCreateDto) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *LocationCreateDto) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *LocationCreateDto) UnsetPostalCode() {
	o.PostalCode.Unset()
}

// GetCountryId returns the CountryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationCreateDto) GetCountryId() string {
	if o == nil || IsNil(o.CountryId.Get()) {
		var ret string
		return ret
	}
	return *o.CountryId.Get()
}

// GetCountryIdOk returns a tuple with the CountryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationCreateDto) GetCountryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryId.Get(), o.CountryId.IsSet()
}

// HasCountryId returns a boolean if a field has been set.
func (o *LocationCreateDto) HasCountryId() bool {
	if o != nil && o.CountryId.IsSet() {
		return true
	}

	return false
}

// SetCountryId gets a reference to the given NullableString and assigns it to the CountryId field.
func (o *LocationCreateDto) SetCountryId(v string) {
	o.CountryId.Set(&v)
}
// SetCountryIdNil sets the value for CountryId to be an explicit nil
func (o *LocationCreateDto) SetCountryIdNil() {
	o.CountryId.Set(nil)
}

// UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
func (o *LocationCreateDto) UnsetCountryId() {
	o.CountryId.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationCreateDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationCreateDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *LocationCreateDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *LocationCreateDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *LocationCreateDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *LocationCreateDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *LocationCreateDto) GetLongitude() float64 {
	if o == nil || IsNil(o.Longitude) {
		var ret float64
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationCreateDto) GetLongitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *LocationCreateDto) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float64 and assigns it to the Longitude field.
func (o *LocationCreateDto) SetLongitude(v float64) {
	o.Longitude = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *LocationCreateDto) GetLatitude() float64 {
	if o == nil || IsNil(o.Latitude) {
		var ret float64
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationCreateDto) GetLatitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *LocationCreateDto) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float64 and assigns it to the Latitude field.
func (o *LocationCreateDto) SetLatitude(v float64) {
	o.Latitude = &v
}

// GetIsRoutable returns the IsRoutable field value if set, zero value otherwise.
func (o *LocationCreateDto) GetIsRoutable() bool {
	if o == nil || IsNil(o.IsRoutable) {
		var ret bool
		return ret
	}
	return *o.IsRoutable
}

// GetIsRoutableOk returns a tuple with the IsRoutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationCreateDto) GetIsRoutableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRoutable) {
		return nil, false
	}
	return o.IsRoutable, true
}

// HasIsRoutable returns a boolean if a field has been set.
func (o *LocationCreateDto) HasIsRoutable() bool {
	if o != nil && !IsNil(o.IsRoutable) {
		return true
	}

	return false
}

// SetIsRoutable gets a reference to the given bool and assigns it to the IsRoutable field.
func (o *LocationCreateDto) SetIsRoutable(v bool) {
	o.IsRoutable = &v
}

// GetIsGlobalPrimary returns the IsGlobalPrimary field value if set, zero value otherwise.
func (o *LocationCreateDto) GetIsGlobalPrimary() bool {
	if o == nil || IsNil(o.IsGlobalPrimary) {
		var ret bool
		return ret
	}
	return *o.IsGlobalPrimary
}

// GetIsGlobalPrimaryOk returns a tuple with the IsGlobalPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationCreateDto) GetIsGlobalPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGlobalPrimary) {
		return nil, false
	}
	return o.IsGlobalPrimary, true
}

// HasIsGlobalPrimary returns a boolean if a field has been set.
func (o *LocationCreateDto) HasIsGlobalPrimary() bool {
	if o != nil && !IsNil(o.IsGlobalPrimary) {
		return true
	}

	return false
}

// SetIsGlobalPrimary gets a reference to the given bool and assigns it to the IsGlobalPrimary field.
func (o *LocationCreateDto) SetIsGlobalPrimary(v bool) {
	o.IsGlobalPrimary = &v
}

// GetIsCountryPrimary returns the IsCountryPrimary field value if set, zero value otherwise.
func (o *LocationCreateDto) GetIsCountryPrimary() bool {
	if o == nil || IsNil(o.IsCountryPrimary) {
		var ret bool
		return ret
	}
	return *o.IsCountryPrimary
}

// GetIsCountryPrimaryOk returns a tuple with the IsCountryPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationCreateDto) GetIsCountryPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCountryPrimary) {
		return nil, false
	}
	return o.IsCountryPrimary, true
}

// HasIsCountryPrimary returns a boolean if a field has been set.
func (o *LocationCreateDto) HasIsCountryPrimary() bool {
	if o != nil && !IsNil(o.IsCountryPrimary) {
		return true
	}

	return false
}

// SetIsCountryPrimary gets a reference to the given bool and assigns it to the IsCountryPrimary field.
func (o *LocationCreateDto) SetIsCountryPrimary(v bool) {
	o.IsCountryPrimary = &v
}

// GetCanGenerateLabels returns the CanGenerateLabels field value if set, zero value otherwise.
func (o *LocationCreateDto) GetCanGenerateLabels() bool {
	if o == nil || IsNil(o.CanGenerateLabels) {
		var ret bool
		return ret
	}
	return *o.CanGenerateLabels
}

// GetCanGenerateLabelsOk returns a tuple with the CanGenerateLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationCreateDto) GetCanGenerateLabelsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanGenerateLabels) {
		return nil, false
	}
	return o.CanGenerateLabels, true
}

// HasCanGenerateLabels returns a boolean if a field has been set.
func (o *LocationCreateDto) HasCanGenerateLabels() bool {
	if o != nil && !IsNil(o.CanGenerateLabels) {
		return true
	}

	return false
}

// SetCanGenerateLabels gets a reference to the given bool and assigns it to the CanGenerateLabels field.
func (o *LocationCreateDto) SetCanGenerateLabels(v bool) {
	o.CanGenerateLabels = &v
}

// GetIsDefaultSenderAddress returns the IsDefaultSenderAddress field value if set, zero value otherwise.
func (o *LocationCreateDto) GetIsDefaultSenderAddress() bool {
	if o == nil || IsNil(o.IsDefaultSenderAddress) {
		var ret bool
		return ret
	}
	return *o.IsDefaultSenderAddress
}

// GetIsDefaultSenderAddressOk returns a tuple with the IsDefaultSenderAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationCreateDto) GetIsDefaultSenderAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultSenderAddress) {
		return nil, false
	}
	return o.IsDefaultSenderAddress, true
}

// HasIsDefaultSenderAddress returns a boolean if a field has been set.
func (o *LocationCreateDto) HasIsDefaultSenderAddress() bool {
	if o != nil && !IsNil(o.IsDefaultSenderAddress) {
		return true
	}

	return false
}

// SetIsDefaultSenderAddress gets a reference to the given bool and assigns it to the IsDefaultSenderAddress field.
func (o *LocationCreateDto) SetIsDefaultSenderAddress(v bool) {
	o.IsDefaultSenderAddress = &v
}

// GetIsDefaultReturnAddress returns the IsDefaultReturnAddress field value if set, zero value otherwise.
func (o *LocationCreateDto) GetIsDefaultReturnAddress() bool {
	if o == nil || IsNil(o.IsDefaultReturnAddress) {
		var ret bool
		return ret
	}
	return *o.IsDefaultReturnAddress
}

// GetIsDefaultReturnAddressOk returns a tuple with the IsDefaultReturnAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationCreateDto) GetIsDefaultReturnAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultReturnAddress) {
		return nil, false
	}
	return o.IsDefaultReturnAddress, true
}

// HasIsDefaultReturnAddress returns a boolean if a field has been set.
func (o *LocationCreateDto) HasIsDefaultReturnAddress() bool {
	if o != nil && !IsNil(o.IsDefaultReturnAddress) {
		return true
	}

	return false
}

// SetIsDefaultReturnAddress gets a reference to the given bool and assigns it to the IsDefaultReturnAddress field.
func (o *LocationCreateDto) SetIsDefaultReturnAddress(v bool) {
	o.IsDefaultReturnAddress = &v
}

// GetIsDefaultSuppingLocation returns the IsDefaultSuppingLocation field value if set, zero value otherwise.
func (o *LocationCreateDto) GetIsDefaultSuppingLocation() bool {
	if o == nil || IsNil(o.IsDefaultSuppingLocation) {
		var ret bool
		return ret
	}
	return *o.IsDefaultSuppingLocation
}

// GetIsDefaultSuppingLocationOk returns a tuple with the IsDefaultSuppingLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationCreateDto) GetIsDefaultSuppingLocationOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultSuppingLocation) {
		return nil, false
	}
	return o.IsDefaultSuppingLocation, true
}

// HasIsDefaultSuppingLocation returns a boolean if a field has been set.
func (o *LocationCreateDto) HasIsDefaultSuppingLocation() bool {
	if o != nil && !IsNil(o.IsDefaultSuppingLocation) {
		return true
	}

	return false
}

// SetIsDefaultSuppingLocation gets a reference to the given bool and assigns it to the IsDefaultSuppingLocation field.
func (o *LocationCreateDto) SetIsDefaultSuppingLocation(v bool) {
	o.IsDefaultSuppingLocation = &v
}

func (o LocationCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
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

type NullableLocationCreateDto struct {
	value *LocationCreateDto
	isSet bool
}

func (v NullableLocationCreateDto) Get() *LocationCreateDto {
	return v.value
}

func (v *NullableLocationCreateDto) Set(val *LocationCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationCreateDto(val *LocationCreateDto) *NullableLocationCreateDto {
	return &NullableLocationCreateDto{value: val, isSet: true}
}

func (v NullableLocationCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


