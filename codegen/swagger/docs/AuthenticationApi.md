# \AuthenticationApi

All URIs are relative to *https://owner-api.teslamotors.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOauthToken**](AuthenticationApi.md#CreateOauthToken) | **Post** /oauth/token | Get an Access Token


# **CreateOauthToken**
> CreateAccessTokenResponse CreateOauthToken(ctx, body)
Get an Access Token

Performs the login. Takes in an plain text email and password, matching the owner's login information for [https://my.teslamotors.com/user/login](https://my.teslamotors.com/user/login). Returns a `access_token` which is passed along as a header with all future requests to authenticate the user. You must provide the `Authorization: Bearer {access_token}` header in all other requests. The current client ID and secret are [available here](http://pastebin.com/YiLPDggh)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateAccessTokenRequest**](CreateAccessTokenRequest.md)|  | 

### Return type

[**CreateAccessTokenResponse**](CreateAccessTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

