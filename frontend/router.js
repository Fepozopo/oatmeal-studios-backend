import { createRouter, createWebHistory } from 'vue-router';

import HomePage from './HomePage.vue';
import LoginPage from './LoginPage.vue';
import UserMaintenance from './UserMaintenance.vue';
import CustomerMaintenance from './CustomerMaintenance.vue';
import OrderEntryPage from './OrderEntryPage.vue';
import OrderDetailsPage from './OrderDetailsPage.vue';
import CreateCustomerPage from './CreateCustomerPage.vue';
import UpdateCustomerPage from './UpdateCustomerPage.vue';


const routes = [
    { path: '/login', component: LoginPage },
    { path: '/home', component: HomePage, meta: { requiresAuth: true } },
    { path: '/users', component: UserMaintenance, meta: { requiresAuth: true } },
    { path: '/customers', component: CustomerMaintenance, meta: { requiresAuth: true } },
    { path: '/customers/new', component: CreateCustomerPage, meta: { requiresAuth: true } },
    { path: '/order-entry', component: OrderEntryPage, meta: { requiresAuth: true } },
    { path: '/order-details', component: OrderDetailsPage, meta: { requiresAuth: true } },
    { path: '/customers/:id/edit', component: UpdateCustomerPage, meta: { requiresAuth: true } },
    { path: '/', redirect: '/home' },
];


const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    const isAuthenticated = !!localStorage.getItem('token');
    if (to.meta.requiresAuth && !isAuthenticated) {
        next('/login');
    } else if (to.path === '/login' && isAuthenticated) {
        next('/home');
    } else {
        next();
    }
});

export default router;