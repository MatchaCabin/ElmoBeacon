import {defineStore} from "pinia";
import {ref} from "vue";
import {model} from "../../wailsjs/go/models.ts";
import {GetSetting, GetUserList, SetSetting} from "../../wailsjs/go/handler/App";
import {ElNotification} from "element-plus";
import User = model.User;

const getSettingUserId = async () => {
    let settingUserId = 0

    await GetSetting("lastUserId").then(res => {
        if (res) {
            settingUserId = parseInt(res)
        }
    }).catch(err => {
        ElNotification({
            title: 'Error',
            message: err,
            type: 'error',
            position: 'top-left',
        })
    })

    return settingUserId
}

export const useUserStore = defineStore('user', () => {
    const userId = ref<number>()
    const userList = ref<User[]>([])

    const updateUserId = async (newUserId: number) => {
        await SetSetting("lastUserId", newUserId.toString()).then(() => {
            userId.value = newUserId
        }).catch(err => {
            ElNotification({
                title: 'Error',
                message: err,
                type: 'error',
                position: 'top-left',
            })
        })
    }

    const updateUserList = async () => {
        await GetUserList().then((res) => {
            if (res) {
                userList.value = res;
                if (!userId.value) {
                    userId.value = userList.value[0].id;
                }
            }
        }).catch(err => {
            ElNotification({
                title: 'Error',
                message: err,
                type: 'error',
                position: 'top-left',
            })
        })
    }

    const init = async () => {
        userId.value = await getSettingUserId()
        await updateUserList()
    }

    return {userId, userList, updateUserId, updateUserList, init}
})