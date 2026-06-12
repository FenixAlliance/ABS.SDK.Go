# DeliveryNoteCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ShipmentId** | Pointer to **NullableString** |  | [optional] 
**ProofOfDeliveryId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDeliveryNoteCreateDto

`func NewDeliveryNoteCreateDto() *DeliveryNoteCreateDto`

NewDeliveryNoteCreateDto instantiates a new DeliveryNoteCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryNoteCreateDtoWithDefaults

`func NewDeliveryNoteCreateDtoWithDefaults() *DeliveryNoteCreateDto`

NewDeliveryNoteCreateDtoWithDefaults instantiates a new DeliveryNoteCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeliveryNoteCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliveryNoteCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliveryNoteCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeliveryNoteCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *DeliveryNoteCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DeliveryNoteCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DeliveryNoteCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DeliveryNoteCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *DeliveryNoteCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DeliveryNoteCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DeliveryNoteCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DeliveryNoteCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *DeliveryNoteCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *DeliveryNoteCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *DeliveryNoteCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeliveryNoteCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeliveryNoteCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeliveryNoteCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DeliveryNoteCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DeliveryNoteCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetShipmentId

`func (o *DeliveryNoteCreateDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *DeliveryNoteCreateDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *DeliveryNoteCreateDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *DeliveryNoteCreateDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *DeliveryNoteCreateDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *DeliveryNoteCreateDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetProofOfDeliveryId

`func (o *DeliveryNoteCreateDto) GetProofOfDeliveryId() string`

GetProofOfDeliveryId returns the ProofOfDeliveryId field if non-nil, zero value otherwise.

### GetProofOfDeliveryIdOk

`func (o *DeliveryNoteCreateDto) GetProofOfDeliveryIdOk() (*string, bool)`

GetProofOfDeliveryIdOk returns a tuple with the ProofOfDeliveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofOfDeliveryId

`func (o *DeliveryNoteCreateDto) SetProofOfDeliveryId(v string)`

SetProofOfDeliveryId sets ProofOfDeliveryId field to given value.

### HasProofOfDeliveryId

`func (o *DeliveryNoteCreateDto) HasProofOfDeliveryId() bool`

HasProofOfDeliveryId returns a boolean if a field has been set.

### SetProofOfDeliveryIdNil

`func (o *DeliveryNoteCreateDto) SetProofOfDeliveryIdNil(b bool)`

 SetProofOfDeliveryIdNil sets the value for ProofOfDeliveryId to be an explicit nil

### UnsetProofOfDeliveryId
`func (o *DeliveryNoteCreateDto) UnsetProofOfDeliveryId()`

UnsetProofOfDeliveryId ensures that no value is present for ProofOfDeliveryId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


