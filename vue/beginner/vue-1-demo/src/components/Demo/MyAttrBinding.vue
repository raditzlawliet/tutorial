<!--
Here we are reactively binding element attributes / properties to the state.
The :title syntax is short for v-bind:title.
-->

<script setup>
import { ref } from 'vue'

const isRed = ref(true)
const color = ref('green')

function toggleRed() {
  isRed.value = !isRed.value
}

function toggleColor() {
  color.value = color.value === 'green' ? 'blue' : 'green'
}
</script>

<template>
  <!--
  class bindings have special support for objects and arrays
  in addition to plain strings
  -->
  <p :class="{ red: isRed }" @click="toggleRed">
    This should be red... but click me to toggle it.
  </p>

  <!-- style bindings also support object and arrays -->
  <p :style="{ color }" @click="toggleColor">
    This should be green, and should toggle between green and blue on click.
  </p>
</template>

<style scoped>
.red {
  color: red;
}
</style>
