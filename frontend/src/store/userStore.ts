import {defineStore} from "pinia";
import {ref} from "vue";
import {model} from "../../wailsjs/go/models.ts";
import {GetUserList} from "../../wailsjs/go/handler/App";
import User = model.User;
import {ElNotification} from "element-plus";

export const useUserStore = defineStore('user', () => {
    const user = ref<User>()
    const userList = ref<User[]>([])

    const updateUserList = () => {
        GetUserList().then((res) => {
            if (res) {
                userList.value = res;
                if (!user.value) {
                    user.value = userList.value[0];
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

    updateUserList()

    return {user, userList, updateUserList}
})