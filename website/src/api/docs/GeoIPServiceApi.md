# GeoIPServiceApi

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**geoIPServiceGetAnonymousIp**](#geoipservicegetanonymousip) | **GET** /v1/geoip/anonymous-ip/{ipAddress} | |
|[**geoIPServiceGetAsn**](#geoipservicegetasn) | **GET** /v1/geoip/asn/{ipAddress} | |
|[**geoIPServiceGetCity**](#geoipservicegetcity) | **GET** /v1/geoip/city/{ipAddress} | |
|[**geoIPServiceGetConnectionType**](#geoipservicegetconnectiontype) | **GET** /v1/geoip/connection-type/{ipAddress} | |
|[**geoIPServiceGetCountry**](#geoipservicegetcountry) | **GET** /v1/geoip/country/{ipAddress} | |
|[**geoIPServiceGetDomain**](#geoipservicegetdomain) | **GET** /v1/geoip/domain/{ipAddress} | |
|[**geoIPServiceGetEnterprise**](#geoipservicegetenterprise) | **GET** /v1/geoip/enterprise/{ipAddress} | |
|[**geoIPServiceGetIsp**](#geoipservicegetisp) | **GET** /v1/geoip/isp/{ipAddress} | |
|[**geoIPServiceGetMyIp**](#geoipservicegetmyip) | **GET** /v1/geoip/myip | |

# **geoIPServiceGetAnonymousIp**
> V1GetAnonymousIpResponse geoIPServiceGetAnonymousIp()


### Example

```typescript
import {
    GeoIPServiceApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new GeoIPServiceApi(configuration);

let ipAddress: string; // (default to undefined)

const { status, data } = await apiInstance.geoIPServiceGetAnonymousIp(
    ipAddress
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ipAddress** | [**string**] |  | defaults to undefined|


### Return type

**V1GetAnonymousIpResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | A successful response. |  -  |
|**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **geoIPServiceGetAsn**
> V1GetAsnResponse geoIPServiceGetAsn()


### Example

```typescript
import {
    GeoIPServiceApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new GeoIPServiceApi(configuration);

let ipAddress: string; // (default to undefined)

const { status, data } = await apiInstance.geoIPServiceGetAsn(
    ipAddress
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ipAddress** | [**string**] |  | defaults to undefined|


### Return type

**V1GetAsnResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | A successful response. |  -  |
|**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **geoIPServiceGetCity**
> V1GetCityResponse geoIPServiceGetCity()


### Example

```typescript
import {
    GeoIPServiceApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new GeoIPServiceApi(configuration);

let ipAddress: string; // (default to undefined)

const { status, data } = await apiInstance.geoIPServiceGetCity(
    ipAddress
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ipAddress** | [**string**] |  | defaults to undefined|


### Return type

**V1GetCityResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | A successful response. |  -  |
|**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **geoIPServiceGetConnectionType**
> V1GetConnectionTypeResponse geoIPServiceGetConnectionType()


### Example

```typescript
import {
    GeoIPServiceApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new GeoIPServiceApi(configuration);

let ipAddress: string; // (default to undefined)

const { status, data } = await apiInstance.geoIPServiceGetConnectionType(
    ipAddress
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ipAddress** | [**string**] |  | defaults to undefined|


### Return type

**V1GetConnectionTypeResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | A successful response. |  -  |
|**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **geoIPServiceGetCountry**
> V1GetCountryResponse geoIPServiceGetCountry()


### Example

```typescript
import {
    GeoIPServiceApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new GeoIPServiceApi(configuration);

let ipAddress: string; // (default to undefined)

const { status, data } = await apiInstance.geoIPServiceGetCountry(
    ipAddress
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ipAddress** | [**string**] |  | defaults to undefined|


### Return type

**V1GetCountryResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | A successful response. |  -  |
|**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **geoIPServiceGetDomain**
> V1GetDomainResponse geoIPServiceGetDomain()


### Example

```typescript
import {
    GeoIPServiceApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new GeoIPServiceApi(configuration);

let ipAddress: string; // (default to undefined)

const { status, data } = await apiInstance.geoIPServiceGetDomain(
    ipAddress
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ipAddress** | [**string**] |  | defaults to undefined|


### Return type

**V1GetDomainResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | A successful response. |  -  |
|**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **geoIPServiceGetEnterprise**
> V1GetEnterpriseResponse geoIPServiceGetEnterprise()


### Example

```typescript
import {
    GeoIPServiceApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new GeoIPServiceApi(configuration);

let ipAddress: string; // (default to undefined)

const { status, data } = await apiInstance.geoIPServiceGetEnterprise(
    ipAddress
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ipAddress** | [**string**] |  | defaults to undefined|


### Return type

**V1GetEnterpriseResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | A successful response. |  -  |
|**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **geoIPServiceGetIsp**
> V1GetIspResponse geoIPServiceGetIsp()


### Example

```typescript
import {
    GeoIPServiceApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new GeoIPServiceApi(configuration);

let ipAddress: string; // (default to undefined)

const { status, data } = await apiInstance.geoIPServiceGetIsp(
    ipAddress
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ipAddress** | [**string**] |  | defaults to undefined|


### Return type

**V1GetIspResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | A successful response. |  -  |
|**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **geoIPServiceGetMyIp**
> V1GetMyIpResponse geoIPServiceGetMyIp()


### Example

```typescript
import {
    GeoIPServiceApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new GeoIPServiceApi(configuration);

const { status, data } = await apiInstance.geoIPServiceGetMyIp();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**V1GetMyIpResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | A successful response. |  -  |
|**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

