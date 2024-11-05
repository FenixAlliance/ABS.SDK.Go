# SuiteLicenseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**LicenseString** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**AvailableSeats** | Pointer to **int32** |  | [optional] 
**TotalSeats** | Pointer to **int32** |  | [optional] 

## Methods

### NewSuiteLicenseDto

`func NewSuiteLicenseDto() *SuiteLicenseDto`

NewSuiteLicenseDto instantiates a new SuiteLicenseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuiteLicenseDtoWithDefaults

`func NewSuiteLicenseDtoWithDefaults() *SuiteLicenseDto`

NewSuiteLicenseDtoWithDefaults instantiates a new SuiteLicenseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SuiteLicenseDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuiteLicenseDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuiteLicenseDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SuiteLicenseDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SuiteLicenseDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SuiteLicenseDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SuiteLicenseDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SuiteLicenseDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SuiteLicenseDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SuiteLicenseDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SuiteLicenseDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SuiteLicenseDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *SuiteLicenseDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SuiteLicenseDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SuiteLicenseDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SuiteLicenseDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SuiteLicenseDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SuiteLicenseDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetLicenseString

`func (o *SuiteLicenseDto) GetLicenseString() string`

GetLicenseString returns the LicenseString field if non-nil, zero value otherwise.

### GetLicenseStringOk

`func (o *SuiteLicenseDto) GetLicenseStringOk() (*string, bool)`

GetLicenseStringOk returns a tuple with the LicenseString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseString

`func (o *SuiteLicenseDto) SetLicenseString(v string)`

SetLicenseString sets LicenseString field to given value.

### HasLicenseString

`func (o *SuiteLicenseDto) HasLicenseString() bool`

HasLicenseString returns a boolean if a field has been set.

### SetLicenseStringNil

`func (o *SuiteLicenseDto) SetLicenseStringNil(b bool)`

 SetLicenseStringNil sets the value for LicenseString to be an explicit nil

### UnsetLicenseString
`func (o *SuiteLicenseDto) UnsetLicenseString()`

UnsetLicenseString ensures that no value is present for LicenseString, not even an explicit nil
### GetEnrollmentId

`func (o *SuiteLicenseDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *SuiteLicenseDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *SuiteLicenseDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *SuiteLicenseDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *SuiteLicenseDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *SuiteLicenseDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetExpirationDate

`func (o *SuiteLicenseDto) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *SuiteLicenseDto) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *SuiteLicenseDto) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *SuiteLicenseDto) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetAvailableSeats

`func (o *SuiteLicenseDto) GetAvailableSeats() int32`

GetAvailableSeats returns the AvailableSeats field if non-nil, zero value otherwise.

### GetAvailableSeatsOk

`func (o *SuiteLicenseDto) GetAvailableSeatsOk() (*int32, bool)`

GetAvailableSeatsOk returns a tuple with the AvailableSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSeats

`func (o *SuiteLicenseDto) SetAvailableSeats(v int32)`

SetAvailableSeats sets AvailableSeats field to given value.

### HasAvailableSeats

`func (o *SuiteLicenseDto) HasAvailableSeats() bool`

HasAvailableSeats returns a boolean if a field has been set.

### GetTotalSeats

`func (o *SuiteLicenseDto) GetTotalSeats() int32`

GetTotalSeats returns the TotalSeats field if non-nil, zero value otherwise.

### GetTotalSeatsOk

`func (o *SuiteLicenseDto) GetTotalSeatsOk() (*int32, bool)`

GetTotalSeatsOk returns a tuple with the TotalSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSeats

`func (o *SuiteLicenseDto) SetTotalSeats(v int32)`

SetTotalSeats sets TotalSeats field to given value.

### HasTotalSeats

`func (o *SuiteLicenseDto) HasTotalSeats() bool`

HasTotalSeats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


