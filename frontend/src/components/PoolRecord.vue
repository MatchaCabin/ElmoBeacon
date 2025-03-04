<script setup lang="ts">
const props = defineProps({
  name: {type: String, required: true},
  count: {type: Number, required: true},
  timestamp: {type: Number, required: true},
  isMissing: {type: Boolean, required: true},
})


const getBgColor = () => {
  if (props.count < 50) {
    return 'bg-cyan-400'
  } else if (props.count < 60) {
    return 'bg-lime-500'
  } else if (props.count < 65) {
    return 'bg-orange-500'
  } else {
    return 'bg-red-600'
  }
}
</script>

<template>
  <div :class="['w-52 h-8 relative shadow-xl rounded-md shrink-0 select-none flex justify-center items-center text-white',getBgColor()]">
    <el-tooltip effect="dark" :content="new Date(timestamp*1000).toLocaleString()" placement="top">
      <div>{{ `${name}「${count}」` }}</div>
    </el-tooltip>
    <div v-if="isMissing" class="absolute right-1 w-7 h-7 rounded-full border-gray-500 border-2 flex items-center justify-center">
      <span class="text-gray-500 -rotate-45 transform">歪</span>
    </div>
  </div>
</template>