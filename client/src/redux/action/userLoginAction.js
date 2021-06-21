const setEmail = email =>{
    return {
            type: "USER_LOGIN_SET_EMAIL",
            payload: {
                email: email
            }
        } 
}

const setPasword = password =>{
    return {
            type: "USER_LOGIN_SET_PASSWORD",
            payload: {
                password: password
            }
        } 
}

const resetForm = () => {
    return {
        type: "USER_LOGIN_RESET_FORM"
    }
}
