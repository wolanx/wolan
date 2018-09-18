import React, { Component } from 'react'
import axios from 'axios'

import { Table, Loading } from 'element-react'

export default class extends Component {
    constructor (props) {
        super(props)
        this.state = {
            loading: true,
            columns: [
                {
                    label: 'coinbase_name',
                    prop: 'coinbase_name',
                    width: 180
                },
                {
                    label: 'coinbase_price_view',
                    prop: 'spec.coinbase_price_view'
                },
                {
                    label: 'change_pct_view',
                    prop: 'spec.change_pct_view',
                },
                {
                    label: 'volume_24h_view',
                    prop: 'spec.volume_24h_view',
                },
                {
                    label: 'amount_total_view',
                    prop: 'spec.amount_total_view',
                },
            ],
            list: [],
        }
    }

    render () {
        return (
            <div>
                <Loading loading={this.state.loading}>
                    <Table
                        style={{width: '100%'}} stripe={true} border={true}
                        columns={this.state.columns} data={this.state.list}
                    />
                </Loading>
                {/*<ul>
                {this.state.list.map((v, k) => {
                return <li key={k}>{v.coinbase_name}</li>
                })}
                </ul>*/}
            </div>
        )
    }

    componentDidMount () {
        axios.get('http://localhost:777/api/market/coin').then((res) => {
            res = res.data
            // console.log(res.data)
            this.setState({
                loading: false,
                list: res.data,
            })
        }).catch((err) => {
            console.log(err.status)
        })
    }
}
