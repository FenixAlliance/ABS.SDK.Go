# JobOfferFieldRecordCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**JobFieldId** | **string** |  | 
**JobOfferId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJobOfferFieldRecordCreateDto

`func NewJobOfferFieldRecordCreateDto(jobFieldId string, ) *JobOfferFieldRecordCreateDto`

NewJobOfferFieldRecordCreateDto instantiates a new JobOfferFieldRecordCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobOfferFieldRecordCreateDtoWithDefaults

`func NewJobOfferFieldRecordCreateDtoWithDefaults() *JobOfferFieldRecordCreateDto`

NewJobOfferFieldRecordCreateDtoWithDefaults instantiates a new JobOfferFieldRecordCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobOfferFieldRecordCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobOfferFieldRecordCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobOfferFieldRecordCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobOfferFieldRecordCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *JobOfferFieldRecordCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *JobOfferFieldRecordCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *JobOfferFieldRecordCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *JobOfferFieldRecordCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetJobFieldId

`func (o *JobOfferFieldRecordCreateDto) GetJobFieldId() string`

GetJobFieldId returns the JobFieldId field if non-nil, zero value otherwise.

### GetJobFieldIdOk

`func (o *JobOfferFieldRecordCreateDto) GetJobFieldIdOk() (*string, bool)`

GetJobFieldIdOk returns a tuple with the JobFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFieldId

`func (o *JobOfferFieldRecordCreateDto) SetJobFieldId(v string)`

SetJobFieldId sets JobFieldId field to given value.


### GetJobOfferId

`func (o *JobOfferFieldRecordCreateDto) GetJobOfferId() string`

GetJobOfferId returns the JobOfferId field if non-nil, zero value otherwise.

### GetJobOfferIdOk

`func (o *JobOfferFieldRecordCreateDto) GetJobOfferIdOk() (*string, bool)`

GetJobOfferIdOk returns a tuple with the JobOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobOfferId

`func (o *JobOfferFieldRecordCreateDto) SetJobOfferId(v string)`

SetJobOfferId sets JobOfferId field to given value.

### HasJobOfferId

`func (o *JobOfferFieldRecordCreateDto) HasJobOfferId() bool`

HasJobOfferId returns a boolean if a field has been set.

### SetJobOfferIdNil

`func (o *JobOfferFieldRecordCreateDto) SetJobOfferIdNil(b bool)`

 SetJobOfferIdNil sets the value for JobOfferId to be an explicit nil

### UnsetJobOfferId
`func (o *JobOfferFieldRecordCreateDto) UnsetJobOfferId()`

UnsetJobOfferId ensures that no value is present for JobOfferId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


