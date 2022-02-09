# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetManager**](DefaultApi.md#GetManager) | **Get** /redfish/v1/Managers/{managerId}/ | 
[**GetManagerVirtualMedia**](DefaultApi.md#GetManagerVirtualMedia) | **Get** /redfish/v1/Managers/{managerId}/VirtualMedia/{virtualMediaId}/ | 
[**GetRoot**](DefaultApi.md#GetRoot) | **Get** /redfish/v1/ | 
[**GetSystem**](DefaultApi.md#GetSystem) | **Get** /redfish/v1/Systems/{systemId}/ | 
[**InsertVirtualMedia**](DefaultApi.md#InsertVirtualMedia) | **Post** /redfish/v1/Managers/{managerId}/VirtualMedia/{virtualMediaId}/Actions/VirtualMedia.InsertMedia | 
[**ListManagerVirtualMedia**](DefaultApi.md#ListManagerVirtualMedia) | **Get** /redfish/v1/Managers/{managerId}/VirtualMedia/ | 
[**ListManagers**](DefaultApi.md#ListManagers) | **Get** /redfish/v1/Managers/ | 
[**ListSystems**](DefaultApi.md#ListSystems) | **Get** /redfish/v1/Systems/ | 
[**ResetSystem**](DefaultApi.md#ResetSystem) | **Post** /redfish/v1/Systems/{ComputerSystemId}/Actions/ComputerSystem.Reset/ | 
[**SetSystem**](DefaultApi.md#SetSystem) | **Patch** /redfish/v1/Systems/{systemId}/ | 



## GetManager

> Manager GetManager(ctx, managerId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    managerId := "managerId_example" // string | ID of resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetManager(context.Background(), managerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManager`: Manager
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managerId** | **string** | ID of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Manager**](Manager.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagerVirtualMedia

> VirtualMedia GetManagerVirtualMedia(ctx, managerId, virtualMediaId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    managerId := "managerId_example" // string | ID of resource
    virtualMediaId := "virtualMediaId_example" // string | ID of resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetManagerVirtualMedia(context.Background(), managerId, virtualMediaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetManagerVirtualMedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagerVirtualMedia`: VirtualMedia
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetManagerVirtualMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managerId** | **string** | ID of resource | 
**virtualMediaId** | **string** | ID of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagerVirtualMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VirtualMedia**](VirtualMedia.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoot

> Root GetRoot(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoot`: Root
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRoot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRootRequest struct via the builder pattern


### Return type

[**Root**](Root.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystem

> ComputerSystem GetSystem(ctx, systemId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | ID of resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSystem(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystem`: ComputerSystem
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputerSystem**](ComputerSystem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertVirtualMedia

> RedfishError InsertVirtualMedia(ctx, managerId, virtualMediaId).InsertMediaRequestBody(insertMediaRequestBody).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    managerId := "managerId_example" // string | ID of resource
    virtualMediaId := "virtualMediaId_example" // string | ID of resource
    insertMediaRequestBody := *openapiclient.NewInsertMediaRequestBody("Image_example") // InsertMediaRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.InsertVirtualMedia(context.Background(), managerId, virtualMediaId).InsertMediaRequestBody(insertMediaRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InsertVirtualMedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsertVirtualMedia`: RedfishError
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.InsertVirtualMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managerId** | **string** | ID of resource | 
**virtualMediaId** | **string** | ID of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsertVirtualMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **insertMediaRequestBody** | [**InsertMediaRequestBody**](InsertMediaRequestBody.md) |  | 

### Return type

[**RedfishError**](RedfishError.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListManagerVirtualMedia

> Collection ListManagerVirtualMedia(ctx, managerId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    managerId := "managerId_example" // string | ID of resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListManagerVirtualMedia(context.Background(), managerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListManagerVirtualMedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListManagerVirtualMedia`: Collection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListManagerVirtualMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managerId** | **string** | ID of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiListManagerVirtualMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListManagers

> Collection ListManagers(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListManagers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListManagers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListManagers`: Collection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListManagers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListManagersRequest struct via the builder pattern


### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystems

> Collection ListSystems(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSystems(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSystems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSystems`: Collection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSystems`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSystemsRequest struct via the builder pattern


### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetSystem

> RedfishError ResetSystem(ctx, computerSystemId).ResetRequestBody(resetRequestBody).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    computerSystemId := "computerSystemId_example" // string | 
    resetRequestBody := *openapiclient.NewResetRequestBody() // ResetRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ResetSystem(context.Background(), computerSystemId).ResetRequestBody(resetRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ResetSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetSystem`: RedfishError
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ResetSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**computerSystemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resetRequestBody** | [**ResetRequestBody**](ResetRequestBody.md) |  | 

### Return type

[**RedfishError**](RedfishError.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSystem

> ComputerSystem SetSystem(ctx, systemId).ComputerSystem(computerSystem).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    systemId := "systemId_example" // string | ID of resource
    computerSystem := *openapiclient.NewComputerSystem("Name_example", "OdataType_example", "OdataId_example") // ComputerSystem | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SetSystem(context.Background(), systemId).ComputerSystem(computerSystem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSystem`: ComputerSystem
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SetSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ID of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **computerSystem** | [**ComputerSystem**](ComputerSystem.md) |  | 

### Return type

[**ComputerSystem**](ComputerSystem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

