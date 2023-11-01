# \BuildApi

All URIs are relative to *http://localhost:3003/api/v1/dynamic*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetArtifact**](BuildApi.md#GetArtifact) | **Get** /artifacts/{artifactId} | Reads information about an artifact.
[**GetArtifactData**](BuildApi.md#GetArtifactData) | **Get** /artifacts/{artifactId}/data | Reads the data for an artifact.
[**GetBuild**](BuildApi.md#GetBuild) | **Get** /builds/{buildId} | Reads the current build graph for a build.
[**GetJob**](BuildApi.md#GetJob) | **Get** /jobs/{jobId} | Reads information about a job.
[**GetJobGraph**](BuildApi.md#GetJobGraph) | **Get** /jobs/{jobId}/graph | Reads information about a job&#39;s graph.
[**GetLogData**](BuildApi.md#GetLogData) | **Get** /logs/{logDescriptorId}/data | Reads part of a log.
[**GetLogDescriptor**](BuildApi.md#GetLogDescriptor) | **Get** /logs/{logDescriptorId} | Fetches a Log Descriptor containing information about part of a log.
[**ListArtifacts**](BuildApi.md#ListArtifacts) | **Get** /builds/{buildId}/artifacts | Reads information about all or some artifacts from a build.
[**Ping**](BuildApi.md#Ping) | **Get** /ping | Checks for connectivity to the Dynamic API.



## GetArtifact

> Artifact GetArtifact(ctx, artifactId).Execute()

Reads information about an artifact.



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
    artifactId := "artifact:5238115e-070a-44fe-bce0-b43582583eff" // string | The ID of the artifact to read information about.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildApi.GetArtifact(context.Background(), artifactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildApi.GetArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifact`: Artifact
    fmt.Fprintf(os.Stdout, "Response from `BuildApi.GetArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artifactId** | **string** | The ID of the artifact to read information about. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Artifact**](Artifact.md)

### Authorization

[jwt_build_token](../README.md#jwt_build_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactData

> *os.File GetArtifactData(ctx, artifactId).Execute()

Reads the data for an artifact.



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
    artifactId := "artifact:5238115e-070a-44fe-bce0-b43582583eff" // string | The ID of the artifact to read information about.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildApi.GetArtifactData(context.Background(), artifactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildApi.GetArtifactData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifactData`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `BuildApi.GetArtifactData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artifactId** | **string** | The ID of the artifact to read information about. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[jwt_build_token](../README.md#jwt_build_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuild

> BuildGraph GetBuild(ctx, buildId).Execute()

Reads the current build graph for a build.



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
    buildId := "build:4738115e-070a-44fe-bce0-b43582583eaa" // string | The ID of the build to read.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildApi.GetBuild(context.Background(), buildId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildApi.GetBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuild`: BuildGraph
    fmt.Fprintf(os.Stdout, "Response from `BuildApi.GetBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | The ID of the build to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuildGraph**](BuildGraph.md)

### Authorization

[jwt_build_token](../README.md#jwt_build_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJob

> Job GetJob(ctx, jobId).Execute()

Reads information about a job.



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
    jobId := "job:5238115e-070a-44fe-bce0-b43582583eff" // string | The ID of the job to read.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildApi.GetJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildApi.GetJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJob`: Job
    fmt.Fprintf(os.Stdout, "Response from `BuildApi.GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The ID of the job to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Job**](Job.md)

### Authorization

[jwt_build_token](../README.md#jwt_build_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobGraph

> JobGraph GetJobGraph(ctx, jobId).Execute()

Reads information about a job's graph.



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
    jobId := "job:5238115e-070a-44fe-bce0-b43582583eff" // string | The ID of the job to read.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildApi.GetJobGraph(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildApi.GetJobGraph``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobGraph`: JobGraph
    fmt.Fprintf(os.Stdout, "Response from `BuildApi.GetJobGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The ID of the job to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobGraph**](JobGraph.md)

### Authorization

[jwt_build_token](../README.md#jwt_build_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogData

> *os.File GetLogData(ctx, logDescriptorId).Start(start).Plaintext(plaintext).Expand(expand).Execute()

Reads part of a log.



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
    logDescriptorId := "log-descriptor:5238115e-070a-44fe-bce0-b43582583eff" // string | A Log Descriptor ID identifying the part of the log to read data from.
    start := int32(56) // int32 | If provided, only log entries with sequence numbers greater than or equal to start will be returned. Note that expand and start cannot both be specified. (optional)
    plaintext := true // bool | True if the log data should be streamed as plain text, or false to stream a series of log entries as JSON. (optional)
    expand := true // bool | True to expand nested log blocks in the returned data, or false to only include a summary. Note that expand and start cannot both be specified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildApi.GetLogData(context.Background(), logDescriptorId).Start(start).Plaintext(plaintext).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildApi.GetLogData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogData`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `BuildApi.GetLogData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logDescriptorId** | **string** | A Log Descriptor ID identifying the part of the log to read data from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | If provided, only log entries with sequence numbers greater than or equal to start will be returned. Note that expand and start cannot both be specified. | 
 **plaintext** | **bool** | True if the log data should be streamed as plain text, or false to stream a series of log entries as JSON. | 
 **expand** | **bool** | True to expand nested log blocks in the returned data, or false to only include a summary. Note that expand and start cannot both be specified. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[jwt_build_token](../README.md#jwt_build_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogDescriptor

> LogDescriptor GetLogDescriptor(ctx, logDescriptorId).Execute()

Fetches a Log Descriptor containing information about part of a log.



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
    logDescriptorId := "log-descriptor:5238115e-070a-44fe-bce0-b43582583eff" // string | The ID of the log descriptor to read.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildApi.GetLogDescriptor(context.Background(), logDescriptorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildApi.GetLogDescriptor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogDescriptor`: LogDescriptor
    fmt.Fprintf(os.Stdout, "Response from `BuildApi.GetLogDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logDescriptorId** | **string** | The ID of the log descriptor to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogDescriptor**](LogDescriptor.md)

### Authorization

[jwt_build_token](../README.md#jwt_build_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArtifacts

> ArtifactsPaginatedResponse ListArtifacts(ctx, buildId).Workflow(workflow).JobName(jobName).GroupName(groupName).Cursor(cursor).Limit(limit).Execute()

Reads information about all or some artifacts from a build.



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
    buildId := "build:4738115e-070a-44fe-bce0-b43582583eaa" // string | The ID of the build to read artifacts for.
    workflow := "build-all" // string | If provided, only artifacts produced by this workflow will be returned. (optional)
    jobName := "build-all" // string | If provided, only artifacts produced by this job will be returned. (optional)
    groupName := "reports" // string | If provided, only artifacts associated with this artifact group name will be returned. This name much match the name provided in the build definition. (optional)
    cursor := "cursor_example" // string | An opaque value obtained from a prior results page that can be used to request the next or previous page of results. If not specified then the first page of results will be returned. (optional)
    limit := int32(56) // int32 | The maximum number of results to return from this call. Additional results will be available in other pages via the returned cursor values. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildApi.ListArtifacts(context.Background(), buildId).Workflow(workflow).JobName(jobName).GroupName(groupName).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildApi.ListArtifacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArtifacts`: ArtifactsPaginatedResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildApi.ListArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | The ID of the build to read artifacts for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflow** | **string** | If provided, only artifacts produced by this workflow will be returned. | 
 **jobName** | **string** | If provided, only artifacts produced by this job will be returned. | 
 **groupName** | **string** | If provided, only artifacts associated with this artifact group name will be returned. This name much match the name provided in the build definition. | 
 **cursor** | **string** | An opaque value obtained from a prior results page that can be used to request the next or previous page of results. If not specified then the first page of results will be returned. | 
 **limit** | **int32** | The maximum number of results to return from this call. Additional results will be available in other pages via the returned cursor values. | 

### Return type

[**ArtifactsPaginatedResponse**](ArtifactsPaginatedResponse.md)

### Authorization

[jwt_build_token](../README.md#jwt_build_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ping

> Ping(ctx).Execute()

Checks for connectivity to the Dynamic API.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BuildApi.Ping(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildApi.Ping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPingRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[jwt_build_token](../README.md#jwt_build_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

