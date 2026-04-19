<script setup>
// 1. Impor yang dibutuhkan
import { createRouter, createWebHistory, RouterView, RouterLink } from 'vue-router'
import { getCurrentInstance } from 'vue'

// 2. Impor komponen halaman Anda
import welcome from './welcome.vue'
import Dashboard from './MainDashboard.vue'

// 3. Definisi rute
const routes = [
  { path: '/', component: welcome },
  { path: '/dashboard', component: Dashboard }
]

// 4. Inisialisasi router
const router = createRouter({
  history: createWebHistory(),
  routes
})

// 5. Daftarkan router ke instance aplikasi saat ini
const app = getCurrentInstance().proxy.$.appContext.app
if (!app._context.provides.router) {
  app.use(router)
}
</script>

<template>
  <nav>
    <RouterLink to="/">Home</RouterLink> | 
    <RouterLink to="/dashboard">Dashboard</RouterLink>
  </nav>

  <hr>

  <RouterView />
</template>