<template>
    <div class="login-wrapper">
        <div class="header">
            <img src="/logo.png" alt="oatmeal studios logo" class="logo" />
            <div class="breadcrumb">HOME</div>
        </div>
        <div class="login-box">
            <div class="login-title">ADMIN LOGIN</div>
            <form class="login-form" @submit.prevent="onLogin">
                <div class="form-row">
                    <label for="username">Username:</label>
                    <input id="username" v-model="username" type="text" />
                </div>
                <div class="form-row">
                    <label for="password">Password:</label>
                    <input id="password" v-model="password" type="password" />
                </div>
                <button type="submit">Login</button>
            </form>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
const username = ref('');
const password = ref('');
const onLogin = () => {
    // TODO: Implement login logic
    alert(`Username: ${username.value}\nPassword: ${password.value}`);
};
</script>

<style scoped>
.login-wrapper {
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
    height: 85px
        /* min-height: 70px; */
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

.login-box {
    margin: 1.5rem 0 0 0.7rem;
    background: none;
    box-shadow: none;
    border: none;
    width: auto;
    min-width: 0;
    padding: 0;
}

.login-title {
    font-weight: bold;
    font-size: 1.1rem;
    margin-bottom: 1rem;
    color: #111;
    letter-spacing: 0.5px;
}

.login-form {
    display: flex;
    flex-direction: column;
    gap: 0.3rem;
    align-items: flex-start;
    margin-top: 0.2rem;
}

.form-row {
    display: flex;
    align-items: center;
    gap: 0.3rem;
    margin-bottom: 0.1rem;
}

label {
    color: #f00;
    font-size: 0.95rem;
    min-width: 70px;
    text-align: left;
    font-weight: normal;
    margin-right: 0.2rem;
}

input[type="text"],
input[type="password"] {
    width: 150px;
    padding: 2px 4px;
    font-size: 1rem;
    border: 1px solid #aaa;
    border-radius: 2px;
    background: #fff;
}

button[type="submit"] {
    width: 60px;
    margin-left: 70px;
    padding: 2px 0;
    font-size: 1rem;
    background: #eee;
    border: 1px solid #888;
    border-radius: 2px;
    cursor: pointer;
    margin-top: 0.1rem;
}
</style>
