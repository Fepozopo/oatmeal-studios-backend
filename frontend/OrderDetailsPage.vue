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
                                customerData.address_2.String }}<br />
                            </span>
                            <span v-if="customerData.city || customerData.state || customerData.zip_code">
                                {{ customerData.city }}, {{ customerData.state }} {{ customerData.zip_code }}
                            </span><br />
                            <span v-if="customerData.country">{{ customerData.country }}</span>
                        </span>
                    </div>
                    <div class="location-info-block">
                        <span class="order-label">Location:</span>
                        <span class="order-value">
                            <span v-if="locationData.location_num && locationData.location_num.Valid">{{
                                locationData.location_num.Int32 }}</span><br />
                            <span v-if="locationData.business_name">{{ locationData.business_name }}</span><br />
                            <span v-if="locationData.address_1">{{ locationData.address_1 }}</span><br />
                            <span v-if="locationData.address_2 && locationData.address_2.Valid">{{
                                locationData.address_2.String }}<br />
                            </span>
                            <span v-if="locationData.city || locationData.state || locationData.zip_code">
                                {{ locationData.city }}, {{ locationData.state }} {{ locationData.zip_code }}
                            </span><br />
                            <span v-if="locationData.country">{{ locationData.country }}</span>
                        </span>
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
                </div>
                <div class="order-row">
                    <span class="order-label">Status:</span>
                    <select class="order-input" v-model="status">
                        <option>APPROVED</option>
                        <option>FUTURE SHIP</option>
                        <option>HOLD</option>
                        <option>SENT TO SHIPPING</option>
                        <option>INVOICED</option>
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
                    <span class="order-label" style="margin-left:2rem;">Write date:</span>
                    <input class="order-input" v-model="writeDate" type="date" />
                    <span class="order-label" style="margin-left:2rem;">Ship date:</span>
                    <input class="order-input" v-model="shipDate" type="date" />
                </div>
                <div class="order-row">
                    <span class="order-label">PO #:</span>
                    <input class="order-input" v-model="poNumber" />
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
                    <select class="order-input" v-model="shipVia">
                        <option></option>
                        <option>UPS GROUND</option>
                        <option>UPS 3-DAY</option>
                        <option>FEDEX GROUND</option>
                        <option>FEDEX 3-DAY</option>
                        <option>USPS GROUND ADVANTAGE</option>
                        <option>USPS PRIORITY</option>
                    </select>
                    <span class="order-label" style="margin-left:2rem;">Free shipping:</span>
                    <label><input type="checkbox" v-model="freeShippingProduct" /> Product</label>
                    <label style="margin-left:1rem;"><input type="checkbox" v-model="freeShippingDisplays" />
                        Displays</label>
                </div>
                <div class="order-row">
                    <span class="order-label">Ship Amount:</span>
                    <input class="order-input" v-model.number="shipAmount" type="number" min="0" step="0.01" />
                </div>
                <div class="order-row">
                    <span class="order-label">Apply to commission:</span>
                    <label><input type="radio" value="Y" v-model="applyCommission" /> Y</label>
                    <label style="margin-left:1rem;"><input type="radio" value="N" v-model="applyCommission" />
                        N</label>
                </div>
                <div class="order-row">
                    <span class="order-label">Special instructions:</span>
                    <textarea class="order-input wide" v-model="specialInstructions" rows="4"
                        style="width:600px;"></textarea>
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
                <div class="order-items-table">
                    <div
                        class="order-items-header"
                        :style="{ gridTemplateColumns: orderItemsGridTemplate }"
                    >
                        <span style="min-width: 60px;"></span>
                        <span v-if="hasPlanogram">Pocket #</span>
                        <span>Item #</span>
                        <span>Qty</span>
                        <span>List Price</span>
                        <span>Discount %</span>
                        <span>Discount Price</span>
                        <span>Total</span>
                    </div>
                    <div
                        v-for="(item, idx) in lineItems"
                        :key="idx"
                        class="order-items-row"
                        :style="{ gridTemplateColumns: orderItemsGridTemplate, alignItems: 'center' }"
                    >
                        <div class="row-controls-horizontal">
                            <button @click="removeLineItem(idx)" :disabled="lineItems.length === 1"
                                class="row-btn-small">-</button>
                            <span class="row-number">{{ idx + 1 }}</span>
                        </div>
                        <input
                            v-if="hasPlanogram"
                            class="order-input tiny"
                            v-model="item.pocket"
                            type="number"
                            min="1"
                            @blur="onPocketChange(idx)"
                            @change="onPocketChange(idx)"
                            placeholder="Pocket #"
                        />
                        <input
                            class="order-input tiny"
                            v-model="item.itemNumber"
                            @blur="onItemNumberBlur(idx)"
                            @change="onSkuChange(idx)"
                        />
                        <input class="order-input tiny" v-model="item.qty" />
                        <input class="order-input tiny" v-model="item.listPrice" />
                        <input class="order-input tiny" v-model="item.discountPct"
                            @keydown.tab="onDiscountTab(idx, $event)" />
                        <span class="order-value">{{ formatCurrency(discountPrice(item)) }}</span>
                        <span class="order-value">{{ formatCurrency(lineTotal(item)) }}</span>
                    </div>
                    <div class="order-items-row" style="align-items:center;">
                        <div class="row-controls-horizontal">
                            <button @click="addLineItem" class="row-btn-small">+</button>
                        </div>
                        <span style="flex:1"></span>
                    </div>
                </div>
                <!-- Order Totals Section -->
                <div class="order-row">
                    <span style="margin-left:auto; min-width: 650px; text-align:right; display:block;">Order total
                        <b>{{ formatCurrency(itemTotal) }}</b></span>
                </div>
                <div class="order-row">
                    <span style="margin-left:auto; min-width: 650px; text-align:right; display:block;">Shipping amount
                        <b>{{ formatCurrency(shipAmount) }}</b></span>
                </div>
                <div class="order-row">
                    <span style="margin-left:auto; min-width: 650px; text-align:right; display:block;">Total
                        <b>{{ formatCurrency(grandTotal) }}</b></span>
                </div>
                <div class="order-row">
                    <button class="continue-btn">SUBMIT</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { useSalesReps } from './useSalesReps.js';
import { ref, onMounted, computed, nextTick, watch } from 'vue';
const salesperson = ref("");
const { salesReps, fetchSalesReps } = useSalesReps();
import { useRoute, useRouter } from 'vue-router';


const router = useRouter();
const route = useRoute();
const goHome = () => {
    router.push('/home');
};

// Computed property for item total
const itemTotal = computed(() => {
    return lineItems.value.reduce((sum, item) => {
        const total = parseFloat(item.qty) ? lineTotal(item) : 0;
        return sum + total;
    }, 0);
});

const status = ref("APPROVED");
const type = ref("REORDER");
const method = ref("ONLINE");

const terms = ref("");
const shipVia = ref("UPS GROUND");
const freeShippingProduct = ref(false);
const freeShippingDisplays = ref(false);
const shipAmount = ref(0);
// Grand total includes item total and shipping amount
const grandTotal = computed(() => {
    const shipping = parseFloat(shipAmount.value) || 0;
    return itemTotal.value + shipping;
});

const applyCommission = ref("Y");

const freeDisplay = ref("N");
const defaultQty = ref(6);
const defaultDiscount = ref(0);

const orderNumber = ref(generateOrderNumber());
const customerId = route.query.customerId;
const locationId = route.query.locationId;

const customerData = ref({});
const locationData = ref({});
const customerLink = computed(() =>
    customerData.value.id ? `/customers/${customerData.value.id}` : '#'
);

function getPSTDateString() {
    return new Date().toLocaleDateString('en-CA', { timeZone: 'America/Los_Angeles' });
}
const writeDate = ref(getPSTDateString());
const shipDate = ref(getPSTDateString());

function discountPrice(item) {
    const price = parseFloat(item.listPrice) || 0;
    const pct = parseFloat(item.discountPct) || 0;
    return price * (1 - pct / 100);
}
function lineTotal(item) {
    const qty = parseFloat(item.qty) || 0;
    return discountPrice(item) * qty;
}
function formatCurrency(val) {
    if (isNaN(val)) return '';
    return '$' + val.toFixed(2);
}

// --- Planogram logic ---
const hasPlanogram = ref(false);
const planogramId = ref(null);
const planogramPockets = ref([]); // [{pocket_number, sku, ...}]

// Fetch planogram for this location, then fetch pockets if exists
async function fetchPlanogramAndPockets() {
  if (!locationId) return;
  // Get planogram assigned to this location
  const planogramRes = await fetch(`/api/planograms/${locationId}/planograms`);
  if (planogramRes.ok) {
    const planogram = await planogramRes.json();
    if (planogram && planogram.id) {
      hasPlanogram.value = true;
      planogramId.value = planogram.id;
      // Fetch pockets for this planogram
      const pocketsRes = await fetch(`/api/planograms/${planogram.id}/pockets`);
      if (pocketsRes.ok) {
        planogramPockets.value = await pocketsRes.json();
      } else {
        planogramPockets.value = [];
      }
    } else {
      hasPlanogram.value = false;
      planogramId.value = null;
      planogramPockets.value = [];
    }
  } else {
    hasPlanogram.value = false;
    planogramId.value = null;
    planogramPockets.value = [];
  }
}

// When a pocket is selected, auto-fill the itemNumber (SKU) for that line
async function onPocketChange(idx) {
  const item = lineItems.value[idx];
  const pocketNum = item.pocket;
  if (!pocketNum) {
    // If cleared, just clear fields
    item.qty = '';
    item.discountPct = '';
    item.listPrice = '';
    return;
  }
  // Only check for valid pocket if planogram is present
  if (hasPlanogram.value) {
    const pockets = Array.isArray(planogramPockets.value) ? planogramPockets.value : [];
    const pocket = pockets.find(p => String(p.pocket_number) === String(pocketNum));
    if (!pocket) {
      alert(`Pocket #${pocketNum} does not exist in the assigned planogram.`);
      item.pocket = '';
      item.qty = '';
      item.discountPct = '';
      item.listPrice = '';
      return;
    }
    if (pocket.sku && pocket.sku.Valid) {
      item.itemNumber = pocket.sku.String;
      // Auto-fetch product info and set qty, discount, price
      try {
        const res = await fetch(`/api/products/sku/${encodeURIComponent(item.itemNumber)}`);
        if (res.ok) {
          const product = await res.json();
          if ((product.status || '').trim() !== 'INACTIVE') {
            item.qty = defaultQty.value;
            item.discountPct = defaultDiscount.value;
            item.listPrice = typeof product.price === 'object' && product.price !== null
              ? (product.price.Float64 ?? product.price.value ?? '')
              : product.price;
          } else {
            item.qty = '';
            item.discountPct = '';
            item.listPrice = '';
          }
        } else {
          item.qty = '';
          item.discountPct = '';
          item.listPrice = '';
        }
      } catch {
        item.qty = '';
        item.discountPct = '';
        item.listPrice = '';
      }
    } else {
      // If pocket exists but no SKU, just clear fields
      item.qty = '';
      item.discountPct = '';
      item.listPrice = '';
    }
  }
}

// When SKU is changed, update pocket selection if it matches a pocket, and fetch product info
async function onSkuChange(idx) {
  const item = lineItems.value[idx];
  const sku = (item.itemNumber || '').trim();
  if (!sku) {
    item.pocket = '';
    item.qty = '';
    item.discountPct = '';
    item.listPrice = '';
    return;
  }
  // If planogram, try to match pocket
  if (hasPlanogram.value) {
    const pockets = Array.isArray(planogramPockets.value) ? planogramPockets.value : [];
    const pocket = pockets.find(p => p.sku && p.sku.Valid && p.sku.String === sku);
    item.pocket = pocket ? pocket.pocket_number : '';
  }
  // Fetch product info
  try {
    const res = await fetch(`/api/products/sku/${encodeURIComponent(sku)}`);
    if (res.ok) {
      const product = await res.json();
      if ((product.status || '').trim() !== 'INACTIVE') {
        item.qty = defaultQty.value;
        item.discountPct = defaultDiscount.value;
        item.listPrice = typeof product.price === 'object' && product.price !== null
          ? (product.price.Float64 ?? product.price.value ?? '')
          : product.price;
      } else {
        item.qty = '';
        item.discountPct = '';
        item.listPrice = '';
      }
    } else {
      item.qty = '';
      item.discountPct = '';
      item.listPrice = '';
    }
  } catch {
    item.qty = '';
    item.discountPct = '';
    item.listPrice = '';
  }
}

// Utility to get SKU for a pocket number (if needed in template)
function getSkuForPocket(pocketNum) {
  const pockets = Array.isArray(planogramPockets.value) ? planogramPockets.value : [];
  const pocket = pockets.find(p => String(p.pocket_number) === String(pocketNum));
  return pocket && pocket.sku && pocket.sku.Valid ? pocket.sku.String : '';
}

// --- Line items state ---
const lineItems = ref([
    { pocket: '', itemNumber: '', qty: '', description: '', listPrice: '', discountPct: '', discountPrice: '', total: '' }
]);

function addLineItem() {
    // If planogram, set defaults for new line
    let newItem = { pocket: '', itemNumber: '', qty: '', description: '', listPrice: '', discountPct: '', discountPrice: '', total: '' };
    if (hasPlanogram.value) {
      newItem.qty = defaultQty.value;
      newItem.discountPct = defaultDiscount.value;
      newItem.listPrice = 0;
    }
    lineItems.value.push(newItem);
    nextTick(() => {
        // Focus the first input of the new row
        const rows = document.querySelectorAll('.order-items-row');
        if (rows.length > 0) {
            const lastRow = rows[rows.length - 1];
            const firstInput = lastRow.querySelector('input');
            if (firstInput) firstInput.focus();
        }
    });
}

function addLineItemAt(idx) {
    lineItems.value.splice(idx + 1, 0, { pocket: '', itemNumber: '', qty: '', description: '', listPrice: '', discountPct: '', discountPrice: '', total: '' });
    nextTick(() => {
        // Focus the first input of the new row
        const rows = document.querySelectorAll('.order-items-row');
        if (rows.length > idx + 1) {
            const newRow = rows[idx + 1];
            const firstInput = newRow.querySelector('input');
            if (firstInput) firstInput.focus();
        }
    });
}

function removeLineItem(idx) {
    if (lineItems.value.length > 1) {
        lineItems.value.splice(idx, 1);
    }
}

function onDiscountTab(idx, event) {
    // Only add if tabbing out of last discount % input
    if (idx === lineItems.value.length - 1 && !event.shiftKey) {
        addLineItem();
    }
}

function handleItemNumberChange(idx) {
    const item = lineItems.value[idx];
    const itemNumber = item.itemNumber?.trim();
    if (!itemNumber) return;

    fetch(`/api/products/sku/${encodeURIComponent(itemNumber)}`)
        .then(async res => {
            if (!res.ok) {
                alert(`Product with SKU "${itemNumber}" not found.`);
                item.listPrice = '';
                item.qty = '';
                item.discountPct = '';
                return;
            }
            const product = await res.json();
            // Prevent adding inactive products
            if ((product.status || '').trim() === 'INACTIVE') {
                alert(`Product "${itemNumber}" is inactive and cannot be added to the order.`);
                item.listPrice = '';
                item.qty = '';
                item.discountPct = '';
                return;
            }
            // Set fields from defaults and product
            item.qty = defaultQty.value;
            item.discountPct = defaultDiscount.value;
            item.listPrice = typeof product.price === 'object' && product.price !== null
                ? (product.price.Float64 ?? product.price.value ?? '')
                : product.price;
        })
        .catch(() => {
            alert(`Error looking up product "${itemNumber}".`);
        });
}

function onItemNumberBlur(idx) {
    const item = lineItems.value[idx];
    const itemNumber = item.itemNumber?.trim();
    if (!itemNumber) return;

    fetch(`/api/products/sku/${encodeURIComponent(itemNumber)}`)
        .then(async res => {
            if (!res.ok) {
                alert(`Product with SKU "${itemNumber}" not found.`);
                item.listPrice = '';
                item.qty = '';
                item.discountPct = '';
                return;
            }
            const product = await res.json();
            // Prevent adding inactive products
            if ((product.status || '').trim() === 'INACTIVE') {
                alert(`Product "${itemNumber}" is inactive and cannot be added to the order.`);
                item.listPrice = '';
                item.qty = '';
                item.discountPct = '';
                return;
            }
            // Set fields from defaults and product
            item.qty = defaultQty.value;
            item.discountPct = defaultDiscount.value;
            item.listPrice = typeof product.price === 'object' && product.price !== null
                ? (product.price.Float64 ?? product.price.value ?? '')
                : product.price;
        })
        .catch(() => {
            alert(`Error looking up product "${itemNumber}".`);
        });
}

onMounted(async () => {
    if (customerId) {
        const res = await fetch(`/api/customers/${customerId}`);
        if (res.ok) {
            customerData.value = await res.json();
            // Set free shipping checkboxes if customer gets free shipping
            if (customerData.value.free_shipping) {
                freeShippingProduct.value = true;
                freeShippingDisplays.value = true;
            }
            // Set default terms if available
            if (customerData.value.terms) {
                terms.value = customerData.value.terms;
            }
            // Set default discount if available
            if (typeof customerData.value.discount !== 'undefined' && customerData.value.discount !== null) {
                defaultDiscount.value = customerData.value.discount;
            }
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
    // Fetch planogram and pockets for this location
    await fetchPlanogramAndPockets();
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

const orderItemsGridTemplate = computed(() =>
  hasPlanogram.value
    ? '60px 100px 100px 80px 1.5fr 100px 100px 100px 100px'
    : '60px 100px 80px 1.5fr 100px 100px 100px 100px'
);
</script>

<style scoped>
.order-items-table {
    background: #dbdbdb;
    padding: 1rem 0.5rem 0.5rem 0.5rem;
    border-radius: 6px;
    width: 100%;
    min-width: 900px;
}

.order-items-header {
    display: grid;
    font-weight: bold;
    margin-bottom: 0.5rem;
    align-items: center;
}

.order-items-row {
    display: grid;
    margin-bottom: 0.5rem;
    align-items: center;
}

.row-controls {
    display: flex;
    flex-direction: column;
    align-items: center;
    min-width: 2.5em;
}

.row-controls-horizontal {
    display: flex;
    flex-direction: row;
    align-items: center;
    min-width: 2.5em;
    gap: 0.2em;
}

.row-btn {
    width: 2em;
    height: 2em;
    font-size: 1.2em;
    background: #eaeaea;
    border: 1px solid #aaa;
    border-radius: 6px;
    margin: 0.1em 0;
    cursor: pointer;
    padding: 0;
}

.row-btn-small {
    width: 1.5em;
    height: 1.5em;
    font-size: 1em;
    background: #eaeaea;
    border: 1px solid #aaa;
    border-radius: 6px;
    margin: 0 0.1em 0 0;
    cursor: pointer;
    padding: 0;
}

.row-btn:disabled,
.row-btn-small:disabled {
    opacity: 0.5;
    cursor: not-allowed;
}

.row-number {
    font-size: 1.1em;
    font-weight: 500;
    margin: 0.2em 0;
    min-width: 1.2em;
    text-align: center;
}

.order-input.tiny {
    width: 90%;
    min-width: 0;
}

.order-input.wide {
    width: 98%;
}

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
    background: #dbdbdb;
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
