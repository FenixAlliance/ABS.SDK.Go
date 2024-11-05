/*
HolderService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.1.4089
Contact: support@fenix-alliance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserSettingsUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSettingsUpdateDto{}

// UserSettingsUpdateDto struct for UserSettingsUpdateDto
type UserSettingsUpdateDto struct {
	PageSize NullableInt32 `json:"pageSize,omitempty"`
	DateFormat string `json:"dateFormat"`
	CurrencyFormat string `json:"currencyFormat"`
	DateTimeFormat string `json:"dateTimeFormat"`
	SiteTheme int32 `json:"siteTheme"`
}

type _UserSettingsUpdateDto UserSettingsUpdateDto

// NewUserSettingsUpdateDto instantiates a new UserSettingsUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSettingsUpdateDto(dateFormat string, currencyFormat string, dateTimeFormat string, siteTheme int32) *UserSettingsUpdateDto {
	this := UserSettingsUpdateDto{}
	this.DateFormat = dateFormat
	this.CurrencyFormat = currencyFormat
	this.DateTimeFormat = dateTimeFormat
	this.SiteTheme = siteTheme
	return &this
}

// NewUserSettingsUpdateDtoWithDefaults instantiates a new UserSettingsUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSettingsUpdateDtoWithDefaults() *UserSettingsUpdateDto {
	this := UserSettingsUpdateDto{}
	return &this
}

// GetPageSize returns the PageSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSettingsUpdateDto) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize.Get()) {
		var ret int32
		return ret
	}
	return *o.PageSize.Get()
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSettingsUpdateDto) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageSize.Get(), o.PageSize.IsSet()
}

// HasPageSize returns a boolean if a field has been set.
func (o *UserSettingsUpdateDto) HasPageSize() bool {
	if o != nil && o.PageSize.IsSet() {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given NullableInt32 and assigns it to the PageSize field.
func (o *UserSettingsUpdateDto) SetPageSize(v int32) {
	o.PageSize.Set(&v)
}
// SetPageSizeNil sets the value for PageSize to be an explicit nil
func (o *UserSettingsUpdateDto) SetPageSizeNil() {
	o.PageSize.Set(nil)
}

// UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
func (o *UserSettingsUpdateDto) UnsetPageSize() {
	o.PageSize.Unset()
}

// GetDateFormat returns the DateFormat field value
func (o *UserSettingsUpdateDto) GetDateFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateFormat
}

// GetDateFormatOk returns a tuple with the DateFormat field value
// and a boolean to check if the value has been set.
func (o *UserSettingsUpdateDto) GetDateFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateFormat, true
}

// SetDateFormat sets field value
func (o *UserSettingsUpdateDto) SetDateFormat(v string) {
	o.DateFormat = v
}

// GetCurrencyFormat returns the CurrencyFormat field value
func (o *UserSettingsUpdateDto) GetCurrencyFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyFormat
}

// GetCurrencyFormatOk returns a tuple with the CurrencyFormat field value
// and a boolean to check if the value has been set.
func (o *UserSettingsUpdateDto) GetCurrencyFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyFormat, true
}

// SetCurrencyFormat sets field value
func (o *UserSettingsUpdateDto) SetCurrencyFormat(v string) {
	o.CurrencyFormat = v
}

// GetDateTimeFormat returns the DateTimeFormat field value
func (o *UserSettingsUpdateDto) GetDateTimeFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateTimeFormat
}

// GetDateTimeFormatOk returns a tuple with the DateTimeFormat field value
// and a boolean to check if the value has been set.
func (o *UserSettingsUpdateDto) GetDateTimeFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateTimeFormat, true
}

// SetDateTimeFormat sets field value
func (o *UserSettingsUpdateDto) SetDateTimeFormat(v string) {
	o.DateTimeFormat = v
}

// GetSiteTheme returns the SiteTheme field value
func (o *UserSettingsUpdateDto) GetSiteTheme() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SiteTheme
}

// GetSiteThemeOk returns a tuple with the SiteTheme field value
// and a boolean to check if the value has been set.
func (o *UserSettingsUpdateDto) GetSiteThemeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteTheme, true
}

// SetSiteTheme sets field value
func (o *UserSettingsUpdateDto) SetSiteTheme(v int32) {
	o.SiteTheme = v
}

func (o UserSettingsUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSettingsUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PageSize.IsSet() {
		toSerialize["pageSize"] = o.PageSize.Get()
	}
	toSerialize["dateFormat"] = o.DateFormat
	toSerialize["currencyFormat"] = o.CurrencyFormat
	toSerialize["dateTimeFormat"] = o.DateTimeFormat
	toSerialize["siteTheme"] = o.SiteTheme
	return toSerialize, nil
}

func (o *UserSettingsUpdateDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dateFormat",
		"currencyFormat",
		"dateTimeFormat",
		"siteTheme",
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

	varUserSettingsUpdateDto := _UserSettingsUpdateDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSettingsUpdateDto)

	if err != nil {
		return err
	}

	*o = UserSettingsUpdateDto(varUserSettingsUpdateDto)

	return err
}

type NullableUserSettingsUpdateDto struct {
	value *UserSettingsUpdateDto
	isSet bool
}

func (v NullableUserSettingsUpdateDto) Get() *UserSettingsUpdateDto {
	return v.value
}

func (v *NullableUserSettingsUpdateDto) Set(val *UserSettingsUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSettingsUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSettingsUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSettingsUpdateDto(val *UserSettingsUpdateDto) *NullableUserSettingsUpdateDto {
	return &NullableUserSettingsUpdateDto{value: val, isSet: true}
}

func (v NullableUserSettingsUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSettingsUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


