import * as task from '../services/task'

export default {

    namespace: 'task',

    state: {
        list: [1, 2, 3],
    },
    reducers: {
        setList (state, action) {
            console.log('setList', action)
            return {
                ...state,
                list: action.payload,
            }
        },
    },
    effects: {
        * getList (action, { call, put }) {
            console.log('getList')
            // yield call(delay, 1000)
            const data = yield call(task.getList)
            console.log('data', data)
            yield put({ type: 'setList', payload: ['qwe', 13, 12, 3, 123, 123, 3123] })
        },
    },

}
