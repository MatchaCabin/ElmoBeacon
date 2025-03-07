export default {
    window: {
        title: 'ElmoBeacon'
    },
    server:{
        cn:'China',
        us:'America',
        intl:'Global',
        jp:'Japan',
        kr:'Korea',
        tw:'Asia'
    },
    gacha:{
        type:{
            1:'Standard Procurement',
            3:'Targeted Procurement',
            4:'Military Upgrade',
            5:'Beginner Procurement',
            6:'Custom Procurement - Dolls',
            7:'Custom Procurement - Weapon',
            8:'Mystery Box'
        },
        statistic:{
            totalCount: 'Total Counter',
            pityCount: 'Pity Counter',
            rank5Data: '5-star Data',
            rank4Data: '4-star Data',
            rank3Data: '3-star Data',
            rank5Avg: 'Avg Pulls per 5-star',
            upRank5Avg: 'Avg Pulls per Up 5-star',
            nonUpRate: 'NonUp 5-star Rate',
        },
        records:{
            title:'Pull Records',
            tip:'Record names are retrieved from the game client. Therefore, regardless of your language selection, the Chinese server will only display Simplified Chinese, while the international server cannot show Simplified Chinese (if you select Simplified Chinese, it will use Traditional Chinese as a substitute).'
        }
    },
    sync: {
        button: {
            title: 'Synchronize Records',
            tip: 'Pull records from the server and stops when it matches the latest record in the local database'
        },
        loading: 'Syncing...',
        result: {
            success: {
                title: '{server} {uid} Synchronization Success',
                changed: '{count} new records added for {poolType}',
                unchanged: 'No new records added'
            },
            error: {
                cn: 'Synchronization Error(CN)',
                os: 'Synchronization Error(OS)'
            }
        }
    },
    version: {
        update: {
            notify: 'There is a new version available, do you want to update?',
            latest: 'Already the latest version.',
            confirm: 'Yes',
            cancel: 'No'
        },
    }
}