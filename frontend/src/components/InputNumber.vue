<template>
    <div class="number">
        <div class="code">
            {{code}}
        </div>
        <input @change="setNumber" v-model="phone" :style="{ paddingLeft: padd }" type="tel" placeholder="_ _ _  _ _ _  _ _  _ _" class="number-input"/>
        <div class="title">
            Phone Number
        </div>
    </div>
</template>

<style scoped>
.number {
    display: flex ;  
    position: relative ; 
    height: 60px ; 
}
.number-input{
    position: absolute ; 
    padding: 15px ; 
    width: calc(100% - 55px); 
    top: 0 ; 
    left: 0 ; 
    background: var(--bg-color) ; 
    outline: none ; 
    color: white ; 
    font-size: 16px ; 
    border: 1px solid grey ;  
    border-radius: 5px ; 
}
.number-input:hover{
    border: 1px solid var(--primary-color) ; 
}
.code{
    position: absolute ; 
    z-index: 1 ; 
    top: 16px ; 
    left: 15px ; 
    color: white ; 
}
.title{
    position: absolute ; 
    background: var(--bg-color) ; 
    top: -10px ; 
    left: 15px ; 
    color: var(--option-color) ; 
    font-size: 13px; 
}
.number-input:hover + .title{
    color: var(--primary-color) ; 
}

</style>
 
<script setup>
    import {ref , computed } from 'vue'
    
    const phone = ref('') 
    
    const padd = computed(() => {
        return (40 + (props.code.length -2) * 10) + 'px'
    })

    const props = defineProps({
        code: String 
    })

    const emit = defineEmits(['getNumber'])

    
    function setNumber() {
        const clean = props.code + phone.value.replace(/\s+/g, '')
        emit('getNumber' , clean) 
    }
</script>
