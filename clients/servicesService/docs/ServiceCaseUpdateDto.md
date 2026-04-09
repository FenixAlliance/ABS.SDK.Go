# ServiceCaseUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**Taxable** | Pointer to **bool** |  | [optional] 
**WorkLocation** | Pointer to **NullableString** |  | [optional] 
**ServiceId** | Pointer to **NullableString** |  | [optional] 
**ServiceQueueId** | Pointer to **NullableString** |  | [optional] 
**ServiceCaseTypeId** | Pointer to **NullableString** |  | [optional] 
**ServiceLevelAgreementId** | Pointer to **NullableString** |  | [optional] 
**IndividualId** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableString** |  | [optional] 
**AccountHolderId** | Pointer to **NullableString** |  | [optional] 
**ReceiverBusinessId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**TerritoryId** | Pointer to **NullableString** |  | [optional] 
**PriceListId** | Pointer to **NullableString** |  | [optional] 
**PromisedStartDate** | Pointer to **time.Time** |  | [optional] 
**PromisedEndDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewServiceCaseUpdateDto

`func NewServiceCaseUpdateDto() *ServiceCaseUpdateDto`

NewServiceCaseUpdateDto instantiates a new ServiceCaseUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCaseUpdateDtoWithDefaults

`func NewServiceCaseUpdateDtoWithDefaults() *ServiceCaseUpdateDto`

NewServiceCaseUpdateDtoWithDefaults instantiates a new ServiceCaseUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ServiceCaseUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServiceCaseUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServiceCaseUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ServiceCaseUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ServiceCaseUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ServiceCaseUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *ServiceCaseUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceCaseUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceCaseUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceCaseUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ServiceCaseUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ServiceCaseUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInstructions

`func (o *ServiceCaseUpdateDto) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ServiceCaseUpdateDto) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ServiceCaseUpdateDto) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ServiceCaseUpdateDto) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *ServiceCaseUpdateDto) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *ServiceCaseUpdateDto) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetTaxable

`func (o *ServiceCaseUpdateDto) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *ServiceCaseUpdateDto) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *ServiceCaseUpdateDto) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *ServiceCaseUpdateDto) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetWorkLocation

`func (o *ServiceCaseUpdateDto) GetWorkLocation() string`

GetWorkLocation returns the WorkLocation field if non-nil, zero value otherwise.

### GetWorkLocationOk

`func (o *ServiceCaseUpdateDto) GetWorkLocationOk() (*string, bool)`

GetWorkLocationOk returns a tuple with the WorkLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkLocation

`func (o *ServiceCaseUpdateDto) SetWorkLocation(v string)`

SetWorkLocation sets WorkLocation field to given value.

### HasWorkLocation

`func (o *ServiceCaseUpdateDto) HasWorkLocation() bool`

HasWorkLocation returns a boolean if a field has been set.

### SetWorkLocationNil

`func (o *ServiceCaseUpdateDto) SetWorkLocationNil(b bool)`

 SetWorkLocationNil sets the value for WorkLocation to be an explicit nil

### UnsetWorkLocation
`func (o *ServiceCaseUpdateDto) UnsetWorkLocation()`

UnsetWorkLocation ensures that no value is present for WorkLocation, not even an explicit nil
### GetServiceId

`func (o *ServiceCaseUpdateDto) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceCaseUpdateDto) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceCaseUpdateDto) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServiceCaseUpdateDto) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### SetServiceIdNil

`func (o *ServiceCaseUpdateDto) SetServiceIdNil(b bool)`

 SetServiceIdNil sets the value for ServiceId to be an explicit nil

### UnsetServiceId
`func (o *ServiceCaseUpdateDto) UnsetServiceId()`

UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil
### GetServiceQueueId

`func (o *ServiceCaseUpdateDto) GetServiceQueueId() string`

GetServiceQueueId returns the ServiceQueueId field if non-nil, zero value otherwise.

### GetServiceQueueIdOk

`func (o *ServiceCaseUpdateDto) GetServiceQueueIdOk() (*string, bool)`

GetServiceQueueIdOk returns a tuple with the ServiceQueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceQueueId

`func (o *ServiceCaseUpdateDto) SetServiceQueueId(v string)`

SetServiceQueueId sets ServiceQueueId field to given value.

### HasServiceQueueId

`func (o *ServiceCaseUpdateDto) HasServiceQueueId() bool`

HasServiceQueueId returns a boolean if a field has been set.

### SetServiceQueueIdNil

`func (o *ServiceCaseUpdateDto) SetServiceQueueIdNil(b bool)`

 SetServiceQueueIdNil sets the value for ServiceQueueId to be an explicit nil

### UnsetServiceQueueId
`func (o *ServiceCaseUpdateDto) UnsetServiceQueueId()`

UnsetServiceQueueId ensures that no value is present for ServiceQueueId, not even an explicit nil
### GetServiceCaseTypeId

`func (o *ServiceCaseUpdateDto) GetServiceCaseTypeId() string`

GetServiceCaseTypeId returns the ServiceCaseTypeId field if non-nil, zero value otherwise.

### GetServiceCaseTypeIdOk

`func (o *ServiceCaseUpdateDto) GetServiceCaseTypeIdOk() (*string, bool)`

GetServiceCaseTypeIdOk returns a tuple with the ServiceCaseTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCaseTypeId

`func (o *ServiceCaseUpdateDto) SetServiceCaseTypeId(v string)`

SetServiceCaseTypeId sets ServiceCaseTypeId field to given value.

### HasServiceCaseTypeId

`func (o *ServiceCaseUpdateDto) HasServiceCaseTypeId() bool`

HasServiceCaseTypeId returns a boolean if a field has been set.

### SetServiceCaseTypeIdNil

`func (o *ServiceCaseUpdateDto) SetServiceCaseTypeIdNil(b bool)`

 SetServiceCaseTypeIdNil sets the value for ServiceCaseTypeId to be an explicit nil

### UnsetServiceCaseTypeId
`func (o *ServiceCaseUpdateDto) UnsetServiceCaseTypeId()`

UnsetServiceCaseTypeId ensures that no value is present for ServiceCaseTypeId, not even an explicit nil
### GetServiceLevelAgreementId

`func (o *ServiceCaseUpdateDto) GetServiceLevelAgreementId() string`

GetServiceLevelAgreementId returns the ServiceLevelAgreementId field if non-nil, zero value otherwise.

### GetServiceLevelAgreementIdOk

`func (o *ServiceCaseUpdateDto) GetServiceLevelAgreementIdOk() (*string, bool)`

GetServiceLevelAgreementIdOk returns a tuple with the ServiceLevelAgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelAgreementId

`func (o *ServiceCaseUpdateDto) SetServiceLevelAgreementId(v string)`

SetServiceLevelAgreementId sets ServiceLevelAgreementId field to given value.

### HasServiceLevelAgreementId

`func (o *ServiceCaseUpdateDto) HasServiceLevelAgreementId() bool`

HasServiceLevelAgreementId returns a boolean if a field has been set.

### SetServiceLevelAgreementIdNil

`func (o *ServiceCaseUpdateDto) SetServiceLevelAgreementIdNil(b bool)`

 SetServiceLevelAgreementIdNil sets the value for ServiceLevelAgreementId to be an explicit nil

### UnsetServiceLevelAgreementId
`func (o *ServiceCaseUpdateDto) UnsetServiceLevelAgreementId()`

UnsetServiceLevelAgreementId ensures that no value is present for ServiceLevelAgreementId, not even an explicit nil
### GetIndividualId

`func (o *ServiceCaseUpdateDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *ServiceCaseUpdateDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *ServiceCaseUpdateDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *ServiceCaseUpdateDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *ServiceCaseUpdateDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *ServiceCaseUpdateDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *ServiceCaseUpdateDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ServiceCaseUpdateDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ServiceCaseUpdateDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ServiceCaseUpdateDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *ServiceCaseUpdateDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *ServiceCaseUpdateDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetAccountHolderId

`func (o *ServiceCaseUpdateDto) GetAccountHolderId() string`

GetAccountHolderId returns the AccountHolderId field if non-nil, zero value otherwise.

### GetAccountHolderIdOk

`func (o *ServiceCaseUpdateDto) GetAccountHolderIdOk() (*string, bool)`

GetAccountHolderIdOk returns a tuple with the AccountHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderId

`func (o *ServiceCaseUpdateDto) SetAccountHolderId(v string)`

SetAccountHolderId sets AccountHolderId field to given value.

### HasAccountHolderId

`func (o *ServiceCaseUpdateDto) HasAccountHolderId() bool`

HasAccountHolderId returns a boolean if a field has been set.

### SetAccountHolderIdNil

`func (o *ServiceCaseUpdateDto) SetAccountHolderIdNil(b bool)`

 SetAccountHolderIdNil sets the value for AccountHolderId to be an explicit nil

### UnsetAccountHolderId
`func (o *ServiceCaseUpdateDto) UnsetAccountHolderId()`

UnsetAccountHolderId ensures that no value is present for AccountHolderId, not even an explicit nil
### GetReceiverBusinessId

`func (o *ServiceCaseUpdateDto) GetReceiverBusinessId() string`

GetReceiverBusinessId returns the ReceiverBusinessId field if non-nil, zero value otherwise.

### GetReceiverBusinessIdOk

`func (o *ServiceCaseUpdateDto) GetReceiverBusinessIdOk() (*string, bool)`

GetReceiverBusinessIdOk returns a tuple with the ReceiverBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverBusinessId

`func (o *ServiceCaseUpdateDto) SetReceiverBusinessId(v string)`

SetReceiverBusinessId sets ReceiverBusinessId field to given value.

### HasReceiverBusinessId

`func (o *ServiceCaseUpdateDto) HasReceiverBusinessId() bool`

HasReceiverBusinessId returns a boolean if a field has been set.

### SetReceiverBusinessIdNil

`func (o *ServiceCaseUpdateDto) SetReceiverBusinessIdNil(b bool)`

 SetReceiverBusinessIdNil sets the value for ReceiverBusinessId to be an explicit nil

### UnsetReceiverBusinessId
`func (o *ServiceCaseUpdateDto) UnsetReceiverBusinessId()`

UnsetReceiverBusinessId ensures that no value is present for ReceiverBusinessId, not even an explicit nil
### GetCurrencyId

`func (o *ServiceCaseUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ServiceCaseUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ServiceCaseUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ServiceCaseUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ServiceCaseUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ServiceCaseUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTerritoryId

`func (o *ServiceCaseUpdateDto) GetTerritoryId() string`

GetTerritoryId returns the TerritoryId field if non-nil, zero value otherwise.

### GetTerritoryIdOk

`func (o *ServiceCaseUpdateDto) GetTerritoryIdOk() (*string, bool)`

GetTerritoryIdOk returns a tuple with the TerritoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerritoryId

`func (o *ServiceCaseUpdateDto) SetTerritoryId(v string)`

SetTerritoryId sets TerritoryId field to given value.

### HasTerritoryId

`func (o *ServiceCaseUpdateDto) HasTerritoryId() bool`

HasTerritoryId returns a boolean if a field has been set.

### SetTerritoryIdNil

`func (o *ServiceCaseUpdateDto) SetTerritoryIdNil(b bool)`

 SetTerritoryIdNil sets the value for TerritoryId to be an explicit nil

### UnsetTerritoryId
`func (o *ServiceCaseUpdateDto) UnsetTerritoryId()`

UnsetTerritoryId ensures that no value is present for TerritoryId, not even an explicit nil
### GetPriceListId

`func (o *ServiceCaseUpdateDto) GetPriceListId() string`

GetPriceListId returns the PriceListId field if non-nil, zero value otherwise.

### GetPriceListIdOk

`func (o *ServiceCaseUpdateDto) GetPriceListIdOk() (*string, bool)`

GetPriceListIdOk returns a tuple with the PriceListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListId

`func (o *ServiceCaseUpdateDto) SetPriceListId(v string)`

SetPriceListId sets PriceListId field to given value.

### HasPriceListId

`func (o *ServiceCaseUpdateDto) HasPriceListId() bool`

HasPriceListId returns a boolean if a field has been set.

### SetPriceListIdNil

`func (o *ServiceCaseUpdateDto) SetPriceListIdNil(b bool)`

 SetPriceListIdNil sets the value for PriceListId to be an explicit nil

### UnsetPriceListId
`func (o *ServiceCaseUpdateDto) UnsetPriceListId()`

UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
### GetPromisedStartDate

`func (o *ServiceCaseUpdateDto) GetPromisedStartDate() time.Time`

GetPromisedStartDate returns the PromisedStartDate field if non-nil, zero value otherwise.

### GetPromisedStartDateOk

`func (o *ServiceCaseUpdateDto) GetPromisedStartDateOk() (*time.Time, bool)`

GetPromisedStartDateOk returns a tuple with the PromisedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedStartDate

`func (o *ServiceCaseUpdateDto) SetPromisedStartDate(v time.Time)`

SetPromisedStartDate sets PromisedStartDate field to given value.

### HasPromisedStartDate

`func (o *ServiceCaseUpdateDto) HasPromisedStartDate() bool`

HasPromisedStartDate returns a boolean if a field has been set.

### GetPromisedEndDate

`func (o *ServiceCaseUpdateDto) GetPromisedEndDate() time.Time`

GetPromisedEndDate returns the PromisedEndDate field if non-nil, zero value otherwise.

### GetPromisedEndDateOk

`func (o *ServiceCaseUpdateDto) GetPromisedEndDateOk() (*time.Time, bool)`

GetPromisedEndDateOk returns a tuple with the PromisedEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedEndDate

`func (o *ServiceCaseUpdateDto) SetPromisedEndDate(v time.Time)`

SetPromisedEndDate sets PromisedEndDate field to given value.

### HasPromisedEndDate

`func (o *ServiceCaseUpdateDto) HasPromisedEndDate() bool`

HasPromisedEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


