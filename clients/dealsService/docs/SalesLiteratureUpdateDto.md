# SalesLiteratureUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ModifiedDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrolmentId** | Pointer to **NullableString** |  | [optional] 
**SalesLiteratureTypeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSalesLiteratureUpdateDto

`func NewSalesLiteratureUpdateDto() *SalesLiteratureUpdateDto`

NewSalesLiteratureUpdateDto instantiates a new SalesLiteratureUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesLiteratureUpdateDtoWithDefaults

`func NewSalesLiteratureUpdateDtoWithDefaults() *SalesLiteratureUpdateDto`

NewSalesLiteratureUpdateDtoWithDefaults instantiates a new SalesLiteratureUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SalesLiteratureUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SalesLiteratureUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SalesLiteratureUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SalesLiteratureUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SalesLiteratureUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SalesLiteratureUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetContent

`func (o *SalesLiteratureUpdateDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SalesLiteratureUpdateDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SalesLiteratureUpdateDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SalesLiteratureUpdateDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *SalesLiteratureUpdateDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *SalesLiteratureUpdateDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetDescription

`func (o *SalesLiteratureUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SalesLiteratureUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SalesLiteratureUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SalesLiteratureUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SalesLiteratureUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SalesLiteratureUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetModifiedDate

`func (o *SalesLiteratureUpdateDto) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *SalesLiteratureUpdateDto) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *SalesLiteratureUpdateDto) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *SalesLiteratureUpdateDto) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *SalesLiteratureUpdateDto) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *SalesLiteratureUpdateDto) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *SalesLiteratureUpdateDto) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *SalesLiteratureUpdateDto) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetTenantId

`func (o *SalesLiteratureUpdateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SalesLiteratureUpdateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SalesLiteratureUpdateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SalesLiteratureUpdateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SalesLiteratureUpdateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SalesLiteratureUpdateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrolmentId

`func (o *SalesLiteratureUpdateDto) GetEnrolmentId() string`

GetEnrolmentId returns the EnrolmentId field if non-nil, zero value otherwise.

### GetEnrolmentIdOk

`func (o *SalesLiteratureUpdateDto) GetEnrolmentIdOk() (*string, bool)`

GetEnrolmentIdOk returns a tuple with the EnrolmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolmentId

`func (o *SalesLiteratureUpdateDto) SetEnrolmentId(v string)`

SetEnrolmentId sets EnrolmentId field to given value.

### HasEnrolmentId

`func (o *SalesLiteratureUpdateDto) HasEnrolmentId() bool`

HasEnrolmentId returns a boolean if a field has been set.

### SetEnrolmentIdNil

`func (o *SalesLiteratureUpdateDto) SetEnrolmentIdNil(b bool)`

 SetEnrolmentIdNil sets the value for EnrolmentId to be an explicit nil

### UnsetEnrolmentId
`func (o *SalesLiteratureUpdateDto) UnsetEnrolmentId()`

UnsetEnrolmentId ensures that no value is present for EnrolmentId, not even an explicit nil
### GetSalesLiteratureTypeId

`func (o *SalesLiteratureUpdateDto) GetSalesLiteratureTypeId() string`

GetSalesLiteratureTypeId returns the SalesLiteratureTypeId field if non-nil, zero value otherwise.

### GetSalesLiteratureTypeIdOk

`func (o *SalesLiteratureUpdateDto) GetSalesLiteratureTypeIdOk() (*string, bool)`

GetSalesLiteratureTypeIdOk returns a tuple with the SalesLiteratureTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesLiteratureTypeId

`func (o *SalesLiteratureUpdateDto) SetSalesLiteratureTypeId(v string)`

SetSalesLiteratureTypeId sets SalesLiteratureTypeId field to given value.

### HasSalesLiteratureTypeId

`func (o *SalesLiteratureUpdateDto) HasSalesLiteratureTypeId() bool`

HasSalesLiteratureTypeId returns a boolean if a field has been set.

### SetSalesLiteratureTypeIdNil

`func (o *SalesLiteratureUpdateDto) SetSalesLiteratureTypeIdNil(b bool)`

 SetSalesLiteratureTypeIdNil sets the value for SalesLiteratureTypeId to be an explicit nil

### UnsetSalesLiteratureTypeId
`func (o *SalesLiteratureUpdateDto) UnsetSalesLiteratureTypeId()`

UnsetSalesLiteratureTypeId ensures that no value is present for SalesLiteratureTypeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


