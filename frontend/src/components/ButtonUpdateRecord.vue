<script setup lang="ts">
import {ElLoading, ElNotification} from "element-plus";
import {UpdateRecord} from "../../wailsjs/go/handler/App";
import {useI18n} from "vue-i18n";
import {useUserStore} from "../store/userStore.ts";

const {t} = useI18n()
const userStore =useUserStore()

const incrementalUpdate = () => {
  const loading = ElLoading.service({
    lock: true,
    text: t('record.update.incremental.loading'),
    background: 'rgba(0, 0, 0, 0.7)',
  })

  UpdateRecord().then(() => {
    userStore.updateUserList()
    ElNotification({
      title: 'Success',
      message: 'Records Updated',
      type: 'success',
      position: 'top-left',
    })
  }).catch(err => {
    ElNotification({
      title: 'Error',
      message: err,
      type: 'error',
      position: 'top-left',
    })
  }).finally(() => {
    loading.close()
  })
}
</script>

<template>
  <el-tooltip :content="$t('record.update.incremental.tip')" placement="top">
    <el-button type="primary" @click="incrementalUpdate">{{ $t('record.update.button') }}</el-button>
  </el-tooltip>
</template>