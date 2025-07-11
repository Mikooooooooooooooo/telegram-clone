<template>
    <div class="input">
        <input :placeholder="text" :disabled="disabled" @keyup.enter="returnCode" v-model="value"/>
    </div>
</template>

<style scoped>
input{
    outline: none ; 
    width: calc(100% - 32px) ; 
    border: 1px solid var(--border-color) ; 
    background: var(--bg-color) ; 
    padding: 15px ; 
    color: var(--option-color) ; 
    font-size: 16px ; 
    border-radius: 10px ; 
    transition: all ease .2s ; 
}
input:hover{
    color: var(--primary-color) ; 
    border-color: var(--primary-color) ;  
}
</style>

<script setup>
    import {ref} from 'vue'
    const props = defineProps({
        text: String , 
        disabled: Boolean
    })  

    const value = ref("")

    const emit = defineEmits(['getValue'])

    function returnCode() {
        emit('getValue' , value.value)
    } 
</script>
