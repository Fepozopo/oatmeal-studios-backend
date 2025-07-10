<template>
  <div class="update-customer-wrapper">
    <div v-if="successMessage" class="success-message">{{ successMessage }}</div>
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
              {{ loc.location_num && loc.business_name ? loc.location_num + ' - ' + loc.business_name : loc.business_name || loc.location_num || 'Location' }}
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
          <select v-model="location.state">
            <option value="">Select State</option>
            <option v-for="abbr in stateOptions" :key="abbr" :value="abbr">{{ abbr }}</option>
          </select>
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
          <label>Sales Rep:</label>
          <select v-model="location.sales_rep">
            <option value="">Select Sales Rep</option>
            <option v-for="rep in salesReps" :key="rep.id" :value="rep.code">
              {{ rep.code && (rep.first_name || rep.last_name) ? rep.code + ' - ' + [rep.first_name, rep.last_name].filter(Boolean).join(' ') : rep.code }}
            </option>
          </select>
        </div>
        <div class="form-row">
          <label>Notes:</label>
          <textarea v-model="location.notes"></textarea>
        </div>
        <div class="form-row">
          <label>Planogram:</label>
          <select v-model="selectedPlanogramId">
            <option value="">None</option>
            <option v-for="plan in planograms" :key="plan.id" :value="plan.id">
              {{ plan.name }} ({{ plan.num_pockets }} pockets)
            </option>
          </select>
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

const stateOptions = [
  'AK', 'AL', 'AR', 'AZ', 'CA', 'CO', 'CT', 'DC', 'DE', 'FL',
  'GA', 'HI', 'IA', 'ID', 'IL', 'IN', 'KS', 'KY', 'LA', 'MA',
  'MD', 'ME', 'MI', 'MN', 'MO', 'MS', 'MT', 'NC', 'ND', 'NE',
  'NH', 'NJ', 'NM', 'NV', 'NY', 'OH', 'OK', 'OR', 'PA', 'RI',
  'SC', 'SD', 'TN', 'TX', 'UT', 'VA', 'VT', 'WA', 'WI', 'WV'
];

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
  sales_rep: '',
  notes: ''
});
const salesReps = ref([]);
const planograms = ref([]);
const selectedPlanogramId = ref("");
// Fetch all planograms for dropdown
const fetchPlanograms = async () => {
  const res = await fetch('/api/planograms');
  if (res.ok) {
    const data = await res.json();
    planograms.value = data;
  }
};

// Fetch assigned planogram for a location
const fetchAssignedPlanogram = async (locationId) => {
  if (!locationId) {
    selectedPlanogramId.value = "";
    return;
  }
  const res = await fetch(`/api/planograms/${locationId}/planograms`);
  if (res.ok) {
    const data = await res.json();
    if (data && data.id) {
      selectedPlanogramId.value = data.id;
    } else {
      selectedPlanogramId.value = "";
    }
  } else {
    selectedPlanogramId.value = "";
  }
};
// Fetch sales reps for dropdown
const fetchSalesReps = async () => {
  const res = await fetch('/api/sales-reps');
  if (res.ok) {
    const data = await res.json();
    // Normalize code and name fields
    salesReps.value = data.map(rep => ({
      id: rep.id,
      code: rep.code || rep.rep_code || '',
      first_name: rep.first_name || '',
      last_name: rep.last_name || ''
    }));
  }
};
const successMessage = ref('');
function showSuccess(msg) {
  successMessage.value = msg;
  setTimeout(() => {
    successMessage.value = '';
  }, 2500);
}

const fetchCustomer = async () => {
  const id = route.params.id;
  const res = await fetch(`/api/customers/${id}`);
  if (res.ok) {
    const data = await res.json();
    // Normalize fields to ensure they are strings or numbers, not objects
    const normalize = (val) => {
      if (val === null || val === undefined) return '';
      if (typeof val === 'object') {
        // Handle Go sql.NullString pattern
        if ('String' in val && 'Valid' in val) return val.Valid ? val.String : '';
        if ('value' in val) return val.value;
        if ('text' in val) return val.text;
        // fallback: show JSON string for debugging
        try {
          return JSON.stringify(val);
        } catch {
          return '';
        }
      }
      return val;
    };
    customer.value = {
      ...data,
      business_name: normalize(data.business_name),
      contact_name: normalize(data.contact_name),
      email: normalize(data.email),
      phone: normalize(data.phone),
      address_1: normalize(data.address_1),
      address_2: normalize(data.address_2),
      city: normalize(data.city),
      state: normalize(data.state),
      zip_code: normalize(data.zip_code),
      country: normalize(data.country),
      notes: normalize(data.notes),
      terms: normalize(data.terms),
      discount: data.discount ?? 0,
      commission: data.commission ?? 0,
      free_shipping: !!data.free_shipping
    };
  }
};
const fetchLocations = async () => {
  const id = route.params.id;
  const res = await fetch(`/api/customers/${id}/locations`);
  if (res.ok) {
    const data = await res.json();
    const normalize = (val) => {
      if (val === null || val === undefined) return '';
      if (typeof val === 'object') {
        if ('String' in val && 'Valid' in val) return val.Valid ? val.String : '';
        if ('value' in val) return val.value;
        if ('text' in val) return val.text;
        try {
          return JSON.stringify(val);
        } catch {
          return '';
        }
      }
      return val;
    };
    locations.value = data.map(loc => ({
      ...loc,
      location_num: (typeof loc.location_num === 'object' && 'Int32' in loc.location_num && 'Valid' in loc.location_num)
        ? (loc.location_num.Valid ? loc.location_num.Int32 : '')
        : normalize(loc.location_num),
      business_name: normalize(loc.business_name),
      contact_name: normalize(loc.contact_name),
      phone: normalize(loc.phone),
      address_1: normalize(loc.address_1),
      address_2: normalize(loc.address_2),
      city: normalize(loc.city),
      state: normalize(loc.state),
      zip_code: normalize(loc.zip_code),
      country: normalize(loc.country),
      notes: normalize(loc.notes)
    })).sort((a, b) => {
      const numA = Number(a.location_num);
      const numB = Number(b.location_num);
      if (!isNaN(numA) && !isNaN(numB)) return numA - numB;
      return String(a.location_num).localeCompare(String(b.location_num));
    });
  }
};

onMounted(async () => {
  await fetchCustomer();
  await fetchLocations();
  await fetchSalesReps();
  await fetchPlanograms();
});

watch(selectedLocationId, async (val) => {
  const normalize = (val, key) => {
    if (val === null || val === undefined) return '';
    if (key === 'location_num' && typeof val === 'object' && 'Int32' in val && 'Valid' in val) {
      return val.Valid ? val.Int32 : '';
    }
    if (typeof val === 'object') {
      if ('String' in val && 'Valid' in val) return val.Valid ? val.String : '';
      if ('value' in val) return val.value;
      if ('text' in val) return val.text;
      try {
        return JSON.stringify(val);
      } catch {
        return '';
      }
    }
    return val;
  };
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
      sales_rep: '',
      notes: ''
    };
    selectedPlanogramId.value = "";
  } else {
    const loc = locations.value.find(l => l.id == val);
    if (loc) {
      location.value = {
        ...loc,
        location_num: normalize(loc.location_num, 'location_num'),
        business_name: normalize(loc.business_name),
        contact_name: normalize(loc.contact_name),
        phone: normalize(loc.phone),
        address_1: normalize(loc.address_1),
        address_2: normalize(loc.address_2),
        city: normalize(loc.city),
        state: normalize(loc.state),
        zip_code: normalize(loc.zip_code),
        country: normalize(loc.country),
        sales_rep: normalize(loc.sales_rep),
        notes: normalize(loc.notes)
      };
      await fetchAssignedPlanogram(loc.id);
    }
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
  console.log('Attempting to update customer with id:', id);
  if (!id || isNaN(Number(id))) {
    alert('Invalid customer ID. Cannot update.');
    return;
  }
  const updatedCustomer = { ...customer.value };
  try {
    const res = await fetch(`/api/customers/${id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(updatedCustomer)
    });
    if (!res.ok) {
      const msg = await res.text();
      alert('Failed to update customer: ' + msg);
      return;
    }
    // Only refresh if update succeeded
    await fetchCustomer();
    showSuccess('Customer updated successfully!');
  } catch (err) {
    alert('Failed to update customer: ' + (err?.message || err));
  }
}
async function saveLocation() {
  const id = route.params.id;
  let isUpdate = false;
  // Ensure sales_rep is a string or undefined
  if (location.value.sales_rep && typeof location.value.sales_rep !== 'string') {
    location.value.sales_rep = String(location.value.sales_rep);
  }
  let locationId = null;
  if (selectedLocationId.value) {
    // update
    isUpdate = true;
    // Remove sales_rep if empty
    const payload = { ...location.value };
    if (payload.sales_rep === undefined || payload.sales_rep === null) delete payload.sales_rep;
    await fetch(`/api/customers/locations/${selectedLocationId.value}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    });
    locationId = selectedLocationId.value;
  } else {
    // create
    const payload = { ...location.value, customer_id: Number(id) };
    if (payload.sales_rep === undefined || payload.sales_rep === null) delete payload.sales_rep;
    const res = await fetch(`/api/customers/${id}/locations`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    });
    if (res.ok) {
      const data = await res.json();
      locationId = data.id;
    }
  }

  // Assign planogram if selected
  if (selectedPlanogramId.value && locationId) {
    await fetch(`/api/planograms/${selectedPlanogramId.value}/assign`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ planogram_id: Number(selectedPlanogramId.value), customer_location_id: Number(locationId) })
    });
  }

  await fetchLocations();
  selectedLocationId.value = '';
  showSuccess(isUpdate ? 'Location updated successfully!' : 'Location added successfully!');
}
</script>

<style scoped>
.update-customer-wrapper {
  min-height: 100vh;
  background: #dbdbdb;
  margin: 0;
  padding: 0;
}
.success-message {
  background: #d4edda;
  color: #155724;
  border: 1px solid #c3e6cb;
  padding: 0.75rem 1.25rem;
  margin: 1rem auto 0.5rem auto;
  border-radius: 4px;
  width: 600px;
  font-size: 1.05rem;
  text-align: center;
  font-family: sans-serif;
  font-weight: 500;
  box-shadow: 0 2px 6px rgba(0,0,0,0.04);
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
