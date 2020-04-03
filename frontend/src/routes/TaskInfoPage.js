import React from 'react'
import { connect } from 'dva'
import { Button, Descriptions, PageHeader } from 'antd'

@connect()
export default class TaskInfoPage extends React.Component {

    taskName = this.props.match.params.name

    render () {
        return (
            <div>
                <PageHeader
                    onBack={() => window.history.back()}
                    title={this.taskName}
                    subTitle="This is a subtitle"
                    extra={[
                        <Button key="3">Operation</Button>,
                        <Button key="2">Operation</Button>,
                        <Button key="1" type="primary">
                            Primary
                        </Button>,
                    ]}
                >
                    <Descriptions size="small" column={3}>
                        <Descriptions.Item label="Created">Lili Qu</Descriptions.Item>
                        <Descriptions.Item label="Association">
                            <a>421421</a>
                        </Descriptions.Item>
                        <Descriptions.Item label="Creation Time">2017-01-10</Descriptions.Item>
                        <Descriptions.Item label="Effective Time">2017-10-10</Descriptions.Item>
                        <Descriptions.Item label="Remarks">
                            Gonghu Road, Xihu District, Hangzhou, Zhejiang, China
                        </Descriptions.Item>
                    </Descriptions>
                </PageHeader>
                <div>task info {this.taskName}</div>
            </div>
        )
    }

}
