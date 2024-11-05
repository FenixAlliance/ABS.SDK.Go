# MarketingListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
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

### NewMarketingListDto

`func NewMarketingListDto() *MarketingListDto`

NewMarketingListDto instantiates a new MarketingListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingListDtoWithDefaults

`func NewMarketingListDtoWithDefaults() *MarketingListDto`

NewMarketingListDtoWithDefaults instantiates a new MarketingListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MarketingListDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingListDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingListDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MarketingListDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MarketingListDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *MarketingListDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MarketingListDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MarketingListDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MarketingListDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *MarketingListDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *MarketingListDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetLocked

`func (o *MarketingListDto) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *MarketingListDto) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *MarketingListDto) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *MarketingListDto) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetName

`func (o *MarketingListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MarketingListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MarketingListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPurpose

`func (o *MarketingListDto) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *MarketingListDto) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *MarketingListDto) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *MarketingListDto) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *MarketingListDto) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *MarketingListDto) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetDescription

`func (o *MarketingListDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketingListDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketingListDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketingListDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MarketingListDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MarketingListDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSource

`func (o *MarketingListDto) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MarketingListDto) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MarketingListDto) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MarketingListDto) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *MarketingListDto) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *MarketingListDto) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetCost

`func (o *MarketingListDto) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *MarketingListDto) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *MarketingListDto) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *MarketingListDto) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetModifiedOn

`func (o *MarketingListDto) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *MarketingListDto) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *MarketingListDto) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *MarketingListDto) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetLastUsedOn

`func (o *MarketingListDto) GetLastUsedOn() time.Time`

GetLastUsedOn returns the LastUsedOn field if non-nil, zero value otherwise.

### GetLastUsedOnOk

`func (o *MarketingListDto) GetLastUsedOnOk() (*time.Time, bool)`

GetLastUsedOnOk returns a tuple with the LastUsedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedOn

`func (o *MarketingListDto) SetLastUsedOn(v time.Time)`

SetLastUsedOn sets LastUsedOn field to given value.

### HasLastUsedOn

`func (o *MarketingListDto) HasLastUsedOn() bool`

HasLastUsedOn returns a boolean if a field has been set.

### GetCurrencyId

`func (o *MarketingListDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *MarketingListDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *MarketingListDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *MarketingListDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *MarketingListDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *MarketingListDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTenantId

`func (o *MarketingListDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MarketingListDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MarketingListDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *MarketingListDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *MarketingListDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *MarketingListDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrolmentId

`func (o *MarketingListDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *MarketingListDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *MarketingListDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *MarketingListDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *MarketingListDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *MarketingListDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
### GetMarketingListType

`func (o *MarketingListDto) GetMarketingListType() int32`

GetMarketingListType returns the MarketingListType field if non-nil, zero value otherwise.

### GetMarketingListTypeOk

`func (o *MarketingListDto) GetMarketingListTypeOk() (*int32, bool)`

GetMarketingListTypeOk returns a tuple with the MarketingListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingListType

`func (o *MarketingListDto) SetMarketingListType(v int32)`

SetMarketingListType sets MarketingListType field to given value.

### HasMarketingListType

`func (o *MarketingListDto) HasMarketingListType() bool`

HasMarketingListType returns a boolean if a field has been set.

### GetMarketingListTarget

`func (o *MarketingListDto) GetMarketingListTarget() int32`

GetMarketingListTarget returns the MarketingListTarget field if non-nil, zero value otherwise.

### GetMarketingListTargetOk

`func (o *MarketingListDto) GetMarketingListTargetOk() (*int32, bool)`

GetMarketingListTargetOk returns a tuple with the MarketingListTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingListTarget

`func (o *MarketingListDto) SetMarketingListTarget(v int32)`

SetMarketingListTarget sets MarketingListTarget field to given value.

### HasMarketingListTarget

`func (o *MarketingListDto) HasMarketingListTarget() bool`

HasMarketingListTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


