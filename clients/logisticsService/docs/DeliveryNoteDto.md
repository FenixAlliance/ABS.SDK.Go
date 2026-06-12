# DeliveryNoteDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ShipmentId** | Pointer to **NullableString** |  | [optional] 
**ProofOfDeliveryId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDeliveryNoteDto

`func NewDeliveryNoteDto() *DeliveryNoteDto`

NewDeliveryNoteDto instantiates a new DeliveryNoteDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryNoteDtoWithDefaults

`func NewDeliveryNoteDtoWithDefaults() *DeliveryNoteDto`

NewDeliveryNoteDtoWithDefaults instantiates a new DeliveryNoteDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeliveryNoteDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliveryNoteDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliveryNoteDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeliveryNoteDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DeliveryNoteDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DeliveryNoteDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *DeliveryNoteDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DeliveryNoteDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DeliveryNoteDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DeliveryNoteDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *DeliveryNoteDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *DeliveryNoteDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *DeliveryNoteDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DeliveryNoteDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DeliveryNoteDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DeliveryNoteDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *DeliveryNoteDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *DeliveryNoteDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *DeliveryNoteDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeliveryNoteDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeliveryNoteDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeliveryNoteDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DeliveryNoteDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DeliveryNoteDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetShipmentId

`func (o *DeliveryNoteDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *DeliveryNoteDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *DeliveryNoteDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *DeliveryNoteDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *DeliveryNoteDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *DeliveryNoteDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetProofOfDeliveryId

`func (o *DeliveryNoteDto) GetProofOfDeliveryId() string`

GetProofOfDeliveryId returns the ProofOfDeliveryId field if non-nil, zero value otherwise.

### GetProofOfDeliveryIdOk

`func (o *DeliveryNoteDto) GetProofOfDeliveryIdOk() (*string, bool)`

GetProofOfDeliveryIdOk returns a tuple with the ProofOfDeliveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofOfDeliveryId

`func (o *DeliveryNoteDto) SetProofOfDeliveryId(v string)`

SetProofOfDeliveryId sets ProofOfDeliveryId field to given value.

### HasProofOfDeliveryId

`func (o *DeliveryNoteDto) HasProofOfDeliveryId() bool`

HasProofOfDeliveryId returns a boolean if a field has been set.

### SetProofOfDeliveryIdNil

`func (o *DeliveryNoteDto) SetProofOfDeliveryIdNil(b bool)`

 SetProofOfDeliveryIdNil sets the value for ProofOfDeliveryId to be an explicit nil

### UnsetProofOfDeliveryId
`func (o *DeliveryNoteDto) UnsetProofOfDeliveryId()`

UnsetProofOfDeliveryId ensures that no value is present for ProofOfDeliveryId, not even an explicit nil
### GetTenantId

`func (o *DeliveryNoteDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DeliveryNoteDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DeliveryNoteDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DeliveryNoteDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DeliveryNoteDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DeliveryNoteDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


