# ProofOfDeliveryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**DocumentNumber** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
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
**SignedBy** | Pointer to **NullableString** |  | [optional] 
**SignerIdentification** | Pointer to **NullableString** |  | [optional] 
**SignatureDate** | Pointer to **NullableTime** |  | [optional] 
**DigitalSignatureReference** | Pointer to **NullableString** |  | [optional] 
**OverallCondition** | Pointer to **NullableString** |  | [optional] 
**TotalQuantityDelivered** | Pointer to **NullableInt32** |  | [optional] 
**TotalQuantityRejected** | Pointer to **NullableInt32** |  | [optional] 
**Remarks** | Pointer to **NullableString** |  | [optional] 
**PhotoEvidenceUri** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Lines** | Pointer to [**[]ProofOfDeliveryLineDto**](ProofOfDeliveryLineDto.md) |  | [optional] 
**DeliveryNoteIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProofOfDeliveryDto

`func NewProofOfDeliveryDto() *ProofOfDeliveryDto`

NewProofOfDeliveryDto instantiates a new ProofOfDeliveryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProofOfDeliveryDtoWithDefaults

`func NewProofOfDeliveryDtoWithDefaults() *ProofOfDeliveryDto`

NewProofOfDeliveryDtoWithDefaults instantiates a new ProofOfDeliveryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProofOfDeliveryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProofOfDeliveryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProofOfDeliveryDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProofOfDeliveryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProofOfDeliveryDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProofOfDeliveryDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ProofOfDeliveryDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProofOfDeliveryDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProofOfDeliveryDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProofOfDeliveryDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ProofOfDeliveryDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ProofOfDeliveryDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetDocumentNumber

`func (o *ProofOfDeliveryDto) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *ProofOfDeliveryDto) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *ProofOfDeliveryDto) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *ProofOfDeliveryDto) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *ProofOfDeliveryDto) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *ProofOfDeliveryDto) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetStatus

`func (o *ProofOfDeliveryDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProofOfDeliveryDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProofOfDeliveryDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProofOfDeliveryDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ProofOfDeliveryDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ProofOfDeliveryDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetShipmentId

`func (o *ProofOfDeliveryDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *ProofOfDeliveryDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *ProofOfDeliveryDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *ProofOfDeliveryDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *ProofOfDeliveryDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *ProofOfDeliveryDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetBillOfLadingId

`func (o *ProofOfDeliveryDto) GetBillOfLadingId() string`

GetBillOfLadingId returns the BillOfLadingId field if non-nil, zero value otherwise.

### GetBillOfLadingIdOk

`func (o *ProofOfDeliveryDto) GetBillOfLadingIdOk() (*string, bool)`

GetBillOfLadingIdOk returns a tuple with the BillOfLadingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingId

`func (o *ProofOfDeliveryDto) SetBillOfLadingId(v string)`

SetBillOfLadingId sets BillOfLadingId field to given value.

### HasBillOfLadingId

`func (o *ProofOfDeliveryDto) HasBillOfLadingId() bool`

HasBillOfLadingId returns a boolean if a field has been set.

### SetBillOfLadingIdNil

`func (o *ProofOfDeliveryDto) SetBillOfLadingIdNil(b bool)`

 SetBillOfLadingIdNil sets the value for BillOfLadingId to be an explicit nil

### UnsetBillOfLadingId
`func (o *ProofOfDeliveryDto) UnsetBillOfLadingId()`

UnsetBillOfLadingId ensures that no value is present for BillOfLadingId, not even an explicit nil
### GetSeawayBillId

`func (o *ProofOfDeliveryDto) GetSeawayBillId() string`

GetSeawayBillId returns the SeawayBillId field if non-nil, zero value otherwise.

### GetSeawayBillIdOk

`func (o *ProofOfDeliveryDto) GetSeawayBillIdOk() (*string, bool)`

GetSeawayBillIdOk returns a tuple with the SeawayBillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeawayBillId

`func (o *ProofOfDeliveryDto) SetSeawayBillId(v string)`

SetSeawayBillId sets SeawayBillId field to given value.

### HasSeawayBillId

`func (o *ProofOfDeliveryDto) HasSeawayBillId() bool`

HasSeawayBillId returns a boolean if a field has been set.

### SetSeawayBillIdNil

`func (o *ProofOfDeliveryDto) SetSeawayBillIdNil(b bool)`

 SetSeawayBillIdNil sets the value for SeawayBillId to be an explicit nil

### UnsetSeawayBillId
`func (o *ProofOfDeliveryDto) UnsetSeawayBillId()`

UnsetSeawayBillId ensures that no value is present for SeawayBillId, not even an explicit nil
### GetAirwayBillId

`func (o *ProofOfDeliveryDto) GetAirwayBillId() string`

GetAirwayBillId returns the AirwayBillId field if non-nil, zero value otherwise.

### GetAirwayBillIdOk

`func (o *ProofOfDeliveryDto) GetAirwayBillIdOk() (*string, bool)`

GetAirwayBillIdOk returns a tuple with the AirwayBillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirwayBillId

`func (o *ProofOfDeliveryDto) SetAirwayBillId(v string)`

SetAirwayBillId sets AirwayBillId field to given value.

### HasAirwayBillId

`func (o *ProofOfDeliveryDto) HasAirwayBillId() bool`

HasAirwayBillId returns a boolean if a field has been set.

### SetAirwayBillIdNil

`func (o *ProofOfDeliveryDto) SetAirwayBillIdNil(b bool)`

 SetAirwayBillIdNil sets the value for AirwayBillId to be an explicit nil

### UnsetAirwayBillId
`func (o *ProofOfDeliveryDto) UnsetAirwayBillId()`

UnsetAirwayBillId ensures that no value is present for AirwayBillId, not even an explicit nil
### GetRoadWaybillId

`func (o *ProofOfDeliveryDto) GetRoadWaybillId() string`

GetRoadWaybillId returns the RoadWaybillId field if non-nil, zero value otherwise.

### GetRoadWaybillIdOk

`func (o *ProofOfDeliveryDto) GetRoadWaybillIdOk() (*string, bool)`

GetRoadWaybillIdOk returns a tuple with the RoadWaybillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoadWaybillId

`func (o *ProofOfDeliveryDto) SetRoadWaybillId(v string)`

SetRoadWaybillId sets RoadWaybillId field to given value.

### HasRoadWaybillId

`func (o *ProofOfDeliveryDto) HasRoadWaybillId() bool`

HasRoadWaybillId returns a boolean if a field has been set.

### SetRoadWaybillIdNil

`func (o *ProofOfDeliveryDto) SetRoadWaybillIdNil(b bool)`

 SetRoadWaybillIdNil sets the value for RoadWaybillId to be an explicit nil

### UnsetRoadWaybillId
`func (o *ProofOfDeliveryDto) UnsetRoadWaybillId()`

UnsetRoadWaybillId ensures that no value is present for RoadWaybillId, not even an explicit nil
### GetRailWaybillId

`func (o *ProofOfDeliveryDto) GetRailWaybillId() string`

GetRailWaybillId returns the RailWaybillId field if non-nil, zero value otherwise.

### GetRailWaybillIdOk

`func (o *ProofOfDeliveryDto) GetRailWaybillIdOk() (*string, bool)`

GetRailWaybillIdOk returns a tuple with the RailWaybillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailWaybillId

`func (o *ProofOfDeliveryDto) SetRailWaybillId(v string)`

SetRailWaybillId sets RailWaybillId field to given value.

### HasRailWaybillId

`func (o *ProofOfDeliveryDto) HasRailWaybillId() bool`

HasRailWaybillId returns a boolean if a field has been set.

### SetRailWaybillIdNil

`func (o *ProofOfDeliveryDto) SetRailWaybillIdNil(b bool)`

 SetRailWaybillIdNil sets the value for RailWaybillId to be an explicit nil

### UnsetRailWaybillId
`func (o *ProofOfDeliveryDto) UnsetRailWaybillId()`

UnsetRailWaybillId ensures that no value is present for RailWaybillId, not even an explicit nil
### GetTruckTripId

`func (o *ProofOfDeliveryDto) GetTruckTripId() string`

GetTruckTripId returns the TruckTripId field if non-nil, zero value otherwise.

### GetTruckTripIdOk

`func (o *ProofOfDeliveryDto) GetTruckTripIdOk() (*string, bool)`

GetTruckTripIdOk returns a tuple with the TruckTripId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckTripId

`func (o *ProofOfDeliveryDto) SetTruckTripId(v string)`

SetTruckTripId sets TruckTripId field to given value.

### HasTruckTripId

`func (o *ProofOfDeliveryDto) HasTruckTripId() bool`

HasTruckTripId returns a boolean if a field has been set.

### SetTruckTripIdNil

`func (o *ProofOfDeliveryDto) SetTruckTripIdNil(b bool)`

 SetTruckTripIdNil sets the value for TruckTripId to be an explicit nil

### UnsetTruckTripId
`func (o *ProofOfDeliveryDto) UnsetTruckTripId()`

UnsetTruckTripId ensures that no value is present for TruckTripId, not even an explicit nil
### GetRecipientName

`func (o *ProofOfDeliveryDto) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *ProofOfDeliveryDto) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *ProofOfDeliveryDto) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.

### HasRecipientName

`func (o *ProofOfDeliveryDto) HasRecipientName() bool`

HasRecipientName returns a boolean if a field has been set.

### SetRecipientNameNil

`func (o *ProofOfDeliveryDto) SetRecipientNameNil(b bool)`

 SetRecipientNameNil sets the value for RecipientName to be an explicit nil

### UnsetRecipientName
`func (o *ProofOfDeliveryDto) UnsetRecipientName()`

UnsetRecipientName ensures that no value is present for RecipientName, not even an explicit nil
### GetRecipientCompanyContactId

`func (o *ProofOfDeliveryDto) GetRecipientCompanyContactId() string`

GetRecipientCompanyContactId returns the RecipientCompanyContactId field if non-nil, zero value otherwise.

### GetRecipientCompanyContactIdOk

`func (o *ProofOfDeliveryDto) GetRecipientCompanyContactIdOk() (*string, bool)`

GetRecipientCompanyContactIdOk returns a tuple with the RecipientCompanyContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientCompanyContactId

`func (o *ProofOfDeliveryDto) SetRecipientCompanyContactId(v string)`

SetRecipientCompanyContactId sets RecipientCompanyContactId field to given value.

### HasRecipientCompanyContactId

`func (o *ProofOfDeliveryDto) HasRecipientCompanyContactId() bool`

HasRecipientCompanyContactId returns a boolean if a field has been set.

### SetRecipientCompanyContactIdNil

`func (o *ProofOfDeliveryDto) SetRecipientCompanyContactIdNil(b bool)`

 SetRecipientCompanyContactIdNil sets the value for RecipientCompanyContactId to be an explicit nil

### UnsetRecipientCompanyContactId
`func (o *ProofOfDeliveryDto) UnsetRecipientCompanyContactId()`

UnsetRecipientCompanyContactId ensures that no value is present for RecipientCompanyContactId, not even an explicit nil
### GetDeliveryAddress

`func (o *ProofOfDeliveryDto) GetDeliveryAddress() string`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *ProofOfDeliveryDto) GetDeliveryAddressOk() (*string, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *ProofOfDeliveryDto) SetDeliveryAddress(v string)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *ProofOfDeliveryDto) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### SetDeliveryAddressNil

`func (o *ProofOfDeliveryDto) SetDeliveryAddressNil(b bool)`

 SetDeliveryAddressNil sets the value for DeliveryAddress to be an explicit nil

### UnsetDeliveryAddress
`func (o *ProofOfDeliveryDto) UnsetDeliveryAddress()`

UnsetDeliveryAddress ensures that no value is present for DeliveryAddress, not even an explicit nil
### GetDeliveryDate

`func (o *ProofOfDeliveryDto) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *ProofOfDeliveryDto) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *ProofOfDeliveryDto) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *ProofOfDeliveryDto) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### SetDeliveryDateNil

`func (o *ProofOfDeliveryDto) SetDeliveryDateNil(b bool)`

 SetDeliveryDateNil sets the value for DeliveryDate to be an explicit nil

### UnsetDeliveryDate
`func (o *ProofOfDeliveryDto) UnsetDeliveryDate()`

UnsetDeliveryDate ensures that no value is present for DeliveryDate, not even an explicit nil
### GetDeliveryTime

`func (o *ProofOfDeliveryDto) GetDeliveryTime() string`

GetDeliveryTime returns the DeliveryTime field if non-nil, zero value otherwise.

### GetDeliveryTimeOk

`func (o *ProofOfDeliveryDto) GetDeliveryTimeOk() (*string, bool)`

GetDeliveryTimeOk returns a tuple with the DeliveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTime

`func (o *ProofOfDeliveryDto) SetDeliveryTime(v string)`

SetDeliveryTime sets DeliveryTime field to given value.

### HasDeliveryTime

`func (o *ProofOfDeliveryDto) HasDeliveryTime() bool`

HasDeliveryTime returns a boolean if a field has been set.

### SetDeliveryTimeNil

`func (o *ProofOfDeliveryDto) SetDeliveryTimeNil(b bool)`

 SetDeliveryTimeNil sets the value for DeliveryTime to be an explicit nil

### UnsetDeliveryTime
`func (o *ProofOfDeliveryDto) UnsetDeliveryTime()`

UnsetDeliveryTime ensures that no value is present for DeliveryTime, not even an explicit nil
### GetSignedBy

`func (o *ProofOfDeliveryDto) GetSignedBy() string`

GetSignedBy returns the SignedBy field if non-nil, zero value otherwise.

### GetSignedByOk

`func (o *ProofOfDeliveryDto) GetSignedByOk() (*string, bool)`

GetSignedByOk returns a tuple with the SignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedBy

`func (o *ProofOfDeliveryDto) SetSignedBy(v string)`

SetSignedBy sets SignedBy field to given value.

### HasSignedBy

`func (o *ProofOfDeliveryDto) HasSignedBy() bool`

HasSignedBy returns a boolean if a field has been set.

### SetSignedByNil

`func (o *ProofOfDeliveryDto) SetSignedByNil(b bool)`

 SetSignedByNil sets the value for SignedBy to be an explicit nil

### UnsetSignedBy
`func (o *ProofOfDeliveryDto) UnsetSignedBy()`

UnsetSignedBy ensures that no value is present for SignedBy, not even an explicit nil
### GetSignerIdentification

`func (o *ProofOfDeliveryDto) GetSignerIdentification() string`

GetSignerIdentification returns the SignerIdentification field if non-nil, zero value otherwise.

### GetSignerIdentificationOk

`func (o *ProofOfDeliveryDto) GetSignerIdentificationOk() (*string, bool)`

GetSignerIdentificationOk returns a tuple with the SignerIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerIdentification

`func (o *ProofOfDeliveryDto) SetSignerIdentification(v string)`

SetSignerIdentification sets SignerIdentification field to given value.

### HasSignerIdentification

`func (o *ProofOfDeliveryDto) HasSignerIdentification() bool`

HasSignerIdentification returns a boolean if a field has been set.

### SetSignerIdentificationNil

`func (o *ProofOfDeliveryDto) SetSignerIdentificationNil(b bool)`

 SetSignerIdentificationNil sets the value for SignerIdentification to be an explicit nil

### UnsetSignerIdentification
`func (o *ProofOfDeliveryDto) UnsetSignerIdentification()`

UnsetSignerIdentification ensures that no value is present for SignerIdentification, not even an explicit nil
### GetSignatureDate

`func (o *ProofOfDeliveryDto) GetSignatureDate() time.Time`

GetSignatureDate returns the SignatureDate field if non-nil, zero value otherwise.

### GetSignatureDateOk

`func (o *ProofOfDeliveryDto) GetSignatureDateOk() (*time.Time, bool)`

GetSignatureDateOk returns a tuple with the SignatureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureDate

`func (o *ProofOfDeliveryDto) SetSignatureDate(v time.Time)`

SetSignatureDate sets SignatureDate field to given value.

### HasSignatureDate

`func (o *ProofOfDeliveryDto) HasSignatureDate() bool`

HasSignatureDate returns a boolean if a field has been set.

### SetSignatureDateNil

`func (o *ProofOfDeliveryDto) SetSignatureDateNil(b bool)`

 SetSignatureDateNil sets the value for SignatureDate to be an explicit nil

### UnsetSignatureDate
`func (o *ProofOfDeliveryDto) UnsetSignatureDate()`

UnsetSignatureDate ensures that no value is present for SignatureDate, not even an explicit nil
### GetDigitalSignatureReference

`func (o *ProofOfDeliveryDto) GetDigitalSignatureReference() string`

GetDigitalSignatureReference returns the DigitalSignatureReference field if non-nil, zero value otherwise.

### GetDigitalSignatureReferenceOk

`func (o *ProofOfDeliveryDto) GetDigitalSignatureReferenceOk() (*string, bool)`

GetDigitalSignatureReferenceOk returns a tuple with the DigitalSignatureReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalSignatureReference

`func (o *ProofOfDeliveryDto) SetDigitalSignatureReference(v string)`

SetDigitalSignatureReference sets DigitalSignatureReference field to given value.

### HasDigitalSignatureReference

`func (o *ProofOfDeliveryDto) HasDigitalSignatureReference() bool`

HasDigitalSignatureReference returns a boolean if a field has been set.

### SetDigitalSignatureReferenceNil

`func (o *ProofOfDeliveryDto) SetDigitalSignatureReferenceNil(b bool)`

 SetDigitalSignatureReferenceNil sets the value for DigitalSignatureReference to be an explicit nil

### UnsetDigitalSignatureReference
`func (o *ProofOfDeliveryDto) UnsetDigitalSignatureReference()`

UnsetDigitalSignatureReference ensures that no value is present for DigitalSignatureReference, not even an explicit nil
### GetOverallCondition

`func (o *ProofOfDeliveryDto) GetOverallCondition() string`

GetOverallCondition returns the OverallCondition field if non-nil, zero value otherwise.

### GetOverallConditionOk

`func (o *ProofOfDeliveryDto) GetOverallConditionOk() (*string, bool)`

GetOverallConditionOk returns a tuple with the OverallCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallCondition

`func (o *ProofOfDeliveryDto) SetOverallCondition(v string)`

SetOverallCondition sets OverallCondition field to given value.

### HasOverallCondition

`func (o *ProofOfDeliveryDto) HasOverallCondition() bool`

HasOverallCondition returns a boolean if a field has been set.

### SetOverallConditionNil

`func (o *ProofOfDeliveryDto) SetOverallConditionNil(b bool)`

 SetOverallConditionNil sets the value for OverallCondition to be an explicit nil

### UnsetOverallCondition
`func (o *ProofOfDeliveryDto) UnsetOverallCondition()`

UnsetOverallCondition ensures that no value is present for OverallCondition, not even an explicit nil
### GetTotalQuantityDelivered

`func (o *ProofOfDeliveryDto) GetTotalQuantityDelivered() int32`

GetTotalQuantityDelivered returns the TotalQuantityDelivered field if non-nil, zero value otherwise.

### GetTotalQuantityDeliveredOk

`func (o *ProofOfDeliveryDto) GetTotalQuantityDeliveredOk() (*int32, bool)`

GetTotalQuantityDeliveredOk returns a tuple with the TotalQuantityDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuantityDelivered

`func (o *ProofOfDeliveryDto) SetTotalQuantityDelivered(v int32)`

SetTotalQuantityDelivered sets TotalQuantityDelivered field to given value.

### HasTotalQuantityDelivered

`func (o *ProofOfDeliveryDto) HasTotalQuantityDelivered() bool`

HasTotalQuantityDelivered returns a boolean if a field has been set.

### SetTotalQuantityDeliveredNil

`func (o *ProofOfDeliveryDto) SetTotalQuantityDeliveredNil(b bool)`

 SetTotalQuantityDeliveredNil sets the value for TotalQuantityDelivered to be an explicit nil

### UnsetTotalQuantityDelivered
`func (o *ProofOfDeliveryDto) UnsetTotalQuantityDelivered()`

UnsetTotalQuantityDelivered ensures that no value is present for TotalQuantityDelivered, not even an explicit nil
### GetTotalQuantityRejected

`func (o *ProofOfDeliveryDto) GetTotalQuantityRejected() int32`

GetTotalQuantityRejected returns the TotalQuantityRejected field if non-nil, zero value otherwise.

### GetTotalQuantityRejectedOk

`func (o *ProofOfDeliveryDto) GetTotalQuantityRejectedOk() (*int32, bool)`

GetTotalQuantityRejectedOk returns a tuple with the TotalQuantityRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuantityRejected

`func (o *ProofOfDeliveryDto) SetTotalQuantityRejected(v int32)`

SetTotalQuantityRejected sets TotalQuantityRejected field to given value.

### HasTotalQuantityRejected

`func (o *ProofOfDeliveryDto) HasTotalQuantityRejected() bool`

HasTotalQuantityRejected returns a boolean if a field has been set.

### SetTotalQuantityRejectedNil

`func (o *ProofOfDeliveryDto) SetTotalQuantityRejectedNil(b bool)`

 SetTotalQuantityRejectedNil sets the value for TotalQuantityRejected to be an explicit nil

### UnsetTotalQuantityRejected
`func (o *ProofOfDeliveryDto) UnsetTotalQuantityRejected()`

UnsetTotalQuantityRejected ensures that no value is present for TotalQuantityRejected, not even an explicit nil
### GetRemarks

`func (o *ProofOfDeliveryDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ProofOfDeliveryDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ProofOfDeliveryDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ProofOfDeliveryDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *ProofOfDeliveryDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *ProofOfDeliveryDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetPhotoEvidenceUri

`func (o *ProofOfDeliveryDto) GetPhotoEvidenceUri() string`

GetPhotoEvidenceUri returns the PhotoEvidenceUri field if non-nil, zero value otherwise.

### GetPhotoEvidenceUriOk

`func (o *ProofOfDeliveryDto) GetPhotoEvidenceUriOk() (*string, bool)`

GetPhotoEvidenceUriOk returns a tuple with the PhotoEvidenceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoEvidenceUri

`func (o *ProofOfDeliveryDto) SetPhotoEvidenceUri(v string)`

SetPhotoEvidenceUri sets PhotoEvidenceUri field to given value.

### HasPhotoEvidenceUri

`func (o *ProofOfDeliveryDto) HasPhotoEvidenceUri() bool`

HasPhotoEvidenceUri returns a boolean if a field has been set.

### SetPhotoEvidenceUriNil

`func (o *ProofOfDeliveryDto) SetPhotoEvidenceUriNil(b bool)`

 SetPhotoEvidenceUriNil sets the value for PhotoEvidenceUri to be an explicit nil

### UnsetPhotoEvidenceUri
`func (o *ProofOfDeliveryDto) UnsetPhotoEvidenceUri()`

UnsetPhotoEvidenceUri ensures that no value is present for PhotoEvidenceUri, not even an explicit nil
### GetTenantId

`func (o *ProofOfDeliveryDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ProofOfDeliveryDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ProofOfDeliveryDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ProofOfDeliveryDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ProofOfDeliveryDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ProofOfDeliveryDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *ProofOfDeliveryDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ProofOfDeliveryDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ProofOfDeliveryDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ProofOfDeliveryDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ProofOfDeliveryDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ProofOfDeliveryDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetLines

`func (o *ProofOfDeliveryDto) GetLines() []ProofOfDeliveryLineDto`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *ProofOfDeliveryDto) GetLinesOk() (*[]ProofOfDeliveryLineDto, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *ProofOfDeliveryDto) SetLines(v []ProofOfDeliveryLineDto)`

SetLines sets Lines field to given value.

### HasLines

`func (o *ProofOfDeliveryDto) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *ProofOfDeliveryDto) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *ProofOfDeliveryDto) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil
### GetDeliveryNoteIds

`func (o *ProofOfDeliveryDto) GetDeliveryNoteIds() []string`

GetDeliveryNoteIds returns the DeliveryNoteIds field if non-nil, zero value otherwise.

### GetDeliveryNoteIdsOk

`func (o *ProofOfDeliveryDto) GetDeliveryNoteIdsOk() (*[]string, bool)`

GetDeliveryNoteIdsOk returns a tuple with the DeliveryNoteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryNoteIds

`func (o *ProofOfDeliveryDto) SetDeliveryNoteIds(v []string)`

SetDeliveryNoteIds sets DeliveryNoteIds field to given value.

### HasDeliveryNoteIds

`func (o *ProofOfDeliveryDto) HasDeliveryNoteIds() bool`

HasDeliveryNoteIds returns a boolean if a field has been set.

### SetDeliveryNoteIdsNil

`func (o *ProofOfDeliveryDto) SetDeliveryNoteIdsNil(b bool)`

 SetDeliveryNoteIdsNil sets the value for DeliveryNoteIds to be an explicit nil

### UnsetDeliveryNoteIds
`func (o *ProofOfDeliveryDto) UnsetDeliveryNoteIds()`

UnsetDeliveryNoteIds ensures that no value is present for DeliveryNoteIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


