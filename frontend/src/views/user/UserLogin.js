import axios from 'axios'
import { Button, Form, Input, Message } from 'element-react'
import PropTypes from 'prop-types'
import React, { Component } from 'react'
import { connect } from 'react-redux'

class UserLogin extends Component {
    static contextTypes = {
        router: PropTypes.object.isRequired,
    }

    constructor (props) {
        super(props)

        this.state = {
            form: {
                username: '13816076964',
                password: 'qwerqwer',
            },
            rules: {
                username: [
                    {required: true, message: '请输入正确的账号', trigger: 'blur'},
                ],
                password: [
                    {required: true, message: '请输入正确的密码', trigger: 'blur'},
                    {
                        validator: (rule, value, callback) => {
                            if (value.length < 4) {
                                callback(new Error('请输入正确的密码'))
                            } else {
                                callback()
                            }
                        }, trigger: 'change'
                    }
                ],
            }
        }
    }

    render () {
        return (
            <Form ref="form" model={this.state.form} rules={this.state.rules} labelWidth="100">
                <Form.Item label="账号" prop="username">
                    <Input type="username" value={this.state.form.username} onChange={this.onChange.bind(this, 'username')} placeholder="请输入您的手机号或者邮箱"/>
                </Form.Item>
                <Form.Item label="密码" prop="password">
                    <Input type="password" value={this.state.form.password} onChange={this.onChange.bind(this, 'password')} placeholder="请输入您的登录密码" autoComplete="off"/>
                </Form.Item>
                <Form.Item>
                    <Button type="primary" onClick={this.handleSubmit.bind(this)}>登录</Button>
                </Form.Item>
            </Form>
        )
    }

    componentDidMount () {
        document.title = '登录' + window.title_prefix
    }

    handleSubmit (e) {
        e.preventDefault()

        this.refs.form.validate((valid) => {
            if (valid) {
                let type = 2
                let username = this.state.form.username
                if (!/@/.test(this.state.form.username)) {
                    type = 1
                    username = '+86-' + this.state.form.username
                }

                axios.post('api/user/login', {
                    type: type,
                    username: username,
                    password: this.state.form.password,
                }).then(res => {
                    res = res.data
                    console.log(res.data)
                    this.props.storeUserSet(res.data)
                    {
                        const params = new URLSearchParams(this.props.location.search)
                        this.context.router.history.push(params.get('from') || '/')
                    }
                }, res => {
                    res = res.response.data
                    console.log(res.message)
                })
            } else {
                console.log('error submit!!')
                return false
            }
        })
    }

    onChange (key, value) {
        this.setState({
            form: Object.assign({}, this.state.form, {[key]: value})
        })
    }
}

export default connect(
    (state) => {
        return {
            user: state.user
        }
    },
    (dispatch) => {
        return {
            storeUserSet: (user) => dispatch({type: 'user/set', val: user})
        }
    }
)(UserLogin)
