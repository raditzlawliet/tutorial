<script setup>
import { mdiExcavator } from '@mdi/js'
import { ref } from 'vue'

const showDialog = ref(false)
</script>

<template>
  <main>
    <v-container fluid class="fill-height">
      <v-row>
        <v-col cols="12" class="d-flex justify-center">
          <v-btn :prepend-icon="mdiExcavator" @click="showDialog = true">Click me!</v-btn>
        </v-col>
      </v-row>
    </v-container>

    <v-dialog max-width="500" v-model="showDialog">
      <v-card title="Congratulation">
        <v-card-text> Now you just need to explore it more deeply. </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn text="Ok" @click="showDialog = false"></v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </main>
</template>
