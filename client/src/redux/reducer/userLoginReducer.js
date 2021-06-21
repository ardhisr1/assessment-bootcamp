const initialState = {
    email: "",
    password: "",
    token:""
}

const userLoginReducer = (state = initialState, action) => {
    switch(action.type){
        case "RESET_FORM":
            return {
                ...initialState
            }
        case "USER_LOGIN_SET_EMAIL":
            return {
                ...state,
                email: action.payload.email
            }
        case "USER_LOGIN_SET_PASSWORD":
            return {
                ...state,
                password: action.payload.password
            }
        case "USER_LOGIN_SET_TOKEN":
            return {
                ...state,
                token: action.payload.token
            }
        default:
            return state
    }
}

export default userLoginReducer;