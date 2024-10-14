# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClaimMspLicence**](MSPsLicensesApi.md#ClaimMspLicence) | **Post** /api/v1/msps/{msp_id}/claim | claimMspLicence
[**ListMspLicenses**](MSPsLicensesApi.md#ListMspLicenses) | **Get** /api/v1/msps/{msp_id}/licenses | listMspLicenses
[**ListMspOrgLicenses**](MSPsLicensesApi.md#ListMspOrgLicenses) | **Get** /api/v1/msps/{msp_id}/stats/licenses | listMspOrgLicenses
[**MoveOrDeleteMspLicenseToAnotherOrg**](MSPsLicensesApi.md#MoveOrDeleteMspLicenseToAnotherOrg) | **Put** /api/v1/msps/{msp_id}/licenses | moveOrDeleteMspLicenseToAnotherOrg

# **ClaimMspLicence**
> InlineResponse20017 ClaimMspLicence(ctx, mspId, optional)
claimMspLicence

Claim an Order by Activation Code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
 **optional** | ***MSPsLicensesApiClaimMspLicenceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsLicensesApiClaimMspLicenceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MspIdClaimBody**](MspIdClaimBody.md)|  | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMspLicenses**
> InlineResponse20018 ListMspLicenses(ctx, mspId)
listMspLicenses

Get List of Msp Licenses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMspOrgLicenses**
> InlineResponse20018 ListMspOrgLicenses(ctx, mspId)
listMspOrgLicenses

Get List of MSP Licences

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveOrDeleteMspLicenseToAnotherOrg**
> MoveOrDeleteMspLicenseToAnotherOrg(ctx, mspId, optional)
moveOrDeleteMspLicenseToAnotherOrg

Move or Delete MSP Licenses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
 **optional** | ***MSPsLicensesApiMoveOrDeleteMspLicenseToAnotherOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsLicensesApiMoveOrDeleteMspLicenseToAnotherOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MspIdLicensesBody**](MspIdLicensesBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

