import { Layout, Menu } from 'antd'
import { DesktopOutlined, MenuFoldOutlined, MenuUnfoldOutlined, PieChartOutlined } from '@ant-design/icons'
import React, { Component } from 'react'

import css from './layout.less'
import IndexPage from './IndexPage'
import { Link, Route, Switch } from 'dva/router'
import TaskInfoPage from './TaskInfoPage'
import TaskListPage from './TaskListPage'

const { Header, Content, Footer, Sider } = Layout

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
                <Sider trigger={null} collapsible collapsed={this.state.collapsed}>
                    <Link className={css.logo} to={'/'}>Logo</Link>
                    <Menu mode="inline">
                        <Menu.Item key="1">
                            <Link to={'/task/list'}>
                                <DesktopOutlined/>
                                <span>Task List</span>
                            </Link>
                        </Menu.Item>
                        <Menu.Item key="2">
                            <Link to={'/task/info'}>
                                <PieChartOutlined/>
                                <span>Task Info</span>
                            </Link>
                        </Menu.Item>
                    </Menu>
                </Sider>
                <Layout>
                    <Header className={css.header} style={{ padding: 0 }}>
                        {React.createElement(this.state.collapsed ? MenuUnfoldOutlined : MenuFoldOutlined, {
                            className: 'trigger',
                            onClick: this.toggle,
                        })}
                    </Header>
                    <Content className={css.content}>
                        <Switch>
                            <Route exact path="/" component={IndexPage}/>
                            <Route exact path="/task/list" component={TaskListPage}/>
                            <Route exact path="/task/:name" component={TaskInfoPage}/>
                        </Switch>
                    </Content>
                    <Footer className={css.footer}>©2020</Footer>
                </Layout>
            </Layout>
        )
    }
}
