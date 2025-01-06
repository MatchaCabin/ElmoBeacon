<script setup lang="ts">
import {useI18n} from "vue-i18n";
import RecordLine from "./RecordLine.vue";
import {handler} from "../../wailsjs/go/models.ts";
import PoolInfo = handler.PoolInfo;

const {t} = useI18n()
const props = defineProps({poolInfo: {type: PoolInfo, required: true}})

const getTitleByPoolType = (): string => {
  switch (props.poolInfo.poolType) {
    case 1:
      return t("pool.type1.title")
    case 3:
      return t("pool.type3.title")
    case 4:
      return t("pool.type4.title")
    case 5:
      return t("pool.type5.title")
    case 6:
      return t("pool.type6.title")
    case 7:
      return t("pool.type7.title")
    case 8:
      return t("pool.type8.title")
    default:
      return t("pool.unknown.title")
  }
}

</script>

<template>
  <div class="w-[424px] shrink-0 flex flex-col px-1 pt-2 rounded-lg bg-blue-50">
    <div class="h-8 shrink-0 text-xl font-bold flex justify-between"><span>{{ getTitleByPoolType() }}</span><span>{{ poolInfo.totalCount }}</span></div>
    <div class="border-b-2 border-black"></div>
    <div class="h-8 shrink-0 flex justify-between text-lg text-red-500"><span>五星平均</span><span>{{
        Math.round(poolInfo.rank5Avg * 10) / 10
      }}{{ poolInfo.poolType === 3 || poolInfo.poolType === 4 ? ` [${Math.round(poolInfo.rank5UpAvg * 10) / 10},${Math.round(poolInfo.missingRate * 1000) / 10 + '%'}]` : '' }}</span></div>
    <div class="h-8 shrink-0 flex justify-between text-lg text-orange-500"><span>五星统计</span><span>{{ `${poolInfo.rank5Count} [${(poolInfo.rank5Rate * 100).toFixed(2)}%]` }}</span></div>
    <div class="h-8 shrink-0 flex justify-between text-lg text-purple-500"><span>四星统计</span><span>{{ `${poolInfo.rank4Count} [${(poolInfo.rank4Rate * 100).toFixed(2)}%]` }}</span></div>
    <div class="h-8 shrink-0 flex justify-between text-lg text-blue-500"><span>三星统计</span><span>{{ `${poolInfo.rank3Count} [${(poolInfo.rank3Rate * 100).toFixed(2)}%]` }}</span></div>

    <div class="border-b-2 border-black"></div>
    <div class="flex flex-col gap-1 py-1">
      <RecordLine icon="" name="已垫" :count="poolInfo.storedCount" :isMissing="false"/>
      <RecordLine v-for="record in poolInfo.recordList" :icon="record.Icon" :name="record.Name" :count="record.Count" :isMissing="record.IsMissing"/>
    </div>
  </div>
</template>