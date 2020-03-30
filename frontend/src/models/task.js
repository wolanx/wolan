import * as task from '../services/task'

export default {

    namespace: 'task',

    state: {
        list: [],
    },
    reducers: {
        setList (state, action) {
            return {
                ...state,
                list: action.payload,
            }
        },
    },
    effects: {
        * getList (action, { call, put }) {
            // yield call(delay, 1000)
            const { data: res } = yield call(task.getList)
            console.log('data', res.data.listTask)
            yield put({ type: 'setList', payload: res.data.listTask })
        },
    },

}
