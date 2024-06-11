<script setup lang="ts">
import type { Flag } from '@/entities/Flag'
import TableRow from './TableRow.vue'

const props = defineProps<{ flags: Flag[] }>()
</script>

<template>
  <div class="relative overflow-x-auto">
    <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
      <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
        <tr>
          <th scope="col" class="px-6 py-3">ID</th>
          <th scope="col" class="px-6 py-3">Flag</th>
          <th scope="col" class="px-6 py-3">Active</th>
          <th scope="col" class="px-6 py-3">Options</th>
        </tr>
      </thead>
      <tbody>
        <TableRow v-for="flag in props.flags" :key="flag.id" :flag="flag" />
      </tbody>
    </table>
  </div>
</template>
