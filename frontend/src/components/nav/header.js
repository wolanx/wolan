import axios from 'axios'
import { Menu } from 'element-react'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Link } from 'react-router-dom'

class App extends Component {
    render () {
        const user = this.props.user

        return (
            <div className="wol-header">
                <Menu className="el-menu-demo" mode="horizontal">
                    <Menu.Item index={'index'}><Link to="/">首页</Link></Menu.Item>
                    <Menu.Item index={'task'}><Link to="/task/list">Task</Link></Menu.Item>
                    {
                        user.id ? (
                            <Menu.Item index={'logout'}><a onClick={this.logout.bind(this)}>{user.phone}注销</a></Menu.Item>
                        ) : (
                            <Menu.Item index={'login'}><Link to="/user/login">登录</Link></Menu.Item>
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
            user: state.user,
        }
    },
    (dispatch) => {
        return {
            storeUserLogout: () => dispatch({type: 'user/logout'}),
        }
    },
)(App)
