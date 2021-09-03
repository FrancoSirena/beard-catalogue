import * as Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import App from './App.vue'

const client = axios.create({
  baseURL: "http://localhost:8080/api/v1",
});

const app = Vue.createApp(App)
app.use(VueAxios, client);
app.mount('#app')


