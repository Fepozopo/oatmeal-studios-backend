<template>
  <div class="update-customer-wrapper">
    <div class="header">
      <img src="/logo.png" alt="oatmeal studios logo" class="logo" />
      <div class="breadcrumb" @click="goHome" style="cursor:pointer">HOME</div>
    </div>
    <div class="form-section">
      <div class="section-title">
        CUSTOMER MAINTENANCE <span v-if="customer">- #{{ customer.id }}</span>
      </div>
      <form v-if="customer" @submit.prevent="updateCustomer">
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
          <input v-model="customer.state" />
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
        <button class="submit-btn" type="submit">UPDATE CUSTOMER</button>
      </form>
    </div>

    <div class="form-section locations-section">
      <div class="section-title">LOCATIONS</div>
      <div v-if="locations.length">
        <div class="form-row">
          <label>Choose Location:</label>
          <select v-model="selectedLocationId" @change="onLocationSelect">
            <option value="">New Location</option>
            <option v-for="loc in locations" :key="loc.id" :value="loc.id">
              {{ loc.location_num ? ('#' + loc.location_num + ' - ') : '' }}{{ loc.business_name }}
            </option>
          </select>
        </div>
      </div>
      <form @submit.prevent="saveLocation">
        <div class="form-row">
          <label>Location Number:</label>
          <input v-model="location.location_num" type="number" />
        </div>
        <div class="form-row">
          <label>Location Name:</label>
          <input v-model="location.business_name" />
        </div>
        <div class="form-row">
          <label>Contact Name:</label>
          <input v-model="location.contact_name" />
        </div>
        <div class="form-row">
          <label>Phone:</label>
          <input v-model="location.phone" />
        </div>
        <div class="form-row">
          <label>Address 1:</label>
          <input v-model="location.address_1" />
        </div>
        <div class="form-row">
          <label>Address 2:</label>
          <input v-model="location.address_2" />
        </div>
        <div class="form-row">
          <label>City:</label>
          <input v-model="location.city" />
        </div>
        <div class="form-row">
          <label>State:</label>
          <input v-model="location.state" />
        </div>
        <div class="form-row">
          <label>Zip Code:</label>
          <input v-model="location.zip_code" />
        </div>
        <div class="form-row">
          <label>Country:</label>
          <input v-model="location.country" maxlength="3" @input="onLocationCountryInput" />
        </div>
        <div class="form-row">
          <label>Notes:</label>
          <textarea v-model="location.notes"></textarea>
        </div>
        <button class="submit-btn" type="submit">{{ selectedLocationId ? 'UPDATE' : 'ADD' }} LOCATION</button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router';
const router = useRouter();
const route = useRoute();
const goHome = () => router.push('/home');

const termsOptions = [
  'NET 30',
  'NET 60',
  'NET 90',
  'PREPAID',
  'COD',
  'CREDIT CARD',
  'EOM',
  'DUE ON RECEIPT'
];

const customer = ref(null);
const locations = ref([]);
const selectedLocationId = ref('');
const location = ref({
  location_num: '',
  business_name: '',
  contact_name: '',
  phone: '',
  address_1: '',
  address_2: '',
  city: '',
  state: '',
  zip_code: '',
  country: 'USA',
  notes: ''
});

const fetchCustomer = async () => {
  const id = route.params.id;
  const res = await fetch(`/api/customers/${id}`);
  if (res.ok) customer.value = await res.json();
};
const fetchLocations = async () => {
  const id = route.params.id;
  const res = await fetch(`/api/customers/${id}/locations`);
  if (res.ok) locations.value = await res.json();
};

onMounted(async () => {
  await fetchCustomer();
  await fetchLocations();
});

watch(selectedLocationId, (val) => {
  if (!val) {
    location.value = {
      location_num: '',
      business_name: '',
      contact_name: '',
      phone: '',
      address_1: '',
      address_2: '',
      city: '',
      state: '',
      zip_code: '',
      country: 'USA',
      notes: ''
    };
  } else {
    const loc = locations.value.find(l => l.id == val);
    if (loc) location.value = { ...loc };
  }
});

function onCountryInput(e) {
  if (e.target.value.length > 3) customer.value.country = e.target.value.slice(0, 3);
}
function onLocationCountryInput(e) {
  if (e.target.value.length > 3) location.value.country = e.target.value.slice(0, 3);
}

async function updateCustomer() {
  const id = route.params.id;
  await fetch(`/api/customers/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(customer.value)
  });
  await fetchCustomer();
}
async function saveLocation() {
  const id = route.params.id;
  if (selectedLocationId.value) {
    // update
    await fetch(`/api/customers/${id}/locations/${selectedLocationId.value}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(location.value)
    });
  } else {
    // create
    await fetch(`/api/customers/${id}/locations`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(location.value)
    });
  }
  await fetchLocations();
  selectedLocationId.value = '';
}
</script>

<style scoped>
.update-customer-wrapper {
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
.form-row textarea,
.form-row select {
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
