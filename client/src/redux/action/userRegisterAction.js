import clientApi from "../../api/clientApi"

const setFullname = fullname =>{
    return {
            type: "USER_REGISTER_SET_FULLNAME",
            payload: {
                fullname: fullname
            }
        } 
}

const setAddress = address =>{
    return {
            type: "USER_REGISTER_SET_ADDRESS",
            payload: {
                address: address
            }
        } 
}

const setEmail = email =>{
    return {
            type: "USER_REGISTER_SET_EMAIL",
            payload: {
                email: email
            }
        } 
}

const setPassword = password =>{
    return {
            type: "USER_REGISTER_SET_PASSWORD",
            payload: {
                password: password
            }
        } 
}

const resetForm = () => {
    return {
        type: "USER_REGISTER_RESET_FORM"
    }
}

const userRegister = (fullname, address, email, password) => async dispatch => {
    try {
        console.log("Registering user..")

        const newData = {
            "fullname":fullname,
            "address":address,
            "email":email,
            "password":password
        }

        const newUser = await clientApi({
            method: "POST",
            url:"/user/register",
            data:newData
        })

        console.log("Register Success")
    } catch (error) {
        console.log(error)
    }
}

const userRegisterAction = {
    setAddress,
    setEmail,
    setFullname,
    setPassword,
    userRegister,
    resetForm,
}

export default userRegisterAction;