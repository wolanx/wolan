import dva from 'dva'
import './index.css'

const createHistory = require('history').createHashHistory

// 1. Initialize
const app = dva({
    history: createHistory(),
})

// 2. Plugins
// app.use({});

// 3. Model
app.model(require('./models/example').default)
app.model(require('./models/task').default)

// 4. Router
app.router(require('./router').default)

// 5. Start
app.start('#root')
