<script setup>
import { ref, computed } from 'vue'
import MenuCard from '../components/MenuCard.vue'
import CartSidebar from '../components/CartSidebar.vue'
import PaymentModal from '../components/PaymentModal.vue'

const menus = ref([
  { id: 1, name: 'Nasi Goreng Ayam', price: 20000, image: '/image/Nasigoreng_ayam.png' },
  { id: 2, name: 'Ayam Geprek', price: 18000, image: '/image/Ayam_geprek.png' },
  { id: 3, name: 'Mie Ayam Bakso', price: 15000, image: '/image/Mieayam_bakso.png' },
  { id: 4, name: 'Soto Ayam', price: 15000, image: '/image/Soto_ayam.png' },
  { id: 5, name: 'Ayam Bakar', price: 22000, image: '/image/Ayam_bakar.png' },
  { id: 6, name: 'Tahu Goreng', price: 5000, image: '/image/Tahu_goreng.png' },
])

const cart = ref([])
const showModal = ref(false)

const addToCart = (menu) => {
  const existingItem = cart.value.find(item => item.id === menu.id)
  if (existingItem) {
    existingItem.qty++
  } else {
    cart.value.push({ ...menu, qty: 1 })
  }
}

const removeFromCart = (id) => {
  const index = cart.value.findIndex(item => item.id === id)
  if (index !== -1) {
    if (cart.value[index].qty > 1) {
      cart.value[index].qty--
    } else {
      cart.value.splice(index, 1)
    }
  }
}

const totalHarga = computed(() => {
  return cart.value.reduce((total, item) => total + (item.price * item.qty), 0)
})

const handleBayarClick = () => {
  if (cart.value.length === 0) return
  showModal.value = true
}

const prosesPembayaran = (metode) => {
  showModal.value = false
  alert(metode === 'qris' ? "✅ Scan Berhasil!" : "💵 Silakan ke Kasir!")
  cart.value = []
}
</script>

<template>
  <div class="min-h-screen bg-[#FDF8F3] flex font-sans text-gray-800">
    <div class="w-2/3 p-8 overflow-y-auto h-screen relative">
      <div class="flex justify-between items-center mb-10 bg-white p-6 rounded-3xl shadow-sm border border-[#EBC898]/30 sticky top-0 z-20">
        <div>
          <h1 class="text-4xl font-black text-transparent bg-clip-text bg-gradient-to-r from-[#9C6307] to-[#6E4403]">Kantin Tak Terlihat</h1>
          <p class="text-gray-500 font-medium mt-1">Silakan sentuh layar untuk memesan makanan</p>
        </div>
        <div class="bg-[#FDF8F3] text-[#9C6307] px-4 py-2 rounded-2xl font-bold tracking-widest text-sm border border-[#EBC898]">KANTIN TERBAEK</div>
      </div>

      <div class="grid grid-cols-3 gap-6 pb-10">
        <MenuCard v-for="menu in menus" :key="menu.id" :menu="menu" @add="addToCart" />
      </div>
    </div>

    <CartSidebar 
      :cart="cart" 
      :total="totalHarga" 
      @bayar="handleBayarClick" 
      @kurang="removeFromCart"
      @tambah="addToCart" 
    />
    
    <PaymentModal :show="showModal" :total="totalHarga" @close="showModal = false" @pilih-metode="prosesPembayaran" />
  </div>
</template>