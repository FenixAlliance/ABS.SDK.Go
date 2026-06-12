# ActivityRecordCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DueDate** | Pointer to **NullableTime** |  | [optional] 
**ActivityFeedId** | Pointer to **NullableString** |  | [optional] 
**ActivityTypeId** | Pointer to **NullableString** |  | [optional] 
**ParentActivityId** | Pointer to **NullableString** |  | [optional] 
**InChargeEnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewActivityRecordCreateDto

`func NewActivityRecordCreateDto() *ActivityRecordCreateDto`

NewActivityRecordCreateDto instantiates a new ActivityRecordCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityRecordCreateDtoWithDefaults

`func NewActivityRecordCreateDtoWithDefaults() *ActivityRecordCreateDto`

NewActivityRecordCreateDtoWithDefaults instantiates a new ActivityRecordCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityRecordCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityRecordCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityRecordCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityRecordCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ActivityRecordCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActivityRecordCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActivityRecordCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ActivityRecordCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *ActivityRecordCreateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityRecordCreateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityRecordCreateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActivityRecordCreateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ActivityRecordCreateDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ActivityRecordCreateDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *ActivityRecordCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ActivityRecordCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ActivityRecordCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ActivityRecordCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ActivityRecordCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ActivityRecordCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *ActivityRecordCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActivityRecordCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActivityRecordCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActivityRecordCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ActivityRecordCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ActivityRecordCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDueDate

`func (o *ActivityRecordCreateDto) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *ActivityRecordCreateDto) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *ActivityRecordCreateDto) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *ActivityRecordCreateDto) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *ActivityRecordCreateDto) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *ActivityRecordCreateDto) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetActivityFeedId

`func (o *ActivityRecordCreateDto) GetActivityFeedId() string`

GetActivityFeedId returns the ActivityFeedId field if non-nil, zero value otherwise.

### GetActivityFeedIdOk

`func (o *ActivityRecordCreateDto) GetActivityFeedIdOk() (*string, bool)`

GetActivityFeedIdOk returns a tuple with the ActivityFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityFeedId

`func (o *ActivityRecordCreateDto) SetActivityFeedId(v string)`

SetActivityFeedId sets ActivityFeedId field to given value.

### HasActivityFeedId

`func (o *ActivityRecordCreateDto) HasActivityFeedId() bool`

HasActivityFeedId returns a boolean if a field has been set.

### SetActivityFeedIdNil

`func (o *ActivityRecordCreateDto) SetActivityFeedIdNil(b bool)`

 SetActivityFeedIdNil sets the value for ActivityFeedId to be an explicit nil

### UnsetActivityFeedId
`func (o *ActivityRecordCreateDto) UnsetActivityFeedId()`

UnsetActivityFeedId ensures that no value is present for ActivityFeedId, not even an explicit nil
### GetActivityTypeId

`func (o *ActivityRecordCreateDto) GetActivityTypeId() string`

GetActivityTypeId returns the ActivityTypeId field if non-nil, zero value otherwise.

### GetActivityTypeIdOk

`func (o *ActivityRecordCreateDto) GetActivityTypeIdOk() (*string, bool)`

GetActivityTypeIdOk returns a tuple with the ActivityTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTypeId

`func (o *ActivityRecordCreateDto) SetActivityTypeId(v string)`

SetActivityTypeId sets ActivityTypeId field to given value.

### HasActivityTypeId

`func (o *ActivityRecordCreateDto) HasActivityTypeId() bool`

HasActivityTypeId returns a boolean if a field has been set.

### SetActivityTypeIdNil

`func (o *ActivityRecordCreateDto) SetActivityTypeIdNil(b bool)`

 SetActivityTypeIdNil sets the value for ActivityTypeId to be an explicit nil

### UnsetActivityTypeId
`func (o *ActivityRecordCreateDto) UnsetActivityTypeId()`

UnsetActivityTypeId ensures that no value is present for ActivityTypeId, not even an explicit nil
### GetParentActivityId

`func (o *ActivityRecordCreateDto) GetParentActivityId() string`

GetParentActivityId returns the ParentActivityId field if non-nil, zero value otherwise.

### GetParentActivityIdOk

`func (o *ActivityRecordCreateDto) GetParentActivityIdOk() (*string, bool)`

GetParentActivityIdOk returns a tuple with the ParentActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentActivityId

`func (o *ActivityRecordCreateDto) SetParentActivityId(v string)`

SetParentActivityId sets ParentActivityId field to given value.

### HasParentActivityId

`func (o *ActivityRecordCreateDto) HasParentActivityId() bool`

HasParentActivityId returns a boolean if a field has been set.

### SetParentActivityIdNil

`func (o *ActivityRecordCreateDto) SetParentActivityIdNil(b bool)`

 SetParentActivityIdNil sets the value for ParentActivityId to be an explicit nil

### UnsetParentActivityId
`func (o *ActivityRecordCreateDto) UnsetParentActivityId()`

UnsetParentActivityId ensures that no value is present for ParentActivityId, not even an explicit nil
### GetInChargeEnrollmentId

`func (o *ActivityRecordCreateDto) GetInChargeEnrollmentId() string`

GetInChargeEnrollmentId returns the InChargeEnrollmentId field if non-nil, zero value otherwise.

### GetInChargeEnrollmentIdOk

`func (o *ActivityRecordCreateDto) GetInChargeEnrollmentIdOk() (*string, bool)`

GetInChargeEnrollmentIdOk returns a tuple with the InChargeEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInChargeEnrollmentId

`func (o *ActivityRecordCreateDto) SetInChargeEnrollmentId(v string)`

SetInChargeEnrollmentId sets InChargeEnrollmentId field to given value.

### HasInChargeEnrollmentId

`func (o *ActivityRecordCreateDto) HasInChargeEnrollmentId() bool`

HasInChargeEnrollmentId returns a boolean if a field has been set.

### SetInChargeEnrollmentIdNil

`func (o *ActivityRecordCreateDto) SetInChargeEnrollmentIdNil(b bool)`

 SetInChargeEnrollmentIdNil sets the value for InChargeEnrollmentId to be an explicit nil

### UnsetInChargeEnrollmentId
`func (o *ActivityRecordCreateDto) UnsetInChargeEnrollmentId()`

UnsetInChargeEnrollmentId ensures that no value is present for InChargeEnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


