<script setup lang="ts">
import { useThemeStore } from '../store/themeStore'
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const themeStore = useThemeStore()
const showPicker = ref(false)
const tempColors = ref({
  from: themeStore.gradientFrom,
  to: themeStore.gradientTo
})

const openPicker = () => {
  showPicker.value = true
  // 初始化临时颜色为当前主题色
  tempColors.value = {
    from: themeStore.gradientFrom,
    to: themeStore.gradientTo
  }
}

const applyTheme = () => {
  themeStore.updateTheme(tempColors.value.from, tempColors.value.to)
  showPicker.value = false
}

// 常用颜色选项
const colorOptions = [
  ['#0c1a32', '#1c1917'], // 默认深蓝
  ['#2e1065', '#4a044e'], // 深紫
  ['#052e16', '#064e3b'], // 森林绿
  ['#7c2d12', '#9a3412'], // 日落橙
  ['#1e1b4b', '#3730a3'], // 午夜紫
  ['#164e63', '#0c4a6e'], // 海洋蓝
  ['#3f0f3f', '#5e0b5e'], // 宝石紫
  ['#1a2e05', '#2b4009'], // 橄榄绿
  ['#5a0e42', '#831843'], // 酒红
  ['#0f172a', '#1e293b']  // 深灰蓝
]
</script>

<template>
  <div class="relative">
    <!-- 调色板按钮 -->
    <button
      @click="openPicker"
      class="p-2 rounded-full hover:bg-gray-700 transition-colors flex items-center"
      :title="t('theme.change')"
    >
      <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01" />
      </svg>
    </button>

    <!-- 颜色选择器弹窗 -->
    <div v-if="showPicker" class="absolute right-0 mt-2 w-80 p-4 bg-gray-800 rounded-lg shadow-xl z-50 border border-gray-700">
      <h3 class="text-lg font-semibold text-gray-200 mb-4">{{ t('theme.custom') }}</h3>
      
      <div class="mb-4">
        <label class="block text-sm font-medium text-gray-300 mb-2">{{ t('theme.from') }}</label>
        <div class="flex gap-3">
          <input type="color" v-model="tempColors.from" class="w-14 h-14 cursor-pointer rounded-lg border border-gray-600">
          <div class="flex-1">
            <div class="w-full h-10 rounded-lg mb-2 border border-gray-600" 
                 :style="{ background: tempColors.from }"></div>
            <div class="text-xs text-gray-400 text-center">{{ tempColors.from }}</div>
          </div>
        </div>
      </div>
      
      <div class="mb-6">
        <label class="block text-sm font-medium text-gray-300 mb-2">{{ t('theme.to') }}</label>
        <div class="flex gap-3">
          <input type="color" v-model="tempColors.to" class="w-14 h-14 cursor-pointer rounded-lg border border-gray-600">
          <div class="flex-1">
            <div class="w-full h-10 rounded-lg mb-2 border border-gray-600" 
                 :style="{ background: tempColors.to }"></div>
            <div class="text-xs text-gray-400 text-center">{{ tempColors.to }}</div>
          </div>
        </div>
      </div>
      
      <div class="mb-6">
        <h4 class="text-sm font-medium text-gray-300 mb-3">{{ t('theme.preview') }}</h4>
        <div class="h-16 rounded-lg overflow-hidden border border-gray-700 shadow-lg"
             :style="{ background: `linear-gradient(to right, ${tempColors.from}, ${tempColors.to})` }">
        </div>
      </div>
      
      <div class="mb-6">
        <h4 class="text-sm font-medium text-gray-300 mb-2">{{ t('theme.popular') }}</h4>
        <div class="grid grid-cols-5 gap-2">
          <button 
            v-for="(colors, index) in colorOptions" 
            :key="index"
            @click="[tempColors.from = colors[0], tempColors.to = colors[1]]"
            class="h-8 rounded hover:scale-105 transition-transform"
            :style="{ background: `linear-gradient(to right, ${colors[0]}, ${colors[1]})` }"
            :title="colors[0] + ' → ' + colors[1]"
          ></button>
        </div>
      </div>
      
      <div class="flex justify-between pt-3 border-t border-gray-700">
        <button 
          @click="showPicker = false" 
          class="px-4 py-2 bg-gray-700 rounded hover:bg-gray-600 transition-colors text-sm"
        >
          {{ t('theme.cancel') }}
        </button>
        <button 
          @click="applyTheme" 
          class="px-4 py-2 bg-gradient-to-r from-blue-600 to-indigo-600 rounded hover:opacity-90 transition-opacity text-sm"
        >
          {{ t('theme.apply') }}
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
button:focus {
  outline: none;
  box-shadow: 0 0 0 3px rgba(66, 153, 225, 0.5);
}

input[type="color"] {
  -webkit-appearance: none;
  border: none;
}

input[type="color"]::-webkit-color-swatch-wrapper {
  padding: 0;
}

input[type="color"]::-webkit-color-swatch {
  border: none;
  border-radius: 8px;
}
</style>