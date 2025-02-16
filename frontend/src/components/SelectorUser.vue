<script setup lang="ts">
import {useUserStore} from "../store/userStore.ts";

const userStore =useUserStore()

</script>

<template>
  <el-dropdown>
    <i-mdi-person class="text-xl text-white mr-2"/>
    <template #dropdown v-if="userStore.user&&userStore.userList">
      <el-dropdown-menu>
        <el-dropdown-item v-for="user in userStore.userList" @click="userStore.user=user">
          <span :class="user.id==userStore.user.id?'text-red-400':''">{{ $t(`server.${user.server}`)+` ${user.uid}` }}</span>
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>