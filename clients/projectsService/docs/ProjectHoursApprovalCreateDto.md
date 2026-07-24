# ProjectHoursApprovalCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**RequesterContactId** | Pointer to **NullableString** |  | [optional] 
**ApproverContactId** | Pointer to **NullableString** |  | [optional] 
**ProjectPeriodId** | Pointer to **NullableString** |  | [optional] 
**Comments** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProjectHoursApprovalCreateDto

`func NewProjectHoursApprovalCreateDto() *ProjectHoursApprovalCreateDto`

NewProjectHoursApprovalCreateDto instantiates a new ProjectHoursApprovalCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectHoursApprovalCreateDtoWithDefaults

`func NewProjectHoursApprovalCreateDtoWithDefaults() *ProjectHoursApprovalCreateDto`

NewProjectHoursApprovalCreateDtoWithDefaults instantiates a new ProjectHoursApprovalCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectHoursApprovalCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectHoursApprovalCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectHoursApprovalCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectHoursApprovalCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ProjectHoursApprovalCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProjectHoursApprovalCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProjectHoursApprovalCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProjectHoursApprovalCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetRequesterContactId

`func (o *ProjectHoursApprovalCreateDto) GetRequesterContactId() string`

GetRequesterContactId returns the RequesterContactId field if non-nil, zero value otherwise.

### GetRequesterContactIdOk

`func (o *ProjectHoursApprovalCreateDto) GetRequesterContactIdOk() (*string, bool)`

GetRequesterContactIdOk returns a tuple with the RequesterContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterContactId

`func (o *ProjectHoursApprovalCreateDto) SetRequesterContactId(v string)`

SetRequesterContactId sets RequesterContactId field to given value.

### HasRequesterContactId

`func (o *ProjectHoursApprovalCreateDto) HasRequesterContactId() bool`

HasRequesterContactId returns a boolean if a field has been set.

### SetRequesterContactIdNil

`func (o *ProjectHoursApprovalCreateDto) SetRequesterContactIdNil(b bool)`

 SetRequesterContactIdNil sets the value for RequesterContactId to be an explicit nil

### UnsetRequesterContactId
`func (o *ProjectHoursApprovalCreateDto) UnsetRequesterContactId()`

UnsetRequesterContactId ensures that no value is present for RequesterContactId, not even an explicit nil
### GetApproverContactId

`func (o *ProjectHoursApprovalCreateDto) GetApproverContactId() string`

GetApproverContactId returns the ApproverContactId field if non-nil, zero value otherwise.

### GetApproverContactIdOk

`func (o *ProjectHoursApprovalCreateDto) GetApproverContactIdOk() (*string, bool)`

GetApproverContactIdOk returns a tuple with the ApproverContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverContactId

`func (o *ProjectHoursApprovalCreateDto) SetApproverContactId(v string)`

SetApproverContactId sets ApproverContactId field to given value.

### HasApproverContactId

`func (o *ProjectHoursApprovalCreateDto) HasApproverContactId() bool`

HasApproverContactId returns a boolean if a field has been set.

### SetApproverContactIdNil

`func (o *ProjectHoursApprovalCreateDto) SetApproverContactIdNil(b bool)`

 SetApproverContactIdNil sets the value for ApproverContactId to be an explicit nil

### UnsetApproverContactId
`func (o *ProjectHoursApprovalCreateDto) UnsetApproverContactId()`

UnsetApproverContactId ensures that no value is present for ApproverContactId, not even an explicit nil
### GetProjectPeriodId

`func (o *ProjectHoursApprovalCreateDto) GetProjectPeriodId() string`

GetProjectPeriodId returns the ProjectPeriodId field if non-nil, zero value otherwise.

### GetProjectPeriodIdOk

`func (o *ProjectHoursApprovalCreateDto) GetProjectPeriodIdOk() (*string, bool)`

GetProjectPeriodIdOk returns a tuple with the ProjectPeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPeriodId

`func (o *ProjectHoursApprovalCreateDto) SetProjectPeriodId(v string)`

SetProjectPeriodId sets ProjectPeriodId field to given value.

### HasProjectPeriodId

`func (o *ProjectHoursApprovalCreateDto) HasProjectPeriodId() bool`

HasProjectPeriodId returns a boolean if a field has been set.

### SetProjectPeriodIdNil

`func (o *ProjectHoursApprovalCreateDto) SetProjectPeriodIdNil(b bool)`

 SetProjectPeriodIdNil sets the value for ProjectPeriodId to be an explicit nil

### UnsetProjectPeriodId
`func (o *ProjectHoursApprovalCreateDto) UnsetProjectPeriodId()`

UnsetProjectPeriodId ensures that no value is present for ProjectPeriodId, not even an explicit nil
### GetComments

`func (o *ProjectHoursApprovalCreateDto) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ProjectHoursApprovalCreateDto) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ProjectHoursApprovalCreateDto) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ProjectHoursApprovalCreateDto) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *ProjectHoursApprovalCreateDto) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *ProjectHoursApprovalCreateDto) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


