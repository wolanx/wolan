import request, { graphql } from '../utils/request'
import { gql } from 'apollo-boost'

export function getList () {
    return graphql({
        query: gql`
            {
                listTask{
                    name
                    version
                }
            }
        `
    })
}

export function getPost () {
    let values = {}
    return request('/api/users', {
        method: 'POST',
        body: JSON.stringify(values),
    })

    // return request(`http://localhost:8080/graphql`, {
    //     method: 'POST',
    //     body: `{"query":"{\\n  listTask{\\n    name\\n  }\\n}\\n","variables":null,"operationName":null}`,
    // })
}

