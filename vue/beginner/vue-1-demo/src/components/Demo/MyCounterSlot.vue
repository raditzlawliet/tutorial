<script setup>
import { ref } from 'vue'

const count = ref(0)
</script>

<template>
  <slot name="header">Count</slot>: {{ count }}
  <button @click="count++">
    <slot>+1</slot>
  </button>
</template>

<style lang="css" scoped>
button {
  background: #000;
  color: #fff;
}
</style>
