import fetch from 'dva/fetch'
import ApolloClient from 'apollo-boost'

const client = new ApolloClient({
    uri: "${window.location.origin}/graphql",
})

// function delay (timeout) {
//     return new Promise(resolve => setTimeout(resolve, timeout))
// }

function parseJSON (response) {
    return response.json()
}

function checkStatus (response) {
    console.log('response', response)
    if (response.status >= 200 && response.status < 300) {
        return response
    }

    const error = new Error(response.statusText)
    error.response = response
    throw error
}

/**
 * Requests a URL, returning a promise.
 *
 * @param  {string} url       The URL we want to request
 * @param  {object} [options] The options we want to pass to "fetch"
 * @return {object}           An object containing either "data" or "err"
 */
export default function request (url, options) {
    return fetch(url, options)
        .then(checkStatus)
        .then(parseJSON)
        .then(data => ({ data }))
        .catch(err => ({ err }))
}

export function graphql (options) {
    return client.query(options)
        .then(data => ({ data }))
        .catch(err => ({ err }))
}
