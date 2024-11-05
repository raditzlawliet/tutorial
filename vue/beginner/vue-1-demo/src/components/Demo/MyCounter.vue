<script setup>
import { ref } from 'vue'

const count = ref(0)
</script>

<template>
  Count: {{ count }}
  <button @click="count++">+1</button>
</template>
