<template>
  <div class="create-customer-wrapper">
    <div class="header">
      <img src="/logo.png" alt="oatmeal studios logo" class="logo" />
      <div class="breadcrumb" @click="goHome" style="cursor:pointer">HOME</div>
    </div>
    <div class="form-section">
      <div class="section-title">CUSTOMER MAINTENANCE</div>
      <form @submit.prevent="submitCustomer">
        <div class="form-row">
          <label>Name:</label>
          <input v-model="customer.business_name" required />
        </div>
        <div class="form-row">
          <label>Contact Name:</label>
          <input v-model="customer.contact_name" />
        </div>
        <div class="form-row">
          <label>Email:</label>
          <input v-model="customer.email" type="email" />
        </div>
        <div class="form-row">
          <label>Phone:</label>
          <input v-model="customer.phone" />
        </div>
        <div class="form-row">
          <label>Address 1:</label>
          <input v-model="customer.address_1" />
        </div>
        <div class="form-row">
          <label>Address 2:</label>
          <input v-model="customer.address_2" />
        </div>
        <div class="form-row">
          <label>City:</label>
          <input v-model="customer.city" />
        </div>
        <div class="form-row">
          <label>State:</label>
          <select v-model="customer.state">
            <option value="">Select State</option>
            <option v-for="abbr in stateOptions" :key="abbr" :value="abbr">{{ abbr }}</option>
          </select>
        </div>
        <div class="form-row">
          <label>Zip Code:</label>
          <input v-model="customer.zip_code" />
        </div>
        <div class="form-row">
          <label>Country:</label>
          <input v-model="customer.country" maxlength="3" @input="onCountryInput" />
        </div>
        <div class="form-row">
          <label>Terms:</label>
          <select v-model="customer.terms">
            <option v-for="option in termsOptions" :key="option" :value="option">{{ option }}</option>
          </select>
        </div>
        <div class="form-row">
          <label>Discount %:</label>
          <input v-model.number="customer.discount" type="number" step="0.01" />
        </div>
        <div class="form-row">
          <label>Commission %:</label>
          <input v-model.number="customer.commission" type="number" step="0.01" />
        </div>
        <div class="form-row">
          <label>Notes:</label>
          <textarea v-model="customer.notes"></textarea>
        </div>
        <div class="form-row">
          <label>Free Shipping:</label>
          <input type="checkbox" v-model="customer.free_shipping" />
        </div>
        <button class="submit-btn" type="submit">ADD CUSTOMER</button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
const router = useRouter();
const goHome = () => router.push('/home');

const termsOptions = [
  'CREDIT CARD',
  'NET 30',
  'NET 60',
  'NET 90',
];

const stateOptions = [
  'AK', 'AL', 'AR', 'AZ', 'CA', 'CO', 'CT', 'DC', 'DE', 'FL',
  'GA', 'HI', 'IA', 'ID', 'IL', 'IN', 'KS', 'KY', 'LA', 'MA',
  'MD', 'ME', 'MI', 'MN', 'MO', 'MS', 'MT', 'NC', 'ND', 'NE',
  'NH', 'NJ', 'NM', 'NV', 'NY', 'OH', 'OK', 'OR', 'PA', 'RI',
  'SC', 'SD', 'TN', 'TX', 'UT', 'VA', 'VT', 'WA', 'WI', 'WV'
];

const customer = ref({
  business_name: '',
  contact_name: '',
  email: '',
  phone: '',
  address_1: '',
  address_2: '',
  city: '',
  state: '',
  zip_code: '',
  country: 'USA',
  terms: 'NET 30',
  discount: 0,
  commission: 20,
  notes: '',
  free_shipping: false
});

function onCountryInput(e) {
  // Only allow 3 characters
  if (e.target.value.length > 3) {
    customer.value.country = e.target.value.slice(0, 3);
  }
}

async function submitCustomer() {
  try {
    const res = await fetch('/api/customers', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(customer.value)
    });
    if (!res.ok) throw new Error('Failed to create customer');
    const data = await res.json();
    // Redirect to update page with new customer id
    router.push(`/customers/${data.id}/edit`);
  } catch (err) {
    alert('Failed to create customer');
  }
}
</script>

<style scoped>
.create-customer-wrapper {
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
.form-section {
  margin-top: 2rem;
  margin-left: 1rem;
  display: flex;
  flex-direction: column;
  gap: 0.7rem;
  width: 600px;
}
.section-title {
  font-size: 1.1rem;
  font-weight: bold;
  margin-bottom: 0.5rem;
}
.form-row {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
}
.form-row label {
  width: 140px;
  font-size: 0.95rem;
  font-weight: 500;
}
.form-row input,
.form-row textarea {
  flex: 1;
  padding: 2px 6px;
  font-size: 1rem;
}
.submit-btn {
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
  margin-top: 1rem;
}
.locations-section {
  margin-top: 2.5rem;
}
</style>
