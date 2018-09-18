import React, { Component } from 'react'
import { connect } from 'react-redux'

// React component
class Counter extends Component {
    render () {
        const {value, onIncreaseClick} = this.props
        return (
            <div>
                <span>{value}</span>
                <button onClick={onIncreaseClick}>Increase</button>
            </div>
        )
    }
}

export default connect(
    (state) => {
        return {
            value: state.counter.count
        }
    },
    (dispatch) => {
        return {
            onIncreaseClick: () => dispatch({type: 'increase'})
        }
    }
)(Counter)