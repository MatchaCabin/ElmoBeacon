<script setup lang="ts">
import {onMounted, ref} from "vue";
import {GetLatestVersion, GetVersion, UpdateSelf} from "../../wailsjs/go/handler/App";
import {ElLoading, ElMessageBox} from "element-plus";
import {NotifyError, NotifySuccess} from "../utils/notify.ts";
import {useI18n} from "vue-i18n";

const {t} = useI18n()
const version = ref('dev')

const checkUpdate = () => {
  GetLatestVersion().then(latestVersion => {
    if (latestVersion != version.value) {
      ElMessageBox.confirm(t('version.update.notify'), latestVersion, {confirmButtonText: t('version.update.confirm'), cancelButtonText: t('version.update.cancel'), type: 'info',}).then(() => {
        const loading = ElLoading.service({lock: true, text: `Update to ${latestVersion}...`, background: 'rgba(0, 0, 0, 0.7)'})
        UpdateSelf().catch(err => {
          NotifyError('Error', err)
        }).finally(() => {
          loading.close()
        })
      })
    } else {
      NotifySuccess(latestVersion, t('version.update.latest'))
    }
  })
}

onMounted(async () => {
  await GetVersion().then(res => {
    if (res) {
      version.value = res;
    }
  })
  if (version.value != "dev") {
    checkUpdate()
  }
})
</script>

<template>
  <el-tag class="ml-2 cursor-pointer" size="small" type="success" effect="light" @click="checkUpdate">
    {{ version }}
  </el-tag>
</template>