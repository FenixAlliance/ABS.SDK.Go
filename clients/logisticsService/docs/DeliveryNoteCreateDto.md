# DeliveryNoteCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ShipmentID** | Pointer to **NullableString** |  | [optional] 
**ProofOfDeliveryID** | Pointer to **NullableString** |  | [optional] 

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
### GetShipmentID

`func (o *DeliveryNoteCreateDto) GetShipmentID() string`

GetShipmentID returns the ShipmentID field if non-nil, zero value otherwise.

### GetShipmentIDOk

`func (o *DeliveryNoteCreateDto) GetShipmentIDOk() (*string, bool)`

GetShipmentIDOk returns a tuple with the ShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentID

`func (o *DeliveryNoteCreateDto) SetShipmentID(v string)`

SetShipmentID sets ShipmentID field to given value.

### HasShipmentID

`func (o *DeliveryNoteCreateDto) HasShipmentID() bool`

HasShipmentID returns a boolean if a field has been set.

### SetShipmentIDNil

`func (o *DeliveryNoteCreateDto) SetShipmentIDNil(b bool)`

 SetShipmentIDNil sets the value for ShipmentID to be an explicit nil

### UnsetShipmentID
`func (o *DeliveryNoteCreateDto) UnsetShipmentID()`

UnsetShipmentID ensures that no value is present for ShipmentID, not even an explicit nil
### GetProofOfDeliveryID

`func (o *DeliveryNoteCreateDto) GetProofOfDeliveryID() string`

GetProofOfDeliveryID returns the ProofOfDeliveryID field if non-nil, zero value otherwise.

### GetProofOfDeliveryIDOk

`func (o *DeliveryNoteCreateDto) GetProofOfDeliveryIDOk() (*string, bool)`

GetProofOfDeliveryIDOk returns a tuple with the ProofOfDeliveryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofOfDeliveryID

`func (o *DeliveryNoteCreateDto) SetProofOfDeliveryID(v string)`

SetProofOfDeliveryID sets ProofOfDeliveryID field to given value.

### HasProofOfDeliveryID

`func (o *DeliveryNoteCreateDto) HasProofOfDeliveryID() bool`

HasProofOfDeliveryID returns a boolean if a field has been set.

### SetProofOfDeliveryIDNil

`func (o *DeliveryNoteCreateDto) SetProofOfDeliveryIDNil(b bool)`

 SetProofOfDeliveryIDNil sets the value for ProofOfDeliveryID to be an explicit nil

### UnsetProofOfDeliveryID
`func (o *DeliveryNoteCreateDto) UnsetProofOfDeliveryID()`

UnsetProofOfDeliveryID ensures that no value is present for ProofOfDeliveryID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


