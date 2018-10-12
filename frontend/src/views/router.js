import axios from 'axios'
import React, { Component } from 'react'
import { Provider } from 'react-redux'
import { BrowserRouter as Router, Link, Route, Switch } from 'react-router-dom'
import { createStore } from 'redux'

import '../assets/less/main.less'
import LayoutBase from '../components/layout/layout'
import reducers from '../reducers'
import ExchangeInfo from './exchange/ExchangeInfo'
import Home from './other/HomeIndex'
import OptionalCp from './user/OptionalCp'
import UserLogin from './user/UserLogin'
import UserReg from './user/UserReg'
import TaskList from './task/TaskList'
import TaskInfo from './task/TaskInfo'

// axios.defaults.withCredentials = true
axios.interceptors.request.use(function (config) {
    if (!/^http/.test(config.url)) {
        if (process.env.NODE_ENV === 'development') {
            config.url = 'http://localhost:23456/' + config.url
        } else {
            config.url = '/' + config.url
        }
    }

    return config
})
/*axios.interceptors.response.use((response) => {
    return response;
}, function (error) {
    // Do something with response error
    if (error.response.status === 401) {
        console.log('unauthorized, logging out ...');
        auth.logout();
        router.replace('/auth/login');
    }
    return Promise.reject(error.response);
});*/

let store = createStore(
    reducers,
    window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()
)

export default class extends Component {
    constructor (props) {
        super(props)

        this.state = {
            loginChecked: false,
        }
    }

    render () {
        if (!this.state.loginChecked) {
            this.loginCheck()
        }

        return (
            <div>
                {this.state.loginChecked === false ? (
                    <div>Loading..</div>
                ) : (
                    <Provider store={store}>
                        <Router>
                            <Switch>
                                <LayoutBase exact path="/" component={Home}/>

                                <LayoutBase path="/task/list" component={TaskList}/>
                                <LayoutBase path="/task/:id" component={TaskInfo}/>

                                <LayoutBase path="/coin/:scope([\d:]+)" component={TaskList}/>
                                <LayoutBase path="/exchange/:id" component={ExchangeInfo}/>

                                <LayoutBase path="/user/login" component={UserLogin}/>
                                <LayoutBase path="/user/reg" component={UserReg}/>
                                <LayoutBase path="/user/optional/coinpair" component={OptionalCp} needLogin={true}/>

                            </Switch>
                        </Router>
                    </Provider>
                )}
            </div>
        )
    }

    loginCheck () {
        axios.get('api/user/info').then(res => {
            res = res.data
            // console.log(res.data)
            store.dispatch({type: 'user/set', val: res.data})
            this.setState({
                loginChecked: true,
            })
        }, res => {
            console.log(res)
            res = res.response.data
            // console.log(res.message)
            this.setState({
                loginChecked: true,
            })
        })
    }
}
