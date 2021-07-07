import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/homepage/Home.vue'

<<
<< << < HEAD
const routes = [{
            path: '/',
            name: 'Home',
            component: Home
        },
        {
            path: '/login',
            name: 'Login',
            // route level code-splitting
            // this generates a separate chunk (about.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            component: () =>
                import ( /* webpackChunkName: "about" */ '../views/Login.vue')
        } ===
        === =
        const routes = [{
                path: '/',
                name: 'Home',
                component: Home
            },
            {
                path: '/login',
                name: 'Login',
                // route level code-splitting
                // this generates a separate chunk (about.[hash].js) for this route
                // which is lazy-loaded when the route is visited.
                component: () =>
                    import ( /* webpackChunkName: "about" */ '../views/login/Login.vue')
            } >>>
            >>> > be4eb13b35184a727a4803c08806fd8e01ad6984
        ]

        const router = createRouter({
            history: createWebHistory(process.env.BASE_URL),
            routes
        })

        export default router