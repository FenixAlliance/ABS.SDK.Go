# MarketingListUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locked** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Purpose** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**Cost** | Pointer to **float64** |  | [optional] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] 
**LastUsedOn** | Pointer to **time.Time** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrolmentId** | Pointer to **NullableString** |  | [optional] 
**MarketingListType** | Pointer to **int32** |  | [optional] 
**MarketingListTarget** | Pointer to **int32** |  | [optional] 

## Methods

### NewMarketingListUpdateDto

`func NewMarketingListUpdateDto() *MarketingListUpdateDto`

NewMarketingListUpdateDto instantiates a new MarketingListUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingListUpdateDtoWithDefaults

`func NewMarketingListUpdateDtoWithDefaults() *MarketingListUpdateDto`

NewMarketingListUpdateDtoWithDefaults instantiates a new MarketingListUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocked

`func (o *MarketingListUpdateDto) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *MarketingListUpdateDto) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *MarketingListUpdateDto) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *MarketingListUpdateDto) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetName

`func (o *MarketingListUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingListUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingListUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingListUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MarketingListUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MarketingListUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPurpose

`func (o *MarketingListUpdateDto) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *MarketingListUpdateDto) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *MarketingListUpdateDto) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *MarketingListUpdateDto) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *MarketingListUpdateDto) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *MarketingListUpdateDto) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetDescription

`func (o *MarketingListUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketingListUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketingListUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketingListUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MarketingListUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MarketingListUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSource

`func (o *MarketingListUpdateDto) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MarketingListUpdateDto) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MarketingListUpdateDto) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MarketingListUpdateDto) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *MarketingListUpdateDto) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *MarketingListUpdateDto) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetCost

`func (o *MarketingListUpdateDto) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *MarketingListUpdateDto) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *MarketingListUpdateDto) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *MarketingListUpdateDto) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetModifiedOn

`func (o *MarketingListUpdateDto) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *MarketingListUpdateDto) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *MarketingListUpdateDto) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *MarketingListUpdateDto) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetLastUsedOn

`func (o *MarketingListUpdateDto) GetLastUsedOn() time.Time`

GetLastUsedOn returns the LastUsedOn field if non-nil, zero value otherwise.

### GetLastUsedOnOk

`func (o *MarketingListUpdateDto) GetLastUsedOnOk() (*time.Time, bool)`

GetLastUsedOnOk returns a tuple with the LastUsedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedOn

`func (o *MarketingListUpdateDto) SetLastUsedOn(v time.Time)`

SetLastUsedOn sets LastUsedOn field to given value.

### HasLastUsedOn

`func (o *MarketingListUpdateDto) HasLastUsedOn() bool`

HasLastUsedOn returns a boolean if a field has been set.

### GetCurrencyId

`func (o *MarketingListUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *MarketingListUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *MarketingListUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *MarketingListUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *MarketingListUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *MarketingListUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTenantId

`func (o *MarketingListUpdateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MarketingListUpdateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MarketingListUpdateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *MarketingListUpdateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *MarketingListUpdateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *MarketingListUpdateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrolmentId

`func (o *MarketingListUpdateDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *MarketingListUpdateDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *MarketingListUpdateDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *MarketingListUpdateDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *MarketingListUpdateDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *MarketingListUpdateDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
### GetMarketingListType

`func (o *MarketingListUpdateDto) GetMarketingListType() int32`

GetMarketingListType returns the MarketingListType field if non-nil, zero value otherwise.

### GetMarketingListTypeOk

`func (o *MarketingListUpdateDto) GetMarketingListTypeOk() (*int32, bool)`

GetMarketingListTypeOk returns a tuple with the MarketingListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingListType

`func (o *MarketingListUpdateDto) SetMarketingListType(v int32)`

SetMarketingListType sets MarketingListType field to given value.

### HasMarketingListType

`func (o *MarketingListUpdateDto) HasMarketingListType() bool`

HasMarketingListType returns a boolean if a field has been set.

### GetMarketingListTarget

`func (o *MarketingListUpdateDto) GetMarketingListTarget() int32`

GetMarketingListTarget returns the MarketingListTarget field if non-nil, zero value otherwise.

### GetMarketingListTargetOk

`func (o *MarketingListUpdateDto) GetMarketingListTargetOk() (*int32, bool)`

GetMarketingListTargetOk returns a tuple with the MarketingListTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingListTarget

`func (o *MarketingListUpdateDto) SetMarketingListTarget(v int32)`

SetMarketingListTarget sets MarketingListTarget field to given value.

### HasMarketingListTarget

`func (o *MarketingListUpdateDto) HasMarketingListTarget() bool`

HasMarketingListTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


