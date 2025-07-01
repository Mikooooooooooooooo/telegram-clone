<template>
    <div class="dropdown">
        <div class="title">
            {{title}} 
        </div>
        <div class="chosen" @click="showHandler" >
            {{chosen?.country}} 
        </div>
        <img v-if="isShow" @click="showHandler" src="../assets/arrow-up.svg" />
        <img v-else src="../assets/arrow-down.svg" /> 
        <div class="list" v-if="isShow">
            <div class="item" @click="setItem(option)" v-for="option in options">
                <div class="name">
                    {{option.country}}
                </div>
                <div class="code">
                    {{option.code}}
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
img {
    position: absolute ; 
    right: 15px ; 
    top: 17px ; 
    width: 16px ; 
    height: 16px ; 
}
.dropdown {
    position: relative ; 
    border-radius: 5px ; 
    border: 1px solid grey ; 
    transition: all .1s ease ; 
}
.title{ 
    font-size: 13px ; 
    position: absolute ; 
    color: var(--option-color) ;
    top: -10px ; 
    left: 15px ; 
    padding: 0 2px ; 
    background: var(--bg-color) ; 
    transition: all .1s ease ; 
}
.chosen{
    padding: 15px ; 
    color: white; 
    font-size: 16px ; 
}
.dropdown:hover{
    border: 1px solid var(--primary-color) ; 
}
.dropdown:hover .title{
    color: var(--primary-color) ; 
}
.list{
    position: absolute ; 
    width: 100% ; 
    max-height: 30vh ; 
    overflow-y: auto ; 
    background: var(--bg-color) ; 
    box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.2);
    top: calc(100% + 10px) ; 
    z-index: 2 ; 
}
.item{
    display: flex ; 
    padding: 15px ; 
    justify-content: space-between ; 
}
.item:hover{
    cursor: pointer ; 
    background: #303030 ;  
}
.name {
    color: white ; 
}
.code {
    color: var(--option-color) ; 
}

</style>

<script setup>
import { ref, onMounted } from 'vue'

const props = defineProps({
    title: String , 
    options: {
        type: Array , 
        required: true 
    }
})

const emit = defineEmits(['setCode'])

const chosen = ref(null)
const isShow = ref(false)

function showHandler() {
   isShow.value = !isShow.value 
}

function setItem(item) { 
    chosen.value = item ; 
    emit('setCode' , chosen.value)
    isShow.value = false ; 
}

onMounted(async () => {
  if (props.options.length > 0) {
    chosen.value = props.options[0]
  } 
})
</script>
