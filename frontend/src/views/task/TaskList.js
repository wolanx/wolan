import axios from 'axios'
import { Button, Loading, Message, Table, Tag } from 'element-react'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Link } from 'react-router-dom'

class CoinInfo extends Component {
    state = {
        list: [],
        columns: [
            {
                type: 'index'
            },
            {
                label: 'name',
                prop: 'name',
                width: 200,
                render: (row) => {
                    return (
                        <span>{row.name}</span>
                    )
                }
            },
            {
                label: 'opt',
                prop: 'opt',
                render: (row) => {
                    return (
                        <div>
                            <Link to={`/task/${row.name}`}><Button size="small">Info</Button></Link>
                        </div>
                    )
                }
            }
        ],
    }

    render () {
        return (
            <div>
                <Table
                    style={{width: '100%'}}
                    columns={this.state.columns}
                    data={this.state.list}
                    border={true}
                    highlightCurrentRow={true}
                />
            </div>
        )
    }

    componentDidMount () {
        axios.get(`api/task/list`).then(res => {
            res = res.data
            console.log(res.data)

            this.setState({
                list: res.data,
            })
        }, res => {
            res = res.response.data
            console.log(res.message)
            Message.error(res.message)
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