# MarketingCampaignCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
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

### NewMarketingCampaignCreateDto

`func NewMarketingCampaignCreateDto() *MarketingCampaignCreateDto`

NewMarketingCampaignCreateDto instantiates a new MarketingCampaignCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingCampaignCreateDtoWithDefaults

`func NewMarketingCampaignCreateDtoWithDefaults() *MarketingCampaignCreateDto`

NewMarketingCampaignCreateDtoWithDefaults instantiates a new MarketingCampaignCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MarketingCampaignCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingCampaignCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingCampaignCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingCampaignCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *MarketingCampaignCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MarketingCampaignCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MarketingCampaignCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MarketingCampaignCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *MarketingCampaignCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingCampaignCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingCampaignCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingCampaignCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MarketingCampaignCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MarketingCampaignCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOffer

`func (o *MarketingCampaignCreateDto) GetOffer() string`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *MarketingCampaignCreateDto) GetOfferOk() (*string, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *MarketingCampaignCreateDto) SetOffer(v string)`

SetOffer sets Offer field to given value.

### HasOffer

`func (o *MarketingCampaignCreateDto) HasOffer() bool`

HasOffer returns a boolean if a field has been set.

### SetOfferNil

`func (o *MarketingCampaignCreateDto) SetOfferNil(b bool)`

 SetOfferNil sets the value for Offer to be an explicit nil

### UnsetOffer
`func (o *MarketingCampaignCreateDto) UnsetOffer()`

UnsetOffer ensures that no value is present for Offer, not even an explicit nil
### GetActive

`func (o *MarketingCampaignCreateDto) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MarketingCampaignCreateDto) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MarketingCampaignCreateDto) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *MarketingCampaignCreateDto) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProposedStart

`func (o *MarketingCampaignCreateDto) GetProposedStart() time.Time`

GetProposedStart returns the ProposedStart field if non-nil, zero value otherwise.

### GetProposedStartOk

`func (o *MarketingCampaignCreateDto) GetProposedStartOk() (*time.Time, bool)`

GetProposedStartOk returns a tuple with the ProposedStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedStart

`func (o *MarketingCampaignCreateDto) SetProposedStart(v time.Time)`

SetProposedStart sets ProposedStart field to given value.

### HasProposedStart

`func (o *MarketingCampaignCreateDto) HasProposedStart() bool`

HasProposedStart returns a boolean if a field has been set.

### GetProposedEnd

`func (o *MarketingCampaignCreateDto) GetProposedEnd() time.Time`

GetProposedEnd returns the ProposedEnd field if non-nil, zero value otherwise.

### GetProposedEndOk

`func (o *MarketingCampaignCreateDto) GetProposedEndOk() (*time.Time, bool)`

GetProposedEndOk returns a tuple with the ProposedEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedEnd

`func (o *MarketingCampaignCreateDto) SetProposedEnd(v time.Time)`

SetProposedEnd sets ProposedEnd field to given value.

### HasProposedEnd

`func (o *MarketingCampaignCreateDto) HasProposedEnd() bool`

HasProposedEnd returns a boolean if a field has been set.

### GetActualStart

`func (o *MarketingCampaignCreateDto) GetActualStart() time.Time`

GetActualStart returns the ActualStart field if non-nil, zero value otherwise.

### GetActualStartOk

`func (o *MarketingCampaignCreateDto) GetActualStartOk() (*time.Time, bool)`

GetActualStartOk returns a tuple with the ActualStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualStart

`func (o *MarketingCampaignCreateDto) SetActualStart(v time.Time)`

SetActualStart sets ActualStart field to given value.

### HasActualStart

`func (o *MarketingCampaignCreateDto) HasActualStart() bool`

HasActualStart returns a boolean if a field has been set.

### GetActualEnd

`func (o *MarketingCampaignCreateDto) GetActualEnd() time.Time`

GetActualEnd returns the ActualEnd field if non-nil, zero value otherwise.

### GetActualEndOk

`func (o *MarketingCampaignCreateDto) GetActualEndOk() (*time.Time, bool)`

GetActualEndOk returns a tuple with the ActualEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualEnd

`func (o *MarketingCampaignCreateDto) SetActualEnd(v time.Time)`

SetActualEnd sets ActualEnd field to given value.

### HasActualEnd

`func (o *MarketingCampaignCreateDto) HasActualEnd() bool`

HasActualEnd returns a boolean if a field has been set.

### GetCode

`func (o *MarketingCampaignCreateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MarketingCampaignCreateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MarketingCampaignCreateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *MarketingCampaignCreateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *MarketingCampaignCreateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *MarketingCampaignCreateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetAllocatedBudget

`func (o *MarketingCampaignCreateDto) GetAllocatedBudget() float64`

GetAllocatedBudget returns the AllocatedBudget field if non-nil, zero value otherwise.

### GetAllocatedBudgetOk

`func (o *MarketingCampaignCreateDto) GetAllocatedBudgetOk() (*float64, bool)`

GetAllocatedBudgetOk returns a tuple with the AllocatedBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedBudget

`func (o *MarketingCampaignCreateDto) SetAllocatedBudget(v float64)`

SetAllocatedBudget sets AllocatedBudget field to given value.

### HasAllocatedBudget

`func (o *MarketingCampaignCreateDto) HasAllocatedBudget() bool`

HasAllocatedBudget returns a boolean if a field has been set.

### GetActivityCost

`func (o *MarketingCampaignCreateDto) GetActivityCost() float64`

GetActivityCost returns the ActivityCost field if non-nil, zero value otherwise.

### GetActivityCostOk

`func (o *MarketingCampaignCreateDto) GetActivityCostOk() (*float64, bool)`

GetActivityCostOk returns a tuple with the ActivityCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityCost

`func (o *MarketingCampaignCreateDto) SetActivityCost(v float64)`

SetActivityCost sets ActivityCost field to given value.

### HasActivityCost

`func (o *MarketingCampaignCreateDto) HasActivityCost() bool`

HasActivityCost returns a boolean if a field has been set.

### GetMiscCost

`func (o *MarketingCampaignCreateDto) GetMiscCost() float64`

GetMiscCost returns the MiscCost field if non-nil, zero value otherwise.

### GetMiscCostOk

`func (o *MarketingCampaignCreateDto) GetMiscCostOk() (*float64, bool)`

GetMiscCostOk returns a tuple with the MiscCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiscCost

`func (o *MarketingCampaignCreateDto) SetMiscCost(v float64)`

SetMiscCost sets MiscCost field to given value.

### HasMiscCost

`func (o *MarketingCampaignCreateDto) HasMiscCost() bool`

HasMiscCost returns a boolean if a field has been set.

### GetExpectedResponsePercent

`func (o *MarketingCampaignCreateDto) GetExpectedResponsePercent() float64`

GetExpectedResponsePercent returns the ExpectedResponsePercent field if non-nil, zero value otherwise.

### GetExpectedResponsePercentOk

`func (o *MarketingCampaignCreateDto) GetExpectedResponsePercentOk() (*float64, bool)`

GetExpectedResponsePercentOk returns a tuple with the ExpectedResponsePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResponsePercent

`func (o *MarketingCampaignCreateDto) SetExpectedResponsePercent(v float64)`

SetExpectedResponsePercent sets ExpectedResponsePercent field to given value.

### HasExpectedResponsePercent

`func (o *MarketingCampaignCreateDto) HasExpectedResponsePercent() bool`

HasExpectedResponsePercent returns a boolean if a field has been set.

### GetMarketingAreaId

`func (o *MarketingCampaignCreateDto) GetMarketingAreaId() string`

GetMarketingAreaId returns the MarketingAreaId field if non-nil, zero value otherwise.

### GetMarketingAreaIdOk

`func (o *MarketingCampaignCreateDto) GetMarketingAreaIdOk() (*string, bool)`

GetMarketingAreaIdOk returns a tuple with the MarketingAreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingAreaId

`func (o *MarketingCampaignCreateDto) SetMarketingAreaId(v string)`

SetMarketingAreaId sets MarketingAreaId field to given value.

### HasMarketingAreaId

`func (o *MarketingCampaignCreateDto) HasMarketingAreaId() bool`

HasMarketingAreaId returns a boolean if a field has been set.

### SetMarketingAreaIdNil

`func (o *MarketingCampaignCreateDto) SetMarketingAreaIdNil(b bool)`

 SetMarketingAreaIdNil sets the value for MarketingAreaId to be an explicit nil

### UnsetMarketingAreaId
`func (o *MarketingCampaignCreateDto) UnsetMarketingAreaId()`

UnsetMarketingAreaId ensures that no value is present for MarketingAreaId, not even an explicit nil
### GetCurrencyId

`func (o *MarketingCampaignCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *MarketingCampaignCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *MarketingCampaignCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *MarketingCampaignCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *MarketingCampaignCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *MarketingCampaignCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTenantId

`func (o *MarketingCampaignCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MarketingCampaignCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MarketingCampaignCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *MarketingCampaignCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *MarketingCampaignCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *MarketingCampaignCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrolmentId

`func (o *MarketingCampaignCreateDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *MarketingCampaignCreateDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *MarketingCampaignCreateDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *MarketingCampaignCreateDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *MarketingCampaignCreateDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *MarketingCampaignCreateDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


