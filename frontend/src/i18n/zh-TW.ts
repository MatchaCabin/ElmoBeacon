export default {
    window: {
        title: '艾莫信標'
    },
    server: {
        cn: '陸服',
        us: '美服',
        intl: '全球服',
        jp: '日服',
        kr: '韓服',
        tw: '亞服'
    },
    gacha: {
        type: {
            1: '常規採購',
            3: '定向採購',
            4: '軍備提升',
            5: '初始採購',
            6: '自選採購·人形',
            7: '自選採購·軍備',
            8: '神秘箱'
        },
        statistic: {
            totalCount: '記錄總數',
            pityCount: '保底進度',
            rank5Data: '五星數據',
            rank4Data: '四星數據',
            rank3Data: '三星數據',
            rank5Avg: '五星平均抽數',
            upRank5Avg: 'Up五星平均抽數',
            nonUpRate: '五星歪率',
        },
        records: {
            title: '抽卡記錄'
        }
    },
    sync: {
        button: {
            title: '同步記錄',
            tip: '從伺服器拉取抽卡記錄，匹配到本地數據庫最新記錄後停止'
        },
        loading: '正在同步記錄...',
        result: {
            success: {
                title: '{server} {uid} 同步成功',
                changed: '{poolType} 新增 {count} 條',
                unchanged: '无無新增數據'
            },
            error: {
                cn: '陸服同步出錯',
                os: '國際服同步出錯'
            }
        }
    },
    theme: {
        change: "更改主題",
        custom: "自訂主題",
        from: "起始顏色",
        to: "結束顏色",
        preview: "預覽效果",
        popular: "流行漸變",
        apply: "應用",
        cancel: "取消"
  }
}