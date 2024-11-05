# MarketingListCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
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

### NewMarketingListCreateDto

`func NewMarketingListCreateDto() *MarketingListCreateDto`

NewMarketingListCreateDto instantiates a new MarketingListCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingListCreateDtoWithDefaults

`func NewMarketingListCreateDtoWithDefaults() *MarketingListCreateDto`

NewMarketingListCreateDtoWithDefaults instantiates a new MarketingListCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MarketingListCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingListCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingListCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingListCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *MarketingListCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MarketingListCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MarketingListCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MarketingListCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetLocked

`func (o *MarketingListCreateDto) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *MarketingListCreateDto) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *MarketingListCreateDto) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *MarketingListCreateDto) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetName

`func (o *MarketingListCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingListCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingListCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingListCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MarketingListCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MarketingListCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPurpose

`func (o *MarketingListCreateDto) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *MarketingListCreateDto) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *MarketingListCreateDto) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *MarketingListCreateDto) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *MarketingListCreateDto) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *MarketingListCreateDto) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetDescription

`func (o *MarketingListCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketingListCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketingListCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketingListCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MarketingListCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MarketingListCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSource

`func (o *MarketingListCreateDto) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MarketingListCreateDto) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MarketingListCreateDto) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MarketingListCreateDto) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *MarketingListCreateDto) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *MarketingListCreateDto) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetCost

`func (o *MarketingListCreateDto) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *MarketingListCreateDto) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *MarketingListCreateDto) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *MarketingListCreateDto) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetModifiedOn

`func (o *MarketingListCreateDto) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *MarketingListCreateDto) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *MarketingListCreateDto) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *MarketingListCreateDto) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetLastUsedOn

`func (o *MarketingListCreateDto) GetLastUsedOn() time.Time`

GetLastUsedOn returns the LastUsedOn field if non-nil, zero value otherwise.

### GetLastUsedOnOk

`func (o *MarketingListCreateDto) GetLastUsedOnOk() (*time.Time, bool)`

GetLastUsedOnOk returns a tuple with the LastUsedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedOn

`func (o *MarketingListCreateDto) SetLastUsedOn(v time.Time)`

SetLastUsedOn sets LastUsedOn field to given value.

### HasLastUsedOn

`func (o *MarketingListCreateDto) HasLastUsedOn() bool`

HasLastUsedOn returns a boolean if a field has been set.

### GetCurrencyId

`func (o *MarketingListCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *MarketingListCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *MarketingListCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *MarketingListCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *MarketingListCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *MarketingListCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTenantId

`func (o *MarketingListCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MarketingListCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MarketingListCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *MarketingListCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *MarketingListCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *MarketingListCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrolmentId

`func (o *MarketingListCreateDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *MarketingListCreateDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *MarketingListCreateDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *MarketingListCreateDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *MarketingListCreateDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *MarketingListCreateDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
### GetMarketingListType

`func (o *MarketingListCreateDto) GetMarketingListType() int32`

GetMarketingListType returns the MarketingListType field if non-nil, zero value otherwise.

### GetMarketingListTypeOk

`func (o *MarketingListCreateDto) GetMarketingListTypeOk() (*int32, bool)`

GetMarketingListTypeOk returns a tuple with the MarketingListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingListType

`func (o *MarketingListCreateDto) SetMarketingListType(v int32)`

SetMarketingListType sets MarketingListType field to given value.

### HasMarketingListType

`func (o *MarketingListCreateDto) HasMarketingListType() bool`

HasMarketingListType returns a boolean if a field has been set.

### GetMarketingListTarget

`func (o *MarketingListCreateDto) GetMarketingListTarget() int32`

GetMarketingListTarget returns the MarketingListTarget field if non-nil, zero value otherwise.

### GetMarketingListTargetOk

`func (o *MarketingListCreateDto) GetMarketingListTargetOk() (*int32, bool)`

GetMarketingListTargetOk returns a tuple with the MarketingListTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingListTarget

`func (o *MarketingListCreateDto) SetMarketingListTarget(v int32)`

SetMarketingListTarget sets MarketingListTarget field to given value.

### HasMarketingListTarget

`func (o *MarketingListCreateDto) HasMarketingListTarget() bool`

HasMarketingListTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


