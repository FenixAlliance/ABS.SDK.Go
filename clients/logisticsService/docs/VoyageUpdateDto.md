# VoyageUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VoyageNumber** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**VoyageDirection** | Pointer to **NullableString** |  | [optional] 
**DepartureDate** | Pointer to **NullableTime** |  | [optional] 
**ArrivalDate** | Pointer to **NullableTime** |  | [optional] 
**VesselId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVoyageUpdateDto

`func NewVoyageUpdateDto() *VoyageUpdateDto`

NewVoyageUpdateDto instantiates a new VoyageUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoyageUpdateDtoWithDefaults

`func NewVoyageUpdateDtoWithDefaults() *VoyageUpdateDto`

NewVoyageUpdateDtoWithDefaults instantiates a new VoyageUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoyageNumber

`func (o *VoyageUpdateDto) GetVoyageNumber() string`

GetVoyageNumber returns the VoyageNumber field if non-nil, zero value otherwise.

### GetVoyageNumberOk

`func (o *VoyageUpdateDto) GetVoyageNumberOk() (*string, bool)`

GetVoyageNumberOk returns a tuple with the VoyageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageNumber

`func (o *VoyageUpdateDto) SetVoyageNumber(v string)`

SetVoyageNumber sets VoyageNumber field to given value.

### HasVoyageNumber

`func (o *VoyageUpdateDto) HasVoyageNumber() bool`

HasVoyageNumber returns a boolean if a field has been set.

### SetVoyageNumberNil

`func (o *VoyageUpdateDto) SetVoyageNumberNil(b bool)`

 SetVoyageNumberNil sets the value for VoyageNumber to be an explicit nil

### UnsetVoyageNumber
`func (o *VoyageUpdateDto) UnsetVoyageNumber()`

UnsetVoyageNumber ensures that no value is present for VoyageNumber, not even an explicit nil
### GetTitle

`func (o *VoyageUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VoyageUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VoyageUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VoyageUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *VoyageUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *VoyageUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *VoyageUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VoyageUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VoyageUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VoyageUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VoyageUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VoyageUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVoyageDirection

`func (o *VoyageUpdateDto) GetVoyageDirection() string`

GetVoyageDirection returns the VoyageDirection field if non-nil, zero value otherwise.

### GetVoyageDirectionOk

`func (o *VoyageUpdateDto) GetVoyageDirectionOk() (*string, bool)`

GetVoyageDirectionOk returns a tuple with the VoyageDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageDirection

`func (o *VoyageUpdateDto) SetVoyageDirection(v string)`

SetVoyageDirection sets VoyageDirection field to given value.

### HasVoyageDirection

`func (o *VoyageUpdateDto) HasVoyageDirection() bool`

HasVoyageDirection returns a boolean if a field has been set.

### SetVoyageDirectionNil

`func (o *VoyageUpdateDto) SetVoyageDirectionNil(b bool)`

 SetVoyageDirectionNil sets the value for VoyageDirection to be an explicit nil

### UnsetVoyageDirection
`func (o *VoyageUpdateDto) UnsetVoyageDirection()`

UnsetVoyageDirection ensures that no value is present for VoyageDirection, not even an explicit nil
### GetDepartureDate

`func (o *VoyageUpdateDto) GetDepartureDate() time.Time`

GetDepartureDate returns the DepartureDate field if non-nil, zero value otherwise.

### GetDepartureDateOk

`func (o *VoyageUpdateDto) GetDepartureDateOk() (*time.Time, bool)`

GetDepartureDateOk returns a tuple with the DepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureDate

`func (o *VoyageUpdateDto) SetDepartureDate(v time.Time)`

SetDepartureDate sets DepartureDate field to given value.

### HasDepartureDate

`func (o *VoyageUpdateDto) HasDepartureDate() bool`

HasDepartureDate returns a boolean if a field has been set.

### SetDepartureDateNil

`func (o *VoyageUpdateDto) SetDepartureDateNil(b bool)`

 SetDepartureDateNil sets the value for DepartureDate to be an explicit nil

### UnsetDepartureDate
`func (o *VoyageUpdateDto) UnsetDepartureDate()`

UnsetDepartureDate ensures that no value is present for DepartureDate, not even an explicit nil
### GetArrivalDate

`func (o *VoyageUpdateDto) GetArrivalDate() time.Time`

GetArrivalDate returns the ArrivalDate field if non-nil, zero value otherwise.

### GetArrivalDateOk

`func (o *VoyageUpdateDto) GetArrivalDateOk() (*time.Time, bool)`

GetArrivalDateOk returns a tuple with the ArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalDate

`func (o *VoyageUpdateDto) SetArrivalDate(v time.Time)`

SetArrivalDate sets ArrivalDate field to given value.

### HasArrivalDate

`func (o *VoyageUpdateDto) HasArrivalDate() bool`

HasArrivalDate returns a boolean if a field has been set.

### SetArrivalDateNil

`func (o *VoyageUpdateDto) SetArrivalDateNil(b bool)`

 SetArrivalDateNil sets the value for ArrivalDate to be an explicit nil

### UnsetArrivalDate
`func (o *VoyageUpdateDto) UnsetArrivalDate()`

UnsetArrivalDate ensures that no value is present for ArrivalDate, not even an explicit nil
### GetVesselId

`func (o *VoyageUpdateDto) GetVesselId() string`

GetVesselId returns the VesselId field if non-nil, zero value otherwise.

### GetVesselIdOk

`func (o *VoyageUpdateDto) GetVesselIdOk() (*string, bool)`

GetVesselIdOk returns a tuple with the VesselId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselId

`func (o *VoyageUpdateDto) SetVesselId(v string)`

SetVesselId sets VesselId field to given value.

### HasVesselId

`func (o *VoyageUpdateDto) HasVesselId() bool`

HasVesselId returns a boolean if a field has been set.

### SetVesselIdNil

`func (o *VoyageUpdateDto) SetVesselIdNil(b bool)`

 SetVesselIdNil sets the value for VesselId to be an explicit nil

### UnsetVesselId
`func (o *VoyageUpdateDto) UnsetVesselId()`

UnsetVesselId ensures that no value is present for VesselId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


