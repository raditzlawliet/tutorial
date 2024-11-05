import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useCounterLocalStorageStore = defineStore('counterLocalStorage', () => {
  const count = ref(localStorage.getItem('counterLocalStorage') || 0)
  const doubleCount = computed(() => count.value * 2)
  function increment() {
    count.value++
    localStorage.setItem('counterLocalStorage', count.value)
  }

  return { count, doubleCount, increment }
})
