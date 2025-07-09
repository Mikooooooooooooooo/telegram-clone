<template>
    <div class="auth">
        <div class="container">
            <Logo class="logo" type="m"/>
            <div class="edit">
                <BaseText type="font-main"> {{ phone }} </BaseText>
                <div class="icon_container">
                   <img class="icon" src="@/assets/edit.png" />
                </div>
            </div>
            <BaseText type="subtitle">Enter your telegram login code</BaseText>
            <div class="input">
            </div>
        </div>
    </div>
</template>

<style scoped>
.container{
    padding-top: 10vh ; 
    width: 500px ; 
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
</style>

<script setup>
    import BaseText from '@/components/BaseText.vue'
    import Logo from '@/components/TelegramLogo.vue' 
    import {useRoute} from 'vue-router'

    const route = useRoute() 
    const phone = route.query.phone 
</script>
