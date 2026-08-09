# CapabilityDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Effect** | Pointer to **NullableString** |  | [optional] 
**Risks** | Pointer to **[]string** |  | [optional] 
**Surfaces** | Pointer to **[]string** |  | [optional] 
**RequiredPermission** | Pointer to **NullableString** |  | [optional] 
**Available** | Pointer to **bool** |  | [optional] 
**DeniedReason** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**InputSchema** | Pointer to **map[string]string** |  | [optional] 
**OutputSchema** | Pointer to **map[string]string** |  | [optional] 
**IsOutputCollection** | Pointer to **bool** |  | [optional] 
**RequiredInputs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCapabilityDto

`func NewCapabilityDto() *CapabilityDto`

NewCapabilityDto instantiates a new CapabilityDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityDtoWithDefaults

`func NewCapabilityDtoWithDefaults() *CapabilityDto`

NewCapabilityDtoWithDefaults instantiates a new CapabilityDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CapabilityDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CapabilityDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CapabilityDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CapabilityDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CapabilityDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CapabilityDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CapabilityDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CapabilityDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CapabilityDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CapabilityDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CapabilityDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CapabilityDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetKey

`func (o *CapabilityDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CapabilityDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CapabilityDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CapabilityDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *CapabilityDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *CapabilityDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetName

`func (o *CapabilityDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CapabilityDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CapabilityDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CapabilityDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CapabilityDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CapabilityDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CapabilityDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CapabilityDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CapabilityDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategory

`func (o *CapabilityDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CapabilityDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CapabilityDto) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CapabilityDto) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *CapabilityDto) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *CapabilityDto) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetEffect

`func (o *CapabilityDto) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *CapabilityDto) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *CapabilityDto) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *CapabilityDto) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### SetEffectNil

`func (o *CapabilityDto) SetEffectNil(b bool)`

 SetEffectNil sets the value for Effect to be an explicit nil

### UnsetEffect
`func (o *CapabilityDto) UnsetEffect()`

UnsetEffect ensures that no value is present for Effect, not even an explicit nil
### GetRisks

`func (o *CapabilityDto) GetRisks() []string`

GetRisks returns the Risks field if non-nil, zero value otherwise.

### GetRisksOk

`func (o *CapabilityDto) GetRisksOk() (*[]string, bool)`

GetRisksOk returns a tuple with the Risks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisks

`func (o *CapabilityDto) SetRisks(v []string)`

SetRisks sets Risks field to given value.

### HasRisks

`func (o *CapabilityDto) HasRisks() bool`

HasRisks returns a boolean if a field has been set.

### SetRisksNil

`func (o *CapabilityDto) SetRisksNil(b bool)`

 SetRisksNil sets the value for Risks to be an explicit nil

### UnsetRisks
`func (o *CapabilityDto) UnsetRisks()`

UnsetRisks ensures that no value is present for Risks, not even an explicit nil
### GetSurfaces

`func (o *CapabilityDto) GetSurfaces() []string`

GetSurfaces returns the Surfaces field if non-nil, zero value otherwise.

### GetSurfacesOk

`func (o *CapabilityDto) GetSurfacesOk() (*[]string, bool)`

GetSurfacesOk returns a tuple with the Surfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurfaces

`func (o *CapabilityDto) SetSurfaces(v []string)`

SetSurfaces sets Surfaces field to given value.

### HasSurfaces

`func (o *CapabilityDto) HasSurfaces() bool`

HasSurfaces returns a boolean if a field has been set.

### SetSurfacesNil

`func (o *CapabilityDto) SetSurfacesNil(b bool)`

 SetSurfacesNil sets the value for Surfaces to be an explicit nil

### UnsetSurfaces
`func (o *CapabilityDto) UnsetSurfaces()`

UnsetSurfaces ensures that no value is present for Surfaces, not even an explicit nil
### GetRequiredPermission

`func (o *CapabilityDto) GetRequiredPermission() string`

GetRequiredPermission returns the RequiredPermission field if non-nil, zero value otherwise.

### GetRequiredPermissionOk

`func (o *CapabilityDto) GetRequiredPermissionOk() (*string, bool)`

GetRequiredPermissionOk returns a tuple with the RequiredPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPermission

`func (o *CapabilityDto) SetRequiredPermission(v string)`

SetRequiredPermission sets RequiredPermission field to given value.

### HasRequiredPermission

`func (o *CapabilityDto) HasRequiredPermission() bool`

HasRequiredPermission returns a boolean if a field has been set.

### SetRequiredPermissionNil

`func (o *CapabilityDto) SetRequiredPermissionNil(b bool)`

 SetRequiredPermissionNil sets the value for RequiredPermission to be an explicit nil

### UnsetRequiredPermission
`func (o *CapabilityDto) UnsetRequiredPermission()`

UnsetRequiredPermission ensures that no value is present for RequiredPermission, not even an explicit nil
### GetAvailable

`func (o *CapabilityDto) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CapabilityDto) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CapabilityDto) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CapabilityDto) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetDeniedReason

`func (o *CapabilityDto) GetDeniedReason() string`

GetDeniedReason returns the DeniedReason field if non-nil, zero value otherwise.

### GetDeniedReasonOk

`func (o *CapabilityDto) GetDeniedReasonOk() (*string, bool)`

GetDeniedReasonOk returns a tuple with the DeniedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedReason

`func (o *CapabilityDto) SetDeniedReason(v string)`

SetDeniedReason sets DeniedReason field to given value.

### HasDeniedReason

`func (o *CapabilityDto) HasDeniedReason() bool`

HasDeniedReason returns a boolean if a field has been set.

### SetDeniedReasonNil

`func (o *CapabilityDto) SetDeniedReasonNil(b bool)`

 SetDeniedReasonNil sets the value for DeniedReason to be an explicit nil

### UnsetDeniedReason
`func (o *CapabilityDto) UnsetDeniedReason()`

UnsetDeniedReason ensures that no value is present for DeniedReason, not even an explicit nil
### GetVersion

`func (o *CapabilityDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CapabilityDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CapabilityDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CapabilityDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *CapabilityDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *CapabilityDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetInputSchema

`func (o *CapabilityDto) GetInputSchema() map[string]string`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *CapabilityDto) GetInputSchemaOk() (*map[string]string, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *CapabilityDto) SetInputSchema(v map[string]string)`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *CapabilityDto) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### SetInputSchemaNil

`func (o *CapabilityDto) SetInputSchemaNil(b bool)`

 SetInputSchemaNil sets the value for InputSchema to be an explicit nil

### UnsetInputSchema
`func (o *CapabilityDto) UnsetInputSchema()`

UnsetInputSchema ensures that no value is present for InputSchema, not even an explicit nil
### GetOutputSchema

`func (o *CapabilityDto) GetOutputSchema() map[string]string`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *CapabilityDto) GetOutputSchemaOk() (*map[string]string, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *CapabilityDto) SetOutputSchema(v map[string]string)`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *CapabilityDto) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### SetOutputSchemaNil

`func (o *CapabilityDto) SetOutputSchemaNil(b bool)`

 SetOutputSchemaNil sets the value for OutputSchema to be an explicit nil

### UnsetOutputSchema
`func (o *CapabilityDto) UnsetOutputSchema()`

UnsetOutputSchema ensures that no value is present for OutputSchema, not even an explicit nil
### GetIsOutputCollection

`func (o *CapabilityDto) GetIsOutputCollection() bool`

GetIsOutputCollection returns the IsOutputCollection field if non-nil, zero value otherwise.

### GetIsOutputCollectionOk

`func (o *CapabilityDto) GetIsOutputCollectionOk() (*bool, bool)`

GetIsOutputCollectionOk returns a tuple with the IsOutputCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOutputCollection

`func (o *CapabilityDto) SetIsOutputCollection(v bool)`

SetIsOutputCollection sets IsOutputCollection field to given value.

### HasIsOutputCollection

`func (o *CapabilityDto) HasIsOutputCollection() bool`

HasIsOutputCollection returns a boolean if a field has been set.

### GetRequiredInputs

`func (o *CapabilityDto) GetRequiredInputs() []string`

GetRequiredInputs returns the RequiredInputs field if non-nil, zero value otherwise.

### GetRequiredInputsOk

`func (o *CapabilityDto) GetRequiredInputsOk() (*[]string, bool)`

GetRequiredInputsOk returns a tuple with the RequiredInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInputs

`func (o *CapabilityDto) SetRequiredInputs(v []string)`

SetRequiredInputs sets RequiredInputs field to given value.

### HasRequiredInputs

`func (o *CapabilityDto) HasRequiredInputs() bool`

HasRequiredInputs returns a boolean if a field has been set.

### SetRequiredInputsNil

`func (o *CapabilityDto) SetRequiredInputsNil(b bool)`

 SetRequiredInputsNil sets the value for RequiredInputs to be an explicit nil

### UnsetRequiredInputs
`func (o *CapabilityDto) UnsetRequiredInputs()`

UnsetRequiredInputs ensures that no value is present for RequiredInputs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


