# JournalDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**LedgerId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**FiscalYearId** | Pointer to **NullableString** |  | [optional] 
**JournalTypeId** | Pointer to **NullableString** |  | [optional] 
**ParentJournalId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJournalDto

`func NewJournalDto() *JournalDto`

NewJournalDto instantiates a new JournalDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalDtoWithDefaults

`func NewJournalDtoWithDefaults() *JournalDto`

NewJournalDtoWithDefaults instantiates a new JournalDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JournalDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JournalDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JournalDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JournalDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *JournalDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *JournalDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *JournalDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *JournalDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *JournalDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *JournalDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *JournalDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *JournalDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *JournalDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JournalDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JournalDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JournalDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *JournalDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *JournalDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *JournalDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JournalDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JournalDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JournalDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *JournalDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *JournalDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTenantId

`func (o *JournalDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *JournalDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *JournalDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *JournalDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *JournalDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *JournalDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetLedgerId

`func (o *JournalDto) GetLedgerId() string`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *JournalDto) GetLedgerIdOk() (*string, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *JournalDto) SetLedgerId(v string)`

SetLedgerId sets LedgerId field to given value.

### HasLedgerId

`func (o *JournalDto) HasLedgerId() bool`

HasLedgerId returns a boolean if a field has been set.

### SetLedgerIdNil

`func (o *JournalDto) SetLedgerIdNil(b bool)`

 SetLedgerIdNil sets the value for LedgerId to be an explicit nil

### UnsetLedgerId
`func (o *JournalDto) UnsetLedgerId()`

UnsetLedgerId ensures that no value is present for LedgerId, not even an explicit nil
### GetEnrollmentId

`func (o *JournalDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *JournalDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *JournalDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *JournalDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *JournalDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *JournalDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetFiscalYearId

`func (o *JournalDto) GetFiscalYearId() string`

GetFiscalYearId returns the FiscalYearId field if non-nil, zero value otherwise.

### GetFiscalYearIdOk

`func (o *JournalDto) GetFiscalYearIdOk() (*string, bool)`

GetFiscalYearIdOk returns a tuple with the FiscalYearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearId

`func (o *JournalDto) SetFiscalYearId(v string)`

SetFiscalYearId sets FiscalYearId field to given value.

### HasFiscalYearId

`func (o *JournalDto) HasFiscalYearId() bool`

HasFiscalYearId returns a boolean if a field has been set.

### SetFiscalYearIdNil

`func (o *JournalDto) SetFiscalYearIdNil(b bool)`

 SetFiscalYearIdNil sets the value for FiscalYearId to be an explicit nil

### UnsetFiscalYearId
`func (o *JournalDto) UnsetFiscalYearId()`

UnsetFiscalYearId ensures that no value is present for FiscalYearId, not even an explicit nil
### GetJournalTypeId

`func (o *JournalDto) GetJournalTypeId() string`

GetJournalTypeId returns the JournalTypeId field if non-nil, zero value otherwise.

### GetJournalTypeIdOk

`func (o *JournalDto) GetJournalTypeIdOk() (*string, bool)`

GetJournalTypeIdOk returns a tuple with the JournalTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalTypeId

`func (o *JournalDto) SetJournalTypeId(v string)`

SetJournalTypeId sets JournalTypeId field to given value.

### HasJournalTypeId

`func (o *JournalDto) HasJournalTypeId() bool`

HasJournalTypeId returns a boolean if a field has been set.

### SetJournalTypeIdNil

`func (o *JournalDto) SetJournalTypeIdNil(b bool)`

 SetJournalTypeIdNil sets the value for JournalTypeId to be an explicit nil

### UnsetJournalTypeId
`func (o *JournalDto) UnsetJournalTypeId()`

UnsetJournalTypeId ensures that no value is present for JournalTypeId, not even an explicit nil
### GetParentJournalId

`func (o *JournalDto) GetParentJournalId() string`

GetParentJournalId returns the ParentJournalId field if non-nil, zero value otherwise.

### GetParentJournalIdOk

`func (o *JournalDto) GetParentJournalIdOk() (*string, bool)`

GetParentJournalIdOk returns a tuple with the ParentJournalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJournalId

`func (o *JournalDto) SetParentJournalId(v string)`

SetParentJournalId sets ParentJournalId field to given value.

### HasParentJournalId

`func (o *JournalDto) HasParentJournalId() bool`

HasParentJournalId returns a boolean if a field has been set.

### SetParentJournalIdNil

`func (o *JournalDto) SetParentJournalIdNil(b bool)`

 SetParentJournalIdNil sets the value for ParentJournalId to be an explicit nil

### UnsetParentJournalId
`func (o *JournalDto) UnsetParentJournalId()`

UnsetParentJournalId ensures that no value is present for ParentJournalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


