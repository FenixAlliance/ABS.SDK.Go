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
	"time"
)

// checks if the TenantCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantCreateDto{}

// TenantCreateDto struct for TenantCreateDto
type TenantCreateDto struct {
	Id *string `json:"id,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Duns NullableString `json:"duns,omitempty"`
	Name NullableString `json:"name,omitempty"`
	LegalName NullableString `json:"legalName,omitempty"`
	Email NullableString `json:"email,omitempty"`
	Phone NullableString `json:"phone,omitempty"`
	WebUrl NullableString `json:"webUrl,omitempty"`
	About NullableString `json:"about,omitempty"`
	Handler NullableString `json:"handler,omitempty"`
	CurrencyId NullableString `json:"currencyId,omitempty"`
	LanguageId NullableString `json:"languageId,omitempty"`
	TimezoneId NullableString `json:"timezoneId,omitempty"`
	CityId NullableString `json:"cityId,omitempty"`
	StateId NullableString `json:"stateId,omitempty"`
	CountryId NullableString `json:"countryId,omitempty"`
	TaxId NullableString `json:"taxId,omitempty"`
	AvatarUrl NullableString `json:"avatarUrl,omitempty"`
}

// NewTenantCreateDto instantiates a new TenantCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantCreateDto() *TenantCreateDto {
	this := TenantCreateDto{}
	return &this
}

// NewTenantCreateDtoWithDefaults instantiates a new TenantCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantCreateDtoWithDefaults() *TenantCreateDto {
	this := TenantCreateDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TenantCreateDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreateDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TenantCreateDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TenantCreateDto) SetId(v string) {
	o.Id = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TenantCreateDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreateDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TenantCreateDto) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *TenantCreateDto) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetDuns returns the Duns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantCreateDto) GetDuns() string {
	if o == nil || IsNil(o.Duns.Get()) {
		var ret string
		return ret
	}
	return *o.Duns.Get()
}

// GetDunsOk returns a tuple with the Duns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantCreateDto) GetDunsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duns.Get(), o.Duns.IsSet()
}

// HasDuns returns a boolean if a field has been set.
func (o *TenantCreateDto) HasDuns() bool {
	if o != nil && o.Duns.IsSet() {
		return true
	}

	return false
}

// SetDuns gets a reference to the given NullableString and assigns it to the Duns field.
func (o *TenantCreateDto) SetDuns(v string) {
	o.Duns.Set(&v)
}
// SetDunsNil sets the value for Duns to be an explicit nil
func (o *TenantCreateDto) SetDunsNil() {
	o.Duns.Set(nil)
}

// UnsetDuns ensures that no value is present for Duns, not even an explicit nil
func (o *TenantCreateDto) UnsetDuns() {
	o.Duns.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantCreateDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantCreateDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TenantCreateDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TenantCreateDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TenantCreateDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TenantCreateDto) UnsetName() {
	o.Name.Unset()
}

// GetLegalName returns the LegalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantCreateDto) GetLegalName() string {
	if o == nil || IsNil(o.LegalName.Get()) {
		var ret string
		return ret
	}
	return *o.LegalName.Get()
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantCreateDto) GetLegalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LegalName.Get(), o.LegalName.IsSet()
}

// HasLegalName returns a boolean if a field has been set.
func (o *TenantCreateDto) HasLegalName() bool {
	if o != nil && o.LegalName.IsSet() {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given NullableString and assigns it to the LegalName field.
func (o *TenantCreateDto) SetLegalName(v string) {
	o.LegalName.Set(&v)
}
// SetLegalNameNil sets the value for LegalName to be an explicit nil
func (o *TenantCreateDto) SetLegalNameNil() {
	o.LegalName.Set(nil)
}

// UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
func (o *TenantCreateDto) UnsetLegalName() {
	o.LegalName.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantCreateDto) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantCreateDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *TenantCreateDto) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *TenantCreateDto) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *TenantCreateDto) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *TenantCreateDto) UnsetEmail() {
	o.Email.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantCreateDto) GetPhone() string {
	if o == nil || IsNil(o.Phone.Get()) {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantCreateDto) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *TenantCreateDto) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *TenantCreateDto) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *TenantCreateDto) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *TenantCreateDto) UnsetPhone() {
	o.Phone.Unset()
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantCreateDto) GetWebUrl() string {
	if o == nil || IsNil(o.WebUrl.Get()) {
		var ret string
		return ret
	}
	return *o.WebUrl.Get()
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantCreateDto) GetWebUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebUrl.Get(), o.WebUrl.IsSet()
}

// HasWebUrl returns a boolean if a field has been set.
func (o *TenantCreateDto) HasWebUrl() bool {
	if o != nil && o.WebUrl.IsSet() {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given NullableString and assigns it to the WebUrl field.
func (o *TenantCreateDto) SetWebUrl(v string) {
	o.WebUrl.Set(&v)
}
// SetWebUrlNil sets the value for WebUrl to be an explicit nil
func (o *TenantCreateDto) SetWebUrlNil() {
	o.WebUrl.Set(nil)
}

// UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
func (o *TenantCreateDto) UnsetWebUrl() {
	o.WebUrl.Unset()
}

// GetAbout returns the About field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantCreateDto) GetAbout() string {
	if o == nil || IsNil(o.About.Get()) {
		var ret string
		return ret
	}
	return *o.About.Get()
}

// GetAboutOk returns a tuple with the About field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantCreateDto) GetAboutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.About.Get(), o.About.IsSet()
}

// HasAbout returns a boolean if a field has been set.
func (o *TenantCreateDto) HasAbout() bool {
	if o != nil && o.About.IsSet() {
		return true
	}

	return false
}

// SetAbout gets a reference to the given NullableString and assigns it to the About field.
func (o *TenantCreateDto) SetAbout(v string) {
	o.About.Set(&v)
}
// SetAboutNil sets the value for About to be an explicit nil
func (o *TenantCreateDto) SetAboutNil() {
	o.About.Set(nil)
}

// UnsetAbout ensures that no value is present for About, not even an explicit nil
func (o *TenantCreateDto) UnsetAbout() {
	o.About.Unset()
}

// GetHandler returns the Handler field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantCreateDto) GetHandler() string {
	if o == nil || IsNil(o.Handler.Get()) {
		var ret string
		return ret
	}
	return *o.Handler.Get()
}

// GetHandlerOk returns a tuple with the Handler field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantCreateDto) GetHandlerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Handler.Get(), o.Handler.IsSet()
}

// HasHandler returns a boolean if a field has been set.
func (o *TenantCreateDto) HasHandler() bool {
	if o != nil && o.Handler.IsSet() {
		return true
	}

	return false
}

// SetHandler gets a reference to the given NullableString and assigns it to the Handler field.
func (o *TenantCreateDto) SetHandler(v string) {
	o.Handler.Set(&v)
}
// SetHandlerNil sets the value for Handler to be an explicit nil
func (o *TenantCreateDto) SetHandlerNil() {
	o.Handler.Set(nil)
}

// UnsetHandler ensures that no value is present for Handler, not even an explicit nil
func (o *TenantCreateDto) UnsetHandler() {
	o.Handler.Unset()
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantCreateDto) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId.Get()) {
		var ret string
		return ret
	}
	return *o.CurrencyId.Get()
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantCreateDto) GetCurrencyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrencyId.Get(), o.CurrencyId.IsSet()
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *TenantCreateDto) HasCurrencyId() bool {
	if o != nil && o.CurrencyId.IsSet() {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given NullableString and assigns it to the CurrencyId field.
func (o *TenantCreateDto) SetCurrencyId(v string) {
	o.CurrencyId.Set(&v)
}
// SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil
func (o *TenantCreateDto) SetCurrencyIdNil() {
	o.CurrencyId.Set(nil)
}

// UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
func (o *TenantCreateDto) UnsetCurrencyId() {
	o.CurrencyId.Unset()
}

// GetLanguageId returns the LanguageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantCreateDto) GetLanguageId() string {
	if o == nil || IsNil(o.LanguageId.Get()) {
		var ret string
		return ret
	}
	return *o.LanguageId.Get()
}

// GetLanguageIdOk returns a tuple with the LanguageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantCreateDto) GetLanguageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LanguageId.Get(), o.LanguageId.IsSet()
}

// HasLanguageId returns a boolean if a field has been set.
func (o *TenantCreateDto) HasLanguageId() bool {
	if o != nil && o.LanguageId.IsSet() {
		return true
	}

	return false
}

// SetLanguageId gets a reference to the given NullableString and assigns it to the LanguageId field.
func (o *TenantCreateDto) SetLanguageId(v string) {
	o.LanguageId.Set(&v)
}
// SetLanguageIdNil sets the value for LanguageId to be an explicit nil
func (o *TenantCreateDto) SetLanguageIdNil() {
	o.LanguageId.Set(nil)
}

// UnsetLanguageId ensures that no value is present for LanguageId, not even an explicit nil
func (o *TenantCreateDto) UnsetLanguageId() {
	o.LanguageId.Unset()
}

// GetTimezoneId returns the TimezoneId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantCreateDto) GetTimezoneId() string {
	if o == nil || IsNil(o.TimezoneId.Get()) {
		var ret string
		return ret
	}
	return *o.TimezoneId.Get()
}

// GetTimezoneIdOk returns a tuple with the TimezoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantCreateDto) GetTimezoneIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimezoneId.Get(), o.TimezoneId.IsSet()
}

// HasTimezoneId returns a boolean if a field has been set.
func (o *TenantCreateDto) HasTimezoneId() bool {
	if o != nil && o.TimezoneId.IsSet() {
		return true
	}

	return false
}

// SetTimezoneId gets a reference to the given NullableString and assigns it to the TimezoneId field.
func (o *TenantCreateDto) SetTimezoneId(v string) {
	o.TimezoneId.Set(&v)
}
// SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil
func (o *TenantCreateDto) SetTimezoneIdNil() {
	o.TimezoneId.Set(nil)
}

// UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
func (o *TenantCreateDto) UnsetTimezoneId() {
	o.TimezoneId.Unset()
}

// GetCityId returns the CityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantCreateDto) GetCityId() string {
	if o == nil || IsNil(o.CityId.Get()) {
		var ret string
		return ret
	}
	return *o.CityId.Get()
}

// GetCityIdOk returns a tuple with the CityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantCreateDto) GetCityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CityId.Get(), o.CityId.IsSet()
}

// HasCityId returns a boolean if a field has been set.
func (o *TenantCreateDto) HasCityId() bool {
	if o != nil && o.CityId.IsSet() {
		return true
	}

	return false
}

// SetCityId gets a reference to the given NullableString and assigns it to the CityId field.
func (o *TenantCreateDto) SetCityId(v string) {
	o.CityId.Set(&v)
}
// SetCityIdNil sets the value for CityId to be an explicit nil
func (o *TenantCreateDto) SetCityIdNil() {
	o.CityId.Set(nil)
}

// UnsetCityId ensures that no value is present for CityId, not even an explicit nil
func (o *TenantCreateDto) UnsetCityId() {
	o.CityId.Unset()
}

// GetStateId returns the StateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantCreateDto) GetStateId() string {
	if o == nil || IsNil(o.StateId.Get()) {
		var ret string
		return ret
	}
	return *o.StateId.Get()
}

// GetStateIdOk returns a tuple with the StateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantCreateDto) GetStateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StateId.Get(), o.StateId.IsSet()
}

// HasStateId returns a boolean if a field has been set.
func (o *TenantCreateDto) HasStateId() bool {
	if o != nil && o.StateId.IsSet() {
		return true
	}

	return false
}

// SetStateId gets a reference to the given NullableString and assigns it to the StateId field.
func (o *TenantCreateDto) SetStateId(v string) {
	o.StateId.Set(&v)
}
// SetStateIdNil sets the value for StateId to be an explicit nil
func (o *TenantCreateDto) SetStateIdNil() {
	o.StateId.Set(nil)
}

// UnsetStateId ensures that no value is present for StateId, not even an explicit nil
func (o *TenantCreateDto) UnsetStateId() {
	o.StateId.Unset()
}

// GetCountryId returns the CountryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantCreateDto) GetCountryId() string {
	if o == nil || IsNil(o.CountryId.Get()) {
		var ret string
		return ret
	}
	return *o.CountryId.Get()
}

// GetCountryIdOk returns a tuple with the CountryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantCreateDto) GetCountryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryId.Get(), o.CountryId.IsSet()
}

// HasCountryId returns a boolean if a field has been set.
func (o *TenantCreateDto) HasCountryId() bool {
	if o != nil && o.CountryId.IsSet() {
		return true
	}

	return false
}

// SetCountryId gets a reference to the given NullableString and assigns it to the CountryId field.
func (o *TenantCreateDto) SetCountryId(v string) {
	o.CountryId.Set(&v)
}
// SetCountryIdNil sets the value for CountryId to be an explicit nil
func (o *TenantCreateDto) SetCountryIdNil() {
	o.CountryId.Set(nil)
}

// UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
func (o *TenantCreateDto) UnsetCountryId() {
	o.CountryId.Unset()
}

// GetTaxId returns the TaxId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantCreateDto) GetTaxId() string {
	if o == nil || IsNil(o.TaxId.Get()) {
		var ret string
		return ret
	}
	return *o.TaxId.Get()
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantCreateDto) GetTaxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxId.Get(), o.TaxId.IsSet()
}

// HasTaxId returns a boolean if a field has been set.
func (o *TenantCreateDto) HasTaxId() bool {
	if o != nil && o.TaxId.IsSet() {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given NullableString and assigns it to the TaxId field.
func (o *TenantCreateDto) SetTaxId(v string) {
	o.TaxId.Set(&v)
}
// SetTaxIdNil sets the value for TaxId to be an explicit nil
func (o *TenantCreateDto) SetTaxIdNil() {
	o.TaxId.Set(nil)
}

// UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
func (o *TenantCreateDto) UnsetTaxId() {
	o.TaxId.Unset()
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantCreateDto) GetAvatarUrl() string {
	if o == nil || IsNil(o.AvatarUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AvatarUrl.Get()
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantCreateDto) GetAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarUrl.Get(), o.AvatarUrl.IsSet()
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *TenantCreateDto) HasAvatarUrl() bool {
	if o != nil && o.AvatarUrl.IsSet() {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given NullableString and assigns it to the AvatarUrl field.
func (o *TenantCreateDto) SetAvatarUrl(v string) {
	o.AvatarUrl.Set(&v)
}
// SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil
func (o *TenantCreateDto) SetAvatarUrlNil() {
	o.AvatarUrl.Set(nil)
}

// UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
func (o *TenantCreateDto) UnsetAvatarUrl() {
	o.AvatarUrl.Unset()
}

func (o TenantCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Duns.IsSet() {
		toSerialize["duns"] = o.Duns.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.LegalName.IsSet() {
		toSerialize["legalName"] = o.LegalName.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if o.WebUrl.IsSet() {
		toSerialize["webUrl"] = o.WebUrl.Get()
	}
	if o.About.IsSet() {
		toSerialize["about"] = o.About.Get()
	}
	if o.Handler.IsSet() {
		toSerialize["handler"] = o.Handler.Get()
	}
	if o.CurrencyId.IsSet() {
		toSerialize["currencyId"] = o.CurrencyId.Get()
	}
	if o.LanguageId.IsSet() {
		toSerialize["languageId"] = o.LanguageId.Get()
	}
	if o.TimezoneId.IsSet() {
		toSerialize["timezoneId"] = o.TimezoneId.Get()
	}
	if o.CityId.IsSet() {
		toSerialize["cityId"] = o.CityId.Get()
	}
	if o.StateId.IsSet() {
		toSerialize["stateId"] = o.StateId.Get()
	}
	if o.CountryId.IsSet() {
		toSerialize["countryId"] = o.CountryId.Get()
	}
	if o.TaxId.IsSet() {
		toSerialize["taxId"] = o.TaxId.Get()
	}
	if o.AvatarUrl.IsSet() {
		toSerialize["avatarUrl"] = o.AvatarUrl.Get()
	}
	return toSerialize, nil
}

type NullableTenantCreateDto struct {
	value *TenantCreateDto
	isSet bool
}

func (v NullableTenantCreateDto) Get() *TenantCreateDto {
	return v.value
}

func (v *NullableTenantCreateDto) Set(val *TenantCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantCreateDto(val *TenantCreateDto) *NullableTenantCreateDto {
	return &NullableTenantCreateDto{value: val, isSet: true}
}

func (v NullableTenantCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


