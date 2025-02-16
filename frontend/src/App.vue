<script setup lang="ts">
import AppHeader from "./components/AppHeader.vue";
import AppMain from "./components/AppMain.vue";
import AppFooter from "./components/AppFooter.vue";
</script>

<template>
  <div class="h-dvh w-dvw bg-gradient-to-b from-blue-950 to-stone-900 flex flex-col">
    <AppHeader/>
    <AppMain/>
    <AppFooter/>
  </div>
</template>