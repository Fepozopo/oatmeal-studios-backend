<template>
    <div class="order-details-wrapper">
        <div class="header">
            <img src="/logo.png" alt="oatmeal studios logo" class="logo" />
            <div class="breadcrumb" @click="goHome" style="cursor:pointer">HOME</div>
        </div>
        <div class="order-details-content">
            <div class="order-details-title">ORDER ENTRY</div>
            <div class="order-details-section">
                <div class="order-row">
                    <div style="display:flex; flex-direction:column; align-items:flex-start;">
                        <div>
                            <span class="order-label">Order #:</span>
                            <span class="order-value">{{ orderNumber }}</span>
                        </div>
                        <button class="delete-btn" style="margin-left:0; margin-top:0.5rem;">DELETE</button>
                    </div>
                </div>
                <div class="order-row customer-location-row">
                    <div class="customer-info-block">
                        <span class="order-label">Customer:</span>
                        <span class="order-value">
                            <a v-if="customerData.id" :href="customerLink" class="customer-link">{{ customerData.id
                                }}</a><br />
                            <span v-if="customerData.business_name">{{ customerData.business_name }}</span><br />
                            <span v-if="customerData.address_1">{{ customerData.address_1 }}</span><br />
                            <span v-if="customerData.address_2 && customerData.address_2.Valid">{{
                                customerData.address_2.String }}</span><br />
                            <span v-if="customerData.city || customerData.state || customerData.zip_code">
                                {{ customerData.city }}, {{ customerData.state }} {{ customerData.zip_code }}
                            </span><br />
                            <span v-if="customerData.country">{{ customerData.country }}</span>
                        </span>
                    </div>
                    <div class="location-info-block">
                        <span class="order-label">Location:</span>
                        <span class="order-value">
                            <span v-if="locationData.id">{{ locationData.id }}</span><br />
                            <span v-if="locationData.business_name">{{ locationData.business_name }}</span><br />
                            <span v-if="locationData.address_1">{{ locationData.address_1 }}</span><br />
                            <span v-if="locationData.address_2 && locationData.address_2.Valid">{{
                                locationData.address_2.String }}</span><br />
                            <span v-if="locationData.city || locationData.state || locationData.zip_code">
                                {{ locationData.city }}, {{ locationData.state }} {{ locationData.zip_code }}
                            </span><br />
                            <span v-if="locationData.country">{{ locationData.country }}</span>
                        </span>
                    </div>
                </div>
                <div class="order-row">
                    <span class="order-label">Salesperson:</span>
                    <select class="order-input wide" v-model="salesperson">
                        <option value=""></option>
                        <option v-for="rep in salesReps" :key="rep.rep_code" :value="rep.rep_code">
                            {{ rep.rep_code }} - {{ rep.first_name }} {{ rep.last_name }}
                        </option>
                    </select>
                </div>
                <div class="order-row">
                    <span class="order-label">Status:</span>
                    <select class="order-input" v-model="status">
                        <option>PENDING</option>
                        <option>COMPLETE</option>
                    </select>
                    <span class="order-label" style="margin-left:2rem;">Type:</span>
                    <select class="order-input" v-model="type">
                        <option></option>
                        <option>NEW</option>
                        <option>REORDER</option>
                        <option>CREDIT</option>
                    </select>
                    <span class="order-label" style="margin-left:2rem;">Method:</span>
                    <select class="order-input" v-model="method">
                        <option></option>
                        <option>ONLINE</option>
                        <option>EMAIL</option>
                        <option>PHONE</option>
                    </select>
                </div>
                <div class="order-row">
                    <span class="order-label">Last Order:</span>
                    <input class="order-input" v-model="lastOrder" type="date" />
                    <span class="order-label" style="margin-left:2rem;">Write date:</span>
                    <input class="order-input" v-model="writeDate" type="date" />
                    <span class="order-label" style="margin-left:2rem;">Ship date:</span>
                    <input class="order-input" v-model="shipDate" type="date" />
                </div>
                <div class="order-row">
                    <span class="order-label">PO #:</span>
                    <input class="order-input" v-model="poNumber" />
                    <span class="order-label" style="margin-left:2rem;">PCP:</span>
                    <input class="order-input short" v-model="pcp" />
                    <span class="order-label" style="margin-left:2rem;">Subs:</span>
                    <span>N</span>
                </div>
                <div class="order-row">
                    <span class="order-label">Terms:</span>
                    <select class="order-input" v-model="terms">
                        <option></option>
                        <option>CREDIT CARD</option>
                        <option>NET 30</option>
                        <option>NET 60</option>
                        <option>NET 90</option>
                    </select>
                    <span class="order-label" style="margin-left:2rem;">Ship via:</span>
                    <input class="order-input" v-model="shipVia" />
                    <span class="order-label" style="margin-left:2rem;">Free shipping:</span>
                    <label><input type="checkbox" v-model="freeShippingProduct" /> Product</label>
                    <label style="margin-left:1rem;"><input type="checkbox" v-model="freeShippingDisplays" />
                        Displays</label>
                </div>
                <div class="order-row">
                    <span class="order-label">Ship Amount:</span>
                    <input class="order-input" v-model="shipAmount" />
                </div>
                <div class="order-row">
                    <span class="order-label">Apply to commission:</span>
                    <label><input type="radio" value="Y" v-model="applyCommission" /> Y</label>
                    <label style="margin-left:1rem;"><input type="radio" value="N" v-model="applyCommission" />
                        N</label>
                </div>
                <div class="order-row">
                    <span class="order-label">Restricted items:</span>
                    <input class="order-input wide" v-model="restrictedItems" />
                </div>
                <div class="order-row">
                    <span class="order-label">Special instructions:</span>
                    <textarea class="order-input wide" v-model="specialInstructions"></textarea>
                </div>
                <div class="order-row">
                    <span class="order-label">Default Qty:</span>
                    <input class="order-input short" v-model="defaultQty" />
                    <span class="order-label" style="margin-left:2rem;">Default Discount %:</span>
                    <input class="order-input short" v-model="defaultDiscount" />
                    <span class="order-label" style="margin-left:2rem;">Free display:</span>
                    <label><input type="radio" value="Y" v-model="freeDisplay" /> Y</label>
                    <label style="margin-left:1rem;"><input type="radio" value="N" v-model="freeDisplay" /> N</label>
                </div>
                <div class="order-row">
                    <div style="display:flex;align-items:center;gap:0.5rem;">
                        <button @click="decrementItemCount">-</button>
                        <input class="order-input short" v-model="itemCount" style="width:2.5em;text-align:center;" />
                        <button @click="incrementItemCount">+</button>
                    </div>
                </div>
                <div class="order-items-table">
                    <div class="order-items-header">
                        <span>Pocket #</span>
                        <span>Item #</span>
                        <span>Qty</span>
                        <span>Description</span>
                        <span>List Price</span>
                        <span>Discount %</span>
                        <span>Discount Price</span>
                        <span>Total</span>
                    </div>
                    <div class="order-items-row" v-for="n in itemCount" :key="n">
                        <input class="order-input tiny" />
                        <input class="order-input tiny" />
                        <input class="order-input tiny" />
                        <input class="order-input wide" />
                        <input class="order-input tiny" />
                        <input class="order-input tiny" />
                        <input class="order-input tiny" />
                        <input class="order-input tiny" />
                    </div>
                </div>
                <div class="order-row">
                    <span style="margin-left:auto;">Item total <b>0</b></span>
                </div>
                <div class="order-row">
                    <button class="continue-btn">CONTINUE</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>

import { useSalesReps } from './useSalesReps.js';
import { ref, onMounted, computed } from 'vue';
const salesperson = ref("");
const { salesReps, fetchSalesReps } = useSalesReps();
import { useRoute, useRouter } from 'vue-router';

const router = useRouter();
const route = useRoute();
const goHome = () => {
    router.push('/home');
};

const orderNumber = ref(generateOrderNumber());
const customerId = route.query.customerId;
const locationId = route.query.locationId;

const customerData = ref({});
const locationData = ref({});
const customerLink = computed(() =>
    customerData.value.id ? `/customers/${customerData.value.id}` : '#'
);


onMounted(async () => {
    if (customerId) {
        const res = await fetch(`/api/customers/${customerId}`);
        if (res.ok) {
            customerData.value = await res.json();
        }
    }
    let locationLoaded = false;
    if (customerId && locationId) {
        const res = await fetch(`/api/customers/${customerId}/locations`);
        if (res.ok) {
            const locations = await res.json();
            const location = locations.find(l => String(l.id) === String(locationId));
            if (location) {
                locationData.value = location;
                locationLoaded = true;
            }
        }
    }
    await fetchSalesReps();
    // Debug log to check values
    console.log('locationData:', locationData.value);
    console.log('salesReps:', salesReps.value);
    // Set default salesperson if available
    if (locationLoaded && locationData.value.sales_rep && locationData.value.sales_rep.Valid && salesReps.value.length > 0) {
        salesperson.value = locationData.value.sales_rep.String;
    }
});

function generateOrderNumber() {
    // Generate a random order number
    return Math.floor(202500000 + Math.random() * 100000).toString();
}
</script>

<style scoped>
.order-details-wrapper {
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

.order-details-content {
    margin-top: 1.5rem;
    margin-left: 1.5rem;
}

.order-details-title {
    font-size: 1.2rem;
    font-weight: bold;
    margin-bottom: 1.2rem;
}

.order-details-section {
    background: #eaeaea;
    padding: 1rem;
    border-radius: 6px;
    width: 900px;
    margin-bottom: 2rem;
}


.order-row {
    display: flex;
    align-items: flex-start;
    margin-bottom: 0.7rem;
}

.customer-location-row {
    display: flex;
    flex-direction: row;
    gap: 3rem;
}

.customer-info-block,
.location-info-block {
    display: flex;
    flex-direction: column;
    min-width: 260px;
}

.order-label {
    font-weight: bold;
    width: 120px;
    display: inline-block;
}

.order-value {
    font-family: monospace;
    white-space: pre-line;
}

.customer-link {
    color: #1a0dab;
    text-decoration: underline;
    cursor: pointer;
}

.delete-btn {
    margin-left: 2rem;
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
