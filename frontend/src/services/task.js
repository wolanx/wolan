import request from '../utils/request'

export function getList () {
    return request(`http://localhost:8080/other/aaa`)
}

export function getPost () {
    let values = {}
    return request('/api/users', {
        method: 'POST',
        body: JSON.stringify(values),
    })
}

