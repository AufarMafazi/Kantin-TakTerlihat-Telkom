<!-- MODUL 3: Penerima Pesanan / Layar Dapur -->

<script setup>
import { ref } from 'vue'

// INI DATA DIGANTI FAR MASIH DUMMY
const pesananDapur = ref([
  {
    id: 'TRX-002',
    waktu: '10:32',
    status: 'Sedang Dimasak',
    items: [
      { name: 'Ayam Geprek', qty: 1, notes: 'Pedas sedang' },
      { name: 'Es Teh Manis', qty: 1, notes: 'Sedikit es' }
    ]
  },
  {
    id: 'TRX-004',
    waktu: '10:45',
    status: 'Antrean Baru',
    items: [
      { name: 'Nasi Goreng Ayam', qty: 2, notes: 'Satu tidak pakai kecap' },
      { name: 'Mie Ayam Bakso', qty: 1, notes: '' }
    ]
  },
  {
    id: 'TRX-005',
    waktu: '10:50',
    status: 'Antrean Baru',
    items: [
      { name: 'Soto Ayam', qty: 3, notes: 'Pisah kuah' }
    ]
  }
])

// Fungsi saat Koki menekan tombol Selesai
const pesananSelesai = (idPesanan) => {
  const index = pesananDapur.value.findIndex(p => p.id === idPesanan)
  if (index !== -1) {
    alert(`Pesanan ${idPesanan} sudah selesai dimasak dan siap dipanggil! 🍳`)
    // Hapus dari layar dapur karena sudah selesai
    pesananDapur.value.splice(index, 1)
  }
}

// mengubah status menjadi Sedang Dimasak
const mulaiMasak = (pesanan) => {
  pesanan.status = 'Sedang Dimasak'
}
</script>

<template>
  <div class="min-h-screen bg-[#FDF8F3] font-sans text-gray-800 p-8">
    
    <div class="flex justify-between items-center mb-8 bg-white p-6 rounded-3xl shadow-sm border border-[#EBC898]/30">
      <div>
        <h1 class="text-3xl font-black text-[#6E4403]">Layar Dapur 🧑‍🍳</h1>
        <p class="text-gray-500 font-medium">Daftar pesanan lunas yang harus dimasak</p>
      </div>
      <div class="flex gap-4">
        <router-link to="/kasir" class="bg-gray-100 hover:bg-gray-200 text-gray-700 px-6 py-3 rounded-2xl font-bold transition-colors">
          Lihat Kasir
        </router-link>
        <div class="bg-red-100 text-red-700 px-6 py-3 rounded-2xl font-bold tracking-widest border border-red-200 flex items-center gap-2">
          <span class="w-3 h-3 bg-red-500 rounded-full animate-pulse"></span>
          LIVE KITCHEN
        </div>
      </div>
    </div>

    <div v-if="pesananDapur.length === 0" class="bg-white rounded-3xl border border-[#EBC898]/30 p-16 text-center shadow-sm">
      <div class="text-6xl mb-4">✨</div>
      <h2 class="text-2xl font-bold text-[#6E4403]">Dapur Bersih!</h2>
      <p class="text-gray-500 mt-2">Belum ada pesanan baru yang masuk.</p>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div v-for="pesanan in pesananDapur" :key="pesanan.id" 
           class="bg-white rounded-3xl shadow-md border-t-8 flex flex-col overflow-hidden transition-all hover:shadow-xl"
           :class="pesanan.status === 'Sedang Dimasak' ? 'border-orange-500' : 'border-[#9C6307]'">
        
        <div class="p-5 border-b border-[#EBC898]/30 bg-[#FDF8F3]/50 flex justify-between items-start">
          <div>
            <h3 class="text-2xl font-black text-[#6E4403]">{{ pesanan.id }}</h3>
            <p class="text-sm text-gray-500 font-bold mt-1">🕒 Masuk: {{ pesanan.waktu }}</p>
          </div>
          <span v-if="pesanan.status === 'Sedang Dimasak'" class="bg-orange-100 text-orange-700 text-xs font-bold px-3 py-1 rounded-full border border-orange-200">
            🔥 Dimasak
          </span>
          <span v-else class="bg-blue-100 text-blue-700 text-xs font-bold px-3 py-1 rounded-full border border-blue-200">
            Baru
          </span>
        </div>

        <div class="p-5 flex-grow bg-white">
          <ul class="space-y-4">
            <li v-for="item in pesanan.items" :key="item.name" class="flex items-start gap-3">
              <div class="bg-[#FDF8F3] border border-[#EBC898]/50 text-[#9C6307] w-10 h-10 rounded-xl flex items-center justify-center font-black text-xl flex-shrink-0">
                {{ item.qty }}
              </div>
              <div>
                <p class="font-bold text-lg text-gray-800 leading-tight">{{ item.name }}</p>
                <p v-if="item.notes" class="text-sm text-red-500 font-medium italic mt-1">Catatan: {{ item.notes }}</p>
              </div>
            </li>
          </ul>
        </div>

        <div class="p-5 bg-gray-50 border-t border-[#EBC898]/20 flex gap-3">
          <button v-if="pesanan.status === 'Antrean Baru'" @click="mulaiMasak(pesanan)" 
                  class="flex-1 bg-white border-2 border-[#9C6307] text-[#9C6307] hover:bg-orange-50 font-bold py-3 rounded-xl transition-colors">
            Mulai Masak
          </button>
          
          <button @click="pesananSelesai(pesanan.id)" 
                  class="flex-1 bg-[#9C6307] hover:bg-[#6E4403] text-white font-bold py-3 rounded-xl transition-colors shadow-md">
            Selesai & Panggil
          </button>
        </div>

      </div>
    </div>

  </div>
</template>