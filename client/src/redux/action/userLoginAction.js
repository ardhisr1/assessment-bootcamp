import clientApi from "../../api/clientApi"
import { useHistory } from 'react-router';

const setEmail = email =>{
    return {
            type: "USER_LOGIN_SET_EMAIL",
            payload: {
                email: email
            }
        } 
}

const setPassword = password =>{
    return {
            type: "USER_LOGIN_SET_PASSWORD",
            payload: {
                password: password
            }
        } 
}

const setToken = token =>{
    return {
            type: "USER_LOGIN_SET_TOKEN",
            payload: {
                token: token
            }
        } 
}

const resetForm = () => {
    return {
        type: "USER_LOGIN_RESET_FORM"
    }
}

const userLogin = (email, password) => async dispatch =>{
    try {
        console.log("login user..")

        const newData = {
            "email":email,
            "password":password
        }

        const newUser = await clientApi({
            method: "POST",
            url:"/user/login",
            data:newData
        })

        const token = newUser.response.Data.Authorization
        console.log(token)
        dispatch(setToken(token))
        console.log("Login Success")

    } catch (error) {
        console.log(error)
    }
}

const userLoginAction = {
    setPassword,
    setEmail,
    resetForm,
    userLogin
}

export default userLoginAction;