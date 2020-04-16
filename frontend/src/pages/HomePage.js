import React from 'react'
import { connect } from 'dva'
import { Result } from 'antd'
import MHeader from '../components/MHeader'
import { Link } from 'dva/router'

function HomePage () {
    return (
        <>
            <MHeader title={'Home'}/>
            <Result
                icon={null}
                title="Great, we have done all the operations!"
                extra={<Link to={'/tasks'}>Next</Link>}
            />
        </>
    )
}

HomePage.propTypes = {}

export default connect()(HomePage)
