import {defineStore} from "pinia";
import {ref} from "vue";
import {useI18n} from "vue-i18n";
import {WindowSetTitle} from "../../wailsjs/runtime";
import {GetVersion} from "../../wailsjs/go/handler/App";

export const useLangStore = defineStore('lang', () => {
    const {t, locale} = useI18n()
    const lang = ref('')
    const version = ref('')
    GetVersion().then(res => {
        version.value = res
        console.log(res);
    })

    function setLang(value: string) {
        lang.value = value
        locale.value = value
        WindowSetTitle(`${t('window.title')} ${version.value}`)
    }

    return {lang, setLang}
})