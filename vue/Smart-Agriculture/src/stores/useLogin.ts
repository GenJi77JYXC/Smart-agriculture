import {ref} from 'vue'
import axios from 'axios'

export default function () {
    let username = ref()
    let password = ref()

    async function login(username:string, password:string) {
        try {
            axios.post("http://127.0.0.1:8080/login", {username: username, password: password})
            .then((res) => {
                console.log(res.data)
                console.log("密码是：",res.data.password)
                console.log("用户名是：", res.data.username)
            })
            
        } catch (error) {
            console.log(error)
        }
    }
    

    return {username, password, login}
}