# AssetTransferUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialList** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **NullableString** |  | [optional] 
**Serial** | Pointer to **NullableString** |  | [optional] 
**DestinationLocationId** | Pointer to **NullableString** |  | [optional] 
**DestinationContactId** | Pointer to **NullableString** |  | [optional] 
**DestinationDepartmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAssetTransferUpdateDto

`func NewAssetTransferUpdateDto() *AssetTransferUpdateDto`

NewAssetTransferUpdateDto instantiates a new AssetTransferUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetTransferUpdateDtoWithDefaults

`func NewAssetTransferUpdateDtoWithDefaults() *AssetTransferUpdateDto`

NewAssetTransferUpdateDtoWithDefaults instantiates a new AssetTransferUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialList

`func (o *AssetTransferUpdateDto) GetSerialList() string`

GetSerialList returns the SerialList field if non-nil, zero value otherwise.

### GetSerialListOk

`func (o *AssetTransferUpdateDto) GetSerialListOk() (*string, bool)`

GetSerialListOk returns a tuple with the SerialList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialList

`func (o *AssetTransferUpdateDto) SetSerialList(v string)`

SetSerialList sets SerialList field to given value.

### HasSerialList

`func (o *AssetTransferUpdateDto) HasSerialList() bool`

HasSerialList returns a boolean if a field has been set.

### SetSerialListNil

`func (o *AssetTransferUpdateDto) SetSerialListNil(b bool)`

 SetSerialListNil sets the value for SerialList to be an explicit nil

### UnsetSerialList
`func (o *AssetTransferUpdateDto) UnsetSerialList()`

UnsetSerialList ensures that no value is present for SerialList, not even an explicit nil
### GetQuantity

`func (o *AssetTransferUpdateDto) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AssetTransferUpdateDto) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AssetTransferUpdateDto) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *AssetTransferUpdateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *AssetTransferUpdateDto) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *AssetTransferUpdateDto) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetSerial

`func (o *AssetTransferUpdateDto) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *AssetTransferUpdateDto) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *AssetTransferUpdateDto) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *AssetTransferUpdateDto) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### SetSerialNil

`func (o *AssetTransferUpdateDto) SetSerialNil(b bool)`

 SetSerialNil sets the value for Serial to be an explicit nil

### UnsetSerial
`func (o *AssetTransferUpdateDto) UnsetSerial()`

UnsetSerial ensures that no value is present for Serial, not even an explicit nil
### GetDestinationLocationId

`func (o *AssetTransferUpdateDto) GetDestinationLocationId() string`

GetDestinationLocationId returns the DestinationLocationId field if non-nil, zero value otherwise.

### GetDestinationLocationIdOk

`func (o *AssetTransferUpdateDto) GetDestinationLocationIdOk() (*string, bool)`

GetDestinationLocationIdOk returns a tuple with the DestinationLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLocationId

`func (o *AssetTransferUpdateDto) SetDestinationLocationId(v string)`

SetDestinationLocationId sets DestinationLocationId field to given value.

### HasDestinationLocationId

`func (o *AssetTransferUpdateDto) HasDestinationLocationId() bool`

HasDestinationLocationId returns a boolean if a field has been set.

### SetDestinationLocationIdNil

`func (o *AssetTransferUpdateDto) SetDestinationLocationIdNil(b bool)`

 SetDestinationLocationIdNil sets the value for DestinationLocationId to be an explicit nil

### UnsetDestinationLocationId
`func (o *AssetTransferUpdateDto) UnsetDestinationLocationId()`

UnsetDestinationLocationId ensures that no value is present for DestinationLocationId, not even an explicit nil
### GetDestinationContactId

`func (o *AssetTransferUpdateDto) GetDestinationContactId() string`

GetDestinationContactId returns the DestinationContactId field if non-nil, zero value otherwise.

### GetDestinationContactIdOk

`func (o *AssetTransferUpdateDto) GetDestinationContactIdOk() (*string, bool)`

GetDestinationContactIdOk returns a tuple with the DestinationContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationContactId

`func (o *AssetTransferUpdateDto) SetDestinationContactId(v string)`

SetDestinationContactId sets DestinationContactId field to given value.

### HasDestinationContactId

`func (o *AssetTransferUpdateDto) HasDestinationContactId() bool`

HasDestinationContactId returns a boolean if a field has been set.

### SetDestinationContactIdNil

`func (o *AssetTransferUpdateDto) SetDestinationContactIdNil(b bool)`

 SetDestinationContactIdNil sets the value for DestinationContactId to be an explicit nil

### UnsetDestinationContactId
`func (o *AssetTransferUpdateDto) UnsetDestinationContactId()`

UnsetDestinationContactId ensures that no value is present for DestinationContactId, not even an explicit nil
### GetDestinationDepartmentId

`func (o *AssetTransferUpdateDto) GetDestinationDepartmentId() string`

GetDestinationDepartmentId returns the DestinationDepartmentId field if non-nil, zero value otherwise.

### GetDestinationDepartmentIdOk

`func (o *AssetTransferUpdateDto) GetDestinationDepartmentIdOk() (*string, bool)`

GetDestinationDepartmentIdOk returns a tuple with the DestinationDepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDepartmentId

`func (o *AssetTransferUpdateDto) SetDestinationDepartmentId(v string)`

SetDestinationDepartmentId sets DestinationDepartmentId field to given value.

### HasDestinationDepartmentId

`func (o *AssetTransferUpdateDto) HasDestinationDepartmentId() bool`

HasDestinationDepartmentId returns a boolean if a field has been set.

### SetDestinationDepartmentIdNil

`func (o *AssetTransferUpdateDto) SetDestinationDepartmentIdNil(b bool)`

 SetDestinationDepartmentIdNil sets the value for DestinationDepartmentId to be an explicit nil

### UnsetDestinationDepartmentId
`func (o *AssetTransferUpdateDto) UnsetDestinationDepartmentId()`

UnsetDestinationDepartmentId ensures that no value is present for DestinationDepartmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


