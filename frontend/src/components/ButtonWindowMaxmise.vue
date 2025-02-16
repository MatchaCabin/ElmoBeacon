<script setup lang="ts">
import {WindowIsMaximised, WindowMaximise, WindowUnmaximise} from "../../wailsjs/runtime";

const handleWindowMaximise = async () => {
  await WindowIsMaximised().then(res => res ? WindowUnmaximise() : WindowMaximise())
}
</script>

<template>
  <i-mdi-crop-square class="hover:text-white/90 cursor-pointer" @click="handleWindowMaximise"/>
</template>