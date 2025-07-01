<template> 
    <div class="login">
        <div class="container">
            <telegram-logo></telegram-logo> 
            <base-text type="font-main">Sign in to Telegram</base-text>
            <base-text type="subtitle">Please confirm your country code and enter your phone number.</base-text>
            <dropdown @setCode="updateCode" class="dropdown" title="Country" :options="countries" />  
            <input-number @getNumber="changeNumber" :code="chosen?.code" class="number"/>
            <app-button @click="toAuth" word="next"/>
        </div>
    </div>
</template>

<style scoped>
.container{
    width: 360px ; 
    margin: 0 auto ; 
    padding-top: 7vh ; 
}
.dropdown{ 
    margin-top: 5vh ; 
    margin-bottom: 3vh ; 
}
.number{
    margin-bottom: 1.5vh ; 
}
</style>

<script>
import TelegramLogo from '../components/TelegramLogo.vue'
import BaseText from '../components/BaseText.vue'
import Dropdown from '../components/Dropdown.vue'
import InputNumber from '../components/InputNumber.vue'
import Button from '../components/Button.vue'

export default {
    components: {
        'telegram-logo': TelegramLogo , 
        'dropdown': Dropdown , 
        'base-text': BaseText , 
        'input-number': InputNumber , 
        'app-button' : Button
    } , 
    data() {
        return {
            chosen: {country: 'Kazakstan' , code: '+7'} , 
            countries:[
                        { country: 'Kazakhstan', code: '+7' },
                        { country: 'United States', code: '+1' },
                        { country: 'United Kingdom', code: '+44' },
                        { country: 'Germany', code: '+49' },
                        { country: 'France', code: '+33' },
                        { country: 'Japan', code: '+81' },
                        { country: 'South Korea', code: '+82' },
                        { country: 'India', code: '+91' },
                        { country: 'Brazil', code: '+55' },
                        { country: 'Turkey', code: '+90' },
                        { country: 'Australia', code: '+61' }
                     ] , 
            number: null , 

        }
    } , 
    methods: {
        updateCode(value){ 
            console.log(value)
            this.chosen = value
        } , 
        changeNumber(value){
            console.log(value) 
            this.number = value
        } , 
        toAuth() {
            this.$router.push('/auth')
        }
    }
    
}
</script>
