import axios from 'axios'
import { Loading, Menu } from 'element-react'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Link } from 'react-router-dom'

class App extends Component {
    state = {
        loading: true,
        list: [],
    }

    render () {
        return (
            <div>
                <Menu className="el-menu-demo" mode="horizontal">
                    <Menu.Item index={'index'}><Link to="/news">快讯</Link></Menu.Item>
                    <Menu.Item index={'danbi'}><Link to="/news/list?type=market">行情分析</Link></Menu.Item>
                    <Menu.Item index={'news'}><Link to="/coin/rating-agency">项目评级</Link></Menu.Item>
                    <Menu.Item index={'zixuan'}><Link to="/news/list?type=depth">深度</Link></Menu.Item>
                </Menu>

                <Loading loading={this.state.loading}>
                    <ul>
                        {this.state.list.map((v, k) => (
                            <li key={k}>
                                <hr/>
                                <h1>{v.name}</h1>
                                <h2>{v.desc}</h2>
                            </li>
                        ))}
                    </ul>
                </Loading>
            </div>
        )
    }

    componentDidMount () {
        document.setTitle('项目评级')

        axios.get(`api/coin/rating-agency`).then(res => {
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