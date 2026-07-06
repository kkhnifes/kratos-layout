# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [conf/conf.proto](#conf_conf-proto)
    - [Bootstrap](#kratos-api-Bootstrap)
    - [Data](#kratos-api-Data)
    - [Data.Database](#kratos-api-Data-Database)
    - [Data.Redis](#kratos-api-Data-Redis)
    - [Log](#kratos-api-Log)
    - [Server](#kratos-api-Server)
    - [Server.GRPC](#kratos-api-Server-GRPC)
    - [Server.HTTP](#kratos-api-Server-HTTP)
  
    - [LogLevel](#kratos-api-LogLevel)
  
- [todo/v1/error_reason.proto](#todo_v1_error_reason-proto)
    - [ErrorReason](#todo-v1-ErrorReason)
  
- [todo/v1/todo.proto](#todo_v1_todo-proto)
    - [CreateTodoRequest](#todo-v1-CreateTodoRequest)
    - [DeleteTodoRequest](#todo-v1-DeleteTodoRequest)
    - [GetTodoRequest](#todo-v1-GetTodoRequest)
    - [ListTodosRequest](#todo-v1-ListTodosRequest)
    - [SyncTodoRequest](#todo-v1-SyncTodoRequest)
    - [Todo](#todo-v1-Todo)
    - [TodoEvent](#todo-v1-TodoEvent)
    - [TodoSet](#todo-v1-TodoSet)
    - [UpdateTodoRequest](#todo-v1-UpdateTodoRequest)
    - [WatchTodosRequest](#todo-v1-WatchTodosRequest)
  
    - [TodoService](#todo-v1-TodoService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="conf_conf-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## conf/conf.proto



<a name="kratos-api-Bootstrap"></a>

### Bootstrap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| server | [Server](#kratos-api-Server) |  |  |
| data | [Data](#kratos-api-Data) |  |  |
| log | [Log](#kratos-api-Log) |  |  |






<a name="kratos-api-Data"></a>

### Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database | [Data.Database](#kratos-api-Data-Database) |  |  |
| redis | [Data.Redis](#kratos-api-Data-Redis) |  |  |






<a name="kratos-api-Data-Database"></a>

### Data.Database



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| driver | [string](#string) |  |  |
| source | [string](#string) |  |  |






<a name="kratos-api-Data-Redis"></a>

### Data.Redis



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network | [string](#string) |  |  |
| addr | [string](#string) |  |  |
| read_timeout | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |
| write_timeout | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |






<a name="kratos-api-Log"></a>

### Log



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| log_path | [string](#string) |  |  |
| log_level | [LogLevel](#kratos-api-LogLevel) |  |  |
| max_size | [int32](#int32) |  |  |
| max_keep_days | [int32](#int32) |  |  |
| max_keep_files | [int32](#int32) |  |  |
| compress | [bool](#bool) |  |  |






<a name="kratos-api-Server"></a>

### Server



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| http | [Server.HTTP](#kratos-api-Server-HTTP) |  |  |
| grpc | [Server.GRPC](#kratos-api-Server-GRPC) |  |  |






<a name="kratos-api-Server-GRPC"></a>

### Server.GRPC



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network | [string](#string) |  |  |
| addr | [string](#string) |  |  |
| timeout | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |






<a name="kratos-api-Server-HTTP"></a>

### Server.HTTP



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network | [string](#string) |  |  |
| addr | [string](#string) |  |  |
| timeout | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |





 


<a name="kratos-api-LogLevel"></a>

### LogLevel


| Name | Number | Description |
| ---- | ------ | ----------- |
| Debug | 0 |  |
| Info | 1 |  |
| Warn | 2 |  |
| Error | 3 |  |
| Fatal | 4 |  |


 

 

 



<a name="todo_v1_error_reason-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## todo/v1/error_reason.proto


 


<a name="todo-v1-ErrorReason"></a>

### ErrorReason


| Name | Number | Description |
| ---- | ------ | ----------- |
| TODO_UNSPECIFIED | 0 |  |
| TODO_NOT_FOUND | 1 |  |
| TODO_INVALID_ARGUMENT | 2 |  |


 

 

 



<a name="todo_v1_todo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## todo/v1/todo.proto



<a name="todo-v1-CreateTodoRequest"></a>

### CreateTodoRequest
CreateTodoRequest is the input for TodoService.CreateTodo.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| todo | [Todo](#todo-v1-Todo) |  | The todo item to create. The id, create_time, and update_time fields are ignored on input and populated by the server in the response. |






<a name="todo-v1-DeleteTodoRequest"></a>

### DeleteTodoRequest
DeleteTodoRequest is the input for TodoService.DeleteTodo.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | Unique identifier of the todo item to delete. |






<a name="todo-v1-GetTodoRequest"></a>

### GetTodoRequest
GetTodoRequest is the input for TodoService.GetTodo.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | Unique identifier of the todo item to retrieve. |






<a name="todo-v1-ListTodosRequest"></a>

### ListTodosRequest
ListTodosRequest is the input for TodoService.ListTodos.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_size | [int32](#int32) |  | Optional. Maximum number of todo items to return in a single page. The server may apply a default and a maximum when unset or out of range. |
| page_token | [string](#string) |  | Optional. Token from a previous response&#39;s next_page_token used to fetch the next page. Leave empty to request the first page. |
| filter | [string](#string) |  | Optional. The standard list filter. Supported fields: * `title` (i.e. `title:&#34;bug&#34;`) * `content` (i.e. `content:&#34;docs&#34;`) * `completed` (i.e. `completed` or `NOT completed`) * `create_time` range (i.e. `create_time&gt;=&#34;2026-01-01T00:00:00Z&#34;`) |
| order_by | [string](#string) |  | Optional. A comma-separated list of fields to order by. Supported fields: * `id` * `title` * `create_time` * `update_time` Append ` desc` to a field for descending order, e.g. `create_time desc`. Defaults to ascending order when no direction is supplied. |






<a name="todo-v1-SyncTodoRequest"></a>

### SyncTodoRequest
SyncTodoRequest is a single client message in the SyncTodos bidirectional
stream. The action field selects which operation to perform on the server.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [string](#string) |  | The operation to perform. One of `create`, `update`, or `delete`. |
| todo | [Todo](#todo-v1-Todo) |  | Payload for `create` and `update` actions. Ignored for `delete`. |
| id | [int64](#int64) |  | Identifier of the target todo. Required for `delete` and `update`. |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  | Field mask used by `update` actions to limit which fields are overwritten. Ignored for `create` and `delete`. |






<a name="todo-v1-Todo"></a>

### Todo
Todo is the canonical representation of a todo item.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | Server-assigned unique identifier. Read-only on create. |
| title | [string](#string) |  | Short human-readable title. |
| content | [string](#string) |  | Detailed description or notes for the todo item. |
| completed | [bool](#bool) |  | Whether the todo item has been completed. |
| create_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | Time at which the todo item was created. Server-assigned, read-only. |
| update_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | Time at which the todo item was last modified. Server-assigned, read-only. |






<a name="todo-v1-TodoEvent"></a>

### TodoEvent
TodoEvent is a server-sent message describing a change to a todo item.
It is emitted by both WatchTodos and SyncTodos streams.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [string](#string) |  | The change that occurred. One of `created`, `updated`, or `deleted`. |
| todo | [Todo](#todo-v1-Todo) |  | The todo item after the change. For `deleted` events only the id field is guaranteed to be populated. |
| event_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | Time at which the change was observed by the server. |






<a name="todo-v1-TodoSet"></a>

### TodoSet
TodoSet is a paginated collection of todo items returned by ListTodos.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| todos | [Todo](#todo-v1-Todo) | repeated | The page of todo items in the order requested by ListTodosRequest.order_by. |
| next_page_token | [string](#string) |  | Token used to retrieve the next page via ListTodosRequest.page_token. Empty when there are no more results. |






<a name="todo-v1-UpdateTodoRequest"></a>

### UpdateTodoRequest
UpdateTodoRequest is the input for TodoService.UpdateTodo.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| todo | [Todo](#todo-v1-Todo) |  | The todo carrying the new field values. todo.id identifies the target record and must be set. |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  | Set of field paths in todo to overwrite. Fields not listed are left unchanged. Use a single path of `*` to replace every mutable field. |






<a name="todo-v1-WatchTodosRequest"></a>

### WatchTodosRequest
WatchTodosRequest is the input for TodoService.WatchTodos.
It selects which todo items the resulting event stream observes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_size | [int32](#int32) |  | Optional. Maximum number of todo items returned in the initial snapshot before live events begin streaming. |
| page_token | [string](#string) |  | Optional. Page token used to resume from a prior snapshot. |
| filter | [string](#string) |  | Optional. List filter expression. Same syntax as ListTodosRequest.filter. |
| order_by | [string](#string) |  | Optional. Order specification. Same syntax as ListTodosRequest.order_by. |





 

 

 


<a name="todo-v1-TodoService"></a>

### TodoService
TodoService provides a simple CRUD example for managing todo items.
It demonstrates unary, server-streaming, and bidirectional-streaming RPCs
together with AIP-style list filtering, ordering, and pagination.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateTodo | [CreateTodoRequest](#todo-v1-CreateTodoRequest) | [Todo](#todo-v1-Todo) | CreateTodo creates a new todo item and returns the persisted record with the server-assigned id, create_time, and update_time populated. Returns INVALID_ARGUMENT if the request payload fails validation. |
| GetTodo | [GetTodoRequest](#todo-v1-GetTodoRequest) | [Todo](#todo-v1-Todo) | GetTodo returns a single todo item by its id. Returns NOT_FOUND if no todo exists with the supplied id. |
| ListTodos | [ListTodosRequest](#todo-v1-ListTodosRequest) | [TodoSet](#todo-v1-TodoSet) | ListTodos returns a page of todo items, optionally filtered and ordered. Use the next_page_token from TodoSet to retrieve subsequent pages. Returns INVALID_ARGUMENT if filter, order_by, or page_token are malformed. |
| UpdateTodo | [UpdateTodoRequest](#todo-v1-UpdateTodoRequest) | [Todo](#todo-v1-Todo) | UpdateTodo applies a partial update to an existing todo item using a FieldMask. Only the fields listed in update_mask are overwritten; all other fields are left unchanged. Returns NOT_FOUND if the target todo does not exist, or INVALID_ARGUMENT if update_mask references unknown fields. |
| DeleteTodo | [DeleteTodoRequest](#todo-v1-DeleteTodoRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | DeleteTodo permanently removes a todo item by its id. Returns NOT_FOUND if no todo exists with the supplied id. |
| WatchTodos | [WatchTodosRequest](#todo-v1-WatchTodosRequest) | [TodoEvent](#todo-v1-TodoEvent) stream | WatchTodos opens a server-side stream that emits a TodoEvent for every create, update, or delete that matches the supplied filter and ordering. The stream remains open until the client cancels or the server terminates it. |
| SyncTodos | [SyncTodoRequest](#todo-v1-SyncTodoRequest) stream | [TodoEvent](#todo-v1-TodoEvent) stream | SyncTodos opens a bidirectional stream for two-way synchronization. The client sends SyncTodoRequest messages describing local changes, and the server pushes back TodoEvent messages reflecting the resulting state. The stream stays open until either side closes it. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

