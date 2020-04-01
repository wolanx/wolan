import React from 'react'
import { connect } from 'dva'
import { Card, Col, Row } from 'antd'

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
            <Row gutter={16}>
                {this.props.list.map((v, k) => {
                    return (
                        <Col span={8}>
                            <Card title={v.name}>
                                {v.git.branch}
                            </Card>
                        </Col>
                    )
                })}
            </Row>
        )
    }

    componentDidMount () {
        this.props.getList()
    }
}

export default TaskListPage
