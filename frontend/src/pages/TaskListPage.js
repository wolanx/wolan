import React from 'react'
import { connect } from 'dva'
import { Card, List } from 'antd'
import { Link } from 'dva/router'
import MHeader from '../components/MHeader'
import MPanel from '../components/MPanel'

@connect(({ task }) => ({
    list: task.list
}), dispatch => ({
    getList: () => {
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
                                <Card title={`${item.name} - ${item.sid}`}
                                      extra={<Link to={`/task/${item.sid}`}>More</Link>}>
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
