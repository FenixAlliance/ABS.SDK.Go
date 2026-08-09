# JobOfferFieldRecordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**JobFieldId** | Pointer to **NullableString** |  | [optional] 
**JobOfferId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJobOfferFieldRecordDto

`func NewJobOfferFieldRecordDto() *JobOfferFieldRecordDto`

NewJobOfferFieldRecordDto instantiates a new JobOfferFieldRecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobOfferFieldRecordDtoWithDefaults

`func NewJobOfferFieldRecordDtoWithDefaults() *JobOfferFieldRecordDto`

NewJobOfferFieldRecordDtoWithDefaults instantiates a new JobOfferFieldRecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobOfferFieldRecordDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobOfferFieldRecordDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobOfferFieldRecordDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobOfferFieldRecordDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *JobOfferFieldRecordDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *JobOfferFieldRecordDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *JobOfferFieldRecordDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *JobOfferFieldRecordDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *JobOfferFieldRecordDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *JobOfferFieldRecordDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *JobOfferFieldRecordDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *JobOfferFieldRecordDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetJobFieldId

`func (o *JobOfferFieldRecordDto) GetJobFieldId() string`

GetJobFieldId returns the JobFieldId field if non-nil, zero value otherwise.

### GetJobFieldIdOk

`func (o *JobOfferFieldRecordDto) GetJobFieldIdOk() (*string, bool)`

GetJobFieldIdOk returns a tuple with the JobFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFieldId

`func (o *JobOfferFieldRecordDto) SetJobFieldId(v string)`

SetJobFieldId sets JobFieldId field to given value.

### HasJobFieldId

`func (o *JobOfferFieldRecordDto) HasJobFieldId() bool`

HasJobFieldId returns a boolean if a field has been set.

### SetJobFieldIdNil

`func (o *JobOfferFieldRecordDto) SetJobFieldIdNil(b bool)`

 SetJobFieldIdNil sets the value for JobFieldId to be an explicit nil

### UnsetJobFieldId
`func (o *JobOfferFieldRecordDto) UnsetJobFieldId()`

UnsetJobFieldId ensures that no value is present for JobFieldId, not even an explicit nil
### GetJobOfferId

`func (o *JobOfferFieldRecordDto) GetJobOfferId() string`

GetJobOfferId returns the JobOfferId field if non-nil, zero value otherwise.

### GetJobOfferIdOk

`func (o *JobOfferFieldRecordDto) GetJobOfferIdOk() (*string, bool)`

GetJobOfferIdOk returns a tuple with the JobOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobOfferId

`func (o *JobOfferFieldRecordDto) SetJobOfferId(v string)`

SetJobOfferId sets JobOfferId field to given value.

### HasJobOfferId

`func (o *JobOfferFieldRecordDto) HasJobOfferId() bool`

HasJobOfferId returns a boolean if a field has been set.

### SetJobOfferIdNil

`func (o *JobOfferFieldRecordDto) SetJobOfferIdNil(b bool)`

 SetJobOfferIdNil sets the value for JobOfferId to be an explicit nil

### UnsetJobOfferId
`func (o *JobOfferFieldRecordDto) UnsetJobOfferId()`

UnsetJobOfferId ensures that no value is present for JobOfferId, not even an explicit nil
### GetTenantId

`func (o *JobOfferFieldRecordDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *JobOfferFieldRecordDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *JobOfferFieldRecordDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *JobOfferFieldRecordDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *JobOfferFieldRecordDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *JobOfferFieldRecordDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *JobOfferFieldRecordDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *JobOfferFieldRecordDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *JobOfferFieldRecordDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *JobOfferFieldRecordDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *JobOfferFieldRecordDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *JobOfferFieldRecordDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


