import React from 'react'
import { Router, Route, Switch } from 'dva/router'
import IndexPage from './routes/IndexPage'
import LayoutPage from './routes/LayoutPage'

function RouterConfig ({ history }) {
    return (
        <Router history={history}>
            {/*<Switch>*/}
            {/*    <Route path="/" exact component={IndexPage}/>*/}
            {/*</Switch>*/}
            <LayoutPage/>
        </Router>
    )
}

export default RouterConfig
