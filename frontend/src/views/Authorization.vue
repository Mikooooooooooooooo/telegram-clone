<template>
    <div class="auth">
        <div class="icon">
        </div>
        <div class="edit">
            <BaseText> {{number}} </BaseText>
            <img class="icon" src="" />
        </div>
        <div class="input">
        </div>
    </div>
</template>

<script setup>
    import BaseText from '../components/BaseText.vue'
     
    const props = defineProps({
        number: String
    })
</script>
