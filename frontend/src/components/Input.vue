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
import { ref, watch } from 'vue'

const props = defineProps({
  modelValue: String,
  text: String,
  disabled: Boolean
})

const emit = defineEmits(['update:modelValue'])

function onInput(event) {
  emit('update:modelValue', event.target.value)
}
</script>

<template>
  <div class="input">
    <input
      :placeholder="text"
      :disabled="disabled"
      :value="modelValue"
      @input="onInput"
    />
  </div>
</template>
