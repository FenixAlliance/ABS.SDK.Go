# WorkOrderUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**Taxable** | Pointer to **NullableBool** |  | [optional] 
**WorkLocation** | Pointer to **NullableString** |  | [optional] 
**ProductionPlanId** | Pointer to **NullableString** |  | [optional] 
**WorkOrderTypeId** | Pointer to **NullableString** |  | [optional] 
**WorkstationId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**IndividualId** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableString** |  | [optional] 
**PromisedStartDate** | Pointer to **NullableTime** |  | [optional] 
**PromisedEndDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewWorkOrderUpdateDto

`func NewWorkOrderUpdateDto() *WorkOrderUpdateDto`

NewWorkOrderUpdateDto instantiates a new WorkOrderUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkOrderUpdateDtoWithDefaults

`func NewWorkOrderUpdateDtoWithDefaults() *WorkOrderUpdateDto`

NewWorkOrderUpdateDtoWithDefaults instantiates a new WorkOrderUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *WorkOrderUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkOrderUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkOrderUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkOrderUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WorkOrderUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WorkOrderUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *WorkOrderUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkOrderUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkOrderUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkOrderUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkOrderUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkOrderUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInstructions

`func (o *WorkOrderUpdateDto) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *WorkOrderUpdateDto) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *WorkOrderUpdateDto) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *WorkOrderUpdateDto) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *WorkOrderUpdateDto) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *WorkOrderUpdateDto) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetTaxable

`func (o *WorkOrderUpdateDto) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *WorkOrderUpdateDto) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *WorkOrderUpdateDto) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *WorkOrderUpdateDto) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### SetTaxableNil

`func (o *WorkOrderUpdateDto) SetTaxableNil(b bool)`

 SetTaxableNil sets the value for Taxable to be an explicit nil

### UnsetTaxable
`func (o *WorkOrderUpdateDto) UnsetTaxable()`

UnsetTaxable ensures that no value is present for Taxable, not even an explicit nil
### GetWorkLocation

`func (o *WorkOrderUpdateDto) GetWorkLocation() string`

GetWorkLocation returns the WorkLocation field if non-nil, zero value otherwise.

### GetWorkLocationOk

`func (o *WorkOrderUpdateDto) GetWorkLocationOk() (*string, bool)`

GetWorkLocationOk returns a tuple with the WorkLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkLocation

`func (o *WorkOrderUpdateDto) SetWorkLocation(v string)`

SetWorkLocation sets WorkLocation field to given value.

### HasWorkLocation

`func (o *WorkOrderUpdateDto) HasWorkLocation() bool`

HasWorkLocation returns a boolean if a field has been set.

### SetWorkLocationNil

`func (o *WorkOrderUpdateDto) SetWorkLocationNil(b bool)`

 SetWorkLocationNil sets the value for WorkLocation to be an explicit nil

### UnsetWorkLocation
`func (o *WorkOrderUpdateDto) UnsetWorkLocation()`

UnsetWorkLocation ensures that no value is present for WorkLocation, not even an explicit nil
### GetProductionPlanId

`func (o *WorkOrderUpdateDto) GetProductionPlanId() string`

GetProductionPlanId returns the ProductionPlanId field if non-nil, zero value otherwise.

### GetProductionPlanIdOk

`func (o *WorkOrderUpdateDto) GetProductionPlanIdOk() (*string, bool)`

GetProductionPlanIdOk returns a tuple with the ProductionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionPlanId

`func (o *WorkOrderUpdateDto) SetProductionPlanId(v string)`

SetProductionPlanId sets ProductionPlanId field to given value.

### HasProductionPlanId

`func (o *WorkOrderUpdateDto) HasProductionPlanId() bool`

HasProductionPlanId returns a boolean if a field has been set.

### SetProductionPlanIdNil

`func (o *WorkOrderUpdateDto) SetProductionPlanIdNil(b bool)`

 SetProductionPlanIdNil sets the value for ProductionPlanId to be an explicit nil

### UnsetProductionPlanId
`func (o *WorkOrderUpdateDto) UnsetProductionPlanId()`

UnsetProductionPlanId ensures that no value is present for ProductionPlanId, not even an explicit nil
### GetWorkOrderTypeId

`func (o *WorkOrderUpdateDto) GetWorkOrderTypeId() string`

GetWorkOrderTypeId returns the WorkOrderTypeId field if non-nil, zero value otherwise.

### GetWorkOrderTypeIdOk

`func (o *WorkOrderUpdateDto) GetWorkOrderTypeIdOk() (*string, bool)`

GetWorkOrderTypeIdOk returns a tuple with the WorkOrderTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkOrderTypeId

`func (o *WorkOrderUpdateDto) SetWorkOrderTypeId(v string)`

SetWorkOrderTypeId sets WorkOrderTypeId field to given value.

### HasWorkOrderTypeId

`func (o *WorkOrderUpdateDto) HasWorkOrderTypeId() bool`

HasWorkOrderTypeId returns a boolean if a field has been set.

### SetWorkOrderTypeIdNil

`func (o *WorkOrderUpdateDto) SetWorkOrderTypeIdNil(b bool)`

 SetWorkOrderTypeIdNil sets the value for WorkOrderTypeId to be an explicit nil

### UnsetWorkOrderTypeId
`func (o *WorkOrderUpdateDto) UnsetWorkOrderTypeId()`

UnsetWorkOrderTypeId ensures that no value is present for WorkOrderTypeId, not even an explicit nil
### GetWorkstationId

`func (o *WorkOrderUpdateDto) GetWorkstationId() string`

GetWorkstationId returns the WorkstationId field if non-nil, zero value otherwise.

### GetWorkstationIdOk

`func (o *WorkOrderUpdateDto) GetWorkstationIdOk() (*string, bool)`

GetWorkstationIdOk returns a tuple with the WorkstationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkstationId

`func (o *WorkOrderUpdateDto) SetWorkstationId(v string)`

SetWorkstationId sets WorkstationId field to given value.

### HasWorkstationId

`func (o *WorkOrderUpdateDto) HasWorkstationId() bool`

HasWorkstationId returns a boolean if a field has been set.

### SetWorkstationIdNil

`func (o *WorkOrderUpdateDto) SetWorkstationIdNil(b bool)`

 SetWorkstationIdNil sets the value for WorkstationId to be an explicit nil

### UnsetWorkstationId
`func (o *WorkOrderUpdateDto) UnsetWorkstationId()`

UnsetWorkstationId ensures that no value is present for WorkstationId, not even an explicit nil
### GetCurrencyId

`func (o *WorkOrderUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *WorkOrderUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *WorkOrderUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *WorkOrderUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *WorkOrderUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *WorkOrderUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetIndividualId

`func (o *WorkOrderUpdateDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *WorkOrderUpdateDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *WorkOrderUpdateDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *WorkOrderUpdateDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *WorkOrderUpdateDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *WorkOrderUpdateDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *WorkOrderUpdateDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *WorkOrderUpdateDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *WorkOrderUpdateDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *WorkOrderUpdateDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *WorkOrderUpdateDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *WorkOrderUpdateDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetPromisedStartDate

`func (o *WorkOrderUpdateDto) GetPromisedStartDate() time.Time`

GetPromisedStartDate returns the PromisedStartDate field if non-nil, zero value otherwise.

### GetPromisedStartDateOk

`func (o *WorkOrderUpdateDto) GetPromisedStartDateOk() (*time.Time, bool)`

GetPromisedStartDateOk returns a tuple with the PromisedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedStartDate

`func (o *WorkOrderUpdateDto) SetPromisedStartDate(v time.Time)`

SetPromisedStartDate sets PromisedStartDate field to given value.

### HasPromisedStartDate

`func (o *WorkOrderUpdateDto) HasPromisedStartDate() bool`

HasPromisedStartDate returns a boolean if a field has been set.

### SetPromisedStartDateNil

`func (o *WorkOrderUpdateDto) SetPromisedStartDateNil(b bool)`

 SetPromisedStartDateNil sets the value for PromisedStartDate to be an explicit nil

### UnsetPromisedStartDate
`func (o *WorkOrderUpdateDto) UnsetPromisedStartDate()`

UnsetPromisedStartDate ensures that no value is present for PromisedStartDate, not even an explicit nil
### GetPromisedEndDate

`func (o *WorkOrderUpdateDto) GetPromisedEndDate() time.Time`

GetPromisedEndDate returns the PromisedEndDate field if non-nil, zero value otherwise.

### GetPromisedEndDateOk

`func (o *WorkOrderUpdateDto) GetPromisedEndDateOk() (*time.Time, bool)`

GetPromisedEndDateOk returns a tuple with the PromisedEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedEndDate

`func (o *WorkOrderUpdateDto) SetPromisedEndDate(v time.Time)`

SetPromisedEndDate sets PromisedEndDate field to given value.

### HasPromisedEndDate

`func (o *WorkOrderUpdateDto) HasPromisedEndDate() bool`

HasPromisedEndDate returns a boolean if a field has been set.

### SetPromisedEndDateNil

`func (o *WorkOrderUpdateDto) SetPromisedEndDateNil(b bool)`

 SetPromisedEndDateNil sets the value for PromisedEndDate to be an explicit nil

### UnsetPromisedEndDate
`func (o *WorkOrderUpdateDto) UnsetPromisedEndDate()`

UnsetPromisedEndDate ensures that no value is present for PromisedEndDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


