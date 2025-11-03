import { Select, SelectContent, SelectItem, SelectTrigger } from "@/components/ui/select"
import { Globe, Sun, Moon, Monitor } from "lucide-react"
import { useI18n } from "../i18n"
import { useTheme } from "./theme-provider"
import type { Theme } from "./theme-provider"
import type { Language } from "../i18n/translations"

export const SettingsPanel = () => {
  const { language, setLanguage, t } = useI18n()
  const { theme, setTheme } = useTheme()

  const handleLanguageChange = (value: string) => {
    setLanguage(value as Language)
  }

  const handleThemeChange = (value: string) => {
    setTheme(value as Theme)
  }

  const getLanguageLabel = (lang: Language) => {
    return lang === 'zh' ? '简体中文' : 'English'
  }

  const getThemeLabel = (themeValue: Theme) => {
    switch (themeValue) {
      case 'light':
        return t('settings.light')
      case 'dark':
        return t('settings.dark')
      case 'system':
        return t('settings.system')
      default:
        return themeValue
    }
  }

  const getThemeIcon = (themeValue: Theme) => {
    switch (themeValue) {
      case 'light':
        return <Sun className="w-4 h-4" />
      case 'dark':
        return <Moon className="w-4 h-4" />
      case 'system':
        return <Monitor className="w-4 h-4" />
      default:
        return null
    }
  }

  return (
  <div className="fixed top-4 right-4 z-50 flex items-center gap-6 rounded-2xl bg-white/80 px-5 py-3 shadow-lg backdrop-blur dark:bg-slate-900/80">
      <div className="flex items-center gap-3">
        <span className="text-xs font-medium uppercase tracking-wide text-slate-500 dark:text-slate-400 flex items-center gap-1">
          <Globe className="w-3 h-3" />
          {t('settings.language')}
        </span>
        <Select value={language} onValueChange={handleLanguageChange}>
          <SelectTrigger className="h-8 w-32 text-xs">
            <span className="flex items-center gap-2">
              <Globe className="w-3 h-3" />
              {getLanguageLabel(language)}
            </span>
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="zh">简体中文</SelectItem>
            <SelectItem value="en">English</SelectItem>
          </SelectContent>
        </Select>
      </div>

      <div className="flex items-center gap-3">
        <span className="text-xs font-medium uppercase tracking-wide text-slate-500 dark:text-slate-400 flex items-center gap-1">
          {getThemeIcon(theme)}
          {t('settings.theme')}
        </span>
        <Select value={theme} onValueChange={handleThemeChange}>
          <SelectTrigger className="h-8 w-36 text-xs">
            <span className="flex items-center gap-2">
              {getThemeIcon(theme)}
              {getThemeLabel(theme)}
            </span>
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="light">
              <div className="flex items-center gap-2">
                <Sun className="w-3 h-3" />
                {t('settings.light')}
              </div>
            </SelectItem>
            <SelectItem value="dark">
              <div className="flex items-center gap-2">
                <Moon className="w-3 h-3" />
                {t('settings.dark')}
              </div>
            </SelectItem>
            <SelectItem value="system">
              <div className="flex items-center gap-2">
                <Monitor className="w-3 h-3" />
                {t('settings.system')}
              </div>
            </SelectItem>
          </SelectContent>
        </Select>
      </div>
    </div>
  )
}