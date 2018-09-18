import 'element-theme-default'
import React from 'react'
import ReactDOM from 'react-dom'
import './assets/reset.css'
import registerServiceWorker from './registerServiceWorker'
import Index from './views/root'

/*global log*/
window.log = function (...b) {
    console.log('%c %s', 'color:#f0f;background:#eee;', ...b)
}

ReactDOM.render(<Index/>, document.getElementById('root'))
registerServiceWorker()
