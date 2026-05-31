# ProofOfDeliveryUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**TotalQuantityDelivered** | Pointer to **NullableInt32** |  | [optional] 
**TotalQuantityRejected** | Pointer to **NullableInt32** |  | [optional] 
**Remarks** | Pointer to **NullableString** |  | [optional] 
**PhotoEvidenceUri** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProofOfDeliveryUpdateDto

`func NewProofOfDeliveryUpdateDto() *ProofOfDeliveryUpdateDto`

NewProofOfDeliveryUpdateDto instantiates a new ProofOfDeliveryUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProofOfDeliveryUpdateDtoWithDefaults

`func NewProofOfDeliveryUpdateDtoWithDefaults() *ProofOfDeliveryUpdateDto`

NewProofOfDeliveryUpdateDtoWithDefaults instantiates a new ProofOfDeliveryUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentNumber

`func (o *ProofOfDeliveryUpdateDto) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *ProofOfDeliveryUpdateDto) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *ProofOfDeliveryUpdateDto) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *ProofOfDeliveryUpdateDto) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *ProofOfDeliveryUpdateDto) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *ProofOfDeliveryUpdateDto) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetShipmentId

`func (o *ProofOfDeliveryUpdateDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *ProofOfDeliveryUpdateDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *ProofOfDeliveryUpdateDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *ProofOfDeliveryUpdateDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *ProofOfDeliveryUpdateDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *ProofOfDeliveryUpdateDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetBillOfLadingId

`func (o *ProofOfDeliveryUpdateDto) GetBillOfLadingId() string`

GetBillOfLadingId returns the BillOfLadingId field if non-nil, zero value otherwise.

### GetBillOfLadingIdOk

`func (o *ProofOfDeliveryUpdateDto) GetBillOfLadingIdOk() (*string, bool)`

GetBillOfLadingIdOk returns a tuple with the BillOfLadingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingId

`func (o *ProofOfDeliveryUpdateDto) SetBillOfLadingId(v string)`

SetBillOfLadingId sets BillOfLadingId field to given value.

### HasBillOfLadingId

`func (o *ProofOfDeliveryUpdateDto) HasBillOfLadingId() bool`

HasBillOfLadingId returns a boolean if a field has been set.

### SetBillOfLadingIdNil

`func (o *ProofOfDeliveryUpdateDto) SetBillOfLadingIdNil(b bool)`

 SetBillOfLadingIdNil sets the value for BillOfLadingId to be an explicit nil

### UnsetBillOfLadingId
`func (o *ProofOfDeliveryUpdateDto) UnsetBillOfLadingId()`

UnsetBillOfLadingId ensures that no value is present for BillOfLadingId, not even an explicit nil
### GetSeawayBillId

`func (o *ProofOfDeliveryUpdateDto) GetSeawayBillId() string`

GetSeawayBillId returns the SeawayBillId field if non-nil, zero value otherwise.

### GetSeawayBillIdOk

`func (o *ProofOfDeliveryUpdateDto) GetSeawayBillIdOk() (*string, bool)`

GetSeawayBillIdOk returns a tuple with the SeawayBillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeawayBillId

`func (o *ProofOfDeliveryUpdateDto) SetSeawayBillId(v string)`

SetSeawayBillId sets SeawayBillId field to given value.

### HasSeawayBillId

`func (o *ProofOfDeliveryUpdateDto) HasSeawayBillId() bool`

HasSeawayBillId returns a boolean if a field has been set.

### SetSeawayBillIdNil

`func (o *ProofOfDeliveryUpdateDto) SetSeawayBillIdNil(b bool)`

 SetSeawayBillIdNil sets the value for SeawayBillId to be an explicit nil

### UnsetSeawayBillId
`func (o *ProofOfDeliveryUpdateDto) UnsetSeawayBillId()`

UnsetSeawayBillId ensures that no value is present for SeawayBillId, not even an explicit nil
### GetAirwayBillId

`func (o *ProofOfDeliveryUpdateDto) GetAirwayBillId() string`

GetAirwayBillId returns the AirwayBillId field if non-nil, zero value otherwise.

### GetAirwayBillIdOk

`func (o *ProofOfDeliveryUpdateDto) GetAirwayBillIdOk() (*string, bool)`

GetAirwayBillIdOk returns a tuple with the AirwayBillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirwayBillId

`func (o *ProofOfDeliveryUpdateDto) SetAirwayBillId(v string)`

SetAirwayBillId sets AirwayBillId field to given value.

### HasAirwayBillId

`func (o *ProofOfDeliveryUpdateDto) HasAirwayBillId() bool`

HasAirwayBillId returns a boolean if a field has been set.

### SetAirwayBillIdNil

`func (o *ProofOfDeliveryUpdateDto) SetAirwayBillIdNil(b bool)`

 SetAirwayBillIdNil sets the value for AirwayBillId to be an explicit nil

### UnsetAirwayBillId
`func (o *ProofOfDeliveryUpdateDto) UnsetAirwayBillId()`

UnsetAirwayBillId ensures that no value is present for AirwayBillId, not even an explicit nil
### GetRoadWaybillId

`func (o *ProofOfDeliveryUpdateDto) GetRoadWaybillId() string`

GetRoadWaybillId returns the RoadWaybillId field if non-nil, zero value otherwise.

### GetRoadWaybillIdOk

`func (o *ProofOfDeliveryUpdateDto) GetRoadWaybillIdOk() (*string, bool)`

GetRoadWaybillIdOk returns a tuple with the RoadWaybillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoadWaybillId

`func (o *ProofOfDeliveryUpdateDto) SetRoadWaybillId(v string)`

SetRoadWaybillId sets RoadWaybillId field to given value.

### HasRoadWaybillId

`func (o *ProofOfDeliveryUpdateDto) HasRoadWaybillId() bool`

HasRoadWaybillId returns a boolean if a field has been set.

### SetRoadWaybillIdNil

`func (o *ProofOfDeliveryUpdateDto) SetRoadWaybillIdNil(b bool)`

 SetRoadWaybillIdNil sets the value for RoadWaybillId to be an explicit nil

### UnsetRoadWaybillId
`func (o *ProofOfDeliveryUpdateDto) UnsetRoadWaybillId()`

UnsetRoadWaybillId ensures that no value is present for RoadWaybillId, not even an explicit nil
### GetRailWaybillId

`func (o *ProofOfDeliveryUpdateDto) GetRailWaybillId() string`

GetRailWaybillId returns the RailWaybillId field if non-nil, zero value otherwise.

### GetRailWaybillIdOk

`func (o *ProofOfDeliveryUpdateDto) GetRailWaybillIdOk() (*string, bool)`

GetRailWaybillIdOk returns a tuple with the RailWaybillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailWaybillId

`func (o *ProofOfDeliveryUpdateDto) SetRailWaybillId(v string)`

SetRailWaybillId sets RailWaybillId field to given value.

### HasRailWaybillId

`func (o *ProofOfDeliveryUpdateDto) HasRailWaybillId() bool`

HasRailWaybillId returns a boolean if a field has been set.

### SetRailWaybillIdNil

`func (o *ProofOfDeliveryUpdateDto) SetRailWaybillIdNil(b bool)`

 SetRailWaybillIdNil sets the value for RailWaybillId to be an explicit nil

### UnsetRailWaybillId
`func (o *ProofOfDeliveryUpdateDto) UnsetRailWaybillId()`

UnsetRailWaybillId ensures that no value is present for RailWaybillId, not even an explicit nil
### GetTruckTripId

`func (o *ProofOfDeliveryUpdateDto) GetTruckTripId() string`

GetTruckTripId returns the TruckTripId field if non-nil, zero value otherwise.

### GetTruckTripIdOk

`func (o *ProofOfDeliveryUpdateDto) GetTruckTripIdOk() (*string, bool)`

GetTruckTripIdOk returns a tuple with the TruckTripId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckTripId

`func (o *ProofOfDeliveryUpdateDto) SetTruckTripId(v string)`

SetTruckTripId sets TruckTripId field to given value.

### HasTruckTripId

`func (o *ProofOfDeliveryUpdateDto) HasTruckTripId() bool`

HasTruckTripId returns a boolean if a field has been set.

### SetTruckTripIdNil

`func (o *ProofOfDeliveryUpdateDto) SetTruckTripIdNil(b bool)`

 SetTruckTripIdNil sets the value for TruckTripId to be an explicit nil

### UnsetTruckTripId
`func (o *ProofOfDeliveryUpdateDto) UnsetTruckTripId()`

UnsetTruckTripId ensures that no value is present for TruckTripId, not even an explicit nil
### GetRecipientName

`func (o *ProofOfDeliveryUpdateDto) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *ProofOfDeliveryUpdateDto) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *ProofOfDeliveryUpdateDto) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.

### HasRecipientName

`func (o *ProofOfDeliveryUpdateDto) HasRecipientName() bool`

HasRecipientName returns a boolean if a field has been set.

### SetRecipientNameNil

`func (o *ProofOfDeliveryUpdateDto) SetRecipientNameNil(b bool)`

 SetRecipientNameNil sets the value for RecipientName to be an explicit nil

### UnsetRecipientName
`func (o *ProofOfDeliveryUpdateDto) UnsetRecipientName()`

UnsetRecipientName ensures that no value is present for RecipientName, not even an explicit nil
### GetRecipientCompanyContactId

`func (o *ProofOfDeliveryUpdateDto) GetRecipientCompanyContactId() string`

GetRecipientCompanyContactId returns the RecipientCompanyContactId field if non-nil, zero value otherwise.

### GetRecipientCompanyContactIdOk

`func (o *ProofOfDeliveryUpdateDto) GetRecipientCompanyContactIdOk() (*string, bool)`

GetRecipientCompanyContactIdOk returns a tuple with the RecipientCompanyContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientCompanyContactId

`func (o *ProofOfDeliveryUpdateDto) SetRecipientCompanyContactId(v string)`

SetRecipientCompanyContactId sets RecipientCompanyContactId field to given value.

### HasRecipientCompanyContactId

`func (o *ProofOfDeliveryUpdateDto) HasRecipientCompanyContactId() bool`

HasRecipientCompanyContactId returns a boolean if a field has been set.

### SetRecipientCompanyContactIdNil

`func (o *ProofOfDeliveryUpdateDto) SetRecipientCompanyContactIdNil(b bool)`

 SetRecipientCompanyContactIdNil sets the value for RecipientCompanyContactId to be an explicit nil

### UnsetRecipientCompanyContactId
`func (o *ProofOfDeliveryUpdateDto) UnsetRecipientCompanyContactId()`

UnsetRecipientCompanyContactId ensures that no value is present for RecipientCompanyContactId, not even an explicit nil
### GetDeliveryAddress

`func (o *ProofOfDeliveryUpdateDto) GetDeliveryAddress() string`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *ProofOfDeliveryUpdateDto) GetDeliveryAddressOk() (*string, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *ProofOfDeliveryUpdateDto) SetDeliveryAddress(v string)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *ProofOfDeliveryUpdateDto) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### SetDeliveryAddressNil

`func (o *ProofOfDeliveryUpdateDto) SetDeliveryAddressNil(b bool)`

 SetDeliveryAddressNil sets the value for DeliveryAddress to be an explicit nil

### UnsetDeliveryAddress
`func (o *ProofOfDeliveryUpdateDto) UnsetDeliveryAddress()`

UnsetDeliveryAddress ensures that no value is present for DeliveryAddress, not even an explicit nil
### GetDeliveryDate

`func (o *ProofOfDeliveryUpdateDto) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *ProofOfDeliveryUpdateDto) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *ProofOfDeliveryUpdateDto) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *ProofOfDeliveryUpdateDto) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### SetDeliveryDateNil

`func (o *ProofOfDeliveryUpdateDto) SetDeliveryDateNil(b bool)`

 SetDeliveryDateNil sets the value for DeliveryDate to be an explicit nil

### UnsetDeliveryDate
`func (o *ProofOfDeliveryUpdateDto) UnsetDeliveryDate()`

UnsetDeliveryDate ensures that no value is present for DeliveryDate, not even an explicit nil
### GetDeliveryTime

`func (o *ProofOfDeliveryUpdateDto) GetDeliveryTime() string`

GetDeliveryTime returns the DeliveryTime field if non-nil, zero value otherwise.

### GetDeliveryTimeOk

`func (o *ProofOfDeliveryUpdateDto) GetDeliveryTimeOk() (*string, bool)`

GetDeliveryTimeOk returns a tuple with the DeliveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTime

`func (o *ProofOfDeliveryUpdateDto) SetDeliveryTime(v string)`

SetDeliveryTime sets DeliveryTime field to given value.

### HasDeliveryTime

`func (o *ProofOfDeliveryUpdateDto) HasDeliveryTime() bool`

HasDeliveryTime returns a boolean if a field has been set.

### SetDeliveryTimeNil

`func (o *ProofOfDeliveryUpdateDto) SetDeliveryTimeNil(b bool)`

 SetDeliveryTimeNil sets the value for DeliveryTime to be an explicit nil

### UnsetDeliveryTime
`func (o *ProofOfDeliveryUpdateDto) UnsetDeliveryTime()`

UnsetDeliveryTime ensures that no value is present for DeliveryTime, not even an explicit nil
### GetOverallCondition

`func (o *ProofOfDeliveryUpdateDto) GetOverallCondition() string`

GetOverallCondition returns the OverallCondition field if non-nil, zero value otherwise.

### GetOverallConditionOk

`func (o *ProofOfDeliveryUpdateDto) GetOverallConditionOk() (*string, bool)`

GetOverallConditionOk returns a tuple with the OverallCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallCondition

`func (o *ProofOfDeliveryUpdateDto) SetOverallCondition(v string)`

SetOverallCondition sets OverallCondition field to given value.

### HasOverallCondition

`func (o *ProofOfDeliveryUpdateDto) HasOverallCondition() bool`

HasOverallCondition returns a boolean if a field has been set.

### SetOverallConditionNil

`func (o *ProofOfDeliveryUpdateDto) SetOverallConditionNil(b bool)`

 SetOverallConditionNil sets the value for OverallCondition to be an explicit nil

### UnsetOverallCondition
`func (o *ProofOfDeliveryUpdateDto) UnsetOverallCondition()`

UnsetOverallCondition ensures that no value is present for OverallCondition, not even an explicit nil
### GetTotalQuantityDelivered

`func (o *ProofOfDeliveryUpdateDto) GetTotalQuantityDelivered() int32`

GetTotalQuantityDelivered returns the TotalQuantityDelivered field if non-nil, zero value otherwise.

### GetTotalQuantityDeliveredOk

`func (o *ProofOfDeliveryUpdateDto) GetTotalQuantityDeliveredOk() (*int32, bool)`

GetTotalQuantityDeliveredOk returns a tuple with the TotalQuantityDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuantityDelivered

`func (o *ProofOfDeliveryUpdateDto) SetTotalQuantityDelivered(v int32)`

SetTotalQuantityDelivered sets TotalQuantityDelivered field to given value.

### HasTotalQuantityDelivered

`func (o *ProofOfDeliveryUpdateDto) HasTotalQuantityDelivered() bool`

HasTotalQuantityDelivered returns a boolean if a field has been set.

### SetTotalQuantityDeliveredNil

`func (o *ProofOfDeliveryUpdateDto) SetTotalQuantityDeliveredNil(b bool)`

 SetTotalQuantityDeliveredNil sets the value for TotalQuantityDelivered to be an explicit nil

### UnsetTotalQuantityDelivered
`func (o *ProofOfDeliveryUpdateDto) UnsetTotalQuantityDelivered()`

UnsetTotalQuantityDelivered ensures that no value is present for TotalQuantityDelivered, not even an explicit nil
### GetTotalQuantityRejected

`func (o *ProofOfDeliveryUpdateDto) GetTotalQuantityRejected() int32`

GetTotalQuantityRejected returns the TotalQuantityRejected field if non-nil, zero value otherwise.

### GetTotalQuantityRejectedOk

`func (o *ProofOfDeliveryUpdateDto) GetTotalQuantityRejectedOk() (*int32, bool)`

GetTotalQuantityRejectedOk returns a tuple with the TotalQuantityRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuantityRejected

`func (o *ProofOfDeliveryUpdateDto) SetTotalQuantityRejected(v int32)`

SetTotalQuantityRejected sets TotalQuantityRejected field to given value.

### HasTotalQuantityRejected

`func (o *ProofOfDeliveryUpdateDto) HasTotalQuantityRejected() bool`

HasTotalQuantityRejected returns a boolean if a field has been set.

### SetTotalQuantityRejectedNil

`func (o *ProofOfDeliveryUpdateDto) SetTotalQuantityRejectedNil(b bool)`

 SetTotalQuantityRejectedNil sets the value for TotalQuantityRejected to be an explicit nil

### UnsetTotalQuantityRejected
`func (o *ProofOfDeliveryUpdateDto) UnsetTotalQuantityRejected()`

UnsetTotalQuantityRejected ensures that no value is present for TotalQuantityRejected, not even an explicit nil
### GetRemarks

`func (o *ProofOfDeliveryUpdateDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ProofOfDeliveryUpdateDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ProofOfDeliveryUpdateDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ProofOfDeliveryUpdateDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *ProofOfDeliveryUpdateDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *ProofOfDeliveryUpdateDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetPhotoEvidenceUri

`func (o *ProofOfDeliveryUpdateDto) GetPhotoEvidenceUri() string`

GetPhotoEvidenceUri returns the PhotoEvidenceUri field if non-nil, zero value otherwise.

### GetPhotoEvidenceUriOk

`func (o *ProofOfDeliveryUpdateDto) GetPhotoEvidenceUriOk() (*string, bool)`

GetPhotoEvidenceUriOk returns a tuple with the PhotoEvidenceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoEvidenceUri

`func (o *ProofOfDeliveryUpdateDto) SetPhotoEvidenceUri(v string)`

SetPhotoEvidenceUri sets PhotoEvidenceUri field to given value.

### HasPhotoEvidenceUri

`func (o *ProofOfDeliveryUpdateDto) HasPhotoEvidenceUri() bool`

HasPhotoEvidenceUri returns a boolean if a field has been set.

### SetPhotoEvidenceUriNil

`func (o *ProofOfDeliveryUpdateDto) SetPhotoEvidenceUriNil(b bool)`

 SetPhotoEvidenceUriNil sets the value for PhotoEvidenceUri to be an explicit nil

### UnsetPhotoEvidenceUri
`func (o *ProofOfDeliveryUpdateDto) UnsetPhotoEvidenceUri()`

UnsetPhotoEvidenceUri ensures that no value is present for PhotoEvidenceUri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


