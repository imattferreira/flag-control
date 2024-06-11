<script setup lang="ts">
const props = withDefaults(defineProps<{ bold: boolean }>(), {
  bold: false
})
</script>

<template>
  <th
    v-if="props.bold"
    scope="row"
    class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white"
  >
    <slot />
  </th>
  <td v-else class="px-6 py-4"><slot /></td>
</template>
