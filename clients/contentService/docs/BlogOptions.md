# BlogOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostsPerPage** | Pointer to **int32** |  | [optional] 
**EnableBlogPostComments** | Pointer to **bool** |  | [optional] 

## Methods

### NewBlogOptions

`func NewBlogOptions() *BlogOptions`

NewBlogOptions instantiates a new BlogOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlogOptionsWithDefaults

`func NewBlogOptionsWithDefaults() *BlogOptions`

NewBlogOptionsWithDefaults instantiates a new BlogOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostsPerPage

`func (o *BlogOptions) GetPostsPerPage() int32`

GetPostsPerPage returns the PostsPerPage field if non-nil, zero value otherwise.

### GetPostsPerPageOk

`func (o *BlogOptions) GetPostsPerPageOk() (*int32, bool)`

GetPostsPerPageOk returns a tuple with the PostsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostsPerPage

`func (o *BlogOptions) SetPostsPerPage(v int32)`

SetPostsPerPage sets PostsPerPage field to given value.

### HasPostsPerPage

`func (o *BlogOptions) HasPostsPerPage() bool`

HasPostsPerPage returns a boolean if a field has been set.

### GetEnableBlogPostComments

`func (o *BlogOptions) GetEnableBlogPostComments() bool`

GetEnableBlogPostComments returns the EnableBlogPostComments field if non-nil, zero value otherwise.

### GetEnableBlogPostCommentsOk

`func (o *BlogOptions) GetEnableBlogPostCommentsOk() (*bool, bool)`

GetEnableBlogPostCommentsOk returns a tuple with the EnableBlogPostComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBlogPostComments

`func (o *BlogOptions) SetEnableBlogPostComments(v bool)`

SetEnableBlogPostComments sets EnableBlogPostComments field to given value.

### HasEnableBlogPostComments

`func (o *BlogOptions) HasEnableBlogPostComments() bool`

HasEnableBlogPostComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


