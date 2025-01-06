<script setup lang="ts">
import {onMounted, ref} from "vue";
import RecordPanel from "../components/RecordPanel.vue";
import {GetPoolInfoList, GetUserList, UpdateRecord} from "../../wailsjs/go/handler/App";
import {ElLoading, ElNotification} from "element-plus";
import {handler, model} from "../../wailsjs/go/models.ts";
import User = model.User;
import PoolInfo = handler.PoolInfo;

const userList = ref<User[]>([])
const userId = ref()
const poolInfoList = ref<PoolInfo[]>([])

const getUserList = async () => {
  await GetUserList().then(res => {
    if (res){
      userList.value = res
      if (!userId.value) {
        userId.value = userList.value[0].id
      }
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

const getPoolInfoList = async () => {
  if (userId.value) {
    await GetPoolInfoList(userId.value).then(res => {
      if (res) {
        poolInfoList.value = res
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
}

const updateRecord = async () => {
  const loading = ElLoading.service({
    lock: true,
    text: 'Updating Records...',
    background: 'rgba(0, 0, 0, 0.7)',
  })
  await UpdateRecord().then(() => {
    ElNotification({
      title: 'Success',
      message: 'Records have been updated',
      type: 'success',
      position: 'bottom-right',
    })
  }).catch(err => {
    ElNotification({
      title: 'Error',
      message: err,
      type: 'error',
      position: 'bottom-right',
    })
  })
  await getUserList()
  await getPoolInfoList()
  loading.close()
}

const handleUserChange = async () => {
  await getPoolInfoList()
}

onMounted(async () => {
  await getUserList()
  await getPoolInfoList()
})

</script>

<template>
  <div class="w-full h-full flex flex-col gap-1 overflow-hidden">
    <div class="h-10 flex flex-row justify-between items-center gap-2">
      <div>
        <el-button class="mr-2" type="primary" @click="updateRecord">{{ $t('btn.update.incremental.title') }}</el-button>
      </div>
      <div >
        <el-select class="!w-48" v-model="userId" placement="bottom" @change="handleUserChange">
          <el-option v-for="user in userList" :key="user.id" :label="`${user.server} ${user.uid}`" :value="user.id"/>
        </el-select>
      </div>
    </div>
    <div class="grow flex flex-row items-start gap-4 overflow-auto">
      <RecordPanel v-for="poolInfo in poolInfoList" :poolInfo="poolInfo"/>
    </div>
  </div>
</template>
