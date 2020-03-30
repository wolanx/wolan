import React from 'react'
import { connect } from 'dva'

@connect(state => ({
    list: state.task.list
}), dispatch => ({
    getList: () => {
        console.log('dispatch')
        return dispatch({ type: 'task/getList' })
    },
}))
class TaskListPage extends React.Component {
    render () {
        return (
            <ul>
                {this.props.list.map((v, k) => {
                    return <li key={k}>{v}</li>
                })}
            </ul>
        )
    }

    componentDidMount () {
        this.props.getList()
    }
}

export default TaskListPage
