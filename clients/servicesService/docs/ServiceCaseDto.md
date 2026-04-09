# ServiceCaseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
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

### NewServiceCaseDto

`func NewServiceCaseDto() *ServiceCaseDto`

NewServiceCaseDto instantiates a new ServiceCaseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCaseDtoWithDefaults

`func NewServiceCaseDtoWithDefaults() *ServiceCaseDto`

NewServiceCaseDtoWithDefaults instantiates a new ServiceCaseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceCaseDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceCaseDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceCaseDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceCaseDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ServiceCaseDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ServiceCaseDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ServiceCaseDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ServiceCaseDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ServiceCaseDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ServiceCaseDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ServiceCaseDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ServiceCaseDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *ServiceCaseDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServiceCaseDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServiceCaseDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ServiceCaseDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ServiceCaseDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ServiceCaseDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *ServiceCaseDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceCaseDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceCaseDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceCaseDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ServiceCaseDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ServiceCaseDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInstructions

`func (o *ServiceCaseDto) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ServiceCaseDto) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ServiceCaseDto) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ServiceCaseDto) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *ServiceCaseDto) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *ServiceCaseDto) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetTaxable

`func (o *ServiceCaseDto) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *ServiceCaseDto) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *ServiceCaseDto) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *ServiceCaseDto) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetWorkLocation

`func (o *ServiceCaseDto) GetWorkLocation() string`

GetWorkLocation returns the WorkLocation field if non-nil, zero value otherwise.

### GetWorkLocationOk

`func (o *ServiceCaseDto) GetWorkLocationOk() (*string, bool)`

GetWorkLocationOk returns a tuple with the WorkLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkLocation

`func (o *ServiceCaseDto) SetWorkLocation(v string)`

SetWorkLocation sets WorkLocation field to given value.

### HasWorkLocation

`func (o *ServiceCaseDto) HasWorkLocation() bool`

HasWorkLocation returns a boolean if a field has been set.

### SetWorkLocationNil

`func (o *ServiceCaseDto) SetWorkLocationNil(b bool)`

 SetWorkLocationNil sets the value for WorkLocation to be an explicit nil

### UnsetWorkLocation
`func (o *ServiceCaseDto) UnsetWorkLocation()`

UnsetWorkLocation ensures that no value is present for WorkLocation, not even an explicit nil
### GetServiceId

`func (o *ServiceCaseDto) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceCaseDto) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceCaseDto) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServiceCaseDto) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### SetServiceIdNil

`func (o *ServiceCaseDto) SetServiceIdNil(b bool)`

 SetServiceIdNil sets the value for ServiceId to be an explicit nil

### UnsetServiceId
`func (o *ServiceCaseDto) UnsetServiceId()`

UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil
### GetServiceQueueId

`func (o *ServiceCaseDto) GetServiceQueueId() string`

GetServiceQueueId returns the ServiceQueueId field if non-nil, zero value otherwise.

### GetServiceQueueIdOk

`func (o *ServiceCaseDto) GetServiceQueueIdOk() (*string, bool)`

GetServiceQueueIdOk returns a tuple with the ServiceQueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceQueueId

`func (o *ServiceCaseDto) SetServiceQueueId(v string)`

SetServiceQueueId sets ServiceQueueId field to given value.

### HasServiceQueueId

`func (o *ServiceCaseDto) HasServiceQueueId() bool`

HasServiceQueueId returns a boolean if a field has been set.

### SetServiceQueueIdNil

`func (o *ServiceCaseDto) SetServiceQueueIdNil(b bool)`

 SetServiceQueueIdNil sets the value for ServiceQueueId to be an explicit nil

### UnsetServiceQueueId
`func (o *ServiceCaseDto) UnsetServiceQueueId()`

UnsetServiceQueueId ensures that no value is present for ServiceQueueId, not even an explicit nil
### GetServiceCaseTypeId

`func (o *ServiceCaseDto) GetServiceCaseTypeId() string`

GetServiceCaseTypeId returns the ServiceCaseTypeId field if non-nil, zero value otherwise.

### GetServiceCaseTypeIdOk

`func (o *ServiceCaseDto) GetServiceCaseTypeIdOk() (*string, bool)`

GetServiceCaseTypeIdOk returns a tuple with the ServiceCaseTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCaseTypeId

`func (o *ServiceCaseDto) SetServiceCaseTypeId(v string)`

SetServiceCaseTypeId sets ServiceCaseTypeId field to given value.

### HasServiceCaseTypeId

`func (o *ServiceCaseDto) HasServiceCaseTypeId() bool`

HasServiceCaseTypeId returns a boolean if a field has been set.

### SetServiceCaseTypeIdNil

`func (o *ServiceCaseDto) SetServiceCaseTypeIdNil(b bool)`

 SetServiceCaseTypeIdNil sets the value for ServiceCaseTypeId to be an explicit nil

### UnsetServiceCaseTypeId
`func (o *ServiceCaseDto) UnsetServiceCaseTypeId()`

UnsetServiceCaseTypeId ensures that no value is present for ServiceCaseTypeId, not even an explicit nil
### GetServiceLevelAgreementId

`func (o *ServiceCaseDto) GetServiceLevelAgreementId() string`

GetServiceLevelAgreementId returns the ServiceLevelAgreementId field if non-nil, zero value otherwise.

### GetServiceLevelAgreementIdOk

`func (o *ServiceCaseDto) GetServiceLevelAgreementIdOk() (*string, bool)`

GetServiceLevelAgreementIdOk returns a tuple with the ServiceLevelAgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelAgreementId

`func (o *ServiceCaseDto) SetServiceLevelAgreementId(v string)`

SetServiceLevelAgreementId sets ServiceLevelAgreementId field to given value.

### HasServiceLevelAgreementId

`func (o *ServiceCaseDto) HasServiceLevelAgreementId() bool`

HasServiceLevelAgreementId returns a boolean if a field has been set.

### SetServiceLevelAgreementIdNil

`func (o *ServiceCaseDto) SetServiceLevelAgreementIdNil(b bool)`

 SetServiceLevelAgreementIdNil sets the value for ServiceLevelAgreementId to be an explicit nil

### UnsetServiceLevelAgreementId
`func (o *ServiceCaseDto) UnsetServiceLevelAgreementId()`

UnsetServiceLevelAgreementId ensures that no value is present for ServiceLevelAgreementId, not even an explicit nil
### GetIndividualId

`func (o *ServiceCaseDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *ServiceCaseDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *ServiceCaseDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *ServiceCaseDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *ServiceCaseDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *ServiceCaseDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *ServiceCaseDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ServiceCaseDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ServiceCaseDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ServiceCaseDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *ServiceCaseDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *ServiceCaseDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetAccountHolderId

`func (o *ServiceCaseDto) GetAccountHolderId() string`

GetAccountHolderId returns the AccountHolderId field if non-nil, zero value otherwise.

### GetAccountHolderIdOk

`func (o *ServiceCaseDto) GetAccountHolderIdOk() (*string, bool)`

GetAccountHolderIdOk returns a tuple with the AccountHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderId

`func (o *ServiceCaseDto) SetAccountHolderId(v string)`

SetAccountHolderId sets AccountHolderId field to given value.

### HasAccountHolderId

`func (o *ServiceCaseDto) HasAccountHolderId() bool`

HasAccountHolderId returns a boolean if a field has been set.

### SetAccountHolderIdNil

`func (o *ServiceCaseDto) SetAccountHolderIdNil(b bool)`

 SetAccountHolderIdNil sets the value for AccountHolderId to be an explicit nil

### UnsetAccountHolderId
`func (o *ServiceCaseDto) UnsetAccountHolderId()`

UnsetAccountHolderId ensures that no value is present for AccountHolderId, not even an explicit nil
### GetReceiverBusinessId

`func (o *ServiceCaseDto) GetReceiverBusinessId() string`

GetReceiverBusinessId returns the ReceiverBusinessId field if non-nil, zero value otherwise.

### GetReceiverBusinessIdOk

`func (o *ServiceCaseDto) GetReceiverBusinessIdOk() (*string, bool)`

GetReceiverBusinessIdOk returns a tuple with the ReceiverBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverBusinessId

`func (o *ServiceCaseDto) SetReceiverBusinessId(v string)`

SetReceiverBusinessId sets ReceiverBusinessId field to given value.

### HasReceiverBusinessId

`func (o *ServiceCaseDto) HasReceiverBusinessId() bool`

HasReceiverBusinessId returns a boolean if a field has been set.

### SetReceiverBusinessIdNil

`func (o *ServiceCaseDto) SetReceiverBusinessIdNil(b bool)`

 SetReceiverBusinessIdNil sets the value for ReceiverBusinessId to be an explicit nil

### UnsetReceiverBusinessId
`func (o *ServiceCaseDto) UnsetReceiverBusinessId()`

UnsetReceiverBusinessId ensures that no value is present for ReceiverBusinessId, not even an explicit nil
### GetCurrencyId

`func (o *ServiceCaseDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ServiceCaseDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ServiceCaseDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ServiceCaseDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ServiceCaseDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ServiceCaseDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTerritoryId

`func (o *ServiceCaseDto) GetTerritoryId() string`

GetTerritoryId returns the TerritoryId field if non-nil, zero value otherwise.

### GetTerritoryIdOk

`func (o *ServiceCaseDto) GetTerritoryIdOk() (*string, bool)`

GetTerritoryIdOk returns a tuple with the TerritoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerritoryId

`func (o *ServiceCaseDto) SetTerritoryId(v string)`

SetTerritoryId sets TerritoryId field to given value.

### HasTerritoryId

`func (o *ServiceCaseDto) HasTerritoryId() bool`

HasTerritoryId returns a boolean if a field has been set.

### SetTerritoryIdNil

`func (o *ServiceCaseDto) SetTerritoryIdNil(b bool)`

 SetTerritoryIdNil sets the value for TerritoryId to be an explicit nil

### UnsetTerritoryId
`func (o *ServiceCaseDto) UnsetTerritoryId()`

UnsetTerritoryId ensures that no value is present for TerritoryId, not even an explicit nil
### GetPriceListId

`func (o *ServiceCaseDto) GetPriceListId() string`

GetPriceListId returns the PriceListId field if non-nil, zero value otherwise.

### GetPriceListIdOk

`func (o *ServiceCaseDto) GetPriceListIdOk() (*string, bool)`

GetPriceListIdOk returns a tuple with the PriceListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListId

`func (o *ServiceCaseDto) SetPriceListId(v string)`

SetPriceListId sets PriceListId field to given value.

### HasPriceListId

`func (o *ServiceCaseDto) HasPriceListId() bool`

HasPriceListId returns a boolean if a field has been set.

### SetPriceListIdNil

`func (o *ServiceCaseDto) SetPriceListIdNil(b bool)`

 SetPriceListIdNil sets the value for PriceListId to be an explicit nil

### UnsetPriceListId
`func (o *ServiceCaseDto) UnsetPriceListId()`

UnsetPriceListId ensures that no value is present for PriceListId, not even an explicit nil
### GetPromisedStartDate

`func (o *ServiceCaseDto) GetPromisedStartDate() time.Time`

GetPromisedStartDate returns the PromisedStartDate field if non-nil, zero value otherwise.

### GetPromisedStartDateOk

`func (o *ServiceCaseDto) GetPromisedStartDateOk() (*time.Time, bool)`

GetPromisedStartDateOk returns a tuple with the PromisedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedStartDate

`func (o *ServiceCaseDto) SetPromisedStartDate(v time.Time)`

SetPromisedStartDate sets PromisedStartDate field to given value.

### HasPromisedStartDate

`func (o *ServiceCaseDto) HasPromisedStartDate() bool`

HasPromisedStartDate returns a boolean if a field has been set.

### GetPromisedEndDate

`func (o *ServiceCaseDto) GetPromisedEndDate() time.Time`

GetPromisedEndDate returns the PromisedEndDate field if non-nil, zero value otherwise.

### GetPromisedEndDateOk

`func (o *ServiceCaseDto) GetPromisedEndDateOk() (*time.Time, bool)`

GetPromisedEndDateOk returns a tuple with the PromisedEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedEndDate

`func (o *ServiceCaseDto) SetPromisedEndDate(v time.Time)`

SetPromisedEndDate sets PromisedEndDate field to given value.

### HasPromisedEndDate

`func (o *ServiceCaseDto) HasPromisedEndDate() bool`

HasPromisedEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


