export default {

    namespace: 'layout',

    state: {
        pathNow: '/',
    },

    subscriptions: {
        setup ({ dispatch, history }) {
            history.listen((location) => {
                // console.log('location is: %o', location)
                // console.log('重定向接收参数：%o', location.state)

                // 调用 effects 属性中的 query 方法，并将 location.state 作为参数传递
                // dispatch({
                //     type: 'query',
                //     payload: location.state,
                // })
                dispatch({
                    type: 'save',
                    payload: {
                        pathNow: location.pathname,
                    },
                })
            })
        },
    },

    reducers: {
        save (state, action) {
            // console.log('save', action.payload)
            return { ...state, ...action.payload }
        },
    },

    effects: {
        * fetch ({ payload }, { call, put }) {
            yield put({ type: 'save', payload })
        },
    },

}
