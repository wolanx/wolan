import axios from 'axios'
import { Loading, Table } from 'element-react'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Link } from 'react-router-dom'

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
            this.setState({
                loading: false,
                list: res.data,
            })
        }).catch((err) => {
            console.log(err.status)
        })
    }

    componentWillUnmount () {
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