export const translations = {
  zh: {
    title: 'GeoIP 查询工具',
    subtitle: '查询IP地址的地理位置信息',
    ipQuery: 'IP地址查询',
    inputPlaceholder: '输入IP地址 (例如: 8.8.8.8)',
    yourIp: '您的IP',
    useMyIp: '使用我的IP',
    search: '查询',
    searching: '查询中...',
    geoInfo: '地理位置信息',
    geoInfoDesc: 'IP地址 {{ip}} 的详细信息',
    ipAddress: 'IP地址',
    city: '城市',
    region: '地区',
    country: '国家',
    continent: '大洲',
    latitude: '纬度',
    longitude: '经度',
    timezone: '时区',
    isp: 'ISP',
    organization: '组织',
    mapLocation: '地图位置',
    coordinates: '坐标',
    viewOnGoogleMaps: '在Google Maps中查看',
    errors: {
      invalidIp: '请输入有效的IP地址',
      notFound: '未找到该IP的地理位置信息',
      invalidFormat: '无效的IP地址格式',
      tooManyRequests: '请求过多，请稍后重试',
      searchFailed: '查询失败，请稍后重试'
    },
    settings: {
      language: '语言',
      theme: '主题',
      light: '明亮',
      dark: '黑暗',
      system: '跟随系统'
    },
    footer: {
      dataSourcePrefix: '数据来源：基于',
      providerName: 'MaxMind GeoIP',
      dataSourceSuffix: '数据服务'
    }
  },
  en: {
    title: 'GeoIP Query Tool',
    subtitle: 'Query geographic location information for IP addresses',
    ipQuery: 'IP Address Query',
    inputPlaceholder: 'Enter IP address (e.g.: 8.8.8.8)',
    yourIp: 'Your IP',
    useMyIp: 'Use My IP',
    search: 'Search',
    searching: 'Searching...',
    geoInfo: 'Geographic Information',
    geoInfoDesc: 'Details for IP address {{ip}}',
    ipAddress: 'IP Address',
    city: 'City',
    region: 'Region',
    country: 'Country',
    continent: 'Continent',
    latitude: 'Latitude',
    longitude: 'Longitude',
    timezone: 'Timezone',
    isp: 'ISP',
    organization: 'Organization',
    mapLocation: 'Map Location',
    coordinates: 'Coordinates',
    viewOnGoogleMaps: 'View on Google Maps',
    errors: {
      invalidIp: 'Please enter a valid IP address',
      notFound: 'Geographic location information not found for this IP',
      invalidFormat: 'Invalid IP address format',
      tooManyRequests: 'Too many requests, please try again later',
      searchFailed: 'Search failed, please try again later'
    },
    settings: {
      language: 'Language',
      theme: 'Theme',
      light: 'Light',
      dark: 'Dark',
      system: 'System'
    },
    footer: {
      dataSourcePrefix: 'Data powered by ',
      providerName: 'MaxMind GeoIP services',
      dataSourceSuffix: ''
    }
  }
} as const

export type Language = keyof typeof translations
export type TranslationKey = keyof typeof translations['zh']