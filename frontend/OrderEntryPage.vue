<template>
    <div class="order-entry-wrapper">
        <div class="header">
            <img src="/logo.png" alt="oatmeal studios logo" class="logo" />
            <div class="breadcrumb" @click="goHome" style="cursor:pointer">HOME</div>
        </div>
        <div class="order-entry-content">
            <div class="order-entry-title">ORDER ENTRY</div>
            <div class="order-entry-form">
                <div class="customer-row">
                    <label for="customer-number">Customer #:</label>
                    <input id="customer-number" type="text" class="customer-input" />
                    <select class="customer-select">
                        <option>Select by name</option>
                        <option v-for="customer in customers" :key="customer.id" :value="customer.id">
                            {{ customer.business_name }}
                        </option>
                    </select>
                </div>
                <button class="create-order-btn">CREATE NEW ORDER</button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const goHome = () => {
    router.push('/home');
};

// 1. Reactive variable for customers
const customers = ref([]);

// 2. Fetch customers on mount
onMounted(async () => {
    const res = await fetch('/api/customers');
    if (res.ok) {
        customers.value = await res.json();
    }
});
</script>

<style scoped>
.order-entry-wrapper {
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

.order-entry-content {
    margin-top: 1.5rem;
    margin-left: 1.5rem;
}

.order-entry-title {
    font-size: 1.2rem;
    font-weight: bold;
    margin-bottom: 1.2rem;
}

.order-entry-form {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 0.7rem;
}

.customer-row {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 1rem;
}

.customer-input {
    width: 140px;
    height: 1.3em;
    font-size: 1.2rem;
    padding: 0 0.5em;
    margin-right: 0.5rem;
    border: 1px solid #bdbdbd;
    border-radius: 4px;
}

.customer-select {
    width: 700px;
    height: 1.3em;
    font-size: 1.2rem;
    padding: 0 0.5em;
    border: 1px solid #bdbdbd;
    border-radius: 4px;
    background: #f5f5f5;
    margin-right: 0.5rem;
}

.order-btn-row {
    margin-top: 0.7rem;
    width: 100%;
    display: flex;
    justify-content: flex-start;
}

.create-order-btn {
    background: #eee;
    border: 1px solid #888;
    border-radius: 2px;
    padding: 2px 12px;
    font-size: 1rem;
    cursor: pointer;
    font-family: sans-serif;
    font-weight: normal;
}
</style>
