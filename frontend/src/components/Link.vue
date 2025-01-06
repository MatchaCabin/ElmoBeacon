<script setup lang="ts">
import {BrowserOpenURL} from "../../wailsjs/runtime";

defineProps({
  href: {
    type: String,
    required: true
  }
})
</script>

<template>
  <span class="text-blue-500 cursor-pointer" @click="BrowserOpenURL(href)"><slot></slot></span>
</template>