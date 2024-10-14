# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdminRegistrationInfo**](AdminsApi.md#GetAdminRegistrationInfo) | **Get** /api/v1/register/recaptcha | getAdminRegistrationInfo
[**RegisterNewAdmin**](AdminsApi.md#RegisterNewAdmin) | **Post** /api/v1/register | registerNewAdmin
[**VerifyAdminInvite**](AdminsApi.md#VerifyAdminInvite) | **Post** /api/v1/invite/verify/{token} | verifyAdminInvite
[**VerifyRegistration**](AdminsApi.md#VerifyRegistration) | **Post** /api/v1/register/verify/{token} | verifyRegistration

# **GetAdminRegistrationInfo**
> InlineResponse200 GetAdminRegistrationInfo(ctx, optional)
getAdminRegistrationInfo

Get Registration Information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdminsApiGetAdminRegistrationInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminsApiGetAdminRegistrationInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recaptchaFlavor** | [**optional.Interface of RecaptchaFlavor1**](.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterNewAdmin**
> RegisterNewAdmin(ctx, optional)
registerNewAdmin

Register a new admin and his/her org An email will also be sent to the user with a link to `/verify/register?token={token}`  ### reCAPTCHA Google reCAPTCHA is the choice to prevent bot registration  It needs this   &lt;script src='https://www.google.com/recaptcha/api.js' &gt;&lt;/script&gt;  and this &lt;div&gt; in the desired place ```html <div class=\"g-recaptcha\" data_sitekey=\"6LdAewsTAAAAAE25XKQhPEQ2FiMTft-WrZXQ5NUd\"></div> ```  Use GET /api/v1/register/recaptcha to read the current setting. Response example: ```json {       \"flavor\": \"google\",   \"required\": true,       \"sitekey\": \"6LdAewsTAAAAAE25XKQhPEQ2FiMTft-WrZXQ5NUd\" } ```  ### hCaptcha Alternative to reCAPTCHA is hCaptcha to prevent bot registration  It needs this script  &lt;script src='https://js.hcaptcha.com/1/api.js' async defer &gt;&lt;/script&gt;  and this &lt;div&gt; in the desired place ```html <div class=\"h-recaptcha\" data_sitekey=\"6LdAewsTAAAAAE25XKQhPEQ2FiMTft-WrZXQ5NUd\"></div> ```  Use GET /api/v1/register/recaptcha?recaptcha_flavor=hcaptcha to read the current setting for hcaptcha with reply. Response example: ```json {   \"flavor\": \"hcaptcha\",   \"required\": true,   \"sitekey\": \"6LdAewsTAAAAAE25XKQhPEQ2FiMTft-WrZXQ5NUd\" }\" ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdminsApiRegisterNewAdminOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminsApiRegisterNewAdminOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of V1RegisterBody**](V1RegisterBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyAdminInvite**
> VerifyAdminInvite(ctx, token)
verifyAdminInvite

**Note**: another call to ```GET /api/v1/self``` is required to see the new set of privileges

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

# **VerifyRegistration**
> InlineResponse2001 VerifyRegistration(ctx, token)
verifyRegistration

Verify registration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

