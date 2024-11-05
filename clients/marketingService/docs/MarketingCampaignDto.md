# MarketingCampaignDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Offer** | Pointer to **NullableString** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**ProposedStart** | Pointer to **time.Time** |  | [optional] 
**ProposedEnd** | Pointer to **time.Time** |  | [optional] 
**ActualStart** | Pointer to **time.Time** |  | [optional] 
**ActualEnd** | Pointer to **time.Time** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**AllocatedBudget** | Pointer to **float64** |  | [optional] 
**ActivityCost** | Pointer to **float64** |  | [optional] 
**MiscCost** | Pointer to **float64** |  | [optional] 
**ExpectedResponsePercent** | Pointer to **float64** |  | [optional] 
**MarketingAreaId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrolmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMarketingCampaignDto

`func NewMarketingCampaignDto() *MarketingCampaignDto`

NewMarketingCampaignDto instantiates a new MarketingCampaignDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingCampaignDtoWithDefaults

`func NewMarketingCampaignDtoWithDefaults() *MarketingCampaignDto`

NewMarketingCampaignDtoWithDefaults instantiates a new MarketingCampaignDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MarketingCampaignDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingCampaignDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingCampaignDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingCampaignDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MarketingCampaignDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MarketingCampaignDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *MarketingCampaignDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MarketingCampaignDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MarketingCampaignDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MarketingCampaignDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *MarketingCampaignDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *MarketingCampaignDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *MarketingCampaignDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingCampaignDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingCampaignDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingCampaignDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MarketingCampaignDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MarketingCampaignDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOffer

`func (o *MarketingCampaignDto) GetOffer() string`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *MarketingCampaignDto) GetOfferOk() (*string, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *MarketingCampaignDto) SetOffer(v string)`

SetOffer sets Offer field to given value.

### HasOffer

`func (o *MarketingCampaignDto) HasOffer() bool`

HasOffer returns a boolean if a field has been set.

### SetOfferNil

`func (o *MarketingCampaignDto) SetOfferNil(b bool)`

 SetOfferNil sets the value for Offer to be an explicit nil

### UnsetOffer
`func (o *MarketingCampaignDto) UnsetOffer()`

UnsetOffer ensures that no value is present for Offer, not even an explicit nil
### GetActive

`func (o *MarketingCampaignDto) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MarketingCampaignDto) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MarketingCampaignDto) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *MarketingCampaignDto) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProposedStart

`func (o *MarketingCampaignDto) GetProposedStart() time.Time`

GetProposedStart returns the ProposedStart field if non-nil, zero value otherwise.

### GetProposedStartOk

`func (o *MarketingCampaignDto) GetProposedStartOk() (*time.Time, bool)`

GetProposedStartOk returns a tuple with the ProposedStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedStart

`func (o *MarketingCampaignDto) SetProposedStart(v time.Time)`

SetProposedStart sets ProposedStart field to given value.

### HasProposedStart

`func (o *MarketingCampaignDto) HasProposedStart() bool`

HasProposedStart returns a boolean if a field has been set.

### GetProposedEnd

`func (o *MarketingCampaignDto) GetProposedEnd() time.Time`

GetProposedEnd returns the ProposedEnd field if non-nil, zero value otherwise.

### GetProposedEndOk

`func (o *MarketingCampaignDto) GetProposedEndOk() (*time.Time, bool)`

GetProposedEndOk returns a tuple with the ProposedEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedEnd

`func (o *MarketingCampaignDto) SetProposedEnd(v time.Time)`

SetProposedEnd sets ProposedEnd field to given value.

### HasProposedEnd

`func (o *MarketingCampaignDto) HasProposedEnd() bool`

HasProposedEnd returns a boolean if a field has been set.

### GetActualStart

`func (o *MarketingCampaignDto) GetActualStart() time.Time`

GetActualStart returns the ActualStart field if non-nil, zero value otherwise.

### GetActualStartOk

`func (o *MarketingCampaignDto) GetActualStartOk() (*time.Time, bool)`

GetActualStartOk returns a tuple with the ActualStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualStart

`func (o *MarketingCampaignDto) SetActualStart(v time.Time)`

SetActualStart sets ActualStart field to given value.

### HasActualStart

`func (o *MarketingCampaignDto) HasActualStart() bool`

HasActualStart returns a boolean if a field has been set.

### GetActualEnd

`func (o *MarketingCampaignDto) GetActualEnd() time.Time`

GetActualEnd returns the ActualEnd field if non-nil, zero value otherwise.

### GetActualEndOk

`func (o *MarketingCampaignDto) GetActualEndOk() (*time.Time, bool)`

GetActualEndOk returns a tuple with the ActualEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualEnd

`func (o *MarketingCampaignDto) SetActualEnd(v time.Time)`

SetActualEnd sets ActualEnd field to given value.

### HasActualEnd

`func (o *MarketingCampaignDto) HasActualEnd() bool`

HasActualEnd returns a boolean if a field has been set.

### GetCode

`func (o *MarketingCampaignDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MarketingCampaignDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MarketingCampaignDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *MarketingCampaignDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *MarketingCampaignDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *MarketingCampaignDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetAllocatedBudget

`func (o *MarketingCampaignDto) GetAllocatedBudget() float64`

GetAllocatedBudget returns the AllocatedBudget field if non-nil, zero value otherwise.

### GetAllocatedBudgetOk

`func (o *MarketingCampaignDto) GetAllocatedBudgetOk() (*float64, bool)`

GetAllocatedBudgetOk returns a tuple with the AllocatedBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedBudget

`func (o *MarketingCampaignDto) SetAllocatedBudget(v float64)`

SetAllocatedBudget sets AllocatedBudget field to given value.

### HasAllocatedBudget

`func (o *MarketingCampaignDto) HasAllocatedBudget() bool`

HasAllocatedBudget returns a boolean if a field has been set.

### GetActivityCost

`func (o *MarketingCampaignDto) GetActivityCost() float64`

GetActivityCost returns the ActivityCost field if non-nil, zero value otherwise.

### GetActivityCostOk

`func (o *MarketingCampaignDto) GetActivityCostOk() (*float64, bool)`

GetActivityCostOk returns a tuple with the ActivityCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityCost

`func (o *MarketingCampaignDto) SetActivityCost(v float64)`

SetActivityCost sets ActivityCost field to given value.

### HasActivityCost

`func (o *MarketingCampaignDto) HasActivityCost() bool`

HasActivityCost returns a boolean if a field has been set.

### GetMiscCost

`func (o *MarketingCampaignDto) GetMiscCost() float64`

GetMiscCost returns the MiscCost field if non-nil, zero value otherwise.

### GetMiscCostOk

`func (o *MarketingCampaignDto) GetMiscCostOk() (*float64, bool)`

GetMiscCostOk returns a tuple with the MiscCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiscCost

`func (o *MarketingCampaignDto) SetMiscCost(v float64)`

SetMiscCost sets MiscCost field to given value.

### HasMiscCost

`func (o *MarketingCampaignDto) HasMiscCost() bool`

HasMiscCost returns a boolean if a field has been set.

### GetExpectedResponsePercent

`func (o *MarketingCampaignDto) GetExpectedResponsePercent() float64`

GetExpectedResponsePercent returns the ExpectedResponsePercent field if non-nil, zero value otherwise.

### GetExpectedResponsePercentOk

`func (o *MarketingCampaignDto) GetExpectedResponsePercentOk() (*float64, bool)`

GetExpectedResponsePercentOk returns a tuple with the ExpectedResponsePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResponsePercent

`func (o *MarketingCampaignDto) SetExpectedResponsePercent(v float64)`

SetExpectedResponsePercent sets ExpectedResponsePercent field to given value.

### HasExpectedResponsePercent

`func (o *MarketingCampaignDto) HasExpectedResponsePercent() bool`

HasExpectedResponsePercent returns a boolean if a field has been set.

### GetMarketingAreaId

`func (o *MarketingCampaignDto) GetMarketingAreaId() string`

GetMarketingAreaId returns the MarketingAreaId field if non-nil, zero value otherwise.

### GetMarketingAreaIdOk

`func (o *MarketingCampaignDto) GetMarketingAreaIdOk() (*string, bool)`

GetMarketingAreaIdOk returns a tuple with the MarketingAreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingAreaId

`func (o *MarketingCampaignDto) SetMarketingAreaId(v string)`

SetMarketingAreaId sets MarketingAreaId field to given value.

### HasMarketingAreaId

`func (o *MarketingCampaignDto) HasMarketingAreaId() bool`

HasMarketingAreaId returns a boolean if a field has been set.

### SetMarketingAreaIdNil

`func (o *MarketingCampaignDto) SetMarketingAreaIdNil(b bool)`

 SetMarketingAreaIdNil sets the value for MarketingAreaId to be an explicit nil

### UnsetMarketingAreaId
`func (o *MarketingCampaignDto) UnsetMarketingAreaId()`

UnsetMarketingAreaId ensures that no value is present for MarketingAreaId, not even an explicit nil
### GetCurrencyId

`func (o *MarketingCampaignDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *MarketingCampaignDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *MarketingCampaignDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *MarketingCampaignDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *MarketingCampaignDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *MarketingCampaignDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTenantId

`func (o *MarketingCampaignDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MarketingCampaignDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MarketingCampaignDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *MarketingCampaignDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *MarketingCampaignDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *MarketingCampaignDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrolmentId

`func (o *MarketingCampaignDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *MarketingCampaignDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *MarketingCampaignDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *MarketingCampaignDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *MarketingCampaignDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *MarketingCampaignDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


