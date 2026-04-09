# ServiceCaseCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
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

### NewServiceCaseCreateDto

`func NewServiceCaseCreateDto() *ServiceCaseCreateDto`

NewServiceCaseCreateDto instantiates a new ServiceCaseCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCaseCreateDtoWithDefaults

`func NewServiceCaseCreateDtoWithDefaults() *ServiceCaseCreateDto`

NewServiceCaseCreateDtoWithDefaults instantiates a new ServiceCaseCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceCaseCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceCaseCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceCaseCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceCaseCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ServiceCaseCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ServiceCaseCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ServiceCaseCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ServiceCaseCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *ServiceCaseCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServiceCaseCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServiceCaseCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ServiceCaseCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ServiceCaseCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ServiceCaseCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *ServiceCaseCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceCaseCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceCaseCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceCaseCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ServiceCaseCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ServiceCaseCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInstructions

`func (o *ServiceCaseCreateDto) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ServiceCaseCreateDto) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ServiceCaseCreateDto) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ServiceCaseCreateDto) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *ServiceCaseCreateDto) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *ServiceCaseCreateDto) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetTaxable

`func (o *ServiceCaseCreateDto) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *ServiceCaseCreateDto) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *ServiceCaseCreateDto) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *ServiceCaseCreateDto) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetWorkLocation

`func (o *ServiceCaseCreateDto) GetWorkLocation() string`

GetWorkLocation returns the WorkLocation field if non-nil, zero value otherwise.

### GetWorkLocationOk

`func (o *ServiceCaseCreateDto) GetWorkLocationOk() (*string, bool)`

GetWorkLocationOk returns a tuple with the WorkLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkLocation

`func (o *ServiceCaseCreateDto) SetWorkLocation(v string)`

SetWorkLocation sets WorkLocation field to given value.

### HasWorkLocation

`func (o *ServiceCaseCreateDto) HasWorkLocation() bool`

HasWorkLocation returns a boolean if a field has been set.

### SetWorkLocationNil

`func (o *ServiceCaseCreateDto) SetWorkLocationNil(b bool)`

 SetWorkLocationNil sets the value for WorkLocation to be an explicit nil

### UnsetWorkLocation
`func (o *ServiceCaseCreateDto) UnsetWorkLocation()`

UnsetWorkLocation ensures that no value is present for WorkLocation, not even an explicit nil
### GetServiceId

`func (o *ServiceCaseCreateDto) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceCaseCreateDto) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceCaseCreateDto) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServiceCaseCreateDto) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### SetServiceIdNil

`func (o *ServiceCaseCreateDto) SetServiceIdNil(b bool)`

 SetServiceIdNil sets the value for ServiceId to be an explicit nil

### UnsetServiceId
`func (o *ServiceCaseCreateDto) UnsetServiceId()`

UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil
### GetServiceQueueId

`func (o *ServiceCaseCreateDto) GetServiceQueueId() string`

GetServiceQueueId returns the ServiceQueueId field if non-nil, zero value otherwise.

### GetServiceQueueIdOk

`func (o *ServiceCaseCreateDto) GetServiceQueueIdOk() (*string, bool)`

GetServiceQueueIdOk returns a tuple with the ServiceQueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceQueueId

`func (o *ServiceCaseCreateDto) SetServiceQueueId(v string)`

SetServiceQueueId sets ServiceQueueId field to given value.

### HasServiceQueueId

`func (o *ServiceCaseCreateDto) HasServiceQueueId() bool`

HasServiceQueueId returns a boolean if a field has been set.

### SetServiceQueueIdNil

`func (o *ServiceCaseCreateDto) SetServiceQueueIdNil(b bool)`

 SetServiceQueueIdNil sets the value for ServiceQueueId to be an explicit nil

### UnsetServiceQueueId
`func (o *ServiceCaseCreateDto) UnsetServiceQueueId()`

UnsetServiceQueueId ensures that no value is present for ServiceQueueId, not even an explicit nil
### GetServiceCaseTypeId

`func (o *ServiceCaseCreateDto) GetServiceCaseTypeId() string`

GetServiceCaseTypeId returns the ServiceCaseTypeId field if non-nil, zero value otherwise.

### GetServiceCaseTypeIdOk

`func (o *ServiceCaseCreateDto) GetServiceCaseTypeIdOk() (*string, bool)`

GetServiceCaseTypeIdOk returns a tuple with the ServiceCaseTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCaseTypeId

`func (o *ServiceCaseCreateDto) SetServiceCaseTypeId(v string)`

SetServiceCaseTypeId sets ServiceCaseTypeId field to given value.

### HasServiceCaseTypeId

`func (o *ServiceCaseCreateDto) HasServiceCaseTypeId() bool`

HasServiceCaseTypeId returns a boolean if a field has been set.

### SetServiceCaseTypeIdNil

`func (o *ServiceCaseCreateDto) SetServiceCaseTypeIdNil(b bool)`

 SetServiceCaseTypeIdNil sets the value for ServiceCaseTypeId to be an explicit nil

### UnsetServiceCaseTypeId
`func (o *ServiceCaseCreateDto) UnsetServiceCaseTypeId()`

UnsetServiceCaseTypeId ensures that no value is present for ServiceCaseTypeId, not even an explicit nil
### GetServiceLevelAgreementId

`func (o *ServiceCaseCreateDto) GetServiceLevelAgreementId() string`

GetServiceLevelAgreementId returns the ServiceLevelAgreementId field if non-nil, zero value otherwise.

### GetServiceLevelAgreementIdOk

`func (o *ServiceCaseCreateDto) GetServiceLevelAgreementIdOk() (*string, bool)`

GetServiceLevelAgreementIdOk returns a tuple with the ServiceLevelAgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelAgreementId

`func (o *ServiceCaseCreateDto) SetServiceLevelAgreementId(v string)`

SetServiceLevelAgreementId sets ServiceLevelAgreementId field to given value.

### HasServiceLevelAgreementId

`func (o *ServiceCaseCreateDto) HasServiceLevelAgreementId() bool`

HasServiceLevelAgreementId returns a boolean if a field has been set.

### SetServiceLevelAgreementIdNil

`func (o *ServiceCaseCreateDto) SetServiceLevelAgreementIdNil(b bool)`

 SetServiceLevelAgreementIdNil sets the value for ServiceLevelAgreementId to be an explicit nil

### UnsetServiceLevelAgreementId
`func (o *ServiceCaseCreateDto) UnsetServiceLevelAgreementId()`

UnsetServiceLevelAgreementId ensures that no value is present for ServiceLevelAgreementId, not even an explicit nil
### GetIndividualId

`func (o *ServiceCaseCreateDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *ServiceCaseCreateDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *ServiceCaseCreateDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *ServiceCaseCreateDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *ServiceCaseCreateDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *ServiceCaseCreateDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *ServiceCaseCreateDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ServiceCaseCreateDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ServiceCaseCreateDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ServiceCaseCreateDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *ServiceCaseCreateDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *ServiceCaseCreateDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetAccountHolderId

`func (o *ServiceCaseCreateDto) GetAccountHolderId() string`

GetAccountHolderId returns the AccountHolderId field if non-nil, zero value otherwise.

### GetAccountHolderIdOk

`func (o *ServiceCaseCreateDto) GetAccountHolderIdOk() (*string, bool)`

GetAccountHolderIdOk returns a tuple with the AccountHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderId

`func (o *ServiceCaseCreateDto) SetAccountHolderId(v string)`

SetAccountHolderId sets AccountHolderId field to given value.

### HasAccountHolderId

`func (o *ServiceCaseCreateDto) HasAccountHolderId() bool`

HasAccountHolderId returns a boolean if a field has been set.

### SetAccountHolderIdNil

`func (o *ServiceCaseCreateDto) SetAccountHolderIdNil(b bool)`

 SetAccountHolderIdNil sets the value for AccountHolderId to be an explicit nil

### UnsetAccountHolderId
`func (o *ServiceCaseCreateDto) UnsetAccountHolderId()`

UnsetAccountHolderId ensures that no value is present for AccountHolderId, not even an explicit nil
### GetReceiverBusinessId

`func (o *ServiceCaseCreateDto) GetReceiverBusinessId() string`

GetReceiverBusinessId returns the ReceiverBusinessId field if non-nil, zero value otherwise.

### GetReceiverBusinessIdOk

`func (o *ServiceCaseCreateDto) GetReceiverBusinessIdOk() (*string, bool)`

GetReceiverBusinessIdOk returns a tuple with the ReceiverBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverBusinessId

`func (o *ServiceCaseCreateDto) SetReceiverBusinessId(v string)`

SetReceiverBusinessId sets ReceiverBusinessId field to given value.

### HasReceiverBusinessId

`func (o *ServiceCaseCreateDto) HasReceiverBusinessId() bool`

HasReceiverBusinessId returns a boolean if a field has been set.

### SetReceiverBusinessIdNil

`func (o *ServiceCaseCreateDto) SetReceiverBusinessIdNil(b bool)`

 SetReceiverBusinessIdNil sets the value for ReceiverBusinessId to be an explicit nil

### UnsetReceiverBusinessId
`func (o *ServiceCaseCreateDto) UnsetReceiverBusinessId()`

UnsetReceiverBusinessId ensures that no value is present for ReceiverBusinessId, not even an explicit nil
### GetCurrencyId

`func (o *ServiceCaseCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ServiceCaseCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ServiceCaseCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ServiceCaseCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ServiceCaseCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ServiceCaseCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTerritoryId

`func (o *ServiceCaseCreateDto) GetTerritoryId() string`

GetTerritoryId returns the TerritoryId field if non-nil, zero value otherwise.

### GetTerritoryIdOk

`func (o *ServiceCaseCreateDto) GetTerritoryIdOk() (*string, bool)`

GetTerritoryIdOk returns a tuple with the TerritoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerritoryId

`func (o *ServiceCaseCreateDto) SetTerritoryId(v string)`

SetTerritoryId sets TerritoryId field to given value.

### HasTerritoryId

`func (o *ServiceCaseCreateDto) HasTerritoryId() bool`

HasTerritoryId returns a boolean if a field has been set.

### SetTerritoryIdNil

`func (o *ServiceCaseCreateDto) SetTerritoryIdNil(b bool)`

 SetTerritoryIdNil sets the value for TerritoryId to be an explicit nil

### UnsetTerritoryId
`func (o *ServiceCaseCreateDto) UnsetTerritoryId()`

UnsetTerritoryId ensures that no value is present for TerritoryId, not even an explicit nil
### GetPriceListId

`func (o *ServiceCaseCreateDto) GetPriceListId() string`

GetPriceListId returns the PriceListId field if non-nil, zero value otherwise.

### GetPriceListIdOk

`func (o *ServiceCaseCreateDto) GetPriceListIdOk() (*string, bool)`

GetPriceListIdOk returns a tuple with the PriceListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListId

`func (o *ServiceCaseCreateDto) SetPriceListId(v string)`

SetPriceListId sets PriceListId field to given value.

### HasPriceListId

`func (o *ServiceCaseCreateDto) HasPriceListId() bool`

HasPriceListId returns a boolean if a field has been set.

### SetPriceListIdNil

`func (o *ServiceCaseCreateDto) SetPriceListIdNil(b bool)`

 SetPriceListIdNil sets the value for PriceListId to be an explicit nil

### UnsetPriceListId
`func (o *ServiceCaseCreateDto) UnsetPriceListId()`

UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
### GetPromisedStartDate

`func (o *ServiceCaseCreateDto) GetPromisedStartDate() time.Time`

GetPromisedStartDate returns the PromisedStartDate field if non-nil, zero value otherwise.

### GetPromisedStartDateOk

`func (o *ServiceCaseCreateDto) GetPromisedStartDateOk() (*time.Time, bool)`

GetPromisedStartDateOk returns a tuple with the PromisedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedStartDate

`func (o *ServiceCaseCreateDto) SetPromisedStartDate(v time.Time)`

SetPromisedStartDate sets PromisedStartDate field to given value.

### HasPromisedStartDate

`func (o *ServiceCaseCreateDto) HasPromisedStartDate() bool`

HasPromisedStartDate returns a boolean if a field has been set.

### GetPromisedEndDate

`func (o *ServiceCaseCreateDto) GetPromisedEndDate() time.Time`

GetPromisedEndDate returns the PromisedEndDate field if non-nil, zero value otherwise.

### GetPromisedEndDateOk

`func (o *ServiceCaseCreateDto) GetPromisedEndDateOk() (*time.Time, bool)`

GetPromisedEndDateOk returns a tuple with the PromisedEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedEndDate

`func (o *ServiceCaseCreateDto) SetPromisedEndDate(v time.Time)`

SetPromisedEndDate sets PromisedEndDate field to given value.

### HasPromisedEndDate

`func (o *ServiceCaseCreateDto) HasPromisedEndDate() bool`

HasPromisedEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


