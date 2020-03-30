import React from 'react'
import { Router } from 'dva/router'
import LayoutPage from './routes/LayoutPage'

function RouterConfig ({ history }) {
    return (
        <Router history={history}>
            <LayoutPage/>
        </Router>
    )
}

export default RouterConfig
