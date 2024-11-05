# ProjectHoursApprovalCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**RequesterContactID** | Pointer to **NullableString** |  | [optional] 
**ApproverContactID** | Pointer to **NullableString** |  | [optional] 
**ProjectPeriodID** | Pointer to **NullableString** |  | [optional] 
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

### GetRequesterContactID

`func (o *ProjectHoursApprovalCreateDto) GetRequesterContactID() string`

GetRequesterContactID returns the RequesterContactID field if non-nil, zero value otherwise.

### GetRequesterContactIDOk

`func (o *ProjectHoursApprovalCreateDto) GetRequesterContactIDOk() (*string, bool)`

GetRequesterContactIDOk returns a tuple with the RequesterContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterContactID

`func (o *ProjectHoursApprovalCreateDto) SetRequesterContactID(v string)`

SetRequesterContactID sets RequesterContactID field to given value.

### HasRequesterContactID

`func (o *ProjectHoursApprovalCreateDto) HasRequesterContactID() bool`

HasRequesterContactID returns a boolean if a field has been set.

### SetRequesterContactIDNil

`func (o *ProjectHoursApprovalCreateDto) SetRequesterContactIDNil(b bool)`

 SetRequesterContactIDNil sets the value for RequesterContactID to be an explicit nil

### UnsetRequesterContactID
`func (o *ProjectHoursApprovalCreateDto) UnsetRequesterContactID()`

UnsetRequesterContactID ensures that no value is present for RequesterContactID, not even an explicit nil
### GetApproverContactID

`func (o *ProjectHoursApprovalCreateDto) GetApproverContactID() string`

GetApproverContactID returns the ApproverContactID field if non-nil, zero value otherwise.

### GetApproverContactIDOk

`func (o *ProjectHoursApprovalCreateDto) GetApproverContactIDOk() (*string, bool)`

GetApproverContactIDOk returns a tuple with the ApproverContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverContactID

`func (o *ProjectHoursApprovalCreateDto) SetApproverContactID(v string)`

SetApproverContactID sets ApproverContactID field to given value.

### HasApproverContactID

`func (o *ProjectHoursApprovalCreateDto) HasApproverContactID() bool`

HasApproverContactID returns a boolean if a field has been set.

### SetApproverContactIDNil

`func (o *ProjectHoursApprovalCreateDto) SetApproverContactIDNil(b bool)`

 SetApproverContactIDNil sets the value for ApproverContactID to be an explicit nil

### UnsetApproverContactID
`func (o *ProjectHoursApprovalCreateDto) UnsetApproverContactID()`

UnsetApproverContactID ensures that no value is present for ApproverContactID, not even an explicit nil
### GetProjectPeriodID

`func (o *ProjectHoursApprovalCreateDto) GetProjectPeriodID() string`

GetProjectPeriodID returns the ProjectPeriodID field if non-nil, zero value otherwise.

### GetProjectPeriodIDOk

`func (o *ProjectHoursApprovalCreateDto) GetProjectPeriodIDOk() (*string, bool)`

GetProjectPeriodIDOk returns a tuple with the ProjectPeriodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPeriodID

`func (o *ProjectHoursApprovalCreateDto) SetProjectPeriodID(v string)`

SetProjectPeriodID sets ProjectPeriodID field to given value.

### HasProjectPeriodID

`func (o *ProjectHoursApprovalCreateDto) HasProjectPeriodID() bool`

HasProjectPeriodID returns a boolean if a field has been set.

### SetProjectPeriodIDNil

`func (o *ProjectHoursApprovalCreateDto) SetProjectPeriodIDNil(b bool)`

 SetProjectPeriodIDNil sets the value for ProjectPeriodID to be an explicit nil

### UnsetProjectPeriodID
`func (o *ProjectHoursApprovalCreateDto) UnsetProjectPeriodID()`

UnsetProjectPeriodID ensures that no value is present for ProjectPeriodID, not even an explicit nil
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


