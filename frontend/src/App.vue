<script setup lang="ts">
import AppHeader from "./components/AppHeader.vue";
import AppMain from "./components/AppMain.vue";
import AppFooter from "./components/AppFooter.vue";
import {useLangStore} from "./store/langStore.ts";
import {useUserStore} from "./store/userStore.ts";
import {usePoolStore} from "./store/poolStore.ts";
import {onBeforeMount, watch} from "vue";

const langStore = useLangStore()
const userStore = useUserStore()
const poolStore = usePoolStore()

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
  <div class="h-dvh w-dvw bg-gradient-to-b from-blue-950 to-stone-900 flex flex-col">
    <AppHeader/>
    <AppMain/>
    <AppFooter/>
  </div>
</template>