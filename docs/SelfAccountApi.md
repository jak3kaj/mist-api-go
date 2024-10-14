# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSelf**](SelfAccountApi.md#DeleteSelf) | **Delete** /api/v1/self | deleteSelf
[**GetSelf**](SelfAccountApi.md#GetSelf) | **Get** /api/v1/self | getSelf
[**UpdateSelf**](SelfAccountApi.md#UpdateSelf) | **Put** /api/v1/self | updateSelf
[**UpdateSelfEmail**](SelfAccountApi.md#UpdateSelfEmail) | **Post** /api/v1/self/update | updateSelfEmail
[**VerifySelfEmail**](SelfAccountApi.md#VerifySelfEmail) | **Get** /api/v1/self/update/verify/{token} | verifySelfEmail

# **DeleteSelf**
> DeleteSelf(ctx, )
deleteSelf

To delete ones account and every associated with it. The effects:  the account would be deleted any orphaned Org (that only has this account as admin) will be deleted along with all data with Org (sites, wlans, devices) will be gone.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSelf**
> InlineResponse20012 GetSelf(ctx, )
getSelf

Get ‘whoami’ and privileges (which org and which sites I have access to)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSelf**
> InlineResponse20012 UpdateSelf(ctx, optional)
updateSelf

update Account Information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SelfAccountApiUpdateSelfOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SelfAccountApiUpdateSelfOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of V1SelfBody**](V1SelfBody.md)| Request Body | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSelfEmail**
> UpdateSelfEmail(ctx, optional)
updateSelfEmail

Change Email We require the user to verify that they actually own the email address they intend to change it to.  After the API call, the user will receive an email to the new email address with a link like https://manage.mist.com/verify/update?expire=:exp_time&email=:admin_email&token=:token  Upon clicking the link, the user is provided with a login page to authenticate using existing credentials. After successful login, the email address of the user gets updated  **Note**: The request parameter email can be used by UI to validate that the current session (if any) belongs to the admin or provide a login page (by pre-populating the email on login screen). UI can also use the request parameter expire to validate token expiry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SelfAccountApiUpdateSelfEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SelfAccountApiUpdateSelfEmailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SelfUpdateBody**](SelfUpdateBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifySelfEmail**
> VerifySelfEmail(ctx, token)
verifySelfEmail

Verify Email change

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

