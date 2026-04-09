# LedgerTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**LedgerClass** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLedgerTypeDto

`func NewLedgerTypeDto() *LedgerTypeDto`

NewLedgerTypeDto instantiates a new LedgerTypeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerTypeDtoWithDefaults

`func NewLedgerTypeDtoWithDefaults() *LedgerTypeDto`

NewLedgerTypeDtoWithDefaults instantiates a new LedgerTypeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LedgerTypeDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedgerTypeDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LedgerTypeDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LedgerTypeDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LedgerTypeDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LedgerTypeDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *LedgerTypeDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LedgerTypeDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LedgerTypeDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LedgerTypeDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *LedgerTypeDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *LedgerTypeDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *LedgerTypeDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LedgerTypeDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LedgerTypeDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LedgerTypeDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LedgerTypeDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LedgerTypeDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLedgerClass

`func (o *LedgerTypeDto) GetLedgerClass() string`

GetLedgerClass returns the LedgerClass field if non-nil, zero value otherwise.

### GetLedgerClassOk

`func (o *LedgerTypeDto) GetLedgerClassOk() (*string, bool)`

GetLedgerClassOk returns a tuple with the LedgerClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerClass

`func (o *LedgerTypeDto) SetLedgerClass(v string)`

SetLedgerClass sets LedgerClass field to given value.

### HasLedgerClass

`func (o *LedgerTypeDto) HasLedgerClass() bool`

HasLedgerClass returns a boolean if a field has been set.

### GetTenantId

`func (o *LedgerTypeDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LedgerTypeDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LedgerTypeDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *LedgerTypeDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *LedgerTypeDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *LedgerTypeDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *LedgerTypeDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *LedgerTypeDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *LedgerTypeDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *LedgerTypeDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *LedgerTypeDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *LedgerTypeDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


