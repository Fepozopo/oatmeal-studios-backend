<template>
    <div class="user-maint-wrapper">
        <div class="header">
            <img src="/logo.png" alt="oatmeal studios logo" class="logo" />
            <div class="breadcrumb" @click="goHome" style="cursor:pointer">HOME</div>
        </div>
        <div class="user-maint-content">
            <div class="user-maint-title">USER MAINTENANCE</div>
            <button class="user-btn">CREATE NEW USER</button>
            <div class="user-edit-label">OR EDIT EXISTING USER</div>
            <select class="user-select">
                <option>Existing users</option>
            </select>
            <button class="user-btn">CONTINUE</button>
        </div>
    </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
const router = useRouter();
const goHome = () => {
    router.push('/home');
};
</script>

<style scoped>
.user-maint-wrapper {
    min-height: 100vh;
    background: #dbdbdb;
    margin: 0;
    padding: 0;
}

.header {
    background: #ffd16a;
    padding: 0.5rem 0 0.5rem 0.5rem;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    width: 100vw;
    box-sizing: border-box;
    height: 85px;
}

.logo {
    height: 64px;
    margin-bottom: 0;
    margin-right: 0;
}

.breadcrumb {
    font-size: 0.85rem;
    color: #fffefe;
    margin-left: 0.1rem;
    margin-top: 0;
    margin-bottom: 0;
    font-family: sans-serif;
    font-weight: 500;
    letter-spacing: 0.5px;
}

.user-maint-content {
    margin-top: 2rem;
    margin-left: 1rem;
    display: flex;
    flex-direction: column;
    gap: 0.7rem;
    width: 320px;
}

.user-maint-title {
    font-size: 1.1rem;
    font-weight: bold;
    margin-bottom: 0.5rem;
}

.user-btn {
    width: 150px;
    text-align: left;
    padding: 4px 10px;
    font-size: 1rem;
    background: #eee;
    border: 1px solid #888;
    border-radius: 2px;
    cursor: pointer;
    font-family: sans-serif;
    font-weight: normal;
    box-sizing: border-box;
}

.user-edit-label {
    font-size: 0.95rem;
    margin-top: 0.5rem;
}

.user-select {
    width: 180px;
    padding: 2px 6px;
    font-size: 1rem;
    margin-bottom: 0.2rem;
}
</style>
