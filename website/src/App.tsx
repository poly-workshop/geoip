import React, { useState, useEffect } from "react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "./components/ui/badge"
import { Separator } from "./components/ui/separator"
import { MapPin, Globe, Building, Network } from "lucide-react"
import { GeoIPServiceApi, Configuration } from "@/api"
import type { V1GetMyIpResponse, V1GetCityResponse, V1Names } from "@/api"
import { useI18n } from "./i18n"
import { SettingsPanel } from "./components/settings-panel"

interface GeoLocationInfo {
  ip: string
  cityNames?: V1Names
  regionNames?: V1Names
  countryNames?: V1Names
  continentNames?: V1Names
  location?: {
    latitude: number
    longitude: number
  }
  timezone?: string
  isp?: string
  organization?: string
  asn?: string
}

function App() {
  const { t, language } = useI18n()
  const [ipAddress, setIpAddress] = useState("")
  const [geoInfo, setGeoInfo] = useState<GeoLocationInfo | null>(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState("")
  const [myIp, setMyIp] = useState("")

  // 创建 API client 实例
  const geoipApi = React.useMemo(() => new GeoIPServiceApi(new Configuration({
    basePath: window.location.origin + "/api"
  })), [])

  const getLocalizedName = React.useCallback((names?: V1Names) => {
    if (!names) return undefined
    if (language === 'zh' && names.zhCn) return names.zhCn
    if (language === 'en' && names.en) return names.en

    const fallbackOrder: Array<keyof V1Names> = ['en', 'zhCn', 'fr', 'es', 'de', 'ja', 'ptBr', 'ru']
    for (const key of fallbackOrder) {
      const value = names[key]
      if (value) {
        return value
      }
    }
    return undefined
  }, [language])

  const localizedNames = React.useMemo(() => {
    if (!geoInfo) return null
    return {
      city: getLocalizedName(geoInfo.cityNames),
      region: getLocalizedName(geoInfo.regionNames),
      country: getLocalizedName(geoInfo.countryNames),
      continent: getLocalizedName(geoInfo.continentNames)
    }
  }, [geoInfo, getLocalizedName])

  // 获取当前IP
  useEffect(() => {
    const fetchMyIp = async () => {
      try {
        const response = await geoipApi.geoIPServiceGetMyIp()
        const data = response.data as V1GetMyIpResponse
        if (data.ipAddress) {
          setMyIp(data.ipAddress)
          setIpAddress(data.ipAddress)
        }
      } catch (error) {
        console.error("Failed to fetch my IP:", error)
      }
    }

    fetchMyIp()
  }, [geoipApi])

  const handleSearch = async () => {
    if (!ipAddress.trim()) {
      setError(t('errors.invalidIp'))
      return
    }

    setLoading(true)
    setError("")
    setGeoInfo(null)

    try {
      const response = await geoipApi.geoIPServiceGetCity(ipAddress.trim())
      const data = response.data as V1GetCityResponse
      
      if (data.city) {
        const primarySubdivision = data.city.subdivisions && data.city.subdivisions.length > 0
          ? data.city.subdivisions[0]
          : undefined

        const info: GeoLocationInfo = {
          ip: ipAddress.trim(),
          cityNames: data.city.cityRecord?.names,
          regionNames: primarySubdivision?.names,
          countryNames: data.city.country?.names,
          continentNames: data.city.continent?.names,
          timezone: data.city.location?.timeZone,
          location: data.city.location ? {
            latitude: data.city.location.latitude || 0,
            longitude: data.city.location.longitude || 0
          } : undefined
        }
        setGeoInfo(info)
      } else {
        setError(t('errors.notFound'))
      }
    } catch (error: unknown) {
      console.error("API Error:", error)
      const errorResponse = error as { response?: { status?: number } }
      if (errorResponse.response?.status === 404) {
        setError(t('errors.notFound'))
      } else if (errorResponse.response?.status === 400) {
        setError(t('errors.invalidFormat'))
      } else if (errorResponse.response?.status === 429) {
        setError(t('errors.tooManyRequests'))
      } else {
        setError(t('errors.searchFailed'))
      }
    } finally {
      setLoading(false)
    }
  }

  const handleUseMyIp = () => {
    if (myIp) {
      setIpAddress(myIp)
    }
  }

  return (
    <div className="min-h-screen bg-slate-100 dark:bg-slate-900 p-4 flex flex-col">
      <SettingsPanel />
      <div className="container mx-auto max-w-4xl pt-16 flex-1">
        <div className="text-center mb-12 mt-8">
          <h1 className="text-4xl font-bold text-slate-900 dark:text-white mb-2">
            {t('title')}
          </h1>
          <p className="text-slate-600 dark:text-slate-300">
            {t('subtitle')}
          </p>
        </div>

        <Card className="mb-8">
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Globe className="w-5 h-5" />
              {t('ipQuery')}
            </CardTitle>
            <CardDescription>
              {t('subtitle')}
              {myIp && (
                <div className="mt-2 flex items-center gap-2">
                  <span className="text-sm">{t('yourIp')}: {myIp}</span>
                  <Button 
                    variant="outline" 
                    size="sm" 
                    onClick={handleUseMyIp}
                  >
                    {t('useMyIp')}
                  </Button>
                </div>
              )}
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="flex gap-4">
              <Input
                placeholder={t('inputPlaceholder')}
                value={ipAddress}
                onChange={(e) => setIpAddress(e.target.value)}
                onKeyPress={(e) => e.key === 'Enter' && handleSearch()}
                className="flex-1"
              />
              <Button 
                onClick={handleSearch} 
                disabled={loading}
                className="min-w-[100px]"
              >
                {loading ? t('searching') : t('search')}
              </Button>
            </div>
            {error && (
              <div className="mt-4 p-3 bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-md text-red-700 dark:text-red-300 text-sm">
                {error}
              </div>
            )}
          </CardContent>
        </Card>

        {geoInfo && (
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <MapPin className="w-5 h-5" />
                {t('geoInfo')}
              </CardTitle>
              <CardDescription>
                {t('geoInfoDesc', { ip: geoInfo.ip })}
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-6">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div className="space-y-3">
                  <div className="flex items-center justify-between">
                    <span className="text-sm font-medium text-slate-600 dark:text-slate-300">{t('ipAddress')}</span>
                    <Badge variant="secondary">{geoInfo.ip}</Badge>
                  </div>
                  
                  {localizedNames?.city && (
                    <div className="flex items-center justify-between">
                      <span className="text-sm font-medium text-slate-600 dark:text-slate-300">{t('city')}</span>
                      <span className="text-sm">{localizedNames.city}</span>
                    </div>
                  )}
                  
                  {localizedNames?.region && (
                    <div className="flex items-center justify-between">
                      <span className="text-sm font-medium text-slate-600 dark:text-slate-300">{t('region')}</span>
                      <span className="text-sm">{localizedNames.region}</span>
                    </div>
                  )}
                  
                  {localizedNames?.country && (
                    <div className="flex items-center justify-between">
                      <span className="text-sm font-medium text-slate-600 dark:text-slate-300">{t('country')}</span>
                      <span className="text-sm">{localizedNames.country}</span>
                    </div>
                  )}

                  {localizedNames?.continent && (
                    <div className="flex items-center justify-between">
                      <span className="text-sm font-medium text-slate-600 dark:text-slate-300">{t('continent')}</span>
                      <span className="text-sm">{localizedNames.continent}</span>
                    </div>
                  )}
                </div>

                <div className="space-y-3">
                  {geoInfo.location && (
                    <>
                      <div className="flex items-center justify-between">
                        <span className="text-sm font-medium text-slate-600 dark:text-slate-300">{t('latitude')}</span>
                        <span className="text-sm">{geoInfo.location.latitude.toFixed(4)}</span>
                      </div>
                      
                      <div className="flex items-center justify-between">
                        <span className="text-sm font-medium text-slate-600 dark:text-slate-300">{t('longitude')}</span>
                        <span className="text-sm">{geoInfo.location.longitude.toFixed(4)}</span>
                      </div>
                    </>
                  )}
                  
                  {geoInfo.timezone && (
                    <div className="flex items-center justify-between">
                      <span className="text-sm font-medium text-slate-600 dark:text-slate-300">{t('timezone')}</span>
                      <span className="text-sm">{geoInfo.timezone}</span>
                    </div>
                  )}
                  
                  {geoInfo.isp && (
                    <div className="flex items-center justify-between">
                      <span className="text-sm font-medium text-slate-600 dark:text-slate-300 flex items-center gap-1">
                        <Network className="w-4 h-4" />
                        {t('isp')}
                      </span>
                      <span className="text-sm">{geoInfo.isp}</span>
                    </div>
                  )}
                  
                  {geoInfo.organization && (
                    <div className="flex items-center justify-between">
                      <span className="text-sm font-medium text-slate-600 dark:text-slate-300 flex items-center gap-1">
                        <Building className="w-4 h-4" />
                        {t('organization')}
                      </span>
                      <span className="text-sm">{geoInfo.organization}</span>
                    </div>
                  )}
                </div>
              </div>

              {geoInfo.location && (
                <>
                  <Separator />
                  <div>
                    <h3 className="text-sm font-medium text-slate-600 dark:text-slate-300 mb-2">{t('mapLocation')}</h3>
                    <div className="bg-slate-50 dark:bg-slate-800 p-4 rounded-lg text-center">
                      <p className="text-sm text-slate-600 dark:text-slate-300">
                        {t('coordinates')}: {geoInfo.location.latitude.toFixed(4)}, {geoInfo.location.longitude.toFixed(4)}
                      </p>
                      <a 
                        href={`https://www.google.com/maps?q=${geoInfo.location.latitude},${geoInfo.location.longitude}`}
                        target="_blank"
                        rel="noopener noreferrer"
                        className="text-blue-600 dark:text-blue-400 hover:text-blue-800 dark:hover:text-blue-300 text-sm mt-2 inline-block"
                      >
                        {t('viewOnGoogleMaps')} →
                      </a>
                    </div>
                  </div>
                </>
              )}
            </CardContent>
          </Card>
        )}
      </div>
      <footer className="mt-12 text-center text-xs text-slate-500 dark:text-slate-400">
        {t('footer.dataSourcePrefix')}
        <a
          href="https://www.maxmind.com/"
          target="_blank"
          rel="noopener noreferrer"
          className="mx-1 text-blue-600 hover:text-blue-700 dark:text-blue-400 dark:hover:text-blue-300 underline-offset-4 hover:underline"
        >
          {t('footer.providerName')}
        </a>
        {t('footer.dataSourceSuffix')}
      </footer>
    </div>
  )
}

export default App
