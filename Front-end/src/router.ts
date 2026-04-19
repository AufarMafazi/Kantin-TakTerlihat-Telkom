import { createRouter, createWebHistory } from 'vue-router'
import SelfServiceView from './views/SelfServiceView.vue'
import PembayaranView from './views/PembayaranView.vue'
import PenerimaPesananView from './views/PenerimaPesananView.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: SelfServiceView }, // Ini halaman utama saat web dibuka
    { path: '/kasir', component: PembayaranView },
    { path: '/dapur', component: PenerimaPesananView }
  ]
})

export default router