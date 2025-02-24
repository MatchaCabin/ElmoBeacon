import {ElNotification} from "element-plus";

export function NotifySuccess(title:string,message:string){
    ElNotification({
        title: title,
        dangerouslyUseHTMLString: true,
        message: message,
        type: 'success',
        position: 'bottom-right',
    })
}

export function NotifyError(title:string,err:any){
    ElNotification({
        title: title,
        dangerouslyUseHTMLString: true,
        message: err,
        type: 'error',
        position: 'bottom-right',
    })
}