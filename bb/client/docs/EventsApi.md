# \EventsApi

All URIs are relative to *http://localhost:3003/api/v1/dynamic*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvents**](EventsApi.md#GetEvents) | **Get** /builds/{buildId}/events | Reads events relating to a build.



## GetEvents

> []Event GetEvents(ctx, buildId).Last(last).Limit(limit).Execute()

Reads events relating to a build.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
    buildId := "build:4738115e-070a-44fe-bce0-b43582583eaa" // string | The ID of the build to read events for
    last := int64(350) // int64 | The sequence number of the last event previously received; only events with higher sequence numbers than this will be returned. (optional)
    limit := int64(100) // int64 | The maximum number of events to return (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEvents(context.Background(), buildId).Last(last).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvents`: []Event
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | The ID of the build to read events for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **last** | **int64** | The sequence number of the last event previously received; only events with higher sequence numbers than this will be returned. | 
 **limit** | **int64** | The maximum number of events to return | 

### Return type

[**[]Event**](Event.md)

### Authorization

[jwt_build_token](../README.md#jwt_build_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

