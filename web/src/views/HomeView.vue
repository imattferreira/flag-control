<script setup lang="ts">
import api from '@/api'
import FlagsTable from '@/components/FlagsTable/FlagsTable.vue'

const flags = await api.flags.getAll()
</script>

<template>
  <main class="max-w-screen-md">
    <FlagsTable :flags="flags" />
  </main>
</template>
