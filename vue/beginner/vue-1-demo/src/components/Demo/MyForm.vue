<script setup>
import { ref } from 'vue'

const text = ref('Edit me')
const checked = ref(true)
const selected = ref('A')

const buttonText = ref('Press me!')
function changeButtonText() {
  buttonText.value = 'Ouch!'
}

const showed = ref(false)
function showValue() {
  showed.value = !showed.value
}
</script>

<template>
  <h2>Text Input</h2>
  <input v-model="text" />
  <p>{{ text }}</p>

  <h2>Checkbox</h2>
  <input type="checkbox" id="checkbox" v-model="checked" />
  <label for="checkbox">Checked: {{ checked }}</label>

  <h2>Select</h2>
  <select v-model="selected">
    <option disabled value="">Please select one</option>
    <option>A</option>
    <option>B</option>
    <option>C</option>
  </select>
  <p>Selected: {{ selected }}</p>

  <button @click="changeButtonText">{{ buttonText }}</button>

  <p></p>
  <button @click="showValue">Press Me Again</button>
  <div v-show="showed">
    <p>{{ text }}</p>
    <p>Checked: {{ checked }}</p>
    <p>Selected: {{ selected }}</p>
  </div>
</template>
