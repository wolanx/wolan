import axios from 'axios'
import { Loading, Menu } from 'element-react'
import $ from 'jquery'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Link } from 'react-router-dom'

class App extends Component {
    state = {
        loading: true,
        list: [],
        last_id: 0,
    }

    constructor (props) {
        super(props)
        this.handleScroll = this.handleScroll.bind(this)
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

                <ul>
                    {this.state.list.map((v, k) => (
                        <li key={k}>
                            <hr/>
                            <h1>{v.publish_ts_view} - {v.title}</h1>
                            <h2>{v.summary}</h2>
                        </li>
                    ))}
                    <Loading loading={this.state.loading}/>
                </ul>
            </div>
        )
    }

    componentDidMount () {
        document.setTitle('快讯')

        this.fetchData()

        window.addEventListener('scroll', this.handleScroll, false)
    }

    componentWillUnmount () {
        window.removeEventListener('scroll', this.handleScroll, false)
    }

    fetchData () {
        this.setState({
            loading: true,
        })

        axios.get(`api/news/list?type=flash&last_id=${this.state.last_id}`).then(res => {
            res = res.data

            let temp = this.state.list.concat(res.data)

            this.setState({
                loading: false,
                list: temp,
                last_id: temp[temp.length - 1].id,
            })
        }, res => {
            res = res.response.data
            console.log(res.message)
        })
    }

    handleScroll () {
        if (this.state.loading) {
            return
        }
        if (($(window).scrollTop() + $(window).height()) >= ($(document).height() - 50)) {
            this.fetchData()
        }
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