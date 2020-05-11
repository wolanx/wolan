import { graphql } from '../utils/request'
import { gql } from 'apollo-boost'

function getList () {
    return graphql({
        query: gql`
            {
                taskList{
                    sid
                    name
                    version
                    git{
                        url
                        branch
                    }
                }
            }
        `
    })
}

function getOne (sid) {
    console.log('getOne', sid)
    return graphql({
        query: gql`
            {
                taskGetBySid(sid: "${sid}"){
                    sid
                    name
                    version
                    git{
                        url
                        branch
                    }
                }
            }
        `
    })
}

export default {

    namespace: 'task',

    state: {
        list: [],
        info: undefined,
    },
    reducers: {
        setList (state, action) {
            return {
                ...state,
                list: action.payload,
            }
        },
        setInfo (state, action) {
            return {
                ...state,
                info: action.payload,
            }
        },
    },
    effects: {
        * getList (action, { call, put }) {
            // yield call(delay, 1000)
            const { data: res } = yield call(getList)
            console.log('data', res.data.taskList)
            yield put({ type: 'setList', payload: res.data.taskList })
        },
        * getOne (action, { call, put }) {
            const { data: res } = yield call(getOne, '01-demo')
            console.log('data', res.data.taskGetBySid)
            yield put({ type: 'setInfo', payload: res.data.taskGetBySid })
        },
    },

}
