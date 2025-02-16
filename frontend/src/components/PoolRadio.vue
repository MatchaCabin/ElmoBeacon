<script setup lang="ts">
import {usePoolStore} from "../store/poolStore.ts";

const props = defineProps({poolType: {type: Number, required: true}})
const poolStore = usePoolStore()

const handleClick = () => {
  poolStore.poolType = props.poolType
}
</script>

<template>
  <div :class="['h-24 p-2 rounded-md text-lg select-none cursor-pointer transition-colors ease-in-out duration-500 flex flex-row items-center',poolStore.poolType===poolType?'text-yellow-500 bg-white/10':'text-white hover:text-blue-500 hover:bg-white/10']" @click="handleClick">
    {{ $t(`gacha.type.${poolType}`) }}
  </div>
</template>