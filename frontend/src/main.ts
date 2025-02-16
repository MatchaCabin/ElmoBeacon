import {createApp} from 'vue'
import './style.css'
import 'element-plus/theme-chalk/el-notification.css'
import 'element-plus/theme-chalk/el-loading.css'
import 'element-plus/theme-chalk/el-message-box.css'
import 'element-plus/theme-chalk/el-button.css'
import App from './App.vue'
import {createPinia} from 'pinia'
import {createI18n} from "vue-i18n";
import zhCN from './i18n/zh-CN.ts';
import zhTW from './i18n/zh-TW.ts';
import en from './i18n/en';
import ja from './i18n/ja';
import kr from './i18n/kr.ts';

const pinia = createPinia()

const i18n = createI18n({
    composition: true,
    locale: 'en',
    fallbackLocale: 'en',
    messages: {
        "zh-CN": zhCN,
        "zh-TW": zhTW,
        "en": en,
        "ja": ja,
        "kr": kr,
    }
})

createApp(App).use(pinia).use(i18n).mount('#app')
