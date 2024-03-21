<script lang="ts" setup>
import useLogin from '../stores/useLogin'
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import Password from 'primevue/password';





const {username, password, login} = useLogin()



</script>

<template>
    <div>
        <div class="card flex justify-content-center">
            <div class="flex flex-column gap-2">
                <label for="username">Username</label>
                <InputText id="username" v-model="username" aria-describedby="username-help" />
                <small id="username-help">Enter your username to reset your password.</small>
            </div>
        </div>
        <div class="card flex justify-content-center">
            <Password v-model="password" :feedback="false" />
        </div>
        <Button @click="login(username, password)">提交</Button>
    </div>
    
</template>
