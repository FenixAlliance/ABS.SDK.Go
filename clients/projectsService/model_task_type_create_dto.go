/*
ProjectsService

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

// checks if the TaskTypeCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskTypeCreateDto{}

// TaskTypeCreateDto struct for TaskTypeCreateDto
type TaskTypeCreateDto struct {
	Id *string `json:"id,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Title NullableString `json:"title,omitempty"`
	TaskCategoryID NullableString `json:"taskCategoryID,omitempty"`
	DisplayInTimeTracker *bool `json:"displayInTimeTracker,omitempty"`
	RequiresDescription *bool `json:"requiresDescription,omitempty"`
}

// NewTaskTypeCreateDto instantiates a new TaskTypeCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskTypeCreateDto() *TaskTypeCreateDto {
	this := TaskTypeCreateDto{}
	return &this
}

// NewTaskTypeCreateDtoWithDefaults instantiates a new TaskTypeCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskTypeCreateDtoWithDefaults() *TaskTypeCreateDto {
	this := TaskTypeCreateDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaskTypeCreateDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskTypeCreateDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaskTypeCreateDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TaskTypeCreateDto) SetId(v string) {
	o.Id = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TaskTypeCreateDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskTypeCreateDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TaskTypeCreateDto) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *TaskTypeCreateDto) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskTypeCreateDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskTypeCreateDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *TaskTypeCreateDto) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *TaskTypeCreateDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *TaskTypeCreateDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *TaskTypeCreateDto) UnsetTitle() {
	o.Title.Unset()
}

// GetTaskCategoryID returns the TaskCategoryID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskTypeCreateDto) GetTaskCategoryID() string {
	if o == nil || IsNil(o.TaskCategoryID.Get()) {
		var ret string
		return ret
	}
	return *o.TaskCategoryID.Get()
}

// GetTaskCategoryIDOk returns a tuple with the TaskCategoryID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskTypeCreateDto) GetTaskCategoryIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskCategoryID.Get(), o.TaskCategoryID.IsSet()
}

// HasTaskCategoryID returns a boolean if a field has been set.
func (o *TaskTypeCreateDto) HasTaskCategoryID() bool {
	if o != nil && o.TaskCategoryID.IsSet() {
		return true
	}

	return false
}

// SetTaskCategoryID gets a reference to the given NullableString and assigns it to the TaskCategoryID field.
func (o *TaskTypeCreateDto) SetTaskCategoryID(v string) {
	o.TaskCategoryID.Set(&v)
}
// SetTaskCategoryIDNil sets the value for TaskCategoryID to be an explicit nil
func (o *TaskTypeCreateDto) SetTaskCategoryIDNil() {
	o.TaskCategoryID.Set(nil)
}

// UnsetTaskCategoryID ensures that no value is present for TaskCategoryID, not even an explicit nil
func (o *TaskTypeCreateDto) UnsetTaskCategoryID() {
	o.TaskCategoryID.Unset()
}

// GetDisplayInTimeTracker returns the DisplayInTimeTracker field value if set, zero value otherwise.
func (o *TaskTypeCreateDto) GetDisplayInTimeTracker() bool {
	if o == nil || IsNil(o.DisplayInTimeTracker) {
		var ret bool
		return ret
	}
	return *o.DisplayInTimeTracker
}

// GetDisplayInTimeTrackerOk returns a tuple with the DisplayInTimeTracker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskTypeCreateDto) GetDisplayInTimeTrackerOk() (*bool, bool) {
	if o == nil || IsNil(o.DisplayInTimeTracker) {
		return nil, false
	}
	return o.DisplayInTimeTracker, true
}

// HasDisplayInTimeTracker returns a boolean if a field has been set.
func (o *TaskTypeCreateDto) HasDisplayInTimeTracker() bool {
	if o != nil && !IsNil(o.DisplayInTimeTracker) {
		return true
	}

	return false
}

// SetDisplayInTimeTracker gets a reference to the given bool and assigns it to the DisplayInTimeTracker field.
func (o *TaskTypeCreateDto) SetDisplayInTimeTracker(v bool) {
	o.DisplayInTimeTracker = &v
}

// GetRequiresDescription returns the RequiresDescription field value if set, zero value otherwise.
func (o *TaskTypeCreateDto) GetRequiresDescription() bool {
	if o == nil || IsNil(o.RequiresDescription) {
		var ret bool
		return ret
	}
	return *o.RequiresDescription
}

// GetRequiresDescriptionOk returns a tuple with the RequiresDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskTypeCreateDto) GetRequiresDescriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiresDescription) {
		return nil, false
	}
	return o.RequiresDescription, true
}

// HasRequiresDescription returns a boolean if a field has been set.
func (o *TaskTypeCreateDto) HasRequiresDescription() bool {
	if o != nil && !IsNil(o.RequiresDescription) {
		return true
	}

	return false
}

// SetRequiresDescription gets a reference to the given bool and assigns it to the RequiresDescription field.
func (o *TaskTypeCreateDto) SetRequiresDescription(v bool) {
	o.RequiresDescription = &v
}

func (o TaskTypeCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskTypeCreateDto) ToMap() (map[string]interface{}, error) {
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
	if o.TaskCategoryID.IsSet() {
		toSerialize["taskCategoryID"] = o.TaskCategoryID.Get()
	}
	if !IsNil(o.DisplayInTimeTracker) {
		toSerialize["displayInTimeTracker"] = o.DisplayInTimeTracker
	}
	if !IsNil(o.RequiresDescription) {
		toSerialize["requiresDescription"] = o.RequiresDescription
	}
	return toSerialize, nil
}

type NullableTaskTypeCreateDto struct {
	value *TaskTypeCreateDto
	isSet bool
}

func (v NullableTaskTypeCreateDto) Get() *TaskTypeCreateDto {
	return v.value
}

func (v *NullableTaskTypeCreateDto) Set(val *TaskTypeCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskTypeCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskTypeCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskTypeCreateDto(val *TaskTypeCreateDto) *NullableTaskTypeCreateDto {
	return &NullableTaskTypeCreateDto{value: val, isSet: true}
}

func (v NullableTaskTypeCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskTypeCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


