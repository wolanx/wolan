import React from 'react'
import { connect } from 'dva'
import { Badge, Descriptions } from 'antd'
import MHeader from '../components/MHeader'
import MPanel from '../components/MPanel'

@connect(({ task }) => ({
    info: task.info
}), dispatch => ({
    doLoad: (sid) => {
        console.log('doLoad')
        return dispatch({ type: 'task/getOne', payload: { sid: sid } })
    }
}))
export default class TaskInfoPage extends React.Component {

    state = {
        sid: this.props.match.params.sid
    }

    static getDerivedStateFromProps (nextProps, prevState) {
        if (prevState.sid !== nextProps.match.params.sid) {
            return {
                sid: nextProps.match.params.sid
            }
        }

        return null
    }

    componentDidUpdate (prevProps, prevState, snapshot) {
        if (this.state.sid !== prevState.sid) {
            this.props.doLoad(this.state.sid)
        }
    }

    componentDidMount () {
        this.props.doLoad(this.state.sid)
    }

    render () {
        return (
            <>
                <MHeader title={'Task details'}/>
                <MPanel title={'Info'}>
                    <Descriptions bordered>
                        <Descriptions.Item label="Name">{this.state.sid}</Descriptions.Item>
                        <Descriptions.Item label="Billing Mode">Prepaid</Descriptions.Item>
                        <Descriptions.Item label="Automatic Renewal">YES</Descriptions.Item>
                        <Descriptions.Item label="Order time">2018-04-24 18:00:00</Descriptions.Item>
                        <Descriptions.Item label="Usage Time" span={2}>
                            2019-04-24 18:00:00
                        </Descriptions.Item>
                        <Descriptions.Item label="Status" span={3}>
                            <Badge status="processing" text="Running"/>
                        </Descriptions.Item>
                        <Descriptions.Item label="Negotiated Amount">$80.00</Descriptions.Item>
                        <Descriptions.Item label="Discount">$20.00</Descriptions.Item>
                        <Descriptions.Item label="Official Receipts">$60.00</Descriptions.Item>
                        <Descriptions.Item label="Config Info">
                            <pre>{JSON.stringify(this.props.info, null, '\t')}</pre>
                        </Descriptions.Item>
                    </Descriptions>
                </MPanel>
            </>
        )
    }
}
