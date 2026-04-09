# InvoiceReferenceCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**ReferencedInvoiceId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInvoiceReferenceCreateDto

`func NewInvoiceReferenceCreateDto() *InvoiceReferenceCreateDto`

NewInvoiceReferenceCreateDto instantiates a new InvoiceReferenceCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceReferenceCreateDtoWithDefaults

`func NewInvoiceReferenceCreateDtoWithDefaults() *InvoiceReferenceCreateDto`

NewInvoiceReferenceCreateDtoWithDefaults instantiates a new InvoiceReferenceCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceReferenceCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceReferenceCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceReferenceCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceReferenceCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *InvoiceReferenceCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InvoiceReferenceCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InvoiceReferenceCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InvoiceReferenceCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetReferencedInvoiceId

`func (o *InvoiceReferenceCreateDto) GetReferencedInvoiceId() string`

GetReferencedInvoiceId returns the ReferencedInvoiceId field if non-nil, zero value otherwise.

### GetReferencedInvoiceIdOk

`func (o *InvoiceReferenceCreateDto) GetReferencedInvoiceIdOk() (*string, bool)`

GetReferencedInvoiceIdOk returns a tuple with the ReferencedInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedInvoiceId

`func (o *InvoiceReferenceCreateDto) SetReferencedInvoiceId(v string)`

SetReferencedInvoiceId sets ReferencedInvoiceId field to given value.

### HasReferencedInvoiceId

`func (o *InvoiceReferenceCreateDto) HasReferencedInvoiceId() bool`

HasReferencedInvoiceId returns a boolean if a field has been set.

### SetReferencedInvoiceIdNil

`func (o *InvoiceReferenceCreateDto) SetReferencedInvoiceIdNil(b bool)`

 SetReferencedInvoiceIdNil sets the value for ReferencedInvoiceId to be an explicit nil

### UnsetReferencedInvoiceId
`func (o *InvoiceReferenceCreateDto) UnsetReferencedInvoiceId()`

UnsetReferencedInvoiceId ensures that no value is present for ReferencedInvoiceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


