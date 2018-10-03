import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Redirect, Route } from 'react-router-dom'

import NavHeader from '../nav/header'

const LayoutBase = ({component: Component, ...rest, needLogin}) => {
    const toLogin = needLogin && !rest.user.id
    window.log('Enter with', rest)

    return (
        <div>
            {toLogin ? (
                <Redirect to={{
                    pathname: '/user/login',
                    search: `from=${rest.path}`,
                }}
                />
            ) : (
                <Route {...rest} render={props => (
                    <div className="DefaultLayout">
                        <NavHeader/>
                        <div className="wol-1200">
                            <Component {...props} />
                        </div>
                        <footer className="wol-footer">Footer</footer>
                    </div>
                )}/>
            )}
        </div>
    )
}

export default connect(
    (state) => {
        return {
            user: state.user
        }
    },
    (dispatch) => {
        return {
            //
        }
    }
)(LayoutBase)
