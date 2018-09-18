import axios from 'axios'
import { Button, Menu } from 'element-react'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Link } from 'react-router-dom'

class App extends Component {
    render () {
        const user = this.props.user

        return (
            <div className="bit-header">
                <Menu className="el-menu-demo" mode="horizontal">
                    <Menu.Item index={'a1'}><Link to="/">首页</Link></Menu.Item>
                    <Menu.Item index={'b2'}><Link to="/market/coin">单币</Link></Menu.Item>
                    {
                        user.id ? (
                            <Menu.Item index={'zhuxiao'}><a onClick={this.logout.bind(this)}>注销{user.phone}</a></Menu.Item>
                        ) : (
                            <div>
                                <Menu.Item index={'c3'}><Link to="/user/login">登录</Link></Menu.Item>
                                <Menu.Item index={'d4'}><Link to="/user/reg">注册</Link></Menu.Item>
                            </div>
                        )
                    }
                </Menu>
            </div>
        )
    }

    logout () {
        axios.get('api/user/logout').then(res => {
            res = res.data
            console.log(res.data)
            this.props.storeUserLogout()
        }, res => {
            res = res.response.data
            console.log(res.message)
        })
    }
}

export default connect(
    (state) => {
        return {
            user: state.user
        }
    },
    (dispatch) => {
        return {
            storeUserLogout: () => dispatch({type: 'user/logout'})
        }
    }
)(App)
