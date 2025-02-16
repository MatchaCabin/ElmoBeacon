<script setup lang="ts">
import PoolMain from "./PoolMain.vue";
import PoolRadio from "./PoolRadio.vue";

const poolTypeList = [1, 3, 4, 5, 6, 7, 8]
</script>

<template>
  <div class="flex-1 overflow-y-auto flex flex-row px-4">
    <div class="w-60 shrink-0 flex flex-col gap-2">
      <PoolRadio v-for="poolType in poolTypeList" :poolType="poolType"/>
    </div>
    <PoolMain/>
  </div>
</template>