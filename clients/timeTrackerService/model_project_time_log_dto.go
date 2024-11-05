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
)

// checks if the ProjectTimeLogDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectTimeLogDto{}

// ProjectTimeLogDto struct for ProjectTimeLogDto
type ProjectTimeLogDto struct {
	Id NullableString `json:"id,omitempty"`
	Timestamp NullableTime `json:"timestamp,omitempty"`
	ProjectId NullableString `json:"projectId,omitempty"`
	ProjectTaskId NullableString `json:"projectTaskId,omitempty"`
	TaskCategoryId NullableString `json:"taskCategoryId,omitempty"`
	ProjectPeriodId NullableString `json:"projectPeriodId,omitempty"`
	ResponsibleContactId NullableString `json:"responsibleContactId,omitempty"`
	CreatorContactId NullableString `json:"creatorContactId,omitempty"`
	RecordType *int32 `json:"recordType,omitempty"`
	TimeStamp *time.Time `json:"timeStamp,omitempty"`
	TimeSpan *string `json:"timeSpan,omitempty"`
	LogDate *time.Time `json:"logDate,omitempty"`
	Comments NullableString `json:"comments,omitempty"`
	Type NullableString `json:"type,omitempty"`
}

// NewProjectTimeLogDto instantiates a new ProjectTimeLogDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectTimeLogDto() *ProjectTimeLogDto {
	this := ProjectTimeLogDto{}
	return &this
}

// NewProjectTimeLogDtoWithDefaults instantiates a new ProjectTimeLogDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectTimeLogDtoWithDefaults() *ProjectTimeLogDto {
	this := ProjectTimeLogDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTimeLogDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTimeLogDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ProjectTimeLogDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ProjectTimeLogDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ProjectTimeLogDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ProjectTimeLogDto) UnsetId() {
	o.Id.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTimeLogDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTimeLogDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ProjectTimeLogDto) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *ProjectTimeLogDto) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *ProjectTimeLogDto) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *ProjectTimeLogDto) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTimeLogDto) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTimeLogDto) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *ProjectTimeLogDto) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableString and assigns it to the ProjectId field.
func (o *ProjectTimeLogDto) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *ProjectTimeLogDto) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *ProjectTimeLogDto) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetProjectTaskId returns the ProjectTaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTimeLogDto) GetProjectTaskId() string {
	if o == nil || IsNil(o.ProjectTaskId.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectTaskId.Get()
}

// GetProjectTaskIdOk returns a tuple with the ProjectTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTimeLogDto) GetProjectTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectTaskId.Get(), o.ProjectTaskId.IsSet()
}

// HasProjectTaskId returns a boolean if a field has been set.
func (o *ProjectTimeLogDto) HasProjectTaskId() bool {
	if o != nil && o.ProjectTaskId.IsSet() {
		return true
	}

	return false
}

// SetProjectTaskId gets a reference to the given NullableString and assigns it to the ProjectTaskId field.
func (o *ProjectTimeLogDto) SetProjectTaskId(v string) {
	o.ProjectTaskId.Set(&v)
}
// SetProjectTaskIdNil sets the value for ProjectTaskId to be an explicit nil
func (o *ProjectTimeLogDto) SetProjectTaskIdNil() {
	o.ProjectTaskId.Set(nil)
}

// UnsetProjectTaskId ensures that no value is present for ProjectTaskId, not even an explicit nil
func (o *ProjectTimeLogDto) UnsetProjectTaskId() {
	o.ProjectTaskId.Unset()
}

// GetTaskCategoryId returns the TaskCategoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTimeLogDto) GetTaskCategoryId() string {
	if o == nil || IsNil(o.TaskCategoryId.Get()) {
		var ret string
		return ret
	}
	return *o.TaskCategoryId.Get()
}

// GetTaskCategoryIdOk returns a tuple with the TaskCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTimeLogDto) GetTaskCategoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskCategoryId.Get(), o.TaskCategoryId.IsSet()
}

// HasTaskCategoryId returns a boolean if a field has been set.
func (o *ProjectTimeLogDto) HasTaskCategoryId() bool {
	if o != nil && o.TaskCategoryId.IsSet() {
		return true
	}

	return false
}

// SetTaskCategoryId gets a reference to the given NullableString and assigns it to the TaskCategoryId field.
func (o *ProjectTimeLogDto) SetTaskCategoryId(v string) {
	o.TaskCategoryId.Set(&v)
}
// SetTaskCategoryIdNil sets the value for TaskCategoryId to be an explicit nil
func (o *ProjectTimeLogDto) SetTaskCategoryIdNil() {
	o.TaskCategoryId.Set(nil)
}

// UnsetTaskCategoryId ensures that no value is present for TaskCategoryId, not even an explicit nil
func (o *ProjectTimeLogDto) UnsetTaskCategoryId() {
	o.TaskCategoryId.Unset()
}

// GetProjectPeriodId returns the ProjectPeriodId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTimeLogDto) GetProjectPeriodId() string {
	if o == nil || IsNil(o.ProjectPeriodId.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectPeriodId.Get()
}

// GetProjectPeriodIdOk returns a tuple with the ProjectPeriodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTimeLogDto) GetProjectPeriodIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectPeriodId.Get(), o.ProjectPeriodId.IsSet()
}

// HasProjectPeriodId returns a boolean if a field has been set.
func (o *ProjectTimeLogDto) HasProjectPeriodId() bool {
	if o != nil && o.ProjectPeriodId.IsSet() {
		return true
	}

	return false
}

// SetProjectPeriodId gets a reference to the given NullableString and assigns it to the ProjectPeriodId field.
func (o *ProjectTimeLogDto) SetProjectPeriodId(v string) {
	o.ProjectPeriodId.Set(&v)
}
// SetProjectPeriodIdNil sets the value for ProjectPeriodId to be an explicit nil
func (o *ProjectTimeLogDto) SetProjectPeriodIdNil() {
	o.ProjectPeriodId.Set(nil)
}

// UnsetProjectPeriodId ensures that no value is present for ProjectPeriodId, not even an explicit nil
func (o *ProjectTimeLogDto) UnsetProjectPeriodId() {
	o.ProjectPeriodId.Unset()
}

// GetResponsibleContactId returns the ResponsibleContactId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTimeLogDto) GetResponsibleContactId() string {
	if o == nil || IsNil(o.ResponsibleContactId.Get()) {
		var ret string
		return ret
	}
	return *o.ResponsibleContactId.Get()
}

// GetResponsibleContactIdOk returns a tuple with the ResponsibleContactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTimeLogDto) GetResponsibleContactIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponsibleContactId.Get(), o.ResponsibleContactId.IsSet()
}

// HasResponsibleContactId returns a boolean if a field has been set.
func (o *ProjectTimeLogDto) HasResponsibleContactId() bool {
	if o != nil && o.ResponsibleContactId.IsSet() {
		return true
	}

	return false
}

// SetResponsibleContactId gets a reference to the given NullableString and assigns it to the ResponsibleContactId field.
func (o *ProjectTimeLogDto) SetResponsibleContactId(v string) {
	o.ResponsibleContactId.Set(&v)
}
// SetResponsibleContactIdNil sets the value for ResponsibleContactId to be an explicit nil
func (o *ProjectTimeLogDto) SetResponsibleContactIdNil() {
	o.ResponsibleContactId.Set(nil)
}

// UnsetResponsibleContactId ensures that no value is present for ResponsibleContactId, not even an explicit nil
func (o *ProjectTimeLogDto) UnsetResponsibleContactId() {
	o.ResponsibleContactId.Unset()
}

// GetCreatorContactId returns the CreatorContactId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTimeLogDto) GetCreatorContactId() string {
	if o == nil || IsNil(o.CreatorContactId.Get()) {
		var ret string
		return ret
	}
	return *o.CreatorContactId.Get()
}

// GetCreatorContactIdOk returns a tuple with the CreatorContactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTimeLogDto) GetCreatorContactIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorContactId.Get(), o.CreatorContactId.IsSet()
}

// HasCreatorContactId returns a boolean if a field has been set.
func (o *ProjectTimeLogDto) HasCreatorContactId() bool {
	if o != nil && o.CreatorContactId.IsSet() {
		return true
	}

	return false
}

// SetCreatorContactId gets a reference to the given NullableString and assigns it to the CreatorContactId field.
func (o *ProjectTimeLogDto) SetCreatorContactId(v string) {
	o.CreatorContactId.Set(&v)
}
// SetCreatorContactIdNil sets the value for CreatorContactId to be an explicit nil
func (o *ProjectTimeLogDto) SetCreatorContactIdNil() {
	o.CreatorContactId.Set(nil)
}

// UnsetCreatorContactId ensures that no value is present for CreatorContactId, not even an explicit nil
func (o *ProjectTimeLogDto) UnsetCreatorContactId() {
	o.CreatorContactId.Unset()
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *ProjectTimeLogDto) GetRecordType() int32 {
	if o == nil || IsNil(o.RecordType) {
		var ret int32
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTimeLogDto) GetRecordTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *ProjectTimeLogDto) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given int32 and assigns it to the RecordType field.
func (o *ProjectTimeLogDto) SetRecordType(v int32) {
	o.RecordType = &v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *ProjectTimeLogDto) GetTimeStamp() time.Time {
	if o == nil || IsNil(o.TimeStamp) {
		var ret time.Time
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTimeLogDto) GetTimeStampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimeStamp) {
		return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *ProjectTimeLogDto) HasTimeStamp() bool {
	if o != nil && !IsNil(o.TimeStamp) {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given time.Time and assigns it to the TimeStamp field.
func (o *ProjectTimeLogDto) SetTimeStamp(v time.Time) {
	o.TimeStamp = &v
}

// GetTimeSpan returns the TimeSpan field value if set, zero value otherwise.
func (o *ProjectTimeLogDto) GetTimeSpan() string {
	if o == nil || IsNil(o.TimeSpan) {
		var ret string
		return ret
	}
	return *o.TimeSpan
}

// GetTimeSpanOk returns a tuple with the TimeSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTimeLogDto) GetTimeSpanOk() (*string, bool) {
	if o == nil || IsNil(o.TimeSpan) {
		return nil, false
	}
	return o.TimeSpan, true
}

// HasTimeSpan returns a boolean if a field has been set.
func (o *ProjectTimeLogDto) HasTimeSpan() bool {
	if o != nil && !IsNil(o.TimeSpan) {
		return true
	}

	return false
}

// SetTimeSpan gets a reference to the given string and assigns it to the TimeSpan field.
func (o *ProjectTimeLogDto) SetTimeSpan(v string) {
	o.TimeSpan = &v
}

// GetLogDate returns the LogDate field value if set, zero value otherwise.
func (o *ProjectTimeLogDto) GetLogDate() time.Time {
	if o == nil || IsNil(o.LogDate) {
		var ret time.Time
		return ret
	}
	return *o.LogDate
}

// GetLogDateOk returns a tuple with the LogDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTimeLogDto) GetLogDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LogDate) {
		return nil, false
	}
	return o.LogDate, true
}

// HasLogDate returns a boolean if a field has been set.
func (o *ProjectTimeLogDto) HasLogDate() bool {
	if o != nil && !IsNil(o.LogDate) {
		return true
	}

	return false
}

// SetLogDate gets a reference to the given time.Time and assigns it to the LogDate field.
func (o *ProjectTimeLogDto) SetLogDate(v time.Time) {
	o.LogDate = &v
}

// GetComments returns the Comments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTimeLogDto) GetComments() string {
	if o == nil || IsNil(o.Comments.Get()) {
		var ret string
		return ret
	}
	return *o.Comments.Get()
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTimeLogDto) GetCommentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comments.Get(), o.Comments.IsSet()
}

// HasComments returns a boolean if a field has been set.
func (o *ProjectTimeLogDto) HasComments() bool {
	if o != nil && o.Comments.IsSet() {
		return true
	}

	return false
}

// SetComments gets a reference to the given NullableString and assigns it to the Comments field.
func (o *ProjectTimeLogDto) SetComments(v string) {
	o.Comments.Set(&v)
}
// SetCommentsNil sets the value for Comments to be an explicit nil
func (o *ProjectTimeLogDto) SetCommentsNil() {
	o.Comments.Set(nil)
}

// UnsetComments ensures that no value is present for Comments, not even an explicit nil
func (o *ProjectTimeLogDto) UnsetComments() {
	o.Comments.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTimeLogDto) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTimeLogDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ProjectTimeLogDto) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *ProjectTimeLogDto) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ProjectTimeLogDto) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ProjectTimeLogDto) UnsetType() {
	o.Type.Unset()
}

func (o ProjectTimeLogDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectTimeLogDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.ProjectId.IsSet() {
		toSerialize["projectId"] = o.ProjectId.Get()
	}
	if o.ProjectTaskId.IsSet() {
		toSerialize["projectTaskId"] = o.ProjectTaskId.Get()
	}
	if o.TaskCategoryId.IsSet() {
		toSerialize["taskCategoryId"] = o.TaskCategoryId.Get()
	}
	if o.ProjectPeriodId.IsSet() {
		toSerialize["projectPeriodId"] = o.ProjectPeriodId.Get()
	}
	if o.ResponsibleContactId.IsSet() {
		toSerialize["responsibleContactId"] = o.ResponsibleContactId.Get()
	}
	if o.CreatorContactId.IsSet() {
		toSerialize["creatorContactId"] = o.CreatorContactId.Get()
	}
	if !IsNil(o.RecordType) {
		toSerialize["recordType"] = o.RecordType
	}
	if !IsNil(o.TimeStamp) {
		toSerialize["timeStamp"] = o.TimeStamp
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
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	return toSerialize, nil
}

type NullableProjectTimeLogDto struct {
	value *ProjectTimeLogDto
	isSet bool
}

func (v NullableProjectTimeLogDto) Get() *ProjectTimeLogDto {
	return v.value
}

func (v *NullableProjectTimeLogDto) Set(val *ProjectTimeLogDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectTimeLogDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectTimeLogDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectTimeLogDto(val *ProjectTimeLogDto) *NullableProjectTimeLogDto {
	return &NullableProjectTimeLogDto{value: val, isSet: true}
}

func (v NullableProjectTimeLogDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectTimeLogDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


