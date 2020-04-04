import { Layout, Menu } from 'antd'
import { DesktopOutlined, PieChartOutlined,LinkOutlined } from '@ant-design/icons'
import React, { Component } from 'react'

import css from '../assets/css/main.less'
import { Link, Route, Switch } from 'dva/router'
import TaskInfoPage from './TaskInfoPage'
import TaskListPage from './TaskListPage'
import IndexPage from './IndexPage'

const { Content, Sider } = Layout

export const NotFound404 = (props) => (
    <div className="whoops-404">
        <h1>没有页面可以匹配</h1>
    </div>
)

let routes = [
    {
        path: '/',
        component: IndexPage,
        exact: true
    },
    {
        path: '/test',
        component: TaskListPage
    },
    {
        path: '/tasks',
        component: TaskListPage
    },
    {
        path: '/task/:name',
        component: TaskInfoPage
    },
    {
        path: '*',
        component: NotFound404
    },
]

export default class LayoutPage extends Component {
    state = {
        collapsed: false,
    }

    toggle = () => {
        this.setState({
            collapsed: !this.state.collapsed,
        })
    }

    render () {
        return (
            <Layout style={{ minHeight: '100vh' }}>
                <Sider collapsible collapsed={this.state.collapsed} onCollapse={this.toggle}>
                    <Link className={css.logo} to={'/'}>Logo</Link>
                    <Menu mode="inline">
                        <Menu.Item key="1">
                            <Link to={'/tasks'}>
                                <DesktopOutlined/>
                                <span>Dashboard</span>
                            </Link>
                        </Menu.Item>
                        <Menu.Item key="2">
                            <Link to={'/task/go-fs'}>
                                <PieChartOutlined/>
                                <span>Task Info</span>
                            </Link>
                        </Menu.Item>
                        <Menu.Item key="3">
                            <Link to={'/no'}>
                                <LinkOutlined/>
                                <span>Not found</span>
                            </Link>
                        </Menu.Item>
                    </Menu>
                </Sider>
                <Layout>
                    <Content>
                        <Switch>
                            {
                                routes.map((v, k) => {
                                    return <Route exact={v.exact} key={k} path={v.path} render={props => (
                                        <v.component {...props}/>
                                    )}/>
                                })
                            }
                        </Switch>
                    </Content>
                </Layout>
            </Layout>
        )
    }
}
