# SuiteLicenseAssignmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Assigned** | Pointer to **bool** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**SuiteLicenseId** | Pointer to **NullableString** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**AvailableSeats** | Pointer to **int32** |  | [optional] 
**TotalSeats** | Pointer to **int32** |  | [optional] 

## Methods

### NewSuiteLicenseAssignmentDto

`func NewSuiteLicenseAssignmentDto() *SuiteLicenseAssignmentDto`

NewSuiteLicenseAssignmentDto instantiates a new SuiteLicenseAssignmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuiteLicenseAssignmentDtoWithDefaults

`func NewSuiteLicenseAssignmentDtoWithDefaults() *SuiteLicenseAssignmentDto`

NewSuiteLicenseAssignmentDtoWithDefaults instantiates a new SuiteLicenseAssignmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SuiteLicenseAssignmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuiteLicenseAssignmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuiteLicenseAssignmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SuiteLicenseAssignmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SuiteLicenseAssignmentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SuiteLicenseAssignmentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SuiteLicenseAssignmentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SuiteLicenseAssignmentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SuiteLicenseAssignmentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SuiteLicenseAssignmentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SuiteLicenseAssignmentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SuiteLicenseAssignmentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *SuiteLicenseAssignmentDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SuiteLicenseAssignmentDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SuiteLicenseAssignmentDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SuiteLicenseAssignmentDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SuiteLicenseAssignmentDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SuiteLicenseAssignmentDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAssigned

`func (o *SuiteLicenseAssignmentDto) GetAssigned() bool`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *SuiteLicenseAssignmentDto) GetAssignedOk() (*bool, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *SuiteLicenseAssignmentDto) SetAssigned(v bool)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *SuiteLicenseAssignmentDto) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.

### GetEnrollmentId

`func (o *SuiteLicenseAssignmentDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *SuiteLicenseAssignmentDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *SuiteLicenseAssignmentDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *SuiteLicenseAssignmentDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *SuiteLicenseAssignmentDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *SuiteLicenseAssignmentDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetSuiteLicenseId

`func (o *SuiteLicenseAssignmentDto) GetSuiteLicenseId() string`

GetSuiteLicenseId returns the SuiteLicenseId field if non-nil, zero value otherwise.

### GetSuiteLicenseIdOk

`func (o *SuiteLicenseAssignmentDto) GetSuiteLicenseIdOk() (*string, bool)`

GetSuiteLicenseIdOk returns a tuple with the SuiteLicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuiteLicenseId

`func (o *SuiteLicenseAssignmentDto) SetSuiteLicenseId(v string)`

SetSuiteLicenseId sets SuiteLicenseId field to given value.

### HasSuiteLicenseId

`func (o *SuiteLicenseAssignmentDto) HasSuiteLicenseId() bool`

HasSuiteLicenseId returns a boolean if a field has been set.

### SetSuiteLicenseIdNil

`func (o *SuiteLicenseAssignmentDto) SetSuiteLicenseIdNil(b bool)`

 SetSuiteLicenseIdNil sets the value for SuiteLicenseId to be an explicit nil

### UnsetSuiteLicenseId
`func (o *SuiteLicenseAssignmentDto) UnsetSuiteLicenseId()`

UnsetSuiteLicenseId ensures that no value is present for SuiteLicenseId, not even an explicit nil
### GetExpirationDate

`func (o *SuiteLicenseAssignmentDto) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *SuiteLicenseAssignmentDto) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *SuiteLicenseAssignmentDto) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *SuiteLicenseAssignmentDto) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetAvailableSeats

`func (o *SuiteLicenseAssignmentDto) GetAvailableSeats() int32`

GetAvailableSeats returns the AvailableSeats field if non-nil, zero value otherwise.

### GetAvailableSeatsOk

`func (o *SuiteLicenseAssignmentDto) GetAvailableSeatsOk() (*int32, bool)`

GetAvailableSeatsOk returns a tuple with the AvailableSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSeats

`func (o *SuiteLicenseAssignmentDto) SetAvailableSeats(v int32)`

SetAvailableSeats sets AvailableSeats field to given value.

### HasAvailableSeats

`func (o *SuiteLicenseAssignmentDto) HasAvailableSeats() bool`

HasAvailableSeats returns a boolean if a field has been set.

### GetTotalSeats

`func (o *SuiteLicenseAssignmentDto) GetTotalSeats() int32`

GetTotalSeats returns the TotalSeats field if non-nil, zero value otherwise.

### GetTotalSeatsOk

`func (o *SuiteLicenseAssignmentDto) GetTotalSeatsOk() (*int32, bool)`

GetTotalSeatsOk returns a tuple with the TotalSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSeats

`func (o *SuiteLicenseAssignmentDto) SetTotalSeats(v int32)`

SetTotalSeats sets TotalSeats field to given value.

### HasTotalSeats

`func (o *SuiteLicenseAssignmentDto) HasTotalSeats() bool`

HasTotalSeats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


