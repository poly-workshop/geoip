# V1Enterprise


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**continent** | [**V1Continent**](V1Continent.md) |  | [optional] [default to undefined]
**subdivisions** | [**Array&lt;V1EnterpriseSubdivision&gt;**](V1EnterpriseSubdivision.md) |  | [optional] [default to undefined]
**postal** | [**V1EnterprisePostal**](V1EnterprisePostal.md) |  | [optional] [default to undefined]
**representedCountry** | [**V1RepresentedCountry**](V1RepresentedCountry.md) |  | [optional] [default to undefined]
**country** | [**V1EnterpriseCountryRecord**](V1EnterpriseCountryRecord.md) |  | [optional] [default to undefined]
**registeredCountry** | [**V1CountryRecord**](V1CountryRecord.md) |  | [optional] [default to undefined]
**city** | [**V1EnterpriseCityRecord**](V1EnterpriseCityRecord.md) |  | [optional] [default to undefined]
**location** | [**Geoipv1Location**](Geoipv1Location.md) |  | [optional] [default to undefined]
**traits** | [**V1EnterpriseTraits**](V1EnterpriseTraits.md) |  | [optional] [default to undefined]

## Example

```typescript
import { V1Enterprise } from './api';

const instance: V1Enterprise = {
    continent,
    subdivisions,
    postal,
    representedCountry,
    country,
    registeredCountry,
    city,
    location,
    traits,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
