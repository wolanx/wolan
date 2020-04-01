import { graphql } from '../utils/request'
import { gql } from 'apollo-boost'

function getList () {
    return graphql({
        query: gql`
            {
                listTask{
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
            const { data: res } = yield call(getList)
            console.log('data', res.data.listTask)
            yield put({ type: 'setList', payload: res.data.listTask })
        },
    },

}
