import React from 'react'
import css from '../assets/css/main.less'
import { Breadcrumb, Layout } from 'antd'
import { Link } from 'dva/router'
import { string } from 'prop-types'

const { Header } = Layout

class MHeader extends React.Component {
    render () {
        return (
            <Header className={css.mHeader}>
                <div className={'p1'}>
                    <span>{this.props.title}</span>
                    <span className={'fr'}>admin</span>
                </div>
                <div className={'p2'}>
                    <Breadcrumb className={'fl'} separator={'>'}>
                        <Breadcrumb.Item>
                            <a href="qwe">Tasks</a>
                        </Breadcrumb.Item>
                        <Breadcrumb.Item>
                            <a href="asd">{this.props.title}</a>
                        </Breadcrumb.Item>
                    </Breadcrumb>
                    <Link className={'fr'} to={'/logout'}>Logout</Link>
                </div>
            </Header>
        )
    }
}

MHeader.propTypes = {
    title: string
}

export default MHeader
