# \VehicleCommandsApi

All URIs are relative to *https://owner-api.teslamotors.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSpeedLimit**](VehicleCommandsApi.md#ActivateSpeedLimit) | **Post** /api/1/vehicles/{vehicle_id}/command/speed_limit_activate | Activate Speed Limit
[**CancelSoftwareUpdate**](VehicleCommandsApi.md#CancelSoftwareUpdate) | **Post** /api/1/vehicles/{vehicle_id}/command/cancel_software_update | Cancel Software Update
[**ClearSpeedLimitPin**](VehicleCommandsApi.md#ClearSpeedLimitPin) | **Post** /api/1/vehicles/{vehicle_id}/command/speed_limit_clear_pin | Clear Speed Limit Pin
[**CloseChargePort**](VehicleCommandsApi.md#CloseChargePort) | **Post** /api/1/vehicles/{vehicle_id}/command/charge_port_door_close | Close Charge Port
[**DeactivateSpeedLimit**](VehicleCommandsApi.md#DeactivateSpeedLimit) | **Post** /api/1/vehicles/{vehicle_id}/command/speed_limit_deactivate | Deactivate Speed Limit
[**FlashLights**](VehicleCommandsApi.md#FlashLights) | **Post** /api/1/vehicles/{vehicle_id}/command/flash_lights | Flash Lights
[**HonkHorn**](VehicleCommandsApi.md#HonkHorn) | **Post** /api/1/vehicles/{vehicle_id}/command/honk_horn | Honk Horn
[**LockDoors**](VehicleCommandsApi.md#LockDoors) | **Post** /api/1/vehicles/{vehicle_id}/command/door_lock | Lock Doors
[**NavigationRequest**](VehicleCommandsApi.md#NavigationRequest) | **Post** /api/1/vehicles/{vehicle_id}/command/navigation_request | Send Navigation Request
[**OpenChargePort**](VehicleCommandsApi.md#OpenChargePort) | **Post** /api/1/vehicles/{vehicle_id}/command/charge_port_door_open | Open Charge Port
[**OpenSunroof**](VehicleCommandsApi.md#OpenSunroof) | **Post** /api/1/vehicles/{vehicle_id}/command/sun_roof_control | Move Pano Roof
[**OpenTrunk**](VehicleCommandsApi.md#OpenTrunk) | **Post** /api/1/vehicles/{vehicle_id}/command/actuate_trunk | Open Trunk/Frunk
[**RemoteSeatHeaterRequest**](VehicleCommandsApi.md#RemoteSeatHeaterRequest) | **Post** /api/1/vehicles/{vehicle_id}/command/remote_seat_heater_request | Set Seat Heater Level
[**RemoteStart**](VehicleCommandsApi.md#RemoteStart) | **Post** /api/1/vehicles/{vehicle_id}/command/remote_start_drive | Remote Start
[**RemoteSteeringWheelHeaterRequest**](VehicleCommandsApi.md#RemoteSteeringWheelHeaterRequest) | **Post** /api/1/vehicles/{vehicle_id}/command/remote_steering_wheel_heater_request | Toggle Steering Wheel Heater
[**ResetValetPin**](VehicleCommandsApi.md#ResetValetPin) | **Post** /api/1/vehicles/{vehicle_id}/command/reset_valet_pin | Reset Valet PIN
[**SendStandardChargeLimit**](VehicleCommandsApi.md#SendStandardChargeLimit) | **Post** /api/1/vehicles/{vehicle_id}/command/charge_standard | Set Charge Limit to Standard
[**SetChargeLimit**](VehicleCommandsApi.md#SetChargeLimit) | **Post** /api/1/vehicles/{vehicle_id}/command/set_charge_limit | Set Charge Limit
[**SetMaxChargeLimit**](VehicleCommandsApi.md#SetMaxChargeLimit) | **Post** /api/1/vehicles/{vehicle_id}/command/charge_max_range | Set Charge Limit to Max Range
[**SetMaxDefrost**](VehicleCommandsApi.md#SetMaxDefrost) | **Post** /api/1/vehicles/{vehicle_id}/command/set_preconditioning_max | Set Max Defrost
[**SetSentryMode**](VehicleCommandsApi.md#SetSentryMode) | **Post** /api/1/vehicles/{vehicle_id}/command/set_sentry_mode | Toggle Sentry Mode
[**SetSpeedLimit**](VehicleCommandsApi.md#SetSpeedLimit) | **Post** /api/1/vehicles/{vehicle_id}/command/speed_limit_set_limit | Set Speed Limit
[**SetTemperatures**](VehicleCommandsApi.md#SetTemperatures) | **Post** /api/1/vehicles/{vehicle_id}/command/set_temps | Set Temperature
[**SharetoVehicle**](VehicleCommandsApi.md#SharetoVehicle) | **Post** /api/1/vehicles/{vehicle_id}/command/share | Share data to Vehicle
[**StartCharge**](VehicleCommandsApi.md#StartCharge) | **Post** /api/1/vehicles/{vehicle_id}/command/charge_start | Start Charging
[**StartHVAC**](VehicleCommandsApi.md#StartHVAC) | **Post** /api/1/vehicles/{vehicle_id}/command/auto_conditioning_start | Start HVAC System
[**StartSoftwareUpdate**](VehicleCommandsApi.md#StartSoftwareUpdate) | **Post** /api/1/vehicles/{vehicle_id}/command/schedule_software_update | Start Software Update
[**StopCharge**](VehicleCommandsApi.md#StopCharge) | **Post** /api/1/vehicles/{vehicle_id}/command/charge_stop | Stop Charging
[**StopHVAC**](VehicleCommandsApi.md#StopHVAC) | **Post** /api/1/vehicles/{vehicle_id}/command/auto_conditioning_stop | Stop HVAC System
[**ToggleValetMode**](VehicleCommandsApi.md#ToggleValetMode) | **Post** /api/1/vehicles/{vehicle_id}/command/set_valet_mode | Set Valet Mode
[**TriggerHomelink**](VehicleCommandsApi.md#TriggerHomelink) | **Post** /api/1/vehicles/{vehicle_id}/command/trigger_homelink | Trigger Homelink
[**UnlockDoors**](VehicleCommandsApi.md#UnlockDoors) | **Post** /api/1/vehicles/{vehicle_id}/command/door_unlock | Unlock Doors
[**WakeUpVehicle**](VehicleCommandsApi.md#WakeUpVehicle) | **Post** /api/1/vehicles/{vehicle_id}/wake_up | Wake Up Car
[**WakeUpVehicleCommand**](VehicleCommandsApi.md#WakeUpVehicleCommand) | **Post** /api/1/vehicles/{vehicle_id}/command/wake_up | Wake Up Car
[**WindowControl**](VehicleCommandsApi.md#WindowControl) | **Post** /api/1/vehicles/{vehicle_id}/command/window_control | Window Control


# **ActivateSpeedLimit**
> CommandResponse ActivateSpeedLimit(ctx, vehicleId, body)
Activate Speed Limit

Activates Speed Limit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**SpeedLimitRequest**](SpeedLimitRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelSoftwareUpdate**
> CommandResponse CancelSoftwareUpdate(ctx, vehicleId)
Cancel Software Update

Cancel Software Update

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

# **ClearSpeedLimitPin**
> CommandResponse ClearSpeedLimitPin(ctx, vehicleId, body)
Clear Speed Limit Pin

Clears Speed Limit Pin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**SpeedLimitRequest**](SpeedLimitRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloseChargePort**
> CommandResponse CloseChargePort(ctx, vehicleId)
Close Charge Port

Closes the charge port.

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

# **DeactivateSpeedLimit**
> CommandResponse DeactivateSpeedLimit(ctx, vehicleId, body)
Deactivate Speed Limit

Deactivates Speed Limit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**SpeedLimitRequest**](SpeedLimitRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FlashLights**
> CommandResponse FlashLights(ctx, vehicleId)
Flash Lights

Flash the lights once.

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

# **HonkHorn**
> CommandResponse HonkHorn(ctx, vehicleId)
Honk Horn

Honk the horn once.

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

# **LockDoors**
> CommandResponse LockDoors(ctx, vehicleId)
Lock Doors

Lock the car's doors.

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

# **NavigationRequest**
> CommandResponse NavigationRequest(ctx, vehicleId, body)
Send Navigation Request

Sends Navigation Request to Vehicle

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**NavigationRequestRequest**](NavigationRequestRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenChargePort**
> CommandResponse OpenChargePort(ctx, vehicleId)
Open Charge Port

Opens and unlocks the charge port.

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

# **OpenSunroof**
> CommandResponse OpenSunroof(ctx, vehicleId, body)
Move Pano Roof

Controls the car's panoramic roof, if installed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**MovePanoRoofRequest**](MovePanoRoofRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenTrunk**
> CommandResponse OpenTrunk(ctx, vehicleId, body)
Open Trunk/Frunk

Open the trunk or frunk. Currently inoperable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**OpenTrunkRequest**](OpenTrunkRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoteSeatHeaterRequest**
> CommandResponse RemoteSeatHeaterRequest(ctx, vehicleId, body)
Set Seat Heater Level

Set the heating level of a seat heater

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**RemoteSeatHeaterRequest**](RemoteSeatHeaterRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoteStart**
> CommandResponse RemoteStart(ctx, vehicleId, body)
Remote Start

Start the car for keyless driving. Must start driving within 2 minutes of issuing this request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**RemoteStartRequest**](RemoteStartRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoteSteeringWheelHeaterRequest**
> CommandResponse RemoteSteeringWheelHeaterRequest(ctx, vehicleId, body)
Toggle Steering Wheel Heater

Toggle the steering wheel heater

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**RemoteSteeringWheelHeaterRequest**](RemoteSteeringWheelHeaterRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetValetPin**
> CommandResponse ResetValetPin(ctx, vehicleId)
Reset Valet PIN

Resets the PIN set for valet mode, if set.

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

# **SendStandardChargeLimit**
> CommandResponse SendStandardChargeLimit(ctx, vehicleId)
Set Charge Limit to Standard

Set the charge mode to standard (90% under the new percentage system introduced in 4.5).

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

# **SetChargeLimit**
> CommandResponse SetChargeLimit(ctx, vehicleId, body)
Set Charge Limit

Set the charge limit to a custom percentage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**SetChargeLimitRequest**](SetChargeLimitRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetMaxChargeLimit**
> CommandResponse SetMaxChargeLimit(ctx, vehicleId)
Set Charge Limit to Max Range

Set the charge mode to max range (100% under the new percentage system introduced in 4.5). Use sparingly!

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

# **SetMaxDefrost**
> CommandResponse SetMaxDefrost(ctx, vehicleId, body)
Set Max Defrost

Set Max Defrost

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**MaxDefrostRequest**](MaxDefrostRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSentryMode**
> CommandResponse SetSentryMode(ctx, vehicleId, body)
Toggle Sentry Mode

Toggle Sentry Mode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**SentryModeRequest**](SentryModeRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSpeedLimit**
> CommandResponse SetSpeedLimit(ctx, vehicleId, body)
Set Speed Limit

Sets Speed Limit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**SetSpeedLimitRequest**](SetSpeedLimitRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetTemperatures**
> CommandResponse SetTemperatures(ctx, vehicleId, body)
Set Temperature

Set the temperature target for the HVAC system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**SetTempsRequest**](SetTempsRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SharetoVehicle**
> CommandResponse SharetoVehicle(ctx, vehicleId, body)
Share data to Vehicle

Sends Data to Vehicle (v10 only)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**ShareRequest**](ShareRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartCharge**
> CommandResponse StartCharge(ctx, vehicleId)
Start Charging

Start charging. Must be plugged in, have power available, and not have reached your charge limit.

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

# **StartHVAC**
> CommandResponse StartHVAC(ctx, vehicleId)
Start HVAC System

Start the climate control system. Will cool or heat automatically, depending on set temperature.

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

# **StartSoftwareUpdate**
> CommandResponse StartSoftwareUpdate(ctx, vehicleId)
Start Software Update

Start Software Update

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

# **StopCharge**
> CommandResponse StopCharge(ctx, vehicleId)
Stop Charging

Stop charging. Must already be charging.

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

# **StopHVAC**
> CommandResponse StopHVAC(ctx, vehicleId)
Stop HVAC System

Stop the climate control system.

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

# **ToggleValetMode**
> CommandResponse ToggleValetMode(ctx, vehicleId, body)
Set Valet Mode

Sets valet mode on or off with a PIN to disable it from within the car. Reuses last PIN from previous valet session. Valet Mode limits the car's top speed to 70MPH and 80kW of acceleration power. It also disables Homelink, Bluetooth and Wifi settings, and the ability to disable mobile access to the car. It also hides your favorites, home, and work locations in navigation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**SetValetRequest**](SetValetRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerHomelink**
> CommandResponse TriggerHomelink(ctx, vehicleId, body)
Trigger Homelink

Trigger Homelink

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**TriggerHomelinkRequest**](TriggerHomelinkRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnlockDoors**
> CommandResponse UnlockDoors(ctx, vehicleId)
Unlock Doors

Unlock the car's doors.

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

# **WakeUpVehicle**
> WakeUpResponse WakeUpVehicle(ctx, vehicleId)
Wake Up Car

Wakes up the car from the sleep state. Necessary to get some data from the car.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 

### Return type

[**WakeUpResponse**](WakeUpResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WakeUpVehicleCommand**
> CommandResponse WakeUpVehicleCommand(ctx, vehicleId)
Wake Up Car

Wakes up the car from the sleep state. Necessary to get some data from the car.

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

# **WindowControl**
> CommandResponse WindowControl(ctx, vehicleId, body)
Window Control

Window Control

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vehicleId** | **string**| The id of the Vehicle. | 
  **body** | [**WindowControlRequest**](WindowControlRequest.md)|  | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

