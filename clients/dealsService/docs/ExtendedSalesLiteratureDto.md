# ExtendedSalesLiteratureDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ModifiedDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrolmentId** | Pointer to **NullableString** |  | [optional] 
**SalesLiteratureTypeId** | Pointer to **NullableString** |  | [optional] 
**SalesLiteratureType** | Pointer to [**SalesLiteratureTypeDto**](SalesLiteratureTypeDto.md) |  | [optional] 
**Tenant** | Pointer to [**TenantDto**](TenantDto.md) |  | [optional] 

## Methods

### NewExtendedSalesLiteratureDto

`func NewExtendedSalesLiteratureDto() *ExtendedSalesLiteratureDto`

NewExtendedSalesLiteratureDto instantiates a new ExtendedSalesLiteratureDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedSalesLiteratureDtoWithDefaults

`func NewExtendedSalesLiteratureDtoWithDefaults() *ExtendedSalesLiteratureDto`

NewExtendedSalesLiteratureDtoWithDefaults instantiates a new ExtendedSalesLiteratureDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtendedSalesLiteratureDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtendedSalesLiteratureDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtendedSalesLiteratureDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExtendedSalesLiteratureDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ExtendedSalesLiteratureDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExtendedSalesLiteratureDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ExtendedSalesLiteratureDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExtendedSalesLiteratureDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExtendedSalesLiteratureDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ExtendedSalesLiteratureDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ExtendedSalesLiteratureDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ExtendedSalesLiteratureDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *ExtendedSalesLiteratureDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExtendedSalesLiteratureDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExtendedSalesLiteratureDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ExtendedSalesLiteratureDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ExtendedSalesLiteratureDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ExtendedSalesLiteratureDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetContent

`func (o *ExtendedSalesLiteratureDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ExtendedSalesLiteratureDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ExtendedSalesLiteratureDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ExtendedSalesLiteratureDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *ExtendedSalesLiteratureDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *ExtendedSalesLiteratureDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetDescription

`func (o *ExtendedSalesLiteratureDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtendedSalesLiteratureDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtendedSalesLiteratureDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtendedSalesLiteratureDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExtendedSalesLiteratureDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExtendedSalesLiteratureDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetModifiedDate

`func (o *ExtendedSalesLiteratureDto) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ExtendedSalesLiteratureDto) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ExtendedSalesLiteratureDto) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *ExtendedSalesLiteratureDto) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ExtendedSalesLiteratureDto) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ExtendedSalesLiteratureDto) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ExtendedSalesLiteratureDto) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ExtendedSalesLiteratureDto) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetTenantId

`func (o *ExtendedSalesLiteratureDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ExtendedSalesLiteratureDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ExtendedSalesLiteratureDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ExtendedSalesLiteratureDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ExtendedSalesLiteratureDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ExtendedSalesLiteratureDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrolmentId

`func (o *ExtendedSalesLiteratureDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *ExtendedSalesLiteratureDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *ExtendedSalesLiteratureDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *ExtendedSalesLiteratureDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *ExtendedSalesLiteratureDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *ExtendedSalesLiteratureDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
### GetSalesLiteratureTypeId

`func (o *ExtendedSalesLiteratureDto) GetSalesLiteratureTypeId() string`

GetSalesLiteratureTypeId returns the SalesLiteratureTypeId field if non-nil, zero value otherwise.

### GetSalesLiteratureTypeIdOk

`func (o *ExtendedSalesLiteratureDto) GetSalesLiteratureTypeIdOk() (*string, bool)`

GetSalesLiteratureTypeIdOk returns a tuple with the SalesLiteratureTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesLiteratureTypeId

`func (o *ExtendedSalesLiteratureDto) SetSalesLiteratureTypeId(v string)`

SetSalesLiteratureTypeId sets SalesLiteratureTypeId field to given value.

### HasSalesLiteratureTypeId

`func (o *ExtendedSalesLiteratureDto) HasSalesLiteratureTypeId() bool`

HasSalesLiteratureTypeId returns a boolean if a field has been set.

### SetSalesLiteratureTypeIdNil

`func (o *ExtendedSalesLiteratureDto) SetSalesLiteratureTypeIdNil(b bool)`

 SetSalesLiteratureTypeIdNil sets the value for SalesLiteratureTypeId to be an explicit nil

### UnsetSalesLiteratureTypeId
`func (o *ExtendedSalesLiteratureDto) UnsetSalesLiteratureTypeId()`

UnsetSalesLiteratureTypeId ensures that no value is present for SalesLiteratureTypeId, not even an explicit nil
### GetSalesLiteratureType

`func (o *ExtendedSalesLiteratureDto) GetSalesLiteratureType() SalesLiteratureTypeDto`

GetSalesLiteratureType returns the SalesLiteratureType field if non-nil, zero value otherwise.

### GetSalesLiteratureTypeOk

`func (o *ExtendedSalesLiteratureDto) GetSalesLiteratureTypeOk() (*SalesLiteratureTypeDto, bool)`

GetSalesLiteratureTypeOk returns a tuple with the SalesLiteratureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesLiteratureType

`func (o *ExtendedSalesLiteratureDto) SetSalesLiteratureType(v SalesLiteratureTypeDto)`

SetSalesLiteratureType sets SalesLiteratureType field to given value.

### HasSalesLiteratureType

`func (o *ExtendedSalesLiteratureDto) HasSalesLiteratureType() bool`

HasSalesLiteratureType returns a boolean if a field has been set.

### GetTenant

`func (o *ExtendedSalesLiteratureDto) GetTenant() TenantDto`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ExtendedSalesLiteratureDto) GetTenantOk() (*TenantDto, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ExtendedSalesLiteratureDto) SetTenant(v TenantDto)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ExtendedSalesLiteratureDto) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


