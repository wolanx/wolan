import axios from 'axios'
import { Loading } from 'element-react'
import React, { Component } from 'react'
import { connect } from 'react-redux'

class App extends Component {
    state = {
        loading: true,
        info: {},
    }

    render () {
        let info = this.state.info
        return (
            <div className="bit-news-info">
                <Loading loading={this.state.loading}>
                    <h1>{info.title}</h1>
                    <hr/>
                    <article dangerouslySetInnerHTML={{__html: info.content}}>
                    </article>
                </Loading>
            </div>
        )
    }

    componentDidMount () {
        const params = new URLSearchParams(this.props.location.search)
        let id = params.get('id')

        axios.get(`api/news/info?id=${id}`).then(res => {
            res = res.data

            document.setTitle(res.data.title)
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