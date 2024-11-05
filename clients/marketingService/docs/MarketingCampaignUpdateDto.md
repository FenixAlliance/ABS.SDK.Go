# MarketingCampaignUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewMarketingCampaignUpdateDto

`func NewMarketingCampaignUpdateDto() *MarketingCampaignUpdateDto`

NewMarketingCampaignUpdateDto instantiates a new MarketingCampaignUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingCampaignUpdateDtoWithDefaults

`func NewMarketingCampaignUpdateDtoWithDefaults() *MarketingCampaignUpdateDto`

NewMarketingCampaignUpdateDtoWithDefaults instantiates a new MarketingCampaignUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MarketingCampaignUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingCampaignUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingCampaignUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingCampaignUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MarketingCampaignUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MarketingCampaignUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOffer

`func (o *MarketingCampaignUpdateDto) GetOffer() string`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *MarketingCampaignUpdateDto) GetOfferOk() (*string, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *MarketingCampaignUpdateDto) SetOffer(v string)`

SetOffer sets Offer field to given value.

### HasOffer

`func (o *MarketingCampaignUpdateDto) HasOffer() bool`

HasOffer returns a boolean if a field has been set.

### SetOfferNil

`func (o *MarketingCampaignUpdateDto) SetOfferNil(b bool)`

 SetOfferNil sets the value for Offer to be an explicit nil

### UnsetOffer
`func (o *MarketingCampaignUpdateDto) UnsetOffer()`

UnsetOffer ensures that no value is present for Offer, not even an explicit nil
### GetActive

`func (o *MarketingCampaignUpdateDto) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MarketingCampaignUpdateDto) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MarketingCampaignUpdateDto) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *MarketingCampaignUpdateDto) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProposedStart

`func (o *MarketingCampaignUpdateDto) GetProposedStart() time.Time`

GetProposedStart returns the ProposedStart field if non-nil, zero value otherwise.

### GetProposedStartOk

`func (o *MarketingCampaignUpdateDto) GetProposedStartOk() (*time.Time, bool)`

GetProposedStartOk returns a tuple with the ProposedStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedStart

`func (o *MarketingCampaignUpdateDto) SetProposedStart(v time.Time)`

SetProposedStart sets ProposedStart field to given value.

### HasProposedStart

`func (o *MarketingCampaignUpdateDto) HasProposedStart() bool`

HasProposedStart returns a boolean if a field has been set.

### GetProposedEnd

`func (o *MarketingCampaignUpdateDto) GetProposedEnd() time.Time`

GetProposedEnd returns the ProposedEnd field if non-nil, zero value otherwise.

### GetProposedEndOk

`func (o *MarketingCampaignUpdateDto) GetProposedEndOk() (*time.Time, bool)`

GetProposedEndOk returns a tuple with the ProposedEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedEnd

`func (o *MarketingCampaignUpdateDto) SetProposedEnd(v time.Time)`

SetProposedEnd sets ProposedEnd field to given value.

### HasProposedEnd

`func (o *MarketingCampaignUpdateDto) HasProposedEnd() bool`

HasProposedEnd returns a boolean if a field has been set.

### GetActualStart

`func (o *MarketingCampaignUpdateDto) GetActualStart() time.Time`

GetActualStart returns the ActualStart field if non-nil, zero value otherwise.

### GetActualStartOk

`func (o *MarketingCampaignUpdateDto) GetActualStartOk() (*time.Time, bool)`

GetActualStartOk returns a tuple with the ActualStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualStart

`func (o *MarketingCampaignUpdateDto) SetActualStart(v time.Time)`

SetActualStart sets ActualStart field to given value.

### HasActualStart

`func (o *MarketingCampaignUpdateDto) HasActualStart() bool`

HasActualStart returns a boolean if a field has been set.

### GetActualEnd

`func (o *MarketingCampaignUpdateDto) GetActualEnd() time.Time`

GetActualEnd returns the ActualEnd field if non-nil, zero value otherwise.

### GetActualEndOk

`func (o *MarketingCampaignUpdateDto) GetActualEndOk() (*time.Time, bool)`

GetActualEndOk returns a tuple with the ActualEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualEnd

`func (o *MarketingCampaignUpdateDto) SetActualEnd(v time.Time)`

SetActualEnd sets ActualEnd field to given value.

### HasActualEnd

`func (o *MarketingCampaignUpdateDto) HasActualEnd() bool`

HasActualEnd returns a boolean if a field has been set.

### GetCode

`func (o *MarketingCampaignUpdateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MarketingCampaignUpdateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MarketingCampaignUpdateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *MarketingCampaignUpdateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *MarketingCampaignUpdateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *MarketingCampaignUpdateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetAllocatedBudget

`func (o *MarketingCampaignUpdateDto) GetAllocatedBudget() float64`

GetAllocatedBudget returns the AllocatedBudget field if non-nil, zero value otherwise.

### GetAllocatedBudgetOk

`func (o *MarketingCampaignUpdateDto) GetAllocatedBudgetOk() (*float64, bool)`

GetAllocatedBudgetOk returns a tuple with the AllocatedBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedBudget

`func (o *MarketingCampaignUpdateDto) SetAllocatedBudget(v float64)`

SetAllocatedBudget sets AllocatedBudget field to given value.

### HasAllocatedBudget

`func (o *MarketingCampaignUpdateDto) HasAllocatedBudget() bool`

HasAllocatedBudget returns a boolean if a field has been set.

### GetActivityCost

`func (o *MarketingCampaignUpdateDto) GetActivityCost() float64`

GetActivityCost returns the ActivityCost field if non-nil, zero value otherwise.

### GetActivityCostOk

`func (o *MarketingCampaignUpdateDto) GetActivityCostOk() (*float64, bool)`

GetActivityCostOk returns a tuple with the ActivityCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityCost

`func (o *MarketingCampaignUpdateDto) SetActivityCost(v float64)`

SetActivityCost sets ActivityCost field to given value.

### HasActivityCost

`func (o *MarketingCampaignUpdateDto) HasActivityCost() bool`

HasActivityCost returns a boolean if a field has been set.

### GetMiscCost

`func (o *MarketingCampaignUpdateDto) GetMiscCost() float64`

GetMiscCost returns the MiscCost field if non-nil, zero value otherwise.

### GetMiscCostOk

`func (o *MarketingCampaignUpdateDto) GetMiscCostOk() (*float64, bool)`

GetMiscCostOk returns a tuple with the MiscCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiscCost

`func (o *MarketingCampaignUpdateDto) SetMiscCost(v float64)`

SetMiscCost sets MiscCost field to given value.

### HasMiscCost

`func (o *MarketingCampaignUpdateDto) HasMiscCost() bool`

HasMiscCost returns a boolean if a field has been set.

### GetExpectedResponsePercent

`func (o *MarketingCampaignUpdateDto) GetExpectedResponsePercent() float64`

GetExpectedResponsePercent returns the ExpectedResponsePercent field if non-nil, zero value otherwise.

### GetExpectedResponsePercentOk

`func (o *MarketingCampaignUpdateDto) GetExpectedResponsePercentOk() (*float64, bool)`

GetExpectedResponsePercentOk returns a tuple with the ExpectedResponsePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResponsePercent

`func (o *MarketingCampaignUpdateDto) SetExpectedResponsePercent(v float64)`

SetExpectedResponsePercent sets ExpectedResponsePercent field to given value.

### HasExpectedResponsePercent

`func (o *MarketingCampaignUpdateDto) HasExpectedResponsePercent() bool`

HasExpectedResponsePercent returns a boolean if a field has been set.

### GetMarketingAreaId

`func (o *MarketingCampaignUpdateDto) GetMarketingAreaId() string`

GetMarketingAreaId returns the MarketingAreaId field if non-nil, zero value otherwise.

### GetMarketingAreaIdOk

`func (o *MarketingCampaignUpdateDto) GetMarketingAreaIdOk() (*string, bool)`

GetMarketingAreaIdOk returns a tuple with the MarketingAreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingAreaId

`func (o *MarketingCampaignUpdateDto) SetMarketingAreaId(v string)`

SetMarketingAreaId sets MarketingAreaId field to given value.

### HasMarketingAreaId

`func (o *MarketingCampaignUpdateDto) HasMarketingAreaId() bool`

HasMarketingAreaId returns a boolean if a field has been set.

### SetMarketingAreaIdNil

`func (o *MarketingCampaignUpdateDto) SetMarketingAreaIdNil(b bool)`

 SetMarketingAreaIdNil sets the value for MarketingAreaId to be an explicit nil

### UnsetMarketingAreaId
`func (o *MarketingCampaignUpdateDto) UnsetMarketingAreaId()`

UnsetMarketingAreaId ensures that no value is present for MarketingAreaId, not even an explicit nil
### GetCurrencyId

`func (o *MarketingCampaignUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *MarketingCampaignUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *MarketingCampaignUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *MarketingCampaignUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *MarketingCampaignUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *MarketingCampaignUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTenantId

`func (o *MarketingCampaignUpdateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MarketingCampaignUpdateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MarketingCampaignUpdateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *MarketingCampaignUpdateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *MarketingCampaignUpdateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *MarketingCampaignUpdateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrolmentId

`func (o *MarketingCampaignUpdateDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *MarketingCampaignUpdateDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *MarketingCampaignUpdateDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *MarketingCampaignUpdateDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *MarketingCampaignUpdateDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *MarketingCampaignUpdateDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


