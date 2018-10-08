import axios from 'axios'
import { Button, Loading, Table } from 'element-react'
import React, { Component } from 'react'
import { connect } from 'react-redux'

class CoinInfo extends Component {
    id = this.props.match.params.id
    state = {
        info: {},
    }

    render () {
        const info = this.state.info
        return (
            <div>
                <h1>info {info.id}</h1>
                <Button onClick={this.start.bind(this)}>Run</Button>
            </div>
        )
    }

    componentDidMount () {
        axios.get(`api/task/${this.id}`).then(res => {
            res = res.data
            console.log(res.data)

            this.setState({
                // info: {},
            })
        }, res => {
            res = res.response.data
            console.log(res.message)
        })
    }

    start () {
        axios.post(`api/task/${this.id}/run`, {
            name: this.id
        }).then(res => {
            res = res.data
            console.log(res.data)
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
    },
)(CoinInfo)