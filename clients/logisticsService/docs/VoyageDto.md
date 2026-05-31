# VoyageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**VoyageNumber** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**VoyageStatus** | Pointer to **NullableString** |  | [optional] 
**VoyageDirection** | Pointer to **NullableString** |  | [optional] 
**DepartureDate** | Pointer to **NullableTime** |  | [optional] 
**ArrivalDate** | Pointer to **NullableTime** |  | [optional] 
**ActualDepartureDate** | Pointer to **NullableTime** |  | [optional] 
**ActualArrivalDate** | Pointer to **NullableTime** |  | [optional] 
**VesselId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVoyageDto

`func NewVoyageDto() *VoyageDto`

NewVoyageDto instantiates a new VoyageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoyageDtoWithDefaults

`func NewVoyageDtoWithDefaults() *VoyageDto`

NewVoyageDtoWithDefaults instantiates a new VoyageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VoyageDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoyageDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoyageDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoyageDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VoyageDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VoyageDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *VoyageDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VoyageDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VoyageDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *VoyageDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *VoyageDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *VoyageDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetVoyageNumber

`func (o *VoyageDto) GetVoyageNumber() string`

GetVoyageNumber returns the VoyageNumber field if non-nil, zero value otherwise.

### GetVoyageNumberOk

`func (o *VoyageDto) GetVoyageNumberOk() (*string, bool)`

GetVoyageNumberOk returns a tuple with the VoyageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageNumber

`func (o *VoyageDto) SetVoyageNumber(v string)`

SetVoyageNumber sets VoyageNumber field to given value.

### HasVoyageNumber

`func (o *VoyageDto) HasVoyageNumber() bool`

HasVoyageNumber returns a boolean if a field has been set.

### SetVoyageNumberNil

`func (o *VoyageDto) SetVoyageNumberNil(b bool)`

 SetVoyageNumberNil sets the value for VoyageNumber to be an explicit nil

### UnsetVoyageNumber
`func (o *VoyageDto) UnsetVoyageNumber()`

UnsetVoyageNumber ensures that no value is present for VoyageNumber, not even an explicit nil
### GetTitle

`func (o *VoyageDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VoyageDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VoyageDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VoyageDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *VoyageDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *VoyageDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *VoyageDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VoyageDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VoyageDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VoyageDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VoyageDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VoyageDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVoyageStatus

`func (o *VoyageDto) GetVoyageStatus() string`

GetVoyageStatus returns the VoyageStatus field if non-nil, zero value otherwise.

### GetVoyageStatusOk

`func (o *VoyageDto) GetVoyageStatusOk() (*string, bool)`

GetVoyageStatusOk returns a tuple with the VoyageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageStatus

`func (o *VoyageDto) SetVoyageStatus(v string)`

SetVoyageStatus sets VoyageStatus field to given value.

### HasVoyageStatus

`func (o *VoyageDto) HasVoyageStatus() bool`

HasVoyageStatus returns a boolean if a field has been set.

### SetVoyageStatusNil

`func (o *VoyageDto) SetVoyageStatusNil(b bool)`

 SetVoyageStatusNil sets the value for VoyageStatus to be an explicit nil

### UnsetVoyageStatus
`func (o *VoyageDto) UnsetVoyageStatus()`

UnsetVoyageStatus ensures that no value is present for VoyageStatus, not even an explicit nil
### GetVoyageDirection

`func (o *VoyageDto) GetVoyageDirection() string`

GetVoyageDirection returns the VoyageDirection field if non-nil, zero value otherwise.

### GetVoyageDirectionOk

`func (o *VoyageDto) GetVoyageDirectionOk() (*string, bool)`

GetVoyageDirectionOk returns a tuple with the VoyageDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageDirection

`func (o *VoyageDto) SetVoyageDirection(v string)`

SetVoyageDirection sets VoyageDirection field to given value.

### HasVoyageDirection

`func (o *VoyageDto) HasVoyageDirection() bool`

HasVoyageDirection returns a boolean if a field has been set.

### SetVoyageDirectionNil

`func (o *VoyageDto) SetVoyageDirectionNil(b bool)`

 SetVoyageDirectionNil sets the value for VoyageDirection to be an explicit nil

### UnsetVoyageDirection
`func (o *VoyageDto) UnsetVoyageDirection()`

UnsetVoyageDirection ensures that no value is present for VoyageDirection, not even an explicit nil
### GetDepartureDate

`func (o *VoyageDto) GetDepartureDate() time.Time`

GetDepartureDate returns the DepartureDate field if non-nil, zero value otherwise.

### GetDepartureDateOk

`func (o *VoyageDto) GetDepartureDateOk() (*time.Time, bool)`

GetDepartureDateOk returns a tuple with the DepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureDate

`func (o *VoyageDto) SetDepartureDate(v time.Time)`

SetDepartureDate sets DepartureDate field to given value.

### HasDepartureDate

`func (o *VoyageDto) HasDepartureDate() bool`

HasDepartureDate returns a boolean if a field has been set.

### SetDepartureDateNil

`func (o *VoyageDto) SetDepartureDateNil(b bool)`

 SetDepartureDateNil sets the value for DepartureDate to be an explicit nil

### UnsetDepartureDate
`func (o *VoyageDto) UnsetDepartureDate()`

UnsetDepartureDate ensures that no value is present for DepartureDate, not even an explicit nil
### GetArrivalDate

`func (o *VoyageDto) GetArrivalDate() time.Time`

GetArrivalDate returns the ArrivalDate field if non-nil, zero value otherwise.

### GetArrivalDateOk

`func (o *VoyageDto) GetArrivalDateOk() (*time.Time, bool)`

GetArrivalDateOk returns a tuple with the ArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalDate

`func (o *VoyageDto) SetArrivalDate(v time.Time)`

SetArrivalDate sets ArrivalDate field to given value.

### HasArrivalDate

`func (o *VoyageDto) HasArrivalDate() bool`

HasArrivalDate returns a boolean if a field has been set.

### SetArrivalDateNil

`func (o *VoyageDto) SetArrivalDateNil(b bool)`

 SetArrivalDateNil sets the value for ArrivalDate to be an explicit nil

### UnsetArrivalDate
`func (o *VoyageDto) UnsetArrivalDate()`

UnsetArrivalDate ensures that no value is present for ArrivalDate, not even an explicit nil
### GetActualDepartureDate

`func (o *VoyageDto) GetActualDepartureDate() time.Time`

GetActualDepartureDate returns the ActualDepartureDate field if non-nil, zero value otherwise.

### GetActualDepartureDateOk

`func (o *VoyageDto) GetActualDepartureDateOk() (*time.Time, bool)`

GetActualDepartureDateOk returns a tuple with the ActualDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDepartureDate

`func (o *VoyageDto) SetActualDepartureDate(v time.Time)`

SetActualDepartureDate sets ActualDepartureDate field to given value.

### HasActualDepartureDate

`func (o *VoyageDto) HasActualDepartureDate() bool`

HasActualDepartureDate returns a boolean if a field has been set.

### SetActualDepartureDateNil

`func (o *VoyageDto) SetActualDepartureDateNil(b bool)`

 SetActualDepartureDateNil sets the value for ActualDepartureDate to be an explicit nil

### UnsetActualDepartureDate
`func (o *VoyageDto) UnsetActualDepartureDate()`

UnsetActualDepartureDate ensures that no value is present for ActualDepartureDate, not even an explicit nil
### GetActualArrivalDate

`func (o *VoyageDto) GetActualArrivalDate() time.Time`

GetActualArrivalDate returns the ActualArrivalDate field if non-nil, zero value otherwise.

### GetActualArrivalDateOk

`func (o *VoyageDto) GetActualArrivalDateOk() (*time.Time, bool)`

GetActualArrivalDateOk returns a tuple with the ActualArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualArrivalDate

`func (o *VoyageDto) SetActualArrivalDate(v time.Time)`

SetActualArrivalDate sets ActualArrivalDate field to given value.

### HasActualArrivalDate

`func (o *VoyageDto) HasActualArrivalDate() bool`

HasActualArrivalDate returns a boolean if a field has been set.

### SetActualArrivalDateNil

`func (o *VoyageDto) SetActualArrivalDateNil(b bool)`

 SetActualArrivalDateNil sets the value for ActualArrivalDate to be an explicit nil

### UnsetActualArrivalDate
`func (o *VoyageDto) UnsetActualArrivalDate()`

UnsetActualArrivalDate ensures that no value is present for ActualArrivalDate, not even an explicit nil
### GetVesselId

`func (o *VoyageDto) GetVesselId() string`

GetVesselId returns the VesselId field if non-nil, zero value otherwise.

### GetVesselIdOk

`func (o *VoyageDto) GetVesselIdOk() (*string, bool)`

GetVesselIdOk returns a tuple with the VesselId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselId

`func (o *VoyageDto) SetVesselId(v string)`

SetVesselId sets VesselId field to given value.

### HasVesselId

`func (o *VoyageDto) HasVesselId() bool`

HasVesselId returns a boolean if a field has been set.

### SetVesselIdNil

`func (o *VoyageDto) SetVesselIdNil(b bool)`

 SetVesselIdNil sets the value for VesselId to be an explicit nil

### UnsetVesselId
`func (o *VoyageDto) UnsetVesselId()`

UnsetVesselId ensures that no value is present for VesselId, not even an explicit nil
### GetTenantId

`func (o *VoyageDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *VoyageDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *VoyageDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *VoyageDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *VoyageDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *VoyageDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *VoyageDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *VoyageDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *VoyageDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *VoyageDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *VoyageDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *VoyageDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


