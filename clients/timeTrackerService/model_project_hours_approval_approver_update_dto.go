/*
TimeTrackerService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.1.4089
Contact: support@fenix-alliance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ProjectHoursApprovalApproverUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectHoursApprovalApproverUpdateDto{}

// ProjectHoursApprovalApproverUpdateDto struct for ProjectHoursApprovalApproverUpdateDto
type ProjectHoursApprovalApproverUpdateDto struct {
	ApproverContactID NullableString `json:"approverContactID,omitempty"`
}

// NewProjectHoursApprovalApproverUpdateDto instantiates a new ProjectHoursApprovalApproverUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectHoursApprovalApproverUpdateDto() *ProjectHoursApprovalApproverUpdateDto {
	this := ProjectHoursApprovalApproverUpdateDto{}
	return &this
}

// NewProjectHoursApprovalApproverUpdateDtoWithDefaults instantiates a new ProjectHoursApprovalApproverUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectHoursApprovalApproverUpdateDtoWithDefaults() *ProjectHoursApprovalApproverUpdateDto {
	this := ProjectHoursApprovalApproverUpdateDto{}
	return &this
}

// GetApproverContactID returns the ApproverContactID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectHoursApprovalApproverUpdateDto) GetApproverContactID() string {
	if o == nil || IsNil(o.ApproverContactID.Get()) {
		var ret string
		return ret
	}
	return *o.ApproverContactID.Get()
}

// GetApproverContactIDOk returns a tuple with the ApproverContactID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectHoursApprovalApproverUpdateDto) GetApproverContactIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApproverContactID.Get(), o.ApproverContactID.IsSet()
}

// HasApproverContactID returns a boolean if a field has been set.
func (o *ProjectHoursApprovalApproverUpdateDto) HasApproverContactID() bool {
	if o != nil && o.ApproverContactID.IsSet() {
		return true
	}

	return false
}

// SetApproverContactID gets a reference to the given NullableString and assigns it to the ApproverContactID field.
func (o *ProjectHoursApprovalApproverUpdateDto) SetApproverContactID(v string) {
	o.ApproverContactID.Set(&v)
}
// SetApproverContactIDNil sets the value for ApproverContactID to be an explicit nil
func (o *ProjectHoursApprovalApproverUpdateDto) SetApproverContactIDNil() {
	o.ApproverContactID.Set(nil)
}

// UnsetApproverContactID ensures that no value is present for ApproverContactID, not even an explicit nil
func (o *ProjectHoursApprovalApproverUpdateDto) UnsetApproverContactID() {
	o.ApproverContactID.Unset()
}

func (o ProjectHoursApprovalApproverUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectHoursApprovalApproverUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApproverContactID.IsSet() {
		toSerialize["approverContactID"] = o.ApproverContactID.Get()
	}
	return toSerialize, nil
}

type NullableProjectHoursApprovalApproverUpdateDto struct {
	value *ProjectHoursApprovalApproverUpdateDto
	isSet bool
}

func (v NullableProjectHoursApprovalApproverUpdateDto) Get() *ProjectHoursApprovalApproverUpdateDto {
	return v.value
}

func (v *NullableProjectHoursApprovalApproverUpdateDto) Set(val *ProjectHoursApprovalApproverUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectHoursApprovalApproverUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectHoursApprovalApproverUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectHoursApprovalApproverUpdateDto(val *ProjectHoursApprovalApproverUpdateDto) *NullableProjectHoursApprovalApproverUpdateDto {
	return &NullableProjectHoursApprovalApproverUpdateDto{value: val, isSet: true}
}

func (v NullableProjectHoursApprovalApproverUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectHoursApprovalApproverUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

