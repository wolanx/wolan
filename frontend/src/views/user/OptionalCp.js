import axios from 'axios'
import { Loading, Table } from 'element-react'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Link } from 'react-router-dom'

import WS from '../../components/ws/WS'

class Row extends Component {
    state = this.props.obj
    subName = `${this.props.obj.coinbase_id}:${this.props.obj.coinquote_id}:${this.props.obj.exchange_id}`

    render () {
        return (
            <tr style={{'lineHeight': '40px'}}>
                <td><Link to={`/coin/${this.state.coinbase_id}:${this.state.coinquote_id}:${this.state.exchange_id}`}>{this.state.name}</Link></td>
                <td>{this.state.exchange_name_view}</td>
                <td><span className={`num ${this.state.color} ${this.state.bgColor}`}>{this.state.spec.coinbase_price_view}</span></td>
                <td><span className={`num ${this.state.color} ${this.state.bgColor}`}>{this.state.spec.change_pct_view}%</span></td>
            </tr>
        )
    }

    componentDidMount () {
        WS.withMarketPriceSpec(this.subName, (obj) => {
            let t = this.state
            t.color = obj.change_pct >= 0 ? 'up' : 'down'
            t.bgColor = obj.price > t.spec.coinbase_price_view ? 'bg-up' : 'bg-down'

            t.spec.coinbase_price_view = obj.price
            t.spec.change_pct_view = obj.change_pct
            this.setState({t})
        })
    }

    componentWillUnmount () {
        WS.withMarketPriceSpec(this.subName, null)
    }
}

class App extends Component {
    constructor (props) {
        super(props)
        this.state = {
            loading: true,
            list: [],
            listSub: [],
        }
    }

    render () {
        return (
            <div>
                <Loading loading={this.state.loading}>
                    <table>
                        <thead>
                        <tr>
                            <th width="120">name</th>
                            <th width="100">exchange</th>
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
        document.title = '自选' + window.title_prefix

        axios.get('api/user/optional/coinpair').then((res) => {
            res = res.data
            const listSub = WS.subMarketPriceSpecSet(res.data.map(v => {
                return `${v.coinbase_id}:${v.coinquote_id}:${v.exchange_id}`
            }))
            this.setState({
                loading: false,
                list: res.data,
                listSub: listSub,
            })
        }).catch((err) => {
            console.log(err.status)
        })
    }

    componentWillUnmount () {
        WS.subMarketPriceSpecSet([])
    }
}

export default connect(
    (state) => {
        return {
            //
        }
    },
    (dispatch) => {
        return {
            //
        }
    }
)(App)