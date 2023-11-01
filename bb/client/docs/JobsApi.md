# \JobsApi

All URIs are relative to *http://localhost:3003/api/v1/dynamic*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJobs**](JobsApi.md#CreateJobs) | **Post** /builds/{buildId}/jobs | Creates and add a set of jobs to a build.



## CreateJobs

> []JobGraph CreateJobs(ctx, buildId).BuildDefinition(buildDefinition).Execute()

Creates and add a set of jobs to a build.



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
    buildId := "build:4738115e-070a-44fe-bce0-b43582583eaa" // string | The ID of the build to which these jobs will be added
    buildDefinition := *openapiclient.NewBuildDefinition("Version_example", []openapiclient.JobDefinition{*openapiclient.NewJobDefinition("generate-code, or workflow1.generate-code", "StepExecution_example", map[string]SecretStringDefinition{"key": *openapiclient.NewSecretStringDefinition()}, []openapiclient.StepDefinition{*openapiclient.NewStepDefinition("compile", []string{"Commands_example"})})}) // BuildDefinition | Definitions for a set of jobs to create and add to a build

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.CreateJobs(context.Background(), buildId).BuildDefinition(buildDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.CreateJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJobs`: []JobGraph
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.CreateJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | The ID of the build to which these jobs will be added | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildDefinition** | [**BuildDefinition**](BuildDefinition.md) | Definitions for a set of jobs to create and add to a build | 

### Return type

[**[]JobGraph**](JobGraph.md)

### Authorization

[jwt_build_token](../README.md#jwt_build_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

