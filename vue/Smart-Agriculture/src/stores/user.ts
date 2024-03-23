import {ref} from 'vue'
import axios from 'axios'
 // 引入defineStore用于创建store
 import {defineStore} from 'pinia'


export const useUserStore = defineStore('user', ()=>{
    
    const username = ref('')
    const password = ref('')
    console.log(username, '-------------', password)

    

    async function login() {
        try {
            let data = new FormData();
            data.append('username', username.value)
            data.append('password', password.value)
            console.log('post 之前')
            await axios.post('http://127.0.0.1:8080/login', data)
            .then(res => {
                console.log("后端返回的数据",res.data)
            })
            console.log('post 之后')
        } catch (error) {
            console.log(error)
        }
    }

    return {username, password, login}
})
