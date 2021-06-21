import React,{useEffect} from 'react';
import { Form, Button } from 'react-bootstrap';
import { useHistory } from 'react-router';
import { useDispatch, useSelector } from 'react-redux';
import userLoginAction from '../redux/action/userLoginAction';

const Login = () => {
    const loginData = useSelector(state=> state.userLogin)
    const dispatch = useDispatch()
    const history = useHistory()

    useEffect(() => {
        console.log(loginData)
    }, [])
    const loginHandler = e => {
        e.preventDefault()
        dispatch(userLoginAction.userLogin(
            loginData.email,
            loginData.password
        ))
        history.push("/dashboard");
    }
    return (
        <div>
        <Form onSubmit={loginHandler}>
        <Form.Group controlId="formBasicEmail">
            <Form.Label>Email address</Form.Label>
            <Form.Control 
            type="email" 
            placeholder="Enter email" 
            value={loginData.email}
            onChange={e=> dispatch(userLoginAction.setEmail(e.target.value))} />
            <Form.Text className="text-muted">
            We'll never share your email with anyone else.
            </Form.Text>
        </Form.Group>

        <Form.Group controlId="formBasicPassword">
            <Form.Label>Password</Form.Label>
            <Form.Control 
            ype="password" 
            placeholder="Password" 
            value={loginData.password}
            onChange={e => dispatch(userLoginAction.setPassword(e.target.value))}/>
        </Form.Group>
        <Form.Group controlId="formBasicCheckbox">
            <Form.Check type="checkbox" label="Check me out" />
        </Form.Group>
        <Button variant="primary" type="submit">
            Submit
        </Button>
        </Form>
        </div>
    )
}

export default Login;