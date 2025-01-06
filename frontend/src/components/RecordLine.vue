<script setup lang="ts">
import {CloseBold} from "@element-plus/icons-vue";

defineProps({
  icon: {
    type: String,
    required: true
  },
  name: {
    type: String,
    required: true
  },
  count: {
    type: Number,
    required: true
  },
  isMissing: {
    type: Boolean,
    required: true
  }
})
</script>

<template>
  <div class="h-8 flex flex-row items-center gap-1 border-gray-200">
    <div class="h-8 w-8">
      <el-image class="h-8 w-8" v-if="icon!=''" :src="icon" fit="cover">
        <template #error>
          <el-image class="h-8 w-8" src="Avatar_Head_Unknown_Spine.png" fit="cover"/>
        </template>
      </el-image>
    </div>
    <div class="w-80 relative mr-auto ">
      <!-- Background div with a specific width and height -->
      <div :style="{ width: `${count*4}px` }" :class="['h-8 max-w-80 rounded',count<55?'bg-blue-400':'bg-red-500/80']"></div>

      <!-- Name displayed outside the div, left aligned -->
      <div class="absolute top-0 h-8  flex items-center">
        {{ name }}
      </div>
    </div>
    <!--    <div :class="['mr-auto bg-red-500/50 max-w-20',`w-[${count}px]`]">1</div>-->
    <el-icon class="!text-red-400 !text-3xl" v-if="isMissing">
      <CloseBold/>
    </el-icon>
    <div class="w-8 text-right">{{ count }}</div>
  </div>
</template>