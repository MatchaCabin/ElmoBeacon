<script setup lang="ts">
import {ElLoading, ElNotification} from "element-plus";
import {UpdateRecord} from "../../wailsjs/go/handler/App";
import {useI18n} from "vue-i18n";
import {useUserStore} from "../store/userStore.ts";
import {usePoolStore} from "../store/poolStore.ts";

const {t} = useI18n()
const userStore =useUserStore()
const poolStore =usePoolStore()

const incrementalUpdate = () => {
  const loading = ElLoading.service({
    lock: true,
    text: t('record.update.incremental.loading'),
    background: 'rgba(0, 0, 0, 0.7)',
  })

  UpdateRecord().then(async () => {
    const lastUserId =userStore.userId
    await userStore.updateUserList()
    if (lastUserId===userStore.userId){
      await poolStore.updatePoolInfo()
    }

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