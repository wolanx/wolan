import { Layout, Menu } from 'antd'
import { DatabaseOutlined, HomeOutlined, LinkOutlined } from '@ant-design/icons'
import React, { Component } from 'react'

import css from '../assets/css/main.less'
import { Link, Route, Switch, withRouter } from 'dva/router'
import TaskInfoPage from './TaskInfoPage'
import TaskListPage from './TaskListPage'
import IndexPage from './HomePage'
import { connect } from 'dva'

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
        path: '/task/:sid',
        component: TaskInfoPage
    },
    {
        path: '*',
        component: NotFound404
    },
]

@withRouter
@connect(state => ({
    pathNow: state.layout.pathNow
}), dispatch => ({
    //
}))
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
        let selectedKeys = []
        if (this.props.pathNow === '/') {
            selectedKeys.push('home')
        }
        if (this.props.pathNow.search('/task') === 0) {
            selectedKeys.push('task')
        }
        return (
            <Layout className={css.layout}>
                <Sider collapsible collapsed={this.state.collapsed} onCollapse={this.toggle}>
                    <Link className={css.logo} to={'/'}>Logo {this.props.pathNow}</Link>
                    <Menu mode="inline" selectedKeys={selectedKeys} theme='dark'>
                        <Menu.Item key="home">
                            <Link to={'/'}>
                                <HomeOutlined/>
                                <span>Home</span>
                            </Link>
                        </Menu.Item>
                        <Menu.ItemGroup key="g1" title="LOCAL"/>
                        <Menu.Item key="task">
                            <Link to={'/tasks'}>
                                <DatabaseOutlined/>
                                <span>Tasks</span>
                            </Link>
                        </Menu.Item>
                        <Menu.ItemGroup key="g2" title="SETTINGS"/>
                        <Menu.Item key="3">
                            <Link to={'/no'}>
                                <LinkOutlined/>
                                <span>Setting</span>
                            </Link>
                        </Menu.Item>
                    </Menu>
                </Sider>
                <Layout>
                    <Content>
                        <Switch>
                            {
                                routes.map((v, k) => {
                                    return <Route exact={v.exact} key={k} path={v.path} render={props => {
                                        // console.log(v.path, props.match, this.state.pathNow)
                                        return <v.component {...props}/>
                                    }}/>
                                })
                            }
                        </Switch>
                    </Content>
                </Layout>
            </Layout>
        )
    }

    componentDidUpdate (prevProps, prevState, snapshot) {
        // let pathname = this.props.location.pathname
        // console.log(pathname)
    }
}
