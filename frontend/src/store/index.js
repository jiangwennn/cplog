import { createStore } from "vuex";

export default createStore({
    // 定义所需要的状态
    state() {
        return {
            appName: 'cplog',
            colors: {
                brand: '#df1a29',
            },
            path: '/',
        }
    },
    // 同步修改state，都是方法，不能异步
    // 第一个参数state 第二个参数是需要修改的值
    mutations: {
        setName(state, payload) {
            state.appName = payload
        }
    },
    // 异步提交mutation
    // 第一个参数是store 第二个参数是修改的值
    actions: {
        
    },
    // 模块化
    modules: {},
})