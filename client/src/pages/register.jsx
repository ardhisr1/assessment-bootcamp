import React,{useEffect} from 'react';
import { Form, Button } from 'react-bootstrap';
import { useHistory } from 'react-router';
import { useDispatch, useSelector } from 'react-redux';
import userRegisterAction from '../redux/action/userRegisterAction';

const Register = () => {
    const registerData = useSelector(state => state.userRegister)
    const dispatch = useDispatch()
    const history = useHistory()

    const registerHandler = e => {
        e.preventDefault()
        dispatch(userRegisterAction.userRegister(
            registerData.fullname,
            registerData.address,
            registerData.email,
            registerData.password
        ))
        history.push("/login");
    }
    return (
        <div>
        <Form onSubmit={registerHandler}>
        <Form.Group controlId="formBasicEmail">
            <Form.Label>Full Name</Form.Label>
            <Form.Control 
            type="text" 
            placeholder="Enter full name" 
            value={registerData.fullname}
            onChange={e => dispatch(userRegisterAction.setFullname(e.target.value))}/>
        </Form.Group>
        <Form.Group controlId="formBasicEmail">
            <Form.Label>Address</Form.Label>
            <Form.Control 
            type="text" 
            placeholder="Enter address" 
            value={registerData.address}
            onChange={e => dispatch(userRegisterAction.setAddress(e.target.value))} />
        </Form.Group>
        <Form.Group controlId="formBasicEmail">
            <Form.Label>Email address</Form.Label>
            <Form.Control 
            type="email" 
            placeholder="Enter email"v
            value={registerData.email}
            onChange={e=> dispatch(userRegisterAction.setEmail(e.target.value))} />
            <Form.Text className="text-muted">
            We'll never share your email with anyone else.
            </Form.Text>
        </Form.Group>
        <Form.Group controlId="formBasicPassword">
            <Form.Label>Password</Form.Label>
            <Form.Control 
            type="password" 
            placeholder="Password"  
            value={registerData.password}
            onChange={e => dispatch(userRegisterAction.setPassword(e.target.value))}/>
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

export default Register;