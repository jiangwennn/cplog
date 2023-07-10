import { createStore } from "vuex";
import { LogPrint } from '_wailsjs/runtime/runtime.js'
import History from '@/store/modules/history.js'


export default createStore({
    // 定义所需要的状态
    state() {
        return {
            appName: 'cplog',
            colors: {
                brand: '#df1a29',
            },
        }
    },
    // 同步修改state，都是方法，不能异步
    // 第一个参数state 第二个参数是需要修改的值
    mutations: {
        setName(state, payload) {
            LogPrint("少时诵诗书")
            state.appName = payload
        }
    },
    // 异步提交mutation
    // 第一个参数是store 第二个参数是修改的值
    actions: {
        
    },
    // 模块化
    modules: {
        history: History
    },
})