<script setup lang="ts">
import AppHeader from "./components/AppHeader.vue";
import AppMain from "./components/AppMain.vue";
import AppFooter from "./components/AppFooter.vue";
import {useLangStore} from "./store/langStore.ts";
import {useUserStore} from "./store/userStore.ts";
import {usePoolStore} from "./store/poolStore.ts";
import {useThemeStore} from "./store/themeStore.ts"; // 新增导入
import {onBeforeMount, watch} from "vue";

const langStore = useLangStore()
const userStore = useUserStore()
const poolStore = usePoolStore()
useThemeStore() // 新增主题store

onBeforeMount(async () => {
  await langStore.init()
  await userStore.init()
  await poolStore.init()
})

watch(() => [langStore.lang, userStore.userId, poolStore.poolType], async () => {
  await poolStore.updatePoolInfo()
})
</script>

<template>
  <div class="h-dvh w-dvw flex flex-col" 
       :style="{ background: `linear-gradient(to bottom, var(--theme-from), var(--theme-to))` }">
    <AppHeader/>
    <AppMain/>
    <AppFooter/>
  </div>
</template>

<style>
/* 定义默认主题变量 */
:root {
  --theme-from: #0c1a32;
  --theme-to: #1c1917;
}
</style>