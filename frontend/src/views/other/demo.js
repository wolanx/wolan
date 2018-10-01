import axios from 'axios'
import React, { Component } from 'react'
import { connect } from 'react-redux'

class App extends Component {
    render () {
        return (
            <div>
                <span>123</span>
            </div>
        )
    }

    componentDidMount () {
        document.title = '123' + window.title_prefix
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