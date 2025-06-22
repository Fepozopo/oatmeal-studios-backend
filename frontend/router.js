import { createRouter, createWebHistory } from 'vue-router';
import HomePage from './HomePage.vue';
import LoginPage from './LoginPage.vue';


const routes = [
    { path: '/login', component: LoginPage },
    { path: '/home', component: HomePage, meta: { requiresAuth: true } },
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