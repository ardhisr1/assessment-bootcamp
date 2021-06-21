import Axios from "axios"

const clientApi = Axios.create({
    baseURL:"http://localhost:9000"
});

export default clientApi;
