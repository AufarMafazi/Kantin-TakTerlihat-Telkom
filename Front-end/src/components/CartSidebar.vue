<script setup>
defineProps(['cart', 'total'])
defineEmits(['bayar', 'tambah', 'kurang'])

const formatRp = (angka) => {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(angka)
}
</script>

<template>
  <div class="w-1/3 bg-white shadow-2xl z-10 flex flex-col h-screen border-l border-[#EBC898]/30">
    <div class="p-6 border-b border-[#EBC898]/30 bg-white">
      <h2 class="text-2xl font-black text-[#6E4403] flex items-center gap-2">
        🛒 Pesanan Saya
      </h2>
    </div>

    <div class="flex-grow p-4 overflow-y-auto bg-[#FDF8F3]/50">
      <div v-if="cart.length === 0" class="flex flex-col items-center justify-center h-full text-gray-400">
        <p class="font-medium text-lg">Keranjang masih kosong</p>
      </div>
      
      <div v-else class="space-y-3">
        <div v-for="item in cart" :key="item.id" class="flex justify-between items-center bg-white p-3 rounded-2xl border border-[#EBC898]/50 shadow-sm">
          <div class="flex gap-3 items-center w-full">
            <div class="w-16 h-16 flex-shrink-0 rounded-xl overflow-hidden bg-gray-100">
              <img :src="item.image" :alt="item.name" class="w-full h-full object-cover">
            </div>
            
            <div class="flex-grow">
              <p class="font-bold text-[#6E4403] text-sm leading-tight mb-2">{{ item.name }}</p>
              
              <div class="flex justify-between items-center">
                <div class="flex items-center gap-3 bg-[#FDF8F3] rounded-lg p-1 border border-[#EBC898]/40">
                  <button @click="$emit('kurang', item.id)" class="w-7 h-7 flex items-center justify-center bg-white rounded-md shadow-sm text-[#9C6307] font-bold hover:bg-[#9C6307] hover:text-white transition-colors">-</button>
                  <span class="font-black text-[#9C6307] min-w-[20px] text-center">{{ item.qty }}</span>
                  <button @click="$emit('tambah', item)" class="w-7 h-7 flex items-center justify-center bg-white rounded-md shadow-sm text-[#9C6307] font-bold hover:bg-[#9C6307] hover:text-white transition-colors">+</button>
                </div>
                <p class="font-bold text-[#6E4403] text-sm">{{ formatRp(item.price * item.qty) }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="p-6 bg-white border-t border-[#EBC898]/30">
      <div class="flex justify-between items-center mb-6">
        <span class="text-[#6E4403] font-bold">Total Tagihan</span>
        <span class="text-3xl font-black text-[#9C6307]">{{ formatRp(total) }}</span>
      </div>
      <button @click="$emit('bayar')"
        class="w-full bg-[#9C6307] hover:bg-[#6E4403] text-white font-bold text-xl py-4 rounded-2xl shadow-lg transition-transform active:scale-95">
        Bayar Sekarang 💳
      </button>
    </div>
  </div>
</template>