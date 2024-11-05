# SalesLiteratureCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ModifiedDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrolmentId** | Pointer to **NullableString** |  | [optional] 
**SalesLiteratureTypeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSalesLiteratureCreateDto

`func NewSalesLiteratureCreateDto() *SalesLiteratureCreateDto`

NewSalesLiteratureCreateDto instantiates a new SalesLiteratureCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesLiteratureCreateDtoWithDefaults

`func NewSalesLiteratureCreateDtoWithDefaults() *SalesLiteratureCreateDto`

NewSalesLiteratureCreateDtoWithDefaults instantiates a new SalesLiteratureCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SalesLiteratureCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SalesLiteratureCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SalesLiteratureCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SalesLiteratureCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SalesLiteratureCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SalesLiteratureCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SalesLiteratureCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SalesLiteratureCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *SalesLiteratureCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SalesLiteratureCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SalesLiteratureCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SalesLiteratureCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SalesLiteratureCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SalesLiteratureCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetContent

`func (o *SalesLiteratureCreateDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SalesLiteratureCreateDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SalesLiteratureCreateDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SalesLiteratureCreateDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *SalesLiteratureCreateDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *SalesLiteratureCreateDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetDescription

`func (o *SalesLiteratureCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SalesLiteratureCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SalesLiteratureCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SalesLiteratureCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SalesLiteratureCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SalesLiteratureCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetModifiedDate

`func (o *SalesLiteratureCreateDto) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *SalesLiteratureCreateDto) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *SalesLiteratureCreateDto) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *SalesLiteratureCreateDto) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *SalesLiteratureCreateDto) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *SalesLiteratureCreateDto) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *SalesLiteratureCreateDto) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *SalesLiteratureCreateDto) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetTenantId

`func (o *SalesLiteratureCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SalesLiteratureCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SalesLiteratureCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SalesLiteratureCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SalesLiteratureCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SalesLiteratureCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrolmentId

`func (o *SalesLiteratureCreateDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *SalesLiteratureCreateDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *SalesLiteratureCreateDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *SalesLiteratureCreateDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *SalesLiteratureCreateDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *SalesLiteratureCreateDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
### GetSalesLiteratureTypeId

`func (o *SalesLiteratureCreateDto) GetSalesLiteratureTypeId() string`

GetSalesLiteratureTypeId returns the SalesLiteratureTypeId field if non-nil, zero value otherwise.

### GetSalesLiteratureTypeIdOk

`func (o *SalesLiteratureCreateDto) GetSalesLiteratureTypeIdOk() (*string, bool)`

GetSalesLiteratureTypeIdOk returns a tuple with the SalesLiteratureTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesLiteratureTypeId

`func (o *SalesLiteratureCreateDto) SetSalesLiteratureTypeId(v string)`

SetSalesLiteratureTypeId sets SalesLiteratureTypeId field to given value.

### HasSalesLiteratureTypeId

`func (o *SalesLiteratureCreateDto) HasSalesLiteratureTypeId() bool`

HasSalesLiteratureTypeId returns a boolean if a field has been set.

### SetSalesLiteratureTypeIdNil

`func (o *SalesLiteratureCreateDto) SetSalesLiteratureTypeIdNil(b bool)`

 SetSalesLiteratureTypeIdNil sets the value for SalesLiteratureTypeId to be an explicit nil

### UnsetSalesLiteratureTypeId
`func (o *SalesLiteratureCreateDto) UnsetSalesLiteratureTypeId()`

UnsetSalesLiteratureTypeId ensures that no value is present for SalesLiteratureTypeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


