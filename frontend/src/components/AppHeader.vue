<script setup lang="ts">
import {WindowIsMaximised, WindowMaximise, WindowUnmaximise} from "../../wailsjs/runtime";
import ButtonVersion from "./ButtonVersion.vue";
import ButtonWindowMinimise from "./ButtonWindowMinimise.vue";
import ButtonWindowMaxmise from "./ButtonWindowMaxmise.vue";
import ButtonWindowClose from "./ButtonWindowClose.vue";
import SelectorLanguage from "./SelectorLanguage.vue";
import SelectorUser from "./SelectorUser.vue";
import ButtonColorPalette from "./ButtonColorPalette.vue"; // 新增导入

const handleWindowMaximise = async () => {
  await WindowIsMaximised().then(res => res ? WindowUnmaximise() : WindowMaximise())
}
</script>

<template>
  <div style="--wails-draggable:drag" class="h-20 shrink-0 px-8 text-lg text-stone-400 flex flex-row justify-between items-center gap-2" @dblclick.self="handleWindowMaximise">
    <div class="text-2xl font-bold select-none mr-auto bg-gradient-to-l from-orange-500 to-slate-50 bg-clip-text text-transparent">
      <span>{{ $t('window.title') }}</span>
      <ButtonVersion/>
    </div>
    
    <SelectorUser/>
    <ButtonColorPalette/> <!-- 新增的调色板按钮 -->
    <SelectorLanguage/>
    
    <ButtonWindowMinimise/>
    <ButtonWindowMaxmise/>
    <ButtonWindowClose/>
  </div>
</template>