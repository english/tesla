# \MediaControlApi

All URIs are relative to *https://owner-api.teslamotors.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MediaNextFavorite**](MediaControlApi.md#MediaNextFavorite) | **Post** /api/1/vehicles/{vehicle_id}/command/media_next_fav | Next Favorite
[**MediaNextTrack**](MediaControlApi.md#MediaNextTrack) | **Post** /api/1/vehicles/{vehicle_id}/command/media_next_track | Next Track
[**MediaPrevFavorite**](MediaControlApi.md#MediaPrevFavorite) | **Post** /api/1/vehicles/{vehicle_id}/command/media_prev_fav | Previous Favorite
[**MediaPrevTrack**](MediaControlApi.md#MediaPrevTrack) | **Post** /api/1/vehicles/{vehicle_id}/command/media_prev_track | Previous Track
[**MediaTogglePlayback**](MediaControlApi.md#MediaTogglePlayback) | **Post** /api/1/vehicles/{vehicle_id}/command/media_toggle_playback | Pause/Play Media
[**MediaVolumeDown**](MediaControlApi.md#MediaVolumeDown) | **Post** /api/1/vehicles/{vehicle_id}/command/media_volume_down | Volume Down
[**MediaVolumeUp**](MediaControlApi.md#MediaVolumeUp) | **Post** /api/1/vehicles/{vehicle_id}/command/media_volume_up | Volume Up


# **MediaNextFavorite**
> CommandResponse MediaNextFavorite(ctx, vehicleId)
Next Favorite

Next Favorite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MediaNextTrack**
> CommandResponse MediaNextTrack(ctx, vehicleId)
Next Track

Next Track

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MediaPrevFavorite**
> CommandResponse MediaPrevFavorite(ctx, vehicleId)
Previous Favorite

Previous Favorite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MediaPrevTrack**
> CommandResponse MediaPrevTrack(ctx, vehicleId)
Previous Track

Previous Track

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MediaTogglePlayback**
> CommandResponse MediaTogglePlayback(ctx, vehicleId)
Pause/Play Media

Pause/Play Media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MediaVolumeDown**
> CommandResponse MediaVolumeDown(ctx, vehicleId)
Volume Down

Volume Down

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MediaVolumeUp**
> CommandResponse MediaVolumeUp(ctx, vehicleId)
Volume Up

Volume Up

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

