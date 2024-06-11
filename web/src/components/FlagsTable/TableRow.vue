<script setup lang="ts">
import type { Flag } from '@/entities/Flag'
import TableRowItem from './TableRowItem.vue'

const props = defineProps<{ flag: Flag }>()
</script>

<template>
  <tr class="bg-white dark:bg-gray-800">
    <TableRowItem>{{ props.flag.id }}</TableRowItem>
    <TableRowItem bold>{{ props.flag.name }}</TableRowItem>
    <TableRowItem>{{ props.flag.active ? 'Active' : 'Disabled' }}</TableRowItem>
  </tr>
</template>
