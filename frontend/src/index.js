import 'element-theme-default'
import React from 'react'
import ReactDOM from 'react-dom'
import './assets/reset.css'
import registerServiceWorker from './registerServiceWorker'
import Root from './views/router'

window.title_prefix = ' - wolan'
window.log = function (...b) {
    console.log('%c %s', 'color:#f0f;background:#eee;', ...b)
}
document.setTitle = function (name) {
    document.title = name + window.title_prefix
}

ReactDOM.render(<Root/>, document.getElementById('root'))
// registerServiceWorker()
