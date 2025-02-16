<script setup lang="ts">
import {useLangStore} from "../store/langStore.ts";
import {usePoolStore} from "../store/poolStore.ts";
import {useUserStore} from "../store/userStore.ts";

const langOptions = [
  {label: '简体中文', value: 'zh-CN'},
  {label: '繁體中文', value: 'zh-TW'},
  {label: 'English', value: 'en'},
  {label: '日本語', value: 'ja'},
  {label: '한국어', value: 'kr'},
]
const langStore = useLangStore()
const poolStore = usePoolStore()
const userStore = useUserStore()
const handleLangChange = async (newLang:string)=>{
  await langStore.updateLang(newLang)
  await poolStore.updatePoolInfo(userStore.user,poolStore.poolType)
}
</script>

<template>
  <el-dropdown>
    <i-mdi-translate class="text-lg text-white mr-2"/>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item v-for="lang in langOptions" @click="handleLangChange(lang.value)">
          <span :class="lang.value==langStore.lang?'text-red-400':''">{{ lang.label }}</span>
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>