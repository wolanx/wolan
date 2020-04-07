import React from 'react'
import { connect } from 'dva'
import { Badge, Descriptions } from 'antd'
import MHeader from '../components/MHeader'
import MPanel from '../components/MPanel'

@connect()
export default class TaskInfoPage extends React.Component {

    taskName = this.props.match.params.name

    render () {
        return (
            <>
                <MHeader title={'Task details'}/>
                <MPanel title={'Info'}>
                    <Descriptions bordered>
                        <Descriptions.Item label="Name">{this.taskName}</Descriptions.Item>
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
                            Data disk type: MongoDB
                            <br/>
                            Database version: 3.4
                            <br/>
                            Package: dds.mongo.mid
                            <br/>
                            Storage space: 10 GB
                            <br/>
                            Replication factor: 3
                            <br/>
                            Region: East China 1<br/>
                        </Descriptions.Item>
                    </Descriptions>
                </MPanel>
            </>
        )
    }

}
