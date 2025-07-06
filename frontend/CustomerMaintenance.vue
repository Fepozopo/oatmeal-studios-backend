<template>
    <div class="customer-maint-wrapper">
        <div class="header">
            <img src="/logo.png" alt="oatmeal studios logo" class="logo" />
            <div class="breadcrumb" @click="goHome" style="cursor:pointer">HOME</div>
        </div>
        <div class="customer-maint-content">
            <div class="customer-maint-title">CUSTOMER MAINTENANCE</div>
            <button class="customer-btn" @click="goToCreateCustomer">CREATE NEW CUSTOMER</button>
            <div class="customer-edit-section">
                <div class="customer-edit-label">OR EDIT EXISTING CUSTOMER</div>
                <div class="customer-edit-row">
                    <span class="customer-edit-label">Customer #:</span>
                    <input class="customer-input" type="text" v-model="customerNumber" />
                    <select class="customer-select" v-model="selectedCustomer">
                        <option value="" disabled>Select by name</option>
                        <option v-for="customer in customers" :key="customer.id" :value="customer.id">
                            {{ customer.business_name }} - {{ customer.id }}
                        </option>
                    </select>
                </div>
                <button class="customer-btn" style="margin-top: 0.5rem;" @click="editCustomer">EDIT CUSTOMER</button>
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

const customerNumber = ref('');
const selectedCustomer = ref('');
const customers = ref([]);

const goToCreateCustomer = () => {
    router.push('/customers/new');
};

const fetchCustomers = async () => {
    try {
        const res = await fetch('/api/customers');
        if (!res.ok) throw new Error('Failed to fetch customers');
        const data = await res.json();
        // Sort by business_name (case-insensitive)
        customers.value = data.sort((a, b) =>
            a.business_name.localeCompare(b.business_name, undefined, { sensitivity: 'base' })
        );
    } catch (err) {
        customers.value = [];
    }
};

onMounted(fetchCustomers);

const editCustomer = () => {
    const id = customerNumber.value || selectedCustomer.value;
    if (id) {
        router.push(`/customers/${id}/edit`);
    } else {
        alert('Please enter or select a customer.');
    }
};
</script>

<style scoped>
.customer-maint-wrapper {
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

.customer-maint-content {
    margin-top: 2rem;
    margin-left: 1rem;
    display: flex;
    flex-direction: column;
    gap: 0.7rem;
    width: 600px;
}

.customer-maint-title {
    font-size: 1.1rem;
    font-weight: bold;
    margin-bottom: 0.5rem;
}

.customer-btn {
    width: 170px;
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

.customer-edit-section {
    margin-top: 1.5rem;
}

.customer-edit-label {
    font-size: 0.95rem;
    margin-right: 0.5rem;
    white-space: nowrap;
}

.customer-edit-row {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    margin-top: 0.5rem;
}

.customer-input {
    width: 100px;
    padding: 2px 6px;
    font-size: 1rem;
}

.customer-select {
    width: 700px;
    padding: 2px 6px;
    font-size: 1rem;
}
</style>
