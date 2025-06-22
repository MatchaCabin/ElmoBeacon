import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useThemeStore = defineStore('theme', () => {
  // 尝试从localStorage获取保存的颜色
  const savedFrom = localStorage.getItem('themeFrom')
  const savedTo = localStorage.getItem('themeTo')
  
  // 默认渐变起始颜色 (对应原 blue-950)
  const gradientFrom = ref(savedFrom || '#0c1a32')
  // 默认渐变结束颜色 (对应原 stone-900)
  const gradientTo = ref(savedTo || '#1c1917')
  
  // 更新主题颜色的方法
  const updateTheme = (from: string, to: string) => {
    gradientFrom.value = from
    gradientTo.value = to
    
    // 更新CSS变量
    document.documentElement.style.setProperty('--theme-from', from)
    document.documentElement.style.setProperty('--theme-to', to)
    
    // 保存到localStorage
    localStorage.setItem('themeFrom', from)
    localStorage.setItem('themeTo', to)
  }

  // 初始化时应用存储的颜色
  if (savedFrom && savedTo) {
    document.documentElement.style.setProperty('--theme-from', savedFrom)
    document.documentElement.style.setProperty('--theme-to', savedTo)
  }

  return { gradientFrom, gradientTo, updateTheme }
})