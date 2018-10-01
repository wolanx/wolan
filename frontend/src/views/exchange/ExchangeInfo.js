import axios from 'axios'
import React, { Component } from 'react'
import { connect } from 'react-redux'

class App extends Component {
    exid = this.props.match.params.id
    state = {
        info: {}
    }

    render () {
        let info = this.state.info

        return (
            <div>
                <h1>{info.name_view}</h1>
                <article>{info.content}</article>
            </div>
        )
    }

    componentDidMount () {
        axios.get(`api/exchange/${this.exid}/info`).then(res => {
            res = res.data

            document.setTitle(`${res.data.name} - 交易所`)
            this.setState({
                loading: false,
                info: res.data,
            })
        }, res => {
            res = res.response.data
            console.log(res.message)
        })
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