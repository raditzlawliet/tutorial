<script setup>
import { useCounterLocalStorageStore } from '@/stores/counterLocalStorage'

const store = useCounterLocalStorageStore()
</script>

<template>
  <div>MyCounterPinia</div>
  <div>Count: {{ store.count }}</div>
  <div>Double Count: {{ store.doubleCount }}</div>
  <button @click="store.increment">+1</button>
</template>
