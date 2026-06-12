# ActivityRecordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Completed** | Pointer to **bool** |  | [optional] 
**DueDate** | Pointer to **time.Time** |  | [optional] 
**ActivityFeedId** | Pointer to **NullableString** |  | [optional] 
**ActivityTypeId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 
**ParentActivityId** | Pointer to **NullableString** |  | [optional] 
**InChargeEnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewActivityRecordDto

`func NewActivityRecordDto() *ActivityRecordDto`

NewActivityRecordDto instantiates a new ActivityRecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityRecordDtoWithDefaults

`func NewActivityRecordDtoWithDefaults() *ActivityRecordDto`

NewActivityRecordDtoWithDefaults instantiates a new ActivityRecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityRecordDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityRecordDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityRecordDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityRecordDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ActivityRecordDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ActivityRecordDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ActivityRecordDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActivityRecordDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActivityRecordDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ActivityRecordDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ActivityRecordDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ActivityRecordDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetType

`func (o *ActivityRecordDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityRecordDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityRecordDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActivityRecordDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ActivityRecordDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ActivityRecordDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *ActivityRecordDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ActivityRecordDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ActivityRecordDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ActivityRecordDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ActivityRecordDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ActivityRecordDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *ActivityRecordDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActivityRecordDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActivityRecordDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActivityRecordDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ActivityRecordDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ActivityRecordDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCompleted

`func (o *ActivityRecordDto) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *ActivityRecordDto) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *ActivityRecordDto) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *ActivityRecordDto) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetDueDate

`func (o *ActivityRecordDto) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *ActivityRecordDto) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *ActivityRecordDto) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *ActivityRecordDto) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetActivityFeedId

`func (o *ActivityRecordDto) GetActivityFeedId() string`

GetActivityFeedId returns the ActivityFeedId field if non-nil, zero value otherwise.

### GetActivityFeedIdOk

`func (o *ActivityRecordDto) GetActivityFeedIdOk() (*string, bool)`

GetActivityFeedIdOk returns a tuple with the ActivityFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityFeedId

`func (o *ActivityRecordDto) SetActivityFeedId(v string)`

SetActivityFeedId sets ActivityFeedId field to given value.

### HasActivityFeedId

`func (o *ActivityRecordDto) HasActivityFeedId() bool`

HasActivityFeedId returns a boolean if a field has been set.

### SetActivityFeedIdNil

`func (o *ActivityRecordDto) SetActivityFeedIdNil(b bool)`

 SetActivityFeedIdNil sets the value for ActivityFeedId to be an explicit nil

### UnsetActivityFeedId
`func (o *ActivityRecordDto) UnsetActivityFeedId()`

UnsetActivityFeedId ensures that no value is present for ActivityFeedId, not even an explicit nil
### GetActivityTypeId

`func (o *ActivityRecordDto) GetActivityTypeId() string`

GetActivityTypeId returns the ActivityTypeId field if non-nil, zero value otherwise.

### GetActivityTypeIdOk

`func (o *ActivityRecordDto) GetActivityTypeIdOk() (*string, bool)`

GetActivityTypeIdOk returns a tuple with the ActivityTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTypeId

`func (o *ActivityRecordDto) SetActivityTypeId(v string)`

SetActivityTypeId sets ActivityTypeId field to given value.

### HasActivityTypeId

`func (o *ActivityRecordDto) HasActivityTypeId() bool`

HasActivityTypeId returns a boolean if a field has been set.

### SetActivityTypeIdNil

`func (o *ActivityRecordDto) SetActivityTypeIdNil(b bool)`

 SetActivityTypeIdNil sets the value for ActivityTypeId to be an explicit nil

### UnsetActivityTypeId
`func (o *ActivityRecordDto) UnsetActivityTypeId()`

UnsetActivityTypeId ensures that no value is present for ActivityTypeId, not even an explicit nil
### GetTenantId

`func (o *ActivityRecordDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ActivityRecordDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ActivityRecordDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ActivityRecordDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ActivityRecordDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ActivityRecordDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *ActivityRecordDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ActivityRecordDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ActivityRecordDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ActivityRecordDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ActivityRecordDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ActivityRecordDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetSocialProfileId

`func (o *ActivityRecordDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *ActivityRecordDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *ActivityRecordDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *ActivityRecordDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *ActivityRecordDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *ActivityRecordDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetParentActivityId

`func (o *ActivityRecordDto) GetParentActivityId() string`

GetParentActivityId returns the ParentActivityId field if non-nil, zero value otherwise.

### GetParentActivityIdOk

`func (o *ActivityRecordDto) GetParentActivityIdOk() (*string, bool)`

GetParentActivityIdOk returns a tuple with the ParentActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentActivityId

`func (o *ActivityRecordDto) SetParentActivityId(v string)`

SetParentActivityId sets ParentActivityId field to given value.

### HasParentActivityId

`func (o *ActivityRecordDto) HasParentActivityId() bool`

HasParentActivityId returns a boolean if a field has been set.

### SetParentActivityIdNil

`func (o *ActivityRecordDto) SetParentActivityIdNil(b bool)`

 SetParentActivityIdNil sets the value for ParentActivityId to be an explicit nil

### UnsetParentActivityId
`func (o *ActivityRecordDto) UnsetParentActivityId()`

UnsetParentActivityId ensures that no value is present for ParentActivityId, not even an explicit nil
### GetInChargeEnrollmentId

`func (o *ActivityRecordDto) GetInChargeEnrollmentId() string`

GetInChargeEnrollmentId returns the InChargeEnrollmentId field if non-nil, zero value otherwise.

### GetInChargeEnrollmentIdOk

`func (o *ActivityRecordDto) GetInChargeEnrollmentIdOk() (*string, bool)`

GetInChargeEnrollmentIdOk returns a tuple with the InChargeEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInChargeEnrollmentId

`func (o *ActivityRecordDto) SetInChargeEnrollmentId(v string)`

SetInChargeEnrollmentId sets InChargeEnrollmentId field to given value.

### HasInChargeEnrollmentId

`func (o *ActivityRecordDto) HasInChargeEnrollmentId() bool`

HasInChargeEnrollmentId returns a boolean if a field has been set.

### SetInChargeEnrollmentIdNil

`func (o *ActivityRecordDto) SetInChargeEnrollmentIdNil(b bool)`

 SetInChargeEnrollmentIdNil sets the value for InChargeEnrollmentId to be an explicit nil

### UnsetInChargeEnrollmentId
`func (o *ActivityRecordDto) UnsetInChargeEnrollmentId()`

UnsetInChargeEnrollmentId ensures that no value is present for InChargeEnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


