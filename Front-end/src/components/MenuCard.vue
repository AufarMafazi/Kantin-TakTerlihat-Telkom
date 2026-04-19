<script setup>
defineProps(['menu'])
defineEmits(['add'])

const formatRp = (angka) => {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(angka)
}
</script>

<template>
  <div class="bg-white rounded-3xl shadow-md hover:shadow-xl transition-all duration-300 border border-[#EBC898]/40 overflow-hidden flex flex-col group cursor-pointer" @click="$emit('add', menu)">
    <div class="h-48 w-full overflow-hidden bg-gray-100 relative">
      <img :src="menu.image" :alt="menu.name" class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-500">
    </div>

    <div class="p-5 flex flex-col flex-grow bg-white">
      <h3 class="font-extrabold text-lg mb-1 text-gray-800">{{ menu.name }}</h3>
      <p class="text-[#9C6307] font-black text-xl mb-4">{{ formatRp(menu.price) }}</p>
      
      <button class="mt-auto w-full bg-[#FDF8F3] hover:bg-[#9C6307] hover:text-white text-[#9C6307] font-bold py-3 rounded-2xl transition-colors border border-[#EBC898]">
        + Tambah ke Pesanan
      </button>
    </div>
  </div>
</template>