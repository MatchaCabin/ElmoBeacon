<script setup lang="ts">
import {ref} from "vue";
import {SetSetting} from "../../wailsjs/go/handler/App";
import {ElNotification} from "element-plus";
import {useLangStore} from "../store/lang.ts";

const langStore = useLangStore()
// const lang = ref('')
const langOptions = ref([
  {label: '简体中文', value: 'zh-CN'},
  {label: '繁體中文', value: 'zh-TW'},
  {label: 'English', value: 'en'},
  {label: '日本語', value: 'ja'},
  {label: '한국어', value: 'ko'},
])
const handleLangChange = (value: string) => {
  SetSetting('lang', value).then(() => {
    langStore.setLang(value)
    ElNotification({
      title: 'Success',
      message: 'Language have been updated',
      type: 'success',
      position: 'bottom-right',
    })
  }).catch(err => {
    ElNotification({
      title: 'Error',
      message: err,
      type: 'error',
      position: 'bottom-right',
    })
  })
}

// onMounted(async () => {
//   await GetSettings().then(res => {
//     if (res){
//       res.forEach(item => {
//         switch (item.key) {
//           case 'lang':
//             lang.value = item.value;
//         }
//       })
//     }
//   })
// })

</script>

<template>
  <div class="w-96 flex flex-col">
    <div class="w-full flex flex-row justify-between items-center">
      <div class="w-32">{{ $t('setting.lang.title') }}</div>
      <div>
        <el-select class="!w-32" v-model="langStore.lang" @change="handleLangChange">
          <el-option v-for="item in langOptions" :key="item.value" :label="item.label" :value="item.value"/>
        </el-select>
      </div>
    </div>
  </div>
</template>