# AssetTransferDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessId** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordId** | Pointer to **NullableString** |  | [optional] 
**AssetId** | Pointer to **NullableString** |  | [optional] 
**AssetName** | Pointer to **NullableString** |  | [optional] 
**IsRootTransfer** | Pointer to **bool** |  | [optional] 
**SerialList** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **NullableString** |  | [optional] 
**Serial** | Pointer to **NullableString** |  | [optional] 
**PreviousAssetTransferId** | Pointer to **NullableString** |  | [optional] 
**SourceLocationId** | Pointer to **NullableString** |  | [optional] 
**SourceLocationName** | Pointer to **NullableString** |  | [optional] 
**DestinationLocationId** | Pointer to **NullableString** |  | [optional] 
**DestinationLocationName** | Pointer to **NullableString** |  | [optional] 
**SourceContactId** | Pointer to **NullableString** |  | [optional] 
**SourceContactName** | Pointer to **NullableString** |  | [optional] 
**DestinationContactId** | Pointer to **NullableString** |  | [optional] 
**DestinationContactName** | Pointer to **NullableString** |  | [optional] 
**SourceDepartmentId** | Pointer to **NullableString** |  | [optional] 
**SourceDepartmentName** | Pointer to **NullableString** |  | [optional] 
**DestinationDepartmentId** | Pointer to **NullableString** |  | [optional] 
**DestinationDepartmentName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAssetTransferDto

`func NewAssetTransferDto() *AssetTransferDto`

NewAssetTransferDto instantiates a new AssetTransferDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetTransferDtoWithDefaults

`func NewAssetTransferDtoWithDefaults() *AssetTransferDto`

NewAssetTransferDtoWithDefaults instantiates a new AssetTransferDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssetTransferDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetTransferDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetTransferDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetTransferDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AssetTransferDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AssetTransferDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *AssetTransferDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AssetTransferDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AssetTransferDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AssetTransferDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBusinessId

`func (o *AssetTransferDto) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *AssetTransferDto) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *AssetTransferDto) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *AssetTransferDto) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### SetBusinessIdNil

`func (o *AssetTransferDto) SetBusinessIdNil(b bool)`

 SetBusinessIdNil sets the value for BusinessId to be an explicit nil

### UnsetBusinessId
`func (o *AssetTransferDto) UnsetBusinessId()`

UnsetBusinessId ensures that no value is present for BusinessId, not even an explicit nil
### GetBusinessProfileRecordId

`func (o *AssetTransferDto) GetBusinessProfileRecordId() string`

GetBusinessProfileRecordId returns the BusinessProfileRecordId field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIdOk

`func (o *AssetTransferDto) GetBusinessProfileRecordIdOk() (*string, bool)`

GetBusinessProfileRecordIdOk returns a tuple with the BusinessProfileRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordId

`func (o *AssetTransferDto) SetBusinessProfileRecordId(v string)`

SetBusinessProfileRecordId sets BusinessProfileRecordId field to given value.

### HasBusinessProfileRecordId

`func (o *AssetTransferDto) HasBusinessProfileRecordId() bool`

HasBusinessProfileRecordId returns a boolean if a field has been set.

### SetBusinessProfileRecordIdNil

`func (o *AssetTransferDto) SetBusinessProfileRecordIdNil(b bool)`

 SetBusinessProfileRecordIdNil sets the value for BusinessProfileRecordId to be an explicit nil

### UnsetBusinessProfileRecordId
`func (o *AssetTransferDto) UnsetBusinessProfileRecordId()`

UnsetBusinessProfileRecordId ensures that no value is present for BusinessProfileRecordId, not even an explicit nil
### GetAssetId

`func (o *AssetTransferDto) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetTransferDto) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetTransferDto) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AssetTransferDto) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### SetAssetIdNil

`func (o *AssetTransferDto) SetAssetIdNil(b bool)`

 SetAssetIdNil sets the value for AssetId to be an explicit nil

### UnsetAssetId
`func (o *AssetTransferDto) UnsetAssetId()`

UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
### GetAssetName

`func (o *AssetTransferDto) GetAssetName() string`

GetAssetName returns the AssetName field if non-nil, zero value otherwise.

### GetAssetNameOk

`func (o *AssetTransferDto) GetAssetNameOk() (*string, bool)`

GetAssetNameOk returns a tuple with the AssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetName

`func (o *AssetTransferDto) SetAssetName(v string)`

SetAssetName sets AssetName field to given value.

### HasAssetName

`func (o *AssetTransferDto) HasAssetName() bool`

HasAssetName returns a boolean if a field has been set.

### SetAssetNameNil

`func (o *AssetTransferDto) SetAssetNameNil(b bool)`

 SetAssetNameNil sets the value for AssetName to be an explicit nil

### UnsetAssetName
`func (o *AssetTransferDto) UnsetAssetName()`

UnsetAssetName ensures that no value is present for AssetName, not even an explicit nil
### GetIsRootTransfer

`func (o *AssetTransferDto) GetIsRootTransfer() bool`

GetIsRootTransfer returns the IsRootTransfer field if non-nil, zero value otherwise.

### GetIsRootTransferOk

`func (o *AssetTransferDto) GetIsRootTransferOk() (*bool, bool)`

GetIsRootTransferOk returns a tuple with the IsRootTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRootTransfer

`func (o *AssetTransferDto) SetIsRootTransfer(v bool)`

SetIsRootTransfer sets IsRootTransfer field to given value.

### HasIsRootTransfer

`func (o *AssetTransferDto) HasIsRootTransfer() bool`

HasIsRootTransfer returns a boolean if a field has been set.

### GetSerialList

`func (o *AssetTransferDto) GetSerialList() string`

GetSerialList returns the SerialList field if non-nil, zero value otherwise.

### GetSerialListOk

`func (o *AssetTransferDto) GetSerialListOk() (*string, bool)`

GetSerialListOk returns a tuple with the SerialList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialList

`func (o *AssetTransferDto) SetSerialList(v string)`

SetSerialList sets SerialList field to given value.

### HasSerialList

`func (o *AssetTransferDto) HasSerialList() bool`

HasSerialList returns a boolean if a field has been set.

### SetSerialListNil

`func (o *AssetTransferDto) SetSerialListNil(b bool)`

 SetSerialListNil sets the value for SerialList to be an explicit nil

### UnsetSerialList
`func (o *AssetTransferDto) UnsetSerialList()`

UnsetSerialList ensures that no value is present for SerialList, not even an explicit nil
### GetQuantity

`func (o *AssetTransferDto) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AssetTransferDto) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AssetTransferDto) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *AssetTransferDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *AssetTransferDto) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *AssetTransferDto) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetSerial

`func (o *AssetTransferDto) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *AssetTransferDto) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *AssetTransferDto) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *AssetTransferDto) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### SetSerialNil

`func (o *AssetTransferDto) SetSerialNil(b bool)`

 SetSerialNil sets the value for Serial to be an explicit nil

### UnsetSerial
`func (o *AssetTransferDto) UnsetSerial()`

UnsetSerial ensures that no value is present for Serial, not even an explicit nil
### GetPreviousAssetTransferId

`func (o *AssetTransferDto) GetPreviousAssetTransferId() string`

GetPreviousAssetTransferId returns the PreviousAssetTransferId field if non-nil, zero value otherwise.

### GetPreviousAssetTransferIdOk

`func (o *AssetTransferDto) GetPreviousAssetTransferIdOk() (*string, bool)`

GetPreviousAssetTransferIdOk returns a tuple with the PreviousAssetTransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAssetTransferId

`func (o *AssetTransferDto) SetPreviousAssetTransferId(v string)`

SetPreviousAssetTransferId sets PreviousAssetTransferId field to given value.

### HasPreviousAssetTransferId

`func (o *AssetTransferDto) HasPreviousAssetTransferId() bool`

HasPreviousAssetTransferId returns a boolean if a field has been set.

### SetPreviousAssetTransferIdNil

`func (o *AssetTransferDto) SetPreviousAssetTransferIdNil(b bool)`

 SetPreviousAssetTransferIdNil sets the value for PreviousAssetTransferId to be an explicit nil

### UnsetPreviousAssetTransferId
`func (o *AssetTransferDto) UnsetPreviousAssetTransferId()`

UnsetPreviousAssetTransferId ensures that no value is present for PreviousAssetTransferId, not even an explicit nil
### GetSourceLocationId

`func (o *AssetTransferDto) GetSourceLocationId() string`

GetSourceLocationId returns the SourceLocationId field if non-nil, zero value otherwise.

### GetSourceLocationIdOk

`func (o *AssetTransferDto) GetSourceLocationIdOk() (*string, bool)`

GetSourceLocationIdOk returns a tuple with the SourceLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLocationId

`func (o *AssetTransferDto) SetSourceLocationId(v string)`

SetSourceLocationId sets SourceLocationId field to given value.

### HasSourceLocationId

`func (o *AssetTransferDto) HasSourceLocationId() bool`

HasSourceLocationId returns a boolean if a field has been set.

### SetSourceLocationIdNil

`func (o *AssetTransferDto) SetSourceLocationIdNil(b bool)`

 SetSourceLocationIdNil sets the value for SourceLocationId to be an explicit nil

### UnsetSourceLocationId
`func (o *AssetTransferDto) UnsetSourceLocationId()`

UnsetSourceLocationId ensures that no value is present for SourceLocationId, not even an explicit nil
### GetSourceLocationName

`func (o *AssetTransferDto) GetSourceLocationName() string`

GetSourceLocationName returns the SourceLocationName field if non-nil, zero value otherwise.

### GetSourceLocationNameOk

`func (o *AssetTransferDto) GetSourceLocationNameOk() (*string, bool)`

GetSourceLocationNameOk returns a tuple with the SourceLocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLocationName

`func (o *AssetTransferDto) SetSourceLocationName(v string)`

SetSourceLocationName sets SourceLocationName field to given value.

### HasSourceLocationName

`func (o *AssetTransferDto) HasSourceLocationName() bool`

HasSourceLocationName returns a boolean if a field has been set.

### SetSourceLocationNameNil

`func (o *AssetTransferDto) SetSourceLocationNameNil(b bool)`

 SetSourceLocationNameNil sets the value for SourceLocationName to be an explicit nil

### UnsetSourceLocationName
`func (o *AssetTransferDto) UnsetSourceLocationName()`

UnsetSourceLocationName ensures that no value is present for SourceLocationName, not even an explicit nil
### GetDestinationLocationId

`func (o *AssetTransferDto) GetDestinationLocationId() string`

GetDestinationLocationId returns the DestinationLocationId field if non-nil, zero value otherwise.

### GetDestinationLocationIdOk

`func (o *AssetTransferDto) GetDestinationLocationIdOk() (*string, bool)`

GetDestinationLocationIdOk returns a tuple with the DestinationLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLocationId

`func (o *AssetTransferDto) SetDestinationLocationId(v string)`

SetDestinationLocationId sets DestinationLocationId field to given value.

### HasDestinationLocationId

`func (o *AssetTransferDto) HasDestinationLocationId() bool`

HasDestinationLocationId returns a boolean if a field has been set.

### SetDestinationLocationIdNil

`func (o *AssetTransferDto) SetDestinationLocationIdNil(b bool)`

 SetDestinationLocationIdNil sets the value for DestinationLocationId to be an explicit nil

### UnsetDestinationLocationId
`func (o *AssetTransferDto) UnsetDestinationLocationId()`

UnsetDestinationLocationId ensures that no value is present for DestinationLocationId, not even an explicit nil
### GetDestinationLocationName

`func (o *AssetTransferDto) GetDestinationLocationName() string`

GetDestinationLocationName returns the DestinationLocationName field if non-nil, zero value otherwise.

### GetDestinationLocationNameOk

`func (o *AssetTransferDto) GetDestinationLocationNameOk() (*string, bool)`

GetDestinationLocationNameOk returns a tuple with the DestinationLocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLocationName

`func (o *AssetTransferDto) SetDestinationLocationName(v string)`

SetDestinationLocationName sets DestinationLocationName field to given value.

### HasDestinationLocationName

`func (o *AssetTransferDto) HasDestinationLocationName() bool`

HasDestinationLocationName returns a boolean if a field has been set.

### SetDestinationLocationNameNil

`func (o *AssetTransferDto) SetDestinationLocationNameNil(b bool)`

 SetDestinationLocationNameNil sets the value for DestinationLocationName to be an explicit nil

### UnsetDestinationLocationName
`func (o *AssetTransferDto) UnsetDestinationLocationName()`

UnsetDestinationLocationName ensures that no value is present for DestinationLocationName, not even an explicit nil
### GetSourceContactId

`func (o *AssetTransferDto) GetSourceContactId() string`

GetSourceContactId returns the SourceContactId field if non-nil, zero value otherwise.

### GetSourceContactIdOk

`func (o *AssetTransferDto) GetSourceContactIdOk() (*string, bool)`

GetSourceContactIdOk returns a tuple with the SourceContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceContactId

`func (o *AssetTransferDto) SetSourceContactId(v string)`

SetSourceContactId sets SourceContactId field to given value.

### HasSourceContactId

`func (o *AssetTransferDto) HasSourceContactId() bool`

HasSourceContactId returns a boolean if a field has been set.

### SetSourceContactIdNil

`func (o *AssetTransferDto) SetSourceContactIdNil(b bool)`

 SetSourceContactIdNil sets the value for SourceContactId to be an explicit nil

### UnsetSourceContactId
`func (o *AssetTransferDto) UnsetSourceContactId()`

UnsetSourceContactId ensures that no value is present for SourceContactId, not even an explicit nil
### GetSourceContactName

`func (o *AssetTransferDto) GetSourceContactName() string`

GetSourceContactName returns the SourceContactName field if non-nil, zero value otherwise.

### GetSourceContactNameOk

`func (o *AssetTransferDto) GetSourceContactNameOk() (*string, bool)`

GetSourceContactNameOk returns a tuple with the SourceContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceContactName

`func (o *AssetTransferDto) SetSourceContactName(v string)`

SetSourceContactName sets SourceContactName field to given value.

### HasSourceContactName

`func (o *AssetTransferDto) HasSourceContactName() bool`

HasSourceContactName returns a boolean if a field has been set.

### SetSourceContactNameNil

`func (o *AssetTransferDto) SetSourceContactNameNil(b bool)`

 SetSourceContactNameNil sets the value for SourceContactName to be an explicit nil

### UnsetSourceContactName
`func (o *AssetTransferDto) UnsetSourceContactName()`

UnsetSourceContactName ensures that no value is present for SourceContactName, not even an explicit nil
### GetDestinationContactId

`func (o *AssetTransferDto) GetDestinationContactId() string`

GetDestinationContactId returns the DestinationContactId field if non-nil, zero value otherwise.

### GetDestinationContactIdOk

`func (o *AssetTransferDto) GetDestinationContactIdOk() (*string, bool)`

GetDestinationContactIdOk returns a tuple with the DestinationContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationContactId

`func (o *AssetTransferDto) SetDestinationContactId(v string)`

SetDestinationContactId sets DestinationContactId field to given value.

### HasDestinationContactId

`func (o *AssetTransferDto) HasDestinationContactId() bool`

HasDestinationContactId returns a boolean if a field has been set.

### SetDestinationContactIdNil

`func (o *AssetTransferDto) SetDestinationContactIdNil(b bool)`

 SetDestinationContactIdNil sets the value for DestinationContactId to be an explicit nil

### UnsetDestinationContactId
`func (o *AssetTransferDto) UnsetDestinationContactId()`

UnsetDestinationContactId ensures that no value is present for DestinationContactId, not even an explicit nil
### GetDestinationContactName

`func (o *AssetTransferDto) GetDestinationContactName() string`

GetDestinationContactName returns the DestinationContactName field if non-nil, zero value otherwise.

### GetDestinationContactNameOk

`func (o *AssetTransferDto) GetDestinationContactNameOk() (*string, bool)`

GetDestinationContactNameOk returns a tuple with the DestinationContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationContactName

`func (o *AssetTransferDto) SetDestinationContactName(v string)`

SetDestinationContactName sets DestinationContactName field to given value.

### HasDestinationContactName

`func (o *AssetTransferDto) HasDestinationContactName() bool`

HasDestinationContactName returns a boolean if a field has been set.

### SetDestinationContactNameNil

`func (o *AssetTransferDto) SetDestinationContactNameNil(b bool)`

 SetDestinationContactNameNil sets the value for DestinationContactName to be an explicit nil

### UnsetDestinationContactName
`func (o *AssetTransferDto) UnsetDestinationContactName()`

UnsetDestinationContactName ensures that no value is present for DestinationContactName, not even an explicit nil
### GetSourceDepartmentId

`func (o *AssetTransferDto) GetSourceDepartmentId() string`

GetSourceDepartmentId returns the SourceDepartmentId field if non-nil, zero value otherwise.

### GetSourceDepartmentIdOk

`func (o *AssetTransferDto) GetSourceDepartmentIdOk() (*string, bool)`

GetSourceDepartmentIdOk returns a tuple with the SourceDepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDepartmentId

`func (o *AssetTransferDto) SetSourceDepartmentId(v string)`

SetSourceDepartmentId sets SourceDepartmentId field to given value.

### HasSourceDepartmentId

`func (o *AssetTransferDto) HasSourceDepartmentId() bool`

HasSourceDepartmentId returns a boolean if a field has been set.

### SetSourceDepartmentIdNil

`func (o *AssetTransferDto) SetSourceDepartmentIdNil(b bool)`

 SetSourceDepartmentIdNil sets the value for SourceDepartmentId to be an explicit nil

### UnsetSourceDepartmentId
`func (o *AssetTransferDto) UnsetSourceDepartmentId()`

UnsetSourceDepartmentId ensures that no value is present for SourceDepartmentId, not even an explicit nil
### GetSourceDepartmentName

`func (o *AssetTransferDto) GetSourceDepartmentName() string`

GetSourceDepartmentName returns the SourceDepartmentName field if non-nil, zero value otherwise.

### GetSourceDepartmentNameOk

`func (o *AssetTransferDto) GetSourceDepartmentNameOk() (*string, bool)`

GetSourceDepartmentNameOk returns a tuple with the SourceDepartmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDepartmentName

`func (o *AssetTransferDto) SetSourceDepartmentName(v string)`

SetSourceDepartmentName sets SourceDepartmentName field to given value.

### HasSourceDepartmentName

`func (o *AssetTransferDto) HasSourceDepartmentName() bool`

HasSourceDepartmentName returns a boolean if a field has been set.

### SetSourceDepartmentNameNil

`func (o *AssetTransferDto) SetSourceDepartmentNameNil(b bool)`

 SetSourceDepartmentNameNil sets the value for SourceDepartmentName to be an explicit nil

### UnsetSourceDepartmentName
`func (o *AssetTransferDto) UnsetSourceDepartmentName()`

UnsetSourceDepartmentName ensures that no value is present for SourceDepartmentName, not even an explicit nil
### GetDestinationDepartmentId

`func (o *AssetTransferDto) GetDestinationDepartmentId() string`

GetDestinationDepartmentId returns the DestinationDepartmentId field if non-nil, zero value otherwise.

### GetDestinationDepartmentIdOk

`func (o *AssetTransferDto) GetDestinationDepartmentIdOk() (*string, bool)`

GetDestinationDepartmentIdOk returns a tuple with the DestinationDepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDepartmentId

`func (o *AssetTransferDto) SetDestinationDepartmentId(v string)`

SetDestinationDepartmentId sets DestinationDepartmentId field to given value.

### HasDestinationDepartmentId

`func (o *AssetTransferDto) HasDestinationDepartmentId() bool`

HasDestinationDepartmentId returns a boolean if a field has been set.

### SetDestinationDepartmentIdNil

`func (o *AssetTransferDto) SetDestinationDepartmentIdNil(b bool)`

 SetDestinationDepartmentIdNil sets the value for DestinationDepartmentId to be an explicit nil

### UnsetDestinationDepartmentId
`func (o *AssetTransferDto) UnsetDestinationDepartmentId()`

UnsetDestinationDepartmentId ensures that no value is present for DestinationDepartmentId, not even an explicit nil
### GetDestinationDepartmentName

`func (o *AssetTransferDto) GetDestinationDepartmentName() string`

GetDestinationDepartmentName returns the DestinationDepartmentName field if non-nil, zero value otherwise.

### GetDestinationDepartmentNameOk

`func (o *AssetTransferDto) GetDestinationDepartmentNameOk() (*string, bool)`

GetDestinationDepartmentNameOk returns a tuple with the DestinationDepartmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDepartmentName

`func (o *AssetTransferDto) SetDestinationDepartmentName(v string)`

SetDestinationDepartmentName sets DestinationDepartmentName field to given value.

### HasDestinationDepartmentName

`func (o *AssetTransferDto) HasDestinationDepartmentName() bool`

HasDestinationDepartmentName returns a boolean if a field has been set.

### SetDestinationDepartmentNameNil

`func (o *AssetTransferDto) SetDestinationDepartmentNameNil(b bool)`

 SetDestinationDepartmentNameNil sets the value for DestinationDepartmentName to be an explicit nil

### UnsetDestinationDepartmentName
`func (o *AssetTransferDto) UnsetDestinationDepartmentName()`

UnsetDestinationDepartmentName ensures that no value is present for DestinationDepartmentName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


