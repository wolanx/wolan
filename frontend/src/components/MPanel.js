import React from 'react'
import { string } from 'prop-types'
import css from '../assets/css/main.less'

class MPanel extends React.Component {
    render () {
        return (
            <section className={css.mPanel}>
                <div className={'p1'}>Title</div>
                <div className={'p2'}>
                    {this.props.children}
                </div>
            </section>
        )
    }
}

MPanel.propTypes = {
    title: string
}

export default MPanel
