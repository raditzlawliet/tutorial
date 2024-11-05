<script setup>
import { useCounterStore } from '@/stores/counter.js'

const store = useCounterStore()
</script>

<template>
  <div>MyCounterPinia</div>
  <div>Count: {{ store.count }}</div>
  <div>Double Count: {{ store.doubleCount }}</div>
  <button @click="store.increment">+1</button>
</template>
