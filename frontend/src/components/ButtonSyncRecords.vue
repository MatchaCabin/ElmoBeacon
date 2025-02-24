<script setup lang="ts">
import {ElLoading} from "element-plus";
import {SyncRecordsCN, SyncRecordsOS} from "../../wailsjs/go/handler/App";
import {useI18n} from "vue-i18n";
import {useUserStore} from "../store/userStore.ts";
import {usePoolStore} from "../store/poolStore.ts";
import {service} from "../../wailsjs/go/models.ts";
import {NotifyError, NotifySuccess} from "../utils/notify.ts";

const {t} = useI18n()
const userStore = useUserStore()
const poolStore = usePoolStore()

const handleSyncResult = (res: service.SyncResult) => {
  NotifySuccess(
      t('sync.result.success.title', {server: t(`server.${res.Server}`),uid:res.Uid}),
      res.DiffList ? res.DiffList.map(diff => {
        return t('sync.result.success.changed', {poolType: t(`gacha.type.${diff.PoolType}`), count: diff.Count})
      }).join("<br/>") : t('sync.result.success.unchanged')
  )
}

const syncRecords = async () => {
  const loading = ElLoading.service({
    lock: true,
    text: t('sync.loading'),
    background: 'rgba(0, 0, 0, 0.7)',
  })

  await SyncRecordsCN().then(handleSyncResult).catch(err => NotifyError(t('sync.result.error.cn'), err))
  await SyncRecordsOS().then(handleSyncResult).catch(err => NotifyError(t('sync.result.error.os'), err))

  const lastUserId = userStore.userId
  await userStore.updateUserList()
  if (lastUserId === userStore.userId) {
    await poolStore.updatePoolInfo()
  }

  loading.close()
}
</script>

<template>
  <el-tooltip :content="$t('sync.button.tip')" placement="top">
    <el-button type="primary" @click="syncRecords">{{ $t('sync.button.title') }}</el-button>
  </el-tooltip>
</template>