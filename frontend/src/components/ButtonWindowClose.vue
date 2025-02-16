<script setup lang="ts">
import {Quit} from "../../wailsjs/runtime";
</script>

<template>
  <i-mdi-close class="hover:text-white/90 cursor-pointer" @click="Quit"/>
</template>
