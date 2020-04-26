# \VehiclesApi

All URIs are relative to *https://owner-api.teslamotors.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVehicle**](VehiclesApi.md#GetVehicle) | **Get** /api/1/vehicles/{vehicle_id} | Retrieve a vehicle
[**GetVehicleChargeState**](VehiclesApi.md#GetVehicleChargeState) | **Get** /api/1/vehicles/{vehicle_id}/data_request/charge_state | Charge State
[**GetVehicleClimateState**](VehiclesApi.md#GetVehicleClimateState) | **Get** /api/1/vehicles/{vehicle_id}/data_request/climate_state | Climate Settings
[**GetVehicleConfig**](VehiclesApi.md#GetVehicleConfig) | **Get** /api/1/vehicles/{vehicle_id}/data_request/vehicle_config | Vehicle Config
[**GetVehicleData**](VehiclesApi.md#GetVehicleData) | **Get** /api/1/vehicles/{vehicle_id}/data | Vehicle Data
[**GetVehicleDriveState**](VehiclesApi.md#GetVehicleDriveState) | **Get** /api/1/vehicles/{vehicle_id}/data_request/drive_state | Driving and Position
[**GetVehicleMobileEnabled**](VehiclesApi.md#GetVehicleMobileEnabled) | **Get** /api/1/vehicles/{vehicle_id}/mobile_enabled | Mobile Access
[**GetVehicleState**](VehiclesApi.md#GetVehicleState) | **Get** /api/1/vehicles/{vehicle_id}/data_request/vehicle_state | Vehicle State
[**GetVehicles**](VehiclesApi.md#GetVehicles) | **Get** /api/1/vehicles | List all Vehicles
[**GetVehilceGuiSettings**](VehiclesApi.md#GetVehilceGuiSettings) | **Get** /api/1/vehicles/{vehicle_id}/data_request/gui_settings | GUI Settings
[**NearbyChargers**](VehiclesApi.md#NearbyChargers) | **Get** /api/1/vehicles/{vehicle_id}/nearby_charging_sites | Get Nearby Chargers


# **GetVehicle**
> VehicleResponse GetVehicle(ctx, vehicleId)
Retrieve a vehicle

Retrieve a specific vehicle

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**VehicleResponse**](VehicleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVehicleChargeState**
> ChargeStateResponse GetVehicleChargeState(ctx, vehicleId)
Charge State

Returns the state of charge in the battery.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**ChargeStateResponse**](ChargeStateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVehicleClimateState**
> ClimateSettingsResponse GetVehicleClimateState(ctx, vehicleId)
Climate Settings

Returns the current temperature and climate control state.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**ClimateSettingsResponse**](ClimateSettingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVehicleConfig**
> VehicleConfigResponse GetVehicleConfig(ctx, vehicleId)
Vehicle Config

Returns the vehicle's configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**VehicleConfigResponse**](VehicleConfigResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVehicleData**
> VehicleDataResponse GetVehicleData(ctx, vehicleId)
Vehicle Data

Returns all vehicle ∂ata, potentially cached

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**VehicleDataResponse**](VehicleDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVehicleDriveState**
> DriveStateResponse GetVehicleDriveState(ctx, vehicleId)
Driving and Position

Returns the driving and position state of the vehicle.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**DriveStateResponse**](DriveStateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVehicleMobileEnabled**
> MobileAccessResponse GetVehicleMobileEnabled(ctx, vehicleId)
Mobile Access

Determines if mobile access to the vehicle is enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**MobileAccessResponse**](MobileAccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVehicleState**
> VehicleStateResponse GetVehicleState(ctx, vehicleId)
Vehicle State

Returns the vehicle's physical state, such as which doors are open.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**VehicleStateResponse**](VehicleStateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVehicles**
> GetVehiclesResponse GetVehicles(ctx, )
List all Vehicles

Retrieve a list of your owned vehicles (includes vehicles not yet shipped!)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetVehiclesResponse**](GetVehiclesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVehilceGuiSettings**
> GuistateResponse GetVehilceGuiSettings(ctx, vehicleId)
GUI Settings

Returns various information about the GUI settings of the car, such as unit format and range display.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**GuistateResponse**](GuistateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NearbyChargers**
> NearbyChargerResponse NearbyChargers(ctx, vehicleId)
Get Nearby Chargers

Get a list of nearby charging sites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**NearbyChargerResponse**](NearbyChargerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

