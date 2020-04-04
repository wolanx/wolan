import React from 'react'
import { connect } from 'dva'
import { Card, List } from 'antd'
import { Link } from 'dva/router'
import MHeader from '../components/MHeader'
import MPanel from '../components/MPanel'

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
            <>
                <MHeader title={'Task list'}/>
                <MPanel>
                    <List
                        grid={{ gutter: 16, column: 4 }}
                        dataSource={this.props.list}
                        renderItem={item => (
                            <List.Item>
                                <Card title={item.name} extra={<Link to={`/task/${item.name}`}>More</Link>}>
                                    {item.git.branch}
                                </Card>
                            </List.Item>
                        )}
                    />
                </MPanel>
            </>
        )
    }

    componentDidMount () {
        this.props.getList()
    }
}

export default TaskListPage
