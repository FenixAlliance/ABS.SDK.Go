# ShareClassDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ForexRates** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShareClassDto

`func NewShareClassDto() *ShareClassDto`

NewShareClassDto instantiates a new ShareClassDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareClassDtoWithDefaults

`func NewShareClassDtoWithDefaults() *ShareClassDto`

NewShareClassDtoWithDefaults instantiates a new ShareClassDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShareClassDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShareClassDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShareClassDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShareClassDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ShareClassDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ShareClassDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ShareClassDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShareClassDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShareClassDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ShareClassDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ShareClassDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ShareClassDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *ShareClassDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShareClassDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShareClassDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShareClassDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ShareClassDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ShareClassDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *ShareClassDto) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ShareClassDto) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ShareClassDto) SetValue(v bool)`

SetValue sets Value field to given value.

### HasValue

`func (o *ShareClassDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDescription

`func (o *ShareClassDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShareClassDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShareClassDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShareClassDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ShareClassDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ShareClassDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetForexRates

`func (o *ShareClassDto) GetForexRates() string`

GetForexRates returns the ForexRates field if non-nil, zero value otherwise.

### GetForexRatesOk

`func (o *ShareClassDto) GetForexRatesOk() (*string, bool)`

GetForexRatesOk returns a tuple with the ForexRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRates

`func (o *ShareClassDto) SetForexRates(v string)`

SetForexRates sets ForexRates field to given value.

### HasForexRates

`func (o *ShareClassDto) HasForexRates() bool`

HasForexRates returns a boolean if a field has been set.

### SetForexRatesNil

`func (o *ShareClassDto) SetForexRatesNil(b bool)`

 SetForexRatesNil sets the value for ForexRates to be an explicit nil

### UnsetForexRates
`func (o *ShareClassDto) UnsetForexRates()`

UnsetForexRates ensures that no value is present for ForexRates, not even an explicit nil
### GetCurrencyId

`func (o *ShareClassDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ShareClassDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ShareClassDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ShareClassDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ShareClassDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ShareClassDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTenantId

`func (o *ShareClassDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ShareClassDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ShareClassDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ShareClassDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ShareClassDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ShareClassDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *ShareClassDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ShareClassDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ShareClassDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ShareClassDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ShareClassDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ShareClassDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


