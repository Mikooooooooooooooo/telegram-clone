<template>
    <div class="input"> 
        <img src="@/assets/search-interface-symbol.png" />
        <input placeholder="Search" type="text"/>
    </div>
</template>

<style scoped>
.input {
    display: flex ; 
    position: relative ; 
}
input{
    width: 100% ; 
    background: #181818 ; 
    outline: none ;
    font-size: 18px ; 
    padding: 10px ; 
    padding-left: 40px ; 
    border: solid 1px #aaa ; 
    border-radius: 20px ; 
    color: #aaa ; 
    font-weight: 500 ; 
}
img{
    width: 20px ; 
    height: 20px ; 
    position: absolute ; 
    top: 12px ;
    left: 12px ; 
}
</style>

<script setup>

</script>
