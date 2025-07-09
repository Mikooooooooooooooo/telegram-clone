<template>
    <div class="input">
        <input :placeholder="text" v-model="value"/>
    </div>
</template>

<script setup>
    const props = defineProps({
        text: String
    })  

    const value = ""
</script>
