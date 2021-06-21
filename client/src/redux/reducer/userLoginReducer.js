const initialState = {
    email: "",
    password: "",
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
        default:
            return state
    }
}

export default userLoginReducer;