<script setup>
import { ref } from 'vue'
import MyTodoItem from './MyTodoItem.vue'

const groceryList = ref([
  { id: 0, text: 'Vegetables' },
  { id: 1, text: 'Cheese' },
  { id: 2, text: 'Whatever else humans are supposed to eat' },
])
</script>

<template>
  <ol>
    <!--
      We are providing each todo-item with the todo object
      it's representing, so that its content can be dynamic.
      We also need to provide each component with a "key",
      which is explained in the guide section on v-for.
    -->
    <MyTodoItem
      v-for="item in groceryList"
      :todo="item"
      :key="item.id"
    ></MyTodoItem>
  </ol>
</template>
