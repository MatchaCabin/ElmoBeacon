export default {
    window: {
        title: '艾莫信标'
    },
    server: {
        cn: '国服',
        us: '美服',
        intl: '全球服',
        jp: '日服',
        kr: '韩服',
        tw: '亚服'
    },
    gacha: {
        type: {
            1: '常规采购',
            3: '定向采购',
            4: '军备提升',
            5: '初始采购',
            6: '自选采购·人形',
            7: '自选采购·军备',
            8: '神秘箱'
        },
        statistic: {
            totalCount: '记录总数',
            pityCount: '保底进度',
            rank5Data: '五星数据',
            rank4Data: '四星数据',
            rank3Data: '三星数据',
            rank5Avg: '五星平均抽数',
            upRank5Avg: 'Up五星平均抽数',
            nonUpRate: '五星歪率',
        },
        records: {
            title: '抽卡记录',
            tip: '记录名称读取自游戏本体，因此不管你选择什么语言，国服只会显示简体中文，国际服无法显示简体中文(如果你选择了简体中文，将使用繁体中文代替)'
        }
    },
    sync: {
        button: {
            title: '同步记录',
            tip: '从服务器拉取抽卡记录，匹配到本地数据库最新记录后停止'
        },
        loading: '正在同步记录...',
        result: {
            success: {
                title: '{server} {uid} 同步成功',
                changed: '{poolType} 新增 {count} 条',
                unchanged: '无新增数据'
            },
            error: {
                cn: '国服同步出错',
                os: '国际服同步出错'
            }
        }
    },
    version: {
        update: {
            notify: '有新的版本可用，是否更新？',
            latest: '已是最新版本',
            confirm: '是',
            cancel: '否'
        },
    }
}