<script setup>
defineProps({
  show: Boolean,
  total: Number
})

defineEmits(['close', 'pilih-metode'])

const formatRp = (angka) => {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(angka)
}
</script>

<template>
  <div v-if="show" class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm z-50 flex items-center justify-center animate-fade-in">
    <div class="bg-white rounded-3xl shadow-2xl w-full max-w-lg overflow-hidden transform transition-all scale-100 border-2 border-[#EBC898]">
      <div class="bg-[#FDF8F3] p-6 border-b border-[#EBC898]/50 flex justify-between items-center">
        <h3 class="text-2xl font-black text-[#6E4403]">Pilih Pembayaran</h3>
        <button @click="$emit('close')" class="text-gray-400 hover:text-red-500 bg-white hover:bg-red-50 rounded-full p-2 transition-colors border border-gray-200">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
        </button>
      </div>

      <div class="p-8 text-center">
        <p class="text-gray-500 font-medium mb-2">Total Tagihan Anda:</p>
        <p class="text-4xl font-black text-[#9C6307] mb-8">{{ formatRp(total) }}</p>

        <div class="grid grid-cols-2 gap-4">
          <button @click="$emit('pilih-metode', 'qris')" class="border-2 border-blue-100 bg-blue-50 hover:bg-blue-500 hover:border-blue-500 group rounded-2xl p-6 transition-all flex flex-col items-center gap-3">
            <div class="text-4xl">📱</div>
            <span class="font-bold text-blue-600 group-hover:text-white">QRIS / E-Wallet</span>
          </button>

          <button @click="$emit('pilih-metode', 'cash')" class="border-2 border-emerald-100 bg-emerald-50 hover:bg-emerald-500 hover:border-emerald-500 group rounded-2xl p-6 transition-all flex flex-col items-center gap-3">
            <div class="text-4xl">💵</div>
            <span class="font-bold text-emerald-600 group-hover:text-white">Bayar di Kasir</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.animate-fade-in { animation: fadeIn 0.3s ease-out; }
@keyframes fadeIn { from { opacity: 0; } to { opacity: 1; } }
</style>