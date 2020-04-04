import React from 'react'
import { connect } from 'dva'
import { Card, Col, Row } from 'antd'
import { Link } from 'dva/router'
import MHeader from '../components/MHeader'

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
                <Row gutter={[16, 16]} style={{ margin: 0 }}>
                    {this.props.list.map((v, k) => {
                        return (
                            <Col span={6} key={k}>
                                <Card title={v.name} extra={<Link to={`/task/${v.name}`}>More</Link>}>
                                    {v.git.branch}
                                </Card>
                            </Col>
                        )
                    })}
                </Row>
            </>
        )
    }

    componentDidMount () {
        this.props.getList()
    }
}

export default TaskListPage
