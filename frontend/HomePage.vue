<template>
    <div class="home-wrapper">
        <div class="header">
            <img src="/logo.png" alt="oatmeal studios logo" class="logo" />
            <div class="breadcrumb" @click="goHome" style="cursor:pointer">HOME</div>
        </div>
        <div class="menu-area">
            <div class="menu-row">
                <div class="menu-col">
                    <div class="menu-col-title">MAINTENANCE</div>
                    <button class="menu-btn">Users</button>
                    <button class="menu-btn">Customers</button>
                    <button class="menu-btn">Rep Groups</button>
                    <button class="menu-btn">Items</button>
                    <button class="menu-btn">Pack Contents</button>
                    <button class="menu-btn">Prod Groups</button>
                    <button class="menu-btn">Categories</button>
                    <button class="menu-btn">Artists</button>
                    <button class="menu-btn">Royalties</button>
                    <button class="menu-btn">Inventory</button>
                    <button class="menu-btn">Inventory Status</button>
                    <button class="menu-btn">Planograms</button>
                    <button class="menu-btn">Planogram Reset</button>
                    <button class="menu-btn">Upload Reports</button>
                    <button class="menu-btn">Inventory Move</button>
                </div>
                <div class="menu-col">
                    <div class="menu-col-title">ORDERS</div>
                    <button class="menu-btn">Enter New Orders</button>
                    <button class="menu-btn">Change Orders</button>
                    <button class="menu-btn">Release Orders</button>
                    <button class="menu-btn">Print Pick Lists And Packing Slips</button>
                    <button class="menu-btn">Sign Off Shipments</button>
                    <button class="menu-btn">Invoice Shipped Orders</button>
                    <button class="menu-btn">Pending Orders/ Credits</button>
                    <button class="menu-btn">Download Orders/ Credits</button>
                </div>
                <div class="menu-col">
                    <div class="menu-col-title">ACCOUNTING</div>
                    <button class="menu-btn">Transactions</button>
                    <button class="menu-btn">Rep Deductions</button>
                    <button class="menu-btn">Rep Reinstates</button>
                    <button class="menu-btn">Commissions</button>
                    <button class="menu-btn">Royalties</button>
                    <button class="menu-btn">Statements</button>
                    <button class="menu-btn">Collections</button>
                    <button class="menu-btn">Reinstate</button>
                </div>
                <div class="menu-col">
                    <div class="menu-col-title">REPORTS</div>
                    <button class="menu-btn">Inventory</button>
                    <button class="menu-btn">Inventory By Day</button>
                    <button class="menu-btn">Card Transactions</button>
                    <button class="menu-btn">Customer Balance</button>
                    <button class="menu-btn">Commission</button>
                    <button class="menu-btn">Planograms</button>
                    <button class="menu-btn">Performance</button>
                    <button class="menu-btn">Performance Notecards</button>
                    <button class="menu-btn">Performance Offramp</button>
                    <button class="menu-btn">Performance Master</button>
                    <button class="menu-btn">Sales</button>
                    <button class="menu-btn">Sales Summary</button>
                    <button class="menu-btn">Sales Reps</button>
                    <button class="menu-btn">Rep Codes</button>
                    <button class="menu-btn">Orders Invoiced</button>
                    <button class="menu-btn">Inactive</button>
                    <button class="menu-btn">Top 50</button>
                    <button class="menu-btn">Product Sales</button>
                    <button class="menu-btn">Unused Credits</button>
                    <button class="menu-btn">Email Addresses</button>
                </div>
                <div class="menu-col">
                    <div class="menu-col-title">REPORTS</div>
                    <button class="menu-btn">MTO</button>
                    <button class="menu-btn">MTO - Offramp</button>
                    <button class="menu-btn">Transactions</button>
                    <button class="menu-btn">Best Sellers</button>
                    <button class="menu-btn">Best Sellers Offramp</button>
                    <button class="menu-btn">Edanalysis</button>
                    <button class="menu-btn">Open Orders</button>
                    <button class="menu-btn">Portal</button>
                    <button class="menu-btn">Current Plano</button>
                    <button class="menu-btn">Net Shipped Territory</button>
                    <button class="menu-btn">Net Shipped Writer</button>
                    <button class="menu-btn">Net Shipped Detail</button>
                    <button class="menu-btn">Net Shipped Detail Offramp</button>
                    <button class="menu-btn">Account List</button>
                    <button class="menu-btn">Distributor Accounts</button>
                    <button class="menu-btn">Mailing Labels</button>
                    <button class="menu-btn">Inventory Check</button>
                    <button class="menu-btn">Locations</button>
                    <button class="menu-btn">Artist Audit</button>
                </div>
            </div>
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
.home-wrapper {
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

.menu-area {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    margin-top: 1.5rem;
    margin-left: 0.7rem;
}

.menu-row {
    display: flex;
    flex-direction: row;
    gap: 2.5rem;
}

.menu-col {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    background: none;
    box-shadow: none;
    border: none;
    min-width: 180px;
    padding: 0;
}

.menu-col-title {
    color: #f00;
    font-weight: bold;
    font-size: 1rem;
    margin-bottom: 0.2rem;
    margin-top: 0.1rem;
    letter-spacing: 0.5px;
    text-align: left;
}

.menu-btn {
    width: 170px;
    text-align: left;
    margin-bottom: 0.2rem;
    padding: 2px 8px;
    font-size: 1rem;
    background: #eee;
    border: 1px solid #888;
    border-radius: 2px;
    cursor: pointer;
    font-family: sans-serif;
    font-weight: normal;
    box-sizing: border-box;
}

.menu-btn:last-child {
    margin-bottom: 0;
}
</style>
