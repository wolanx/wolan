import request from '../utils/request'

export function getList () {
    // return request(`http://localhost:8080/other/aaa`)
    return request(`http://localhost:8080/graphql`,{
        method: 'POST',
        body: `{"query":"{\\n  listTask{\\n    name\\n  }\\n}\\n","variables":null,"operationName":null}`,
    })
}

export function getPost () {
    let values = {}
    return request('/api/users', {
        method: 'POST',
        body: JSON.stringify(values),
    })
}

