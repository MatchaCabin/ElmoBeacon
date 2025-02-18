import {defineStore} from "pinia";
import {ref} from "vue";
import {useI18n} from "vue-i18n";
import {GetSetting, SetSetting} from "../../wailsjs/go/handler/App";
import {ElNotification} from "element-plus";

const getSettingLang = async () => {
    let settingLang = ''

    await GetSetting("lang").then(res => {
        if (res && ['zh-CN', 'zh-TW', 'en', 'ja', 'kr'].includes(res)) {
            settingLang = res
        }
    }).catch(err => {
        ElNotification({
            title: 'Error',
            message: err,
            type: 'error',
            position: 'top-left',
        })
    })

    return settingLang
}

const getDefaultLang = () => {
    let defaultLang = 'en'
    const browserLang = navigator.language

    if (['zh-HK', 'zh-MO', 'zh-TW'].includes(browserLang)) {
        return 'zh-TW'
    } else if (browserLang.startsWith('zh')) {
        return 'zh-CN'
    } else if (browserLang.startsWith('ja')) {
        return 'ja'
    } else if (browserLang.startsWith('kr')) {
        return 'kr'
    }

    return defaultLang
}

export const useLangStore = defineStore('lang', () => {
    const {locale} = useI18n()
    const lang = ref('')

    const updateLang = async (newLang: string) => {
        if (lang.value !== newLang) {
            await SetSetting('lang', newLang).then(() => {
                lang.value = newLang
                locale.value = newLang
            }).catch(err => {
                ElNotification({
                    title: 'Error',
                    message: err,
                    type: 'error',
                    position: 'top-left',
                })
            })
        }
    }

    const init = async () => {
        const settingLang = await getSettingLang()
        if (settingLang) {
            lang.value = settingLang
            locale.value = settingLang
        } else {
            const defaultLang = getDefaultLang()
            await updateLang(defaultLang)
        }
    }

    return {lang, updateLang, init}
})