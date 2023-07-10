import { createStore } from 'vuex';
import { LogPrint, EventsOn } from '_wailsjs/runtime/runtime.js'
import { GetContents } from '_wailsjs/go/app/Clipboard'

export default {
    namespaced: true,
    state: () => {
        return {
            list: {},
        }
    },
    mutations: {
        pushList(state, payload) {
            state.list[payload.Id] = payload
        },
        setList(state, payload) {
            LogPrint("哈哈哈哈哈哈哈哈哈哈")
            state.list = payload
        }
    },
    actions: {
        start(state) {
            // 初始化获取所有内容
            GetContents().then(result => {
                state.commit('setList', result)
            })
            // 监听剪切板新内容事件
            EventsOn('NewClip', content => {
                state.commit('pushList', content)
            })

        }
    },

}