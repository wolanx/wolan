import { combineReducers } from 'redux'
import todos from './todos'
import visibilityFilter from './visibilityFilter'

export default combineReducers({
    msg: (state) => {
        return 'testsetset'
    },
    user: (state = {}, action) => {
        const count = state.count
        switch (action.type) {
            case 'user/set':
                return action.val
            case 'user/logout':
                return {}
            default:
                return state
        }
    },
    counter: function counter (state = {count: 0}, action) {
        const count = state.count
        switch (action.type) {
            case 'increase':
                return {count: count + 1}
            default:
                return state
        }
    },
    todos,
    visibilityFilter
})