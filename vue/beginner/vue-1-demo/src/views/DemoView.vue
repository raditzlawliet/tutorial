<script setup>
import MyCounter from '../components/Demo/MyCounter.vue'
import MyCounterSlot from '../components/Demo/MyCounterSlot.vue'
import MyHelloWorld from '../components/Demo/MyHelloWorld.vue'
import MyTodoList from '../components/Demo/MyTodoList.vue'
import MyTodoListPlain from '../components/Demo/MyTodoListPlain.vue'
import MyAttrBinding from '../components/Demo/MyAttrBinding.vue'
import MyForm from '../components/Demo/MyForm.vue'
import MyCounterPinia from '@/components/Demo/MyCounterPinia.vue'
import MyCounterPiniaLocalStorage from '@/components/Demo/MyCounterPiniaLocalStorage.vue'
</script>

<template>
  <main>
    <hr />
    <MyCounter />
    <MyCounter />

    <hr />
    <MyCounterSlot>
      <template v-slot:header>
        <!-- content for the header slot -->
        Angka
      </template>
      Tambah 1
    </MyCounterSlot>

    <hr />
    <MyCounterSlot>
      <template v-slot:header>
        <!-- content for the header slot -->
        Angka
      </template>
      <template #default> Tambah 1 </template>
    </MyCounterSlot>

    <hr />
    <MyHelloWorld />

    <hr />
    <MyTodoList />

    <hr />
    <MyTodoListPlain />

    <hr />
    <MyAttrBinding />

    <hr />
    <MyForm />

    <hr />
    <MyCounterPinia />
    <MyCounterPinia />

    <hr />
    <MyCounterPiniaLocalStorage />
    <MyCounterPiniaLocalStorage />
  </main>
</template>
