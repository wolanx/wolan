import axios from 'axios'
import React, { Component } from 'react'
import { Provider } from 'react-redux'
import { BrowserRouter as Router, Link, Route, Switch } from 'react-router-dom'
import { createStore } from 'redux'

import '../assets/less/main.less'
import LayoutBase from '../components/layout/layout'
import WS from '../components/ws/WS'
import reducers from '../reducers'
import CoinInfo from './coin/CoinInfo'
import RatingAgency from './coin/RatingAgency'
import ExchangeInfo from './exchange/ExchangeInfo'
import MarketCoin from './market/MarketCoin'
import MarketExchange from './market/MarketExchange'
import NewsFlash from './news/NewsFlash'
import NewsInfo from './news/NewsInfo'
import NewsList from './news/NewsList'
import Home from './other/HomeIndex'
import OptionalCp from './user/OptionalCp'
import UserLogin from './user/UserLogin'
import UserReg from './user/UserReg'

axios.defaults.withCredentials = true
axios.interceptors.request.use(function (config) {
    if (!/^http/.test(config.url)) {
        // config.url = 'https://www.bitdata.com.cn/' + config.url
        config.url = 'http://localhost:777/' + config.url
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

window.ws = new WS()

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

                                <LayoutBase path="/market/coin" component={MarketCoin}/>
                                <LayoutBase path="/market/exchange" component={MarketExchange}/>

                                <LayoutBase path="/coin/:scope([\d:]+)" component={CoinInfo}/>
                                <LayoutBase path="/exchange/:id" component={ExchangeInfo}/>

                                <LayoutBase path="/user/login" component={UserLogin}/>
                                <LayoutBase path="/user/reg" component={UserReg}/>
                                <LayoutBase path="/user/optional/coinpair" component={OptionalCp} needLogin={true}/>

                                <LayoutBase exact path="/news" component={NewsFlash}/>
                                <LayoutBase path="/news/list" component={NewsList}/>
                                <LayoutBase path="/news/info" component={NewsInfo}/>
                                <LayoutBase path="/coin/rating-agency" component={RatingAgency}/>
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
            res = res.response.data
            // console.log(res.message)
            this.setState({
                loginChecked: true,
            })
        })
    }
}
