# WorkOrderDto

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
**PromisedStartDate** | Pointer to **time.Time** |  | [optional] 
**PromisedEndDate** | Pointer to **time.Time** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**IndividualId** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableString** |  | [optional] 
**ProductionPlanId** | Pointer to **NullableString** |  | [optional] 
**WorkOrderTypeId** | Pointer to **NullableString** |  | [optional] 
**WorkstationId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWorkOrderDto

`func NewWorkOrderDto() *WorkOrderDto`

NewWorkOrderDto instantiates a new WorkOrderDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkOrderDtoWithDefaults

`func NewWorkOrderDtoWithDefaults() *WorkOrderDto`

NewWorkOrderDtoWithDefaults instantiates a new WorkOrderDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkOrderDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkOrderDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkOrderDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkOrderDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WorkOrderDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WorkOrderDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *WorkOrderDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WorkOrderDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WorkOrderDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WorkOrderDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *WorkOrderDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *WorkOrderDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *WorkOrderDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkOrderDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkOrderDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkOrderDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WorkOrderDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WorkOrderDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *WorkOrderDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkOrderDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkOrderDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkOrderDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkOrderDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkOrderDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInstructions

`func (o *WorkOrderDto) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *WorkOrderDto) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *WorkOrderDto) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *WorkOrderDto) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *WorkOrderDto) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *WorkOrderDto) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetTaxable

`func (o *WorkOrderDto) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *WorkOrderDto) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *WorkOrderDto) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *WorkOrderDto) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetWorkLocation

`func (o *WorkOrderDto) GetWorkLocation() string`

GetWorkLocation returns the WorkLocation field if non-nil, zero value otherwise.

### GetWorkLocationOk

`func (o *WorkOrderDto) GetWorkLocationOk() (*string, bool)`

GetWorkLocationOk returns a tuple with the WorkLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkLocation

`func (o *WorkOrderDto) SetWorkLocation(v string)`

SetWorkLocation sets WorkLocation field to given value.

### HasWorkLocation

`func (o *WorkOrderDto) HasWorkLocation() bool`

HasWorkLocation returns a boolean if a field has been set.

### SetWorkLocationNil

`func (o *WorkOrderDto) SetWorkLocationNil(b bool)`

 SetWorkLocationNil sets the value for WorkLocation to be an explicit nil

### UnsetWorkLocation
`func (o *WorkOrderDto) UnsetWorkLocation()`

UnsetWorkLocation ensures that no value is present for WorkLocation, not even an explicit nil
### GetPromisedStartDate

`func (o *WorkOrderDto) GetPromisedStartDate() time.Time`

GetPromisedStartDate returns the PromisedStartDate field if non-nil, zero value otherwise.

### GetPromisedStartDateOk

`func (o *WorkOrderDto) GetPromisedStartDateOk() (*time.Time, bool)`

GetPromisedStartDateOk returns a tuple with the PromisedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedStartDate

`func (o *WorkOrderDto) SetPromisedStartDate(v time.Time)`

SetPromisedStartDate sets PromisedStartDate field to given value.

### HasPromisedStartDate

`func (o *WorkOrderDto) HasPromisedStartDate() bool`

HasPromisedStartDate returns a boolean if a field has been set.

### GetPromisedEndDate

`func (o *WorkOrderDto) GetPromisedEndDate() time.Time`

GetPromisedEndDate returns the PromisedEndDate field if non-nil, zero value otherwise.

### GetPromisedEndDateOk

`func (o *WorkOrderDto) GetPromisedEndDateOk() (*time.Time, bool)`

GetPromisedEndDateOk returns a tuple with the PromisedEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedEndDate

`func (o *WorkOrderDto) SetPromisedEndDate(v time.Time)`

SetPromisedEndDate sets PromisedEndDate field to given value.

### HasPromisedEndDate

`func (o *WorkOrderDto) HasPromisedEndDate() bool`

HasPromisedEndDate returns a boolean if a field has been set.

### GetCurrencyId

`func (o *WorkOrderDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *WorkOrderDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *WorkOrderDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *WorkOrderDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *WorkOrderDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *WorkOrderDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetIndividualId

`func (o *WorkOrderDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *WorkOrderDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *WorkOrderDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *WorkOrderDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *WorkOrderDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *WorkOrderDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *WorkOrderDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *WorkOrderDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *WorkOrderDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *WorkOrderDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *WorkOrderDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *WorkOrderDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetProductionPlanId

`func (o *WorkOrderDto) GetProductionPlanId() string`

GetProductionPlanId returns the ProductionPlanId field if non-nil, zero value otherwise.

### GetProductionPlanIdOk

`func (o *WorkOrderDto) GetProductionPlanIdOk() (*string, bool)`

GetProductionPlanIdOk returns a tuple with the ProductionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionPlanId

`func (o *WorkOrderDto) SetProductionPlanId(v string)`

SetProductionPlanId sets ProductionPlanId field to given value.

### HasProductionPlanId

`func (o *WorkOrderDto) HasProductionPlanId() bool`

HasProductionPlanId returns a boolean if a field has been set.

### SetProductionPlanIdNil

`func (o *WorkOrderDto) SetProductionPlanIdNil(b bool)`

 SetProductionPlanIdNil sets the value for ProductionPlanId to be an explicit nil

### UnsetProductionPlanId
`func (o *WorkOrderDto) UnsetProductionPlanId()`

UnsetProductionPlanId ensures that no value is present for ProductionPlanId, not even an explicit nil
### GetWorkOrderTypeId

`func (o *WorkOrderDto) GetWorkOrderTypeId() string`

GetWorkOrderTypeId returns the WorkOrderTypeId field if non-nil, zero value otherwise.

### GetWorkOrderTypeIdOk

`func (o *WorkOrderDto) GetWorkOrderTypeIdOk() (*string, bool)`

GetWorkOrderTypeIdOk returns a tuple with the WorkOrderTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkOrderTypeId

`func (o *WorkOrderDto) SetWorkOrderTypeId(v string)`

SetWorkOrderTypeId sets WorkOrderTypeId field to given value.

### HasWorkOrderTypeId

`func (o *WorkOrderDto) HasWorkOrderTypeId() bool`

HasWorkOrderTypeId returns a boolean if a field has been set.

### SetWorkOrderTypeIdNil

`func (o *WorkOrderDto) SetWorkOrderTypeIdNil(b bool)`

 SetWorkOrderTypeIdNil sets the value for WorkOrderTypeId to be an explicit nil

### UnsetWorkOrderTypeId
`func (o *WorkOrderDto) UnsetWorkOrderTypeId()`

UnsetWorkOrderTypeId ensures that no value is present for WorkOrderTypeId, not even an explicit nil
### GetWorkstationId

`func (o *WorkOrderDto) GetWorkstationId() string`

GetWorkstationId returns the WorkstationId field if non-nil, zero value otherwise.

### GetWorkstationIdOk

`func (o *WorkOrderDto) GetWorkstationIdOk() (*string, bool)`

GetWorkstationIdOk returns a tuple with the WorkstationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkstationId

`func (o *WorkOrderDto) SetWorkstationId(v string)`

SetWorkstationId sets WorkstationId field to given value.

### HasWorkstationId

`func (o *WorkOrderDto) HasWorkstationId() bool`

HasWorkstationId returns a boolean if a field has been set.

### SetWorkstationIdNil

`func (o *WorkOrderDto) SetWorkstationIdNil(b bool)`

 SetWorkstationIdNil sets the value for WorkstationId to be an explicit nil

### UnsetWorkstationId
`func (o *WorkOrderDto) UnsetWorkstationId()`

UnsetWorkstationId ensures that no value is present for WorkstationId, not even an explicit nil
### GetTenantId

`func (o *WorkOrderDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *WorkOrderDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *WorkOrderDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *WorkOrderDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *WorkOrderDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *WorkOrderDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


