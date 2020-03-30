import request from '../utils/request'

export function getList () {
    return request(`http://localhost:8080`)
}
