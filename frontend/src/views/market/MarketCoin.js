import axios from 'axios'
import { Button, Loading, Table } from 'element-react'
import React, { Component } from 'react'
import { Link } from 'react-router-dom'

class Row extends Component {
    state = this.props.obj
    subName = '' + this.props.obj.coinbase_id

    render () {
        return (
            <tr style={{'lineHeight': '40px'}}>
                <td><Link to={`/coin/${this.state.coinbase_id}`}>{this.state.coinbase_name}</Link></td>
                <td><span className={`num ${this.state.color} ${this.state.bgColor}`}>{this.state.spec.coinbase_price_view}</span></td>
                <td><span className={`num ${this.state.color} ${this.state.bgColor}`}>{this.state.spec.change_pct_view}%</span></td>
            </tr>
        )
    }
}

export default class extends Component {
    state = {
        loading: true,
        list: [],
        listSub: [],
    }

    render () {
        return (
            <div>
                <Loading loading={this.state.loading}>
                    <table>
                        <thead>
                        <tr>
                            <th width="120">name</th>
                            <th width="100">price</th>
                            <th>pct</th>
                        </tr>
                        </thead>
                        <tbody>
                        {this.state.list.map((v, k) => {
                            return (
                                <Row key={k} obj={v}/>
                            )
                        })}
                        </tbody>
                    </table>
                </Loading>
            </div>
        )
    }

    componentDidMount () {
        document.setTitle('单币')

        axios.get('api/market/coin').then(res => {
            res = res.data
            this.setState({
                loading: false,
                list: res.data,
            })
        }, res => {
            res = res.response.data
            console.log(res.message)
        })
    }
}
