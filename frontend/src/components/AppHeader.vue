<script setup lang="ts">
import {WindowIsMaximised, WindowMaximise, WindowUnmaximise} from "../../wailsjs/runtime";
import ButtonVersion from "./ButtonVersion.vue";
import ButtonWindowMinimise from "./ButtonWindowMinimise.vue";
import ButtonWindowMaxmise from "./ButtonWindowMaxmise.vue";
import ButtonWindowClose from "./ButtonWindowClose.vue";
import SelectorLanguage from "./SelectorLanguage.vue";
import SelectorUser from "./SelectorUser.vue";

const handleWindowMaximise = async () => {
  await WindowIsMaximised().then(res => res ? WindowUnmaximise() : WindowMaximise())
}
</script>

<template>
  <div style="--wails-draggable:drag" class="h-20 shrink-0 px-8 text-lg text-stone-400 flex flex-row justify-between items-center gap-2" @dblclick.self="handleWindowMaximise">
    <div class="text-2xl font-bold select-none mr-auto bg-gradient-to-tl from-slate-800 via-violet-500 to-zinc-400 bg-clip-text text-transparent">
      <span>{{ $t('window.title') }}</span>
      <ButtonVersion/>
    </div>
    <SelectorUser/>
    <SelectorLanguage/>
    <ButtonWindowMinimise/>
    <ButtonWindowMaxmise/>
    <ButtonWindowClose/>
  </div>
</template>