<template>
    <div class="main">
        <div class="slide">
            <div class="slide_header">
                <div class="menu">
                    <img src="@/assets/menu.png"/>
                </div>
                <div class="search">
                    <SearchInput/>
                </div>
            </div>
            <div class="list">
            </div>
        </div>
        <div class="chat">
            <div class="chat_header">
            </div>
            <div class="chat_content">
            </div>
        </div>
    </div>
</template>

<style scoped>
.main{
    display: flex ; 
    height: 100vh ; 
}
.slide{
    padding-top: 10px ; 
    width: 450px ; 
}
.slide_header{
    display: flex ; 
    align-items: center ; 
    width: 400px ; 
    margin: 0 auto ; 
    gap: 20px ; 
}
.search{
    width: 100% ; 
}

</style>

<script setup>
import SearchInput from "@/components/SearchInput.vue"

</script>
