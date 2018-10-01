import axios from 'axios'
import { Loading, Table } from 'element-react'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Link } from 'react-router-dom'

class CoinInfo extends Component {
    state = {
        list: [],
    }

    render () {
        return (
            <div>
                <ul>
                    {this.state.list.map((v, k) => {
                        return (
                            <li key={k}>
                                <Link to={`/task/${v.name}`}>{v.name}</Link>
                            </li>
                        )
                    })}
                </ul>
            </div>
        )
    }

    componentDidMount () {
        axios.get(`api/task/list`).then(res => {
            res = res.data
            console.log(res.data)

            this.setState({
                list: res.data,
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
    },
)(CoinInfo)