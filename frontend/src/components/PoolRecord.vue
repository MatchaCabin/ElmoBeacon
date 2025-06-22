<script setup lang="ts">
import { usePoolStore } from '../store/poolStore'


const poolStore = usePoolStore()

const props = defineProps({
  name: {type: String, required: true},
  count: {type: Number, required: true},
  timestamp: {type: Number, required: true},
  isMissing: {type: Boolean, required: true},
})


const getBgColor = () => {
  //80抽保底|70抽保底|50抽保底|无保底
  if (poolStore.poolType==1||poolStore.poolType==3||poolStore.poolType==6){
      if (props.count <= 58) {
        return 'bg-green-500'
      } else if (props.count < 66) {
        return 'bg-cyan-400'
      } else {
        return 'bg-red-600'
      }
  }else if(poolStore.poolType==4||poolStore.poolType==5||poolStore.poolType==7){
      if (props.count <= 50) {
        return 'bg-green-500'
      } else if (props.count < 58) {
        return 'bg-cyan-400'
      } else {
        return 'bg-red-600'
      }
  }else if(poolStore.poolType==5){
      if (props.count < 50) {
        return 'bg-green-500'
      } else{
        return 'bg-red-600'
      }
  }else{
    if (props.count<=100){
      return 'bg-green-500'
    }else if(props.count<=400){
      return 'bg-cyan-400'
    }else{
      return 'bg-red-600'
    }
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