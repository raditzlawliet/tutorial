<script setup>
import {
  mdiStar,
  mdiCakeVariant,
  mdiCheckboxMarkedCircle,
  mdiAccountCircle,
  mdiDelete,
} from '@mdi/js'
</script>
<template>
  <div class="text-center">
    <v-chip class="ma-2" color="indigo" :prepend-icon="mdiAccountCircle"> Mike </v-chip>

    <v-chip :append-icon="mdiStar" class="ma-2" color="orange"> Premium </v-chip>

    <v-chip :append-icon="mdiCakeVariant" class="ma-2" color="primary"> 1 Year </v-chip>

    <v-chip class="ma-2" color="green">
      <template v-slot:prepend>
        <v-avatar class="green-darken-4"> 1 </v-avatar>
      </template>
      Years
    </v-chip>

    <v-chip
      :model-value="true"
      class="ma-2"
      color="teal"
      :prepend-icon="mdiCheckboxMarkedCircle"
      closable
    >
      Confirmed
    </v-chip>

    <v-chip
      :model-value="true"
      class="ma-2"
      :close-icon="mdiDelete"
      color="teal"
      :prepend-icon="mdiCheckboxMarkedCircle"
      closable
    >
      Confirmed
    </v-chip>
  </div>
</template>
