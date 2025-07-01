import {createRouter , createWebHistory} from 'vue-router' 
import LoginPage from '../views/LoginPage.vue' 
import AuthPage from '../views/Authorization.vue'

const routes = [
    {
        path: '/' , 
        name: 'Login' , 
        component: LoginPage
    } , 
    {
        path: '/auth' , 
        name: 'Authorization' , 
        component: AuthPage
    }
]

const router = createRouter({
    history: createWebHistory() , 
    routes
})

export default router
