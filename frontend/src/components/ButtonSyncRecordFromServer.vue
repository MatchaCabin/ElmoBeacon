<script setup lang="ts">
import {ElLoading, ElNotification} from "element-plus";
import {SyncRecordsFromServer} from "../../wailsjs/go/handler/App";
import {useI18n} from "vue-i18n";
import {useUserStore} from "../store/userStore.ts";
import {usePoolStore} from "../store/poolStore.ts";

const {t} = useI18n()
const userStore = useUserStore()
const poolStore = usePoolStore()

const syncRecordFromServer = () => {
  const loading = ElLoading.service({
    lock: true,
    text: "Sync Records...",
    background: 'rgba(0, 0, 0, 0.7)',
  })

  SyncRecordsFromServer().then(async res => {
    const lastUserId = userStore.userId
    await userStore.updateUserList()
    if (lastUserId === userStore.userId) {
      await poolStore.updatePoolInfo()
    }

    ElNotification({
      title: 'Success',
      dangerouslyUseHTMLString: true,
      message: res ? res.map(item => {
        return `${t(`server.${item.Server}`)} ${item.Uid} ${t(`gacha.type.${item.PoolType}`)} ${item.Count}`
      }).join("<br/>") : 'Nothing updated',
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
    <el-button type="primary" @click="syncRecordFromServer">{{ $t('record.update.button') }}</el-button>
  </el-tooltip>
</template>