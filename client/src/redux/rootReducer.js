import { combineReducers } from "redux";

import userRegister from "./reducer/userRegisterReducer"
import userLogin from "./reducer/userLoginReducer"

const rootReducer = combineReducers({
    userRegister: userRegister,
    userLogin: userLogin,
})

export default rootReducer;