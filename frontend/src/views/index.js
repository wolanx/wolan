import React, { Component } from 'react'
import axios from 'axios'

export default class App extends Component {
    constructor (props) {
        super(props)
        this.state = {
            msg: 'qweqwe',
            list: [],
        }
    }

    render () {
        return (
            <div>
                <h2>Index{this.state.msg}</h2>
                <ul>
                    {this.state.list.map((v, k) => {
                        return <li key={k}>{v}</li>
                    })}
                </ul>
            </div>
        )
    }

    componentDidMount () {
        axios.get('http://localhost:23456').then((res) => {
            console.log(res.data.arr)
            this.setState({
                list: res.data.arr,
            })
        }).catch((err) => {
            console.log(err.status)
        })
    }

}