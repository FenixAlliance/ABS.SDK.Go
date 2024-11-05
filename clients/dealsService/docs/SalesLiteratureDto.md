# SalesLiteratureDto

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

## Methods

### NewSalesLiteratureDto

`func NewSalesLiteratureDto() *SalesLiteratureDto`

NewSalesLiteratureDto instantiates a new SalesLiteratureDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesLiteratureDtoWithDefaults

`func NewSalesLiteratureDtoWithDefaults() *SalesLiteratureDto`

NewSalesLiteratureDtoWithDefaults instantiates a new SalesLiteratureDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SalesLiteratureDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SalesLiteratureDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SalesLiteratureDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SalesLiteratureDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SalesLiteratureDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SalesLiteratureDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SalesLiteratureDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SalesLiteratureDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SalesLiteratureDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SalesLiteratureDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SalesLiteratureDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SalesLiteratureDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *SalesLiteratureDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SalesLiteratureDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SalesLiteratureDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SalesLiteratureDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SalesLiteratureDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SalesLiteratureDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetContent

`func (o *SalesLiteratureDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SalesLiteratureDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SalesLiteratureDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SalesLiteratureDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *SalesLiteratureDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *SalesLiteratureDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetDescription

`func (o *SalesLiteratureDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SalesLiteratureDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SalesLiteratureDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SalesLiteratureDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SalesLiteratureDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SalesLiteratureDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetModifiedDate

`func (o *SalesLiteratureDto) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *SalesLiteratureDto) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *SalesLiteratureDto) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *SalesLiteratureDto) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *SalesLiteratureDto) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *SalesLiteratureDto) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *SalesLiteratureDto) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *SalesLiteratureDto) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetTenantId

`func (o *SalesLiteratureDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SalesLiteratureDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SalesLiteratureDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SalesLiteratureDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SalesLiteratureDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SalesLiteratureDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrolmentId

`func (o *SalesLiteratureDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *SalesLiteratureDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *SalesLiteratureDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *SalesLiteratureDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *SalesLiteratureDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *SalesLiteratureDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
### GetSalesLiteratureTypeId

`func (o *SalesLiteratureDto) GetSalesLiteratureTypeId() string`

GetSalesLiteratureTypeId returns the SalesLiteratureTypeId field if non-nil, zero value otherwise.

### GetSalesLiteratureTypeIdOk

`func (o *SalesLiteratureDto) GetSalesLiteratureTypeIdOk() (*string, bool)`

GetSalesLiteratureTypeIdOk returns a tuple with the SalesLiteratureTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesLiteratureTypeId

`func (o *SalesLiteratureDto) SetSalesLiteratureTypeId(v string)`

SetSalesLiteratureTypeId sets SalesLiteratureTypeId field to given value.

### HasSalesLiteratureTypeId

`func (o *SalesLiteratureDto) HasSalesLiteratureTypeId() bool`

HasSalesLiteratureTypeId returns a boolean if a field has been set.

### SetSalesLiteratureTypeIdNil

`func (o *SalesLiteratureDto) SetSalesLiteratureTypeIdNil(b bool)`

 SetSalesLiteratureTypeIdNil sets the value for SalesLiteratureTypeId to be an explicit nil

### UnsetSalesLiteratureTypeId
`func (o *SalesLiteratureDto) UnsetSalesLiteratureTypeId()`

UnsetSalesLiteratureTypeId ensures that no value is present for SalesLiteratureTypeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


