<script setup lang="ts">
import {onMounted, ref} from "vue";
import HomeView from "./views/HomeView.vue";
import DataView from "./views/DataView.vue";
import SettingView from "./views/SettingView.vue";
import {Histogram, HomeFilled, Setting} from "@element-plus/icons-vue";
import {CheckUpdate, GetSettings, SetSetting, UpdateSelf} from "../wailsjs/go/handler/App";
import {ElLoading, ElMessageBox, ElNotification} from "element-plus";
import {useLangStore} from "./store/lang.ts";

const viewIndex = ref(0)
const langStore = useLangStore()

const initLang = async () => {
  await GetSettings().then(res => {
    if (res) {
      res.forEach(item => {
        if (item.key == "lang") {
          langStore.setLang(item.value)
        }
      })
    } else {
      let lang =''
      switch (navigator.language) {
        case "zh-CN":
          lang = 'zh-CN'
          break
        case "zh-TW":
          lang = 'zh-TW'
          break
        case "zh-HK":
          lang = 'zh-TW'
          break
        case "en-US":
          lang = 'en'
          break
        case "en-GB":
          lang = 'en'
          break
        case "ja-JP":
          lang = 'ja'
          break
        case "ko-KR":
          lang = 'ko'
          break
        default:
          lang = 'en'
      }
      langStore.setLang(lang)
      SetSetting('lang', lang).then(() => {
        langStore.setLang(lang)
      }).catch(err => {
        ElNotification({
          title: 'Error',
          message: err,
          type: 'error',
          position: 'bottom-right',
        })
      })
    }
  }).catch(err => {
    ElNotification({
      title: 'Error',
      message: err,
      type: 'error',
      position: 'bottom-right',
    })
  })
}

const checkUpdate = () => {
  CheckUpdate().then(res => {
    if (res) {
      ElMessageBox.confirm('There is a new version available, do you want to update?', 'Warning', {confirmButtonText: 'OK', cancelButtonText: 'Cancel', type: 'warning',}).then(() => {
        const loading = ElLoading.service({
          lock: true,
          text: 'Updating...',
          background: 'rgba(0, 0, 0, 0.7)',
        })
        UpdateSelf().catch(err => {
          ElNotification({
            title: 'Error',
            message: err,
            type: 'error',
            position: 'bottom-right',
          })
        }).finally(() => {
          loading.close()
        })
      })
    }
  })
}

onMounted(async () => {
  await initLang()
  checkUpdate()
})

</script>

<template>
  <div class="h-dvh w-dvw grow flex flex-row p-1 overflow-hidden">
    <div class="w-16 shrink-0 flex flex-col gap-4">
      <el-button class="!text-3xl !m-0 !h-12" :icon="HomeFilled" text @click="viewIndex=0"/>
      <el-button class="!text-3xl !m-0 !h-12" :icon="Histogram" text @click="viewIndex=1"/>
      <div class="mb-auto"></div>
      <el-button class="!text-3xl !m-0 !h-12" :icon="Setting" text @click="viewIndex=2"/>
    </div>
    <div class="grow p-1 overflow-hidden">
      <HomeView v-if="viewIndex===0"/>
      <DataView v-if="viewIndex===1"/>
      <SettingView v-if="viewIndex===2"/>
    </div>
  </div>
</template>