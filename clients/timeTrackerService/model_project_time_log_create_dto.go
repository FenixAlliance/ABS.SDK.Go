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
	"time"
	"bytes"
	"fmt"
)

// checks if the ProjectTimeLogCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectTimeLogCreateDto{}

// ProjectTimeLogCreateDto struct for ProjectTimeLogCreateDto
type ProjectTimeLogCreateDto struct {
	Id *string `json:"id,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	TimeSpan *string `json:"timeSpan,omitempty"`
	LogDate *time.Time `json:"logDate,omitempty"`
	Comments NullableString `json:"comments,omitempty"`
	ProjectTaskID string `json:"projectTaskID"`
	ProjectPeriodID string `json:"projectPeriodID"`
	ProjectTimeLogRecordType *int32 `json:"projectTimeLogRecordType,omitempty"`
	ProjectID NullableString `json:"projectID,omitempty"`
}

type _ProjectTimeLogCreateDto ProjectTimeLogCreateDto

// NewProjectTimeLogCreateDto instantiates a new ProjectTimeLogCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectTimeLogCreateDto(projectTaskID string, projectPeriodID string) *ProjectTimeLogCreateDto {
	this := ProjectTimeLogCreateDto{}
	this.ProjectTaskID = projectTaskID
	this.ProjectPeriodID = projectPeriodID
	return &this
}

// NewProjectTimeLogCreateDtoWithDefaults instantiates a new ProjectTimeLogCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectTimeLogCreateDtoWithDefaults() *ProjectTimeLogCreateDto {
	this := ProjectTimeLogCreateDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectTimeLogCreateDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTimeLogCreateDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectTimeLogCreateDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProjectTimeLogCreateDto) SetId(v string) {
	o.Id = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ProjectTimeLogCreateDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTimeLogCreateDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ProjectTimeLogCreateDto) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *ProjectTimeLogCreateDto) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTimeSpan returns the TimeSpan field value if set, zero value otherwise.
func (o *ProjectTimeLogCreateDto) GetTimeSpan() string {
	if o == nil || IsNil(o.TimeSpan) {
		var ret string
		return ret
	}
	return *o.TimeSpan
}

// GetTimeSpanOk returns a tuple with the TimeSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTimeLogCreateDto) GetTimeSpanOk() (*string, bool) {
	if o == nil || IsNil(o.TimeSpan) {
		return nil, false
	}
	return o.TimeSpan, true
}

// HasTimeSpan returns a boolean if a field has been set.
func (o *ProjectTimeLogCreateDto) HasTimeSpan() bool {
	if o != nil && !IsNil(o.TimeSpan) {
		return true
	}

	return false
}

// SetTimeSpan gets a reference to the given string and assigns it to the TimeSpan field.
func (o *ProjectTimeLogCreateDto) SetTimeSpan(v string) {
	o.TimeSpan = &v
}

// GetLogDate returns the LogDate field value if set, zero value otherwise.
func (o *ProjectTimeLogCreateDto) GetLogDate() time.Time {
	if o == nil || IsNil(o.LogDate) {
		var ret time.Time
		return ret
	}
	return *o.LogDate
}

// GetLogDateOk returns a tuple with the LogDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTimeLogCreateDto) GetLogDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LogDate) {
		return nil, false
	}
	return o.LogDate, true
}

// HasLogDate returns a boolean if a field has been set.
func (o *ProjectTimeLogCreateDto) HasLogDate() bool {
	if o != nil && !IsNil(o.LogDate) {
		return true
	}

	return false
}

// SetLogDate gets a reference to the given time.Time and assigns it to the LogDate field.
func (o *ProjectTimeLogCreateDto) SetLogDate(v time.Time) {
	o.LogDate = &v
}

// GetComments returns the Comments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTimeLogCreateDto) GetComments() string {
	if o == nil || IsNil(o.Comments.Get()) {
		var ret string
		return ret
	}
	return *o.Comments.Get()
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTimeLogCreateDto) GetCommentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comments.Get(), o.Comments.IsSet()
}

// HasComments returns a boolean if a field has been set.
func (o *ProjectTimeLogCreateDto) HasComments() bool {
	if o != nil && o.Comments.IsSet() {
		return true
	}

	return false
}

// SetComments gets a reference to the given NullableString and assigns it to the Comments field.
func (o *ProjectTimeLogCreateDto) SetComments(v string) {
	o.Comments.Set(&v)
}
// SetCommentsNil sets the value for Comments to be an explicit nil
func (o *ProjectTimeLogCreateDto) SetCommentsNil() {
	o.Comments.Set(nil)
}

// UnsetComments ensures that no value is present for Comments, not even an explicit nil
func (o *ProjectTimeLogCreateDto) UnsetComments() {
	o.Comments.Unset()
}

// GetProjectTaskID returns the ProjectTaskID field value
func (o *ProjectTimeLogCreateDto) GetProjectTaskID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectTaskID
}

// GetProjectTaskIDOk returns a tuple with the ProjectTaskID field value
// and a boolean to check if the value has been set.
func (o *ProjectTimeLogCreateDto) GetProjectTaskIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectTaskID, true
}

// SetProjectTaskID sets field value
func (o *ProjectTimeLogCreateDto) SetProjectTaskID(v string) {
	o.ProjectTaskID = v
}

// GetProjectPeriodID returns the ProjectPeriodID field value
func (o *ProjectTimeLogCreateDto) GetProjectPeriodID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectPeriodID
}

// GetProjectPeriodIDOk returns a tuple with the ProjectPeriodID field value
// and a boolean to check if the value has been set.
func (o *ProjectTimeLogCreateDto) GetProjectPeriodIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectPeriodID, true
}

// SetProjectPeriodID sets field value
func (o *ProjectTimeLogCreateDto) SetProjectPeriodID(v string) {
	o.ProjectPeriodID = v
}

// GetProjectTimeLogRecordType returns the ProjectTimeLogRecordType field value if set, zero value otherwise.
func (o *ProjectTimeLogCreateDto) GetProjectTimeLogRecordType() int32 {
	if o == nil || IsNil(o.ProjectTimeLogRecordType) {
		var ret int32
		return ret
	}
	return *o.ProjectTimeLogRecordType
}

// GetProjectTimeLogRecordTypeOk returns a tuple with the ProjectTimeLogRecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTimeLogCreateDto) GetProjectTimeLogRecordTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectTimeLogRecordType) {
		return nil, false
	}
	return o.ProjectTimeLogRecordType, true
}

// HasProjectTimeLogRecordType returns a boolean if a field has been set.
func (o *ProjectTimeLogCreateDto) HasProjectTimeLogRecordType() bool {
	if o != nil && !IsNil(o.ProjectTimeLogRecordType) {
		return true
	}

	return false
}

// SetProjectTimeLogRecordType gets a reference to the given int32 and assigns it to the ProjectTimeLogRecordType field.
func (o *ProjectTimeLogCreateDto) SetProjectTimeLogRecordType(v int32) {
	o.ProjectTimeLogRecordType = &v
}

// GetProjectID returns the ProjectID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTimeLogCreateDto) GetProjectID() string {
	if o == nil || IsNil(o.ProjectID.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectID.Get()
}

// GetProjectIDOk returns a tuple with the ProjectID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTimeLogCreateDto) GetProjectIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectID.Get(), o.ProjectID.IsSet()
}

// HasProjectID returns a boolean if a field has been set.
func (o *ProjectTimeLogCreateDto) HasProjectID() bool {
	if o != nil && o.ProjectID.IsSet() {
		return true
	}

	return false
}

// SetProjectID gets a reference to the given NullableString and assigns it to the ProjectID field.
func (o *ProjectTimeLogCreateDto) SetProjectID(v string) {
	o.ProjectID.Set(&v)
}
// SetProjectIDNil sets the value for ProjectID to be an explicit nil
func (o *ProjectTimeLogCreateDto) SetProjectIDNil() {
	o.ProjectID.Set(nil)
}

// UnsetProjectID ensures that no value is present for ProjectID, not even an explicit nil
func (o *ProjectTimeLogCreateDto) UnsetProjectID() {
	o.ProjectID.Unset()
}

func (o ProjectTimeLogCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectTimeLogCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.TimeSpan) {
		toSerialize["timeSpan"] = o.TimeSpan
	}
	if !IsNil(o.LogDate) {
		toSerialize["logDate"] = o.LogDate
	}
	if o.Comments.IsSet() {
		toSerialize["comments"] = o.Comments.Get()
	}
	toSerialize["projectTaskID"] = o.ProjectTaskID
	toSerialize["projectPeriodID"] = o.ProjectPeriodID
	if !IsNil(o.ProjectTimeLogRecordType) {
		toSerialize["projectTimeLogRecordType"] = o.ProjectTimeLogRecordType
	}
	if o.ProjectID.IsSet() {
		toSerialize["projectID"] = o.ProjectID.Get()
	}
	return toSerialize, nil
}

func (o *ProjectTimeLogCreateDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"projectTaskID",
		"projectPeriodID",
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

	varProjectTimeLogCreateDto := _ProjectTimeLogCreateDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectTimeLogCreateDto)

	if err != nil {
		return err
	}

	*o = ProjectTimeLogCreateDto(varProjectTimeLogCreateDto)

	return err
}

type NullableProjectTimeLogCreateDto struct {
	value *ProjectTimeLogCreateDto
	isSet bool
}

func (v NullableProjectTimeLogCreateDto) Get() *ProjectTimeLogCreateDto {
	return v.value
}

func (v *NullableProjectTimeLogCreateDto) Set(val *ProjectTimeLogCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectTimeLogCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectTimeLogCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectTimeLogCreateDto(val *ProjectTimeLogCreateDto) *NullableProjectTimeLogCreateDto {
	return &NullableProjectTimeLogCreateDto{value: val, isSet: true}
}

func (v NullableProjectTimeLogCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectTimeLogCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


