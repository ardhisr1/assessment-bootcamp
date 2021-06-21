const initialState = {
    fullname: "",
    address:"",
    email:"",
    password:"",
}

const userRegisterReducer = (state = initialState, action) =>{
    switch(action.type){
        case "RESET_FORM":
            return {
                ...initialState
            }
        case "USER_REGISTER_SET_FULLNAME":
            return {
                ...state,
                fullname: action.payload.fullname
            }
        case "USER_REGISTER_SET_ADDRESS":
            return {
                ...state,
                address: action.payload.address
            }
        case "USER_REGISTER_SET_EMAIL":
            return {
                ...state,
                email: action.payload.email
            }
        case "USER_REGISTER_SET_PASSWORD":
            return {
                ...state,
                password: action.payload.password
            }
        default:
            return state
    }
}

export default userRegisterReducer;