<script lang="ts" setup>

// import axios from 'axios';
// import { ref } from 'vue';
import {useUserStore} from '@/stores/user'

const userStore = useUserStore();

// const username = ref('')
// const password = ref('')

// function login() {

//     try {
//         let data = new FormData();
//         data.append('username', username.value)
//         data.append('password', password.value)
//         console.log(username.value, password.value)
//         console.log('post 之前')
//         console.log(data)
//              axios.post('http://127.0.0.1:8080/login', data)
//             .then(res => {
//                 console.log("后端返回的数据",res.data.data)
//             })
//             console.log('post 之后')
//         } catch (error) {
//             console.log('axios出错')
//             console.log(error)
//         }
//     }



</script>

<template>
    <div class="login-container">
    <h2>登录页</h2>
    <div class="p-inputgroup">
      <span class="p-inputgroup-addon">
        <i class="pi pi-user"></i>
      </span>
      <InputText v-model="userStore.username" placeholder="请输入用户名" />
    </div>
    <div class="p-inputgroup">
      <span class="p-inputgroup-addon">
        <i class="pi pi-key"></i>
      </span>
      <Password v-model="userStore.password" placeholder="请输入密码" />
    </div>
    <Button label="登录" @click="userStore.login()" />
  </div>
    
</template>
