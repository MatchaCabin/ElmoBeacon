import {defineStore} from "pinia";
import {ref, watchEffect} from "vue";
import {handler, model} from "../../wailsjs/go/models.ts";
import {GetPoolInfo} from "../../wailsjs/go/handler/App";
import {ElNotification} from "element-plus";
import {useUserStore} from "./userStore.ts";
import PoolInfo = handler.PoolInfo;
import User = model.User;

export const usePoolStore = defineStore('pool', () => {
    const poolType = ref(1)
    const poolInfo = ref<PoolInfo>()

    const userStore = useUserStore()

    const updatePoolInfo = async (newUser: User | undefined, newPoolType: number) => {
        if (newUser) {
            await GetPoolInfo(newUser.id, newPoolType).then(res => {
                poolInfo.value = res
            }).catch(err => {
                ElNotification({
                    title: 'Error',
                    message: err,
                    type: 'error',
                    position: 'top-left',
                })
            })
        }
    }

    updatePoolInfo(userStore.user, poolType.value)

    watchEffect(async () => {
        await updatePoolInfo(userStore.user, poolType.value)
    })

    return {poolType, poolInfo, updatePoolInfo}
})