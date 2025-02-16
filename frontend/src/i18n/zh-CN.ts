export default {
    window: {
        title: '艾莫信标'
    },
    server:{
        cn:'国服',
        us:'美服',
        intl:'全球服',
        jp:'日服',
        kr:'韩服',
        tw:'亚服'
    },
    gacha:{
        type:{
            1:'常规采购',
            3:'定向采购',
            4:'军备提升',
            5:'初始采购',
            6:'自选采购·人形',
            7:'自选采购·军备',
            8:'神秘箱'
        },
        statistic:{
            totalCount: '记录总数',
            pityCount: '保底进度',
            rank5Data: '五星数据',
            rank4Data: '四星数据',
            rank3Data: '三星数据',
            rank5Avg: '五星平均抽数',
            upRank5Avg: 'Up五星平均抽数',
            nonUpRate: '五星歪率',
        },
        records:{
            title:'抽卡记录',
            tip:'记录名称读取自游戏本体，因此不管你选择什么语言，国服只会显示简体中文，国际服无法显示简体中文(如果你选择了简体中文，将使用繁体中文代替)'
        }
    },
    record:{
        update:{
            button:'更新记录',
            incremental: {
                button: '增量更新',
                tip: '从服务器逐步拉取抽卡记录，匹配到本地数据库最新记录后停止',
                loading:'增量更新中...',
            },
            full:{
                button: '全量更新',
                tip: '从服务器拉取全部抽卡记录，与本地数据库同步，通常用于纠错',
                loading:'全量更新中...',
            }
        },
    }
}