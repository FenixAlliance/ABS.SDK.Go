# AssetValueAmendDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**AssetId** | Pointer to **NullableString** |  | [optional] 
**AssetName** | Pointer to **NullableString** |  | [optional] 
**PreviousValue** | Pointer to **float64** |  | [optional] 
**NewValue** | Pointer to **float64** |  | [optional] 
**AmendmentAmount** | Pointer to **float64** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**AmendmentDate** | Pointer to **time.Time** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAssetValueAmendDto

`func NewAssetValueAmendDto() *AssetValueAmendDto`

NewAssetValueAmendDto instantiates a new AssetValueAmendDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetValueAmendDtoWithDefaults

`func NewAssetValueAmendDtoWithDefaults() *AssetValueAmendDto`

NewAssetValueAmendDtoWithDefaults instantiates a new AssetValueAmendDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssetValueAmendDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetValueAmendDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetValueAmendDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetValueAmendDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AssetValueAmendDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AssetValueAmendDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *AssetValueAmendDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AssetValueAmendDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AssetValueAmendDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AssetValueAmendDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTenantId

`func (o *AssetValueAmendDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AssetValueAmendDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AssetValueAmendDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AssetValueAmendDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AssetValueAmendDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AssetValueAmendDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *AssetValueAmendDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *AssetValueAmendDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *AssetValueAmendDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *AssetValueAmendDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *AssetValueAmendDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *AssetValueAmendDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetAssetId

`func (o *AssetValueAmendDto) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetValueAmendDto) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetValueAmendDto) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AssetValueAmendDto) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### SetAssetIdNil

`func (o *AssetValueAmendDto) SetAssetIdNil(b bool)`

 SetAssetIdNil sets the value for AssetId to be an explicit nil

### UnsetAssetId
`func (o *AssetValueAmendDto) UnsetAssetId()`

UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
### GetAssetName

`func (o *AssetValueAmendDto) GetAssetName() string`

GetAssetName returns the AssetName field if non-nil, zero value otherwise.

### GetAssetNameOk

`func (o *AssetValueAmendDto) GetAssetNameOk() (*string, bool)`

GetAssetNameOk returns a tuple with the AssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetName

`func (o *AssetValueAmendDto) SetAssetName(v string)`

SetAssetName sets AssetName field to given value.

### HasAssetName

`func (o *AssetValueAmendDto) HasAssetName() bool`

HasAssetName returns a boolean if a field has been set.

### SetAssetNameNil

`func (o *AssetValueAmendDto) SetAssetNameNil(b bool)`

 SetAssetNameNil sets the value for AssetName to be an explicit nil

### UnsetAssetName
`func (o *AssetValueAmendDto) UnsetAssetName()`

UnsetAssetName ensures that no value is present for AssetName, not even an explicit nil
### GetPreviousValue

`func (o *AssetValueAmendDto) GetPreviousValue() float64`

GetPreviousValue returns the PreviousValue field if non-nil, zero value otherwise.

### GetPreviousValueOk

`func (o *AssetValueAmendDto) GetPreviousValueOk() (*float64, bool)`

GetPreviousValueOk returns a tuple with the PreviousValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousValue

`func (o *AssetValueAmendDto) SetPreviousValue(v float64)`

SetPreviousValue sets PreviousValue field to given value.

### HasPreviousValue

`func (o *AssetValueAmendDto) HasPreviousValue() bool`

HasPreviousValue returns a boolean if a field has been set.

### GetNewValue

`func (o *AssetValueAmendDto) GetNewValue() float64`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *AssetValueAmendDto) GetNewValueOk() (*float64, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *AssetValueAmendDto) SetNewValue(v float64)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *AssetValueAmendDto) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetAmendmentAmount

`func (o *AssetValueAmendDto) GetAmendmentAmount() float64`

GetAmendmentAmount returns the AmendmentAmount field if non-nil, zero value otherwise.

### GetAmendmentAmountOk

`func (o *AssetValueAmendDto) GetAmendmentAmountOk() (*float64, bool)`

GetAmendmentAmountOk returns a tuple with the AmendmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendmentAmount

`func (o *AssetValueAmendDto) SetAmendmentAmount(v float64)`

SetAmendmentAmount sets AmendmentAmount field to given value.

### HasAmendmentAmount

`func (o *AssetValueAmendDto) HasAmendmentAmount() bool`

HasAmendmentAmount returns a boolean if a field has been set.

### GetReason

`func (o *AssetValueAmendDto) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AssetValueAmendDto) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AssetValueAmendDto) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AssetValueAmendDto) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *AssetValueAmendDto) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *AssetValueAmendDto) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetAmendmentDate

`func (o *AssetValueAmendDto) GetAmendmentDate() time.Time`

GetAmendmentDate returns the AmendmentDate field if non-nil, zero value otherwise.

### GetAmendmentDateOk

`func (o *AssetValueAmendDto) GetAmendmentDateOk() (*time.Time, bool)`

GetAmendmentDateOk returns a tuple with the AmendmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendmentDate

`func (o *AssetValueAmendDto) SetAmendmentDate(v time.Time)`

SetAmendmentDate sets AmendmentDate field to given value.

### HasAmendmentDate

`func (o *AssetValueAmendDto) HasAmendmentDate() bool`

HasAmendmentDate returns a boolean if a field has been set.

### GetCurrencyId

`func (o *AssetValueAmendDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AssetValueAmendDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AssetValueAmendDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *AssetValueAmendDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *AssetValueAmendDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *AssetValueAmendDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


