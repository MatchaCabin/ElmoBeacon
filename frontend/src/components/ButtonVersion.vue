<script setup lang="ts">
import {onMounted, ref} from "vue";
import {GetLatestVersion, GetVersion, UpdateSelf} from "../../wailsjs/go/handler/App";
import {ElLoading, ElMessageBox, ElNotification} from "element-plus";

const version = ref('dev')

const checkUpdate = () => {
  GetLatestVersion().then(latestVersion => {
    if (latestVersion!=version.value) {
      ElMessageBox.confirm('There is a new version available, do you want to update?', latestVersion, {confirmButtonText: 'OK', cancelButtonText: 'Cancel', type: 'info',}).then(() => {
        const loading = ElLoading.service({
          lock: true,
          text: `Update to ${latestVersion}...`,
          background: 'rgba(0, 0, 0, 0.7)',
        })
        UpdateSelf().catch(err => {
          ElNotification({
            title: 'Error',
            message: err,
            type: 'error',
            position: 'top-left',
          })
        }).finally(() => {
          loading.close()
        })
      })
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