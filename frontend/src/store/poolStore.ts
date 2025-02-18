import {defineStore} from "pinia";
import {ref} from "vue";
import {handler} from "../../wailsjs/go/models.ts";
import {GetPoolInfo} from "../../wailsjs/go/handler/App";
import {ElNotification} from "element-plus";
import {useUserStore} from "./userStore.ts";
import PoolInfo = handler.PoolInfo;

export const usePoolStore = defineStore('pool', () => {
    const poolType = ref(1)
    const poolInfo = ref<PoolInfo>()

    const userStore = useUserStore()

    const updatePoolInfo = async () => {
        if (userStore.userId) {
            await GetPoolInfo(userStore.userId, poolType.value).then(res => {
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

    const init = async () => {
        await updatePoolInfo()
    }

    return {poolType, poolInfo, updatePoolInfo, init}
})