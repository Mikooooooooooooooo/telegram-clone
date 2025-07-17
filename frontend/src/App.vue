<template>
    <div id="telegram">
        <router-view></router-view>
    </div>
</template> 

<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/api'

const router = useRouter()


onMounted(async () => {
  const token = localStorage.getItem('token')
  if (token) {
    try {
      const { data } = await api.get('/me')
      console.log('Welcome back', data.phone)
      router.push('/main')
    } catch (err) {
      console.error('Token invalid:', err)
      localStorage.removeItem('token')
    }
  }
})

</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Roboto:ital,wght@0,100..900;1,100..900&display=swap');
*{
    padding: 0 ; 
    margin: 0 ; 
    font-family: "Roboto", sans-serif;
}
#app{
    background-color: var(--bg-color) ; 
    min-height: 100vh ; 
}
</style>
