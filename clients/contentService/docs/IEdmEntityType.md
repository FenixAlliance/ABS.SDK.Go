# IEdmEntityType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeKind** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 
**IsAbstract** | Pointer to **bool** |  | [optional] [readonly] 
**IsOpen** | Pointer to **bool** |  | [optional] [readonly] 
**BaseType** | Pointer to [**IEdmStructuredType**](IEdmStructuredType.md) |  | [optional] 
**DeclaredProperties** | Pointer to [**[]IEdmProperty**](IEdmProperty.md) |  | [optional] [readonly] 
**SchemaElementKind** | Pointer to **string** |  | [optional] [readonly] 
**Namespace** | Pointer to **NullableString** |  | [optional] [readonly] 
**DeclaredKey** | Pointer to [**[]IEdmStructuralProperty**](IEdmStructuralProperty.md) |  | [optional] [readonly] 
**HasStream** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewIEdmEntityType

`func NewIEdmEntityType() *IEdmEntityType`

NewIEdmEntityType instantiates a new IEdmEntityType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIEdmEntityTypeWithDefaults

`func NewIEdmEntityTypeWithDefaults() *IEdmEntityType`

NewIEdmEntityTypeWithDefaults instantiates a new IEdmEntityType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeKind

`func (o *IEdmEntityType) GetTypeKind() string`

GetTypeKind returns the TypeKind field if non-nil, zero value otherwise.

### GetTypeKindOk

`func (o *IEdmEntityType) GetTypeKindOk() (*string, bool)`

GetTypeKindOk returns a tuple with the TypeKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeKind

`func (o *IEdmEntityType) SetTypeKind(v string)`

SetTypeKind sets TypeKind field to given value.

### HasTypeKind

`func (o *IEdmEntityType) HasTypeKind() bool`

HasTypeKind returns a boolean if a field has been set.

### GetName

`func (o *IEdmEntityType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IEdmEntityType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IEdmEntityType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IEdmEntityType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IEdmEntityType) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IEdmEntityType) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsAbstract

`func (o *IEdmEntityType) GetIsAbstract() bool`

GetIsAbstract returns the IsAbstract field if non-nil, zero value otherwise.

### GetIsAbstractOk

`func (o *IEdmEntityType) GetIsAbstractOk() (*bool, bool)`

GetIsAbstractOk returns a tuple with the IsAbstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAbstract

`func (o *IEdmEntityType) SetIsAbstract(v bool)`

SetIsAbstract sets IsAbstract field to given value.

### HasIsAbstract

`func (o *IEdmEntityType) HasIsAbstract() bool`

HasIsAbstract returns a boolean if a field has been set.

### GetIsOpen

`func (o *IEdmEntityType) GetIsOpen() bool`

GetIsOpen returns the IsOpen field if non-nil, zero value otherwise.

### GetIsOpenOk

`func (o *IEdmEntityType) GetIsOpenOk() (*bool, bool)`

GetIsOpenOk returns a tuple with the IsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpen

`func (o *IEdmEntityType) SetIsOpen(v bool)`

SetIsOpen sets IsOpen field to given value.

### HasIsOpen

`func (o *IEdmEntityType) HasIsOpen() bool`

HasIsOpen returns a boolean if a field has been set.

### GetBaseType

`func (o *IEdmEntityType) GetBaseType() IEdmStructuredType`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *IEdmEntityType) GetBaseTypeOk() (*IEdmStructuredType, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *IEdmEntityType) SetBaseType(v IEdmStructuredType)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *IEdmEntityType) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetDeclaredProperties

`func (o *IEdmEntityType) GetDeclaredProperties() []IEdmProperty`

GetDeclaredProperties returns the DeclaredProperties field if non-nil, zero value otherwise.

### GetDeclaredPropertiesOk

`func (o *IEdmEntityType) GetDeclaredPropertiesOk() (*[]IEdmProperty, bool)`

GetDeclaredPropertiesOk returns a tuple with the DeclaredProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredProperties

`func (o *IEdmEntityType) SetDeclaredProperties(v []IEdmProperty)`

SetDeclaredProperties sets DeclaredProperties field to given value.

### HasDeclaredProperties

`func (o *IEdmEntityType) HasDeclaredProperties() bool`

HasDeclaredProperties returns a boolean if a field has been set.

### SetDeclaredPropertiesNil

`func (o *IEdmEntityType) SetDeclaredPropertiesNil(b bool)`

 SetDeclaredPropertiesNil sets the value for DeclaredProperties to be an explicit nil

### UnsetDeclaredProperties
`func (o *IEdmEntityType) UnsetDeclaredProperties()`

UnsetDeclaredProperties ensures that no value is present for DeclaredProperties, not even an explicit nil
### GetSchemaElementKind

`func (o *IEdmEntityType) GetSchemaElementKind() string`

GetSchemaElementKind returns the SchemaElementKind field if non-nil, zero value otherwise.

### GetSchemaElementKindOk

`func (o *IEdmEntityType) GetSchemaElementKindOk() (*string, bool)`

GetSchemaElementKindOk returns a tuple with the SchemaElementKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaElementKind

`func (o *IEdmEntityType) SetSchemaElementKind(v string)`

SetSchemaElementKind sets SchemaElementKind field to given value.

### HasSchemaElementKind

`func (o *IEdmEntityType) HasSchemaElementKind() bool`

HasSchemaElementKind returns a boolean if a field has been set.

### GetNamespace

`func (o *IEdmEntityType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IEdmEntityType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IEdmEntityType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IEdmEntityType) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *IEdmEntityType) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *IEdmEntityType) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetDeclaredKey

`func (o *IEdmEntityType) GetDeclaredKey() []IEdmStructuralProperty`

GetDeclaredKey returns the DeclaredKey field if non-nil, zero value otherwise.

### GetDeclaredKeyOk

`func (o *IEdmEntityType) GetDeclaredKeyOk() (*[]IEdmStructuralProperty, bool)`

GetDeclaredKeyOk returns a tuple with the DeclaredKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredKey

`func (o *IEdmEntityType) SetDeclaredKey(v []IEdmStructuralProperty)`

SetDeclaredKey sets DeclaredKey field to given value.

### HasDeclaredKey

`func (o *IEdmEntityType) HasDeclaredKey() bool`

HasDeclaredKey returns a boolean if a field has been set.

### SetDeclaredKeyNil

`func (o *IEdmEntityType) SetDeclaredKeyNil(b bool)`

 SetDeclaredKeyNil sets the value for DeclaredKey to be an explicit nil

### UnsetDeclaredKey
`func (o *IEdmEntityType) UnsetDeclaredKey()`

UnsetDeclaredKey ensures that no value is present for DeclaredKey, not even an explicit nil
### GetHasStream

`func (o *IEdmEntityType) GetHasStream() bool`

GetHasStream returns the HasStream field if non-nil, zero value otherwise.

### GetHasStreamOk

`func (o *IEdmEntityType) GetHasStreamOk() (*bool, bool)`

GetHasStreamOk returns a tuple with the HasStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStream

`func (o *IEdmEntityType) SetHasStream(v bool)`

SetHasStream sets HasStream field to given value.

### HasHasStream

`func (o *IEdmEntityType) HasHasStream() bool`

HasHasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


