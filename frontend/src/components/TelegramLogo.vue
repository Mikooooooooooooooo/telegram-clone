<template>
    <div :class="['telegram_logo' , type]">
        <img src="../assets/telegram_logo.svg"/>
    </div>
</template>

<script setup>
    const props = defineProps({
        type: String 
    })
</script>

<style scoped>

    .telegram_logo{
        display: flex ; 
        align-items: center ; 
        justify-content: center ; 
    }
    img {
        width: 160px ; 
        height: 160px ; 
    }
    .m img{
        width: 120px ; 
        height: 120px ; 
    }
</style>

