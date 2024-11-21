# FileObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Verfication** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **int32** |  | [optional] 
**Overwrite** | Pointer to **bool** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 

## Methods

### NewFileObject

`func NewFileObject() *FileObject`

NewFileObject instantiates a new FileObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileObjectWithDefaults

`func NewFileObjectWithDefaults() *FileObject`

NewFileObjectWithDefaults instantiates a new FileObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FileObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContent

`func (o *FileObject) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FileObject) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FileObject) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *FileObject) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetVerfication

`func (o *FileObject) GetVerfication() string`

GetVerfication returns the Verfication field if non-nil, zero value otherwise.

### GetVerficationOk

`func (o *FileObject) GetVerficationOk() (*string, bool)`

GetVerficationOk returns a tuple with the Verfication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerfication

`func (o *FileObject) SetVerfication(v string)`

SetVerfication sets Verfication field to given value.

### HasVerfication

`func (o *FileObject) HasVerfication() bool`

HasVerfication returns a boolean if a field has been set.

### GetPath

`func (o *FileObject) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileObject) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileObject) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FileObject) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMode

`func (o *FileObject) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FileObject) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FileObject) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FileObject) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOverwrite

`func (o *FileObject) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *FileObject) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *FileObject) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *FileObject) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetUser

`func (o *FileObject) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FileObject) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FileObject) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *FileObject) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetGroup

`func (o *FileObject) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *FileObject) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *FileObject) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *FileObject) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


