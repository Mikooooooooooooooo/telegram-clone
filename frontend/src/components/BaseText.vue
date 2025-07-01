<template> 
    <p :class="['base-text' , type]">
        <slot/> 
    </p>
</template>

<script setup> 
    defineProps(
        {type: String}
    )
</script>

<style scoped>
.base-text {
    font-size: 1rem ; 
    color: var(--text-color) ; 
    text-align: center ; 
}
.font-main{
    font-size: 2rem ; 
    font-weight: bold ; 
    margin-bottom: 2vh ; 
    line-height: 110% ; 
}
.title{ 
    font-weight: bold ; 
    font-size: 1.5rem ; 
}
.subtitle{
    color: var(--option-color) ;  
    margin: 0 auto ;
    width: 250px ; 
}
</style>
