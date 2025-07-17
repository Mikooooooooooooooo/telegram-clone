<template>
    <div class="register">
        <div class="container">
            <Logo class="logo" type="m"/>
            <div class="edit">
                <BaseText type="font-main"> {{ formatPhone(phone) }}</BaseText>
                <div class="icon_container">
                   <img @click="goBack" class="icon" src="@/assets/edit.png" />
                </div>
            </div>
            <BaseText type="subtitle">Create your telegram login code</BaseText>
            <Input v-model="code" :disabled="isDisabled" text="Code" class="input" />
            <Input v-model="confirm" :disabled="isDisabled" text="Confirm" class="confirm" />
            <Button word="Register" @click="register"/>
        </div>
    </div>
</template>

<style scoped>
.container{
    text-align: center ; 
    padding-top: 10vh ; 
    width: 350px ; 
    margin: 0 auto ; 
}
.logo{
    margin: 5vh 0 ; 
}
img{
    width: 24px ; 
    height: 24px; 
}
.edit{
    display: flex ; 
    align-items: center ; 
    justify-content: center ; 
    gap: 10px ; 
}
.icon_container{
    height: 35.2px ; 
}
.input{
    margin-top: 5vh ; 
    width: 100% ; 
}
.confirm{
    margin: 2vh 0; 
}

</style>

<script setup>
    import {ref} from 'vue'
    import BaseText from '@/components/BaseText.vue'
    import Logo from '@/components/TelegramLogo.vue' 
    import Input from '@/components/Input.vue' 
    import Button from '@/components/Button.vue'
    import api from '@/api'
    import { useRoute, useRouter } from 'vue-router'

    const route = useRoute() 
    const router = useRouter() 
    const phone = route.query.phone 
    const code = ref("")
    const confirm = ref("")
    const isDisabled = ref(false)

    
    function formatPhone(raw) {
      const digits = raw.replace(/\D/g, '')
      if (!digits.startsWith('7')) return raw
    
      const parts = [
        '+7',
        digits.slice(1, 4),
        digits.slice(4, 7),
        digits.slice(7, 9),
        digits.slice(9)
      ].filter(Boolean)
    
      return parts.join(' ')
    }
    
    function register() {
        if (code.value !== confirm.value) {
            alert("Passwords do not match")
            return
        }

        isDisabled.value = true 

        api.post('/register', {
            phone: phone,
            password: code.value
        })
        .then(res => {
            console.log("Registration successful")
            localStorage.setItem('token' , res.data.token)
            router.push('/main')
        })
        .catch(err => {
            console.error("Registration failed", err)
        })
        .finally(() => {
            isDisabled.value = false 
        })
    }

    function goBack(){
        router.push('/')
    }

</script>
