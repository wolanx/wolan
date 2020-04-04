import React from 'react'
import css from '../routes/Layout.less'
import { Layout } from 'antd'
import { string } from 'prop-types'

const { Header } = Layout

class MHeader extends React.Component {
    render () {
        return (
            <Header className={css.header} style={{ padding: 0 }}>
                {this.props.title}
            </Header>
        )
    }
}

MHeader.propTypes = {
    title: string
}

export default MHeader
