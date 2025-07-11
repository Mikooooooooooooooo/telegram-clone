<template>
    <div class="auth">
        <div class="container">
            <Logo class="logo" type="m"/>
            <div class="edit">
                <BaseText type="font-main"> {{ formatPhone(phone) }}</BaseText>
                <div class="icon_container">
                   <img class="icon" src="@/assets/edit.png" />
                </div>
            </div>
            <BaseText type="subtitle">Enter your telegram login code</BaseText>
            <Input @getValue="getPass" text="Code" class="input" :disabled="isDisabled"/>
            <div class="error" v-if="error">{{error}}</div>
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
.error {
  color: #8f3838; 
  font-size: 14px;
  margin-top: 8px;
}
</style>

<script setup>
    import {ref} from 'vue'
    import BaseText from '@/components/BaseText.vue'
    import Logo from '@/components/TelegramLogo.vue' 
    import Input from '@/components/Input.vue' 
    import {useRoute} from 'vue-router'
    import api from '@/api'

    const route = useRoute() 
    const phone = route.query.phone 
    const error = ref(null)
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

    function getPass(code){
        isDisabled.value = true 
        api.post('/login' , {phone: phone , password: code})
            .then(res => {
                console.log("Login success")
            })
            .catch(err => {
                error.value = "Wrong code. Try again"
                console.error("Wrong password")
                setTimeout(() => {
                    error.value = null
                }, 3000)
            })
            .finally(() => {
                isDisabled.value = false
            })
    }

</script>
