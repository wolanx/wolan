import axios from 'axios'
import { Loading, Table } from 'element-react'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import TradingView from '../../components/tradingview/TradingView'

class CoinInfo extends Component {
    cbid = 0
    cqid = 0
    exid = 0

    state = {
        loading: true,
        info: {},
    }

    constructor (props) {
        super(props)

        let [cbid, cqid = 2392, exid = 1] = props.match.params.scope.split(':')
        this.cbid = cbid
        this.cqid = cqid
        this.exid = exid

        window.log('scope', cbid, cqid, exid)
    }

    render () {
        const info = this.state.info
        return this.state.loading ? (
            <div>Loading</div>
        ) : (
            <div>
                <h1>{info.coinbase_name}</h1>
                <TradingView cpname={`${info.coinbase_name}/${info.coinquote_name} ${info.exchange_name_view}`} cbid={this.cbid} cqid={this.cqid} exid={this.exid}/>
            </div>
        )
    }

    componentDidMount () {
        document.setTitle('单币')

        axios.get(`api/coin/${this.cbid}/info?coinquote_id=${this.cqid}&exchange_id=${this.exid}`).then(res => {
            res = res.data
            console.log(res.data)

            const info = res.data
            document.setTitle(`${info.coinbase_name}/${info.coinquote_name} ${info.exchange_name_view}`)
            this.setState({
                loading: false,
                info: info,
            })
        }, res => {
            res = res.response.data
            console.log(res.message)
        })
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
)(CoinInfo)