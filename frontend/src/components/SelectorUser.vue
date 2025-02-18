<script setup lang="ts">
import {useUserStore} from "../store/userStore.ts";

const userStore =useUserStore()

</script>

<template>
  <el-dropdown v-if="userStore.userId&&userStore.userList">
    <i-mdi-person class="text-xl text-white mr-2"/>
    <template #dropdown >
      <el-dropdown-menu>
        <el-dropdown-item v-for="user in userStore.userList" @click="userStore.updateUserId(user.id)">
          <span :class="user.id==userStore.userId?'text-red-400':''">{{ $t(`server.${user.server}`)+` ${user.uid}` }}</span>
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>