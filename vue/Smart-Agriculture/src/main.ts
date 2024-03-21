

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import PrimeVue from 'primevue/config'

import Password from 'primevue/password';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import 'primevue/resources/themes/saga-blue/theme.css'; // 根据你选择的主题替换
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';



const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(PrimeVue)

app.component("Password", Password)
app.component("InputText", InputText)
app.component("Button", Button)


app.mount('#app')
