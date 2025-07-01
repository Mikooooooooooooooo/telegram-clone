<template>
    <div class="button"> 
        {{props.word}}
    </div>
</template>

<style scoped>
.button{
    width: 100% ; 
    height: 5.5vh ; 
    text-transform: uppercase ;
    background: var(--primary-color) ; 
    display: flex ; 
    align-items: center ; 
    border: 1px solid var(--primary-color) ; 
    justify-content: center ; 
    color: white ; 
    font-size: 16px ; 
    font-weight: 600 ; 
    border-radius: 10px ; 
    cursor: pointer ; 
    transition: .1s all ease ; 
}
.button:hover{
    background: #7b67db ; 
}
</style>

<script setup>
    
    const props = defineProps({
        word: String 
    })
</script>

