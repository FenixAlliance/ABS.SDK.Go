# ActivityRecordUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Completed** | Pointer to **NullableBool** |  | [optional] 
**DueDate** | Pointer to **NullableTime** |  | [optional] 
**ActivityTypeId** | Pointer to **NullableString** |  | [optional] 
**ParentActivityId** | Pointer to **NullableString** |  | [optional] 
**InChargeEnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewActivityRecordUpdateDto

`func NewActivityRecordUpdateDto() *ActivityRecordUpdateDto`

NewActivityRecordUpdateDto instantiates a new ActivityRecordUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityRecordUpdateDtoWithDefaults

`func NewActivityRecordUpdateDtoWithDefaults() *ActivityRecordUpdateDto`

NewActivityRecordUpdateDtoWithDefaults instantiates a new ActivityRecordUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActivityRecordUpdateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityRecordUpdateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityRecordUpdateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActivityRecordUpdateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ActivityRecordUpdateDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ActivityRecordUpdateDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *ActivityRecordUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ActivityRecordUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ActivityRecordUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ActivityRecordUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ActivityRecordUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ActivityRecordUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *ActivityRecordUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActivityRecordUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActivityRecordUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActivityRecordUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ActivityRecordUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ActivityRecordUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCompleted

`func (o *ActivityRecordUpdateDto) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *ActivityRecordUpdateDto) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *ActivityRecordUpdateDto) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *ActivityRecordUpdateDto) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### SetCompletedNil

`func (o *ActivityRecordUpdateDto) SetCompletedNil(b bool)`

 SetCompletedNil sets the value for Completed to be an explicit nil

### UnsetCompleted
`func (o *ActivityRecordUpdateDto) UnsetCompleted()`

UnsetCompleted ensures that no value is present for Completed, not even an explicit nil
### GetDueDate

`func (o *ActivityRecordUpdateDto) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *ActivityRecordUpdateDto) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *ActivityRecordUpdateDto) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *ActivityRecordUpdateDto) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *ActivityRecordUpdateDto) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *ActivityRecordUpdateDto) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetActivityTypeId

`func (o *ActivityRecordUpdateDto) GetActivityTypeId() string`

GetActivityTypeId returns the ActivityTypeId field if non-nil, zero value otherwise.

### GetActivityTypeIdOk

`func (o *ActivityRecordUpdateDto) GetActivityTypeIdOk() (*string, bool)`

GetActivityTypeIdOk returns a tuple with the ActivityTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTypeId

`func (o *ActivityRecordUpdateDto) SetActivityTypeId(v string)`

SetActivityTypeId sets ActivityTypeId field to given value.

### HasActivityTypeId

`func (o *ActivityRecordUpdateDto) HasActivityTypeId() bool`

HasActivityTypeId returns a boolean if a field has been set.

### SetActivityTypeIdNil

`func (o *ActivityRecordUpdateDto) SetActivityTypeIdNil(b bool)`

 SetActivityTypeIdNil sets the value for ActivityTypeId to be an explicit nil

### UnsetActivityTypeId
`func (o *ActivityRecordUpdateDto) UnsetActivityTypeId()`

UnsetActivityTypeId ensures that no value is present for ActivityTypeId, not even an explicit nil
### GetParentActivityId

`func (o *ActivityRecordUpdateDto) GetParentActivityId() string`

GetParentActivityId returns the ParentActivityId field if non-nil, zero value otherwise.

### GetParentActivityIdOk

`func (o *ActivityRecordUpdateDto) GetParentActivityIdOk() (*string, bool)`

GetParentActivityIdOk returns a tuple with the ParentActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentActivityId

`func (o *ActivityRecordUpdateDto) SetParentActivityId(v string)`

SetParentActivityId sets ParentActivityId field to given value.

### HasParentActivityId

`func (o *ActivityRecordUpdateDto) HasParentActivityId() bool`

HasParentActivityId returns a boolean if a field has been set.

### SetParentActivityIdNil

`func (o *ActivityRecordUpdateDto) SetParentActivityIdNil(b bool)`

 SetParentActivityIdNil sets the value for ParentActivityId to be an explicit nil

### UnsetParentActivityId
`func (o *ActivityRecordUpdateDto) UnsetParentActivityId()`

UnsetParentActivityId ensures that no value is present for ParentActivityId, not even an explicit nil
### GetInChargeEnrollmentId

`func (o *ActivityRecordUpdateDto) GetInChargeEnrollmentId() string`

GetInChargeEnrollmentId returns the InChargeEnrollmentId field if non-nil, zero value otherwise.

### GetInChargeEnrollmentIdOk

`func (o *ActivityRecordUpdateDto) GetInChargeEnrollmentIdOk() (*string, bool)`

GetInChargeEnrollmentIdOk returns a tuple with the InChargeEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInChargeEnrollmentId

`func (o *ActivityRecordUpdateDto) SetInChargeEnrollmentId(v string)`

SetInChargeEnrollmentId sets InChargeEnrollmentId field to given value.

### HasInChargeEnrollmentId

`func (o *ActivityRecordUpdateDto) HasInChargeEnrollmentId() bool`

HasInChargeEnrollmentId returns a boolean if a field has been set.

### SetInChargeEnrollmentIdNil

`func (o *ActivityRecordUpdateDto) SetInChargeEnrollmentIdNil(b bool)`

 SetInChargeEnrollmentIdNil sets the value for InChargeEnrollmentId to be an explicit nil

### UnsetInChargeEnrollmentId
`func (o *ActivityRecordUpdateDto) UnsetInChargeEnrollmentId()`

UnsetInChargeEnrollmentId ensures that no value is present for InChargeEnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


