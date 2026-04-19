<!-- MODUL 2: Pembayaran / Layar Kasir -->

<script setup>
import { ref } from 'vue'

const formatRp = (angka) => {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(angka)
}

// INI DATA DIGANTI FAR MASIH DUMMY
const antreanPesanan = ref([
  {
    id: 'TRX-001',
    waktu: '10:30',
    metode: 'CASH',
    status: 'Menunggu Pembayaran',
    total: 35000,
    items: [
      { name: 'Nasi Goreng Ayam', qty: 1 },
      { name: 'Es Teh Manis', qty: 1 }
    ]
  },
  {
    id: 'TRX-002',
    waktu: '10:32',
    metode: 'QRIS',
    status: 'Lunas - Diproses Dapur',
    total: 18000,
    items: [
      { name: 'Ayam Geprek', qty: 1 }
    ]
  },
  {
    id: 'TRX-003',
    waktu: '10:35',
    metode: 'CASH',
    status: 'Menunggu Pembayaran',
    total: 42000,
    items: [
      { name: 'Soto Ayam', qty: 2 },
      { name: 'Tahu Goreng', qty: 2 }
    ]
  }
])


const konfirmasiLunas = (idPesanan) => {
  const pesanan = antreanPesanan.value.find(p => p.id === idPesanan)
  if (pesanan) {
    pesanan.status = 'Lunas - Diproses Dapur'
    alert(`Transaksi ${idPesanan} berhasil dikonfirmasi Lunas! 💰`)
  }
}
</script>

<template>
  <div class="min-h-screen bg-[#FDF8F3] font-sans text-gray-800 p-8">
    
    <div class="flex justify-between items-center mb-8 bg-white p-6 rounded-3xl shadow-sm border border-[#EBC898]/30">
      <div>
        <h1 class="text-3xl font-black text-[#6E4403]">Dashboard Kasir 💻</h1>
        <p class="text-gray-500 font-medium">Kelola pembayaran dan antrean pesanan</p>
      </div>
      <div class="flex gap-4">
        <router-link to="/" class="bg-gray-100 hover:bg-gray-200 text-gray-700 px-6 py-3 rounded-2xl font-bold transition-colors">
          Buka Layar Kiosk
        </router-link>
        <div class="bg-emerald-100 text-emerald-700 px-6 py-3 rounded-2xl font-bold tracking-widest border border-emerald-200">
          KASIR AKTIF
        </div>
      </div>
    </div>

    <div class="bg-white rounded-3xl shadow-sm border border-[#EBC898]/30 overflow-hidden">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-[#FDF8F3] border-b border-[#EBC898]/30">
            <th class="p-5 font-bold text-[#6E4403]">ID Transaksi</th>
            <th class="p-5 font-bold text-[#6E4403]">Waktu & Detail Pesanan</th>
            <th class="p-5 font-bold text-[#6E4403]">Metode</th>
            <th class="p-5 font-bold text-[#6E4403]">Total Tagihan</th>
            <th class="p-5 font-bold text-[#6E4403]">Status</th>
            <th class="p-5 font-bold text-center text-[#6E4403]">Aksi Kasir</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-[#EBC898]/20">
          
          <tr v-for="pesanan in antreanPesanan" :key="pesanan.id" class="hover:bg-gray-50 transition-colors">
            
            <td class="p-5 font-bold text-gray-800">{{ pesanan.id }}</td>
            
            <td class="p-5">
              <span class="text-xs text-gray-400 font-bold block mb-1">Pukul {{ pesanan.waktu }}</span>
              <ul class="text-sm font-medium text-gray-600 space-y-1">
                <li v-for="item in pesanan.items" :key="item.name">
                  • {{ item.name }} <span class="text-[#9C6307] font-bold">x{{ item.qty }}</span>
                </li>
              </ul>
            </td>

            <td class="p-5">
              <span v-if="pesanan.metode === 'QRIS'" class="bg-blue-100 text-blue-700 py-1 px-3 rounded-lg text-xs font-bold border border-blue-200">📱 QRIS</span>
              <span v-else class="bg-emerald-100 text-emerald-700 py-1 px-3 rounded-lg text-xs font-bold border border-emerald-200">💵 CASH</span>
            </td>

            <td class="p-5 font-black text-[#9C6307] text-lg">{{ formatRp(pesanan.total) }}</td>

            <td class="p-5">
              <span v-if="pesanan.status.includes('Menunggu')" class="flex items-center gap-2 text-orange-600 font-bold text-sm bg-orange-50 py-2 px-3 rounded-xl w-fit">
                <span class="w-2 h-2 rounded-full bg-orange-500 animate-pulse"></span> {{ pesanan.status }}
              </span>
              <span v-else class="flex items-center gap-2 text-emerald-600 font-bold text-sm bg-emerald-50 py-2 px-3 rounded-xl w-fit">
                <span class="w-2 h-2 rounded-full bg-emerald-500"></span> {{ pesanan.status }}
              </span>
            </td>

            <td class="p-5 text-center">
              <button 
                v-if="pesanan.status.includes('Menunggu')" 
                @click="konfirmasiLunas(pesanan.id)"
                class="bg-[#9C6307] hover:bg-[#6E4403] text-white font-bold py-2 px-4 rounded-xl text-sm transition-colors shadow-md w-full">
                Terima Uang & Lunas
              </button>
              
              <span v-else class="text-gray-400 font-bold text-sm italic">
                Selesai ✓
              </span>
            </td>

          </tr>
        </tbody>
      </table>
    </div>

  </div>
</template>