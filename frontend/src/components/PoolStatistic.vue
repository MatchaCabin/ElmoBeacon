<script setup lang="ts">
defineProps({
  title: {type: String, required: true},
  value: {type: Number, required: false},
  note: {type: String, required: false},
})
</script>

<template>
  <div class="select-none w-64 h-24 p-2 rounded-lg shadow-xl flex flex-col items-center bg-blue-500/10">
    <div>{{ title }}</div>
    <div v-if="value!==undefined" class="text-2xl">{{ Math.round(value * 10) / 10 }}</div>
    <div v-if="note" :class="value!==undefined?'':'text-2xl'">{{ note }}</div>
  </div>
</template>