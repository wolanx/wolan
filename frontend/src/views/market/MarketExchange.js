import axios from 'axios'
import { Button, Loading, Table } from 'element-react'
import React, { Component } from 'react'
import { Link } from 'react-router-dom'

export default class extends Component {
    state = {
        loading: true,
        list: [],
    }

    render () {
        return (
            <div>
                <Loading loading={this.state.loading}>
                    <table>
                        <thead>
                        <tr>
                            <th width="120">name</th>
                            <th width="100">volume</th>
                            <th>coinpair_num</th>
                            <th>grade</th>
                            <th>tag_views</th>
                        </tr>
                        </thead>
                        <tbody>
                        {this.state.list.map((v, k) => {
                            return (
                                <tr key={k}>
                                    <td><Link to={`/exchange/${v.id}`}>{v.name_view}</Link></td>
                                    <td>{v.volume}</td>
                                    <td>{v.coinpair_num}</td>
                                    <td>{v.grade}</td>
                                    <td>{v.tag_views.map((v, k) => (<button key={k}>{v}</button>))}</td>
                                </tr>
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

        axios.get('api/market/exchange').then(res => {
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
