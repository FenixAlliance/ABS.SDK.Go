# ProofOfDeliveryCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**DocumentNumber** | Pointer to **NullableString** |  | [optional] 
**ShipmentId** | Pointer to **NullableString** |  | [optional] 
**BillOfLadingId** | Pointer to **NullableString** |  | [optional] 
**SeawayBillId** | Pointer to **NullableString** |  | [optional] 
**AirwayBillId** | Pointer to **NullableString** |  | [optional] 
**RoadWaybillId** | Pointer to **NullableString** |  | [optional] 
**RailWaybillId** | Pointer to **NullableString** |  | [optional] 
**TruckTripId** | Pointer to **NullableString** |  | [optional] 
**RecipientName** | Pointer to **NullableString** |  | [optional] 
**RecipientCompanyContactId** | Pointer to **NullableString** |  | [optional] 
**DeliveryAddress** | Pointer to **NullableString** |  | [optional] 
**DeliveryDate** | Pointer to **NullableTime** |  | [optional] 
**DeliveryTime** | Pointer to **NullableString** |  | [optional] 
**OverallCondition** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProofOfDeliveryCreateDto

`func NewProofOfDeliveryCreateDto() *ProofOfDeliveryCreateDto`

NewProofOfDeliveryCreateDto instantiates a new ProofOfDeliveryCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProofOfDeliveryCreateDtoWithDefaults

`func NewProofOfDeliveryCreateDtoWithDefaults() *ProofOfDeliveryCreateDto`

NewProofOfDeliveryCreateDtoWithDefaults instantiates a new ProofOfDeliveryCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProofOfDeliveryCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProofOfDeliveryCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProofOfDeliveryCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProofOfDeliveryCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ProofOfDeliveryCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProofOfDeliveryCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProofOfDeliveryCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProofOfDeliveryCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDocumentNumber

`func (o *ProofOfDeliveryCreateDto) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *ProofOfDeliveryCreateDto) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *ProofOfDeliveryCreateDto) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *ProofOfDeliveryCreateDto) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *ProofOfDeliveryCreateDto) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *ProofOfDeliveryCreateDto) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetShipmentId

`func (o *ProofOfDeliveryCreateDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *ProofOfDeliveryCreateDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *ProofOfDeliveryCreateDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *ProofOfDeliveryCreateDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *ProofOfDeliveryCreateDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *ProofOfDeliveryCreateDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetBillOfLadingId

`func (o *ProofOfDeliveryCreateDto) GetBillOfLadingId() string`

GetBillOfLadingId returns the BillOfLadingId field if non-nil, zero value otherwise.

### GetBillOfLadingIdOk

`func (o *ProofOfDeliveryCreateDto) GetBillOfLadingIdOk() (*string, bool)`

GetBillOfLadingIdOk returns a tuple with the BillOfLadingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingId

`func (o *ProofOfDeliveryCreateDto) SetBillOfLadingId(v string)`

SetBillOfLadingId sets BillOfLadingId field to given value.

### HasBillOfLadingId

`func (o *ProofOfDeliveryCreateDto) HasBillOfLadingId() bool`

HasBillOfLadingId returns a boolean if a field has been set.

### SetBillOfLadingIdNil

`func (o *ProofOfDeliveryCreateDto) SetBillOfLadingIdNil(b bool)`

 SetBillOfLadingIdNil sets the value for BillOfLadingId to be an explicit nil

### UnsetBillOfLadingId
`func (o *ProofOfDeliveryCreateDto) UnsetBillOfLadingId()`

UnsetBillOfLadingId ensures that no value is present for BillOfLadingId, not even an explicit nil
### GetSeawayBillId

`func (o *ProofOfDeliveryCreateDto) GetSeawayBillId() string`

GetSeawayBillId returns the SeawayBillId field if non-nil, zero value otherwise.

### GetSeawayBillIdOk

`func (o *ProofOfDeliveryCreateDto) GetSeawayBillIdOk() (*string, bool)`

GetSeawayBillIdOk returns a tuple with the SeawayBillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeawayBillId

`func (o *ProofOfDeliveryCreateDto) SetSeawayBillId(v string)`

SetSeawayBillId sets SeawayBillId field to given value.

### HasSeawayBillId

`func (o *ProofOfDeliveryCreateDto) HasSeawayBillId() bool`

HasSeawayBillId returns a boolean if a field has been set.

### SetSeawayBillIdNil

`func (o *ProofOfDeliveryCreateDto) SetSeawayBillIdNil(b bool)`

 SetSeawayBillIdNil sets the value for SeawayBillId to be an explicit nil

### UnsetSeawayBillId
`func (o *ProofOfDeliveryCreateDto) UnsetSeawayBillId()`

UnsetSeawayBillId ensures that no value is present for SeawayBillId, not even an explicit nil
### GetAirwayBillId

`func (o *ProofOfDeliveryCreateDto) GetAirwayBillId() string`

GetAirwayBillId returns the AirwayBillId field if non-nil, zero value otherwise.

### GetAirwayBillIdOk

`func (o *ProofOfDeliveryCreateDto) GetAirwayBillIdOk() (*string, bool)`

GetAirwayBillIdOk returns a tuple with the AirwayBillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirwayBillId

`func (o *ProofOfDeliveryCreateDto) SetAirwayBillId(v string)`

SetAirwayBillId sets AirwayBillId field to given value.

### HasAirwayBillId

`func (o *ProofOfDeliveryCreateDto) HasAirwayBillId() bool`

HasAirwayBillId returns a boolean if a field has been set.

### SetAirwayBillIdNil

`func (o *ProofOfDeliveryCreateDto) SetAirwayBillIdNil(b bool)`

 SetAirwayBillIdNil sets the value for AirwayBillId to be an explicit nil

### UnsetAirwayBillId
`func (o *ProofOfDeliveryCreateDto) UnsetAirwayBillId()`

UnsetAirwayBillId ensures that no value is present for AirwayBillId, not even an explicit nil
### GetRoadWaybillId

`func (o *ProofOfDeliveryCreateDto) GetRoadWaybillId() string`

GetRoadWaybillId returns the RoadWaybillId field if non-nil, zero value otherwise.

### GetRoadWaybillIdOk

`func (o *ProofOfDeliveryCreateDto) GetRoadWaybillIdOk() (*string, bool)`

GetRoadWaybillIdOk returns a tuple with the RoadWaybillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoadWaybillId

`func (o *ProofOfDeliveryCreateDto) SetRoadWaybillId(v string)`

SetRoadWaybillId sets RoadWaybillId field to given value.

### HasRoadWaybillId

`func (o *ProofOfDeliveryCreateDto) HasRoadWaybillId() bool`

HasRoadWaybillId returns a boolean if a field has been set.

### SetRoadWaybillIdNil

`func (o *ProofOfDeliveryCreateDto) SetRoadWaybillIdNil(b bool)`

 SetRoadWaybillIdNil sets the value for RoadWaybillId to be an explicit nil

### UnsetRoadWaybillId
`func (o *ProofOfDeliveryCreateDto) UnsetRoadWaybillId()`

UnsetRoadWaybillId ensures that no value is present for RoadWaybillId, not even an explicit nil
### GetRailWaybillId

`func (o *ProofOfDeliveryCreateDto) GetRailWaybillId() string`

GetRailWaybillId returns the RailWaybillId field if non-nil, zero value otherwise.

### GetRailWaybillIdOk

`func (o *ProofOfDeliveryCreateDto) GetRailWaybillIdOk() (*string, bool)`

GetRailWaybillIdOk returns a tuple with the RailWaybillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailWaybillId

`func (o *ProofOfDeliveryCreateDto) SetRailWaybillId(v string)`

SetRailWaybillId sets RailWaybillId field to given value.

### HasRailWaybillId

`func (o *ProofOfDeliveryCreateDto) HasRailWaybillId() bool`

HasRailWaybillId returns a boolean if a field has been set.

### SetRailWaybillIdNil

`func (o *ProofOfDeliveryCreateDto) SetRailWaybillIdNil(b bool)`

 SetRailWaybillIdNil sets the value for RailWaybillId to be an explicit nil

### UnsetRailWaybillId
`func (o *ProofOfDeliveryCreateDto) UnsetRailWaybillId()`

UnsetRailWaybillId ensures that no value is present for RailWaybillId, not even an explicit nil
### GetTruckTripId

`func (o *ProofOfDeliveryCreateDto) GetTruckTripId() string`

GetTruckTripId returns the TruckTripId field if non-nil, zero value otherwise.

### GetTruckTripIdOk

`func (o *ProofOfDeliveryCreateDto) GetTruckTripIdOk() (*string, bool)`

GetTruckTripIdOk returns a tuple with the TruckTripId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckTripId

`func (o *ProofOfDeliveryCreateDto) SetTruckTripId(v string)`

SetTruckTripId sets TruckTripId field to given value.

### HasTruckTripId

`func (o *ProofOfDeliveryCreateDto) HasTruckTripId() bool`

HasTruckTripId returns a boolean if a field has been set.

### SetTruckTripIdNil

`func (o *ProofOfDeliveryCreateDto) SetTruckTripIdNil(b bool)`

 SetTruckTripIdNil sets the value for TruckTripId to be an explicit nil

### UnsetTruckTripId
`func (o *ProofOfDeliveryCreateDto) UnsetTruckTripId()`

UnsetTruckTripId ensures that no value is present for TruckTripId, not even an explicit nil
### GetRecipientName

`func (o *ProofOfDeliveryCreateDto) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *ProofOfDeliveryCreateDto) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *ProofOfDeliveryCreateDto) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.

### HasRecipientName

`func (o *ProofOfDeliveryCreateDto) HasRecipientName() bool`

HasRecipientName returns a boolean if a field has been set.

### SetRecipientNameNil

`func (o *ProofOfDeliveryCreateDto) SetRecipientNameNil(b bool)`

 SetRecipientNameNil sets the value for RecipientName to be an explicit nil

### UnsetRecipientName
`func (o *ProofOfDeliveryCreateDto) UnsetRecipientName()`

UnsetRecipientName ensures that no value is present for RecipientName, not even an explicit nil
### GetRecipientCompanyContactId

`func (o *ProofOfDeliveryCreateDto) GetRecipientCompanyContactId() string`

GetRecipientCompanyContactId returns the RecipientCompanyContactId field if non-nil, zero value otherwise.

### GetRecipientCompanyContactIdOk

`func (o *ProofOfDeliveryCreateDto) GetRecipientCompanyContactIdOk() (*string, bool)`

GetRecipientCompanyContactIdOk returns a tuple with the RecipientCompanyContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientCompanyContactId

`func (o *ProofOfDeliveryCreateDto) SetRecipientCompanyContactId(v string)`

SetRecipientCompanyContactId sets RecipientCompanyContactId field to given value.

### HasRecipientCompanyContactId

`func (o *ProofOfDeliveryCreateDto) HasRecipientCompanyContactId() bool`

HasRecipientCompanyContactId returns a boolean if a field has been set.

### SetRecipientCompanyContactIdNil

`func (o *ProofOfDeliveryCreateDto) SetRecipientCompanyContactIdNil(b bool)`

 SetRecipientCompanyContactIdNil sets the value for RecipientCompanyContactId to be an explicit nil

### UnsetRecipientCompanyContactId
`func (o *ProofOfDeliveryCreateDto) UnsetRecipientCompanyContactId()`

UnsetRecipientCompanyContactId ensures that no value is present for RecipientCompanyContactId, not even an explicit nil
### GetDeliveryAddress

`func (o *ProofOfDeliveryCreateDto) GetDeliveryAddress() string`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *ProofOfDeliveryCreateDto) GetDeliveryAddressOk() (*string, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *ProofOfDeliveryCreateDto) SetDeliveryAddress(v string)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *ProofOfDeliveryCreateDto) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### SetDeliveryAddressNil

`func (o *ProofOfDeliveryCreateDto) SetDeliveryAddressNil(b bool)`

 SetDeliveryAddressNil sets the value for DeliveryAddress to be an explicit nil

### UnsetDeliveryAddress
`func (o *ProofOfDeliveryCreateDto) UnsetDeliveryAddress()`

UnsetDeliveryAddress ensures that no value is present for DeliveryAddress, not even an explicit nil
### GetDeliveryDate

`func (o *ProofOfDeliveryCreateDto) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *ProofOfDeliveryCreateDto) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *ProofOfDeliveryCreateDto) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *ProofOfDeliveryCreateDto) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### SetDeliveryDateNil

`func (o *ProofOfDeliveryCreateDto) SetDeliveryDateNil(b bool)`

 SetDeliveryDateNil sets the value for DeliveryDate to be an explicit nil

### UnsetDeliveryDate
`func (o *ProofOfDeliveryCreateDto) UnsetDeliveryDate()`

UnsetDeliveryDate ensures that no value is present for DeliveryDate, not even an explicit nil
### GetDeliveryTime

`func (o *ProofOfDeliveryCreateDto) GetDeliveryTime() string`

GetDeliveryTime returns the DeliveryTime field if non-nil, zero value otherwise.

### GetDeliveryTimeOk

`func (o *ProofOfDeliveryCreateDto) GetDeliveryTimeOk() (*string, bool)`

GetDeliveryTimeOk returns a tuple with the DeliveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTime

`func (o *ProofOfDeliveryCreateDto) SetDeliveryTime(v string)`

SetDeliveryTime sets DeliveryTime field to given value.

### HasDeliveryTime

`func (o *ProofOfDeliveryCreateDto) HasDeliveryTime() bool`

HasDeliveryTime returns a boolean if a field has been set.

### SetDeliveryTimeNil

`func (o *ProofOfDeliveryCreateDto) SetDeliveryTimeNil(b bool)`

 SetDeliveryTimeNil sets the value for DeliveryTime to be an explicit nil

### UnsetDeliveryTime
`func (o *ProofOfDeliveryCreateDto) UnsetDeliveryTime()`

UnsetDeliveryTime ensures that no value is present for DeliveryTime, not even an explicit nil
### GetOverallCondition

`func (o *ProofOfDeliveryCreateDto) GetOverallCondition() string`

GetOverallCondition returns the OverallCondition field if non-nil, zero value otherwise.

### GetOverallConditionOk

`func (o *ProofOfDeliveryCreateDto) GetOverallConditionOk() (*string, bool)`

GetOverallConditionOk returns a tuple with the OverallCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallCondition

`func (o *ProofOfDeliveryCreateDto) SetOverallCondition(v string)`

SetOverallCondition sets OverallCondition field to given value.

### HasOverallCondition

`func (o *ProofOfDeliveryCreateDto) HasOverallCondition() bool`

HasOverallCondition returns a boolean if a field has been set.

### SetOverallConditionNil

`func (o *ProofOfDeliveryCreateDto) SetOverallConditionNil(b bool)`

 SetOverallConditionNil sets the value for OverallCondition to be an explicit nil

### UnsetOverallCondition
`func (o *ProofOfDeliveryCreateDto) UnsetOverallCondition()`

UnsetOverallCondition ensures that no value is present for OverallCondition, not even an explicit nil
### GetRemarks

`func (o *ProofOfDeliveryCreateDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ProofOfDeliveryCreateDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ProofOfDeliveryCreateDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ProofOfDeliveryCreateDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *ProofOfDeliveryCreateDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *ProofOfDeliveryCreateDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


