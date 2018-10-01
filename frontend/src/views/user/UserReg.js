import React, { Component } from 'react'

export default class extends Component {
    render () {
        return (
            <div>
                <h1>reg</h1>
            </div>
        )
    }

    componentDidMount () {
        document.title = '注册' + window.title_prefix
    }
}
