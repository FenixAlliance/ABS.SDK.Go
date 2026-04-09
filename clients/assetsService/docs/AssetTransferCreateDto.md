# AssetTransferCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | Pointer to **NullableString** |  | [optional] 
**IsRootTransfer** | Pointer to **bool** |  | [optional] 
**SerialList** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **NullableString** |  | [optional] 
**Serial** | Pointer to **NullableString** |  | [optional] 
**PreviousAssetTransferId** | Pointer to **NullableString** |  | [optional] 
**SourceLocationId** | Pointer to **NullableString** |  | [optional] 
**DestinationLocationId** | Pointer to **NullableString** |  | [optional] 
**SourceContactId** | Pointer to **NullableString** |  | [optional] 
**DestinationContactId** | Pointer to **NullableString** |  | [optional] 
**SourceDepartmentId** | Pointer to **NullableString** |  | [optional] 
**DestinationDepartmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAssetTransferCreateDto

`func NewAssetTransferCreateDto() *AssetTransferCreateDto`

NewAssetTransferCreateDto instantiates a new AssetTransferCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetTransferCreateDtoWithDefaults

`func NewAssetTransferCreateDtoWithDefaults() *AssetTransferCreateDto`

NewAssetTransferCreateDtoWithDefaults instantiates a new AssetTransferCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *AssetTransferCreateDto) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetTransferCreateDto) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetTransferCreateDto) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AssetTransferCreateDto) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### SetAssetIdNil

`func (o *AssetTransferCreateDto) SetAssetIdNil(b bool)`

 SetAssetIdNil sets the value for AssetId to be an explicit nil

### UnsetAssetId
`func (o *AssetTransferCreateDto) UnsetAssetId()`

UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
### GetIsRootTransfer

`func (o *AssetTransferCreateDto) GetIsRootTransfer() bool`

GetIsRootTransfer returns the IsRootTransfer field if non-nil, zero value otherwise.

### GetIsRootTransferOk

`func (o *AssetTransferCreateDto) GetIsRootTransferOk() (*bool, bool)`

GetIsRootTransferOk returns a tuple with the IsRootTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRootTransfer

`func (o *AssetTransferCreateDto) SetIsRootTransfer(v bool)`

SetIsRootTransfer sets IsRootTransfer field to given value.

### HasIsRootTransfer

`func (o *AssetTransferCreateDto) HasIsRootTransfer() bool`

HasIsRootTransfer returns a boolean if a field has been set.

### GetSerialList

`func (o *AssetTransferCreateDto) GetSerialList() string`

GetSerialList returns the SerialList field if non-nil, zero value otherwise.

### GetSerialListOk

`func (o *AssetTransferCreateDto) GetSerialListOk() (*string, bool)`

GetSerialListOk returns a tuple with the SerialList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialList

`func (o *AssetTransferCreateDto) SetSerialList(v string)`

SetSerialList sets SerialList field to given value.

### HasSerialList

`func (o *AssetTransferCreateDto) HasSerialList() bool`

HasSerialList returns a boolean if a field has been set.

### SetSerialListNil

`func (o *AssetTransferCreateDto) SetSerialListNil(b bool)`

 SetSerialListNil sets the value for SerialList to be an explicit nil

### UnsetSerialList
`func (o *AssetTransferCreateDto) UnsetSerialList()`

UnsetSerialList ensures that no value is present for SerialList, not even an explicit nil
### GetQuantity

`func (o *AssetTransferCreateDto) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AssetTransferCreateDto) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AssetTransferCreateDto) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *AssetTransferCreateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *AssetTransferCreateDto) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *AssetTransferCreateDto) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetSerial

`func (o *AssetTransferCreateDto) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *AssetTransferCreateDto) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *AssetTransferCreateDto) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *AssetTransferCreateDto) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### SetSerialNil

`func (o *AssetTransferCreateDto) SetSerialNil(b bool)`

 SetSerialNil sets the value for Serial to be an explicit nil

### UnsetSerial
`func (o *AssetTransferCreateDto) UnsetSerial()`

UnsetSerial ensures that no value is present for Serial, not even an explicit nil
### GetPreviousAssetTransferId

`func (o *AssetTransferCreateDto) GetPreviousAssetTransferId() string`

GetPreviousAssetTransferId returns the PreviousAssetTransferId field if non-nil, zero value otherwise.

### GetPreviousAssetTransferIdOk

`func (o *AssetTransferCreateDto) GetPreviousAssetTransferIdOk() (*string, bool)`

GetPreviousAssetTransferIdOk returns a tuple with the PreviousAssetTransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAssetTransferId

`func (o *AssetTransferCreateDto) SetPreviousAssetTransferId(v string)`

SetPreviousAssetTransferId sets PreviousAssetTransferId field to given value.

### HasPreviousAssetTransferId

`func (o *AssetTransferCreateDto) HasPreviousAssetTransferId() bool`

HasPreviousAssetTransferId returns a boolean if a field has been set.

### SetPreviousAssetTransferIdNil

`func (o *AssetTransferCreateDto) SetPreviousAssetTransferIdNil(b bool)`

 SetPreviousAssetTransferIdNil sets the value for PreviousAssetTransferId to be an explicit nil

### UnsetPreviousAssetTransferId
`func (o *AssetTransferCreateDto) UnsetPreviousAssetTransferId()`

UnsetPreviousAssetTransferId ensures that no value is present for PreviousAssetTransferId, not even an explicit nil
### GetSourceLocationId

`func (o *AssetTransferCreateDto) GetSourceLocationId() string`

GetSourceLocationId returns the SourceLocationId field if non-nil, zero value otherwise.

### GetSourceLocationIdOk

`func (o *AssetTransferCreateDto) GetSourceLocationIdOk() (*string, bool)`

GetSourceLocationIdOk returns a tuple with the SourceLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLocationId

`func (o *AssetTransferCreateDto) SetSourceLocationId(v string)`

SetSourceLocationId sets SourceLocationId field to given value.

### HasSourceLocationId

`func (o *AssetTransferCreateDto) HasSourceLocationId() bool`

HasSourceLocationId returns a boolean if a field has been set.

### SetSourceLocationIdNil

`func (o *AssetTransferCreateDto) SetSourceLocationIdNil(b bool)`

 SetSourceLocationIdNil sets the value for SourceLocationId to be an explicit nil

### UnsetSourceLocationId
`func (o *AssetTransferCreateDto) UnsetSourceLocationId()`

UnsetSourceLocationId ensures that no value is present for SourceLocationId, not even an explicit nil
### GetDestinationLocationId

`func (o *AssetTransferCreateDto) GetDestinationLocationId() string`

GetDestinationLocationId returns the DestinationLocationId field if non-nil, zero value otherwise.

### GetDestinationLocationIdOk

`func (o *AssetTransferCreateDto) GetDestinationLocationIdOk() (*string, bool)`

GetDestinationLocationIdOk returns a tuple with the DestinationLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLocationId

`func (o *AssetTransferCreateDto) SetDestinationLocationId(v string)`

SetDestinationLocationId sets DestinationLocationId field to given value.

### HasDestinationLocationId

`func (o *AssetTransferCreateDto) HasDestinationLocationId() bool`

HasDestinationLocationId returns a boolean if a field has been set.

### SetDestinationLocationIdNil

`func (o *AssetTransferCreateDto) SetDestinationLocationIdNil(b bool)`

 SetDestinationLocationIdNil sets the value for DestinationLocationId to be an explicit nil

### UnsetDestinationLocationId
`func (o *AssetTransferCreateDto) UnsetDestinationLocationId()`

UnsetDestinationLocationId ensures that no value is present for DestinationLocationId, not even an explicit nil
### GetSourceContactId

`func (o *AssetTransferCreateDto) GetSourceContactId() string`

GetSourceContactId returns the SourceContactId field if non-nil, zero value otherwise.

### GetSourceContactIdOk

`func (o *AssetTransferCreateDto) GetSourceContactIdOk() (*string, bool)`

GetSourceContactIdOk returns a tuple with the SourceContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceContactId

`func (o *AssetTransferCreateDto) SetSourceContactId(v string)`

SetSourceContactId sets SourceContactId field to given value.

### HasSourceContactId

`func (o *AssetTransferCreateDto) HasSourceContactId() bool`

HasSourceContactId returns a boolean if a field has been set.

### SetSourceContactIdNil

`func (o *AssetTransferCreateDto) SetSourceContactIdNil(b bool)`

 SetSourceContactIdNil sets the value for SourceContactId to be an explicit nil

### UnsetSourceContactId
`func (o *AssetTransferCreateDto) UnsetSourceContactId()`

UnsetSourceContactId ensures that no value is present for SourceContactId, not even an explicit nil
### GetDestinationContactId

`func (o *AssetTransferCreateDto) GetDestinationContactId() string`

GetDestinationContactId returns the DestinationContactId field if non-nil, zero value otherwise.

### GetDestinationContactIdOk

`func (o *AssetTransferCreateDto) GetDestinationContactIdOk() (*string, bool)`

GetDestinationContactIdOk returns a tuple with the DestinationContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationContactId

`func (o *AssetTransferCreateDto) SetDestinationContactId(v string)`

SetDestinationContactId sets DestinationContactId field to given value.

### HasDestinationContactId

`func (o *AssetTransferCreateDto) HasDestinationContactId() bool`

HasDestinationContactId returns a boolean if a field has been set.

### SetDestinationContactIdNil

`func (o *AssetTransferCreateDto) SetDestinationContactIdNil(b bool)`

 SetDestinationContactIdNil sets the value for DestinationContactId to be an explicit nil

### UnsetDestinationContactId
`func (o *AssetTransferCreateDto) UnsetDestinationContactId()`

UnsetDestinationContactId ensures that no value is present for DestinationContactId, not even an explicit nil
### GetSourceDepartmentId

`func (o *AssetTransferCreateDto) GetSourceDepartmentId() string`

GetSourceDepartmentId returns the SourceDepartmentId field if non-nil, zero value otherwise.

### GetSourceDepartmentIdOk

`func (o *AssetTransferCreateDto) GetSourceDepartmentIdOk() (*string, bool)`

GetSourceDepartmentIdOk returns a tuple with the SourceDepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDepartmentId

`func (o *AssetTransferCreateDto) SetSourceDepartmentId(v string)`

SetSourceDepartmentId sets SourceDepartmentId field to given value.

### HasSourceDepartmentId

`func (o *AssetTransferCreateDto) HasSourceDepartmentId() bool`

HasSourceDepartmentId returns a boolean if a field has been set.

### SetSourceDepartmentIdNil

`func (o *AssetTransferCreateDto) SetSourceDepartmentIdNil(b bool)`

 SetSourceDepartmentIdNil sets the value for SourceDepartmentId to be an explicit nil

### UnsetSourceDepartmentId
`func (o *AssetTransferCreateDto) UnsetSourceDepartmentId()`

UnsetSourceDepartmentId ensures that no value is present for SourceDepartmentId, not even an explicit nil
### GetDestinationDepartmentId

`func (o *AssetTransferCreateDto) GetDestinationDepartmentId() string`

GetDestinationDepartmentId returns the DestinationDepartmentId field if non-nil, zero value otherwise.

### GetDestinationDepartmentIdOk

`func (o *AssetTransferCreateDto) GetDestinationDepartmentIdOk() (*string, bool)`

GetDestinationDepartmentIdOk returns a tuple with the DestinationDepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDepartmentId

`func (o *AssetTransferCreateDto) SetDestinationDepartmentId(v string)`

SetDestinationDepartmentId sets DestinationDepartmentId field to given value.

### HasDestinationDepartmentId

`func (o *AssetTransferCreateDto) HasDestinationDepartmentId() bool`

HasDestinationDepartmentId returns a boolean if a field has been set.

### SetDestinationDepartmentIdNil

`func (o *AssetTransferCreateDto) SetDestinationDepartmentIdNil(b bool)`

 SetDestinationDepartmentIdNil sets the value for DestinationDepartmentId to be an explicit nil

### UnsetDestinationDepartmentId
`func (o *AssetTransferCreateDto) UnsetDestinationDepartmentId()`

UnsetDestinationDepartmentId ensures that no value is present for DestinationDepartmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


