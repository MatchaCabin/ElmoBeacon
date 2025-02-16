<script setup lang="ts">
import {usePoolStore} from "../store/poolStore.ts";
import PoolStatistic from "./PoolStatistic.vue";
import PoolRecord from "./PoolRecord.vue";

const poolStore = usePoolStore()

</script>

<template>
  <div class="w-full px-2 flex flex-col gap-4">
    <template v-if="poolStore.poolInfo">
      <div class="flex flex-row flex-wrap gap-2">
        <PoolStatistic class="text-cyan-400" :title="$t('gacha.statistic.totalCount')" :value="poolStore.poolInfo.totalCount"/>
        <PoolStatistic class="text-lime-400" :title="$t('gacha.statistic.pityCount')" :value="poolStore.poolInfo.storedCount"/>
        <PoolStatistic class="text-yellow-400" :title="$t('gacha.statistic.rank5Data')" :value="poolStore.poolInfo.rank5Count" :note="(poolStore.poolInfo.rank5Rate*100).toFixed(2)+'%'"/>
        <PoolStatistic class="text-purple-400" :title="$t('gacha.statistic.rank4Data')" :value="poolStore.poolInfo.rank4Count" :note="(poolStore.poolInfo.rank4Rate*100).toFixed(2)+'%'"/>
        <PoolStatistic class="text-blue-400" :title="$t('gacha.statistic.rank3Data')" :value="poolStore.poolInfo.rank3Count" :note="(poolStore.poolInfo.rank3Rate*100).toFixed(2)+'%'"/>
        <PoolStatistic class="text-red-500" :title="$t('gacha.statistic.rank5Avg')" :value="poolStore.poolInfo.rank5Avg"/>
        <PoolStatistic v-if="poolStore.poolType==3||poolStore.poolType==4" class="text-red-500" :title="$t('gacha.statistic.upRank5Avg')" :value="poolStore.poolInfo.rank5UpAvg"/>
        <PoolStatistic v-if="poolStore.poolType==3||poolStore.poolType==4" class="text-red-500" :title="$t('gacha.statistic.nonUpRate')" :note="Math.round(poolStore.poolInfo.missingRate*1000) /10+'%'"/>
      </div>

      <div class="h-12 w-full bg-blue-500/20 shadow-md select-none text-white text-center rounded-md flex flex-row gap-2 justify-center items-center">
        <span>{{ $t('gacha.records.title') }}</span>
        <el-tooltip effect="light" :content="$t('gacha.records.tip')" placement="top">
          <i-mdi-help class="text-sm"/>
          <template #content>
            <div class="text-lg max-w-96">{{$t('gacha.records.tip')}}</div>
          </template>
        </el-tooltip>
      </div>

      <el-scrollbar>
        <div class="w-full flex gap-x-2 gap-y-2 flex-wrap">
          <PoolRecord v-for="record in poolStore.poolInfo.recordList" :name="record.Name" :count="record.Count" :is-missing="record.IsMissing"/>
        </div>
      </el-scrollbar>
    </template>
  </div>
</template>