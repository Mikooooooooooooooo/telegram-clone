import {createRouter , createWebHistory} from 'vue-router' 
import PhonePage from '../views/PhonePage.vue' 
import LoginPage from '../views/LoginPage.vue'
import RegisterPage from '../views/RegisterPage.vue'
import MainPage from '../views/MainPage.vue'

const routes = [
    {
        path: '/' , 
        name: 'Phone' , 
        component: PhonePage 
    } , 
    {
        path: '/login' , 
        name: 'Login' , 
        component: LoginPage 
    } , 
    {
        path: '/register' , 
        name: 'Register' , 
        component : RegisterPage , 
    } ,
    {
        path: '/main' , 
        name: 'main' , 
        component: MainPage , 
    }
]

const router = createRouter({
    history: createWebHistory() , 
    routes
})

export default router
