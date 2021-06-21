import React,{useEffect} from 'react';
import { Form, Button } from 'react-bootstrap';
import { useDispatch, useSelector } from 'react-redux';
import userRegisterAction from '../redux/action/userRegisterAction';

const Register = () => {
    const registerData = useSelector(state => state.userRegister)
    const dispatch = useDispatch()

    useEffect({
        
    },[])

    const handler= (e) => {
        e.preventDefault()
        userRegisterAction.userRegister(
            registerData.fullname
        )
    }
    return (
        <div>
        <Form>
        <Form.Group controlId="formBasicEmail">
            <Form.Label>Full Name</Form.Label>
            <Form.Control type="text" placeholder="Enter full name" />
        </Form.Group>
        <Form.Group controlId="formBasicEmail">
            <Form.Label>Address</Form.Label>
            <Form.Control type="text" placeholder="Enter address" />
        </Form.Group>
        <Form.Group controlId="formBasicEmail">
            <Form.Label>Email address</Form.Label>
            <Form.Control type="email" placeholder="Enter email" />
            <Form.Text className="text-muted">
            We'll never share your email with anyone else.
            </Form.Text>
        </Form.Group>
        <Form.Group controlId="formBasicPassword">
            <Form.Label>Password</Form.Label>
            <Form.Control type="password" placeholder="Password" />
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