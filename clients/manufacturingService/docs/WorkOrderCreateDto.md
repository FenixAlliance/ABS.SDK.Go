# WorkOrderCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**ProductionPlanId** | Pointer to **NullableString** |  | [optional] 
**WorkOrderTypeId** | Pointer to **NullableString** |  | [optional] 
**WorkstationId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**IndividualId** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableString** |  | [optional] 
**PromisedStartDate** | Pointer to **NullableTime** |  | [optional] 
**PromisedEndDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewWorkOrderCreateDto

`func NewWorkOrderCreateDto(title string, ) *WorkOrderCreateDto`

NewWorkOrderCreateDto instantiates a new WorkOrderCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkOrderCreateDtoWithDefaults

`func NewWorkOrderCreateDtoWithDefaults() *WorkOrderCreateDto`

NewWorkOrderCreateDtoWithDefaults instantiates a new WorkOrderCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkOrderCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkOrderCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkOrderCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkOrderCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *WorkOrderCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WorkOrderCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WorkOrderCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WorkOrderCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *WorkOrderCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkOrderCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkOrderCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *WorkOrderCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkOrderCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkOrderCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkOrderCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkOrderCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkOrderCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInstructions

`func (o *WorkOrderCreateDto) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *WorkOrderCreateDto) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *WorkOrderCreateDto) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *WorkOrderCreateDto) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *WorkOrderCreateDto) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *WorkOrderCreateDto) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetProductionPlanId

`func (o *WorkOrderCreateDto) GetProductionPlanId() string`

GetProductionPlanId returns the ProductionPlanId field if non-nil, zero value otherwise.

### GetProductionPlanIdOk

`func (o *WorkOrderCreateDto) GetProductionPlanIdOk() (*string, bool)`

GetProductionPlanIdOk returns a tuple with the ProductionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionPlanId

`func (o *WorkOrderCreateDto) SetProductionPlanId(v string)`

SetProductionPlanId sets ProductionPlanId field to given value.

### HasProductionPlanId

`func (o *WorkOrderCreateDto) HasProductionPlanId() bool`

HasProductionPlanId returns a boolean if a field has been set.

### SetProductionPlanIdNil

`func (o *WorkOrderCreateDto) SetProductionPlanIdNil(b bool)`

 SetProductionPlanIdNil sets the value for ProductionPlanId to be an explicit nil

### UnsetProductionPlanId
`func (o *WorkOrderCreateDto) UnsetProductionPlanId()`

UnsetProductionPlanId ensures that no value is present for ProductionPlanId, not even an explicit nil
### GetWorkOrderTypeId

`func (o *WorkOrderCreateDto) GetWorkOrderTypeId() string`

GetWorkOrderTypeId returns the WorkOrderTypeId field if non-nil, zero value otherwise.

### GetWorkOrderTypeIdOk

`func (o *WorkOrderCreateDto) GetWorkOrderTypeIdOk() (*string, bool)`

GetWorkOrderTypeIdOk returns a tuple with the WorkOrderTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkOrderTypeId

`func (o *WorkOrderCreateDto) SetWorkOrderTypeId(v string)`

SetWorkOrderTypeId sets WorkOrderTypeId field to given value.

### HasWorkOrderTypeId

`func (o *WorkOrderCreateDto) HasWorkOrderTypeId() bool`

HasWorkOrderTypeId returns a boolean if a field has been set.

### SetWorkOrderTypeIdNil

`func (o *WorkOrderCreateDto) SetWorkOrderTypeIdNil(b bool)`

 SetWorkOrderTypeIdNil sets the value for WorkOrderTypeId to be an explicit nil

### UnsetWorkOrderTypeId
`func (o *WorkOrderCreateDto) UnsetWorkOrderTypeId()`

UnsetWorkOrderTypeId ensures that no value is present for WorkOrderTypeId, not even an explicit nil
### GetWorkstationId

`func (o *WorkOrderCreateDto) GetWorkstationId() string`

GetWorkstationId returns the WorkstationId field if non-nil, zero value otherwise.

### GetWorkstationIdOk

`func (o *WorkOrderCreateDto) GetWorkstationIdOk() (*string, bool)`

GetWorkstationIdOk returns a tuple with the WorkstationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkstationId

`func (o *WorkOrderCreateDto) SetWorkstationId(v string)`

SetWorkstationId sets WorkstationId field to given value.

### HasWorkstationId

`func (o *WorkOrderCreateDto) HasWorkstationId() bool`

HasWorkstationId returns a boolean if a field has been set.

### SetWorkstationIdNil

`func (o *WorkOrderCreateDto) SetWorkstationIdNil(b bool)`

 SetWorkstationIdNil sets the value for WorkstationId to be an explicit nil

### UnsetWorkstationId
`func (o *WorkOrderCreateDto) UnsetWorkstationId()`

UnsetWorkstationId ensures that no value is present for WorkstationId, not even an explicit nil
### GetCurrencyId

`func (o *WorkOrderCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *WorkOrderCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *WorkOrderCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *WorkOrderCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *WorkOrderCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *WorkOrderCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetIndividualId

`func (o *WorkOrderCreateDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *WorkOrderCreateDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *WorkOrderCreateDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *WorkOrderCreateDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *WorkOrderCreateDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *WorkOrderCreateDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *WorkOrderCreateDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *WorkOrderCreateDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *WorkOrderCreateDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *WorkOrderCreateDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *WorkOrderCreateDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *WorkOrderCreateDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetPromisedStartDate

`func (o *WorkOrderCreateDto) GetPromisedStartDate() time.Time`

GetPromisedStartDate returns the PromisedStartDate field if non-nil, zero value otherwise.

### GetPromisedStartDateOk

`func (o *WorkOrderCreateDto) GetPromisedStartDateOk() (*time.Time, bool)`

GetPromisedStartDateOk returns a tuple with the PromisedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedStartDate

`func (o *WorkOrderCreateDto) SetPromisedStartDate(v time.Time)`

SetPromisedStartDate sets PromisedStartDate field to given value.

### HasPromisedStartDate

`func (o *WorkOrderCreateDto) HasPromisedStartDate() bool`

HasPromisedStartDate returns a boolean if a field has been set.

### SetPromisedStartDateNil

`func (o *WorkOrderCreateDto) SetPromisedStartDateNil(b bool)`

 SetPromisedStartDateNil sets the value for PromisedStartDate to be an explicit nil

### UnsetPromisedStartDate
`func (o *WorkOrderCreateDto) UnsetPromisedStartDate()`

UnsetPromisedStartDate ensures that no value is present for PromisedStartDate, not even an explicit nil
### GetPromisedEndDate

`func (o *WorkOrderCreateDto) GetPromisedEndDate() time.Time`

GetPromisedEndDate returns the PromisedEndDate field if non-nil, zero value otherwise.

### GetPromisedEndDateOk

`func (o *WorkOrderCreateDto) GetPromisedEndDateOk() (*time.Time, bool)`

GetPromisedEndDateOk returns a tuple with the PromisedEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedEndDate

`func (o *WorkOrderCreateDto) SetPromisedEndDate(v time.Time)`

SetPromisedEndDate sets PromisedEndDate field to given value.

### HasPromisedEndDate

`func (o *WorkOrderCreateDto) HasPromisedEndDate() bool`

HasPromisedEndDate returns a boolean if a field has been set.

### SetPromisedEndDateNil

`func (o *WorkOrderCreateDto) SetPromisedEndDateNil(b bool)`

 SetPromisedEndDateNil sets the value for PromisedEndDate to be an explicit nil

### UnsetPromisedEndDate
`func (o *WorkOrderCreateDto) UnsetPromisedEndDate()`

UnsetPromisedEndDate ensures that no value is present for PromisedEndDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


